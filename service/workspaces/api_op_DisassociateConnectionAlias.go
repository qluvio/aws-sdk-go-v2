// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a connection alias from a directory. Disassociating a connection
// alias disables cross-Region redirection between two directories in different AWS
// Regions. For more information, see  Cross-Region Redirection for Amazon
// WorkSpaces
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html).
// Before performing this operation, call  DescribeConnectionAliases
// (https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeConnectionAliases.html)
// to make sure that the current state of the connection alias is CREATED.
func (c *Client) DisassociateConnectionAlias(ctx context.Context, params *DisassociateConnectionAliasInput, optFns ...func(*Options)) (*DisassociateConnectionAliasOutput, error) {
	if params == nil {
		params = &DisassociateConnectionAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateConnectionAlias", params, optFns, addOperationDisassociateConnectionAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateConnectionAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateConnectionAliasInput struct {

	// The identifier of the connection alias to disassociate.
	//
	// This member is required.
	AliasId *string
}

type DisassociateConnectionAliasOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateConnectionAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateConnectionAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateConnectionAlias{}, middleware.After)
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
	addOpDisassociateConnectionAliasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateConnectionAlias(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateConnectionAlias(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DisassociateConnectionAlias",
	}
}