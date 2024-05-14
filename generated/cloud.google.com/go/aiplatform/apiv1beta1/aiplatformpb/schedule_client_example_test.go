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

func ExampleNewScheduleClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleScheduleClient_CreateSchedule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#CreateScheduleRequest.
	req := &aiplatformpb.CreateScheduleRequest{
		Parent: "",
		Schedule: &aiplatformpb.Schedule{
			Cron: "",
			CreatePipelineJobRequest: &aiplatformpb.CreatePipelineJobRequest{
				Parent: "",
				PipelineJob: &aiplatformpb.PipelineJob{...}
				PipelineJobId: "",
			}
			CreateModelMonitoringJobRequest: &aiplatformpb.CreateModelMonitoringJobRequest{
				Parent: "",
				ModelMonitoringJob: &aiplatformpb.ModelMonitoringJob{...}
				ModelMonitoringJobId: "",
			}
			CreateNotebookExecutionJobRequest: &aiplatformpb.CreateNotebookExecutionJobRequest{
				Parent: "",
				NotebookExecutionJob: &aiplatformpb.NotebookExecutionJob{...}
				NotebookExecutionJobId: "",
			}
			Name: "",
			DisplayName: "",
			StartTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			EndTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			MaxRunCount: "",
			StartedRunCount: "",
			State: "",
			CreateTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			UpdateTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			NextRunTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			LastPauseTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			LastResumeTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			MaxConcurrentRunCount: "",
			AllowQueueing: "",
			CatchUp: "",
			LastScheduledRunResponse: "",
		}
	}
	resp, err := c.CreateSchedule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleScheduleClient_DeleteSchedule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#DeleteScheduleRequest.
	req := &aiplatformpb.DeleteScheduleRequest{
		Name: "",
	}
	op, err := c.DeleteSchedule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleScheduleClient_GetSchedule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#GetScheduleRequest.
	req := &aiplatformpb.GetScheduleRequest{
		Name: "",
	}
	resp, err := c.GetSchedule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleScheduleClient_ListSchedules() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#ListSchedulesRequest.
	req := &aiplatformpb.ListSchedulesRequest{
		Parent: "",
		Filter: "",
		PageSize: "",
		PageToken: "",
		OrderBy: "",
	}
	it := c.ListSchedules(ctx, req)
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
		_ = it.Response.(*aiplatformpb.ListSchedulesResponse)
	}
}

func ExampleScheduleClient_PauseSchedule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#PauseScheduleRequest.
	req := &aiplatformpb.PauseScheduleRequest{
		Name: "",
	}
	err = c.PauseSchedule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleScheduleClient_ResumeSchedule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#ResumeScheduleRequest.
	req := &aiplatformpb.ResumeScheduleRequest{
		Name: "",
		CatchUp: "",
	}
	err = c.ResumeSchedule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleScheduleClient_UpdateSchedule() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#UpdateScheduleRequest.
	req := &aiplatformpb.UpdateScheduleRequest{
		Schedule: &aiplatformpb.Schedule{
			Cron: "",
			CreatePipelineJobRequest: &aiplatformpb.CreatePipelineJobRequest{
				Parent: "",
				PipelineJob: &aiplatformpb.PipelineJob{...}
				PipelineJobId: "",
			}
			CreateModelMonitoringJobRequest: &aiplatformpb.CreateModelMonitoringJobRequest{
				Parent: "",
				ModelMonitoringJob: &aiplatformpb.ModelMonitoringJob{...}
				ModelMonitoringJobId: "",
			}
			CreateNotebookExecutionJobRequest: &aiplatformpb.CreateNotebookExecutionJobRequest{
				Parent: "",
				NotebookExecutionJob: &aiplatformpb.NotebookExecutionJob{...}
				NotebookExecutionJobId: "",
			}
			Name: "",
			DisplayName: "",
			StartTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			EndTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			MaxRunCount: "",
			StartedRunCount: "",
			State: "",
			CreateTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			UpdateTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			NextRunTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			LastPauseTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			LastResumeTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			MaxConcurrentRunCount: "",
			AllowQueueing: "",
			CatchUp: "",
			LastScheduledRunResponse: "",
		}
		UpdateMask: &aiplatformpb.FieldMask{
			Paths: "",
		}
	}
	resp, err := c.UpdateSchedule(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleScheduleClient_GetLocation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
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

func ExampleScheduleClient_ListLocations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
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

func ExampleScheduleClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
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

func ExampleScheduleClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
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

func ExampleScheduleClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
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

func ExampleScheduleClient_CancelOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
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

func ExampleScheduleClient_DeleteOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
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

func ExampleScheduleClient_GetOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
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

func ExampleScheduleClient_ListOperations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
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

func ExampleScheduleClient_WaitOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewScheduleClient(ctx)
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