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

var newDisksClientHook clientHook

// DisksCallOptions contains the retry settings for each method of DisksClient.
type DisksCallOptions struct {
	AddResourcePolicies []gax.CallOption
	AggregatedList []gax.CallOption
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

func defaultDisksGRPCClientOptions() []option.ClientOption {
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

func defaultDisksCallOptions() *DisksCallOptions {
	return &DisksCallOptions{
		AddResourcePolicies: []gax.CallOption{
		},
		AggregatedList: []gax.CallOption{
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

// internalDisksClient is an interface that defines the methods available from Google Compute Engine API.
type internalDisksClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddResourcePolicies(context.Context, *computepb.AddResourcePoliciesDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	AggregatedList(context.Context, *computepb.AggregatedListDisksRequest, ...gax.CallOption) *DisksScopedListPairIterator
	BulkInsert(context.Context, *computepb.BulkInsertDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	CreateSnapshot(context.Context, *computepb.CreateSnapshotDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	Delete(context.Context, *computepb.DeleteDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetDiskRequest, ...gax.CallOption) (*computepb.Disk, error)
	GetIamPolicy(context.Context, *computepb.GetIamPolicyDiskRequest, ...gax.CallOption) (*computepb.Policy, error)
	Insert(context.Context, *computepb.InsertDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListDisksRequest, ...gax.CallOption) *DiskIterator
	RemoveResourcePolicies(context.Context, *computepb.RemoveResourcePoliciesDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	Resize(context.Context, *computepb.ResizeDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetIamPolicy(context.Context, *computepb.SetIamPolicyDiskRequest, ...gax.CallOption) (*computepb.Policy, error)
	SetLabels(context.Context, *computepb.SetLabelsDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	StartAsyncReplication(context.Context, *computepb.StartAsyncReplicationDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	StopAsyncReplication(context.Context, *computepb.StopAsyncReplicationDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	StopGroupAsyncReplication(context.Context, *computepb.StopGroupAsyncReplicationDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	TestIamPermissions(context.Context, *computepb.TestIamPermissionsDiskRequest, ...gax.CallOption) (*computepb.TestPermissionsResponse, error)
	Update(context.Context, *computepb.UpdateDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// DisksClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Disks API.
type DisksClient struct {
	// The internal transport-dependent client.
	internalClient internalDisksClient

	// The call options for this service.
	CallOptions *DisksCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *DisksClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *DisksClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *DisksClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddResourcePolicies adds existing resource policies to a disk. You can only add one policy which will be applied to this disk for scheduling snapshot creation.
func (c *DisksClient) AddResourcePolicies(ctx context.Context, req *computepb.AddResourcePoliciesDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddResourcePolicies(ctx, req, opts...)
}

// AggregatedList retrieves an aggregated list of persistent disks. To prevent failure, Google recommends that you set the returnPartialSuccess parameter to true.
func (c *DisksClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListDisksRequest, opts ...gax.CallOption) *DisksScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// BulkInsert bulk create a set of disks.
func (c *DisksClient) BulkInsert(ctx context.Context, req *computepb.BulkInsertDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.BulkInsert(ctx, req, opts...)
}

// CreateSnapshot creates a snapshot of a specified persistent disk. For regular snapshot creation, consider using snapshots.insert instead, as that method supports more features, such as creating snapshots in a project different from the source disk project.
func (c *DisksClient) CreateSnapshot(ctx context.Context, req *computepb.CreateSnapshotDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.CreateSnapshot(ctx, req, opts...)
}

// Delete deletes the specified persistent disk. Deleting a disk removes its data permanently and is irreversible. However, deleting a disk does not delete any snapshots previously made from the disk. You must separately delete snapshots.
func (c *DisksClient) Delete(ctx context.Context, req *computepb.DeleteDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified persistent disk.
func (c *DisksClient) Get(ctx context.Context, req *computepb.GetDiskRequest, opts ...gax.CallOption) (*computepb.Disk, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists.
func (c *DisksClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// Insert creates a persistent disk in the specified project using the data in the request. You can create a disk from a source (sourceImage, sourceSnapshot, or sourceDisk) or create an empty 500 GB data disk by omitting all properties. You can also create a disk that is larger than the default size by specifying the sizeGb property.
func (c *DisksClient) Insert(ctx context.Context, req *computepb.InsertDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of persistent disks contained within the specified zone.
func (c *DisksClient) List(ctx context.Context, req *computepb.ListDisksRequest, opts ...gax.CallOption) *DiskIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// RemoveResourcePolicies removes resource policies from a disk.
func (c *DisksClient) RemoveResourcePolicies(ctx context.Context, req *computepb.RemoveResourcePoliciesDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.RemoveResourcePolicies(ctx, req, opts...)
}

// Resize resizes the specified persistent disk. You can only increase the size of the disk.
func (c *DisksClient) Resize(ctx context.Context, req *computepb.ResizeDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Resize(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy.
func (c *DisksClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// SetLabels sets the labels on a disk. To learn more about labels, read the Labeling Resources documentation.
func (c *DisksClient) SetLabels(ctx context.Context, req *computepb.SetLabelsDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetLabels(ctx, req, opts...)
}

// StartAsyncReplication starts asynchronous replication. Must be invoked on the primary disk.
func (c *DisksClient) StartAsyncReplication(ctx context.Context, req *computepb.StartAsyncReplicationDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.StartAsyncReplication(ctx, req, opts...)
}

// StopAsyncReplication stops asynchronous replication. Can be invoked either on the primary or on the secondary disk.
func (c *DisksClient) StopAsyncReplication(ctx context.Context, req *computepb.StopAsyncReplicationDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.StopAsyncReplication(ctx, req, opts...)
}

// StopGroupAsyncReplication stops asynchronous replication for a consistency group of disks. Can be invoked either in the primary or secondary scope.
func (c *DisksClient) StopGroupAsyncReplication(ctx context.Context, req *computepb.StopGroupAsyncReplicationDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.StopGroupAsyncReplication(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *DisksClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsDiskRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// Update updates the specified disk with the data included in the request. The update is performed only on selected fields included as part of update-mask. Only the following fields can be modified: user_license.
func (c *DisksClient) Update(ctx context.Context, req *computepb.UpdateDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Update(ctx, req, opts...)
}

// disksGRPCClient is a client for interacting with Google Compute Engine API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type disksGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing DisksClient
	CallOptions **DisksCallOptions

	// The gRPC API client.
	disksClient computepb.DisksClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewDisksClient creates a new disks client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The Disks API.
func NewDisksClient(ctx context.Context, opts ...option.ClientOption) (*DisksClient, error) {
	clientOpts := defaultDisksGRPCClientOptions()
	if newDisksClientHook != nil {
		hookOpts, err := newDisksClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := DisksClient{CallOptions: defaultDisksCallOptions()}

	c := &disksGRPCClient{
		connPool:    connPool,
		disksClient: computepb.NewDisksClient(connPool),
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
func (c *disksGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *disksGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *disksGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *disksGRPCClient) AddResourcePolicies(ctx context.Context, req *computepb.AddResourcePoliciesDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AddResourcePolicies[0:len((*c.CallOptions).AddResourcePolicies):len((*c.CallOptions).AddResourcePolicies)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.AddResourcePolicies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListDisksRequest, opts ...gax.CallOption) *DisksScopedListPairIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedList[0:len((*c.CallOptions).AggregatedList):len((*c.CallOptions).AggregatedList)], opts...)
	it := &DisksScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListDisksRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]DisksScopedListPair, string, error) {
		resp := &computepb.DiskAggregatedList{}
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
			resp, err = c.disksClient.AggregatedList(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp

		elems := make([]DisksScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, DisksScopedListPair{k, v})
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

func (c *disksGRPCClient) BulkInsert(ctx context.Context, req *computepb.BulkInsertDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).BulkInsert[0:len((*c.CallOptions).BulkInsert):len((*c.CallOptions).BulkInsert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.BulkInsert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) CreateSnapshot(ctx context.Context, req *computepb.CreateSnapshotDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateSnapshot[0:len((*c.CallOptions).CreateSnapshot):len((*c.CallOptions).CreateSnapshot)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.CreateSnapshot(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) Delete(ctx context.Context, req *computepb.DeleteDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Delete[0:len((*c.CallOptions).Delete):len((*c.CallOptions).Delete)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.Delete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) Get(ctx context.Context, req *computepb.GetDiskRequest, opts ...gax.CallOption) (*computepb.Disk, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	var resp *computepb.Disk
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.Get(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *computepb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) Insert(ctx context.Context, req *computepb.InsertDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Insert[0:len((*c.CallOptions).Insert):len((*c.CallOptions).Insert)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.Insert(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) List(ctx context.Context, req *computepb.ListDisksRequest, opts ...gax.CallOption) *DiskIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).List[0:len((*c.CallOptions).List):len((*c.CallOptions).List)], opts...)
	it := &DiskIterator{}
	req = proto.Clone(req).(*computepb.ListDisksRequest)
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
			resp, err = c.disksClient.List(ctx, req, settings.GRPC...)
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

func (c *disksGRPCClient) RemoveResourcePolicies(ctx context.Context, req *computepb.RemoveResourcePoliciesDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RemoveResourcePolicies[0:len((*c.CallOptions).RemoveResourcePolicies):len((*c.CallOptions).RemoveResourcePolicies)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.RemoveResourcePolicies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) Resize(ctx context.Context, req *computepb.ResizeDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Resize[0:len((*c.CallOptions).Resize):len((*c.CallOptions).Resize)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.Resize(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *computepb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) SetLabels(ctx context.Context, req *computepb.SetLabelsDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetLabels[0:len((*c.CallOptions).SetLabels):len((*c.CallOptions).SetLabels)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.SetLabels(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) StartAsyncReplication(ctx context.Context, req *computepb.StartAsyncReplicationDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).StartAsyncReplication[0:len((*c.CallOptions).StartAsyncReplication):len((*c.CallOptions).StartAsyncReplication)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.StartAsyncReplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) StopAsyncReplication(ctx context.Context, req *computepb.StopAsyncReplicationDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).StopAsyncReplication[0:len((*c.CallOptions).StopAsyncReplication):len((*c.CallOptions).StopAsyncReplication)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.StopAsyncReplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) StopGroupAsyncReplication(ctx context.Context, req *computepb.StopGroupAsyncReplicationDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).StopGroupAsyncReplication[0:len((*c.CallOptions).StopGroupAsyncReplication):len((*c.CallOptions).StopGroupAsyncReplication)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.StopGroupAsyncReplication(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsDiskRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *computepb.TestPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *disksGRPCClient) Update(ctx context.Context, req *computepb.UpdateDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "disk", url.QueryEscape(req.GetDisk()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).Update[0:len((*c.CallOptions).Update):len((*c.CallOptions).Update)], opts...)
	var resp *computepb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.disksClient.Update(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
