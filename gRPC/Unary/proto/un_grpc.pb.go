// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: proto/un.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UnaryClient is the client API for Unary service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UnaryClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type unaryClient struct {
	cc grpc.ClientConnInterface
}

func NewUnaryClient(cc grpc.ClientConnInterface) UnaryClient {
	return &unaryClient{cc}
}

func (c *unaryClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/Unary/hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnaryServer is the server API for Unary service.
// All implementations must embed UnimplementedUnaryServer
// for forward compatibility
type UnaryServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedUnaryServer()
}

// UnimplementedUnaryServer must be embedded to have forward compatible implementations.
type UnimplementedUnaryServer struct {
}

func (UnimplementedUnaryServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedUnaryServer) mustEmbedUnimplementedUnaryServer() {}

// UnsafeUnaryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UnaryServer will
// result in compilation errors.
type UnsafeUnaryServer interface {
	mustEmbedUnimplementedUnaryServer()
}

func RegisterUnaryServer(s grpc.ServiceRegistrar, srv UnaryServer) {
	s.RegisterService(&Unary_ServiceDesc, srv)
}

func _Unary_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnaryServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Unary/hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnaryServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Unary_ServiceDesc is the grpc.ServiceDesc for Unary service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Unary_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Unary",
	HandlerType: (*UnaryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "hello",
			Handler:    _Unary_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/un.proto",
}
