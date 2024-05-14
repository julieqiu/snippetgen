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

var newRoutersClientHook clientHook

// RoutersCallOptions contains the retry settings for each method of RoutersClient.
type RoutersCallOptions struct {
	AggregatedList []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	GetNatIpInfo []gax.CallOption
	GetNatMappingInfo []gax.CallOption
	GetRouterStatus []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	Patch []gax.CallOption
	Preview []gax.CallOption
	Update []gax.CallOption
}

func defaultRoutersGRPCClientOptions() []option.ClientOption {
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

func defaultRoutersCallOptions() *RoutersCallOptions {
	return &RoutersCallOptions{
		AggregatedList: []gax.CallOption{
		},
		Delete: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		GetNatIpInfo: []gax.CallOption{
		},
		GetNatMappingInfo: []gax.CallOption{
		},
		GetRouterStatus: []gax.CallOption{
		},
		Insert: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
		Patch: []gax.CallOption{
		},
		Preview: []gax.CallOption{
		},
		Update: []gax.CallOption{
		},
	}
}

// internalRoutersClient is an interface that defines the methods available from Google Compute Engine API.
type internalRoutersClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListRoutersRequest, ...gax.CallOption) *RoutersScopedListPairIterator
	Delete(context.Context, *computepb.DeleteRouterRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetRouterRequest, ...gax.CallOption) (*computepb.Router, error)
	GetNatIpInfo(context.Context, *computepb.GetNatIpInfoRouterRequest, ...gax.CallOption) (*computepb.NatIpInfoResponse, error)
	GetNatMappingInfo(context.Context, *computepb.GetNatMappingInfoRoutersRequest, ...gax.CallOption) *VmEndpointNatMappingsIterator
	GetRouterStatus(context.Context, *computepb.GetRouterStatusRouterRequest, ...gax.CallOption) (*computepb.RouterStatusResponse, error)
	Insert(context.Context, *computepb.InsertRouterRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListRoutersRequest, ...gax.CallOption) *RouterIterator
	Patch(context.Context, *computepb.PatchRouterRequest, ...gax.CallOption) (*computepb.Operation, error)
	Preview(context.Context, *computepb.PreviewRouterRequest, ...gax.CallOption) (*computepb.RoutersPreviewResponse, error)
	Update(context.Context, *computepb.UpdateRouterRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// RoutersClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Routers API.
type RoutersClient struct {
	// The internal transport-dependent client.
	internalClient internalRoutersClient

	// The call options for this service.
	CallOptions *RoutersCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RoutersClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RoutersClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RoutersClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves an aggregated list of routers. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *RoutersClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListRoutersRequest, opts ...gax.CallOption) *RoutersScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified Router resource.
func (c *RoutersClient) Delete(ctx context.Context, req *computepb.DeleteRouterRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified Router resource.
func (c *RoutersClient) Get(ctx context.Context, req *computepb.GetRouterRequest, opts ...gax.CallOption) (*computepb.Router, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetNatIpInfo retrieves runtime NAT IP information.
func (c *RoutersClient) GetNatIpInfo(ctx context.Context, req *computepb.GetNatIpInfoRouterRequest, opts ...gax.CallOption) (*computepb.NatIpInfoResponse, error) {
	return c.internalClient.GetNatIpInfo(ctx, req, opts...)
}

// GetNatMappingInfo retrieves runtime Nat mapping information of VM endpoints.
func (c *RoutersClient) GetNatMappingInfo(ctx context.Context, req *computepb.GetNatMappingInfoRoutersRequest, opts ...gax.CallOption) *VmEndpointNatMappingsIterator {
	return c.internalClient.GetNatMappingInfo(ctx, req, opts...)
}

// GetRouterStatus retrieves runtime information of the specified router.
func (c *RoutersClient) GetRouterStatus(ctx context.Context, req *computepb.GetRouterStatusRouterRequest, opts ...gax.CallOption) (*computepb.RouterStatusResponse, error) {
	return c.internalClient.GetRouterStatus(ctx, req, opts...)
}

// Insert creates a Router resource in the specified project and region using the data included in the request.
func (c *RoutersClient) Insert(ctx context.Context, req *computepb.InsertRouterRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of Router resources available to the specified project.
func (c *RoutersClient) List(ctx context.Context, req *computepb.ListRoutersRequest, opts ...gax.CallOption) *RouterIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified Router resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *RoutersClient) Patch(ctx context.Context, req *computepb.PatchRouterRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// Preview preview fields auto-generated during router create and update operations. Calling this method does NOT create or update the router.
func (c *RoutersClient) Preview(ctx context.Context, req *computepb.PreviewRouterRequest, opts ...gax.CallOption) (*computepb.RoutersPreviewResponse, error) {
	return c.internalClient.Preview(ctx, req, opts...)
}

// Update updates the specified Router resource with the data included in the request. This method conforms to PUT semantics, which requests that the state of the target resource be created or replaced with the state defined by the representation enclosed in the request message payload.
func (c *RoutersClient) Update(ctx context.Context, req *computepb.UpdateRouterRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Update(ctx, req, opts...)
}

// routersGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type routersGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing RoutersClient
	CallOptions **RoutersCallOptions

	// The gRPC API client.
	routersClient computepb.RoutersClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewRoutersClient creates a new routers client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The Routers API.
func NewRoutersClient(ctx context.Context, opts ...option.ClientOption) (*RoutersClient, error) {
	clientOpts := defaultRoutersGRPCClientOptions()
	if newRoutersClientHook != nil {
		hookOpts, err := newRoutersClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RoutersClient{CallOptions: defaultRoutersCallOptions()}

	c := &routersGRPCClient{
		connPool:    connPool,
		routersClient: computepb.NewRoutersClient(connPool),
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
func (c *routersGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *routersGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *routersGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *routersGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListRoutersRequest, opts ...gax.CallOption) *RoutersScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &RoutersScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListRoutersRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]RoutersScopedListPair, string, error) {
		resp := &computepb.RouterAggregatedList{}
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
			resp, err = c.routersClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]RoutersScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, RoutersScopedListPair{k, v})
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

func (c *routersGRPCClient) Delete(ctx context.Context, req *computepb.DeleteRouterRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "router", url.QueryEscape(req.GetRouter()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.routersClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *routersGRPCClient) Get(ctx context.Context, req *computepb.GetRouterRequest, opts ...gax.CallOption) (*computepb.Router, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "router", url.QueryEscape(req.GetRouter()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.Router
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.routersClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *routersGRPCClient) GetNatIpInfo(ctx context.Context, req *computepb.GetNatIpInfoRouterRequest, opts ...gax.CallOption) (*computepb.NatIpInfoResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "router", url.QueryEscape(req.GetRouter()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetNatIpInfo[0:len((*c.CallOptions).GetNatIpInfo):len((*c.CallOptions).GetNatIpInfo)], opts...)
	var resp *computepb.NatIpInfoResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.routersClient.GetNatIpInfo(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *routersGRPCClient) GetNatMappingInfo(ctx context.Context, req *computepb.GetNatMappingInfoRoutersRequest, opts ...gax.CallOption) *VmEndpointNatMappingsIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "router", url.QueryEscape(req.GetRouter()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetNatMappingInfo[0:len((*c.CallOptions).GetNatMappingInfo):len((*c.CallOptions).GetNatMappingInfo)], opts...)
	it := &VmEndpointNatMappingsIterator{}
	req = proto.Clone(req).(*computepb.GetNatMappingInfoRoutersRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.VmEndpointNatMappings, string, error) {
		resp := &computepb.VmEndpointNatMappingsList{}
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
			resp, err = c.routersClient.GetNatMappingInfo(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetResult(), resp.GetNextPageToken(), nil
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

func (c *routersGRPCClient) GetRouterStatus(ctx context.Context, req *computepb.GetRouterStatusRouterRequest, opts ...gax.CallOption) (*computepb.RouterStatusResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "router", url.QueryEscape(req.GetRouter()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetRouterStatus[0:len((*c.CallOptions).GetRouterStatus):len((*c.CallOptions).GetRouterStatus)], opts...)
	var resp *computepb.RouterStatusResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.routersClient.GetRouterStatus(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *routersGRPCClient) Insert(ctx context.Context, req *computepb.InsertRouterRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.routersClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *routersGRPCClient) List(ctx context.Context, req *computepb.ListRoutersRequest, opts ...gax.CallOption) *RouterIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &RouterIterator{}
	req = proto.Clone(req).(*computepb.ListRoutersRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.Router, string, error) {
		resp := &computepb.RouterList{}
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
			resp, err = c.routersClient.List(ctx, req, settings.GRPC...)
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

func (c *routersGRPCClient) Patch(ctx context.Context, req *computepb.PatchRouterRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "router", url.QueryEscape(req.GetRouter()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.routersClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *routersGRPCClient) Preview(ctx context.Context, req *computepb.PreviewRouterRequest, opts ...gax.CallOption) (*computepb.RoutersPreviewResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "router", url.QueryEscape(req.GetRouter()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Preview[0:len((*c.CallOptions).Preview):len((*c.CallOptions).Preview)], opts...)
	var resp *computepb.RoutersPreviewResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.routersClient.Preview(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *routersGRPCClient) Update(ctx context.Context, req *computepb.UpdateRouterRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "router", url.QueryEscape(req.GetRouter()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Update[0:len((*c.CallOptions).Update):len((*c.CallOptions).Update)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.routersClient.Update(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
