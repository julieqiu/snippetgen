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


package computepb

import (
	"context"
	"fmt"
	"math"
	"net/url"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var newFirewallPoliciesClientHook clientHook

// FirewallPoliciesCallOptions contains the retry settings for each method of FirewallPoliciesClient.
type FirewallPoliciesCallOptions struct {
	AddAssociation []gax.CallOption
	AddRule []gax.CallOption
	CloneRules []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	GetAssociation []gax.CallOption
	GetIamPolicy []gax.CallOption
	GetRule []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	ListAssociations []gax.CallOption
	Move []gax.CallOption
	Patch []gax.CallOption
	PatchRule []gax.CallOption
	RemoveAssociation []gax.CallOption
	RemoveRule []gax.CallOption
	SetIamPolicy []gax.CallOption
	TestIamPermissions []gax.CallOption
}

func defaultFirewallPoliciesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("compute.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("compute.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("compute.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultFirewallPoliciesCallOptions() *FirewallPoliciesCallOptions {
	return &FirewallPoliciesCallOptions{
		AddAssociation: []gax.CallOption{
		},
		AddRule: []gax.CallOption{
		},
		CloneRules: []gax.CallOption{
		},
		Delete: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		GetAssociation: []gax.CallOption{
		},
		GetIamPolicy: []gax.CallOption{
		},
		GetRule: []gax.CallOption{
		},
		Insert: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
		ListAssociations: []gax.CallOption{
		},
		Move: []gax.CallOption{
		},
		Patch: []gax.CallOption{
		},
		PatchRule: []gax.CallOption{
		},
		RemoveAssociation: []gax.CallOption{
		},
		RemoveRule: []gax.CallOption{
		},
		SetIamPolicy: []gax.CallOption{
		},
		TestIamPermissions: []gax.CallOption{
		},
	}
}

// internalFirewallPoliciesClient is an interface that defines the methods available from Google Compute Engine API.
type internalFirewallPoliciesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddAssociation(context.Context, *computepb.AddAssociationFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	AddRule(context.Context, *computepb.AddRuleFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	CloneRules(context.Context, *computepb.CloneRulesFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	Delete(context.Context, *computepb.DeleteFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetFirewallPolicyRequest, ...gax.CallOption) (*computepb.FirewallPolicy, error)
	GetAssociation(context.Context, *computepb.GetAssociationFirewallPolicyRequest, ...gax.CallOption) (*computepb.FirewallPolicyAssociation, error)
	GetIamPolicy(context.Context, *computepb.GetIamPolicyFirewallPolicyRequest, ...gax.CallOption) (*computepb.Policy, error)
	GetRule(context.Context, *computepb.GetRuleFirewallPolicyRequest, ...gax.CallOption) (*computepb.FirewallPolicyRule, error)
	Insert(context.Context, *computepb.InsertFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListFirewallPoliciesRequest, ...gax.CallOption) *FirewallPolicyIterator
	ListAssociations(context.Context, *computepb.ListAssociationsFirewallPolicyRequest, ...gax.CallOption) (*computepb.FirewallPoliciesListAssociationsResponse, error)
	Move(context.Context, *computepb.MoveFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	Patch(context.Context, *computepb.PatchFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	PatchRule(context.Context, *computepb.PatchRuleFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	RemoveAssociation(context.Context, *computepb.RemoveAssociationFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	RemoveRule(context.Context, *computepb.RemoveRuleFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetIamPolicy(context.Context, *computepb.SetIamPolicyFirewallPolicyRequest, ...gax.CallOption) (*computepb.Policy, error)
	TestIamPermissions(context.Context, *computepb.TestIamPermissionsFirewallPolicyRequest, ...gax.CallOption) (*computepb.TestPermissionsResponse, error)
}

// FirewallPoliciesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The FirewallPolicies API.
type FirewallPoliciesClient struct {
	// The internal transport-dependent client.
	internalClient internalFirewallPoliciesClient

	// The call options for this service.
	CallOptions *FirewallPoliciesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FirewallPoliciesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FirewallPoliciesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *FirewallPoliciesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddAssociation inserts an association for the specified firewall policy.
func (c *FirewallPoliciesClient) AddAssociation(ctx context.Context, req *computepb.AddAssociationFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddAssociation(ctx, req, opts...)
}

// AddRule inserts a rule into a firewall policy.
func (c *FirewallPoliciesClient) AddRule(ctx context.Context, req *computepb.AddRuleFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddRule(ctx, req, opts...)
}

// CloneRules copies rules to the specified firewall policy.
func (c *FirewallPoliciesClient) CloneRules(ctx context.Context, req *computepb.CloneRulesFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.CloneRules(ctx, req, opts...)
}

// Delete deletes the specified policy.
func (c *FirewallPoliciesClient) Delete(ctx context.Context, req *computepb.DeleteFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified firewall policy.
func (c *FirewallPoliciesClient) Get(ctx context.Context, req *computepb.GetFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPolicy, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetAssociation gets an association with the specified name.
func (c *FirewallPoliciesClient) GetAssociation(ctx context.Context, req *computepb.GetAssociationFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPolicyAssociation, error) {
	return c.internalClient.GetAssociation(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists.
func (c *FirewallPoliciesClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// GetRule gets a rule of the specified priority.
func (c *FirewallPoliciesClient) GetRule(ctx context.Context, req *computepb.GetRuleFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPolicyRule, error) {
	return c.internalClient.GetRule(ctx, req, opts...)
}

// Insert creates a new policy in the specified project using the data included in the request.
func (c *FirewallPoliciesClient) Insert(ctx context.Context, req *computepb.InsertFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List lists all the policies that have been configured for the specified folder or organization.
func (c *FirewallPoliciesClient) List(ctx context.Context, req *computepb.ListFirewallPoliciesRequest, opts ...gax.CallOption) *FirewallPolicyIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// ListAssociations lists associations of a specified target, i.e., organization or folder.
func (c *FirewallPoliciesClient) ListAssociations(ctx context.Context, req *computepb.ListAssociationsFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPoliciesListAssociationsResponse, error) {
	return c.internalClient.ListAssociations(ctx, req, opts...)
}

// Move moves the specified firewall policy.
func (c *FirewallPoliciesClient) Move(ctx context.Context, req *computepb.MoveFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Move(ctx, req, opts...)
}

// Patch patches the specified policy with the data included in the request.
func (c *FirewallPoliciesClient) Patch(ctx context.Context, req *computepb.PatchFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// PatchRule patches a rule of the specified priority.
func (c *FirewallPoliciesClient) PatchRule(ctx context.Context, req *computepb.PatchRuleFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.PatchRule(ctx, req, opts...)
}

// RemoveAssociation removes an association for the specified firewall policy.
func (c *FirewallPoliciesClient) RemoveAssociation(ctx context.Context, req *computepb.RemoveAssociationFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.RemoveAssociation(ctx, req, opts...)
}

// RemoveRule deletes a rule of the specified priority.
func (c *FirewallPoliciesClient) RemoveRule(ctx context.Context, req *computepb.RemoveRuleFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.RemoveRule(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy.
func (c *FirewallPoliciesClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *FirewallPoliciesClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// firewallPoliciesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type firewallPoliciesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing FirewallPoliciesClient
	CallOptions **FirewallPoliciesCallOptions

	// The gRPC API client.
	firewallPoliciesClient computepb.FirewallPoliciesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewFirewallPoliciesClient creates a new firewall policies client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The FirewallPolicies API.
func NewFirewallPoliciesClient(ctx context.Context, opts ...option.ClientOption) (*FirewallPoliciesClient, error) {
	clientOpts := defaultFirewallPoliciesGRPCClientOptions()
	if newFirewallPoliciesClientHook != nil {
		hookOpts, err := newFirewallPoliciesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := FirewallPoliciesClient{CallOptions: defaultFirewallPoliciesCallOptions()}

	c := &firewallPoliciesGRPCClient{
		connPool:    connPool,
		firewallPoliciesClient: computepb.NewFirewallPoliciesClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *firewallPoliciesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *firewallPoliciesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *firewallPoliciesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *firewallPoliciesGRPCClient) AddAssociation(ctx context.Context, req *computepb.AddAssociationFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AddAssociation[0:len((*c.CallOptions).AddAssociation):len((*c.CallOptions).AddAssociation)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.AddAssociation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) AddRule(ctx context.Context, req *computepb.AddRuleFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AddRule[0:len((*c.CallOptions).AddRule):len((*c.CallOptions).AddRule)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.AddRule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) CloneRules(ctx context.Context, req *computepb.CloneRulesFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CloneRules[0:len((*c.CallOptions).CloneRules):len((*c.CallOptions).CloneRules)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.CloneRules(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) Get(ctx context.Context, req *computepb.GetFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPolicy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.FirewallPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) GetAssociation(ctx context.Context, req *computepb.GetAssociationFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPolicyAssociation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetAssociation[0:len((*c.CallOptions).GetAssociation):len((*c.CallOptions).GetAssociation)], opts...)
	var resp *computepb.FirewallPolicyAssociation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.GetAssociation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *computepb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) GetRule(ctx context.Context, req *computepb.GetRuleFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPolicyRule, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetRule[0:len((*c.CallOptions).GetRule):len((*c.CallOptions).GetRule)], opts...)
	var resp *computepb.FirewallPolicyRule
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.GetRule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) Insert(ctx context.Context, req *computepb.InsertFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) List(ctx context.Context, req *computepb.ListFirewallPoliciesRequest, opts ...gax.CallOption) *FirewallPolicyIterator {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &FirewallPolicyIterator{}
	req = proto.Clone(req).(*computepb.ListFirewallPoliciesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.FirewallPolicy, string, error) {
		resp := &computepb.FirewallPolicyList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = math.MaxInt32
		} else if pageSize != 0 {
			req.MaxResults = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.firewallPoliciesClient.List(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *firewallPoliciesGRPCClient) ListAssociations(ctx context.Context, req *computepb.ListAssociationsFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPoliciesListAssociationsResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListAssociations[0:len((*c.CallOptions).ListAssociations):len((*c.CallOptions).ListAssociations)], opts...)
	var resp *computepb.FirewallPoliciesListAssociationsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.ListAssociations(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) Move(ctx context.Context, req *computepb.MoveFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Move[0:len((*c.CallOptions).Move):len((*c.CallOptions).Move)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.Move(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) Patch(ctx context.Context, req *computepb.PatchFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) PatchRule(ctx context.Context, req *computepb.PatchRuleFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).PatchRule[0:len((*c.CallOptions).PatchRule):len((*c.CallOptions).PatchRule)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.PatchRule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) RemoveAssociation(ctx context.Context, req *computepb.RemoveAssociationFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RemoveAssociation[0:len((*c.CallOptions).RemoveAssociation):len((*c.CallOptions).RemoveAssociation)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.RemoveAssociation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) RemoveRule(ctx context.Context, req *computepb.RemoveRuleFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RemoveRule[0:len((*c.CallOptions).RemoveRule):len((*c.CallOptions).RemoveRule)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.RemoveRule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *computepb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firewallPoliciesGRPCClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *computepb.TestPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firewallPoliciesClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
