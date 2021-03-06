// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sub_protoc/hello_world.proto

package helloworld

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa7fb5dc084ce51f, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa7fb5dc084ce51f, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

func init() { proto.RegisterFile("sub_protoc/hello_world.proto", fileDescriptor_fa7fb5dc084ce51f) }

var fileDescriptor_fa7fb5dc084ce51f = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4b, 0xc5, 0x30,
	0x10, 0xc6, 0x2d, 0x88, 0x4f, 0x0f, 0x41, 0xc8, 0x20, 0x45, 0x1d, 0x24, 0x83, 0x38, 0x25, 0xa0,
	0xb8, 0x3a, 0xbc, 0x45, 0xdd, 0x1e, 0x75, 0x10, 0x5c, 0x4a, 0x5a, 0x8f, 0x53, 0xb8, 0xf6, 0x62,
	0xd2, 0xa2, 0xfd, 0xef, 0x25, 0xc1, 0xd2, 0x0e, 0x6f, 0xfb, 0xee, 0xcb, 0xef, 0xc7, 0x91, 0x83,
	0xab, 0x38, 0x36, 0xb5, 0x0f, 0x32, 0x48, 0x6b, 0x3f, 0x91, 0x59, 0xea, 0x1f, 0x09, 0xfc, 0x61,
	0x72, 0xa5, 0x20, 0x57, 0xb9, 0xd1, 0x1a, 0x4e, 0x9f, 0xd3, 0x54, 0xe1, 0xf7, 0x88, 0x71, 0x50,
	0x0a, 0x0e, 0x7b, 0xd7, 0x61, 0x59, 0x5c, 0x17, 0xb7, 0x27, 0x55, 0xce, 0xfa, 0x06, 0xe0, 0x9f,
	0xf1, 0x3c, 0xa9, 0x12, 0x36, 0x1d, 0xc6, 0xe8, 0x68, 0x86, 0xe6, 0xf1, 0xee, 0x05, 0x36, 0x4f,
	0x01, 0x71, 0xc0, 0xa0, 0x1e, 0xe1, 0xf8, 0xd5, 0x4d, 0xd9, 0x52, 0xa5, 0x59, 0xf6, 0x99, 0xf5,
	0xb2, 0x8b, 0xf3, 0x3d, 0x2f, 0x9e, 0x27, 0x7d, 0xb0, 0x25, 0xb8, 0xfc, 0x12, 0x43, 0xc1, 0xb7,
	0x06, 0x7f, 0x5d, 0xe7, 0x19, 0xe3, 0x8a, 0xdd, 0x9e, 0x65, 0xf8, 0x2d, 0xe5, 0x5d, 0xfa, 0xd2,
	0xae, 0x78, 0x7f, 0x20, 0x11, 0x62, 0x34, 0x24, 0xec, 0x7a, 0x32, 0x12, 0xc8, 0x26, 0xdd, 0xce,
	0xba, 0x5d, 0xf4, 0x55, 0x6c, 0x8e, 0xf2, 0x49, 0xee, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8f,
	0x53, 0xb3, 0xf4, 0x32, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sub_protoc/hello_world.proto",
}
