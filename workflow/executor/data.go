package executor

import (
	"context"
	"encoding/json"
	"fmt"

	"k8s.io/utils/pointer"

	"github.com/argoproj/argo-workflows/v3/workflow/data"
)

func (we *WorkflowExecutor) Data(ctx context.Context) error {
	dataTemplate := we.Template.Data
	if dataTemplate == nil {
		return nil
	}

	transformedData, err := data.ProcessData(dataTemplate, newExecutorDataSourceProcessor(ctx, we))
	if err != nil {
		return fmt.Errorf("unable to process data template: %w", err)
	}

	return we.processOutput(ctx, transformedData)
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
