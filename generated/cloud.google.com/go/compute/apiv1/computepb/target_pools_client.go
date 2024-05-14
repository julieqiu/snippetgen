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

var newTargetPoolsClientHook clientHook

// TargetPoolsCallOptions contains the retry settings for each method of TargetPoolsClient.
type TargetPoolsCallOptions struct {
	AddHealthCheck []gax.CallOption
	AddInstance []gax.CallOption
	AggregatedList []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	GetHealth []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	RemoveHealthCheck []gax.CallOption
	RemoveInstance []gax.CallOption
	SetBackup []gax.CallOption
	SetSecurityPolicy []gax.CallOption
}

func defaultTargetPoolsGRPCClientOptions() []option.ClientOption {
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

func defaultTargetPoolsCallOptions() *TargetPoolsCallOptions {
	return &TargetPoolsCallOptions{
		AddHealthCheck: []gax.CallOption{
		},
		AddInstance: []gax.CallOption{
		},
		AggregatedList: []gax.CallOption{
		},
		Delete: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		GetHealth: []gax.CallOption{
		},
		Insert: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
		RemoveHealthCheck: []gax.CallOption{
		},
		RemoveInstance: []gax.CallOption{
		},
		SetBackup: []gax.CallOption{
		},
		SetSecurityPolicy: []gax.CallOption{
		},
	}
}

// internalTargetPoolsClient is an interface that defines the methods available from Google Compute Engine API.
type internalTargetPoolsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddHealthCheck(context.Context, *computepb.AddHealthCheckTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
	AddInstance(context.Context, *computepb.AddInstanceTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
	AggregatedList(context.Context, *computepb.AggregatedListTargetPoolsRequest, ...gax.CallOption) *TargetPoolsScopedListPairIterator
	Delete(context.Context, *computepb.DeleteTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetTargetPoolRequest, ...gax.CallOption) (*computepb.TargetPool, error)
	GetHealth(context.Context, *computepb.GetHealthTargetPoolRequest, ...gax.CallOption) (*computepb.TargetPoolInstanceHealth, error)
	Insert(context.Context, *computepb.InsertTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListTargetPoolsRequest, ...gax.CallOption) *TargetPoolIterator
	RemoveHealthCheck(context.Context, *computepb.RemoveHealthCheckTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
	RemoveInstance(context.Context, *computepb.RemoveInstanceTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetBackup(context.Context, *computepb.SetBackupTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetSecurityPolicy(context.Context, *computepb.SetSecurityPolicyTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// TargetPoolsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The TargetPools API.
type TargetPoolsClient struct {
	// The internal transport-dependent client.
	internalClient internalTargetPoolsClient

	// The call options for this service.
	CallOptions *TargetPoolsCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TargetPoolsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TargetPoolsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TargetPoolsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddHealthCheck adds health check URLs to a target pool.
func (c *TargetPoolsClient) AddHealthCheck(ctx context.Context, req *computepb.AddHealthCheckTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddHealthCheck(ctx, req, opts...)
}

// AddInstance adds an instance to a target pool.
func (c *TargetPoolsClient) AddInstance(ctx context.Context, req *computepb.AddInstanceTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddInstance(ctx, req, opts...)
}

// AggregatedList retrieves an aggregated list of target pools. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *TargetPoolsClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetPoolsRequest, opts ...gax.CallOption) *TargetPoolsScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified target pool.
func (c *TargetPoolsClient) Delete(ctx context.Context, req *computepb.DeleteTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified target pool.
func (c *TargetPoolsClient) Get(ctx context.Context, req *computepb.GetTargetPoolRequest, opts ...gax.CallOption) (*computepb.TargetPool, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetHealth gets the most recent health check results for each IP for the instance that is referenced by the given target pool.
func (c *TargetPoolsClient) GetHealth(ctx context.Context, req *computepb.GetHealthTargetPoolRequest, opts ...gax.CallOption) (*computepb.TargetPoolInstanceHealth, error) {
	return c.internalClient.GetHealth(ctx, req, opts...)
}

// Insert creates a target pool in the specified project and region using the data included in the request.
func (c *TargetPoolsClient) Insert(ctx context.Context, req *computepb.InsertTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of target pools available to the specified project and region.
func (c *TargetPoolsClient) List(ctx context.Context, req *computepb.ListTargetPoolsRequest, opts ...gax.CallOption) *TargetPoolIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// RemoveHealthCheck removes health check URL from a target pool.
func (c *TargetPoolsClient) RemoveHealthCheck(ctx context.Context, req *computepb.RemoveHealthCheckTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.RemoveHealthCheck(ctx, req, opts...)
}

// RemoveInstance removes instance URL from a target pool.
func (c *TargetPoolsClient) RemoveInstance(ctx context.Context, req *computepb.RemoveInstanceTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.RemoveInstance(ctx, req, opts...)
}

// SetBackup changes a backup target pool’s configurations.
func (c *TargetPoolsClient) SetBackup(ctx context.Context, req *computepb.SetBackupTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetBackup(ctx, req, opts...)
}

// SetSecurityPolicy sets the Google Cloud Armor security policy for the specified target pool. For more information, see Google Cloud Armor Overview
func (c *TargetPoolsClient) SetSecurityPolicy(ctx context.Context, req *computepb.SetSecurityPolicyTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetSecurityPolicy(ctx, req, opts...)
}

// targetPoolsGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type targetPoolsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TargetPoolsClient
	CallOptions **TargetPoolsCallOptions

	// The gRPC API client.
	targetPoolsClient computepb.TargetPoolsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewTargetPoolsClient creates a new target pools client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The TargetPools API.
func NewTargetPoolsClient(ctx context.Context, opts ...option.ClientOption) (*TargetPoolsClient, error) {
	clientOpts := defaultTargetPoolsGRPCClientOptions()
	if newTargetPoolsClientHook != nil {
		hookOpts, err := newTargetPoolsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TargetPoolsClient{CallOptions: defaultTargetPoolsCallOptions()}

	c := &targetPoolsGRPCClient{
		connPool:    connPool,
		targetPoolsClient: computepb.NewTargetPoolsClient(connPool),
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
func (c *targetPoolsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *targetPoolsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *targetPoolsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *targetPoolsGRPCClient) AddHealthCheck(ctx context.Context, req *computepb.AddHealthCheckTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_pool", url.QueryEscape(req.GetTargetPool()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AddHealthCheck[0:len((*c.CallOptions).AddHealthCheck):len((*c.CallOptions).AddHealthCheck)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetPoolsClient.AddHealthCheck(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetPoolsGRPCClient) AddInstance(ctx context.Context, req *computepb.AddInstanceTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_pool", url.QueryEscape(req.GetTargetPool()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AddInstance[0:len((*c.CallOptions).AddInstance):len((*c.CallOptions).AddInstance)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetPoolsClient.AddInstance(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetPoolsGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetPoolsRequest, opts ...gax.CallOption) *TargetPoolsScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &TargetPoolsScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListTargetPoolsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]TargetPoolsScopedListPair, string, error) {
		resp := &computepb.TargetPoolAggregatedList{}
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
			resp, err = c.targetPoolsClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]TargetPoolsScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, TargetPoolsScopedListPair{k, v})
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

func (c *targetPoolsGRPCClient) Delete(ctx context.Context, req *computepb.DeleteTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_pool", url.QueryEscape(req.GetTargetPool()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetPoolsClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetPoolsGRPCClient) Get(ctx context.Context, req *computepb.GetTargetPoolRequest, opts ...gax.CallOption) (*computepb.TargetPool, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_pool", url.QueryEscape(req.GetTargetPool()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.TargetPool
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetPoolsClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetPoolsGRPCClient) GetHealth(ctx context.Context, req *computepb.GetHealthTargetPoolRequest, opts ...gax.CallOption) (*computepb.TargetPoolInstanceHealth, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_pool", url.QueryEscape(req.GetTargetPool()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetHealth[0:len((*c.CallOptions).GetHealth):len((*c.CallOptions).GetHealth)], opts...)
	var resp *computepb.TargetPoolInstanceHealth
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetPoolsClient.GetHealth(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetPoolsGRPCClient) Insert(ctx context.Context, req *computepb.InsertTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetPoolsClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetPoolsGRPCClient) List(ctx context.Context, req *computepb.ListTargetPoolsRequest, opts ...gax.CallOption) *TargetPoolIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &TargetPoolIterator{}
	req = proto.Clone(req).(*computepb.ListTargetPoolsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.TargetPool, string, error) {
		resp := &computepb.TargetPoolList{}
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
			resp, err = c.targetPoolsClient.List(ctx, req, settings.GRPC...)
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

func (c *targetPoolsGRPCClient) RemoveHealthCheck(ctx context.Context, req *computepb.RemoveHealthCheckTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_pool", url.QueryEscape(req.GetTargetPool()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RemoveHealthCheck[0:len((*c.CallOptions).RemoveHealthCheck):len((*c.CallOptions).RemoveHealthCheck)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetPoolsClient.RemoveHealthCheck(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetPoolsGRPCClient) RemoveInstance(ctx context.Context, req *computepb.RemoveInstanceTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_pool", url.QueryEscape(req.GetTargetPool()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RemoveInstance[0:len((*c.CallOptions).RemoveInstance):len((*c.CallOptions).RemoveInstance)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetPoolsClient.RemoveInstance(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetPoolsGRPCClient) SetBackup(ctx context.Context, req *computepb.SetBackupTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_pool", url.QueryEscape(req.GetTargetPool()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetBackup[0:len((*c.CallOptions).SetBackup):len((*c.CallOptions).SetBackup)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetPoolsClient.SetBackup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *targetPoolsGRPCClient) SetSecurityPolicy(ctx context.Context, req *computepb.SetSecurityPolicyTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "target_pool", url.QueryEscape(req.GetTargetPool()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetSecurityPolicy[0:len((*c.CallOptions).SetSecurityPolicy):len((*c.CallOptions).SetSecurityPolicy)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.targetPoolsClient.SetSecurityPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
