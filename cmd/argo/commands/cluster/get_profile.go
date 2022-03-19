package cluster

import (
	"context"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
	"sigs.k8s.io/yaml"

	"github.com/argoproj/argo-workflows/v3/workflow/common"
)

func newGetProfileCommand() *cobra.Command {
	var (
		kubeconfig             string
		remoteServer           string
		remoteContext          string
		localNamespace         string
		remoteNamespace        string
		remoteInstallNamespace string
	)
	cmd := &cobra.Command{
		Use: "get-profile local_cluster remote_cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			if len(args) != 2 {
				cmd.HelpFunc()(cmd, args)
				os.Exit(1)
			}

			localCluster, remoteCluster := args[0], args[1]

			if remoteContext == "" {
				remoteContext = remoteCluster
			}
			if remoteInstallNamespace == "" {
				remoteInstallNamespace = remoteNamespace
			}

			clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
				&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
				&clientcmd.ConfigOverrides{CurrentContext: remoteContext},
			)

			remoteServiceAccountName := remoteServiceAccountName(localNamespace, localCluster, remoteNamespace)

			config, err := clientConfig.ClientConfig()
			if err != nil {
				return err
			}

			if remoteServer == "" {
				remoteServer = config.Host
			}

			client := kubernetes.NewForConfigOrDie(config)

			serviceAccount, err := client.CoreV1().ServiceAccounts(remoteInstallNamespace).Get(ctx, remoteServiceAccountName, metav1.GetOptions{})
			if err != nil {
				return err
			}
			secretName := serviceAccount.Secrets[0].Name

			secret, err := client.CoreV1().Secrets(remoteInstallNamespace).Get(ctx, secretName, metav1.GetOptions{})
			if err != nil {
				return err
			}

			ca := secret.Data[apiv1.ServiceAccountRootCAKey]
			token := secret.Data[apiv1.ServiceAccountTokenKey]

			kubeconfig := api.Config{
				Kind:       "Config",
				APIVersion: "v1",
				Clusters: map[string]*api.Cluster{
					remoteCluster: {Server: remoteServer, CertificateAuthorityData: ca},
				},
				AuthInfos: map[string]*api.AuthInfo{
					remoteServiceAccountName: {Token: string(token)},
				},
				Contexts: map[string]*api.Context{
					remoteCluster: {Cluster: remoteCluster, AuthInfo: remoteServiceAccountName},
				},
				CurrentContext: remoteCluster,
			}

			data, err := clientcmd.Write(kubeconfig)
			if err != nil {
				return err
			}

			profile := &apiv1.Secret{
				TypeMeta: metav1.TypeMeta{Kind: "Secret", APIVersion: "v1"},
				ObjectMeta: metav1.ObjectMeta{
					Name: remoteServiceAccountName,
					Labels: map[string]string{
						common.LabelKeyCluster:           remoteCluster,
						common.LabelKeyWorkflowNamespace: localNamespace,
					},
					Annotations: map[string]string{
						common.AnnotationKeyNamespace: remoteNamespace,
					},
				},
				Data: map[string][]byte{"kubeconfig": data},
			}

			data, err = yaml.Marshal(profile)
			if err != nil {
				return err
			}

			println("ALEX", remoteServer)

			_, _ = os.Stdout.WriteString("# This is an auto-generated file. DO NOT EDIT\n")
			_, _ = os.Stdout.Write(data)

			return nil
		},
	}
	cmd.Flags().StringVar(&kubeconfig, "kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "kubeconfig")
	cmd.Flags().StringVar(&remoteServer, "remote-server", "", "URL for remote server")
	cmd.Flags().StringVar(&remoteContext, "remote-context", "", "remote context")
	cmd.Flags().StringVar(&localNamespace, "local-namespace", "", "restrict to this local namespace (empty for all namespaces)")
	cmd.Flags().StringVar(&remoteNamespace, "remote-namespace", "", "restrict the this remote namespace (empty for all namespaces)")
	cmd.Flags().StringVar(&remoteInstallNamespace, "remote-install-namespace", "", "the remote namespace that the service account is created in")
	return cmd
}
