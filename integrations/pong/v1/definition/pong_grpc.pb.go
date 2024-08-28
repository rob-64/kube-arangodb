// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: integrations/pong/v1/definition/pong.proto

package definition

import (
	context "context"
	definition "github.com/arangodb/kube-arangodb/integrations/shared/v1/definition"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PongV1Client is the client API for PongV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PongV1Client interface {
	Ping(ctx context.Context, in *definition.Empty, opts ...grpc.CallOption) (*PongV1PingResponse, error)
}

type pongV1Client struct {
	cc grpc.ClientConnInterface
}

func NewPongV1Client(cc grpc.ClientConnInterface) PongV1Client {
	return &pongV1Client{cc}
}

func (c *pongV1Client) Ping(ctx context.Context, in *definition.Empty, opts ...grpc.CallOption) (*PongV1PingResponse, error) {
	out := new(PongV1PingResponse)
	err := c.cc.Invoke(ctx, "/pong.PongV1/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PongV1Server is the server API for PongV1 service.
// All implementations must embed UnimplementedPongV1Server
// for forward compatibility
type PongV1Server interface {
	Ping(context.Context, *definition.Empty) (*PongV1PingResponse, error)
	mustEmbedUnimplementedPongV1Server()
}

// UnimplementedPongV1Server must be embedded to have forward compatible implementations.
type UnimplementedPongV1Server struct {
}

func (UnimplementedPongV1Server) Ping(context.Context, *definition.Empty) (*PongV1PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedPongV1Server) mustEmbedUnimplementedPongV1Server() {}

// UnsafePongV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PongV1Server will
// result in compilation errors.
type UnsafePongV1Server interface {
	mustEmbedUnimplementedPongV1Server()
}

func RegisterPongV1Server(s grpc.ServiceRegistrar, srv PongV1Server) {
	s.RegisterService(&PongV1_ServiceDesc, srv)
}

func _PongV1_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(definition.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PongV1Server).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pong.PongV1/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PongV1Server).Ping(ctx, req.(*definition.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PongV1_ServiceDesc is the grpc.ServiceDesc for PongV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PongV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pong.PongV1",
	HandlerType: (*PongV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PongV1_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "integrations/pong/v1/definition/pong.proto",
}
