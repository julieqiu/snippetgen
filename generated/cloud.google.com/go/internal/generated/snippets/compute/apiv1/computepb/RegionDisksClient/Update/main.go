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

// [START compute_v1_generated_RegionDisks_Update_sync]

package main

import (
	"context"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
)

func main() {
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

// [END compute_v1_generated_RegionDisks_Update_sync]
