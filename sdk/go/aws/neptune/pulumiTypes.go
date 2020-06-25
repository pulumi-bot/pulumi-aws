// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClusterParameterGroupParameter struct {
	// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod *string `pulumi:"applyMethod"`
	// The name of the neptune parameter.
	Name string `pulumi:"name"`
	// The value of the neptune parameter.
	Value string `pulumi:"value"`
}

// ClusterParameterGroupParameterInput is an input type that accepts ClusterParameterGroupParameterArgs and ClusterParameterGroupParameterOutput values.
// You can construct a concrete instance of `ClusterParameterGroupParameterInput` via:
//
//          ClusterParameterGroupParameterArgs{...}
type ClusterParameterGroupParameterInput interface {
	pulumi.Input

	ToClusterParameterGroupParameterOutput() ClusterParameterGroupParameterOutput
	ToClusterParameterGroupParameterOutputWithContext(context.Context) ClusterParameterGroupParameterOutput
}

type ClusterParameterGroupParameterArgs struct {
	// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod pulumi.StringPtrInput `pulumi:"applyMethod"`
	// The name of the neptune parameter.
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the neptune parameter.
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
//          ClusterParameterGroupParameterArray{ ClusterParameterGroupParameterArgs{...} }
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

// The name of the neptune parameter.
func (o ClusterParameterGroupParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterParameterGroupParameter) string { return v.Name }).(pulumi.StringOutput)
}

// The value of the neptune parameter.
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

type ParameterGroupParameter struct {
	// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod *string `pulumi:"applyMethod"`
	// The name of the Neptune parameter.
	Name string `pulumi:"name"`
	// The value of the Neptune parameter.
	Value string `pulumi:"value"`
}

// ParameterGroupParameterInput is an input type that accepts ParameterGroupParameterArgs and ParameterGroupParameterOutput values.
// You can construct a concrete instance of `ParameterGroupParameterInput` via:
//
//          ParameterGroupParameterArgs{...}
type ParameterGroupParameterInput interface {
	pulumi.Input

	ToParameterGroupParameterOutput() ParameterGroupParameterOutput
	ToParameterGroupParameterOutputWithContext(context.Context) ParameterGroupParameterOutput
}

type ParameterGroupParameterArgs struct {
	// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod pulumi.StringPtrInput `pulumi:"applyMethod"`
	// The name of the Neptune parameter.
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the Neptune parameter.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ParameterGroupParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameter)(nil)).Elem()
}

func (i ParameterGroupParameterArgs) ToParameterGroupParameterOutput() ParameterGroupParameterOutput {
	return i.ToParameterGroupParameterOutputWithContext(context.Background())
}

func (i ParameterGroupParameterArgs) ToParameterGroupParameterOutputWithContext(ctx context.Context) ParameterGroupParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterOutput)
}

// ParameterGroupParameterArrayInput is an input type that accepts ParameterGroupParameterArray and ParameterGroupParameterArrayOutput values.
// You can construct a concrete instance of `ParameterGroupParameterArrayInput` via:
//
//          ParameterGroupParameterArray{ ParameterGroupParameterArgs{...} }
type ParameterGroupParameterArrayInput interface {
	pulumi.Input

	ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput
	ToParameterGroupParameterArrayOutputWithContext(context.Context) ParameterGroupParameterArrayOutput
}

type ParameterGroupParameterArray []ParameterGroupParameterInput

func (ParameterGroupParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameter)(nil)).Elem()
}

func (i ParameterGroupParameterArray) ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput {
	return i.ToParameterGroupParameterArrayOutputWithContext(context.Background())
}

func (i ParameterGroupParameterArray) ToParameterGroupParameterArrayOutputWithContext(ctx context.Context) ParameterGroupParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterArrayOutput)
}

type ParameterGroupParameterOutput struct{ *pulumi.OutputState }

func (ParameterGroupParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameter)(nil)).Elem()
}

func (o ParameterGroupParameterOutput) ToParameterGroupParameterOutput() ParameterGroupParameterOutput {
	return o
}

func (o ParameterGroupParameterOutput) ToParameterGroupParameterOutputWithContext(ctx context.Context) ParameterGroupParameterOutput {
	return o
}

// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
func (o ParameterGroupParameterOutput) ApplyMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterGroupParameter) *string { return v.ApplyMethod }).(pulumi.StringPtrOutput)
}

// The name of the Neptune parameter.
func (o ParameterGroupParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupParameter) string { return v.Name }).(pulumi.StringOutput)
}

// The value of the Neptune parameter.
func (o ParameterGroupParameterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupParameter) string { return v.Value }).(pulumi.StringOutput)
}

type ParameterGroupParameterArrayOutput struct{ *pulumi.OutputState }

func (ParameterGroupParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameter)(nil)).Elem()
}

func (o ParameterGroupParameterArrayOutput) ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput {
	return o
}

func (o ParameterGroupParameterArrayOutput) ToParameterGroupParameterArrayOutputWithContext(ctx context.Context) ParameterGroupParameterArrayOutput {
	return o
}

func (o ParameterGroupParameterArrayOutput) Index(i pulumi.IntInput) ParameterGroupParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterGroupParameter {
		return vs[0].([]ParameterGroupParameter)[vs[1].(int)]
	}).(ParameterGroupParameterOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterParameterGroupParameterOutput{})
	pulumi.RegisterOutputType(ClusterParameterGroupParameterArrayOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterArrayOutput{})
}
