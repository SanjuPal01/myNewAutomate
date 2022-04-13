// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chef/automate/api/interservice/event_feed (interfaces: EventFeedServiceClient,EventFeedService_EventExportClient,EventFeedServiceServer,EventFeedService_EventExportServer)

// Package event_feed is a generated GoMock package.
package event_feed

import (
	context "context"
	event "github.com/chef/automate/api/interservice/event"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockEventFeedServiceClient is a mock of EventFeedServiceClient interface
type MockEventFeedServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockEventFeedServiceClientMockRecorder
}

// MockEventFeedServiceClientMockRecorder is the mock recorder for MockEventFeedServiceClient
type MockEventFeedServiceClientMockRecorder struct {
	mock *MockEventFeedServiceClient
}

// NewMockEventFeedServiceClient creates a new mock instance
func NewMockEventFeedServiceClient(ctrl *gomock.Controller) *MockEventFeedServiceClient {
	mock := &MockEventFeedServiceClient{ctrl: ctrl}
	mock.recorder = &MockEventFeedServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventFeedServiceClient) EXPECT() *MockEventFeedServiceClientMockRecorder {
	return m.recorder
}

// EventExport mocks base method
func (m *MockEventFeedServiceClient) EventExport(arg0 context.Context, arg1 *EventExportRequest, arg2 ...grpc.CallOption) (EventFeedService_EventExportClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EventExport", varargs...)
	ret0, _ := ret[0].(EventFeedService_EventExportClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EventExport indicates an expected call of EventExport
func (mr *MockEventFeedServiceClientMockRecorder) EventExport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventExport", reflect.TypeOf((*MockEventFeedServiceClient)(nil).EventExport), varargs...)
}

// GetFeed mocks base method
func (m *MockEventFeedServiceClient) GetFeed(arg0 context.Context, arg1 *FeedRequest, arg2 ...grpc.CallOption) (*FeedResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFeed", varargs...)
	ret0, _ := ret[0].(*FeedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeed indicates an expected call of GetFeed
func (mr *MockEventFeedServiceClientMockRecorder) GetFeed(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeed", reflect.TypeOf((*MockEventFeedServiceClient)(nil).GetFeed), varargs...)
}

// GetFeedSummary mocks base method
func (m *MockEventFeedServiceClient) GetFeedSummary(arg0 context.Context, arg1 *FeedSummaryRequest, arg2 ...grpc.CallOption) (*FeedSummaryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFeedSummary", varargs...)
	ret0, _ := ret[0].(*FeedSummaryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeedSummary indicates an expected call of GetFeedSummary
func (mr *MockEventFeedServiceClientMockRecorder) GetFeedSummary(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeedSummary", reflect.TypeOf((*MockEventFeedServiceClient)(nil).GetFeedSummary), varargs...)
}

// GetFeedTimeline mocks base method
func (m *MockEventFeedServiceClient) GetFeedTimeline(arg0 context.Context, arg1 *FeedTimelineRequest, arg2 ...grpc.CallOption) (*FeedTimelineResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFeedTimeline", varargs...)
	ret0, _ := ret[0].(*FeedTimelineResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeedTimeline indicates an expected call of GetFeedTimeline
func (mr *MockEventFeedServiceClientMockRecorder) GetFeedTimeline(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeedTimeline", reflect.TypeOf((*MockEventFeedServiceClient)(nil).GetFeedTimeline), varargs...)
}

// HandleEvent mocks base method
func (m *MockEventFeedServiceClient) HandleEvent(arg0 context.Context, arg1 *event.EventMsg, arg2 ...grpc.CallOption) (*event.EventResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HandleEvent", varargs...)
	ret0, _ := ret[0].(*event.EventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleEvent indicates an expected call of HandleEvent
func (mr *MockEventFeedServiceClientMockRecorder) HandleEvent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEvent", reflect.TypeOf((*MockEventFeedServiceClient)(nil).HandleEvent), varargs...)
}

// MockEventFeedService_EventExportClient is a mock of EventFeedService_EventExportClient interface
type MockEventFeedService_EventExportClient struct {
	ctrl     *gomock.Controller
	recorder *MockEventFeedService_EventExportClientMockRecorder
}

// MockEventFeedService_EventExportClientMockRecorder is the mock recorder for MockEventFeedService_EventExportClient
type MockEventFeedService_EventExportClientMockRecorder struct {
	mock *MockEventFeedService_EventExportClient
}

// NewMockEventFeedService_EventExportClient creates a new mock instance
func NewMockEventFeedService_EventExportClient(ctrl *gomock.Controller) *MockEventFeedService_EventExportClient {
	mock := &MockEventFeedService_EventExportClient{ctrl: ctrl}
	mock.recorder = &MockEventFeedService_EventExportClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventFeedService_EventExportClient) EXPECT() *MockEventFeedService_EventExportClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockEventFeedService_EventExportClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockEventFeedService_EventExportClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockEventFeedService_EventExportClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockEventFeedService_EventExportClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockEventFeedService_EventExportClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockEventFeedService_EventExportClient)(nil).Context))
}

// Header mocks base method
func (m *MockEventFeedService_EventExportClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockEventFeedService_EventExportClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockEventFeedService_EventExportClient)(nil).Header))
}

// Recv mocks base method
func (m *MockEventFeedService_EventExportClient) Recv() (*EventExportResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*EventExportResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockEventFeedService_EventExportClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockEventFeedService_EventExportClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockEventFeedService_EventExportClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockEventFeedService_EventExportClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockEventFeedService_EventExportClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockEventFeedService_EventExportClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockEventFeedService_EventExportClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockEventFeedService_EventExportClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockEventFeedService_EventExportClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockEventFeedService_EventExportClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockEventFeedService_EventExportClient)(nil).Trailer))
}

// MockEventFeedServiceServer is a mock of EventFeedServiceServer interface
type MockEventFeedServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockEventFeedServiceServerMockRecorder
}

// MockEventFeedServiceServerMockRecorder is the mock recorder for MockEventFeedServiceServer
type MockEventFeedServiceServerMockRecorder struct {
	mock *MockEventFeedServiceServer
}

// NewMockEventFeedServiceServer creates a new mock instance
func NewMockEventFeedServiceServer(ctrl *gomock.Controller) *MockEventFeedServiceServer {
	mock := &MockEventFeedServiceServer{ctrl: ctrl}
	mock.recorder = &MockEventFeedServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventFeedServiceServer) EXPECT() *MockEventFeedServiceServerMockRecorder {
	return m.recorder
}

// EventExport mocks base method
func (m *MockEventFeedServiceServer) EventExport(arg0 *EventExportRequest, arg1 EventFeedService_EventExportServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventExport", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EventExport indicates an expected call of EventExport
func (mr *MockEventFeedServiceServerMockRecorder) EventExport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventExport", reflect.TypeOf((*MockEventFeedServiceServer)(nil).EventExport), arg0, arg1)
}

// GetFeed mocks base method
func (m *MockEventFeedServiceServer) GetFeed(arg0 context.Context, arg1 *FeedRequest) (*FeedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeed", arg0, arg1)
	ret0, _ := ret[0].(*FeedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeed indicates an expected call of GetFeed
func (mr *MockEventFeedServiceServerMockRecorder) GetFeed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeed", reflect.TypeOf((*MockEventFeedServiceServer)(nil).GetFeed), arg0, arg1)
}

// GetFeedSummary mocks base method
func (m *MockEventFeedServiceServer) GetFeedSummary(arg0 context.Context, arg1 *FeedSummaryRequest) (*FeedSummaryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeedSummary", arg0, arg1)
	ret0, _ := ret[0].(*FeedSummaryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeedSummary indicates an expected call of GetFeedSummary
func (mr *MockEventFeedServiceServerMockRecorder) GetFeedSummary(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeedSummary", reflect.TypeOf((*MockEventFeedServiceServer)(nil).GetFeedSummary), arg0, arg1)
}

// GetFeedTimeline mocks base method
func (m *MockEventFeedServiceServer) GetFeedTimeline(arg0 context.Context, arg1 *FeedTimelineRequest) (*FeedTimelineResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeedTimeline", arg0, arg1)
	ret0, _ := ret[0].(*FeedTimelineResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeedTimeline indicates an expected call of GetFeedTimeline
func (mr *MockEventFeedServiceServerMockRecorder) GetFeedTimeline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeedTimeline", reflect.TypeOf((*MockEventFeedServiceServer)(nil).GetFeedTimeline), arg0, arg1)
}

// HandleEvent mocks base method
func (m *MockEventFeedServiceServer) HandleEvent(arg0 context.Context, arg1 *event.EventMsg) (*event.EventResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleEvent", arg0, arg1)
	ret0, _ := ret[0].(*event.EventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleEvent indicates an expected call of HandleEvent
func (mr *MockEventFeedServiceServerMockRecorder) HandleEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEvent", reflect.TypeOf((*MockEventFeedServiceServer)(nil).HandleEvent), arg0, arg1)
}

// MockEventFeedService_EventExportServer is a mock of EventFeedService_EventExportServer interface
type MockEventFeedService_EventExportServer struct {
	ctrl     *gomock.Controller
	recorder *MockEventFeedService_EventExportServerMockRecorder
}

// MockEventFeedService_EventExportServerMockRecorder is the mock recorder for MockEventFeedService_EventExportServer
type MockEventFeedService_EventExportServerMockRecorder struct {
	mock *MockEventFeedService_EventExportServer
}

// NewMockEventFeedService_EventExportServer creates a new mock instance
func NewMockEventFeedService_EventExportServer(ctrl *gomock.Controller) *MockEventFeedService_EventExportServer {
	mock := &MockEventFeedService_EventExportServer{ctrl: ctrl}
	mock.recorder = &MockEventFeedService_EventExportServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventFeedService_EventExportServer) EXPECT() *MockEventFeedService_EventExportServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockEventFeedService_EventExportServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockEventFeedService_EventExportServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockEventFeedService_EventExportServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockEventFeedService_EventExportServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockEventFeedService_EventExportServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockEventFeedService_EventExportServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockEventFeedService_EventExportServer) Send(arg0 *EventExportResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockEventFeedService_EventExportServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockEventFeedService_EventExportServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockEventFeedService_EventExportServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockEventFeedService_EventExportServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockEventFeedService_EventExportServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockEventFeedService_EventExportServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockEventFeedService_EventExportServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockEventFeedService_EventExportServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockEventFeedService_EventExportServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockEventFeedService_EventExportServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockEventFeedService_EventExportServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockEventFeedService_EventExportServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockEventFeedService_EventExportServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockEventFeedService_EventExportServer)(nil).SetTrailer), arg0)
}