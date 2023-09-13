// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: video/v1/kik_video_service.proto

package video

import (
	_go "github.com/sokkit-io/xiphias-model-common/generated/go"
	_ "github.com/sokkit-io/xiphias-model-common/generated/go/kikoptions"
	v1 "github.com/sokkit-io/xiphias-model-common/generated/go/video/v1"
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

type BatchedGetConvoStateResponse_Result int32

const (
	BatchedGetConvoStateResponse_OK BatchedGetConvoStateResponse_Result = 0
)

// Enum value maps for BatchedGetConvoStateResponse_Result.
var (
	BatchedGetConvoStateResponse_Result_name = map[int32]string{
		0: "OK",
	}
	BatchedGetConvoStateResponse_Result_value = map[string]int32{
		"OK": 0,
	}
)

func (x BatchedGetConvoStateResponse_Result) Enum() *BatchedGetConvoStateResponse_Result {
	p := new(BatchedGetConvoStateResponse_Result)
	*p = x
	return p
}

func (x BatchedGetConvoStateResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BatchedGetConvoStateResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_video_v1_kik_video_service_proto_enumTypes[0].Descriptor()
}

func (BatchedGetConvoStateResponse_Result) Type() protoreflect.EnumType {
	return &file_video_v1_kik_video_service_proto_enumTypes[0]
}

func (x BatchedGetConvoStateResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BatchedGetConvoStateResponse_Result.Descriptor instead.
func (BatchedGetConvoStateResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_video_v1_kik_video_service_proto_rawDescGZIP(), []int{1, 0}
}

type JoinConvoConferenceResponse_Result int32

const (
	// Note that clients MUST check for an OK result as new error results may be added in the future.
	JoinConvoConferenceResponse_OK JoinConvoConferenceResponse_Result = 0
	// The conference has reached its maximum number of participants
	JoinConvoConferenceResponse_FULL JoinConvoConferenceResponse_Result = 1
	// The user is not allowed to join the conference. This MAY be caused by (but is not limited
	// to), not being a member of the group.
	JoinConvoConferenceResponse_NOT_ALLOWED JoinConvoConferenceResponse_Result = 2
)

// Enum value maps for JoinConvoConferenceResponse_Result.
var (
	JoinConvoConferenceResponse_Result_name = map[int32]string{
		0: "OK",
		1: "FULL",
		2: "NOT_ALLOWED",
	}
	JoinConvoConferenceResponse_Result_value = map[string]int32{
		"OK":          0,
		"FULL":        1,
		"NOT_ALLOWED": 2,
	}
)

func (x JoinConvoConferenceResponse_Result) Enum() *JoinConvoConferenceResponse_Result {
	p := new(JoinConvoConferenceResponse_Result)
	*p = x
	return p
}

func (x JoinConvoConferenceResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JoinConvoConferenceResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_video_v1_kik_video_service_proto_enumTypes[1].Descriptor()
}

func (JoinConvoConferenceResponse_Result) Type() protoreflect.EnumType {
	return &file_video_v1_kik_video_service_proto_enumTypes[1]
}

func (x JoinConvoConferenceResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JoinConvoConferenceResponse_Result.Descriptor instead.
func (JoinConvoConferenceResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_video_v1_kik_video_service_proto_rawDescGZIP(), []int{3, 0}
}

type BatchedGetConvoStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConvoIds []*v1.ConvoId `protobuf:"bytes,1,rep,name=convo_ids,json=convoIds,proto3" json:"convo_ids,omitempty"`
}

func (x *BatchedGetConvoStateRequest) Reset() {
	*x = BatchedGetConvoStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_kik_video_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchedGetConvoStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchedGetConvoStateRequest) ProtoMessage() {}

func (x *BatchedGetConvoStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_kik_video_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchedGetConvoStateRequest.ProtoReflect.Descriptor instead.
func (*BatchedGetConvoStateRequest) Descriptor() ([]byte, []int) {
	return file_video_v1_kik_video_service_proto_rawDescGZIP(), []int{0}
}

func (x *BatchedGetConvoStateRequest) GetConvoIds() []*v1.ConvoId {
	if x != nil {
		return x.ConvoIds
	}
	return nil
}

type BatchedGetConvoStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result BatchedGetConvoStateResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mobile.video.v1.BatchedGetConvoStateResponse_Result" json:"result,omitempty"`
	// This list SHOULD be the exact same set of convo ids requested (but order is not guaranteed).
	// A ConvoVideoState message with no users listed SHOULD be returned if there is no conference
	// associated with the convo.
	LatestStates []*v1.ConvoVideoState `protobuf:"bytes,2,rep,name=latest_states,json=latestStates,proto3" json:"latest_states,omitempty"`
}

func (x *BatchedGetConvoStateResponse) Reset() {
	*x = BatchedGetConvoStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_kik_video_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchedGetConvoStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchedGetConvoStateResponse) ProtoMessage() {}

func (x *BatchedGetConvoStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_kik_video_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchedGetConvoStateResponse.ProtoReflect.Descriptor instead.
func (*BatchedGetConvoStateResponse) Descriptor() ([]byte, []int) {
	return file_video_v1_kik_video_service_proto_rawDescGZIP(), []int{1}
}

func (x *BatchedGetConvoStateResponse) GetResult() BatchedGetConvoStateResponse_Result {
	if x != nil {
		return x.Result
	}
	return BatchedGetConvoStateResponse_OK
}

func (x *BatchedGetConvoStateResponse) GetLatestStates() []*v1.ConvoVideoState {
	if x != nil {
		return x.LatestStates
	}
	return nil
}

type JoinConvoConferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConvoId *v1.ConvoId `protobuf:"bytes,1,opt,name=convo_id,json=convoId,proto3" json:"convo_id,omitempty"`
}

func (x *JoinConvoConferenceRequest) Reset() {
	*x = JoinConvoConferenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_kik_video_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinConvoConferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinConvoConferenceRequest) ProtoMessage() {}

func (x *JoinConvoConferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_kik_video_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinConvoConferenceRequest.ProtoReflect.Descriptor instead.
func (*JoinConvoConferenceRequest) Descriptor() ([]byte, []int) {
	return file_video_v1_kik_video_service_proto_rawDescGZIP(), []int{2}
}

func (x *JoinConvoConferenceRequest) GetConvoId() *v1.ConvoId {
	if x != nil {
		return x.ConvoId
	}
	return nil
}

type JoinConvoConferenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result JoinConvoConferenceResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mobile.video.v1.JoinConvoConferenceResponse_Result" json:"result,omitempty"`
	// This MUST be present for the following results, but MAY be absent otherwise:
	//  -   OK
	//  -   FULL
	LatestState *v1.ConvoVideoState `protobuf:"bytes,2,opt,name=latest_state,json=latestState,proto3" json:"latest_state,omitempty"`
	// This MAY be null if the user is NOT allowed to join the video
	ConnectionInfo *v1.ConferenceConnectionInfo `protobuf:"bytes,3,opt,name=connection_info,json=connectionInfo,proto3" json:"connection_info,omitempty"`
	// This token must be provided to the media server when the client connects to it.
	// This SHOULD be present if connection_info is present.
	ConnectionToken *v1.MediaServerConnectionToken `protobuf:"bytes,4,opt,name=connection_token,json=connectionToken,proto3" json:"connection_token,omitempty"`
}

func (x *JoinConvoConferenceResponse) Reset() {
	*x = JoinConvoConferenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_kik_video_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinConvoConferenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinConvoConferenceResponse) ProtoMessage() {}

func (x *JoinConvoConferenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_kik_video_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinConvoConferenceResponse.ProtoReflect.Descriptor instead.
func (*JoinConvoConferenceResponse) Descriptor() ([]byte, []int) {
	return file_video_v1_kik_video_service_proto_rawDescGZIP(), []int{3}
}

func (x *JoinConvoConferenceResponse) GetResult() JoinConvoConferenceResponse_Result {
	if x != nil {
		return x.Result
	}
	return JoinConvoConferenceResponse_OK
}

func (x *JoinConvoConferenceResponse) GetLatestState() *v1.ConvoVideoState {
	if x != nil {
		return x.LatestState
	}
	return nil
}

func (x *JoinConvoConferenceResponse) GetConnectionInfo() *v1.ConferenceConnectionInfo {
	if x != nil {
		return x.ConnectionInfo
	}
	return nil
}

func (x *JoinConvoConferenceResponse) GetConnectionToken() *v1.MediaServerConnectionToken {
	if x != nil {
		return x.ConnectionToken
	}
	return nil
}

type LeaveConvoConferenceNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConvoId *v1.ConvoId `protobuf:"bytes,1,opt,name=convo_id,json=convoId,proto3" json:"convo_id,omitempty"`
	// This is the token returned to the client through this KikVideo service via JoinConvoConferenceFromVcsResponse
	// The client should cache this token and ALWAYS pass this token for onLeave.
	// This token is intended for the rounds' video conference server to sort out race conditions. For example
	// client turns the video on/off too quickly and the leave may arrive after the rejoin. In this case, the sequenctial nature
	// of the token will block the previous onLeave (note the client on leave event send leave via kikvideo and also requests disconnect
	// from the media server directly).
	ConnectionToken *v1.MediaServerConnectionToken `protobuf:"bytes,4,opt,name=connection_token,json=connectionToken,proto3" json:"connection_token,omitempty"`
}

func (x *LeaveConvoConferenceNotification) Reset() {
	*x = LeaveConvoConferenceNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_kik_video_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveConvoConferenceNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveConvoConferenceNotification) ProtoMessage() {}

func (x *LeaveConvoConferenceNotification) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_kik_video_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveConvoConferenceNotification.ProtoReflect.Descriptor instead.
func (*LeaveConvoConferenceNotification) Descriptor() ([]byte, []int) {
	return file_video_v1_kik_video_service_proto_rawDescGZIP(), []int{4}
}

func (x *LeaveConvoConferenceNotification) GetConvoId() *v1.ConvoId {
	if x != nil {
		return x.ConvoId
	}
	return nil
}

func (x *LeaveConvoConferenceNotification) GetConnectionToken() *v1.MediaServerConnectionToken {
	if x != nil {
		return x.ConnectionToken
	}
	return nil
}

var File_video_v1_kik_video_service_proto protoreflect.FileDescriptor

var file_video_v1_kik_video_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x69, 0x6b, 0x5f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x76, 0x31, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a,
	0x1b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x6f,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x76, 0x6f, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x49, 0x64, 0x42, 0x0b, 0xca, 0x9d, 0x25, 0x07, 0x08,
	0x01, 0x78, 0x01, 0x80, 0x01, 0x14, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x76, 0x6f, 0x49, 0x64, 0x73,
	0x22, 0xd0, 0x01, 0x0a, 0x1c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x76, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x34, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x76, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x50, 0x0a, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x09, 0xca, 0x9d, 0x25, 0x05, 0x08, 0x00,
	0x80, 0x01, 0x14, 0x52, 0x0c, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x22, 0x10, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x00, 0x22, 0x5f, 0x0a, 0x1a, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x6f,
	0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3b, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x76, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x49, 0x64, 0x42, 0x06, 0xca,
	0x9d, 0x25, 0x02, 0x08, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x76, 0x6f, 0x49, 0x64, 0x4a, 0x04,
	0x08, 0x02, 0x10, 0x03, 0x22, 0x88, 0x03, 0x0a, 0x1b, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e,
	0x76, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x6f,
	0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x43, 0x0a, 0x0c, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x52, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x56, 0x0a, 0x10, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x2b, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02,
	0x4f, 0x4b, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x02, 0x22,
	0xc3, 0x01, 0x0a, 0x20, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x43, 0x6f,
	0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x76, 0x6f, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x49, 0x64,
	0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x76, 0x6f, 0x49,
	0x64, 0x12, 0x56, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a,
	0x04, 0x08, 0x03, 0x10, 0x04, 0x32, 0xd4, 0x02, 0x0a, 0x08, 0x4b, 0x69, 0x6b, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x12, 0x73, 0x0a, 0x14, 0x42, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x64, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x64, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x13, 0x4a, 0x6f, 0x69, 0x6e, 0x43,
	0x6f, 0x6e, 0x76, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2b,
	0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x16, 0x4f, 0x6e, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x31, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x6f,
	0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x5b, 0x0a, 0x14,
	0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x78, 0x69, 0x70, 0x68, 0x69,
	0x61, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_video_v1_kik_video_service_proto_rawDescOnce sync.Once
	file_video_v1_kik_video_service_proto_rawDescData = file_video_v1_kik_video_service_proto_rawDesc
)

func file_video_v1_kik_video_service_proto_rawDescGZIP() []byte {
	file_video_v1_kik_video_service_proto_rawDescOnce.Do(func() {
		file_video_v1_kik_video_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_v1_kik_video_service_proto_rawDescData)
	})
	return file_video_v1_kik_video_service_proto_rawDescData
}

var file_video_v1_kik_video_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_video_v1_kik_video_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_video_v1_kik_video_service_proto_goTypes = []interface{}{
	(BatchedGetConvoStateResponse_Result)(0), // 0: mobile.video.v1.BatchedGetConvoStateResponse.Result
	(JoinConvoConferenceResponse_Result)(0),  // 1: mobile.video.v1.JoinConvoConferenceResponse.Result
	(*BatchedGetConvoStateRequest)(nil),      // 2: mobile.video.v1.BatchedGetConvoStateRequest
	(*BatchedGetConvoStateResponse)(nil),     // 3: mobile.video.v1.BatchedGetConvoStateResponse
	(*JoinConvoConferenceRequest)(nil),       // 4: mobile.video.v1.JoinConvoConferenceRequest
	(*JoinConvoConferenceResponse)(nil),      // 5: mobile.video.v1.JoinConvoConferenceResponse
	(*LeaveConvoConferenceNotification)(nil), // 6: mobile.video.v1.LeaveConvoConferenceNotification
	(*v1.ConvoId)(nil),                       // 7: common.video.v1.ConvoId
	(*v1.ConvoVideoState)(nil),               // 8: common.video.v1.ConvoVideoState
	(*v1.ConferenceConnectionInfo)(nil),      // 9: common.video.v1.ConferenceConnectionInfo
	(*v1.MediaServerConnectionToken)(nil),    // 10: common.video.v1.MediaServerConnectionToken
	(*_go.VoidResponse)(nil),                 // 11: common.VoidResponse
}
var file_video_v1_kik_video_service_proto_depIdxs = []int32{
	7,  // 0: mobile.video.v1.BatchedGetConvoStateRequest.convo_ids:type_name -> common.video.v1.ConvoId
	0,  // 1: mobile.video.v1.BatchedGetConvoStateResponse.result:type_name -> mobile.video.v1.BatchedGetConvoStateResponse.Result
	8,  // 2: mobile.video.v1.BatchedGetConvoStateResponse.latest_states:type_name -> common.video.v1.ConvoVideoState
	7,  // 3: mobile.video.v1.JoinConvoConferenceRequest.convo_id:type_name -> common.video.v1.ConvoId
	1,  // 4: mobile.video.v1.JoinConvoConferenceResponse.result:type_name -> mobile.video.v1.JoinConvoConferenceResponse.Result
	8,  // 5: mobile.video.v1.JoinConvoConferenceResponse.latest_state:type_name -> common.video.v1.ConvoVideoState
	9,  // 6: mobile.video.v1.JoinConvoConferenceResponse.connection_info:type_name -> common.video.v1.ConferenceConnectionInfo
	10, // 7: mobile.video.v1.JoinConvoConferenceResponse.connection_token:type_name -> common.video.v1.MediaServerConnectionToken
	7,  // 8: mobile.video.v1.LeaveConvoConferenceNotification.convo_id:type_name -> common.video.v1.ConvoId
	10, // 9: mobile.video.v1.LeaveConvoConferenceNotification.connection_token:type_name -> common.video.v1.MediaServerConnectionToken
	2,  // 10: mobile.video.v1.KikVideo.BatchedGetConvoState:input_type -> mobile.video.v1.BatchedGetConvoStateRequest
	4,  // 11: mobile.video.v1.KikVideo.JoinConvoConference:input_type -> mobile.video.v1.JoinConvoConferenceRequest
	6,  // 12: mobile.video.v1.KikVideo.OnLeaveConvoConference:input_type -> mobile.video.v1.LeaveConvoConferenceNotification
	3,  // 13: mobile.video.v1.KikVideo.BatchedGetConvoState:output_type -> mobile.video.v1.BatchedGetConvoStateResponse
	5,  // 14: mobile.video.v1.KikVideo.JoinConvoConference:output_type -> mobile.video.v1.JoinConvoConferenceResponse
	11, // 15: mobile.video.v1.KikVideo.OnLeaveConvoConference:output_type -> common.VoidResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_video_v1_kik_video_service_proto_init() }
func file_video_v1_kik_video_service_proto_init() {
	if File_video_v1_kik_video_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_v1_kik_video_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchedGetConvoStateRequest); i {
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
		file_video_v1_kik_video_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchedGetConvoStateResponse); i {
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
		file_video_v1_kik_video_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinConvoConferenceRequest); i {
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
		file_video_v1_kik_video_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinConvoConferenceResponse); i {
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
		file_video_v1_kik_video_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveConvoConferenceNotification); i {
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
			RawDescriptor: file_video_v1_kik_video_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_v1_kik_video_service_proto_goTypes,
		DependencyIndexes: file_video_v1_kik_video_service_proto_depIdxs,
		EnumInfos:         file_video_v1_kik_video_service_proto_enumTypes,
		MessageInfos:      file_video_v1_kik_video_service_proto_msgTypes,
	}.Build()
	File_video_v1_kik_video_service_proto = out.File
	file_video_v1_kik_video_service_proto_rawDesc = nil
	file_video_v1_kik_video_service_proto_goTypes = nil
	file_video_v1_kik_video_service_proto_depIdxs = nil
}
