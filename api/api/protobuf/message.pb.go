// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: api/api/message_v1/message.proto

package protobuf

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

type MessageType int32

const (
	MessageType__       MessageType = 0
	MessageType_Text    MessageType = 1
	MessageType_Sticker MessageType = 2
	MessageType_Video   MessageType = 3
	MessageType_Image   MessageType = 4
	MessageType_GIF     MessageType = 5
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "_",
		1: "Text",
		2: "Sticker",
		3: "Video",
		4: "Image",
		5: "GIF",
	}
	MessageType_value = map[string]int32{
		"_":       0,
		"Text":    1,
		"Sticker": 2,
		"Video":   3,
		"Image":   4,
		"GIF":     5,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_api_message_v1_message_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_api_api_message_v1_message_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_api_api_message_v1_message_proto_rawDescGZIP(), []int{0}
}

type GetMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId uint64 `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
}

func (x *GetMessageRequest) Reset() {
	*x = GetMessageRequest{}
	mi := &file_api_api_message_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequest) ProtoMessage() {}

func (x *GetMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_message_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequest.ProtoReflect.Descriptor instead.
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_api_message_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *GetMessageRequest) GetMessageId() uint64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type GetRandomMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId uint64      `protobuf:"varint,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
	Type   MessageType `protobuf:"varint,2,opt,name=type,proto3,enum=api.MessageType" json:"type,omitempty"`
}

func (x *GetRandomMessageRequest) Reset() {
	*x = GetRandomMessageRequest{}
	mi := &file_api_api_message_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRandomMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRandomMessageRequest) ProtoMessage() {}

func (x *GetRandomMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_message_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRandomMessageRequest.ProtoReflect.Descriptor instead.
func (*GetRandomMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_api_message_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *GetRandomMessageRequest) GetChatId() uint64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *GetRandomMessageRequest) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType__
}

type AddMessageCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId  uint64      `protobuf:"varint,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
	Content string      `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Type    MessageType `protobuf:"varint,3,opt,name=type,proto3,enum=api.MessageType" json:"type,omitempty"`
}

func (x *AddMessageCommand) Reset() {
	*x = AddMessageCommand{}
	mi := &file_api_api_message_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddMessageCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMessageCommand) ProtoMessage() {}

func (x *AddMessageCommand) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_message_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMessageCommand.ProtoReflect.Descriptor instead.
func (*AddMessageCommand) Descriptor() ([]byte, []int) {
	return file_api_api_message_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *AddMessageCommand) GetChatId() uint64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *AddMessageCommand) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AddMessageCommand) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType__
}

type AddMessageCommandResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId uint64 `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
}

func (x *AddMessageCommandResult) Reset() {
	*x = AddMessageCommandResult{}
	mi := &file_api_api_message_v1_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddMessageCommandResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMessageCommandResult) ProtoMessage() {}

func (x *AddMessageCommandResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_message_v1_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMessageCommandResult.ProtoReflect.Descriptor instead.
func (*AddMessageCommandResult) Descriptor() ([]byte, []int) {
	return file_api_api_message_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *AddMessageCommandResult) GetMessageId() uint64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string      `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Type    MessageType `protobuf:"varint,3,opt,name=type,proto3,enum=api.MessageType" json:"type,omitempty"`
	Sender  *SenderInfo `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	mi := &file_api_api_message_v1_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_message_v1_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_api_api_message_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *MessageResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MessageResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MessageResponse) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType__
}

func (x *MessageResponse) GetSender() *SenderInfo {
	if x != nil {
		return x.Sender
	}
	return nil
}

type SenderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SenderInfo) Reset() {
	*x = SenderInfo{}
	mi := &file_api_api_message_v1_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SenderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SenderInfo) ProtoMessage() {}

func (x *SenderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_message_v1_message_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SenderInfo.ProtoReflect.Descriptor instead.
func (*SenderInfo) Descriptor() ([]byte, []int) {
	return file_api_api_message_v1_message_proto_rawDescGZIP(), []int{5}
}

func (x *SenderInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SenderInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_api_message_v1_message_proto protoreflect.FileDescriptor

var file_api_api_message_v1_message_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x31, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x6b, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x37, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x0f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x50, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x05, 0x0a, 0x01, 0x5f, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x03,
	0x12, 0x09, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x47,
	0x49, 0x46, 0x10, 0x05, 0x22, 0x04, 0x08, 0x06, 0x10, 0x0a, 0x32, 0xd8, 0x01, 0x0a, 0x0e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x42, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x1d, 0x5a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_message_v1_message_proto_rawDescOnce sync.Once
	file_api_api_message_v1_message_proto_rawDescData = file_api_api_message_v1_message_proto_rawDesc
)

func file_api_api_message_v1_message_proto_rawDescGZIP() []byte {
	file_api_api_message_v1_message_proto_rawDescOnce.Do(func() {
		file_api_api_message_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_message_v1_message_proto_rawDescData)
	})
	return file_api_api_message_v1_message_proto_rawDescData
}

var file_api_api_message_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_api_message_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_api_message_v1_message_proto_goTypes = []any{
	(MessageType)(0),                // 0: api.MessageType
	(*GetMessageRequest)(nil),       // 1: api.GetMessageRequest
	(*GetRandomMessageRequest)(nil), // 2: api.GetRandomMessageRequest
	(*AddMessageCommand)(nil),       // 3: api.AddMessageCommand
	(*AddMessageCommandResult)(nil), // 4: api.AddMessageCommandResult
	(*MessageResponse)(nil),         // 5: api.MessageResponse
	(*SenderInfo)(nil),              // 6: api.SenderInfo
}
var file_api_api_message_v1_message_proto_depIdxs = []int32{
	0, // 0: api.GetRandomMessageRequest.type:type_name -> api.MessageType
	0, // 1: api.AddMessageCommand.type:type_name -> api.MessageType
	0, // 2: api.MessageResponse.type:type_name -> api.MessageType
	6, // 3: api.MessageResponse.sender:type_name -> api.SenderInfo
	1, // 4: api.MessageService.GetMessage:input_type -> api.GetMessageRequest
	2, // 5: api.MessageService.GetRandomMessage:input_type -> api.GetRandomMessageRequest
	3, // 6: api.MessageService.AddMessage:input_type -> api.AddMessageCommand
	5, // 7: api.MessageService.GetMessage:output_type -> api.MessageResponse
	5, // 8: api.MessageService.GetRandomMessage:output_type -> api.MessageResponse
	4, // 9: api.MessageService.AddMessage:output_type -> api.AddMessageCommandResult
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_api_message_v1_message_proto_init() }
func file_api_api_message_v1_message_proto_init() {
	if File_api_api_message_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_api_message_v1_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_message_v1_message_proto_goTypes,
		DependencyIndexes: file_api_api_message_v1_message_proto_depIdxs,
		EnumInfos:         file_api_api_message_v1_message_proto_enumTypes,
		MessageInfos:      file_api_api_message_v1_message_proto_msgTypes,
	}.Build()
	File_api_api_message_v1_message_proto = out.File
	file_api_api_message_v1_message_proto_rawDesc = nil
	file_api_api_message_v1_message_proto_goTypes = nil
	file_api_api_message_v1_message_proto_depIdxs = nil
}
