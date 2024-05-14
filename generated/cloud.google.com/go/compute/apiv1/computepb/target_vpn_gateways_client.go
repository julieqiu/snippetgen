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

var newTargetVpnGatewaysClientHook clientHook

// TargetVpnGatewaysCallOptions contains the retry settings for each method of TargetVpnGatewaysClient.
type TargetVpnGatewaysCallOptions struct {
	AggregatedList []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	SetLabels []gax.CallOption
}

func defaultTargetVpnGatewaysGRPCClientOptions() []option.ClientOption {
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

func defaultTargetVpnGatewaysCallOptions() *TargetVpnGatewaysCallOptions {
	return &TargetVpnGatewaysCallOptions{
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
		SetLabels: []gax.CallOption{
		},
	}
}

// internalTargetVpnGatewaysClient is an interface that defines the methods available from Google Compute Engine API.
type internalTargetVpnGatewaysClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListTargetVpnGatewaysRequest, ...gax.CallOption) *TargetVpnGatewaysScopedListPairIterator
	Delete(context.Context, *computepb.DeleteTargetVpnGatewayRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetTargetVpnGatewayRequest, ...gax.CallOption) (*computepb.TargetVpnGateway, error)
	Insert(context.Context, *computepb.InsertTargetVpnGatewayRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListTargetVpnGatewaysRequest, ...gax.CallOption) *TargetVpnGatewayIterator
	SetLabels(context.Context, *computepb.SetLabelsTargetVpnGatewayRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// TargetVpnGatewaysClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The TargetVpnGateways API.
type TargetVpnGatewaysClient struct {
	// The internal transport-dependent client.
	internalClient internalTargetVpnGatewaysClient

	// The call options for this service.
	CallOptions *TargetVpnGatewaysCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TargetVpnGatewaysClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TargetVpnGatewaysClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TargetVpnGatewaysClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves an aggregated list of target VPN gateways. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *TargetVpnGatewaysClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetVpnGatewaysRequest, opts ...gax.CallOption) *TargetVpnGatewaysScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified target VPN gateway.
func (c *TargetVpnGatewaysClient) Delete(ctx context.Context, req *computepb.DeleteTargetVpnGatewayRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified target VPN gateway.
func (c *TargetVpnGatewaysClient) Get(ctx context.Context, req *computepb.GetTargetVpnGatewayRequest, opts ...gax.CallOption) (*computepb.TargetVpnGateway, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a target VPN gateway in the specified project and region using the data included in the request.
func (c *TargetVpnGatewaysClient) Insert(ctx context.Context, req *computepb.InsertTargetVpnGatewayRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of target VPN gateways available to the specified project and region.
func (c *TargetVpnGatewaysClient) List(ctx context.Context, req *computepb.ListTargetVpnGatewaysRequest, opts ...gax.CallOption) *TargetVpnGatewayIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// SetLabels sets the labels on a TargetVpnGateway. To learn more about labels, read the Labeling Resources documentation.
func (c *TargetVpnGatewaysClient) SetLabels(ctx context.Context, req *computepb.SetLabelsTargetVpnGatewayRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetLabels(ctx, req, opts...)
}

// targetVpnGatewaysGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type targetVpnGatewaysGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TargetVpnGatewaysClient
	CallOptions **TargetVpnGatewaysCallOptions

	// The gRPC API client.
	targetVpnGatewaysClient computepb.TargetVpnGatewaysClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewTargetVpnGatewaysClient creates a new target vpn gateways client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The TargetVpnGateways API.
func NewTargetVpnGatewaysClient(ctx context.Context, opts ...option.ClientOption) (*TargetVpnGatewaysClient, error) {
	clientOpts := defaultTargetVpnGatewaysGRPCClientOptions()
	if newTargetVpnGatewaysClientHook != nil {
		hookOpts, err := newTargetVpnGatewaysClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TargetVpnGatewaysClient{CallOptions: defaultTargetVpnGatewaysCallOptions()}

	c := &targetVpnGatewaysGRPCClient{
		connPool:    connPool,
		targetVpnGatewaysClient: computepb.NewTargetVpnGatewaysClient(connPool),
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
func (c *targetVpnGatewaysGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *targetVpnGatewaysGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *targetVpnGatewaysGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *targetVpnGatewaysGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetVpnGatewaysRequest, opts ...gax.CallOption) *TargetVpnGatewaysScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &TargetVpnGatewaysScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListTargetVpnGatewaysRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]TargetVpnGatewaysScopedListPair, string, error) {
		resp := &computepb.TargetVpnGatewayAggregatedList{}
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
			resp, err = c.targetVpnGatewaysClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]TargetVpnGatewaysScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, TargetVpnGatewaysScopedListPair{k, v})
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

func (c *targetVpnGatewaysGRPCClient) Delete(ctx context.Context, req *computepb.DeleteTargetVpnGatewayRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_vpn_gateway", url.QueryEscape(req.GetTargetVpnGateway()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetVpnGatewaysClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetVpnGatewaysGRPCClient) Get(ctx context.Context, req *computepb.GetTargetVpnGatewayRequest, opts ...gax.CallOption) (*computepb.TargetVpnGateway, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_vpn_gateway", url.QueryEscape(req.GetTargetVpnGateway()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.TargetVpnGateway
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetVpnGatewaysClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetVpnGatewaysGRPCClient) Insert(ctx context.Context, req *computepb.InsertTargetVpnGatewayRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetVpnGatewaysClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetVpnGatewaysGRPCClient) List(ctx context.Context, req *computepb.ListTargetVpnGatewaysRequest, opts ...gax.CallOption) *TargetVpnGatewayIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &TargetVpnGatewayIterator{}
	req = proto.Clone(req).(*computepb.ListTargetVpnGatewaysRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.TargetVpnGateway, string, error) {
		resp := &computepb.TargetVpnGatewayList{}
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
			resp, err = c.targetVpnGatewaysClient.List(ctx, req, settings.GRPC...)
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

func (c *targetVpnGatewaysGRPCClient) SetLabels(ctx context.Context, req *computepb.SetLabelsTargetVpnGatewayRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetLabels[0:len((*c.CallOptions).SetLabels):len((*c.CallOptions).SetLabels)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetVpnGatewaysClient.SetLabels(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
