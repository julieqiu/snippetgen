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

func ExampleNewFirewallPoliciesClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleFirewallPoliciesClient_AddAssociation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AddAssociationFirewallPolicyRequest.
	req := &computepb.AddAssociationFirewallPolicyRequest{
		FirewallPolicy: "",
		FirewallPolicyAssociationResource: &computepb.FirewallPolicyAssociation{
			AttachmentTarget: "",
			DisplayName: "",
			FirewallPolicyId: "",
			Name: "",
			ShortName: "",
		}
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

func ExampleFirewallPoliciesClient_AddRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AddRuleFirewallPolicyRequest.
	req := &computepb.AddRuleFirewallPolicyRequest{
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
		RequestId: "",
	}
	resp, err := c.AddRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallPoliciesClient_CloneRules() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#CloneRulesFirewallPolicyRequest.
	req := &computepb.CloneRulesFirewallPolicyRequest{
		FirewallPolicy: "",
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

func ExampleFirewallPoliciesClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteFirewallPolicyRequest.
	req := &computepb.DeleteFirewallPolicyRequest{
		FirewallPolicy: "",
		RequestId: "",
	}
	resp, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallPoliciesClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetFirewallPolicyRequest.
	req := &computepb.GetFirewallPolicyRequest{
		FirewallPolicy: "",
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallPoliciesClient_GetAssociation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetAssociationFirewallPolicyRequest.
	req := &computepb.GetAssociationFirewallPolicyRequest{
		FirewallPolicy: "",
		Name: "",
	}
	resp, err := c.GetAssociation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallPoliciesClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetIamPolicyFirewallPolicyRequest.
	req := &computepb.GetIamPolicyFirewallPolicyRequest{
		OptionsRequestedPolicyVersion: "",
		Resource: "",
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallPoliciesClient_GetRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetRuleFirewallPolicyRequest.
	req := &computepb.GetRuleFirewallPolicyRequest{
		FirewallPolicy: "",
		Priority: "",
	}
	resp, err := c.GetRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallPoliciesClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertFirewallPolicyRequest.
	req := &computepb.InsertFirewallPolicyRequest{
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
		ParentId: "",
		RequestId: "",
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallPoliciesClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListFirewallPoliciesRequest.
	req := &computepb.ListFirewallPoliciesRequest{
		Filter: "",
		MaxResults: "",
		OrderBy: "",
		PageToken: "",
		ParentId: "",
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

func ExampleFirewallPoliciesClient_ListAssociations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListAssociationsFirewallPolicyRequest.
	req := &computepb.ListAssociationsFirewallPolicyRequest{
		TargetResource: "",
	}
	resp, err := c.ListAssociations(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallPoliciesClient_Move() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#MoveFirewallPolicyRequest.
	req := &computepb.MoveFirewallPolicyRequest{
		FirewallPolicy: "",
		ParentId: "",
		RequestId: "",
	}
	resp, err := c.Move(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallPoliciesClient_Patch() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchFirewallPolicyRequest.
	req := &computepb.PatchFirewallPolicyRequest{
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
		RequestId: "",
	}
	resp, err := c.Patch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallPoliciesClient_PatchRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchRuleFirewallPolicyRequest.
	req := &computepb.PatchRuleFirewallPolicyRequest{
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
		RequestId: "",
	}
	resp, err := c.PatchRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallPoliciesClient_RemoveAssociation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#RemoveAssociationFirewallPolicyRequest.
	req := &computepb.RemoveAssociationFirewallPolicyRequest{
		FirewallPolicy: "",
		Name: "",
		RequestId: "",
	}
	resp, err := c.RemoveAssociation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallPoliciesClient_RemoveRule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#RemoveRuleFirewallPolicyRequest.
	req := &computepb.RemoveRuleFirewallPolicyRequest{
		FirewallPolicy: "",
		Priority: "",
		RequestId: "",
	}
	resp, err := c.RemoveRule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirewallPoliciesClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#SetIamPolicyFirewallPolicyRequest.
	req := &computepb.SetIamPolicyFirewallPolicyRequest{
		GlobalOrganizationSetPolicyRequestResource: &computepb.GlobalOrganizationSetPolicyRequest{
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

func ExampleFirewallPoliciesClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewFirewallPoliciesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#TestIamPermissionsFirewallPolicyRequest.
	req := &computepb.TestIamPermissionsFirewallPolicyRequest{
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
