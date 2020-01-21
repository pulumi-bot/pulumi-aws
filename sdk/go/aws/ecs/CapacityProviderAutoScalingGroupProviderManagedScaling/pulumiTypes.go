// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package CapacityProviderAutoScalingGroupProviderManagedScaling

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type CapacityProviderAutoScalingGroupProviderManagedScaling struct {
	MaximumScalingStepSize *int `pulumi:"maximumScalingStepSize"`
	MinimumScalingStepSize *int `pulumi:"minimumScalingStepSize"`
	Status *string `pulumi:"status"`
	TargetCapacity *int `pulumi:"targetCapacity"`
}

type CapacityProviderAutoScalingGroupProviderManagedScalingInput interface {
	pulumi.Input

	ToCapacityProviderAutoScalingGroupProviderManagedScalingOutput() CapacityProviderAutoScalingGroupProviderManagedScalingOutput
	ToCapacityProviderAutoScalingGroupProviderManagedScalingOutputWithContext(context.Context) CapacityProviderAutoScalingGroupProviderManagedScalingOutput
}

type CapacityProviderAutoScalingGroupProviderManagedScalingArgs struct {
	MaximumScalingStepSize pulumi.IntPtrInput `pulumi:"maximumScalingStepSize"`
	MinimumScalingStepSize pulumi.IntPtrInput `pulumi:"minimumScalingStepSize"`
	Status pulumi.StringPtrInput `pulumi:"status"`
	TargetCapacity pulumi.IntPtrInput `pulumi:"targetCapacity"`
}

func (CapacityProviderAutoScalingGroupProviderManagedScalingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityProviderAutoScalingGroupProviderManagedScaling)(nil)).Elem()
}

func (i CapacityProviderAutoScalingGroupProviderManagedScalingArgs) ToCapacityProviderAutoScalingGroupProviderManagedScalingOutput() CapacityProviderAutoScalingGroupProviderManagedScalingOutput {
	return i.ToCapacityProviderAutoScalingGroupProviderManagedScalingOutputWithContext(context.Background())
}

func (i CapacityProviderAutoScalingGroupProviderManagedScalingArgs) ToCapacityProviderAutoScalingGroupProviderManagedScalingOutputWithContext(ctx context.Context) CapacityProviderAutoScalingGroupProviderManagedScalingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityProviderAutoScalingGroupProviderManagedScalingOutput)
}

func (i CapacityProviderAutoScalingGroupProviderManagedScalingArgs) ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput() CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput {
	return i.ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutputWithContext(context.Background())
}

func (i CapacityProviderAutoScalingGroupProviderManagedScalingArgs) ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutputWithContext(ctx context.Context) CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityProviderAutoScalingGroupProviderManagedScalingOutput).ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutputWithContext(ctx)
}

type CapacityProviderAutoScalingGroupProviderManagedScalingPtrInput interface {
	pulumi.Input

	ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput() CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput
	ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutputWithContext(context.Context) CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput
}

type capacityProviderAutoScalingGroupProviderManagedScalingPtrType CapacityProviderAutoScalingGroupProviderManagedScalingArgs

func CapacityProviderAutoScalingGroupProviderManagedScalingPtr(v *CapacityProviderAutoScalingGroupProviderManagedScalingArgs) CapacityProviderAutoScalingGroupProviderManagedScalingPtrInput {	return (*capacityProviderAutoScalingGroupProviderManagedScalingPtrType)(v)
}

func (*capacityProviderAutoScalingGroupProviderManagedScalingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityProviderAutoScalingGroupProviderManagedScaling)(nil)).Elem()
}

func (i *capacityProviderAutoScalingGroupProviderManagedScalingPtrType) ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput() CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput {
	return i.ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutputWithContext(context.Background())
}

func (i *capacityProviderAutoScalingGroupProviderManagedScalingPtrType) ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutputWithContext(ctx context.Context) CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput)
}

type CapacityProviderAutoScalingGroupProviderManagedScalingOutput struct { *pulumi.OutputState }

func (CapacityProviderAutoScalingGroupProviderManagedScalingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityProviderAutoScalingGroupProviderManagedScaling)(nil)).Elem()
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingOutput) ToCapacityProviderAutoScalingGroupProviderManagedScalingOutput() CapacityProviderAutoScalingGroupProviderManagedScalingOutput {
	return o
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingOutput) ToCapacityProviderAutoScalingGroupProviderManagedScalingOutputWithContext(ctx context.Context) CapacityProviderAutoScalingGroupProviderManagedScalingOutput {
	return o
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingOutput) ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput() CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput {
	return o.ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutputWithContext(context.Background())
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingOutput) ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutputWithContext(ctx context.Context) CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput {
	return o.ApplyT(func(v CapacityProviderAutoScalingGroupProviderManagedScaling) *CapacityProviderAutoScalingGroupProviderManagedScaling {
		return &v
	}).(CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput)
}
func (o CapacityProviderAutoScalingGroupProviderManagedScalingOutput) MaximumScalingStepSize() pulumi.IntPtrOutput {
	return o.ApplyT(func (v CapacityProviderAutoScalingGroupProviderManagedScaling) *int { return v.MaximumScalingStepSize }).(pulumi.IntPtrOutput)
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingOutput) MinimumScalingStepSize() pulumi.IntPtrOutput {
	return o.ApplyT(func (v CapacityProviderAutoScalingGroupProviderManagedScaling) *int { return v.MinimumScalingStepSize }).(pulumi.IntPtrOutput)
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func (v CapacityProviderAutoScalingGroupProviderManagedScaling) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingOutput) TargetCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func (v CapacityProviderAutoScalingGroupProviderManagedScaling) *int { return v.TargetCapacity }).(pulumi.IntPtrOutput)
}

type CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput struct { *pulumi.OutputState}

func (CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityProviderAutoScalingGroupProviderManagedScaling)(nil)).Elem()
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput) ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput() CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput {
	return o
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput) ToCapacityProviderAutoScalingGroupProviderManagedScalingPtrOutputWithContext(ctx context.Context) CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput {
	return o
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput) Elem() CapacityProviderAutoScalingGroupProviderManagedScalingOutput {
	return o.ApplyT(func (v *CapacityProviderAutoScalingGroupProviderManagedScaling) CapacityProviderAutoScalingGroupProviderManagedScaling { return *v }).(CapacityProviderAutoScalingGroupProviderManagedScalingOutput)
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput) MaximumScalingStepSize() pulumi.IntPtrOutput {
	return o.ApplyT(func (v CapacityProviderAutoScalingGroupProviderManagedScaling) *int { return v.MaximumScalingStepSize }).(pulumi.IntPtrOutput)
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput) MinimumScalingStepSize() pulumi.IntPtrOutput {
	return o.ApplyT(func (v CapacityProviderAutoScalingGroupProviderManagedScaling) *int { return v.MinimumScalingStepSize }).(pulumi.IntPtrOutput)
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func (v CapacityProviderAutoScalingGroupProviderManagedScaling) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput) TargetCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func (v CapacityProviderAutoScalingGroupProviderManagedScaling) *int { return v.TargetCapacity }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CapacityProviderAutoScalingGroupProviderManagedScalingOutput{})
	pulumi.RegisterOutputType(CapacityProviderAutoScalingGroupProviderManagedScalingPtrOutput{})
}
