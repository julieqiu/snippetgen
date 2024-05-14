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

func ExampleNewStoragePoolsClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewStoragePoolsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleStoragePoolsClient_AggregatedList() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewStoragePoolsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AggregatedListStoragePoolsRequest.
	req := &computepb.AggregatedListStoragePoolsRequest{
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
		_ = it.Response.(*computepb.StoragePoolAggregatedList)
	}
}

func ExampleStoragePoolsClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewStoragePoolsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteStoragePoolRequest.
	req := &computepb.DeleteStoragePoolRequest{
		Project: "",
		RequestId: "",
		StoragePool: "",
		Zone: "",
	}
	resp, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStoragePoolsClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewStoragePoolsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetStoragePoolRequest.
	req := &computepb.GetStoragePoolRequest{
		Project: "",
		StoragePool: "",
		Zone: "",
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStoragePoolsClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewStoragePoolsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetIamPolicyStoragePoolRequest.
	req := &computepb.GetIamPolicyStoragePoolRequest{
		OptionsRequestedPolicyVersion: "",
		Project: "",
		Resource: "",
		Zone: "",
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStoragePoolsClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewStoragePoolsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertStoragePoolRequest.
	req := &computepb.InsertStoragePoolRequest{
		Project: "",
		RequestId: "",
		StoragePoolResource: &computepb.StoragePool{
			CapacityProvisioningType: "",
			CreationTimestamp: "",
			Description: "",
			Id: "",
			Kind: "",
			LabelFingerprint: "",
			Labels: "",
			Name: "",
			PerformanceProvisioningType: "",
			PoolProvisionedCapacityGb: "",
			PoolProvisionedIops: "",
			PoolProvisionedThroughput: "",
			ResourceStatus: &computepb.StoragePoolResourceStatus{
				DiskCount: "",
				LastResizeTimestamp: "",
				MaxTotalProvisionedDiskCapacityGb: "",
				PoolUsedCapacityBytes: "",
				PoolUsedIops: "",
				PoolUsedThroughput: "",
				PoolUserWrittenBytes: "",
				TotalProvisionedDiskCapacityGb: "",
				TotalProvisionedDiskIops: "",
				TotalProvisionedDiskThroughput: "",
			}
			SelfLink: "",
			SelfLinkWithId: "",
			State: "",
			Status: &computepb.StoragePoolResourceStatus{
				DiskCount: "",
				LastResizeTimestamp: "",
				MaxTotalProvisionedDiskCapacityGb: "",
				PoolUsedCapacityBytes: "",
				PoolUsedIops: "",
				PoolUsedThroughput: "",
				PoolUserWrittenBytes: "",
				TotalProvisionedDiskCapacityGb: "",
				TotalProvisionedDiskIops: "",
				TotalProvisionedDiskThroughput: "",
			}
			StoragePoolType: "",
			Zone: "",
		}
		Zone: "",
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStoragePoolsClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewStoragePoolsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListStoragePoolsRequest.
	req := &computepb.ListStoragePoolsRequest{
		Filter: "",
		MaxResults: "",
		OrderBy: "",
		PageToken: "",
		Project: "",
		ReturnPartialSuccess: "",
		Zone: "",
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
		_ = it.Response.(*computepb.StoragePoolList)
	}
}

func ExampleStoragePoolsClient_ListDisks() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewStoragePoolsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListDisksStoragePoolsRequest.
	req := &computepb.ListDisksStoragePoolsRequest{
		Filter: "",
		MaxResults: "",
		OrderBy: "",
		PageToken: "",
		Project: "",
		ReturnPartialSuccess: "",
		StoragePool: "",
		Zone: "",
	}
	it := c.ListDisks(ctx, req)
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
		_ = it.Response.(*computepb.StoragePoolListDisks)
	}
}

func ExampleStoragePoolsClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewStoragePoolsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetIamPolicyStoragePoolRequest.
	req := &computepb.SetIamPolicyStoragePoolRequest{
		Project: "",
		Resource: "",
		Zone: "",
		ZoneSetPolicyRequestResource: &computepb.ZoneSetPolicyRequest{
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
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStoragePoolsClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewStoragePoolsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#TestIamPermissionsStoragePoolRequest.
	req := &computepb.TestIamPermissionsStoragePoolRequest{
		Project: "",
		Resource: "",
		TestPermissionsRequestResource: &computepb.TestPermissionsRequest{
			Permissions: "",
		}
		Zone: "",
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStoragePoolsClient_Update() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewStoragePoolsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#UpdateStoragePoolRequest.
	req := &computepb.UpdateStoragePoolRequest{
		Project: "",
		RequestId: "",
		StoragePool: "",
		StoragePoolResource: &computepb.StoragePool{
			CapacityProvisioningType: "",
			CreationTimestamp: "",
			Description: "",
			Id: "",
			Kind: "",
			LabelFingerprint: "",
			Labels: "",
			Name: "",
			PerformanceProvisioningType: "",
			PoolProvisionedCapacityGb: "",
			PoolProvisionedIops: "",
			PoolProvisionedThroughput: "",
			ResourceStatus: &computepb.StoragePoolResourceStatus{
				DiskCount: "",
				LastResizeTimestamp: "",
				MaxTotalProvisionedDiskCapacityGb: "",
				PoolUsedCapacityBytes: "",
				PoolUsedIops: "",
				PoolUsedThroughput: "",
				PoolUserWrittenBytes: "",
				TotalProvisionedDiskCapacityGb: "",
				TotalProvisionedDiskIops: "",
				TotalProvisionedDiskThroughput: "",
			}
			SelfLink: "",
			SelfLinkWithId: "",
			State: "",
			Status: &computepb.StoragePoolResourceStatus{
				DiskCount: "",
				LastResizeTimestamp: "",
				MaxTotalProvisionedDiskCapacityGb: "",
				PoolUsedCapacityBytes: "",
				PoolUsedIops: "",
				PoolUsedThroughput: "",
				PoolUserWrittenBytes: "",
				TotalProvisionedDiskCapacityGb: "",
				TotalProvisionedDiskIops: "",
				TotalProvisionedDiskThroughput: "",
			}
			StoragePoolType: "",
			Zone: "",
		}
		UpdateMask: "",
		Zone: "",
	}
	resp, err := c.Update(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}