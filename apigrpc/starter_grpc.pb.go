// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apigrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GoStarterClient is the client API for GoStarter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoStarterClient interface {
	// Health checking that determines whether backend instance responds properly.
	Healthcheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Returns list of users
	FindUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FindUsersResponse, error)
	// Creates a new user
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// Returns a single user by ID.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type goStarterClient struct {
	cc grpc.ClientConnInterface
}

func NewGoStarterClient(cc grpc.ClientConnInterface) GoStarterClient {
	return &goStarterClient{cc}
}

func (c *goStarterClient) Healthcheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/starter.apigrpc.GoStarter/Healthcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goStarterClient) FindUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FindUsersResponse, error) {
	out := new(FindUsersResponse)
	err := c.cc.Invoke(ctx, "/starter.apigrpc.GoStarter/FindUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goStarterClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/starter.apigrpc.GoStarter/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goStarterClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/starter.apigrpc.GoStarter/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoStarterServer is the server API for GoStarter service.
// All implementations must embed UnimplementedGoStarterServer
// for forward compatibility
type GoStarterServer interface {
	// Health checking that determines whether backend instance responds properly.
	Healthcheck(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// Returns list of users
	FindUsers(context.Context, *emptypb.Empty) (*FindUsersResponse, error)
	// Creates a new user
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// Returns a single user by ID.
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	mustEmbedUnimplementedGoStarterServer()
}

// UnimplementedGoStarterServer must be embedded to have forward compatible implementations.
type UnimplementedGoStarterServer struct {
}

func (UnimplementedGoStarterServer) Healthcheck(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthcheck not implemented")
}
func (UnimplementedGoStarterServer) FindUsers(context.Context, *emptypb.Empty) (*FindUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUsers not implemented")
}
func (UnimplementedGoStarterServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedGoStarterServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedGoStarterServer) mustEmbedUnimplementedGoStarterServer() {}

// UnsafeGoStarterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoStarterServer will
// result in compilation errors.
type UnsafeGoStarterServer interface {
	mustEmbedUnimplementedGoStarterServer()
}

func RegisterGoStarterServer(s grpc.ServiceRegistrar, srv GoStarterServer) {
	s.RegisterService(&GoStarter_ServiceDesc, srv)
}

func _GoStarter_Healthcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoStarterServer).Healthcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starter.apigrpc.GoStarter/Healthcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoStarterServer).Healthcheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoStarter_FindUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoStarterServer).FindUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starter.apigrpc.GoStarter/FindUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoStarterServer).FindUsers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoStarter_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoStarterServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starter.apigrpc.GoStarter/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoStarterServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoStarter_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoStarterServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starter.apigrpc.GoStarter/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoStarterServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoStarter_ServiceDesc is the grpc.ServiceDesc for GoStarter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoStarter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "starter.apigrpc.GoStarter",
	HandlerType: (*GoStarterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthcheck",
			Handler:    _GoStarter_Healthcheck_Handler,
		},
		{
			MethodName: "FindUsers",
			Handler:    _GoStarter_FindUsers_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _GoStarter_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _GoStarter_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "starter.proto",
}
