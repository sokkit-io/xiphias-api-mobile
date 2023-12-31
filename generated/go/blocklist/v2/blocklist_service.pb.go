// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: blocklist/v2/blocklist_service.proto

package blocklist

import (
	v2 "github.com/sokkit-io/xiphias-model-common/generated/go/common/v2"
	_ "github.com/sokkit-io/xiphias-model-common/generated/go/kikoptions"
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

type BlockResponse_Result int32

const (
	BlockResponse_OK BlockResponse_Result = 0
)

// Enum value maps for BlockResponse_Result.
var (
	BlockResponse_Result_name = map[int32]string{
		0: "OK",
	}
	BlockResponse_Result_value = map[string]int32{
		"OK": 0,
	}
)

func (x BlockResponse_Result) Enum() *BlockResponse_Result {
	p := new(BlockResponse_Result)
	*p = x
	return p
}

func (x BlockResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_blocklist_v2_blocklist_service_proto_enumTypes[0].Descriptor()
}

func (BlockResponse_Result) Type() protoreflect.EnumType {
	return &file_blocklist_v2_blocklist_service_proto_enumTypes[0]
}

func (x BlockResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockResponse_Result.Descriptor instead.
func (BlockResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_blocklist_v2_blocklist_service_proto_rawDescGZIP(), []int{1, 0}
}

type UnblockResponse_Result int32

const (
	UnblockResponse_OK UnblockResponse_Result = 0
)

// Enum value maps for UnblockResponse_Result.
var (
	UnblockResponse_Result_name = map[int32]string{
		0: "OK",
	}
	UnblockResponse_Result_value = map[string]int32{
		"OK": 0,
	}
)

func (x UnblockResponse_Result) Enum() *UnblockResponse_Result {
	p := new(UnblockResponse_Result)
	*p = x
	return p
}

func (x UnblockResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UnblockResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_blocklist_v2_blocklist_service_proto_enumTypes[1].Descriptor()
}

func (UnblockResponse_Result) Type() protoreflect.EnumType {
	return &file_blocklist_v2_blocklist_service_proto_enumTypes[1]
}

func (x UnblockResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UnblockResponse_Result.Descriptor instead.
func (UnblockResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_blocklist_v2_blocklist_service_proto_rawDescGZIP(), []int{3, 0}
}

type BlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blockee *v2.PersonaId `protobuf:"bytes,1,opt,name=blockee,proto3" json:"blockee,omitempty"`
}

func (x *BlockRequest) Reset() {
	*x = BlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRequest) ProtoMessage() {}

func (x *BlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRequest.ProtoReflect.Descriptor instead.
func (*BlockRequest) Descriptor() ([]byte, []int) {
	return file_blocklist_v2_blocklist_service_proto_rawDescGZIP(), []int{0}
}

func (x *BlockRequest) GetBlockee() *v2.PersonaId {
	if x != nil {
		return x.Blockee
	}
	return nil
}

type BlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result BlockResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mobile.blocklist.v2.BlockResponse_Result" json:"result,omitempty"`
}

func (x *BlockResponse) Reset() {
	*x = BlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResponse) ProtoMessage() {}

func (x *BlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResponse.ProtoReflect.Descriptor instead.
func (*BlockResponse) Descriptor() ([]byte, []int) {
	return file_blocklist_v2_blocklist_service_proto_rawDescGZIP(), []int{1}
}

func (x *BlockResponse) GetResult() BlockResponse_Result {
	if x != nil {
		return x.Result
	}
	return BlockResponse_OK
}

type UnblockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blockee *v2.PersonaId `protobuf:"bytes,1,opt,name=blockee,proto3" json:"blockee,omitempty"`
}

func (x *UnblockRequest) Reset() {
	*x = UnblockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnblockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnblockRequest) ProtoMessage() {}

func (x *UnblockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnblockRequest.ProtoReflect.Descriptor instead.
func (*UnblockRequest) Descriptor() ([]byte, []int) {
	return file_blocklist_v2_blocklist_service_proto_rawDescGZIP(), []int{2}
}

func (x *UnblockRequest) GetBlockee() *v2.PersonaId {
	if x != nil {
		return x.Blockee
	}
	return nil
}

type UnblockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result UnblockResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mobile.blocklist.v2.UnblockResponse_Result" json:"result,omitempty"`
}

func (x *UnblockResponse) Reset() {
	*x = UnblockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnblockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnblockResponse) ProtoMessage() {}

func (x *UnblockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnblockResponse.ProtoReflect.Descriptor instead.
func (*UnblockResponse) Descriptor() ([]byte, []int) {
	return file_blocklist_v2_blocklist_service_proto_rawDescGZIP(), []int{3}
}

func (x *UnblockResponse) GetResult() UnblockResponse_Result {
	if x != nil {
		return x.Result
	}
	return UnblockResponse_OK
}

type IsBlockedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blockee *v2.PersonaId `protobuf:"bytes,1,opt,name=blockee,proto3" json:"blockee,omitempty"`
}

func (x *IsBlockedRequest) Reset() {
	*x = IsBlockedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsBlockedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsBlockedRequest) ProtoMessage() {}

func (x *IsBlockedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsBlockedRequest.ProtoReflect.Descriptor instead.
func (*IsBlockedRequest) Descriptor() ([]byte, []int) {
	return file_blocklist_v2_blocklist_service_proto_rawDescGZIP(), []int{4}
}

func (x *IsBlockedRequest) GetBlockee() *v2.PersonaId {
	if x != nil {
		return x.Blockee
	}
	return nil
}

// Returns true if the given blocker has blocked the given blockee.
type IsBlockedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsBlocked bool `protobuf:"varint,1,opt,name=is_blocked,json=isBlocked,proto3" json:"is_blocked,omitempty"`
}

func (x *IsBlockedResponse) Reset() {
	*x = IsBlockedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsBlockedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsBlockedResponse) ProtoMessage() {}

func (x *IsBlockedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsBlockedResponse.ProtoReflect.Descriptor instead.
func (*IsBlockedResponse) Descriptor() ([]byte, []int) {
	return file_blocklist_v2_blocklist_service_proto_rawDescGZIP(), []int{5}
}

func (x *IsBlockedResponse) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

// Request a stream of all blocked PersonaIds for a given user
type GetBlocklistStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBlocklistStreamRequest) Reset() {
	*x = GetBlocklistStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocklistStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocklistStreamRequest) ProtoMessage() {}

func (x *GetBlocklistStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlocklistStreamRequest.ProtoReflect.Descriptor instead.
func (*GetBlocklistStreamRequest) Descriptor() ([]byte, []int) {
	return file_blocklist_v2_blocklist_service_proto_rawDescGZIP(), []int{6}
}

// Contains a page of a blocklist. Clients should accumulate a list of blockees from the stream to receive the
// entire blocklist.
type GetBlocklistStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Page of blockees
	Blockees []*BlocklistEntry `protobuf:"bytes,1,rep,name=blockees,proto3" json:"blockees,omitempty"`
}

func (x *GetBlocklistStreamResponse) Reset() {
	*x = GetBlocklistStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocklistStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocklistStreamResponse) ProtoMessage() {}

func (x *GetBlocklistStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlocklistStreamResponse.ProtoReflect.Descriptor instead.
func (*GetBlocklistStreamResponse) Descriptor() ([]byte, []int) {
	return file_blocklist_v2_blocklist_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetBlocklistStreamResponse) GetBlockees() []*BlocklistEntry {
	if x != nil {
		return x.Blockees
	}
	return nil
}

type BlocklistEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blockee      *v2.PersonaId `protobuf:"bytes,1,opt,name=blockee,proto3" json:"blockee,omitempty"`
	DoNotDisplay bool          `protobuf:"varint,2,opt,name=do_not_display,json=doNotDisplay,proto3" json:"do_not_display,omitempty"`
}

func (x *BlocklistEntry) Reset() {
	*x = BlocklistEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlocklistEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlocklistEntry) ProtoMessage() {}

func (x *BlocklistEntry) ProtoReflect() protoreflect.Message {
	mi := &file_blocklist_v2_blocklist_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlocklistEntry.ProtoReflect.Descriptor instead.
func (*BlocklistEntry) Descriptor() ([]byte, []int) {
	return file_blocklist_v2_blocklist_service_proto_rawDescGZIP(), []int{8}
}

func (x *BlocklistEntry) GetBlockee() *v2.PersonaId {
	if x != nil {
		return x.Blockee
	}
	return nil
}

func (x *BlocklistEntry) GetDoNotDisplay() bool {
	if x != nil {
		return x.DoNotDisplay
	}
	return false
}

var File_blocklist_v2_blocklist_service_proto protoreflect.FileDescriptor

var file_blocklist_v2_blocklist_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x19, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x32, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a,
	0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x49, 0x64, 0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08, 0x01, 0x52, 0x07, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x65, 0x22, 0x64, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x10, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x22, 0x48, 0x0a, 0x0e, 0x55,
	0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x49, 0x64, 0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08, 0x01, 0x52, 0x07, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x65, 0x22, 0x68, 0x0a, 0x0f, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x55,
	0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x10, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x22,
	0x4a, 0x0a, 0x10, 0x49, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x49, 0x64, 0x42, 0x06, 0xca, 0x9d, 0x25, 0x02,
	0x08, 0x01, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x65, 0x22, 0x32, 0x0a, 0x11, 0x49,
	0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22,
	0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x69, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x0a, 0xca, 0x9d, 0x25, 0x06, 0x78, 0x00, 0x80, 0x01, 0xfa, 0x01, 0x52, 0x08, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x65, 0x73, 0x22, 0x6e, 0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x36, 0x0a, 0x07, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x49, 0x64,
	0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08, 0x01, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x65, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x6f, 0x4e, 0x6f, 0x74,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x32, 0x86, 0x03, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x21,
	0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x07, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x23, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x6e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x09, 0x49,
	0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x25, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x49,
	0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x77, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2e, 0x2e,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x42, 0x67, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d,
	0x69, 0x6f, 0x2f, 0x78, 0x69, 0x70, 0x68, 0x69, 0x61, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x67, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x32, 0x3b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_blocklist_v2_blocklist_service_proto_rawDescOnce sync.Once
	file_blocklist_v2_blocklist_service_proto_rawDescData = file_blocklist_v2_blocklist_service_proto_rawDesc
)

func file_blocklist_v2_blocklist_service_proto_rawDescGZIP() []byte {
	file_blocklist_v2_blocklist_service_proto_rawDescOnce.Do(func() {
		file_blocklist_v2_blocklist_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_blocklist_v2_blocklist_service_proto_rawDescData)
	})
	return file_blocklist_v2_blocklist_service_proto_rawDescData
}

var file_blocklist_v2_blocklist_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_blocklist_v2_blocklist_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_blocklist_v2_blocklist_service_proto_goTypes = []interface{}{
	(BlockResponse_Result)(0),          // 0: mobile.blocklist.v2.BlockResponse.Result
	(UnblockResponse_Result)(0),        // 1: mobile.blocklist.v2.UnblockResponse.Result
	(*BlockRequest)(nil),               // 2: mobile.blocklist.v2.BlockRequest
	(*BlockResponse)(nil),              // 3: mobile.blocklist.v2.BlockResponse
	(*UnblockRequest)(nil),             // 4: mobile.blocklist.v2.UnblockRequest
	(*UnblockResponse)(nil),            // 5: mobile.blocklist.v2.UnblockResponse
	(*IsBlockedRequest)(nil),           // 6: mobile.blocklist.v2.IsBlockedRequest
	(*IsBlockedResponse)(nil),          // 7: mobile.blocklist.v2.IsBlockedResponse
	(*GetBlocklistStreamRequest)(nil),  // 8: mobile.blocklist.v2.GetBlocklistStreamRequest
	(*GetBlocklistStreamResponse)(nil), // 9: mobile.blocklist.v2.GetBlocklistStreamResponse
	(*BlocklistEntry)(nil),             // 10: mobile.blocklist.v2.BlocklistEntry
	(*v2.PersonaId)(nil),               // 11: common.v2.PersonaId
}
var file_blocklist_v2_blocklist_service_proto_depIdxs = []int32{
	11, // 0: mobile.blocklist.v2.BlockRequest.blockee:type_name -> common.v2.PersonaId
	0,  // 1: mobile.blocklist.v2.BlockResponse.result:type_name -> mobile.blocklist.v2.BlockResponse.Result
	11, // 2: mobile.blocklist.v2.UnblockRequest.blockee:type_name -> common.v2.PersonaId
	1,  // 3: mobile.blocklist.v2.UnblockResponse.result:type_name -> mobile.blocklist.v2.UnblockResponse.Result
	11, // 4: mobile.blocklist.v2.IsBlockedRequest.blockee:type_name -> common.v2.PersonaId
	10, // 5: mobile.blocklist.v2.GetBlocklistStreamResponse.blockees:type_name -> mobile.blocklist.v2.BlocklistEntry
	11, // 6: mobile.blocklist.v2.BlocklistEntry.blockee:type_name -> common.v2.PersonaId
	2,  // 7: mobile.blocklist.v2.Blocklist.Block:input_type -> mobile.blocklist.v2.BlockRequest
	4,  // 8: mobile.blocklist.v2.Blocklist.Unblock:input_type -> mobile.blocklist.v2.UnblockRequest
	6,  // 9: mobile.blocklist.v2.Blocklist.IsBlocked:input_type -> mobile.blocklist.v2.IsBlockedRequest
	8,  // 10: mobile.blocklist.v2.Blocklist.GetBlocklistStream:input_type -> mobile.blocklist.v2.GetBlocklistStreamRequest
	3,  // 11: mobile.blocklist.v2.Blocklist.Block:output_type -> mobile.blocklist.v2.BlockResponse
	5,  // 12: mobile.blocklist.v2.Blocklist.Unblock:output_type -> mobile.blocklist.v2.UnblockResponse
	7,  // 13: mobile.blocklist.v2.Blocklist.IsBlocked:output_type -> mobile.blocklist.v2.IsBlockedResponse
	9,  // 14: mobile.blocklist.v2.Blocklist.GetBlocklistStream:output_type -> mobile.blocklist.v2.GetBlocklistStreamResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_blocklist_v2_blocklist_service_proto_init() }
func file_blocklist_v2_blocklist_service_proto_init() {
	if File_blocklist_v2_blocklist_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blocklist_v2_blocklist_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRequest); i {
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
		file_blocklist_v2_blocklist_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResponse); i {
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
		file_blocklist_v2_blocklist_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnblockRequest); i {
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
		file_blocklist_v2_blocklist_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnblockResponse); i {
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
		file_blocklist_v2_blocklist_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsBlockedRequest); i {
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
		file_blocklist_v2_blocklist_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsBlockedResponse); i {
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
		file_blocklist_v2_blocklist_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlocklistStreamRequest); i {
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
		file_blocklist_v2_blocklist_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlocklistStreamResponse); i {
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
		file_blocklist_v2_blocklist_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlocklistEntry); i {
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
			RawDescriptor: file_blocklist_v2_blocklist_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blocklist_v2_blocklist_service_proto_goTypes,
		DependencyIndexes: file_blocklist_v2_blocklist_service_proto_depIdxs,
		EnumInfos:         file_blocklist_v2_blocklist_service_proto_enumTypes,
		MessageInfos:      file_blocklist_v2_blocklist_service_proto_msgTypes,
	}.Build()
	File_blocklist_v2_blocklist_service_proto = out.File
	file_blocklist_v2_blocklist_service_proto_rawDesc = nil
	file_blocklist_v2_blocklist_service_proto_goTypes = nil
	file_blocklist_v2_blocklist_service_proto_depIdxs = nil
}
