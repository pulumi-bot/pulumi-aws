// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClusterParameterGroupParameter struct {
	// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod *string `pulumi:"applyMethod"`
	// The name of the documentDB cluster parameter group. If omitted, this provider will assign a random, unique name.
	Name string `pulumi:"name"`
	// The value of the documentDB parameter.
	Value string `pulumi:"value"`
}

// ClusterParameterGroupParameterInput is an input type that accepts ClusterParameterGroupParameterArgs and ClusterParameterGroupParameterOutput values.
// You can construct a concrete instance of `ClusterParameterGroupParameterInput` via:
//
// 		 ClusterParameterGroupParameterArgs{...}
//
type ClusterParameterGroupParameterInput interface {
	pulumi.Input

	ToClusterParameterGroupParameterOutput() ClusterParameterGroupParameterOutput
	ToClusterParameterGroupParameterOutputWithContext(context.Context) ClusterParameterGroupParameterOutput
}

type ClusterParameterGroupParameterArgs struct {
	// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod pulumi.StringPtrInput `pulumi:"applyMethod"`
	// The name of the documentDB cluster parameter group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the documentDB parameter.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ClusterParameterGroupParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParameterGroupParameter)(nil)).Elem()
}

func (i ClusterParameterGroupParameterArgs) ToClusterParameterGroupParameterOutput() ClusterParameterGroupParameterOutput {
	return i.ToClusterParameterGroupParameterOutputWithContext(context.Background())
}

func (i ClusterParameterGroupParameterArgs) ToClusterParameterGroupParameterOutputWithContext(ctx context.Context) ClusterParameterGroupParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupParameterOutput)
}

// ClusterParameterGroupParameterArrayInput is an input type that accepts ClusterParameterGroupParameterArray and ClusterParameterGroupParameterArrayOutput values.
// You can construct a concrete instance of `ClusterParameterGroupParameterArrayInput` via:
//
// 		 ClusterParameterGroupParameterArray{ ClusterParameterGroupParameterArgs{...} }
//
type ClusterParameterGroupParameterArrayInput interface {
	pulumi.Input

	ToClusterParameterGroupParameterArrayOutput() ClusterParameterGroupParameterArrayOutput
	ToClusterParameterGroupParameterArrayOutputWithContext(context.Context) ClusterParameterGroupParameterArrayOutput
}

type ClusterParameterGroupParameterArray []ClusterParameterGroupParameterInput

func (ClusterParameterGroupParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterParameterGroupParameter)(nil)).Elem()
}

func (i ClusterParameterGroupParameterArray) ToClusterParameterGroupParameterArrayOutput() ClusterParameterGroupParameterArrayOutput {
	return i.ToClusterParameterGroupParameterArrayOutputWithContext(context.Background())
}

func (i ClusterParameterGroupParameterArray) ToClusterParameterGroupParameterArrayOutputWithContext(ctx context.Context) ClusterParameterGroupParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupParameterArrayOutput)
}

type ClusterParameterGroupParameterOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParameterGroupParameter)(nil)).Elem()
}

func (o ClusterParameterGroupParameterOutput) ToClusterParameterGroupParameterOutput() ClusterParameterGroupParameterOutput {
	return o
}

func (o ClusterParameterGroupParameterOutput) ToClusterParameterGroupParameterOutputWithContext(ctx context.Context) ClusterParameterGroupParameterOutput {
	return o
}

// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
func (o ClusterParameterGroupParameterOutput) ApplyMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterParameterGroupParameter) *string { return v.ApplyMethod }).(pulumi.StringPtrOutput)
}

// The name of the documentDB cluster parameter group. If omitted, this provider will assign a random, unique name.
func (o ClusterParameterGroupParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterParameterGroupParameter) string { return v.Name }).(pulumi.StringOutput)
}

// The value of the documentDB parameter.
func (o ClusterParameterGroupParameterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterParameterGroupParameter) string { return v.Value }).(pulumi.StringOutput)
}

type ClusterParameterGroupParameterArrayOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterParameterGroupParameter)(nil)).Elem()
}

func (o ClusterParameterGroupParameterArrayOutput) ToClusterParameterGroupParameterArrayOutput() ClusterParameterGroupParameterArrayOutput {
	return o
}

func (o ClusterParameterGroupParameterArrayOutput) ToClusterParameterGroupParameterArrayOutputWithContext(ctx context.Context) ClusterParameterGroupParameterArrayOutput {
	return o
}

func (o ClusterParameterGroupParameterArrayOutput) Index(i pulumi.IntInput) ClusterParameterGroupParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterParameterGroupParameter {
		return vs[0].([]ClusterParameterGroupParameter)[vs[1].(int)]
	}).(ClusterParameterGroupParameterOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterParameterGroupParameterOutput{})
	pulumi.RegisterOutputType(ClusterParameterGroupParameterArrayOutput{})
}
