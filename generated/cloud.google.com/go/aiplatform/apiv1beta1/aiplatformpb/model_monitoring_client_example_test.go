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


package aiplatformpb_test

import (
	"context"

	aiplatformpb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

func ExampleNewModelMonitoringClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleModelMonitoringClient_CreateModelMonitor() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#CreateModelMonitorRequest.
	req := &aiplatformpb.CreateModelMonitorRequest{
		Parent: "",
		ModelMonitor: &aiplatformpb.ModelMonitor{
			TabularObjective: "",
			Name: "",
			DisplayName: "",
			ModelMonitoringTarget: "",
			TrainingDataset: &aiplatformpb.ModelMonitoringInput{
				ColumnizedDataset: "",
				BatchPredictionOutput: "",
				VertexEndpointLogs: "",
				TimeInterval: &intervalpb.Interval{...}
				TimeOffset: "",
			}
			NotificationSpec: &aiplatformpb.ModelMonitoringNotificationSpec{
				EmailConfig: "",
				EnableCloudLogging: "",
				NotificationChannelConfigs: "",
			}
			OutputSpec: &aiplatformpb.ModelMonitoringOutputSpec{
				GcsBaseDirectory: &aiplatformpb.GcsDestination{...}
			}
			ExplanationSpec: &aiplatformpb.ExplanationSpec{
				Parameters: &aiplatformpb.ExplanationParameters{...}
				Metadata: &aiplatformpb.ExplanationMetadata{...}
			}
			ModelMonitoringSchema: &aiplatformpb.ModelMonitoringSchema{
				FeatureFields: "",
				PredictionFields: "",
				GroundTruthFields: "",
			}
			CreateTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			UpdateTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
		}
		ModelMonitorId: "",
	}
	op, err := c.CreateModelMonitor(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelMonitoringClient_CreateModelMonitoringJob() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#CreateModelMonitoringJobRequest.
	req := &aiplatformpb.CreateModelMonitoringJobRequest{
		Parent: "",
		ModelMonitoringJob: &aiplatformpb.ModelMonitoringJob{
			Name: "",
			DisplayName: "",
			ModelMonitoringSpec: &aiplatformpb.ModelMonitoringSpec{
				ObjectiveSpec: &aiplatformpb.ModelMonitoringObjectiveSpec{...}
				NotificationSpec: &aiplatformpb.ModelMonitoringNotificationSpec{...}
				OutputSpec: &aiplatformpb.ModelMonitoringOutputSpec{...}
			}
			CreateTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			UpdateTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			State: "",
			Schedule: "",
			JobExecutionDetail: &aiplatformpb.ModelMonitoringJobExecutionDetail{
				BaselineDatasets: "",
				TargetDatasets: "",
				ObjectiveStatus: "",
				Error: &statuspb.Status{...}
			}
			ScheduleTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
		}
		ModelMonitoringJobId: "",
	}
	resp, err := c.CreateModelMonitoringJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelMonitoringClient_DeleteModelMonitor() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#DeleteModelMonitorRequest.
	req := &aiplatformpb.DeleteModelMonitorRequest{
		Name: "",
		Force: "",
	}
	op, err := c.DeleteModelMonitor(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleModelMonitoringClient_DeleteModelMonitoringJob() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#DeleteModelMonitoringJobRequest.
	req := &aiplatformpb.DeleteModelMonitoringJobRequest{
		Name: "",
	}
	op, err := c.DeleteModelMonitoringJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleModelMonitoringClient_GetModelMonitor() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#GetModelMonitorRequest.
	req := &aiplatformpb.GetModelMonitorRequest{
		Name: "",
	}
	resp, err := c.GetModelMonitor(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelMonitoringClient_GetModelMonitoringJob() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#GetModelMonitoringJobRequest.
	req := &aiplatformpb.GetModelMonitoringJobRequest{
		Name: "",
	}
	resp, err := c.GetModelMonitoringJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelMonitoringClient_ListModelMonitoringJobs() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#ListModelMonitoringJobsRequest.
	req := &aiplatformpb.ListModelMonitoringJobsRequest{
		Parent: "",
		Filter: "",
		PageSize: "",
		PageToken: "",
		ReadMask: &aiplatformpb.FieldMask{
			Paths: "",
		}
	}
	it := c.ListModelMonitoringJobs(ctx, req)
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
		_ = it.Response.(*aiplatformpb.ListModelMonitoringJobsResponse)
	}
}

func ExampleModelMonitoringClient_ListModelMonitors() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#ListModelMonitorsRequest.
	req := &aiplatformpb.ListModelMonitorsRequest{
		Parent: "",
		Filter: "",
		PageSize: "",
		PageToken: "",
		ReadMask: &aiplatformpb.FieldMask{
			Paths: "",
		}
	}
	it := c.ListModelMonitors(ctx, req)
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
		_ = it.Response.(*aiplatformpb.ListModelMonitorsResponse)
	}
}

func ExampleModelMonitoringClient_SearchModelMonitoringAlerts() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#SearchModelMonitoringAlertsRequest.
	req := &aiplatformpb.SearchModelMonitoringAlertsRequest{
		ModelMonitor: "",
		ModelMonitoringJob: "",
		AlertTimeInterval: &aiplatformpb.Interval{
			StartTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			EndTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
		}
		StatsName: "",
		ObjectiveType: "",
		PageSize: "",
		PageToken: "",
	}
	it := c.SearchModelMonitoringAlerts(ctx, req)
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
		_ = it.Response.(*aiplatformpb.SearchModelMonitoringAlertsResponse)
	}
}

func ExampleModelMonitoringClient_SearchModelMonitoringStats() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#SearchModelMonitoringStatsRequest.
	req := &aiplatformpb.SearchModelMonitoringStatsRequest{
		ModelMonitor: "",
		StatsFilter: &aiplatformpb.SearchModelMonitoringStatsFilter{
			TabularStatsFilter: "",
		}
		TimeInterval: &aiplatformpb.Interval{
			StartTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			EndTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
		}
		PageSize: "",
		PageToken: "",
	}
	it := c.SearchModelMonitoringStats(ctx, req)
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
		_ = it.Response.(*aiplatformpb.SearchModelMonitoringStatsResponse)
	}
}

func ExampleModelMonitoringClient_UpdateModelMonitor() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#UpdateModelMonitorRequest.
	req := &aiplatformpb.UpdateModelMonitorRequest{
		ModelMonitor: &aiplatformpb.ModelMonitor{
			TabularObjective: "",
			Name: "",
			DisplayName: "",
			ModelMonitoringTarget: "",
			TrainingDataset: &aiplatformpb.ModelMonitoringInput{
				ColumnizedDataset: "",
				BatchPredictionOutput: "",
				VertexEndpointLogs: "",
				TimeInterval: &intervalpb.Interval{...}
				TimeOffset: "",
			}
			NotificationSpec: &aiplatformpb.ModelMonitoringNotificationSpec{
				EmailConfig: "",
				EnableCloudLogging: "",
				NotificationChannelConfigs: "",
			}
			OutputSpec: &aiplatformpb.ModelMonitoringOutputSpec{
				GcsBaseDirectory: &aiplatformpb.GcsDestination{...}
			}
			ExplanationSpec: &aiplatformpb.ExplanationSpec{
				Parameters: &aiplatformpb.ExplanationParameters{...}
				Metadata: &aiplatformpb.ExplanationMetadata{...}
			}
			ModelMonitoringSchema: &aiplatformpb.ModelMonitoringSchema{
				FeatureFields: "",
				PredictionFields: "",
				GroundTruthFields: "",
			}
			CreateTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			UpdateTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
		}
		UpdateMask: &aiplatformpb.FieldMask{
			Paths: "",
		}
	}
	op, err := c.UpdateModelMonitor(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelMonitoringClient_GetLocation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#GetLocationRequest.
	req := &locationpb.GetLocationRequest{
		Name: "",
	}
	resp, err := c.GetLocation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelMonitoringClient_ListLocations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#ListLocationsRequest.
	req := &locationpb.ListLocationsRequest{
		Name: "",
		Filter: "",
		PageSize: "",
		PageToken: "",
	}
	it := c.ListLocations(ctx, req)
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
		_ = it.Response.(*locationpb.ListLocationsResponse)
	}
}

func ExampleModelMonitoringClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/iam/apiv1/iampb#GetIamPolicyRequest.
	req := &iampb.GetIamPolicyRequest{
		Resource: "",
		Options: "",
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelMonitoringClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/iam/apiv1/iampb#SetIamPolicyRequest.
	req := &iampb.SetIamPolicyRequest{
		Resource: "",
		Policy: "",
		UpdateMask: &iampb.FieldMask{
			Paths: "",
		}
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelMonitoringClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/iam/apiv1/iampb#TestIamPermissionsRequest.
	req := &iampb.TestIamPermissionsRequest{
		Resource: "",
		Permissions: "",
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelMonitoringClient_CancelOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#CancelOperationRequest.
	req := &longrunningpb.CancelOperationRequest{
		Name: "",
	}
	err = c.CancelOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleModelMonitoringClient_DeleteOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#DeleteOperationRequest.
	req := &longrunningpb.DeleteOperationRequest{
		Name: "",
	}
	err = c.DeleteOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleModelMonitoringClient_GetOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#GetOperationRequest.
	req := &longrunningpb.GetOperationRequest{
		Name: "",
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelMonitoringClient_ListOperations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#ListOperationsRequest.
	req := &longrunningpb.ListOperationsRequest{
		Name: "",
		Filter: "",
		PageSize: "",
		PageToken: "",
	}
	it := c.ListOperations(ctx, req)
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
		_ = it.Response.(*longrunningpb.ListOperationsResponse)
	}
}

func ExampleModelMonitoringClient_WaitOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewModelMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#WaitOperationRequest.
	req := &longrunningpb.WaitOperationRequest{
		Name: "",
		Timeout: &longrunningpb.Duration{
			Seconds: "",
			Nanos: "",
		}
	}
	resp, err := c.WaitOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
