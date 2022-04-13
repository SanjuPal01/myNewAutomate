// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: interservice/cds/service/cds.proto

package service

import (
	context "context"
	request "github.com/chef/automate/api/external/cds/request"
	response "github.com/chef/automate/api/external/cds/response"
	common "github.com/chef/automate/api/external/common"
	version "github.com/chef/automate/api/external/common/version"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

var File_interservice_cds_service_cds_proto protoreflect.FileDescriptor

var file_interservice_cds_service_cds_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x64, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x5f, 0x63, 0x64, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x1f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x64, 0x73, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x64, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc9, 0x05, 0x0a, 0x12, 0x41, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x43, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x71, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x6d, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2b, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x1a, 0x2c, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x7b, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x32, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x64, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x73, 0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x32, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x24, 0x2e, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x71, 0x0a, 0x10, 0x49, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2d, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x1a, 0x2e, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73, 0x2e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x6c, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x2a, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x64, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x2b, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x73,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x64, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_interservice_cds_service_cds_proto_goTypes = []interface{}{
	(*version.VersionInfoRequest)(nil),  // 0: chef.automate.api.common.version.VersionInfoRequest
	(*request.ContentItems)(nil),        // 1: chef.automate.api.cds.request.ContentItems
	(*request.InstallContentItem)(nil),  // 2: chef.automate.api.cds.request.InstallContentItem
	(*request.DownloadContentItem)(nil), // 3: chef.automate.api.cds.request.DownloadContentItem
	(*request.ContentEnabled)(nil),      // 4: chef.automate.api.cds.request.ContentEnabled
	(*request.Credentials)(nil),         // 5: chef.automate.api.cds.request.Credentials
	(*version.VersionInfo)(nil),         // 6: chef.automate.api.common.version.VersionInfo
	(*response.ContentItems)(nil),       // 7: chef.automate.api.cds.response.ContentItems
	(*response.InstallContentItem)(nil), // 8: chef.automate.api.cds.response.InstallContentItem
	(*common.ExportData)(nil),           // 9: chef.automate.api.common.ExportData
	(*response.ContentEnabled)(nil),     // 10: chef.automate.api.cds.response.ContentEnabled
	(*response.Credentials)(nil),        // 11: chef.automate.api.cds.response.Credentials
}
var file_interservice_cds_service_cds_proto_depIdxs = []int32{
	0,  // 0: chef.automate.domain.automate_cds.service.AutomateCdsService.GetVersion:input_type -> chef.automate.api.common.version.VersionInfoRequest
	1,  // 1: chef.automate.domain.automate_cds.service.AutomateCdsService.ListContentItems:input_type -> chef.automate.api.cds.request.ContentItems
	2,  // 2: chef.automate.domain.automate_cds.service.AutomateCdsService.InstallContentItem:input_type -> chef.automate.api.cds.request.InstallContentItem
	3,  // 3: chef.automate.domain.automate_cds.service.AutomateCdsService.DownloadContentItem:input_type -> chef.automate.api.cds.request.DownloadContentItem
	4,  // 4: chef.automate.domain.automate_cds.service.AutomateCdsService.IsContentEnabled:input_type -> chef.automate.api.cds.request.ContentEnabled
	5,  // 5: chef.automate.domain.automate_cds.service.AutomateCdsService.SubmitCredentials:input_type -> chef.automate.api.cds.request.Credentials
	6,  // 6: chef.automate.domain.automate_cds.service.AutomateCdsService.GetVersion:output_type -> chef.automate.api.common.version.VersionInfo
	7,  // 7: chef.automate.domain.automate_cds.service.AutomateCdsService.ListContentItems:output_type -> chef.automate.api.cds.response.ContentItems
	8,  // 8: chef.automate.domain.automate_cds.service.AutomateCdsService.InstallContentItem:output_type -> chef.automate.api.cds.response.InstallContentItem
	9,  // 9: chef.automate.domain.automate_cds.service.AutomateCdsService.DownloadContentItem:output_type -> chef.automate.api.common.ExportData
	10, // 10: chef.automate.domain.automate_cds.service.AutomateCdsService.IsContentEnabled:output_type -> chef.automate.api.cds.response.ContentEnabled
	11, // 11: chef.automate.domain.automate_cds.service.AutomateCdsService.SubmitCredentials:output_type -> chef.automate.api.cds.response.Credentials
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_interservice_cds_service_cds_proto_init() }
func file_interservice_cds_service_cds_proto_init() {
	if File_interservice_cds_service_cds_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interservice_cds_service_cds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interservice_cds_service_cds_proto_goTypes,
		DependencyIndexes: file_interservice_cds_service_cds_proto_depIdxs,
	}.Build()
	File_interservice_cds_service_cds_proto = out.File
	file_interservice_cds_service_cds_proto_rawDesc = nil
	file_interservice_cds_service_cds_proto_goTypes = nil
	file_interservice_cds_service_cds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AutomateCdsServiceClient is the client API for AutomateCdsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AutomateCdsServiceClient interface {
	GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error)
	ListContentItems(ctx context.Context, in *request.ContentItems, opts ...grpc.CallOption) (*response.ContentItems, error)
	InstallContentItem(ctx context.Context, in *request.InstallContentItem, opts ...grpc.CallOption) (*response.InstallContentItem, error)
	DownloadContentItem(ctx context.Context, in *request.DownloadContentItem, opts ...grpc.CallOption) (AutomateCdsService_DownloadContentItemClient, error)
	IsContentEnabled(ctx context.Context, in *request.ContentEnabled, opts ...grpc.CallOption) (*response.ContentEnabled, error)
	SubmitCredentials(ctx context.Context, in *request.Credentials, opts ...grpc.CallOption) (*response.Credentials, error)
}

type automateCdsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAutomateCdsServiceClient(cc grpc.ClientConnInterface) AutomateCdsServiceClient {
	return &automateCdsServiceClient{cc}
}

func (c *automateCdsServiceClient) GetVersion(ctx context.Context, in *version.VersionInfoRequest, opts ...grpc.CallOption) (*version.VersionInfo, error) {
	out := new(version.VersionInfo)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.automate_cds.service.AutomateCdsService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automateCdsServiceClient) ListContentItems(ctx context.Context, in *request.ContentItems, opts ...grpc.CallOption) (*response.ContentItems, error) {
	out := new(response.ContentItems)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.automate_cds.service.AutomateCdsService/ListContentItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automateCdsServiceClient) InstallContentItem(ctx context.Context, in *request.InstallContentItem, opts ...grpc.CallOption) (*response.InstallContentItem, error) {
	out := new(response.InstallContentItem)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.automate_cds.service.AutomateCdsService/InstallContentItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automateCdsServiceClient) DownloadContentItem(ctx context.Context, in *request.DownloadContentItem, opts ...grpc.CallOption) (AutomateCdsService_DownloadContentItemClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AutomateCdsService_serviceDesc.Streams[0], "/chef.automate.domain.automate_cds.service.AutomateCdsService/DownloadContentItem", opts...)
	if err != nil {
		return nil, err
	}
	x := &automateCdsServiceDownloadContentItemClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AutomateCdsService_DownloadContentItemClient interface {
	Recv() (*common.ExportData, error)
	grpc.ClientStream
}

type automateCdsServiceDownloadContentItemClient struct {
	grpc.ClientStream
}

func (x *automateCdsServiceDownloadContentItemClient) Recv() (*common.ExportData, error) {
	m := new(common.ExportData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *automateCdsServiceClient) IsContentEnabled(ctx context.Context, in *request.ContentEnabled, opts ...grpc.CallOption) (*response.ContentEnabled, error) {
	out := new(response.ContentEnabled)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.automate_cds.service.AutomateCdsService/IsContentEnabled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automateCdsServiceClient) SubmitCredentials(ctx context.Context, in *request.Credentials, opts ...grpc.CallOption) (*response.Credentials, error) {
	out := new(response.Credentials)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.automate_cds.service.AutomateCdsService/SubmitCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutomateCdsServiceServer is the server API for AutomateCdsService service.
type AutomateCdsServiceServer interface {
	GetVersion(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error)
	ListContentItems(context.Context, *request.ContentItems) (*response.ContentItems, error)
	InstallContentItem(context.Context, *request.InstallContentItem) (*response.InstallContentItem, error)
	DownloadContentItem(*request.DownloadContentItem, AutomateCdsService_DownloadContentItemServer) error
	IsContentEnabled(context.Context, *request.ContentEnabled) (*response.ContentEnabled, error)
	SubmitCredentials(context.Context, *request.Credentials) (*response.Credentials, error)
}

// UnimplementedAutomateCdsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAutomateCdsServiceServer struct {
}

func (*UnimplementedAutomateCdsServiceServer) GetVersion(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (*UnimplementedAutomateCdsServiceServer) ListContentItems(context.Context, *request.ContentItems) (*response.ContentItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContentItems not implemented")
}
func (*UnimplementedAutomateCdsServiceServer) InstallContentItem(context.Context, *request.InstallContentItem) (*response.InstallContentItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallContentItem not implemented")
}
func (*UnimplementedAutomateCdsServiceServer) DownloadContentItem(*request.DownloadContentItem, AutomateCdsService_DownloadContentItemServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadContentItem not implemented")
}
func (*UnimplementedAutomateCdsServiceServer) IsContentEnabled(context.Context, *request.ContentEnabled) (*response.ContentEnabled, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsContentEnabled not implemented")
}
func (*UnimplementedAutomateCdsServiceServer) SubmitCredentials(context.Context, *request.Credentials) (*response.Credentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitCredentials not implemented")
}

func RegisterAutomateCdsServiceServer(s *grpc.Server, srv AutomateCdsServiceServer) {
	s.RegisterService(&_AutomateCdsService_serviceDesc, srv)
}

func _AutomateCdsService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(version.VersionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomateCdsServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.automate_cds.service.AutomateCdsService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomateCdsServiceServer).GetVersion(ctx, req.(*version.VersionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomateCdsService_ListContentItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ContentItems)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomateCdsServiceServer).ListContentItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.automate_cds.service.AutomateCdsService/ListContentItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomateCdsServiceServer).ListContentItems(ctx, req.(*request.ContentItems))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomateCdsService_InstallContentItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.InstallContentItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomateCdsServiceServer).InstallContentItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.automate_cds.service.AutomateCdsService/InstallContentItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomateCdsServiceServer).InstallContentItem(ctx, req.(*request.InstallContentItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomateCdsService_DownloadContentItem_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(request.DownloadContentItem)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AutomateCdsServiceServer).DownloadContentItem(m, &automateCdsServiceDownloadContentItemServer{stream})
}

type AutomateCdsService_DownloadContentItemServer interface {
	Send(*common.ExportData) error
	grpc.ServerStream
}

type automateCdsServiceDownloadContentItemServer struct {
	grpc.ServerStream
}

func (x *automateCdsServiceDownloadContentItemServer) Send(m *common.ExportData) error {
	return x.ServerStream.SendMsg(m)
}

func _AutomateCdsService_IsContentEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ContentEnabled)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomateCdsServiceServer).IsContentEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.automate_cds.service.AutomateCdsService/IsContentEnabled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomateCdsServiceServer).IsContentEnabled(ctx, req.(*request.ContentEnabled))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomateCdsService_SubmitCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomateCdsServiceServer).SubmitCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.automate_cds.service.AutomateCdsService/SubmitCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomateCdsServiceServer).SubmitCredentials(ctx, req.(*request.Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

var _AutomateCdsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.automate_cds.service.AutomateCdsService",
	HandlerType: (*AutomateCdsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _AutomateCdsService_GetVersion_Handler,
		},
		{
			MethodName: "ListContentItems",
			Handler:    _AutomateCdsService_ListContentItems_Handler,
		},
		{
			MethodName: "InstallContentItem",
			Handler:    _AutomateCdsService_InstallContentItem_Handler,
		},
		{
			MethodName: "IsContentEnabled",
			Handler:    _AutomateCdsService_IsContentEnabled_Handler,
		},
		{
			MethodName: "SubmitCredentials",
			Handler:    _AutomateCdsService_SubmitCredentials_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadContentItem",
			Handler:       _AutomateCdsService_DownloadContentItem_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "interservice/cds/service/cds.proto",
}