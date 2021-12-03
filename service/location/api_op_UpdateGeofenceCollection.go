// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified properties of a given geofence collection.
func (c *Client) UpdateGeofenceCollection(ctx context.Context, params *UpdateGeofenceCollectionInput, optFns ...func(*Options)) (*UpdateGeofenceCollectionOutput, error) {
	if params == nil {
		params = &UpdateGeofenceCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGeofenceCollection", params, optFns, c.addOperationUpdateGeofenceCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGeofenceCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGeofenceCollectionInput struct {

	// The name of the geofence collection to update.
	//
	// This member is required.
	CollectionName *string

	// Updates the description for the geofence collection.
	Description *string

	// Updates the pricing plan for the geofence collection. For more information about
	// each pricing plan option restrictions, see Amazon Location Service pricing
	// (https://aws.amazon.com/location/pricing/).
	PricingPlan types.PricingPlan

	// Updates the data provider for the geofence collection. A required value for the
	// following pricing plans: MobileAssetTracking| MobileAssetManagement For more
	// information about data providers
	// (https://aws.amazon.com/location/data-providers/) and pricing plans
	// (https://aws.amazon.com/location/pricing/), see the Amazon Location Service
	// product page. This can only be updated when updating the PricingPlan in the same
	// request. Amazon Location Service uses PricingPlanDataSource to calculate billing
	// for your geofence collection. Your data won't be shared with the data provider,
	// and will remain in your AWS account and Region unless you move it.
	PricingPlanDataSource *string

	noSmithyDocumentSerde
}

type UpdateGeofenceCollectionOutput struct {

	// The Amazon Resource Name (ARN) of the updated geofence collection. Used to
	// specify a resource across AWS.
	//
	// * Format example:
	// arn:aws:geo:region:account-id:geofence-collection/ExampleGeofenceCollection
	//
	// This member is required.
	CollectionArn *string

	// The name of the updated geofence collection.
	//
	// This member is required.
	CollectionName *string

	// The time when the geofence collection was last updated in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// This member is required.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGeofenceCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGeofenceCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGeofenceCollection{}, middleware.After)
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
	if err = addEndpointPrefix_opUpdateGeofenceCollectionMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateGeofenceCollectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGeofenceCollection(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateGeofenceCollectionMiddleware struct {
}

func (*endpointPrefix_opUpdateGeofenceCollectionMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateGeofenceCollectionMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "geofencing." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opUpdateGeofenceCollectionMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opUpdateGeofenceCollectionMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateGeofenceCollection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "UpdateGeofenceCollection",
	}
}
