// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: judger/v1/judger.proto

package v1

import (
	proto "github.com/Jecosine/alioth-kratos/api/proto"
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

type PingJudgerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingJudgerRequest) Reset() {
	*x = PingJudgerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judger_v1_judger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingJudgerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingJudgerRequest) ProtoMessage() {}

func (x *PingJudgerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_judger_v1_judger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingJudgerRequest.ProtoReflect.Descriptor instead.
func (*PingJudgerRequest) Descriptor() ([]byte, []int) {
	return file_judger_v1_judger_proto_rawDescGZIP(), []int{0}
}

type PingJudgerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data *JudgerStatus `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PingJudgerReply) Reset() {
	*x = PingJudgerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judger_v1_judger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingJudgerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingJudgerReply) ProtoMessage() {}

func (x *PingJudgerReply) ProtoReflect() protoreflect.Message {
	mi := &file_judger_v1_judger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingJudgerReply.ProtoReflect.Descriptor instead.
func (*PingJudgerReply) Descriptor() ([]byte, []int) {
	return file_judger_v1_judger_proto_rawDescGZIP(), []int{1}
}

func (x *PingJudgerReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PingJudgerReply) GetData() *JudgerStatus {
	if x != nil {
		return x.Data
	}
	return nil
}

// JudgerStatus is a struct representing running status of current judger
type JudgerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status             int64   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	CurrentTasksAmount int64   `protobuf:"varint,2,opt,name=current_tasks_amount,json=currentTasksAmount,proto3" json:"current_tasks_amount,omitempty"`
	EstimateTime       float64 `protobuf:"fixed64,3,opt,name=estimate_time,json=estimateTime,proto3" json:"estimate_time,omitempty"`
}

func (x *JudgerStatus) Reset() {
	*x = JudgerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judger_v1_judger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgerStatus) ProtoMessage() {}

func (x *JudgerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_judger_v1_judger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgerStatus.ProtoReflect.Descriptor instead.
func (*JudgerStatus) Descriptor() ([]byte, []int) {
	return file_judger_v1_judger_proto_rawDescGZIP(), []int{2}
}

func (x *JudgerStatus) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *JudgerStatus) GetCurrentTasksAmount() int64 {
	if x != nil {
		return x.CurrentTasksAmount
	}
	return 0
}

func (x *JudgerStatus) GetEstimateTime() float64 {
	if x != nil {
		return x.EstimateTime
	}
	return 0
}

type SubmitJudgerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *proto.JudgeRequestProto `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *SubmitJudgerRequest) Reset() {
	*x = SubmitJudgerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judger_v1_judger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitJudgerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitJudgerRequest) ProtoMessage() {}

func (x *SubmitJudgerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_judger_v1_judger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitJudgerRequest.ProtoReflect.Descriptor instead.
func (*SubmitJudgerRequest) Descriptor() ([]byte, []int) {
	return file_judger_v1_judger_proto_rawDescGZIP(), []int{3}
}

func (x *SubmitJudgerRequest) GetPayload() *proto.JudgeRequestProto {
	if x != nil {
		return x.Payload
	}
	return nil
}

type SubmitJudgerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Received string `protobuf:"bytes,2,opt,name=received,proto3" json:"received,omitempty"`
}

func (x *SubmitJudgerReply) Reset() {
	*x = SubmitJudgerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judger_v1_judger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitJudgerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitJudgerReply) ProtoMessage() {}

func (x *SubmitJudgerReply) ProtoReflect() protoreflect.Message {
	mi := &file_judger_v1_judger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitJudgerReply.ProtoReflect.Descriptor instead.
func (*SubmitJudgerReply) Descriptor() ([]byte, []int) {
	return file_judger_v1_judger_proto_rawDescGZIP(), []int{4}
}

func (x *SubmitJudgerReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SubmitJudgerReply) GetReceived() string {
	if x != nil {
		return x.Received
	}
	return ""
}

type DeleteJudgerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteJudgerRequest) Reset() {
	*x = DeleteJudgerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judger_v1_judger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJudgerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJudgerRequest) ProtoMessage() {}

func (x *DeleteJudgerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_judger_v1_judger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJudgerRequest.ProtoReflect.Descriptor instead.
func (*DeleteJudgerRequest) Descriptor() ([]byte, []int) {
	return file_judger_v1_judger_proto_rawDescGZIP(), []int{5}
}

type DeleteJudgerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteJudgerReply) Reset() {
	*x = DeleteJudgerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judger_v1_judger_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJudgerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJudgerReply) ProtoMessage() {}

func (x *DeleteJudgerReply) ProtoReflect() protoreflect.Message {
	mi := &file_judger_v1_judger_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJudgerReply.ProtoReflect.Descriptor instead.
func (*DeleteJudgerReply) Descriptor() ([]byte, []int) {
	return file_judger_v1_judger_proto_rawDescGZIP(), []int{6}
}

type GetJudgerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetJudgerRequest) Reset() {
	*x = GetJudgerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judger_v1_judger_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJudgerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJudgerRequest) ProtoMessage() {}

func (x *GetJudgerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_judger_v1_judger_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJudgerRequest.ProtoReflect.Descriptor instead.
func (*GetJudgerRequest) Descriptor() ([]byte, []int) {
	return file_judger_v1_judger_proto_rawDescGZIP(), []int{7}
}

type GetJudgerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetJudgerReply) Reset() {
	*x = GetJudgerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judger_v1_judger_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJudgerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJudgerReply) ProtoMessage() {}

func (x *GetJudgerReply) ProtoReflect() protoreflect.Message {
	mi := &file_judger_v1_judger_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJudgerReply.ProtoReflect.Descriptor instead.
func (*GetJudgerReply) Descriptor() ([]byte, []int) {
	return file_judger_v1_judger_proto_rawDescGZIP(), []int{8}
}

type ListJudgerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListJudgerRequest) Reset() {
	*x = ListJudgerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judger_v1_judger_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJudgerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJudgerRequest) ProtoMessage() {}

func (x *ListJudgerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_judger_v1_judger_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJudgerRequest.ProtoReflect.Descriptor instead.
func (*ListJudgerRequest) Descriptor() ([]byte, []int) {
	return file_judger_v1_judger_proto_rawDescGZIP(), []int{9}
}

type ListJudgerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListJudgerReply) Reset() {
	*x = ListJudgerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judger_v1_judger_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJudgerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJudgerReply) ProtoMessage() {}

func (x *ListJudgerReply) ProtoReflect() protoreflect.Message {
	mi := &file_judger_v1_judger_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJudgerReply.ProtoReflect.Descriptor instead.
func (*ListJudgerReply) Descriptor() ([]byte, []int) {
	return file_judger_v1_judger_proto_rawDescGZIP(), []int{10}
}

var File_judger_v1_judger_proto protoreflect.FileDescriptor

var file_judger_v1_judger_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x75, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x13, 0x0a, 0x11, 0x50, 0x69, 0x6e, 0x67, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x52, 0x0a, 0x0f, 0x50, 0x69, 0x6e, 0x67, 0x4a, 0x75,
	0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x75,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7d, 0x0a, 0x0c, 0x4a, 0x75,
	0x64, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x65, 0x73, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x13, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x30, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x43, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4a, 0x75, 0x64, 0x67,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4a, 0x75,
	0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x32, 0xd3, 0x01, 0x0a, 0x06, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x4a, 0x75, 0x64, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c,
	0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x69, 0x0a, 0x06,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x75, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4a, 0x75, 0x64,
	0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x65, 0x63, 0x6f, 0x73, 0x69, 0x6e, 0x65, 0x2f, 0x61,
	0x6c, 0x69, 0x6f, 0x74, 0x68, 0x2d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_judger_v1_judger_proto_rawDescOnce sync.Once
	file_judger_v1_judger_proto_rawDescData = file_judger_v1_judger_proto_rawDesc
)

func file_judger_v1_judger_proto_rawDescGZIP() []byte {
	file_judger_v1_judger_proto_rawDescOnce.Do(func() {
		file_judger_v1_judger_proto_rawDescData = protoimpl.X.CompressGZIP(file_judger_v1_judger_proto_rawDescData)
	})
	return file_judger_v1_judger_proto_rawDescData
}

var file_judger_v1_judger_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_judger_v1_judger_proto_goTypes = []interface{}{
	(*PingJudgerRequest)(nil),       // 0: api.judger.v1.PingJudgerRequest
	(*PingJudgerReply)(nil),         // 1: api.judger.v1.PingJudgerReply
	(*JudgerStatus)(nil),            // 2: api.judger.v1.JudgerStatus
	(*SubmitJudgerRequest)(nil),     // 3: api.judger.v1.SubmitJudgerRequest
	(*SubmitJudgerReply)(nil),       // 4: api.judger.v1.SubmitJudgerReply
	(*DeleteJudgerRequest)(nil),     // 5: api.judger.v1.DeleteJudgerRequest
	(*DeleteJudgerReply)(nil),       // 6: api.judger.v1.DeleteJudgerReply
	(*GetJudgerRequest)(nil),        // 7: api.judger.v1.GetJudgerRequest
	(*GetJudgerReply)(nil),          // 8: api.judger.v1.GetJudgerReply
	(*ListJudgerRequest)(nil),       // 9: api.judger.v1.ListJudgerRequest
	(*ListJudgerReply)(nil),         // 10: api.judger.v1.ListJudgerReply
	(*proto.JudgeRequestProto)(nil), // 11: api.JudgeRequestProto
}
var file_judger_v1_judger_proto_depIdxs = []int32{
	2,  // 0: api.judger.v1.PingJudgerReply.data:type_name -> api.judger.v1.JudgerStatus
	11, // 1: api.judger.v1.SubmitJudgerRequest.payload:type_name -> api.JudgeRequestProto
	0,  // 2: api.judger.v1.Judger.Ping:input_type -> api.judger.v1.PingJudgerRequest
	3,  // 3: api.judger.v1.Judger.Submit:input_type -> api.judger.v1.SubmitJudgerRequest
	1,  // 4: api.judger.v1.Judger.Ping:output_type -> api.judger.v1.PingJudgerReply
	4,  // 5: api.judger.v1.Judger.Submit:output_type -> api.judger.v1.SubmitJudgerReply
	4,  // [4:6] is the sub-list for method output_type
	2,  // [2:4] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_judger_v1_judger_proto_init() }
func file_judger_v1_judger_proto_init() {
	if File_judger_v1_judger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_judger_v1_judger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingJudgerRequest); i {
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
		file_judger_v1_judger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingJudgerReply); i {
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
		file_judger_v1_judger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JudgerStatus); i {
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
		file_judger_v1_judger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitJudgerRequest); i {
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
		file_judger_v1_judger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitJudgerReply); i {
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
		file_judger_v1_judger_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJudgerRequest); i {
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
		file_judger_v1_judger_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJudgerReply); i {
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
		file_judger_v1_judger_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJudgerRequest); i {
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
		file_judger_v1_judger_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJudgerReply); i {
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
		file_judger_v1_judger_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJudgerRequest); i {
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
		file_judger_v1_judger_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJudgerReply); i {
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
			RawDescriptor: file_judger_v1_judger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_judger_v1_judger_proto_goTypes,
		DependencyIndexes: file_judger_v1_judger_proto_depIdxs,
		MessageInfos:      file_judger_v1_judger_proto_msgTypes,
	}.Build()
	File_judger_v1_judger_proto = out.File
	file_judger_v1_judger_proto_rawDesc = nil
	file_judger_v1_judger_proto_goTypes = nil
	file_judger_v1_judger_proto_depIdxs = nil
}
