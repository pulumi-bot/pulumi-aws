// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClusterParameterGroupParameter struct {
	// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod *string `pulumi:"applyMethod"`
	// The name of the neptune parameter.
	Name string `pulumi:"name"`
	// The value of the neptune parameter.
	Value string `pulumi:"value"`
}

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

type ClusterParameterGroupParameterArgs struct {
	// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod *string `pulumi:"applyMethod"`
	// The name of the neptune parameter.
	Name string `pulumi:"name"`
	// The value of the neptune parameter.
	Value string `pulumi:"value"`
}

type ClusterParameterGroupParameterArgsInput interface {
	pulumi.Input

	ToClusterParameterGroupParameterArgsOutput() ClusterParameterGroupParameterArgsOutput
	ToClusterParameterGroupParameterArgsOutputWithContext(context.Context) ClusterParameterGroupParameterArgsOutput
}

type ClusterParameterGroupParameterArgsArgs struct {
	// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod pulumi.StringPtrInput `pulumi:"applyMethod"`
	// The name of the neptune parameter.
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the neptune parameter.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ClusterParameterGroupParameterArgsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParameterGroupParameterArgs)(nil)).Elem()
}

func (i ClusterParameterGroupParameterArgsArgs) ToClusterParameterGroupParameterArgsOutput() ClusterParameterGroupParameterArgsOutput {
	return i.ToClusterParameterGroupParameterArgsOutputWithContext(context.Background())
}

func (i ClusterParameterGroupParameterArgsArgs) ToClusterParameterGroupParameterArgsOutputWithContext(ctx context.Context) ClusterParameterGroupParameterArgsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupParameterArgsOutput)
}

type ClusterParameterGroupParameterArgsArrayInput interface {
	pulumi.Input

	ToClusterParameterGroupParameterArgsArrayOutput() ClusterParameterGroupParameterArgsArrayOutput
	ToClusterParameterGroupParameterArgsArrayOutputWithContext(context.Context) ClusterParameterGroupParameterArgsArrayOutput
}

type ClusterParameterGroupParameterArgsArray []ClusterParameterGroupParameterArgsInput

func (ClusterParameterGroupParameterArgsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterParameterGroupParameterArgs)(nil)).Elem()
}

func (i ClusterParameterGroupParameterArgsArray) ToClusterParameterGroupParameterArgsArrayOutput() ClusterParameterGroupParameterArgsArrayOutput {
	return i.ToClusterParameterGroupParameterArgsArrayOutputWithContext(context.Background())
}

func (i ClusterParameterGroupParameterArgsArray) ToClusterParameterGroupParameterArgsArrayOutputWithContext(ctx context.Context) ClusterParameterGroupParameterArgsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupParameterArgsArrayOutput)
}

type ClusterParameterGroupParameterArgsOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupParameterArgsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParameterGroupParameterArgs)(nil)).Elem()
}

func (o ClusterParameterGroupParameterArgsOutput) ToClusterParameterGroupParameterArgsOutput() ClusterParameterGroupParameterArgsOutput {
	return o
}

func (o ClusterParameterGroupParameterArgsOutput) ToClusterParameterGroupParameterArgsOutputWithContext(ctx context.Context) ClusterParameterGroupParameterArgsOutput {
	return o
}

// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
func (o ClusterParameterGroupParameterArgsOutput) ApplyMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterParameterGroupParameterArgs) *string { return v.ApplyMethod }).(pulumi.StringPtrOutput)
}

// The name of the neptune parameter.
func (o ClusterParameterGroupParameterArgsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterParameterGroupParameterArgs) string { return v.Name }).(pulumi.StringOutput)
}

// The value of the neptune parameter.
func (o ClusterParameterGroupParameterArgsOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterParameterGroupParameterArgs) string { return v.Value }).(pulumi.StringOutput)
}

type ClusterParameterGroupParameterArgsArrayOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupParameterArgsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterParameterGroupParameterArgs)(nil)).Elem()
}

func (o ClusterParameterGroupParameterArgsArrayOutput) ToClusterParameterGroupParameterArgsArrayOutput() ClusterParameterGroupParameterArgsArrayOutput {
	return o
}

func (o ClusterParameterGroupParameterArgsArrayOutput) ToClusterParameterGroupParameterArgsArrayOutputWithContext(ctx context.Context) ClusterParameterGroupParameterArgsArrayOutput {
	return o
}

func (o ClusterParameterGroupParameterArgsArrayOutput) Index(i pulumi.IntInput) ClusterParameterGroupParameterArgsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterParameterGroupParameterArgs {
		return vs[0].([]ClusterParameterGroupParameterArgs)[vs[1].(int)]
	}).(ClusterParameterGroupParameterArgsOutput)
}

type ClusterParameterGroupParameterState struct {
	// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod *string `pulumi:"applyMethod"`
	// The name of the neptune parameter.
	Name string `pulumi:"name"`
	// The value of the neptune parameter.
	Value string `pulumi:"value"`
}

type ClusterParameterGroupParameterStateInput interface {
	pulumi.Input

	ToClusterParameterGroupParameterStateOutput() ClusterParameterGroupParameterStateOutput
	ToClusterParameterGroupParameterStateOutputWithContext(context.Context) ClusterParameterGroupParameterStateOutput
}

type ClusterParameterGroupParameterStateArgs struct {
	// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod pulumi.StringPtrInput `pulumi:"applyMethod"`
	// The name of the neptune parameter.
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the neptune parameter.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ClusterParameterGroupParameterStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParameterGroupParameterState)(nil)).Elem()
}

func (i ClusterParameterGroupParameterStateArgs) ToClusterParameterGroupParameterStateOutput() ClusterParameterGroupParameterStateOutput {
	return i.ToClusterParameterGroupParameterStateOutputWithContext(context.Background())
}

func (i ClusterParameterGroupParameterStateArgs) ToClusterParameterGroupParameterStateOutputWithContext(ctx context.Context) ClusterParameterGroupParameterStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupParameterStateOutput)
}

type ClusterParameterGroupParameterStateArrayInput interface {
	pulumi.Input

	ToClusterParameterGroupParameterStateArrayOutput() ClusterParameterGroupParameterStateArrayOutput
	ToClusterParameterGroupParameterStateArrayOutputWithContext(context.Context) ClusterParameterGroupParameterStateArrayOutput
}

type ClusterParameterGroupParameterStateArray []ClusterParameterGroupParameterStateInput

func (ClusterParameterGroupParameterStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterParameterGroupParameterState)(nil)).Elem()
}

func (i ClusterParameterGroupParameterStateArray) ToClusterParameterGroupParameterStateArrayOutput() ClusterParameterGroupParameterStateArrayOutput {
	return i.ToClusterParameterGroupParameterStateArrayOutputWithContext(context.Background())
}

func (i ClusterParameterGroupParameterStateArray) ToClusterParameterGroupParameterStateArrayOutputWithContext(ctx context.Context) ClusterParameterGroupParameterStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupParameterStateArrayOutput)
}

type ClusterParameterGroupParameterStateOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupParameterStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParameterGroupParameterState)(nil)).Elem()
}

func (o ClusterParameterGroupParameterStateOutput) ToClusterParameterGroupParameterStateOutput() ClusterParameterGroupParameterStateOutput {
	return o
}

func (o ClusterParameterGroupParameterStateOutput) ToClusterParameterGroupParameterStateOutputWithContext(ctx context.Context) ClusterParameterGroupParameterStateOutput {
	return o
}

// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
func (o ClusterParameterGroupParameterStateOutput) ApplyMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterParameterGroupParameterState) *string { return v.ApplyMethod }).(pulumi.StringPtrOutput)
}

// The name of the neptune parameter.
func (o ClusterParameterGroupParameterStateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterParameterGroupParameterState) string { return v.Name }).(pulumi.StringOutput)
}

// The value of the neptune parameter.
func (o ClusterParameterGroupParameterStateOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterParameterGroupParameterState) string { return v.Value }).(pulumi.StringOutput)
}

type ClusterParameterGroupParameterStateArrayOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupParameterStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterParameterGroupParameterState)(nil)).Elem()
}

func (o ClusterParameterGroupParameterStateArrayOutput) ToClusterParameterGroupParameterStateArrayOutput() ClusterParameterGroupParameterStateArrayOutput {
	return o
}

func (o ClusterParameterGroupParameterStateArrayOutput) ToClusterParameterGroupParameterStateArrayOutputWithContext(ctx context.Context) ClusterParameterGroupParameterStateArrayOutput {
	return o
}

func (o ClusterParameterGroupParameterStateArrayOutput) Index(i pulumi.IntInput) ClusterParameterGroupParameterStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterParameterGroupParameterState {
		return vs[0].([]ClusterParameterGroupParameterState)[vs[1].(int)]
	}).(ClusterParameterGroupParameterStateOutput)
}

type ParameterGroupParameter struct {
	// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod *string `pulumi:"applyMethod"`
	// The name of the Neptune parameter.
	Name string `pulumi:"name"`
	// The value of the Neptune parameter.
	Value string `pulumi:"value"`
}

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

type ParameterGroupParameterArgs struct {
	// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod *string `pulumi:"applyMethod"`
	// The name of the Neptune parameter.
	Name string `pulumi:"name"`
	// The value of the Neptune parameter.
	Value string `pulumi:"value"`
}

type ParameterGroupParameterArgsInput interface {
	pulumi.Input

	ToParameterGroupParameterArgsOutput() ParameterGroupParameterArgsOutput
	ToParameterGroupParameterArgsOutputWithContext(context.Context) ParameterGroupParameterArgsOutput
}

type ParameterGroupParameterArgsArgs struct {
	// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod pulumi.StringPtrInput `pulumi:"applyMethod"`
	// The name of the Neptune parameter.
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the Neptune parameter.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ParameterGroupParameterArgsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameterArgs)(nil)).Elem()
}

func (i ParameterGroupParameterArgsArgs) ToParameterGroupParameterArgsOutput() ParameterGroupParameterArgsOutput {
	return i.ToParameterGroupParameterArgsOutputWithContext(context.Background())
}

func (i ParameterGroupParameterArgsArgs) ToParameterGroupParameterArgsOutputWithContext(ctx context.Context) ParameterGroupParameterArgsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterArgsOutput)
}

type ParameterGroupParameterArgsArrayInput interface {
	pulumi.Input

	ToParameterGroupParameterArgsArrayOutput() ParameterGroupParameterArgsArrayOutput
	ToParameterGroupParameterArgsArrayOutputWithContext(context.Context) ParameterGroupParameterArgsArrayOutput
}

type ParameterGroupParameterArgsArray []ParameterGroupParameterArgsInput

func (ParameterGroupParameterArgsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameterArgs)(nil)).Elem()
}

func (i ParameterGroupParameterArgsArray) ToParameterGroupParameterArgsArrayOutput() ParameterGroupParameterArgsArrayOutput {
	return i.ToParameterGroupParameterArgsArrayOutputWithContext(context.Background())
}

func (i ParameterGroupParameterArgsArray) ToParameterGroupParameterArgsArrayOutputWithContext(ctx context.Context) ParameterGroupParameterArgsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterArgsArrayOutput)
}

type ParameterGroupParameterArgsOutput struct{ *pulumi.OutputState }

func (ParameterGroupParameterArgsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameterArgs)(nil)).Elem()
}

func (o ParameterGroupParameterArgsOutput) ToParameterGroupParameterArgsOutput() ParameterGroupParameterArgsOutput {
	return o
}

func (o ParameterGroupParameterArgsOutput) ToParameterGroupParameterArgsOutputWithContext(ctx context.Context) ParameterGroupParameterArgsOutput {
	return o
}

// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
func (o ParameterGroupParameterArgsOutput) ApplyMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterGroupParameterArgs) *string { return v.ApplyMethod }).(pulumi.StringPtrOutput)
}

// The name of the Neptune parameter.
func (o ParameterGroupParameterArgsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupParameterArgs) string { return v.Name }).(pulumi.StringOutput)
}

// The value of the Neptune parameter.
func (o ParameterGroupParameterArgsOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupParameterArgs) string { return v.Value }).(pulumi.StringOutput)
}

type ParameterGroupParameterArgsArrayOutput struct{ *pulumi.OutputState }

func (ParameterGroupParameterArgsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameterArgs)(nil)).Elem()
}

func (o ParameterGroupParameterArgsArrayOutput) ToParameterGroupParameterArgsArrayOutput() ParameterGroupParameterArgsArrayOutput {
	return o
}

func (o ParameterGroupParameterArgsArrayOutput) ToParameterGroupParameterArgsArrayOutputWithContext(ctx context.Context) ParameterGroupParameterArgsArrayOutput {
	return o
}

func (o ParameterGroupParameterArgsArrayOutput) Index(i pulumi.IntInput) ParameterGroupParameterArgsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterGroupParameterArgs {
		return vs[0].([]ParameterGroupParameterArgs)[vs[1].(int)]
	}).(ParameterGroupParameterArgsOutput)
}

type ParameterGroupParameterState struct {
	// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod *string `pulumi:"applyMethod"`
	// The name of the Neptune parameter.
	Name string `pulumi:"name"`
	// The value of the Neptune parameter.
	Value string `pulumi:"value"`
}

type ParameterGroupParameterStateInput interface {
	pulumi.Input

	ToParameterGroupParameterStateOutput() ParameterGroupParameterStateOutput
	ToParameterGroupParameterStateOutputWithContext(context.Context) ParameterGroupParameterStateOutput
}

type ParameterGroupParameterStateArgs struct {
	// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod pulumi.StringPtrInput `pulumi:"applyMethod"`
	// The name of the Neptune parameter.
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the Neptune parameter.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ParameterGroupParameterStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameterState)(nil)).Elem()
}

func (i ParameterGroupParameterStateArgs) ToParameterGroupParameterStateOutput() ParameterGroupParameterStateOutput {
	return i.ToParameterGroupParameterStateOutputWithContext(context.Background())
}

func (i ParameterGroupParameterStateArgs) ToParameterGroupParameterStateOutputWithContext(ctx context.Context) ParameterGroupParameterStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterStateOutput)
}

type ParameterGroupParameterStateArrayInput interface {
	pulumi.Input

	ToParameterGroupParameterStateArrayOutput() ParameterGroupParameterStateArrayOutput
	ToParameterGroupParameterStateArrayOutputWithContext(context.Context) ParameterGroupParameterStateArrayOutput
}

type ParameterGroupParameterStateArray []ParameterGroupParameterStateInput

func (ParameterGroupParameterStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameterState)(nil)).Elem()
}

func (i ParameterGroupParameterStateArray) ToParameterGroupParameterStateArrayOutput() ParameterGroupParameterStateArrayOutput {
	return i.ToParameterGroupParameterStateArrayOutputWithContext(context.Background())
}

func (i ParameterGroupParameterStateArray) ToParameterGroupParameterStateArrayOutputWithContext(ctx context.Context) ParameterGroupParameterStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterStateArrayOutput)
}

type ParameterGroupParameterStateOutput struct{ *pulumi.OutputState }

func (ParameterGroupParameterStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameterState)(nil)).Elem()
}

func (o ParameterGroupParameterStateOutput) ToParameterGroupParameterStateOutput() ParameterGroupParameterStateOutput {
	return o
}

func (o ParameterGroupParameterStateOutput) ToParameterGroupParameterStateOutputWithContext(ctx context.Context) ParameterGroupParameterStateOutput {
	return o
}

// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
func (o ParameterGroupParameterStateOutput) ApplyMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterGroupParameterState) *string { return v.ApplyMethod }).(pulumi.StringPtrOutput)
}

// The name of the Neptune parameter.
func (o ParameterGroupParameterStateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupParameterState) string { return v.Name }).(pulumi.StringOutput)
}

// The value of the Neptune parameter.
func (o ParameterGroupParameterStateOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupParameterState) string { return v.Value }).(pulumi.StringOutput)
}

type ParameterGroupParameterStateArrayOutput struct{ *pulumi.OutputState }

func (ParameterGroupParameterStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameterState)(nil)).Elem()
}

func (o ParameterGroupParameterStateArrayOutput) ToParameterGroupParameterStateArrayOutput() ParameterGroupParameterStateArrayOutput {
	return o
}

func (o ParameterGroupParameterStateArrayOutput) ToParameterGroupParameterStateArrayOutputWithContext(ctx context.Context) ParameterGroupParameterStateArrayOutput {
	return o
}

func (o ParameterGroupParameterStateArrayOutput) Index(i pulumi.IntInput) ParameterGroupParameterStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterGroupParameterState {
		return vs[0].([]ParameterGroupParameterState)[vs[1].(int)]
	}).(ParameterGroupParameterStateOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterParameterGroupParameterOutput{})
	pulumi.RegisterOutputType(ClusterParameterGroupParameterArrayOutput{})
	pulumi.RegisterOutputType(ClusterParameterGroupParameterArgsOutput{})
	pulumi.RegisterOutputType(ClusterParameterGroupParameterArgsArrayOutput{})
	pulumi.RegisterOutputType(ClusterParameterGroupParameterStateOutput{})
	pulumi.RegisterOutputType(ClusterParameterGroupParameterStateArrayOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterArrayOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterArgsOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterArgsArrayOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterStateOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterStateArrayOutput{})
}
