// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeter.proto

/*
Package greeter is a generated protocol buffer package.

It is generated from these files:
	greeter.proto

It has these top-level messages:
	HelloReq
	HelloRes
*/
package greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type HelloReq struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Count int32  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *HelloReq) Reset()                    { *m = HelloReq{} }
func (m *HelloReq) String() string            { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()               {}
func (*HelloReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloReq) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type HelloRes struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloRes) Reset()                    { *m = HelloRes{} }
func (m *HelloRes) String() string            { return proto.CompactTextString(m) }
func (*HelloRes) ProtoMessage()               {}
func (*HelloRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloRes) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloReq)(nil), "greeter.HelloReq")
	proto.RegisterType((*HelloRes)(nil), "greeter.HelloRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRes, error)
	SayHellos(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (Greeter_SayHellosClient, error)
	GreetMany(ctx context.Context, opts ...grpc.CallOption) (Greeter_GreetManyClient, error)
	GreetChat(ctx context.Context, opts ...grpc.CallOption) (Greeter_GreetChatClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRes, error) {
	out := new(HelloRes)
	err := grpc.Invoke(ctx, "/greeter.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHellos(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (Greeter_SayHellosClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Greeter_serviceDesc.Streams[0], c.cc, "/greeter.Greeter/SayHellos", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHellosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_SayHellosClient interface {
	Recv() (*HelloRes, error)
	grpc.ClientStream
}

type greeterSayHellosClient struct {
	grpc.ClientStream
}

func (x *greeterSayHellosClient) Recv() (*HelloRes, error) {
	m := new(HelloRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) GreetMany(ctx context.Context, opts ...grpc.CallOption) (Greeter_GreetManyClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Greeter_serviceDesc.Streams[1], c.cc, "/greeter.Greeter/GreetMany", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterGreetManyClient{stream}
	return x, nil
}

type Greeter_GreetManyClient interface {
	Send(*HelloReq) error
	CloseAndRecv() (*HelloRes, error)
	grpc.ClientStream
}

type greeterGreetManyClient struct {
	grpc.ClientStream
}

func (x *greeterGreetManyClient) Send(m *HelloReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterGreetManyClient) CloseAndRecv() (*HelloRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) GreetChat(ctx context.Context, opts ...grpc.CallOption) (Greeter_GreetChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Greeter_serviceDesc.Streams[2], c.cc, "/greeter.Greeter/GreetChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterGreetChatClient{stream}
	return x, nil
}

type Greeter_GreetChatClient interface {
	Send(*HelloReq) error
	Recv() (*HelloRes, error)
	grpc.ClientStream
}

type greeterGreetChatClient struct {
	grpc.ClientStream
}

func (x *greeterGreetChatClient) Send(m *HelloReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterGreetChatClient) Recv() (*HelloRes, error) {
	m := new(HelloRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SayHello(context.Context, *HelloReq) (*HelloRes, error)
	SayHellos(*HelloReq, Greeter_SayHellosServer) error
	GreetMany(Greeter_GreetManyServer) error
	GreetChat(Greeter_GreetChatServer) error
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHellos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).SayHellos(m, &greeterSayHellosServer{stream})
}

type Greeter_SayHellosServer interface {
	Send(*HelloRes) error
	grpc.ServerStream
}

type greeterSayHellosServer struct {
	grpc.ServerStream
}

func (x *greeterSayHellosServer) Send(m *HelloRes) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_GreetMany_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).GreetMany(&greeterGreetManyServer{stream})
}

type Greeter_GreetManyServer interface {
	SendAndClose(*HelloRes) error
	Recv() (*HelloReq, error)
	grpc.ServerStream
}

type greeterGreetManyServer struct {
	grpc.ServerStream
}

func (x *greeterGreetManyServer) SendAndClose(m *HelloRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterGreetManyServer) Recv() (*HelloReq, error) {
	m := new(HelloReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_GreetChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).GreetChat(&greeterGreetChatServer{stream})
}

type Greeter_GreetChatServer interface {
	Send(*HelloRes) error
	Recv() (*HelloReq, error)
	grpc.ServerStream
}

type greeterGreetChatServer struct {
	grpc.ServerStream
}

func (x *greeterGreetChatServer) Send(m *HelloRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterGreetChatServer) Recv() (*HelloReq, error) {
	m := new(HelloReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greeter.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHellos",
			Handler:       _Greeter_SayHellos_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GreetMany",
			Handler:       _Greeter_GreetMany_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GreetChat",
			Handler:       _Greeter_GreetChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greeter.proto",
}

func init() { proto.RegisterFile("greeter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2f, 0x4a, 0x4d,
	0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x4c, 0xb8,
	0x38, 0x3c, 0x52, 0x73, 0x72, 0xf2, 0x83, 0x52, 0x0b, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xe4, 0xfc,
	0xd2, 0xbc, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x08, 0x47, 0x49, 0x05, 0xae, 0xab,
	0x58, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d, 0xa6, 0x11, 0xc6, 0x35, 0x7a,
	0xc2, 0xc8, 0xc5, 0xee, 0x0e, 0xb1, 0x47, 0xc8, 0x88, 0x8b, 0x23, 0x38, 0xb1, 0x12, 0xac, 0x49,
	0x48, 0x50, 0x0f, 0xe6, 0x18, 0x98, 0xd5, 0x52, 0x18, 0x42, 0xc5, 0x4a, 0x0c, 0x42, 0xa6, 0x5c,
	0x9c, 0x30, 0x3d, 0xc5, 0xc4, 0x6a, 0x32, 0x60, 0x04, 0x69, 0x03, 0xdb, 0xea, 0x9b, 0x98, 0x57,
	0x49, 0xac, 0x36, 0x0d, 0x46, 0x21, 0x73, 0xa8, 0x36, 0xe7, 0x8c, 0xc4, 0x12, 0xe2, 0xb5, 0x19,
	0x30, 0x3a, 0xa9, 0x70, 0xf1, 0xa4, 0x15, 0x97, 0xe4, 0xc3, 0xa4, 0x9d, 0x78, 0xa0, 0x7e, 0x0e,
	0x00, 0x85, 0x74, 0x00, 0xe3, 0x22, 0x26, 0x66, 0xf7, 0xa0, 0x90, 0x24, 0x36, 0x70, 0xc0, 0x1b,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x73, 0xda, 0xca, 0x65, 0x89, 0x01, 0x00, 0x00,
}
