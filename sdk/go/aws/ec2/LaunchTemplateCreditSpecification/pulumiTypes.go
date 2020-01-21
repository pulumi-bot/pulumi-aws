// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package LaunchTemplateCreditSpecification

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type LaunchTemplateCreditSpecification struct {
	CpuCredits *string `pulumi:"cpuCredits"`
}

type LaunchTemplateCreditSpecificationInput interface {
	pulumi.Input

	ToLaunchTemplateCreditSpecificationOutput() LaunchTemplateCreditSpecificationOutput
	ToLaunchTemplateCreditSpecificationOutputWithContext(context.Context) LaunchTemplateCreditSpecificationOutput
}

type LaunchTemplateCreditSpecificationArgs struct {
	CpuCredits pulumi.StringPtrInput `pulumi:"cpuCredits"`
}

func (LaunchTemplateCreditSpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchTemplateCreditSpecification)(nil)).Elem()
}

func (i LaunchTemplateCreditSpecificationArgs) ToLaunchTemplateCreditSpecificationOutput() LaunchTemplateCreditSpecificationOutput {
	return i.ToLaunchTemplateCreditSpecificationOutputWithContext(context.Background())
}

func (i LaunchTemplateCreditSpecificationArgs) ToLaunchTemplateCreditSpecificationOutputWithContext(ctx context.Context) LaunchTemplateCreditSpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateCreditSpecificationOutput)
}

func (i LaunchTemplateCreditSpecificationArgs) ToLaunchTemplateCreditSpecificationPtrOutput() LaunchTemplateCreditSpecificationPtrOutput {
	return i.ToLaunchTemplateCreditSpecificationPtrOutputWithContext(context.Background())
}

func (i LaunchTemplateCreditSpecificationArgs) ToLaunchTemplateCreditSpecificationPtrOutputWithContext(ctx context.Context) LaunchTemplateCreditSpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateCreditSpecificationOutput).ToLaunchTemplateCreditSpecificationPtrOutputWithContext(ctx)
}

type LaunchTemplateCreditSpecificationPtrInput interface {
	pulumi.Input

	ToLaunchTemplateCreditSpecificationPtrOutput() LaunchTemplateCreditSpecificationPtrOutput
	ToLaunchTemplateCreditSpecificationPtrOutputWithContext(context.Context) LaunchTemplateCreditSpecificationPtrOutput
}

type launchTemplateCreditSpecificationPtrType LaunchTemplateCreditSpecificationArgs

func LaunchTemplateCreditSpecificationPtr(v *LaunchTemplateCreditSpecificationArgs) LaunchTemplateCreditSpecificationPtrInput {	return (*launchTemplateCreditSpecificationPtrType)(v)
}

func (*launchTemplateCreditSpecificationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchTemplateCreditSpecification)(nil)).Elem()
}

func (i *launchTemplateCreditSpecificationPtrType) ToLaunchTemplateCreditSpecificationPtrOutput() LaunchTemplateCreditSpecificationPtrOutput {
	return i.ToLaunchTemplateCreditSpecificationPtrOutputWithContext(context.Background())
}

func (i *launchTemplateCreditSpecificationPtrType) ToLaunchTemplateCreditSpecificationPtrOutputWithContext(ctx context.Context) LaunchTemplateCreditSpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateCreditSpecificationPtrOutput)
}

type LaunchTemplateCreditSpecificationOutput struct { *pulumi.OutputState }

func (LaunchTemplateCreditSpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchTemplateCreditSpecification)(nil)).Elem()
}

func (o LaunchTemplateCreditSpecificationOutput) ToLaunchTemplateCreditSpecificationOutput() LaunchTemplateCreditSpecificationOutput {
	return o
}

func (o LaunchTemplateCreditSpecificationOutput) ToLaunchTemplateCreditSpecificationOutputWithContext(ctx context.Context) LaunchTemplateCreditSpecificationOutput {
	return o
}

func (o LaunchTemplateCreditSpecificationOutput) ToLaunchTemplateCreditSpecificationPtrOutput() LaunchTemplateCreditSpecificationPtrOutput {
	return o.ToLaunchTemplateCreditSpecificationPtrOutputWithContext(context.Background())
}

func (o LaunchTemplateCreditSpecificationOutput) ToLaunchTemplateCreditSpecificationPtrOutputWithContext(ctx context.Context) LaunchTemplateCreditSpecificationPtrOutput {
	return o.ApplyT(func(v LaunchTemplateCreditSpecification) *LaunchTemplateCreditSpecification {
		return &v
	}).(LaunchTemplateCreditSpecificationPtrOutput)
}
func (o LaunchTemplateCreditSpecificationOutput) CpuCredits() pulumi.StringPtrOutput {
	return o.ApplyT(func (v LaunchTemplateCreditSpecification) *string { return v.CpuCredits }).(pulumi.StringPtrOutput)
}

type LaunchTemplateCreditSpecificationPtrOutput struct { *pulumi.OutputState}

func (LaunchTemplateCreditSpecificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchTemplateCreditSpecification)(nil)).Elem()
}

func (o LaunchTemplateCreditSpecificationPtrOutput) ToLaunchTemplateCreditSpecificationPtrOutput() LaunchTemplateCreditSpecificationPtrOutput {
	return o
}

func (o LaunchTemplateCreditSpecificationPtrOutput) ToLaunchTemplateCreditSpecificationPtrOutputWithContext(ctx context.Context) LaunchTemplateCreditSpecificationPtrOutput {
	return o
}

func (o LaunchTemplateCreditSpecificationPtrOutput) Elem() LaunchTemplateCreditSpecificationOutput {
	return o.ApplyT(func (v *LaunchTemplateCreditSpecification) LaunchTemplateCreditSpecification { return *v }).(LaunchTemplateCreditSpecificationOutput)
}

func (o LaunchTemplateCreditSpecificationPtrOutput) CpuCredits() pulumi.StringPtrOutput {
	return o.ApplyT(func (v LaunchTemplateCreditSpecification) *string { return v.CpuCredits }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LaunchTemplateCreditSpecificationOutput{})
	pulumi.RegisterOutputType(LaunchTemplateCreditSpecificationPtrOutput{})
}
