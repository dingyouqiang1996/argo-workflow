package controller

import (
	apiv1 "k8s.io/api/core/v1"
)

/*
These are tho only fields that we ever look at significant in, some are immutable:
metadata:
  name*
  namespace
  labels: {}
  annotations: {}
  deletionTimestamp
spec:
  serviceAccountName
  nodeName
status:
  phase
  message
  podIp
  initContainerStatuses: []
  containerStatuses: []
  conditions: []
*/
// TODO is the only change we really need to care about the phase of the pod?
func significantPodChange(from *apiv1.Pod, to *apiv1.Pod) bool {
	return from.Spec.NodeName != to.Spec.NodeName ||
		from.Status.Phase != to.Status.Phase ||
		from.Status.Message != to.Status.Message ||
		from.Status.PodIP != to.Status.PodIP ||
		significantContainerStatusesChange(from.Status.ContainerStatuses, to.Status.ContainerStatuses) ||
		significantContainerStatusesChange(from.Status.InitContainerStatuses, to.Status.InitContainerStatuses)
}

func significantContainerStatusesChange(from []apiv1.ContainerStatus, to []apiv1.ContainerStatus) bool {
	if len(from) != len(to) {
		return true
	}
	statuses := map[string]apiv1.ContainerStatus{}
	for _, s := range from {
		statuses[s.Name] = s
	}
	for _, s := range to {
		if significantContainerStatusChange(statuses[s.Name], s) {
			return true
		}
	}
	return false
}

func significantContainerStatusChange(from apiv1.ContainerStatus, to apiv1.ContainerStatus) bool {
	return from.Ready != to.Ready || significantContainerStateChange(from.State, to.State)
}

func significantContainerStateChange(from apiv1.ContainerState, to apiv1.ContainerState) bool {
	// waiting has two significant fields and either could potentially change
	return to.Waiting != nil && (from.Waiting == nil || from.Waiting.Message != to.Waiting.Message || from.Waiting.Reason != to.Waiting.Reason) ||
		// running only has one field which is immutable -  so any change is significant
		(to.Running != nil && from.Running == nil) ||
		// I'm assuming this field is immutable - so any change is significant
		(to.Terminated != nil && from.Terminated == nil)
}
