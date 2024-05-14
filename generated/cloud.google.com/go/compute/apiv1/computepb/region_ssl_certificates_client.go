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

var newRegionSslCertificatesClientHook clientHook

// RegionSslCertificatesCallOptions contains the retry settings for each method of RegionSslCertificatesClient.
type RegionSslCertificatesCallOptions struct {
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
}

func defaultRegionSslCertificatesGRPCClientOptions() []option.ClientOption {
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

func defaultRegionSslCertificatesCallOptions() *RegionSslCertificatesCallOptions {
	return &RegionSslCertificatesCallOptions{
		Delete: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		Insert: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
	}
}

// internalRegionSslCertificatesClient is an interface that defines the methods available from Google Compute Engine API.
type internalRegionSslCertificatesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Delete(context.Context, *computepb.DeleteRegionSslCertificateRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetRegionSslCertificateRequest, ...gax.CallOption) (*computepb.SslCertificate, error)
	Insert(context.Context, *computepb.InsertRegionSslCertificateRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListRegionSslCertificatesRequest, ...gax.CallOption) *SslCertificateIterator
}

// RegionSslCertificatesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The RegionSslCertificates API.
type RegionSslCertificatesClient struct {
	// The internal transport-dependent client.
	internalClient internalRegionSslCertificatesClient

	// The call options for this service.
	CallOptions *RegionSslCertificatesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RegionSslCertificatesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RegionSslCertificatesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RegionSslCertificatesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Delete deletes the specified SslCertificate resource in the region.
func (c *RegionSslCertificatesClient) Delete(ctx context.Context, req *computepb.DeleteRegionSslCertificateRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified SslCertificate resource in the specified region. Get a list of available SSL certificates by making a list() request.
func (c *RegionSslCertificatesClient) Get(ctx context.Context, req *computepb.GetRegionSslCertificateRequest, opts ...gax.CallOption) (*computepb.SslCertificate, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a SslCertificate resource in the specified project and region using the data included in the request
func (c *RegionSslCertificatesClient) Insert(ctx context.Context, req *computepb.InsertRegionSslCertificateRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of SslCertificate resources available to the specified project in the specified region.
func (c *RegionSslCertificatesClient) List(ctx context.Context, req *computepb.ListRegionSslCertificatesRequest, opts ...gax.CallOption) *SslCertificateIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// regionSslCertificatesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type regionSslCertificatesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing RegionSslCertificatesClient
	CallOptions **RegionSslCertificatesCallOptions

	// The gRPC API client.
	regionSslCertificatesClient computepb.RegionSslCertificatesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewRegionSslCertificatesClient creates a new region ssl certificates client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The RegionSslCertificates API.
func NewRegionSslCertificatesClient(ctx context.Context, opts ...option.ClientOption) (*RegionSslCertificatesClient, error) {
	clientOpts := defaultRegionSslCertificatesGRPCClientOptions()
	if newRegionSslCertificatesClientHook != nil {
		hookOpts, err := newRegionSslCertificatesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RegionSslCertificatesClient{CallOptions: defaultRegionSslCertificatesCallOptions()}

	c := &regionSslCertificatesGRPCClient{
		connPool:    connPool,
		regionSslCertificatesClient: computepb.NewRegionSslCertificatesClient(connPool),
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
func (c *regionSslCertificatesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *regionSslCertificatesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *regionSslCertificatesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *regionSslCertificatesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteRegionSslCertificateRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "ssl_certificate", url.QueryEscape(req.GetSslCertificate()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionSslCertificatesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionSslCertificatesGRPCClient) Get(ctx context.Context, req *computepb.GetRegionSslCertificateRequest, opts ...gax.CallOption) (*computepb.SslCertificate, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "ssl_certificate", url.QueryEscape(req.GetSslCertificate()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.SslCertificate
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionSslCertificatesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionSslCertificatesGRPCClient) Insert(ctx context.Context, req *computepb.InsertRegionSslCertificateRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionSslCertificatesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionSslCertificatesGRPCClient) List(ctx context.Context, req *computepb.ListRegionSslCertificatesRequest, opts ...gax.CallOption) *SslCertificateIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &SslCertificateIterator{}
	req = proto.Clone(req).(*computepb.ListRegionSslCertificatesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.SslCertificate, string, error) {
		resp := &computepb.SslCertificateList{}
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
			resp, err = c.regionSslCertificatesClient.List(ctx, req, settings.GRPC...)
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