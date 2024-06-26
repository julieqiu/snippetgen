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

var newNetworkFirewallPoliciesClientHook clientHook

// NetworkFirewallPoliciesCallOptions contains the retry settings for each method of NetworkFirewallPoliciesClient.
type NetworkFirewallPoliciesCallOptions struct {
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
	Patch []gax.CallOption
	PatchRule []gax.CallOption
	RemoveAssociation []gax.CallOption
	RemoveRule []gax.CallOption
	SetIamPolicy []gax.CallOption
	TestIamPermissions []gax.CallOption
}

func defaultNetworkFirewallPoliciesGRPCClientOptions() []option.ClientOption {
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

func defaultNetworkFirewallPoliciesCallOptions() *NetworkFirewallPoliciesCallOptions {
	return &NetworkFirewallPoliciesCallOptions{
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

// internalNetworkFirewallPoliciesClient is an interface that defines the methods available from Google Compute Engine API.
type internalNetworkFirewallPoliciesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddAssociation(context.Context, *computepb.AddAssociationNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	AddRule(context.Context, *computepb.AddRuleNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	CloneRules(context.Context, *computepb.CloneRulesNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	Delete(context.Context, *computepb.DeleteNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.FirewallPolicy, error)
	GetAssociation(context.Context, *computepb.GetAssociationNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.FirewallPolicyAssociation, error)
	GetIamPolicy(context.Context, *computepb.GetIamPolicyNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.Policy, error)
	GetRule(context.Context, *computepb.GetRuleNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.FirewallPolicyRule, error)
	Insert(context.Context, *computepb.InsertNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListNetworkFirewallPoliciesRequest, ...gax.CallOption) *FirewallPolicyIterator
	Patch(context.Context, *computepb.PatchNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	PatchRule(context.Context, *computepb.PatchRuleNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	RemoveAssociation(context.Context, *computepb.RemoveAssociationNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	RemoveRule(context.Context, *computepb.RemoveRuleNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetIamPolicy(context.Context, *computepb.SetIamPolicyNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.Policy, error)
	TestIamPermissions(context.Context, *computepb.TestIamPermissionsNetworkFirewallPolicyRequest, ...gax.CallOption) (*computepb.TestPermissionsResponse, error)
}

// NetworkFirewallPoliciesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The NetworkFirewallPolicies API.
type NetworkFirewallPoliciesClient struct {
	// The internal transport-dependent client.
	internalClient internalNetworkFirewallPoliciesClient

	// The call options for this service.
	CallOptions *NetworkFirewallPoliciesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *NetworkFirewallPoliciesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *NetworkFirewallPoliciesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *NetworkFirewallPoliciesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddAssociation inserts an association for the specified firewall policy.
func (c *NetworkFirewallPoliciesClient) AddAssociation(ctx context.Context, req *computepb.AddAssociationNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddAssociation(ctx, req, opts...)
}

// AddRule inserts a rule into a firewall policy.
func (c *NetworkFirewallPoliciesClient) AddRule(ctx context.Context, req *computepb.AddRuleNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddRule(ctx, req, opts...)
}

// CloneRules copies rules to the specified firewall policy.
func (c *NetworkFirewallPoliciesClient) CloneRules(ctx context.Context, req *computepb.CloneRulesNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.CloneRules(ctx, req, opts...)
}

// Delete deletes the specified policy.
func (c *NetworkFirewallPoliciesClient) Delete(ctx context.Context, req *computepb.DeleteNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified network firewall policy.
func (c *NetworkFirewallPoliciesClient) Get(ctx context.Context, req *computepb.GetNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPolicy, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetAssociation gets an association with the specified name.
func (c *NetworkFirewallPoliciesClient) GetAssociation(ctx context.Context, req *computepb.GetAssociationNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPolicyAssociation, error) {
	return c.internalClient.GetAssociation(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists.
func (c *NetworkFirewallPoliciesClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// GetRule gets a rule of the specified priority.
func (c *NetworkFirewallPoliciesClient) GetRule(ctx context.Context, req *computepb.GetRuleNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPolicyRule, error) {
	return c.internalClient.GetRule(ctx, req, opts...)
}

// Insert creates a new policy in the specified project using the data included in the request.
func (c *NetworkFirewallPoliciesClient) Insert(ctx context.Context, req *computepb.InsertNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List lists all the policies that have been configured for the specified project.
func (c *NetworkFirewallPoliciesClient) List(ctx context.Context, req *computepb.ListNetworkFirewallPoliciesRequest, opts ...gax.CallOption) *FirewallPolicyIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified policy with the data included in the request.
func (c *NetworkFirewallPoliciesClient) Patch(ctx context.Context, req *computepb.PatchNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// PatchRule patches a rule of the specified priority.
func (c *NetworkFirewallPoliciesClient) PatchRule(ctx context.Context, req *computepb.PatchRuleNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.PatchRule(ctx, req, opts...)
}

// RemoveAssociation removes an association for the specified firewall policy.
func (c *NetworkFirewallPoliciesClient) RemoveAssociation(ctx context.Context, req *computepb.RemoveAssociationNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.RemoveAssociation(ctx, req, opts...)
}

// RemoveRule deletes a rule of the specified priority.
func (c *NetworkFirewallPoliciesClient) RemoveRule(ctx context.Context, req *computepb.RemoveRuleNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.RemoveRule(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy.
func (c *NetworkFirewallPoliciesClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *NetworkFirewallPoliciesClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// networkFirewallPoliciesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type networkFirewallPoliciesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing NetworkFirewallPoliciesClient
	CallOptions **NetworkFirewallPoliciesCallOptions

	// The gRPC API client.
	networkFirewallPoliciesClient computepb.NetworkFirewallPoliciesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewNetworkFirewallPoliciesClient creates a new network firewall policies client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The NetworkFirewallPolicies API.
func NewNetworkFirewallPoliciesClient(ctx context.Context, opts ...option.ClientOption) (*NetworkFirewallPoliciesClient, error) {
	clientOpts := defaultNetworkFirewallPoliciesGRPCClientOptions()
	if newNetworkFirewallPoliciesClientHook != nil {
		hookOpts, err := newNetworkFirewallPoliciesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := NetworkFirewallPoliciesClient{CallOptions: defaultNetworkFirewallPoliciesCallOptions()}

	c := &networkFirewallPoliciesGRPCClient{
		connPool:    connPool,
		networkFirewallPoliciesClient: computepb.NewNetworkFirewallPoliciesClient(connPool),
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
func (c *networkFirewallPoliciesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *networkFirewallPoliciesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *networkFirewallPoliciesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *networkFirewallPoliciesGRPCClient) AddAssociation(ctx context.Context, req *computepb.AddAssociationNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AddAssociation[0:len((*c.CallOptions).AddAssociation):len((*c.CallOptions).AddAssociation)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.AddAssociation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) AddRule(ctx context.Context, req *computepb.AddRuleNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AddRule[0:len((*c.CallOptions).AddRule):len((*c.CallOptions).AddRule)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.AddRule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) CloneRules(ctx context.Context, req *computepb.CloneRulesNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CloneRules[0:len((*c.CallOptions).CloneRules):len((*c.CallOptions).CloneRules)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.CloneRules(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) Get(ctx context.Context, req *computepb.GetNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPolicy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.FirewallPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) GetAssociation(ctx context.Context, req *computepb.GetAssociationNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPolicyAssociation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetAssociation[0:len((*c.CallOptions).GetAssociation):len((*c.CallOptions).GetAssociation)], opts...)
	var resp *computepb.FirewallPolicyAssociation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.GetAssociation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *computepb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) GetRule(ctx context.Context, req *computepb.GetRuleNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.FirewallPolicyRule, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetRule[0:len((*c.CallOptions).GetRule):len((*c.CallOptions).GetRule)], opts...)
	var resp *computepb.FirewallPolicyRule
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.GetRule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) Insert(ctx context.Context, req *computepb.InsertNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) List(ctx context.Context, req *computepb.ListNetworkFirewallPoliciesRequest, opts ...gax.CallOption) *FirewallPolicyIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &FirewallPolicyIterator{}
	req = proto.Clone(req).(*computepb.ListNetworkFirewallPoliciesRequest)
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
			resp, err = c.networkFirewallPoliciesClient.List(ctx, req, settings.GRPC...)
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

func (c *networkFirewallPoliciesGRPCClient) Patch(ctx context.Context, req *computepb.PatchNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) PatchRule(ctx context.Context, req *computepb.PatchRuleNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).PatchRule[0:len((*c.CallOptions).PatchRule):len((*c.CallOptions).PatchRule)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.PatchRule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) RemoveAssociation(ctx context.Context, req *computepb.RemoveAssociationNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RemoveAssociation[0:len((*c.CallOptions).RemoveAssociation):len((*c.CallOptions).RemoveAssociation)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.RemoveAssociation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) RemoveRule(ctx context.Context, req *computepb.RemoveRuleNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "firewall_policy", url.QueryEscape(req.GetFirewallPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RemoveRule[0:len((*c.CallOptions).RemoveRule):len((*c.CallOptions).RemoveRule)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.RemoveRule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *computepb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkFirewallPoliciesGRPCClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsNetworkFirewallPolicyRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *computepb.TestPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkFirewallPoliciesClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
