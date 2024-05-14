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

// [START compute_v1_generated_RegionBackendServices_Update_sync]

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
	c, err := computepb.NewRegionBackendServicesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#UpdateRegionBackendServiceRequest.
	req := &computepb.UpdateRegionBackendServiceRequest{
		BackendService: "",
		BackendServiceResource: &computepb.BackendService{
			AffinityCookieTtlSec: "",
			Backends: &computepb.Backend{
				BalancingMode: "",
				CapacityScaler: "",
				Description: "",
				Failover: "",
				Group: "",
				MaxConnections: "",
				MaxConnectionsPerEndpoint: "",
				MaxConnectionsPerInstance: "",
				MaxRate: "",
				MaxRatePerEndpoint: "",
				MaxRatePerInstance: "",
				MaxUtilization: "",
				Preference: "",
			}
			CdnPolicy: &computepb.BackendServiceCdnPolicy{
				BypassCacheOnRequestHeaders: &computepb.BackendServiceCdnPolicyBypassCacheOnRequestHeader{...}
				CacheKeyPolicy: &computepb.CacheKeyPolicy{...}
				CacheMode: "",
				ClientTtl: "",
				DefaultTtl: "",
				MaxTtl: "",
				NegativeCaching: "",
				NegativeCachingPolicy: &computepb.BackendServiceCdnPolicyNegativeCachingPolicy{...}
				RequestCoalescing: "",
				ServeWhileStale: "",
				SignedUrlCacheMaxAgeSec: "",
				SignedUrlKeyNames: "",
			}
			CircuitBreakers: &computepb.CircuitBreakers{
				MaxConnections: "",
				MaxPendingRequests: "",
				MaxRequests: "",
				MaxRequestsPerConnection: "",
				MaxRetries: "",
			}
			CompressionMode: "",
			ConnectionDraining: &computepb.ConnectionDraining{
				DrainingTimeoutSec: "",
			}
			ConnectionTrackingPolicy: &computepb.BackendServiceConnectionTrackingPolicy{
				ConnectionPersistenceOnUnhealthyBackends: "",
				EnableStrongAffinity: "",
				IdleTimeoutSec: "",
				TrackingMode: "",
			}
			ConsistentHash: &computepb.ConsistentHashLoadBalancerSettings{
				HttpCookie: &computepb.ConsistentHashLoadBalancerSettingsHttpCookie{...}
				HttpHeaderName: "",
				MinimumRingSize: "",
			}
			CreationTimestamp: "",
			CustomRequestHeaders: "",
			CustomResponseHeaders: "",
			Description: "",
			EdgeSecurityPolicy: "",
			EnableCDN: "",
			FailoverPolicy: &computepb.BackendServiceFailoverPolicy{
				DisableConnectionDrainOnFailover: "",
				DropTrafficIfUnhealthy: "",
				FailoverRatio: "",
			}
			Fingerprint: "",
			HealthChecks: "",
			Iap: &computepb.BackendServiceIAP{
				Enabled: "",
				Oauth2ClientId: "",
				Oauth2ClientSecret: "",
				Oauth2ClientSecretSha256: "",
			}
			Id: "",
			Kind: "",
			LoadBalancingScheme: "",
			LocalityLbPolicies: &computepb.BackendServiceLocalityLoadBalancingPolicyConfig{
				CustomPolicy: &computepb.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy{...}
				Policy: &computepb.BackendServiceLocalityLoadBalancingPolicyConfigPolicy{...}
			}
			LocalityLbPolicy: "",
			LogConfig: &computepb.BackendServiceLogConfig{
				Enable: "",
				OptionalFields: "",
				OptionalMode: "",
				SampleRate: "",
			}
			MaxStreamDuration: &computepb.Duration{
				Nanos: "",
				Seconds: "",
			}
			Metadatas: "",
			Name: "",
			Network: "",
			OutlierDetection: &computepb.OutlierDetection{
				BaseEjectionTime: &computepb.Duration{...}
				ConsecutiveErrors: "",
				ConsecutiveGatewayFailure: "",
				EnforcingConsecutiveErrors: "",
				EnforcingConsecutiveGatewayFailure: "",
				EnforcingSuccessRate: "",
				Interval: &computepb.Duration{...}
				MaxEjectionPercent: "",
				SuccessRateMinimumHosts: "",
				SuccessRateRequestVolume: "",
				SuccessRateStdevFactor: "",
			}
			Port: "",
			PortName: "",
			Protocol: "",
			Region: "",
			SecurityPolicy: "",
			SecuritySettings: &computepb.SecuritySettings{
				AwsV4Authentication: &computepb.AWSV4Signature{...}
				ClientTlsPolicy: "",
				SubjectAltNames: "",
			}
			SelfLink: "",
			ServiceBindings: "",
			ServiceLbPolicy: "",
			SessionAffinity: "",
			Subsetting: &computepb.Subsetting{
				Policy: "",
			}
			TimeoutSec: "",
			UsedBy: &computepb.BackendServiceUsedBy{
				Reference: "",
			}
		}
		Project: "",
		Region: "",
		RequestId: "",
	}
	resp, err := c.Update(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END compute_v1_generated_RegionBackendServices_Update_sync]
