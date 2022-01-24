// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/outposts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an order for an Outpost.
func (c *Client) CreateOrder(ctx context.Context, params *CreateOrderInput, optFns ...func(*Options)) (*CreateOrderOutput, error) {
	if params == nil {
		params = &CreateOrderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOrder", params, optFns, c.addOperationCreateOrderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOrderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOrderInput struct {

	// The line items that make up the order.
	//
	// This member is required.
	LineItems []types.LineItemRequest

	// The ID or the Amazon Resource Name (ARN) of the Outpost.
	//
	// This member is required.
	OutpostIdentifier *string

	// The payment option for the order.
	//
	// This member is required.
	PaymentOption types.PaymentOption

	// The payment terms for the order.
	PaymentTerm types.PaymentTerm

	noSmithyDocumentSerde
}

type CreateOrderOutput struct {

	// Information about this order.
	Order *types.Order

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateOrderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateOrder{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateOrder{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateOrderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOrder(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateOrder(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "outposts",
		OperationName: "CreateOrder",
	}
}