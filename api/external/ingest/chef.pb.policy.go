// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: external/ingest/chef.proto

package ingest

import (
	policy "github.com/chef/automate/api/external/iam/v2/policy"
	request "github.com/chef/automate/api/external/ingest/request"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.ingest.ChefIngester/ProcessChefRun", "ingest:nodes:{entity_uuid}:runs", "ingest:nodes:create", "POST", "/api/v0/ingest/events/chef/run", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Run); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "run_id":
					return m.RunId
				case "entity_uuid":
					return m.EntityUuid
				case "message_version":
					return m.MessageVersion
				case "message_type":
					return m.MessageType
				case "node_name":
					return m.NodeName
				case "organization_name":
					return m.OrganizationName
				case "start_time":
					return m.StartTime
				case "chef_server_fqdn":
					return m.ChefServerFqdn
				case "end_time":
					return m.EndTime
				case "status":
					return m.Status
				case "source":
					return m.Source
				case "policy_name":
					return m.PolicyName
				case "policy_group":
					return m.PolicyGroup
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.ingest.ChefIngester/ProcessChefAction", "ingest:actions", "ingest:actions:create", "POST", "/api/v0/ingest/events/chef/action", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Action); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "message_type":
					return m.MessageType
				case "message_version":
					return m.MessageVersion
				case "entity_name":
					return m.EntityName
				case "entity_type":
					return m.EntityType
				case "task":
					return m.Task
				case "organization_name":
					return m.OrganizationName
				case "remote_hostname":
					return m.RemoteHostname
				case "run_id":
					return m.RunId
				case "node_id":
					return m.NodeId
				case "recorded_at":
					return m.RecordedAt
				case "remote_request_id":
					return m.RemoteRequestId
				case "request_id":
					return m.RequestId
				case "requestor_name":
					return m.RequestorName
				case "requestor_type":
					return m.RequestorType
				case "service_hostname":
					return m.ServiceHostname
				case "user_agent":
					return m.UserAgent
				case "parent_type":
					return m.ParentType
				case "parent_name":
					return m.ParentName
				case "revision_id":
					return m.RevisionId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.ingest.ChefIngester/ProcessNodeDelete", "ingest:nodes", "ingest:nodes:delete", "POST", "/api/v0/ingest/events/chef/nodedelete", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Delete); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "node_name":
					return m.NodeName
				case "organization_name":
					return m.OrganizationName
				case "remote_hostname":
					return m.RemoteHostname
				case "service_hostname":
					return m.ServiceHostname
				case "node_id":
					return m.NodeId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.ingest.ChefIngester/ProcessMultipleNodeDeletes", "ingest:nodes", "ingest:nodes:delete", "POST", "/api/v0/ingest/events/chef/node-multiple-deletes", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.ingest.ChefIngester/ProcessLivenessPing", "ingest:nodes:{entity_uuid}:liveness", "ingest:nodes:create", "POST", "/api/v0/ingest/events/chef/liveness", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Liveness); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "event_type":
					return m.EventType
				case "entity_uuid":
					return m.EntityUuid
				case "chef_server_fqdn":
					return m.ChefServerFqdn
				case "source":
					return m.Source
				case "message_version":
					return m.MessageVersion
				case "organization_name":
					return m.OrganizationName
				case "node_name":
					return m.NodeName
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.ingest.ChefIngester/GetVersion", "system:service:version", "system:serviceVersion:get", "GET", "/api/v0/ingest/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}