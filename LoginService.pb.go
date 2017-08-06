// Code generated by protoc-gen-go. DO NOT EDIT.
// source: LoginService.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	LoginService.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
*/
package protos

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

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "protos.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "protos.LoginResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LoginService service

type LoginServiceClient interface {
	LoginAuthentication(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type loginServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoginServiceClient(cc *grpc.ClientConn) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) LoginAuthentication(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/protos.LoginService/LoginAuthentication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoginService service

type LoginServiceServer interface {
	LoginAuthentication(context.Context, *LoginRequest) (*LoginResponse, error)
}

func RegisterLoginServiceServer(s *grpc.Server, srv LoginServiceServer) {
	s.RegisterService(&_LoginService_serviceDesc, srv)
}

func _LoginService_LoginAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).LoginAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.LoginService/LoginAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).LoginAuthentication(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginAuthentication",
			Handler:    _LoginService_LoginAuthentication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "LoginService.proto",
}

func init() { proto.RegisterFile("LoginService.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xf2, 0xc9, 0x4f, 0xcf,
	0xcc, 0x0b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x03, 0x53, 0xc5, 0x4a, 0x6e, 0x5c, 0x3c, 0x60, 0xd9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x21, 0x29, 0x2e, 0x8e, 0xd2, 0xe2, 0xd4, 0xa2, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x38, 0x1f, 0x24, 0x57, 0x90, 0x58, 0x5c, 0x5c, 0x9e, 0x5f, 0x94, 0x22, 0xc1,
	0x04, 0x91, 0x83, 0xf1, 0x95, 0xd4, 0xb9, 0x78, 0xa1, 0xe6, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x0a, 0x89, 0x71, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16, 0x43, 0x8d, 0x81, 0xf2, 0x8c, 0x42,
	0xa0, 0x16, 0x42, 0x9d, 0x23, 0xe4, 0xc2, 0x25, 0x0c, 0xe6, 0x3b, 0x96, 0x96, 0x64, 0xa4, 0xe6,
	0x95, 0x64, 0x26, 0x27, 0x96, 0x64, 0xe6, 0xe7, 0x09, 0x89, 0x40, 0xdc, 0x59, 0xac, 0x87, 0xec,
	0x3a, 0x29, 0x51, 0x34, 0x51, 0x88, 0x5d, 0x4a, 0x0c, 0x49, 0x10, 0xef, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xfe, 0x83, 0x66, 0x50, 0xeb, 0x00, 0x00, 0x00,
}
