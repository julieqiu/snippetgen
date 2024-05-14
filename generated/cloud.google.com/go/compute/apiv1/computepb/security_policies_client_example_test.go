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

func ExampleNewSecurityPoliciesClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSecurityPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleSecurityPoliciesClient_AddRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSecurityPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AddRuleSecurityPolicyRequest.
	req := &computepb.AddRuleSecurityPolicyRequest{
		Project: "",
		SecurityPolicy: "",
		SecurityPolicyRuleResource: &computepb.SecurityPolicyRule{
			Action: "",
			Description: "",
			HeaderAction: &computepb.SecurityPolicyRuleHttpHeaderAction{
				RequestHeadersToAdds: &computepb.SecurityPolicyRuleHttpHeaderActionHttpHeaderOption{...}
			}
			Kind: "",
			Match: &computepb.SecurityPolicyRuleMatcher{
				Config: &computepb.SecurityPolicyRuleMatcherConfig{...}
				Expr: &computepb.Expr{...}
				ExprOptions: &computepb.SecurityPolicyRuleMatcherExprOptions{...}
				VersionedExpr: "",
			}
			NetworkMatch: &computepb.SecurityPolicyRuleNetworkMatcher{
				DestIpRanges: "",
				DestPorts: "",
				IpProtocols: "",
				SrcAsns: "",
				SrcIpRanges: "",
				SrcPorts: "",
				SrcRegionCodes: "",
				UserDefinedFields: &computepb.SecurityPolicyRuleNetworkMatcherUserDefinedFieldMatch{...}
			}
			PreconfiguredWafConfig: &computepb.SecurityPolicyRulePreconfiguredWafConfig{
				Exclusions: &computepb.SecurityPolicyRulePreconfiguredWafConfigExclusion{...}
			}
			Preview: "",
			Priority: "",
			RateLimitOptions: &computepb.SecurityPolicyRuleRateLimitOptions{
				BanDurationSec: "",
				BanThreshold: &computepb.SecurityPolicyRuleRateLimitOptionsThreshold{...}
				ConformAction: "",
				EnforceOnKey: "",
				EnforceOnKeyConfigs: &computepb.SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig{...}
				EnforceOnKeyName: "",
				ExceedAction: "",
				ExceedRedirectOptions: &computepb.SecurityPolicyRuleRedirectOptions{...}
				RateLimitThreshold: &computepb.SecurityPolicyRuleRateLimitOptionsThreshold{...}
			}
			RedirectOptions: &computepb.SecurityPolicyRuleRedirectOptions{
				Target: "",
				Type: "",
			}
		}
		ValidateOnly: "",
	}
	resp, err := c.AddRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecurityPoliciesClient_AggregatedList() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSecurityPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AggregatedListSecurityPoliciesRequest.
	req := &computepb.AggregatedListSecurityPoliciesRequest{
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
		_ = it.Response.(*computepb.SecurityPoliciesAggregatedList)
	}
}

func ExampleSecurityPoliciesClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSecurityPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteSecurityPolicyRequest.
	req := &computepb.DeleteSecurityPolicyRequest{
		Project: "",
		RequestId: "",
		SecurityPolicy: "",
	}
	resp, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecurityPoliciesClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSecurityPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetSecurityPolicyRequest.
	req := &computepb.GetSecurityPolicyRequest{
		Project: "",
		SecurityPolicy: "",
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecurityPoliciesClient_GetRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSecurityPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetRuleSecurityPolicyRequest.
	req := &computepb.GetRuleSecurityPolicyRequest{
		Priority: "",
		Project: "",
		SecurityPolicy: "",
	}
	resp, err := c.GetRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecurityPoliciesClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSecurityPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertSecurityPolicyRequest.
	req := &computepb.InsertSecurityPolicyRequest{
		Project: "",
		RequestId: "",
		SecurityPolicyResource: &computepb.SecurityPolicy{
			AdaptiveProtectionConfig: &computepb.SecurityPolicyAdaptiveProtectionConfig{
				Layer7DdosDefenseConfig: &computepb.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig{...}
			}
			AdvancedOptionsConfig: &computepb.SecurityPolicyAdvancedOptionsConfig{
				JsonCustomConfig: &computepb.SecurityPolicyAdvancedOptionsConfigJsonCustomConfig{...}
				JsonParsing: "",
				LogLevel: "",
				UserIpRequestHeaders: "",
			}
			CreationTimestamp: "",
			DdosProtectionConfig: &computepb.SecurityPolicyDdosProtectionConfig{
				DdosProtection: "",
			}
			Description: "",
			Fingerprint: "",
			Id: "",
			Kind: "",
			LabelFingerprint: "",
			Labels: "",
			Name: "",
			RecaptchaOptionsConfig: &computepb.SecurityPolicyRecaptchaOptionsConfig{
				RedirectSiteKey: "",
			}
			Region: "",
			Rules: &computepb.SecurityPolicyRule{
				Action: "",
				Description: "",
				HeaderAction: &computepb.SecurityPolicyRuleHttpHeaderAction{...}
				Kind: "",
				Match: &computepb.SecurityPolicyRuleMatcher{...}
				NetworkMatch: &computepb.SecurityPolicyRuleNetworkMatcher{...}
				PreconfiguredWafConfig: &computepb.SecurityPolicyRulePreconfiguredWafConfig{...}
				Preview: "",
				Priority: "",
				RateLimitOptions: &computepb.SecurityPolicyRuleRateLimitOptions{...}
				RedirectOptions: &computepb.SecurityPolicyRuleRedirectOptions{...}
			}
			SelfLink: "",
			Type: "",
			UserDefinedFields: &computepb.SecurityPolicyUserDefinedField{
				Base: "",
				Mask: "",
				Name: "",
				Offset: "",
				Size: "",
			}
		}
		ValidateOnly: "",
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecurityPoliciesClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSecurityPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListSecurityPoliciesRequest.
	req := &computepb.ListSecurityPoliciesRequest{
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
		_ = it.Response.(*computepb.SecurityPolicyList)
	}
}

func ExampleSecurityPoliciesClient_ListPreconfiguredExpressionSets() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSecurityPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListPreconfiguredExpressionSetsSecurityPoliciesRequest.
	req := &computepb.ListPreconfiguredExpressionSetsSecurityPoliciesRequest{
		Filter: "",
		MaxResults: "",
		OrderBy: "",
		PageToken: "",
		Project: "",
		ReturnPartialSuccess: "",
	}
	resp, err := c.ListPreconfiguredExpressionSets(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecurityPoliciesClient_Patch() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSecurityPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchSecurityPolicyRequest.
	req := &computepb.PatchSecurityPolicyRequest{
		Project: "",
		RequestId: "",
		SecurityPolicy: "",
		SecurityPolicyResource: &computepb.SecurityPolicy{
			AdaptiveProtectionConfig: &computepb.SecurityPolicyAdaptiveProtectionConfig{
				Layer7DdosDefenseConfig: &computepb.SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig{...}
			}
			AdvancedOptionsConfig: &computepb.SecurityPolicyAdvancedOptionsConfig{
				JsonCustomConfig: &computepb.SecurityPolicyAdvancedOptionsConfigJsonCustomConfig{...}
				JsonParsing: "",
				LogLevel: "",
				UserIpRequestHeaders: "",
			}
			CreationTimestamp: "",
			DdosProtectionConfig: &computepb.SecurityPolicyDdosProtectionConfig{
				DdosProtection: "",
			}
			Description: "",
			Fingerprint: "",
			Id: "",
			Kind: "",
			LabelFingerprint: "",
			Labels: "",
			Name: "",
			RecaptchaOptionsConfig: &computepb.SecurityPolicyRecaptchaOptionsConfig{
				RedirectSiteKey: "",
			}
			Region: "",
			Rules: &computepb.SecurityPolicyRule{
				Action: "",
				Description: "",
				HeaderAction: &computepb.SecurityPolicyRuleHttpHeaderAction{...}
				Kind: "",
				Match: &computepb.SecurityPolicyRuleMatcher{...}
				NetworkMatch: &computepb.SecurityPolicyRuleNetworkMatcher{...}
				PreconfiguredWafConfig: &computepb.SecurityPolicyRulePreconfiguredWafConfig{...}
				Preview: "",
				Priority: "",
				RateLimitOptions: &computepb.SecurityPolicyRuleRateLimitOptions{...}
				RedirectOptions: &computepb.SecurityPolicyRuleRedirectOptions{...}
			}
			SelfLink: "",
			Type: "",
			UserDefinedFields: &computepb.SecurityPolicyUserDefinedField{
				Base: "",
				Mask: "",
				Name: "",
				Offset: "",
				Size: "",
			}
		}
		UpdateMask: "",
	}
	resp, err := c.Patch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecurityPoliciesClient_PatchRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSecurityPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchRuleSecurityPolicyRequest.
	req := &computepb.PatchRuleSecurityPolicyRequest{
		Priority: "",
		Project: "",
		SecurityPolicy: "",
		SecurityPolicyRuleResource: &computepb.SecurityPolicyRule{
			Action: "",
			Description: "",
			HeaderAction: &computepb.SecurityPolicyRuleHttpHeaderAction{
				RequestHeadersToAdds: &computepb.SecurityPolicyRuleHttpHeaderActionHttpHeaderOption{...}
			}
			Kind: "",
			Match: &computepb.SecurityPolicyRuleMatcher{
				Config: &computepb.SecurityPolicyRuleMatcherConfig{...}
				Expr: &computepb.Expr{...}
				ExprOptions: &computepb.SecurityPolicyRuleMatcherExprOptions{...}
				VersionedExpr: "",
			}
			NetworkMatch: &computepb.SecurityPolicyRuleNetworkMatcher{
				DestIpRanges: "",
				DestPorts: "",
				IpProtocols: "",
				SrcAsns: "",
				SrcIpRanges: "",
				SrcPorts: "",
				SrcRegionCodes: "",
				UserDefinedFields: &computepb.SecurityPolicyRuleNetworkMatcherUserDefinedFieldMatch{...}
			}
			PreconfiguredWafConfig: &computepb.SecurityPolicyRulePreconfiguredWafConfig{
				Exclusions: &computepb.SecurityPolicyRulePreconfiguredWafConfigExclusion{...}
			}
			Preview: "",
			Priority: "",
			RateLimitOptions: &computepb.SecurityPolicyRuleRateLimitOptions{
				BanDurationSec: "",
				BanThreshold: &computepb.SecurityPolicyRuleRateLimitOptionsThreshold{...}
				ConformAction: "",
				EnforceOnKey: "",
				EnforceOnKeyConfigs: &computepb.SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig{...}
				EnforceOnKeyName: "",
				ExceedAction: "",
				ExceedRedirectOptions: &computepb.SecurityPolicyRuleRedirectOptions{...}
				RateLimitThreshold: &computepb.SecurityPolicyRuleRateLimitOptionsThreshold{...}
			}
			RedirectOptions: &computepb.SecurityPolicyRuleRedirectOptions{
				Target: "",
				Type: "",
			}
		}
		UpdateMask: "",
		ValidateOnly: "",
	}
	resp, err := c.PatchRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecurityPoliciesClient_RemoveRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSecurityPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#RemoveRuleSecurityPolicyRequest.
	req := &computepb.RemoveRuleSecurityPolicyRequest{
		Priority: "",
		Project: "",
		SecurityPolicy: "",
	}
	resp, err := c.RemoveRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSecurityPoliciesClient_SetLabels() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewSecurityPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetLabelsSecurityPolicyRequest.
	req := &computepb.SetLabelsSecurityPolicyRequest{
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
