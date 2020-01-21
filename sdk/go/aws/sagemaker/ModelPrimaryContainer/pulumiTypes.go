// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ModelPrimaryContainer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ModelPrimaryContainer struct {
	ContainerHostname *string `pulumi:"containerHostname"`
	Environment map[string]interface{} `pulumi:"environment"`
	Image string `pulumi:"image"`
	ModelDataUrl *string `pulumi:"modelDataUrl"`
}

type ModelPrimaryContainerInput interface {
	pulumi.Input

	ToModelPrimaryContainerOutput() ModelPrimaryContainerOutput
	ToModelPrimaryContainerOutputWithContext(context.Context) ModelPrimaryContainerOutput
}

type ModelPrimaryContainerArgs struct {
	ContainerHostname pulumi.StringPtrInput `pulumi:"containerHostname"`
	Environment pulumi.MapInput `pulumi:"environment"`
	Image pulumi.StringInput `pulumi:"image"`
	ModelDataUrl pulumi.StringPtrInput `pulumi:"modelDataUrl"`
}

func (ModelPrimaryContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelPrimaryContainer)(nil)).Elem()
}

func (i ModelPrimaryContainerArgs) ToModelPrimaryContainerOutput() ModelPrimaryContainerOutput {
	return i.ToModelPrimaryContainerOutputWithContext(context.Background())
}

func (i ModelPrimaryContainerArgs) ToModelPrimaryContainerOutputWithContext(ctx context.Context) ModelPrimaryContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelPrimaryContainerOutput)
}

func (i ModelPrimaryContainerArgs) ToModelPrimaryContainerPtrOutput() ModelPrimaryContainerPtrOutput {
	return i.ToModelPrimaryContainerPtrOutputWithContext(context.Background())
}

func (i ModelPrimaryContainerArgs) ToModelPrimaryContainerPtrOutputWithContext(ctx context.Context) ModelPrimaryContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelPrimaryContainerOutput).ToModelPrimaryContainerPtrOutputWithContext(ctx)
}

type ModelPrimaryContainerPtrInput interface {
	pulumi.Input

	ToModelPrimaryContainerPtrOutput() ModelPrimaryContainerPtrOutput
	ToModelPrimaryContainerPtrOutputWithContext(context.Context) ModelPrimaryContainerPtrOutput
}

type modelPrimaryContainerPtrType ModelPrimaryContainerArgs

func ModelPrimaryContainerPtr(v *ModelPrimaryContainerArgs) ModelPrimaryContainerPtrInput {	return (*modelPrimaryContainerPtrType)(v)
}

func (*modelPrimaryContainerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelPrimaryContainer)(nil)).Elem()
}

func (i *modelPrimaryContainerPtrType) ToModelPrimaryContainerPtrOutput() ModelPrimaryContainerPtrOutput {
	return i.ToModelPrimaryContainerPtrOutputWithContext(context.Background())
}

func (i *modelPrimaryContainerPtrType) ToModelPrimaryContainerPtrOutputWithContext(ctx context.Context) ModelPrimaryContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelPrimaryContainerPtrOutput)
}

type ModelPrimaryContainerOutput struct { *pulumi.OutputState }

func (ModelPrimaryContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModelPrimaryContainer)(nil)).Elem()
}

func (o ModelPrimaryContainerOutput) ToModelPrimaryContainerOutput() ModelPrimaryContainerOutput {
	return o
}

func (o ModelPrimaryContainerOutput) ToModelPrimaryContainerOutputWithContext(ctx context.Context) ModelPrimaryContainerOutput {
	return o
}

func (o ModelPrimaryContainerOutput) ToModelPrimaryContainerPtrOutput() ModelPrimaryContainerPtrOutput {
	return o.ToModelPrimaryContainerPtrOutputWithContext(context.Background())
}

func (o ModelPrimaryContainerOutput) ToModelPrimaryContainerPtrOutputWithContext(ctx context.Context) ModelPrimaryContainerPtrOutput {
	return o.ApplyT(func(v ModelPrimaryContainer) *ModelPrimaryContainer {
		return &v
	}).(ModelPrimaryContainerPtrOutput)
}
func (o ModelPrimaryContainerOutput) ContainerHostname() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ModelPrimaryContainer) *string { return v.ContainerHostname }).(pulumi.StringPtrOutput)
}

func (o ModelPrimaryContainerOutput) Environment() pulumi.MapOutput {
	return o.ApplyT(func (v ModelPrimaryContainer) map[string]interface{} { return v.Environment }).(pulumi.MapOutput)
}

func (o ModelPrimaryContainerOutput) Image() pulumi.StringOutput {
	return o.ApplyT(func (v ModelPrimaryContainer) string { return v.Image }).(pulumi.StringOutput)
}

func (o ModelPrimaryContainerOutput) ModelDataUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ModelPrimaryContainer) *string { return v.ModelDataUrl }).(pulumi.StringPtrOutput)
}

type ModelPrimaryContainerPtrOutput struct { *pulumi.OutputState}

func (ModelPrimaryContainerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelPrimaryContainer)(nil)).Elem()
}

func (o ModelPrimaryContainerPtrOutput) ToModelPrimaryContainerPtrOutput() ModelPrimaryContainerPtrOutput {
	return o
}

func (o ModelPrimaryContainerPtrOutput) ToModelPrimaryContainerPtrOutputWithContext(ctx context.Context) ModelPrimaryContainerPtrOutput {
	return o
}

func (o ModelPrimaryContainerPtrOutput) Elem() ModelPrimaryContainerOutput {
	return o.ApplyT(func (v *ModelPrimaryContainer) ModelPrimaryContainer { return *v }).(ModelPrimaryContainerOutput)
}

func (o ModelPrimaryContainerPtrOutput) ContainerHostname() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ModelPrimaryContainer) *string { return v.ContainerHostname }).(pulumi.StringPtrOutput)
}

func (o ModelPrimaryContainerPtrOutput) Environment() pulumi.MapOutput {
	return o.ApplyT(func (v ModelPrimaryContainer) map[string]interface{} { return v.Environment }).(pulumi.MapOutput)
}

func (o ModelPrimaryContainerPtrOutput) Image() pulumi.StringOutput {
	return o.ApplyT(func (v ModelPrimaryContainer) string { return v.Image }).(pulumi.StringOutput)
}

func (o ModelPrimaryContainerPtrOutput) ModelDataUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ModelPrimaryContainer) *string { return v.ModelDataUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ModelPrimaryContainerOutput{})
	pulumi.RegisterOutputType(ModelPrimaryContainerPtrOutput{})
}
