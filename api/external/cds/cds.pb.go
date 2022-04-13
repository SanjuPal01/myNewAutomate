// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: external/cds/cds.proto

package cds

import (
	context "context"
	_ "github.com/chef/automate/api/external/annotations/iam"
	request "github.com/chef/automate/api/external/cds/request"
	response "github.com/chef/automate/api/external/cds/response"
	common "github.com/chef/automate/api/external/common"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
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

var File_external_cds_cds_proto protoreflect.FileDescriptor

var file_external_cds_cds_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x64, 0x73, 0x2f, 0x63,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x1a,
	0x1f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x64, 0x73, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x64, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa8, 0x07, 0x0a, 0x03, 0x43, 0x64, 0x73,
	0x12, 0xb9, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2b, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x1a, 0x2c, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62,
	0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x8a, 0xb5, 0x18, 0x0f, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x8a, 0xb5, 0x18, 0x14, 0x12, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x3a, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xcc, 0x01, 0x0a,
	0x11, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x12, 0x2a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x2b,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x5e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x3a, 0x01, 0x2a, 0x8a, 0xb5, 0x18, 0x15, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x3a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x8a, 0xb5,
	0x18, 0x19, 0x12, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x3a, 0x61, 0x64, 0x64, 0x12, 0xce, 0x01, 0x0a, 0x10,
	0x49, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x2d, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x1a,
	0x2e, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22,
	0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x65,
	0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x8a, 0xb5, 0x18, 0x15, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x8a, 0xb5, 0x18, 0x1d, 0x12,
	0x1b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x3a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0xcf, 0x01, 0x0a,
	0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x32, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x52, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1e, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x3a, 0x01, 0x2a,
	0x8a, 0xb5, 0x18, 0x0f, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x8a, 0xb5, 0x18, 0x17, 0x12, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x3a, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x3a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x73,
	0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x32, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x24, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x00, 0x30, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x64, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_external_cds_cds_proto_goTypes = []interface{}{
	(*request.ContentItems)(nil),        // 0: chef.automate.api.cds.request.ContentItems
	(*request.Credentials)(nil),         // 1: chef.automate.api.cds.request.Credentials
	(*request.ContentEnabled)(nil),      // 2: chef.automate.api.cds.request.ContentEnabled
	(*request.InstallContentItem)(nil),  // 3: chef.automate.api.cds.request.InstallContentItem
	(*request.DownloadContentItem)(nil), // 4: chef.automate.api.cds.request.DownloadContentItem
	(*response.ContentItems)(nil),       // 5: chef.automate.api.cds.response.ContentItems
	(*response.Credentials)(nil),        // 6: chef.automate.api.cds.response.Credentials
	(*response.ContentEnabled)(nil),     // 7: chef.automate.api.cds.response.ContentEnabled
	(*response.InstallContentItem)(nil), // 8: chef.automate.api.cds.response.InstallContentItem
	(*common.ExportData)(nil),           // 9: chef.automate.api.common.ExportData
}
var file_external_cds_cds_proto_depIdxs = []int32{
	0, // 0: chef.automate.api.cds.Cds.ListContentItems:input_type -> chef.automate.api.cds.request.ContentItems
	1, // 1: chef.automate.api.cds.Cds.SubmitCredentials:input_type -> chef.automate.api.cds.request.Credentials
	2, // 2: chef.automate.api.cds.Cds.IsContentEnabled:input_type -> chef.automate.api.cds.request.ContentEnabled
	3, // 3: chef.automate.api.cds.Cds.InstallContentItem:input_type -> chef.automate.api.cds.request.InstallContentItem
	4, // 4: chef.automate.api.cds.Cds.DownloadContentItem:input_type -> chef.automate.api.cds.request.DownloadContentItem
	5, // 5: chef.automate.api.cds.Cds.ListContentItems:output_type -> chef.automate.api.cds.response.ContentItems
	6, // 6: chef.automate.api.cds.Cds.SubmitCredentials:output_type -> chef.automate.api.cds.response.Credentials
	7, // 7: chef.automate.api.cds.Cds.IsContentEnabled:output_type -> chef.automate.api.cds.response.ContentEnabled
	8, // 8: chef.automate.api.cds.Cds.InstallContentItem:output_type -> chef.automate.api.cds.response.InstallContentItem
	9, // 9: chef.automate.api.cds.Cds.DownloadContentItem:output_type -> chef.automate.api.common.ExportData
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_cds_cds_proto_init() }
func file_external_cds_cds_proto_init() {
	if File_external_cds_cds_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_cds_cds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_cds_cds_proto_goTypes,
		DependencyIndexes: file_external_cds_cds_proto_depIdxs,
	}.Build()
	File_external_cds_cds_proto = out.File
	file_external_cds_cds_proto_rawDesc = nil
	file_external_cds_cds_proto_goTypes = nil
	file_external_cds_cds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CdsClient is the client API for Cds service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CdsClient interface {
	//
	//ListContentItems
	//
	//Returns a list of metadata for each CDS content. Provides a description and current
	//state of each content item.
	//
	//Authorization Action:
	//```
	//content:items:list
	//```
	ListContentItems(ctx context.Context, in *request.ContentItems, opts ...grpc.CallOption) (*response.ContentItems, error)
	//
	//SubmitCredentials
	//
	//Submit a Chef Cloud Credentials to enable content
	//
	//Authorization Action:
	//```
	//content:credentials:add
	//```
	SubmitCredentials(ctx context.Context, in *request.Credentials, opts ...grpc.CallOption) (*response.Credentials, error)
	//
	//IsContentEnabled
	//
	//Check if the content is enable for this Automate instance.
	//
	//Authorization Action:
	//```
	//content:credentials:enabled
	//```
	IsContentEnabled(ctx context.Context, in *request.ContentEnabled, opts ...grpc.CallOption) (*response.ContentEnabled, error)
	//
	//InstallContentItem
	//
	//Installs a content item from its ID
	//
	//Authorization Action:
	//```
	//content:items:install
	//```
	InstallContentItem(ctx context.Context, in *request.InstallContentItem, opts ...grpc.CallOption) (*response.InstallContentItem, error)
	//
	//DownloadContentItem
	//
	//Download a content item from its ID
	//
	//grpc gateway is not able to handle streaming; https://github.com/grpc-ecosystem/grpc-gateway/issues/435
	//so we do not auto-generate the route for download; we instead custom handle with mux
	//
	//Authorization Action:
	//```
	//content:items:download
	//```
	DownloadContentItem(ctx context.Context, in *request.DownloadContentItem, opts ...grpc.CallOption) (Cds_DownloadContentItemClient, error)
}

type cdsClient struct {
	cc grpc.ClientConnInterface
}

func NewCdsClient(cc grpc.ClientConnInterface) CdsClient {
	return &cdsClient{cc}
}

func (c *cdsClient) ListContentItems(ctx context.Context, in *request.ContentItems, opts ...grpc.CallOption) (*response.ContentItems, error) {
	out := new(response.ContentItems)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cds.Cds/ListContentItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cdsClient) SubmitCredentials(ctx context.Context, in *request.Credentials, opts ...grpc.CallOption) (*response.Credentials, error) {
	out := new(response.Credentials)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cds.Cds/SubmitCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cdsClient) IsContentEnabled(ctx context.Context, in *request.ContentEnabled, opts ...grpc.CallOption) (*response.ContentEnabled, error) {
	out := new(response.ContentEnabled)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cds.Cds/IsContentEnabled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cdsClient) InstallContentItem(ctx context.Context, in *request.InstallContentItem, opts ...grpc.CallOption) (*response.InstallContentItem, error) {
	out := new(response.InstallContentItem)
	err := c.cc.Invoke(ctx, "/chef.automate.api.cds.Cds/InstallContentItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cdsClient) DownloadContentItem(ctx context.Context, in *request.DownloadContentItem, opts ...grpc.CallOption) (Cds_DownloadContentItemClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Cds_serviceDesc.Streams[0], "/chef.automate.api.cds.Cds/DownloadContentItem", opts...)
	if err != nil {
		return nil, err
	}
	x := &cdsDownloadContentItemClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cds_DownloadContentItemClient interface {
	Recv() (*common.ExportData, error)
	grpc.ClientStream
}

type cdsDownloadContentItemClient struct {
	grpc.ClientStream
}

func (x *cdsDownloadContentItemClient) Recv() (*common.ExportData, error) {
	m := new(common.ExportData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CdsServer is the server API for Cds service.
type CdsServer interface {
	//
	//ListContentItems
	//
	//Returns a list of metadata for each CDS content. Provides a description and current
	//state of each content item.
	//
	//Authorization Action:
	//```
	//content:items:list
	//```
	ListContentItems(context.Context, *request.ContentItems) (*response.ContentItems, error)
	//
	//SubmitCredentials
	//
	//Submit a Chef Cloud Credentials to enable content
	//
	//Authorization Action:
	//```
	//content:credentials:add
	//```
	SubmitCredentials(context.Context, *request.Credentials) (*response.Credentials, error)
	//
	//IsContentEnabled
	//
	//Check if the content is enable for this Automate instance.
	//
	//Authorization Action:
	//```
	//content:credentials:enabled
	//```
	IsContentEnabled(context.Context, *request.ContentEnabled) (*response.ContentEnabled, error)
	//
	//InstallContentItem
	//
	//Installs a content item from its ID
	//
	//Authorization Action:
	//```
	//content:items:install
	//```
	InstallContentItem(context.Context, *request.InstallContentItem) (*response.InstallContentItem, error)
	//
	//DownloadContentItem
	//
	//Download a content item from its ID
	//
	//grpc gateway is not able to handle streaming; https://github.com/grpc-ecosystem/grpc-gateway/issues/435
	//so we do not auto-generate the route for download; we instead custom handle with mux
	//
	//Authorization Action:
	//```
	//content:items:download
	//```
	DownloadContentItem(*request.DownloadContentItem, Cds_DownloadContentItemServer) error
}

// UnimplementedCdsServer can be embedded to have forward compatible implementations.
type UnimplementedCdsServer struct {
}

func (*UnimplementedCdsServer) ListContentItems(context.Context, *request.ContentItems) (*response.ContentItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContentItems not implemented")
}
func (*UnimplementedCdsServer) SubmitCredentials(context.Context, *request.Credentials) (*response.Credentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitCredentials not implemented")
}
func (*UnimplementedCdsServer) IsContentEnabled(context.Context, *request.ContentEnabled) (*response.ContentEnabled, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsContentEnabled not implemented")
}
func (*UnimplementedCdsServer) InstallContentItem(context.Context, *request.InstallContentItem) (*response.InstallContentItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallContentItem not implemented")
}
func (*UnimplementedCdsServer) DownloadContentItem(*request.DownloadContentItem, Cds_DownloadContentItemServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadContentItem not implemented")
}

func RegisterCdsServer(s *grpc.Server, srv CdsServer) {
	s.RegisterService(&_Cds_serviceDesc, srv)
}

func _Cds_ListContentItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ContentItems)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CdsServer).ListContentItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cds.Cds/ListContentItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CdsServer).ListContentItems(ctx, req.(*request.ContentItems))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cds_SubmitCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CdsServer).SubmitCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cds.Cds/SubmitCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CdsServer).SubmitCredentials(ctx, req.(*request.Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cds_IsContentEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ContentEnabled)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CdsServer).IsContentEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cds.Cds/IsContentEnabled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CdsServer).IsContentEnabled(ctx, req.(*request.ContentEnabled))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cds_InstallContentItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.InstallContentItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CdsServer).InstallContentItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.cds.Cds/InstallContentItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CdsServer).InstallContentItem(ctx, req.(*request.InstallContentItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cds_DownloadContentItem_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(request.DownloadContentItem)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CdsServer).DownloadContentItem(m, &cdsDownloadContentItemServer{stream})
}

type Cds_DownloadContentItemServer interface {
	Send(*common.ExportData) error
	grpc.ServerStream
}

type cdsDownloadContentItemServer struct {
	grpc.ServerStream
}

func (x *cdsDownloadContentItemServer) Send(m *common.ExportData) error {
	return x.ServerStream.SendMsg(m)
}

var _Cds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.cds.Cds",
	HandlerType: (*CdsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListContentItems",
			Handler:    _Cds_ListContentItems_Handler,
		},
		{
			MethodName: "SubmitCredentials",
			Handler:    _Cds_SubmitCredentials_Handler,
		},
		{
			MethodName: "IsContentEnabled",
			Handler:    _Cds_IsContentEnabled_Handler,
		},
		{
			MethodName: "InstallContentItem",
			Handler:    _Cds_InstallContentItem_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadContentItem",
			Handler:       _Cds_DownloadContentItem_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "external/cds/cds.proto",
}