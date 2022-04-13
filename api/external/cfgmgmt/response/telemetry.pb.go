// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: external/cfgmgmt/response/telemetry.proto

package response

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UpdateTelemetryReportedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTelemetryReportedResponse) Reset() {
	*x = UpdateTelemetryReportedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_response_telemetry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTelemetryReportedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTelemetryReportedResponse) ProtoMessage() {}

func (x *UpdateTelemetryReportedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_response_telemetry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTelemetryReportedResponse.ProtoReflect.Descriptor instead.
func (*UpdateTelemetryReportedResponse) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_response_telemetry_proto_rawDescGZIP(), []int{0}
}

type GetNodesUsageCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// number of days since telematics was last posted
	DaysSinceLastPost int64 `protobuf:"varint,1,opt,name=days_since_last_post,json=daysSinceLastPost,proto3" json:"days_since_last_post,omitempty"`
	// unique nodes count in a duration
	NodeCnt int64 `protobuf:"varint,2,opt,name=node_cnt,json=nodeCnt,proto3" json:"node_cnt,omitempty"`
}

func (x *GetNodesUsageCountResponse) Reset() {
	*x = GetNodesUsageCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_response_telemetry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodesUsageCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesUsageCountResponse) ProtoMessage() {}

func (x *GetNodesUsageCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_response_telemetry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesUsageCountResponse.ProtoReflect.Descriptor instead.
func (*GetNodesUsageCountResponse) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_response_telemetry_proto_rawDescGZIP(), []int{1}
}

func (x *GetNodesUsageCountResponse) GetDaysSinceLastPost() int64 {
	if x != nil {
		return x.DaysSinceLastPost
	}
	return 0
}

func (x *GetNodesUsageCountResponse) GetNodeCnt() int64 {
	if x != nil {
		return x.NodeCnt
	}
	return 0
}

var File_external_cfgmgmt_response_telemetry_proto protoreflect.FileDescriptor

var file_external_cfgmgmt_response_telemetry_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67,
	0x6d, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x21, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x68, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x14, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11,
	0x64, 0x61, 0x79, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6e, 0x74, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_cfgmgmt_response_telemetry_proto_rawDescOnce sync.Once
	file_external_cfgmgmt_response_telemetry_proto_rawDescData = file_external_cfgmgmt_response_telemetry_proto_rawDesc
)

func file_external_cfgmgmt_response_telemetry_proto_rawDescGZIP() []byte {
	file_external_cfgmgmt_response_telemetry_proto_rawDescOnce.Do(func() {
		file_external_cfgmgmt_response_telemetry_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_cfgmgmt_response_telemetry_proto_rawDescData)
	})
	return file_external_cfgmgmt_response_telemetry_proto_rawDescData
}

var file_external_cfgmgmt_response_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_external_cfgmgmt_response_telemetry_proto_goTypes = []interface{}{
	(*UpdateTelemetryReportedResponse)(nil), // 0: chef.automate.api.cfgmgmt.response.UpdateTelemetryReportedResponse
	(*GetNodesUsageCountResponse)(nil),      // 1: chef.automate.api.cfgmgmt.response.GetNodesUsageCountResponse
}
var file_external_cfgmgmt_response_telemetry_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_cfgmgmt_response_telemetry_proto_init() }
func file_external_cfgmgmt_response_telemetry_proto_init() {
	if File_external_cfgmgmt_response_telemetry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_cfgmgmt_response_telemetry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTelemetryReportedResponse); i {
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
		file_external_cfgmgmt_response_telemetry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodesUsageCountResponse); i {
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
			RawDescriptor: file_external_cfgmgmt_response_telemetry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_cfgmgmt_response_telemetry_proto_goTypes,
		DependencyIndexes: file_external_cfgmgmt_response_telemetry_proto_depIdxs,
		MessageInfos:      file_external_cfgmgmt_response_telemetry_proto_msgTypes,
	}.Build()
	File_external_cfgmgmt_response_telemetry_proto = out.File
	file_external_cfgmgmt_response_telemetry_proto_rawDesc = nil
	file_external_cfgmgmt_response_telemetry_proto_goTypes = nil
	file_external_cfgmgmt_response_telemetry_proto_depIdxs = nil
}