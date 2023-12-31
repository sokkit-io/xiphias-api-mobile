// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: groups/v1/group_search_service.proto

package groups

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_go "github.com/sokkit-io/xiphias-model-common/generated/go"
	v1 "github.com/sokkit-io/xiphias-model-common/generated/go/groups/v1"
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

type GetGroupSuggestionsResponse_Result int32

const (
	GetGroupSuggestionsResponse_OK                  GetGroupSuggestionsResponse_Result = 0
	GetGroupSuggestionsResponse_RATE_LIMIT_EXCEEDED GetGroupSuggestionsResponse_Result = 1
)

// Enum value maps for GetGroupSuggestionsResponse_Result.
var (
	GetGroupSuggestionsResponse_Result_name = map[int32]string{
		0: "OK",
		1: "RATE_LIMIT_EXCEEDED",
	}
	GetGroupSuggestionsResponse_Result_value = map[string]int32{
		"OK":                  0,
		"RATE_LIMIT_EXCEEDED": 1,
	}
)

func (x GetGroupSuggestionsResponse_Result) Enum() *GetGroupSuggestionsResponse_Result {
	p := new(GetGroupSuggestionsResponse_Result)
	*p = x
	return p
}

func (x GetGroupSuggestionsResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetGroupSuggestionsResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_groups_v1_group_search_service_proto_enumTypes[0].Descriptor()
}

func (GetGroupSuggestionsResponse_Result) Type() protoreflect.EnumType {
	return &file_groups_v1_group_search_service_proto_enumTypes[0]
}

func (x GetGroupSuggestionsResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetGroupSuggestionsResponse_Result.Descriptor instead.
func (GetGroupSuggestionsResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_groups_v1_group_search_service_proto_rawDescGZIP(), []int{2, 0}
}

type FindGroupsResponse_Result int32

const (
	FindGroupsResponse_OK                  FindGroupsResponse_Result = 0
	FindGroupsResponse_RATE_LIMIT_EXCEEDED FindGroupsResponse_Result = 1
)

// Enum value maps for FindGroupsResponse_Result.
var (
	FindGroupsResponse_Result_name = map[int32]string{
		0: "OK",
		1: "RATE_LIMIT_EXCEEDED",
	}
	FindGroupsResponse_Result_value = map[string]int32{
		"OK":                  0,
		"RATE_LIMIT_EXCEEDED": 1,
	}
)

func (x FindGroupsResponse_Result) Enum() *FindGroupsResponse_Result {
	p := new(FindGroupsResponse_Result)
	*p = x
	return p
}

func (x FindGroupsResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FindGroupsResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_groups_v1_group_search_service_proto_enumTypes[1].Descriptor()
}

func (FindGroupsResponse_Result) Type() protoreflect.EnumType {
	return &file_groups_v1_group_search_service_proto_enumTypes[1]
}

func (x FindGroupsResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FindGroupsResponse_Result.Descriptor instead.
func (FindGroupsResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_groups_v1_group_search_service_proto_rawDescGZIP(), []int{4, 0}
}

type PublicGroupJoinToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token []byte `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *PublicGroupJoinToken) Reset() {
	*x = PublicGroupJoinToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_v1_group_search_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicGroupJoinToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicGroupJoinToken) ProtoMessage() {}

func (x *PublicGroupJoinToken) ProtoReflect() protoreflect.Message {
	mi := &file_groups_v1_group_search_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicGroupJoinToken.ProtoReflect.Descriptor instead.
func (*PublicGroupJoinToken) Descriptor() ([]byte, []int) {
	return file_groups_v1_group_search_service_proto_rawDescGZIP(), []int{0}
}

func (x *PublicGroupJoinToken) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

type GetGroupSuggestionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGroupSuggestionsRequest) Reset() {
	*x = GetGroupSuggestionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_v1_group_search_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupSuggestionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupSuggestionsRequest) ProtoMessage() {}

func (x *GetGroupSuggestionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groups_v1_group_search_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupSuggestionsRequest.ProtoReflect.Descriptor instead.
func (*GetGroupSuggestionsRequest) Descriptor() ([]byte, []int) {
	return file_groups_v1_group_search_service_proto_rawDescGZIP(), []int{1}
}

type GetGroupSuggestionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result GetGroupSuggestionsResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mobile.groups.v1.GetGroupSuggestionsResponse_Result" json:"result,omitempty"`
	// Suggested Groups. Expect upward of 25 group suggestions
	Suggestion []*LimitedGroupDetails `protobuf:"bytes,2,rep,name=suggestion,proto3" json:"suggestion,omitempty"`
}

func (x *GetGroupSuggestionsResponse) Reset() {
	*x = GetGroupSuggestionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_v1_group_search_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupSuggestionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupSuggestionsResponse) ProtoMessage() {}

func (x *GetGroupSuggestionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groups_v1_group_search_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupSuggestionsResponse.ProtoReflect.Descriptor instead.
func (*GetGroupSuggestionsResponse) Descriptor() ([]byte, []int) {
	return file_groups_v1_group_search_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetGroupSuggestionsResponse) GetResult() GetGroupSuggestionsResponse_Result {
	if x != nil {
		return x.Result
	}
	return GetGroupSuggestionsResponse_OK
}

func (x *GetGroupSuggestionsResponse) GetSuggestion() []*LimitedGroupDetails {
	if x != nil {
		return x.Suggestion
	}
	return nil
}

// Request for group details, given a simple query string.
type FindGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A simple string search containing a potential portion of a group hashtag.
	// This string does NOT contain special characters to indicate wildcards.
	// Note that this string is specified WITHOUT a '#' character prefix
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *FindGroupsRequest) Reset() {
	*x = FindGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_v1_group_search_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindGroupsRequest) ProtoMessage() {}

func (x *FindGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groups_v1_group_search_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindGroupsRequest.ProtoReflect.Descriptor instead.
func (*FindGroupsRequest) Descriptor() ([]byte, []int) {
	return file_groups_v1_group_search_service_proto_rawDescGZIP(), []int{3}
}

func (x *FindGroupsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// Response for a FindGroups request. Provides either an exact match or availability
// of the search string, as well as a filtered series of other groups that
// partially match the search string.
type FindGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result FindGroupsResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=mobile.groups.v1.FindGroupsResponse_Result" json:"result,omitempty"`
	// Groups that match the search request.  If there is an exact match, it
	// will be the first match record returned.
	Match []*LimitedGroupDetails `protobuf:"bytes,2,rep,name=match,proto3" json:"match,omitempty"`
	// Optional flag to indicate if the exact search term is available for use
	// as the hashtag in a public group creation request.
	IsAvailableForCreation bool `protobuf:"varint,3,opt,name=is_available_for_creation,json=isAvailableForCreation,proto3" json:"is_available_for_creation,omitempty"`
}

func (x *FindGroupsResponse) Reset() {
	*x = FindGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_v1_group_search_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindGroupsResponse) ProtoMessage() {}

func (x *FindGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groups_v1_group_search_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindGroupsResponse.ProtoReflect.Descriptor instead.
func (*FindGroupsResponse) Descriptor() ([]byte, []int) {
	return file_groups_v1_group_search_service_proto_rawDescGZIP(), []int{4}
}

func (x *FindGroupsResponse) GetResult() FindGroupsResponse_Result {
	if x != nil {
		return x.Result
	}
	return FindGroupsResponse_OK
}

func (x *FindGroupsResponse) GetMatch() []*LimitedGroupDetails {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *FindGroupsResponse) GetIsAvailableForCreation() bool {
	if x != nil {
		return x.IsAvailableForCreation
	}
	return false
}

// Group details, containing only enough details to provide display details of that group
type LimitedGroupDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the group.
	// NOTE: This is being kept as optional, as the group join token may render
	//       it obsolete.
	Jid *_go.XiGroupJid `protobuf:"bytes,1,opt,name=jid,proto3" json:"jid,omitempty"`
	// Display data of the group, including hashtag, display name, and display
	// url components.
	DisplayData *v1.GroupDisplayData `protobuf:"bytes,2,opt,name=display_data,json=displayData,proto3" json:"display_data,omitempty"`
	// Current count of members in the group
	MemberCount uint32 `protobuf:"varint,3,opt,name=member_count,json=memberCount,proto3" json:"member_count,omitempty"`
	// The time of the last group activity. If no last activity time was found the field will be empty
	// There is no guarantee regarding accuracy of this time - it means that at some point the group search service
	// was told that this group had activity at this time, but there may have been activity since then
	LastActivityTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=last_activity_time,json=lastActivityTime,proto3" json:"last_activity_time,omitempty"`
	// Maximum number of members group can contain at the present time.
	// Expect that this can change over time for a given group as we work towards larger groups
	// therefore, caching these values for a period of time on the order of minutes should be
	// ok, but it is not recommended to cache them for hours, and definitely not days
	// This value wouldn't normally be 0, but if it was, assume that it means there's a problem with
	// the public groups for which the size is 0 (which would probably be all of them) and that
	// they should be not shown to users.
	MaxGroupSize uint32 `protobuf:"varint,5,opt,name=max_group_size,json=maxGroupSize,proto3" json:"max_group_size,omitempty"`
	// The active members in this group. This field is only relevant if this LimitedGroupDetails
	// is being returned as part of a get trending groups request. If it's not, then its
	// value will always be 0 and it should be ignored.
	ActiveMembers uint32 `protobuf:"varint,6,opt,name=active_members,json=activeMembers,proto3" json:"active_members,omitempty"`
	// Token which must be provided in order to join this group directly.
	// The byte contents of this token must be base64-encoded and provided
	// in the kik-server group join stanza in order to be successful in joining
	// the group.
	// * Provides additional tracking capability in analytics to connect searches
	//   to group joins
	// * Provides additional security possibilities by limiting the options spammers
	//   have in harvesting users from groups (token is only useful to the user
	//   that received it, easier to apply rate limits, etc.)
	//
	GroupJoinToken *PublicGroupJoinToken `protobuf:"bytes,100,opt,name=group_join_token,json=groupJoinToken,proto3" json:"group_join_token,omitempty"`
}

func (x *LimitedGroupDetails) Reset() {
	*x = LimitedGroupDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_v1_group_search_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitedGroupDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitedGroupDetails) ProtoMessage() {}

func (x *LimitedGroupDetails) ProtoReflect() protoreflect.Message {
	mi := &file_groups_v1_group_search_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitedGroupDetails.ProtoReflect.Descriptor instead.
func (*LimitedGroupDetails) Descriptor() ([]byte, []int) {
	return file_groups_v1_group_search_service_proto_rawDescGZIP(), []int{5}
}

func (x *LimitedGroupDetails) GetJid() *_go.XiGroupJid {
	if x != nil {
		return x.Jid
	}
	return nil
}

func (x *LimitedGroupDetails) GetDisplayData() *v1.GroupDisplayData {
	if x != nil {
		return x.DisplayData
	}
	return nil
}

func (x *LimitedGroupDetails) GetMemberCount() uint32 {
	if x != nil {
		return x.MemberCount
	}
	return 0
}

func (x *LimitedGroupDetails) GetLastActivityTime() *timestamp.Timestamp {
	if x != nil {
		return x.LastActivityTime
	}
	return nil
}

func (x *LimitedGroupDetails) GetMaxGroupSize() uint32 {
	if x != nil {
		return x.MaxGroupSize
	}
	return 0
}

func (x *LimitedGroupDetails) GetActiveMembers() uint32 {
	if x != nil {
		return x.ActiveMembers
	}
	return 0
}

func (x *LimitedGroupDetails) GetGroupJoinToken() *PublicGroupJoinToken {
	if x != nil {
		return x.GroupJoinToken
	}
	return nil
}

var File_groups_v1_group_search_service_proto protoreflect.FileDescriptor

var file_groups_v1_group_search_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x14, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x21, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0b,
	0xca, 0x9d, 0x25, 0x07, 0x08, 0x01, 0x28, 0x01, 0x30, 0x80, 0x28, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xe9, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x34, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x51,
	0x0a, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x0a, 0xca, 0x9d, 0x25, 0x06, 0x08,
	0x00, 0x80, 0x01, 0x80, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x29, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49,
	0x54, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x01, 0x22, 0x48, 0x0a, 0x11,
	0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x33, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1d, 0xca, 0x9d, 0x25, 0x19, 0x08, 0x01, 0x12, 0x15, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2e, 0x5f, 0x5d, 0x7b, 0x31, 0x2c, 0x33, 0x32, 0x7d, 0x24, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x87, 0x02, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x46, 0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x09, 0xca, 0x9d, 0x25, 0x05, 0x08, 0x00,
	0x80, 0x01, 0x19, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x39, 0x0a, 0x19, 0x69, 0x73,
	0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x69,
	0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x29, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x41, 0x54, 0x45, 0x5f,
	0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x01,
	0x22, 0xa6, 0x03, 0x0a, 0x13, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2c, 0x0a, 0x03, 0x6a, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x58,
	0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x69, 0x64, 0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08,
	0x00, 0x52, 0x03, 0x6a, 0x69, 0x64, 0x12, 0x4d, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08, 0x00, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x50, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08, 0x00, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61,
	0x78, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x50, 0x0a, 0x10, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xde, 0x01, 0x0a, 0x0b, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x59, 0x0a, 0x0a, 0x46, 0x69, 0x6e,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x23, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x74, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x2e, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x57, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x5a, 0x45, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d,
	0x69, 0x6f, 0x2f, 0x78, 0x69, 0x70, 0x68, 0x69, 0x61, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x67, 0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_groups_v1_group_search_service_proto_rawDescOnce sync.Once
	file_groups_v1_group_search_service_proto_rawDescData = file_groups_v1_group_search_service_proto_rawDesc
)

func file_groups_v1_group_search_service_proto_rawDescGZIP() []byte {
	file_groups_v1_group_search_service_proto_rawDescOnce.Do(func() {
		file_groups_v1_group_search_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_groups_v1_group_search_service_proto_rawDescData)
	})
	return file_groups_v1_group_search_service_proto_rawDescData
}

var file_groups_v1_group_search_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_groups_v1_group_search_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_groups_v1_group_search_service_proto_goTypes = []interface{}{
	(GetGroupSuggestionsResponse_Result)(0), // 0: mobile.groups.v1.GetGroupSuggestionsResponse.Result
	(FindGroupsResponse_Result)(0),          // 1: mobile.groups.v1.FindGroupsResponse.Result
	(*PublicGroupJoinToken)(nil),            // 2: mobile.groups.v1.PublicGroupJoinToken
	(*GetGroupSuggestionsRequest)(nil),      // 3: mobile.groups.v1.GetGroupSuggestionsRequest
	(*GetGroupSuggestionsResponse)(nil),     // 4: mobile.groups.v1.GetGroupSuggestionsResponse
	(*FindGroupsRequest)(nil),               // 5: mobile.groups.v1.FindGroupsRequest
	(*FindGroupsResponse)(nil),              // 6: mobile.groups.v1.FindGroupsResponse
	(*LimitedGroupDetails)(nil),             // 7: mobile.groups.v1.LimitedGroupDetails
	(*_go.XiGroupJid)(nil),                  // 8: common.XiGroupJid
	(*v1.GroupDisplayData)(nil),             // 9: common.groups.v1.GroupDisplayData
	(*timestamp.Timestamp)(nil),             // 10: google.protobuf.Timestamp
}
var file_groups_v1_group_search_service_proto_depIdxs = []int32{
	0,  // 0: mobile.groups.v1.GetGroupSuggestionsResponse.result:type_name -> mobile.groups.v1.GetGroupSuggestionsResponse.Result
	7,  // 1: mobile.groups.v1.GetGroupSuggestionsResponse.suggestion:type_name -> mobile.groups.v1.LimitedGroupDetails
	1,  // 2: mobile.groups.v1.FindGroupsResponse.result:type_name -> mobile.groups.v1.FindGroupsResponse.Result
	7,  // 3: mobile.groups.v1.FindGroupsResponse.match:type_name -> mobile.groups.v1.LimitedGroupDetails
	8,  // 4: mobile.groups.v1.LimitedGroupDetails.jid:type_name -> common.XiGroupJid
	9,  // 5: mobile.groups.v1.LimitedGroupDetails.display_data:type_name -> common.groups.v1.GroupDisplayData
	10, // 6: mobile.groups.v1.LimitedGroupDetails.last_activity_time:type_name -> google.protobuf.Timestamp
	2,  // 7: mobile.groups.v1.LimitedGroupDetails.group_join_token:type_name -> mobile.groups.v1.PublicGroupJoinToken
	5,  // 8: mobile.groups.v1.GroupSearch.FindGroups:input_type -> mobile.groups.v1.FindGroupsRequest
	3,  // 9: mobile.groups.v1.GroupSearch.GetGroupSuggestions:input_type -> mobile.groups.v1.GetGroupSuggestionsRequest
	6,  // 10: mobile.groups.v1.GroupSearch.FindGroups:output_type -> mobile.groups.v1.FindGroupsResponse
	4,  // 11: mobile.groups.v1.GroupSearch.GetGroupSuggestions:output_type -> mobile.groups.v1.GetGroupSuggestionsResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_groups_v1_group_search_service_proto_init() }
func file_groups_v1_group_search_service_proto_init() {
	if File_groups_v1_group_search_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_groups_v1_group_search_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicGroupJoinToken); i {
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
		file_groups_v1_group_search_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupSuggestionsRequest); i {
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
		file_groups_v1_group_search_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupSuggestionsResponse); i {
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
		file_groups_v1_group_search_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindGroupsRequest); i {
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
		file_groups_v1_group_search_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindGroupsResponse); i {
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
		file_groups_v1_group_search_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitedGroupDetails); i {
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
			RawDescriptor: file_groups_v1_group_search_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_groups_v1_group_search_service_proto_goTypes,
		DependencyIndexes: file_groups_v1_group_search_service_proto_depIdxs,
		EnumInfos:         file_groups_v1_group_search_service_proto_enumTypes,
		MessageInfos:      file_groups_v1_group_search_service_proto_msgTypes,
	}.Build()
	File_groups_v1_group_search_service_proto = out.File
	file_groups_v1_group_search_service_proto_rawDesc = nil
	file_groups_v1_group_search_service_proto_goTypes = nil
	file_groups_v1_group_search_service_proto_depIdxs = nil
}
