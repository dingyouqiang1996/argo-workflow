// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/apiclient/info/info.proto

package info

import (
	context "context"
	fmt "fmt"
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
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

type GetInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoRequest) Reset()         { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()    {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96940c93018255fa, []int{0}
}
func (m *GetInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoRequest.Merge(m, src)
}
func (m *GetInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoRequest proto.InternalMessageInfo

type InfoResponse struct {
	ManagedNamespace     string           `protobuf:"bytes,1,opt,name=managedNamespace,proto3" json:"managedNamespace,omitempty"`
	Links                []*v1alpha1.Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	NavColor             string           `protobuf:"bytes,3,opt,name=navColor,proto3" json:"navColor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_96940c93018255fa, []int{1}
}
func (m *InfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResponse.Merge(m, src)
}
func (m *InfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *InfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResponse proto.InternalMessageInfo

func (m *InfoResponse) GetManagedNamespace() string {
	if m != nil {
		return m.ManagedNamespace
	}
	return ""
}

func (m *InfoResponse) GetLinks() []*v1alpha1.Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *InfoResponse) GetNavColor() string {
	if m != nil {
		return m.NavColor
	}
	return ""
}

type GetVersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionRequest) Reset()         { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()    {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96940c93018255fa, []int{2}
}
func (m *GetVersionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetVersionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionRequest.Merge(m, src)
}
func (m *GetVersionRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionRequest proto.InternalMessageInfo

type GetUserInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoRequest) Reset()         { *m = GetUserInfoRequest{} }
func (m *GetUserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoRequest) ProtoMessage()    {}
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96940c93018255fa, []int{3}
}
func (m *GetUserInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetUserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetUserInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetUserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoRequest.Merge(m, src)
}
func (m *GetUserInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetUserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoRequest proto.InternalMessageInfo

type GetUserInfoResponse struct {
	Issuer               string   `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Subject              string   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Groups               []string `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	EmailVerified        bool     `protobuf:"varint,5,opt,name=emailVerified,proto3" json:"emailVerified,omitempty"`
	ServiceAccountName   string   `protobuf:"bytes,6,opt,name=serviceAccountName,proto3" json:"serviceAccountName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoResponse) Reset()         { *m = GetUserInfoResponse{} }
func (m *GetUserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoResponse) ProtoMessage()    {}
func (*GetUserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_96940c93018255fa, []int{4}
}
func (m *GetUserInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetUserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetUserInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetUserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoResponse.Merge(m, src)
}
func (m *GetUserInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetUserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoResponse proto.InternalMessageInfo

func (m *GetUserInfoResponse) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *GetUserInfoResponse) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *GetUserInfoResponse) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *GetUserInfoResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GetUserInfoResponse) GetEmailVerified() bool {
	if m != nil {
		return m.EmailVerified
	}
	return false
}

func (m *GetUserInfoResponse) GetServiceAccountName() string {
	if m != nil {
		return m.ServiceAccountName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetInfoRequest)(nil), "info.GetInfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "info.InfoResponse")
	proto.RegisterType((*GetVersionRequest)(nil), "info.GetVersionRequest")
	proto.RegisterType((*GetUserInfoRequest)(nil), "info.GetUserInfoRequest")
	proto.RegisterType((*GetUserInfoResponse)(nil), "info.GetUserInfoResponse")
}

func init() { proto.RegisterFile("pkg/apiclient/info/info.proto", fileDescriptor_96940c93018255fa) }

var fileDescriptor_96940c93018255fa = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0x36, 0x2d, 0x9b, 0x52, 0xd2, 0x69, 0x44, 0x17, 0x0b, 0xa2, 0x28, 0xe2, 0x10,
	0x21, 0xb1, 0x56, 0x5a, 0xf1, 0x01, 0x50, 0x89, 0x28, 0x12, 0x70, 0x08, 0xa2, 0x07, 0x54, 0x09,
	0x6d, 0x9c, 0x89, 0xbb, 0x8d, 0xb3, 0x6b, 0x76, 0xd7, 0xee, 0x9d, 0x1b, 0x67, 0x7e, 0x85, 0x2f,
	0xe0, 0x84, 0xc4, 0x05, 0x89, 0x1f, 0x40, 0x11, 0x1f, 0x82, 0xbc, 0xb6, 0xd3, 0x04, 0x7a, 0x40,
	0xe2, 0x62, 0xcd, 0xbc, 0x19, 0xbf, 0x7d, 0x6f, 0x67, 0x87, 0x3c, 0x48, 0xe6, 0x51, 0xc0, 0x13,
	0x11, 0xc6, 0x02, 0xa5, 0x0d, 0x84, 0x9c, 0x29, 0xf7, 0x61, 0x89, 0x56, 0x56, 0xc1, 0x56, 0x1e,
	0xfb, 0xf7, 0x23, 0xa5, 0xa2, 0x18, 0xf3, 0xbe, 0x80, 0x4b, 0xa9, 0x2c, 0xb7, 0x42, 0x49, 0x53,
	0xf4, 0xf8, 0x2f, 0x23, 0x61, 0x2f, 0xd2, 0x09, 0x0b, 0xd5, 0x22, 0xe0, 0x3a, 0x52, 0x89, 0x56,
	0x97, 0x2e, 0x78, 0x7c, 0xa5, 0xf4, 0x7c, 0x16, 0xab, 0x2b, 0x13, 0x94, 0xa7, 0x98, 0xa0, 0x82,
	0x82, 0x6c, 0xc0, 0xe3, 0xe4, 0x82, 0x0f, 0x82, 0x08, 0x25, 0x6a, 0x6e, 0x71, 0x5a, 0xd0, 0xf5,
	0x5a, 0x64, 0x7f, 0x88, 0x76, 0x24, 0x67, 0x6a, 0x8c, 0xef, 0x53, 0x34, 0xb6, 0xf7, 0xd9, 0x23,
	0x7b, 0x45, 0x6e, 0x12, 0x25, 0x0d, 0xc2, 0x23, 0xd2, 0x5a, 0x70, 0xc9, 0x23, 0x9c, 0xbe, 0xe2,
	0x0b, 0x34, 0x09, 0x0f, 0x91, 0x7a, 0x5d, 0xaf, 0x7f, 0x6b, 0xfc, 0x17, 0x0e, 0xe7, 0x64, 0x3b,
	0x16, 0x72, 0x6e, 0x68, 0xad, 0x5b, 0xef, 0x37, 0x8f, 0x9f, 0xb3, 0x6b, 0xb5, 0xac, 0x52, 0xeb,
	0x82, 0x77, 0x2b, 0xb5, 0x2c, 0x3b, 0x61, 0xc9, 0x3c, 0x62, 0xb9, 0x60, 0x56, 0xa1, 0xac, 0x12,
	0xcc, 0x5e, 0x08, 0x39, 0x1f, 0x17, 0xa4, 0xe0, 0x93, 0x5d, 0xc9, 0xb3, 0x53, 0x15, 0x2b, 0x4d,
	0xeb, 0x4e, 0xc1, 0x2a, 0xef, 0x1d, 0x92, 0x83, 0x21, 0xda, 0x33, 0xd4, 0x46, 0x28, 0x59, 0x79,
	0x69, 0x13, 0x18, 0xa2, 0x7d, 0x63, 0x50, 0xaf, 0x3b, 0xfc, 0xe6, 0x91, 0xc3, 0x0d, 0xb8, 0x34,
	0x7a, 0x97, 0x34, 0x84, 0x31, 0x29, 0xea, 0xd2, 0x5e, 0x99, 0x01, 0x25, 0x3b, 0x26, 0x9d, 0x5c,
	0x62, 0x68, 0x69, 0xcd, 0x15, 0xaa, 0x34, 0xff, 0x23, 0xd2, 0x2a, 0x4d, 0x0c, 0xad, 0x77, 0xeb,
	0xf9, 0x1f, 0x45, 0x06, 0x6d, 0xb2, 0x8d, 0x0b, 0x2e, 0x62, 0xba, 0xe5, 0xfa, 0x8b, 0x04, 0x1e,
	0x92, 0xdb, 0x2e, 0x38, 0x43, 0x2d, 0x66, 0x02, 0xa7, 0x74, 0xbb, 0xeb, 0xf5, 0x77, 0xc7, 0x9b,
	0x20, 0x30, 0x02, 0x06, 0x75, 0x26, 0x42, 0x7c, 0x1a, 0x86, 0x2a, 0x95, 0x36, 0xbf, 0x5d, 0xda,
	0x70, 0x44, 0x37, 0x54, 0x8e, 0xbf, 0xd4, 0x48, 0x33, 0xb7, 0xf1, 0xba, 0x28, 0xc1, 0x88, 0xec,
	0x94, 0x13, 0x85, 0x36, 0x73, 0x8f, 0x6b, 0x73, 0xc0, 0x3e, 0x14, 0xe8, 0xba, 0xf5, 0x5e, 0xfb,
	0xc3, 0x8f, 0x5f, 0x9f, 0x6a, 0xfb, 0xb0, 0xe7, 0x5e, 0x5d, 0x36, 0x70, 0xaf, 0x12, 0x3e, 0x7a,
	0x84, 0x5c, 0x5f, 0x2a, 0x1c, 0xad, 0xe8, 0x36, 0xaf, 0xd9, 0x1f, 0xfd, 0xff, 0x98, 0x4b, 0xc6,
	0xde, 0x91, 0x13, 0x72, 0x00, 0x77, 0x2a, 0x21, 0x59, 0x79, 0xf8, 0x39, 0x69, 0xae, 0xcd, 0x0c,
	0xe8, 0x4a, 0xcb, 0x1f, 0xd3, 0xf5, 0xef, 0xdd, 0x50, 0x29, 0x5d, 0x52, 0x47, 0x0e, 0xd0, 0xaa,
	0xc8, 0x53, 0x83, 0x3a, 0xef, 0x7e, 0x76, 0xfa, 0x75, 0xd9, 0xf1, 0xbe, 0x2f, 0x3b, 0xde, 0xcf,
	0x65, 0xc7, 0x7b, 0xfb, 0xe4, 0xdf, 0x77, 0x6c, 0x6d, 0x93, 0x27, 0x0d, 0xb7, 0x52, 0x27, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xf9, 0x19, 0x4c, 0xe6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InfoServiceClient is the client API for InfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfoServiceClient interface {
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*v1alpha1.Version, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
}

type infoServiceClient struct {
	cc *grpc.ClientConn
}

func NewInfoServiceClient(cc *grpc.ClientConn) InfoServiceClient {
	return &infoServiceClient{cc}
}

func (c *infoServiceClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/info.InfoService/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*v1alpha1.Version, error) {
	out := new(v1alpha1.Version)
	err := c.cc.Invoke(ctx, "/info.InfoService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, "/info.InfoService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServiceServer is the server API for InfoService service.
type InfoServiceServer interface {
	GetInfo(context.Context, *GetInfoRequest) (*InfoResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*v1alpha1.Version, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
}

// UnimplementedInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInfoServiceServer struct {
}

func (*UnimplementedInfoServiceServer) GetInfo(ctx context.Context, req *GetInfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (*UnimplementedInfoServiceServer) GetVersion(ctx context.Context, req *GetVersionRequest) (*v1alpha1.Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (*UnimplementedInfoServiceServer) GetUserInfo(ctx context.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}

func RegisterInfoServiceServer(s *grpc.Server, srv InfoServiceServer) {
	s.RegisterService(&_InfoService_serviceDesc, srv)
}

func _InfoService_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.InfoService/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.InfoService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/info.InfoService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "info.InfoService",
	HandlerType: (*InfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _InfoService_GetInfo_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _InfoService_GetVersion_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _InfoService_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiclient/info/info.proto",
}

func (m *GetInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *InfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.NavColor) > 0 {
		i -= len(m.NavColor)
		copy(dAtA[i:], m.NavColor)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.NavColor)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Links) > 0 {
		for iNdEx := len(m.Links) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Links[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ManagedNamespace) > 0 {
		i -= len(m.ManagedNamespace)
		copy(dAtA[i:], m.ManagedNamespace)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.ManagedNamespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetVersionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetVersionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetVersionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *GetUserInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetUserInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetUserInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *GetUserInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetUserInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetUserInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ServiceAccountName) > 0 {
		i -= len(m.ServiceAccountName)
		copy(dAtA[i:], m.ServiceAccountName)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.ServiceAccountName)))
		i--
		dAtA[i] = 0x32
	}
	if m.EmailVerified {
		i--
		if m.EmailVerified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Groups) > 0 {
		for iNdEx := len(m.Groups) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Groups[iNdEx])
			copy(dAtA[i:], m.Groups[iNdEx])
			i = encodeVarintInfo(dAtA, i, uint64(len(m.Groups[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetInfoRequest) Size() (n int) {
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

func (m *InfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ManagedNamespace)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovInfo(uint64(l))
		}
	}
	l = len(m.NavColor)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetVersionRequest) Size() (n int) {
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

func (m *GetUserInfoRequest) Size() (n int) {
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

func (m *GetUserInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if len(m.Groups) > 0 {
		for _, s := range m.Groups {
			l = len(s)
			n += 1 + l + sovInfo(uint64(l))
		}
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if m.EmailVerified {
		n += 2
	}
	l = len(m.ServiceAccountName)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInfo(x uint64) (n int) {
	return sovInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
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
			return fmt.Errorf("proto: GetInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInfo
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
func (m *InfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
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
			return fmt.Errorf("proto: InfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ManagedNamespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ManagedNamespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, &v1alpha1.Link{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NavColor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NavColor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInfo
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
func (m *GetVersionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
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
			return fmt.Errorf("proto: GetVersionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetVersionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInfo
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
func (m *GetUserInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
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
			return fmt.Errorf("proto: GetUserInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetUserInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInfo
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
func (m *GetUserInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
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
			return fmt.Errorf("proto: GetUserInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetUserInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Groups = append(m.Groups, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmailVerified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EmailVerified = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceAccountName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceAccountName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInfo
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
func skipInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
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
				return 0, ErrInvalidLengthInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInfo = fmt.Errorf("proto: unexpected end of group")
)
