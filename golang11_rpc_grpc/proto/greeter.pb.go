// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/greeter.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HelloRequestMessage struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequestMessage) Reset()         { *m = HelloRequestMessage{} }
func (m *HelloRequestMessage) String() string { return proto.CompactTextString(m) }
func (*HelloRequestMessage) ProtoMessage()    {}
func (*HelloRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2d8b56b5ea0dd4, []int{0}
}

func (m *HelloRequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequestMessage.Unmarshal(m, b)
}
func (m *HelloRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequestMessage.Marshal(b, m, deterministic)
}
func (m *HelloRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequestMessage.Merge(m, src)
}
func (m *HelloRequestMessage) XXX_Size() int {
	return xxx_messageInfo_HelloRequestMessage.Size(m)
}
func (m *HelloRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequestMessage proto.InternalMessageInfo

func (m *HelloRequestMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponseMessage struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponseMessage) Reset()         { *m = HelloResponseMessage{} }
func (m *HelloResponseMessage) String() string { return proto.CompactTextString(m) }
func (*HelloResponseMessage) ProtoMessage()    {}
func (*HelloResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2d8b56b5ea0dd4, []int{1}
}

func (m *HelloResponseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponseMessage.Unmarshal(m, b)
}
func (m *HelloResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponseMessage.Marshal(b, m, deterministic)
}
func (m *HelloResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponseMessage.Merge(m, src)
}
func (m *HelloResponseMessage) XXX_Size() int {
	return xxx_messageInfo_HelloResponseMessage.Size(m)
}
func (m *HelloResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponseMessage proto.InternalMessageInfo

func (m *HelloResponseMessage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *HelloResponseMessage) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HelloResponseMessage) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequestMessage)(nil), "protocol.HelloRequestMessage")
	proto.RegisterType((*HelloResponseMessage)(nil), "protocol.HelloResponseMessage")
}

func init() { proto.RegisterFile("proto/greeter.proto", fileDescriptor_8a2d8b56b5ea0dd4) }

var fileDescriptor_8a2d8b56b5ea0dd4 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0xd2, 0x03, 0xf3, 0x84, 0x38, 0xc0, 0x54,
	0x72, 0x7e, 0x8e, 0x92, 0x26, 0x97, 0xb0, 0x47, 0x6a, 0x4e, 0x4e, 0x7e, 0x50, 0x6a, 0x61, 0x69,
	0x6a, 0x71, 0x89, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62,
	0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0x14, 0xc0, 0x25, 0x02, 0x55,
	0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x0a, 0x53, 0x2b, 0xc0, 0xc5, 0x9c, 0x5b, 0x9c, 0x0e, 0x55,
	0x0a, 0x62, 0x82, 0x74, 0x27, 0xe7, 0xa7, 0xa4, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x06, 0x81,
	0xd9, 0x20, 0xb1, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x66, 0x88, 0x89, 0x20, 0xb6, 0x51, 0x14, 0x17,
	0xbb, 0x3b, 0xc4, 0x5d, 0x42, 0xfe, 0x5c, 0x3c, 0x60, 0xc3, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93,
	0x53, 0x85, 0x64, 0xf5, 0x60, 0x4e, 0xd4, 0xc3, 0xe2, 0x3e, 0x29, 0x39, 0x0c, 0x69, 0x14, 0x37,
	0x29, 0x31, 0x24, 0xb1, 0x81, 0x15, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x30, 0x2f, 0x96,
	0x75, 0x00, 0x01, 0x00, 0x00,
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
	Helloservice(ctx context.Context, in *HelloRequestMessage, opts ...grpc.CallOption) (*HelloResponseMessage, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Helloservice(ctx context.Context, in *HelloRequestMessage, opts ...grpc.CallOption) (*HelloResponseMessage, error) {
	out := new(HelloResponseMessage)
	err := c.cc.Invoke(ctx, "/protocol.Greeter/Helloservice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	Helloservice(context.Context, *HelloRequestMessage) (*HelloResponseMessage, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) Helloservice(ctx context.Context, req *HelloRequestMessage) (*HelloResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Helloservice not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Helloservice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Helloservice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Greeter/Helloservice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Helloservice(ctx, req.(*HelloRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Helloservice",
			Handler:    _Greeter_Helloservice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/greeter.proto",
}