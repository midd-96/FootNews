// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service_footnews.proto

package pb

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

// FootNewsClient is the client API for FootNews service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FootNewsClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserReponse, error)
}

type footNewsClient struct {
	cc grpc.ClientConnInterface
}

func NewFootNewsClient(cc grpc.ClientConnInterface) FootNewsClient {
	return &footNewsClient{cc}
}

func (c *footNewsClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReponse, error) {
	out := new(CreateUserReponse)
	err := c.cc.Invoke(ctx, "/pb.footNews/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *footNewsClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserReponse, error) {
	out := new(LoginUserReponse)
	err := c.cc.Invoke(ctx, "/pb.footNews/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FootNewsServer is the server API for FootNews service.
// All implementations must embed UnimplementedFootNewsServer
// for forward compatibility
type FootNewsServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserReponse, error)
	mustEmbedUnimplementedFootNewsServer()
}

// UnimplementedFootNewsServer must be embedded to have forward compatible implementations.
type UnimplementedFootNewsServer struct {
}

func (UnimplementedFootNewsServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedFootNewsServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedFootNewsServer) mustEmbedUnimplementedFootNewsServer() {}

// UnsafeFootNewsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FootNewsServer will
// result in compilation errors.
type UnsafeFootNewsServer interface {
	mustEmbedUnimplementedFootNewsServer()
}

func RegisterFootNewsServer(s grpc.ServiceRegistrar, srv FootNewsServer) {
	s.RegisterService(&FootNews_ServiceDesc, srv)
}

func _FootNews_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FootNewsServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.footNews/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FootNewsServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FootNews_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FootNewsServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.footNews/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FootNewsServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FootNews_ServiceDesc is the grpc.ServiceDesc for FootNews service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FootNews_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.footNews",
	HandlerType: (*FootNewsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _FootNews_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _FootNews_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_footnews.proto",
}
