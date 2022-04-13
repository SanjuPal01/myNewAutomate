// -*- mode: protobuf; indent-tabs-mode: t; c-basic-offset: 8; tab-width: 8 -*-

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: config/applications/config_request.proto

package applications

import (
	shared "github.com/chef/automate/api/config/shared"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V1 *ConfigRequest_V1 `protobuf:"bytes,3,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_applications_config_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_applications_config_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_applications_config_request_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if x != nil {
		return x.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sys *ConfigRequest_V1_System `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
}

func (x *ConfigRequest_V1) Reset() {
	*x = ConfigRequest_V1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_applications_config_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1) ProtoMessage() {}

func (x *ConfigRequest_V1) ProtoReflect() protoreflect.Message {
	mi := &file_config_applications_config_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return file_config_applications_config_request_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if x != nil {
		return x.Sys
	}
	return nil
}

type ConfigRequest_V1_System struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mlsa    *shared.Mlsa                     `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Service *ConfigRequest_V1_System_Service `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Tls     *shared.TLSCredentials           `protobuf:"bytes,3,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Log     *ConfigRequest_V1_System_Log     `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	Storage *ConfigRequest_V1_System_Storage `protobuf:"bytes,5,opt,name=storage,proto3" json:"storage,omitempty" toml:"storage,omitempty" mapstructure:"storage,omitempty"`
}

func (x *ConfigRequest_V1_System) Reset() {
	*x = ConfigRequest_V1_System{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_applications_config_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System) ProtoMessage() {}

func (x *ConfigRequest_V1_System) ProtoReflect() protoreflect.Message {
	mi := &file_config_applications_config_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return file_config_applications_config_request_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if x != nil {
		return x.Mlsa
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if x != nil {
		return x.Tls
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetStorage() *ConfigRequest_V1_System_Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	Host        *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"` // The listen host is no longer setable(localhost only)
	Port        *wrapperspb.Int32Value  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	MetricsPort *wrapperspb.Int32Value  `protobuf:"bytes,3,opt,name=metrics_port,json=metricsPort,proto3" json:"metrics_port,omitempty" toml:"metrics_port,omitempty" mapstructure:"metrics_port,omitempty"`
}

func (x *ConfigRequest_V1_System_Service) Reset() {
	*x = ConfigRequest_V1_System_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_applications_config_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Service) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Service) ProtoReflect() protoreflect.Message {
	mi := &file_config_applications_config_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Service.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return file_config_applications_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

// Deprecated: Do not use.
func (x *ConfigRequest_V1_System_Service) GetHost() *wrapperspb.StringValue {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetPort() *wrapperspb.Int32Value {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetMetricsPort() *wrapperspb.Int32Value {
	if x != nil {
		return x.MetricsPort
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Format *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=format,proto3" json:"format,omitempty" toml:"format,omitempty" mapstructure:"format,omitempty"`
	Level  *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
}

func (x *ConfigRequest_V1_System_Log) Reset() {
	*x = ConfigRequest_V1_System_Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_applications_config_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Log) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Log) ProtoReflect() protoreflect.Message {
	mi := &file_config_applications_config_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Log.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return file_config_applications_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 1}
}

func (x *ConfigRequest_V1_System_Log) GetFormat() *wrapperspb.StringValue {
	if x != nil {
		return x.Format
	}
	return nil
}

func (x *ConfigRequest_V1_System_Log) GetLevel() *wrapperspb.StringValue {
	if x != nil {
		return x.Level
	}
	return nil
}

type ConfigRequest_V1_System_Storage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty" toml:"database,omitempty" mapstructure:"database,omitempty"`
	User     *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty" toml:"user,omitempty" mapstructure:"user,omitempty"`
}

func (x *ConfigRequest_V1_System_Storage) Reset() {
	*x = ConfigRequest_V1_System_Storage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_applications_config_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Storage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Storage) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Storage) ProtoReflect() protoreflect.Message {
	mi := &file_config_applications_config_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Storage.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Storage) Descriptor() ([]byte, []int) {
	return file_config_applications_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 2}
}

func (x *ConfigRequest_V1_System_Storage) GetDatabase() *wrapperspb.StringValue {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *ConfigRequest_V1_System_Storage) GetUser() *wrapperspb.StringValue {
	if x != nil {
		return x.User
	}
	return nil
}

var File_config_applications_config_request_proto protoreflect.FileDescriptor

var file_config_applications_config_request_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1a, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x74, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x61, 0x32, 0x2d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x32, 0x63, 0x6f, 0x6e,
	0x66, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x08, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x02, 0x76, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x52, 0x02, 0x76, 0x31, 0x1a, 0xb4, 0x07, 0x0a, 0x02, 0x56,
	0x31, 0x12, 0x4c, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x03, 0x73, 0x79, 0x73, 0x1a,
	0xdf, 0x06, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x34, 0x0a, 0x04, 0x6d, 0x6c,
	0x73, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x6c, 0x73, 0x61, 0x52, 0x04, 0x6d, 0x6c, 0x73, 0x61,
	0x12, 0x5c, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x42, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c,
	0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12, 0x50, 0x0a, 0x03,
	0x6c, 0x6f, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x5c,
	0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x42, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0xea, 0x01, 0x0a,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x47,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x16, 0xc2, 0xf3, 0x18, 0x12, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0x95, 0x4f, 0x1a, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x5a, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x1a, 0xc2, 0xf3, 0x18, 0x16,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x10, 0xa5, 0x9d, 0x01,
	0x1a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x50,
	0x6f, 0x72, 0x74, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x1a, 0x6f, 0x0a, 0x03, 0x4c, 0x6f, 0x67,
	0x12, 0x34, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0x75, 0x0a, 0x07, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x3a, 0x1a, 0xc2, 0xf3, 0x18, 0x16, 0x0a, 0x14, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x32, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_applications_config_request_proto_rawDescOnce sync.Once
	file_config_applications_config_request_proto_rawDescData = file_config_applications_config_request_proto_rawDesc
)

func file_config_applications_config_request_proto_rawDescGZIP() []byte {
	file_config_applications_config_request_proto_rawDescOnce.Do(func() {
		file_config_applications_config_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_applications_config_request_proto_rawDescData)
	})
	return file_config_applications_config_request_proto_rawDescData
}

var file_config_applications_config_request_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_config_applications_config_request_proto_goTypes = []interface{}{
	(*ConfigRequest)(nil),                   // 0: chef.automate.domain.applications.ConfigRequest
	(*ConfigRequest_V1)(nil),                // 1: chef.automate.domain.applications.ConfigRequest.V1
	(*ConfigRequest_V1_System)(nil),         // 2: chef.automate.domain.applications.ConfigRequest.V1.System
	(*ConfigRequest_V1_System_Service)(nil), // 3: chef.automate.domain.applications.ConfigRequest.V1.System.Service
	(*ConfigRequest_V1_System_Log)(nil),     // 4: chef.automate.domain.applications.ConfigRequest.V1.System.Log
	(*ConfigRequest_V1_System_Storage)(nil), // 5: chef.automate.domain.applications.ConfigRequest.V1.System.Storage
	(*shared.Mlsa)(nil),                     // 6: chef.automate.infra.config.Mlsa
	(*shared.TLSCredentials)(nil),           // 7: chef.automate.infra.config.TLSCredentials
	(*wrapperspb.StringValue)(nil),          // 8: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),           // 9: google.protobuf.Int32Value
}
var file_config_applications_config_request_proto_depIdxs = []int32{
	1,  // 0: chef.automate.domain.applications.ConfigRequest.v1:type_name -> chef.automate.domain.applications.ConfigRequest.V1
	2,  // 1: chef.automate.domain.applications.ConfigRequest.V1.sys:type_name -> chef.automate.domain.applications.ConfigRequest.V1.System
	6,  // 2: chef.automate.domain.applications.ConfigRequest.V1.System.mlsa:type_name -> chef.automate.infra.config.Mlsa
	3,  // 3: chef.automate.domain.applications.ConfigRequest.V1.System.service:type_name -> chef.automate.domain.applications.ConfigRequest.V1.System.Service
	7,  // 4: chef.automate.domain.applications.ConfigRequest.V1.System.tls:type_name -> chef.automate.infra.config.TLSCredentials
	4,  // 5: chef.automate.domain.applications.ConfigRequest.V1.System.log:type_name -> chef.automate.domain.applications.ConfigRequest.V1.System.Log
	5,  // 6: chef.automate.domain.applications.ConfigRequest.V1.System.storage:type_name -> chef.automate.domain.applications.ConfigRequest.V1.System.Storage
	8,  // 7: chef.automate.domain.applications.ConfigRequest.V1.System.Service.host:type_name -> google.protobuf.StringValue
	9,  // 8: chef.automate.domain.applications.ConfigRequest.V1.System.Service.port:type_name -> google.protobuf.Int32Value
	9,  // 9: chef.automate.domain.applications.ConfigRequest.V1.System.Service.metrics_port:type_name -> google.protobuf.Int32Value
	8,  // 10: chef.automate.domain.applications.ConfigRequest.V1.System.Log.format:type_name -> google.protobuf.StringValue
	8,  // 11: chef.automate.domain.applications.ConfigRequest.V1.System.Log.level:type_name -> google.protobuf.StringValue
	8,  // 12: chef.automate.domain.applications.ConfigRequest.V1.System.Storage.database:type_name -> google.protobuf.StringValue
	8,  // 13: chef.automate.domain.applications.ConfigRequest.V1.System.Storage.user:type_name -> google.protobuf.StringValue
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_config_applications_config_request_proto_init() }
func file_config_applications_config_request_proto_init() {
	if File_config_applications_config_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_applications_config_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_config_applications_config_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1); i {
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
		file_config_applications_config_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System); i {
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
		file_config_applications_config_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Service); i {
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
		file_config_applications_config_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Log); i {
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
		file_config_applications_config_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Storage); i {
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
			RawDescriptor: file_config_applications_config_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_applications_config_request_proto_goTypes,
		DependencyIndexes: file_config_applications_config_request_proto_depIdxs,
		MessageInfos:      file_config_applications_config_request_proto_msgTypes,
	}.Build()
	File_config_applications_config_request_proto = out.File
	file_config_applications_config_request_proto_rawDesc = nil
	file_config_applications_config_request_proto_goTypes = nil
	file_config_applications_config_request_proto_depIdxs = nil
}