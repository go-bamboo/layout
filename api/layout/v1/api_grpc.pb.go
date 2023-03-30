// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: api/layout/v1/api.proto

package v1

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

// QuantBotV1Client is the client API for QuantBotV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuantBotV1Client interface {
	ChainBotWebHook(ctx context.Context, in *ChainBotWebHookRequest, opts ...grpc.CallOption) (*ChainBotWebHookReply, error)
}

type quantBotV1Client struct {
	cc grpc.ClientConnInterface
}

func NewQuantBotV1Client(cc grpc.ClientConnInterface) QuantBotV1Client {
	return &quantBotV1Client{cc}
}

func (c *quantBotV1Client) ChainBotWebHook(ctx context.Context, in *ChainBotWebHookRequest, opts ...grpc.CallOption) (*ChainBotWebHookReply, error) {
	out := new(ChainBotWebHookReply)
	err := c.cc.Invoke(ctx, "/api.layout.v1.QuantBotV1/ChainBotWebHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuantBotV1Server is the server API for QuantBotV1 service.
// All implementations must embed UnimplementedQuantBotV1Server
// for forward compatibility
type QuantBotV1Server interface {
	ChainBotWebHook(context.Context, *ChainBotWebHookRequest) (*ChainBotWebHookReply, error)
	mustEmbedUnimplementedQuantBotV1Server()
}

// UnimplementedQuantBotV1Server must be embedded to have forward compatible implementations.
type UnimplementedQuantBotV1Server struct {
}

func (UnimplementedQuantBotV1Server) ChainBotWebHook(context.Context, *ChainBotWebHookRequest) (*ChainBotWebHookReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainBotWebHook not implemented")
}
func (UnimplementedQuantBotV1Server) mustEmbedUnimplementedQuantBotV1Server() {}

// UnsafeQuantBotV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuantBotV1Server will
// result in compilation errors.
type UnsafeQuantBotV1Server interface {
	mustEmbedUnimplementedQuantBotV1Server()
}

func RegisterQuantBotV1Server(s grpc.ServiceRegistrar, srv QuantBotV1Server) {
	s.RegisterService(&QuantBotV1_ServiceDesc, srv)
}

func _QuantBotV1_ChainBotWebHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainBotWebHookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuantBotV1Server).ChainBotWebHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.layout.v1.QuantBotV1/ChainBotWebHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuantBotV1Server).ChainBotWebHook(ctx, req.(*ChainBotWebHookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuantBotV1_ServiceDesc is the grpc.ServiceDesc for QuantBotV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuantBotV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.layout.v1.QuantBotV1",
	HandlerType: (*QuantBotV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChainBotWebHook",
			Handler:    _QuantBotV1_ChainBotWebHook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/layout/v1/api.proto",
}
