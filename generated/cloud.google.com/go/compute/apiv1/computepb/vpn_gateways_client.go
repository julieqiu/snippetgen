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

var newVpnGatewaysClientHook clientHook

// VpnGatewaysCallOptions contains the retry settings for each method of VpnGatewaysClient.
type VpnGatewaysCallOptions struct {
	AggregatedList []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	GetStatus []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	SetLabels []gax.CallOption
	TestIamPermissions []gax.CallOption
}

func defaultVpnGatewaysGRPCClientOptions() []option.ClientOption {
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

func defaultVpnGatewaysCallOptions() *VpnGatewaysCallOptions {
	return &VpnGatewaysCallOptions{
		AggregatedList: []gax.CallOption{
		},
		Delete: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		GetStatus: []gax.CallOption{
		},
		Insert: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
		SetLabels: []gax.CallOption{
		},
		TestIamPermissions: []gax.CallOption{
		},
	}
}

// internalVpnGatewaysClient is an interface that defines the methods available from Google Compute Engine API.
type internalVpnGatewaysClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListVpnGatewaysRequest, ...gax.CallOption) *VpnGatewaysScopedListPairIterator
	Delete(context.Context, *computepb.DeleteVpnGatewayRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetVpnGatewayRequest, ...gax.CallOption) (*computepb.VpnGateway, error)
	GetStatus(context.Context, *computepb.GetStatusVpnGatewayRequest, ...gax.CallOption) (*computepb.VpnGatewaysGetStatusResponse, error)
	Insert(context.Context, *computepb.InsertVpnGatewayRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListVpnGatewaysRequest, ...gax.CallOption) *VpnGatewayIterator
	SetLabels(context.Context, *computepb.SetLabelsVpnGatewayRequest, ...gax.CallOption) (*computepb.Operation, error)
	TestIamPermissions(context.Context, *computepb.TestIamPermissionsVpnGatewayRequest, ...gax.CallOption) (*computepb.TestPermissionsResponse, error)
}

// VpnGatewaysClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The VpnGateways API.
type VpnGatewaysClient struct {
	// The internal transport-dependent client.
	internalClient internalVpnGatewaysClient

	// The call options for this service.
	CallOptions *VpnGatewaysCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *VpnGatewaysClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *VpnGatewaysClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *VpnGatewaysClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves an aggregated list of VPN gateways. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *VpnGatewaysClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListVpnGatewaysRequest, opts ...gax.CallOption) *VpnGatewaysScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified VPN gateway.
func (c *VpnGatewaysClient) Delete(ctx context.Context, req *computepb.DeleteVpnGatewayRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified VPN gateway.
func (c *VpnGatewaysClient) Get(ctx context.Context, req *computepb.GetVpnGatewayRequest, opts ...gax.CallOption) (*computepb.VpnGateway, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetStatus returns the status for the specified VPN gateway.
func (c *VpnGatewaysClient) GetStatus(ctx context.Context, req *computepb.GetStatusVpnGatewayRequest, opts ...gax.CallOption) (*computepb.VpnGatewaysGetStatusResponse, error) {
	return c.internalClient.GetStatus(ctx, req, opts...)
}

// Insert creates a VPN gateway in the specified project and region using the data included in the request.
func (c *VpnGatewaysClient) Insert(ctx context.Context, req *computepb.InsertVpnGatewayRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of VPN gateways available to the specified project and region.
func (c *VpnGatewaysClient) List(ctx context.Context, req *computepb.ListVpnGatewaysRequest, opts ...gax.CallOption) *VpnGatewayIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// SetLabels sets the labels on a VpnGateway. To learn more about labels, read the Labeling Resources documentation.
func (c *VpnGatewaysClient) SetLabels(ctx context.Context, req *computepb.SetLabelsVpnGatewayRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetLabels(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *VpnGatewaysClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsVpnGatewayRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// vpnGatewaysGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type vpnGatewaysGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing VpnGatewaysClient
	CallOptions **VpnGatewaysCallOptions

	// The gRPC API client.
	vpnGatewaysClient computepb.VpnGatewaysClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewVpnGatewaysClient creates a new vpn gateways client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The VpnGateways API.
func NewVpnGatewaysClient(ctx context.Context, opts ...option.ClientOption) (*VpnGatewaysClient, error) {
	clientOpts := defaultVpnGatewaysGRPCClientOptions()
	if newVpnGatewaysClientHook != nil {
		hookOpts, err := newVpnGatewaysClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := VpnGatewaysClient{CallOptions: defaultVpnGatewaysCallOptions()}

	c := &vpnGatewaysGRPCClient{
		connPool:    connPool,
		vpnGatewaysClient: computepb.NewVpnGatewaysClient(connPool),
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
func (c *vpnGatewaysGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *vpnGatewaysGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *vpnGatewaysGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *vpnGatewaysGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListVpnGatewaysRequest, opts ...gax.CallOption) *VpnGatewaysScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &VpnGatewaysScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListVpnGatewaysRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]VpnGatewaysScopedListPair, string, error) {
		resp := &computepb.VpnGatewayAggregatedList{}
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
			resp, err = c.vpnGatewaysClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]VpnGatewaysScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, VpnGatewaysScopedListPair{k, v})
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

func (c *vpnGatewaysGRPCClient) Delete(ctx context.Context, req *computepb.DeleteVpnGatewayRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "vpn_gateway", url.QueryEscape(req.GetVpnGateway()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vpnGatewaysClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vpnGatewaysGRPCClient) Get(ctx context.Context, req *computepb.GetVpnGatewayRequest, opts ...gax.CallOption) (*computepb.VpnGateway, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "vpn_gateway", url.QueryEscape(req.GetVpnGateway()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.VpnGateway
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vpnGatewaysClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vpnGatewaysGRPCClient) GetStatus(ctx context.Context, req *computepb.GetStatusVpnGatewayRequest, opts ...gax.CallOption) (*computepb.VpnGatewaysGetStatusResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "vpn_gateway", url.QueryEscape(req.GetVpnGateway()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetStatus[0:len((*c.CallOptions).GetStatus):len((*c.CallOptions).GetStatus)], opts...)
	var resp *computepb.VpnGatewaysGetStatusResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vpnGatewaysClient.GetStatus(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vpnGatewaysGRPCClient) Insert(ctx context.Context, req *computepb.InsertVpnGatewayRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vpnGatewaysClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vpnGatewaysGRPCClient) List(ctx context.Context, req *computepb.ListVpnGatewaysRequest, opts ...gax.CallOption) *VpnGatewayIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &VpnGatewayIterator{}
	req = proto.Clone(req).(*computepb.ListVpnGatewaysRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.VpnGateway, string, error) {
		resp := &computepb.VpnGatewayList{}
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
			resp, err = c.vpnGatewaysClient.List(ctx, req, settings.GRPC...)
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

func (c *vpnGatewaysGRPCClient) SetLabels(ctx context.Context, req *computepb.SetLabelsVpnGatewayRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetLabels[0:len((*c.CallOptions).SetLabels):len((*c.CallOptions).SetLabels)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vpnGatewaysClient.SetLabels(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *vpnGatewaysGRPCClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsVpnGatewayRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *computepb.TestPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.vpnGatewaysClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
