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
	"sort"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var newForwardingRulesClientHook clientHook

// ForwardingRulesCallOptions contains the retry settings for each method of ForwardingRulesClient.
type ForwardingRulesCallOptions struct {
	AggregatedList []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	Patch []gax.CallOption
	SetLabels []gax.CallOption
	SetTarget []gax.CallOption
}

func defaultForwardingRulesGRPCClientOptions() []option.ClientOption {
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

func defaultForwardingRulesCallOptions() *ForwardingRulesCallOptions {
	return &ForwardingRulesCallOptions{
		AggregatedList: []gax.CallOption{
		},
		Delete: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		Insert: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
		Patch: []gax.CallOption{
		},
		SetLabels: []gax.CallOption{
		},
		SetTarget: []gax.CallOption{
		},
	}
}

// internalForwardingRulesClient is an interface that defines the methods available from Google Compute Engine API.
type internalForwardingRulesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListForwardingRulesRequest, ...gax.CallOption) *ForwardingRulesScopedListPairIterator
	Delete(context.Context, *computepb.DeleteForwardingRuleRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetForwardingRuleRequest, ...gax.CallOption) (*computepb.ForwardingRule, error)
	Insert(context.Context, *computepb.InsertForwardingRuleRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListForwardingRulesRequest, ...gax.CallOption) *ForwardingRuleIterator
	Patch(context.Context, *computepb.PatchForwardingRuleRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetLabels(context.Context, *computepb.SetLabelsForwardingRuleRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetTarget(context.Context, *computepb.SetTargetForwardingRuleRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// ForwardingRulesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The ForwardingRules API.
type ForwardingRulesClient struct {
	// The internal transport-dependent client.
	internalClient internalForwardingRulesClient

	// The call options for this service.
	CallOptions *ForwardingRulesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ForwardingRulesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ForwardingRulesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ForwardingRulesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves an aggregated list of forwarding rules. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *ForwardingRulesClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListForwardingRulesRequest, opts ...gax.CallOption) *ForwardingRulesScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified ForwardingRule resource.
func (c *ForwardingRulesClient) Delete(ctx context.Context, req *computepb.DeleteForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified ForwardingRule resource.
func (c *ForwardingRulesClient) Get(ctx context.Context, req *computepb.GetForwardingRuleRequest, opts ...gax.CallOption) (*computepb.ForwardingRule, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a ForwardingRule resource in the specified project and region using the data included in the request.
func (c *ForwardingRulesClient) Insert(ctx context.Context, req *computepb.InsertForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of ForwardingRule resources available to the specified project and region.
func (c *ForwardingRulesClient) List(ctx context.Context, req *computepb.ListForwardingRulesRequest, opts ...gax.CallOption) *ForwardingRuleIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch updates the specified forwarding rule with the data included in the request. This method supports PATCH semantics and uses the JSON merge patch format and processing rules. Currently, you can only patch the network_tier field.
func (c *ForwardingRulesClient) Patch(ctx context.Context, req *computepb.PatchForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetLabels sets the labels on the specified resource. To learn more about labels, read the Labeling Resources documentation.
func (c *ForwardingRulesClient) SetLabels(ctx context.Context, req *computepb.SetLabelsForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetLabels(ctx, req, opts...)
}

// SetTarget changes target URL for forwarding rule. The new target should be of the same type as the old target.
func (c *ForwardingRulesClient) SetTarget(ctx context.Context, req *computepb.SetTargetForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetTarget(ctx, req, opts...)
}

// forwardingRulesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type forwardingRulesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ForwardingRulesClient
	CallOptions **ForwardingRulesCallOptions

	// The gRPC API client.
	forwardingRulesClient computepb.ForwardingRulesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewForwardingRulesClient creates a new forwarding rules client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The ForwardingRules API.
func NewForwardingRulesClient(ctx context.Context, opts ...option.ClientOption) (*ForwardingRulesClient, error) {
	clientOpts := defaultForwardingRulesGRPCClientOptions()
	if newForwardingRulesClientHook != nil {
		hookOpts, err := newForwardingRulesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ForwardingRulesClient{CallOptions: defaultForwardingRulesCallOptions()}

	c := &forwardingRulesGRPCClient{
		connPool:    connPool,
		forwardingRulesClient: computepb.NewForwardingRulesClient(connPool),
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
func (c *forwardingRulesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *forwardingRulesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *forwardingRulesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *forwardingRulesGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListForwardingRulesRequest, opts ...gax.CallOption) *ForwardingRulesScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &ForwardingRulesScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListForwardingRulesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]ForwardingRulesScopedListPair, string, error) {
		resp := &computepb.ForwardingRuleAggregatedList{}
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
			resp, err = c.forwardingRulesClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]ForwardingRulesScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, ForwardingRulesScopedListPair{k, v})
		}
		sort.Slice(elems, func(i, j int) bool { return elems[i].Key < elems[j].Key } )

		return elems, resp.GetNextPageToken(), nil
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

func (c *forwardingRulesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "forwarding_rule", url.QueryEscape(req.GetForwardingRule()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.forwardingRulesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *forwardingRulesGRPCClient) Get(ctx context.Context, req *computepb.GetForwardingRuleRequest, opts ...gax.CallOption) (*computepb.ForwardingRule, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "forwarding_rule", url.QueryEscape(req.GetForwardingRule()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.ForwardingRule
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.forwardingRulesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *forwardingRulesGRPCClient) Insert(ctx context.Context, req *computepb.InsertForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.forwardingRulesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *forwardingRulesGRPCClient) List(ctx context.Context, req *computepb.ListForwardingRulesRequest, opts ...gax.CallOption) *ForwardingRuleIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &ForwardingRuleIterator{}
	req = proto.Clone(req).(*computepb.ListForwardingRulesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.ForwardingRule, string, error) {
		resp := &computepb.ForwardingRuleList{}
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
			resp, err = c.forwardingRulesClient.List(ctx, req, settings.GRPC...)
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

func (c *forwardingRulesGRPCClient) Patch(ctx context.Context, req *computepb.PatchForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "forwarding_rule", url.QueryEscape(req.GetForwardingRule()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.forwardingRulesClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *forwardingRulesGRPCClient) SetLabels(ctx context.Context, req *computepb.SetLabelsForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetLabels[0:len((*c.CallOptions).SetLabels):len((*c.CallOptions).SetLabels)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.forwardingRulesClient.SetLabels(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *forwardingRulesGRPCClient) SetTarget(ctx context.Context, req *computepb.SetTargetForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "forwarding_rule", url.QueryEscape(req.GetForwardingRule()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetTarget[0:len((*c.CallOptions).SetTarget):len((*c.CallOptions).SetTarget)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.forwardingRulesClient.SetTarget(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}