// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: automate-gateway/api/legacy/legacy.proto

package legacy

import policy "github.com/chef/automate/api/external/iam/v2/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.legacy.LegacyDataCollector/Status", "ingest:status", "ingest:status:get", "GET", "/api/v0/events/data-collector", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}