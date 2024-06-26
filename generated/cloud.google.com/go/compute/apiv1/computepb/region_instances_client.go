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
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
)

var newRegionInstancesClientHook clientHook

// RegionInstancesCallOptions contains the retry settings for each method of RegionInstancesClient.
type RegionInstancesCallOptions struct {
	BulkInsert []gax.CallOption
}

func defaultRegionInstancesGRPCClientOptions() []option.ClientOption {
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

func defaultRegionInstancesCallOptions() *RegionInstancesCallOptions {
	return &RegionInstancesCallOptions{
		BulkInsert: []gax.CallOption{
		},
	}
}

// internalRegionInstancesClient is an interface that defines the methods available from Google Compute Engine API.
type internalRegionInstancesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	BulkInsert(context.Context, *computepb.BulkInsertRegionInstanceRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// RegionInstancesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The RegionInstances API.
type RegionInstancesClient struct {
	// The internal transport-dependent client.
	internalClient internalRegionInstancesClient

	// The call options for this service.
	CallOptions *RegionInstancesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RegionInstancesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RegionInstancesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RegionInstancesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// BulkInsert creates multiple instances in a given region. Count specifies the number of instances to create.
func (c *RegionInstancesClient) BulkInsert(ctx context.Context, req *computepb.BulkInsertRegionInstanceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.BulkInsert(ctx, req, opts...)
}

// regionInstancesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type regionInstancesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing RegionInstancesClient
	CallOptions **RegionInstancesCallOptions

	// The gRPC API client.
	regionInstancesClient computepb.RegionInstancesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewRegionInstancesClient creates a new region instances client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The RegionInstances API.
func NewRegionInstancesClient(ctx context.Context, opts ...option.ClientOption) (*RegionInstancesClient, error) {
	clientOpts := defaultRegionInstancesGRPCClientOptions()
	if newRegionInstancesClientHook != nil {
		hookOpts, err := newRegionInstancesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RegionInstancesClient{CallOptions: defaultRegionInstancesCallOptions()}

	c := &regionInstancesGRPCClient{
		connPool:    connPool,
		regionInstancesClient: computepb.NewRegionInstancesClient(connPool),
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
func (c *regionInstancesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *regionInstancesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *regionInstancesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *regionInstancesGRPCClient) BulkInsert(ctx context.Context, req *computepb.BulkInsertRegionInstanceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).BulkInsert[0:len((*c.CallOptions).BulkInsert):len((*c.CallOptions).BulkInsert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionInstancesClient.BulkInsert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
