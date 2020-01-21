// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getResourceShareFilter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetResourceShareFilter struct {
	// The name of the tag key to filter on.
	Name string `pulumi:"name"`
	// The value of the tag key.
	Values []string `pulumi:"values"`
}

type GetResourceShareFilterInput interface {
	pulumi.Input

	ToGetResourceShareFilterOutput() GetResourceShareFilterOutput
	ToGetResourceShareFilterOutputWithContext(context.Context) GetResourceShareFilterOutput
}

type GetResourceShareFilterArgs struct {
	// The name of the tag key to filter on.
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the tag key.
	Values pulumi.StringArrayInput `pulumi:"values"`
}

func (GetResourceShareFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourceShareFilter)(nil)).Elem()
}

func (i GetResourceShareFilterArgs) ToGetResourceShareFilterOutput() GetResourceShareFilterOutput {
	return i.ToGetResourceShareFilterOutputWithContext(context.Background())
}

func (i GetResourceShareFilterArgs) ToGetResourceShareFilterOutputWithContext(ctx context.Context) GetResourceShareFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetResourceShareFilterOutput)
}

type GetResourceShareFilterArrayInput interface {
	pulumi.Input

	ToGetResourceShareFilterArrayOutput() GetResourceShareFilterArrayOutput
	ToGetResourceShareFilterArrayOutputWithContext(context.Context) GetResourceShareFilterArrayOutput
}

type GetResourceShareFilterArray []GetResourceShareFilterInput

func (GetResourceShareFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetResourceShareFilter)(nil)).Elem()
}

func (i GetResourceShareFilterArray) ToGetResourceShareFilterArrayOutput() GetResourceShareFilterArrayOutput {
	return i.ToGetResourceShareFilterArrayOutputWithContext(context.Background())
}

func (i GetResourceShareFilterArray) ToGetResourceShareFilterArrayOutputWithContext(ctx context.Context) GetResourceShareFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetResourceShareFilterArrayOutput)
}

type GetResourceShareFilterOutput struct { *pulumi.OutputState }

func (GetResourceShareFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourceShareFilter)(nil)).Elem()
}

func (o GetResourceShareFilterOutput) ToGetResourceShareFilterOutput() GetResourceShareFilterOutput {
	return o
}

func (o GetResourceShareFilterOutput) ToGetResourceShareFilterOutputWithContext(ctx context.Context) GetResourceShareFilterOutput {
	return o
}

// The name of the tag key to filter on.
func (o GetResourceShareFilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v GetResourceShareFilter) string { return v.Name }).(pulumi.StringOutput)
}

// The value of the tag key.
func (o GetResourceShareFilterOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func (v GetResourceShareFilter) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type GetResourceShareFilterArrayOutput struct { *pulumi.OutputState}

func (GetResourceShareFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetResourceShareFilter)(nil)).Elem()
}

func (o GetResourceShareFilterArrayOutput) ToGetResourceShareFilterArrayOutput() GetResourceShareFilterArrayOutput {
	return o
}

func (o GetResourceShareFilterArrayOutput) ToGetResourceShareFilterArrayOutputWithContext(ctx context.Context) GetResourceShareFilterArrayOutput {
	return o
}

func (o GetResourceShareFilterArrayOutput) Index(i pulumi.IntInput) GetResourceShareFilterOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetResourceShareFilter {
		return vs[0].([]GetResourceShareFilter)[vs[1].(int)]
	}).(GetResourceShareFilterOutput)
}

func init() {
	pulumi.RegisterOutputType(GetResourceShareFilterOutput{})
	pulumi.RegisterOutputType(GetResourceShareFilterArrayOutput{})
}
