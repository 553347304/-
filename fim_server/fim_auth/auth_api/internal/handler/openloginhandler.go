package handler

import (
	"fim_server/fim_auth/auth_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"fim_server/common/response"
	"fim_server/fim_auth/auth_api/internal/logic"
	"fim_server/fim_auth/auth_api/internal/svc"
)

func open_loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.OpenLoginInfoRequest
		err := httpx.Parse(r, &req)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewOpen_loginLogic(r.Context(), svcCtx)

		resp, err := l.Open_login(&req)
		response.Response(r, w, resp, err)
	}
}
