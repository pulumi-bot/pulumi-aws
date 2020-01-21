// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package LaunchTemplateElasticGpuSpecification

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type LaunchTemplateElasticGpuSpecification struct {
	// Accelerator type.
	Type string `pulumi:"type"`
}

type LaunchTemplateElasticGpuSpecificationInput interface {
	pulumi.Input

	ToLaunchTemplateElasticGpuSpecificationOutput() LaunchTemplateElasticGpuSpecificationOutput
	ToLaunchTemplateElasticGpuSpecificationOutputWithContext(context.Context) LaunchTemplateElasticGpuSpecificationOutput
}

type LaunchTemplateElasticGpuSpecificationArgs struct {
	// Accelerator type.
	Type pulumi.StringInput `pulumi:"type"`
}

func (LaunchTemplateElasticGpuSpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchTemplateElasticGpuSpecification)(nil)).Elem()
}

func (i LaunchTemplateElasticGpuSpecificationArgs) ToLaunchTemplateElasticGpuSpecificationOutput() LaunchTemplateElasticGpuSpecificationOutput {
	return i.ToLaunchTemplateElasticGpuSpecificationOutputWithContext(context.Background())
}

func (i LaunchTemplateElasticGpuSpecificationArgs) ToLaunchTemplateElasticGpuSpecificationOutputWithContext(ctx context.Context) LaunchTemplateElasticGpuSpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateElasticGpuSpecificationOutput)
}

type LaunchTemplateElasticGpuSpecificationArrayInput interface {
	pulumi.Input

	ToLaunchTemplateElasticGpuSpecificationArrayOutput() LaunchTemplateElasticGpuSpecificationArrayOutput
	ToLaunchTemplateElasticGpuSpecificationArrayOutputWithContext(context.Context) LaunchTemplateElasticGpuSpecificationArrayOutput
}

type LaunchTemplateElasticGpuSpecificationArray []LaunchTemplateElasticGpuSpecificationInput

func (LaunchTemplateElasticGpuSpecificationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LaunchTemplateElasticGpuSpecification)(nil)).Elem()
}

func (i LaunchTemplateElasticGpuSpecificationArray) ToLaunchTemplateElasticGpuSpecificationArrayOutput() LaunchTemplateElasticGpuSpecificationArrayOutput {
	return i.ToLaunchTemplateElasticGpuSpecificationArrayOutputWithContext(context.Background())
}

func (i LaunchTemplateElasticGpuSpecificationArray) ToLaunchTemplateElasticGpuSpecificationArrayOutputWithContext(ctx context.Context) LaunchTemplateElasticGpuSpecificationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateElasticGpuSpecificationArrayOutput)
}

type LaunchTemplateElasticGpuSpecificationOutput struct { *pulumi.OutputState }

func (LaunchTemplateElasticGpuSpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchTemplateElasticGpuSpecification)(nil)).Elem()
}

func (o LaunchTemplateElasticGpuSpecificationOutput) ToLaunchTemplateElasticGpuSpecificationOutput() LaunchTemplateElasticGpuSpecificationOutput {
	return o
}

func (o LaunchTemplateElasticGpuSpecificationOutput) ToLaunchTemplateElasticGpuSpecificationOutputWithContext(ctx context.Context) LaunchTemplateElasticGpuSpecificationOutput {
	return o
}

// Accelerator type.
func (o LaunchTemplateElasticGpuSpecificationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func (v LaunchTemplateElasticGpuSpecification) string { return v.Type }).(pulumi.StringOutput)
}

type LaunchTemplateElasticGpuSpecificationArrayOutput struct { *pulumi.OutputState}

func (LaunchTemplateElasticGpuSpecificationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LaunchTemplateElasticGpuSpecification)(nil)).Elem()
}

func (o LaunchTemplateElasticGpuSpecificationArrayOutput) ToLaunchTemplateElasticGpuSpecificationArrayOutput() LaunchTemplateElasticGpuSpecificationArrayOutput {
	return o
}

func (o LaunchTemplateElasticGpuSpecificationArrayOutput) ToLaunchTemplateElasticGpuSpecificationArrayOutputWithContext(ctx context.Context) LaunchTemplateElasticGpuSpecificationArrayOutput {
	return o
}

func (o LaunchTemplateElasticGpuSpecificationArrayOutput) Index(i pulumi.IntInput) LaunchTemplateElasticGpuSpecificationOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) LaunchTemplateElasticGpuSpecification {
		return vs[0].([]LaunchTemplateElasticGpuSpecification)[vs[1].(int)]
	}).(LaunchTemplateElasticGpuSpecificationOutput)
}

func init() {
	pulumi.RegisterOutputType(LaunchTemplateElasticGpuSpecificationOutput{})
	pulumi.RegisterOutputType(LaunchTemplateElasticGpuSpecificationArrayOutput{})
}
