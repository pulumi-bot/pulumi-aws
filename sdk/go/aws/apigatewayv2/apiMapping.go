// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Amazon API Gateway Version 2 API mapping.
// More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigatewayv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigatewayv2.NewApiMapping(ctx, "example", &apigatewayv2.ApiMappingArgs{
// 			ApiId:      pulumi.Any(aws_apigatewayv2_api.Example.Id),
// 			DomainName: pulumi.Any(aws_apigatewayv2_domain_name.Example.Id),
// 			Stage:      pulumi.Any(aws_apigatewayv2_stage.Example.Id),
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
// `aws_apigatewayv2_api_mapping` can be imported by using the API mapping identifier and domain name, e.g.
//
// ```sh
//  $ pulumi import aws:apigatewayv2/apiMapping:ApiMapping example 1122334/ws-api.example.com
// ```
type ApiMapping struct {
	pulumi.CustomResourceState

	// The API identifier.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The [API mapping key](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-mapping-template-reference.html).
	ApiMappingKey pulumi.StringPtrOutput `pulumi:"apiMappingKey"`
	// The domain name. Use the `apigatewayv2.DomainName` resource to configure a domain name.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The API stage. Use the `apigatewayv2.Stage` resource to configure an API stage.
	Stage pulumi.StringOutput `pulumi:"stage"`
}

// NewApiMapping registers a new resource with the given unique name, arguments, and options.
func NewApiMapping(ctx *pulumi.Context,
	name string, args *ApiMappingArgs, opts ...pulumi.ResourceOption) (*ApiMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.Stage == nil {
		return nil, errors.New("invalid value for required argument 'Stage'")
	}
	var resource ApiMapping
	err := ctx.RegisterResource("aws:apigatewayv2/apiMapping:ApiMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiMapping gets an existing ApiMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiMappingState, opts ...pulumi.ResourceOption) (*ApiMapping, error) {
	var resource ApiMapping
	err := ctx.ReadResource("aws:apigatewayv2/apiMapping:ApiMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiMapping resources.
type apiMappingState struct {
	// The API identifier.
	ApiId *string `pulumi:"apiId"`
	// The [API mapping key](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-mapping-template-reference.html).
	ApiMappingKey *string `pulumi:"apiMappingKey"`
	// The domain name. Use the `apigatewayv2.DomainName` resource to configure a domain name.
	DomainName *string `pulumi:"domainName"`
	// The API stage. Use the `apigatewayv2.Stage` resource to configure an API stage.
	Stage *string `pulumi:"stage"`
}

type ApiMappingState struct {
	// The API identifier.
	ApiId pulumi.StringPtrInput
	// The [API mapping key](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-mapping-template-reference.html).
	ApiMappingKey pulumi.StringPtrInput
	// The domain name. Use the `apigatewayv2.DomainName` resource to configure a domain name.
	DomainName pulumi.StringPtrInput
	// The API stage. Use the `apigatewayv2.Stage` resource to configure an API stage.
	Stage pulumi.StringPtrInput
}

func (ApiMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiMappingState)(nil)).Elem()
}

type apiMappingArgs struct {
	// The API identifier.
	ApiId string `pulumi:"apiId"`
	// The [API mapping key](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-mapping-template-reference.html).
	ApiMappingKey *string `pulumi:"apiMappingKey"`
	// The domain name. Use the `apigatewayv2.DomainName` resource to configure a domain name.
	DomainName string `pulumi:"domainName"`
	// The API stage. Use the `apigatewayv2.Stage` resource to configure an API stage.
	Stage string `pulumi:"stage"`
}

// The set of arguments for constructing a ApiMapping resource.
type ApiMappingArgs struct {
	// The API identifier.
	ApiId pulumi.StringInput
	// The [API mapping key](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-mapping-template-reference.html).
	ApiMappingKey pulumi.StringPtrInput
	// The domain name. Use the `apigatewayv2.DomainName` resource to configure a domain name.
	DomainName pulumi.StringInput
	// The API stage. Use the `apigatewayv2.Stage` resource to configure an API stage.
	Stage pulumi.StringInput
}

func (ApiMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiMappingArgs)(nil)).Elem()
}

type ApiMappingInput interface {
	pulumi.Input

	ToApiMappingOutput() ApiMappingOutput
	ToApiMappingOutputWithContext(ctx context.Context) ApiMappingOutput
}

func (*ApiMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiMapping)(nil))
}

func (i *ApiMapping) ToApiMappingOutput() ApiMappingOutput {
	return i.ToApiMappingOutputWithContext(context.Background())
}

func (i *ApiMapping) ToApiMappingOutputWithContext(ctx context.Context) ApiMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiMappingOutput)
}

func (i *ApiMapping) ToApiMappingPtrOutput() ApiMappingPtrOutput {
	return i.ToApiMappingPtrOutputWithContext(context.Background())
}

func (i *ApiMapping) ToApiMappingPtrOutputWithContext(ctx context.Context) ApiMappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiMappingPtrOutput)
}

type ApiMappingPtrInput interface {
	pulumi.Input

	ToApiMappingPtrOutput() ApiMappingPtrOutput
	ToApiMappingPtrOutputWithContext(ctx context.Context) ApiMappingPtrOutput
}

type apiMappingPtrType ApiMappingArgs

func (*apiMappingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiMapping)(nil))
}

func (i *apiMappingPtrType) ToApiMappingPtrOutput() ApiMappingPtrOutput {
	return i.ToApiMappingPtrOutputWithContext(context.Background())
}

func (i *apiMappingPtrType) ToApiMappingPtrOutputWithContext(ctx context.Context) ApiMappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiMappingOutput).ToApiMappingPtrOutput()
}

// ApiMappingArrayInput is an input type that accepts ApiMappingArray and ApiMappingArrayOutput values.
// You can construct a concrete instance of `ApiMappingArrayInput` via:
//
//          ApiMappingArray{ ApiMappingArgs{...} }
type ApiMappingArrayInput interface {
	pulumi.Input

	ToApiMappingArrayOutput() ApiMappingArrayOutput
	ToApiMappingArrayOutputWithContext(context.Context) ApiMappingArrayOutput
}

type ApiMappingArray []ApiMappingInput

func (ApiMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiMapping)(nil))
}

func (i ApiMappingArray) ToApiMappingArrayOutput() ApiMappingArrayOutput {
	return i.ToApiMappingArrayOutputWithContext(context.Background())
}

func (i ApiMappingArray) ToApiMappingArrayOutputWithContext(ctx context.Context) ApiMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiMappingArrayOutput)
}

// ApiMappingMapInput is an input type that accepts ApiMappingMap and ApiMappingMapOutput values.
// You can construct a concrete instance of `ApiMappingMapInput` via:
//
//          ApiMappingMap{ "key": ApiMappingArgs{...} }
type ApiMappingMapInput interface {
	pulumi.Input

	ToApiMappingMapOutput() ApiMappingMapOutput
	ToApiMappingMapOutputWithContext(context.Context) ApiMappingMapOutput
}

type ApiMappingMap map[string]ApiMappingInput

func (ApiMappingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiMapping)(nil))
}

func (i ApiMappingMap) ToApiMappingMapOutput() ApiMappingMapOutput {
	return i.ToApiMappingMapOutputWithContext(context.Background())
}

func (i ApiMappingMap) ToApiMappingMapOutputWithContext(ctx context.Context) ApiMappingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiMappingMapOutput)
}

type ApiMappingOutput struct {
	*pulumi.OutputState
}

func (ApiMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiMapping)(nil))
}

func (o ApiMappingOutput) ToApiMappingOutput() ApiMappingOutput {
	return o
}

func (o ApiMappingOutput) ToApiMappingOutputWithContext(ctx context.Context) ApiMappingOutput {
	return o
}

func (o ApiMappingOutput) ToApiMappingPtrOutput() ApiMappingPtrOutput {
	return o.ToApiMappingPtrOutputWithContext(context.Background())
}

func (o ApiMappingOutput) ToApiMappingPtrOutputWithContext(ctx context.Context) ApiMappingPtrOutput {
	return o.ApplyT(func(v ApiMapping) *ApiMapping {
		return &v
	}).(ApiMappingPtrOutput)
}

type ApiMappingPtrOutput struct {
	*pulumi.OutputState
}

func (ApiMappingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiMapping)(nil))
}

func (o ApiMappingPtrOutput) ToApiMappingPtrOutput() ApiMappingPtrOutput {
	return o
}

func (o ApiMappingPtrOutput) ToApiMappingPtrOutputWithContext(ctx context.Context) ApiMappingPtrOutput {
	return o
}

type ApiMappingArrayOutput struct{ *pulumi.OutputState }

func (ApiMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiMapping)(nil))
}

func (o ApiMappingArrayOutput) ToApiMappingArrayOutput() ApiMappingArrayOutput {
	return o
}

func (o ApiMappingArrayOutput) ToApiMappingArrayOutputWithContext(ctx context.Context) ApiMappingArrayOutput {
	return o
}

func (o ApiMappingArrayOutput) Index(i pulumi.IntInput) ApiMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiMapping {
		return vs[0].([]ApiMapping)[vs[1].(int)]
	}).(ApiMappingOutput)
}

type ApiMappingMapOutput struct{ *pulumi.OutputState }

func (ApiMappingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiMapping)(nil))
}

func (o ApiMappingMapOutput) ToApiMappingMapOutput() ApiMappingMapOutput {
	return o
}

func (o ApiMappingMapOutput) ToApiMappingMapOutputWithContext(ctx context.Context) ApiMappingMapOutput {
	return o
}

func (o ApiMappingMapOutput) MapIndex(k pulumi.StringInput) ApiMappingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiMapping {
		return vs[0].(map[string]ApiMapping)[vs[1].(string)]
	}).(ApiMappingOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiMappingOutput{})
	pulumi.RegisterOutputType(ApiMappingPtrOutput{})
	pulumi.RegisterOutputType(ApiMappingArrayOutput{})
	pulumi.RegisterOutputType(ApiMappingMapOutput{})
}
