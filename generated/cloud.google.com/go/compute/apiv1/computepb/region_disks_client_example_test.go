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

func ExampleNewRegionDisksClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleRegionDisksClient_AddResourcePolicies() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AddResourcePoliciesRegionDiskRequest.
	req := &computepb.AddResourcePoliciesRegionDiskRequest{
		Disk: "",
		Project: "",
		Region: "",
		RegionDisksAddResourcePoliciesRequestResource: &computepb.RegionDisksAddResourcePoliciesRequest{
			ResourcePolicies: "",
		}
		RequestId: "",
	}
	resp, err := c.AddResourcePolicies(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionDisksClient_BulkInsert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#BulkInsertRegionDiskRequest.
	req := &computepb.BulkInsertRegionDiskRequest{
		BulkInsertDiskResourceResource: &computepb.BulkInsertDiskResource{
			SourceConsistencyGroupPolicy: "",
		}
		Project: "",
		Region: "",
		RequestId: "",
	}
	resp, err := c.BulkInsert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionDisksClient_CreateSnapshot() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#CreateSnapshotRegionDiskRequest.
	req := &computepb.CreateSnapshotRegionDiskRequest{
		Disk: "",
		Project: "",
		Region: "",
		RequestId: "",
		SnapshotResource: &computepb.Snapshot{
			Architecture: "",
			AutoCreated: "",
			ChainName: "",
			CreationSizeBytes: "",
			CreationTimestamp: "",
			Description: "",
			DiskSizeGb: "",
			DownloadBytes: "",
			EnableConfidentialCompute: "",
			GuestOsFeatures: &computepb.GuestOsFeature{
				Type: "",
			}
			Id: "",
			Kind: "",
			LabelFingerprint: "",
			Labels: "",
			LicenseCodes: "",
			Licenses: "",
			LocationHint: "",
			Name: "",
			SatisfiesPzi: "",
			SatisfiesPzs: "",
			SelfLink: "",
			SnapshotEncryptionKey: &computepb.CustomerEncryptionKey{
				KmsKeyName: "",
				KmsKeyServiceAccount: "",
				RawKey: "",
				RsaEncryptedKey: "",
				Sha256: "",
			}
			SnapshotType: "",
			SourceDisk: "",
			SourceDiskEncryptionKey: &computepb.CustomerEncryptionKey{
				KmsKeyName: "",
				KmsKeyServiceAccount: "",
				RawKey: "",
				RsaEncryptedKey: "",
				Sha256: "",
			}
			SourceDiskForRecoveryCheckpoint: "",
			SourceDiskId: "",
			SourceInstantSnapshot: "",
			SourceInstantSnapshotEncryptionKey: &computepb.CustomerEncryptionKey{
				KmsKeyName: "",
				KmsKeyServiceAccount: "",
				RawKey: "",
				RsaEncryptedKey: "",
				Sha256: "",
			}
			SourceInstantSnapshotId: "",
			SourceSnapshotSchedulePolicy: "",
			SourceSnapshotSchedulePolicyId: "",
			Status: "",
			StorageBytes: "",
			StorageBytesStatus: "",
			StorageLocations: "",
		}
	}
	resp, err := c.CreateSnapshot(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionDisksClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteRegionDiskRequest.
	req := &computepb.DeleteRegionDiskRequest{
		Disk: "",
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

func ExampleRegionDisksClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetRegionDiskRequest.
	req := &computepb.GetRegionDiskRequest{
		Disk: "",
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

func ExampleRegionDisksClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetIamPolicyRegionDiskRequest.
	req := &computepb.GetIamPolicyRegionDiskRequest{
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

func ExampleRegionDisksClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertRegionDiskRequest.
	req := &computepb.InsertRegionDiskRequest{
		DiskResource: &computepb.Disk{
			Architecture: "",
			AsyncPrimaryDisk: &computepb.DiskAsyncReplication{
				ConsistencyGroupPolicy: "",
				ConsistencyGroupPolicyId: "",
				Disk: "",
				DiskId: "",
			}
			AsyncSecondaryDisks: "",
			CreationTimestamp: "",
			Description: "",
			DiskEncryptionKey: &computepb.CustomerEncryptionKey{
				KmsKeyName: "",
				KmsKeyServiceAccount: "",
				RawKey: "",
				RsaEncryptedKey: "",
				Sha256: "",
			}
			EnableConfidentialCompute: "",
			GuestOsFeatures: &computepb.GuestOsFeature{
				Type: "",
			}
			Id: "",
			Kind: "",
			LabelFingerprint: "",
			Labels: "",
			LastAttachTimestamp: "",
			LastDetachTimestamp: "",
			LicenseCodes: "",
			Licenses: "",
			LocationHint: "",
			Name: "",
			Options: "",
			Params: &computepb.DiskParams{
				ResourceManagerTags: "",
			}
			PhysicalBlockSizeBytes: "",
			ProvisionedIops: "",
			ProvisionedThroughput: "",
			Region: "",
			ReplicaZones: "",
			ResourcePolicies: "",
			ResourceStatus: &computepb.DiskResourceStatus{
				AsyncPrimaryDisk: &computepb.DiskResourceStatusAsyncReplicationStatus{...}
				AsyncSecondaryDisks: "",
			}
			SatisfiesPzi: "",
			SatisfiesPzs: "",
			SelfLink: "",
			SizeGb: "",
			SourceConsistencyGroupPolicy: "",
			SourceConsistencyGroupPolicyId: "",
			SourceDisk: "",
			SourceDiskId: "",
			SourceImage: "",
			SourceImageEncryptionKey: &computepb.CustomerEncryptionKey{
				KmsKeyName: "",
				KmsKeyServiceAccount: "",
				RawKey: "",
				RsaEncryptedKey: "",
				Sha256: "",
			}
			SourceImageId: "",
			SourceInstantSnapshot: "",
			SourceInstantSnapshotId: "",
			SourceSnapshot: "",
			SourceSnapshotEncryptionKey: &computepb.CustomerEncryptionKey{
				KmsKeyName: "",
				KmsKeyServiceAccount: "",
				RawKey: "",
				RsaEncryptedKey: "",
				Sha256: "",
			}
			SourceSnapshotId: "",
			SourceStorageObject: "",
			Status: "",
			StoragePool: "",
			Type: "",
			Users: "",
			Zone: "",
		}
		Project: "",
		Region: "",
		RequestId: "",
		SourceImage: "",
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionDisksClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListRegionDisksRequest.
	req := &computepb.ListRegionDisksRequest{
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
		_ = it.Response.(*computepb.DiskList)
	}
}

func ExampleRegionDisksClient_RemoveResourcePolicies() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#RemoveResourcePoliciesRegionDiskRequest.
	req := &computepb.RemoveResourcePoliciesRegionDiskRequest{
		Disk: "",
		Project: "",
		Region: "",
		RegionDisksRemoveResourcePoliciesRequestResource: &computepb.RegionDisksRemoveResourcePoliciesRequest{
			ResourcePolicies: "",
		}
		RequestId: "",
	}
	resp, err := c.RemoveResourcePolicies(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionDisksClient_Resize() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ResizeRegionDiskRequest.
	req := &computepb.ResizeRegionDiskRequest{
		Disk: "",
		Project: "",
		Region: "",
		RegionDisksResizeRequestResource: &computepb.RegionDisksResizeRequest{
			SizeGb: "",
		}
		RequestId: "",
	}
	resp, err := c.Resize(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionDisksClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetIamPolicyRegionDiskRequest.
	req := &computepb.SetIamPolicyRegionDiskRequest{
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

func ExampleRegionDisksClient_SetLabels() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetLabelsRegionDiskRequest.
	req := &computepb.SetLabelsRegionDiskRequest{
		Project: "",
		Region: "",
		RegionSetLabelsRequestResource: &computepb.RegionSetLabelsRequest{
			LabelFingerprint: "",
			Labels: "",
		}
		RequestId: "",
		Resource: "",
	}
	resp, err := c.SetLabels(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionDisksClient_StartAsyncReplication() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#StartAsyncReplicationRegionDiskRequest.
	req := &computepb.StartAsyncReplicationRegionDiskRequest{
		Disk: "",
		Project: "",
		Region: "",
		RegionDisksStartAsyncReplicationRequestResource: &computepb.RegionDisksStartAsyncReplicationRequest{
			AsyncSecondaryDisk: "",
		}
		RequestId: "",
	}
	resp, err := c.StartAsyncReplication(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionDisksClient_StopAsyncReplication() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#StopAsyncReplicationRegionDiskRequest.
	req := &computepb.StopAsyncReplicationRegionDiskRequest{
		Disk: "",
		Project: "",
		Region: "",
		RequestId: "",
	}
	resp, err := c.StopAsyncReplication(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionDisksClient_StopGroupAsyncReplication() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#StopGroupAsyncReplicationRegionDiskRequest.
	req := &computepb.StopGroupAsyncReplicationRegionDiskRequest{
		DisksStopGroupAsyncReplicationResourceResource: &computepb.DisksStopGroupAsyncReplicationResource{
			ResourcePolicy: "",
		}
		Project: "",
		Region: "",
		RequestId: "",
	}
	resp, err := c.StopGroupAsyncReplication(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionDisksClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#TestIamPermissionsRegionDiskRequest.
	req := &computepb.TestIamPermissionsRegionDiskRequest{
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

func ExampleRegionDisksClient_Update() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionDisksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#UpdateRegionDiskRequest.
	req := &computepb.UpdateRegionDiskRequest{
		Disk: "",
		DiskResource: &computepb.Disk{
			Architecture: "",
			AsyncPrimaryDisk: &computepb.DiskAsyncReplication{
				ConsistencyGroupPolicy: "",
				ConsistencyGroupPolicyId: "",
				Disk: "",
				DiskId: "",
			}
			AsyncSecondaryDisks: "",
			CreationTimestamp: "",
			Description: "",
			DiskEncryptionKey: &computepb.CustomerEncryptionKey{
				KmsKeyName: "",
				KmsKeyServiceAccount: "",
				RawKey: "",
				RsaEncryptedKey: "",
				Sha256: "",
			}
			EnableConfidentialCompute: "",
			GuestOsFeatures: &computepb.GuestOsFeature{
				Type: "",
			}
			Id: "",
			Kind: "",
			LabelFingerprint: "",
			Labels: "",
			LastAttachTimestamp: "",
			LastDetachTimestamp: "",
			LicenseCodes: "",
			Licenses: "",
			LocationHint: "",
			Name: "",
			Options: "",
			Params: &computepb.DiskParams{
				ResourceManagerTags: "",
			}
			PhysicalBlockSizeBytes: "",
			ProvisionedIops: "",
			ProvisionedThroughput: "",
			Region: "",
			ReplicaZones: "",
			ResourcePolicies: "",
			ResourceStatus: &computepb.DiskResourceStatus{
				AsyncPrimaryDisk: &computepb.DiskResourceStatusAsyncReplicationStatus{...}
				AsyncSecondaryDisks: "",
			}
			SatisfiesPzi: "",
			SatisfiesPzs: "",
			SelfLink: "",
			SizeGb: "",
			SourceConsistencyGroupPolicy: "",
			SourceConsistencyGroupPolicyId: "",
			SourceDisk: "",
			SourceDiskId: "",
			SourceImage: "",
			SourceImageEncryptionKey: &computepb.CustomerEncryptionKey{
				KmsKeyName: "",
				KmsKeyServiceAccount: "",
				RawKey: "",
				RsaEncryptedKey: "",
				Sha256: "",
			}
			SourceImageId: "",
			SourceInstantSnapshot: "",
			SourceInstantSnapshotId: "",
			SourceSnapshot: "",
			SourceSnapshotEncryptionKey: &computepb.CustomerEncryptionKey{
				KmsKeyName: "",
				KmsKeyServiceAccount: "",
				RawKey: "",
				RsaEncryptedKey: "",
				Sha256: "",
			}
			SourceSnapshotId: "",
			SourceStorageObject: "",
			Status: "",
			StoragePool: "",
			Type: "",
			Users: "",
			Zone: "",
		}
		Paths: "",
		Project: "",
		Region: "",
		RequestId: "",
		UpdateMask: "",
	}
	resp, err := c.Update(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
