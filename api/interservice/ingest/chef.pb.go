// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: interservice/ingest/chef.proto

package ingest

import (
	context "context"
	request "github.com/chef/automate/api/external/ingest/request"
	response "github.com/chef/automate/api/external/ingest/response"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Version message
//
// The ingest-service version constructed with:
// * Service name
// * Built time
// * Semantic version
// * Git SHA
type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty" toml:"version,omitempty" mapstructure:"version,omitempty"`
	Built   string `protobuf:"bytes,1,opt,name=built,proto3" json:"built,omitempty" toml:"built,omitempty" mapstructure:"built,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	Sha     string `protobuf:"bytes,4,opt,name=sha,proto3" json:"sha,omitempty" toml:"sha,omitempty" mapstructure:"sha,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_ingest_chef_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_ingest_chef_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_interservice_ingest_chef_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Version) GetBuilt() string {
	if x != nil {
		return x.Built
	}
	return ""
}

func (x *Version) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Version) GetSha() string {
	if x != nil {
		return x.Sha
	}
	return ""
}

type VersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VersionRequest) Reset() {
	*x = VersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_ingest_chef_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionRequest) ProtoMessage() {}

func (x *VersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_ingest_chef_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionRequest.ProtoReflect.Descriptor instead.
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return file_interservice_ingest_chef_proto_rawDescGZIP(), []int{1}
}

var File_interservice_ingest_chef_proto protoreflect.FileDescriptor

var file_interservice_ingest_chef_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x6c,
	0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x68, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x68, 0x61, 0x22, 0x10, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xc8, 0x07, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x66, 0x49, 0x6e,
	0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x96, 0x01,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x66, 0x52, 0x75, 0x6e,
	0x12, 0x25, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x52, 0x75, 0x6e, 0x1a, 0x39, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x66, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x30, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f,
	0x72, 0x75, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0xa2, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x43, 0x68, 0x65, 0x66, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3c, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x43, 0x68, 0x65, 0x66, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x68, 0x65,
	0x66, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0xa6, 0x01, 0x0a, 0x13,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x2a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x1a,
	0x3a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x76, 0x65, 0x6e,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73,
	0x73, 0x3a, 0x01, 0x2a, 0x12, 0x9f, 0x01, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x73, 0x12, 0x3b, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x44, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa6, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x3c, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x68, 0x65,
	0x66, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x7f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x30, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e,
	0x67, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_ingest_chef_proto_rawDescOnce sync.Once
	file_interservice_ingest_chef_proto_rawDescData = file_interservice_ingest_chef_proto_rawDesc
)

func file_interservice_ingest_chef_proto_rawDescGZIP() []byte {
	file_interservice_ingest_chef_proto_rawDescOnce.Do(func() {
		file_interservice_ingest_chef_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_ingest_chef_proto_rawDescData)
	})
	return file_interservice_ingest_chef_proto_rawDescData
}

var file_interservice_ingest_chef_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_interservice_ingest_chef_proto_goTypes = []interface{}{
	(*Version)(nil),                                    // 0: chef.automate.domain.ingest.Version
	(*VersionRequest)(nil),                             // 1: chef.automate.domain.ingest.VersionRequest
	(*request.Run)(nil),                                // 2: chef.automate.api.ingest.request.Run
	(*request.Action)(nil),                             // 3: chef.automate.api.ingest.request.Action
	(*request.Liveness)(nil),                           // 4: chef.automate.api.ingest.request.Liveness
	(*request.MultipleNodeDeleteRequest)(nil),          // 5: chef.automate.api.ingest.request.MultipleNodeDeleteRequest
	(*request.Delete)(nil),                             // 6: chef.automate.api.ingest.request.Delete
	(*response.ProcessChefRunResponse)(nil),            // 7: chef.automate.api.ingest.response.ProcessChefRunResponse
	(*response.ProcessChefActionResponse)(nil),         // 8: chef.automate.api.ingest.response.ProcessChefActionResponse
	(*response.ProcessLivenessResponse)(nil),           // 9: chef.automate.api.ingest.response.ProcessLivenessResponse
	(*response.ProcessMultipleNodeDeleteResponse)(nil), // 10: chef.automate.api.ingest.response.ProcessMultipleNodeDeleteResponse
	(*response.ProcessNodeDeleteResponse)(nil),         // 11: chef.automate.api.ingest.response.ProcessNodeDeleteResponse
}
var file_interservice_ingest_chef_proto_depIdxs = []int32{
	2,  // 0: chef.automate.domain.ingest.ChefIngesterService.ProcessChefRun:input_type -> chef.automate.api.ingest.request.Run
	3,  // 1: chef.automate.domain.ingest.ChefIngesterService.ProcessChefAction:input_type -> chef.automate.api.ingest.request.Action
	4,  // 2: chef.automate.domain.ingest.ChefIngesterService.ProcessLivenessPing:input_type -> chef.automate.api.ingest.request.Liveness
	5,  // 3: chef.automate.domain.ingest.ChefIngesterService.ProcessMultipleNodeDeletes:input_type -> chef.automate.api.ingest.request.MultipleNodeDeleteRequest
	6,  // 4: chef.automate.domain.ingest.ChefIngesterService.ProcessNodeDelete:input_type -> chef.automate.api.ingest.request.Delete
	1,  // 5: chef.automate.domain.ingest.ChefIngesterService.GetVersion:input_type -> chef.automate.domain.ingest.VersionRequest
	7,  // 6: chef.automate.domain.ingest.ChefIngesterService.ProcessChefRun:output_type -> chef.automate.api.ingest.response.ProcessChefRunResponse
	8,  // 7: chef.automate.domain.ingest.ChefIngesterService.ProcessChefAction:output_type -> chef.automate.api.ingest.response.ProcessChefActionResponse
	9,  // 8: chef.automate.domain.ingest.ChefIngesterService.ProcessLivenessPing:output_type -> chef.automate.api.ingest.response.ProcessLivenessResponse
	10, // 9: chef.automate.domain.ingest.ChefIngesterService.ProcessMultipleNodeDeletes:output_type -> chef.automate.api.ingest.response.ProcessMultipleNodeDeleteResponse
	11, // 10: chef.automate.domain.ingest.ChefIngesterService.ProcessNodeDelete:output_type -> chef.automate.api.ingest.response.ProcessNodeDeleteResponse
	0,  // 11: chef.automate.domain.ingest.ChefIngesterService.GetVersion:output_type -> chef.automate.domain.ingest.Version
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_interservice_ingest_chef_proto_init() }
func file_interservice_ingest_chef_proto_init() {
	if File_interservice_ingest_chef_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_ingest_chef_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_interservice_ingest_chef_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionRequest); i {
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
			RawDescriptor: file_interservice_ingest_chef_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interservice_ingest_chef_proto_goTypes,
		DependencyIndexes: file_interservice_ingest_chef_proto_depIdxs,
		MessageInfos:      file_interservice_ingest_chef_proto_msgTypes,
	}.Build()
	File_interservice_ingest_chef_proto = out.File
	file_interservice_ingest_chef_proto_rawDesc = nil
	file_interservice_ingest_chef_proto_goTypes = nil
	file_interservice_ingest_chef_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChefIngesterServiceClient is the client API for ChefIngesterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChefIngesterServiceClient interface {
	ProcessChefRun(ctx context.Context, in *request.Run, opts ...grpc.CallOption) (*response.ProcessChefRunResponse, error)
	ProcessChefAction(ctx context.Context, in *request.Action, opts ...grpc.CallOption) (*response.ProcessChefActionResponse, error)
	ProcessLivenessPing(ctx context.Context, in *request.Liveness, opts ...grpc.CallOption) (*response.ProcessLivenessResponse, error)
	ProcessMultipleNodeDeletes(ctx context.Context, in *request.MultipleNodeDeleteRequest, opts ...grpc.CallOption) (*response.ProcessMultipleNodeDeleteResponse, error)
	ProcessNodeDelete(ctx context.Context, in *request.Delete, opts ...grpc.CallOption) (*response.ProcessNodeDeleteResponse, error)
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*Version, error)
}

type chefIngesterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChefIngesterServiceClient(cc grpc.ClientConnInterface) ChefIngesterServiceClient {
	return &chefIngesterServiceClient{cc}
}

func (c *chefIngesterServiceClient) ProcessChefRun(ctx context.Context, in *request.Run, opts ...grpc.CallOption) (*response.ProcessChefRunResponse, error) {
	out := new(response.ProcessChefRunResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngesterService/ProcessChefRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterServiceClient) ProcessChefAction(ctx context.Context, in *request.Action, opts ...grpc.CallOption) (*response.ProcessChefActionResponse, error) {
	out := new(response.ProcessChefActionResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngesterService/ProcessChefAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterServiceClient) ProcessLivenessPing(ctx context.Context, in *request.Liveness, opts ...grpc.CallOption) (*response.ProcessLivenessResponse, error) {
	out := new(response.ProcessLivenessResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngesterService/ProcessLivenessPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterServiceClient) ProcessMultipleNodeDeletes(ctx context.Context, in *request.MultipleNodeDeleteRequest, opts ...grpc.CallOption) (*response.ProcessMultipleNodeDeleteResponse, error) {
	out := new(response.ProcessMultipleNodeDeleteResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngesterService/ProcessMultipleNodeDeletes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterServiceClient) ProcessNodeDelete(ctx context.Context, in *request.Delete, opts ...grpc.CallOption) (*response.ProcessNodeDeleteResponse, error) {
	out := new(response.ProcessNodeDeleteResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngesterService/ProcessNodeDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterServiceClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngesterService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChefIngesterServiceServer is the server API for ChefIngesterService service.
type ChefIngesterServiceServer interface {
	ProcessChefRun(context.Context, *request.Run) (*response.ProcessChefRunResponse, error)
	ProcessChefAction(context.Context, *request.Action) (*response.ProcessChefActionResponse, error)
	ProcessLivenessPing(context.Context, *request.Liveness) (*response.ProcessLivenessResponse, error)
	ProcessMultipleNodeDeletes(context.Context, *request.MultipleNodeDeleteRequest) (*response.ProcessMultipleNodeDeleteResponse, error)
	ProcessNodeDelete(context.Context, *request.Delete) (*response.ProcessNodeDeleteResponse, error)
	GetVersion(context.Context, *VersionRequest) (*Version, error)
}

// UnimplementedChefIngesterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChefIngesterServiceServer struct {
}

func (*UnimplementedChefIngesterServiceServer) ProcessChefRun(context.Context, *request.Run) (*response.ProcessChefRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessChefRun not implemented")
}
func (*UnimplementedChefIngesterServiceServer) ProcessChefAction(context.Context, *request.Action) (*response.ProcessChefActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessChefAction not implemented")
}
func (*UnimplementedChefIngesterServiceServer) ProcessLivenessPing(context.Context, *request.Liveness) (*response.ProcessLivenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessLivenessPing not implemented")
}
func (*UnimplementedChefIngesterServiceServer) ProcessMultipleNodeDeletes(context.Context, *request.MultipleNodeDeleteRequest) (*response.ProcessMultipleNodeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessMultipleNodeDeletes not implemented")
}
func (*UnimplementedChefIngesterServiceServer) ProcessNodeDelete(context.Context, *request.Delete) (*response.ProcessNodeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessNodeDelete not implemented")
}
func (*UnimplementedChefIngesterServiceServer) GetVersion(context.Context, *VersionRequest) (*Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}

func RegisterChefIngesterServiceServer(s *grpc.Server, srv ChefIngesterServiceServer) {
	s.RegisterService(&_ChefIngesterService_serviceDesc, srv)
}

func _ChefIngesterService_ProcessChefRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Run)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServiceServer).ProcessChefRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngesterService/ProcessChefRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServiceServer).ProcessChefRun(ctx, req.(*request.Run))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngesterService_ProcessChefAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Action)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServiceServer).ProcessChefAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngesterService/ProcessChefAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServiceServer).ProcessChefAction(ctx, req.(*request.Action))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngesterService_ProcessLivenessPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Liveness)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServiceServer).ProcessLivenessPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngesterService/ProcessLivenessPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServiceServer).ProcessLivenessPing(ctx, req.(*request.Liveness))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngesterService_ProcessMultipleNodeDeletes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.MultipleNodeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServiceServer).ProcessMultipleNodeDeletes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngesterService/ProcessMultipleNodeDeletes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServiceServer).ProcessMultipleNodeDeletes(ctx, req.(*request.MultipleNodeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngesterService_ProcessNodeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Delete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServiceServer).ProcessNodeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngesterService/ProcessNodeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServiceServer).ProcessNodeDelete(ctx, req.(*request.Delete))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngesterService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngesterService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServiceServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChefIngesterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.ingest.ChefIngesterService",
	HandlerType: (*ChefIngesterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessChefRun",
			Handler:    _ChefIngesterService_ProcessChefRun_Handler,
		},
		{
			MethodName: "ProcessChefAction",
			Handler:    _ChefIngesterService_ProcessChefAction_Handler,
		},
		{
			MethodName: "ProcessLivenessPing",
			Handler:    _ChefIngesterService_ProcessLivenessPing_Handler,
		},
		{
			MethodName: "ProcessMultipleNodeDeletes",
			Handler:    _ChefIngesterService_ProcessMultipleNodeDeletes_Handler,
		},
		{
			MethodName: "ProcessNodeDelete",
			Handler:    _ChefIngesterService_ProcessNodeDelete_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ChefIngesterService_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interservice/ingest/chef.proto",
}
