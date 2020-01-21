// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ComputeEnvironmentComputeResourcesLaunchTemplate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ComputeEnvironmentComputeResourcesLaunchTemplate struct {
	// ID of the launch template. You must specify either the launch template ID or launch template name in the request, but not both.
	LaunchTemplateId *string `pulumi:"launchTemplateId"`
	// Name of the launch template.
	LaunchTemplateName *string `pulumi:"launchTemplateName"`
	// The version number of the launch template. Default: The default version of the launch template.
	Version *string `pulumi:"version"`
}

type ComputeEnvironmentComputeResourcesLaunchTemplateInput interface {
	pulumi.Input

	ToComputeEnvironmentComputeResourcesLaunchTemplateOutput() ComputeEnvironmentComputeResourcesLaunchTemplateOutput
	ToComputeEnvironmentComputeResourcesLaunchTemplateOutputWithContext(context.Context) ComputeEnvironmentComputeResourcesLaunchTemplateOutput
}

type ComputeEnvironmentComputeResourcesLaunchTemplateArgs struct {
	// ID of the launch template. You must specify either the launch template ID or launch template name in the request, but not both.
	LaunchTemplateId pulumi.StringPtrInput `pulumi:"launchTemplateId"`
	// Name of the launch template.
	LaunchTemplateName pulumi.StringPtrInput `pulumi:"launchTemplateName"`
	// The version number of the launch template. Default: The default version of the launch template.
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (ComputeEnvironmentComputeResourcesLaunchTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeEnvironmentComputeResourcesLaunchTemplate)(nil)).Elem()
}

func (i ComputeEnvironmentComputeResourcesLaunchTemplateArgs) ToComputeEnvironmentComputeResourcesLaunchTemplateOutput() ComputeEnvironmentComputeResourcesLaunchTemplateOutput {
	return i.ToComputeEnvironmentComputeResourcesLaunchTemplateOutputWithContext(context.Background())
}

func (i ComputeEnvironmentComputeResourcesLaunchTemplateArgs) ToComputeEnvironmentComputeResourcesLaunchTemplateOutputWithContext(ctx context.Context) ComputeEnvironmentComputeResourcesLaunchTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeEnvironmentComputeResourcesLaunchTemplateOutput)
}

func (i ComputeEnvironmentComputeResourcesLaunchTemplateArgs) ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput() ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput {
	return i.ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutputWithContext(context.Background())
}

func (i ComputeEnvironmentComputeResourcesLaunchTemplateArgs) ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutputWithContext(ctx context.Context) ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeEnvironmentComputeResourcesLaunchTemplateOutput).ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutputWithContext(ctx)
}

type ComputeEnvironmentComputeResourcesLaunchTemplatePtrInput interface {
	pulumi.Input

	ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput() ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput
	ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutputWithContext(context.Context) ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput
}

type computeEnvironmentComputeResourcesLaunchTemplatePtrType ComputeEnvironmentComputeResourcesLaunchTemplateArgs

func ComputeEnvironmentComputeResourcesLaunchTemplatePtr(v *ComputeEnvironmentComputeResourcesLaunchTemplateArgs) ComputeEnvironmentComputeResourcesLaunchTemplatePtrInput {	return (*computeEnvironmentComputeResourcesLaunchTemplatePtrType)(v)
}

func (*computeEnvironmentComputeResourcesLaunchTemplatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeEnvironmentComputeResourcesLaunchTemplate)(nil)).Elem()
}

func (i *computeEnvironmentComputeResourcesLaunchTemplatePtrType) ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput() ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput {
	return i.ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutputWithContext(context.Background())
}

func (i *computeEnvironmentComputeResourcesLaunchTemplatePtrType) ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutputWithContext(ctx context.Context) ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput)
}

type ComputeEnvironmentComputeResourcesLaunchTemplateOutput struct { *pulumi.OutputState }

func (ComputeEnvironmentComputeResourcesLaunchTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeEnvironmentComputeResourcesLaunchTemplate)(nil)).Elem()
}

func (o ComputeEnvironmentComputeResourcesLaunchTemplateOutput) ToComputeEnvironmentComputeResourcesLaunchTemplateOutput() ComputeEnvironmentComputeResourcesLaunchTemplateOutput {
	return o
}

func (o ComputeEnvironmentComputeResourcesLaunchTemplateOutput) ToComputeEnvironmentComputeResourcesLaunchTemplateOutputWithContext(ctx context.Context) ComputeEnvironmentComputeResourcesLaunchTemplateOutput {
	return o
}

func (o ComputeEnvironmentComputeResourcesLaunchTemplateOutput) ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput() ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput {
	return o.ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutputWithContext(context.Background())
}

func (o ComputeEnvironmentComputeResourcesLaunchTemplateOutput) ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutputWithContext(ctx context.Context) ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput {
	return o.ApplyT(func(v ComputeEnvironmentComputeResourcesLaunchTemplate) *ComputeEnvironmentComputeResourcesLaunchTemplate {
		return &v
	}).(ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput)
}
// ID of the launch template. You must specify either the launch template ID or launch template name in the request, but not both.
func (o ComputeEnvironmentComputeResourcesLaunchTemplateOutput) LaunchTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ComputeEnvironmentComputeResourcesLaunchTemplate) *string { return v.LaunchTemplateId }).(pulumi.StringPtrOutput)
}

// Name of the launch template.
func (o ComputeEnvironmentComputeResourcesLaunchTemplateOutput) LaunchTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ComputeEnvironmentComputeResourcesLaunchTemplate) *string { return v.LaunchTemplateName }).(pulumi.StringPtrOutput)
}

// The version number of the launch template. Default: The default version of the launch template.
func (o ComputeEnvironmentComputeResourcesLaunchTemplateOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ComputeEnvironmentComputeResourcesLaunchTemplate) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput struct { *pulumi.OutputState}

func (ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeEnvironmentComputeResourcesLaunchTemplate)(nil)).Elem()
}

func (o ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput) ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput() ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput {
	return o
}

func (o ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput) ToComputeEnvironmentComputeResourcesLaunchTemplatePtrOutputWithContext(ctx context.Context) ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput {
	return o
}

func (o ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput) Elem() ComputeEnvironmentComputeResourcesLaunchTemplateOutput {
	return o.ApplyT(func (v *ComputeEnvironmentComputeResourcesLaunchTemplate) ComputeEnvironmentComputeResourcesLaunchTemplate { return *v }).(ComputeEnvironmentComputeResourcesLaunchTemplateOutput)
}

// ID of the launch template. You must specify either the launch template ID or launch template name in the request, but not both.
func (o ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput) LaunchTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ComputeEnvironmentComputeResourcesLaunchTemplate) *string { return v.LaunchTemplateId }).(pulumi.StringPtrOutput)
}

// Name of the launch template.
func (o ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput) LaunchTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ComputeEnvironmentComputeResourcesLaunchTemplate) *string { return v.LaunchTemplateName }).(pulumi.StringPtrOutput)
}

// The version number of the launch template. Default: The default version of the launch template.
func (o ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ComputeEnvironmentComputeResourcesLaunchTemplate) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ComputeEnvironmentComputeResourcesLaunchTemplateOutput{})
	pulumi.RegisterOutputType(ComputeEnvironmentComputeResourcesLaunchTemplatePtrOutput{})
}
