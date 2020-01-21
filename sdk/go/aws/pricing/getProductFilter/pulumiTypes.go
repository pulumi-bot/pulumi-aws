// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getProductFilter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetProductFilter struct {
	// The product attribute name that you want to filter on.
	Field string `pulumi:"field"`
	// The product attribute value that you want to filter on.
	Value string `pulumi:"value"`
}

type GetProductFilterInput interface {
	pulumi.Input

	ToGetProductFilterOutput() GetProductFilterOutput
	ToGetProductFilterOutputWithContext(context.Context) GetProductFilterOutput
}

type GetProductFilterArgs struct {
	// The product attribute name that you want to filter on.
	Field pulumi.StringInput `pulumi:"field"`
	// The product attribute value that you want to filter on.
	Value pulumi.StringInput `pulumi:"value"`
}

func (GetProductFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductFilter)(nil)).Elem()
}

func (i GetProductFilterArgs) ToGetProductFilterOutput() GetProductFilterOutput {
	return i.ToGetProductFilterOutputWithContext(context.Background())
}

func (i GetProductFilterArgs) ToGetProductFilterOutputWithContext(ctx context.Context) GetProductFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProductFilterOutput)
}

type GetProductFilterArrayInput interface {
	pulumi.Input

	ToGetProductFilterArrayOutput() GetProductFilterArrayOutput
	ToGetProductFilterArrayOutputWithContext(context.Context) GetProductFilterArrayOutput
}

type GetProductFilterArray []GetProductFilterInput

func (GetProductFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProductFilter)(nil)).Elem()
}

func (i GetProductFilterArray) ToGetProductFilterArrayOutput() GetProductFilterArrayOutput {
	return i.ToGetProductFilterArrayOutputWithContext(context.Background())
}

func (i GetProductFilterArray) ToGetProductFilterArrayOutputWithContext(ctx context.Context) GetProductFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProductFilterArrayOutput)
}

type GetProductFilterOutput struct { *pulumi.OutputState }

func (GetProductFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductFilter)(nil)).Elem()
}

func (o GetProductFilterOutput) ToGetProductFilterOutput() GetProductFilterOutput {
	return o
}

func (o GetProductFilterOutput) ToGetProductFilterOutputWithContext(ctx context.Context) GetProductFilterOutput {
	return o
}

// The product attribute name that you want to filter on.
func (o GetProductFilterOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func (v GetProductFilter) string { return v.Field }).(pulumi.StringOutput)
}

// The product attribute value that you want to filter on.
func (o GetProductFilterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func (v GetProductFilter) string { return v.Value }).(pulumi.StringOutput)
}

type GetProductFilterArrayOutput struct { *pulumi.OutputState}

func (GetProductFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProductFilter)(nil)).Elem()
}

func (o GetProductFilterArrayOutput) ToGetProductFilterArrayOutput() GetProductFilterArrayOutput {
	return o
}

func (o GetProductFilterArrayOutput) ToGetProductFilterArrayOutputWithContext(ctx context.Context) GetProductFilterArrayOutput {
	return o
}

func (o GetProductFilterArrayOutput) Index(i pulumi.IntInput) GetProductFilterOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetProductFilter {
		return vs[0].([]GetProductFilter)[vs[1].(int)]
	}).(GetProductFilterOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProductFilterOutput{})
	pulumi.RegisterOutputType(GetProductFilterArrayOutput{})
}
