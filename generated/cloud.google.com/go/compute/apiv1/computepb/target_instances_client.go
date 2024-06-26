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

var newTargetInstancesClientHook clientHook

// TargetInstancesCallOptions contains the retry settings for each method of TargetInstancesClient.
type TargetInstancesCallOptions struct {
	AggregatedList []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	SetSecurityPolicy []gax.CallOption
}

func defaultTargetInstancesGRPCClientOptions() []option.ClientOption {
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

func defaultTargetInstancesCallOptions() *TargetInstancesCallOptions {
	return &TargetInstancesCallOptions{
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
		SetSecurityPolicy: []gax.CallOption{
		},
	}
}

// internalTargetInstancesClient is an interface that defines the methods available from Google Compute Engine API.
type internalTargetInstancesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListTargetInstancesRequest, ...gax.CallOption) *TargetInstancesScopedListPairIterator
	Delete(context.Context, *computepb.DeleteTargetInstanceRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetTargetInstanceRequest, ...gax.CallOption) (*computepb.TargetInstance, error)
	Insert(context.Context, *computepb.InsertTargetInstanceRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListTargetInstancesRequest, ...gax.CallOption) *TargetInstanceIterator
	SetSecurityPolicy(context.Context, *computepb.SetSecurityPolicyTargetInstanceRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// TargetInstancesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The TargetInstances API.
type TargetInstancesClient struct {
	// The internal transport-dependent client.
	internalClient internalTargetInstancesClient

	// The call options for this service.
	CallOptions *TargetInstancesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TargetInstancesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TargetInstancesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TargetInstancesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves an aggregated list of target instances. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *TargetInstancesClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetInstancesRequest, opts ...gax.CallOption) *TargetInstancesScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified TargetInstance resource.
func (c *TargetInstancesClient) Delete(ctx context.Context, req *computepb.DeleteTargetInstanceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified TargetInstance resource.
func (c *TargetInstancesClient) Get(ctx context.Context, req *computepb.GetTargetInstanceRequest, opts ...gax.CallOption) (*computepb.TargetInstance, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a TargetInstance resource in the specified project and zone using the data included in the request.
func (c *TargetInstancesClient) Insert(ctx context.Context, req *computepb.InsertTargetInstanceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of TargetInstance resources available to the specified project and zone.
func (c *TargetInstancesClient) List(ctx context.Context, req *computepb.ListTargetInstancesRequest, opts ...gax.CallOption) *TargetInstanceIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// SetSecurityPolicy sets the Google Cloud Armor security policy for the specified target instance. For more information, see Google Cloud Armor Overview
func (c *TargetInstancesClient) SetSecurityPolicy(ctx context.Context, req *computepb.SetSecurityPolicyTargetInstanceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetSecurityPolicy(ctx, req, opts...)
}

// targetInstancesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type targetInstancesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TargetInstancesClient
	CallOptions **TargetInstancesCallOptions

	// The gRPC API client.
	targetInstancesClient computepb.TargetInstancesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewTargetInstancesClient creates a new target instances client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The TargetInstances API.
func NewTargetInstancesClient(ctx context.Context, opts ...option.ClientOption) (*TargetInstancesClient, error) {
	clientOpts := defaultTargetInstancesGRPCClientOptions()
	if newTargetInstancesClientHook != nil {
		hookOpts, err := newTargetInstancesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TargetInstancesClient{CallOptions: defaultTargetInstancesCallOptions()}

	c := &targetInstancesGRPCClient{
		connPool:    connPool,
		targetInstancesClient: computepb.NewTargetInstancesClient(connPool),
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
func (c *targetInstancesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *targetInstancesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *targetInstancesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *targetInstancesGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetInstancesRequest, opts ...gax.CallOption) *TargetInstancesScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &TargetInstancesScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListTargetInstancesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]TargetInstancesScopedListPair, string, error) {
		resp := &computepb.TargetInstanceAggregatedList{}
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
			resp, err = c.targetInstancesClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]TargetInstancesScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, TargetInstancesScopedListPair{k, v})
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

func (c *targetInstancesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteTargetInstanceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "target_instance", url.QueryEscape(req.GetTargetInstance()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetInstancesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetInstancesGRPCClient) Get(ctx context.Context, req *computepb.GetTargetInstanceRequest, opts ...gax.CallOption) (*computepb.TargetInstance, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "target_instance", url.QueryEscape(req.GetTargetInstance()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.TargetInstance
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetInstancesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetInstancesGRPCClient) Insert(ctx context.Context, req *computepb.InsertTargetInstanceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetInstancesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetInstancesGRPCClient) List(ctx context.Context, req *computepb.ListTargetInstancesRequest, opts ...gax.CallOption) *TargetInstanceIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &TargetInstanceIterator{}
	req = proto.Clone(req).(*computepb.ListTargetInstancesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.TargetInstance, string, error) {
		resp := &computepb.TargetInstanceList{}
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
			resp, err = c.targetInstancesClient.List(ctx, req, settings.GRPC...)
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

func (c *targetInstancesGRPCClient) SetSecurityPolicy(ctx context.Context, req *computepb.SetSecurityPolicyTargetInstanceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "target_instance", url.QueryEscape(req.GetTargetInstance()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetSecurityPolicy[0:len((*c.CallOptions).SetSecurityPolicy):len((*c.CallOptions).SetSecurityPolicy)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetInstancesClient.SetSecurityPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
