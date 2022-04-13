// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: external/iam/v2/tokens.proto

package v2

import (
	"context"

	request "github.com/chef/automate/api/external/iam/v2/request"
	response "github.com/chef/automate/api/external/iam/v2/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the TokensServer interface (at compile time)
var _ TokensServer = &TokensServerMock{}

// NewTokensServerMock gives you a fresh instance of TokensServerMock.
func NewTokensServerMock() *TokensServerMock {
	return &TokensServerMock{validateRequests: true}
}

// NewTokensServerMockWithoutValidation gives you a fresh instance of
// TokensServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewTokensServerMockWithoutValidation() *TokensServerMock {
	return &TokensServerMock{}
}

// TokensServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type TokensServerMock struct {
	validateRequests bool
	CreateTokenFunc  func(context.Context, *request.CreateTokenReq) (*response.CreateTokenResp, error)
	GetTokenFunc     func(context.Context, *request.GetTokenReq) (*response.GetTokenResp, error)
	UpdateTokenFunc  func(context.Context, *request.UpdateTokenReq) (*response.UpdateTokenResp, error)
	DeleteTokenFunc  func(context.Context, *request.DeleteTokenReq) (*response.DeleteTokenResp, error)
	ListTokensFunc   func(context.Context, *request.ListTokensReq) (*response.ListTokensResp, error)
}

func (m *TokensServerMock) CreateToken(ctx context.Context, req *request.CreateTokenReq) (*response.CreateTokenResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.CreateTokenFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'CreateToken' not implemented")
}

func (m *TokensServerMock) GetToken(ctx context.Context, req *request.GetTokenReq) (*response.GetTokenResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetTokenFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetToken' not implemented")
}

func (m *TokensServerMock) UpdateToken(ctx context.Context, req *request.UpdateTokenReq) (*response.UpdateTokenResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.UpdateTokenFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'UpdateToken' not implemented")
}

func (m *TokensServerMock) DeleteToken(ctx context.Context, req *request.DeleteTokenReq) (*response.DeleteTokenResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.DeleteTokenFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'DeleteToken' not implemented")
}

func (m *TokensServerMock) ListTokens(ctx context.Context, req *request.ListTokensReq) (*response.ListTokensResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ListTokensFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ListTokens' not implemented")
}

// Reset resets all overridden functions
func (m *TokensServerMock) Reset() {
	m.CreateTokenFunc = nil
	m.GetTokenFunc = nil
	m.UpdateTokenFunc = nil
	m.DeleteTokenFunc = nil
	m.ListTokensFunc = nil
}
