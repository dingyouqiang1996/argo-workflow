// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	_v1 "k8s.io/api/core/v1"
	_v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson2ae774dcDecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha1(in *jlexer.Lexer, out *CronWorkflowStatus) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "active":
			if in.IsNull() {
				in.Skip()
				out.Active = nil
			} else {
				in.Delim('[')
				if out.Active == nil {
					if !in.IsDelim(']') {
						out.Active = make([]_v1.ObjectReference, 0, 0)
					} else {
						out.Active = []_v1.ObjectReference{}
					}
				} else {
					out.Active = (out.Active)[:0]
				}
				for !in.IsDelim(']') {
					var v1 _v1.ObjectReference
					easyjson2ae774dcDecodeK8sIoApiCoreV1(in, &v1)
					out.Active = append(out.Active, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "lastScheduledTime":
			if in.IsNull() {
				in.Skip()
				out.LastScheduledTime = nil
			} else {
				if out.LastScheduledTime == nil {
					out.LastScheduledTime = new(_v11.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastScheduledTime).UnmarshalJSON(data))
				}
			}
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make(Conditions, 0, 1)
					} else {
						out.Conditions = Conditions{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v2 Condition
					(v2).UnmarshalEasyJSON(in)
					out.Conditions = append(out.Conditions, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2ae774dcEncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha1(out *jwriter.Writer, in CronWorkflowStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Active) != 0 {
		const prefix string = ",\"active\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v3, v4 := range in.Active {
				if v3 > 0 {
					out.RawByte(',')
				}
				easyjson2ae774dcEncodeK8sIoApiCoreV1(out, v4)
			}
			out.RawByte(']')
		}
	}
	if in.LastScheduledTime != nil {
		const prefix string = ",\"lastScheduledTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.LastScheduledTime).MarshalJSON())
	}
	if len(in.Conditions) != 0 {
		const prefix string = ",\"conditions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Conditions {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CronWorkflowStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2ae774dcEncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CronWorkflowStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2ae774dcEncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CronWorkflowStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data, CoerceToString: true}
	easyjson2ae774dcDecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CronWorkflowStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2ae774dcDecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha1(l, v)
}
func easyjson2ae774dcDecodeK8sIoApiCoreV1(in *jlexer.Lexer, out *_v1.ObjectReference) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "kind":
			out.Kind = string(in.String())
		case "namespace":
			out.Namespace = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "uid":
			out.UID = types.UID(in.String())
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "resourceVersion":
			out.ResourceVersion = string(in.String())
		case "fieldPath":
			out.FieldPath = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2ae774dcEncodeK8sIoApiCoreV1(out *jwriter.Writer, in _v1.ObjectReference) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Kind))
	}
	if in.Namespace != "" {
		const prefix string = ",\"namespace\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Namespace))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.UID != "" {
		const prefix string = ",\"uid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UID))
	}
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.APIVersion))
	}
	if in.ResourceVersion != "" {
		const prefix string = ",\"resourceVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ResourceVersion))
	}
	if in.FieldPath != "" {
		const prefix string = ",\"fieldPath\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FieldPath))
	}
	out.RawByte('}')
}
func easyjson2ae774dcDecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha11(in *jlexer.Lexer, out *CronWorkflowSpec) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "workflowSpec":
			(out.WorkflowSpec).UnmarshalEasyJSON(in)
		case "schedule":
			out.Schedule = string(in.String())
		case "concurrencyPolicy":
			out.ConcurrencyPolicy = ConcurrencyPolicy(in.String())
		case "suspend":
			out.Suspend = bool(in.Bool())
		case "startingDeadlineSeconds":
			if in.IsNull() {
				in.Skip()
				out.StartingDeadlineSeconds = nil
			} else {
				if out.StartingDeadlineSeconds == nil {
					out.StartingDeadlineSeconds = new(int64)
				}
				*out.StartingDeadlineSeconds = int64(in.Int64())
			}
		case "successfulJobsHistoryLimit":
			if in.IsNull() {
				in.Skip()
				out.SuccessfulJobsHistoryLimit = nil
			} else {
				if out.SuccessfulJobsHistoryLimit == nil {
					out.SuccessfulJobsHistoryLimit = new(int32)
				}
				*out.SuccessfulJobsHistoryLimit = int32(in.Int32())
			}
		case "failedJobsHistoryLimit":
			if in.IsNull() {
				in.Skip()
				out.FailedJobsHistoryLimit = nil
			} else {
				if out.FailedJobsHistoryLimit == nil {
					out.FailedJobsHistoryLimit = new(int32)
				}
				*out.FailedJobsHistoryLimit = int32(in.Int32())
			}
		case "timezone":
			out.Timezone = string(in.String())
		case "workflowMetadata":
			if in.IsNull() {
				in.Skip()
				out.WorkflowMetadata = nil
			} else {
				if out.WorkflowMetadata == nil {
					out.WorkflowMetadata = new(_v11.ObjectMeta)
				}
				easyjson2ae774dcDecodeK8sIoApimachineryPkgApisMetaV1(in, out.WorkflowMetadata)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2ae774dcEncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha11(out *jwriter.Writer, in CronWorkflowSpec) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"workflowSpec\":"
		out.RawString(prefix[1:])
		(in.WorkflowSpec).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"schedule\":"
		out.RawString(prefix)
		out.String(string(in.Schedule))
	}
	if in.ConcurrencyPolicy != "" {
		const prefix string = ",\"concurrencyPolicy\":"
		out.RawString(prefix)
		out.String(string(in.ConcurrencyPolicy))
	}
	if in.Suspend {
		const prefix string = ",\"suspend\":"
		out.RawString(prefix)
		out.Bool(bool(in.Suspend))
	}
	if in.StartingDeadlineSeconds != nil {
		const prefix string = ",\"startingDeadlineSeconds\":"
		out.RawString(prefix)
		out.Int64(int64(*in.StartingDeadlineSeconds))
	}
	if in.SuccessfulJobsHistoryLimit != nil {
		const prefix string = ",\"successfulJobsHistoryLimit\":"
		out.RawString(prefix)
		out.Int32(int32(*in.SuccessfulJobsHistoryLimit))
	}
	if in.FailedJobsHistoryLimit != nil {
		const prefix string = ",\"failedJobsHistoryLimit\":"
		out.RawString(prefix)
		out.Int32(int32(*in.FailedJobsHistoryLimit))
	}
	if in.Timezone != "" {
		const prefix string = ",\"timezone\":"
		out.RawString(prefix)
		out.String(string(in.Timezone))
	}
	if in.WorkflowMetadata != nil {
		const prefix string = ",\"workflowMetadata\":"
		out.RawString(prefix)
		easyjson2ae774dcEncodeK8sIoApimachineryPkgApisMetaV1(out, *in.WorkflowMetadata)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CronWorkflowSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2ae774dcEncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CronWorkflowSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2ae774dcEncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CronWorkflowSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data, CoerceToString: true}
	easyjson2ae774dcDecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CronWorkflowSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2ae774dcDecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha11(l, v)
}
func easyjson2ae774dcDecodeK8sIoApimachineryPkgApisMetaV1(in *jlexer.Lexer, out *_v11.ObjectMeta) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "generateName":
			out.GenerateName = string(in.String())
		case "namespace":
			out.Namespace = string(in.String())
		case "selfLink":
			out.SelfLink = string(in.String())
		case "uid":
			out.UID = types.UID(in.String())
		case "resourceVersion":
			out.ResourceVersion = string(in.String())
		case "generation":
			out.Generation = int64(in.Int64())
		case "creationTimestamp":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreationTimestamp).UnmarshalJSON(data))
			}
		case "deletionTimestamp":
			if in.IsNull() {
				in.Skip()
				out.DeletionTimestamp = nil
			} else {
				if out.DeletionTimestamp == nil {
					out.DeletionTimestamp = new(_v11.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.DeletionTimestamp).UnmarshalJSON(data))
				}
			}
		case "deletionGracePeriodSeconds":
			if in.IsNull() {
				in.Skip()
				out.DeletionGracePeriodSeconds = nil
			} else {
				if out.DeletionGracePeriodSeconds == nil {
					out.DeletionGracePeriodSeconds = new(int64)
				}
				*out.DeletionGracePeriodSeconds = int64(in.Int64())
			}
		case "labels":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Labels = make(map[string]string)
				} else {
					out.Labels = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v7 string
					v7 = string(in.String())
					(out.Labels)[key] = v7
					in.WantComma()
				}
				in.Delim('}')
			}
		case "annotations":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Annotations = make(map[string]string)
				} else {
					out.Annotations = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v8 string
					v8 = string(in.String())
					(out.Annotations)[key] = v8
					in.WantComma()
				}
				in.Delim('}')
			}
		case "ownerReferences":
			if in.IsNull() {
				in.Skip()
				out.OwnerReferences = nil
			} else {
				in.Delim('[')
				if out.OwnerReferences == nil {
					if !in.IsDelim(']') {
						out.OwnerReferences = make([]_v11.OwnerReference, 0, 0)
					} else {
						out.OwnerReferences = []_v11.OwnerReference{}
					}
				} else {
					out.OwnerReferences = (out.OwnerReferences)[:0]
				}
				for !in.IsDelim(']') {
					var v9 _v11.OwnerReference
					easyjson2ae774dcDecodeK8sIoApimachineryPkgApisMetaV11(in, &v9)
					out.OwnerReferences = append(out.OwnerReferences, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "finalizers":
			if in.IsNull() {
				in.Skip()
				out.Finalizers = nil
			} else {
				in.Delim('[')
				if out.Finalizers == nil {
					if !in.IsDelim(']') {
						out.Finalizers = make([]string, 0, 4)
					} else {
						out.Finalizers = []string{}
					}
				} else {
					out.Finalizers = (out.Finalizers)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.Finalizers = append(out.Finalizers, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "clusterName":
			out.ClusterName = string(in.String())
		case "managedFields":
			if in.IsNull() {
				in.Skip()
				out.ManagedFields = nil
			} else {
				in.Delim('[')
				if out.ManagedFields == nil {
					if !in.IsDelim(']') {
						out.ManagedFields = make([]_v11.ManagedFieldsEntry, 0, 0)
					} else {
						out.ManagedFields = []_v11.ManagedFieldsEntry{}
					}
				} else {
					out.ManagedFields = (out.ManagedFields)[:0]
				}
				for !in.IsDelim(']') {
					var v11 _v11.ManagedFieldsEntry
					easyjson2ae774dcDecodeK8sIoApimachineryPkgApisMetaV12(in, &v11)
					out.ManagedFields = append(out.ManagedFields, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2ae774dcEncodeK8sIoApimachineryPkgApisMetaV1(out *jwriter.Writer, in _v11.ObjectMeta) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.GenerateName != "" {
		const prefix string = ",\"generateName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.GenerateName))
	}
	if in.Namespace != "" {
		const prefix string = ",\"namespace\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Namespace))
	}
	if in.SelfLink != "" {
		const prefix string = ",\"selfLink\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SelfLink))
	}
	if in.UID != "" {
		const prefix string = ",\"uid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UID))
	}
	if in.ResourceVersion != "" {
		const prefix string = ",\"resourceVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ResourceVersion))
	}
	if in.Generation != 0 {
		const prefix string = ",\"generation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Generation))
	}
	if true {
		const prefix string = ",\"creationTimestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.CreationTimestamp).MarshalJSON())
	}
	if in.DeletionTimestamp != nil {
		const prefix string = ",\"deletionTimestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.DeletionTimestamp).MarshalJSON())
	}
	if in.DeletionGracePeriodSeconds != nil {
		const prefix string = ",\"deletionGracePeriodSeconds\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(*in.DeletionGracePeriodSeconds))
	}
	if len(in.Labels) != 0 {
		const prefix string = ",\"labels\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v12First := true
			for v12Name, v12Value := range in.Labels {
				if v12First {
					v12First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v12Name))
				out.RawByte(':')
				out.String(string(v12Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.Annotations) != 0 {
		const prefix string = ",\"annotations\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v13First := true
			for v13Name, v13Value := range in.Annotations {
				if v13First {
					v13First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v13Name))
				out.RawByte(':')
				out.String(string(v13Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.OwnerReferences) != 0 {
		const prefix string = ",\"ownerReferences\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v14, v15 := range in.OwnerReferences {
				if v14 > 0 {
					out.RawByte(',')
				}
				easyjson2ae774dcEncodeK8sIoApimachineryPkgApisMetaV11(out, v15)
			}
			out.RawByte(']')
		}
	}
	if len(in.Finalizers) != 0 {
		const prefix string = ",\"finalizers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v16, v17 := range in.Finalizers {
				if v16 > 0 {
					out.RawByte(',')
				}
				out.String(string(v17))
			}
			out.RawByte(']')
		}
	}
	if in.ClusterName != "" {
		const prefix string = ",\"clusterName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ClusterName))
	}
	if len(in.ManagedFields) != 0 {
		const prefix string = ",\"managedFields\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v18, v19 := range in.ManagedFields {
				if v18 > 0 {
					out.RawByte(',')
				}
				easyjson2ae774dcEncodeK8sIoApimachineryPkgApisMetaV12(out, v19)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson2ae774dcDecodeK8sIoApimachineryPkgApisMetaV12(in *jlexer.Lexer, out *_v11.ManagedFieldsEntry) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "manager":
			out.Manager = string(in.String())
		case "operation":
			out.Operation = _v11.ManagedFieldsOperationType(in.String())
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "time":
			if in.IsNull() {
				in.Skip()
				out.Time = nil
			} else {
				if out.Time == nil {
					out.Time = new(_v11.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Time).UnmarshalJSON(data))
				}
			}
		case "fieldsType":
			out.FieldsType = string(in.String())
		case "fieldsV1":
			if in.IsNull() {
				in.Skip()
				out.FieldsV1 = nil
			} else {
				if out.FieldsV1 == nil {
					out.FieldsV1 = new(_v11.FieldsV1)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.FieldsV1).UnmarshalJSON(data))
				}
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2ae774dcEncodeK8sIoApimachineryPkgApisMetaV12(out *jwriter.Writer, in _v11.ManagedFieldsEntry) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Manager != "" {
		const prefix string = ",\"manager\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Manager))
	}
	if in.Operation != "" {
		const prefix string = ",\"operation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Operation))
	}
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.APIVersion))
	}
	if in.Time != nil {
		const prefix string = ",\"time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.Time).MarshalJSON())
	}
	if in.FieldsType != "" {
		const prefix string = ",\"fieldsType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FieldsType))
	}
	if in.FieldsV1 != nil {
		const prefix string = ",\"fieldsV1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.FieldsV1).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson2ae774dcDecodeK8sIoApimachineryPkgApisMetaV11(in *jlexer.Lexer, out *_v11.OwnerReference) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "uid":
			out.UID = types.UID(in.String())
		case "controller":
			if in.IsNull() {
				in.Skip()
				out.Controller = nil
			} else {
				if out.Controller == nil {
					out.Controller = new(bool)
				}
				*out.Controller = bool(in.Bool())
			}
		case "blockOwnerDeletion":
			if in.IsNull() {
				in.Skip()
				out.BlockOwnerDeletion = nil
			} else {
				if out.BlockOwnerDeletion == nil {
					out.BlockOwnerDeletion = new(bool)
				}
				*out.BlockOwnerDeletion = bool(in.Bool())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2ae774dcEncodeK8sIoApimachineryPkgApisMetaV11(out *jwriter.Writer, in _v11.OwnerReference) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"apiVersion\":"
		out.RawString(prefix[1:])
		out.String(string(in.APIVersion))
	}
	{
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"uid\":"
		out.RawString(prefix)
		out.String(string(in.UID))
	}
	if in.Controller != nil {
		const prefix string = ",\"controller\":"
		out.RawString(prefix)
		out.Bool(bool(*in.Controller))
	}
	if in.BlockOwnerDeletion != nil {
		const prefix string = ",\"blockOwnerDeletion\":"
		out.RawString(prefix)
		out.Bool(bool(*in.BlockOwnerDeletion))
	}
	out.RawByte('}')
}
func easyjson2ae774dcDecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha12(in *jlexer.Lexer, out *CronWorkflowList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "metadata":
			easyjson2ae774dcDecodeK8sIoApimachineryPkgApisMetaV13(in, &out.ListMeta)
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]CronWorkflow, 0, 0)
					} else {
						out.Items = []CronWorkflow{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v20 CronWorkflow
					(v20).UnmarshalEasyJSON(in)
					out.Items = append(out.Items, v20)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "kind":
			out.Kind = string(in.String())
		case "apiVersion":
			out.APIVersion = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2ae774dcEncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha12(out *jwriter.Writer, in CronWorkflowList) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"metadata\":"
		out.RawString(prefix[1:])
		easyjson2ae774dcEncodeK8sIoApimachineryPkgApisMetaV13(out, in.ListMeta)
	}
	{
		const prefix string = ",\"items\":"
		out.RawString(prefix)
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v21, v22 := range in.Items {
				if v21 > 0 {
					out.RawByte(',')
				}
				(v22).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		out.RawString(prefix)
		out.String(string(in.APIVersion))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CronWorkflowList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2ae774dcEncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CronWorkflowList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2ae774dcEncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CronWorkflowList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data, CoerceToString: true}
	easyjson2ae774dcDecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CronWorkflowList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2ae774dcDecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha12(l, v)
}
func easyjson2ae774dcDecodeK8sIoApimachineryPkgApisMetaV13(in *jlexer.Lexer, out *_v11.ListMeta) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "selfLink":
			out.SelfLink = string(in.String())
		case "resourceVersion":
			out.ResourceVersion = string(in.String())
		case "continue":
			out.Continue = string(in.String())
		case "remainingItemCount":
			if in.IsNull() {
				in.Skip()
				out.RemainingItemCount = nil
			} else {
				if out.RemainingItemCount == nil {
					out.RemainingItemCount = new(int64)
				}
				*out.RemainingItemCount = int64(in.Int64())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2ae774dcEncodeK8sIoApimachineryPkgApisMetaV13(out *jwriter.Writer, in _v11.ListMeta) {
	out.RawByte('{')
	first := true
	_ = first
	if in.SelfLink != "" {
		const prefix string = ",\"selfLink\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.SelfLink))
	}
	if in.ResourceVersion != "" {
		const prefix string = ",\"resourceVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ResourceVersion))
	}
	if in.Continue != "" {
		const prefix string = ",\"continue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Continue))
	}
	if in.RemainingItemCount != nil {
		const prefix string = ",\"remainingItemCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(*in.RemainingItemCount))
	}
	out.RawByte('}')
}
func easyjson2ae774dcDecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha13(in *jlexer.Lexer, out *CronWorkflow) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "metadata":
			easyjson2ae774dcDecodeK8sIoApimachineryPkgApisMetaV1(in, &out.ObjectMeta)
		case "spec":
			(out.Spec).UnmarshalEasyJSON(in)
		case "status":
			(out.Status).UnmarshalEasyJSON(in)
		case "kind":
			out.Kind = string(in.String())
		case "apiVersion":
			out.APIVersion = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2ae774dcEncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha13(out *jwriter.Writer, in CronWorkflow) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"metadata\":"
		out.RawString(prefix[1:])
		easyjson2ae774dcEncodeK8sIoApimachineryPkgApisMetaV1(out, in.ObjectMeta)
	}
	{
		const prefix string = ",\"spec\":"
		out.RawString(prefix)
		(in.Spec).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		(in.Status).MarshalEasyJSON(out)
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		out.RawString(prefix)
		out.String(string(in.Kind))
	}
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		out.RawString(prefix)
		out.String(string(in.APIVersion))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CronWorkflow) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2ae774dcEncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CronWorkflow) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2ae774dcEncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CronWorkflow) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data, CoerceToString: true}
	easyjson2ae774dcDecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CronWorkflow) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2ae774dcDecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha13(l, v)
}
