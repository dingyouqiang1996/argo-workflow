package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CronWorkflow is the definition of a scheduled workflow resource
// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CronWorkflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              WorkflowSpec        `json:"spec" protobuf:"bytes,2,opt,name=spec"`
	Status            CronWorkflowStatus  `json:"status" protobuf:"bytes,3,opt,name=status"`
	Options           CronWorkflowOptions `json:"options" protobuf:"bytes,4,opt,name=options"`
}

// CronWorkflowList is list of CronWorkflow resources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CronWorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`
	Items           []CronWorkflow `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type CronWorkflowStatus struct {
	// Active is a list of active workflows stemming from this CronWorkflow
	Active []v1.ObjectReference `json:"active,omitempty" protobuf:"bytes,1,rep,name=active"`
	// LastScheduleTime is the last time the CronWorkflow was scheduled
	LastScheduledTime *metav1.Time `json:"lastScheduledTime,omitempty" protobuf:"bytes,2,opt,name=lastScheduledTime"`
}

// CronWorkflowOptions is the schedule of when to run CronWorkflows
type CronWorkflowOptions struct {
	// Schedule is a schedule to run the Workflow in Cron format
	Schedule string `json:"schedule" protobuf:"bytes,1,opt,name=schedule"`
	// RuntimeGenerateName is the name generator the Workflow will be run with. Mutually exclusive with RuntimeName
	RuntimeGenerateName string `json:"runtimeGenerateName" protobuf:"bytes,2,opt,name=runtimeGenerateName"`
	// ConcurrencyPolicy is the K8s-style concurrency policy that will be used
	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrencyPolicy,omitempty" protobuf:"bytes,3,opt,name=concurrencyPolicy,casttype=ConcurrencyPolicy"`
	// Suspend is a flag that will stop new CronWorkflows from running if set to true
	Suspend bool `json:"suspend,omitempty" protobuf:"varint,4,opt,name=suspend"`
	// StartingDeadlineSeconds is the K8s-style deadline that will limit the time a CronWorkflow will be run after its
	// original scheduled time if it is missed.
	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds,omitempty" protobuf:"varint,5,opt,name=startingDeadlineSeconds"`
}

type ConcurrencyPolicy string

const (
	AllowConcurrent   ConcurrencyPolicy = "Allow"
	ForbidConcurrent  ConcurrencyPolicy = "Forbid"
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)
