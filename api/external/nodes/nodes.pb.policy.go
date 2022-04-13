// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: external/nodes/nodes.proto

package nodes

import policy "github.com/chef/automate/api/external/iam/v2/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/Create", "infra:nodes", "infra:nodes:create", "POST", "/api/v0/nodes", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Node); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "platform":
					return m.Platform
				case "platform_version":
					return m.PlatformVersion
				case "manager":
					return m.Manager
				case "status":
					return m.Status
				case "connection_error":
					return m.ConnectionError
				case "state":
					return m.State
				case "name_prefix":
					return m.NamePrefix
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/Read", "infra:nodes:{id}", "infra:nodes:get", "GET", "/api/v0/nodes/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/Update", "infra:nodes:{id}", "infra:nodes:update", "PUT", "/api/v0/nodes/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Node); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "platform":
					return m.Platform
				case "platform_version":
					return m.PlatformVersion
				case "manager":
					return m.Manager
				case "status":
					return m.Status
				case "connection_error":
					return m.ConnectionError
				case "state":
					return m.State
				case "name_prefix":
					return m.NamePrefix
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/Delete", "infra:nodes:{id}", "infra:nodes:delete", "DELETE", "/api/v0/nodes/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/BulkDeleteById", "infra:nodes", "infra:nodes:delete", "POST", "/api/v0/nodes/delete/ids", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/List", "infra:nodes", "infra:nodes:list", "POST", "/api/v0/nodes/search", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/Rerun", "infra:nodes:{id}", "infra:nodes:rerun", "GET", "/api/v0/nodes/rerun/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/BulkDelete", "infra:nodes", "infra:nodes:delete", "POST", "/api/v0/nodes/delete", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.nodes.v1.NodesService/BulkCreate", "infra:nodes", "infra:nodes:create", "POST", "/api/v0/nodes/bulk-create", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
