// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.9.0
// source: rpc/user.proto

package user_rpc

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

const (
	User_UserCreate_FullMethodName     = "/user_rpc.User/UserCreate"
	User_UserInfo_FullMethodName       = "/user_rpc.User/UserInfo"
	User_UserBaseInfo_FullMethodName   = "/user_rpc.User/UserBaseInfo"
	User_UserListInfo_FullMethodName   = "/user_rpc.User/UserListInfo"
	User_IsFriend_FullMethodName       = "/user_rpc.User/IsFriend"
	User_FriendList_FullMethodName     = "/user_rpc.User/FriendList"
	User_UserOnlineList_FullMethodName = "/user_rpc.User/UserOnlineList"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	UserBaseInfo(ctx context.Context, in *UserBaseInfoRequest, opts ...grpc.CallOption) (*UserBaseInfoResponse, error)
	UserListInfo(ctx context.Context, in *UserListInfoRequest, opts ...grpc.CallOption) (*UserListInfoResponse, error)
	IsFriend(ctx context.Context, in *IsFriendRequest, opts ...grpc.CallOption) (*IsFriendResponse, error)
	FriendList(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error)
	UserOnlineList(ctx context.Context, in *UserOnlineListRequest, opts ...grpc.CallOption) (*UserOnlineListResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, User_UserCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, User_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserBaseInfo(ctx context.Context, in *UserBaseInfoRequest, opts ...grpc.CallOption) (*UserBaseInfoResponse, error) {
	out := new(UserBaseInfoResponse)
	err := c.cc.Invoke(ctx, User_UserBaseInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserListInfo(ctx context.Context, in *UserListInfoRequest, opts ...grpc.CallOption) (*UserListInfoResponse, error) {
	out := new(UserListInfoResponse)
	err := c.cc.Invoke(ctx, User_UserListInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) IsFriend(ctx context.Context, in *IsFriendRequest, opts ...grpc.CallOption) (*IsFriendResponse, error) {
	out := new(IsFriendResponse)
	err := c.cc.Invoke(ctx, User_IsFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) FriendList(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error) {
	out := new(FriendListResponse)
	err := c.cc.Invoke(ctx, User_FriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserOnlineList(ctx context.Context, in *UserOnlineListRequest, opts ...grpc.CallOption) (*UserOnlineListResponse, error) {
	out := new(UserOnlineListResponse)
	err := c.cc.Invoke(ctx, User_UserOnlineList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	UserBaseInfo(context.Context, *UserBaseInfoRequest) (*UserBaseInfoResponse, error)
	UserListInfo(context.Context, *UserListInfoRequest) (*UserListInfoResponse, error)
	IsFriend(context.Context, *IsFriendRequest) (*IsFriendResponse, error)
	FriendList(context.Context, *FriendListRequest) (*FriendListResponse, error)
	UserOnlineList(context.Context, *UserOnlineListRequest) (*UserOnlineListResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreate not implemented")
}
func (UnimplementedUserServer) UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedUserServer) UserBaseInfo(context.Context, *UserBaseInfoRequest) (*UserBaseInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserBaseInfo not implemented")
}
func (UnimplementedUserServer) UserListInfo(context.Context, *UserListInfoRequest) (*UserListInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserListInfo not implemented")
}
func (UnimplementedUserServer) IsFriend(context.Context, *IsFriendRequest) (*IsFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFriend not implemented")
}
func (UnimplementedUserServer) FriendList(context.Context, *FriendListRequest) (*FriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendList not implemented")
}
func (UnimplementedUserServer) UserOnlineList(context.Context, *UserOnlineListRequest) (*UserOnlineListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserOnlineList not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_UserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserCreate(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _User_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _User_UserBaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBaseInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserBaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserBaseInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserBaseInfo(ctx, req.(*UserBaseInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _User_UserListInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserListInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserListInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserListInfo(ctx, req.(*UserListInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _User_IsFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).IsFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_IsFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).IsFriend(ctx, req.(*IsFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _User_FriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_FriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FriendList(ctx, req.(*FriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _User_UserOnlineList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserOnlineListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserOnlineList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserOnlineList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserOnlineList(ctx, req.(*UserOnlineListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_rpc.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserCreate",
			Handler:    _User_UserCreate_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _User_UserInfo_Handler,
		},
		{
			MethodName: "UserBaseInfo",
			Handler:    _User_UserBaseInfo_Handler,
		},
		{
			MethodName: "UserListInfo",
			Handler:    _User_UserListInfo_Handler,
		},
		{
			MethodName: "IsFriend",
			Handler:    _User_IsFriend_Handler,
		},
		{
			MethodName: "FriendList",
			Handler:    _User_FriendList_Handler,
		},
		{
			MethodName: "UserOnlineList",
			Handler:    _User_UserOnlineList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/user.proto",
}
