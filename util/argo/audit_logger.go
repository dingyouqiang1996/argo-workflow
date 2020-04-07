package argo

import (
	"fmt"
	"time"

	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"

	log "github.com/sirupsen/logrus"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type AuditLogger struct {
	kIf       kubernetes.Interface
	component string
	ns        string
}

type EventInfo struct {
	Type   string
	Reason string
}

type ObjectRef struct {
	Name            string
	Namespace       string
	ResourceVersion string
	UID             types.UID
}

const (
	EventReasonWorkflowRunning       = "WorkflowRunning"
	EventReasonWorkflowSucceeded     = "WorkflowSucceeded"
	EventReasonWorkflowFailed        = "WorkflowFailed"
	EventReasonWorkflowTimedOut      = "WorkflowTimedOut"
	EventReasonWorkflowNodeSucceeded = "WorkflowNodeSucceeded"
	EventReasonWorkflowNodeFailed    = "WorkflowNodeFailed"
	EventReasonWorkflowNodeError     = "WorkflowNodeError"
)

func (l *AuditLogger) logEvent(objMeta ObjectRef, gvk schema.GroupVersionKind, info EventInfo, message string, annotations map[string]string) {
	logCtx := log.WithFields(log.Fields{
		"type":   info.Type,
		"reason": info.Reason,
	})
	switch gvk.Kind {
	case "Workflow":
		logCtx = logCtx.WithField("workflow", objMeta.Name)
	default:
		logCtx = logCtx.WithField("name", objMeta.Name)
	}
	t := metav1.Time{Time: time.Now()}
	event := v1.Event{
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprintf("%v.%x", objMeta.Name, t.UnixNano()),
			Annotations: annotations,
		},
		Source: v1.EventSource{
			Component: l.component,
		},
		InvolvedObject: v1.ObjectReference{
			Kind:            gvk.Kind,
			Name:            objMeta.Name,
			Namespace:       objMeta.Namespace,
			ResourceVersion: objMeta.ResourceVersion,
			APIVersion:      gvk.Version,
			UID:             objMeta.UID,
		},

		FirstTimestamp: t,
		LastTimestamp:  t,
		Count:          1,
		Message:        message,
		Type:           info.Type,
		Reason:         info.Reason,
	}
	logCtx.WithField("event", event).Debug()
	_, err := l.kIf.CoreV1().Events(objMeta.Namespace).Create(&event)
	if err != nil {
		logCtx.Errorf("Unable to create audit event: %v", err)
		return
	}
}

func (l *AuditLogger) LogWorkflowEvent(workflow *v1alpha1.Workflow, info EventInfo, message string, annotations map[string]string) {
	objectMeta := ObjectRef{
		Name:            workflow.ObjectMeta.Name,
		Namespace:       workflow.ObjectMeta.Namespace,
		ResourceVersion: workflow.ObjectMeta.ResourceVersion,
		UID:             workflow.ObjectMeta.UID,
	}
	l.logEvent(objectMeta, v1alpha1.WorkflowSchemaGroupVersionKind, info, message, annotations)
}

func NewAuditLogger(ns string, kIf kubernetes.Interface, component string) *AuditLogger {
	return &AuditLogger{
		ns:        ns,
		kIf:       kIf,
		component: component,
	}
}
