// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"fim_server/service/api/group/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(src *rest.Server, serverCtx *svc.ServiceContext) {
	src.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/group/group",
				Handler: GroupCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/group/group",
				Handler: GroupUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/group/:id",
				Handler: GroupInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/group/group/:id",
				Handler: GroupDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/member",
				Handler: GroupMemberHandler(serverCtx),
			},
		},
	)
}