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

var newRegionHealthCheckServicesClientHook clientHook

// RegionHealthCheckServicesCallOptions contains the retry settings for each method of RegionHealthCheckServicesClient.
type RegionHealthCheckServicesCallOptions struct {
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	Patch []gax.CallOption
}

func defaultRegionHealthCheckServicesGRPCClientOptions() []option.ClientOption {
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

func defaultRegionHealthCheckServicesCallOptions() *RegionHealthCheckServicesCallOptions {
	return &RegionHealthCheckServicesCallOptions{
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
	}
}

// internalRegionHealthCheckServicesClient is an interface that defines the methods available from Google Compute Engine API.
type internalRegionHealthCheckServicesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Delete(context.Context, *computepb.DeleteRegionHealthCheckServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetRegionHealthCheckServiceRequest, ...gax.CallOption) (*computepb.HealthCheckService, error)
	Insert(context.Context, *computepb.InsertRegionHealthCheckServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListRegionHealthCheckServicesRequest, ...gax.CallOption) *HealthCheckServiceIterator
	Patch(context.Context, *computepb.PatchRegionHealthCheckServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// RegionHealthCheckServicesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The RegionHealthCheckServices API.
type RegionHealthCheckServicesClient struct {
	// The internal transport-dependent client.
	internalClient internalRegionHealthCheckServicesClient

	// The call options for this service.
	CallOptions *RegionHealthCheckServicesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RegionHealthCheckServicesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RegionHealthCheckServicesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RegionHealthCheckServicesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Delete deletes the specified regional HealthCheckService.
func (c *RegionHealthCheckServicesClient) Delete(ctx context.Context, req *computepb.DeleteRegionHealthCheckServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified regional HealthCheckService resource.
func (c *RegionHealthCheckServicesClient) Get(ctx context.Context, req *computepb.GetRegionHealthCheckServiceRequest, opts ...gax.CallOption) (*computepb.HealthCheckService, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a regional HealthCheckService resource in the specified project and region using the data included in the request.
func (c *RegionHealthCheckServicesClient) Insert(ctx context.Context, req *computepb.InsertRegionHealthCheckServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List lists all the HealthCheckService resources that have been configured for the specified project in the given region.
func (c *RegionHealthCheckServicesClient) List(ctx context.Context, req *computepb.ListRegionHealthCheckServicesRequest, opts ...gax.CallOption) *HealthCheckServiceIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch updates the specified regional HealthCheckService resource with the data included in the request. This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (c *RegionHealthCheckServicesClient) Patch(ctx context.Context, req *computepb.PatchRegionHealthCheckServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// regionHealthCheckServicesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type regionHealthCheckServicesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing RegionHealthCheckServicesClient
	CallOptions **RegionHealthCheckServicesCallOptions

	// The gRPC API client.
	regionHealthCheckServicesClient computepb.RegionHealthCheckServicesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewRegionHealthCheckServicesClient creates a new region health check services client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The RegionHealthCheckServices API.
func NewRegionHealthCheckServicesClient(ctx context.Context, opts ...option.ClientOption) (*RegionHealthCheckServicesClient, error) {
	clientOpts := defaultRegionHealthCheckServicesGRPCClientOptions()
	if newRegionHealthCheckServicesClientHook != nil {
		hookOpts, err := newRegionHealthCheckServicesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RegionHealthCheckServicesClient{CallOptions: defaultRegionHealthCheckServicesCallOptions()}

	c := &regionHealthCheckServicesGRPCClient{
		connPool:    connPool,
		regionHealthCheckServicesClient: computepb.NewRegionHealthCheckServicesClient(connPool),
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
func (c *regionHealthCheckServicesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *regionHealthCheckServicesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *regionHealthCheckServicesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *regionHealthCheckServicesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteRegionHealthCheckServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "health_check_service", url.QueryEscape(req.GetHealthCheckService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionHealthCheckServicesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionHealthCheckServicesGRPCClient) Get(ctx context.Context, req *computepb.GetRegionHealthCheckServiceRequest, opts ...gax.CallOption) (*computepb.HealthCheckService, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "health_check_service", url.QueryEscape(req.GetHealthCheckService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.HealthCheckService
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionHealthCheckServicesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionHealthCheckServicesGRPCClient) Insert(ctx context.Context, req *computepb.InsertRegionHealthCheckServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionHealthCheckServicesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionHealthCheckServicesGRPCClient) List(ctx context.Context, req *computepb.ListRegionHealthCheckServicesRequest, opts ...gax.CallOption) *HealthCheckServiceIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &HealthCheckServiceIterator{}
	req = proto.Clone(req).(*computepb.ListRegionHealthCheckServicesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.HealthCheckService, string, error) {
		resp := &computepb.HealthCheckServicesList{}
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
			resp, err = c.regionHealthCheckServicesClient.List(ctx, req, settings.GRPC...)
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

func (c *regionHealthCheckServicesGRPCClient) Patch(ctx context.Context, req *computepb.PatchRegionHealthCheckServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "health_check_service", url.QueryEscape(req.GetHealthCheckService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionHealthCheckServicesClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
