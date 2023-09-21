// Code generated by smithy-go-codegen DO NOT EDIT.

package ssooidc

import (
	"context"
<<<<<<< HEAD
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
=======
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
>>>>>>> main
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates device authorization by requesting a pair of verification codes from
// the authorization service.
func (c *Client) StartDeviceAuthorization(ctx context.Context, params *StartDeviceAuthorizationInput, optFns ...func(*Options)) (*StartDeviceAuthorizationOutput, error) {
	if params == nil {
		params = &StartDeviceAuthorizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDeviceAuthorization", params, optFns, c.addOperationStartDeviceAuthorizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDeviceAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDeviceAuthorizationInput struct {

	// The unique identifier string for the client that is registered with IAM
	// Identity Center. This value should come from the persisted result of the
	// RegisterClient API operation.
	//
	// This member is required.
	ClientId *string

	// A secret string that is generated for the client. This value should come from
	// the persisted result of the RegisterClient API operation.
	//
	// This member is required.
	ClientSecret *string

	// The URL for the AWS access portal. For more information, see Using the AWS
	// access portal (https://docs.aws.amazon.com/singlesignon/latest/userguide/using-the-portal.html)
	// in the IAM Identity Center User Guide.
	//
	// This member is required.
	StartUrl *string

	noSmithyDocumentSerde
}

type StartDeviceAuthorizationOutput struct {

	// The short-lived code that is used by the device when polling for a session
	// token.
	DeviceCode *string

	// Indicates the number of seconds in which the verification code will become
	// invalid.
	ExpiresIn int32

	// Indicates the number of seconds the client must wait between attempts when
	// polling for a session.
	Interval int32

	// A one-time user verification code. This is needed to authorize an in-use device.
	UserCode *string

	// The URI of the verification page that takes the userCode to authorize the
	// device.
	VerificationUri *string

	// An alternate URL that the client can use to automatically launch a browser.
	// This process skips the manual step in which the user visits the verification
	// page and enters their code.
	VerificationUriComplete *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDeviceAuthorizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartDeviceAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartDeviceAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
<<<<<<< HEAD
=======
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
>>>>>>> main
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
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
<<<<<<< HEAD
=======
	if err = addStartDeviceAuthorizationResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
>>>>>>> main
	if err = addOpStartDeviceAuthorizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDeviceAuthorization(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
<<<<<<< HEAD
=======
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
>>>>>>> main
	return nil
}

func newServiceMetadataMiddleware_opStartDeviceAuthorization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartDeviceAuthorization",
	}
}
<<<<<<< HEAD
=======

type opStartDeviceAuthorizationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opStartDeviceAuthorizationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartDeviceAuthorizationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "awsssooidc"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "awsssooidc"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("awsssooidc")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addStartDeviceAuthorizationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opStartDeviceAuthorizationResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
>>>>>>> main
