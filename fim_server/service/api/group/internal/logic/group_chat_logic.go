package logic

import (
	"context"

	"fim_server/service/api/group/internal/svc"
	"fim_server/service/api/group/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupChatLogic {
	return &GroupChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupChatLogic) GroupChat(req *types.GroupChatRequest) (resp *types.GroupChatResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
