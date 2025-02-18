// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package entityresolutioniface provides an interface to enable mocking the AWS EntityResolution service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package entityresolutioniface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/entityresolution"
)

// EntityResolutionAPI provides an interface to enable mocking the
// entityresolution.EntityResolution service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS EntityResolution.
//	func myFunc(svc entityresolutioniface.EntityResolutionAPI) bool {
//	    // Make svc.CreateMatchingWorkflow request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := entityresolution.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockEntityResolutionClient struct {
//	    entityresolutioniface.EntityResolutionAPI
//	}
//	func (m *mockEntityResolutionClient) CreateMatchingWorkflow(input *entityresolution.CreateMatchingWorkflowInput) (*entityresolution.CreateMatchingWorkflowOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockEntityResolutionClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type EntityResolutionAPI interface {
	CreateMatchingWorkflow(*entityresolution.CreateMatchingWorkflowInput) (*entityresolution.CreateMatchingWorkflowOutput, error)
	CreateMatchingWorkflowWithContext(aws.Context, *entityresolution.CreateMatchingWorkflowInput, ...request.Option) (*entityresolution.CreateMatchingWorkflowOutput, error)
	CreateMatchingWorkflowRequest(*entityresolution.CreateMatchingWorkflowInput) (*request.Request, *entityresolution.CreateMatchingWorkflowOutput)

	CreateSchemaMapping(*entityresolution.CreateSchemaMappingInput) (*entityresolution.CreateSchemaMappingOutput, error)
	CreateSchemaMappingWithContext(aws.Context, *entityresolution.CreateSchemaMappingInput, ...request.Option) (*entityresolution.CreateSchemaMappingOutput, error)
	CreateSchemaMappingRequest(*entityresolution.CreateSchemaMappingInput) (*request.Request, *entityresolution.CreateSchemaMappingOutput)

	DeleteMatchingWorkflow(*entityresolution.DeleteMatchingWorkflowInput) (*entityresolution.DeleteMatchingWorkflowOutput, error)
	DeleteMatchingWorkflowWithContext(aws.Context, *entityresolution.DeleteMatchingWorkflowInput, ...request.Option) (*entityresolution.DeleteMatchingWorkflowOutput, error)
	DeleteMatchingWorkflowRequest(*entityresolution.DeleteMatchingWorkflowInput) (*request.Request, *entityresolution.DeleteMatchingWorkflowOutput)

	DeleteSchemaMapping(*entityresolution.DeleteSchemaMappingInput) (*entityresolution.DeleteSchemaMappingOutput, error)
	DeleteSchemaMappingWithContext(aws.Context, *entityresolution.DeleteSchemaMappingInput, ...request.Option) (*entityresolution.DeleteSchemaMappingOutput, error)
	DeleteSchemaMappingRequest(*entityresolution.DeleteSchemaMappingInput) (*request.Request, *entityresolution.DeleteSchemaMappingOutput)

	GetMatchId(*entityresolution.GetMatchIdInput) (*entityresolution.GetMatchIdOutput, error)
	GetMatchIdWithContext(aws.Context, *entityresolution.GetMatchIdInput, ...request.Option) (*entityresolution.GetMatchIdOutput, error)
	GetMatchIdRequest(*entityresolution.GetMatchIdInput) (*request.Request, *entityresolution.GetMatchIdOutput)

	GetMatchingJob(*entityresolution.GetMatchingJobInput) (*entityresolution.GetMatchingJobOutput, error)
	GetMatchingJobWithContext(aws.Context, *entityresolution.GetMatchingJobInput, ...request.Option) (*entityresolution.GetMatchingJobOutput, error)
	GetMatchingJobRequest(*entityresolution.GetMatchingJobInput) (*request.Request, *entityresolution.GetMatchingJobOutput)

	GetMatchingWorkflow(*entityresolution.GetMatchingWorkflowInput) (*entityresolution.GetMatchingWorkflowOutput, error)
	GetMatchingWorkflowWithContext(aws.Context, *entityresolution.GetMatchingWorkflowInput, ...request.Option) (*entityresolution.GetMatchingWorkflowOutput, error)
	GetMatchingWorkflowRequest(*entityresolution.GetMatchingWorkflowInput) (*request.Request, *entityresolution.GetMatchingWorkflowOutput)

	GetSchemaMapping(*entityresolution.GetSchemaMappingInput) (*entityresolution.GetSchemaMappingOutput, error)
	GetSchemaMappingWithContext(aws.Context, *entityresolution.GetSchemaMappingInput, ...request.Option) (*entityresolution.GetSchemaMappingOutput, error)
	GetSchemaMappingRequest(*entityresolution.GetSchemaMappingInput) (*request.Request, *entityresolution.GetSchemaMappingOutput)

	ListMatchingJobs(*entityresolution.ListMatchingJobsInput) (*entityresolution.ListMatchingJobsOutput, error)
	ListMatchingJobsWithContext(aws.Context, *entityresolution.ListMatchingJobsInput, ...request.Option) (*entityresolution.ListMatchingJobsOutput, error)
	ListMatchingJobsRequest(*entityresolution.ListMatchingJobsInput) (*request.Request, *entityresolution.ListMatchingJobsOutput)

	ListMatchingJobsPages(*entityresolution.ListMatchingJobsInput, func(*entityresolution.ListMatchingJobsOutput, bool) bool) error
	ListMatchingJobsPagesWithContext(aws.Context, *entityresolution.ListMatchingJobsInput, func(*entityresolution.ListMatchingJobsOutput, bool) bool, ...request.Option) error

	ListMatchingWorkflows(*entityresolution.ListMatchingWorkflowsInput) (*entityresolution.ListMatchingWorkflowsOutput, error)
	ListMatchingWorkflowsWithContext(aws.Context, *entityresolution.ListMatchingWorkflowsInput, ...request.Option) (*entityresolution.ListMatchingWorkflowsOutput, error)
	ListMatchingWorkflowsRequest(*entityresolution.ListMatchingWorkflowsInput) (*request.Request, *entityresolution.ListMatchingWorkflowsOutput)

	ListMatchingWorkflowsPages(*entityresolution.ListMatchingWorkflowsInput, func(*entityresolution.ListMatchingWorkflowsOutput, bool) bool) error
	ListMatchingWorkflowsPagesWithContext(aws.Context, *entityresolution.ListMatchingWorkflowsInput, func(*entityresolution.ListMatchingWorkflowsOutput, bool) bool, ...request.Option) error

	ListSchemaMappings(*entityresolution.ListSchemaMappingsInput) (*entityresolution.ListSchemaMappingsOutput, error)
	ListSchemaMappingsWithContext(aws.Context, *entityresolution.ListSchemaMappingsInput, ...request.Option) (*entityresolution.ListSchemaMappingsOutput, error)
	ListSchemaMappingsRequest(*entityresolution.ListSchemaMappingsInput) (*request.Request, *entityresolution.ListSchemaMappingsOutput)

	ListSchemaMappingsPages(*entityresolution.ListSchemaMappingsInput, func(*entityresolution.ListSchemaMappingsOutput, bool) bool) error
	ListSchemaMappingsPagesWithContext(aws.Context, *entityresolution.ListSchemaMappingsInput, func(*entityresolution.ListSchemaMappingsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*entityresolution.ListTagsForResourceInput) (*entityresolution.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *entityresolution.ListTagsForResourceInput, ...request.Option) (*entityresolution.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*entityresolution.ListTagsForResourceInput) (*request.Request, *entityresolution.ListTagsForResourceOutput)

	StartMatchingJob(*entityresolution.StartMatchingJobInput) (*entityresolution.StartMatchingJobOutput, error)
	StartMatchingJobWithContext(aws.Context, *entityresolution.StartMatchingJobInput, ...request.Option) (*entityresolution.StartMatchingJobOutput, error)
	StartMatchingJobRequest(*entityresolution.StartMatchingJobInput) (*request.Request, *entityresolution.StartMatchingJobOutput)

	TagResource(*entityresolution.TagResourceInput) (*entityresolution.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *entityresolution.TagResourceInput, ...request.Option) (*entityresolution.TagResourceOutput, error)
	TagResourceRequest(*entityresolution.TagResourceInput) (*request.Request, *entityresolution.TagResourceOutput)

	UntagResource(*entityresolution.UntagResourceInput) (*entityresolution.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *entityresolution.UntagResourceInput, ...request.Option) (*entityresolution.UntagResourceOutput, error)
	UntagResourceRequest(*entityresolution.UntagResourceInput) (*request.Request, *entityresolution.UntagResourceOutput)

	UpdateMatchingWorkflow(*entityresolution.UpdateMatchingWorkflowInput) (*entityresolution.UpdateMatchingWorkflowOutput, error)
	UpdateMatchingWorkflowWithContext(aws.Context, *entityresolution.UpdateMatchingWorkflowInput, ...request.Option) (*entityresolution.UpdateMatchingWorkflowOutput, error)
	UpdateMatchingWorkflowRequest(*entityresolution.UpdateMatchingWorkflowInput) (*request.Request, *entityresolution.UpdateMatchingWorkflowOutput)
}

var _ EntityResolutionAPI = (*entityresolution.EntityResolution)(nil)
