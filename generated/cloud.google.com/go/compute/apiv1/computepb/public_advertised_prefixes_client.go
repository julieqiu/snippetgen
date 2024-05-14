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

var newPublicAdvertisedPrefixesClientHook clientHook

// PublicAdvertisedPrefixesCallOptions contains the retry settings for each method of PublicAdvertisedPrefixesClient.
type PublicAdvertisedPrefixesCallOptions struct {
	Announce []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	Patch []gax.CallOption
	Withdraw []gax.CallOption
}

func defaultPublicAdvertisedPrefixesGRPCClientOptions() []option.ClientOption {
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

func defaultPublicAdvertisedPrefixesCallOptions() *PublicAdvertisedPrefixesCallOptions {
	return &PublicAdvertisedPrefixesCallOptions{
		Announce: []gax.CallOption{
		},
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
		Withdraw: []gax.CallOption{
		},
	}
}

// internalPublicAdvertisedPrefixesClient is an interface that defines the methods available from Google Compute Engine API.
type internalPublicAdvertisedPrefixesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Announce(context.Context, *computepb.AnnouncePublicAdvertisedPrefixeRequest, ...gax.CallOption) (*computepb.Operation, error)
	Delete(context.Context, *computepb.DeletePublicAdvertisedPrefixeRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetPublicAdvertisedPrefixeRequest, ...gax.CallOption) (*computepb.PublicAdvertisedPrefix, error)
	Insert(context.Context, *computepb.InsertPublicAdvertisedPrefixeRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListPublicAdvertisedPrefixesRequest, ...gax.CallOption) *PublicAdvertisedPrefixIterator
	Patch(context.Context, *computepb.PatchPublicAdvertisedPrefixeRequest, ...gax.CallOption) (*computepb.Operation, error)
	Withdraw(context.Context, *computepb.WithdrawPublicAdvertisedPrefixeRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// PublicAdvertisedPrefixesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The PublicAdvertisedPrefixes API.
type PublicAdvertisedPrefixesClient struct {
	// The internal transport-dependent client.
	internalClient internalPublicAdvertisedPrefixesClient

	// The call options for this service.
	CallOptions *PublicAdvertisedPrefixesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PublicAdvertisedPrefixesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PublicAdvertisedPrefixesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PublicAdvertisedPrefixesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Announce announces the specified PublicAdvertisedPrefix
func (c *PublicAdvertisedPrefixesClient) Announce(ctx context.Context, req *computepb.AnnouncePublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Announce(ctx, req, opts...)
}

// Delete deletes the specified PublicAdvertisedPrefix
func (c *PublicAdvertisedPrefixesClient) Delete(ctx context.Context, req *computepb.DeletePublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified PublicAdvertisedPrefix resource.
func (c *PublicAdvertisedPrefixesClient) Get(ctx context.Context, req *computepb.GetPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.PublicAdvertisedPrefix, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a PublicAdvertisedPrefix in the specified project using the parameters that are included in the request.
func (c *PublicAdvertisedPrefixesClient) Insert(ctx context.Context, req *computepb.InsertPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List lists the PublicAdvertisedPrefixes for a project.
func (c *PublicAdvertisedPrefixesClient) List(ctx context.Context, req *computepb.ListPublicAdvertisedPrefixesRequest, opts ...gax.CallOption) *PublicAdvertisedPrefixIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified Router resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *PublicAdvertisedPrefixesClient) Patch(ctx context.Context, req *computepb.PatchPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// Withdraw withdraws the specified PublicAdvertisedPrefix
func (c *PublicAdvertisedPrefixesClient) Withdraw(ctx context.Context, req *computepb.WithdrawPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Withdraw(ctx, req, opts...)
}

// publicAdvertisedPrefixesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type publicAdvertisedPrefixesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing PublicAdvertisedPrefixesClient
	CallOptions **PublicAdvertisedPrefixesCallOptions

	// The gRPC API client.
	publicAdvertisedPrefixesClient computepb.PublicAdvertisedPrefixesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewPublicAdvertisedPrefixesClient creates a new public advertised prefixes client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The PublicAdvertisedPrefixes API.
func NewPublicAdvertisedPrefixesClient(ctx context.Context, opts ...option.ClientOption) (*PublicAdvertisedPrefixesClient, error) {
	clientOpts := defaultPublicAdvertisedPrefixesGRPCClientOptions()
	if newPublicAdvertisedPrefixesClientHook != nil {
		hookOpts, err := newPublicAdvertisedPrefixesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PublicAdvertisedPrefixesClient{CallOptions: defaultPublicAdvertisedPrefixesCallOptions()}

	c := &publicAdvertisedPrefixesGRPCClient{
		connPool:    connPool,
		publicAdvertisedPrefixesClient: computepb.NewPublicAdvertisedPrefixesClient(connPool),
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
func (c *publicAdvertisedPrefixesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *publicAdvertisedPrefixesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *publicAdvertisedPrefixesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *publicAdvertisedPrefixesGRPCClient) Announce(ctx context.Context, req *computepb.AnnouncePublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "public_advertised_prefix", url.QueryEscape(req.GetPublicAdvertisedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Announce[0:len((*c.CallOptions).Announce):len((*c.CallOptions).Announce)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publicAdvertisedPrefixesClient.Announce(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publicAdvertisedPrefixesGRPCClient) Delete(ctx context.Context, req *computepb.DeletePublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "public_advertised_prefix", url.QueryEscape(req.GetPublicAdvertisedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publicAdvertisedPrefixesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publicAdvertisedPrefixesGRPCClient) Get(ctx context.Context, req *computepb.GetPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.PublicAdvertisedPrefix, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "public_advertised_prefix", url.QueryEscape(req.GetPublicAdvertisedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.PublicAdvertisedPrefix
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publicAdvertisedPrefixesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publicAdvertisedPrefixesGRPCClient) Insert(ctx context.Context, req *computepb.InsertPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publicAdvertisedPrefixesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publicAdvertisedPrefixesGRPCClient) List(ctx context.Context, req *computepb.ListPublicAdvertisedPrefixesRequest, opts ...gax.CallOption) *PublicAdvertisedPrefixIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &PublicAdvertisedPrefixIterator{}
	req = proto.Clone(req).(*computepb.ListPublicAdvertisedPrefixesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.PublicAdvertisedPrefix, string, error) {
		resp := &computepb.PublicAdvertisedPrefixList{}
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
			resp, err = c.publicAdvertisedPrefixesClient.List(ctx, req, settings.GRPC...)
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

func (c *publicAdvertisedPrefixesGRPCClient) Patch(ctx context.Context, req *computepb.PatchPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "public_advertised_prefix", url.QueryEscape(req.GetPublicAdvertisedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publicAdvertisedPrefixesClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publicAdvertisedPrefixesGRPCClient) Withdraw(ctx context.Context, req *computepb.WithdrawPublicAdvertisedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "public_advertised_prefix", url.QueryEscape(req.GetPublicAdvertisedPrefix()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Withdraw[0:len((*c.CallOptions).Withdraw):len((*c.CallOptions).Withdraw)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publicAdvertisedPrefixesClient.Withdraw(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
