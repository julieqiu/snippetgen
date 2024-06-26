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

var newRegionSecurityPoliciesClientHook clientHook

// RegionSecurityPoliciesCallOptions contains the retry settings for each method of RegionSecurityPoliciesClient.
type RegionSecurityPoliciesCallOptions struct {
	AddRule []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	GetRule []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	Patch []gax.CallOption
	PatchRule []gax.CallOption
	RemoveRule []gax.CallOption
}

func defaultRegionSecurityPoliciesGRPCClientOptions() []option.ClientOption {
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

func defaultRegionSecurityPoliciesCallOptions() *RegionSecurityPoliciesCallOptions {
	return &RegionSecurityPoliciesCallOptions{
		AddRule: []gax.CallOption{
		},
		Delete: []gax.CallOption{
		},
		Get: []gax.CallOption{
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
		RemoveRule: []gax.CallOption{
		},
	}
}

// internalRegionSecurityPoliciesClient is an interface that defines the methods available from Google Compute Engine API.
type internalRegionSecurityPoliciesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddRule(context.Context, *computepb.AddRuleRegionSecurityPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	Delete(context.Context, *computepb.DeleteRegionSecurityPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetRegionSecurityPolicyRequest, ...gax.CallOption) (*computepb.SecurityPolicy, error)
	GetRule(context.Context, *computepb.GetRuleRegionSecurityPolicyRequest, ...gax.CallOption) (*computepb.SecurityPolicyRule, error)
	Insert(context.Context, *computepb.InsertRegionSecurityPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListRegionSecurityPoliciesRequest, ...gax.CallOption) *SecurityPolicyIterator
	Patch(context.Context, *computepb.PatchRegionSecurityPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	PatchRule(context.Context, *computepb.PatchRuleRegionSecurityPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
	RemoveRule(context.Context, *computepb.RemoveRuleRegionSecurityPolicyRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// RegionSecurityPoliciesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The RegionSecurityPolicies API.
type RegionSecurityPoliciesClient struct {
	// The internal transport-dependent client.
	internalClient internalRegionSecurityPoliciesClient

	// The call options for this service.
	CallOptions *RegionSecurityPoliciesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RegionSecurityPoliciesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RegionSecurityPoliciesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RegionSecurityPoliciesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddRule inserts a rule into a security policy.
func (c *RegionSecurityPoliciesClient) AddRule(ctx context.Context, req *computepb.AddRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddRule(ctx, req, opts...)
}

// Delete deletes the specified policy.
func (c *RegionSecurityPoliciesClient) Delete(ctx context.Context, req *computepb.DeleteRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get list all of the ordered rules present in a single specified policy.
func (c *RegionSecurityPoliciesClient) Get(ctx context.Context, req *computepb.GetRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.SecurityPolicy, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetRule gets a rule at the specified priority.
func (c *RegionSecurityPoliciesClient) GetRule(ctx context.Context, req *computepb.GetRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.SecurityPolicyRule, error) {
	return c.internalClient.GetRule(ctx, req, opts...)
}

// Insert creates a new policy in the specified project using the data included in the request.
func (c *RegionSecurityPoliciesClient) Insert(ctx context.Context, req *computepb.InsertRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List list all the policies that have been configured for the specified project and region.
func (c *RegionSecurityPoliciesClient) List(ctx context.Context, req *computepb.ListRegionSecurityPoliciesRequest, opts ...gax.CallOption) *SecurityPolicyIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified policy with the data included in the request. To clear fields in the policy, leave the fields empty and specify them in the updateMask. This cannot be used to be update the rules in the policy. Please use the per rule methods like addRule, patchRule, and removeRule instead.
func (c *RegionSecurityPoliciesClient) Patch(ctx context.Context, req *computepb.PatchRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// PatchRule patches a rule at the specified priority. To clear fields in the rule, leave the fields empty and specify them in the updateMask.
func (c *RegionSecurityPoliciesClient) PatchRule(ctx context.Context, req *computepb.PatchRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.PatchRule(ctx, req, opts...)
}

// RemoveRule deletes a rule at the specified priority.
func (c *RegionSecurityPoliciesClient) RemoveRule(ctx context.Context, req *computepb.RemoveRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.RemoveRule(ctx, req, opts...)
}

// regionSecurityPoliciesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type regionSecurityPoliciesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing RegionSecurityPoliciesClient
	CallOptions **RegionSecurityPoliciesCallOptions

	// The gRPC API client.
	regionSecurityPoliciesClient computepb.RegionSecurityPoliciesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewRegionSecurityPoliciesClient creates a new region security policies client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The RegionSecurityPolicies API.
func NewRegionSecurityPoliciesClient(ctx context.Context, opts ...option.ClientOption) (*RegionSecurityPoliciesClient, error) {
	clientOpts := defaultRegionSecurityPoliciesGRPCClientOptions()
	if newRegionSecurityPoliciesClientHook != nil {
		hookOpts, err := newRegionSecurityPoliciesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RegionSecurityPoliciesClient{CallOptions: defaultRegionSecurityPoliciesCallOptions()}

	c := &regionSecurityPoliciesGRPCClient{
		connPool:    connPool,
		regionSecurityPoliciesClient: computepb.NewRegionSecurityPoliciesClient(connPool),
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
func (c *regionSecurityPoliciesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *regionSecurityPoliciesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *regionSecurityPoliciesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *regionSecurityPoliciesGRPCClient) AddRule(ctx context.Context, req *computepb.AddRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AddRule[0:len((*c.CallOptions).AddRule):len((*c.CallOptions).AddRule)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionSecurityPoliciesClient.AddRule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionSecurityPoliciesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionSecurityPoliciesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionSecurityPoliciesGRPCClient) Get(ctx context.Context, req *computepb.GetRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.SecurityPolicy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.SecurityPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionSecurityPoliciesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionSecurityPoliciesGRPCClient) GetRule(ctx context.Context, req *computepb.GetRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.SecurityPolicyRule, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetRule[0:len((*c.CallOptions).GetRule):len((*c.CallOptions).GetRule)], opts...)
	var resp *computepb.SecurityPolicyRule
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionSecurityPoliciesClient.GetRule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionSecurityPoliciesGRPCClient) Insert(ctx context.Context, req *computepb.InsertRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionSecurityPoliciesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionSecurityPoliciesGRPCClient) List(ctx context.Context, req *computepb.ListRegionSecurityPoliciesRequest, opts ...gax.CallOption) *SecurityPolicyIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &SecurityPolicyIterator{}
	req = proto.Clone(req).(*computepb.ListRegionSecurityPoliciesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.SecurityPolicy, string, error) {
		resp := &computepb.SecurityPolicyList{}
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
			resp, err = c.regionSecurityPoliciesClient.List(ctx, req, settings.GRPC...)
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

func (c *regionSecurityPoliciesGRPCClient) Patch(ctx context.Context, req *computepb.PatchRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionSecurityPoliciesClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionSecurityPoliciesGRPCClient) PatchRule(ctx context.Context, req *computepb.PatchRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).PatchRule[0:len((*c.CallOptions).PatchRule):len((*c.CallOptions).PatchRule)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionSecurityPoliciesClient.PatchRule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionSecurityPoliciesGRPCClient) RemoveRule(ctx context.Context, req *computepb.RemoveRuleRegionSecurityPolicyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "security_policy", url.QueryEscape(req.GetSecurityPolicy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RemoveRule[0:len((*c.CallOptions).RemoveRule):len((*c.CallOptions).RemoveRule)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionSecurityPoliciesClient.RemoveRule(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
