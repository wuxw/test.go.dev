// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userfind.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
	_ "test.go.dev/grpc-hello-world/proto/google/api"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserFindRequest struct {
	Referer              string   `protobuf:"bytes,1,opt,name=referer,proto3" json:"referer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserFindRequest) Reset()         { *m = UserFindRequest{} }
func (m *UserFindRequest) String() string { return proto.CompactTextString(m) }
func (*UserFindRequest) ProtoMessage()    {}
func (*UserFindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f86da53a2dc07a38, []int{0}
}

func (m *UserFindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserFindRequest.Unmarshal(m, b)
}
func (m *UserFindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserFindRequest.Marshal(b, m, deterministic)
}
func (m *UserFindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFindRequest.Merge(m, src)
}
func (m *UserFindRequest) XXX_Size() int {
	return xxx_messageInfo_UserFindRequest.Size(m)
}
func (m *UserFindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserFindRequest proto.InternalMessageInfo

func (m *UserFindRequest) GetReferer() string {
	if m != nil {
		return m.Referer
	}
	return ""
}

type UserFindResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserFindResponse) Reset()         { *m = UserFindResponse{} }
func (m *UserFindResponse) String() string { return proto.CompactTextString(m) }
func (*UserFindResponse) ProtoMessage()    {}
func (*UserFindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f86da53a2dc07a38, []int{1}
}

func (m *UserFindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserFindResponse.Unmarshal(m, b)
}
func (m *UserFindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserFindResponse.Marshal(b, m, deterministic)
}
func (m *UserFindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFindResponse.Merge(m, src)
}
func (m *UserFindResponse) XXX_Size() int {
	return xxx_messageInfo_UserFindResponse.Size(m)
}
func (m *UserFindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserFindResponse proto.InternalMessageInfo

func (m *UserFindResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*UserFindRequest)(nil), "proto.UserFindRequest")
	proto.RegisterType((*UserFindResponse)(nil), "proto.UserFindResponse")
}

func init() { proto.RegisterFile("userfind.proto", fileDescriptor_f86da53a2dc07a38) }

var fileDescriptor_f86da53a2dc07a38 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2d, 0x4e, 0x2d,
	0x4a, 0xcb, 0xcc, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x32,
	0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25,
	0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x45, 0x4a, 0xda, 0x5c, 0xfc, 0xa1, 0xc5, 0xa9, 0x45,
	0x6e, 0x99, 0x79, 0x29, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x12, 0x5c, 0xec, 0x45,
	0xa9, 0x69, 0xa9, 0x45, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x92,
	0x0e, 0x97, 0x00, 0x42, 0x71, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x48, 0x75, 0x6e, 0x6a, 0x71,
	0x71, 0x62, 0x7a, 0x2a, 0x4c, 0x35, 0x94, 0x6b, 0x94, 0xc8, 0xc5, 0x01, 0x53, 0x2d, 0x14, 0xca,
	0xc5, 0x1d, 0x9c, 0x58, 0x09, 0xe7, 0x8a, 0x41, 0x6c, 0xd7, 0x43, 0xb3, 0x5a, 0x4a, 0x1c, 0x43,
	0x1c, 0x62, 0x8b, 0x92, 0x68, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xf8, 0x85, 0xb8, 0xf4, 0x41, 0x9e,
	0x8c, 0x07, 0xf9, 0xd2, 0x8a, 0x51, 0x2b, 0x89, 0x0d, 0xac, 0xdc, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xa4, 0x18, 0x8f, 0xa9, 0xfb, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserFindClient is the client API for UserFind service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserFindClient interface {
	SayUserFind(ctx context.Context, in *UserFindRequest, opts ...grpc.CallOption) (*UserFindResponse, error)
}

type userFindClient struct {
	cc *grpc.ClientConn
}

func NewUserFindClient(cc *grpc.ClientConn) UserFindClient {
	return &userFindClient{cc}
}

func (c *userFindClient) SayUserFind(ctx context.Context, in *UserFindRequest, opts ...grpc.CallOption) (*UserFindResponse, error) {
	out := new(UserFindResponse)
	err := c.cc.Invoke(ctx, "/proto.UserFind/SayUserFind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserFindServer is the server API for UserFind service.
type UserFindServer interface {
	SayUserFind(context.Context, *UserFindRequest) (*UserFindResponse, error)
}

func RegisterUserFindServer(s *grpc.Server, srv UserFindServer) {
	s.RegisterService(&_UserFind_serviceDesc, srv)
}

func _UserFind_SayUserFind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFindServer).SayUserFind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserFind/SayUserFind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFindServer).SayUserFind(ctx, req.(*UserFindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserFind_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserFind",
	HandlerType: (*UserFindServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayUserFind",
			Handler:    _UserFind_SayUserFind_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userfind.proto",
}