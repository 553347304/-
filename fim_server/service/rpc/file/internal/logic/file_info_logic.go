package logic

import (
	"context"
	"fim_server/models/file_models"
	"fim_server/service/rpc/file/file_rpc"
	"fim_server/service/rpc/file/internal/svc"
	"fim_server/utils/stores/logs"
	"path"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileInfoLogic {
	return &FileInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileInfoLogic) FileInfo(in *file_rpc.FileInfoRequest) (*file_rpc.FileInfoResponse, error) {
	// todo: add your logic here and delete this line

	uid := strings.Replace(in.FileId, "/api/file/", "", -1)
	var fileModel file_models.FileModel
	err := l.svcCtx.DB.Take(&fileModel, "uid = ?", uid).Error
	if err != nil {
		return nil, logs.Error("文件不存在", uid)
	}

	return &file_rpc.FileInfoResponse{
		Name: fileModel.Name,
		Hash: fileModel.Hash,
		Size: fileModel.Size,
		Ext:  path.Ext(fileModel.Name),
	}, nil
}