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


package computepb_test

import (
	"context"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/iterator"
)

func ExampleNewUrlMapsClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewUrlMapsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleUrlMapsClient_AggregatedList() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewUrlMapsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AggregatedListUrlMapsRequest.
	req := &computepb.AggregatedListUrlMapsRequest{
		Filter: "",
		IncludeAllScopes: "",
		MaxResults: "",
		OrderBy: "",
		PageToken: "",
		Project: "",
		ReturnPartialSuccess: "",
		ServiceProjectNumber: "",
	}
	it := c.AggregatedList(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*computepb.UrlMapsAggregatedList)
	}
}

func ExampleUrlMapsClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewUrlMapsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteUrlMapRequest.
	req := &computepb.DeleteUrlMapRequest{
		Project: "",
		RequestId: "",
		UrlMap: "",
	}
	resp, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleUrlMapsClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewUrlMapsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetUrlMapRequest.
	req := &computepb.GetUrlMapRequest{
		Project: "",
		UrlMap: "",
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleUrlMapsClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewUrlMapsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertUrlMapRequest.
	req := &computepb.InsertUrlMapRequest{
		Project: "",
		RequestId: "",
		UrlMapResource: &computepb.UrlMap{
			CreationTimestamp: "",
			DefaultRouteAction: &computepb.HttpRouteAction{
				CorsPolicy: &computepb.CorsPolicy{...}
				FaultInjectionPolicy: &computepb.HttpFaultInjection{...}
				MaxStreamDuration: &computepb.Duration{...}
				RequestMirrorPolicy: &computepb.RequestMirrorPolicy{...}
				RetryPolicy: &computepb.HttpRetryPolicy{...}
				Timeout: &computepb.Duration{...}
				UrlRewrite: &computepb.UrlRewrite{...}
				WeightedBackendServices: &computepb.WeightedBackendService{...}
			}
			DefaultService: "",
			DefaultUrlRedirect: &computepb.HttpRedirectAction{
				HostRedirect: "",
				HttpsRedirect: "",
				PathRedirect: "",
				PrefixRedirect: "",
				RedirectResponseCode: "",
				StripQuery: "",
			}
			Description: "",
			Fingerprint: "",
			HeaderAction: &computepb.HttpHeaderAction{
				RequestHeadersToAdd: &computepb.HttpHeaderOption{...}
				RequestHeadersToRemove: "",
				ResponseHeadersToAdd: &computepb.HttpHeaderOption{...}
				ResponseHeadersToRemove: "",
			}
			HostRules: &computepb.HostRule{
				Description: "",
				Hosts: "",
				PathMatcher: "",
			}
			Id: "",
			Kind: "",
			Name: "",
			PathMatchers: &computepb.PathMatcher{
				DefaultRouteAction: &computepb.HttpRouteAction{...}
				DefaultService: "",
				DefaultUrlRedirect: &computepb.HttpRedirectAction{...}
				Description: "",
				HeaderAction: &computepb.HttpHeaderAction{...}
				Name: "",
				PathRules: &computepb.PathRule{...}
				RouteRules: &computepb.HttpRouteRule{...}
			}
			Region: "",
			SelfLink: "",
			Tests: &computepb.UrlMapTest{
				Description: "",
				ExpectedOutputUrl: "",
				ExpectedRedirectResponseCode: "",
				Headers: &computepb.UrlMapTestHeader{...}
				Host: "",
				Path: "",
				Service: "",
			}
		}
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleUrlMapsClient_InvalidateCache() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewUrlMapsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InvalidateCacheUrlMapRequest.
	req := &computepb.InvalidateCacheUrlMapRequest{
		CacheInvalidationRuleResource: &computepb.CacheInvalidationRule{
			Host: "",
			Path: "",
		}
		Project: "",
		RequestId: "",
		UrlMap: "",
	}
	resp, err := c.InvalidateCache(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleUrlMapsClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewUrlMapsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListUrlMapsRequest.
	req := &computepb.ListUrlMapsRequest{
		Filter: "",
		MaxResults: "",
		OrderBy: "",
		PageToken: "",
		Project: "",
		ReturnPartialSuccess: "",
	}
	it := c.List(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*computepb.UrlMapList)
	}
}

func ExampleUrlMapsClient_Patch() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewUrlMapsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchUrlMapRequest.
	req := &computepb.PatchUrlMapRequest{
		Project: "",
		RequestId: "",
		UrlMap: "",
		UrlMapResource: &computepb.UrlMap{
			CreationTimestamp: "",
			DefaultRouteAction: &computepb.HttpRouteAction{
				CorsPolicy: &computepb.CorsPolicy{...}
				FaultInjectionPolicy: &computepb.HttpFaultInjection{...}
				MaxStreamDuration: &computepb.Duration{...}
				RequestMirrorPolicy: &computepb.RequestMirrorPolicy{...}
				RetryPolicy: &computepb.HttpRetryPolicy{...}
				Timeout: &computepb.Duration{...}
				UrlRewrite: &computepb.UrlRewrite{...}
				WeightedBackendServices: &computepb.WeightedBackendService{...}
			}
			DefaultService: "",
			DefaultUrlRedirect: &computepb.HttpRedirectAction{
				HostRedirect: "",
				HttpsRedirect: "",
				PathRedirect: "",
				PrefixRedirect: "",
				RedirectResponseCode: "",
				StripQuery: "",
			}
			Description: "",
			Fingerprint: "",
			HeaderAction: &computepb.HttpHeaderAction{
				RequestHeadersToAdd: &computepb.HttpHeaderOption{...}
				RequestHeadersToRemove: "",
				ResponseHeadersToAdd: &computepb.HttpHeaderOption{...}
				ResponseHeadersToRemove: "",
			}
			HostRules: &computepb.HostRule{
				Description: "",
				Hosts: "",
				PathMatcher: "",
			}
			Id: "",
			Kind: "",
			Name: "",
			PathMatchers: &computepb.PathMatcher{
				DefaultRouteAction: &computepb.HttpRouteAction{...}
				DefaultService: "",
				DefaultUrlRedirect: &computepb.HttpRedirectAction{...}
				Description: "",
				HeaderAction: &computepb.HttpHeaderAction{...}
				Name: "",
				PathRules: &computepb.PathRule{...}
				RouteRules: &computepb.HttpRouteRule{...}
			}
			Region: "",
			SelfLink: "",
			Tests: &computepb.UrlMapTest{
				Description: "",
				ExpectedOutputUrl: "",
				ExpectedRedirectResponseCode: "",
				Headers: &computepb.UrlMapTestHeader{...}
				Host: "",
				Path: "",
				Service: "",
			}
		}
	}
	resp, err := c.Patch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleUrlMapsClient_Update() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewUrlMapsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#UpdateUrlMapRequest.
	req := &computepb.UpdateUrlMapRequest{
		Project: "",
		RequestId: "",
		UrlMap: "",
		UrlMapResource: &computepb.UrlMap{
			CreationTimestamp: "",
			DefaultRouteAction: &computepb.HttpRouteAction{
				CorsPolicy: &computepb.CorsPolicy{...}
				FaultInjectionPolicy: &computepb.HttpFaultInjection{...}
				MaxStreamDuration: &computepb.Duration{...}
				RequestMirrorPolicy: &computepb.RequestMirrorPolicy{...}
				RetryPolicy: &computepb.HttpRetryPolicy{...}
				Timeout: &computepb.Duration{...}
				UrlRewrite: &computepb.UrlRewrite{...}
				WeightedBackendServices: &computepb.WeightedBackendService{...}
			}
			DefaultService: "",
			DefaultUrlRedirect: &computepb.HttpRedirectAction{
				HostRedirect: "",
				HttpsRedirect: "",
				PathRedirect: "",
				PrefixRedirect: "",
				RedirectResponseCode: "",
				StripQuery: "",
			}
			Description: "",
			Fingerprint: "",
			HeaderAction: &computepb.HttpHeaderAction{
				RequestHeadersToAdd: &computepb.HttpHeaderOption{...}
				RequestHeadersToRemove: "",
				ResponseHeadersToAdd: &computepb.HttpHeaderOption{...}
				ResponseHeadersToRemove: "",
			}
			HostRules: &computepb.HostRule{
				Description: "",
				Hosts: "",
				PathMatcher: "",
			}
			Id: "",
			Kind: "",
			Name: "",
			PathMatchers: &computepb.PathMatcher{
				DefaultRouteAction: &computepb.HttpRouteAction{...}
				DefaultService: "",
				DefaultUrlRedirect: &computepb.HttpRedirectAction{...}
				Description: "",
				HeaderAction: &computepb.HttpHeaderAction{...}
				Name: "",
				PathRules: &computepb.PathRule{...}
				RouteRules: &computepb.HttpRouteRule{...}
			}
			Region: "",
			SelfLink: "",
			Tests: &computepb.UrlMapTest{
				Description: "",
				ExpectedOutputUrl: "",
				ExpectedRedirectResponseCode: "",
				Headers: &computepb.UrlMapTestHeader{...}
				Host: "",
				Path: "",
				Service: "",
			}
		}
	}
	resp, err := c.Update(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleUrlMapsClient_Validate() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewUrlMapsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ValidateUrlMapRequest.
	req := &computepb.ValidateUrlMapRequest{
		Project: "",
		UrlMap: "",
		UrlMapsValidateRequestResource: &computepb.UrlMapsValidateRequest{
			LoadBalancingSchemes: "",
			Resource: &computepb.UrlMap{
				CreationTimestamp: "",
				DefaultRouteAction: &computepb.HttpRouteAction{...}
				DefaultService: "",
				DefaultUrlRedirect: &computepb.HttpRedirectAction{...}
				Description: "",
				Fingerprint: "",
				HeaderAction: &computepb.HttpHeaderAction{...}
				HostRules: &computepb.HostRule{...}
				Id: "",
				Kind: "",
				Name: "",
				PathMatchers: &computepb.PathMatcher{...}
				Region: "",
				SelfLink: "",
				Tests: &computepb.UrlMapTest{...}
			}
		}
	}
	resp, err := c.Validate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}