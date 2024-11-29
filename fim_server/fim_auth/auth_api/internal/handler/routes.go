// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"fim_server/fim_auth/auth_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/auth/authentication",
				Handler: authenticationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/auth/login",
				Handler: loginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/auth/logout",
				Handler: logoutHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/auth/open_login",
				Handler: open_loginHandler(serverCtx),
			},
		},
	)
}
