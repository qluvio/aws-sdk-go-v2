// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Sends the forgot device request, as an administrator.
type AdminForgetDeviceInput struct {
	_ struct{} `type:"structure"`

	// The device key.
	//
	// DeviceKey is a required field
	DeviceKey *string `min:"1" type:"string" required:"true"`

	// The user pool ID.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The user name.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s AdminForgetDeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AdminForgetDeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AdminForgetDeviceInput"}

	if s.DeviceKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceKey"))
	}
	if s.DeviceKey != nil && len(*s.DeviceKey) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeviceKey", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AdminForgetDeviceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AdminForgetDeviceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAdminForgetDevice = "AdminForgetDevice"

// AdminForgetDeviceRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Forgets the device, as an administrator.
//
// Requires developer credentials.
//
//    // Example sending a request using AdminForgetDeviceRequest.
//    req := client.AdminForgetDeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminForgetDevice
func (c *Client) AdminForgetDeviceRequest(input *AdminForgetDeviceInput) AdminForgetDeviceRequest {
	op := &aws.Operation{
		Name:       opAdminForgetDevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AdminForgetDeviceInput{}
	}

	req := c.newRequest(op, input, &AdminForgetDeviceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AdminForgetDeviceRequest{Request: req, Input: input, Copy: c.AdminForgetDeviceRequest}
}

// AdminForgetDeviceRequest is the request type for the
// AdminForgetDevice API operation.
type AdminForgetDeviceRequest struct {
	*aws.Request
	Input *AdminForgetDeviceInput
	Copy  func(*AdminForgetDeviceInput) AdminForgetDeviceRequest
}

// Send marshals and sends the AdminForgetDevice API request.
func (r AdminForgetDeviceRequest) Send(ctx context.Context) (*AdminForgetDeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminForgetDeviceResponse{
		AdminForgetDeviceOutput: r.Request.Data.(*AdminForgetDeviceOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminForgetDeviceResponse is the response type for the
// AdminForgetDevice API operation.
type AdminForgetDeviceResponse struct {
	*AdminForgetDeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminForgetDevice request.
func (r *AdminForgetDeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
