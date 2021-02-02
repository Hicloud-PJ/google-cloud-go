// Copyright 2021 Google LLC
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

package datalabeling_test

import (
	"context"

	datalabeling "cloud.google.com/go/datalabeling/apiv1beta1"
	"google.golang.org/api/iterator"
	datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_CreateDataset() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.CreateDatasetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateDataset(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetDataset() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.GetDatasetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetDataset(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListDatasets() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.ListDatasetsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListDatasets(ctx, req)
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
	}
}

func ExampleClient_DeleteDataset() {
	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.DeleteDatasetRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteDataset(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ImportData() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.ImportDataRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ImportData(ctx, req)
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

func ExampleClient_ExportData() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.ExportDataRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ExportData(ctx, req)
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

func ExampleClient_GetDataItem() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.GetDataItemRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetDataItem(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListDataItems() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.ListDataItemsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListDataItems(ctx, req)
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
	}
}

func ExampleClient_GetAnnotatedDataset() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.GetAnnotatedDatasetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetAnnotatedDataset(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListAnnotatedDatasets() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.ListAnnotatedDatasetsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListAnnotatedDatasets(ctx, req)
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
	}
}

func ExampleClient_DeleteAnnotatedDataset() {
	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.DeleteAnnotatedDatasetRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteAnnotatedDataset(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_LabelImage() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.LabelImageRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.LabelImage(ctx, req)
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

func ExampleClient_LabelVideo() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.LabelVideoRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.LabelVideo(ctx, req)
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

func ExampleClient_LabelText() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.LabelTextRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.LabelText(ctx, req)
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

func ExampleClient_GetExample() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.GetExampleRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetExample(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListExamples() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.ListExamplesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListExamples(ctx, req)
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
	}
}

func ExampleClient_CreateAnnotationSpecSet() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.CreateAnnotationSpecSetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateAnnotationSpecSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetAnnotationSpecSet() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.GetAnnotationSpecSetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetAnnotationSpecSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListAnnotationSpecSets() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.ListAnnotationSpecSetsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListAnnotationSpecSets(ctx, req)
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
	}
}

func ExampleClient_DeleteAnnotationSpecSet() {
	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.DeleteAnnotationSpecSetRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteAnnotationSpecSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_CreateInstruction() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.CreateInstructionRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateInstruction(ctx, req)
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

func ExampleClient_GetInstruction() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.GetInstructionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetInstruction(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListInstructions() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.ListInstructionsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListInstructions(ctx, req)
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
	}
}

func ExampleClient_DeleteInstruction() {
	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.DeleteInstructionRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteInstruction(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetEvaluation() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.GetEvaluationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetEvaluation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_SearchEvaluations() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.SearchEvaluationsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.SearchEvaluations(ctx, req)
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
	}
}

func ExampleClient_SearchExampleComparisons() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.SearchExampleComparisonsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.SearchExampleComparisons(ctx, req)
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
	}
}

func ExampleClient_CreateEvaluationJob() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.CreateEvaluationJobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateEvaluationJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateEvaluationJob() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.UpdateEvaluationJobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateEvaluationJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetEvaluationJob() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.GetEvaluationJobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetEvaluationJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_PauseEvaluationJob() {
	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.PauseEvaluationJobRequest{
		// TODO: Fill request struct fields.
	}
	err = c.PauseEvaluationJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ResumeEvaluationJob() {
	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.ResumeEvaluationJobRequest{
		// TODO: Fill request struct fields.
	}
	err = c.ResumeEvaluationJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_DeleteEvaluationJob() {
	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.DeleteEvaluationJobRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteEvaluationJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ListEvaluationJobs() {
	// import datalabelingpb "google.golang.org/genproto/googleapis/cloud/datalabeling/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datalabelingpb.ListEvaluationJobsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListEvaluationJobs(ctx, req)
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
	}
}