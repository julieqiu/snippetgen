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

func ExampleNewNetworkFirewallPoliciesClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNetworkFirewallPoliciesClient_AddAssociation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AddAssociationNetworkFirewallPolicyRequest.
	req := &computepb.AddAssociationNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		FirewallPolicyAssociationResource: &computepb.FirewallPolicyAssociation{
			AttachmentTarget: "",
			DisplayName: "",
			FirewallPolicyId: "",
			Name: "",
			ShortName: "",
		}
		Project: "",
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

func ExampleNetworkFirewallPoliciesClient_AddRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AddRuleNetworkFirewallPolicyRequest.
	req := &computepb.AddRuleNetworkFirewallPolicyRequest{
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
		RequestId: "",
	}
	resp, err := c.AddRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkFirewallPoliciesClient_CloneRules() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#CloneRulesNetworkFirewallPolicyRequest.
	req := &computepb.CloneRulesNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		Project: "",
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

func ExampleNetworkFirewallPoliciesClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteNetworkFirewallPolicyRequest.
	req := &computepb.DeleteNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
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

func ExampleNetworkFirewallPoliciesClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetNetworkFirewallPolicyRequest.
	req := &computepb.GetNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		Project: "",
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkFirewallPoliciesClient_GetAssociation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetAssociationNetworkFirewallPolicyRequest.
	req := &computepb.GetAssociationNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		Name: "",
		Project: "",
	}
	resp, err := c.GetAssociation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkFirewallPoliciesClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetIamPolicyNetworkFirewallPolicyRequest.
	req := &computepb.GetIamPolicyNetworkFirewallPolicyRequest{
		OptionsRequestedPolicyVersion: "",
		Project: "",
		Resource: "",
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkFirewallPoliciesClient_GetRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetRuleNetworkFirewallPolicyRequest.
	req := &computepb.GetRuleNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		Priority: "",
		Project: "",
	}
	resp, err := c.GetRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkFirewallPoliciesClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertNetworkFirewallPolicyRequest.
	req := &computepb.InsertNetworkFirewallPolicyRequest{
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
		RequestId: "",
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkFirewallPoliciesClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListNetworkFirewallPoliciesRequest.
	req := &computepb.ListNetworkFirewallPoliciesRequest{
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
		_ = it.Response.(*computepb.FirewallPolicyList)
	}
}

func ExampleNetworkFirewallPoliciesClient_Patch() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchNetworkFirewallPolicyRequest.
	req := &computepb.PatchNetworkFirewallPolicyRequest{
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
		RequestId: "",
	}
	resp, err := c.Patch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkFirewallPoliciesClient_PatchRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchRuleNetworkFirewallPolicyRequest.
	req := &computepb.PatchRuleNetworkFirewallPolicyRequest{
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
		RequestId: "",
	}
	resp, err := c.PatchRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkFirewallPoliciesClient_RemoveAssociation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#RemoveAssociationNetworkFirewallPolicyRequest.
	req := &computepb.RemoveAssociationNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		Name: "",
		Project: "",
		RequestId: "",
	}
	resp, err := c.RemoveAssociation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkFirewallPoliciesClient_RemoveRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#RemoveRuleNetworkFirewallPolicyRequest.
	req := &computepb.RemoveRuleNetworkFirewallPolicyRequest{
		FirewallPolicy: "",
		Priority: "",
		Project: "",
		RequestId: "",
	}
	resp, err := c.RemoveRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkFirewallPoliciesClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetIamPolicyNetworkFirewallPolicyRequest.
	req := &computepb.SetIamPolicyNetworkFirewallPolicyRequest{
		GlobalSetPolicyRequestResource: &computepb.GlobalSetPolicyRequest{
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
		Project: "",
		Resource: "",
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNetworkFirewallPoliciesClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewNetworkFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#TestIamPermissionsNetworkFirewallPolicyRequest.
	req := &computepb.TestIamPermissionsNetworkFirewallPolicyRequest{
		Project: "",
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
