// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// AppPositionClient is the client API for AppPosition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppPositionClient interface {
	AppTopCategory(ctx context.Context, in *Date, opts ...grpc.CallOption) (*CategoryToMaxPosition, error)
}

type appPositionClient struct {
	cc grpc.ClientConnInterface
}

func NewAppPositionClient(cc grpc.ClientConnInterface) AppPositionClient {
	return &appPositionClient{cc}
}

func (c *appPositionClient) AppTopCategory(ctx context.Context, in *Date, opts ...grpc.CallOption) (*CategoryToMaxPosition, error) {
	out := new(CategoryToMaxPosition)
	err := c.cc.Invoke(ctx, "/grpc.AppPosition/AppTopCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppPositionServer is the server API for AppPosition service.
// All implementations must embed UnimplementedAppPositionServer
// for forward compatibility
type AppPositionServer interface {
	AppTopCategory(context.Context, *Date) (*CategoryToMaxPosition, error)
	mustEmbedUnimplementedAppPositionServer()
}

// UnimplementedAppPositionServer must be embedded to have forward compatible implementations.
type UnimplementedAppPositionServer struct {
}

func (UnimplementedAppPositionServer) AppTopCategory(context.Context, *Date) (*CategoryToMaxPosition, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppTopCategory not implemented")
}
func (UnimplementedAppPositionServer) mustEmbedUnimplementedAppPositionServer() {}

// UnsafeAppPositionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppPositionServer will
// result in compilation errors.
type UnsafeAppPositionServer interface {
	mustEmbedUnimplementedAppPositionServer()
}

func RegisterAppPositionServer(s grpc.ServiceRegistrar, srv AppPositionServer) {
	s.RegisterService(&AppPosition_ServiceDesc, srv)
}

func _AppPosition_AppTopCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Date)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppPositionServer).AppTopCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AppPosition/AppTopCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppPositionServer).AppTopCategory(ctx, req.(*Date))
	}
	return interceptor(ctx, in, info, handler)
}

// AppPosition_ServiceDesc is the grpc.ServiceDesc for AppPosition service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppPosition_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.AppPosition",
	HandlerType: (*AppPositionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppTopCategory",
			Handler:    _AppPosition_AppTopCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app_position.proto",
}