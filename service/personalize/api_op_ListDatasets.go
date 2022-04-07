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

// Returns the list of datasets contained in the given dataset group. The response
// provides the properties for each dataset, including the Amazon Resource Name
// (ARN). For more information on datasets, see CreateDataset
// (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataset.html).
func (c *Client) ListDatasets(ctx context.Context, params *ListDatasetsInput, optFns ...func(*Options)) (*ListDatasetsOutput, error) {
	if params == nil {
		params = &ListDatasetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatasets", params, optFns, c.addOperationListDatasetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatasetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatasetsInput struct {

	// The Amazon Resource Name (ARN) of the dataset group that contains the datasets
	// to list.
	DatasetGroupArn *string

	// The maximum number of datasets to return.
	MaxResults *int32

	// A token returned from the previous call to ListDatasetImportJobs for getting the
	// next set of dataset import jobs (if they exist).
	NextToken *string

	noSmithyDocumentSerde
}

type ListDatasetsOutput struct {

	// An array of Dataset objects. Each object provides metadata information.
	Datasets []types.DatasetSummary

	// A token for getting the next set of datasets (if they exist).
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDatasetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDatasets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDatasets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDatasets(options.Region), middleware.Before); err != nil {
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

// ListDatasetsAPIClient is a client that implements the ListDatasets operation.
type ListDatasetsAPIClient interface {
	ListDatasets(context.Context, *ListDatasetsInput, ...func(*Options)) (*ListDatasetsOutput, error)
}

var _ ListDatasetsAPIClient = (*Client)(nil)

// ListDatasetsPaginatorOptions is the paginator options for ListDatasets
type ListDatasetsPaginatorOptions struct {
	// The maximum number of datasets to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDatasetsPaginator is a paginator for ListDatasets
type ListDatasetsPaginator struct {
	options   ListDatasetsPaginatorOptions
	client    ListDatasetsAPIClient
	params    *ListDatasetsInput
	nextToken *string
	firstPage bool
}

// NewListDatasetsPaginator returns a new ListDatasetsPaginator
func NewListDatasetsPaginator(client ListDatasetsAPIClient, params *ListDatasetsInput, optFns ...func(*ListDatasetsPaginatorOptions)) *ListDatasetsPaginator {
	if params == nil {
		params = &ListDatasetsInput{}
	}

	options := ListDatasetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDatasetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDatasetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDatasets page.
func (p *ListDatasetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDatasetsOutput, error) {
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

	result, err := p.client.ListDatasets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDatasets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListDatasets",
	}
}
