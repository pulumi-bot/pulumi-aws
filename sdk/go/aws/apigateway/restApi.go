// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an API Gateway REST API.
//
// > **Note:** Amazon API Gateway Version 1 resources are used for creating and deploying REST APIs. To create and deploy WebSocket and HTTP APIs, use Amazon API Gateway Version 2.
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigateway.NewRestApi(ctx, "myDemoAPI", &apigateway.RestApiArgs{
// 			Description: pulumi.String("This is my API for demonstration purposes"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Regional Endpoint Type
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigateway.NewRestApi(ctx, "example", &apigateway.RestApiArgs{
// 			EndpointConfiguration: &apigateway.RestApiEndpointConfigurationArgs{
// 				Types: pulumi.String("REGIONAL"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// `aws_api_gateway_rest_api` can be imported by using the REST API ID, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/restApi:RestApi example 12345abcde
// ```
type RestApi struct {
	pulumi.CustomResourceState

	// The source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER.
	ApiKeySource pulumi.StringPtrOutput `pulumi:"apiKeySource"`
	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes pulumi.StringArrayOutput `pulumi:"binaryMediaTypes"`
	// An OpenAPI specification that defines the set of routes and integrations to create as part of the REST API.
	Body pulumi.StringPtrOutput `pulumi:"body"`
	// The creation date of the REST API
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The description of the REST API
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Nested argument defining API endpoint configuration including endpoint type. Defined below.
	EndpointConfiguration RestApiEndpointConfigurationOutput `pulumi:"endpointConfiguration"`
	// The execution ARN part to be used in `lambdaPermission`'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
	ExecutionArn pulumi.StringOutput `pulumi:"executionArn"`
	// Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default).
	MinimumCompressionSize pulumi.IntPtrOutput `pulumi:"minimumCompressionSize"`
	// The name of the REST API
	Name pulumi.StringOutput `pulumi:"name"`
	// JSON formatted policy document that controls access to the API Gateway. This provider will only perform drift detection of its value when present in a configuration. It is recommended to use the `apigateway.RestApiPolicy` resource instead.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The resource ID of the REST API's root
	RootResourceId pulumi.StringOutput `pulumi:"rootResourceId"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewRestApi registers a new resource with the given unique name, arguments, and options.
func NewRestApi(ctx *pulumi.Context,
	name string, args *RestApiArgs, opts ...pulumi.ResourceOption) (*RestApi, error) {
	if args == nil {
		args = &RestApiArgs{}
	}

	var resource RestApi
	err := ctx.RegisterResource("aws:apigateway/restApi:RestApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRestApi gets an existing RestApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRestApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RestApiState, opts ...pulumi.ResourceOption) (*RestApi, error) {
	var resource RestApi
	err := ctx.ReadResource("aws:apigateway/restApi:RestApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RestApi resources.
type restApiState struct {
	// The source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER.
	ApiKeySource *string `pulumi:"apiKeySource"`
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes []string `pulumi:"binaryMediaTypes"`
	// An OpenAPI specification that defines the set of routes and integrations to create as part of the REST API.
	Body *string `pulumi:"body"`
	// The creation date of the REST API
	CreatedDate *string `pulumi:"createdDate"`
	// The description of the REST API
	Description *string `pulumi:"description"`
	// Nested argument defining API endpoint configuration including endpoint type. Defined below.
	EndpointConfiguration *RestApiEndpointConfiguration `pulumi:"endpointConfiguration"`
	// The execution ARN part to be used in `lambdaPermission`'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
	ExecutionArn *string `pulumi:"executionArn"`
	// Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default).
	MinimumCompressionSize *int `pulumi:"minimumCompressionSize"`
	// The name of the REST API
	Name *string `pulumi:"name"`
	// JSON formatted policy document that controls access to the API Gateway. This provider will only perform drift detection of its value when present in a configuration. It is recommended to use the `apigateway.RestApiPolicy` resource instead.
	Policy *string `pulumi:"policy"`
	// The resource ID of the REST API's root
	RootResourceId *string `pulumi:"rootResourceId"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type RestApiState struct {
	// The source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER.
	ApiKeySource pulumi.StringPtrInput
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes pulumi.StringArrayInput
	// An OpenAPI specification that defines the set of routes and integrations to create as part of the REST API.
	Body pulumi.StringPtrInput
	// The creation date of the REST API
	CreatedDate pulumi.StringPtrInput
	// The description of the REST API
	Description pulumi.StringPtrInput
	// Nested argument defining API endpoint configuration including endpoint type. Defined below.
	EndpointConfiguration RestApiEndpointConfigurationPtrInput
	// The execution ARN part to be used in `lambdaPermission`'s `sourceArn`
	// when allowing API Gateway to invoke a Lambda function,
	// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
	ExecutionArn pulumi.StringPtrInput
	// Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default).
	MinimumCompressionSize pulumi.IntPtrInput
	// The name of the REST API
	Name pulumi.StringPtrInput
	// JSON formatted policy document that controls access to the API Gateway. This provider will only perform drift detection of its value when present in a configuration. It is recommended to use the `apigateway.RestApiPolicy` resource instead.
	Policy pulumi.StringPtrInput
	// The resource ID of the REST API's root
	RootResourceId pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (RestApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*restApiState)(nil)).Elem()
}

type restApiArgs struct {
	// The source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER.
	ApiKeySource *string `pulumi:"apiKeySource"`
	// The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes []string `pulumi:"binaryMediaTypes"`
	// An OpenAPI specification that defines the set of routes and integrations to create as part of the REST API.
	Body *string `pulumi:"body"`
	// The description of the REST API
	Description *string `pulumi:"description"`
	// Nested argument defining API endpoint configuration including endpoint type. Defined below.
	EndpointConfiguration *RestApiEndpointConfiguration `pulumi:"endpointConfiguration"`
	// Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default).
	MinimumCompressionSize *int `pulumi:"minimumCompressionSize"`
	// The name of the REST API
	Name *string `pulumi:"name"`
	// JSON formatted policy document that controls access to the API Gateway. This provider will only perform drift detection of its value when present in a configuration. It is recommended to use the `apigateway.RestApiPolicy` resource instead.
	Policy *string `pulumi:"policy"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a RestApi resource.
type RestApiArgs struct {
	// The source of the API key for requests. Valid values are HEADER (default) and AUTHORIZER.
	ApiKeySource pulumi.StringPtrInput
	// The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes pulumi.StringArrayInput
	// An OpenAPI specification that defines the set of routes and integrations to create as part of the REST API.
	Body pulumi.StringPtrInput
	// The description of the REST API
	Description pulumi.StringPtrInput
	// Nested argument defining API endpoint configuration including endpoint type. Defined below.
	EndpointConfiguration RestApiEndpointConfigurationPtrInput
	// Minimum response size to compress for the REST API. Integer between -1 and 10485760 (10MB). Setting a value greater than -1 will enable compression, -1 disables compression (default).
	MinimumCompressionSize pulumi.IntPtrInput
	// The name of the REST API
	Name pulumi.StringPtrInput
	// JSON formatted policy document that controls access to the API Gateway. This provider will only perform drift detection of its value when present in a configuration. It is recommended to use the `apigateway.RestApiPolicy` resource instead.
	Policy pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (RestApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restApiArgs)(nil)).Elem()
}

type RestApiInput interface {
	pulumi.Input

	ToRestApiOutput() RestApiOutput
	ToRestApiOutputWithContext(ctx context.Context) RestApiOutput
}

type RestApiPtrInput interface {
	pulumi.Input

	ToRestApiPtrOutput() RestApiPtrOutput
	ToRestApiPtrOutputWithContext(ctx context.Context) RestApiPtrOutput
}

func (RestApi) ElementType() reflect.Type {
	return reflect.TypeOf((*RestApi)(nil)).Elem()
}

func (i RestApi) ToRestApiOutput() RestApiOutput {
	return i.ToRestApiOutputWithContext(context.Background())
}

func (i RestApi) ToRestApiOutputWithContext(ctx context.Context) RestApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestApiOutput)
}

func (i RestApi) ToRestApiPtrOutput() RestApiPtrOutput {
	return i.ToRestApiPtrOutputWithContext(context.Background())
}

func (i RestApi) ToRestApiPtrOutputWithContext(ctx context.Context) RestApiPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestApiPtrOutput)
}

type RestApiOutput struct {
	*pulumi.OutputState
}

func (RestApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestApiOutput)(nil)).Elem()
}

func (o RestApiOutput) ToRestApiOutput() RestApiOutput {
	return o
}

func (o RestApiOutput) ToRestApiOutputWithContext(ctx context.Context) RestApiOutput {
	return o
}

type RestApiPtrOutput struct {
	*pulumi.OutputState
}

func (RestApiPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestApi)(nil)).Elem()
}

func (o RestApiPtrOutput) ToRestApiPtrOutput() RestApiPtrOutput {
	return o
}

func (o RestApiPtrOutput) ToRestApiPtrOutputWithContext(ctx context.Context) RestApiPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RestApiOutput{})
	pulumi.RegisterOutputType(RestApiPtrOutput{})
}
