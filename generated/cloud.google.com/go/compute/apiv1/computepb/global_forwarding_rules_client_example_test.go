// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.


package computepb_test

import (
	"context"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/iterator"
)

func ExampleNewGlobalForwardingRulesClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewGlobalForwardingRulesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleGlobalForwardingRulesClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewGlobalForwardingRulesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteGlobalForwardingRuleRequest.
	req := &computepb.DeleteGlobalForwardingRuleRequest{
		ForwardingRule: "",
		Project: "",
		RequestId: "",
	}
	resp, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGlobalForwardingRulesClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewGlobalForwardingRulesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetGlobalForwardingRuleRequest.
	req := &computepb.GetGlobalForwardingRuleRequest{
		ForwardingRule: "",
		Project: "",
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGlobalForwardingRulesClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewGlobalForwardingRulesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertGlobalForwardingRuleRequest.
	req := &computepb.InsertGlobalForwardingRuleRequest{
		ForwardingRuleResource: &computepb.ForwardingRule{
			IPAddress: "",
			IPProtocol: "",
			AllPorts: "",
			AllowGlobalAccess: "",
			AllowPscGlobalAccess: "",
			BackendService: "",
			BaseForwardingRule: "",
			CreationTimestamp: "",
			Description: "",
			Fingerprint: "",
			Id: "",
			IpCollection: "",
			IpVersion: "",
			IsMirroringCollector: "",
			Kind: "",
			LabelFingerprint: "",
			Labels: "",
			LoadBalancingScheme: "",
			MetadataFilters: &computepb.MetadataFilter{
				FilterLabels: &computepb.MetadataFilterLabelMatch{...}
				FilterMatchCriteria: "",
			}
			Name: "",
			Network: "",
			NetworkTier: "",
			NoAutomateDnsZone: "",
			PortRange: "",
			Ports: "",
			PscConnectionId: "",
			PscConnectionStatus: "",
			Region: "",
			SelfLink: "",
			ServiceDirectoryRegistrations: &computepb.ForwardingRuleServiceDirectoryRegistration{
				Namespace: "",
				Service: "",
				ServiceDirectoryRegion: "",
			}
			ServiceLabel: "",
			ServiceName: "",
			SourceIpRanges: "",
			Subnetwork: "",
			Target: "",
		}
		Project: "",
		RequestId: "",
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGlobalForwardingRulesClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewGlobalForwardingRulesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListGlobalForwardingRulesRequest.
	req := &computepb.ListGlobalForwardingRulesRequest{
		Filter: "",
		MaxResults: "",
		OrderBy: "",
		PageToken: "",
		Project: "",
		ReturnPartialSuccess: "",
	}
	it := c.List(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*computepb.ForwardingRuleList)
	}
}

func ExampleGlobalForwardingRulesClient_Patch() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewGlobalForwardingRulesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchGlobalForwardingRuleRequest.
	req := &computepb.PatchGlobalForwardingRuleRequest{
		ForwardingRule: "",
		ForwardingRuleResource: &computepb.ForwardingRule{
			IPAddress: "",
			IPProtocol: "",
			AllPorts: "",
			AllowGlobalAccess: "",
			AllowPscGlobalAccess: "",
			BackendService: "",
			BaseForwardingRule: "",
			CreationTimestamp: "",
			Description: "",
			Fingerprint: "",
			Id: "",
			IpCollection: "",
			IpVersion: "",
			IsMirroringCollector: "",
			Kind: "",
			LabelFingerprint: "",
			Labels: "",
			LoadBalancingScheme: "",
			MetadataFilters: &computepb.MetadataFilter{
				FilterLabels: &computepb.MetadataFilterLabelMatch{...}
				FilterMatchCriteria: "",
			}
			Name: "",
			Network: "",
			NetworkTier: "",
			NoAutomateDnsZone: "",
			PortRange: "",
			Ports: "",
			PscConnectionId: "",
			PscConnectionStatus: "",
			Region: "",
			SelfLink: "",
			ServiceDirectoryRegistrations: &computepb.ForwardingRuleServiceDirectoryRegistration{
				Namespace: "",
				Service: "",
				ServiceDirectoryRegion: "",
			}
			ServiceLabel: "",
			ServiceName: "",
			SourceIpRanges: "",
			Subnetwork: "",
			Target: "",
		}
		Project: "",
		RequestId: "",
	}
	resp, err := c.Patch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGlobalForwardingRulesClient_SetLabels() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewGlobalForwardingRulesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetLabelsGlobalForwardingRuleRequest.
	req := &computepb.SetLabelsGlobalForwardingRuleRequest{
		GlobalSetLabelsRequestResource: &computepb.GlobalSetLabelsRequest{
			LabelFingerprint: "",
			Labels: "",
		}
		Project: "",
		Resource: "",
	}
	resp, err := c.SetLabels(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleGlobalForwardingRulesClient_SetTarget() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewGlobalForwardingRulesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetTargetGlobalForwardingRuleRequest.
	req := &computepb.SetTargetGlobalForwardingRuleRequest{
		ForwardingRule: "",
		Project: "",
		RequestId: "",
		TargetReferenceResource: &computepb.TargetReference{
			Target: "",
		}
	}
	resp, err := c.SetTarget(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
