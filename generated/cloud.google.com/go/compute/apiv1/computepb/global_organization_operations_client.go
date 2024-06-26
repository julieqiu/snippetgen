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

var newGlobalOrganizationOperationsClientHook clientHook

// GlobalOrganizationOperationsCallOptions contains the retry settings for each method of GlobalOrganizationOperationsClient.
type GlobalOrganizationOperationsCallOptions struct {
	Delete []gax.CallOption
	Get []gax.CallOption
	List []gax.CallOption
}

func defaultGlobalOrganizationOperationsGRPCClientOptions() []option.ClientOption {
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

func defaultGlobalOrganizationOperationsCallOptions() *GlobalOrganizationOperationsCallOptions {
	return &GlobalOrganizationOperationsCallOptions{
		Delete: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
	}
}

// internalGlobalOrganizationOperationsClient is an interface that defines the methods available from Google Compute Engine API.
type internalGlobalOrganizationOperationsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Delete(context.Context, *computepb.DeleteGlobalOrganizationOperationRequest, ...gax.CallOption) (*computepb.DeleteGlobalOrganizationOperationResponse, error)
	Get(context.Context, *computepb.GetGlobalOrganizationOperationRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListGlobalOrganizationOperationsRequest, ...gax.CallOption) *OperationIterator
}

// GlobalOrganizationOperationsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The GlobalOrganizationOperations API.
type GlobalOrganizationOperationsClient struct {
	// The internal transport-dependent client.
	internalClient internalGlobalOrganizationOperationsClient

	// The call options for this service.
	CallOptions *GlobalOrganizationOperationsCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *GlobalOrganizationOperationsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *GlobalOrganizationOperationsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *GlobalOrganizationOperationsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Delete deletes the specified Operations resource.
func (c *GlobalOrganizationOperationsClient) Delete(ctx context.Context, req *computepb.DeleteGlobalOrganizationOperationRequest, opts ...gax.CallOption) (*computepb.DeleteGlobalOrganizationOperationResponse, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get retrieves the specified Operations resource. Gets a list of operations by making a list() request.
func (c *GlobalOrganizationOperationsClient) Get(ctx context.Context, req *computepb.GetGlobalOrganizationOperationRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// List retrieves a list of Operation resources contained within the specified organization.
func (c *GlobalOrganizationOperationsClient) List(ctx context.Context, req *computepb.ListGlobalOrganizationOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// globalOrganizationOperationsGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type globalOrganizationOperationsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing GlobalOrganizationOperationsClient
	CallOptions **GlobalOrganizationOperationsCallOptions

	// The gRPC API client.
	globalOrganizationOperationsClient computepb.GlobalOrganizationOperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewGlobalOrganizationOperationsClient creates a new global organization operations client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The GlobalOrganizationOperations API.
func NewGlobalOrganizationOperationsClient(ctx context.Context, opts ...option.ClientOption) (*GlobalOrganizationOperationsClient, error) {
	clientOpts := defaultGlobalOrganizationOperationsGRPCClientOptions()
	if newGlobalOrganizationOperationsClientHook != nil {
		hookOpts, err := newGlobalOrganizationOperationsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := GlobalOrganizationOperationsClient{CallOptions: defaultGlobalOrganizationOperationsCallOptions()}

	c := &globalOrganizationOperationsGRPCClient{
		connPool:    connPool,
		globalOrganizationOperationsClient: computepb.NewGlobalOrganizationOperationsClient(connPool),
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
func (c *globalOrganizationOperationsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *globalOrganizationOperationsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *globalOrganizationOperationsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *globalOrganizationOperationsGRPCClient) Delete(ctx context.Context, req *computepb.DeleteGlobalOrganizationOperationRequest, opts ...gax.CallOption) (*computepb.DeleteGlobalOrganizationOperationResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "operation", url.QueryEscape(req.GetOperation()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.DeleteGlobalOrganizationOperationResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.globalOrganizationOperationsClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *globalOrganizationOperationsGRPCClient) Get(ctx context.Context, req *computepb.GetGlobalOrganizationOperationRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "operation", url.QueryEscape(req.GetOperation()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.globalOrganizationOperationsClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *globalOrganizationOperationsGRPCClient) List(ctx context.Context, req *computepb.ListGlobalOrganizationOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*computepb.ListGlobalOrganizationOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.Operation, string, error) {
		resp := &computepb.OperationList{}
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
			resp, err = c.globalOrganizationOperationsClient.List(ctx, req, settings.GRPC...)
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
