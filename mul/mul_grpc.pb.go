// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mul.proto

package mul

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

// MulServiceClient is the client API for MulService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MulServiceClient interface {
	Mul2Numbers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type mulServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMulServiceClient(cc grpc.ClientConnInterface) MulServiceClient {
	return &mulServiceClient{cc}
}

func (c *mulServiceClient) Mul2Numbers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mul.MulService/Mul2Numbers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MulServiceServer is the server API for MulService service.
// All implementations should embed UnimplementedMulServiceServer
// for forward compatibility
type MulServiceServer interface {
	Mul2Numbers(context.Context, *Request) (*Response, error)
}

// UnimplementedMulServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMulServiceServer struct {
}

func (UnimplementedMulServiceServer) Mul2Numbers(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mul2Numbers not implemented")
}

// UnsafeMulServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MulServiceServer will
// result in compilation errors.
type UnsafeMulServiceServer interface {
	mustEmbedUnimplementedMulServiceServer()
}

func RegisterMulServiceServer(s grpc.ServiceRegistrar, srv MulServiceServer) {
	s.RegisterService(&MulService_ServiceDesc, srv)
}

func _MulService_Mul2Numbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MulServiceServer).Mul2Numbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mul.MulService/Mul2Numbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MulServiceServer).Mul2Numbers(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// MulService_ServiceDesc is the grpc.ServiceDesc for MulService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MulService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mul.MulService",
	HandlerType: (*MulServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mul2Numbers",
			Handler:    _MulService_Mul2Numbers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mul.proto",
}
