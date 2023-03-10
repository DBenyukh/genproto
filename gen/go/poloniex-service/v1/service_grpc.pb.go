// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: poloniex-service/v1/service.proto

package poloniex_service

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

// PoloniexServiceClient is the client API for PoloniexService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PoloniexServiceClient interface {
	GetRT(ctx context.Context, in *GetRtRequest, opts ...grpc.CallOption) (*GetRtResponse, error)
	GetKL(ctx context.Context, in *GetKlRequest, opts ...grpc.CallOption) (*GetKlResponse, error)
}

type poloniexServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPoloniexServiceClient(cc grpc.ClientConnInterface) PoloniexServiceClient {
	return &poloniexServiceClient{cc}
}

func (c *poloniexServiceClient) GetRT(ctx context.Context, in *GetRtRequest, opts ...grpc.CallOption) (*GetRtResponse, error) {
	out := new(GetRtResponse)
	err := c.cc.Invoke(ctx, "/poloniex.v1.PoloniexService/GetRT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poloniexServiceClient) GetKL(ctx context.Context, in *GetKlRequest, opts ...grpc.CallOption) (*GetKlResponse, error) {
	out := new(GetKlResponse)
	err := c.cc.Invoke(ctx, "/poloniex.v1.PoloniexService/GetKL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoloniexServiceServer is the server API for PoloniexService service.
// All implementations must embed UnimplementedPoloniexServiceServer
// for forward compatibility
type PoloniexServiceServer interface {
	GetRT(context.Context, *GetRtRequest) (*GetRtResponse, error)
	GetKL(context.Context, *GetKlRequest) (*GetKlResponse, error)
	mustEmbedUnimplementedPoloniexServiceServer()
}

// UnimplementedPoloniexServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPoloniexServiceServer struct {
}

func (UnimplementedPoloniexServiceServer) GetRT(context.Context, *GetRtRequest) (*GetRtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRT not implemented")
}
func (UnimplementedPoloniexServiceServer) GetKL(context.Context, *GetKlRequest) (*GetKlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKL not implemented")
}
func (UnimplementedPoloniexServiceServer) mustEmbedUnimplementedPoloniexServiceServer() {}

// UnsafePoloniexServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoloniexServiceServer will
// result in compilation errors.
type UnsafePoloniexServiceServer interface {
	mustEmbedUnimplementedPoloniexServiceServer()
}

func RegisterPoloniexServiceServer(s grpc.ServiceRegistrar, srv PoloniexServiceServer) {
	s.RegisterService(&PoloniexService_ServiceDesc, srv)
}

func _PoloniexService_GetRT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoloniexServiceServer).GetRT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poloniex.v1.PoloniexService/GetRT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoloniexServiceServer).GetRT(ctx, req.(*GetRtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoloniexService_GetKL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoloniexServiceServer).GetKL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poloniex.v1.PoloniexService/GetKL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoloniexServiceServer).GetKL(ctx, req.(*GetKlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PoloniexService_ServiceDesc is the grpc.ServiceDesc for PoloniexService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PoloniexService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "poloniex.v1.PoloniexService",
	HandlerType: (*PoloniexServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRT",
			Handler:    _PoloniexService_GetRT_Handler,
		},
		{
			MethodName: "GetKL",
			Handler:    _PoloniexService_GetKL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "poloniex-service/v1/service.proto",
}
