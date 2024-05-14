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

func ExampleNewVizierClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleVizierClient_AddTrialMeasurement() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#AddTrialMeasurementRequest.
	req := &aiplatformpb.AddTrialMeasurementRequest{
		TrialName: "",
		Measurement: &aiplatformpb.Measurement{
			ElapsedDuration: &durationpb.Duration{
				Seconds: "",
				Nanos: "",
			}
			StepCount: "",
			Metrics: "",
		}
	}
	resp, err := c.AddTrialMeasurement(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleVizierClient_CheckTrialEarlyStoppingState() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#CheckTrialEarlyStoppingStateRequest.
	req := &aiplatformpb.CheckTrialEarlyStoppingStateRequest{
		TrialName: "",
	}
	op, err := c.CheckTrialEarlyStoppingState(ctx, req)
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

func ExampleVizierClient_CompleteTrial() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#CompleteTrialRequest.
	req := &aiplatformpb.CompleteTrialRequest{
		Name: "",
		FinalMeasurement: &aiplatformpb.Measurement{
			ElapsedDuration: &durationpb.Duration{
				Seconds: "",
				Nanos: "",
			}
			StepCount: "",
			Metrics: "",
		}
		TrialInfeasible: "",
		InfeasibleReason: "",
	}
	resp, err := c.CompleteTrial(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleVizierClient_CreateStudy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#CreateStudyRequest.
	req := &aiplatformpb.CreateStudyRequest{
		Parent: "",
		Study: &aiplatformpb.Study{
			Name: "",
			DisplayName: "",
			StudySpec: &aiplatformpb.StudySpec{
				DecayCurveStoppingSpec: "",
				MedianAutomatedStoppingSpec: "",
				ConvexStopConfig: "",
				ConvexAutomatedStoppingSpec: "",
				Metrics: "",
				Parameters: "",
				Algorithm: "",
				ObservationNoise: "",
				MeasurementSelectionType: "",
				TransferLearningConfig: "",
				StudyStoppingConfig: "",
			}
			State: "",
			CreateTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			InactiveReason: "",
		}
	}
	resp, err := c.CreateStudy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleVizierClient_CreateTrial() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#CreateTrialRequest.
	req := &aiplatformpb.CreateTrialRequest{
		Parent: "",
		Trial: &aiplatformpb.Trial{
			Name: "",
			Id: "",
			State: "",
			Parameters: "",
			FinalMeasurement: &aiplatformpb.Measurement{
				ElapsedDuration: &durationpb.Duration{...}
				StepCount: "",
				Metrics: "",
			}
			Measurements: &aiplatformpb.Measurement{
				ElapsedDuration: &durationpb.Duration{...}
				StepCount: "",
				Metrics: "",
			}
			StartTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			EndTime: &timestamppb.Timestamp{
				Seconds: "",
				Nanos: "",
			}
			ClientId: "",
			InfeasibleReason: "",
			CustomJob: "",
			WebAccessUris: "",
		}
	}
	resp, err := c.CreateTrial(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleVizierClient_DeleteStudy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#DeleteStudyRequest.
	req := &aiplatformpb.DeleteStudyRequest{
		Name: "",
	}
	err = c.DeleteStudy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleVizierClient_DeleteTrial() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#DeleteTrialRequest.
	req := &aiplatformpb.DeleteTrialRequest{
		Name: "",
	}
	err = c.DeleteTrial(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleVizierClient_GetStudy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#GetStudyRequest.
	req := &aiplatformpb.GetStudyRequest{
		Name: "",
	}
	resp, err := c.GetStudy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleVizierClient_GetTrial() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#GetTrialRequest.
	req := &aiplatformpb.GetTrialRequest{
		Name: "",
	}
	resp, err := c.GetTrial(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleVizierClient_ListOptimalTrials() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#ListOptimalTrialsRequest.
	req := &aiplatformpb.ListOptimalTrialsRequest{
		Parent: "",
	}
	resp, err := c.ListOptimalTrials(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleVizierClient_ListStudies() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#ListStudiesRequest.
	req := &aiplatformpb.ListStudiesRequest{
		Parent: "",
		PageToken: "",
		PageSize: "",
	}
	it := c.ListStudies(ctx, req)
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
		_ = it.Response.(*aiplatformpb.ListStudiesResponse)
	}
}

func ExampleVizierClient_ListTrials() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#ListTrialsRequest.
	req := &aiplatformpb.ListTrialsRequest{
		Parent: "",
		PageToken: "",
		PageSize: "",
	}
	it := c.ListTrials(ctx, req)
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
		_ = it.Response.(*aiplatformpb.ListTrialsResponse)
	}
}

func ExampleVizierClient_LookupStudy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#LookupStudyRequest.
	req := &aiplatformpb.LookupStudyRequest{
		Parent: "",
		DisplayName: "",
	}
	resp, err := c.LookupStudy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleVizierClient_StopTrial() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#StopTrialRequest.
	req := &aiplatformpb.StopTrialRequest{
		Name: "",
	}
	resp, err := c.StopTrial(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleVizierClient_SuggestTrials() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Fill request struct fields.
	// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb#SuggestTrialsRequest.
	req := &aiplatformpb.SuggestTrialsRequest{
		Parent: "",
		SuggestionCount: "",
		ClientId: "",
		Contexts: &aiplatformpb.TrialContext{
			Description: "",
			Parameters: "",
		}
	}
	op, err := c.SuggestTrials(ctx, req)
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

func ExampleVizierClient_GetLocation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
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

func ExampleVizierClient_ListLocations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
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

func ExampleVizierClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
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

func ExampleVizierClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
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

func ExampleVizierClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
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

func ExampleVizierClient_CancelOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
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

func ExampleVizierClient_DeleteOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
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

func ExampleVizierClient_GetOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
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

func ExampleVizierClient_ListOperations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
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

func ExampleVizierClient_WaitOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatformpb.NewVizierClient(ctx)
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
