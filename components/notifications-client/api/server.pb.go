// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: server.proto

package api

import (
	context "context"
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

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x13, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0b, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd4, 0x04,
	0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x37, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x14, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a,
	0x17, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x1a, 0x1e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x21, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x1a, 0x21, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x1e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12,
	0x23, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x55, 0x52, 0x4c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x52, 0x4c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_server_proto_goTypes = []interface{}{
	(*Event)(nil),                 // 0: notifications.Event
	(*Rule)(nil),                  // 1: notifications.Rule
	(*RuleIdentifier)(nil),        // 2: notifications.RuleIdentifier
	(*Empty)(nil),                 // 3: notifications.Empty
	(*URLValidationRequest)(nil),  // 4: notifications.URLValidationRequest
	(*VersionRequest)(nil),        // 5: notifications.VersionRequest
	(*Response)(nil),              // 6: notifications.Response
	(*RuleAddResponse)(nil),       // 7: notifications.RuleAddResponse
	(*RuleDeleteResponse)(nil),    // 8: notifications.RuleDeleteResponse
	(*RuleUpdateResponse)(nil),    // 9: notifications.RuleUpdateResponse
	(*RuleGetResponse)(nil),       // 10: notifications.RuleGetResponse
	(*RuleListResponse)(nil),      // 11: notifications.RuleListResponse
	(*URLValidationResponse)(nil), // 12: notifications.URLValidationResponse
	(*VersionResponse)(nil),       // 13: notifications.VersionResponse
}
var file_server_proto_depIdxs = []int32{
	0,  // 0: notifications.Notifications.Notify:input_type -> notifications.Event
	1,  // 1: notifications.Notifications.AddRule:input_type -> notifications.Rule
	2,  // 2: notifications.Notifications.DeleteRule:input_type -> notifications.RuleIdentifier
	1,  // 3: notifications.Notifications.UpdateRule:input_type -> notifications.Rule
	2,  // 4: notifications.Notifications.GetRule:input_type -> notifications.RuleIdentifier
	3,  // 5: notifications.Notifications.ListRules:input_type -> notifications.Empty
	4,  // 6: notifications.Notifications.ValidateWebhook:input_type -> notifications.URLValidationRequest
	5,  // 7: notifications.Notifications.Version:input_type -> notifications.VersionRequest
	6,  // 8: notifications.Notifications.Notify:output_type -> notifications.Response
	7,  // 9: notifications.Notifications.AddRule:output_type -> notifications.RuleAddResponse
	8,  // 10: notifications.Notifications.DeleteRule:output_type -> notifications.RuleDeleteResponse
	9,  // 11: notifications.Notifications.UpdateRule:output_type -> notifications.RuleUpdateResponse
	10, // 12: notifications.Notifications.GetRule:output_type -> notifications.RuleGetResponse
	11, // 13: notifications.Notifications.ListRules:output_type -> notifications.RuleListResponse
	12, // 14: notifications.Notifications.ValidateWebhook:output_type -> notifications.URLValidationResponse
	13, // 15: notifications.Notifications.Version:output_type -> notifications.VersionResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	file_notifications_proto_init()
	file_rules_proto_init()
	file_health_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotificationsClient is the client API for Notifications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationsClient interface {
	// Publish a notification
	Notify(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error)
	// Manage notification alerting rules
	AddRule(ctx context.Context, in *Rule, opts ...grpc.CallOption) (*RuleAddResponse, error)
	DeleteRule(ctx context.Context, in *RuleIdentifier, opts ...grpc.CallOption) (*RuleDeleteResponse, error)
	UpdateRule(ctx context.Context, in *Rule, opts ...grpc.CallOption) (*RuleUpdateResponse, error)
	GetRule(ctx context.Context, in *RuleIdentifier, opts ...grpc.CallOption) (*RuleGetResponse, error)
	ListRules(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RuleListResponse, error)
	ValidateWebhook(ctx context.Context, in *URLValidationRequest, opts ...grpc.CallOption) (*URLValidationResponse, error)
	// Health checks and metadata
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type notificationsClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationsClient(cc grpc.ClientConnInterface) NotificationsClient {
	return &notificationsClient{cc}
}

func (c *notificationsClient) Notify(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) AddRule(ctx context.Context, in *Rule, opts ...grpc.CallOption) (*RuleAddResponse, error) {
	out := new(RuleAddResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/AddRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) DeleteRule(ctx context.Context, in *RuleIdentifier, opts ...grpc.CallOption) (*RuleDeleteResponse, error) {
	out := new(RuleDeleteResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/DeleteRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) UpdateRule(ctx context.Context, in *Rule, opts ...grpc.CallOption) (*RuleUpdateResponse, error) {
	out := new(RuleUpdateResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/UpdateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) GetRule(ctx context.Context, in *RuleIdentifier, opts ...grpc.CallOption) (*RuleGetResponse, error) {
	out := new(RuleGetResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/GetRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) ListRules(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RuleListResponse, error) {
	out := new(RuleListResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/ListRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) ValidateWebhook(ctx context.Context, in *URLValidationRequest, opts ...grpc.CallOption) (*URLValidationResponse, error) {
	out := new(URLValidationResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/ValidateWebhook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/notifications.Notifications/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServer is the server API for Notifications service.
type NotificationsServer interface {
	// Publish a notification
	Notify(context.Context, *Event) (*Response, error)
	// Manage notification alerting rules
	AddRule(context.Context, *Rule) (*RuleAddResponse, error)
	DeleteRule(context.Context, *RuleIdentifier) (*RuleDeleteResponse, error)
	UpdateRule(context.Context, *Rule) (*RuleUpdateResponse, error)
	GetRule(context.Context, *RuleIdentifier) (*RuleGetResponse, error)
	ListRules(context.Context, *Empty) (*RuleListResponse, error)
	ValidateWebhook(context.Context, *URLValidationRequest) (*URLValidationResponse, error)
	// Health checks and metadata
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

// UnimplementedNotificationsServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationsServer struct {
}

func (*UnimplementedNotificationsServer) Notify(context.Context, *Event) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (*UnimplementedNotificationsServer) AddRule(context.Context, *Rule) (*RuleAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRule not implemented")
}
func (*UnimplementedNotificationsServer) DeleteRule(context.Context, *RuleIdentifier) (*RuleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRule not implemented")
}
func (*UnimplementedNotificationsServer) UpdateRule(context.Context, *Rule) (*RuleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRule not implemented")
}
func (*UnimplementedNotificationsServer) GetRule(context.Context, *RuleIdentifier) (*RuleGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRule not implemented")
}
func (*UnimplementedNotificationsServer) ListRules(context.Context, *Empty) (*RuleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRules not implemented")
}
func (*UnimplementedNotificationsServer) ValidateWebhook(context.Context, *URLValidationRequest) (*URLValidationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateWebhook not implemented")
}
func (*UnimplementedNotificationsServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}

func RegisterNotificationsServer(s *grpc.Server, srv NotificationsServer) {
	s.RegisterService(&_Notifications_serviceDesc, srv)
}

func _Notifications_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).Notify(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_AddRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).AddRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/AddRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).AddRule(ctx, req.(*Rule))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_DeleteRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).DeleteRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/DeleteRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).DeleteRule(ctx, req.(*RuleIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_UpdateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).UpdateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/UpdateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).UpdateRule(ctx, req.(*Rule))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_GetRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).GetRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/GetRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).GetRule(ctx, req.(*RuleIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_ListRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).ListRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/ListRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).ListRules(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_ValidateWebhook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URLValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).ValidateWebhook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/ValidateWebhook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).ValidateWebhook(ctx, req.(*URLValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifications.Notifications/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notifications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notifications.Notifications",
	HandlerType: (*NotificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notify",
			Handler:    _Notifications_Notify_Handler,
		},
		{
			MethodName: "AddRule",
			Handler:    _Notifications_AddRule_Handler,
		},
		{
			MethodName: "DeleteRule",
			Handler:    _Notifications_DeleteRule_Handler,
		},
		{
			MethodName: "UpdateRule",
			Handler:    _Notifications_UpdateRule_Handler,
		},
		{
			MethodName: "GetRule",
			Handler:    _Notifications_GetRule_Handler,
		},
		{
			MethodName: "ListRules",
			Handler:    _Notifications_ListRules_Handler,
		},
		{
			MethodName: "ValidateWebhook",
			Handler:    _Notifications_ValidateWebhook_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Notifications_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}