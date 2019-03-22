package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/argoproj/argo/errors"
	"github.com/argoproj/argo/pkg/apis/workflow"
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/ghodss/yaml"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasttemplate"
	apiv1 "k8s.io/api/core/v1"
	apierr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

// FindOverlappingVolume looks an artifact path, checks if it overlaps with any
// user specified volumeMounts in the template, and returns the deepest volumeMount
// (if any).
func FindOverlappingVolume(tmpl *wfv1.Template, path string) *apiv1.VolumeMount {
	if tmpl.Container == nil {
		return nil
	}
	var volMnt *apiv1.VolumeMount
	deepestLen := 0
	for _, mnt := range tmpl.Container.VolumeMounts {
		if !strings.HasPrefix(path, mnt.MountPath) {
			continue
		}
		if len(mnt.MountPath) > deepestLen {
			volMnt = &mnt
			deepestLen = len(mnt.MountPath)
		}
	}
	return volMnt
}

// KillPodContainer is a convenience function to issue a kill signal to a container in a pod
// It gives a 15 second grace period before issuing SIGKILL
// NOTE: this only works with containers that have sh
func KillPodContainer(restConfig *rest.Config, namespace string, pod string, container string) error {
	exec, err := ExecPodContainer(restConfig, namespace, pod, container, true, true, "sh", "-c", "kill 1; sleep 15; kill -9 1")
	if err != nil {
		return err
	}
	// Stream will initiate the command. We do want to wait for the result so we launch as a goroutine
	go func() {
		_, _, err := GetExecutorOutput(exec)
		if err != nil {
			log.Warnf("Kill command failed (expected to fail with 137): %v", err)
			return
		}
		log.Infof("Kill of %s (%s) successfully issued", pod, container)
	}()
	return nil
}

// ExecPodContainer runs a command in a container in a pod and returns the remotecommand.Executor
func ExecPodContainer(restConfig *rest.Config, namespace string, pod string, container string, stdout bool, stderr bool, command ...string) (remotecommand.Executor, error) {
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, errors.InternalWrapError(err)
	}

	execRequest := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(pod).
		Namespace(namespace).
		SubResource("exec").
		Param("container", container).
		Param("stdout", fmt.Sprintf("%v", stdout)).
		Param("stderr", fmt.Sprintf("%v", stderr)).
		Param("tty", "false")

	for _, cmd := range command {
		execRequest = execRequest.Param("command", cmd)
	}

	log.Info(execRequest.URL())
	exec, err := remotecommand.NewSPDYExecutor(restConfig, "POST", execRequest.URL())
	if err != nil {
		return nil, errors.InternalWrapError(err)
	}
	return exec, nil
}

// GetExecutorOutput returns the output of an remotecommand.Executor
func GetExecutorOutput(exec remotecommand.Executor) (*bytes.Buffer, *bytes.Buffer, error) {
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	err := exec.Stream(remotecommand.StreamOptions{
		Stdout: &stdOut,
		Stderr: &stdErr,
		Tty:    false,
	})
	if err != nil {
		return nil, nil, errors.InternalWrapError(err)
	}
	return &stdOut, &stdErr, nil
}

// ProcessArgs sets in the inputs, the values either passed via arguments, or the hardwired values
// It substitutes:
// * parameters in the template from the arguments
// * global parameters (e.g. {{workflow.parameters.XX}}, {{workflow.name}}, {{workflow.status}})
// * local parameters (e.g. {{pod.name}})
func ProcessArgs(tmpl *wfv1.Template, args wfv1.Arguments, globalParams, localParams map[string]string, validateOnly bool) (*wfv1.Template, error) {
	// For each input parameter:
	// 1) check if was supplied as argument. if so use the supplied value from arg
	// 2) if not, use default value.
	// 3) if no default value, it is an error
	tmpl = tmpl.DeepCopy()
	for i, inParam := range tmpl.Inputs.Parameters {
		if inParam.Default != nil {
			// first set to default value
			inParam.Value = inParam.Default
		}
		// overwrite value from argument (if supplied)
		argParam := args.GetParameterByName(inParam.Name)
		if argParam != nil && argParam.Value != nil {
			newValue := *argParam.Value
			inParam.Value = &newValue
		}
		if inParam.Value == nil {
			return nil, errors.Errorf(errors.CodeBadRequest, "inputs.parameters.%s was not supplied", inParam.Name)
		}
		tmpl.Inputs.Parameters[i] = inParam
	}

	// Performs substitutions of input artifacts
	newInputArtifacts := make([]wfv1.Artifact, len(tmpl.Inputs.Artifacts))
	for i, inArt := range tmpl.Inputs.Artifacts {
		// if artifact has hard-wired location, we prefer that
		if inArt.HasLocation() {
			newInputArtifacts[i] = inArt
			continue
		}
		// artifact must be supplied
		argArt := args.GetArtifactByName(inArt.Name)
		if !inArt.Optional && argArt == nil {
			return nil, errors.Errorf(errors.CodeBadRequest, "inputs.artifacts.%s was not supplied", inArt.Name)
		}
		if !inArt.Optional && !argArt.HasLocation() && !validateOnly {
			return nil, errors.Errorf(errors.CodeBadRequest, "inputs.artifacts.%s missing location information", inArt.Name)
		}
		argArt.Path = inArt.Path
		argArt.Mode = inArt.Mode
		newInputArtifacts[i] = *argArt
	}
	tmpl.Inputs.Artifacts = newInputArtifacts

	return substituteParams(tmpl, globalParams, localParams)
}

// substituteParams returns a new copy of the template with global, pod, and input parameters substituted
func substituteParams(tmpl *wfv1.Template, globalParams, localParams map[string]string) (*wfv1.Template, error) {
	tmplBytes, err := json.Marshal(tmpl)
	if err != nil {
		return nil, errors.InternalWrapError(err)
	}
	// First replace globals & locals, then replace inputs because globals could be referenced in the inputs
	replaceMap := make(map[string]string)
	for k, v := range globalParams {
		replaceMap[k] = v
	}
	for k, v := range localParams {
		replaceMap[k] = v
	}
	fstTmpl := fasttemplate.New(string(tmplBytes), "{{", "}}")
	globalReplacedTmplStr, err := Replace(fstTmpl, replaceMap, true)
	if err != nil {
		return nil, err
	}
	var globalReplacedTmpl wfv1.Template
	err = json.Unmarshal([]byte(globalReplacedTmplStr), &globalReplacedTmpl)
	if err != nil {
		return nil, errors.InternalWrapError(err)
	}
	// Now replace the rest of substitutions (the ones that can be made) in the template
	replaceMap = make(map[string]string)
	for _, inParam := range globalReplacedTmpl.Inputs.Parameters {
		if inParam.Value == nil {
			return nil, errors.InternalErrorf("inputs.parameters.%s had no value", inParam.Name)
		}
		replaceMap["inputs.parameters."+inParam.Name] = *inParam.Value
	}
	fstTmpl = fasttemplate.New(globalReplacedTmplStr, "{{", "}}")
	s, err := Replace(fstTmpl, replaceMap, true)
	if err != nil {
		return nil, err
	}
	var newTmpl wfv1.Template
	err = json.Unmarshal([]byte(s), &newTmpl)
	if err != nil {
		return nil, errors.InternalWrapError(err)
	}
	return &newTmpl, nil
}

// Replace executes basic string substitution of a template with replacement values.
// allowUnresolved indicates whether or not it is acceptable to have unresolved variables
// remaining in the substituted template. prefixFilter will apply the replacements only
// to variables with the specified prefix
func Replace(fstTmpl *fasttemplate.Template, replaceMap map[string]string, allowUnresolved bool) (string, error) {
	var unresolvedErr error
	replacedTmpl := fstTmpl.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
		replacement, ok := replaceMap[tag]
		if !ok {
			if allowUnresolved {
				// just write the same string back
				return w.Write([]byte(fmt.Sprintf("{{%s}}", tag)))
			}
			unresolvedErr = errors.Errorf(errors.CodeBadRequest, "failed to resolve {{%s}}", tag)
			return 0, nil
		}
		// The following escapes any special characters (e.g. newlines, tabs, etc...)
		// in preparation for substitution
		replacement = strconv.Quote(replacement)
		replacement = replacement[1 : len(replacement)-1]
		return w.Write([]byte(replacement))
	})
	if unresolvedErr != nil {
		return "", unresolvedErr
	}
	return replacedTmpl, nil
}

// RunCommand is a convenience function to run/log a command and log the stderr upon failure
func RunCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmdStr := strings.Join(cmd.Args, " ")
	log.Info(cmdStr)
	cmdOutput, err := cmd.Output()
	if err != nil {
		exErr := err.(*exec.ExitError)
		errOutput := string(exErr.Stderr)
		log.Errorf("`%s` failed: %s", cmdStr, errOutput)
		return errors.InternalError(strings.TrimSpace(errOutput))
	}
	if strings.Contains(string(cmdOutput), "Failed") {
		return errors.InternalError(string(cmdOutput))

	}
	return nil
}

const patchRetries = 5

// AddPodAnnotation adds an annotation to pod
func AddPodAnnotation(c kubernetes.Interface, podName, namespace, key, value string) error {
	return addPodMetadata(c, "annotations", podName, namespace, key, value)
}

// AddPodLabel adds an label to pod
func AddPodLabel(c kubernetes.Interface, podName, namespace, key, value string) error {
	return addPodMetadata(c, "labels", podName, namespace, key, value)
}

// addPodMetadata is helper to either add a pod label or annotation to the pod
func addPodMetadata(c kubernetes.Interface, field, podName, namespace, key, value string) error {
	metadata := map[string]interface{}{
		"metadata": map[string]interface{}{
			field: map[string]string{
				key: value,
			},
		},
	}
	var err error
	patch, err := json.Marshal(metadata)
	if err != nil {
		return errors.InternalWrapError(err)
	}
	for attempt := 0; attempt < patchRetries; attempt++ {
		_, err = c.CoreV1().Pods(namespace).Patch(podName, types.MergePatchType, patch)
		if err != nil {
			if !apierr.IsConflict(err) {
				return err
			}
		} else {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	return err
}

// IsPodTemplate returns whether the template corresponds to a pod
func IsPodTemplate(tmpl *wfv1.Template) bool {
	if tmpl.Container != nil || tmpl.Script != nil || tmpl.Resource != nil {
		return true
	}
	return false
}

// GetTaskAncestry returns a list of taskNames which are ancestors of this task
func GetTaskAncestry(taskName string, tasks []wfv1.DAGTask) []string {
	taskByName := make(map[string]wfv1.DAGTask)
	for _, task := range tasks {
		taskByName[task.Name] = task
	}

	visited := make(map[string]bool)
	var getAncestry func(s string)
	getAncestry = func(currTask string) {
		task := taskByName[currTask]
		for _, depTask := range task.Dependencies {
			getAncestry(depTask)
		}
		if currTask != taskName {
			visited[currTask] = true
		}
	}

	getAncestry(taskName)
	ancestry := make([]string, 0)
	for ancestor := range visited {
		ancestry = append(ancestry, ancestor)
	}
	return ancestry
}

var yamlSeparator = regexp.MustCompile("\\n---")

// SplitYAMLFile is a helper to split a body into multiple workflow objects
func SplitYAMLFile(body []byte, strict bool) ([]wfv1.Workflow, error) {
	manifestsStrings := yamlSeparator.Split(string(body), -1)
	manifests := make([]wfv1.Workflow, 0)
	for _, manifestStr := range manifestsStrings {
		if strings.TrimSpace(manifestStr) == "" {
			continue
		}
		var wf wfv1.Workflow
		var opts []yaml.JSONOpt
		if strict {
			opts = append(opts, yaml.DisallowUnknownFields) // nolint
		}
		err := yaml.Unmarshal([]byte(manifestStr), &wf, opts...)
		if wf.Kind != "" && wf.Kind != workflow.Kind {
			// If we get here, it was a k8s manifest which was not of type 'Workflow'
			// We ignore these since we only care about Workflow manifests.
			continue
		}
		if err != nil {
			return nil, errors.New(errors.CodeBadRequest, err.Error())
		}
		manifests = append(manifests, wf)
	}
	return manifests, nil
}
