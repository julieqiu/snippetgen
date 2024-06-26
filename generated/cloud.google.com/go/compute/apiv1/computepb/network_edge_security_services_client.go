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

var newNetworkEdgeSecurityServicesClientHook clientHook

// NetworkEdgeSecurityServicesCallOptions contains the retry settings for each method of NetworkEdgeSecurityServicesClient.
type NetworkEdgeSecurityServicesCallOptions struct {
	AggregatedList []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	Patch []gax.CallOption
}

func defaultNetworkEdgeSecurityServicesGRPCClientOptions() []option.ClientOption {
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

func defaultNetworkEdgeSecurityServicesCallOptions() *NetworkEdgeSecurityServicesCallOptions {
	return &NetworkEdgeSecurityServicesCallOptions{
		AggregatedList: []gax.CallOption{
		},
		Delete: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		Insert: []gax.CallOption{
		},
		Patch: []gax.CallOption{
		},
	}
}

// internalNetworkEdgeSecurityServicesClient is an interface that defines the methods available from Google Compute Engine API.
type internalNetworkEdgeSecurityServicesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListNetworkEdgeSecurityServicesRequest, ...gax.CallOption) *NetworkEdgeSecurityServicesScopedListPairIterator
	Delete(context.Context, *computepb.DeleteNetworkEdgeSecurityServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetNetworkEdgeSecurityServiceRequest, ...gax.CallOption) (*computepb.NetworkEdgeSecurityService, error)
	Insert(context.Context, *computepb.InsertNetworkEdgeSecurityServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	Patch(context.Context, *computepb.PatchNetworkEdgeSecurityServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// NetworkEdgeSecurityServicesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The NetworkEdgeSecurityServices API.
type NetworkEdgeSecurityServicesClient struct {
	// The internal transport-dependent client.
	internalClient internalNetworkEdgeSecurityServicesClient

	// The call options for this service.
	CallOptions *NetworkEdgeSecurityServicesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *NetworkEdgeSecurityServicesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *NetworkEdgeSecurityServicesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *NetworkEdgeSecurityServicesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves the list of all NetworkEdgeSecurityService resources available to the specified project. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *NetworkEdgeSecurityServicesClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListNetworkEdgeSecurityServicesRequest, opts ...gax.CallOption) *NetworkEdgeSecurityServicesScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified service.
func (c *NetworkEdgeSecurityServicesClient) Delete(ctx context.Context, req *computepb.DeleteNetworkEdgeSecurityServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get gets a specified NetworkEdgeSecurityService.
func (c *NetworkEdgeSecurityServicesClient) Get(ctx context.Context, req *computepb.GetNetworkEdgeSecurityServiceRequest, opts ...gax.CallOption) (*computepb.NetworkEdgeSecurityService, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a new service in the specified project using the data included in the request.
func (c *NetworkEdgeSecurityServicesClient) Insert(ctx context.Context, req *computepb.InsertNetworkEdgeSecurityServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// Patch patches the specified policy with the data included in the request.
func (c *NetworkEdgeSecurityServicesClient) Patch(ctx context.Context, req *computepb.PatchNetworkEdgeSecurityServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// networkEdgeSecurityServicesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type networkEdgeSecurityServicesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing NetworkEdgeSecurityServicesClient
	CallOptions **NetworkEdgeSecurityServicesCallOptions

	// The gRPC API client.
	networkEdgeSecurityServicesClient computepb.NetworkEdgeSecurityServicesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewNetworkEdgeSecurityServicesClient creates a new network edge security services client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The NetworkEdgeSecurityServices API.
func NewNetworkEdgeSecurityServicesClient(ctx context.Context, opts ...option.ClientOption) (*NetworkEdgeSecurityServicesClient, error) {
	clientOpts := defaultNetworkEdgeSecurityServicesGRPCClientOptions()
	if newNetworkEdgeSecurityServicesClientHook != nil {
		hookOpts, err := newNetworkEdgeSecurityServicesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := NetworkEdgeSecurityServicesClient{CallOptions: defaultNetworkEdgeSecurityServicesCallOptions()}

	c := &networkEdgeSecurityServicesGRPCClient{
		connPool:    connPool,
		networkEdgeSecurityServicesClient: computepb.NewNetworkEdgeSecurityServicesClient(connPool),
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
func (c *networkEdgeSecurityServicesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *networkEdgeSecurityServicesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *networkEdgeSecurityServicesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *networkEdgeSecurityServicesGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListNetworkEdgeSecurityServicesRequest, opts ...gax.CallOption) *NetworkEdgeSecurityServicesScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &NetworkEdgeSecurityServicesScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListNetworkEdgeSecurityServicesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]NetworkEdgeSecurityServicesScopedListPair, string, error) {
		resp := &computepb.NetworkEdgeSecurityServiceAggregatedList{}
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
			resp, err = c.networkEdgeSecurityServicesClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]NetworkEdgeSecurityServicesScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, NetworkEdgeSecurityServicesScopedListPair{k, v})
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

func (c *networkEdgeSecurityServicesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteNetworkEdgeSecurityServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "network_edge_security_service", url.QueryEscape(req.GetNetworkEdgeSecurityService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkEdgeSecurityServicesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkEdgeSecurityServicesGRPCClient) Get(ctx context.Context, req *computepb.GetNetworkEdgeSecurityServiceRequest, opts ...gax.CallOption) (*computepb.NetworkEdgeSecurityService, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "network_edge_security_service", url.QueryEscape(req.GetNetworkEdgeSecurityService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.NetworkEdgeSecurityService
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkEdgeSecurityServicesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkEdgeSecurityServicesGRPCClient) Insert(ctx context.Context, req *computepb.InsertNetworkEdgeSecurityServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkEdgeSecurityServicesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *networkEdgeSecurityServicesGRPCClient) Patch(ctx context.Context, req *computepb.PatchNetworkEdgeSecurityServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "network_edge_security_service", url.QueryEscape(req.GetNetworkEdgeSecurityService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.networkEdgeSecurityServicesClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
