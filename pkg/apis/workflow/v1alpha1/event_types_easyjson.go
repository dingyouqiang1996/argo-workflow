// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha1(in *jlexer.Lexer, out *WorkflowEventBindingSpec) {
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
		case "event":
			(out.Event).UnmarshalEasyJSON(in)
		case "submit":
			if in.IsNull() {
				in.Skip()
				out.Submit = nil
			} else {
				if out.Submit == nil {
					out.Submit = new(Submit)
				}
				(*out.Submit).UnmarshalEasyJSON(in)
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
func easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha1(out *jwriter.Writer, in WorkflowEventBindingSpec) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"event\":"
		out.RawString(prefix[1:])
		(in.Event).MarshalEasyJSON(out)
	}
	if in.Submit != nil {
		const prefix string = ",\"submit\":"
		out.RawString(prefix)
		(*in.Submit).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WorkflowEventBindingSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WorkflowEventBindingSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WorkflowEventBindingSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WorkflowEventBindingSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha1(l, v)
}
func easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha11(in *jlexer.Lexer, out *WorkflowEventBindingList) {
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
			easyjsonDc809a28DecodeK8sIoApimachineryPkgApisMetaV1(in, &out.ListMeta)
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]WorkflowEventBinding, 0, 0)
					} else {
						out.Items = []WorkflowEventBinding{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v1 WorkflowEventBinding
					(v1).UnmarshalEasyJSON(in)
					out.Items = append(out.Items, v1)
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
func easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha11(out *jwriter.Writer, in WorkflowEventBindingList) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"metadata\":"
		out.RawString(prefix[1:])
		easyjsonDc809a28EncodeK8sIoApimachineryPkgApisMetaV1(out, in.ListMeta)
	}
	{
		const prefix string = ",\"items\":"
		out.RawString(prefix)
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Items {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
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
func (v WorkflowEventBindingList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WorkflowEventBindingList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WorkflowEventBindingList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WorkflowEventBindingList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha11(l, v)
}
func easyjsonDc809a28DecodeK8sIoApimachineryPkgApisMetaV1(in *jlexer.Lexer, out *_v1.ListMeta) {
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
func easyjsonDc809a28EncodeK8sIoApimachineryPkgApisMetaV1(out *jwriter.Writer, in _v1.ListMeta) {
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
func easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha12(in *jlexer.Lexer, out *WorkflowEventBinding) {
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
			easyjsonDc809a28DecodeK8sIoApimachineryPkgApisMetaV11(in, &out.ObjectMeta)
		case "spec":
			(out.Spec).UnmarshalEasyJSON(in)
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
func easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha12(out *jwriter.Writer, in WorkflowEventBinding) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"metadata\":"
		out.RawString(prefix[1:])
		easyjsonDc809a28EncodeK8sIoApimachineryPkgApisMetaV11(out, in.ObjectMeta)
	}
	{
		const prefix string = ",\"spec\":"
		out.RawString(prefix)
		(in.Spec).MarshalEasyJSON(out)
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
func (v WorkflowEventBinding) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WorkflowEventBinding) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WorkflowEventBinding) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WorkflowEventBinding) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha12(l, v)
}
func easyjsonDc809a28DecodeK8sIoApimachineryPkgApisMetaV11(in *jlexer.Lexer, out *_v1.ObjectMeta) {
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
					out.DeletionTimestamp = new(_v1.Time)
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
					var v4 string
					v4 = string(in.String())
					(out.Labels)[key] = v4
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
					var v5 string
					v5 = string(in.String())
					(out.Annotations)[key] = v5
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
						out.OwnerReferences = make([]_v1.OwnerReference, 0, 0)
					} else {
						out.OwnerReferences = []_v1.OwnerReference{}
					}
				} else {
					out.OwnerReferences = (out.OwnerReferences)[:0]
				}
				for !in.IsDelim(']') {
					var v6 _v1.OwnerReference
					easyjsonDc809a28DecodeK8sIoApimachineryPkgApisMetaV12(in, &v6)
					out.OwnerReferences = append(out.OwnerReferences, v6)
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
					var v7 string
					v7 = string(in.String())
					out.Finalizers = append(out.Finalizers, v7)
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
						out.ManagedFields = make([]_v1.ManagedFieldsEntry, 0, 0)
					} else {
						out.ManagedFields = []_v1.ManagedFieldsEntry{}
					}
				} else {
					out.ManagedFields = (out.ManagedFields)[:0]
				}
				for !in.IsDelim(']') {
					var v8 _v1.ManagedFieldsEntry
					easyjsonDc809a28DecodeK8sIoApimachineryPkgApisMetaV13(in, &v8)
					out.ManagedFields = append(out.ManagedFields, v8)
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
func easyjsonDc809a28EncodeK8sIoApimachineryPkgApisMetaV11(out *jwriter.Writer, in _v1.ObjectMeta) {
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
			v9First := true
			for v9Name, v9Value := range in.Labels {
				if v9First {
					v9First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v9Name))
				out.RawByte(':')
				out.String(string(v9Value))
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
			v10First := true
			for v10Name, v10Value := range in.Annotations {
				if v10First {
					v10First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v10Name))
				out.RawByte(':')
				out.String(string(v10Value))
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
			for v11, v12 := range in.OwnerReferences {
				if v11 > 0 {
					out.RawByte(',')
				}
				easyjsonDc809a28EncodeK8sIoApimachineryPkgApisMetaV12(out, v12)
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
			for v13, v14 := range in.Finalizers {
				if v13 > 0 {
					out.RawByte(',')
				}
				out.String(string(v14))
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
			for v15, v16 := range in.ManagedFields {
				if v15 > 0 {
					out.RawByte(',')
				}
				easyjsonDc809a28EncodeK8sIoApimachineryPkgApisMetaV13(out, v16)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonDc809a28DecodeK8sIoApimachineryPkgApisMetaV13(in *jlexer.Lexer, out *_v1.ManagedFieldsEntry) {
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
			out.Operation = _v1.ManagedFieldsOperationType(in.String())
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "time":
			if in.IsNull() {
				in.Skip()
				out.Time = nil
			} else {
				if out.Time == nil {
					out.Time = new(_v1.Time)
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
					out.FieldsV1 = new(_v1.FieldsV1)
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
func easyjsonDc809a28EncodeK8sIoApimachineryPkgApisMetaV13(out *jwriter.Writer, in _v1.ManagedFieldsEntry) {
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
func easyjsonDc809a28DecodeK8sIoApimachineryPkgApisMetaV12(in *jlexer.Lexer, out *_v1.OwnerReference) {
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
func easyjsonDc809a28EncodeK8sIoApimachineryPkgApisMetaV12(out *jwriter.Writer, in _v1.OwnerReference) {
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
func easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha13(in *jlexer.Lexer, out *Submit) {
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
		case "workflowTemplateRef":
			(out.WorkflowTemplateRef).UnmarshalEasyJSON(in)
		case "arguments":
			if in.IsNull() {
				in.Skip()
				out.Arguments = nil
			} else {
				if out.Arguments == nil {
					out.Arguments = new(Arguments)
				}
				(*out.Arguments).UnmarshalEasyJSON(in)
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
func easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha13(out *jwriter.Writer, in Submit) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"workflowTemplateRef\":"
		out.RawString(prefix[1:])
		(in.WorkflowTemplateRef).MarshalEasyJSON(out)
	}
	if in.Arguments != nil {
		const prefix string = ",\"arguments\":"
		out.RawString(prefix)
		(*in.Arguments).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Submit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Submit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Submit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Submit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha13(l, v)
}
func easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha14(in *jlexer.Lexer, out *Event) {
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
		case "selector":
			out.Selector = string(in.String())
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
func easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha14(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"selector\":"
		out.RawString(prefix[1:])
		out.String(string(in.Selector))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha14(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDc809a28EncodeGithubComArgoprojArgoPkgApisWorkflowV1alpha14(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha14(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDc809a28DecodeGithubComArgoprojArgoPkgApisWorkflowV1alpha14(l, v)
}
