// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an API Gateway Authorizer.
type Authorizer struct {
	pulumi.CustomResourceState

	// The credentials required for the authorizer.
	// To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
	AuthorizerCredentials pulumi.StringPtrOutput `pulumi:"authorizerCredentials"`
	// The TTL of cached authorizer results in seconds.
	// Defaults to `300`.
	AuthorizerResultTtlInSeconds pulumi.IntPtrOutput `pulumi:"authorizerResultTtlInSeconds"`
	// The authorizer's Uniform Resource Identifier (URI).
	// This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
	// e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
	AuthorizerUri pulumi.StringPtrOutput `pulumi:"authorizerUri"`
	// The source of the identity in an incoming request.
	// Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
	IdentitySource pulumi.StringPtrOutput `pulumi:"identitySource"`
	// A validation expression for the incoming identity.
	// For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
	// against this expression, and will proceed if the token matches. If the token doesn't match,
	// the client receives a 401 Unauthorized response.
	IdentityValidationExpression pulumi.StringPtrOutput `pulumi:"identityValidationExpression"`
	// The name of the authorizer
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of the Amazon Cognito user pool ARNs.
	// Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
	ProviderArns pulumi.StringArrayOutput `pulumi:"providerArns"`
	// The ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
	// Defaults to `TOKEN`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewAuthorizer registers a new resource with the given unique name, arguments, and options.
func NewAuthorizer(ctx *pulumi.Context,
	name string, args *AuthorizerArgs, opts ...pulumi.ResourceOption) (*Authorizer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.RestApi == nil {
		return nil, errors.New("invalid value for required argument 'RestApi'")
	}
	var resource Authorizer
	err := ctx.RegisterResource("aws:apigateway/authorizer:Authorizer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizer gets an existing Authorizer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizerState, opts ...pulumi.ResourceOption) (*Authorizer, error) {
	var resource Authorizer
	err := ctx.ReadResource("aws:apigateway/authorizer:Authorizer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Authorizer resources.
type authorizerState struct {
	// The credentials required for the authorizer.
	// To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
	AuthorizerCredentials *string `pulumi:"authorizerCredentials"`
	// The TTL of cached authorizer results in seconds.
	// Defaults to `300`.
	AuthorizerResultTtlInSeconds *int `pulumi:"authorizerResultTtlInSeconds"`
	// The authorizer's Uniform Resource Identifier (URI).
	// This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
	// e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
	AuthorizerUri *string `pulumi:"authorizerUri"`
	// The source of the identity in an incoming request.
	// Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
	IdentitySource *string `pulumi:"identitySource"`
	// A validation expression for the incoming identity.
	// For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
	// against this expression, and will proceed if the token matches. If the token doesn't match,
	// the client receives a 401 Unauthorized response.
	IdentityValidationExpression *string `pulumi:"identityValidationExpression"`
	// The name of the authorizer
	Name *string `pulumi:"name"`
	// A list of the Amazon Cognito user pool ARNs.
	// Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
	ProviderArns []string `pulumi:"providerArns"`
	// The ID of the associated REST API
	RestApi *string `pulumi:"restApi"`
	// The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
	// Defaults to `TOKEN`.
	Type *string `pulumi:"type"`
}

type AuthorizerState struct {
	// The credentials required for the authorizer.
	// To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
	AuthorizerCredentials pulumi.StringPtrInput
	// The TTL of cached authorizer results in seconds.
	// Defaults to `300`.
	AuthorizerResultTtlInSeconds pulumi.IntPtrInput
	// The authorizer's Uniform Resource Identifier (URI).
	// This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
	// e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
	AuthorizerUri pulumi.StringPtrInput
	// The source of the identity in an incoming request.
	// Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
	IdentitySource pulumi.StringPtrInput
	// A validation expression for the incoming identity.
	// For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
	// against this expression, and will proceed if the token matches. If the token doesn't match,
	// the client receives a 401 Unauthorized response.
	IdentityValidationExpression pulumi.StringPtrInput
	// The name of the authorizer
	Name pulumi.StringPtrInput
	// A list of the Amazon Cognito user pool ARNs.
	// Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
	ProviderArns pulumi.StringArrayInput
	// The ID of the associated REST API
	RestApi pulumi.StringPtrInput
	// The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
	// Defaults to `TOKEN`.
	Type pulumi.StringPtrInput
}

func (AuthorizerState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizerState)(nil)).Elem()
}

type authorizerArgs struct {
	// The credentials required for the authorizer.
	// To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
	AuthorizerCredentials *string `pulumi:"authorizerCredentials"`
	// The TTL of cached authorizer results in seconds.
	// Defaults to `300`.
	AuthorizerResultTtlInSeconds *int `pulumi:"authorizerResultTtlInSeconds"`
	// The authorizer's Uniform Resource Identifier (URI).
	// This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
	// e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
	AuthorizerUri *string `pulumi:"authorizerUri"`
	// The source of the identity in an incoming request.
	// Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
	IdentitySource *string `pulumi:"identitySource"`
	// A validation expression for the incoming identity.
	// For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
	// against this expression, and will proceed if the token matches. If the token doesn't match,
	// the client receives a 401 Unauthorized response.
	IdentityValidationExpression *string `pulumi:"identityValidationExpression"`
	// The name of the authorizer
	Name *string `pulumi:"name"`
	// A list of the Amazon Cognito user pool ARNs.
	// Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
	ProviderArns []string `pulumi:"providerArns"`
	// The ID of the associated REST API
	RestApi interface{} `pulumi:"restApi"`
	// The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
	// Defaults to `TOKEN`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Authorizer resource.
type AuthorizerArgs struct {
	// The credentials required for the authorizer.
	// To specify an IAM Role for API Gateway to assume, use the IAM Role ARN.
	AuthorizerCredentials pulumi.StringPtrInput
	// The TTL of cached authorizer results in seconds.
	// Defaults to `300`.
	AuthorizerResultTtlInSeconds pulumi.IntPtrInput
	// The authorizer's Uniform Resource Identifier (URI).
	// This must be a well-formed Lambda function URI in the form of `arn:aws:apigateway:{region}:lambda:path/{service_api}`,
	// e.g. `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:012345678912:function:my-function/invocations`
	AuthorizerUri pulumi.StringPtrInput
	// The source of the identity in an incoming request.
	// Defaults to `method.request.header.Authorization`. For `REQUEST` type, this may be a comma-separated list of values, including headers, query string parameters and stage variables - e.g. `"method.request.header.SomeHeaderName,method.request.querystring.SomeQueryStringName,stageVariables.SomeStageVariableName"`
	IdentitySource pulumi.StringPtrInput
	// A validation expression for the incoming identity.
	// For `TOKEN` type, this value should be a regular expression. The incoming token from the client is matched
	// against this expression, and will proceed if the token matches. If the token doesn't match,
	// the client receives a 401 Unauthorized response.
	IdentityValidationExpression pulumi.StringPtrInput
	// The name of the authorizer
	Name pulumi.StringPtrInput
	// A list of the Amazon Cognito user pool ARNs.
	// Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`.
	ProviderArns pulumi.StringArrayInput
	// The ID of the associated REST API
	RestApi pulumi.Input
	// The type of the authorizer. Possible values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, or `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.
	// Defaults to `TOKEN`.
	Type pulumi.StringPtrInput
}

func (AuthorizerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizerArgs)(nil)).Elem()
}
