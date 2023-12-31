// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: config/v1/feature_config_service.proto

package config

import (
	duration "github.com/golang/protobuf/ptypes/duration"
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

// GetFeatureConfigRequest is an empty request, that may contain fields in the
// future.
//
// The service _may_ use caller information provided by the Xiphias Bridge via
// headers.
type GetFeatureConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFeatureConfigsRequest) Reset() {
	*x = GetFeatureConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1_feature_config_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeatureConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureConfigsRequest) ProtoMessage() {}

func (x *GetFeatureConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_feature_config_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureConfigsRequest.ProtoReflect.Descriptor instead.
func (*GetFeatureConfigsRequest) Descriptor() ([]byte, []int) {
	return file_config_v1_feature_config_service_proto_rawDescGZIP(), []int{0}
}

// GetFeatureConfigResponse contains a set of configurations that _may_ be
// tailored to the caller.
type GetFeatureConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The default maximum size for both public and private groups.
	MaxGroupSize uint32 `protobuf:"varint,10,opt,name=max_group_size,json=maxGroupSize,proto3" json:"max_group_size,omitempty"`
	// Switch for ProfileService updating a user's legacy roster last updated
	// timestamp and legacy profile pic timestamp when that user's UserRosterEntry
	// entity is updated.
	// When this switch is true, the client should retrieve a user's UserRosterEntry
	// entity whenever that user's profile pic timestamp is updated in their legacy
	// roster info.
	UserRosterEntryUpdatesLegacyRosterTimestamp bool `protobuf:"varint,11,opt,name=user_roster_entry_updates_legacy_roster_timestamp,json=userRosterEntryUpdatesLegacyRosterTimestamp,proto3" json:"user_roster_entry_updates_legacy_roster_timestamp,omitempty"`
	// The following boolean are using by the client and server.
	// It controls (enable/disables) the call to EntityService.GetTrustedBots
	// which we know will be obsolete after Harmony 3 is in place and is_trusted
	// is supported properly via harmony roster subscription sync
	EntityServiceGetTrustedBotsEnabled bool `protobuf:"varint,12,opt,name=entity_service_get_trusted_bots_enabled,json=entityServiceGetTrustedBotsEnabled,proto3" json:"entity_service_get_trusted_bots_enabled,omitempty"`
	// When the entity_service_get_trusted_bots_enabled is enabled, this is the duration to
	// lapse before the client should invoke another call to EntityService.getTrustedBots
	MinDurationBetweenPullEntityServiceGetTrustedBots *duration.Duration `protobuf:"bytes,13,opt,name=min_duration_between_pull_entity_service_get_trusted_bots,json=minDurationBetweenPullEntityServiceGetTrustedBots,proto3" json:"min_duration_between_pull_entity_service_get_trusted_bots,omitempty"`
	// The max number of interests an user can select
	MaxUserInterests uint32 `protobuf:"varint,14,opt,name=max_user_interests,json=maxUserInterests,proto3" json:"max_user_interests,omitempty"`
	// The following is the current list of user interests that will be offered to the user
	// for selection.
	// Verbiage returned is localized using the connecting device locale.
	// Before locale is fully supported, English is returned.
	CurrentUserInterestList []*UserInterestListItem `protobuf:"bytes,15,rep,name=current_user_interest_list,json=currentUserInterestList,proto3" json:"current_user_interest_list,omitempty"`
	// The list of chat interests
	// Verbiage is what mobile client will display to the user
	// When user picks the interests to chat about, the interest ids
	// are attached to findPartner API.
	CurrentChatInterestList []*ChatInterestListItem `protobuf:"bytes,16,rep,name=current_chat_interest_list,json=currentChatInterestList,proto3" json:"current_chat_interest_list,omitempty"`
}

func (x *GetFeatureConfigsResponse) Reset() {
	*x = GetFeatureConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1_feature_config_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeatureConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureConfigsResponse) ProtoMessage() {}

func (x *GetFeatureConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_feature_config_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureConfigsResponse.ProtoReflect.Descriptor instead.
func (*GetFeatureConfigsResponse) Descriptor() ([]byte, []int) {
	return file_config_v1_feature_config_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetFeatureConfigsResponse) GetMaxGroupSize() uint32 {
	if x != nil {
		return x.MaxGroupSize
	}
	return 0
}

func (x *GetFeatureConfigsResponse) GetUserRosterEntryUpdatesLegacyRosterTimestamp() bool {
	if x != nil {
		return x.UserRosterEntryUpdatesLegacyRosterTimestamp
	}
	return false
}

func (x *GetFeatureConfigsResponse) GetEntityServiceGetTrustedBotsEnabled() bool {
	if x != nil {
		return x.EntityServiceGetTrustedBotsEnabled
	}
	return false
}

func (x *GetFeatureConfigsResponse) GetMinDurationBetweenPullEntityServiceGetTrustedBots() *duration.Duration {
	if x != nil {
		return x.MinDurationBetweenPullEntityServiceGetTrustedBots
	}
	return nil
}

func (x *GetFeatureConfigsResponse) GetMaxUserInterests() uint32 {
	if x != nil {
		return x.MaxUserInterests
	}
	return 0
}

func (x *GetFeatureConfigsResponse) GetCurrentUserInterestList() []*UserInterestListItem {
	if x != nil {
		return x.CurrentUserInterestList
	}
	return nil
}

func (x *GetFeatureConfigsResponse) GetCurrentChatInterestList() []*ChatInterestListItem {
	if x != nil {
		return x.CurrentChatInterestList
	}
	return nil
}

type UserInterestListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LocalizedVerbiage string `protobuf:"bytes,2,opt,name=localized_verbiage,json=localizedVerbiage,proto3" json:"localized_verbiage,omitempty"`
}

func (x *UserInterestListItem) Reset() {
	*x = UserInterestListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1_feature_config_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInterestListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInterestListItem) ProtoMessage() {}

func (x *UserInterestListItem) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_feature_config_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInterestListItem.ProtoReflect.Descriptor instead.
func (*UserInterestListItem) Descriptor() ([]byte, []int) {
	return file_config_v1_feature_config_service_proto_rawDescGZIP(), []int{2}
}

func (x *UserInterestListItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInterestListItem) GetLocalizedVerbiage() string {
	if x != nil {
		return x.LocalizedVerbiage
	}
	return ""
}

type ChatInterestListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LocalizedInterest string `protobuf:"bytes,2,opt,name=localized_interest,json=localizedInterest,proto3" json:"localized_interest,omitempty"`
	// Optional and only returned when used internally
	// via GetAllChatInterests rpc
	LocalizedIceBreaker string `protobuf:"bytes,3,opt,name=localized_ice_breaker,json=localizedIceBreaker,proto3" json:"localized_ice_breaker,omitempty"`
	// This is a unicode emoji string used by clients to display emojis related to this chat interest
	// There could be 0 or more emojis in this string
	// The emojis (if present) will be in javascript format \uXXX
	// see: https://unicode.org/emoji/charts/full-emoji-list.html
	InterestEmoji string `protobuf:"bytes,4,opt,name=interest_emoji,json=interestEmoji,proto3" json:"interest_emoji,omitempty"`
}

func (x *ChatInterestListItem) Reset() {
	*x = ChatInterestListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1_feature_config_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInterestListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInterestListItem) ProtoMessage() {}

func (x *ChatInterestListItem) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_feature_config_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInterestListItem.ProtoReflect.Descriptor instead.
func (*ChatInterestListItem) Descriptor() ([]byte, []int) {
	return file_config_v1_feature_config_service_proto_rawDescGZIP(), []int{3}
}

func (x *ChatInterestListItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChatInterestListItem) GetLocalizedInterest() string {
	if x != nil {
		return x.LocalizedInterest
	}
	return ""
}

func (x *ChatInterestListItem) GetLocalizedIceBreaker() string {
	if x != nil {
		return x.LocalizedIceBreaker
	}
	return ""
}

func (x *ChatInterestListItem) GetInterestEmoji() string {
	if x != nil {
		return x.InterestEmoji
	}
	return ""
}

// GetAllUserInterestIdsRequest is empty for now, but might include fields in the future
type GetAllUserInterestIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllUserInterestIdsRequest) Reset() {
	*x = GetAllUserInterestIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1_feature_config_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserInterestIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserInterestIdsRequest) ProtoMessage() {}

func (x *GetAllUserInterestIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_feature_config_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserInterestIdsRequest.ProtoReflect.Descriptor instead.
func (*GetAllUserInterestIdsRequest) Descriptor() ([]byte, []int) {
	return file_config_v1_feature_config_service_proto_rawDescGZIP(), []int{4}
}

// This list includes the current interests that an user can pick if he is to update this profile
// This also include all the interests we ever had offered to the client
//
// Don't see the need to return page tokens, when do will add it to the request object and will
// return the page token here
type GetAllUserInterestIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetAllUserInterestIdsResponse) Reset() {
	*x = GetAllUserInterestIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1_feature_config_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserInterestIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserInterestIdsResponse) ProtoMessage() {}

func (x *GetAllUserInterestIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_feature_config_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserInterestIdsResponse.ProtoReflect.Descriptor instead.
func (*GetAllUserInterestIdsResponse) Descriptor() ([]byte, []int) {
	return file_config_v1_feature_config_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllUserInterestIdsResponse) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetAllChatInterestsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllChatInterestsRequest) Reset() {
	*x = GetAllChatInterestsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1_feature_config_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllChatInterestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllChatInterestsRequest) ProtoMessage() {}

func (x *GetAllChatInterestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_feature_config_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllChatInterestsRequest.ProtoReflect.Descriptor instead.
func (*GetAllChatInterestsRequest) Descriptor() ([]byte, []int) {
	return file_config_v1_feature_config_service_proto_rawDescGZIP(), []int{6}
}

// This list includes all the currently pickable and interests picked for all
// Users in the waiting queue.
type GetAllChatInterestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatInterests []*ChatInterestListItem `protobuf:"bytes,1,rep,name=chat_interests,json=chatInterests,proto3" json:"chat_interests,omitempty"`
}

func (x *GetAllChatInterestsResponse) Reset() {
	*x = GetAllChatInterestsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_v1_feature_config_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllChatInterestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllChatInterestsResponse) ProtoMessage() {}

func (x *GetAllChatInterestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_v1_feature_config_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllChatInterestsResponse.ProtoReflect.Descriptor instead.
func (*GetAllChatInterestsResponse) Descriptor() ([]byte, []int) {
	return file_config_v1_feature_config_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllChatInterestsResponse) GetChatInterests() []*ChatInterestListItem {
	if x != nil {
		return x.ChatInterests
	}
	return nil
}

var File_config_v1_feature_config_service_proto protoreflect.FileDescriptor

var file_config_v1_feature_config_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xbc, 0x05, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x04, 0xca, 0x9d, 0x25, 0x00, 0x52, 0x0c, 0x6d,
	0x61, 0x78, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x66, 0x0a, 0x31, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f,
	0x72, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x2b, 0x75, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x4c, 0x65,
	0x67, 0x61, 0x63, 0x79, 0x52, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x53, 0x0a, 0x27, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x6f, 0x74, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x22, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x42, 0x6f, 0x74,
	0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x97, 0x01, 0x0a, 0x39, 0x6d, 0x69, 0x6e,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65,
	0x6e, 0x5f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x6f, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08, 0x01, 0x52,
	0x31, 0x6d, 0x69, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x74, 0x77,
	0x65, 0x65, 0x6e, 0x50, 0x75, 0x6c, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x42, 0x6f,
	0x74, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10,
	0x6d, 0x61, 0x78, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73,
	0x12, 0x70, 0x0a, 0x1a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x0b, 0xca, 0x9d,
	0x25, 0x07, 0x08, 0x01, 0x78, 0x01, 0x80, 0x01, 0x64, 0x52, 0x17, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x6e, 0x0a, 0x1a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x09,
	0xca, 0x9d, 0x25, 0x05, 0x78, 0x01, 0x80, 0x01, 0x64, 0x52, 0x17, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x55, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x62, 0x69, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x56,
	0x65, 0x72, 0x62, 0x69, 0x61, 0x67, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x74,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2d, 0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x12,
	0x32, 0x0a, 0x15, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x69, 0x63, 0x65,
	0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x49, 0x63, 0x65, 0x42, 0x72, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f,
	0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x73, 0x74, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0b, 0xca, 0x9d, 0x25, 0x07, 0x08, 0x01,
	0x78, 0x01, 0x80, 0x01, 0x64, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x79, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x0b, 0xca, 0x9d, 0x25, 0x07, 0x08, 0x01, 0x78,
	0x01, 0x80, 0x01, 0x64, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x73, 0x74, 0x73, 0x32, 0xf1, 0x02, 0x0a, 0x0d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x6e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x2a, 0x2e, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x49, 0x64, 0x73, 0x12, 0x2e,
	0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x65, 0x73, 0x74, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x65, 0x73, 0x74, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x74, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73, 0x12, 0x2c, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x62, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x6b,
	0x69, 0x6b, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x72, 0x70, 0x63, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x78, 0x69, 0x70, 0x68, 0x69,
	0x61, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_config_v1_feature_config_service_proto_rawDescOnce sync.Once
	file_config_v1_feature_config_service_proto_rawDescData = file_config_v1_feature_config_service_proto_rawDesc
)

func file_config_v1_feature_config_service_proto_rawDescGZIP() []byte {
	file_config_v1_feature_config_service_proto_rawDescOnce.Do(func() {
		file_config_v1_feature_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_v1_feature_config_service_proto_rawDescData)
	})
	return file_config_v1_feature_config_service_proto_rawDescData
}

var file_config_v1_feature_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_config_v1_feature_config_service_proto_goTypes = []interface{}{
	(*GetFeatureConfigsRequest)(nil),      // 0: mobile.config.v1.GetFeatureConfigsRequest
	(*GetFeatureConfigsResponse)(nil),     // 1: mobile.config.v1.GetFeatureConfigsResponse
	(*UserInterestListItem)(nil),          // 2: mobile.config.v1.UserInterestListItem
	(*ChatInterestListItem)(nil),          // 3: mobile.config.v1.ChatInterestListItem
	(*GetAllUserInterestIdsRequest)(nil),  // 4: mobile.config.v1.GetAllUserInterestIdsRequest
	(*GetAllUserInterestIdsResponse)(nil), // 5: mobile.config.v1.GetAllUserInterestIdsResponse
	(*GetAllChatInterestsRequest)(nil),    // 6: mobile.config.v1.GetAllChatInterestsRequest
	(*GetAllChatInterestsResponse)(nil),   // 7: mobile.config.v1.GetAllChatInterestsResponse
	(*duration.Duration)(nil),             // 8: google.protobuf.Duration
}
var file_config_v1_feature_config_service_proto_depIdxs = []int32{
	8, // 0: mobile.config.v1.GetFeatureConfigsResponse.min_duration_between_pull_entity_service_get_trusted_bots:type_name -> google.protobuf.Duration
	2, // 1: mobile.config.v1.GetFeatureConfigsResponse.current_user_interest_list:type_name -> mobile.config.v1.UserInterestListItem
	3, // 2: mobile.config.v1.GetFeatureConfigsResponse.current_chat_interest_list:type_name -> mobile.config.v1.ChatInterestListItem
	3, // 3: mobile.config.v1.GetAllChatInterestsResponse.chat_interests:type_name -> mobile.config.v1.ChatInterestListItem
	0, // 4: mobile.config.v1.FeatureConfig.GetFeatureConfigs:input_type -> mobile.config.v1.GetFeatureConfigsRequest
	4, // 5: mobile.config.v1.FeatureConfig.GetAllUserInterestIds:input_type -> mobile.config.v1.GetAllUserInterestIdsRequest
	6, // 6: mobile.config.v1.FeatureConfig.GetAllChatInterests:input_type -> mobile.config.v1.GetAllChatInterestsRequest
	1, // 7: mobile.config.v1.FeatureConfig.GetFeatureConfigs:output_type -> mobile.config.v1.GetFeatureConfigsResponse
	5, // 8: mobile.config.v1.FeatureConfig.GetAllUserInterestIds:output_type -> mobile.config.v1.GetAllUserInterestIdsResponse
	7, // 9: mobile.config.v1.FeatureConfig.GetAllChatInterests:output_type -> mobile.config.v1.GetAllChatInterestsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_config_v1_feature_config_service_proto_init() }
func file_config_v1_feature_config_service_proto_init() {
	if File_config_v1_feature_config_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_v1_feature_config_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeatureConfigsRequest); i {
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
		file_config_v1_feature_config_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeatureConfigsResponse); i {
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
		file_config_v1_feature_config_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInterestListItem); i {
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
		file_config_v1_feature_config_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatInterestListItem); i {
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
		file_config_v1_feature_config_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserInterestIdsRequest); i {
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
		file_config_v1_feature_config_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserInterestIdsResponse); i {
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
		file_config_v1_feature_config_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllChatInterestsRequest); i {
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
		file_config_v1_feature_config_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllChatInterestsResponse); i {
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
			RawDescriptor: file_config_v1_feature_config_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_v1_feature_config_service_proto_goTypes,
		DependencyIndexes: file_config_v1_feature_config_service_proto_depIdxs,
		MessageInfos:      file_config_v1_feature_config_service_proto_msgTypes,
	}.Build()
	File_config_v1_feature_config_service_proto = out.File
	file_config_v1_feature_config_service_proto_rawDesc = nil
	file_config_v1_feature_config_service_proto_goTypes = nil
	file_config_v1_feature_config_service_proto_depIdxs = nil
}
