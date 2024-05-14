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

func ExampleNewRegionNetworkFirewallPoliciesClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleRegionNetworkFirewallPoliciesClient_AddAssociation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AddAssociationRegionNetworkFirewallPolicyRequest.
	req := &computepb.AddAssociationRegionNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		FirewallPolicyAssociationResource: &computepb.FirewallPolicyAssociation{
			AttachmentTarget: "",
			DisplayName: "",
			FirewallPolicyId: "",
			Name: "",
			ShortName: "",
		}
		Project: "",
		Region: "",
		ReplaceExistingAssociation: "",
		RequestId: "",
	}
	resp, err := c.AddAssociation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionNetworkFirewallPoliciesClient_AddRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AddRuleRegionNetworkFirewallPolicyRequest.
	req := &computepb.AddRuleRegionNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		FirewallPolicyRuleResource: &computepb.FirewallPolicyRule{
			Action: "",
			Description: "",
			Direction: "",
			Disabled: "",
			EnableLogging: "",
			Kind: "",
			Match: &computepb.FirewallPolicyRuleMatcher{
				DestAddressGroups: "",
				DestFqdns: "",
				DestIpRanges: "",
				DestRegionCodes: "",
				DestThreatIntelligences: "",
				Layer4Configs: &computepb.FirewallPolicyRuleMatcherLayer4Config{...}
				SrcAddressGroups: "",
				SrcFqdns: "",
				SrcIpRanges: "",
				SrcRegionCodes: "",
				SrcSecureTags: &computepb.FirewallPolicyRuleSecureTag{...}
				SrcThreatIntelligences: "",
			}
			Priority: "",
			RuleName: "",
			RuleTupleCount: "",
			SecurityProfileGroup: "",
			TargetResources: "",
			TargetSecureTags: &computepb.FirewallPolicyRuleSecureTag{
				Name: "",
				State: "",
			}
			TargetServiceAccounts: "",
			TlsInspect: "",
		}
		MaxPriority: "",
		MinPriority: "",
		Project: "",
		Region: "",
		RequestId: "",
	}
	resp, err := c.AddRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionNetworkFirewallPoliciesClient_CloneRules() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#CloneRulesRegionNetworkFirewallPolicyRequest.
	req := &computepb.CloneRulesRegionNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		Project: "",
		Region: "",
		RequestId: "",
		SourceFirewallPolicy: "",
	}
	resp, err := c.CloneRules(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionNetworkFirewallPoliciesClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteRegionNetworkFirewallPolicyRequest.
	req := &computepb.DeleteRegionNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
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

func ExampleRegionNetworkFirewallPoliciesClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetRegionNetworkFirewallPolicyRequest.
	req := &computepb.GetRegionNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
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

func ExampleRegionNetworkFirewallPoliciesClient_GetAssociation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetAssociationRegionNetworkFirewallPolicyRequest.
	req := &computepb.GetAssociationRegionNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		Name: "",
		Project: "",
		Region: "",
	}
	resp, err := c.GetAssociation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionNetworkFirewallPoliciesClient_GetEffectiveFirewalls() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetEffectiveFirewallsRegionNetworkFirewallPolicyRequest.
	req := &computepb.GetEffectiveFirewallsRegionNetworkFirewallPolicyRequest{
		Network: "",
		Project: "",
		Region: "",
	}
	resp, err := c.GetEffectiveFirewalls(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionNetworkFirewallPoliciesClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetIamPolicyRegionNetworkFirewallPolicyRequest.
	req := &computepb.GetIamPolicyRegionNetworkFirewallPolicyRequest{
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

func ExampleRegionNetworkFirewallPoliciesClient_GetRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetRuleRegionNetworkFirewallPolicyRequest.
	req := &computepb.GetRuleRegionNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		Priority: "",
		Project: "",
		Region: "",
	}
	resp, err := c.GetRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionNetworkFirewallPoliciesClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertRegionNetworkFirewallPolicyRequest.
	req := &computepb.InsertRegionNetworkFirewallPolicyRequest{
		FirewallPolicyResource: &computepb.FirewallPolicy{
			Associations: &computepb.FirewallPolicyAssociation{
				AttachmentTarget: "",
				DisplayName: "",
				FirewallPolicyId: "",
				Name: "",
				ShortName: "",
			}
			CreationTimestamp: "",
			Description: "",
			DisplayName: "",
			Fingerprint: "",
			Id: "",
			Kind: "",
			Name: "",
			Parent: "",
			Region: "",
			RuleTupleCount: "",
			Rules: &computepb.FirewallPolicyRule{
				Action: "",
				Description: "",
				Direction: "",
				Disabled: "",
				EnableLogging: "",
				Kind: "",
				Match: &computepb.FirewallPolicyRuleMatcher{...}
				Priority: "",
				RuleName: "",
				RuleTupleCount: "",
				SecurityProfileGroup: "",
				TargetResources: "",
				TargetSecureTags: &computepb.FirewallPolicyRuleSecureTag{...}
				TargetServiceAccounts: "",
				TlsInspect: "",
			}
			SelfLink: "",
			SelfLinkWithId: "",
			ShortName: "",
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

func ExampleRegionNetworkFirewallPoliciesClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListRegionNetworkFirewallPoliciesRequest.
	req := &computepb.ListRegionNetworkFirewallPoliciesRequest{
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
		_ = it.Response.(*computepb.FirewallPolicyList)
	}
}

func ExampleRegionNetworkFirewallPoliciesClient_Patch() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchRegionNetworkFirewallPolicyRequest.
	req := &computepb.PatchRegionNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		FirewallPolicyResource: &computepb.FirewallPolicy{
			Associations: &computepb.FirewallPolicyAssociation{
				AttachmentTarget: "",
				DisplayName: "",
				FirewallPolicyId: "",
				Name: "",
				ShortName: "",
			}
			CreationTimestamp: "",
			Description: "",
			DisplayName: "",
			Fingerprint: "",
			Id: "",
			Kind: "",
			Name: "",
			Parent: "",
			Region: "",
			RuleTupleCount: "",
			Rules: &computepb.FirewallPolicyRule{
				Action: "",
				Description: "",
				Direction: "",
				Disabled: "",
				EnableLogging: "",
				Kind: "",
				Match: &computepb.FirewallPolicyRuleMatcher{...}
				Priority: "",
				RuleName: "",
				RuleTupleCount: "",
				SecurityProfileGroup: "",
				TargetResources: "",
				TargetSecureTags: &computepb.FirewallPolicyRuleSecureTag{...}
				TargetServiceAccounts: "",
				TlsInspect: "",
			}
			SelfLink: "",
			SelfLinkWithId: "",
			ShortName: "",
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

func ExampleRegionNetworkFirewallPoliciesClient_PatchRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchRuleRegionNetworkFirewallPolicyRequest.
	req := &computepb.PatchRuleRegionNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		FirewallPolicyRuleResource: &computepb.FirewallPolicyRule{
			Action: "",
			Description: "",
			Direction: "",
			Disabled: "",
			EnableLogging: "",
			Kind: "",
			Match: &computepb.FirewallPolicyRuleMatcher{
				DestAddressGroups: "",
				DestFqdns: "",
				DestIpRanges: "",
				DestRegionCodes: "",
				DestThreatIntelligences: "",
				Layer4Configs: &computepb.FirewallPolicyRuleMatcherLayer4Config{...}
				SrcAddressGroups: "",
				SrcFqdns: "",
				SrcIpRanges: "",
				SrcRegionCodes: "",
				SrcSecureTags: &computepb.FirewallPolicyRuleSecureTag{...}
				SrcThreatIntelligences: "",
			}
			Priority: "",
			RuleName: "",
			RuleTupleCount: "",
			SecurityProfileGroup: "",
			TargetResources: "",
			TargetSecureTags: &computepb.FirewallPolicyRuleSecureTag{
				Name: "",
				State: "",
			}
			TargetServiceAccounts: "",
			TlsInspect: "",
		}
		Priority: "",
		Project: "",
		Region: "",
		RequestId: "",
	}
	resp, err := c.PatchRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionNetworkFirewallPoliciesClient_RemoveAssociation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#RemoveAssociationRegionNetworkFirewallPolicyRequest.
	req := &computepb.RemoveAssociationRegionNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		Name: "",
		Project: "",
		Region: "",
		RequestId: "",
	}
	resp, err := c.RemoveAssociation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionNetworkFirewallPoliciesClient_RemoveRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#RemoveRuleRegionNetworkFirewallPolicyRequest.
	req := &computepb.RemoveRuleRegionNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		Priority: "",
		Project: "",
		Region: "",
		RequestId: "",
	}
	resp, err := c.RemoveRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionNetworkFirewallPoliciesClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetIamPolicyRegionNetworkFirewallPolicyRequest.
	req := &computepb.SetIamPolicyRegionNetworkFirewallPolicyRequest{
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

func ExampleRegionNetworkFirewallPoliciesClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewRegionNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#TestIamPermissionsRegionNetworkFirewallPolicyRequest.
	req := &computepb.TestIamPermissionsRegionNetworkFirewallPolicyRequest{
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