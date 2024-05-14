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

var newTargetTcpProxiesClientHook clientHook

// TargetTcpProxiesCallOptions contains the retry settings for each method of TargetTcpProxiesClient.
type TargetTcpProxiesCallOptions struct {
	AggregatedList []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	SetBackendService []gax.CallOption
	SetProxyHeader []gax.CallOption
}

func defaultTargetTcpProxiesGRPCClientOptions() []option.ClientOption {
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

func defaultTargetTcpProxiesCallOptions() *TargetTcpProxiesCallOptions {
	return &TargetTcpProxiesCallOptions{
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
		SetBackendService: []gax.CallOption{
		},
		SetProxyHeader: []gax.CallOption{
		},
	}
}

// internalTargetTcpProxiesClient is an interface that defines the methods available from Google Compute Engine API.
type internalTargetTcpProxiesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListTargetTcpProxiesRequest, ...gax.CallOption) *TargetTcpProxiesScopedListPairIterator
	Delete(context.Context, *computepb.DeleteTargetTcpProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetTargetTcpProxyRequest, ...gax.CallOption) (*computepb.TargetTcpProxy, error)
	Insert(context.Context, *computepb.InsertTargetTcpProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListTargetTcpProxiesRequest, ...gax.CallOption) *TargetTcpProxyIterator
	SetBackendService(context.Context, *computepb.SetBackendServiceTargetTcpProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetProxyHeader(context.Context, *computepb.SetProxyHeaderTargetTcpProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// TargetTcpProxiesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The TargetTcpProxies API.
type TargetTcpProxiesClient struct {
	// The internal transport-dependent client.
	internalClient internalTargetTcpProxiesClient

	// The call options for this service.
	CallOptions *TargetTcpProxiesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TargetTcpProxiesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TargetTcpProxiesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TargetTcpProxiesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves the list of all TargetTcpProxy resources, regional and global, available to the specified project. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *TargetTcpProxiesClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetTcpProxiesRequest, opts ...gax.CallOption) *TargetTcpProxiesScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified TargetTcpProxy resource.
func (c *TargetTcpProxiesClient) Delete(ctx context.Context, req *computepb.DeleteTargetTcpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified TargetTcpProxy resource.
func (c *TargetTcpProxiesClient) Get(ctx context.Context, req *computepb.GetTargetTcpProxyRequest, opts ...gax.CallOption) (*computepb.TargetTcpProxy, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a TargetTcpProxy resource in the specified project using the data included in the request.
func (c *TargetTcpProxiesClient) Insert(ctx context.Context, req *computepb.InsertTargetTcpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of TargetTcpProxy resources available to the specified project.
func (c *TargetTcpProxiesClient) List(ctx context.Context, req *computepb.ListTargetTcpProxiesRequest, opts ...gax.CallOption) *TargetTcpProxyIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// SetBackendService changes the BackendService for TargetTcpProxy.
func (c *TargetTcpProxiesClient) SetBackendService(ctx context.Context, req *computepb.SetBackendServiceTargetTcpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetBackendService(ctx, req, opts...)
}

// SetProxyHeader changes the ProxyHeaderType for TargetTcpProxy.
func (c *TargetTcpProxiesClient) SetProxyHeader(ctx context.Context, req *computepb.SetProxyHeaderTargetTcpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetProxyHeader(ctx, req, opts...)
}

// targetTcpProxiesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type targetTcpProxiesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TargetTcpProxiesClient
	CallOptions **TargetTcpProxiesCallOptions

	// The gRPC API client.
	targetTcpProxiesClient computepb.TargetTcpProxiesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewTargetTcpProxiesClient creates a new target tcp proxies client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The TargetTcpProxies API.
func NewTargetTcpProxiesClient(ctx context.Context, opts ...option.ClientOption) (*TargetTcpProxiesClient, error) {
	clientOpts := defaultTargetTcpProxiesGRPCClientOptions()
	if newTargetTcpProxiesClientHook != nil {
		hookOpts, err := newTargetTcpProxiesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TargetTcpProxiesClient{CallOptions: defaultTargetTcpProxiesCallOptions()}

	c := &targetTcpProxiesGRPCClient{
		connPool:    connPool,
		targetTcpProxiesClient: computepb.NewTargetTcpProxiesClient(connPool),
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
func (c *targetTcpProxiesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *targetTcpProxiesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *targetTcpProxiesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *targetTcpProxiesGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetTcpProxiesRequest, opts ...gax.CallOption) *TargetTcpProxiesScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &TargetTcpProxiesScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListTargetTcpProxiesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]TargetTcpProxiesScopedListPair, string, error) {
		resp := &computepb.TargetTcpProxyAggregatedList{}
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
			resp, err = c.targetTcpProxiesClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]TargetTcpProxiesScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, TargetTcpProxiesScopedListPair{k, v})
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

func (c *targetTcpProxiesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteTargetTcpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_tcp_proxy", url.QueryEscape(req.GetTargetTcpProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetTcpProxiesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetTcpProxiesGRPCClient) Get(ctx context.Context, req *computepb.GetTargetTcpProxyRequest, opts ...gax.CallOption) (*computepb.TargetTcpProxy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_tcp_proxy", url.QueryEscape(req.GetTargetTcpProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.TargetTcpProxy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetTcpProxiesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetTcpProxiesGRPCClient) Insert(ctx context.Context, req *computepb.InsertTargetTcpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetTcpProxiesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetTcpProxiesGRPCClient) List(ctx context.Context, req *computepb.ListTargetTcpProxiesRequest, opts ...gax.CallOption) *TargetTcpProxyIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &TargetTcpProxyIterator{}
	req = proto.Clone(req).(*computepb.ListTargetTcpProxiesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.TargetTcpProxy, string, error) {
		resp := &computepb.TargetTcpProxyList{}
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
			resp, err = c.targetTcpProxiesClient.List(ctx, req, settings.GRPC...)
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

func (c *targetTcpProxiesGRPCClient) SetBackendService(ctx context.Context, req *computepb.SetBackendServiceTargetTcpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_tcp_proxy", url.QueryEscape(req.GetTargetTcpProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetBackendService[0:len((*c.CallOptions).SetBackendService):len((*c.CallOptions).SetBackendService)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetTcpProxiesClient.SetBackendService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetTcpProxiesGRPCClient) SetProxyHeader(ctx context.Context, req *computepb.SetProxyHeaderTargetTcpProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_tcp_proxy", url.QueryEscape(req.GetTargetTcpProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetProxyHeader[0:len((*c.CallOptions).SetProxyHeader):len((*c.CallOptions).SetProxyHeader)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetTcpProxiesClient.SetProxyHeader(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}