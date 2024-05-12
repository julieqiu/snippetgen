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


package secretmanagerpb

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	iampb "cloud.google.com/go/iam/apiv1/iampb"
	secretmanagerpb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var newSecretManagerClientHook clientHook

// SecretManagerCallOptions contains the retry settings for each method of SecretManagerClient.
type SecretManagerCallOptions struct {
	ListSecrets []gax.CallOption
	CreateSecret []gax.CallOption
	AddSecretVersion []gax.CallOption
	GetSecret []gax.CallOption
	UpdateSecret []gax.CallOption
	DeleteSecret []gax.CallOption
	ListSecretVersions []gax.CallOption
	GetSecretVersion []gax.CallOption
	AccessSecretVersion []gax.CallOption
	DisableSecretVersion []gax.CallOption
	EnableSecretVersion []gax.CallOption
	DestroySecretVersion []gax.CallOption
	SetIamPolicy []gax.CallOption
	GetIamPolicy []gax.CallOption
	TestIamPermissions []gax.CallOption
}

func defaultSecretManagerGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("secretmanager.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("secretmanager.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("secretmanager.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://secretmanager.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSecretManagerCallOptions() *SecretManagerCallOptions {
	return &SecretManagerCallOptions{
		ListSecrets: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		CreateSecret: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		AddSecretVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetSecret: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateSecret: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteSecret: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListSecretVersions: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetSecretVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		AccessSecretVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.ResourceExhausted,
				}, gax.Backoff{
					Initial:    2000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 2.00,
				})
			}),
		},
		DisableSecretVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		EnableSecretVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DestroySecretVersion: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		SetIamPolicy: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetIamPolicy: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		TestIamPermissions: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalSecretManagerClient is an interface that defines the methods available from Secret Manager API.
type internalSecretManagerClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListSecrets(context.Context, *secretmanagerpb.ListSecretsRequest, ...gax.CallOption) *SecretIterator
	CreateSecret(context.Context, *secretmanagerpb.CreateSecretRequest, ...gax.CallOption) (*secretmanagerpb.Secret, error)
	AddSecretVersion(context.Context, *secretmanagerpb.AddSecretVersionRequest, ...gax.CallOption) (*secretmanagerpb.SecretVersion, error)
	GetSecret(context.Context, *secretmanagerpb.GetSecretRequest, ...gax.CallOption) (*secretmanagerpb.Secret, error)
	UpdateSecret(context.Context, *secretmanagerpb.UpdateSecretRequest, ...gax.CallOption) (*secretmanagerpb.Secret, error)
	DeleteSecret(context.Context, *secretmanagerpb.DeleteSecretRequest, ...gax.CallOption) error
	ListSecretVersions(context.Context, *secretmanagerpb.ListSecretVersionsRequest, ...gax.CallOption) *SecretVersionIterator
	GetSecretVersion(context.Context, *secretmanagerpb.GetSecretVersionRequest, ...gax.CallOption) (*secretmanagerpb.SecretVersion, error)
	AccessSecretVersion(context.Context, *secretmanagerpb.AccessSecretVersionRequest, ...gax.CallOption) (*secretmanagerpb.AccessSecretVersionResponse, error)
	DisableSecretVersion(context.Context, *secretmanagerpb.DisableSecretVersionRequest, ...gax.CallOption) (*secretmanagerpb.SecretVersion, error)
	EnableSecretVersion(context.Context, *secretmanagerpb.EnableSecretVersionRequest, ...gax.CallOption) (*secretmanagerpb.SecretVersion, error)
	DestroySecretVersion(context.Context, *secretmanagerpb.DestroySecretVersionRequest, ...gax.CallOption) (*secretmanagerpb.SecretVersion, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
}

// SecretManagerClient is a client for interacting with Secret Manager API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Secret Manager Service
//
// Manages secrets and operations using those secrets. Implements a REST
// model with the following objects:
//
//   Secret
//
//   SecretVersion
type SecretManagerClient struct {
	// The internal transport-dependent client.
	internalClient internalSecretManagerClient

	// The call options for this service.
	CallOptions *SecretManagerCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SecretManagerClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SecretManagerClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *SecretManagerClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListSecrets lists Secrets.
func (c *SecretManagerClient) ListSecrets(ctx context.Context, req *secretmanagerpb.ListSecretsRequest, opts ...gax.CallOption) *SecretIterator {
	return c.internalClient.ListSecrets(ctx, req, opts...)
}

// CreateSecret creates a new Secret containing no SecretVersions.
func (c *SecretManagerClient) CreateSecret(ctx context.Context, req *secretmanagerpb.CreateSecretRequest, opts ...gax.CallOption) (*secretmanagerpb.Secret, error) {
	return c.internalClient.CreateSecret(ctx, req, opts...)
}

// AddSecretVersion creates a new SecretVersion containing secret data and attaches
// it to an existing Secret.
func (c *SecretManagerClient) AddSecretVersion(ctx context.Context, req *secretmanagerpb.AddSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.SecretVersion, error) {
	return c.internalClient.AddSecretVersion(ctx, req, opts...)
}

// GetSecret gets metadata for a given Secret.
func (c *SecretManagerClient) GetSecret(ctx context.Context, req *secretmanagerpb.GetSecretRequest, opts ...gax.CallOption) (*secretmanagerpb.Secret, error) {
	return c.internalClient.GetSecret(ctx, req, opts...)
}

// UpdateSecret updates metadata of an existing Secret.
func (c *SecretManagerClient) UpdateSecret(ctx context.Context, req *secretmanagerpb.UpdateSecretRequest, opts ...gax.CallOption) (*secretmanagerpb.Secret, error) {
	return c.internalClient.UpdateSecret(ctx, req, opts...)
}

// DeleteSecret deletes a Secret.
func (c *SecretManagerClient) DeleteSecret(ctx context.Context, req *secretmanagerpb.DeleteSecretRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteSecret(ctx, req, opts...)
}

// ListSecretVersions lists SecretVersions. This call does not return secret
// data.
func (c *SecretManagerClient) ListSecretVersions(ctx context.Context, req *secretmanagerpb.ListSecretVersionsRequest, opts ...gax.CallOption) *SecretVersionIterator {
	return c.internalClient.ListSecretVersions(ctx, req, opts...)
}

// GetSecretVersion gets metadata for a SecretVersion.
//
// projects/*/secrets/*/versions/latest is an alias to the most recently
// created SecretVersion.
func (c *SecretManagerClient) GetSecretVersion(ctx context.Context, req *secretmanagerpb.GetSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.SecretVersion, error) {
	return c.internalClient.GetSecretVersion(ctx, req, opts...)
}

// AccessSecretVersion accesses a SecretVersion. This call returns the secret data.
//
// projects/*/secrets/*/versions/latest is an alias to the most recently
// created SecretVersion.
func (c *SecretManagerClient) AccessSecretVersion(ctx context.Context, req *secretmanagerpb.AccessSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.AccessSecretVersionResponse, error) {
	return c.internalClient.AccessSecretVersion(ctx, req, opts...)
}

// DisableSecretVersion disables a SecretVersion.
//
// Sets the state of the SecretVersion to
// DISABLED.
func (c *SecretManagerClient) DisableSecretVersion(ctx context.Context, req *secretmanagerpb.DisableSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.SecretVersion, error) {
	return c.internalClient.DisableSecretVersion(ctx, req, opts...)
}

// EnableSecretVersion enables a SecretVersion.
//
// Sets the state of the SecretVersion to
// ENABLED.
func (c *SecretManagerClient) EnableSecretVersion(ctx context.Context, req *secretmanagerpb.EnableSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.SecretVersion, error) {
	return c.internalClient.EnableSecretVersion(ctx, req, opts...)
}

// DestroySecretVersion destroys a SecretVersion.
//
// Sets the state of the SecretVersion to
// DESTROYED and irrevocably destroys the
// secret data.
func (c *SecretManagerClient) DestroySecretVersion(ctx context.Context, req *secretmanagerpb.DestroySecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.SecretVersion, error) {
	return c.internalClient.DestroySecretVersion(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified secret. Replaces any
// existing policy.
//
// Permissions on SecretVersions are enforced according
// to the policy set on the associated Secret.
func (c *SecretManagerClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a secret.
// Returns empty policy if the secret exists and does not have a policy set.
func (c *SecretManagerClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has for the specified secret.
// If the secret does not exist, this call returns an empty set of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building permission-aware
// UIs and command-line tools, not for authorization checking. This operation
// may “fail open” without warning.
func (c *SecretManagerClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// secretManagerGRPCClient is a client for interacting with Secret Manager API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type secretManagerGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing SecretManagerClient
	CallOptions **SecretManagerCallOptions

	// The gRPC API client.
	secretManagerClient secretmanagerpb.SecretManagerServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewSecretManagerClient creates a new secret manager service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Secret Manager Service
//
// Manages secrets and operations using those secrets. Implements a REST
// model with the following objects:
//
//   Secret
//
//   SecretVersion
func NewSecretManagerClient(ctx context.Context, opts ...option.ClientOption) (*SecretManagerClient, error) {
	clientOpts := defaultSecretManagerGRPCClientOptions()
	if newSecretManagerClientHook != nil {
		hookOpts, err := newSecretManagerClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := SecretManagerClient{CallOptions: defaultSecretManagerCallOptions()}

	c := &secretManagerGRPCClient{
		connPool:    connPool,
		secretManagerClient: secretmanagerpb.NewSecretManagerServiceClient(connPool),
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
func (c *secretManagerGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *secretManagerGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *secretManagerGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *secretManagerGRPCClient) ListSecrets(ctx context.Context, req *secretmanagerpb.ListSecretsRequest, opts ...gax.CallOption) *SecretIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListSecrets[0:len((*c.CallOptions).ListSecrets):len((*c.CallOptions).ListSecrets)], opts...)
	it := &SecretIterator{}
	req = proto.Clone(req).(*secretmanagerpb.ListSecretsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*secretmanagerpb.Secret, string, error) {
		resp := &secretmanagerpb.ListSecretsResponse{}
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
			resp, err = c.secretManagerClient.ListSecrets(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetSecrets(), resp.GetNextPageToken(), nil
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

func (c *secretManagerGRPCClient) CreateSecret(ctx context.Context, req *secretmanagerpb.CreateSecretRequest, opts ...gax.CallOption) (*secretmanagerpb.Secret, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateSecret[0:len((*c.CallOptions).CreateSecret):len((*c.CallOptions).CreateSecret)], opts...)
	var resp *secretmanagerpb.Secret
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.secretManagerClient.CreateSecret(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *secretManagerGRPCClient) AddSecretVersion(ctx context.Context, req *secretmanagerpb.AddSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.SecretVersion, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AddSecretVersion[0:len((*c.CallOptions).AddSecretVersion):len((*c.CallOptions).AddSecretVersion)], opts...)
	var resp *secretmanagerpb.SecretVersion
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.secretManagerClient.AddSecretVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *secretManagerGRPCClient) GetSecret(ctx context.Context, req *secretmanagerpb.GetSecretRequest, opts ...gax.CallOption) (*secretmanagerpb.Secret, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetSecret[0:len((*c.CallOptions).GetSecret):len((*c.CallOptions).GetSecret)], opts...)
	var resp *secretmanagerpb.Secret
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.secretManagerClient.GetSecret(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *secretManagerGRPCClient) UpdateSecret(ctx context.Context, req *secretmanagerpb.UpdateSecretRequest, opts ...gax.CallOption) (*secretmanagerpb.Secret, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "secret.name", url.QueryEscape(req.GetSecret().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateSecret[0:len((*c.CallOptions).UpdateSecret):len((*c.CallOptions).UpdateSecret)], opts...)
	var resp *secretmanagerpb.Secret
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.secretManagerClient.UpdateSecret(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *secretManagerGRPCClient) DeleteSecret(ctx context.Context, req *secretmanagerpb.DeleteSecretRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteSecret[0:len((*c.CallOptions).DeleteSecret):len((*c.CallOptions).DeleteSecret)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.secretManagerClient.DeleteSecret(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *secretManagerGRPCClient) ListSecretVersions(ctx context.Context, req *secretmanagerpb.ListSecretVersionsRequest, opts ...gax.CallOption) *SecretVersionIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListSecretVersions[0:len((*c.CallOptions).ListSecretVersions):len((*c.CallOptions).ListSecretVersions)], opts...)
	it := &SecretVersionIterator{}
	req = proto.Clone(req).(*secretmanagerpb.ListSecretVersionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*secretmanagerpb.SecretVersion, string, error) {
		resp := &secretmanagerpb.ListSecretVersionsResponse{}
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
			resp, err = c.secretManagerClient.ListSecretVersions(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetVersions(), resp.GetNextPageToken(), nil
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

func (c *secretManagerGRPCClient) GetSecretVersion(ctx context.Context, req *secretmanagerpb.GetSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.SecretVersion, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetSecretVersion[0:len((*c.CallOptions).GetSecretVersion):len((*c.CallOptions).GetSecretVersion)], opts...)
	var resp *secretmanagerpb.SecretVersion
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.secretManagerClient.GetSecretVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *secretManagerGRPCClient) AccessSecretVersion(ctx context.Context, req *secretmanagerpb.AccessSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.AccessSecretVersionResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AccessSecretVersion[0:len((*c.CallOptions).AccessSecretVersion):len((*c.CallOptions).AccessSecretVersion)], opts...)
	var resp *secretmanagerpb.AccessSecretVersionResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.secretManagerClient.AccessSecretVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *secretManagerGRPCClient) DisableSecretVersion(ctx context.Context, req *secretmanagerpb.DisableSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.SecretVersion, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DisableSecretVersion[0:len((*c.CallOptions).DisableSecretVersion):len((*c.CallOptions).DisableSecretVersion)], opts...)
	var resp *secretmanagerpb.SecretVersion
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.secretManagerClient.DisableSecretVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *secretManagerGRPCClient) EnableSecretVersion(ctx context.Context, req *secretmanagerpb.EnableSecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.SecretVersion, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).EnableSecretVersion[0:len((*c.CallOptions).EnableSecretVersion):len((*c.CallOptions).EnableSecretVersion)], opts...)
	var resp *secretmanagerpb.SecretVersion
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.secretManagerClient.EnableSecretVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *secretManagerGRPCClient) DestroySecretVersion(ctx context.Context, req *secretmanagerpb.DestroySecretVersionRequest, opts ...gax.CallOption) (*secretmanagerpb.SecretVersion, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DestroySecretVersion[0:len((*c.CallOptions).DestroySecretVersion):len((*c.CallOptions).DestroySecretVersion)], opts...)
	var resp *secretmanagerpb.SecretVersion
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.secretManagerClient.DestroySecretVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *secretManagerGRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.secretManagerClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *secretManagerGRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.secretManagerClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *secretManagerGRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.secretManagerClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
