// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contact.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	contact.proto

It has these top-level messages:
	StringMessage
*/
package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type StringMessage struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StringMessage) Reset()                    { *m = StringMessage{} }
func (m *StringMessage) String() string            { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()               {}
func (*StringMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StringMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StringMessage)(nil), "example.StringMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for YourService service

type YourServiceClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type yourServiceClient struct {
	cc *grpc.ClientConn
}

func NewYourServiceClient(cc *grpc.ClientConn) YourServiceClient {
	return &yourServiceClient{cc}
}

func (c *yourServiceClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := grpc.Invoke(ctx, "/example.YourService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for YourService service

type YourServiceServer interface {
	Echo(context.Context, *StringMessage) (*StringMessage, error)
}

func RegisterYourServiceServer(s *grpc.Server, srv YourServiceServer) {
	s.RegisterService(&_YourService_serviceDesc, srv)
}

func _YourService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YourServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.YourService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YourServiceServer).Echo(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _YourService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.YourService",
	HandlerType: (*YourServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _YourService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact.proto",
}

func init() { proto.RegisterFile("contact.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0x2b,
	0x49, 0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xad, 0x48, 0xcc, 0x2d,
	0xc8, 0x49, 0x95, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f,
	0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x53, 0x52, 0xe5, 0xe2,
	0x0d, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe1,
	0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x8c,
	0x92, 0xb8, 0xb8, 0x23, 0xf3, 0x4b, 0x8b, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x82,
	0xb9, 0x58, 0x5c, 0x93, 0x33, 0xf2, 0x85, 0xc4, 0xf4, 0xa0, 0xb6, 0xe8, 0xa1, 0x18, 0x22, 0x85,
	0x43, 0x5c, 0x49, 0xba, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xa2, 0x4a, 0x02, 0xfa, 0x65, 0x86, 0xfa,
	0x50, 0x25, 0xfa, 0xa9, 0xc9, 0x19, 0xf9, 0x56, 0x8c, 0x5a, 0x49, 0x6c, 0x60, 0x17, 0x19, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x8a, 0xfd, 0xcd, 0xc9, 0x00, 0x00, 0x00,
}
