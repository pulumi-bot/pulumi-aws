// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ProjectVpcConfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ProjectVpcConfig struct {
	// The security group IDs to assign to running builds.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The subnet IDs within which to run builds.
	Subnets []string `pulumi:"subnets"`
	// The ID of the VPC within which to run builds.
	VpcId string `pulumi:"vpcId"`
}

type ProjectVpcConfigInput interface {
	pulumi.Input

	ToProjectVpcConfigOutput() ProjectVpcConfigOutput
	ToProjectVpcConfigOutputWithContext(context.Context) ProjectVpcConfigOutput
}

type ProjectVpcConfigArgs struct {
	// The security group IDs to assign to running builds.
	SecurityGroupIds pulumi.StringArrayInput `pulumi:"securityGroupIds"`
	// The subnet IDs within which to run builds.
	Subnets pulumi.StringArrayInput `pulumi:"subnets"`
	// The ID of the VPC within which to run builds.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

func (ProjectVpcConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectVpcConfig)(nil)).Elem()
}

func (i ProjectVpcConfigArgs) ToProjectVpcConfigOutput() ProjectVpcConfigOutput {
	return i.ToProjectVpcConfigOutputWithContext(context.Background())
}

func (i ProjectVpcConfigArgs) ToProjectVpcConfigOutputWithContext(ctx context.Context) ProjectVpcConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectVpcConfigOutput)
}

func (i ProjectVpcConfigArgs) ToProjectVpcConfigPtrOutput() ProjectVpcConfigPtrOutput {
	return i.ToProjectVpcConfigPtrOutputWithContext(context.Background())
}

func (i ProjectVpcConfigArgs) ToProjectVpcConfigPtrOutputWithContext(ctx context.Context) ProjectVpcConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectVpcConfigOutput).ToProjectVpcConfigPtrOutputWithContext(ctx)
}

type ProjectVpcConfigPtrInput interface {
	pulumi.Input

	ToProjectVpcConfigPtrOutput() ProjectVpcConfigPtrOutput
	ToProjectVpcConfigPtrOutputWithContext(context.Context) ProjectVpcConfigPtrOutput
}

type projectVpcConfigPtrType ProjectVpcConfigArgs

func ProjectVpcConfigPtr(v *ProjectVpcConfigArgs) ProjectVpcConfigPtrInput {	return (*projectVpcConfigPtrType)(v)
}

func (*projectVpcConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectVpcConfig)(nil)).Elem()
}

func (i *projectVpcConfigPtrType) ToProjectVpcConfigPtrOutput() ProjectVpcConfigPtrOutput {
	return i.ToProjectVpcConfigPtrOutputWithContext(context.Background())
}

func (i *projectVpcConfigPtrType) ToProjectVpcConfigPtrOutputWithContext(ctx context.Context) ProjectVpcConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectVpcConfigPtrOutput)
}

type ProjectVpcConfigOutput struct { *pulumi.OutputState }

func (ProjectVpcConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectVpcConfig)(nil)).Elem()
}

func (o ProjectVpcConfigOutput) ToProjectVpcConfigOutput() ProjectVpcConfigOutput {
	return o
}

func (o ProjectVpcConfigOutput) ToProjectVpcConfigOutputWithContext(ctx context.Context) ProjectVpcConfigOutput {
	return o
}

func (o ProjectVpcConfigOutput) ToProjectVpcConfigPtrOutput() ProjectVpcConfigPtrOutput {
	return o.ToProjectVpcConfigPtrOutputWithContext(context.Background())
}

func (o ProjectVpcConfigOutput) ToProjectVpcConfigPtrOutputWithContext(ctx context.Context) ProjectVpcConfigPtrOutput {
	return o.ApplyT(func(v ProjectVpcConfig) *ProjectVpcConfig {
		return &v
	}).(ProjectVpcConfigPtrOutput)
}
// The security group IDs to assign to running builds.
func (o ProjectVpcConfigOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ProjectVpcConfig) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The subnet IDs within which to run builds.
func (o ProjectVpcConfigOutput) Subnets() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ProjectVpcConfig) []string { return v.Subnets }).(pulumi.StringArrayOutput)
}

// The ID of the VPC within which to run builds.
func (o ProjectVpcConfigOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func (v ProjectVpcConfig) string { return v.VpcId }).(pulumi.StringOutput)
}

type ProjectVpcConfigPtrOutput struct { *pulumi.OutputState}

func (ProjectVpcConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectVpcConfig)(nil)).Elem()
}

func (o ProjectVpcConfigPtrOutput) ToProjectVpcConfigPtrOutput() ProjectVpcConfigPtrOutput {
	return o
}

func (o ProjectVpcConfigPtrOutput) ToProjectVpcConfigPtrOutputWithContext(ctx context.Context) ProjectVpcConfigPtrOutput {
	return o
}

func (o ProjectVpcConfigPtrOutput) Elem() ProjectVpcConfigOutput {
	return o.ApplyT(func (v *ProjectVpcConfig) ProjectVpcConfig { return *v }).(ProjectVpcConfigOutput)
}

// The security group IDs to assign to running builds.
func (o ProjectVpcConfigPtrOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ProjectVpcConfig) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The subnet IDs within which to run builds.
func (o ProjectVpcConfigPtrOutput) Subnets() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ProjectVpcConfig) []string { return v.Subnets }).(pulumi.StringArrayOutput)
}

// The ID of the VPC within which to run builds.
func (o ProjectVpcConfigPtrOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func (v ProjectVpcConfig) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProjectVpcConfigOutput{})
	pulumi.RegisterOutputType(ProjectVpcConfigPtrOutput{})
}
