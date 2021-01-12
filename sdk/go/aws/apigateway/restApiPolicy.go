// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an API Gateway REST API Policy.
//
// > **Note:** Amazon API Gateway Version 1 resources are used for creating and deploying REST APIs. To create and deploy WebSocket and HTTP APIs, use Amazon API Gateway Version 2 [resources](https://www.terraform.io/docs/providers/aws/r/apigatewayv2_api.html).
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testRestApi, err := apigateway.NewRestApi(ctx, "testRestApi", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewRestApiPolicy(ctx, "testRestApiPolicy", &apigateway.RestApiPolicyArgs{
// 			RestApiId: testRestApi.ID(),
// 			Policy: testRestApi.Arn.ApplyT(func(arn string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"AWS\": \"*\"\n", "      },\n", "      \"Action\": \"execute-api:Invoke\",\n", "      \"Resource\": \"", arn, "\",\n", "      \"Condition\": {\n", "        \"IpAddress\": {\n", "          \"aws:SourceIp\": \"123.123.123.123/32\"\n", "        }\n", "      }\n", "    }\n", "  ]\n", "}\n"), nil
// 			}).(pulumi.StringOutput),
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
// `aws_api_gateway_rest_api_policy` can be imported by using the REST API ID, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/restApiPolicy:RestApiPolicy example 12345abcde
// ```
type RestApiPolicy struct {
	pulumi.CustomResourceState

	Policy pulumi.StringOutput `pulumi:"policy"`
	// The ID of the REST API.
	RestApiId pulumi.StringOutput `pulumi:"restApiId"`
}

// NewRestApiPolicy registers a new resource with the given unique name, arguments, and options.
func NewRestApiPolicy(ctx *pulumi.Context,
	name string, args *RestApiPolicyArgs, opts ...pulumi.ResourceOption) (*RestApiPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.RestApiId == nil {
		return nil, errors.New("invalid value for required argument 'RestApiId'")
	}
	var resource RestApiPolicy
	err := ctx.RegisterResource("aws:apigateway/restApiPolicy:RestApiPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRestApiPolicy gets an existing RestApiPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRestApiPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RestApiPolicyState, opts ...pulumi.ResourceOption) (*RestApiPolicy, error) {
	var resource RestApiPolicy
	err := ctx.ReadResource("aws:apigateway/restApiPolicy:RestApiPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RestApiPolicy resources.
type restApiPolicyState struct {
	Policy *string `pulumi:"policy"`
	// The ID of the REST API.
	RestApiId *string `pulumi:"restApiId"`
}

type RestApiPolicyState struct {
	Policy pulumi.StringPtrInput
	// The ID of the REST API.
	RestApiId pulumi.StringPtrInput
}

func (RestApiPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*restApiPolicyState)(nil)).Elem()
}

type restApiPolicyArgs struct {
	Policy string `pulumi:"policy"`
	// The ID of the REST API.
	RestApiId string `pulumi:"restApiId"`
}

// The set of arguments for constructing a RestApiPolicy resource.
type RestApiPolicyArgs struct {
	Policy pulumi.StringInput
	// The ID of the REST API.
	RestApiId pulumi.StringInput
}

func (RestApiPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restApiPolicyArgs)(nil)).Elem()
}

type RestApiPolicyInput interface {
	pulumi.Input

	ToRestApiPolicyOutput() RestApiPolicyOutput
	ToRestApiPolicyOutputWithContext(ctx context.Context) RestApiPolicyOutput
}

func (*RestApiPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*RestApiPolicy)(nil))
}

func (i *RestApiPolicy) ToRestApiPolicyOutput() RestApiPolicyOutput {
	return i.ToRestApiPolicyOutputWithContext(context.Background())
}

func (i *RestApiPolicy) ToRestApiPolicyOutputWithContext(ctx context.Context) RestApiPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestApiPolicyOutput)
}

func (i *RestApiPolicy) ToRestApiPolicyPtrOutput() RestApiPolicyPtrOutput {
	return i.ToRestApiPolicyPtrOutputWithContext(context.Background())
}

func (i *RestApiPolicy) ToRestApiPolicyPtrOutputWithContext(ctx context.Context) RestApiPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestApiPolicyPtrOutput)
}

type RestApiPolicyPtrInput interface {
	pulumi.Input

	ToRestApiPolicyPtrOutput() RestApiPolicyPtrOutput
	ToRestApiPolicyPtrOutputWithContext(ctx context.Context) RestApiPolicyPtrOutput
}

type restApiPolicyPtrType RestApiPolicyArgs

func (*restApiPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RestApiPolicy)(nil))
}

func (i *restApiPolicyPtrType) ToRestApiPolicyPtrOutput() RestApiPolicyPtrOutput {
	return i.ToRestApiPolicyPtrOutputWithContext(context.Background())
}

func (i *restApiPolicyPtrType) ToRestApiPolicyPtrOutputWithContext(ctx context.Context) RestApiPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestApiPolicyOutput).ToRestApiPolicyPtrOutput()
}

// RestApiPolicyArrayInput is an input type that accepts RestApiPolicyArray and RestApiPolicyArrayOutput values.
// You can construct a concrete instance of `RestApiPolicyArrayInput` via:
//
//          RestApiPolicyArray{ RestApiPolicyArgs{...} }
type RestApiPolicyArrayInput interface {
	pulumi.Input

	ToRestApiPolicyArrayOutput() RestApiPolicyArrayOutput
	ToRestApiPolicyArrayOutputWithContext(context.Context) RestApiPolicyArrayOutput
}

type RestApiPolicyArray []RestApiPolicyInput

func (RestApiPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RestApiPolicy)(nil))
}

func (i RestApiPolicyArray) ToRestApiPolicyArrayOutput() RestApiPolicyArrayOutput {
	return i.ToRestApiPolicyArrayOutputWithContext(context.Background())
}

func (i RestApiPolicyArray) ToRestApiPolicyArrayOutputWithContext(ctx context.Context) RestApiPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestApiPolicyArrayOutput)
}

// RestApiPolicyMapInput is an input type that accepts RestApiPolicyMap and RestApiPolicyMapOutput values.
// You can construct a concrete instance of `RestApiPolicyMapInput` via:
//
//          RestApiPolicyMap{ "key": RestApiPolicyArgs{...} }
type RestApiPolicyMapInput interface {
	pulumi.Input

	ToRestApiPolicyMapOutput() RestApiPolicyMapOutput
	ToRestApiPolicyMapOutputWithContext(context.Context) RestApiPolicyMapOutput
}

type RestApiPolicyMap map[string]RestApiPolicyInput

func (RestApiPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RestApiPolicy)(nil))
}

func (i RestApiPolicyMap) ToRestApiPolicyMapOutput() RestApiPolicyMapOutput {
	return i.ToRestApiPolicyMapOutputWithContext(context.Background())
}

func (i RestApiPolicyMap) ToRestApiPolicyMapOutputWithContext(ctx context.Context) RestApiPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestApiPolicyMapOutput)
}

type RestApiPolicyOutput struct {
	*pulumi.OutputState
}

func (RestApiPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestApiPolicy)(nil))
}

func (o RestApiPolicyOutput) ToRestApiPolicyOutput() RestApiPolicyOutput {
	return o
}

func (o RestApiPolicyOutput) ToRestApiPolicyOutputWithContext(ctx context.Context) RestApiPolicyOutput {
	return o
}

func (o RestApiPolicyOutput) ToRestApiPolicyPtrOutput() RestApiPolicyPtrOutput {
	return o.ToRestApiPolicyPtrOutputWithContext(context.Background())
}

func (o RestApiPolicyOutput) ToRestApiPolicyPtrOutputWithContext(ctx context.Context) RestApiPolicyPtrOutput {
	return o.ApplyT(func(v RestApiPolicy) *RestApiPolicy {
		return &v
	}).(RestApiPolicyPtrOutput)
}

type RestApiPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (RestApiPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestApiPolicy)(nil))
}

func (o RestApiPolicyPtrOutput) ToRestApiPolicyPtrOutput() RestApiPolicyPtrOutput {
	return o
}

func (o RestApiPolicyPtrOutput) ToRestApiPolicyPtrOutputWithContext(ctx context.Context) RestApiPolicyPtrOutput {
	return o
}

type RestApiPolicyArrayOutput struct{ *pulumi.OutputState }

func (RestApiPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RestApiPolicy)(nil))
}

func (o RestApiPolicyArrayOutput) ToRestApiPolicyArrayOutput() RestApiPolicyArrayOutput {
	return o
}

func (o RestApiPolicyArrayOutput) ToRestApiPolicyArrayOutputWithContext(ctx context.Context) RestApiPolicyArrayOutput {
	return o
}

func (o RestApiPolicyArrayOutput) Index(i pulumi.IntInput) RestApiPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RestApiPolicy {
		return vs[0].([]RestApiPolicy)[vs[1].(int)]
	}).(RestApiPolicyOutput)
}

type RestApiPolicyMapOutput struct{ *pulumi.OutputState }

func (RestApiPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RestApiPolicy)(nil))
}

func (o RestApiPolicyMapOutput) ToRestApiPolicyMapOutput() RestApiPolicyMapOutput {
	return o
}

func (o RestApiPolicyMapOutput) ToRestApiPolicyMapOutputWithContext(ctx context.Context) RestApiPolicyMapOutput {
	return o
}

func (o RestApiPolicyMapOutput) MapIndex(k pulumi.StringInput) RestApiPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RestApiPolicy {
		return vs[0].(map[string]RestApiPolicy)[vs[1].(string)]
	}).(RestApiPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(RestApiPolicyOutput{})
	pulumi.RegisterOutputType(RestApiPolicyPtrOutput{})
	pulumi.RegisterOutputType(RestApiPolicyArrayOutput{})
	pulumi.RegisterOutputType(RestApiPolicyMapOutput{})
}
