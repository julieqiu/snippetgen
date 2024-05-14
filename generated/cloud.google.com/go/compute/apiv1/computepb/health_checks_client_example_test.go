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

func ExampleNewHealthChecksClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewHealthChecksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleHealthChecksClient_AggregatedList() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewHealthChecksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#AggregatedListHealthChecksRequest.
	req := &computepb.AggregatedListHealthChecksRequest{
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
		_ = it.Response.(*computepb.HealthChecksAggregatedList)
	}
}

func ExampleHealthChecksClient_Delete() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewHealthChecksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#DeleteHealthCheckRequest.
	req := &computepb.DeleteHealthCheckRequest{
		HealthCheck: "",
		Project: "",
		RequestId: "",
	}
	resp, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHealthChecksClient_Get() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewHealthChecksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#GetHealthCheckRequest.
	req := &computepb.GetHealthCheckRequest{
		HealthCheck: "",
		Project: "",
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHealthChecksClient_Insert() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewHealthChecksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#InsertHealthCheckRequest.
	req := &computepb.InsertHealthCheckRequest{
		HealthCheckResource: &computepb.HealthCheck{
			CheckIntervalSec: "",
			CreationTimestamp: "",
			Description: "",
			GrpcHealthCheck: &computepb.GRPCHealthCheck{
				GrpcServiceName: "",
				Port: "",
				PortName: "",
				PortSpecification: "",
			}
			HealthyThreshold: "",
			Http2HealthCheck: &computepb.HTTP2HealthCheck{
				Host: "",
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				RequestPath: "",
				Response: "",
			}
			HttpHealthCheck: &computepb.HTTPHealthCheck{
				Host: "",
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				RequestPath: "",
				Response: "",
			}
			HttpsHealthCheck: &computepb.HTTPSHealthCheck{
				Host: "",
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				RequestPath: "",
				Response: "",
			}
			Id: "",
			Kind: "",
			LogConfig: &computepb.HealthCheckLogConfig{
				Enable: "",
			}
			Name: "",
			Region: "",
			SelfLink: "",
			SslHealthCheck: &computepb.SSLHealthCheck{
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				Request: "",
				Response: "",
			}
			TcpHealthCheck: &computepb.TCPHealthCheck{
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				Request: "",
				Response: "",
			}
			TimeoutSec: "",
			Type: "",
			UnhealthyThreshold: "",
		}
		Project: "",
		RequestId: "",
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHealthChecksClient_List() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewHealthChecksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#ListHealthChecksRequest.
	req := &computepb.ListHealthChecksRequest{
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
		_ = it.Response.(*computepb.HealthCheckList)
	}
}

func ExampleHealthChecksClient_Patch() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewHealthChecksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#PatchHealthCheckRequest.
	req := &computepb.PatchHealthCheckRequest{
		HealthCheck: "",
		HealthCheckResource: &computepb.HealthCheck{
			CheckIntervalSec: "",
			CreationTimestamp: "",
			Description: "",
			GrpcHealthCheck: &computepb.GRPCHealthCheck{
				GrpcServiceName: "",
				Port: "",
				PortName: "",
				PortSpecification: "",
			}
			HealthyThreshold: "",
			Http2HealthCheck: &computepb.HTTP2HealthCheck{
				Host: "",
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				RequestPath: "",
				Response: "",
			}
			HttpHealthCheck: &computepb.HTTPHealthCheck{
				Host: "",
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				RequestPath: "",
				Response: "",
			}
			HttpsHealthCheck: &computepb.HTTPSHealthCheck{
				Host: "",
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				RequestPath: "",
				Response: "",
			}
			Id: "",
			Kind: "",
			LogConfig: &computepb.HealthCheckLogConfig{
				Enable: "",
			}
			Name: "",
			Region: "",
			SelfLink: "",
			SslHealthCheck: &computepb.SSLHealthCheck{
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				Request: "",
				Response: "",
			}
			TcpHealthCheck: &computepb.TCPHealthCheck{
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				Request: "",
				Response: "",
			}
			TimeoutSec: "",
			Type: "",
			UnhealthyThreshold: "",
		}
		Project: "",
		RequestId: "",
	}
	resp, err := c.Patch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleHealthChecksClient_Update() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := computepb.NewHealthChecksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/compute/apiv1/computepb#UpdateHealthCheckRequest.
	req := &computepb.UpdateHealthCheckRequest{
		HealthCheck: "",
		HealthCheckResource: &computepb.HealthCheck{
			CheckIntervalSec: "",
			CreationTimestamp: "",
			Description: "",
			GrpcHealthCheck: &computepb.GRPCHealthCheck{
				GrpcServiceName: "",
				Port: "",
				PortName: "",
				PortSpecification: "",
			}
			HealthyThreshold: "",
			Http2HealthCheck: &computepb.HTTP2HealthCheck{
				Host: "",
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				RequestPath: "",
				Response: "",
			}
			HttpHealthCheck: &computepb.HTTPHealthCheck{
				Host: "",
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				RequestPath: "",
				Response: "",
			}
			HttpsHealthCheck: &computepb.HTTPSHealthCheck{
				Host: "",
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				RequestPath: "",
				Response: "",
			}
			Id: "",
			Kind: "",
			LogConfig: &computepb.HealthCheckLogConfig{
				Enable: "",
			}
			Name: "",
			Region: "",
			SelfLink: "",
			SslHealthCheck: &computepb.SSLHealthCheck{
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				Request: "",
				Response: "",
			}
			TcpHealthCheck: &computepb.TCPHealthCheck{
				Port: "",
				PortName: "",
				PortSpecification: "",
				ProxyHeader: "",
				Request: "",
				Response: "",
			}
			TimeoutSec: "",
			Type: "",
			UnhealthyThreshold: "",
		}
		Project: "",
		RequestId: "",
	}
	resp, err := c.Update(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
