// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DecryptInput struct {
	_ struct{} `type:"structure"`

	// Ciphertext to be decrypted. The blob includes metadata.
	//
	// CiphertextBlob is automatically base64 encoded/decoded by the SDK.
	//
	// CiphertextBlob is a required field
	CiphertextBlob []byte `min:"1" type:"blob" required:"true"`

	// The encryption context. If this was specified in the Encrypt function, it
	// must be specified here or the decryption operation will fail. For more information,
	// see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context).
	EncryptionContext map[string]string `type:"map"`

	// A list of grant tokens.
	//
	// For more information, see Grant Tokens (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []string `type:"list"`
}

// String returns the string representation
func (s DecryptInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DecryptInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DecryptInput"}

	if s.CiphertextBlob == nil {
		invalidParams.Add(aws.NewErrParamRequired("CiphertextBlob"))
	}
	if s.CiphertextBlob != nil && len(s.CiphertextBlob) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CiphertextBlob", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DecryptOutput struct {
	_ struct{} `type:"structure"`

	// ARN of the key used to perform the decryption. This value is returned if
	// no errors are encountered during the operation.
	KeyId *string `min:"1" type:"string"`

	// Decrypted plaintext data. When you use the HTTP API or the AWS CLI, the value
	// is Base64-encoded. Otherwise, it is not encoded.
	//
	// Plaintext is automatically base64 encoded/decoded by the SDK.
	Plaintext []byte `min:"1" type:"blob" sensitive:"true"`
}

// String returns the string representation
func (s DecryptOutput) String() string {
	return awsutil.Prettify(s)
}

const opDecrypt = "Decrypt"

// DecryptRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Decrypts ciphertext. Ciphertext is plaintext that has been previously encrypted
// by using any of the following operations:
//
//    * GenerateDataKey
//
//    * GenerateDataKeyWithoutPlaintext
//
//    * Encrypt
//
// Whenever possible, use key policies to give users permission to call the
// Decrypt operation on the CMK, instead of IAM policies. Otherwise, you might
// create an IAM user policy that gives the user Decrypt permission on all CMKs.
// This user could decrypt ciphertext that was encrypted by CMKs in other accounts
// if the key policy for the cross-account CMK permits it. If you must use an
// IAM policy for Decrypt permissions, limit the user to particular CMKs or
// particular trusted accounts.
//
// The result of this operation varies with the key state of the CMK. For details,
// see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using DecryptRequest.
//    req := client.DecryptRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/Decrypt
func (c *Client) DecryptRequest(input *DecryptInput) DecryptRequest {
	op := &aws.Operation{
		Name:       opDecrypt,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DecryptInput{}
	}

	req := c.newRequest(op, input, &DecryptOutput{})
	return DecryptRequest{Request: req, Input: input, Copy: c.DecryptRequest}
}

// DecryptRequest is the request type for the
// Decrypt API operation.
type DecryptRequest struct {
	*aws.Request
	Input *DecryptInput
	Copy  func(*DecryptInput) DecryptRequest
}

// Send marshals and sends the Decrypt API request.
func (r DecryptRequest) Send(ctx context.Context) (*DecryptResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DecryptResponse{
		DecryptOutput: r.Request.Data.(*DecryptOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DecryptResponse is the response type for the
// Decrypt API operation.
type DecryptResponse struct {
	*DecryptOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// Decrypt request.
func (r *DecryptResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
