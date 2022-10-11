// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: user_service.proto

package __

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

// UserRegistrationClient is the client API for UserRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRegistrationClient interface {
	Registration(ctx context.Context, in *UserRegistrationRequest, opts ...grpc.CallOption) (*User, error)
}

type userRegistrationClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRegistrationClient(cc grpc.ClientConnInterface) UserRegistrationClient {
	return &userRegistrationClient{cc}
}

func (c *userRegistrationClient) Registration(ctx context.Context, in *UserRegistrationRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.UserRegistration/Registration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRegistrationServer is the server API for UserRegistration service.
// All implementations must embed UnimplementedUserRegistrationServer
// for forward compatibility
type UserRegistrationServer interface {
	Registration(context.Context, *UserRegistrationRequest) (*User, error)
	mustEmbedUnimplementedUserRegistrationServer()
}

// UnimplementedUserRegistrationServer must be embedded to have forward compatible implementations.
type UnimplementedUserRegistrationServer struct {
}

func (UnimplementedUserRegistrationServer) Registration(context.Context, *UserRegistrationRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}
func (UnimplementedUserRegistrationServer) mustEmbedUnimplementedUserRegistrationServer() {}

// UnsafeUserRegistrationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRegistrationServer will
// result in compilation errors.
type UnsafeUserRegistrationServer interface {
	mustEmbedUnimplementedUserRegistrationServer()
}

func RegisterUserRegistrationServer(s grpc.ServiceRegistrar, srv UserRegistrationServer) {
	s.RegisterService(&UserRegistration_ServiceDesc, srv)
}

func _UserRegistration_Registration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistrationServer).Registration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserRegistration/Registration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistrationServer).Registration(ctx, req.(*UserRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRegistration_ServiceDesc is the grpc.ServiceDesc for UserRegistration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRegistration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserRegistration",
	HandlerType: (*UserRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registration",
			Handler:    _UserRegistration_Registration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}