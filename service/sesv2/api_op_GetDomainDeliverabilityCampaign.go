// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve all the deliverability data for a specific campaign. This data is
// available for a campaign only if the campaign sent email by using a domain that
// the Deliverability dashboard is enabled for.
func (c *Client) GetDomainDeliverabilityCampaign(ctx context.Context, params *GetDomainDeliverabilityCampaignInput, optFns ...func(*Options)) (*GetDomainDeliverabilityCampaignOutput, error) {
	if params == nil {
		params = &GetDomainDeliverabilityCampaignInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDomainDeliverabilityCampaign", params, optFns, addOperationGetDomainDeliverabilityCampaignMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDomainDeliverabilityCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieve all the deliverability data for a specific campaign. This data is
// available for a campaign only if the campaign sent email by using a domain that
// the Deliverability dashboard is enabled for (PutDeliverabilityDashboardOption
// operation).
type GetDomainDeliverabilityCampaignInput struct {

	// The unique identifier for the campaign. The Deliverability dashboard
	// automatically generates and assigns this identifier to a campaign.
	//
	// This member is required.
	CampaignId *string
}

// An object that contains all the deliverability data for a specific campaign.
// This data is available for a campaign only if the campaign sent email by using a
// domain that the Deliverability dashboard is enabled for.
type GetDomainDeliverabilityCampaignOutput struct {

	// An object that contains the deliverability data for the campaign.
	//
	// This member is required.
	DomainDeliverabilityCampaign *types.DomainDeliverabilityCampaign

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDomainDeliverabilityCampaignMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDomainDeliverabilityCampaign{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDomainDeliverabilityCampaign{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addOpGetDomainDeliverabilityCampaignValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomainDeliverabilityCampaign(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDomainDeliverabilityCampaign(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetDomainDeliverabilityCampaign",
	}
}