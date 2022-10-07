// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.2
// source: activity.proto

package services

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

// ActivityClient is the client API for Activity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityClient interface {
	CreateActivity(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Response, error)
	GetActivitys(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Activity_GetActivitysClient, error)
	GetActivity(ctx context.Context, in *ActivityId, opts ...grpc.CallOption) (*Activity, error)
	EditActivity(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Response, error)
	DeleteActivity(ctx context.Context, in *ActivityId, opts ...grpc.CallOption) (*Response, error)
}

type activityClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityClient(cc grpc.ClientConnInterface) ActivityClient {
	return &activityClient{cc}
}

func (c *activityClient) CreateActivity(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/activity.activity/CreateActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) GetActivitys(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Activity_GetActivitysClient, error) {
	stream, err := c.cc.NewStream(ctx, &Activity_ServiceDesc.Streams[0], "/activity.activity/GetActivitys", opts...)
	if err != nil {
		return nil, err
	}
	x := &activityGetActivitysClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Activity_GetActivitysClient interface {
	Recv() (*Activity, error)
	grpc.ClientStream
}

type activityGetActivitysClient struct {
	grpc.ClientStream
}

func (x *activityGetActivitysClient) Recv() (*Activity, error) {
	m := new(Activity)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *activityClient) GetActivity(ctx context.Context, in *ActivityId, opts ...grpc.CallOption) (*Activity, error) {
	out := new(Activity)
	err := c.cc.Invoke(ctx, "/activity.activity/GetActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) EditActivity(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/activity.activity/EditActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) DeleteActivity(ctx context.Context, in *ActivityId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/activity.activity/DeleteActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServer is the server API for Activity service.
// All implementations must embed UnimplementedActivityServer
// for forward compatibility
type ActivityServer interface {
	CreateActivity(context.Context, *Activity) (*Response, error)
	GetActivitys(*Empty, Activity_GetActivitysServer) error
	GetActivity(context.Context, *ActivityId) (*Activity, error)
	EditActivity(context.Context, *Activity) (*Response, error)
	DeleteActivity(context.Context, *ActivityId) (*Response, error)
	mustEmbedUnimplementedActivityServer()
}

// UnimplementedActivityServer must be embedded to have forward compatible implementations.
type UnimplementedActivityServer struct {
}

func (UnimplementedActivityServer) CreateActivity(context.Context, *Activity) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActivity not implemented")
}
func (UnimplementedActivityServer) GetActivitys(*Empty, Activity_GetActivitysServer) error {
	return status.Errorf(codes.Unimplemented, "method GetActivitys not implemented")
}
func (UnimplementedActivityServer) GetActivity(context.Context, *ActivityId) (*Activity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivity not implemented")
}
func (UnimplementedActivityServer) EditActivity(context.Context, *Activity) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditActivity not implemented")
}
func (UnimplementedActivityServer) DeleteActivity(context.Context, *ActivityId) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteActivity not implemented")
}
func (UnimplementedActivityServer) mustEmbedUnimplementedActivityServer() {}

// UnsafeActivityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServer will
// result in compilation errors.
type UnsafeActivityServer interface {
	mustEmbedUnimplementedActivityServer()
}

func RegisterActivityServer(s grpc.ServiceRegistrar, srv ActivityServer) {
	s.RegisterService(&Activity_ServiceDesc, srv)
}

func _Activity_CreateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Activity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).CreateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.activity/CreateActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).CreateActivity(ctx, req.(*Activity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_GetActivitys_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActivityServer).GetActivitys(m, &activityGetActivitysServer{stream})
}

type Activity_GetActivitysServer interface {
	Send(*Activity) error
	grpc.ServerStream
}

type activityGetActivitysServer struct {
	grpc.ServerStream
}

func (x *activityGetActivitysServer) Send(m *Activity) error {
	return x.ServerStream.SendMsg(m)
}

func _Activity_GetActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).GetActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.activity/GetActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).GetActivity(ctx, req.(*ActivityId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_EditActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Activity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).EditActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.activity/EditActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).EditActivity(ctx, req.(*Activity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_DeleteActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).DeleteActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.activity/DeleteActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).DeleteActivity(ctx, req.(*ActivityId))
	}
	return interceptor(ctx, in, info, handler)
}

// Activity_ServiceDesc is the grpc.ServiceDesc for Activity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Activity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "activity.activity",
	HandlerType: (*ActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateActivity",
			Handler:    _Activity_CreateActivity_Handler,
		},
		{
			MethodName: "GetActivity",
			Handler:    _Activity_GetActivity_Handler,
		},
		{
			MethodName: "EditActivity",
			Handler:    _Activity_EditActivity_Handler,
		},
		{
			MethodName: "DeleteActivity",
			Handler:    _Activity_DeleteActivity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetActivitys",
			Handler:       _Activity_GetActivitys_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "activity.proto",
}