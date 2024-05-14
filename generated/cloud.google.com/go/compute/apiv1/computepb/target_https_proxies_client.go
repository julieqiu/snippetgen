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

var newTargetHttpsProxiesClientHook clientHook

// TargetHttpsProxiesCallOptions contains the retry settings for each method of TargetHttpsProxiesClient.
type TargetHttpsProxiesCallOptions struct {
	AggregatedList []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	Patch []gax.CallOption
	SetCertificateMap []gax.CallOption
	SetQuicOverride []gax.CallOption
	SetSslCertificates []gax.CallOption
	SetSslPolicy []gax.CallOption
	SetUrlMap []gax.CallOption
}

func defaultTargetHttpsProxiesGRPCClientOptions() []option.ClientOption {
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

func defaultTargetHttpsProxiesCallOptions() *TargetHttpsProxiesCallOptions {
	return &TargetHttpsProxiesCallOptions{
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
		SetCertificateMap: []gax.CallOption{
		},
		SetQuicOverride: []gax.CallOption{
		},
		SetSslCertificates: []gax.CallOption{
		},
		SetSslPolicy: []gax.CallOption{
		},
		SetUrlMap: []gax.CallOption{
		},
	}
}

// internalTargetHttpsProxiesClient is an interface that defines the methods available from Google Compute Engine API.
type internalTargetHttpsProxiesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListTargetHttpsProxiesRequest, ...gax.CallOption) *TargetHttpsProxiesScopedListPairIterator
	Delete(context.Context, *computepb.DeleteTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.TargetHttpsProxy, error)
	Insert(context.Context, *computepb.InsertTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListTargetHttpsProxiesRequest, ...gax.CallOption) *TargetHttpsProxyIterator
	Patch(context.Context, *computepb.PatchTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetCertificateMap(context.Context, *computepb.SetCertificateMapTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetQuicOverride(context.Context, *computepb.SetQuicOverrideTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetSslCertificates(context.Context, *computepb.SetSslCertificatesTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetSslPolicy(context.Context, *computepb.SetSslPolicyTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetUrlMap(context.Context, *computepb.SetUrlMapTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// TargetHttpsProxiesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The TargetHttpsProxies API.
type TargetHttpsProxiesClient struct {
	// The internal transport-dependent client.
	internalClient internalTargetHttpsProxiesClient

	// The call options for this service.
	CallOptions *TargetHttpsProxiesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TargetHttpsProxiesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TargetHttpsProxiesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TargetHttpsProxiesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves the list of all TargetHttpsProxy resources, regional and global, available to the specified project. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *TargetHttpsProxiesClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetHttpsProxiesRequest, opts ...gax.CallOption) *TargetHttpsProxiesScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified TargetHttpsProxy resource.
func (c *TargetHttpsProxiesClient) Delete(ctx context.Context, req *computepb.DeleteTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified TargetHttpsProxy resource.
func (c *TargetHttpsProxiesClient) Get(ctx context.Context, req *computepb.GetTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.TargetHttpsProxy, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a TargetHttpsProxy resource in the specified project using the data included in the request.
func (c *TargetHttpsProxiesClient) Insert(ctx context.Context, req *computepb.InsertTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of TargetHttpsProxy resources available to the specified project.
func (c *TargetHttpsProxiesClient) List(ctx context.Context, req *computepb.ListTargetHttpsProxiesRequest, opts ...gax.CallOption) *TargetHttpsProxyIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified TargetHttpsProxy resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *TargetHttpsProxiesClient) Patch(ctx context.Context, req *computepb.PatchTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetCertificateMap changes the Certificate Map for TargetHttpsProxy.
func (c *TargetHttpsProxiesClient) SetCertificateMap(ctx context.Context, req *computepb.SetCertificateMapTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetCertificateMap(ctx, req, opts...)
}

// SetQuicOverride sets the QUIC override policy for TargetHttpsProxy.
func (c *TargetHttpsProxiesClient) SetQuicOverride(ctx context.Context, req *computepb.SetQuicOverrideTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetQuicOverride(ctx, req, opts...)
}

// SetSslCertificates replaces SslCertificates for TargetHttpsProxy.
func (c *TargetHttpsProxiesClient) SetSslCertificates(ctx context.Context, req *computepb.SetSslCertificatesTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetSslCertificates(ctx, req, opts...)
}

// SetSslPolicy sets the SSL policy for TargetHttpsProxy. The SSL policy specifies the server-side support for SSL features. This affects connections between clients and the HTTPS proxy load balancer. They do not affect the connection between the load balancer and the backends.
func (c *TargetHttpsProxiesClient) SetSslPolicy(ctx context.Context, req *computepb.SetSslPolicyTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetSslPolicy(ctx, req, opts...)
}

// SetUrlMap changes the URL map for TargetHttpsProxy.
func (c *TargetHttpsProxiesClient) SetUrlMap(ctx context.Context, req *computepb.SetUrlMapTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetUrlMap(ctx, req, opts...)
}

// targetHttpsProxiesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type targetHttpsProxiesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TargetHttpsProxiesClient
	CallOptions **TargetHttpsProxiesCallOptions

	// The gRPC API client.
	targetHttpsProxiesClient computepb.TargetHttpsProxiesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewTargetHttpsProxiesClient creates a new target https proxies client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The TargetHttpsProxies API.
func NewTargetHttpsProxiesClient(ctx context.Context, opts ...option.ClientOption) (*TargetHttpsProxiesClient, error) {
	clientOpts := defaultTargetHttpsProxiesGRPCClientOptions()
	if newTargetHttpsProxiesClientHook != nil {
		hookOpts, err := newTargetHttpsProxiesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TargetHttpsProxiesClient{CallOptions: defaultTargetHttpsProxiesCallOptions()}

	c := &targetHttpsProxiesGRPCClient{
		connPool:    connPool,
		targetHttpsProxiesClient: computepb.NewTargetHttpsProxiesClient(connPool),
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
func (c *targetHttpsProxiesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *targetHttpsProxiesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *targetHttpsProxiesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *targetHttpsProxiesGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetHttpsProxiesRequest, opts ...gax.CallOption) *TargetHttpsProxiesScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &TargetHttpsProxiesScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListTargetHttpsProxiesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]TargetHttpsProxiesScopedListPair, string, error) {
		resp := &computepb.TargetHttpsProxyAggregatedList{}
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
			resp, err = c.targetHttpsProxiesClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]TargetHttpsProxiesScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, TargetHttpsProxiesScopedListPair{k, v})
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

func (c *targetHttpsProxiesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_https_proxy", url.QueryEscape(req.GetTargetHttpsProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpsProxiesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetHttpsProxiesGRPCClient) Get(ctx context.Context, req *computepb.GetTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.TargetHttpsProxy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_https_proxy", url.QueryEscape(req.GetTargetHttpsProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.TargetHttpsProxy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpsProxiesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetHttpsProxiesGRPCClient) Insert(ctx context.Context, req *computepb.InsertTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpsProxiesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetHttpsProxiesGRPCClient) List(ctx context.Context, req *computepb.ListTargetHttpsProxiesRequest, opts ...gax.CallOption) *TargetHttpsProxyIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &TargetHttpsProxyIterator{}
	req = proto.Clone(req).(*computepb.ListTargetHttpsProxiesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.TargetHttpsProxy, string, error) {
		resp := &computepb.TargetHttpsProxyList{}
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
			resp, err = c.targetHttpsProxiesClient.List(ctx, req, settings.GRPC...)
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

func (c *targetHttpsProxiesGRPCClient) Patch(ctx context.Context, req *computepb.PatchTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_https_proxy", url.QueryEscape(req.GetTargetHttpsProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpsProxiesClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetHttpsProxiesGRPCClient) SetCertificateMap(ctx context.Context, req *computepb.SetCertificateMapTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_https_proxy", url.QueryEscape(req.GetTargetHttpsProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetCertificateMap[0:len((*c.CallOptions).SetCertificateMap):len((*c.CallOptions).SetCertificateMap)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpsProxiesClient.SetCertificateMap(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetHttpsProxiesGRPCClient) SetQuicOverride(ctx context.Context, req *computepb.SetQuicOverrideTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_https_proxy", url.QueryEscape(req.GetTargetHttpsProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetQuicOverride[0:len((*c.CallOptions).SetQuicOverride):len((*c.CallOptions).SetQuicOverride)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpsProxiesClient.SetQuicOverride(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetHttpsProxiesGRPCClient) SetSslCertificates(ctx context.Context, req *computepb.SetSslCertificatesTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_https_proxy", url.QueryEscape(req.GetTargetHttpsProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetSslCertificates[0:len((*c.CallOptions).SetSslCertificates):len((*c.CallOptions).SetSslCertificates)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpsProxiesClient.SetSslCertificates(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetHttpsProxiesGRPCClient) SetSslPolicy(ctx context.Context, req *computepb.SetSslPolicyTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_https_proxy", url.QueryEscape(req.GetTargetHttpsProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetSslPolicy[0:len((*c.CallOptions).SetSslPolicy):len((*c.CallOptions).SetSslPolicy)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpsProxiesClient.SetSslPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetHttpsProxiesGRPCClient) SetUrlMap(ctx context.Context, req *computepb.SetUrlMapTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "target_https_proxy", url.QueryEscape(req.GetTargetHttpsProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetUrlMap[0:len((*c.CallOptions).SetUrlMap):len((*c.CallOptions).SetUrlMap)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetHttpsProxiesClient.SetUrlMap(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}