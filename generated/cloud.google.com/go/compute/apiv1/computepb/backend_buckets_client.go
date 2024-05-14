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

var newBackendBucketsClientHook clientHook

// BackendBucketsCallOptions contains the retry settings for each method of BackendBucketsClient.
type BackendBucketsCallOptions struct {
	AddSignedUrlKey []gax.CallOption
	Delete []gax.CallOption
	DeleteSignedUrlKey []gax.CallOption
	Get []gax.CallOption
	GetIamPolicy []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	Patch []gax.CallOption
	SetEdgeSecurityPolicy []gax.CallOption
	SetIamPolicy []gax.CallOption
	TestIamPermissions []gax.CallOption
	Update []gax.CallOption
}

func defaultBackendBucketsGRPCClientOptions() []option.ClientOption {
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

func defaultBackendBucketsCallOptions() *BackendBucketsCallOptions {
	return &BackendBucketsCallOptions{
		AddSignedUrlKey: []gax.CallOption{
		},
		Delete: []gax.CallOption{
		},
		DeleteSignedUrlKey: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		GetIamPolicy: []gax.CallOption{
		},
		Insert: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
		Patch: []gax.CallOption{
		},
		SetEdgeSecurityPolicy: []gax.CallOption{
		},
		SetIamPolicy: []gax.CallOption{
		},
		TestIamPermissions: []gax.CallOption{
		},
		Update: []gax.CallOption{
		},
	}
}

// internalBackendBucketsClient is an interface that defines the methods available from Google Compute Engine API.
type internalBackendBucketsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddSignedUrlKey(context.Context, *computepb.AddSignedUrlKeyBackendBucketRequest, ...gax.CallOption) (*computepb.Operation, error)
	Delete(context.Context, *computepb.DeleteBackendBucketRequest, ...gax.CallOption) (*computepb.Operation, error)
	DeleteSignedUrlKey(context.Context, *computepb.DeleteSignedUrlKeyBackendBucketRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetBackendBucketRequest, ...gax.CallOption) (*computepb.BackendBucket, error)
	GetIamPolicy(context.Context, *computepb.GetIamPolicyBackendBucketRequest, ...gax.CallOption) (*computepb.Policy, error)
	Insert(context.Context, *computepb.InsertBackendBucketRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListBackendBucketsRequest, ...gax.CallOption) *BackendBucketIterator
	Patch(context.Context, *computepb.PatchBackendBucketRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetEdgeSecurityPolicy(context.Context, *computepb.SetEdgeSecurityPolicyBackendBucketRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetIamPolicy(context.Context, *computepb.SetIamPolicyBackendBucketRequest, ...gax.CallOption) (*computepb.Policy, error)
	TestIamPermissions(context.Context, *computepb.TestIamPermissionsBackendBucketRequest, ...gax.CallOption) (*computepb.TestPermissionsResponse, error)
	Update(context.Context, *computepb.UpdateBackendBucketRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// BackendBucketsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The BackendBuckets API.
type BackendBucketsClient struct {
	// The internal transport-dependent client.
	internalClient internalBackendBucketsClient

	// The call options for this service.
	CallOptions *BackendBucketsCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BackendBucketsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BackendBucketsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *BackendBucketsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddSignedUrlKey adds a key for validating requests with signed URLs for this backend bucket.
func (c *BackendBucketsClient) AddSignedUrlKey(ctx context.Context, req *computepb.AddSignedUrlKeyBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddSignedUrlKey(ctx, req, opts...)
}

// Delete deletes the specified BackendBucket resource.
func (c *BackendBucketsClient) Delete(ctx context.Context, req *computepb.DeleteBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// DeleteSignedUrlKey deletes a key for validating requests with signed URLs for this backend bucket.
func (c *BackendBucketsClient) DeleteSignedUrlKey(ctx context.Context, req *computepb.DeleteSignedUrlKeyBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.DeleteSignedUrlKey(ctx, req, opts...)
}

// Get returns the specified BackendBucket resource.
func (c *BackendBucketsClient) Get(ctx context.Context, req *computepb.GetBackendBucketRequest, opts ...gax.CallOption) (*computepb.BackendBucket, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists.
func (c *BackendBucketsClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyBackendBucketRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// Insert creates a BackendBucket resource in the specified project using the data included in the request.
func (c *BackendBucketsClient) Insert(ctx context.Context, req *computepb.InsertBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of BackendBucket resources available to the specified project.
func (c *BackendBucketsClient) List(ctx context.Context, req *computepb.ListBackendBucketsRequest, opts ...gax.CallOption) *BackendBucketIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch updates the specified BackendBucket resource with the data included in the request. This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (c *BackendBucketsClient) Patch(ctx context.Context, req *computepb.PatchBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetEdgeSecurityPolicy sets the edge security policy for the specified backend bucket.
func (c *BackendBucketsClient) SetEdgeSecurityPolicy(ctx context.Context, req *computepb.SetEdgeSecurityPolicyBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetEdgeSecurityPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy.
func (c *BackendBucketsClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyBackendBucketRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *BackendBucketsClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsBackendBucketRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// Update updates the specified BackendBucket resource with the data included in the request.
func (c *BackendBucketsClient) Update(ctx context.Context, req *computepb.UpdateBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Update(ctx, req, opts...)
}

// backendBucketsGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type backendBucketsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing BackendBucketsClient
	CallOptions **BackendBucketsCallOptions

	// The gRPC API client.
	backendBucketsClient computepb.BackendBucketsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewBackendBucketsClient creates a new backend buckets client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The BackendBuckets API.
func NewBackendBucketsClient(ctx context.Context, opts ...option.ClientOption) (*BackendBucketsClient, error) {
	clientOpts := defaultBackendBucketsGRPCClientOptions()
	if newBackendBucketsClientHook != nil {
		hookOpts, err := newBackendBucketsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := BackendBucketsClient{CallOptions: defaultBackendBucketsCallOptions()}

	c := &backendBucketsGRPCClient{
		connPool:    connPool,
		backendBucketsClient: computepb.NewBackendBucketsClient(connPool),
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
func (c *backendBucketsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *backendBucketsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *backendBucketsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *backendBucketsGRPCClient) AddSignedUrlKey(ctx context.Context, req *computepb.AddSignedUrlKeyBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AddSignedUrlKey[0:len((*c.CallOptions).AddSignedUrlKey):len((*c.CallOptions).AddSignedUrlKey)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendBucketsClient.AddSignedUrlKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendBucketsGRPCClient) Delete(ctx context.Context, req *computepb.DeleteBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendBucketsClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendBucketsGRPCClient) DeleteSignedUrlKey(ctx context.Context, req *computepb.DeleteSignedUrlKeyBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteSignedUrlKey[0:len((*c.CallOptions).DeleteSignedUrlKey):len((*c.CallOptions).DeleteSignedUrlKey)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendBucketsClient.DeleteSignedUrlKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendBucketsGRPCClient) Get(ctx context.Context, req *computepb.GetBackendBucketRequest, opts ...gax.CallOption) (*computepb.BackendBucket, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.BackendBucket
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendBucketsClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendBucketsGRPCClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyBackendBucketRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *computepb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendBucketsClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendBucketsGRPCClient) Insert(ctx context.Context, req *computepb.InsertBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendBucketsClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendBucketsGRPCClient) List(ctx context.Context, req *computepb.ListBackendBucketsRequest, opts ...gax.CallOption) *BackendBucketIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &BackendBucketIterator{}
	req = proto.Clone(req).(*computepb.ListBackendBucketsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.BackendBucket, string, error) {
		resp := &computepb.BackendBucketList{}
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
			resp, err = c.backendBucketsClient.List(ctx, req, settings.GRPC...)
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

func (c *backendBucketsGRPCClient) Patch(ctx context.Context, req *computepb.PatchBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendBucketsClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendBucketsGRPCClient) SetEdgeSecurityPolicy(ctx context.Context, req *computepb.SetEdgeSecurityPolicyBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetEdgeSecurityPolicy[0:len((*c.CallOptions).SetEdgeSecurityPolicy):len((*c.CallOptions).SetEdgeSecurityPolicy)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendBucketsClient.SetEdgeSecurityPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendBucketsGRPCClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyBackendBucketRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *computepb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendBucketsClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendBucketsGRPCClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsBackendBucketRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *computepb.TestPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendBucketsClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *backendBucketsGRPCClient) Update(ctx context.Context, req *computepb.UpdateBackendBucketRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "backend_bucket", url.QueryEscape(req.GetBackendBucket()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Update[0:len((*c.CallOptions).Update):len((*c.CallOptions).Update)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.backendBucketsClient.Update(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
