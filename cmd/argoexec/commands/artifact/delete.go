package artifact

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/argoproj/argo-workflows/v3/cmd/argo/commands/client"
	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	workflow "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	wfv1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	executor "github.com/argoproj/argo-workflows/v3/workflow/artifacts"
	"github.com/argoproj/argo-workflows/v3/workflow/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func NewArtifactDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:          "delete",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			namespace := client.Namespace()
			clientConfig := client.GetConfig()

			if podName, ok := os.LookupEnv(common.EnvVarArtifactPodName); ok {

				config, err := clientConfig.ClientConfig()
				workflowInterface := workflow.NewForConfigOrDie(config)
				checkErr(err)

				artifactGCTaskInterface := workflowInterface.ArgoprojV1alpha1().WorkflowArtifactGCTasks(namespace)
				labelSelector := fmt.Sprintf("%s = %s", common.LabelKeyArtifactGCPodName, podName)

				taskList, err := artifactGCTaskInterface.List(context.Background(), metav1.ListOptions{LabelSelector: labelSelector})
				checkErr(err)

				deleteArtifacts(taskList, cmd, artifactGCTaskInterface)
			}
			return nil
		},
	}
}

func deleteArtifacts(taskList *v1alpha1.WorkflowArtifactGCTaskList, cmd *cobra.Command, artifactGCTaskInterface wfv1alpha1.WorkflowArtifactGCTaskInterface) {
	for _, task := range taskList.Items {
		task.Status.ArtifactResultsByNode = make(map[string]v1alpha1.ArtifactResultNodeStatus)
		for nodeName, artifactNodeSpec := range task.Spec.ArtifactsByNode {

			var archiveLocation *v1alpha1.ArtifactLocation
			artResultNodeStatus := v1alpha1.ArtifactResultNodeStatus{ArtifactResults: make(map[string]v1alpha1.ArtifactResult)}
			if artifactNodeSpec.ArchiveLocation != nil {
				archiveLocation = artifactNodeSpec.ArchiveLocation
			}

			var resources resources
			resources.Files = make(map[string][]byte) // same resources for every artifact
			for _, artifact := range artifactNodeSpec.Artifacts {
				if archiveLocation != nil {
					artifact.Relocate(archiveLocation)
				}

				drv, err := executor.NewDriver(cmd.Context(), &artifact, resources)
				checkErr(err)

				err = drv.Delete(&artifact)
				if err != nil {
					errString := err.Error()
					artResultNodeStatus.ArtifactResults[artifact.Name] = v1alpha1.ArtifactResult{Name: artifact.Name, Success: false, Error: &errString}
				} else {
					artResultNodeStatus.ArtifactResults[artifact.Name] = v1alpha1.ArtifactResult{Name: artifact.Name, Success: true, Error: nil}
				}
			}

			task.Status.ArtifactResultsByNode[nodeName] = artResultNodeStatus
		}
		patch, err := json.Marshal(map[string]interface{}{"status": v1alpha1.ArtifactGCSStatus{ArtifactResultsByNode: task.Status.ArtifactResultsByNode}})
		checkErr(err)
		_, err = artifactGCTaskInterface.Patch(context.Background(), task.Name, types.MergePatchType, patch, metav1.PatchOptions{}, "status")
		checkErr(err)
	}
}

type resources struct {
	Files map[string][]byte
}

func (r resources) GetSecret(ctx context.Context, name, key string) (string, error) {
	// path := filepath.Join("/Users/dpadhiar/go/src/github.com/argoproj/argo-workflows", name, key)
	path := filepath.Join(common.SecretVolMountPath, name, key)
	if file, ok := r.Files[path]; ok {
		return string(file), nil
	}

	file, err := os.ReadFile(path)
	r.Files[path] = file
	return string(file), err
}

func (r resources) GetConfigMapKey(ctx context.Context, name, key string) (string, error) {
	return "", fmt.Errorf("not supported")
}
