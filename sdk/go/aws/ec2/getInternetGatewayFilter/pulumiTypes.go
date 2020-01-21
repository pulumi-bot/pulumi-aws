// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getInternetGatewayFilter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetInternetGatewayFilter struct {
	// The name of the field to filter by, as defined by
	// [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInternetGateways.html).
	Name string `pulumi:"name"`
	// Set of values that are accepted for the given field.
	// An Internet Gateway will be selected if any one of the given values matches.
	Values []string `pulumi:"values"`
}

type GetInternetGatewayFilterInput interface {
	pulumi.Input

	ToGetInternetGatewayFilterOutput() GetInternetGatewayFilterOutput
	ToGetInternetGatewayFilterOutputWithContext(context.Context) GetInternetGatewayFilterOutput
}

type GetInternetGatewayFilterArgs struct {
	// The name of the field to filter by, as defined by
	// [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInternetGateways.html).
	Name pulumi.StringInput `pulumi:"name"`
	// Set of values that are accepted for the given field.
	// An Internet Gateway will be selected if any one of the given values matches.
	Values pulumi.StringArrayInput `pulumi:"values"`
}

func (GetInternetGatewayFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInternetGatewayFilter)(nil)).Elem()
}

func (i GetInternetGatewayFilterArgs) ToGetInternetGatewayFilterOutput() GetInternetGatewayFilterOutput {
	return i.ToGetInternetGatewayFilterOutputWithContext(context.Background())
}

func (i GetInternetGatewayFilterArgs) ToGetInternetGatewayFilterOutputWithContext(ctx context.Context) GetInternetGatewayFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInternetGatewayFilterOutput)
}

type GetInternetGatewayFilterArrayInput interface {
	pulumi.Input

	ToGetInternetGatewayFilterArrayOutput() GetInternetGatewayFilterArrayOutput
	ToGetInternetGatewayFilterArrayOutputWithContext(context.Context) GetInternetGatewayFilterArrayOutput
}

type GetInternetGatewayFilterArray []GetInternetGatewayFilterInput

func (GetInternetGatewayFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInternetGatewayFilter)(nil)).Elem()
}

func (i GetInternetGatewayFilterArray) ToGetInternetGatewayFilterArrayOutput() GetInternetGatewayFilterArrayOutput {
	return i.ToGetInternetGatewayFilterArrayOutputWithContext(context.Background())
}

func (i GetInternetGatewayFilterArray) ToGetInternetGatewayFilterArrayOutputWithContext(ctx context.Context) GetInternetGatewayFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInternetGatewayFilterArrayOutput)
}

type GetInternetGatewayFilterOutput struct { *pulumi.OutputState }

func (GetInternetGatewayFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInternetGatewayFilter)(nil)).Elem()
}

func (o GetInternetGatewayFilterOutput) ToGetInternetGatewayFilterOutput() GetInternetGatewayFilterOutput {
	return o
}

func (o GetInternetGatewayFilterOutput) ToGetInternetGatewayFilterOutputWithContext(ctx context.Context) GetInternetGatewayFilterOutput {
	return o
}

// The name of the field to filter by, as defined by
// [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInternetGateways.html).
func (o GetInternetGatewayFilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v GetInternetGatewayFilter) string { return v.Name }).(pulumi.StringOutput)
}

// Set of values that are accepted for the given field.
// An Internet Gateway will be selected if any one of the given values matches.
func (o GetInternetGatewayFilterOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func (v GetInternetGatewayFilter) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type GetInternetGatewayFilterArrayOutput struct { *pulumi.OutputState}

func (GetInternetGatewayFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInternetGatewayFilter)(nil)).Elem()
}

func (o GetInternetGatewayFilterArrayOutput) ToGetInternetGatewayFilterArrayOutput() GetInternetGatewayFilterArrayOutput {
	return o
}

func (o GetInternetGatewayFilterArrayOutput) ToGetInternetGatewayFilterArrayOutputWithContext(ctx context.Context) GetInternetGatewayFilterArrayOutput {
	return o
}

func (o GetInternetGatewayFilterArrayOutput) Index(i pulumi.IntInput) GetInternetGatewayFilterOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetInternetGatewayFilter {
		return vs[0].([]GetInternetGatewayFilter)[vs[1].(int)]
	}).(GetInternetGatewayFilterOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInternetGatewayFilterOutput{})
	pulumi.RegisterOutputType(GetInternetGatewayFilterArrayOutput{})
}
