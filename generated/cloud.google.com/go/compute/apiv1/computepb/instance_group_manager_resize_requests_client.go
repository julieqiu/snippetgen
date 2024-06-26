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

var newInstanceGroupManagerResizeRequestsClientHook clientHook

// InstanceGroupManagerResizeRequestsCallOptions contains the retry settings for each method of InstanceGroupManagerResizeRequestsClient.
type InstanceGroupManagerResizeRequestsCallOptions struct {
	Cancel []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
}

func defaultInstanceGroupManagerResizeRequestsGRPCClientOptions() []option.ClientOption {
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

func defaultInstanceGroupManagerResizeRequestsCallOptions() *InstanceGroupManagerResizeRequestsCallOptions {
	return &InstanceGroupManagerResizeRequestsCallOptions{
		Cancel: []gax.CallOption{
		},
		Delete: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		Insert: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
	}
}

// internalInstanceGroupManagerResizeRequestsClient is an interface that defines the methods available from Google Compute Engine API.
type internalInstanceGroupManagerResizeRequestsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Cancel(context.Context, *computepb.CancelInstanceGroupManagerResizeRequestRequest, ...gax.CallOption) (*computepb.Operation, error)
	Delete(context.Context, *computepb.DeleteInstanceGroupManagerResizeRequestRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetInstanceGroupManagerResizeRequestRequest, ...gax.CallOption) (*computepb.InstanceGroupManagerResizeRequest, error)
	Insert(context.Context, *computepb.InsertInstanceGroupManagerResizeRequestRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListInstanceGroupManagerResizeRequestsRequest, ...gax.CallOption) *InstanceGroupManagerResizeRequestIterator
}

// InstanceGroupManagerResizeRequestsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The InstanceGroupManagerResizeRequests API.
type InstanceGroupManagerResizeRequestsClient struct {
	// The internal transport-dependent client.
	internalClient internalInstanceGroupManagerResizeRequestsClient

	// The call options for this service.
	CallOptions *InstanceGroupManagerResizeRequestsCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *InstanceGroupManagerResizeRequestsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *InstanceGroupManagerResizeRequestsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *InstanceGroupManagerResizeRequestsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Cancel cancels the specified resize request and removes it from the queue. Cancelled resize request does no longer wait for the resources to be provisioned. Cancel is only possible for requests that are accepted in the queue.
func (c *InstanceGroupManagerResizeRequestsClient) Cancel(ctx context.Context, req *computepb.CancelInstanceGroupManagerResizeRequestRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Cancel(ctx, req, opts...)
}

// Delete deletes the specified, inactive resize request. Requests that are still active cannot be deleted. Deleting request does not delete instances that were provisioned previously.
func (c *InstanceGroupManagerResizeRequestsClient) Delete(ctx context.Context, req *computepb.DeleteInstanceGroupManagerResizeRequestRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns all of the details about the specified resize request.
func (c *InstanceGroupManagerResizeRequestsClient) Get(ctx context.Context, req *computepb.GetInstanceGroupManagerResizeRequestRequest, opts ...gax.CallOption) (*computepb.InstanceGroupManagerResizeRequest, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a new resize request that starts provisioning VMs immediately or queues VM creation.
func (c *InstanceGroupManagerResizeRequestsClient) Insert(ctx context.Context, req *computepb.InsertInstanceGroupManagerResizeRequestRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of resize requests that are contained in the managed instance group.
func (c *InstanceGroupManagerResizeRequestsClient) List(ctx context.Context, req *computepb.ListInstanceGroupManagerResizeRequestsRequest, opts ...gax.CallOption) *InstanceGroupManagerResizeRequestIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// instanceGroupManagerResizeRequestsGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type instanceGroupManagerResizeRequestsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing InstanceGroupManagerResizeRequestsClient
	CallOptions **InstanceGroupManagerResizeRequestsCallOptions

	// The gRPC API client.
	instanceGroupManagerResizeRequestsClient computepb.InstanceGroupManagerResizeRequestsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewInstanceGroupManagerResizeRequestsClient creates a new instance group manager resize requests client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The InstanceGroupManagerResizeRequests API.
func NewInstanceGroupManagerResizeRequestsClient(ctx context.Context, opts ...option.ClientOption) (*InstanceGroupManagerResizeRequestsClient, error) {
	clientOpts := defaultInstanceGroupManagerResizeRequestsGRPCClientOptions()
	if newInstanceGroupManagerResizeRequestsClientHook != nil {
		hookOpts, err := newInstanceGroupManagerResizeRequestsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := InstanceGroupManagerResizeRequestsClient{CallOptions: defaultInstanceGroupManagerResizeRequestsCallOptions()}

	c := &instanceGroupManagerResizeRequestsGRPCClient{
		connPool:    connPool,
		instanceGroupManagerResizeRequestsClient: computepb.NewInstanceGroupManagerResizeRequestsClient(connPool),
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
func (c *instanceGroupManagerResizeRequestsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *instanceGroupManagerResizeRequestsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *instanceGroupManagerResizeRequestsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *instanceGroupManagerResizeRequestsGRPCClient) Cancel(ctx context.Context, req *computepb.CancelInstanceGroupManagerResizeRequestRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "instance_group_manager", url.QueryEscape(req.GetInstanceGroupManager()), "resize_request", url.QueryEscape(req.GetResizeRequest()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Cancel[0:len((*c.CallOptions).Cancel):len((*c.CallOptions).Cancel)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.instanceGroupManagerResizeRequestsClient.Cancel(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *instanceGroupManagerResizeRequestsGRPCClient) Delete(ctx context.Context, req *computepb.DeleteInstanceGroupManagerResizeRequestRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "instance_group_manager", url.QueryEscape(req.GetInstanceGroupManager()), "resize_request", url.QueryEscape(req.GetResizeRequest()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.instanceGroupManagerResizeRequestsClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *instanceGroupManagerResizeRequestsGRPCClient) Get(ctx context.Context, req *computepb.GetInstanceGroupManagerResizeRequestRequest, opts ...gax.CallOption) (*computepb.InstanceGroupManagerResizeRequest, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "instance_group_manager", url.QueryEscape(req.GetInstanceGroupManager()), "resize_request", url.QueryEscape(req.GetResizeRequest()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.InstanceGroupManagerResizeRequest
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.instanceGroupManagerResizeRequestsClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *instanceGroupManagerResizeRequestsGRPCClient) Insert(ctx context.Context, req *computepb.InsertInstanceGroupManagerResizeRequestRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "instance_group_manager", url.QueryEscape(req.GetInstanceGroupManager()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.instanceGroupManagerResizeRequestsClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *instanceGroupManagerResizeRequestsGRPCClient) List(ctx context.Context, req *computepb.ListInstanceGroupManagerResizeRequestsRequest, opts ...gax.CallOption) *InstanceGroupManagerResizeRequestIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "instance_group_manager", url.QueryEscape(req.GetInstanceGroupManager()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &InstanceGroupManagerResizeRequestIterator{}
	req = proto.Clone(req).(*computepb.ListInstanceGroupManagerResizeRequestsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.InstanceGroupManagerResizeRequest, string, error) {
		resp := &computepb.InstanceGroupManagerResizeRequestsListResponse{}
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
			resp, err = c.instanceGroupManagerResizeRequestsClient.List(ctx, req, settings.GRPC...)
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
