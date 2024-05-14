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

var newInterconnectsClientHook clientHook

// InterconnectsCallOptions contains the retry settings for each method of InterconnectsClient.
type InterconnectsCallOptions struct {
	Delete []gax.CallOption
	Get []gax.CallOption
	GetDiagnostics []gax.CallOption
	GetMacsecConfig []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	Patch []gax.CallOption
	SetLabels []gax.CallOption
}

func defaultInterconnectsGRPCClientOptions() []option.ClientOption {
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

func defaultInterconnectsCallOptions() *InterconnectsCallOptions {
	return &InterconnectsCallOptions{
		Delete: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		GetDiagnostics: []gax.CallOption{
		},
		GetMacsecConfig: []gax.CallOption{
		},
		Insert: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
		Patch: []gax.CallOption{
		},
		SetLabels: []gax.CallOption{
		},
	}
}

// internalInterconnectsClient is an interface that defines the methods available from Google Compute Engine API.
type internalInterconnectsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Delete(context.Context, *computepb.DeleteInterconnectRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetInterconnectRequest, ...gax.CallOption) (*computepb.Interconnect, error)
	GetDiagnostics(context.Context, *computepb.GetDiagnosticsInterconnectRequest, ...gax.CallOption) (*computepb.InterconnectsGetDiagnosticsResponse, error)
	GetMacsecConfig(context.Context, *computepb.GetMacsecConfigInterconnectRequest, ...gax.CallOption) (*computepb.InterconnectsGetMacsecConfigResponse, error)
	Insert(context.Context, *computepb.InsertInterconnectRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListInterconnectsRequest, ...gax.CallOption) *InterconnectIterator
	Patch(context.Context, *computepb.PatchInterconnectRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetLabels(context.Context, *computepb.SetLabelsInterconnectRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// InterconnectsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Interconnects API.
type InterconnectsClient struct {
	// The internal transport-dependent client.
	internalClient internalInterconnectsClient

	// The call options for this service.
	CallOptions *InterconnectsCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *InterconnectsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *InterconnectsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *InterconnectsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Delete deletes the specified Interconnect.
func (c *InterconnectsClient) Delete(ctx context.Context, req *computepb.DeleteInterconnectRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified Interconnect. Get a list of available Interconnects by making a list() request.
func (c *InterconnectsClient) Get(ctx context.Context, req *computepb.GetInterconnectRequest, opts ...gax.CallOption) (*computepb.Interconnect, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetDiagnostics returns the interconnectDiagnostics for the specified Interconnect. In the event of a global outage, do not use this API to make decisions about where to redirect your network traffic. Unlike a VLAN attachment, which is regional, a Cloud Interconnect connection is a global resource. A global outage can prevent this API from functioning properly.
func (c *InterconnectsClient) GetDiagnostics(ctx context.Context, req *computepb.GetDiagnosticsInterconnectRequest, opts ...gax.CallOption) (*computepb.InterconnectsGetDiagnosticsResponse, error) {
	return c.internalClient.GetDiagnostics(ctx, req, opts...)
}

// GetMacsecConfig returns the interconnectMacsecConfig for the specified Interconnect.
func (c *InterconnectsClient) GetMacsecConfig(ctx context.Context, req *computepb.GetMacsecConfigInterconnectRequest, opts ...gax.CallOption) (*computepb.InterconnectsGetMacsecConfigResponse, error) {
	return c.internalClient.GetMacsecConfig(ctx, req, opts...)
}

// Insert creates an Interconnect in the specified project using the data included in the request.
func (c *InterconnectsClient) Insert(ctx context.Context, req *computepb.InsertInterconnectRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of Interconnects available to the specified project.
func (c *InterconnectsClient) List(ctx context.Context, req *computepb.ListInterconnectsRequest, opts ...gax.CallOption) *InterconnectIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch updates the specified Interconnect with the data included in the request. This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (c *InterconnectsClient) Patch(ctx context.Context, req *computepb.PatchInterconnectRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetLabels sets the labels on an Interconnect. To learn more about labels, read the Labeling Resources documentation.
func (c *InterconnectsClient) SetLabels(ctx context.Context, req *computepb.SetLabelsInterconnectRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetLabels(ctx, req, opts...)
}

// interconnectsGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type interconnectsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing InterconnectsClient
	CallOptions **InterconnectsCallOptions

	// The gRPC API client.
	interconnectsClient computepb.InterconnectsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewInterconnectsClient creates a new interconnects client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The Interconnects API.
func NewInterconnectsClient(ctx context.Context, opts ...option.ClientOption) (*InterconnectsClient, error) {
	clientOpts := defaultInterconnectsGRPCClientOptions()
	if newInterconnectsClientHook != nil {
		hookOpts, err := newInterconnectsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := InterconnectsClient{CallOptions: defaultInterconnectsCallOptions()}

	c := &interconnectsGRPCClient{
		connPool:    connPool,
		interconnectsClient: computepb.NewInterconnectsClient(connPool),
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
func (c *interconnectsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *interconnectsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *interconnectsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *interconnectsGRPCClient) Delete(ctx context.Context, req *computepb.DeleteInterconnectRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "interconnect", url.QueryEscape(req.GetInterconnect()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.interconnectsClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *interconnectsGRPCClient) Get(ctx context.Context, req *computepb.GetInterconnectRequest, opts ...gax.CallOption) (*computepb.Interconnect, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "interconnect", url.QueryEscape(req.GetInterconnect()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.Interconnect
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.interconnectsClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *interconnectsGRPCClient) GetDiagnostics(ctx context.Context, req *computepb.GetDiagnosticsInterconnectRequest, opts ...gax.CallOption) (*computepb.InterconnectsGetDiagnosticsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "interconnect", url.QueryEscape(req.GetInterconnect()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetDiagnostics[0:len((*c.CallOptions).GetDiagnostics):len((*c.CallOptions).GetDiagnostics)], opts...)
	var resp *computepb.InterconnectsGetDiagnosticsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.interconnectsClient.GetDiagnostics(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *interconnectsGRPCClient) GetMacsecConfig(ctx context.Context, req *computepb.GetMacsecConfigInterconnectRequest, opts ...gax.CallOption) (*computepb.InterconnectsGetMacsecConfigResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "interconnect", url.QueryEscape(req.GetInterconnect()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetMacsecConfig[0:len((*c.CallOptions).GetMacsecConfig):len((*c.CallOptions).GetMacsecConfig)], opts...)
	var resp *computepb.InterconnectsGetMacsecConfigResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.interconnectsClient.GetMacsecConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *interconnectsGRPCClient) Insert(ctx context.Context, req *computepb.InsertInterconnectRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.interconnectsClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *interconnectsGRPCClient) List(ctx context.Context, req *computepb.ListInterconnectsRequest, opts ...gax.CallOption) *InterconnectIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &InterconnectIterator{}
	req = proto.Clone(req).(*computepb.ListInterconnectsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.Interconnect, string, error) {
		resp := &computepb.InterconnectList{}
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
			resp, err = c.interconnectsClient.List(ctx, req, settings.GRPC...)
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

func (c *interconnectsGRPCClient) Patch(ctx context.Context, req *computepb.PatchInterconnectRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "interconnect", url.QueryEscape(req.GetInterconnect()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.interconnectsClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *interconnectsGRPCClient) SetLabels(ctx context.Context, req *computepb.SetLabelsInterconnectRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetLabels[0:len((*c.CallOptions).SetLabels):len((*c.CallOptions).SetLabels)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.interconnectsClient.SetLabels(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
