// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user_rpc.proto

package users

import (
	"context"

	"fim_server/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserCreateRequest  = user_rpc.UserCreateRequest
	UserCreateResponse = user_rpc.UserCreateResponse
	UserInfoRequest    = user_rpc.UserInfoRequest
	UserInfoResponse   = user_rpc.UserInfoResponse

	Users interface {
		UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
		UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	}

	defaultUsers struct {
		cli zrpc.Client
	}
)

func NewUsers(cli zrpc.Client) Users {
	return &defaultUsers{
		cli: cli,
	}
}

func (m *defaultUsers) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	client := user_rpc.NewUsersClient(m.cli.Conn())
	return client.UserCreate(ctx, in, opts...)
}

func (m *defaultUsers) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user_rpc.NewUsersClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}
