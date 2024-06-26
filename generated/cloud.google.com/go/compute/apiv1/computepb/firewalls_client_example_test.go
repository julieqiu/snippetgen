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

func ExampleNewFirewallsClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleFirewallsClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteFirewallRequest.
	req := &computepb.DeleteFirewallRequest{
		Firewall: "",
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

func ExampleFirewallsClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetFirewallRequest.
	req := &computepb.GetFirewallRequest{
		Firewall: "",
		Project: "",
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallsClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertFirewallRequest.
	req := &computepb.InsertFirewallRequest{
		FirewallResource: &computepb.Firewall{
			Allowed: &computepb.Allowed{
				IPProtocol: "",
				Ports: "",
			}
			CreationTimestamp: "",
			Denied: &computepb.Denied{
				IPProtocol: "",
				Ports: "",
			}
			Description: "",
			DestinationRanges: "",
			Direction: "",
			Disabled: "",
			Id: "",
			Kind: "",
			LogConfig: &computepb.FirewallLogConfig{
				Enable: "",
				Metadata: "",
			}
			Name: "",
			Network: "",
			Priority: "",
			SelfLink: "",
			SourceRanges: "",
			SourceServiceAccounts: "",
			SourceTags: "",
			TargetServiceAccounts: "",
			TargetTags: "",
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

func ExampleFirewallsClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListFirewallsRequest.
	req := &computepb.ListFirewallsRequest{
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
		_ = it.Response.(*computepb.FirewallList)
	}
}

func ExampleFirewallsClient_Patch() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchFirewallRequest.
	req := &computepb.PatchFirewallRequest{
		Firewall: "",
		FirewallResource: &computepb.Firewall{
			Allowed: &computepb.Allowed{
				IPProtocol: "",
				Ports: "",
			}
			CreationTimestamp: "",
			Denied: &computepb.Denied{
				IPProtocol: "",
				Ports: "",
			}
			Description: "",
			DestinationRanges: "",
			Direction: "",
			Disabled: "",
			Id: "",
			Kind: "",
			LogConfig: &computepb.FirewallLogConfig{
				Enable: "",
				Metadata: "",
			}
			Name: "",
			Network: "",
			Priority: "",
			SelfLink: "",
			SourceRanges: "",
			SourceServiceAccounts: "",
			SourceTags: "",
			TargetServiceAccounts: "",
			TargetTags: "",
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

func ExampleFirewallsClient_Update() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#UpdateFirewallRequest.
	req := &computepb.UpdateFirewallRequest{
		Firewall: "",
		FirewallResource: &computepb.Firewall{
			Allowed: &computepb.Allowed{
				IPProtocol: "",
				Ports: "",
			}
			CreationTimestamp: "",
			Denied: &computepb.Denied{
				IPProtocol: "",
				Ports: "",
			}
			Description: "",
			DestinationRanges: "",
			Direction: "",
			Disabled: "",
			Id: "",
			Kind: "",
			LogConfig: &computepb.FirewallLogConfig{
				Enable: "",
				Metadata: "",
			}
			Name: "",
			Network: "",
			Priority: "",
			SelfLink: "",
			SourceRanges: "",
			SourceServiceAccounts: "",
			SourceTags: "",
			TargetServiceAccounts: "",
			TargetTags: "",
		}
		Project: "",
		RequestId: "",
	}
	resp, err := c.Update(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
