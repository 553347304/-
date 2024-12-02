// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package user

import (
	"context"

	"fim_server/service/rpc/user/user_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserCreateRequest    = user_rpc.UserCreateRequest
	UserCreateResponse   = user_rpc.UserCreateResponse
	UserInfo             = user_rpc.UserInfo
	UserInfoRequest      = user_rpc.UserInfoRequest
	UserInfoResponse     = user_rpc.UserInfoResponse
	UserListInfoRequest  = user_rpc.UserListInfoRequest
	UserListInfoResponse = user_rpc.UserListInfoResponse

	User interface {
		UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
		UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
		UserListInfo(ctx context.Context, in *UserListInfoRequest, opts ...grpc.CallOption) (*UserListInfoResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	client := user_rpc.NewUserClient(m.cli.Conn())
	return client.UserCreate(ctx, in, opts...)
}

func (m *defaultUser) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user_rpc.NewUserClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultUser) UserListInfo(ctx context.Context, in *UserListInfoRequest, opts ...grpc.CallOption) (*UserListInfoResponse, error) {
	client := user_rpc.NewUserClient(m.cli.Conn())
	return client.UserListInfo(ctx, in, opts...)
}