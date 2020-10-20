// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/apiclient/workflowarchive/workflow-archive.proto

package workflowarchive

import (
	context "context"
	fmt "fmt"
	v1alpha1 "github.com/argoproj/argo/v3/pkg/apis/workflow/v1alpha1"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ListArchivedWorkflowsRequest struct {
	ListOptions          *v1.ListOptions `protobuf:"bytes,1,opt,name=listOptions,proto3" json:"listOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListArchivedWorkflowsRequest) Reset()         { *m = ListArchivedWorkflowsRequest{} }
func (m *ListArchivedWorkflowsRequest) String() string { return proto.CompactTextString(m) }
func (*ListArchivedWorkflowsRequest) ProtoMessage()    {}
func (*ListArchivedWorkflowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca9a2d33e8bb19, []int{0}
}
func (m *ListArchivedWorkflowsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListArchivedWorkflowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListArchivedWorkflowsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListArchivedWorkflowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArchivedWorkflowsRequest.Merge(m, src)
}
func (m *ListArchivedWorkflowsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListArchivedWorkflowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArchivedWorkflowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListArchivedWorkflowsRequest proto.InternalMessageInfo

func (m *ListArchivedWorkflowsRequest) GetListOptions() *v1.ListOptions {
	if m != nil {
		return m.ListOptions
	}
	return nil
}

type GetArchivedWorkflowRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArchivedWorkflowRequest) Reset()         { *m = GetArchivedWorkflowRequest{} }
func (m *GetArchivedWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*GetArchivedWorkflowRequest) ProtoMessage()    {}
func (*GetArchivedWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca9a2d33e8bb19, []int{1}
}
func (m *GetArchivedWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetArchivedWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetArchivedWorkflowRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetArchivedWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArchivedWorkflowRequest.Merge(m, src)
}
func (m *GetArchivedWorkflowRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetArchivedWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArchivedWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArchivedWorkflowRequest proto.InternalMessageInfo

func (m *GetArchivedWorkflowRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type DeleteArchivedWorkflowRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteArchivedWorkflowRequest) Reset()         { *m = DeleteArchivedWorkflowRequest{} }
func (m *DeleteArchivedWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteArchivedWorkflowRequest) ProtoMessage()    {}
func (*DeleteArchivedWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca9a2d33e8bb19, []int{2}
}
func (m *DeleteArchivedWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeleteArchivedWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeleteArchivedWorkflowRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeleteArchivedWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArchivedWorkflowRequest.Merge(m, src)
}
func (m *DeleteArchivedWorkflowRequest) XXX_Size() int {
	return m.Size()
}
func (m *DeleteArchivedWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArchivedWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArchivedWorkflowRequest proto.InternalMessageInfo

func (m *DeleteArchivedWorkflowRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type ArchivedWorkflowDeletedResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArchivedWorkflowDeletedResponse) Reset()         { *m = ArchivedWorkflowDeletedResponse{} }
func (m *ArchivedWorkflowDeletedResponse) String() string { return proto.CompactTextString(m) }
func (*ArchivedWorkflowDeletedResponse) ProtoMessage()    {}
func (*ArchivedWorkflowDeletedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca9a2d33e8bb19, []int{3}
}
func (m *ArchivedWorkflowDeletedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ArchivedWorkflowDeletedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ArchivedWorkflowDeletedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ArchivedWorkflowDeletedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchivedWorkflowDeletedResponse.Merge(m, src)
}
func (m *ArchivedWorkflowDeletedResponse) XXX_Size() int {
	return m.Size()
}
func (m *ArchivedWorkflowDeletedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchivedWorkflowDeletedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArchivedWorkflowDeletedResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListArchivedWorkflowsRequest)(nil), "workflowarchive.ListArchivedWorkflowsRequest")
	proto.RegisterType((*GetArchivedWorkflowRequest)(nil), "workflowarchive.GetArchivedWorkflowRequest")
	proto.RegisterType((*DeleteArchivedWorkflowRequest)(nil), "workflowarchive.DeleteArchivedWorkflowRequest")
	proto.RegisterType((*ArchivedWorkflowDeletedResponse)(nil), "workflowarchive.ArchivedWorkflowDeletedResponse")
}

func init() {
	proto.RegisterFile("pkg/apiclient/workflowarchive/workflow-archive.proto", fileDescriptor_95ca9a2d33e8bb19)
}

var fileDescriptor_95ca9a2d33e8bb19 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x8a, 0xd4, 0x40,
	0x10, 0xc6, 0x69, 0x15, 0xc1, 0xde, 0x83, 0xd2, 0xe2, 0x1f, 0xc2, 0x38, 0xae, 0x39, 0x2d, 0xca,
	0x54, 0x9b, 0x9d, 0x45, 0x3c, 0xaa, 0x0c, 0x7a, 0x59, 0x10, 0x66, 0x0f, 0x82, 0xb7, 0xde, 0xa4,
	0xcc, 0xb4, 0x93, 0x49, 0xb7, 0xdd, 0x3d, 0x3d, 0x88, 0x78, 0xf1, 0x15, 0x7c, 0x85, 0x3d, 0x8b,
	0x8f, 0xe1, 0x51, 0xf0, 0x05, 0x64, 0xf0, 0x41, 0x24, 0x99, 0x64, 0xb3, 0x24, 0xf3, 0x47, 0xf0,
	0x56, 0xe9, 0x54, 0x7d, 0xf5, 0xe3, 0xfb, 0xd2, 0xa1, 0x47, 0x7a, 0x9a, 0x72, 0xa1, 0x65, 0x9c,
	0x49, 0xcc, 0x1d, 0x5f, 0x28, 0x33, 0x7d, 0x97, 0xa9, 0x85, 0x30, 0xf1, 0x44, 0x7a, 0x3c, 0x7f,
	0x1e, 0x54, 0x07, 0xa0, 0x8d, 0x72, 0x8a, 0x5d, 0x6f, 0xf5, 0x05, 0xbd, 0x54, 0xa9, 0x34, 0xc3,
	0x42, 0x89, 0x8b, 0x3c, 0x57, 0x4e, 0x38, 0xa9, 0x72, 0xbb, 0x6a, 0x0f, 0x8e, 0xa6, 0x4f, 0x2d,
	0x48, 0x55, 0xbc, 0x9d, 0x89, 0x78, 0x22, 0x73, 0x34, 0x1f, 0x79, 0xb5, 0xd8, 0xf2, 0x19, 0x3a,
	0xc1, 0x7d, 0xc4, 0x53, 0xcc, 0xd1, 0x08, 0x87, 0x49, 0x35, 0xf5, 0x32, 0x95, 0x6e, 0x32, 0x3f,
	0x85, 0x58, 0xcd, 0xb8, 0x30, 0xa9, 0xd2, 0x46, 0xbd, 0x2f, 0x0b, 0xee, 0x87, 0xcd, 0x74, 0x4d,
	0xc2, 0x7d, 0x24, 0x32, 0x3d, 0x11, 0x1d, 0x9d, 0xd0, 0xd2, 0xde, 0xb1, 0xb4, 0xee, 0xf9, 0x0a,
	0x35, 0x79, 0x53, 0x0d, 0xd8, 0x31, 0x7e, 0x98, 0xa3, 0x75, 0xec, 0x84, 0xee, 0x65, 0xd2, 0xba,
	0xd7, 0xba, 0x44, 0xbe, 0x4b, 0xf6, 0xc9, 0xc1, 0xde, 0x61, 0x04, 0x2b, 0x66, 0xb8, 0xc8, 0x0c,
	0x7a, 0x9a, 0x16, 0x07, 0x16, 0x0a, 0x66, 0xf0, 0x11, 0x1c, 0x37, 0x83, 0xe3, 0x8b, 0x2a, 0x21,
	0xd0, 0xe0, 0x15, 0x76, 0x76, 0xd6, 0x2b, 0x6f, 0xd0, 0xcb, 0x73, 0x99, 0x94, 0xab, 0xae, 0x8d,
	0x8b, 0x32, 0x8c, 0xe8, 0xbd, 0x11, 0x66, 0xe8, 0xf0, 0xdf, 0x47, 0x1e, 0xd0, 0xfb, 0xed, 0xe6,
	0x95, 0x44, 0x32, 0x46, 0xab, 0x55, 0x6e, 0xf1, 0xf0, 0xec, 0x0a, 0xbd, 0xd3, 0xee, 0x39, 0x41,
	0xe3, 0x65, 0x8c, 0xec, 0x3b, 0xa1, 0xb7, 0xd6, 0xfa, 0xc2, 0x06, 0xd0, 0x8a, 0x17, 0xb6, 0xf9,
	0x17, 0x8c, 0xa0, 0x09, 0x0a, 0xea, 0xa0, 0xca, 0x02, 0xfc, 0xb0, 0xb1, 0xac, 0xd6, 0x84, 0x3a,
	0x28, 0xa8, 0x95, 0x0a, 0xf5, 0x30, 0xfc, 0xf2, 0xeb, 0xcf, 0xd7, 0x4b, 0x3d, 0x16, 0x94, 0xdf,
	0x90, 0x8f, 0x78, 0xb5, 0x3b, 0x19, 0x2c, 0xce, 0xc1, 0xbe, 0x11, 0x7a, 0x73, 0x8d, 0xab, 0xec,
	0x51, 0x07, 0x78, 0xb3, 0xf7, 0xc1, 0xb3, 0xff, 0xc5, 0x0d, 0x0f, 0x4a, 0xd4, 0x90, 0xed, 0x6f,
	0x46, 0xe5, 0x9f, 0xe6, 0x32, 0xf9, 0xcc, 0xce, 0x08, 0xbd, 0xbd, 0x3e, 0x56, 0x06, 0x1d, 0xe6,
	0xad, 0xf9, 0x07, 0x8f, 0x3b, 0xfd, 0x3b, 0xc2, 0xaf, 0x31, 0x1f, 0xee, 0xc4, 0x7c, 0x31, 0xfa,
	0xb1, 0xec, 0x93, 0x9f, 0xcb, 0x3e, 0xf9, 0xbd, 0xec, 0x93, 0xb7, 0x4f, 0x36, 0xde, 0xbb, 0xad,
	0xff, 0x8a, 0xd3, 0xab, 0xe5, 0x75, 0x1b, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xf5, 0xa8,
	0x70, 0x53, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArchivedWorkflowServiceClient is the client API for ArchivedWorkflowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArchivedWorkflowServiceClient interface {
	ListArchivedWorkflows(ctx context.Context, in *ListArchivedWorkflowsRequest, opts ...grpc.CallOption) (*v1alpha1.WorkflowList, error)
	GetArchivedWorkflow(ctx context.Context, in *GetArchivedWorkflowRequest, opts ...grpc.CallOption) (*v1alpha1.Workflow, error)
	DeleteArchivedWorkflow(ctx context.Context, in *DeleteArchivedWorkflowRequest, opts ...grpc.CallOption) (*ArchivedWorkflowDeletedResponse, error)
}

type archivedWorkflowServiceClient struct {
	cc *grpc.ClientConn
}

func NewArchivedWorkflowServiceClient(cc *grpc.ClientConn) ArchivedWorkflowServiceClient {
	return &archivedWorkflowServiceClient{cc}
}

func (c *archivedWorkflowServiceClient) ListArchivedWorkflows(ctx context.Context, in *ListArchivedWorkflowsRequest, opts ...grpc.CallOption) (*v1alpha1.WorkflowList, error) {
	out := new(v1alpha1.WorkflowList)
	err := c.cc.Invoke(ctx, "/workflowarchive.ArchivedWorkflowService/ListArchivedWorkflows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *archivedWorkflowServiceClient) GetArchivedWorkflow(ctx context.Context, in *GetArchivedWorkflowRequest, opts ...grpc.CallOption) (*v1alpha1.Workflow, error) {
	out := new(v1alpha1.Workflow)
	err := c.cc.Invoke(ctx, "/workflowarchive.ArchivedWorkflowService/GetArchivedWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *archivedWorkflowServiceClient) DeleteArchivedWorkflow(ctx context.Context, in *DeleteArchivedWorkflowRequest, opts ...grpc.CallOption) (*ArchivedWorkflowDeletedResponse, error) {
	out := new(ArchivedWorkflowDeletedResponse)
	err := c.cc.Invoke(ctx, "/workflowarchive.ArchivedWorkflowService/DeleteArchivedWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArchivedWorkflowServiceServer is the server API for ArchivedWorkflowService service.
type ArchivedWorkflowServiceServer interface {
	ListArchivedWorkflows(context.Context, *ListArchivedWorkflowsRequest) (*v1alpha1.WorkflowList, error)
	GetArchivedWorkflow(context.Context, *GetArchivedWorkflowRequest) (*v1alpha1.Workflow, error)
	DeleteArchivedWorkflow(context.Context, *DeleteArchivedWorkflowRequest) (*ArchivedWorkflowDeletedResponse, error)
}

// UnimplementedArchivedWorkflowServiceServer can be embedded to have forward compatible implementations.
type UnimplementedArchivedWorkflowServiceServer struct {
}

func (*UnimplementedArchivedWorkflowServiceServer) ListArchivedWorkflows(ctx context.Context, req *ListArchivedWorkflowsRequest) (*v1alpha1.WorkflowList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArchivedWorkflows not implemented")
}
func (*UnimplementedArchivedWorkflowServiceServer) GetArchivedWorkflow(ctx context.Context, req *GetArchivedWorkflowRequest) (*v1alpha1.Workflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArchivedWorkflow not implemented")
}
func (*UnimplementedArchivedWorkflowServiceServer) DeleteArchivedWorkflow(ctx context.Context, req *DeleteArchivedWorkflowRequest) (*ArchivedWorkflowDeletedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArchivedWorkflow not implemented")
}

func RegisterArchivedWorkflowServiceServer(s *grpc.Server, srv ArchivedWorkflowServiceServer) {
	s.RegisterService(&_ArchivedWorkflowService_serviceDesc, srv)
}

func _ArchivedWorkflowService_ListArchivedWorkflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArchivedWorkflowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchivedWorkflowServiceServer).ListArchivedWorkflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowarchive.ArchivedWorkflowService/ListArchivedWorkflows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchivedWorkflowServiceServer).ListArchivedWorkflows(ctx, req.(*ListArchivedWorkflowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArchivedWorkflowService_GetArchivedWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArchivedWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchivedWorkflowServiceServer).GetArchivedWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowarchive.ArchivedWorkflowService/GetArchivedWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchivedWorkflowServiceServer).GetArchivedWorkflow(ctx, req.(*GetArchivedWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArchivedWorkflowService_DeleteArchivedWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArchivedWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchivedWorkflowServiceServer).DeleteArchivedWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflowarchive.ArchivedWorkflowService/DeleteArchivedWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchivedWorkflowServiceServer).DeleteArchivedWorkflow(ctx, req.(*DeleteArchivedWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArchivedWorkflowService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "workflowarchive.ArchivedWorkflowService",
	HandlerType: (*ArchivedWorkflowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListArchivedWorkflows",
			Handler:    _ArchivedWorkflowService_ListArchivedWorkflows_Handler,
		},
		{
			MethodName: "GetArchivedWorkflow",
			Handler:    _ArchivedWorkflowService_GetArchivedWorkflow_Handler,
		},
		{
			MethodName: "DeleteArchivedWorkflow",
			Handler:    _ArchivedWorkflowService_DeleteArchivedWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiclient/workflowarchive/workflow-archive.proto",
}

func (m *ListArchivedWorkflowsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListArchivedWorkflowsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListArchivedWorkflowsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ListOptions != nil {
		{
			size, err := m.ListOptions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWorkflowArchive(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetArchivedWorkflowRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetArchivedWorkflowRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetArchivedWorkflowRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintWorkflowArchive(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeleteArchivedWorkflowRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteArchivedWorkflowRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeleteArchivedWorkflowRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintWorkflowArchive(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ArchivedWorkflowDeletedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArchivedWorkflowDeletedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ArchivedWorkflowDeletedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorkflowArchive(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorkflowArchive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListArchivedWorkflowsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ListOptions != nil {
		l = m.ListOptions.Size()
		n += 1 + l + sovWorkflowArchive(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetArchivedWorkflowRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovWorkflowArchive(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DeleteArchivedWorkflowRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovWorkflowArchive(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ArchivedWorkflowDeletedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWorkflowArchive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorkflowArchive(x uint64) (n int) {
	return sovWorkflowArchive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListArchivedWorkflowsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkflowArchive
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListArchivedWorkflowsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListArchivedWorkflowsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowArchive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ListOptions == nil {
				m.ListOptions = &v1.ListOptions{}
			}
			if err := m.ListOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkflowArchive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetArchivedWorkflowRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkflowArchive
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetArchivedWorkflowRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetArchivedWorkflowRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowArchive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkflowArchive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeleteArchivedWorkflowRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkflowArchive
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeleteArchivedWorkflowRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteArchivedWorkflowRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowArchive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkflowArchive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ArchivedWorkflowDeletedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkflowArchive
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ArchivedWorkflowDeletedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArchivedWorkflowDeletedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipWorkflowArchive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkflowArchive
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWorkflowArchive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkflowArchive
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorkflowArchive
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorkflowArchive
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthWorkflowArchive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWorkflowArchive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWorkflowArchive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWorkflowArchive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkflowArchive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWorkflowArchive = fmt.Errorf("proto: unexpected end of group")
)
