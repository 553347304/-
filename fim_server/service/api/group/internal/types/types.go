// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

import "fim_server/models/mgorm"

type EmptyResponse struct {
}

type GroupAddRequest struct {
	UserId    uint      `header:"User-Id"`
	GroupId   uint      `json:"group_id"`
	ValidInfo ValidInfo `json:"valid_info,optional"`
}

type GroupAddResponse struct {
}

type GroupBanUpdateRequest struct {
	UserId   uint `header:"User-Id"`
	GroupId  uint `json:"group_id"`
	MemberId uint `json:"member_id"`
	BanTime  int  `json:"ban_time"` // 禁言时间
}

type GroupBanUpdateResponse struct {
}

type GroupChatRequest struct {
	Header
}

type GroupChatResponse struct {
}

type GroupCreateRequest struct {
	UserId     uint   `header:"User-Id"`
	Mode       int8   `json:"mode,optional"` // 模式 1 直接创建   2 创建模式
	Name       string `json:"name,optional"`
	IsSearch   bool   `json:"is_search,optional"`
	Size       int    `json:"size,optional"`
	UserIdList []uint `json:"user_id_list,optional"`
}

type GroupCreateResponse struct {
}

type GroupDeleteRequest struct {
	UserId uint `header:"User-Id"`
	Id     int8 `path:"id"`
}

type GroupDeleteResponse struct {
}

type GroupFriendsInfo struct {
	UserId    uint   `json:"user_id"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
	IsInGroup bool   `json:"is_in_group"`
}

type GroupFriendsListRequest struct {
	UserId uint `header:"User-Id"`
	Id     uint `form:"id"`
}

type GroupFriendsListResponse struct {
	Total int64              `json:"total"`
	List  []GroupFriendsInfo `json:"list"`
}

type GroupHistoryDeleteRequest struct {
	Header
	ParamsPath
	IdList []uint `json:"id_list"`
}

type GroupHistoryRequest struct {
	Header
	ParamsPath
	PageInfo
}

type GroupHistoryResponse struct {
}

type GroupInfoRequest struct {
	UserId uint `header:"User-Id"`
	Id     int8 `path:"id"`
}

type GroupInfoResponse struct {
	GroupId          uint       `json:"group_id"`
	Name             string     `json:"name"`
	Sign             string     `json:"sign"`
	Avatar           string     `json:"avatar"`
	MemberCount      int        `json:"member_count"`
	MemberOnlinCount int        `json:"member_onlin_count"`
	Leader           UserInfo   `json:"leader"` // 群主
	AdminList        []UserInfo `json:"admin_list"`
	Role             int8       `json:"role"`     // 角色   1 群主 2 群管理员 3 群成员
	IsBan            bool       `json:"is_time"`  // is禁言
	BanTime          *int       `json:"ban_time"` // 禁言时间 单位分钟
}

type GroupMeInfo struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	MemberTatal int64  `json:"member_tatal"`
	Role        int8   `json:"role"`
	Mode        int8   `json:"mode"`
}

type GroupMeListRequest struct {
	Header
	PageInfo
	Mode int8 `form:"mode"` // 1 我创建的群 | 2 我加入的群
}

type GroupMeListResponse struct {
	Total int64         `json:"total"`
	List  []GroupMeInfo `json:"list"`
}

type GroupMemberAddRequest struct {
	UserId       uint   `header:"User-Id"`
	Id           uint   `json:"id"`
	MemberIdList []uint `json:"member_id_list"`
}

type GroupMemberAddResponse struct {
}

type GroupMemberDeleteRequest struct {
	UserId   uint `header:"User-Id"`
	Id       uint `form:"id"`
	MemberId uint `form:"member_id"`
}

type GroupMemberDeleteResponse struct {
}

type GroupMemberInfo struct {
	UserId         uint   `json:"user_id"`
	Name           string `json:"name"`
	Avatar         string `json:"avatar"`
	InOnline       bool   `json:"in_online"`
	Role           int8   `json:"role"`
	MemberName     string `json:"member_name"`
	CreatedAt      string `json:"created_at"`
	NewMessageDate string `json:"new_message_date"`
}

type GroupMemberInfoRequest struct {
	UserId uint `header:"User-Id"`
	Id     uint `path:"id"`
}

type GroupMemberNameRequest struct {
	UserId   uint   `header:"User-Id"`
	Id       uint   `json:"id"`
	MemberId uint   `json:"member_id"`
	Name     string `json:"name"`
}

type GroupMemberNameResponse struct {
}

type GroupMemberRequest struct {
	UserId uint   `header:"User-Id"`
	Id     uint   `form:"id"`
	Page   int    `form:"page,optional"`
	Limit  int    `form:"limit,optional"`
	Sort   string `form:"sort,optional"`
}

type GroupMemberResponse struct {
	Total int64             `json:"total"`
	List  []GroupMemberInfo `json:"list"`
}

type GroupMemberRoleRequest struct {
	UserId   uint `header:"User-Id"`
	Id       uint `json:"id"`
	MemberId uint `json:"member_id"`
	Role     int8 `json:"role"`
}

type GroupMemberRoleResponse struct {
}

type GroupSearchInfo struct {
	GroupId         uint   `json:"group_id"`
	Name            string `json:"name"`
	Sign            string `json:"sign"`
	Avatar          string `json:"avatar"`
	IsInGroup       bool   `json:"is_in_group"`
	UserCount       int    `json:"user_count"`
	UserOnlineCount int    `json:"user_online_count"`
}

type GroupSearchListRequest struct {
	UserId uint   `header:"User-Id"`
	Id     string `form:"id"`
	Key    string `form:"key"`
	Page   int    `form:"page,optional"`
	Limit  int    `form:"limit,optional"`
}

type GroupSearchListResponse struct {
	Total int64             `json:"total"`
	List  []GroupSearchInfo `json:"list"`
}

type GroupSessionRequest struct {
	Header
	PageInfo
}

type GroupSessionResponse struct {
	GroupId           uint   `json:"group_id"`
	Name              string `json:"name"`
	Avatar            string `json:"avatar"`
	NewMessageDate    string `json:"new_message_date"`
	NewMessagePreview string `json:"new_message_preview"`
}

type GroupTopRequest struct {
	Header
	GroupId uint `json:"group_id"`
	Top     bool `json:"top"`
}

type GroupTopResponse struct {
}

type GroupUpdateRequest struct {
	UserId             uint      `header:"User-Id"`
	Id                 int8      `json:"id"`
	Name               string    `json:"name,optional"`                 // 群名
	Avatar             string    `json:"avatar,optional"`               // 群头像
	Sign               string    `json:"sign,optional"`                 // 群简介
	IsSearch           *bool     `json:"is_search,optional"`            // is搜索
	IsInvite           *bool     `json:"is_invite,optional"`            // is邀请
	IsTemporarySession *bool     `json:"is_temporary_session,optional"` // is临时会话
	IsBan              *bool     `json:"is_time,optional"`              // is禁言
	Valid              int8      `json:"valid，optional"`
	ValidInfo          ValidInfo `json:"valid_info,optional"`
}

type GroupUpdateResponse struct {
}

type GroupValidInfo struct {
	ID         uint      `json:"id"`
	UserId     uint      `header:"User-Id"`
	GroupId    uint      `json:"group_id"`
	UserAvatar string    `json:"user_avatar"`
	UserName   string    `json:"user_name"`
	Name       string    `json:"name"`
	Status     int8      `json:"status"` // 状态
	Valid      int8      `json:"valid，optional"`
	ValidInfo  ValidInfo `json:"valid_info,optional"`
	Type       int8      `json:"type"` // 1 加群 2 退群
	CreatedAt  string    `json:"created_at"`
}

type GroupValidIssueRequest struct {
	UserId uint `header:"User-Id"`
	Id     uint `path:"id"`
}

type GroupValidIssueResponse struct {
	Valid     int8      `json:"valid，optional"`
	ValidInfo ValidInfo `json:"valid_info,optional"`
}

type GroupValidListRequest struct {
	UserId  uint `header:"User-Id"`
	GroupId uint `form:"group_id"`
	Page    int  `form:"page,optional"`
	Limit   int  `form:"limit,optional"`
}

type GroupValidListResponse struct {
	Total int64            `json:"total"`
	List  []GroupValidInfo `json:"list"`
}

type GroupValidStatusRequest struct {
	UserId  uint `header:"User-Id"`
	VaildId uint `json:"vaild_id"`
	Status  int8 `json:"status"`
}

type GroupValidStatusResponse struct {
}

type Header struct {
	UserId uint `header:"User-Id"`
}

type PageInfo struct {
	Key   string `form:"key,optional"`
	Page  int    `form:"page,optional"`
	Limit int    `form:"limit,optional"`
}

type ParamsPath struct {
	Id uint `path:"id"`
}

type UserInfo struct {
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type ValidInfo struct {
	Issue  mgorm.String `json:"issue,optional"`
	Answer mgorm.String `json:"answer,optional"`
}
