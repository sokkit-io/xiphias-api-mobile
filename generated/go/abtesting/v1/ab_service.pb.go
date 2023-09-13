// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: abtesting/v1/ab_service.proto

package abtesting

import (
	_go "github.com/sokkit-io/xiphias-model-common/generated/go"
	v1 "github.com/sokkit-io/xiphias-model-common/generated/go/abtesting/v1"
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

type GetExperimentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The jid for the user we are fetching experiments for
	// NOTE: "This will not be part of the request long term - the server bridge will send the jid as a GRPC header to service." - Dan
	Jid *_go.XiBareUserJid `protobuf:"bytes,1,opt,name=jid,proto3" json:"jid,omitempty"`
	// Set of service selected (and only service selected) Experiment's the user already knows it is in.
	// We *ARE NOT* including pre-registration selected experiments at this time.
	ParticipatingExperiments []*v1.Experiment `protobuf:"bytes,2,rep,name=participating_experiments,json=participatingExperiments,proto3" json:"participating_experiments,omitempty"`
}

func (x *GetExperimentsRequest) Reset() {
	*x = GetExperimentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_abtesting_v1_ab_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExperimentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExperimentsRequest) ProtoMessage() {}

func (x *GetExperimentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_abtesting_v1_ab_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExperimentsRequest.ProtoReflect.Descriptor instead.
func (*GetExperimentsRequest) Descriptor() ([]byte, []int) {
	return file_abtesting_v1_ab_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetExperimentsRequest) GetJid() *_go.XiBareUserJid {
	if x != nil {
		return x.Jid
	}
	return nil
}

func (x *GetExperimentsRequest) GetParticipatingExperiments() []*v1.Experiment {
	if x != nil {
		return x.ParticipatingExperiments
	}
	return nil
}

type GetExperimentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Set of Experiment's associated with the requested user
	ParticipatingExperiments []*v1.Experiment `protobuf:"bytes,1,rep,name=participating_experiments,json=participatingExperiments,proto3" json:"participating_experiments,omitempty"`
	// Determines whether or not this response should direct clients to send the metrics update.
	// This should be TRUE in the majority of scenarios except for when there is a need to handle
	// client metrics load issues.
	// Specifically, this should be only used in overload scenarios to prevent heartbeats for users not in any
	// experiments. However it is technically possible to be used at any time.
	ShouldSendMetricsEvent bool `protobuf:"varint,2,opt,name=should_send_metrics_event,json=shouldSendMetricsEvent,proto3" json:"should_send_metrics_event,omitempty"`
}

func (x *GetExperimentsResponse) Reset() {
	*x = GetExperimentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_abtesting_v1_ab_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExperimentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExperimentsResponse) ProtoMessage() {}

func (x *GetExperimentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_abtesting_v1_ab_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExperimentsResponse.ProtoReflect.Descriptor instead.
func (*GetExperimentsResponse) Descriptor() ([]byte, []int) {
	return file_abtesting_v1_ab_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetExperimentsResponse) GetParticipatingExperiments() []*v1.Experiment {
	if x != nil {
		return x.ParticipatingExperiments
	}
	return nil
}

func (x *GetExperimentsResponse) GetShouldSendMetricsEvent() bool {
	if x != nil {
		return x.ShouldSendMetricsEvent
	}
	return false
}

var File_abtesting_v1_ab_service_proto protoreflect.FileDescriptor

var file_abtesting_v1_ab_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x62, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x03, 0x6a, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x58, 0x69, 0x42, 0x61, 0x72, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x4a, 0x69, 0x64, 0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08, 0x01, 0x52,
	0x03, 0x6a, 0x69, 0x64, 0x12, 0x67, 0x0a, 0x19, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x09, 0xca, 0x9d, 0x25, 0x05, 0x08, 0x00,
	0x80, 0x01, 0x64, 0x52, 0x18, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xb1, 0x01,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x19, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x18, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64,
	0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x73, 0x68, 0x6f, 0x75, 0x6c,
	0x64, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x32, 0x85, 0x01, 0x0a, 0x09, 0x41, 0x62, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x78, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2a,
	0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x64, 0x0a, 0x15, 0x63, 0x6f, 0x6d,
	0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72,
	0x70, 0x63, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x78, 0x69, 0x70, 0x68, 0x69, 0x61, 0x73,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x62, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_abtesting_v1_ab_service_proto_rawDescOnce sync.Once
	file_abtesting_v1_ab_service_proto_rawDescData = file_abtesting_v1_ab_service_proto_rawDesc
)

func file_abtesting_v1_ab_service_proto_rawDescGZIP() []byte {
	file_abtesting_v1_ab_service_proto_rawDescOnce.Do(func() {
		file_abtesting_v1_ab_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_abtesting_v1_ab_service_proto_rawDescData)
	})
	return file_abtesting_v1_ab_service_proto_rawDescData
}

var file_abtesting_v1_ab_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_abtesting_v1_ab_service_proto_goTypes = []interface{}{
	(*GetExperimentsRequest)(nil),  // 0: mobile.abtesting.v1.GetExperimentsRequest
	(*GetExperimentsResponse)(nil), // 1: mobile.abtesting.v1.GetExperimentsResponse
	(*_go.XiBareUserJid)(nil),      // 2: common.XiBareUserJid
	(*v1.Experiment)(nil),          // 3: common.abtesting.v1.Experiment
}
var file_abtesting_v1_ab_service_proto_depIdxs = []int32{
	2, // 0: mobile.abtesting.v1.GetExperimentsRequest.jid:type_name -> common.XiBareUserJid
	3, // 1: mobile.abtesting.v1.GetExperimentsRequest.participating_experiments:type_name -> common.abtesting.v1.Experiment
	3, // 2: mobile.abtesting.v1.GetExperimentsResponse.participating_experiments:type_name -> common.abtesting.v1.Experiment
	0, // 3: mobile.abtesting.v1.AbTesting.GetParticipatingExperiments:input_type -> mobile.abtesting.v1.GetExperimentsRequest
	1, // 4: mobile.abtesting.v1.AbTesting.GetParticipatingExperiments:output_type -> mobile.abtesting.v1.GetExperimentsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_abtesting_v1_ab_service_proto_init() }
func file_abtesting_v1_ab_service_proto_init() {
	if File_abtesting_v1_ab_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_abtesting_v1_ab_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExperimentsRequest); i {
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
		file_abtesting_v1_ab_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExperimentsResponse); i {
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
			RawDescriptor: file_abtesting_v1_ab_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_abtesting_v1_ab_service_proto_goTypes,
		DependencyIndexes: file_abtesting_v1_ab_service_proto_depIdxs,
		MessageInfos:      file_abtesting_v1_ab_service_proto_msgTypes,
	}.Build()
	File_abtesting_v1_ab_service_proto = out.File
	file_abtesting_v1_ab_service_proto_rawDesc = nil
	file_abtesting_v1_ab_service_proto_goTypes = nil
	file_abtesting_v1_ab_service_proto_depIdxs = nil
}