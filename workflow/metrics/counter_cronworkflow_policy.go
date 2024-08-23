package metrics

import (
	"context"

	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

const (
	nameCronPolicy = `cronworkflows_concurrencypolicy_triggered`
)

func addCronWfPolicyCounter(_ context.Context, m *Metrics) error {
	return m.createInstrument(int64Counter,
		nameCronPolicy,
		"Total number of times CronWorkflow concurrencyPolicy has triggered",
		"{cronworkflow}",
		withAsBuiltIn(),
	)
}

func (m *Metrics) CronWfPolicy(ctx context.Context, name, namespace string, policy wfv1.ConcurrencyPolicy) {
	m.addInt(ctx, nameCronPolicy, 1, instAttribs{
		{name: labelCronWFName, value: name},
		{name: labelWorkflowNamespace, value: namespace},
		{name: labelConcurrencyPolicy, value: string(policy)},
	})
}
