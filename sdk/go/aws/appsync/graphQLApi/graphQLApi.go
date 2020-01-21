// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package graphQLApi

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/appsync/GraphQLApiAdditionalAuthenticationProvider"
	"https:/github.com/pulumi/pulumi-aws/appsync/GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig"
	"https:/github.com/pulumi/pulumi-aws/appsync/GraphQLApiAdditionalAuthenticationProviderUserPoolConfig"
	"https:/github.com/pulumi/pulumi-aws/appsync/GraphQLApiLogConfig"
	"https:/github.com/pulumi/pulumi-aws/appsync/GraphQLApiOpenidConnectConfig"
	"https:/github.com/pulumi/pulumi-aws/appsync/GraphQLApiUserPoolConfig"
)

// Provides an AppSync GraphQL API.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appsync_graphql_api.html.markdown.
type GraphQLApi struct {
	pulumi.CustomResourceState

	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders appsyncGraphQLApiAdditionalAuthenticationProvider.GraphQLApiAdditionalAuthenticationProviderArrayOutput `pulumi:"additionalAuthenticationProviders"`
	// The ARN
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
	AuthenticationType pulumi.StringOutput `pulumi:"authenticationType"`
	// Nested argument containing logging configuration. Defined below.
	LogConfig appsyncGraphQLApiLogConfig.GraphQLApiLogConfigPtrOutput `pulumi:"logConfig"`
	// A user-supplied name for the GraphqlApi.
	Name pulumi.StringOutput `pulumi:"name"`
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig appsyncGraphQLApiOpenidConnectConfig.GraphQLApiOpenidConnectConfigPtrOutput `pulumi:"openidConnectConfig"`
	// The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema pulumi.StringPtrOutput `pulumi:"schema"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
	Uris pulumi.StringMapOutput `pulumi:"uris"`
	// The Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig appsyncGraphQLApiUserPoolConfig.GraphQLApiUserPoolConfigPtrOutput `pulumi:"userPoolConfig"`
}

// NewGraphQLApi registers a new resource with the given unique name, arguments, and options.
func NewGraphQLApi(ctx *pulumi.Context,
	name string, args *GraphQLApiArgs, opts ...pulumi.ResourceOption) (*GraphQLApi, error) {
	if args == nil || args.AuthenticationType == nil {
		return nil, errors.New("missing required argument 'AuthenticationType'")
	}
	if args == nil {
		args = &GraphQLApiArgs{}
	}
	var resource GraphQLApi
	err := ctx.RegisterResource("aws:appsync/graphQLApi:GraphQLApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGraphQLApi gets an existing GraphQLApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGraphQLApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphQLApiState, opts ...pulumi.ResourceOption) (*GraphQLApi, error) {
	var resource GraphQLApi
	err := ctx.ReadResource("aws:appsync/graphQLApi:GraphQLApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GraphQLApi resources.
type graphQLApiState struct {
	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders []appsyncGraphQLApiAdditionalAuthenticationProvider.GraphQLApiAdditionalAuthenticationProvider `pulumi:"additionalAuthenticationProviders"`
	// The ARN
	Arn *string `pulumi:"arn"`
	// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
	AuthenticationType *string `pulumi:"authenticationType"`
	// Nested argument containing logging configuration. Defined below.
	LogConfig *appsyncGraphQLApiLogConfig.GraphQLApiLogConfig `pulumi:"logConfig"`
	// A user-supplied name for the GraphqlApi.
	Name *string `pulumi:"name"`
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig *appsyncGraphQLApiOpenidConnectConfig.GraphQLApiOpenidConnectConfig `pulumi:"openidConnectConfig"`
	// The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema *string `pulumi:"schema"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
	Uris map[string]string `pulumi:"uris"`
	// The Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig *appsyncGraphQLApiUserPoolConfig.GraphQLApiUserPoolConfig `pulumi:"userPoolConfig"`
}

type GraphQLApiState struct {
	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders appsyncGraphQLApiAdditionalAuthenticationProvider.GraphQLApiAdditionalAuthenticationProviderArrayInput
	// The ARN
	Arn pulumi.StringPtrInput
	// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
	AuthenticationType pulumi.StringPtrInput
	// Nested argument containing logging configuration. Defined below.
	LogConfig appsyncGraphQLApiLogConfig.GraphQLApiLogConfigPtrInput
	// A user-supplied name for the GraphqlApi.
	Name pulumi.StringPtrInput
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig appsyncGraphQLApiOpenidConnectConfig.GraphQLApiOpenidConnectConfigPtrInput
	// The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
	Uris pulumi.StringMapInput
	// The Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig appsyncGraphQLApiUserPoolConfig.GraphQLApiUserPoolConfigPtrInput
}

func (GraphQLApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQLApiState)(nil)).Elem()
}

type graphQLApiArgs struct {
	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders []appsyncGraphQLApiAdditionalAuthenticationProvider.GraphQLApiAdditionalAuthenticationProvider `pulumi:"additionalAuthenticationProviders"`
	// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
	AuthenticationType string `pulumi:"authenticationType"`
	// Nested argument containing logging configuration. Defined below.
	LogConfig *appsyncGraphQLApiLogConfig.GraphQLApiLogConfig `pulumi:"logConfig"`
	// A user-supplied name for the GraphqlApi.
	Name *string `pulumi:"name"`
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig *appsyncGraphQLApiOpenidConnectConfig.GraphQLApiOpenidConnectConfig `pulumi:"openidConnectConfig"`
	// The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema *string `pulumi:"schema"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig *appsyncGraphQLApiUserPoolConfig.GraphQLApiUserPoolConfig `pulumi:"userPoolConfig"`
}

// The set of arguments for constructing a GraphQLApi resource.
type GraphQLApiArgs struct {
	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProviders appsyncGraphQLApiAdditionalAuthenticationProvider.GraphQLApiAdditionalAuthenticationProviderArrayInput
	// The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
	AuthenticationType pulumi.StringInput
	// Nested argument containing logging configuration. Defined below.
	LogConfig appsyncGraphQLApiLogConfig.GraphQLApiLogConfigPtrInput
	// A user-supplied name for the GraphqlApi.
	Name pulumi.StringPtrInput
	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenidConnectConfig appsyncGraphQLApiOpenidConnectConfig.GraphQLApiOpenidConnectConfigPtrInput
	// The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
	Schema pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig appsyncGraphQLApiUserPoolConfig.GraphQLApiUserPoolConfigPtrInput
}

func (GraphQLApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQLApiArgs)(nil)).Elem()
}

