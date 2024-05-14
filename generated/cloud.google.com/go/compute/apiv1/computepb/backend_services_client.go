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

var newBackendServicesClientHook clientHook

// BackendServicesCallOptions contains the retry settings for each method of BackendServicesClient.
type BackendServicesCallOptions struct {
	AddSignedUrlKey []gax.CallOption
	AggregatedList []gax.CallOption
	Delete []gax.CallOption
	DeleteSignedUrlKey []gax.CallOption
	Get []gax.CallOption
	GetHealth []gax.CallOption
	GetIamPolicy []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	ListUsable []gax.CallOption
	Patch []gax.CallOption
	SetEdgeSecurityPolicy []gax.CallOption
	SetIamPolicy []gax.CallOption
	SetSecurityPolicy []gax.CallOption
	TestIamPermissions []gax.CallOption
	Update []gax.CallOption
}

func defaultBackendServicesGRPCClientOptions() []option.ClientOption {
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

func defaultBackendServicesCallOptions() *BackendServicesCallOptions {
	return &BackendServicesCallOptions{
		AddSignedUrlKey: []gax.CallOption{
		},
		AggregatedList: []gax.CallOption{
		},
		Delete: []gax.CallOption{
		},
		DeleteSignedUrlKey: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		GetHealth: []gax.CallOption{
		},
		GetIamPolicy: []gax.CallOption{
		},
		Insert: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
		ListUsable: []gax.CallOption{
		},
		Patch: []gax.CallOption{
		},
		SetEdgeSecurityPolicy: []gax.CallOption{
		},
		SetIamPolicy: []gax.CallOption{
		},
		SetSecurityPolicy: []gax.CallOption{
		},
		TestIamPermissions: []gax.CallOption{
		},
		Update: []gax.CallOption{
		},
	}
}

// internalBackendServicesClient is an interface that defines the methods available from Google Compute Engine API.
type internalBackendServicesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddSignedUrlKey(context.Context, *computepb.AddSignedUrlKeyBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	AggregatedList(context.Context, *computepb.AggregatedListBackendServicesRequest, ...gax.CallOption) *BackendServicesScopedListPairIterator
	Delete(context.Context, *computepb.DeleteBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	DeleteSignedUrlKey(context.Context, *computepb.DeleteSignedUrlKeyBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetBackendServiceRequest, ...gax.CallOption) (*computepb.BackendService, error)
	GetHealth(context.Context, *computepb.GetHealthBackendServiceRequest, ...gax.CallOption) (*computepb.BackendServiceGroupHealth, error)
	GetIamPolicy(context.Context, *computepb.GetIamPolicyBackendServiceRequest, ...gax.CallOption) (*computepb.Policy, error)
	Insert(context.Context, *computepb.InsertBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListBackendServicesRequest, ...gax.CallOption) *BackendServiceIterator
	ListUsable(context.Context, *computepb.ListUsableBackendServicesRequest, ...gax.CallOption) *BackendServiceIterator
	Patch(context.Context, *computepb.PatchBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetEdgeSecurityPolicy(context.Context, *computepb.SetEdgeSecurityPolicyBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetIamPolicy(context.Context, *computepb.SetIamPolicyBackendServiceRequest, ...gax.CallOption) (*computepb.Policy, error)
	SetSecurityPolicy(context.Context, *computepb.SetSecurityPolicyBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
	TestIamPermissions(context.Context, *computepb.TestIamPermissionsBackendServiceRequest, ...gax.CallOption) (*computepb.TestPermissionsResponse, error)
	Update(context.Context, *computepb.UpdateBackendServiceRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// BackendServicesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The BackendServices API.
type BackendServicesClient struct {
	// The internal transport-dependent client.
	internalClient internalBackendServicesClient

	// The call options for this service.
	CallOptions *BackendServicesCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BackendServicesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BackendServicesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *BackendServicesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddSignedUrlKey adds a key for validating requests with signed URLs for this backend service.
func (c *BackendServicesClient) AddSignedUrlKey(ctx context.Context, req *computepb.AddSignedUrlKeyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddSignedUrlKey(ctx, req, opts...)
}

// AggregatedList retrieves the list of all BackendService resources, regional and global, available to the specified project. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *BackendServicesClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListBackendServicesRequest, opts ...gax.CallOption) *BackendServicesScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified BackendService resource.
func (c *BackendServicesClient) Delete(ctx context.Context, req *computepb.DeleteBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// DeleteSignedUrlKey deletes a key for validating requests with signed URLs for this backend service.
func (c *BackendServicesClient) DeleteSignedUrlKey(ctx context.Context, req *computepb.DeleteSignedUrlKeyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.DeleteSignedUrlKey(ctx, req, opts...)
}

// Get returns the specified BackendService resource.
func (c *BackendServicesClient) Get(ctx context.Context, req *computepb.GetBackendServiceRequest, opts ...gax.CallOption) (*computepb.BackendService, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetHealth gets the most recent health check results for this BackendService. Example request body: { “group”: “/zones/us-east1-b/instanceGroups/lb-backend-example” }
func (c *BackendServicesClient) GetHealth(ctx context.Context, req *computepb.GetHealthBackendServiceRequest, opts ...gax.CallOption) (*computepb.BackendServiceGroupHealth, error) {
	return c.internalClient.GetHealth(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists.
func (c *BackendServicesClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// Insert creates a BackendService resource in the specified project using the data included in the request. For more information, see Backend services overview .
func (c *BackendServicesClient) Insert(ctx context.Context, req *computepb.InsertBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of BackendService resources available to the specified project.
func (c *BackendServicesClient) List(ctx context.Context, req *computepb.ListBackendServicesRequest, opts ...gax.CallOption) *BackendServiceIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// ListUsable retrieves an aggregated list of all usable backend services in the specified project.
func (c *BackendServicesClient) ListUsable(ctx context.Context, req *computepb.ListUsableBackendServicesRequest, opts ...gax.CallOption) *BackendServiceIterator {
	return c.internalClient.ListUsable(ctx, req, opts...)
}

// Patch patches the specified BackendService resource with the data included in the request. For more information, see Backend services overview. This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (c *BackendServicesClient) Patch(ctx context.Context, req *computepb.PatchBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetEdgeSecurityPolicy sets the edge security policy for the specified backend service.
func (c *BackendServicesClient) SetEdgeSecurityPolicy(ctx context.Context, req *computepb.SetEdgeSecurityPolicyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetEdgeSecurityPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy.
func (c *BackendServicesClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// SetSecurityPolicy sets the Google Cloud Armor security policy for the specified backend service. For more information, see Google Cloud Armor Overview
func (c *BackendServicesClient) SetSecurityPolicy(ctx context.Context, req *computepb.SetSecurityPolicyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetSecurityPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *BackendServicesClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsBackendServiceRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// Update updates the specified BackendService resource with the data included in the request. For more information, see Backend services overview.
func (c *BackendServicesClient) Update(ctx context.Context, req *computepb.UpdateBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Update(ctx, req, opts...)
}

// backendServicesGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type backendServicesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing BackendServicesClient
	CallOptions **BackendServicesCallOptions

	// The gRPC API client.
	backendServicesClient computepb.BackendServicesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewBackendServicesClient creates a new backend services client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The BackendServices API.
func NewBackendServicesClient(ctx context.Context, opts ...option.ClientOption) (*BackendServicesClient, error) {
	clientOpts := defaultBackendServicesGRPCClientOptions()
	if newBackendServicesClientHook != nil {
		hookOpts, err := newBackendServicesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := BackendServicesClient{CallOptions: defaultBackendServicesCallOptions()}

	c := &backendServicesGRPCClient{
		connPool:    connPool,
		backendServicesClient: computepb.NewBackendServicesClient(connPool),
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
func (c *backendServicesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *backendServicesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *backendServicesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *backendServicesGRPCClient) AddSignedUrlKey(ctx context.Context, req *computepb.AddSignedUrlKeyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_service", url.QueryEscape(req.GetBackendService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AddSignedUrlKey[0:len((*c.CallOptions).AddSignedUrlKey):len((*c.CallOptions).AddSignedUrlKey)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendServicesClient.AddSignedUrlKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendServicesGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListBackendServicesRequest, opts ...gax.CallOption) *BackendServicesScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &BackendServicesScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListBackendServicesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]BackendServicesScopedListPair, string, error) {
		resp := &computepb.BackendServiceAggregatedList{}
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
			resp, err = c.backendServicesClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]BackendServicesScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, BackendServicesScopedListPair{k, v})
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

func (c *backendServicesGRPCClient) Delete(ctx context.Context, req *computepb.DeleteBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_service", url.QueryEscape(req.GetBackendService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendServicesClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendServicesGRPCClient) DeleteSignedUrlKey(ctx context.Context, req *computepb.DeleteSignedUrlKeyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_service", url.QueryEscape(req.GetBackendService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteSignedUrlKey[0:len((*c.CallOptions).DeleteSignedUrlKey):len((*c.CallOptions).DeleteSignedUrlKey)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendServicesClient.DeleteSignedUrlKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendServicesGRPCClient) Get(ctx context.Context, req *computepb.GetBackendServiceRequest, opts ...gax.CallOption) (*computepb.BackendService, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_service", url.QueryEscape(req.GetBackendService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.BackendService
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendServicesClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendServicesGRPCClient) GetHealth(ctx context.Context, req *computepb.GetHealthBackendServiceRequest, opts ...gax.CallOption) (*computepb.BackendServiceGroupHealth, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_service", url.QueryEscape(req.GetBackendService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetHealth[0:len((*c.CallOptions).GetHealth):len((*c.CallOptions).GetHealth)], opts...)
	var resp *computepb.BackendServiceGroupHealth
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendServicesClient.GetHealth(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendServicesGRPCClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *computepb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendServicesClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendServicesGRPCClient) Insert(ctx context.Context, req *computepb.InsertBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendServicesClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendServicesGRPCClient) List(ctx context.Context, req *computepb.ListBackendServicesRequest, opts ...gax.CallOption) *BackendServiceIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &BackendServiceIterator{}
	req = proto.Clone(req).(*computepb.ListBackendServicesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.BackendService, string, error) {
		resp := &computepb.BackendServiceList{}
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
			resp, err = c.backendServicesClient.List(ctx, req, settings.GRPC...)
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

func (c *backendServicesGRPCClient) ListUsable(ctx context.Context, req *computepb.ListUsableBackendServicesRequest, opts ...gax.CallOption) *BackendServiceIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListUsable[0:len((*c.CallOptions).ListUsable):len((*c.CallOptions).ListUsable)], opts...)
	it := &BackendServiceIterator{}
	req = proto.Clone(req).(*computepb.ListUsableBackendServicesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.BackendService, string, error) {
		resp := &computepb.BackendServiceListUsable{}
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
			resp, err = c.backendServicesClient.ListUsable(ctx, req, settings.GRPC...)
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

func (c *backendServicesGRPCClient) Patch(ctx context.Context, req *computepb.PatchBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_service", url.QueryEscape(req.GetBackendService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendServicesClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendServicesGRPCClient) SetEdgeSecurityPolicy(ctx context.Context, req *computepb.SetEdgeSecurityPolicyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_service", url.QueryEscape(req.GetBackendService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetEdgeSecurityPolicy[0:len((*c.CallOptions).SetEdgeSecurityPolicy):len((*c.CallOptions).SetEdgeSecurityPolicy)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendServicesClient.SetEdgeSecurityPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendServicesGRPCClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *computepb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendServicesClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendServicesGRPCClient) SetSecurityPolicy(ctx context.Context, req *computepb.SetSecurityPolicyBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_service", url.QueryEscape(req.GetBackendService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetSecurityPolicy[0:len((*c.CallOptions).SetSecurityPolicy):len((*c.CallOptions).SetSecurityPolicy)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendServicesClient.SetSecurityPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendServicesGRPCClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsBackendServiceRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *computepb.TestPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendServicesClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendServicesGRPCClient) Update(ctx context.Context, req *computepb.UpdateBackendServiceRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_service", url.QueryEscape(req.GetBackendService()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Update[0:len((*c.CallOptions).Update):len((*c.CallOptions).Update)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendServicesClient.Update(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
