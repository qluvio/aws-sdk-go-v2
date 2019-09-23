// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CopyObjectInput struct {
	_ struct{} `type:"structure"`

	// The canned ACL to apply to the object.
	ACL ObjectCannedACL `location:"header" locationName:"x-amz-acl" type:"string" enum:"true"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`

	// Specifies presentational information for the object.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field.
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`

	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`

	// A standard MIME type describing the format of the object data.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// The name of the source bucket and key name of the source object, separated
	// by a slash (/). Must be URL-encoded.
	//
	// CopySource is a required field
	CopySource *string `location:"header" locationName:"x-amz-copy-source" type:"string" required:"true"`

	// Copies the object if its entity tag (ETag) matches the specified tag.
	CopySourceIfMatch *string `location:"header" locationName:"x-amz-copy-source-if-match" type:"string"`

	// Copies the object if it has been modified since the specified time.
	CopySourceIfModifiedSince *time.Time `location:"header" locationName:"x-amz-copy-source-if-modified-since" type:"timestamp"`

	// Copies the object if its entity tag (ETag) is different than the specified
	// ETag.
	CopySourceIfNoneMatch *string `location:"header" locationName:"x-amz-copy-source-if-none-match" type:"string"`

	// Copies the object if it hasn't been modified since the specified time.
	CopySourceIfUnmodifiedSince *time.Time `location:"header" locationName:"x-amz-copy-source-if-unmodified-since" type:"timestamp"`

	// Specifies the algorithm to use when decrypting the source object (e.g., AES256).
	CopySourceSSECustomerAlgorithm *string `location:"header" locationName:"x-amz-copy-source-server-side-encryption-customer-algorithm" type:"string"`

	// Specifies the customer-provided encryption key for Amazon S3 to use to decrypt
	// the source object. The encryption key provided in this header must be one
	// that was used when the source object was created.
	CopySourceSSECustomerKey *string `location:"header" locationName:"x-amz-copy-source-server-side-encryption-customer-key" type:"string" sensitive:"true"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure the encryption
	// key was transmitted without error.
	CopySourceSSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-copy-source-server-side-encryption-customer-key-MD5" type:"string"`

	// The date and time at which the object is no longer cacheable.
	Expires *time.Time `location:"header" locationName:"Expires" type:"timestamp"`

	// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
	GrantFullControl *string `location:"header" locationName:"x-amz-grant-full-control" type:"string"`

	// Allows grantee to read the object data and its metadata.
	GrantRead *string `location:"header" locationName:"x-amz-grant-read" type:"string"`

	// Allows grantee to read the object ACL.
	GrantReadACP *string `location:"header" locationName:"x-amz-grant-read-acp" type:"string"`

	// Allows grantee to write the ACL for the applicable object.
	GrantWriteACP *string `location:"header" locationName:"x-amz-grant-write-acp" type:"string"`

	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// A map of metadata to store with the object in S3.
	Metadata map[string]string `location:"headers" locationName:"x-amz-meta-" type:"map"`

	// Specifies whether the metadata is copied from the source object or replaced
	// with metadata provided in the request.
	MetadataDirective MetadataDirective `location:"header" locationName:"x-amz-metadata-directive" type:"string" enum:"true"`

	// Specifies whether you want to apply a Legal Hold to the copied object.
	ObjectLockLegalHoldStatus ObjectLockLegalHoldStatus `location:"header" locationName:"x-amz-object-lock-legal-hold" type:"string" enum:"true"`

	// The object lock mode that you want to apply to the copied object.
	ObjectLockMode ObjectLockMode `location:"header" locationName:"x-amz-object-lock-mode" type:"string" enum:"true"`

	// The date and time when you want the copied object's object lock to expire.
	ObjectLockRetainUntilDate *time.Time `location:"header" locationName:"x-amz-object-lock-retain-until-date" type:"timestamp" timestampFormat:"iso8601"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// Specifies the algorithm to use to when encrypting the object (e.g., AES256).
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	// data. This value is used to store the object and then it is discarded; Amazon
	// does not store the encryption key. The key must be appropriate for use with
	// the algorithm specified in the x-amz-server-side​-encryption​-customer-algorithm
	// header.
	SSECustomerKey *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key" type:"string" sensitive:"true"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure the encryption
	// key was transmitted without error.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// Specifies the AWS KMS Encryption Context to use for object encryption. The
	// value of this header is a base64-encoded UTF-8 string holding JSON with the
	// encryption context key-value pairs.
	SSEKMSEncryptionContext *string `location:"header" locationName:"x-amz-server-side-encryption-context" type:"string" sensitive:"true"`

	// Specifies the AWS KMS key ID to use for object encryption. All GET and PUT
	// requests for an object protected by AWS KMS will fail if not made via SSL
	// or using SigV4. Documentation on configuring any of the officially supported
	// AWS SDKs and CLI can be found at http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingAWSSDK.html#specify-signature-version
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`

	// The type of storage to use for the object. Defaults to 'STANDARD'.
	StorageClass StorageClass `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"true"`

	// The tag-set for the object destination object this value must be used in
	// conjunction with the TaggingDirective. The tag-set must be encoded as URL
	// Query parameters
	Tagging *string `location:"header" locationName:"x-amz-tagging" type:"string"`

	// Specifies whether the object tag-set are copied from the source object or
	// replaced with tag-set provided in the request.
	TaggingDirective TaggingDirective `location:"header" locationName:"x-amz-tagging-directive" type:"string" enum:"true"`

	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata.
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}

// String returns the string representation
func (s CopyObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyObjectInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.CopySource == nil {
		invalidParams.Add(aws.NewErrParamRequired("CopySource"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *CopyObjectInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

func (s *CopyObjectInput) getCopySourceSSECustomerKey() (v string) {
	if s.CopySourceSSECustomerKey == nil {
		return v
	}
	return *s.CopySourceSSECustomerKey
}

func (s *CopyObjectInput) getSSECustomerKey() (v string) {
	if s.SSECustomerKey == nil {
		return v
	}
	return *s.SSECustomerKey
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CopyObjectInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.ACL) > 0 {
		v := s.ACL

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-acl", v, metadata)
	}
	if s.CacheControl != nil {
		v := *s.CacheControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Cache-Control", protocol.StringValue(v), metadata)
	}
	if s.ContentDisposition != nil {
		v := *s.ContentDisposition

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Disposition", protocol.StringValue(v), metadata)
	}
	if s.ContentEncoding != nil {
		v := *s.ContentEncoding

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Encoding", protocol.StringValue(v), metadata)
	}
	if s.ContentLanguage != nil {
		v := *s.ContentLanguage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Language", protocol.StringValue(v), metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue(v), metadata)
	}
	if s.CopySource != nil {
		v := *s.CopySource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source", protocol.StringValue(v), metadata)
	}
	if s.CopySourceIfMatch != nil {
		v := *s.CopySourceIfMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-if-match", protocol.StringValue(v), metadata)
	}
	if s.CopySourceIfModifiedSince != nil {
		v := *s.CopySourceIfModifiedSince

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-if-modified-since",
			protocol.TimeValue{V: v, Format: protocol.RFC822TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.CopySourceIfNoneMatch != nil {
		v := *s.CopySourceIfNoneMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-if-none-match", protocol.StringValue(v), metadata)
	}
	if s.CopySourceIfUnmodifiedSince != nil {
		v := *s.CopySourceIfUnmodifiedSince

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-if-unmodified-since",
			protocol.TimeValue{V: v, Format: protocol.RFC822TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.CopySourceSSECustomerAlgorithm != nil {
		v := *s.CopySourceSSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.CopySourceSSECustomerKey != nil {
		v := *s.CopySourceSSECustomerKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-server-side-encryption-customer-key", protocol.StringValue(v), metadata)
	}
	if s.CopySourceSSECustomerKeyMD5 != nil {
		v := *s.CopySourceSSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if s.Expires != nil {
		v := *s.Expires

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Expires",
			protocol.TimeValue{V: v, Format: protocol.RFC822TimeFormatName, QuotedFormatTime: false}, metadata)
	}
	if s.GrantFullControl != nil {
		v := *s.GrantFullControl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-full-control", protocol.StringValue(v), metadata)
	}
	if s.GrantRead != nil {
		v := *s.GrantRead

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-read", protocol.StringValue(v), metadata)
	}
	if s.GrantReadACP != nil {
		v := *s.GrantReadACP

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-read-acp", protocol.StringValue(v), metadata)
	}
	if s.GrantWriteACP != nil {
		v := *s.GrantWriteACP

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-grant-write-acp", protocol.StringValue(v), metadata)
	}
	if len(s.MetadataDirective) > 0 {
		v := s.MetadataDirective

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-metadata-directive", v, metadata)
	}
	if len(s.ObjectLockLegalHoldStatus) > 0 {
		v := s.ObjectLockLegalHoldStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-legal-hold", v, metadata)
	}
	if len(s.ObjectLockMode) > 0 {
		v := s.ObjectLockMode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-mode", v, metadata)
	}
	if s.ObjectLockRetainUntilDate != nil {
		v := *s.ObjectLockRetainUntilDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-object-lock-retain-until-date",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: false}, metadata)
	}
	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.SSECustomerAlgorithm != nil {
		v := *s.SSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKey != nil {
		v := *s.SSECustomerKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKeyMD5 != nil {
		v := *s.SSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSEncryptionContext != nil {
		v := *s.SSEKMSEncryptionContext

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-context", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSKeyId != nil {
		v := *s.SSEKMSKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-aws-kms-key-id", protocol.StringValue(v), metadata)
	}
	if len(s.ServerSideEncryption) > 0 {
		v := s.ServerSideEncryption

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption", v, metadata)
	}
	if len(s.StorageClass) > 0 {
		v := s.StorageClass

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-storage-class", v, metadata)
	}
	if s.Tagging != nil {
		v := *s.Tagging

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-tagging", protocol.StringValue(v), metadata)
	}
	if len(s.TaggingDirective) > 0 {
		v := s.TaggingDirective

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-tagging-directive", v, metadata)
	}
	if s.WebsiteRedirectLocation != nil {
		v := *s.WebsiteRedirectLocation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-website-redirect-location", protocol.StringValue(v), metadata)
	}
	if s.Metadata != nil {
		v := s.Metadata

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.HeadersTarget, "x-amz-meta-", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.StringValue(v1))
		}
		ms0.End()

	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	return nil
}

type CopyObjectOutput struct {
	_ struct{} `type:"structure" payload:"CopyObjectResult"`

	CopyObjectResult *CopyObjectResult `type:"structure"`

	CopySourceVersionId *string `location:"header" locationName:"x-amz-copy-source-version-id" type:"string"`

	// If the object expiration is configured, the response includes this header.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// If present, specifies the AWS KMS Encryption Context to use for object encryption.
	// The value of this header is a base64-encoded UTF-8 string holding JSON with
	// the encryption context key-value pairs.
	SSEKMSEncryptionContext *string `location:"header" locationName:"x-amz-server-side-encryption-context" type:"string" sensitive:"true"`

	// If present, specifies the ID of the AWS Key Management Service (KMS) master
	// encryption key that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption ServerSideEncryption `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"true"`

	// Version ID of the newly created copy.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`
}

// String returns the string representation
func (s CopyObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CopyObjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CopySourceVersionId != nil {
		v := *s.CopySourceVersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-copy-source-version-id", protocol.StringValue(v), metadata)
	}
	if s.Expiration != nil {
		v := *s.Expiration

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-expiration", protocol.StringValue(v), metadata)
	}
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	if s.SSECustomerAlgorithm != nil {
		v := *s.SSECustomerAlgorithm

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-algorithm", protocol.StringValue(v), metadata)
	}
	if s.SSECustomerKeyMD5 != nil {
		v := *s.SSECustomerKeyMD5

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-customer-key-MD5", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSEncryptionContext != nil {
		v := *s.SSEKMSEncryptionContext

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-context", protocol.StringValue(v), metadata)
	}
	if s.SSEKMSKeyId != nil {
		v := *s.SSEKMSKeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption-aws-kms-key-id", protocol.StringValue(v), metadata)
	}
	if len(s.ServerSideEncryption) > 0 {
		v := s.ServerSideEncryption

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-server-side-encryption", v, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-version-id", protocol.StringValue(v), metadata)
	}
	if s.CopyObjectResult != nil {
		v := s.CopyObjectResult

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CopyObjectResult", v, metadata)
	}
	return nil
}

const opCopyObject = "CopyObject"

// CopyObjectRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Creates a copy of an object that is already stored in Amazon S3.
//
//    // Example sending a request using CopyObjectRequest.
//    req := client.CopyObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/CopyObject
func (c *Client) CopyObjectRequest(input *CopyObjectInput) CopyObjectRequest {
	op := &aws.Operation{
		Name:       opCopyObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &CopyObjectInput{}
	}

	req := c.newRequest(op, input, &CopyObjectOutput{})
	return CopyObjectRequest{Request: req, Input: input, Copy: c.CopyObjectRequest}
}

// CopyObjectRequest is the request type for the
// CopyObject API operation.
type CopyObjectRequest struct {
	*aws.Request
	Input *CopyObjectInput
	Copy  func(*CopyObjectInput) CopyObjectRequest
}

// Send marshals and sends the CopyObject API request.
func (r CopyObjectRequest) Send(ctx context.Context) (*CopyObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyObjectResponse{
		CopyObjectOutput: r.Request.Data.(*CopyObjectOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyObjectResponse is the response type for the
// CopyObject API operation.
type CopyObjectResponse struct {
	*CopyObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyObject request.
func (r *CopyObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
