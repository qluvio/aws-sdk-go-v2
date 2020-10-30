// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the runtime configuration of a thing.
func (c *Client) UpdateThingRuntimeConfiguration(ctx context.Context, params *UpdateThingRuntimeConfigurationInput, optFns ...func(*Options)) (*UpdateThingRuntimeConfigurationOutput, error) {
	if params == nil {
		params = &UpdateThingRuntimeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateThingRuntimeConfiguration", params, optFns, addOperationUpdateThingRuntimeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateThingRuntimeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateThingRuntimeConfigurationInput struct {

	// The thing name.
	//
	// This member is required.
	ThingName *string

	// Configuration for telemetry service.
	TelemetryConfiguration *types.TelemetryConfigurationUpdate
}

type UpdateThingRuntimeConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateThingRuntimeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateThingRuntimeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateThingRuntimeConfiguration{}, middleware.After)
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
	addOpUpdateThingRuntimeConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateThingRuntimeConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateThingRuntimeConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "UpdateThingRuntimeConfiguration",
	}
}