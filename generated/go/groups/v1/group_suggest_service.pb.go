// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: groups/v1/group_suggest_service.proto

package groups

import (
	v1 "github.com/sokkit-io/xiphias-model-common/generated/go/asset/v1"
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

type GetSuggestedGroupSearchTermsResponse_Result int32

const (
	GetSuggestedGroupSearchTermsResponse_OK GetSuggestedGroupSearchTermsResponse_Result = 0
	// Request denied due to rate limiting. See method description for expected rate limits.
	GetSuggestedGroupSearchTermsResponse_RATE_LIMIT_EXCEEDED GetSuggestedGroupSearchTermsResponse_Result = 1
)

// Enum value maps for GetSuggestedGroupSearchTermsResponse_Result.
var (
	GetSuggestedGroupSearchTermsResponse_Result_name = map[int32]string{
		0: "OK",
		1: "RATE_LIMIT_EXCEEDED",
	}
	GetSuggestedGroupSearchTermsResponse_Result_value = map[string]int32{
		"OK":                  0,
		"RATE_LIMIT_EXCEEDED": 1,
	}
)

func (x GetSuggestedGroupSearchTermsResponse_Result) Enum() *GetSuggestedGroupSearchTermsResponse_Result {
	p := new(GetSuggestedGroupSearchTermsResponse_Result)
	*p = x
	return p
}

func (x GetSuggestedGroupSearchTermsResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetSuggestedGroupSearchTermsResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_groups_v1_group_suggest_service_proto_enumTypes[0].Descriptor()
}

func (GetSuggestedGroupSearchTermsResponse_Result) Type() protoreflect.EnumType {
	return &file_groups_v1_group_suggest_service_proto_enumTypes[0]
}

func (x GetSuggestedGroupSearchTermsResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetSuggestedGroupSearchTermsResponse_Result.Descriptor instead.
func (GetSuggestedGroupSearchTermsResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_groups_v1_group_suggest_service_proto_rawDescGZIP(), []int{1, 0}
}

type GetSuggestedGroupSearchTermsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSuggestedGroupSearchTermsRequest) Reset() {
	*x = GetSuggestedGroupSearchTermsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_v1_group_suggest_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuggestedGroupSearchTermsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuggestedGroupSearchTermsRequest) ProtoMessage() {}

func (x *GetSuggestedGroupSearchTermsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groups_v1_group_suggest_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuggestedGroupSearchTermsRequest.ProtoReflect.Descriptor instead.
func (*GetSuggestedGroupSearchTermsRequest) Descriptor() ([]byte, []int) {
	return file_groups_v1_group_suggest_service_proto_rawDescGZIP(), []int{0}
}

type GetSuggestedGroupSearchTermsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result        GetSuggestedGroupSearchTermsResponse_Result           `protobuf:"varint,1,opt,name=result,proto3,enum=mobile.groups.v1.GetSuggestedGroupSearchTermsResponse_Result" json:"result,omitempty"`
	SuggestedTerm []*GetSuggestedGroupSearchTermsResponse_SuggestedTerm `protobuf:"bytes,2,rep,name=suggested_term,json=suggestedTerm,proto3" json:"suggested_term,omitempty"`
}

func (x *GetSuggestedGroupSearchTermsResponse) Reset() {
	*x = GetSuggestedGroupSearchTermsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_v1_group_suggest_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuggestedGroupSearchTermsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuggestedGroupSearchTermsResponse) ProtoMessage() {}

func (x *GetSuggestedGroupSearchTermsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groups_v1_group_suggest_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuggestedGroupSearchTermsResponse.ProtoReflect.Descriptor instead.
func (*GetSuggestedGroupSearchTermsResponse) Descriptor() ([]byte, []int) {
	return file_groups_v1_group_suggest_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetSuggestedGroupSearchTermsResponse) GetResult() GetSuggestedGroupSearchTermsResponse_Result {
	if x != nil {
		return x.Result
	}
	return GetSuggestedGroupSearchTermsResponse_OK
}

func (x *GetSuggestedGroupSearchTermsResponse) GetSuggestedTerm() []*GetSuggestedGroupSearchTermsResponse_SuggestedTerm {
	if x != nil {
		return x.SuggestedTerm
	}
	return nil
}

type GetSuggestedGroupSearchTermsResponse_SuggestedTerm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A simple string intended as a search term that may be presented to
	// the user as a suggestion to use when searching for groups.  If
	// selected, this term should be provided to the group search service
	// as-is.
	// Note these terms will not have a '#' character prefix, nor any special
	// wildcards. (there are no wildcards included in the spec at this time)
	Term string `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
	// An optional Inner Pic Element
	SuggestedGroupAvatarPic *v1.MediaContent `protobuf:"bytes,2,opt,name=suggested_group_avatar_pic,json=suggestedGroupAvatarPic,proto3" json:"suggested_group_avatar_pic,omitempty"`
	// An optional Suggested Kik Asset Element
	SuggestedGroupKikAsset *v1.MediaContent `protobuf:"bytes,3,opt,name=suggested_group_kik_asset,json=suggestedGroupKikAsset,proto3" json:"suggested_group_kik_asset,omitempty"`
}

func (x *GetSuggestedGroupSearchTermsResponse_SuggestedTerm) Reset() {
	*x = GetSuggestedGroupSearchTermsResponse_SuggestedTerm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_v1_group_suggest_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuggestedGroupSearchTermsResponse_SuggestedTerm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuggestedGroupSearchTermsResponse_SuggestedTerm) ProtoMessage() {}

func (x *GetSuggestedGroupSearchTermsResponse_SuggestedTerm) ProtoReflect() protoreflect.Message {
	mi := &file_groups_v1_group_suggest_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuggestedGroupSearchTermsResponse_SuggestedTerm.ProtoReflect.Descriptor instead.
func (*GetSuggestedGroupSearchTermsResponse_SuggestedTerm) Descriptor() ([]byte, []int) {
	return file_groups_v1_group_suggest_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetSuggestedGroupSearchTermsResponse_SuggestedTerm) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

func (x *GetSuggestedGroupSearchTermsResponse_SuggestedTerm) GetSuggestedGroupAvatarPic() *v1.MediaContent {
	if x != nil {
		return x.SuggestedGroupAvatarPic
	}
	return nil
}

func (x *GetSuggestedGroupSearchTermsResponse_SuggestedTerm) GetSuggestedGroupKikAsset() *v1.MediaContent {
	if x != nil {
		return x.SuggestedGroupKikAsset
	}
	return nil
}

var File_groups_v1_group_suggest_service_proto protoreflect.FileDescriptor

var file_groups_v1_group_suggest_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x25, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9b, 0x04, 0x0a, 0x24, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x55, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x3d, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x76, 0x0a, 0x0e, 0x73, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x44, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x54, 0x65, 0x72, 0x6d, 0x42, 0x09, 0xca, 0x9d, 0x25, 0x05, 0x08, 0x00, 0x80, 0x01,
	0x64, 0x52, 0x0d, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x65, 0x72, 0x6d,
	0x1a, 0xf8, 0x01, 0x0a, 0x0d, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x65,
	0x72, 0x6d, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1d, 0xca, 0x9d, 0x25, 0x19, 0x08, 0x01, 0x12, 0x15, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2e, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x33, 0x32, 0x7d, 0x24, 0x52,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x5a, 0x0a, 0x1a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x17, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x69,
	0x63, 0x12, 0x58, 0x0a, 0x19, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6b, 0x69, 0x6b, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x16, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4b, 0x69, 0x6b, 0x41, 0x73, 0x73, 0x65, 0x74, 0x22, 0x29, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x17, 0x0a,
	0x13, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x45, 0x58, 0x43, 0x45,
	0x45, 0x44, 0x45, 0x44, 0x10, 0x01, 0x32, 0xa0, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x12, 0x8f, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x12, 0x35, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x36, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x57, 0x0a, 0x0e, 0x63, 0x6f, 0x6d,
	0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x5a, 0x45, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d, 0x69,
	0x6f, 0x2f, 0x78, 0x69, 0x70, 0x68, 0x69, 0x61, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67,
	0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_groups_v1_group_suggest_service_proto_rawDescOnce sync.Once
	file_groups_v1_group_suggest_service_proto_rawDescData = file_groups_v1_group_suggest_service_proto_rawDesc
)

func file_groups_v1_group_suggest_service_proto_rawDescGZIP() []byte {
	file_groups_v1_group_suggest_service_proto_rawDescOnce.Do(func() {
		file_groups_v1_group_suggest_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_groups_v1_group_suggest_service_proto_rawDescData)
	})
	return file_groups_v1_group_suggest_service_proto_rawDescData
}

var file_groups_v1_group_suggest_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_groups_v1_group_suggest_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_groups_v1_group_suggest_service_proto_goTypes = []interface{}{
	(GetSuggestedGroupSearchTermsResponse_Result)(0),           // 0: mobile.groups.v1.GetSuggestedGroupSearchTermsResponse.Result
	(*GetSuggestedGroupSearchTermsRequest)(nil),                // 1: mobile.groups.v1.GetSuggestedGroupSearchTermsRequest
	(*GetSuggestedGroupSearchTermsResponse)(nil),               // 2: mobile.groups.v1.GetSuggestedGroupSearchTermsResponse
	(*GetSuggestedGroupSearchTermsResponse_SuggestedTerm)(nil), // 3: mobile.groups.v1.GetSuggestedGroupSearchTermsResponse.SuggestedTerm
	(*v1.MediaContent)(nil),                                    // 4: common.asset.v1.MediaContent
}
var file_groups_v1_group_suggest_service_proto_depIdxs = []int32{
	0, // 0: mobile.groups.v1.GetSuggestedGroupSearchTermsResponse.result:type_name -> mobile.groups.v1.GetSuggestedGroupSearchTermsResponse.Result
	3, // 1: mobile.groups.v1.GetSuggestedGroupSearchTermsResponse.suggested_term:type_name -> mobile.groups.v1.GetSuggestedGroupSearchTermsResponse.SuggestedTerm
	4, // 2: mobile.groups.v1.GetSuggestedGroupSearchTermsResponse.SuggestedTerm.suggested_group_avatar_pic:type_name -> common.asset.v1.MediaContent
	4, // 3: mobile.groups.v1.GetSuggestedGroupSearchTermsResponse.SuggestedTerm.suggested_group_kik_asset:type_name -> common.asset.v1.MediaContent
	1, // 4: mobile.groups.v1.GroupSuggest.GetSuggestedGroupSearchTerms:input_type -> mobile.groups.v1.GetSuggestedGroupSearchTermsRequest
	2, // 5: mobile.groups.v1.GroupSuggest.GetSuggestedGroupSearchTerms:output_type -> mobile.groups.v1.GetSuggestedGroupSearchTermsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_groups_v1_group_suggest_service_proto_init() }
func file_groups_v1_group_suggest_service_proto_init() {
	if File_groups_v1_group_suggest_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_groups_v1_group_suggest_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSuggestedGroupSearchTermsRequest); i {
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
		file_groups_v1_group_suggest_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSuggestedGroupSearchTermsResponse); i {
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
		file_groups_v1_group_suggest_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSuggestedGroupSearchTermsResponse_SuggestedTerm); i {
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
			RawDescriptor: file_groups_v1_group_suggest_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_groups_v1_group_suggest_service_proto_goTypes,
		DependencyIndexes: file_groups_v1_group_suggest_service_proto_depIdxs,
		EnumInfos:         file_groups_v1_group_suggest_service_proto_enumTypes,
		MessageInfos:      file_groups_v1_group_suggest_service_proto_msgTypes,
	}.Build()
	File_groups_v1_group_suggest_service_proto = out.File
	file_groups_v1_group_suggest_service_proto_rawDesc = nil
	file_groups_v1_group_suggest_service_proto_goTypes = nil
	file_groups_v1_group_suggest_service_proto_depIdxs = nil
}
