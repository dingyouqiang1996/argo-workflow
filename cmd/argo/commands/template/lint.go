package template

import (
	"fmt"
	"os"

	"github.com/argoproj/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/argoproj/argo/v2/cmd/argo/commands/client"
	"github.com/argoproj/argo/v2/workflow/lint"
)

func NewLintCommand() *cobra.Command {
	var (
		strict bool
		format string
	)
	var command = &cobra.Command{
		Use:   "lint (DIRECTORY | FILE1 FILE2 FILE3...)",
		Short: "validate a file or directory of workflow template manifests",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, apiClient := client.NewAPIClient()
			fmtr, err := lint.GetFormatter(format)
			errors.CheckError(err)

			res, err := lint.Lint(ctx, &lint.LintOptions{
				ServiceClients: lint.ServiceClients{
					WorkflowTemplatesClient: apiClient.NewWorkflowTemplateServiceClient(),
				},
				Files:            args,
				Strict:           strict,
				DefaultNamespace: client.Namespace(),
				Formatter:        fmtr,
			})
			errors.CheckError(err)

			fmt.Print(res)
			if !res.Success {
				os.Exit(1)
			}
		},
	}

	command.Flags().StringVar(&format, "format", "pretty", "Linting results output format. One of: pretty|simple")
	command.Flags().BoolVar(&strict, "strict", true, "perform strict workflow validation")
	return command
}
