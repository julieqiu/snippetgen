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

func ExampleNewNetworkAttachmentsClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkAttachmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNetworkAttachmentsClient_AggregatedList() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkAttachmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AggregatedListNetworkAttachmentsRequest.
	req := &computepb.AggregatedListNetworkAttachmentsRequest{
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
		_ = it.Response.(*computepb.NetworkAttachmentAggregatedList)
	}
}

func ExampleNetworkAttachmentsClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkAttachmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteNetworkAttachmentRequest.
	req := &computepb.DeleteNetworkAttachmentRequest{
		NetworkAttachment: "",
		Project: "",
		Region: "",
		RequestId: "",
	}
	resp, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkAttachmentsClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkAttachmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetNetworkAttachmentRequest.
	req := &computepb.GetNetworkAttachmentRequest{
		NetworkAttachment: "",
		Project: "",
		Region: "",
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkAttachmentsClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkAttachmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetIamPolicyNetworkAttachmentRequest.
	req := &computepb.GetIamPolicyNetworkAttachmentRequest{
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

func ExampleNetworkAttachmentsClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkAttachmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertNetworkAttachmentRequest.
	req := &computepb.InsertNetworkAttachmentRequest{
		NetworkAttachmentResource: &computepb.NetworkAttachment{
			ConnectionEndpoints: &computepb.NetworkAttachmentConnectedEndpoint{
				IpAddress: "",
				Ipv6Address: "",
				ProjectIdOrNum: "",
				SecondaryIpCidrRanges: "",
				Status: "",
				Subnetwork: "",
				SubnetworkCidrRange: "",
			}
			ConnectionPreference: "",
			CreationTimestamp: "",
			Description: "",
			Fingerprint: "",
			Id: "",
			Kind: "",
			Name: "",
			Network: "",
			ProducerAcceptLists: "",
			ProducerRejectLists: "",
			Region: "",
			SelfLink: "",
			SelfLinkWithId: "",
			Subnetworks: "",
		}
		Project: "",
		Region: "",
		RequestId: "",
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkAttachmentsClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkAttachmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListNetworkAttachmentsRequest.
	req := &computepb.ListNetworkAttachmentsRequest{
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
		_ = it.Response.(*computepb.NetworkAttachmentList)
	}
}

func ExampleNetworkAttachmentsClient_Patch() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkAttachmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchNetworkAttachmentRequest.
	req := &computepb.PatchNetworkAttachmentRequest{
		NetworkAttachment: "",
		NetworkAttachmentResource: &computepb.NetworkAttachment{
			ConnectionEndpoints: &computepb.NetworkAttachmentConnectedEndpoint{
				IpAddress: "",
				Ipv6Address: "",
				ProjectIdOrNum: "",
				SecondaryIpCidrRanges: "",
				Status: "",
				Subnetwork: "",
				SubnetworkCidrRange: "",
			}
			ConnectionPreference: "",
			CreationTimestamp: "",
			Description: "",
			Fingerprint: "",
			Id: "",
			Kind: "",
			Name: "",
			Network: "",
			ProducerAcceptLists: "",
			ProducerRejectLists: "",
			Region: "",
			SelfLink: "",
			SelfLinkWithId: "",
			Subnetworks: "",
		}
		Project: "",
		Region: "",
		RequestId: "",
	}
	resp, err := c.Patch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkAttachmentsClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkAttachmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetIamPolicyNetworkAttachmentRequest.
	req := &computepb.SetIamPolicyNetworkAttachmentRequest{
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

func ExampleNetworkAttachmentsClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkAttachmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#TestIamPermissionsNetworkAttachmentRequest.
	req := &computepb.TestIamPermissionsNetworkAttachmentRequest{
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
