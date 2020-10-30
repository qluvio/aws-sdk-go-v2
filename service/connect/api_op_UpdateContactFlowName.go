// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The name of the contact flow.
func (c *Client) UpdateContactFlowName(ctx context.Context, params *UpdateContactFlowNameInput, optFns ...func(*Options)) (*UpdateContactFlowNameOutput, error) {
	if params == nil {
		params = &UpdateContactFlowNameInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateContactFlowName", params, optFns, addOperationUpdateContactFlowNameMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateContactFlowNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContactFlowNameInput struct {

	// The identifier of the contact flow.
	//
	// This member is required.
	ContactFlowId *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The description of the contact flow.
	Description *string

	// The name of the contact flow.
	Name *string
}

type UpdateContactFlowNameOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateContactFlowNameMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateContactFlowName{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateContactFlowName{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateContactFlowNameValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContactFlowName(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateContactFlowName(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateContactFlowName",
	}
}