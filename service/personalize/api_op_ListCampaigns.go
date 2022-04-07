// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of campaigns that use the given solution. When a solution is not
// specified, all the campaigns associated with the account are listed. The
// response provides the properties for each campaign, including the Amazon
// Resource Name (ARN). For more information on campaigns, see CreateCampaign
// (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateCampaign.html).
func (c *Client) ListCampaigns(ctx context.Context, params *ListCampaignsInput, optFns ...func(*Options)) (*ListCampaignsOutput, error) {
	if params == nil {
		params = &ListCampaignsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCampaigns", params, optFns, c.addOperationListCampaignsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCampaignsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCampaignsInput struct {

	// The maximum number of campaigns to return.
	MaxResults *int32

	// A token returned from the previous call to ListCampaigns
	// (https://docs.aws.amazon.com/personalize/latest/dg/API_ListCampaigns.html) for
	// getting the next set of campaigns (if they exist).
	NextToken *string

	// The Amazon Resource Name (ARN) of the solution to list the campaigns for. When a
	// solution is not specified, all the campaigns associated with the account are
	// listed.
	SolutionArn *string

	noSmithyDocumentSerde
}

type ListCampaignsOutput struct {

	// A list of the campaigns.
	Campaigns []types.CampaignSummary

	// A token for getting the next set of campaigns (if they exist).
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCampaignsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCampaigns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCampaigns{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCampaigns(options.Region), middleware.Before); err != nil {
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

// ListCampaignsAPIClient is a client that implements the ListCampaigns operation.
type ListCampaignsAPIClient interface {
	ListCampaigns(context.Context, *ListCampaignsInput, ...func(*Options)) (*ListCampaignsOutput, error)
}

var _ ListCampaignsAPIClient = (*Client)(nil)

// ListCampaignsPaginatorOptions is the paginator options for ListCampaigns
type ListCampaignsPaginatorOptions struct {
	// The maximum number of campaigns to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCampaignsPaginator is a paginator for ListCampaigns
type ListCampaignsPaginator struct {
	options   ListCampaignsPaginatorOptions
	client    ListCampaignsAPIClient
	params    *ListCampaignsInput
	nextToken *string
	firstPage bool
}

// NewListCampaignsPaginator returns a new ListCampaignsPaginator
func NewListCampaignsPaginator(client ListCampaignsAPIClient, params *ListCampaignsInput, optFns ...func(*ListCampaignsPaginatorOptions)) *ListCampaignsPaginator {
	if params == nil {
		params = &ListCampaignsInput{}
	}

	options := ListCampaignsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCampaignsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCampaignsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCampaigns page.
func (p *ListCampaignsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCampaignsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListCampaigns(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListCampaigns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListCampaigns",
	}
}
