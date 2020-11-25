// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// This API action puts tags on an Amazon S3 on Outposts bucket. To put tags on an
// S3 bucket, see PutBucketTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html) in
// the Amazon Simple Storage Service API. Sets the tags for an Outposts bucket. For
// more information, see Using Amazon S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in the
// Amazon Simple Storage Service Developer Guide. Use tags to organize your AWS
// bill to reflect your own cost structure. To do this, sign up to get your AWS
// account bill with tag key values included. Then, to see the cost of combined
// resources, organize your billing information according to resources with the
// same tag key values. For example, you can tag several resources with a specific
// application name, and then organize your billing information to see the total
// cost of that application across several services. For more information, see Cost
// Allocation and Tagging
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html).
// Within a bucket, if you add a tag that has the same key as an existing tag, the
// new value overwrites the old value. For more information, see Using Cost
// Allocation in Amazon S3 Bucket Tags
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/CostAllocTagging.html). To use
// this operation, you must have permissions to perform the
// s3outposts:PutBucketTagging action. The Outposts bucket owner has this
// permission by default and can grant this permission to others. For more
// information about permissions, see  Permissions Related to Bucket Subresource
// Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
// PutBucketTagging has the following special errors:
//
// * Error code:
// InvalidTagError
//
// * Description: The tag provided was not a valid tag. This error
// can occur if the tag did not pass input validation. For information about tag
// restrictions, see  User-Defined Tag Restrictions
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html)
// and  AWS-Generated Cost Allocation Tag Restrictions
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/aws-tag-restrictions.html).
//
// *
// Error code: MalformedXMLError
//
// * Description: The XML provided does not match
// the schema.
//
// * Error code: OperationAbortedError
//
// * Description: A conflicting
// conditional operation is currently in progress against this resource. Try
// again.
//
// * Error code: InternalError
//
// * Description: The service was unable to
// apply the provided tag to the bucket.
//
// All Amazon S3 on Outposts REST API
// requests for this action require an additional parameter of outpost-id to be
// passed with the request and an S3 on Outposts endpoint hostname prefix instead
// of s3-control. For an example of the request syntax for Amazon S3 on Outposts
// that uses the S3 on Outposts endpoint hostname prefix and the outpost-id derived
// using the access point ARN, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_PutBucketTagging.html#API_control_PutBucketTagging_Examples)
// section below. The following actions are related to PutBucketTagging:
//
// *
// GetBucketTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketTagging.html)
//
// *
// DeleteBucketTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketTagging.html)
func (c *Client) PutBucketTagging(ctx context.Context, params *PutBucketTaggingInput, optFns ...func(*Options)) (*PutBucketTaggingOutput, error) {
	if params == nil {
		params = &PutBucketTaggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketTagging", params, optFns, addOperationPutBucketTaggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketTaggingInput struct {

	// The AWS account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The Amazon Resource Name (ARN) of the bucket. For Amazon S3 on Outposts specify
	// the ARN of the bucket accessed in the format
	// arn:aws:s3-outposts:::outpost//bucket/. For example, to access the bucket
	// reports through outpost my-outpost owned by account 123456789012 in Region
	// us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string

	//
	//
	// This member is required.
	Tagging *types.Tagging
}

type PutBucketTaggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBucketTaggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketTagging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketTagging{}, middleware.After)
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
	if err = addEndpointPrefix_opPutBucketTaggingMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutBucketTaggingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketTagging(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketTaggingUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opPutBucketTaggingMiddleware struct {
}

func (*endpointPrefix_opPutBucketTaggingMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutBucketTaggingMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*PutBucketTaggingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opPutBucketTaggingMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opPutBucketTaggingMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opPutBucketTagging(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketTagging",
	}
}

func copyPutBucketTaggingInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*PutBucketTaggingInput)
	if !ok {
		return nil, fmt.Errorf("expect *PutBucketTaggingInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getPutBucketTaggingARNMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketTaggingInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setPutBucketTaggingARNMember(input interface{}, v string) error {
	in := input.(*PutBucketTaggingInput)
	in.Bucket = &v
	return nil
}
func backFillPutBucketTaggingAccountID(input interface{}, v string) error {
	in := input.(*PutBucketTaggingInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addPutBucketTaggingUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getPutBucketTaggingARNMember,
			BackfillAccountID: backFillPutBucketTaggingAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setPutBucketTaggingARNMember,
			CopyInput:         copyPutBucketTaggingInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}