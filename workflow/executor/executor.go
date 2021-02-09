package executor

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"k8s.io/utils/pointer"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"strings"
	"time"

	common2 "github.com/argoproj/argo/v3/workflow/artifacts/common"

	argofile "github.com/argoproj/pkg/file"
	log "github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"

	"github.com/argoproj/argo/v3/errors"
	wfv1 "github.com/argoproj/argo/v3/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo/v3/util"
	"github.com/argoproj/argo/v3/util/archive"
	errorsutil "github.com/argoproj/argo/v3/util/errors"
	"github.com/argoproj/argo/v3/util/retry"
	waitutil "github.com/argoproj/argo/v3/util/wait"
	artifact "github.com/argoproj/argo/v3/workflow/artifacts"
	"github.com/argoproj/argo/v3/workflow/common"
	os_specific "github.com/argoproj/argo/v3/workflow/executor/os-specific"
)

// ExecutorRetry is a retry backoff settings for WorkflowExecutor
// Run	Seconds
// 0	0.000
// 1	1.000
// 2	2.600
// 3	5.160
// 4	9.256
var ExecutorRetry = wait.Backoff{
	Steps:    5,
	Duration: 1 * time.Second,
	Factor:   1.6,
	Jitter:   0.5,
}

const (
	// This directory temporarily stores the tarballs of the artifacts before uploading
	tempOutArtDir = "/tmp/argo/outputs/artifacts"
)

// WorkflowExecutor is program which runs as the init/wait container
type WorkflowExecutor struct {
	PodName            string
	Template           wfv1.Template
	ClientSet          kubernetes.Interface
	Namespace          string
	PodAnnotationsPath string
	ExecutionControl   *common.ExecutionControl
	RuntimeExecutor    ContainerRuntimeExecutor

	// memoized container ID to prevent multiple lookups
	mainContainerID string
	// memoized configmaps
	memoizedConfigMaps map[string]string
	// memoized secrets
	memoizedSecrets map[string][]byte
	// list of errors that occurred during execution.
	// the first of these is used as the overall message of the node
	errors []error
}

//go:generate mockery -name ContainerRuntimeExecutor

// ContainerRuntimeExecutor is the interface for interacting with a container runtime (e.g. docker)
type ContainerRuntimeExecutor interface {
	// GetFileContents returns the file contents of a file in a container as a string
	GetFileContents(containerID string, sourcePath string) (string, error)

	// CopyFile copies a source file in a container to a local path
	CopyFile(containerID string, sourcePath string, destPath string, compressionLevel int) error

	// GetOutputStream returns the entirety of the container output as a io.Reader
	// Used to capture script results as an output parameter, and to archive container logs
	GetOutputStream(ctx context.Context, containerID string, combinedOutput bool) (io.ReadCloser, error)

	// GetExitCode returns the exit code of the container
	// Used to capture script exit code as an output parameter
	GetExitCode(ctx context.Context, containerID string) (string, error)

	// WaitInit is called before Wait() to signal the executor about an impending Wait call.
	// For most executors this is a noop, and is only used by the the PNS executor
	WaitInit() error

	// Wait waits for the container to complete
	Wait(ctx context.Context, containerID string) error

	// Kill a list of containerIDs first with a SIGTERM then with a SIGKILL after a grace period
	Kill(ctx context.Context, containerIDs []string) error
}

// NewExecutor instantiates a new workflow executor
func NewExecutor(clientset kubernetes.Interface, podName, namespace, podAnnotationsPath string, cre ContainerRuntimeExecutor, template wfv1.Template) WorkflowExecutor {
	return WorkflowExecutor{
		PodName:            podName,
		ClientSet:          clientset,
		Namespace:          namespace,
		PodAnnotationsPath: podAnnotationsPath,
		RuntimeExecutor:    cre,
		Template:           template,
		memoizedConfigMaps: map[string]string{},
		memoizedSecrets:    map[string][]byte{},
		errors:             []error{},
	}
}

// HandleError is a helper to annotate the pod with the error message upon a unexpected executor panic or error
func (we *WorkflowExecutor) HandleError(ctx context.Context) {
	if r := recover(); r != nil {
		_ = we.AddAnnotation(ctx, common.AnnotationKeyNodeMessage, fmt.Sprintf("%v", r))
		util.WriteTeriminateMessage(fmt.Sprintf("%v", r))
		log.Fatalf("executor panic: %+v\n%s", r, debug.Stack())
	} else {
		if len(we.errors) > 0 {
			util.WriteTeriminateMessage(we.errors[0].Error())
			_ = we.AddAnnotation(ctx, common.AnnotationKeyNodeMessage, we.errors[0].Error())
		}
	}
}

// LoadArtifacts loads artifacts from location to a container path
func (we *WorkflowExecutor) LoadArtifacts(ctx context.Context) error {
	log.Infof("Start loading input artifacts...")

	for _, art := range we.Template.Inputs.Artifacts {

		log.Infof("Downloading artifact: %s", art.Name)

		if !art.HasLocationOrKey() {
			if art.Optional {
				log.Warnf("Ignoring optional artifact '%s' which was not supplied", art.Name)
				continue
			} else {
				return errors.Errorf("required artifact %s not supplied", art.Name)
			}
		}
		driverArt, err := we.newDriverArt(&art)
		if err != nil {
			return err
		}
		artDriver, err := we.InitDriver(ctx, driverArt)
		if err != nil {
			return err
		}
		// Determine the file path of where to load the artifact
		if art.Path == "" {
			return errors.InternalErrorf("Artifact %s did not specify a path", art.Name)
		}
		var artPath string
		mnt := common.FindOverlappingVolume(&we.Template, art.Path)
		if mnt == nil {
			artPath = path.Join(common.ExecutorArtifactBaseDir, art.Name)
		} else {
			// If we get here, it means the input artifact path overlaps with an user specified
			// volumeMount in the container. Because we also implement input artifacts as volume
			// mounts, we need to load the artifact into the user specified volume mount,
			// as opposed to the `input-artifacts` volume that is an implementation detail
			// unbeknownst to the user.
			log.Infof("Specified artifact path %s overlaps with volume mount at %s. Extracting to volume mount", art.Path, mnt.MountPath)
			artPath = path.Join(common.ExecutorMainFilesystemDir, art.Path)
		}

		// The artifact is downloaded to a temporary location, after which we determine if
		// the file is a tarball or not. If it is, it is first extracted then renamed to
		// the desired location. If not, it is simply renamed to the location.
		tempArtPath := artPath + ".tmp"
		err = artDriver.Load(driverArt, tempArtPath)
		if err != nil {
			if art.Optional && errors.IsCode(errors.CodeNotFound, err) {
				log.Infof("Skipping optional input artifact that was not found: %s", art.Name)
				continue
			}
			return err
		}

		isTar := false
		isZip := false
		if art.GetArchive().None != nil {
			// explicitly not a tar
			isTar = false
			isZip = false
		} else if art.GetArchive().Tar != nil {
			// explicitly a tar
			isTar = true
		} else if art.GetArchive().Zip != nil {
			// explicitly a zip
			isZip = true
		} else {
			// auto-detect if tarball
			// (don't try to autodetect zip files for backwards compatibility)
			isTar, err = isTarball(tempArtPath)
			if err != nil {
				return err
			}
		}

		if isTar {
			err = untar(tempArtPath, artPath)
			_ = os.Remove(tempArtPath)
		} else if isZip {
			err = unzip(tempArtPath, artPath)
			_ = os.Remove(tempArtPath)
		} else {
			err = os.Rename(tempArtPath, artPath)
		}
		if err != nil {
			return err
		}

		log.Infof("Successfully download file: %s", artPath)
		if art.Mode != nil {
			err = chmod(artPath, *art.Mode, art.RecurseMode)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (we *WorkflowExecutor) ProcessData(ctx context.Context) error {
	dataTemplate := we.Template.Data
	if dataTemplate == nil {
		return nil
	}

	// Once we allow input parameters to the data template, we'll load them here
	var data interface{}
	var err error
	for _, step := range dataTemplate {
		switch {
		case step.WithArtifactPaths != nil:
			data, err = we.processDataArtifacts(ctx, step.WithArtifactPaths)
		case step.Filter != nil:
			data, err = we.processFilter(data, step.Filter)
		case step.Aggregator != nil:
			data, err = we.processAggregator(data, step.Aggregator)
		}

		if err != nil {
			return fmt.Errorf("error processing data step '%s': %w", step.Name, err)
		}
	}

	return we.processOutput(ctx, data)
}

func (we *WorkflowExecutor) processDataArtifacts(ctx context.Context, artifacts *wfv1.WithArtifactPaths) ([]string, error) {
	driverArt, err := we.newDriverArt(&artifacts.Artifact)
	if err != nil {
		return nil, err
	}
	artDriver, err := we.InitDriver(ctx, driverArt)
	if err != nil {
		return nil, err
	}

	var files []string
	files, err = artDriver.ListObjects(&artifacts.Artifact)
	if err != nil {
		return nil, err
	}

	return files, nil

	//


}

func (we *WorkflowExecutor) processFilter(data interface{}, filter *wfv1.Filter) ([]string, error) {
	var files []string
	var ok bool
	if files, ok = data.([]string); !ok {
		// Currently we only support input being []string, but we could easily also do so for [][]string
		return nil, fmt.Errorf("intput is not []string")
	}

	switch fil := filter; {
	case fil.Directory != nil:
		// If recursive is set to false, remove all files that contain a directory
		if fil.Directory.Recursive != nil && !*fil.Directory.Recursive {
			inPlaceFilter(func(file string) bool {
				return !strings.Contains(file, "/")
			}, &files)
		}

		if fil.Directory.Regex != "" {
			re, err := regexp.Compile(fil.Directory.Regex)
			if err != nil {
				return nil, fmt.Errorf("regex '%s' is not valid: %w", fil.Directory.Regex, err)
			}
			inPlaceFilter(func(file string) bool {
				return re.MatchString(file)
			}, &files)
		}
	}

	return files, nil
}


func (we *WorkflowExecutor) processAggregator(data interface{}, aggregator *wfv1.Aggregator) ([][]string, error) {
	var files []string
	var ok bool
	if files, ok = data.([]string); !ok {
		return nil, fmt.Errorf("intput is not []string")
	}

	var aggFiles [][]string
	switch {
	case aggregator.Batch != 0:
		// Starts at -1 because we increment before first file
		filesSeen := -1
		aggFiles = groupBy(func(file string) string {
			filesSeen++
			return fmt.Sprint(filesSeen / aggregator.Batch)
		}, files)
	case aggregator.Regex != "":
		re, err := regexp.Compile(aggregator.Regex)
		if err != nil {
			return nil, fmt.Errorf("regex '%s' is not valid: %w", aggregator.Regex, err)
		}
		aggFiles = groupBy(func(file string) string {
			match := re.FindStringSubmatch(file)
			if len(match) == 1 {
				return match[0]
			}
			return match[1]
		}, files)
	}

	return aggFiles, nil

}

func (we *WorkflowExecutor) processOutput(ctx context.Context, data interface{}) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	we.Template.Outputs.Result = pointer.StringPtr(string(out))
	err = we.AnnotateOutputs(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

// StageFiles will create any files required by script/resource templates
func (we *WorkflowExecutor) StageFiles() error {
	var filePath string
	var body []byte
	switch we.Template.GetType() {
	case wfv1.TemplateTypeScript:
		log.Infof("Loading script source to %s", common.ExecutorScriptSourcePath)
		filePath = common.ExecutorScriptSourcePath
		body = []byte(we.Template.Script.Source)
	case wfv1.TemplateTypeResource:
		log.Infof("Loading manifest to %s", common.ExecutorResourceManifestPath)
		filePath = common.ExecutorResourceManifestPath
		body = []byte(we.Template.Resource.Manifest)
	default:
		return nil
	}
	err := ioutil.WriteFile(filePath, body, 0644)
	if err != nil {
		return errors.InternalWrapError(err)
	}
	return nil
}

// SaveArtifacts uploads artifacts to the archive location
func (we *WorkflowExecutor) SaveArtifacts(ctx context.Context) error {
	if len(we.Template.Outputs.Artifacts) == 0 {
		log.Infof("No output artifacts")
		return nil
	}
	log.Infof("Saving output artifacts")
	mainCtrID, err := we.GetMainContainerID(ctx)
	if err != nil {
		return err
	}

	err = os.MkdirAll(tempOutArtDir, os.ModePerm)
	if err != nil {
		return errors.InternalWrapError(err)
	}

	for i, art := range we.Template.Outputs.Artifacts {
		err := we.saveArtifact(ctx, mainCtrID, &art)
		if err != nil {
			return err
		}
		we.Template.Outputs.Artifacts[i] = art
	}
	return nil
}

func (we *WorkflowExecutor) saveArtifact(ctx context.Context, mainCtrID string, art *wfv1.Artifact) error {
	// Determine the file path of where to find the artifact
	if art.Path == "" {
		return errors.InternalErrorf("Artifact %s did not specify a path", art.Name)
	}
	fileName, localArtPath, err := we.stageArchiveFile(mainCtrID, art)
	if err != nil {
		if art.Optional && errors.IsCode(errors.CodeNotFound, err) {
			log.Warnf("Ignoring optional artifact '%s' which does not exist in path '%s': %v", art.Name, art.Path, err)
			return nil
		}
		return err
	}
	return we.saveArtifactFromFile(ctx, art, fileName, localArtPath)
}

// fileBase is probably path.Base(filePath), but can be something else
func (we *WorkflowExecutor) saveArtifactFromFile(ctx context.Context, art *wfv1.Artifact, fileName, localArtPath string) error {
	if !art.HasKey() {
		key, err := we.Template.ArchiveLocation.GetKey()
		if err != nil {
			return err
		}
		if err = art.SetType(we.Template.ArchiveLocation.Get()); err != nil {
			return err
		}
		if err := art.SetKey(path.Join(key, fileName)); err != nil {
			return err
		}
	}
	driverArt, err := we.newDriverArt(art)
	if err != nil {
		return err
	}
	artDriver, err := we.InitDriver(ctx, driverArt)
	if err != nil {
		return err
	}
	err = artDriver.Save(localArtPath, driverArt)
	if err != nil {
		return err
	}
	we.maybeDeleteLocalArtPath(localArtPath)
	log.Infof("Successfully saved file: %s", localArtPath)
	return nil
}

func (we *WorkflowExecutor) maybeDeleteLocalArtPath(localArtPath string) {
	if os.Getenv("REMOVE_LOCAL_ART_PATH") == "true" {
		log.WithField("localArtPath", localArtPath).Info("deleting local artifact")
		// remove is best effort (the container will go away anyways).
		// we just want reduce peak space usage
		err := os.Remove(localArtPath)
		if err != nil {
			log.Warnf("Failed to remove %s: %v", localArtPath, err)
		}
	} else {
		log.WithField("localArtPath", localArtPath).Info("not deleting local artifact")
	}
}

// stageArchiveFile stages a path in a container for archiving from the wait sidecar.
// Returns a filename and a local path for the upload.
// The filename is incorporated into the final path when uploading it to the artifact repo.
// The local path is the final staging location of the file (or directory) which we will pass
// to the SaveArtifacts call and may be a directory or file.
func (we *WorkflowExecutor) stageArchiveFile(mainCtrID string, art *wfv1.Artifact) (string, string, error) {
	log.Infof("Staging artifact: %s", art.Name)
	strategy := art.Archive
	if strategy == nil {
		// If no strategy is specified, default to the tar strategy
		strategy = &wfv1.ArchiveStrategy{
			Tar: &wfv1.TarStrategy{},
		}
	}
	compressionLevel := gzip.NoCompression
	if strategy.Tar != nil {
		if l := strategy.Tar.CompressionLevel; l != nil {
			compressionLevel = int(*l)
		} else {
			compressionLevel = gzip.DefaultCompression
		}
	}

	if !we.isBaseImagePath(art.Path) {
		// If we get here, we are uploading an artifact from a mirrored volume mount which the wait
		// sidecar has direct access to. We can upload directly from the shared volume mount,
		// instead of copying it from the container.
		mountedArtPath := filepath.Join(common.ExecutorMainFilesystemDir, art.Path)
		log.Infof("Staging %s from mirrored volume mount %s", art.Path, mountedArtPath)
		if strategy.None != nil {
			fileName := filepath.Base(art.Path)
			log.Infof("No compression strategy needed. Staging skipped")
			if !argofile.Exists(mountedArtPath) {
				return "", "", errors.Errorf(errors.CodeNotFound, "%s no such file or directory", art.Path)
			}
			return fileName, mountedArtPath, nil
		}
		fileName := fmt.Sprintf("%s.tgz", art.Name)
		localArtPath := filepath.Join(tempOutArtDir, fileName)
		f, err := os.Create(localArtPath)
		if err != nil {
			return "", "", errors.InternalWrapError(err)
		}
		w := bufio.NewWriter(f)
		err = archive.TarGzToWriter(mountedArtPath, compressionLevel, w)
		if err != nil {
			return "", "", err
		}
		log.Infof("Successfully staged %s from mirrored volume mount %s", art.Path, mountedArtPath)
		return fileName, localArtPath, nil
	}

	fileName := fmt.Sprintf("%s.tgz", art.Name)
	localArtPath := filepath.Join(tempOutArtDir, fileName)
	log.Infof("Copying %s from container base image layer to %s", art.Path, localArtPath)

	err := we.RuntimeExecutor.CopyFile(mainCtrID, art.Path, localArtPath, compressionLevel)
	if err != nil {
		return "", "", err
	}
	if strategy.Tar != nil {
		// NOTE we already tar gzip the file in the executor. So this is a noop.
		return fileName, localArtPath, nil
	}
	// localArtPath now points to a .tgz file, and the archive strategy is *not* tar. We need to untar it
	log.Infof("Untaring %s archive before upload", localArtPath)
	unarchivedArtPath := path.Join(filepath.Dir(localArtPath), art.Name)
	err = untar(localArtPath, unarchivedArtPath)
	if err != nil {
		return "", "", err
	}
	// Delete the tarball
	err = os.Remove(localArtPath)
	if err != nil {
		return "", "", errors.InternalWrapError(err)
	}
	isDir, err := argofile.IsDirectory(unarchivedArtPath)
	if err != nil {
		return "", "", errors.InternalWrapError(err)
	}
	fileName = filepath.Base(art.Path)
	if isDir {
		localArtPath = unarchivedArtPath
	} else {
		// If we are uploading a single file, we need to preserve original filename so that
		// 1. minio client can infer its mime-type, based on file extension
		// 2. the original filename is incorporated into the final path
		localArtPath = path.Join(tempOutArtDir, fileName)
		err = os.Rename(unarchivedArtPath, localArtPath)
		if err != nil {
			return "", "", errors.InternalWrapError(err)
		}
	}
	// In the future, if we were to support other compression formats (e.g. bzip2) or options
	// the logic would go here, and compression would be moved out of the executors
	return fileName, localArtPath, nil
}

// isBaseImagePath checks if the given artifact path resides in the base image layer of the container
// versus a shared volume mount between the wait and main container
func (we *WorkflowExecutor) isBaseImagePath(path string) bool {
	// first check if path overlaps with a user-specified volumeMount
	if common.FindOverlappingVolume(&we.Template, path) != nil {
		return false
	}
	// next check if path overlaps with a shared input-artifact emptyDir mounted by argo
	for _, inArt := range we.Template.Inputs.Artifacts {
		if path == inArt.Path {
			// The input artifact may have been optional and not supplied. If this is the case, the file won't exist on
			// the input artifact volume. Since this function was called, we know that we want to use this path as an
			// output artifact, so we should look for it in the base image path.
			if inArt.Optional && !inArt.HasLocationOrKey() {
				return true
			}
			return false
		}
		if strings.HasPrefix(path, inArt.Path+"/") {
			return false
		}
	}
	return true
}

// SaveParameters will save the content in the specified file path as output parameter value
func (we *WorkflowExecutor) SaveParameters(ctx context.Context) error {
	if len(we.Template.Outputs.Parameters) == 0 {
		log.Infof("No output parameters")
		return nil
	}
	log.Infof("Saving output parameters")
	mainCtrID, err := we.GetMainContainerID(ctx)
	if err != nil {
		return err
	}

	for i, param := range we.Template.Outputs.Parameters {
		log.Infof("Saving path output parameter: %s", param.Name)
		// Determine the file path of where to find the parameter
		if param.ValueFrom == nil || param.ValueFrom.Path == "" {
			continue
		}

		var output *wfv1.AnyString
		if we.isBaseImagePath(param.ValueFrom.Path) {
			executorType := os.Getenv(common.EnvVarContainerRuntimeExecutor)
			if executorType == common.ContainerRuntimeExecutorK8sAPI || executorType == common.ContainerRuntimeExecutorKubelet {
				log.Infof("Copying output parameter %s from base image layer %s is not supported for k8sapi and kubelet executors. "+
					"Consider using an emptyDir volume: https://argoproj.github.io/argo/empty-dir/.", param.Name, param.ValueFrom.Path)
				continue
			}
			log.Infof("Copying %s from base image layer", param.ValueFrom.Path)
			fileContents, err := we.RuntimeExecutor.GetFileContents(mainCtrID, param.ValueFrom.Path)
			if err != nil {
				// We have a default value to use instead of returning an error
				if param.ValueFrom.Default != nil {
					output = param.ValueFrom.Default
				} else {
					return err
				}
			} else {
				output = wfv1.AnyStringPtr(fileContents)
			}
		} else {
			log.Infof("Copying %s from volume mount", param.ValueFrom.Path)
			mountedPath := filepath.Join(common.ExecutorMainFilesystemDir, param.ValueFrom.Path)
			data, err := ioutil.ReadFile(mountedPath)
			if err != nil {
				// We have a default value to use instead of returning an error
				if param.ValueFrom.Default != nil {
					output = param.ValueFrom.Default
				} else {
					return err
				}
			} else {
				output = wfv1.AnyStringPtr(string(data))
			}
		}

		// Trims off a single newline for user convenience
		output = wfv1.AnyStringPtr(strings.TrimSuffix(output.String(), "\n"))
		we.Template.Outputs.Parameters[i].Value = output
		log.Infof("Successfully saved output parameter: %s", param.Name)
	}
	return nil
}

// SaveLogs saves logs
func (we *WorkflowExecutor) SaveLogs(ctx context.Context) (*wfv1.Artifact, error) {
	if !we.Template.ArchiveLocation.IsArchiveLogs() {
		return nil, nil
	}
	log.Infof("Saving logs")
	mainCtrID, err := we.GetMainContainerID(ctx)
	if err != nil {
		return nil, err
	}
	tempLogsDir := "/tmp/argo/outputs/logs"
	err = os.MkdirAll(tempLogsDir, os.ModePerm)
	if err != nil {
		return nil, errors.InternalWrapError(err)
	}
	fileName := "main.log"
	mainLog := path.Join(tempLogsDir, fileName)
	err = we.saveLogToFile(ctx, mainCtrID, mainLog)
	if err != nil {
		return nil, err
	}
	art := &wfv1.Artifact{Name: "main-logs"}
	err = we.saveArtifactFromFile(ctx, art, fileName, mainLog)
	if err != nil {
		return nil, err
	}
	return art, nil
}

// GetSecret will retrieve the Secrets from VolumeMount
func (we *WorkflowExecutor) GetSecret(ctx context.Context, accessKeyName string, accessKey string) (string, error) {
	file, err := ioutil.ReadFile(filepath.Join(common.SecretVolMountPath, accessKeyName, accessKey))
	if err != nil {
		return "", err
	}
	return string(file), nil
}

// saveLogToFile saves the entire log output of a container to a local file
func (we *WorkflowExecutor) saveLogToFile(ctx context.Context, mainCtrID, path string) error {
	outFile, err := os.Create(path)
	if err != nil {
		return errors.InternalWrapError(err)
	}
	defer func() { _ = outFile.Close() }()
	reader, err := we.RuntimeExecutor.GetOutputStream(ctx, mainCtrID, true)
	if err != nil {
		return err
	}
	defer func() { _ = reader.Close() }()
	_, err = io.Copy(outFile, reader)
	if err != nil {
		return errors.InternalWrapError(err)
	}
	return nil
}

func (we *WorkflowExecutor) newDriverArt(art *wfv1.Artifact) (*wfv1.Artifact, error) {
	driverArt := art.DeepCopy()
	err := driverArt.Relocate(we.Template.ArchiveLocation)
	return driverArt, err
}

// InitDriver initializes an instance of an artifact driver
func (we *WorkflowExecutor) InitDriver(ctx context.Context, art *wfv1.Artifact) (common2.ArtifactDriver, error) {
	driver, err := artifact.NewDriver(ctx, art, we)
	if err == artifact.ErrUnsupportedDriver {
		return nil, errors.Errorf(errors.CodeBadRequest, "Unsupported artifact driver for %s", art.Name)
	}
	return driver, err
}

// getPod is a wrapper around the pod interface to get the current pod from kube API server
func (we *WorkflowExecutor) getPod(ctx context.Context) (*apiv1.Pod, error) {
	podsIf := we.ClientSet.CoreV1().Pods(we.Namespace)
	var pod *apiv1.Pod
	err := waitutil.Backoff(ExecutorRetry, func() (bool, error) {
		var err error
		pod, err = podsIf.Get(ctx, we.PodName, metav1.GetOptions{})
		return !errorsutil.IsTransientErr(err), err
	})
	if err != nil {
		return nil, errors.InternalWrapError(err)
	}
	return pod, nil
}

// GetConfigMapKey retrieves a configmap value and memoizes the result
func (we *WorkflowExecutor) GetConfigMapKey(ctx context.Context, name, key string) (string, error) {
	namespace := we.Namespace
	cachedKey := fmt.Sprintf("%s/%s/%s", namespace, name, key)
	if val, ok := we.memoizedConfigMaps[cachedKey]; ok {
		return val, nil
	}
	configmapsIf := we.ClientSet.CoreV1().ConfigMaps(namespace)
	var configmap *apiv1.ConfigMap
	err := waitutil.Backoff(retry.DefaultRetry, func() (bool, error) {
		var err error
		configmap, err = configmapsIf.Get(ctx, name, metav1.GetOptions{})
		return !errorsutil.IsTransientErr(err), err
	})
	if err != nil {
		return "", errors.InternalWrapError(err)
	}
	// memoize all keys in the configmap since it's highly likely we will need to get a
	// subsequent key in the configmap (e.g. username + password) and we can save an API call
	for k, v := range configmap.Data {
		we.memoizedConfigMaps[fmt.Sprintf("%s/%s/%s", namespace, name, k)] = v
	}
	val, ok := we.memoizedConfigMaps[cachedKey]
	if !ok {
		return "", errors.Errorf(errors.CodeBadRequest, "configmap '%s' does not have the key '%s'", name, key)
	}
	return val, nil
}

// GetSecrets retrieves a secret value and memoizes the result
func (we *WorkflowExecutor) GetSecrets(ctx context.Context, namespace, name, key string) ([]byte, error) {
	cachedKey := fmt.Sprintf("%s/%s/%s", namespace, name, key)
	if val, ok := we.memoizedSecrets[cachedKey]; ok {
		return val, nil
	}
	secretsIf := we.ClientSet.CoreV1().Secrets(namespace)
	var secret *apiv1.Secret
	err := waitutil.Backoff(retry.DefaultRetry, func() (bool, error) {
		var err error
		secret, err = secretsIf.Get(ctx, name, metav1.GetOptions{})
		return !errorsutil.IsTransientErr(err), err
	})
	if err != nil {
		return []byte{}, errors.InternalWrapError(err)
	}
	// memoize all keys in the secret since it's highly likely we will need to get a
	// subsequent key in the secret (e.g. username + password) and we can save an API call
	for k, v := range secret.Data {
		we.memoizedSecrets[fmt.Sprintf("%s/%s/%s", namespace, name, k)] = v
	}
	val, ok := we.memoizedSecrets[cachedKey]
	if !ok {
		return []byte{}, errors.Errorf(errors.CodeBadRequest, "secret '%s' does not have the key '%s'", name, key)
	}
	return val, nil
}

// GetMainContainerStatus returns the container status of the main container, nil if the main container does not exist
func (we *WorkflowExecutor) GetMainContainerStatus(ctx context.Context) (*apiv1.ContainerStatus, error) {
	pod, err := we.getPod(ctx)
	if err != nil {
		return nil, err
	}
	for _, ctrStatus := range pod.Status.ContainerStatuses {
		if ctrStatus.Name == common.MainContainerName {
			return &ctrStatus, nil
		}
	}
	return nil, nil
}

// GetMainContainerID returns the container id of the main container
func (we *WorkflowExecutor) GetMainContainerID(ctx context.Context) (string, error) {
	if we.mainContainerID != "" {
		return we.mainContainerID, nil
	}
	ctrStatus, err := we.GetMainContainerStatus(ctx)
	if err != nil {
		return "", err
	}
	if ctrStatus == nil {
		return "", nil
	}
	we.mainContainerID = containerID(ctrStatus.ContainerID)
	return we.mainContainerID, nil
}

// CaptureScriptResult will add the stdout of a script template as output result
func (we *WorkflowExecutor) CaptureScriptResult(ctx context.Context) error {

	if we.ExecutionControl == nil || !we.ExecutionControl.IncludeScriptOutput {
		log.Infof("No Script output reference in workflow. Capturing script output ignored")
		return nil
	}
	if we.Template.Script == nil && we.Template.Container == nil {
		log.Infof("Template type is neither of Script or Container. Capturing script output ignored")
		return nil
	}
	log.Infof("Capturing script output")
	mainContainerID, err := we.GetMainContainerID(ctx)
	if err != nil {
		return err
	}
	reader, err := we.RuntimeExecutor.GetOutputStream(ctx, mainContainerID, false)
	if err != nil {
		return err
	}
	defer func() { _ = reader.Close() }()
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.InternalWrapError(err)
	}
	out := string(bytes)
	// Trims off a single newline for user convenience
	outputLen := len(out)
	if outputLen > 0 && out[outputLen-1] == '\n' {
		out = out[0 : outputLen-1]
	}

	const maxAnnotationSize int = 256 * (1 << 10) // 256 kB
	// A character in a string is a byte
	if len(out) > maxAnnotationSize {
		log.Warnf("Output is larger than the maximum allowed size of 256 kB, only the last 256 kB were saved")
		out = out[len(out)-maxAnnotationSize:]
	}

	we.Template.Outputs.Result = &out
	return nil
}

// CaptureScriptExitCode will add the exit code of a script template as output exit code
func (we *WorkflowExecutor) CaptureScriptExitCode(ctx context.Context) error {
	if we.Template.Script == nil && we.Template.Container == nil {
		log.Infof("Template type is neither of Script or Container. Capturing exit code ignored")
		return nil
	}
	log.Infof("Capturing script exit code")
	mainContainerID, err := we.GetMainContainerID(ctx)
	if err != nil {
		return err
	}
	exitCode, err := we.RuntimeExecutor.GetExitCode(ctx, mainContainerID)
	if err != nil {
		return err
	}

	if exitCode != "" {
		we.Template.Outputs.ExitCode = &exitCode
	}
	return nil
}

// AnnotateOutputs annotation to the pod indicating all the outputs.
func (we *WorkflowExecutor) AnnotateOutputs(ctx context.Context, logArt *wfv1.Artifact) error {
	outputs := we.Template.Outputs.DeepCopy()
	if logArt != nil {
		outputs.Artifacts = append(outputs.Artifacts, *logArt)
	}

	if !outputs.HasOutputs() {
		return nil
	}
	log.Infof("Annotating pod with output")
	outputBytes, err := json.Marshal(outputs)
	if err != nil {
		return errors.InternalWrapError(err)
	}
	log.Infof("Annotation: %s", outputBytes)
	return we.AddAnnotation(ctx, common.AnnotationKeyOutputs, string(outputBytes))
}

// AddError adds an error to the list of encountered errors durign execution
func (we *WorkflowExecutor) AddError(err error) {
	log.Errorf("executor error: %+v", err)
	we.errors = append(we.errors, err)
}

// AddAnnotation adds an annotation to the workflow pod
func (we *WorkflowExecutor) AddAnnotation(ctx context.Context, key, value string) error {
	return common.AddPodAnnotation(ctx, we.ClientSet, we.PodName, we.Namespace, key, value, ExecutorRetry)
}

// isTarball returns whether or not the file is a tarball
func isTarball(filePath string) (bool, error) {
	log.Infof("Detecting if %s is a tarball", filePath)
	f, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer f.Close()
	gzr, err := gzip.NewReader(f)
	if err != nil {
		return false, nil
	}
	defer gzr.Close()
	tarr := tar.NewReader(gzr)
	_, err = tarr.Next()
	return err == nil, nil
}

// untar extracts a tarball to a temporary directory,
// renaming it to the desired location
func untar(tarPath string, destPath string) error {
	decompressor := func(src string, dest string) error {
		_, err := common.RunCommand("tar", "-xf", src, "-C", dest)
		return err
	}

	return unpack(tarPath, destPath, decompressor)
}

// unzip extracts a zip folder to a temporary directory,
// renaming it to the desired location
func unzip(zipPath string, destPath string) error {
	decompressor := func(src string, dest string) error {
		r, err := zip.OpenReader(src)
		if err != nil {
			return err
		}
		defer func() {
			if err := r.Close(); err != nil {
				panic(err)
			}
		}()

		// Closure to address file descriptors issue with all the deferred .Close() methods
		extractAndWriteFile := func(f *zip.File) error {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer func() {
				if err := rc.Close(); err != nil {
					panic(err)
				}
			}()

			path := filepath.Join(dest, f.Name)
			if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
				return fmt.Errorf("%s: Illegal file path", path)
			}

			if f.FileInfo().IsDir() {
				if err = os.MkdirAll(path, f.Mode()); err != nil {
					return err
				}
			} else {
				if err = os.MkdirAll(filepath.Dir(path), f.Mode()); err != nil {
					return err
				}
				f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err != nil {
					return err
				}
				defer func() {
					if err := f.Close(); err != nil {
						panic(err)
					}
				}()

				_, err = io.Copy(f, rc)
				if err != nil {
					return err
				}
			}
			return nil
		}

		for _, f := range r.File {
			if err := extractAndWriteFile(f); err != nil {
				return err
			}
			log.Infof("Extracting file: %s", f.Name)
		}

		log.Infof("Extraction of %s finished!", src)

		return nil
	}

	return unpack(zipPath, destPath, decompressor)
}

// unpack unpacks a compressed file (tarball or zip file) to a temporary directory,
// renaming it to the desired location
// decompression is done using the decompressor closure, that should decompress a tarball or zip file
func unpack(srcPath string, destPath string, decompressor func(string, string) error) error {
	// first extract the tar into a temporary dir
	tmpDir := destPath + ".tmpdir"
	err := os.MkdirAll(tmpDir, os.ModePerm)
	if err != nil {
		return errors.InternalWrapError(err)
	}
	if decompressor != nil {
		if err = decompressor(srcPath, tmpDir); err != nil {
			return err
		}
	}
	// next, decide how we wish to rename the file/dir
	// to the destination path.
	files, err := ioutil.ReadDir(tmpDir)
	if err != nil {
		return errors.InternalWrapError(err)
	}
	if len(files) == 1 {
		// if the tar is comprised of single file or directory,
		// rename that file to the desired location
		filePath := path.Join(tmpDir, files[0].Name())
		err = os.Rename(filePath, destPath)
		if err != nil {
			return errors.InternalWrapError(err)
		}
		err = os.Remove(tmpDir)
		if err != nil {
			return errors.InternalWrapError(err)
		}
	} else {
		// the tar extracted into multiple files. In this case,
		// just rename the temp directory to the dest path
		err = os.Rename(tmpDir, destPath)
		if err != nil {
			return errors.InternalWrapError(err)
		}
	}
	return nil
}

func chmod(artPath string, mode int32, recurse bool) error {
	err := os.Chmod(artPath, os.FileMode(mode))
	if err != nil {
		return errors.InternalWrapError(err)
	}

	if recurse {
		err = filepath.Walk(artPath, func(path string, f os.FileInfo, err error) error {
			return os.Chmod(path, os.FileMode(mode))
		})
		if err != nil {
			return errors.InternalWrapError(err)
		}
	}

	return nil
}

// containerID is a convenience function to strip the 'docker://', 'containerd://' from k8s ContainerID string
func containerID(ctrID string) string {
	schemeIndex := strings.Index(ctrID, "://")
	if schemeIndex == -1 {
		return ctrID
	}
	return ctrID[schemeIndex+3:]
}

// Wait is the sidecar container logic which waits for the main container to complete.
// Also monitors for updates in the pod annotations which may change (e.g. terminate)
// Upon completion, kills any sidecars after it finishes.
func (we *WorkflowExecutor) Wait(ctx context.Context) error {
	err := we.RuntimeExecutor.WaitInit()
	if err != nil {
		return err
	}
	log.Infof("Waiting on main container")
	mainContainerID, err := we.waitMainContainerStart(ctx)
	if err != nil {
		return err
	}
	log.Infof("main container started with container ID: %s", mainContainerID)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	annotationUpdatesCh := we.monitorAnnotations(ctx)
	go we.monitorDeadline(ctx, annotationUpdatesCh)

	err = waitutil.Backoff(ExecutorRetry, func() (bool, error) {
		err := we.RuntimeExecutor.Wait(ctx, mainContainerID)
		return err == nil, err
	})
	if err != nil {
		return err
	}
	log.Infof("Main container completed")
	return nil
}

// waitMainContainerStart waits for the main container to start and returns its container ID.
func (we *WorkflowExecutor) waitMainContainerStart(ctx context.Context) (string, error) {
	for {
		podsIf := we.ClientSet.CoreV1().Pods(we.Namespace)
		fieldSelector := fields.ParseSelectorOrDie(fmt.Sprintf("metadata.name=%s", we.PodName))
		opts := metav1.ListOptions{
			FieldSelector: fieldSelector.String(),
		}

		var watchIf watch.Interface

		err := waitutil.Backoff(ExecutorRetry, func() (bool, error) {
			var err error
			watchIf, err = podsIf.Watch(ctx, opts)
			return !errorsutil.IsTransientErr(err), err
		})
		if err != nil {
			return "", errors.InternalWrapErrorf(err, "Failed to establish pod watch: %v", err)
		}
		for watchEv := range watchIf.ResultChan() {
			if watchEv.Type == watch.Error {
				return "", errors.InternalErrorf("Pod watch error waiting for main to start: %v", watchEv.Object)
			}
			pod, ok := watchEv.Object.(*apiv1.Pod)
			if !ok {
				log.Warnf("Pod watch returned non pod object: %v", watchEv.Object)
				continue
			}
			for _, ctrStatus := range pod.Status.ContainerStatuses {
				if ctrStatus.Name == common.MainContainerName {
					log.Debug(ctrStatus)
					if ctrStatus.State.Waiting != nil {
						// main container is still in waiting status
					} else if ctrStatus.State.Waiting == nil && ctrStatus.State.Running == nil && ctrStatus.State.Terminated == nil {
						// status still not ready, wait
					} else if ctrStatus.ContainerID != "" {
						we.mainContainerID = containerID(ctrStatus.ContainerID)
						return containerID(ctrStatus.ContainerID), nil
					} else {
						// main container in running or terminated state but missing container ID
						return "", errors.InternalError("Main container ID cannot be found")
					}
				}
			}
		}
		log.Warnf("Pod watch closed unexpectedly")
	}
}

func watchFileChanges(ctx context.Context, pollInterval time.Duration, filePath string) <-chan struct{} {
	res := make(chan struct{})
	go func() {
		defer close(res)

		var modTime *time.Time
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			file, err := os.Stat(filePath)
			if err != nil {
				log.Fatal(err)
			}
			newModTime := file.ModTime()
			if modTime != nil && !modTime.Equal(file.ModTime()) {
				res <- struct{}{}
			}
			modTime = &newModTime
			time.Sleep(pollInterval)
		}
	}()
	return res
}

// monitorAnnotations starts a goroutine which monitors for any changes to the pod annotations.
// Emits an event on the returned channel upon any updates
func (we *WorkflowExecutor) monitorAnnotations(ctx context.Context) <-chan struct{} {
	log.Infof("Starting annotations monitor")

	// Create a channel to listen for a SIGUSR2. Upon receiving of the signal, we force reload our annotations
	// directly from kubernetes API. The controller uses this to fast-track notification of annotations
	// instead of waiting for the volume file to get updated (which can take minutes)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os_specific.GetOsSignal())

	we.setExecutionControl(ctx)

	// Create a channel which will notify a listener on new updates to the annotations
	annotationUpdateCh := make(chan struct{})

	annotationChanges := watchFileChanges(ctx, 10*time.Second, we.PodAnnotationsPath)
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Infof("Annotations monitor stopped")
				signal.Stop(sigs)
				close(sigs)
				close(annotationUpdateCh)
				return
			case <-sigs:
				log.Infof("Received update signal. Reloading annotations from API")
				annotationUpdateCh <- struct{}{}
				we.setExecutionControl(ctx)
			case <-annotationChanges:
				log.Infof("%s updated", we.PodAnnotationsPath)
				err := we.LoadExecutionControl()
				if err != nil {
					log.Warnf("Failed to reload execution control from annotations: %v", err)
					continue
				}
				if we.ExecutionControl != nil {
					log.Infof("Execution control reloaded from annotations: %v", *we.ExecutionControl)
				}
				annotationUpdateCh <- struct{}{}
			}
		}
	}()
	return annotationUpdateCh
}

// setExecutionControl sets the execution control information from the pod annotation
func (we *WorkflowExecutor) setExecutionControl(ctx context.Context) {
	pod, err := we.getPod(ctx)
	if err != nil {
		log.Warnf("Failed to set execution control from API server: %v", err)
		return
	}
	execCtlString, ok := pod.ObjectMeta.Annotations[common.AnnotationKeyExecutionControl]
	if !ok {
		we.ExecutionControl = nil
	} else {
		var execCtl common.ExecutionControl
		err = json.Unmarshal([]byte(execCtlString), &execCtl)
		if err != nil {
			log.Errorf("Error unmarshalling '%s': %v", execCtlString, err)
			return
		}
		we.ExecutionControl = &execCtl
		log.Infof("Execution control set from API: %v", *we.ExecutionControl)
	}
}

// monitorDeadline checks to see if we exceeded the deadline for the step and
// terminates the main container if we did
func (we *WorkflowExecutor) monitorDeadline(ctx context.Context, annotationsUpdate <-chan struct{}) {
	log.Infof("Starting deadline monitor")
	for {
		select {
		case <-ctx.Done():
			log.Info("Deadline monitor stopped")
			return
		case <-annotationsUpdate:
		default:
			// TODO(jessesuen): we do not effectively use the annotations update channel yet. Ideally, we
			// should optimize this logic so that we use some type of mutable timer against the deadline
			// value instead of polling.
			if we.ExecutionControl != nil && we.ExecutionControl.Deadline != nil {
				if time.Now().UTC().After(*we.ExecutionControl.Deadline) {
					var message string
					// Zero value of the deadline indicates an intentional cancel vs. a timeout. We treat
					// timeouts as a failure and the pod should be annotated with that error
					if we.ExecutionControl.Deadline.IsZero() {
						message = "terminated"
					} else {
						message = "Step exceeded its deadline"
					}
					log.Info(message)
					_ = we.AddAnnotation(ctx, common.AnnotationKeyNodeMessage, message)
					log.Infof("Killing main container")
					mainContainerID, _ := we.GetMainContainerID(ctx)
					err := we.RuntimeExecutor.Kill(ctx, []string{mainContainerID})
					if err != nil {
						log.Warnf("Failed to kill main container: %v", err)
					}
					return
				}
			}
			time.Sleep(1 * time.Second)
		}
	}
}

// KillSidecars kills any sidecars to the main container
func (we *WorkflowExecutor) KillSidecars(ctx context.Context) error {
	log.Infof("Killing sidecars")
	pod, err := we.getPod(ctx)
	if err != nil {
		return err
	}
	sidecarIDs := make([]string, 0)
	for _, ctrStatus := range pod.Status.ContainerStatuses {
		if ctrStatus.Name == common.MainContainerName || ctrStatus.Name == common.WaitContainerName {
			continue
		}
		if ctrStatus.State.Terminated != nil {
			continue
		}
		containerID := containerID(ctrStatus.ContainerID)
		log.Infof("Killing sidecar %s (%s)", ctrStatus.Name, containerID)
		sidecarIDs = append(sidecarIDs, containerID)
	}
	if len(sidecarIDs) == 0 {
		return nil
	}
	return we.RuntimeExecutor.Kill(ctx, sidecarIDs)
}

// LoadExecutionControl reads the execution control definition from the the Kubernetes downward api annotations volume file
func (we *WorkflowExecutor) LoadExecutionControl() error {
	err := unmarshalAnnotationField(we.PodAnnotationsPath, common.AnnotationKeyExecutionControl, &we.ExecutionControl)
	if err != nil {
		if errors.IsCode(errors.CodeNotFound, err) {
			return nil
		}
		return err
	}
	return nil
}

// LoadTemplate reads the template definition from the the Kubernetes downward api annotations volume file
func LoadTemplate(path string) (*wfv1.Template, error) {
	var tmpl wfv1.Template
	err := unmarshalAnnotationField(path, common.AnnotationKeyTemplate, &tmpl)
	if err != nil {
		return nil, err
	}
	return &tmpl, nil
}

// unmarshalAnnotationField unmarshals the value of an annotation key into the supplied interface
// from the downward api annotation volume file
func unmarshalAnnotationField(filePath string, key string, into interface{}) error {
	// Read the annotation file
	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("ERROR opening annotation file from %s", filePath)
		return errors.InternalWrapError(err)
	}

	defer func() {
		_ = file.Close()
	}()
	reader := bufio.NewReader(file)

	// Prefix of key property in the annotation file
	prefix := fmt.Sprintf("%s=", key)

	for {
		// Read line-by-line
		var buffer bytes.Buffer
		var l []byte
		var isPrefix bool
		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)
			// If we've reached the end of the line, stop reading.
			if !isPrefix {
				break
			}
			// If we're just at the EOF, break
			if err != nil {
				break
			}
		}

		line := buffer.String()

		// Read property
		if strings.HasPrefix(line, prefix) {
			// Trim the prefix
			content := strings.TrimPrefix(line, prefix)

			// This part is a bit tricky in terms of unmarshalling
			// The content in the file will be something like,
			// `"{\"type\":\"container\",\"inputs\":{},\"outputs\":{}}"`
			// which is required to unmarshal twice

			// First unmarshal to a string without escaping characters
			var fieldString string
			err = json.Unmarshal([]byte(content), &fieldString)
			if err != nil {
				log.Errorf("Error unmarshalling annotation into string, %s, %v\n", content, err)
				return errors.InternalWrapError(err)
			}

			// Second unmarshal to a template
			err = json.Unmarshal([]byte(fieldString), into)
			if err != nil {
				log.Errorf("Error unmarshalling annotation into datastructure, %s, %v\n", fieldString, err)
				return errors.InternalWrapError(err)
			}
			return nil
		}

		// The end of the annotation file
		if err == io.EOF {
			break
		}
	}

	if err != io.EOF {
		return errors.InternalWrapError(err)
	}

	// If we reach here, then the key does not exist in the file
	return errors.Errorf(errors.CodeNotFound, "Key %s not found in annotation file: %s", key, filePath)
}
