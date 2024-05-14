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

func ExampleNewSubnetworksClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSubnetworksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleSubnetworksClient_AggregatedList() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSubnetworksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AggregatedListSubnetworksRequest.
	req := &computepb.AggregatedListSubnetworksRequest{
		Filter: "",
		IncludeAllScopes: "",
		MaxResults: "",
		OrderBy: "",
		PageToken: "",
		Project: "",
		ReturnPartialSuccess: "",
		ServiceProjectNumber: "",
	}
	it := c.AggregatedList(ctx, req)
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
		_ = it.Response.(*computepb.SubnetworkAggregatedList)
	}
}

func ExampleSubnetworksClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSubnetworksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteSubnetworkRequest.
	req := &computepb.DeleteSubnetworkRequest{
		Project: "",
		Region: "",
		RequestId: "",
		Subnetwork: "",
	}
	resp, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSubnetworksClient_ExpandIpCidrRange() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSubnetworksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ExpandIpCidrRangeSubnetworkRequest.
	req := &computepb.ExpandIpCidrRangeSubnetworkRequest{
		Project: "",
		Region: "",
		RequestId: "",
		Subnetwork: "",
		SubnetworksExpandIpCidrRangeRequestResource: &computepb.SubnetworksExpandIpCidrRangeRequest{
			IpCidrRange: "",
		}
	}
	resp, err := c.ExpandIpCidrRange(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSubnetworksClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSubnetworksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetSubnetworkRequest.
	req := &computepb.GetSubnetworkRequest{
		Project: "",
		Region: "",
		Subnetwork: "",
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSubnetworksClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSubnetworksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetIamPolicySubnetworkRequest.
	req := &computepb.GetIamPolicySubnetworkRequest{
		OptionsRequestedPolicyVersion: "",
		Project: "",
		Region: "",
		Resource: "",
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSubnetworksClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSubnetworksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertSubnetworkRequest.
	req := &computepb.InsertSubnetworkRequest{
		Project: "",
		Region: "",
		RequestId: "",
		SubnetworkResource: &computepb.Subnetwork{
			CreationTimestamp: "",
			Description: "",
			EnableFlowLogs: "",
			ExternalIpv6Prefix: "",
			Fingerprint: "",
			GatewayAddress: "",
			Id: "",
			InternalIpv6Prefix: "",
			IpCidrRange: "",
			Ipv6AccessType: "",
			Ipv6CidrRange: "",
			Kind: "",
			LogConfig: &computepb.SubnetworkLogConfig{
				AggregationInterval: "",
				Enable: "",
				FilterExpr: "",
				FlowSampling: "",
				Metadata: "",
				MetadataFields: "",
			}
			Name: "",
			Network: "",
			PrivateIpGoogleAccess: "",
			PrivateIpv6GoogleAccess: "",
			Purpose: "",
			Region: "",
			ReservedInternalRange: "",
			Role: "",
			SecondaryIpRanges: &computepb.SubnetworkSecondaryRange{
				IpCidrRange: "",
				RangeName: "",
				ReservedInternalRange: "",
			}
			SelfLink: "",
			StackType: "",
			State: "",
		}
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSubnetworksClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSubnetworksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListSubnetworksRequest.
	req := &computepb.ListSubnetworksRequest{
		Filter: "",
		MaxResults: "",
		OrderBy: "",
		PageToken: "",
		Project: "",
		Region: "",
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
		_ = it.Response.(*computepb.SubnetworkList)
	}
}

func ExampleSubnetworksClient_ListUsable() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSubnetworksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListUsableSubnetworksRequest.
	req := &computepb.ListUsableSubnetworksRequest{
		Filter: "",
		MaxResults: "",
		OrderBy: "",
		PageToken: "",
		Project: "",
		ReturnPartialSuccess: "",
	}
	it := c.ListUsable(ctx, req)
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
		_ = it.Response.(*computepb.UsableSubnetworksAggregatedList)
	}
}

func ExampleSubnetworksClient_Patch() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSubnetworksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchSubnetworkRequest.
	req := &computepb.PatchSubnetworkRequest{
		DrainTimeoutSeconds: "",
		Project: "",
		Region: "",
		RequestId: "",
		Subnetwork: "",
		SubnetworkResource: &computepb.Subnetwork{
			CreationTimestamp: "",
			Description: "",
			EnableFlowLogs: "",
			ExternalIpv6Prefix: "",
			Fingerprint: "",
			GatewayAddress: "",
			Id: "",
			InternalIpv6Prefix: "",
			IpCidrRange: "",
			Ipv6AccessType: "",
			Ipv6CidrRange: "",
			Kind: "",
			LogConfig: &computepb.SubnetworkLogConfig{
				AggregationInterval: "",
				Enable: "",
				FilterExpr: "",
				FlowSampling: "",
				Metadata: "",
				MetadataFields: "",
			}
			Name: "",
			Network: "",
			PrivateIpGoogleAccess: "",
			PrivateIpv6GoogleAccess: "",
			Purpose: "",
			Region: "",
			ReservedInternalRange: "",
			Role: "",
			SecondaryIpRanges: &computepb.SubnetworkSecondaryRange{
				IpCidrRange: "",
				RangeName: "",
				ReservedInternalRange: "",
			}
			SelfLink: "",
			StackType: "",
			State: "",
		}
	}
	resp, err := c.Patch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSubnetworksClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSubnetworksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetIamPolicySubnetworkRequest.
	req := &computepb.SetIamPolicySubnetworkRequest{
		Project: "",
		Region: "",
		RegionSetPolicyRequestResource: &computepb.RegionSetPolicyRequest{
			Bindings: &computepb.Binding{
				BindingId: "",
				Condition: &computepb.Expr{...}
				Members: "",
				Role: "",
			}
			Etag: "",
			Policy: &computepb.Policy{
				AuditConfigs: &computepb.AuditConfig{...}
				Bindings: &computepb.Binding{...}
				Etag: "",
				IamOwned: "",
				Rules: &computepb.Rule{...}
				Version: "",
			}
		}
		Resource: "",
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSubnetworksClient_SetPrivateIpGoogleAccess() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSubnetworksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetPrivateIpGoogleAccessSubnetworkRequest.
	req := &computepb.SetPrivateIpGoogleAccessSubnetworkRequest{
		Project: "",
		Region: "",
		RequestId: "",
		Subnetwork: "",
		SubnetworksSetPrivateIpGoogleAccessRequestResource: &computepb.SubnetworksSetPrivateIpGoogleAccessRequest{
			PrivateIpGoogleAccess: "",
		}
	}
	resp, err := c.SetPrivateIpGoogleAccess(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSubnetworksClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSubnetworksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#TestIamPermissionsSubnetworkRequest.
	req := &computepb.TestIamPermissionsSubnetworkRequest{
		Project: "",
		Region: "",
		Resource: "",
		TestPermissionsRequestResource: &computepb.TestPermissionsRequest{
			Permissions: "",
		}
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
