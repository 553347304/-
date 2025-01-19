package logic

import (
	"context"
	
	"fim_server/service/api/log/internal/svc"
	"fim_server/service/api/log/internal/types"
)

type LogReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogReadLogic {
	return &LogReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogReadLogic) LogRead(req *types.LogReadRequest) (resp *types.Empty, err error) {
	
	// var logModel logs_model.LogModel
	// err = l.svcCtx.DB.Take(&logModel, req.ID).Error
	// if err != nil {
	// 	return nil, errors.New("日志记录不存在")
	// }
	// // 前端要判断一下，如果已经读取了，就不要再调接口了
	// if logModel.IsRead {
	// 	return
	// }
	// l.svcCtx.DB.Model(&logModel).Update("is_read", true)
	return
}
