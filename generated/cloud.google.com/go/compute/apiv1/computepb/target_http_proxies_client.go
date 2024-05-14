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

var newTargetHttpProxiesClientHook clientHook

// TargetHttpProxiesCallOptions contains the retry settings for each method of TargetHttpProxiesClient.
type TargetHttpProxiesCallOptions struct {
	AggregatedList []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	Patch []gax.CallOption
	SetUrlMap []gax.CallOption
}

func defaultTargetHttpProxiesGRPCClientOptions() []option.ClientOption {
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

func defaultTargetHttpProxiesCallOptions() *TargetHttpProxiesCallOptions {
	return &TargetHttpProxiesCallOptions{
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
		SetUrlMap: []gax.CallOption{
		},
	}
}

// internalTargetHttpProxiesClient is an interface that defines the methods available from Google Compute Engine API.
type internalTargetHttpProxiesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListTargetHttpProxiesRequest, ...gax.CallOption) *TargetHttpProxiesScopedListPairIterator
	Delete(context.Context, *computepb.DeleteTargetHttpProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetTargetHttpProxyRequest, ...gax.CallOption) (*computepb.TargetHttpProxy, error)
	Insert(context.Context, *computepb.InsertTargetHttpProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListTargetHttpProxiesRequest, ...gax.CallOption) *TargetHttpProxyIterator
	Patch(context.Context, *computepb.PatchTargetHttpProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetUrlMap(context.Context, *computepb.SetUrlMapTargetHttpProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// TargetHttpProxiesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The TargetHttpProxies API.
type TargetHttpProxiesClient struct {
	// The internal transport-dependent client.
	internalClient internalTargetHttpProxiesClient

	// The call options for this service.
	CallOptions *TargetHttpProxiesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TargetHttpProxiesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TargetHttpProxiesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TargetHttpProxiesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves the list of all TargetHttpProxy resources, regional and global, available to the specified project. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *TargetHttpProxiesClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetHttpProxiesRequest, opts ...gax.CallOption) *TargetHttpProxiesScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified TargetHttpProxy resource.
func (c *TargetHttpProxiesClient) Delete(ctx context.Context, req *computepb.DeleteTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified TargetHttpProxy resource.
func (c *TargetHttpProxiesClient) Get(ctx context.Context, req *computepb.GetTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.TargetHttpProxy, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a TargetHttpProxy resource in the specified project using the data included in the request.
func (c *TargetHttpProxiesClient) Insert(ctx context.Context, req *computepb.InsertTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of TargetHttpProxy resources available to the specified project.
func (c *TargetHttpProxiesClient) List(ctx context.Context, req *computepb.ListTargetHttpProxiesRequest, opts ...gax.CallOption) *TargetHttpProxyIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified TargetHttpProxy resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *TargetHttpProxiesClient) Patch(ctx context.Context, req *computepb.PatchTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetUrlMap changes the URL map for TargetHttpProxy.
func (c *TargetHttpProxiesClient) SetUrlMap(ctx context.Context, req *computepb.SetUrlMapTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetUrlMap(ctx, req, opts...)
}

// targetHttpProxiesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type targetHttpProxiesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TargetHttpProxiesClient
	CallOptions **TargetHttpProxiesCallOptions

	// The gRPC API client.
	targetHttpProxiesClient computepb.TargetHttpProxiesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewTargetHttpProxiesClient creates a new target http proxies client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The TargetHttpProxies API.
func NewTargetHttpProxiesClient(ctx context.Context, opts ...option.ClientOption) (*TargetHttpProxiesClient, error) {
	clientOpts := defaultTargetHttpProxiesGRPCClientOptions()
	if newTargetHttpProxiesClientHook != nil {
		hookOpts, err := newTargetHttpProxiesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TargetHttpProxiesClient{CallOptions: defaultTargetHttpProxiesCallOptions()}

	c := &targetHttpProxiesGRPCClient{
		connPool:    connPool,
		targetHttpProxiesClient: computepb.NewTargetHttpProxiesClient(connPool),
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
func (c *targetHttpProxiesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *targetHttpProxiesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *targetHttpProxiesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *targetHttpProxiesGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetHttpProxiesRequest, opts ...gax.CallOption) *TargetHttpProxiesScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &TargetHttpProxiesScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListTargetHttpProxiesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]TargetHttpProxiesScopedListPair, string, error) {
		resp := &computepb.TargetHttpProxyAggregatedList{}
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
			resp, err = c.targetHttpProxiesClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]TargetHttpProxiesScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, TargetHttpProxiesScopedListPair{k, v})
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

func (c *targetHttpProxiesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_http_proxy", url.QueryEscape(req.GetTargetHttpProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpProxiesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetHttpProxiesGRPCClient) Get(ctx context.Context, req *computepb.GetTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.TargetHttpProxy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_http_proxy", url.QueryEscape(req.GetTargetHttpProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.TargetHttpProxy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpProxiesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetHttpProxiesGRPCClient) Insert(ctx context.Context, req *computepb.InsertTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpProxiesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetHttpProxiesGRPCClient) List(ctx context.Context, req *computepb.ListTargetHttpProxiesRequest, opts ...gax.CallOption) *TargetHttpProxyIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &TargetHttpProxyIterator{}
	req = proto.Clone(req).(*computepb.ListTargetHttpProxiesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.TargetHttpProxy, string, error) {
		resp := &computepb.TargetHttpProxyList{}
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
			resp, err = c.targetHttpProxiesClient.List(ctx, req, settings.GRPC...)
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

func (c *targetHttpProxiesGRPCClient) Patch(ctx context.Context, req *computepb.PatchTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_http_proxy", url.QueryEscape(req.GetTargetHttpProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpProxiesClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetHttpProxiesGRPCClient) SetUrlMap(ctx context.Context, req *computepb.SetUrlMapTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_http_proxy", url.QueryEscape(req.GetTargetHttpProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetUrlMap[0:len((*c.CallOptions).SetUrlMap):len((*c.CallOptions).SetUrlMap)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpProxiesClient.SetUrlMap(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
