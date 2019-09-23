// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Initiates the authorization request, as an administrator.
type AdminInitiateAuthInput struct {
	_ struct{} `type:"structure"`

	// The analytics metadata for collecting Amazon Pinpoint metrics for AdminInitiateAuth
	// calls.
	AnalyticsMetadata *AnalyticsMetadataType `type:"structure"`

	// The authentication flow for this call to execute. The API action will depend
	// on this value. For example:
	//
	//    * REFRESH_TOKEN_AUTH will take in a valid refresh token and return new
	//    tokens.
	//
	//    * USER_SRP_AUTH will take in USERNAME and SRP_A and return the SRP variables
	//    to be used for next challenge execution.
	//
	//    * USER_PASSWORD_AUTH will take in USERNAME and PASSWORD and return the
	//    next challenge or tokens.
	//
	// Valid values include:
	//
	//    * USER_SRP_AUTH: Authentication flow for the Secure Remote Password (SRP)
	//    protocol.
	//
	//    * REFRESH_TOKEN_AUTH/REFRESH_TOKEN: Authentication flow for refreshing
	//    the access token and ID token by supplying a valid refresh token.
	//
	//    * CUSTOM_AUTH: Custom authentication flow.
	//
	//    * ADMIN_NO_SRP_AUTH: Non-SRP authentication flow; you can pass in the
	//    USERNAME and PASSWORD directly if the flow is enabled for calling the
	//    app client.
	//
	//    * USER_PASSWORD_AUTH: Non-SRP authentication flow; USERNAME and PASSWORD
	//    are passed directly. If a user migration Lambda trigger is set, this flow
	//    will invoke the user migration Lambda if the USERNAME is not found in
	//    the user pool.
	//
	// AuthFlow is a required field
	AuthFlow AuthFlowType `type:"string" required:"true" enum:"true"`

	// The authentication parameters. These are inputs corresponding to the AuthFlow
	// that you are invoking. The required values depend on the value of AuthFlow:
	//
	//    * For USER_SRP_AUTH: USERNAME (required), SRP_A (required), SECRET_HASH
	//    (required if the app client is configured with a client secret), DEVICE_KEY
	//
	//    * For REFRESH_TOKEN_AUTH/REFRESH_TOKEN: REFRESH_TOKEN (required), SECRET_HASH
	//    (required if the app client is configured with a client secret), DEVICE_KEY
	//
	//    * For ADMIN_NO_SRP_AUTH: USERNAME (required), SECRET_HASH (if app client
	//    is configured with client secret), PASSWORD (required), DEVICE_KEY
	//
	//    * For CUSTOM_AUTH: USERNAME (required), SECRET_HASH (if app client is
	//    configured with client secret), DEVICE_KEY
	AuthParameters map[string]string `type:"map"`

	// The app client ID.
	//
	// ClientId is a required field
	ClientId *string `min:"1" type:"string" required:"true" sensitive:"true"`

	// This is a random key-value pair map which can contain any key and will be
	// passed to your PreAuthentication Lambda trigger as-is. It can be used to
	// implement additional validations around authentication.
	ClientMetadata map[string]string `type:"map"`

	// Contextual data such as the user's device fingerprint, IP address, or location
	// used for evaluating the risk of an unexpected event by Amazon Cognito advanced
	// security.
	ContextData *ContextDataType `type:"structure"`

	// The ID of the Amazon Cognito user pool.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AdminInitiateAuthInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AdminInitiateAuthInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AdminInitiateAuthInput"}
	if len(s.AuthFlow) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AuthFlow"))
	}

	if s.ClientId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientId"))
	}
	if s.ClientId != nil && len(*s.ClientId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientId", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}
	if s.ContextData != nil {
		if err := s.ContextData.Validate(); err != nil {
			invalidParams.AddNested("ContextData", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Initiates the authentication response, as an administrator.
type AdminInitiateAuthOutput struct {
	_ struct{} `type:"structure"`

	// The result of the authentication response. This is only returned if the caller
	// does not need to pass another challenge. If the caller does need to pass
	// another challenge before it gets tokens, ChallengeName, ChallengeParameters,
	// and Session are returned.
	AuthenticationResult *AuthenticationResultType `type:"structure"`

	// The name of the challenge which you are responding to with this call. This
	// is returned to you in the AdminInitiateAuth response if you need to pass
	// another challenge.
	//
	//    * MFA_SETUP: If MFA is required, users who do not have at least one of
	//    the MFA methods set up are presented with an MFA_SETUP challenge. The
	//    user must set up at least one MFA type to continue to authenticate.
	//
	//    * SELECT_MFA_TYPE: Selects the MFA type. Valid MFA options are SMS_MFA
	//    for text SMS MFA, and SOFTWARE_TOKEN_MFA for TOTP software token MFA.
	//
	//    * SMS_MFA: Next challenge is to supply an SMS_MFA_CODE, delivered via
	//    SMS.
	//
	//    * PASSWORD_VERIFIER: Next challenge is to supply PASSWORD_CLAIM_SIGNATURE,
	//    PASSWORD_CLAIM_SECRET_BLOCK, and TIMESTAMP after the client-side SRP calculations.
	//
	//    * CUSTOM_CHALLENGE: This is returned if your custom authentication flow
	//    determines that the user should pass another challenge before tokens are
	//    issued.
	//
	//    * DEVICE_SRP_AUTH: If device tracking was enabled on your user pool and
	//    the previous challenges were passed, this challenge is returned so that
	//    Amazon Cognito can start tracking this device.
	//
	//    * DEVICE_PASSWORD_VERIFIER: Similar to PASSWORD_VERIFIER, but for devices
	//    only.
	//
	//    * ADMIN_NO_SRP_AUTH: This is returned if you need to authenticate with
	//    USERNAME and PASSWORD directly. An app client must be enabled to use this
	//    flow.
	//
	//    * NEW_PASSWORD_REQUIRED: For users which are required to change their
	//    passwords after successful first login. This challenge should be passed
	//    with NEW_PASSWORD and any other required attributes.
	ChallengeName ChallengeNameType `type:"string" enum:"true"`

	// The challenge parameters. These are returned to you in the AdminInitiateAuth
	// response if you need to pass another challenge. The responses in this parameter
	// should be used to compute inputs to the next call (AdminRespondToAuthChallenge).
	//
	// All challenges require USERNAME and SECRET_HASH (if applicable).
	//
	// The value of the USER_ID_FOR_SRP attribute will be the user's actual username,
	// not an alias (such as email address or phone number), even if you specified
	// an alias in your call to AdminInitiateAuth. This is because, in the AdminRespondToAuthChallenge
	// API ChallengeResponses, the USERNAME attribute cannot be an alias.
	ChallengeParameters map[string]string `type:"map"`

	// The session which should be passed both ways in challenge-response calls
	// to the service. If AdminInitiateAuth or AdminRespondToAuthChallenge API call
	// determines that the caller needs to go through another challenge, they return
	// a session with other challenge parameters. This session should be passed
	// as it is to the next AdminRespondToAuthChallenge API call.
	Session *string `min:"20" type:"string"`
}

// String returns the string representation
func (s AdminInitiateAuthOutput) String() string {
	return awsutil.Prettify(s)
}

const opAdminInitiateAuth = "AdminInitiateAuth"

// AdminInitiateAuthRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Initiates the authentication flow, as an administrator.
//
// Requires developer credentials.
//
//    // Example sending a request using AdminInitiateAuthRequest.
//    req := client.AdminInitiateAuthRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminInitiateAuth
func (c *Client) AdminInitiateAuthRequest(input *AdminInitiateAuthInput) AdminInitiateAuthRequest {
	op := &aws.Operation{
		Name:       opAdminInitiateAuth,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AdminInitiateAuthInput{}
	}

	req := c.newRequest(op, input, &AdminInitiateAuthOutput{})
	return AdminInitiateAuthRequest{Request: req, Input: input, Copy: c.AdminInitiateAuthRequest}
}

// AdminInitiateAuthRequest is the request type for the
// AdminInitiateAuth API operation.
type AdminInitiateAuthRequest struct {
	*aws.Request
	Input *AdminInitiateAuthInput
	Copy  func(*AdminInitiateAuthInput) AdminInitiateAuthRequest
}

// Send marshals and sends the AdminInitiateAuth API request.
func (r AdminInitiateAuthRequest) Send(ctx context.Context) (*AdminInitiateAuthResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminInitiateAuthResponse{
		AdminInitiateAuthOutput: r.Request.Data.(*AdminInitiateAuthOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminInitiateAuthResponse is the response type for the
// AdminInitiateAuth API operation.
type AdminInitiateAuthResponse struct {
	*AdminInitiateAuthOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminInitiateAuth request.
func (r *AdminInitiateAuthResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
