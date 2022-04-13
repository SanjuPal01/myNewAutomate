// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: external/iam/v2/policy.proto

package v2

import (
	"context"

	request "github.com/chef/automate/api/external/iam/v2/request"
	response "github.com/chef/automate/api/external/iam/v2/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the PoliciesServer interface (at compile time)
var _ PoliciesServer = &PoliciesServerMock{}

// NewPoliciesServerMock gives you a fresh instance of PoliciesServerMock.
func NewPoliciesServerMock() *PoliciesServerMock {
	return &PoliciesServerMock{validateRequests: true}
}

// NewPoliciesServerMockWithoutValidation gives you a fresh instance of
// PoliciesServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewPoliciesServerMockWithoutValidation() *PoliciesServerMock {
	return &PoliciesServerMock{}
}

// PoliciesServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type PoliciesServerMock struct {
	validateRequests          bool
	CreatePolicyFunc          func(context.Context, *request.CreatePolicyReq) (*response.CreatePolicyResp, error)
	GetPolicyFunc             func(context.Context, *request.GetPolicyReq) (*response.GetPolicyResp, error)
	ListPoliciesFunc          func(context.Context, *request.ListPoliciesReq) (*response.ListPoliciesResp, error)
	DeletePolicyFunc          func(context.Context, *request.DeletePolicyReq) (*response.DeletePolicyResp, error)
	UpdatePolicyFunc          func(context.Context, *request.UpdatePolicyReq) (*response.UpdatePolicyResp, error)
	GetPolicyVersionFunc      func(context.Context, *request.GetPolicyVersionReq) (*response.GetPolicyVersionResp, error)
	ListPolicyMembersFunc     func(context.Context, *request.ListPolicyMembersReq) (*response.ListPolicyMembersResp, error)
	ReplacePolicyMembersFunc  func(context.Context, *request.ReplacePolicyMembersReq) (*response.ReplacePolicyMembersResp, error)
	RemovePolicyMembersFunc   func(context.Context, *request.RemovePolicyMembersReq) (*response.RemovePolicyMembersResp, error)
	AddPolicyMembersFunc      func(context.Context, *request.AddPolicyMembersReq) (*response.AddPolicyMembersResp, error)
	CreateRoleFunc            func(context.Context, *request.CreateRoleReq) (*response.CreateRoleResp, error)
	ListRolesFunc             func(context.Context, *request.ListRolesReq) (*response.ListRolesResp, error)
	GetRoleFunc               func(context.Context, *request.GetRoleReq) (*response.GetRoleResp, error)
	DeleteRoleFunc            func(context.Context, *request.DeleteRoleReq) (*response.DeleteRoleResp, error)
	UpdateRoleFunc            func(context.Context, *request.UpdateRoleReq) (*response.UpdateRoleResp, error)
	CreateProjectFunc         func(context.Context, *request.CreateProjectReq) (*response.CreateProjectResp, error)
	UpdateProjectFunc         func(context.Context, *request.UpdateProjectReq) (*response.UpdateProjectResp, error)
	GetProjectFunc            func(context.Context, *request.GetProjectReq) (*response.GetProjectResp, error)
	ListProjectsFunc          func(context.Context, *request.ListProjectsReq) (*response.ListProjectsResp, error)
	DeleteProjectFunc         func(context.Context, *request.DeleteProjectReq) (*response.DeleteProjectResp, error)
	IntrospectAllProjectsFunc func(context.Context, *request.ListProjectsReq) (*response.ListProjectsResp, error)
}

func (m *PoliciesServerMock) CreatePolicy(ctx context.Context, req *request.CreatePolicyReq) (*response.CreatePolicyResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.CreatePolicyFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'CreatePolicy' not implemented")
}

func (m *PoliciesServerMock) GetPolicy(ctx context.Context, req *request.GetPolicyReq) (*response.GetPolicyResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetPolicyFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetPolicy' not implemented")
}

func (m *PoliciesServerMock) ListPolicies(ctx context.Context, req *request.ListPoliciesReq) (*response.ListPoliciesResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ListPoliciesFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ListPolicies' not implemented")
}

func (m *PoliciesServerMock) DeletePolicy(ctx context.Context, req *request.DeletePolicyReq) (*response.DeletePolicyResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.DeletePolicyFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'DeletePolicy' not implemented")
}

func (m *PoliciesServerMock) UpdatePolicy(ctx context.Context, req *request.UpdatePolicyReq) (*response.UpdatePolicyResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.UpdatePolicyFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'UpdatePolicy' not implemented")
}

func (m *PoliciesServerMock) GetPolicyVersion(ctx context.Context, req *request.GetPolicyVersionReq) (*response.GetPolicyVersionResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetPolicyVersionFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetPolicyVersion' not implemented")
}

func (m *PoliciesServerMock) ListPolicyMembers(ctx context.Context, req *request.ListPolicyMembersReq) (*response.ListPolicyMembersResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ListPolicyMembersFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ListPolicyMembers' not implemented")
}

func (m *PoliciesServerMock) ReplacePolicyMembers(ctx context.Context, req *request.ReplacePolicyMembersReq) (*response.ReplacePolicyMembersResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ReplacePolicyMembersFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ReplacePolicyMembers' not implemented")
}

func (m *PoliciesServerMock) RemovePolicyMembers(ctx context.Context, req *request.RemovePolicyMembersReq) (*response.RemovePolicyMembersResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.RemovePolicyMembersFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'RemovePolicyMembers' not implemented")
}

func (m *PoliciesServerMock) AddPolicyMembers(ctx context.Context, req *request.AddPolicyMembersReq) (*response.AddPolicyMembersResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.AddPolicyMembersFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'AddPolicyMembers' not implemented")
}

func (m *PoliciesServerMock) CreateRole(ctx context.Context, req *request.CreateRoleReq) (*response.CreateRoleResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.CreateRoleFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'CreateRole' not implemented")
}

func (m *PoliciesServerMock) ListRoles(ctx context.Context, req *request.ListRolesReq) (*response.ListRolesResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ListRolesFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ListRoles' not implemented")
}

func (m *PoliciesServerMock) GetRole(ctx context.Context, req *request.GetRoleReq) (*response.GetRoleResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetRoleFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetRole' not implemented")
}

func (m *PoliciesServerMock) DeleteRole(ctx context.Context, req *request.DeleteRoleReq) (*response.DeleteRoleResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.DeleteRoleFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'DeleteRole' not implemented")
}

func (m *PoliciesServerMock) UpdateRole(ctx context.Context, req *request.UpdateRoleReq) (*response.UpdateRoleResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.UpdateRoleFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'UpdateRole' not implemented")
}

func (m *PoliciesServerMock) CreateProject(ctx context.Context, req *request.CreateProjectReq) (*response.CreateProjectResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.CreateProjectFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'CreateProject' not implemented")
}

func (m *PoliciesServerMock) UpdateProject(ctx context.Context, req *request.UpdateProjectReq) (*response.UpdateProjectResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.UpdateProjectFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'UpdateProject' not implemented")
}

func (m *PoliciesServerMock) GetProject(ctx context.Context, req *request.GetProjectReq) (*response.GetProjectResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetProjectFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetProject' not implemented")
}

func (m *PoliciesServerMock) ListProjects(ctx context.Context, req *request.ListProjectsReq) (*response.ListProjectsResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ListProjectsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ListProjects' not implemented")
}

func (m *PoliciesServerMock) DeleteProject(ctx context.Context, req *request.DeleteProjectReq) (*response.DeleteProjectResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.DeleteProjectFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'DeleteProject' not implemented")
}

func (m *PoliciesServerMock) IntrospectAllProjects(ctx context.Context, req *request.ListProjectsReq) (*response.ListProjectsResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.IntrospectAllProjectsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'IntrospectAllProjects' not implemented")
}

// Reset resets all overridden functions
func (m *PoliciesServerMock) Reset() {
	m.CreatePolicyFunc = nil
	m.GetPolicyFunc = nil
	m.ListPoliciesFunc = nil
	m.DeletePolicyFunc = nil
	m.UpdatePolicyFunc = nil
	m.GetPolicyVersionFunc = nil
	m.ListPolicyMembersFunc = nil
	m.ReplacePolicyMembersFunc = nil
	m.RemovePolicyMembersFunc = nil
	m.AddPolicyMembersFunc = nil
	m.CreateRoleFunc = nil
	m.ListRolesFunc = nil
	m.GetRoleFunc = nil
	m.DeleteRoleFunc = nil
	m.UpdateRoleFunc = nil
	m.CreateProjectFunc = nil
	m.UpdateProjectFunc = nil
	m.GetProjectFunc = nil
	m.ListProjectsFunc = nil
	m.DeleteProjectFunc = nil
	m.IntrospectAllProjectsFunc = nil
}
