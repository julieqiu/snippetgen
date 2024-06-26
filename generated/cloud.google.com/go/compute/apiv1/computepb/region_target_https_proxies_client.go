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

var newRegionTargetHttpsProxiesClientHook clientHook

// RegionTargetHttpsProxiesCallOptions contains the retry settings for each method of RegionTargetHttpsProxiesClient.
type RegionTargetHttpsProxiesCallOptions struct {
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	Patch []gax.CallOption
	SetSslCertificates []gax.CallOption
	SetUrlMap []gax.CallOption
}

func defaultRegionTargetHttpsProxiesGRPCClientOptions() []option.ClientOption {
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

func defaultRegionTargetHttpsProxiesCallOptions() *RegionTargetHttpsProxiesCallOptions {
	return &RegionTargetHttpsProxiesCallOptions{
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
		SetSslCertificates: []gax.CallOption{
		},
		SetUrlMap: []gax.CallOption{
		},
	}
}

// internalRegionTargetHttpsProxiesClient is an interface that defines the methods available from Google Compute Engine API.
type internalRegionTargetHttpsProxiesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Delete(context.Context, *computepb.DeleteRegionTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetRegionTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.TargetHttpsProxy, error)
	Insert(context.Context, *computepb.InsertRegionTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListRegionTargetHttpsProxiesRequest, ...gax.CallOption) *TargetHttpsProxyIterator
	Patch(context.Context, *computepb.PatchRegionTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetSslCertificates(context.Context, *computepb.SetSslCertificatesRegionTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetUrlMap(context.Context, *computepb.SetUrlMapRegionTargetHttpsProxyRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// RegionTargetHttpsProxiesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The RegionTargetHttpsProxies API.
type RegionTargetHttpsProxiesClient struct {
	// The internal transport-dependent client.
	internalClient internalRegionTargetHttpsProxiesClient

	// The call options for this service.
	CallOptions *RegionTargetHttpsProxiesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RegionTargetHttpsProxiesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RegionTargetHttpsProxiesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RegionTargetHttpsProxiesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Delete deletes the specified TargetHttpsProxy resource.
func (c *RegionTargetHttpsProxiesClient) Delete(ctx context.Context, req *computepb.DeleteRegionTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified TargetHttpsProxy resource in the specified region.
func (c *RegionTargetHttpsProxiesClient) Get(ctx context.Context, req *computepb.GetRegionTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.TargetHttpsProxy, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a TargetHttpsProxy resource in the specified project and region using the data included in the request.
func (c *RegionTargetHttpsProxiesClient) Insert(ctx context.Context, req *computepb.InsertRegionTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of TargetHttpsProxy resources available to the specified project in the specified region.
func (c *RegionTargetHttpsProxiesClient) List(ctx context.Context, req *computepb.ListRegionTargetHttpsProxiesRequest, opts ...gax.CallOption) *TargetHttpsProxyIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified regional TargetHttpsProxy resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *RegionTargetHttpsProxiesClient) Patch(ctx context.Context, req *computepb.PatchRegionTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetSslCertificates replaces SslCertificates for TargetHttpsProxy.
func (c *RegionTargetHttpsProxiesClient) SetSslCertificates(ctx context.Context, req *computepb.SetSslCertificatesRegionTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetSslCertificates(ctx, req, opts...)
}

// SetUrlMap changes the URL map for TargetHttpsProxy.
func (c *RegionTargetHttpsProxiesClient) SetUrlMap(ctx context.Context, req *computepb.SetUrlMapRegionTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetUrlMap(ctx, req, opts...)
}

// regionTargetHttpsProxiesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type regionTargetHttpsProxiesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing RegionTargetHttpsProxiesClient
	CallOptions **RegionTargetHttpsProxiesCallOptions

	// The gRPC API client.
	regionTargetHttpsProxiesClient computepb.RegionTargetHttpsProxiesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewRegionTargetHttpsProxiesClient creates a new region target https proxies client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The RegionTargetHttpsProxies API.
func NewRegionTargetHttpsProxiesClient(ctx context.Context, opts ...option.ClientOption) (*RegionTargetHttpsProxiesClient, error) {
	clientOpts := defaultRegionTargetHttpsProxiesGRPCClientOptions()
	if newRegionTargetHttpsProxiesClientHook != nil {
		hookOpts, err := newRegionTargetHttpsProxiesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RegionTargetHttpsProxiesClient{CallOptions: defaultRegionTargetHttpsProxiesCallOptions()}

	c := &regionTargetHttpsProxiesGRPCClient{
		connPool:    connPool,
		regionTargetHttpsProxiesClient: computepb.NewRegionTargetHttpsProxiesClient(connPool),
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
func (c *regionTargetHttpsProxiesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *regionTargetHttpsProxiesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *regionTargetHttpsProxiesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *regionTargetHttpsProxiesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteRegionTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_https_proxy", url.QueryEscape(req.GetTargetHttpsProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionTargetHttpsProxiesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionTargetHttpsProxiesGRPCClient) Get(ctx context.Context, req *computepb.GetRegionTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.TargetHttpsProxy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_https_proxy", url.QueryEscape(req.GetTargetHttpsProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.TargetHttpsProxy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionTargetHttpsProxiesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionTargetHttpsProxiesGRPCClient) Insert(ctx context.Context, req *computepb.InsertRegionTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionTargetHttpsProxiesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionTargetHttpsProxiesGRPCClient) List(ctx context.Context, req *computepb.ListRegionTargetHttpsProxiesRequest, opts ...gax.CallOption) *TargetHttpsProxyIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &TargetHttpsProxyIterator{}
	req = proto.Clone(req).(*computepb.ListRegionTargetHttpsProxiesRequest)
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
			resp, err = c.regionTargetHttpsProxiesClient.List(ctx, req, settings.GRPC...)
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

func (c *regionTargetHttpsProxiesGRPCClient) Patch(ctx context.Context, req *computepb.PatchRegionTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_https_proxy", url.QueryEscape(req.GetTargetHttpsProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionTargetHttpsProxiesClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionTargetHttpsProxiesGRPCClient) SetSslCertificates(ctx context.Context, req *computepb.SetSslCertificatesRegionTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_https_proxy", url.QueryEscape(req.GetTargetHttpsProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetSslCertificates[0:len((*c.CallOptions).SetSslCertificates):len((*c.CallOptions).SetSslCertificates)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionTargetHttpsProxiesClient.SetSslCertificates(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionTargetHttpsProxiesGRPCClient) SetUrlMap(ctx context.Context, req *computepb.SetUrlMapRegionTargetHttpsProxyRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_https_proxy", url.QueryEscape(req.GetTargetHttpsProxy()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetUrlMap[0:len((*c.CallOptions).SetUrlMap):len((*c.CallOptions).SetUrlMap)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionTargetHttpsProxiesClient.SetUrlMap(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
