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

var newInterconnectRemoteLocationsClientHook clientHook

// InterconnectRemoteLocationsCallOptions contains the retry settings for each method of InterconnectRemoteLocationsClient.
type InterconnectRemoteLocationsCallOptions struct {
	Get []gax.CallOption
	List []gax.CallOption
}

func defaultInterconnectRemoteLocationsGRPCClientOptions() []option.ClientOption {
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

func defaultInterconnectRemoteLocationsCallOptions() *InterconnectRemoteLocationsCallOptions {
	return &InterconnectRemoteLocationsCallOptions{
		Get: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
	}
}

// internalInterconnectRemoteLocationsClient is an interface that defines the methods available from Google Compute Engine API.
type internalInterconnectRemoteLocationsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Get(context.Context, *computepb.GetInterconnectRemoteLocationRequest, ...gax.CallOption) (*computepb.InterconnectRemoteLocation, error)
	List(context.Context, *computepb.ListInterconnectRemoteLocationsRequest, ...gax.CallOption) *InterconnectRemoteLocationIterator
}

// InterconnectRemoteLocationsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The InterconnectRemoteLocations API.
type InterconnectRemoteLocationsClient struct {
	// The internal transport-dependent client.
	internalClient internalInterconnectRemoteLocationsClient

	// The call options for this service.
	CallOptions *InterconnectRemoteLocationsCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *InterconnectRemoteLocationsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *InterconnectRemoteLocationsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *InterconnectRemoteLocationsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Get returns the details for the specified interconnect remote location. Gets a list of available interconnect remote locations by making a list() request.
func (c *InterconnectRemoteLocationsClient) Get(ctx context.Context, req *computepb.GetInterconnectRemoteLocationRequest, opts ...gax.CallOption) (*computepb.InterconnectRemoteLocation, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// List retrieves the list of interconnect remote locations available to the specified project.
func (c *InterconnectRemoteLocationsClient) List(ctx context.Context, req *computepb.ListInterconnectRemoteLocationsRequest, opts ...gax.CallOption) *InterconnectRemoteLocationIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// interconnectRemoteLocationsGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type interconnectRemoteLocationsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing InterconnectRemoteLocationsClient
	CallOptions **InterconnectRemoteLocationsCallOptions

	// The gRPC API client.
	interconnectRemoteLocationsClient computepb.InterconnectRemoteLocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewInterconnectRemoteLocationsClient creates a new interconnect remote locations client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The InterconnectRemoteLocations API.
func NewInterconnectRemoteLocationsClient(ctx context.Context, opts ...option.ClientOption) (*InterconnectRemoteLocationsClient, error) {
	clientOpts := defaultInterconnectRemoteLocationsGRPCClientOptions()
	if newInterconnectRemoteLocationsClientHook != nil {
		hookOpts, err := newInterconnectRemoteLocationsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := InterconnectRemoteLocationsClient{CallOptions: defaultInterconnectRemoteLocationsCallOptions()}

	c := &interconnectRemoteLocationsGRPCClient{
		connPool:    connPool,
		interconnectRemoteLocationsClient: computepb.NewInterconnectRemoteLocationsClient(connPool),
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
func (c *interconnectRemoteLocationsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *interconnectRemoteLocationsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *interconnectRemoteLocationsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *interconnectRemoteLocationsGRPCClient) Get(ctx context.Context, req *computepb.GetInterconnectRemoteLocationRequest, opts ...gax.CallOption) (*computepb.InterconnectRemoteLocation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "interconnect_remote_location", url.QueryEscape(req.GetInterconnectRemoteLocation()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.InterconnectRemoteLocation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.interconnectRemoteLocationsClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *interconnectRemoteLocationsGRPCClient) List(ctx context.Context, req *computepb.ListInterconnectRemoteLocationsRequest, opts ...gax.CallOption) *InterconnectRemoteLocationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &InterconnectRemoteLocationIterator{}
	req = proto.Clone(req).(*computepb.ListInterconnectRemoteLocationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.InterconnectRemoteLocation, string, error) {
		resp := &computepb.InterconnectRemoteLocationList{}
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
			resp, err = c.interconnectRemoteLocationsClient.List(ctx, req, settings.GRPC...)
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
