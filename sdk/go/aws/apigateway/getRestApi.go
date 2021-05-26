// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the id and rootResourceId of a REST API in
// API Gateway. To fetch the REST API you must provide a name to match against.
// As there is no unique name constraint on REST APIs this data source will
// error if there is more than one match.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigateway.LookupRestApi(ctx, &apigateway.LookupRestApiArgs{
// 			Name: "my-rest-api",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRestApi(ctx *pulumi.Context, args *LookupRestApiArgs, opts ...pulumi.InvokeOption) (*LookupRestApiResult, error) {
	var rv LookupRestApiResult
	err := ctx.Invoke("aws:apigateway/getRestApi:getRestApi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRestApi.
type LookupRestApiArgs struct {
	// The name of the REST API to look up. If no REST API is found with this name, an error will be returned. If multiple REST APIs are found with this name, an error will be returned.
	Name string `pulumi:"name"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getRestApi.
type LookupRestApiResult struct {
	// The source of the API key for requests.
	ApiKeySource string `pulumi:"apiKeySource"`
	// The ARN of the REST API.
	Arn string `pulumi:"arn"`
	// The list of binary media types supported by the REST API.
	BinaryMediaTypes []string `pulumi:"binaryMediaTypes"`
	// The description of the REST API.
	Description string `pulumi:"description"`
	// The endpoint configuration of this RestApi showing the endpoint types of the API.
	EndpointConfigurations []GetRestApiEndpointConfiguration `pulumi:"endpointConfigurations"`
	// The execution ARN part to be used in `lambdaPermission`'s `sourceArn` when allowing API Gateway to invoke a Lambda function, e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
	ExecutionArn string `pulumi:"executionArn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Minimum response size to compress for the REST API.
	MinimumCompressionSize int    `pulumi:"minimumCompressionSize"`
	Name                   string `pulumi:"name"`
	// JSON formatted policy document that controls access to the API Gateway.
	Policy string `pulumi:"policy"`
	// Set to the ID of the API Gateway Resource on the found REST API where the route matches '/'.
	RootResourceId string `pulumi:"rootResourceId"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

func LookupRestApiApply(ctx *pulumi.Context, args LookupRestApiApplyInput, opts ...pulumi.InvokeOption) LookupRestApiResultOutput {
	return args.ToLookupRestApiApplyOutput().ApplyT(func(v LookupRestApiArgs) (LookupRestApiResult, error) {
		r, err := LookupRestApi(ctx, &v, opts...)
		return *r, err

	}).(LookupRestApiResultOutput)
}

// LookupRestApiApplyInput is an input type that accepts LookupRestApiApplyArgs and LookupRestApiApplyOutput values.
// You can construct a concrete instance of `LookupRestApiApplyInput` via:
//
//          LookupRestApiApplyArgs{...}
type LookupRestApiApplyInput interface {
	pulumi.Input

	ToLookupRestApiApplyOutput() LookupRestApiApplyOutput
	ToLookupRestApiApplyOutputWithContext(context.Context) LookupRestApiApplyOutput
}

// A collection of arguments for invoking getRestApi.
type LookupRestApiApplyArgs struct {
	// The name of the REST API to look up. If no REST API is found with this name, an error will be returned. If multiple REST APIs are found with this name, an error will be returned.
	Name pulumi.StringInput `pulumi:"name"`
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupRestApiApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestApiArgs)(nil)).Elem()
}

func (i LookupRestApiApplyArgs) ToLookupRestApiApplyOutput() LookupRestApiApplyOutput {
	return i.ToLookupRestApiApplyOutputWithContext(context.Background())
}

func (i LookupRestApiApplyArgs) ToLookupRestApiApplyOutputWithContext(ctx context.Context) LookupRestApiApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupRestApiApplyOutput)
}

// A collection of arguments for invoking getRestApi.
type LookupRestApiApplyOutput struct{ *pulumi.OutputState }

func (LookupRestApiApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestApiArgs)(nil)).Elem()
}

func (o LookupRestApiApplyOutput) ToLookupRestApiApplyOutput() LookupRestApiApplyOutput {
	return o
}

func (o LookupRestApiApplyOutput) ToLookupRestApiApplyOutputWithContext(ctx context.Context) LookupRestApiApplyOutput {
	return o
}

// The name of the REST API to look up. If no REST API is found with this name, an error will be returned. If multiple REST APIs are found with this name, an error will be returned.
func (o LookupRestApiApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiArgs) string { return v.Name }).(pulumi.StringOutput)
}

// Key-value map of resource tags.
func (o LookupRestApiApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRestApiArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getRestApi.
type LookupRestApiResultOutput struct{ *pulumi.OutputState }

func (LookupRestApiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestApiResult)(nil)).Elem()
}

func (o LookupRestApiResultOutput) ToLookupRestApiResultOutput() LookupRestApiResultOutput {
	return o
}

func (o LookupRestApiResultOutput) ToLookupRestApiResultOutputWithContext(ctx context.Context) LookupRestApiResultOutput {
	return o
}

// The source of the API key for requests.
func (o LookupRestApiResultOutput) ApiKeySource() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiResult) string { return v.ApiKeySource }).(pulumi.StringOutput)
}

// The ARN of the REST API.
func (o LookupRestApiResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The list of binary media types supported by the REST API.
func (o LookupRestApiResultOutput) BinaryMediaTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRestApiResult) []string { return v.BinaryMediaTypes }).(pulumi.StringArrayOutput)
}

// The description of the REST API.
func (o LookupRestApiResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiResult) string { return v.Description }).(pulumi.StringOutput)
}

// The endpoint configuration of this RestApi showing the endpoint types of the API.
func (o LookupRestApiResultOutput) EndpointConfigurations() GetRestApiEndpointConfigurationArrayOutput {
	return o.ApplyT(func(v LookupRestApiResult) []GetRestApiEndpointConfiguration { return v.EndpointConfigurations }).(GetRestApiEndpointConfigurationArrayOutput)
}

// The execution ARN part to be used in `lambdaPermission`'s `sourceArn` when allowing API Gateway to invoke a Lambda function, e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j`, which can be concatenated with allowed stage, method and resource path.
func (o LookupRestApiResultOutput) ExecutionArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiResult) string { return v.ExecutionArn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRestApiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiResult) string { return v.Id }).(pulumi.StringOutput)
}

// Minimum response size to compress for the REST API.
func (o LookupRestApiResultOutput) MinimumCompressionSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRestApiResult) int { return v.MinimumCompressionSize }).(pulumi.IntOutput)
}

func (o LookupRestApiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiResult) string { return v.Name }).(pulumi.StringOutput)
}

// JSON formatted policy document that controls access to the API Gateway.
func (o LookupRestApiResultOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiResult) string { return v.Policy }).(pulumi.StringOutput)
}

// Set to the ID of the API Gateway Resource on the found REST API where the route matches '/'.
func (o LookupRestApiResultOutput) RootResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestApiResult) string { return v.RootResourceId }).(pulumi.StringOutput)
}

// Key-value map of resource tags.
func (o LookupRestApiResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRestApiResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRestApiApplyOutput{})
	pulumi.RegisterOutputType(LookupRestApiResultOutput{})
}
