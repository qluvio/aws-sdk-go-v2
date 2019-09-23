// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to update user attributes.
type UpdateUserAttributesInput struct {
	_ struct{} `type:"structure"`

	// The access token for the request to update user attributes.
	//
	// AccessToken is a required field
	AccessToken *string `type:"string" required:"true" sensitive:"true"`

	// An array of name-value pairs representing user attributes.
	//
	// For custom attributes, you must prepend the custom: prefix to the attribute
	// name.
	//
	// UserAttributes is a required field
	UserAttributes []AttributeType `type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateUserAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUserAttributesInput"}

	if s.AccessToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessToken"))
	}

	if s.UserAttributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserAttributes"))
	}
	if s.UserAttributes != nil {
		for i, v := range s.UserAttributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UserAttributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server for the request to update user attributes.
type UpdateUserAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The code delivery details list from the server for the request to update
	// user attributes.
	CodeDeliveryDetailsList []CodeDeliveryDetailsType `type:"list"`
}

// String returns the string representation
func (s UpdateUserAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateUserAttributes = "UpdateUserAttributes"

// UpdateUserAttributesRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Allows a user to update a specific attribute (one at a time).
//
//    // Example sending a request using UpdateUserAttributesRequest.
//    req := client.UpdateUserAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/UpdateUserAttributes
func (c *Client) UpdateUserAttributesRequest(input *UpdateUserAttributesInput) UpdateUserAttributesRequest {
	op := &aws.Operation{
		Name:       opUpdateUserAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateUserAttributesInput{}
	}

	req := c.newRequest(op, input, &UpdateUserAttributesOutput{})
	req.Config.Credentials = aws.AnonymousCredentials
	return UpdateUserAttributesRequest{Request: req, Input: input, Copy: c.UpdateUserAttributesRequest}
}

// UpdateUserAttributesRequest is the request type for the
// UpdateUserAttributes API operation.
type UpdateUserAttributesRequest struct {
	*aws.Request
	Input *UpdateUserAttributesInput
	Copy  func(*UpdateUserAttributesInput) UpdateUserAttributesRequest
}

// Send marshals and sends the UpdateUserAttributes API request.
func (r UpdateUserAttributesRequest) Send(ctx context.Context) (*UpdateUserAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserAttributesResponse{
		UpdateUserAttributesOutput: r.Request.Data.(*UpdateUserAttributesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserAttributesResponse is the response type for the
// UpdateUserAttributes API operation.
type UpdateUserAttributesResponse struct {
	*UpdateUserAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUserAttributes request.
func (r *UpdateUserAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
