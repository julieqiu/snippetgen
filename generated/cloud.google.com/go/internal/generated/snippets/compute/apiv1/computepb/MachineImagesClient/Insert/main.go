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

// [START compute_v1_generated_MachineImages_Insert_sync]

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
	c, err := computepb.NewMachineImagesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertMachineImageRequest.
	req := &computepb.InsertMachineImageRequest{
		MachineImageResource: &computepb.MachineImage{
			CreationTimestamp: "",
			Description: "",
			GuestFlush: "",
			Id: "",
			InstanceProperties: &computepb.InstanceProperties{
				AdvancedMachineFeatures: &computepb.AdvancedMachineFeatures{...}
				CanIpForward: "",
				ConfidentialInstanceConfig: &computepb.ConfidentialInstanceConfig{...}
				Description: "",
				Disks: &computepb.AttachedDisk{...}
				GuestAccelerators: &computepb.AcceleratorConfig{...}
				KeyRevocationActionType: "",
				Labels: "",
				MachineType: "",
				Metadata: &computepb.Metadata{...}
				MinCpuPlatform: "",
				NetworkInterfaces: &computepb.NetworkInterface{...}
				NetworkPerformanceConfig: &computepb.NetworkPerformanceConfig{...}
				PrivateIpv6GoogleAccess: "",
				ReservationAffinity: &computepb.ReservationAffinity{...}
				ResourceManagerTags: "",
				ResourcePolicies: "",
				Scheduling: &computepb.Scheduling{...}
				ServiceAccounts: &computepb.ServiceAccount{...}
				ShieldedInstanceConfig: &computepb.ShieldedInstanceConfig{...}
				Tags: &computepb.Tags{...}
			}
			Kind: "",
			MachineImageEncryptionKey: &computepb.CustomerEncryptionKey{
				KmsKeyName: "",
				KmsKeyServiceAccount: "",
				RawKey: "",
				RsaEncryptedKey: "",
				Sha256: "",
			}
			Name: "",
			SatisfiesPzi: "",
			SatisfiesPzs: "",
			SavedDisks: &computepb.SavedDisk{
				Architecture: "",
				Kind: "",
				SourceDisk: "",
				StorageBytes: "",
				StorageBytesStatus: "",
			}
			SelfLink: "",
			SourceDiskEncryptionKeys: &computepb.SourceDiskEncryptionKey{
				DiskEncryptionKey: &computepb.CustomerEncryptionKey{...}
				SourceDisk: "",
			}
			SourceInstance: "",
			SourceInstanceProperties: &computepb.SourceInstanceProperties{
				CanIpForward: "",
				DeletionProtection: "",
				Description: "",
				Disks: &computepb.SavedAttachedDisk{...}
				GuestAccelerators: &computepb.AcceleratorConfig{...}
				KeyRevocationActionType: "",
				Labels: "",
				MachineType: "",
				Metadata: &computepb.Metadata{...}
				MinCpuPlatform: "",
				NetworkInterfaces: &computepb.NetworkInterface{...}
				Scheduling: &computepb.Scheduling{...}
				ServiceAccounts: &computepb.ServiceAccount{...}
				Tags: &computepb.Tags{...}
			}
			Status: "",
			StorageLocations: "",
			TotalStorageBytes: "",
		}
		Project: "",
		RequestId: "",
		SourceInstance: "",
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END compute_v1_generated_MachineImages_Insert_sync]
