package logic

import (
	"context"
	"fim_server/models"
	"fim_server/models/chat_models"
	"fim_server/service/rpc/user/user_rpc"
	"fim_server/utils/src"
	"fim_server/utils/src/sqls"
	"fim_server/utils/stores/logs"
	"fim_server/utils/stores/method"

	"fim_server/service/api/chat/internal/svc"
	"fim_server/service/api/chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatHistoryLogic {
	return &ChatHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type UserInfo struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
type ChatHistory struct {
	ID            uint                  `json:"id"`
	IsMe          bool                  `json:"is_me"`      // 哪条消息是我发的
	CreatedAt     string                `json:"created_at"` // 消息时间
	Message       models.Message        `json:"message"`
	SystemMessage *models.SystemMessage `json:"system_message"`
}

type ChatHistoryResponse struct {
	Total       int           `json:"total"`
	SendUser    UserInfo      `json:"sendUser"`
	ReceiveUser UserInfo      `json:"receive_user"`
	List        []ChatHistory `json:"list"`
}

func (l *ChatHistoryLogic) ChatHistory(req *types.ChatHistoryRequest) (resp *ChatHistoryResponse, err error) {
	// todo: add your logic here and delete this line

	_, err = l.svcCtx.UserRpc.IsFriend(context.Background(), &user_rpc.IsFriendRequest{User1: uint32(req.UserId), User2: uint32(req.FriendId)})
	if err != nil {
		return nil, logs.Error("不是好友")
	}

	chatList, total := sqls.GetList(chat_models.ChatModel{}, sqls.Mysql{
		DB: l.svcCtx.DB.Where("(send_user_id = ? and receive_user_id = ?) or "+
			"(send_user_id = ? and receive_user_id = ?) and id not in "+
			"(select chat_id from user_chat_delete_models where user_id = ?)",
			req.UserId, req.FriendId, req.FriendId, req.UserId, req.UserId),
		PageInfo: src.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "created_at desc",
		},
	})

	var userIdList []uint32
	for _, model := range chatList {
		userIdList = append(userIdList, uint32(model.SendUserId))
		userIdList = append(userIdList, uint32(model.ReceiveUserId))
	}
	userIdList = method.Deduplication(userIdList) // 去重
	// 调用户服务
	response, err := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
		UserIdList: userIdList,
	})
	if err != nil {
		return nil, logs.Error("用户服务错误")
	}

	var list = make([]ChatHistory, 0)
	var sendUser, receiveUser UserInfo
	for i, model := range chatList {

		if i == 0 {
			sendUser = UserInfo{
				ID:     model.SendUserId,
				Name:   response.UserInfo[uint32(model.SendUserId)].Name,
				Avatar: response.UserInfo[uint32(model.SendUserId)].Avatar,
			}
			receiveUser = UserInfo{
				ID:     model.ReceiveUserId,
				Name:   response.UserInfo[uint32(model.ReceiveUserId)].Name,
				Avatar: response.UserInfo[uint32(model.ReceiveUserId)].Avatar,
			}
		}

		info := ChatHistory{
			ID:            model.ID,
			CreatedAt:     model.CreatedAt.String(),
			Message:       model.Message,
			SystemMessage: model.SystemMessage,
		}
		if model.SendUserId == req.UserId {
			info.IsMe = true
		}
		list = append(list, info)

	}
	resp = &ChatHistoryResponse{
		Total:       int(total),
		SendUser:    sendUser,
		ReceiveUser: receiveUser,
		List:        list,
	}

	return
}
