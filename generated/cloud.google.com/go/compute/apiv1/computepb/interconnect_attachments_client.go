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

var newInterconnectAttachmentsClientHook clientHook

// InterconnectAttachmentsCallOptions contains the retry settings for each method of InterconnectAttachmentsClient.
type InterconnectAttachmentsCallOptions struct {
	AggregatedList []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	Patch []gax.CallOption
	SetLabels []gax.CallOption
}

func defaultInterconnectAttachmentsGRPCClientOptions() []option.ClientOption {
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

func defaultInterconnectAttachmentsCallOptions() *InterconnectAttachmentsCallOptions {
	return &InterconnectAttachmentsCallOptions{
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
		Patch: []gax.CallOption{
		},
		SetLabels: []gax.CallOption{
		},
	}
}

// internalInterconnectAttachmentsClient is an interface that defines the methods available from Google Compute Engine API.
type internalInterconnectAttachmentsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListInterconnectAttachmentsRequest, ...gax.CallOption) *InterconnectAttachmentsScopedListPairIterator
	Delete(context.Context, *computepb.DeleteInterconnectAttachmentRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetInterconnectAttachmentRequest, ...gax.CallOption) (*computepb.InterconnectAttachment, error)
	Insert(context.Context, *computepb.InsertInterconnectAttachmentRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListInterconnectAttachmentsRequest, ...gax.CallOption) *InterconnectAttachmentIterator
	Patch(context.Context, *computepb.PatchInterconnectAttachmentRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetLabels(context.Context, *computepb.SetLabelsInterconnectAttachmentRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// InterconnectAttachmentsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The InterconnectAttachments API.
type InterconnectAttachmentsClient struct {
	// The internal transport-dependent client.
	internalClient internalInterconnectAttachmentsClient

	// The call options for this service.
	CallOptions *InterconnectAttachmentsCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *InterconnectAttachmentsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *InterconnectAttachmentsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *InterconnectAttachmentsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves an aggregated list of interconnect attachments. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *InterconnectAttachmentsClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListInterconnectAttachmentsRequest, opts ...gax.CallOption) *InterconnectAttachmentsScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified interconnect attachment.
func (c *InterconnectAttachmentsClient) Delete(ctx context.Context, req *computepb.DeleteInterconnectAttachmentRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified interconnect attachment.
func (c *InterconnectAttachmentsClient) Get(ctx context.Context, req *computepb.GetInterconnectAttachmentRequest, opts ...gax.CallOption) (*computepb.InterconnectAttachment, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates an InterconnectAttachment in the specified project using the data included in the request.
func (c *InterconnectAttachmentsClient) Insert(ctx context.Context, req *computepb.InsertInterconnectAttachmentRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of interconnect attachments contained within the specified region.
func (c *InterconnectAttachmentsClient) List(ctx context.Context, req *computepb.ListInterconnectAttachmentsRequest, opts ...gax.CallOption) *InterconnectAttachmentIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch updates the specified interconnect attachment with the data included in the request. This method supports PATCH semantics and uses the JSON merge patch format and processing rules.
func (c *InterconnectAttachmentsClient) Patch(ctx context.Context, req *computepb.PatchInterconnectAttachmentRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetLabels sets the labels on an InterconnectAttachment. To learn more about labels, read the Labeling Resources documentation.
func (c *InterconnectAttachmentsClient) SetLabels(ctx context.Context, req *computepb.SetLabelsInterconnectAttachmentRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetLabels(ctx, req, opts...)
}

// interconnectAttachmentsGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type interconnectAttachmentsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing InterconnectAttachmentsClient
	CallOptions **InterconnectAttachmentsCallOptions

	// The gRPC API client.
	interconnectAttachmentsClient computepb.InterconnectAttachmentsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewInterconnectAttachmentsClient creates a new interconnect attachments client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The InterconnectAttachments API.
func NewInterconnectAttachmentsClient(ctx context.Context, opts ...option.ClientOption) (*InterconnectAttachmentsClient, error) {
	clientOpts := defaultInterconnectAttachmentsGRPCClientOptions()
	if newInterconnectAttachmentsClientHook != nil {
		hookOpts, err := newInterconnectAttachmentsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := InterconnectAttachmentsClient{CallOptions: defaultInterconnectAttachmentsCallOptions()}

	c := &interconnectAttachmentsGRPCClient{
		connPool:    connPool,
		interconnectAttachmentsClient: computepb.NewInterconnectAttachmentsClient(connPool),
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
func (c *interconnectAttachmentsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *interconnectAttachmentsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *interconnectAttachmentsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *interconnectAttachmentsGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListInterconnectAttachmentsRequest, opts ...gax.CallOption) *InterconnectAttachmentsScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &InterconnectAttachmentsScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListInterconnectAttachmentsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]InterconnectAttachmentsScopedListPair, string, error) {
		resp := &computepb.InterconnectAttachmentAggregatedList{}
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
			resp, err = c.interconnectAttachmentsClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]InterconnectAttachmentsScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, InterconnectAttachmentsScopedListPair{k, v})
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

func (c *interconnectAttachmentsGRPCClient) Delete(ctx context.Context, req *computepb.DeleteInterconnectAttachmentRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "interconnect_attachment", url.QueryEscape(req.GetInterconnectAttachment()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.interconnectAttachmentsClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *interconnectAttachmentsGRPCClient) Get(ctx context.Context, req *computepb.GetInterconnectAttachmentRequest, opts ...gax.CallOption) (*computepb.InterconnectAttachment, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "interconnect_attachment", url.QueryEscape(req.GetInterconnectAttachment()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.InterconnectAttachment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.interconnectAttachmentsClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *interconnectAttachmentsGRPCClient) Insert(ctx context.Context, req *computepb.InsertInterconnectAttachmentRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.interconnectAttachmentsClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *interconnectAttachmentsGRPCClient) List(ctx context.Context, req *computepb.ListInterconnectAttachmentsRequest, opts ...gax.CallOption) *InterconnectAttachmentIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &InterconnectAttachmentIterator{}
	req = proto.Clone(req).(*computepb.ListInterconnectAttachmentsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.InterconnectAttachment, string, error) {
		resp := &computepb.InterconnectAttachmentList{}
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
			resp, err = c.interconnectAttachmentsClient.List(ctx, req, settings.GRPC...)
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

func (c *interconnectAttachmentsGRPCClient) Patch(ctx context.Context, req *computepb.PatchInterconnectAttachmentRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "interconnect_attachment", url.QueryEscape(req.GetInterconnectAttachment()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Patch[0:len((*c.CallOptions).Patch):len((*c.CallOptions).Patch)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.interconnectAttachmentsClient.Patch(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *interconnectAttachmentsGRPCClient) SetLabels(ctx context.Context, req *computepb.SetLabelsInterconnectAttachmentRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetLabels[0:len((*c.CallOptions).SetLabels):len((*c.CallOptions).SetLabels)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.interconnectAttachmentsClient.SetLabels(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
