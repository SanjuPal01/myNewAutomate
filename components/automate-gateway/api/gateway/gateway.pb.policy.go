// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: automate-gateway/api/gateway/gateway.proto

package gateway

import policy "github.com/chef/automate/api/external/iam/v2/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.Gateway/GetVersion", "system:service:version", "system:serviceVersion:get", "GET", "/api/v0/gateway/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.Gateway/GetHealth", "system:health", "system:health:get", "GET", "/api/v0/gateway/health", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
