// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/RemoveTagsFromResourceMessage
type RemoveTagsFromResourceInput struct {
	_ struct{} `type:"structure"`

	// >The Amazon Resource Name (ARN) of the AWS DMS resource the tag is to be
	// removed from.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`

	// The tag key (name) of the tag to be removed.
	//
	// TagKeys is a required field
	TagKeys []string `type:"list" required:"true"`
}

// String returns the string representation
func (s RemoveTagsFromResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveTagsFromResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveTagsFromResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/RemoveTagsFromResourceResponse
type RemoveTagsFromResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveTagsFromResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opRemoveTagsFromResource = "RemoveTagsFromResource"

// RemoveTagsFromResourceRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Removes metadata tags from a DMS resource.
//
//    // Example sending a request using RemoveTagsFromResourceRequest.
//    req := client.RemoveTagsFromResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/RemoveTagsFromResource
func (c *Client) RemoveTagsFromResourceRequest(input *RemoveTagsFromResourceInput) RemoveTagsFromResourceRequest {
	op := &aws.Operation{
		Name:       opRemoveTagsFromResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveTagsFromResourceInput{}
	}

	req := c.newRequest(op, input, &RemoveTagsFromResourceOutput{})
	return RemoveTagsFromResourceRequest{Request: req, Input: input, Copy: c.RemoveTagsFromResourceRequest}
}

// RemoveTagsFromResourceRequest is the request type for the
// RemoveTagsFromResource API operation.
type RemoveTagsFromResourceRequest struct {
	*aws.Request
	Input *RemoveTagsFromResourceInput
	Copy  func(*RemoveTagsFromResourceInput) RemoveTagsFromResourceRequest
}

// Send marshals and sends the RemoveTagsFromResource API request.
func (r RemoveTagsFromResourceRequest) Send(ctx context.Context) (*RemoveTagsFromResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveTagsFromResourceResponse{
		RemoveTagsFromResourceOutput: r.Request.Data.(*RemoveTagsFromResourceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveTagsFromResourceResponse is the response type for the
// RemoveTagsFromResource API operation.
type RemoveTagsFromResourceResponse struct {
	*RemoveTagsFromResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveTagsFromResource request.
func (r *RemoveTagsFromResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}