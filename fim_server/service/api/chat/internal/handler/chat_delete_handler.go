package handler

import (
	"net/http"

	"fim_server/service/api/chat/internal/logic"
	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fim_server/service/server/response"
)

func ChatDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewChatDeleteLogic(r.Context(), svcCtx)
		resp, err := l.ChatDelete(&req)
		response.Response(r, w, resp, err)
	}
}