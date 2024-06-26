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

var newStoragePoolTypesClientHook clientHook

// StoragePoolTypesCallOptions contains the retry settings for each method of StoragePoolTypesClient.
type StoragePoolTypesCallOptions struct {
	AggregatedList []gax.CallOption
	Get []gax.CallOption
	List []gax.CallOption
}

func defaultStoragePoolTypesGRPCClientOptions() []option.ClientOption {
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

func defaultStoragePoolTypesCallOptions() *StoragePoolTypesCallOptions {
	return &StoragePoolTypesCallOptions{
		AggregatedList: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
	}
}

// internalStoragePoolTypesClient is an interface that defines the methods available from Google Compute Engine API.
type internalStoragePoolTypesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListStoragePoolTypesRequest, ...gax.CallOption) *StoragePoolTypesScopedListPairIterator
	Get(context.Context, *computepb.GetStoragePoolTypeRequest, ...gax.CallOption) (*computepb.StoragePoolType, error)
	List(context.Context, *computepb.ListStoragePoolTypesRequest, ...gax.CallOption) *StoragePoolTypeIterator
}

// StoragePoolTypesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The StoragePoolTypes API.
type StoragePoolTypesClient struct {
	// The internal transport-dependent client.
	internalClient internalStoragePoolTypesClient

	// The call options for this service.
	CallOptions *StoragePoolTypesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *StoragePoolTypesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *StoragePoolTypesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *StoragePoolTypesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves an aggregated list of storage pool types. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *StoragePoolTypesClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListStoragePoolTypesRequest, opts ...gax.CallOption) *StoragePoolTypesScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Get returns the specified storage pool type.
func (c *StoragePoolTypesClient) Get(ctx context.Context, req *computepb.GetStoragePoolTypeRequest, opts ...gax.CallOption) (*computepb.StoragePoolType, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// List retrieves a list of storage pool types available to the specified project.
func (c *StoragePoolTypesClient) List(ctx context.Context, req *computepb.ListStoragePoolTypesRequest, opts ...gax.CallOption) *StoragePoolTypeIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// storagePoolTypesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type storagePoolTypesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing StoragePoolTypesClient
	CallOptions **StoragePoolTypesCallOptions

	// The gRPC API client.
	storagePoolTypesClient computepb.StoragePoolTypesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewStoragePoolTypesClient creates a new storage pool types client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The StoragePoolTypes API.
func NewStoragePoolTypesClient(ctx context.Context, opts ...option.ClientOption) (*StoragePoolTypesClient, error) {
	clientOpts := defaultStoragePoolTypesGRPCClientOptions()
	if newStoragePoolTypesClientHook != nil {
		hookOpts, err := newStoragePoolTypesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := StoragePoolTypesClient{CallOptions: defaultStoragePoolTypesCallOptions()}

	c := &storagePoolTypesGRPCClient{
		connPool:    connPool,
		storagePoolTypesClient: computepb.NewStoragePoolTypesClient(connPool),
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
func (c *storagePoolTypesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *storagePoolTypesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *storagePoolTypesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *storagePoolTypesGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListStoragePoolTypesRequest, opts ...gax.CallOption) *StoragePoolTypesScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &StoragePoolTypesScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListStoragePoolTypesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]StoragePoolTypesScopedListPair, string, error) {
		resp := &computepb.StoragePoolTypeAggregatedList{}
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
			resp, err = c.storagePoolTypesClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]StoragePoolTypesScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, StoragePoolTypesScopedListPair{k, v})
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

func (c *storagePoolTypesGRPCClient) Get(ctx context.Context, req *computepb.GetStoragePoolTypeRequest, opts ...gax.CallOption) (*computepb.StoragePoolType, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "storage_pool_type", url.QueryEscape(req.GetStoragePoolType()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.StoragePoolType
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.storagePoolTypesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *storagePoolTypesGRPCClient) List(ctx context.Context, req *computepb.ListStoragePoolTypesRequest, opts ...gax.CallOption) *StoragePoolTypeIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &StoragePoolTypeIterator{}
	req = proto.Clone(req).(*computepb.ListStoragePoolTypesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.StoragePoolType, string, error) {
		resp := &computepb.StoragePoolTypeList{}
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
			resp, err = c.storagePoolTypesClient.List(ctx, req, settings.GRPC...)
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
