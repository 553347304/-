package handler

import (
	"fim_server/fim_file/file_api/internal/logic"
	"fim_server/fim_file/file_api/internal/svc"
	"fim_server/fim_file/file_api/internal/types"
	"fim_server/utils/stores/files"
	"fim_server/utils/stores/logs"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"path"

	"fim_server/common/response"
)

func ImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		// 图片类型
		imageType := r.FormValue("image_type")
		if imageType == "" {
			response.Response(r, w, nil, logs.Error("image_type为空"))
			return
		}

		// 读取文件
		file := files.FormFile(files.File{R: r, Key: "image",
			MaxSize: &svcCtx.Config.File.MaxSize, WhiteEXT: &svcCtx.Config.File.WhiteEXT})
		if file.Error != nil {
			response.Response(r, w, nil, logs.Error(file.Error))
			return
		}

		l := logic.NewImageLogic(r.Context(), svcCtx)
		resp, err := l.Image(&req)

		resp.Url = files.WriteFile(path.Join(svcCtx.Config.File.Path, imageType, file.Name), file.Byte)

		response.Response(r, w, resp, err)
	}
}
