// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartExecutionInput struct {
	_ struct{} `type:"structure"`

	// The string that contains the JSON input data for the execution, for example:
	//
	// "input": "{\"first_name\" : \"test\"}"
	//
	// If you don't include any JSON input data, you still must include the two
	// braces, for example: "input": "{}"
	Input *string `locationName:"input" type:"string" sensitive:"true"`

	// The name of the execution. This name must be unique for your AWS account,
	// region, and state machine for 90 days. For more information, see Limits Related
	// to State Machine Executions (https://docs.aws.amazon.com/step-functions/latest/dg/limits.html#service-limits-state-machine-executions)
	// in the AWS Step Functions Developer Guide.
	//
	// A name must not contain:
	//
	//    * white space
	//
	//    * brackets < > { } [ ]
	//
	//    * wildcard characters ? *
	//
	//    * special characters " # % \ ^ | ~ ` $ & , ; : /
	//
	//    * control characters (U+0000-001F, U+007F-009F)
	Name *string `locationName:"name" min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the state machine to execute.
	//
	// StateMachineArn is a required field
	StateMachineArn *string `locationName:"stateMachineArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartExecutionInput"}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.StateMachineArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("StateMachineArn"))
	}
	if s.StateMachineArn != nil && len(*s.StateMachineArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StateMachineArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartExecutionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that identifies the execution.
	//
	// ExecutionArn is a required field
	ExecutionArn *string `locationName:"executionArn" min:"1" type:"string" required:"true"`

	// The date the execution is started.
	//
	// StartDate is a required field
	StartDate *time.Time `locationName:"startDate" type:"timestamp" required:"true"`
}

// String returns the string representation
func (s StartExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartExecution = "StartExecution"

// StartExecutionRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Starts a state machine execution.
//
// StartExecution is idempotent. If StartExecution is called with the same name
// and input as a running execution, the call will succeed and return the same
// response as the original request. If the execution is closed or if the input
// is different, it will return a 400 ExecutionAlreadyExists error. Names can
// be reused after 90 days.
//
//    // Example sending a request using StartExecutionRequest.
//    req := client.StartExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/StartExecution
func (c *Client) StartExecutionRequest(input *StartExecutionInput) StartExecutionRequest {
	op := &aws.Operation{
		Name:       opStartExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartExecutionInput{}
	}

	req := c.newRequest(op, input, &StartExecutionOutput{})
	return StartExecutionRequest{Request: req, Input: input, Copy: c.StartExecutionRequest}
}

// StartExecutionRequest is the request type for the
// StartExecution API operation.
type StartExecutionRequest struct {
	*aws.Request
	Input *StartExecutionInput
	Copy  func(*StartExecutionInput) StartExecutionRequest
}

// Send marshals and sends the StartExecution API request.
func (r StartExecutionRequest) Send(ctx context.Context) (*StartExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartExecutionResponse{
		StartExecutionOutput: r.Request.Data.(*StartExecutionOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartExecutionResponse is the response type for the
// StartExecution API operation.
type StartExecutionResponse struct {
	*StartExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartExecution request.
func (r *StartExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
