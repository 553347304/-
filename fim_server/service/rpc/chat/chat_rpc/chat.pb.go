// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.9.0
// source: rpc/chat.proto

package chat_rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendUserId    uint32 `protobuf:"varint,1,opt,name=send_user_id,json=sendUserId,proto3" json:"send_user_id,omitempty"`
	ReceiveUserId uint32 `protobuf:"varint,2,opt,name=receive_user_id,json=receiveUserId,proto3" json:"receive_user_id,omitempty"`
	Message       []byte `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`                                  // message json格式
	SystemMessage []byte `protobuf:"bytes,4,opt,name=system_message,json=systemMessage,proto3" json:"system_message,omitempty"` // 系统消息
}

func (x *UserChatRequest) Reset() {
	*x = UserChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChatRequest) ProtoMessage() {}

func (x *UserChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChatRequest.ProtoReflect.Descriptor instead.
func (*UserChatRequest) Descriptor() ([]byte, []int) {
	return file_rpc_chat_proto_rawDescGZIP(), []int{0}
}

func (x *UserChatRequest) GetSendUserId() uint32 {
	if x != nil {
		return x.SendUserId
	}
	return 0
}

func (x *UserChatRequest) GetReceiveUserId() uint32 {
	if x != nil {
		return x.ReceiveUserId
	}
	return 0
}

func (x *UserChatRequest) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *UserChatRequest) GetSystemMessage() []byte {
	if x != nil {
		return x.SystemMessage
	}
	return nil
}

type UserChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserChatResponse) Reset() {
	*x = UserChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChatResponse) ProtoMessage() {}

func (x *UserChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChatResponse.ProtoReflect.Descriptor instead.
func (*UserChatResponse) Descriptor() ([]byte, []int) {
	return file_rpc_chat_proto_rawDescGZIP(), []int{1}
}

func (x *UserChatResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_rpc_chat_proto protoreflect.FileDescriptor

var file_rpc_chat_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x22, 0x9c, 0x01, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x49, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x41,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_chat_proto_rawDescOnce sync.Once
	file_rpc_chat_proto_rawDescData = file_rpc_chat_proto_rawDesc
)

func file_rpc_chat_proto_rawDescGZIP() []byte {
	file_rpc_chat_proto_rawDescOnce.Do(func() {
		file_rpc_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_chat_proto_rawDescData)
	})
	return file_rpc_chat_proto_rawDescData
}

var file_rpc_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_chat_proto_goTypes = []interface{}{
	(*UserChatRequest)(nil),  // 0: user_rpc.UserChatRequest
	(*UserChatResponse)(nil), // 1: user_rpc.UserChatResponse
}
var file_rpc_chat_proto_depIdxs = []int32{
	0, // 0: user_rpc.Chat.UserChat:input_type -> user_rpc.UserChatRequest
	1, // 1: user_rpc.Chat.UserChat:output_type -> user_rpc.UserChatResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_chat_proto_init() }
func file_rpc_chat_proto_init() {
	if File_rpc_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChatRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChatResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_chat_proto_goTypes,
		DependencyIndexes: file_rpc_chat_proto_depIdxs,
		MessageInfos:      file_rpc_chat_proto_msgTypes,
	}.Build()
	File_rpc_chat_proto = out.File
	file_rpc_chat_proto_rawDesc = nil
	file_rpc_chat_proto_goTypes = nil
	file_rpc_chat_proto_depIdxs = nil
}