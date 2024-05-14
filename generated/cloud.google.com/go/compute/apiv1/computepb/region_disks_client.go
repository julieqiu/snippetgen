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

var newRegionDisksClientHook clientHook

// RegionDisksCallOptions contains the retry settings for each method of RegionDisksClient.
type RegionDisksCallOptions struct {
	AddResourcePolicies []gax.CallOption
	BulkInsert []gax.CallOption
	CreateSnapshot []gax.CallOption
	Delete []gax.CallOption
	Get []gax.CallOption
	GetIamPolicy []gax.CallOption
	Insert []gax.CallOption
	List []gax.CallOption
	RemoveResourcePolicies []gax.CallOption
	Resize []gax.CallOption
	SetIamPolicy []gax.CallOption
	SetLabels []gax.CallOption
	StartAsyncReplication []gax.CallOption
	StopAsyncReplication []gax.CallOption
	StopGroupAsyncReplication []gax.CallOption
	TestIamPermissions []gax.CallOption
	Update []gax.CallOption
}

func defaultRegionDisksGRPCClientOptions() []option.ClientOption {
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

func defaultRegionDisksCallOptions() *RegionDisksCallOptions {
	return &RegionDisksCallOptions{
		AddResourcePolicies: []gax.CallOption{
		},
		BulkInsert: []gax.CallOption{
		},
		CreateSnapshot: []gax.CallOption{
		},
		Delete: []gax.CallOption{
		},
		Get: []gax.CallOption{
		},
		GetIamPolicy: []gax.CallOption{
		},
		Insert: []gax.CallOption{
		},
		List: []gax.CallOption{
		},
		RemoveResourcePolicies: []gax.CallOption{
		},
		Resize: []gax.CallOption{
		},
		SetIamPolicy: []gax.CallOption{
		},
		SetLabels: []gax.CallOption{
		},
		StartAsyncReplication: []gax.CallOption{
		},
		StopAsyncReplication: []gax.CallOption{
		},
		StopGroupAsyncReplication: []gax.CallOption{
		},
		TestIamPermissions: []gax.CallOption{
		},
		Update: []gax.CallOption{
		},
	}
}

// internalRegionDisksClient is an interface that defines the methods available from Google Compute Engine API.
type internalRegionDisksClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddResourcePolicies(context.Context, *computepb.AddResourcePoliciesRegionDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	BulkInsert(context.Context, *computepb.BulkInsertRegionDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	CreateSnapshot(context.Context, *computepb.CreateSnapshotRegionDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	Delete(context.Context, *computepb.DeleteRegionDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetRegionDiskRequest, ...gax.CallOption) (*computepb.Disk, error)
	GetIamPolicy(context.Context, *computepb.GetIamPolicyRegionDiskRequest, ...gax.CallOption) (*computepb.Policy, error)
	Insert(context.Context, *computepb.InsertRegionDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListRegionDisksRequest, ...gax.CallOption) *DiskIterator
	RemoveResourcePolicies(context.Context, *computepb.RemoveResourcePoliciesRegionDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	Resize(context.Context, *computepb.ResizeRegionDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetIamPolicy(context.Context, *computepb.SetIamPolicyRegionDiskRequest, ...gax.CallOption) (*computepb.Policy, error)
	SetLabels(context.Context, *computepb.SetLabelsRegionDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	StartAsyncReplication(context.Context, *computepb.StartAsyncReplicationRegionDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	StopAsyncReplication(context.Context, *computepb.StopAsyncReplicationRegionDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	StopGroupAsyncReplication(context.Context, *computepb.StopGroupAsyncReplicationRegionDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	TestIamPermissions(context.Context, *computepb.TestIamPermissionsRegionDiskRequest, ...gax.CallOption) (*computepb.TestPermissionsResponse, error)
	Update(context.Context, *computepb.UpdateRegionDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// RegionDisksClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The RegionDisks API.
type RegionDisksClient struct {
	// The internal transport-dependent client.
	internalClient internalRegionDisksClient

	// The call options for this service.
	CallOptions *RegionDisksCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RegionDisksClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RegionDisksClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RegionDisksClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddResourcePolicies adds existing resource policies to a regional disk. You can only add one policy which will be applied to this disk for scheduling snapshot creation.
func (c *RegionDisksClient) AddResourcePolicies(ctx context.Context, req *computepb.AddResourcePoliciesRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddResourcePolicies(ctx, req, opts...)
}

// BulkInsert bulk create a set of disks.
func (c *RegionDisksClient) BulkInsert(ctx context.Context, req *computepb.BulkInsertRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.BulkInsert(ctx, req, opts...)
}

// CreateSnapshot creates a snapshot of a specified persistent disk. For regular snapshot creation, consider using snapshots.insert instead, as that method supports more features, such as creating snapshots in a project different from the source disk project.
func (c *RegionDisksClient) CreateSnapshot(ctx context.Context, req *computepb.CreateSnapshotRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.CreateSnapshot(ctx, req, opts...)
}

// Delete deletes the specified regional persistent disk. Deleting a regional disk removes all the replicas of its data permanently and is irreversible. However, deleting a disk does not delete any snapshots previously made from the disk. You must separately delete snapshots.
func (c *RegionDisksClient) Delete(ctx context.Context, req *computepb.DeleteRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns a specified regional persistent disk.
func (c *RegionDisksClient) Get(ctx context.Context, req *computepb.GetRegionDiskRequest, opts ...gax.CallOption) (*computepb.Disk, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists.
func (c *RegionDisksClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyRegionDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// Insert creates a persistent regional disk in the specified project using the data included in the request.
func (c *RegionDisksClient) Insert(ctx context.Context, req *computepb.InsertRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of persistent disks contained within the specified region.
func (c *RegionDisksClient) List(ctx context.Context, req *computepb.ListRegionDisksRequest, opts ...gax.CallOption) *DiskIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// RemoveResourcePolicies removes resource policies from a regional disk.
func (c *RegionDisksClient) RemoveResourcePolicies(ctx context.Context, req *computepb.RemoveResourcePoliciesRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.RemoveResourcePolicies(ctx, req, opts...)
}

// Resize resizes the specified regional persistent disk.
func (c *RegionDisksClient) Resize(ctx context.Context, req *computepb.ResizeRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Resize(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy.
func (c *RegionDisksClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyRegionDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// SetLabels sets the labels on the target regional disk.
func (c *RegionDisksClient) SetLabels(ctx context.Context, req *computepb.SetLabelsRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetLabels(ctx, req, opts...)
}

// StartAsyncReplication starts asynchronous replication. Must be invoked on the primary disk.
func (c *RegionDisksClient) StartAsyncReplication(ctx context.Context, req *computepb.StartAsyncReplicationRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.StartAsyncReplication(ctx, req, opts...)
}

// StopAsyncReplication stops asynchronous replication. Can be invoked either on the primary or on the secondary disk.
func (c *RegionDisksClient) StopAsyncReplication(ctx context.Context, req *computepb.StopAsyncReplicationRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.StopAsyncReplication(ctx, req, opts...)
}

// StopGroupAsyncReplication stops asynchronous replication for a consistency group of disks. Can be invoked either in the primary or secondary scope.
func (c *RegionDisksClient) StopGroupAsyncReplication(ctx context.Context, req *computepb.StopGroupAsyncReplicationRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.StopGroupAsyncReplication(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *RegionDisksClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsRegionDiskRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// Update update the specified disk with the data included in the request. Update is performed only on selected fields included as part of update-mask. Only the following fields can be modified: user_license.
func (c *RegionDisksClient) Update(ctx context.Context, req *computepb.UpdateRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Update(ctx, req, opts...)
}

// regionDisksGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type regionDisksGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing RegionDisksClient
	CallOptions **RegionDisksCallOptions

	// The gRPC API client.
	regionDisksClient computepb.RegionDisksClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewRegionDisksClient creates a new region disks client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The RegionDisks API.
func NewRegionDisksClient(ctx context.Context, opts ...option.ClientOption) (*RegionDisksClient, error) {
	clientOpts := defaultRegionDisksGRPCClientOptions()
	if newRegionDisksClientHook != nil {
		hookOpts, err := newRegionDisksClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RegionDisksClient{CallOptions: defaultRegionDisksCallOptions()}

	c := &regionDisksGRPCClient{
		connPool:    connPool,
		regionDisksClient: computepb.NewRegionDisksClient(connPool),
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
func (c *regionDisksGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *regionDisksGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *regionDisksGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *regionDisksGRPCClient) AddResourcePolicies(ctx context.Context, req *computepb.AddResourcePoliciesRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AddResourcePolicies[0:len((*c.CallOptions).AddResourcePolicies):len((*c.CallOptions).AddResourcePolicies)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.AddResourcePolicies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) BulkInsert(ctx context.Context, req *computepb.BulkInsertRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).BulkInsert[0:len((*c.CallOptions).BulkInsert):len((*c.CallOptions).BulkInsert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.BulkInsert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) CreateSnapshot(ctx context.Context, req *computepb.CreateSnapshotRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateSnapshot[0:len((*c.CallOptions).CreateSnapshot):len((*c.CallOptions).CreateSnapshot)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.CreateSnapshot(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) Delete(ctx context.Context, req *computepb.DeleteRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) Get(ctx context.Context, req *computepb.GetRegionDiskRequest, opts ...gax.CallOption) (*computepb.Disk, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.Disk
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyRegionDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *computepb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) Insert(ctx context.Context, req *computepb.InsertRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) List(ctx context.Context, req *computepb.ListRegionDisksRequest, opts ...gax.CallOption) *DiskIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &DiskIterator{}
	req = proto.Clone(req).(*computepb.ListRegionDisksRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.Disk, string, error) {
		resp := &computepb.DiskList{}
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
			resp, err = c.regionDisksClient.List(ctx, req, settings.GRPC...)
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

func (c *regionDisksGRPCClient) RemoveResourcePolicies(ctx context.Context, req *computepb.RemoveResourcePoliciesRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RemoveResourcePolicies[0:len((*c.CallOptions).RemoveResourcePolicies):len((*c.CallOptions).RemoveResourcePolicies)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.RemoveResourcePolicies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) Resize(ctx context.Context, req *computepb.ResizeRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Resize[0:len((*c.CallOptions).Resize):len((*c.CallOptions).Resize)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.Resize(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyRegionDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *computepb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) SetLabels(ctx context.Context, req *computepb.SetLabelsRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetLabels[0:len((*c.CallOptions).SetLabels):len((*c.CallOptions).SetLabels)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.SetLabels(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) StartAsyncReplication(ctx context.Context, req *computepb.StartAsyncReplicationRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).StartAsyncReplication[0:len((*c.CallOptions).StartAsyncReplication):len((*c.CallOptions).StartAsyncReplication)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.StartAsyncReplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) StopAsyncReplication(ctx context.Context, req *computepb.StopAsyncReplicationRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).StopAsyncReplication[0:len((*c.CallOptions).StopAsyncReplication):len((*c.CallOptions).StopAsyncReplication)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.StopAsyncReplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) StopGroupAsyncReplication(ctx context.Context, req *computepb.StopGroupAsyncReplicationRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).StopGroupAsyncReplication[0:len((*c.CallOptions).StopGroupAsyncReplication):len((*c.CallOptions).StopGroupAsyncReplication)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.StopGroupAsyncReplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsRegionDiskRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *computepb.TestPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *regionDisksGRPCClient) Update(ctx context.Context, req *computepb.UpdateRegionDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "region", url.QueryEscape(req.GetRegion()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Update[0:len((*c.CallOptions).Update):len((*c.CallOptions).Update)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.regionDisksClient.Update(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}