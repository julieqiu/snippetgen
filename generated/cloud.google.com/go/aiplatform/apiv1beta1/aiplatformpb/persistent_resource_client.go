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


package aiplatformpb

import (
	"context"
	"fmt"
	"math"
	"net/url"

	aiplatformpb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var newPersistentResourceClientHook clientHook

// PersistentResourceCallOptions contains the retry settings for each method of PersistentResourceClient.
type PersistentResourceCallOptions struct {
	CreatePersistentResource []gax.CallOption
	GetPersistentResource []gax.CallOption
	ListPersistentResources []gax.CallOption
	DeletePersistentResource []gax.CallOption
	UpdatePersistentResource []gax.CallOption
	RebootPersistentResource []gax.CallOption
	GetLocation []gax.CallOption
	ListLocations []gax.CallOption
	GetIamPolicy []gax.CallOption
	SetIamPolicy []gax.CallOption
	TestIamPermissions []gax.CallOption
	CancelOperation []gax.CallOption
	DeleteOperation []gax.CallOption
	GetOperation []gax.CallOption
	ListOperations []gax.CallOption
	WaitOperation []gax.CallOption
}

func defaultPersistentResourceGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("aiplatform.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("aiplatform.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("aiplatform.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://aiplatform.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPersistentResourceCallOptions() *PersistentResourceCallOptions {
	return &PersistentResourceCallOptions{
		CreatePersistentResource: []gax.CallOption{
		},
		GetPersistentResource: []gax.CallOption{
		},
		ListPersistentResources: []gax.CallOption{
		},
		DeletePersistentResource: []gax.CallOption{
		},
		UpdatePersistentResource: []gax.CallOption{
		},
		RebootPersistentResource: []gax.CallOption{
		},
		GetLocation: []gax.CallOption{
		},
		ListLocations: []gax.CallOption{
		},
		GetIamPolicy: []gax.CallOption{
		},
		SetIamPolicy: []gax.CallOption{
		},
		TestIamPermissions: []gax.CallOption{
		},
		CancelOperation: []gax.CallOption{
		},
		DeleteOperation: []gax.CallOption{
		},
		GetOperation: []gax.CallOption{
		},
		ListOperations: []gax.CallOption{
		},
		WaitOperation: []gax.CallOption{
		},
	}
}

// internalPersistentResourceClient is an interface that defines the methods available from Vertex AI API.
type internalPersistentResourceClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreatePersistentResource(context.Context, *aiplatformpb.CreatePersistentResourceRequest, ...gax.CallOption) (*CreatePersistentResourceOperation, error)
	CreatePersistentResourceOperation(name string) *CreatePersistentResourceOperation
	GetPersistentResource(context.Context, *aiplatformpb.GetPersistentResourceRequest, ...gax.CallOption) (*aiplatformpb.PersistentResource, error)
	ListPersistentResources(context.Context, *aiplatformpb.ListPersistentResourcesRequest, ...gax.CallOption) *PersistentResourceIterator
	DeletePersistentResource(context.Context, *aiplatformpb.DeletePersistentResourceRequest, ...gax.CallOption) (*DeletePersistentResourceOperation, error)
	DeletePersistentResourceOperation(name string) *DeletePersistentResourceOperation
	UpdatePersistentResource(context.Context, *aiplatformpb.UpdatePersistentResourceRequest, ...gax.CallOption) (*UpdatePersistentResourceOperation, error)
	UpdatePersistentResourceOperation(name string) *UpdatePersistentResourceOperation
	RebootPersistentResource(context.Context, *aiplatformpb.RebootPersistentResourceRequest, ...gax.CallOption) (*RebootPersistentResourceOperation, error)
	RebootPersistentResourceOperation(name string) *RebootPersistentResourceOperation
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
	WaitOperation(context.Context, *longrunningpb.WaitOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
}

// PersistentResourceClient is a client for interacting with Vertex AI API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for managing Vertex AI’s machine learning PersistentResource.
type PersistentResourceClient struct {
	// The internal transport-dependent client.
	internalClient internalPersistentResourceClient

	// The call options for this service.
	CallOptions *PersistentResourceCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PersistentResourceClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PersistentResourceClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PersistentResourceClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreatePersistentResource creates a PersistentResource.
func (c *PersistentResourceClient) CreatePersistentResource(ctx context.Context, req *aiplatformpb.CreatePersistentResourceRequest, opts ...gax.CallOption) (*CreatePersistentResourceOperation, error) {
	return c.internalClient.CreatePersistentResource(ctx, req, opts...)
}

// CreatePersistentResourceOperation returns a new CreatePersistentResourceOperation from a given name.
// The name must be that of a previously created CreatePersistentResourceOperation, possibly from a different process.
func (c *PersistentResourceClient) CreatePersistentResourceOperation(name string) *CreatePersistentResourceOperation {
	return c.internalClient.CreatePersistentResourceOperation(name)
}

// GetPersistentResource gets a PersistentResource.
func (c *PersistentResourceClient) GetPersistentResource(ctx context.Context, req *aiplatformpb.GetPersistentResourceRequest, opts ...gax.CallOption) (*aiplatformpb.PersistentResource, error) {
	return c.internalClient.GetPersistentResource(ctx, req, opts...)
}

// ListPersistentResources lists PersistentResources in a Location.
func (c *PersistentResourceClient) ListPersistentResources(ctx context.Context, req *aiplatformpb.ListPersistentResourcesRequest, opts ...gax.CallOption) *PersistentResourceIterator {
	return c.internalClient.ListPersistentResources(ctx, req, opts...)
}

// DeletePersistentResource deletes a PersistentResource.
func (c *PersistentResourceClient) DeletePersistentResource(ctx context.Context, req *aiplatformpb.DeletePersistentResourceRequest, opts ...gax.CallOption) (*DeletePersistentResourceOperation, error) {
	return c.internalClient.DeletePersistentResource(ctx, req, opts...)
}

// DeletePersistentResourceOperation returns a new DeletePersistentResourceOperation from a given name.
// The name must be that of a previously created DeletePersistentResourceOperation, possibly from a different process.
func (c *PersistentResourceClient) DeletePersistentResourceOperation(name string) *DeletePersistentResourceOperation {
	return c.internalClient.DeletePersistentResourceOperation(name)
}

// UpdatePersistentResource updates a PersistentResource.
func (c *PersistentResourceClient) UpdatePersistentResource(ctx context.Context, req *aiplatformpb.UpdatePersistentResourceRequest, opts ...gax.CallOption) (*UpdatePersistentResourceOperation, error) {
	return c.internalClient.UpdatePersistentResource(ctx, req, opts...)
}

// UpdatePersistentResourceOperation returns a new UpdatePersistentResourceOperation from a given name.
// The name must be that of a previously created UpdatePersistentResourceOperation, possibly from a different process.
func (c *PersistentResourceClient) UpdatePersistentResourceOperation(name string) *UpdatePersistentResourceOperation {
	return c.internalClient.UpdatePersistentResourceOperation(name)
}

// RebootPersistentResource reboots a PersistentResource.
func (c *PersistentResourceClient) RebootPersistentResource(ctx context.Context, req *aiplatformpb.RebootPersistentResourceRequest, opts ...gax.CallOption) (*RebootPersistentResourceOperation, error) {
	return c.internalClient.RebootPersistentResource(ctx, req, opts...)
}

// RebootPersistentResourceOperation returns a new RebootPersistentResourceOperation from a given name.
// The name must be that of a previously created RebootPersistentResourceOperation, possibly from a different process.
func (c *PersistentResourceClient) RebootPersistentResourceOperation(name string) *RebootPersistentResourceOperation {
	return c.internalClient.RebootPersistentResourceOperation(name)
}

// GetLocation gets information about a location.
func (c *PersistentResourceClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// ListLocations lists information about the supported locations for this service.
func (c *PersistentResourceClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. Returns an empty policy
// if the resource exists and does not have a policy set.
func (c *PersistentResourceClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces
// any existing policy.
//
// Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED
// errors.
func (c *PersistentResourceClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource. If the
// resource does not exist, this will return an empty set of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building
// permission-aware UIs and command-line tools, not for authorization
// checking. This operation may “fail open” without warning.
func (c *PersistentResourceClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *PersistentResourceClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// DeleteOperation is a utility method from google.longrunning.Operations.
func (c *PersistentResourceClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *PersistentResourceClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *PersistentResourceClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// WaitOperation is a utility method from google.longrunning.Operations.
func (c *PersistentResourceClient) WaitOperation(ctx context.Context, req *longrunningpb.WaitOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.WaitOperation(ctx, req, opts...)
}

// persistentResourceGRPCClient is a client for interacting with Vertex AI API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type persistentResourceGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing PersistentResourceClient
	CallOptions **PersistentResourceCallOptions

	// The gRPC API client.
	persistentResourceClient aiplatformpb.PersistentResourceServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	iamPolicyClient iampb.IAMPolicyClient

	locationsClient locationpb.LocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewPersistentResourceClient creates a new persistent resource service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for managing Vertex AI’s machine learning PersistentResource.
func NewPersistentResourceClient(ctx context.Context, opts ...option.ClientOption) (*PersistentResourceClient, error) {
	clientOpts := defaultPersistentResourceGRPCClientOptions()
	if newPersistentResourceClientHook != nil {
		hookOpts, err := newPersistentResourceClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PersistentResourceClient{CallOptions: defaultPersistentResourceCallOptions()}

	c := &persistentResourceGRPCClient{
		connPool:    connPool,
		persistentResourceClient: aiplatformpb.NewPersistentResourceServiceClient(connPool),
		CallOptions: &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
		iamPolicyClient: iampb.NewIAMPolicyClient(connPool),
		locationsClient: locationpb.NewLocationsClient(connPool),

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *persistentResourceGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *persistentResourceGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *persistentResourceGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *persistentResourceGRPCClient) CreatePersistentResource(ctx context.Context, req *aiplatformpb.CreatePersistentResourceRequest, opts ...gax.CallOption) (*CreatePersistentResourceOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreatePersistentResource[0:len((*c.CallOptions).CreatePersistentResource):len((*c.CallOptions).CreatePersistentResource)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.persistentResourceClient.CreatePersistentResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreatePersistentResourceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *persistentResourceGRPCClient) GetPersistentResource(ctx context.Context, req *aiplatformpb.GetPersistentResourceRequest, opts ...gax.CallOption) (*aiplatformpb.PersistentResource, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetPersistentResource[0:len((*c.CallOptions).GetPersistentResource):len((*c.CallOptions).GetPersistentResource)], opts...)
	var resp *aiplatformpb.PersistentResource
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.persistentResourceClient.GetPersistentResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *persistentResourceGRPCClient) ListPersistentResources(ctx context.Context, req *aiplatformpb.ListPersistentResourcesRequest, opts ...gax.CallOption) *PersistentResourceIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListPersistentResources[0:len((*c.CallOptions).ListPersistentResources):len((*c.CallOptions).ListPersistentResources)], opts...)
	it := &PersistentResourceIterator{}
	req = proto.Clone(req).(*aiplatformpb.ListPersistentResourcesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.PersistentResource, string, error) {
		resp := &aiplatformpb.ListPersistentResourcesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.persistentResourceClient.ListPersistentResources(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetPersistentResources(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *persistentResourceGRPCClient) DeletePersistentResource(ctx context.Context, req *aiplatformpb.DeletePersistentResourceRequest, opts ...gax.CallOption) (*DeletePersistentResourceOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeletePersistentResource[0:len((*c.CallOptions).DeletePersistentResource):len((*c.CallOptions).DeletePersistentResource)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.persistentResourceClient.DeletePersistentResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeletePersistentResourceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *persistentResourceGRPCClient) UpdatePersistentResource(ctx context.Context, req *aiplatformpb.UpdatePersistentResourceRequest, opts ...gax.CallOption) (*UpdatePersistentResourceOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "persistent_resource.name", url.QueryEscape(req.GetPersistentResource().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdatePersistentResource[0:len((*c.CallOptions).UpdatePersistentResource):len((*c.CallOptions).UpdatePersistentResource)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.persistentResourceClient.UpdatePersistentResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdatePersistentResourceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *persistentResourceGRPCClient) RebootPersistentResource(ctx context.Context, req *aiplatformpb.RebootPersistentResourceRequest, opts ...gax.CallOption) (*RebootPersistentResourceOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RebootPersistentResource[0:len((*c.CallOptions).RebootPersistentResource):len((*c.CallOptions).RebootPersistentResource)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.persistentResourceClient.RebootPersistentResource(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &RebootPersistentResourceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *persistentResourceGRPCClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
	var resp *locationpb.Location
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.locationsClient.GetLocation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *persistentResourceGRPCClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
	it := &LocationIterator{}
	req = proto.Clone(req).(*locationpb.ListLocationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*locationpb.Location, string, error) {
		resp := &locationpb.ListLocationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.locationsClient.ListLocations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetLocations(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *persistentResourceGRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *persistentResourceGRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *persistentResourceGRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamPolicyClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *persistentResourceGRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.CancelOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *persistentResourceGRPCClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteOperation[0:len((*c.CallOptions).DeleteOperation):len((*c.CallOptions).DeleteOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.DeleteOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *persistentResourceGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *persistentResourceGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *persistentResourceGRPCClient) WaitOperation(ctx context.Context, req *longrunningpb.WaitOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).WaitOperation[0:len((*c.CallOptions).WaitOperation):len((*c.CallOptions).WaitOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.WaitOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreatePersistentResourceOperation returns a new CreatePersistentResourceOperation from a given name.
// The name must be that of a previously created CreatePersistentResourceOperation, possibly from a different process.
func (c *persistentResourceGRPCClient) CreatePersistentResourceOperation(name string) *CreatePersistentResourceOperation {
	return &CreatePersistentResourceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// DeletePersistentResourceOperation returns a new DeletePersistentResourceOperation from a given name.
// The name must be that of a previously created DeletePersistentResourceOperation, possibly from a different process.
func (c *persistentResourceGRPCClient) DeletePersistentResourceOperation(name string) *DeletePersistentResourceOperation {
	return &DeletePersistentResourceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// RebootPersistentResourceOperation returns a new RebootPersistentResourceOperation from a given name.
// The name must be that of a previously created RebootPersistentResourceOperation, possibly from a different process.
func (c *persistentResourceGRPCClient) RebootPersistentResourceOperation(name string) *RebootPersistentResourceOperation {
	return &RebootPersistentResourceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// UpdatePersistentResourceOperation returns a new UpdatePersistentResourceOperation from a given name.
// The name must be that of a previously created UpdatePersistentResourceOperation, possibly from a different process.
func (c *persistentResourceGRPCClient) UpdatePersistentResourceOperation(name string) *UpdatePersistentResourceOperation {
	return &UpdatePersistentResourceOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}