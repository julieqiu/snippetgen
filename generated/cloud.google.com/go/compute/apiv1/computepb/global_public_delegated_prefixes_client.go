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

var newGlobalPublicDelegatedPrefixesClientHook clientHook

// GlobalPublicDelegatedPrefixesCallOptions contains the retry settings for each method of GlobalPublicDelegatedPrefixesClient.
type GlobalPublicDelegatedPrefixesCallOptions struct {
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	Patch []gax.CallOption
}

func defaultGlobalPublicDelegatedPrefixesGRPCClientOptions() []option.ClientOption {
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

func defaultGlobalPublicDelegatedPrefixesCallOptions() *GlobalPublicDelegatedPrefixesCallOptions {
	return &GlobalPublicDelegatedPrefixesCallOptions{
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

// internalGlobalPublicDelegatedPrefixesClient is an interface that defines the methods available from Google Compute Engine API.
type internalGlobalPublicDelegatedPrefixesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Delete(context.Context, *computepb.DeleteGlobalPublicDelegatedPrefixeRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetGlobalPublicDelegatedPrefixeRequest, ...gax.CallOption) (*computepb.PublicDelegatedPrefix, error)
	Insert(context.Context, *computepb.InsertGlobalPublicDelegatedPrefixeRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListGlobalPublicDelegatedPrefixesRequest, ...gax.CallOption) *PublicDelegatedPrefixIterator
	Patch(context.Context, *computepb.PatchGlobalPublicDelegatedPrefixeRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// GlobalPublicDelegatedPrefixesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The GlobalPublicDelegatedPrefixes API.
type GlobalPublicDelegatedPrefixesClient struct {
	// The internal transport-dependent client.
	internalClient internalGlobalPublicDelegatedPrefixesClient

	// The call options for this service.
	CallOptions *GlobalPublicDelegatedPrefixesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *GlobalPublicDelegatedPrefixesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *GlobalPublicDelegatedPrefixesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *GlobalPublicDelegatedPrefixesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Delete deletes the specified global PublicDelegatedPrefix.
func (c *GlobalPublicDelegatedPrefixesClient) Delete(ctx context.Context, req *computepb.DeleteGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified global PublicDelegatedPrefix resource.
func (c *GlobalPublicDelegatedPrefixesClient) Get(ctx context.Context, req *computepb.GetGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.PublicDelegatedPrefix, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a global PublicDelegatedPrefix in the specified project using the parameters that are included in the request.
func (c *GlobalPublicDelegatedPrefixesClient) Insert(ctx context.Context, req *computepb.InsertGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List lists the global PublicDelegatedPrefixes for a project.
func (c *GlobalPublicDelegatedPrefixesClient) List(ctx context.Context, req *computepb.ListGlobalPublicDelegatedPrefixesRequest, opts ...gax.CallOption) *PublicDelegatedPrefixIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified global PublicDelegatedPrefix resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *GlobalPublicDelegatedPrefixesClient) Patch(ctx context.Context, req *computepb.PatchGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// globalPublicDelegatedPrefixesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type globalPublicDelegatedPrefixesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing GlobalPublicDelegatedPrefixesClient
	CallOptions **GlobalPublicDelegatedPrefixesCallOptions

	// The gRPC API client.
	globalPublicDelegatedPrefixesClient computepb.GlobalPublicDelegatedPrefixesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewGlobalPublicDelegatedPrefixesClient creates a new global public delegated prefixes client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The GlobalPublicDelegatedPrefixes API.
func NewGlobalPublicDelegatedPrefixesClient(ctx context.Context, opts ...option.ClientOption) (*GlobalPublicDelegatedPrefixesClient, error) {
	clientOpts := defaultGlobalPublicDelegatedPrefixesGRPCClientOptions()
	if newGlobalPublicDelegatedPrefixesClientHook != nil {
		hookOpts, err := newGlobalPublicDelegatedPrefixesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := GlobalPublicDelegatedPrefixesClient{CallOptions: defaultGlobalPublicDelegatedPrefixesCallOptions()}

	c := &globalPublicDelegatedPrefixesGRPCClient{
		connPool:    connPool,
		globalPublicDelegatedPrefixesClient: computepb.NewGlobalPublicDelegatedPrefixesClient(connPool),
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
func (c *globalPublicDelegatedPrefixesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *globalPublicDelegatedPrefixesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *globalPublicDelegatedPrefixesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *globalPublicDelegatedPrefixesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "public_delegated_prefix", url.QueryEscape(req.GetPublicDelegatedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.globalPublicDelegatedPrefixesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *globalPublicDelegatedPrefixesGRPCClient) Get(ctx context.Context, req *computepb.GetGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.PublicDelegatedPrefix, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "public_delegated_prefix", url.QueryEscape(req.GetPublicDelegatedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.PublicDelegatedPrefix
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.globalPublicDelegatedPrefixesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *globalPublicDelegatedPrefixesGRPCClient) Insert(ctx context.Context, req *computepb.InsertGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.globalPublicDelegatedPrefixesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *globalPublicDelegatedPrefixesGRPCClient) List(ctx context.Context, req *computepb.ListGlobalPublicDelegatedPrefixesRequest, opts ...gax.CallOption) *PublicDelegatedPrefixIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &PublicDelegatedPrefixIterator{}
	req = proto.Clone(req).(*computepb.ListGlobalPublicDelegatedPrefixesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.PublicDelegatedPrefix, string, error) {
		resp := &computepb.PublicDelegatedPrefixList{}
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
			resp, err = c.globalPublicDelegatedPrefixesClient.List(ctx, req, settings.GRPC...)
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

func (c *globalPublicDelegatedPrefixesGRPCClient) Patch(ctx context.Context, req *computepb.PatchGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "public_delegated_prefix", url.QueryEscape(req.GetPublicDelegatedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.globalPublicDelegatedPrefixesClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}