// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.2
// source: userCenterInterface/api/userInfo/v1/user.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_userCenterInterface_api_userInfo_v1_user_proto_rawDescGZIP(), []int{0}
}

type CreateUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateUserReply) Reset() {
	*x = CreateUserReply{}
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReply) ProtoMessage() {}

func (x *CreateUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserReply.ProtoReflect.Descriptor instead.
func (*CreateUserReply) Descriptor() ([]byte, []int) {
	return file_userCenterInterface_api_userInfo_v1_user_proto_rawDescGZIP(), []int{1}
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_userCenterInterface_api_userInfo_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUserReply) Reset() {
	*x = UpdateUserReply{}
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserReply) ProtoMessage() {}

func (x *UpdateUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserReply.ProtoReflect.Descriptor instead.
func (*UpdateUserReply) Descriptor() ([]byte, []int) {
	return file_userCenterInterface_api_userInfo_v1_user_proto_rawDescGZIP(), []int{3}
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_userCenterInterface_api_userInfo_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUserRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserReply) Reset() {
	*x = DeleteUserReply{}
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserReply) ProtoMessage() {}

func (x *DeleteUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserReply.ProtoReflect.Descriptor instead.
func (*DeleteUserReply) Descriptor() ([]byte, []int) {
	return file_userCenterInterface_api_userInfo_v1_user_proto_rawDescGZIP(), []int{5}
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_userCenterInterface_api_userInfo_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type GetUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetUserReply) Reset() {
	*x = GetUserReply{}
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReply) ProtoMessage() {}

func (x *GetUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReply.ProtoReflect.Descriptor instead.
func (*GetUserReply) Descriptor() ([]byte, []int) {
	return file_userCenterInterface_api_userInfo_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserReply) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *GetUserReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ListUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUserRequest) Reset() {
	*x = ListUserRequest{}
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRequest) ProtoMessage() {}

func (x *ListUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRequest.ProtoReflect.Descriptor instead.
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return file_userCenterInterface_api_userInfo_v1_user_proto_rawDescGZIP(), []int{8}
}

type ListUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUserReply) Reset() {
	*x = ListUserReply{}
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserReply) ProtoMessage() {}

func (x *ListUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserReply.ProtoReflect.Descriptor instead.
func (*ListUserReply) Descriptor() ([]byte, []int) {
	return file_userCenterInterface_api_userInfo_v1_user_proto_rawDescGZIP(), []int{9}
}

var File_userCenterInterface_api_userInfo_v1_user_proto protoreflect.FileDescriptor

var file_userCenterInterface_api_userInfo_v1_user_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x23, 0x75, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22,
	0x32, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xc9, 0x05, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x8c, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x36, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x3a, 0x01, 0x2a, 0x22, 0x05, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x91, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x36,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x1a, 0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x8e, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x36, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x2a, 0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x86, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x33, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d,
	0x12, 0x0b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x74, 0x61, 0x67, 0x7d, 0x12, 0x83, 0x01,
	0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x34, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x32, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x42, 0x4f, 0x0a, 0x23, 0x75, 0x73, 0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x26, 0x75, 0x73,
	0x65, 0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userCenterInterface_api_userInfo_v1_user_proto_rawDescOnce sync.Once
	file_userCenterInterface_api_userInfo_v1_user_proto_rawDescData = file_userCenterInterface_api_userInfo_v1_user_proto_rawDesc
)

func file_userCenterInterface_api_userInfo_v1_user_proto_rawDescGZIP() []byte {
	file_userCenterInterface_api_userInfo_v1_user_proto_rawDescOnce.Do(func() {
		file_userCenterInterface_api_userInfo_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_userCenterInterface_api_userInfo_v1_user_proto_rawDescData)
	})
	return file_userCenterInterface_api_userInfo_v1_user_proto_rawDescData
}

var file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_userCenterInterface_api_userInfo_v1_user_proto_goTypes = []any{
	(*CreateUserRequest)(nil), // 0: userCenterInterface.api.userInfo.v1.CreateUserRequest
	(*CreateUserReply)(nil),   // 1: userCenterInterface.api.userInfo.v1.CreateUserReply
	(*UpdateUserRequest)(nil), // 2: userCenterInterface.api.userInfo.v1.UpdateUserRequest
	(*UpdateUserReply)(nil),   // 3: userCenterInterface.api.userInfo.v1.UpdateUserReply
	(*DeleteUserRequest)(nil), // 4: userCenterInterface.api.userInfo.v1.DeleteUserRequest
	(*DeleteUserReply)(nil),   // 5: userCenterInterface.api.userInfo.v1.DeleteUserReply
	(*GetUserRequest)(nil),    // 6: userCenterInterface.api.userInfo.v1.GetUserRequest
	(*GetUserReply)(nil),      // 7: userCenterInterface.api.userInfo.v1.GetUserReply
	(*ListUserRequest)(nil),   // 8: userCenterInterface.api.userInfo.v1.ListUserRequest
	(*ListUserReply)(nil),     // 9: userCenterInterface.api.userInfo.v1.ListUserReply
}
var file_userCenterInterface_api_userInfo_v1_user_proto_depIdxs = []int32{
	0, // 0: userCenterInterface.api.userInfo.v1.User.CreateUser:input_type -> userCenterInterface.api.userInfo.v1.CreateUserRequest
	2, // 1: userCenterInterface.api.userInfo.v1.User.UpdateUser:input_type -> userCenterInterface.api.userInfo.v1.UpdateUserRequest
	4, // 2: userCenterInterface.api.userInfo.v1.User.DeleteUser:input_type -> userCenterInterface.api.userInfo.v1.DeleteUserRequest
	6, // 3: userCenterInterface.api.userInfo.v1.User.GetUser:input_type -> userCenterInterface.api.userInfo.v1.GetUserRequest
	8, // 4: userCenterInterface.api.userInfo.v1.User.ListUser:input_type -> userCenterInterface.api.userInfo.v1.ListUserRequest
	1, // 5: userCenterInterface.api.userInfo.v1.User.CreateUser:output_type -> userCenterInterface.api.userInfo.v1.CreateUserReply
	3, // 6: userCenterInterface.api.userInfo.v1.User.UpdateUser:output_type -> userCenterInterface.api.userInfo.v1.UpdateUserReply
	5, // 7: userCenterInterface.api.userInfo.v1.User.DeleteUser:output_type -> userCenterInterface.api.userInfo.v1.DeleteUserReply
	7, // 8: userCenterInterface.api.userInfo.v1.User.GetUser:output_type -> userCenterInterface.api.userInfo.v1.GetUserReply
	9, // 9: userCenterInterface.api.userInfo.v1.User.ListUser:output_type -> userCenterInterface.api.userInfo.v1.ListUserReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_userCenterInterface_api_userInfo_v1_user_proto_init() }
func file_userCenterInterface_api_userInfo_v1_user_proto_init() {
	if File_userCenterInterface_api_userInfo_v1_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_userCenterInterface_api_userInfo_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userCenterInterface_api_userInfo_v1_user_proto_goTypes,
		DependencyIndexes: file_userCenterInterface_api_userInfo_v1_user_proto_depIdxs,
		MessageInfos:      file_userCenterInterface_api_userInfo_v1_user_proto_msgTypes,
	}.Build()
	File_userCenterInterface_api_userInfo_v1_user_proto = out.File
	file_userCenterInterface_api_userInfo_v1_user_proto_rawDesc = nil
	file_userCenterInterface_api_userInfo_v1_user_proto_goTypes = nil
	file_userCenterInterface_api_userInfo_v1_user_proto_depIdxs = nil
}
