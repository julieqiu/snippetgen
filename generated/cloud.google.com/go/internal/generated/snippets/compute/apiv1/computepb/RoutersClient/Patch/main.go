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

// [START compute_v1_generated_Routers_Patch_sync]

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
	c, err := computepb.NewRoutersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchRouterRequest.
	req := &computepb.PatchRouterRequest{
		Project: "",
		Region: "",
		RequestId: "",
		Router: "",
		RouterResource: &computepb.Router{
			Bgp: &computepb.RouterBgp{
				AdvertiseMode: "",
				AdvertisedGroups: "",
				AdvertisedIpRanges: &computepb.RouterAdvertisedIpRange{...}
				Asn: "",
				IdentifierRange: "",
				KeepaliveInterval: "",
			}
			BgpPeers: &computepb.RouterBgpPeer{
				AdvertiseMode: "",
				AdvertisedGroups: "",
				AdvertisedIpRanges: &computepb.RouterAdvertisedIpRange{...}
				AdvertisedRoutePriority: "",
				Bfd: &computepb.RouterBgpPeerBfd{...}
				CustomLearnedIpRanges: &computepb.RouterBgpPeerCustomLearnedIpRange{...}
				CustomLearnedRoutePriority: "",
				Enable: "",
				EnableIpv4: "",
				EnableIpv6: "",
				ExportPolicies: "",
				ImportPolicies: "",
				InterfaceName: "",
				IpAddress: "",
				Ipv4NexthopAddress: "",
				Ipv6NexthopAddress: "",
				ManagementType: "",
				Md5AuthenticationKeyName: "",
				Name: "",
				PeerAsn: "",
				PeerIpAddress: "",
				PeerIpv4NexthopAddress: "",
				PeerIpv6NexthopAddress: "",
				RouterApplianceInstance: "",
			}
			CreationTimestamp: "",
			Description: "",
			EncryptedInterconnectRouter: "",
			Id: "",
			Interfaces: &computepb.RouterInterface{
				IpRange: "",
				IpVersion: "",
				LinkedInterconnectAttachment: "",
				LinkedVpnTunnel: "",
				ManagementType: "",
				Name: "",
				PrivateIpAddress: "",
				RedundantInterface: "",
				Subnetwork: "",
			}
			Kind: "",
			Md5AuthenticationKeys: &computepb.RouterMd5AuthenticationKey{
				Key: "",
				Name: "",
			}
			Name: "",
			Nats: &computepb.RouterNat{
				AutoNetworkTier: "",
				DrainNatIps: "",
				EnableDynamicPortAllocation: "",
				EnableEndpointIndependentMapping: "",
				EndpointTypes: "",
				IcmpIdleTimeoutSec: "",
				LogConfig: &computepb.RouterNatLogConfig{...}
				MaxPortsPerVm: "",
				MinPortsPerVm: "",
				Name: "",
				NatIpAllocateOption: "",
				NatIps: "",
				Rules: &computepb.RouterNatRule{...}
				SourceSubnetworkIpRangesToNat: "",
				Subnetworks: &computepb.RouterNatSubnetworkToNat{...}
				TcpEstablishedIdleTimeoutSec: "",
				TcpTimeWaitTimeoutSec: "",
				TcpTransitoryIdleTimeoutSec: "",
				Type: "",
				UdpIdleTimeoutSec: "",
			}
			Network: "",
			Region: "",
			SelfLink: "",
		}
	}
	resp, err := c.Patch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END compute_v1_generated_Routers_Patch_sync]
