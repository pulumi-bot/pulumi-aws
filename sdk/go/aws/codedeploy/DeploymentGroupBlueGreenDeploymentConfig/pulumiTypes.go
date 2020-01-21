// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package DeploymentGroupBlueGreenDeploymentConfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/codedeploy/DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption"
	"https:/github.com/pulumi/pulumi-aws/codedeploy/DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption"
	"https:/github.com/pulumi/pulumi-aws/codedeploy/DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess"
)

type DeploymentGroupBlueGreenDeploymentConfig struct {
	// Information about the action to take when newly provisioned instances are ready to receive traffic in a blue/green deployment (documented below).
	DeploymentReadyOption *codedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption `pulumi:"deploymentReadyOption"`
	// Information about how instances are provisioned for a replacement environment in a blue/green deployment (documented below).
	GreenFleetProvisioningOption *codedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption `pulumi:"greenFleetProvisioningOption"`
	// Information about whether to terminate instances in the original fleet during a blue/green deployment (documented below).
	TerminateBlueInstancesOnDeploymentSuccess *codedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess `pulumi:"terminateBlueInstancesOnDeploymentSuccess"`
}

type DeploymentGroupBlueGreenDeploymentConfigInput interface {
	pulumi.Input

	ToDeploymentGroupBlueGreenDeploymentConfigOutput() DeploymentGroupBlueGreenDeploymentConfigOutput
	ToDeploymentGroupBlueGreenDeploymentConfigOutputWithContext(context.Context) DeploymentGroupBlueGreenDeploymentConfigOutput
}

type DeploymentGroupBlueGreenDeploymentConfigArgs struct {
	// Information about the action to take when newly provisioned instances are ready to receive traffic in a blue/green deployment (documented below).
	DeploymentReadyOption codedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionPtrInput `pulumi:"deploymentReadyOption"`
	// Information about how instances are provisioned for a replacement environment in a blue/green deployment (documented below).
	GreenFleetProvisioningOption codedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionPtrInput `pulumi:"greenFleetProvisioningOption"`
	// Information about whether to terminate instances in the original fleet during a blue/green deployment (documented below).
	TerminateBlueInstancesOnDeploymentSuccess codedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessPtrInput `pulumi:"terminateBlueInstancesOnDeploymentSuccess"`
}

func (DeploymentGroupBlueGreenDeploymentConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentGroupBlueGreenDeploymentConfig)(nil)).Elem()
}

func (i DeploymentGroupBlueGreenDeploymentConfigArgs) ToDeploymentGroupBlueGreenDeploymentConfigOutput() DeploymentGroupBlueGreenDeploymentConfigOutput {
	return i.ToDeploymentGroupBlueGreenDeploymentConfigOutputWithContext(context.Background())
}

func (i DeploymentGroupBlueGreenDeploymentConfigArgs) ToDeploymentGroupBlueGreenDeploymentConfigOutputWithContext(ctx context.Context) DeploymentGroupBlueGreenDeploymentConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentGroupBlueGreenDeploymentConfigOutput)
}

func (i DeploymentGroupBlueGreenDeploymentConfigArgs) ToDeploymentGroupBlueGreenDeploymentConfigPtrOutput() DeploymentGroupBlueGreenDeploymentConfigPtrOutput {
	return i.ToDeploymentGroupBlueGreenDeploymentConfigPtrOutputWithContext(context.Background())
}

func (i DeploymentGroupBlueGreenDeploymentConfigArgs) ToDeploymentGroupBlueGreenDeploymentConfigPtrOutputWithContext(ctx context.Context) DeploymentGroupBlueGreenDeploymentConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentGroupBlueGreenDeploymentConfigOutput).ToDeploymentGroupBlueGreenDeploymentConfigPtrOutputWithContext(ctx)
}

type DeploymentGroupBlueGreenDeploymentConfigPtrInput interface {
	pulumi.Input

	ToDeploymentGroupBlueGreenDeploymentConfigPtrOutput() DeploymentGroupBlueGreenDeploymentConfigPtrOutput
	ToDeploymentGroupBlueGreenDeploymentConfigPtrOutputWithContext(context.Context) DeploymentGroupBlueGreenDeploymentConfigPtrOutput
}

type deploymentGroupBlueGreenDeploymentConfigPtrType DeploymentGroupBlueGreenDeploymentConfigArgs

func DeploymentGroupBlueGreenDeploymentConfigPtr(v *DeploymentGroupBlueGreenDeploymentConfigArgs) DeploymentGroupBlueGreenDeploymentConfigPtrInput {	return (*deploymentGroupBlueGreenDeploymentConfigPtrType)(v)
}

func (*deploymentGroupBlueGreenDeploymentConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentGroupBlueGreenDeploymentConfig)(nil)).Elem()
}

func (i *deploymentGroupBlueGreenDeploymentConfigPtrType) ToDeploymentGroupBlueGreenDeploymentConfigPtrOutput() DeploymentGroupBlueGreenDeploymentConfigPtrOutput {
	return i.ToDeploymentGroupBlueGreenDeploymentConfigPtrOutputWithContext(context.Background())
}

func (i *deploymentGroupBlueGreenDeploymentConfigPtrType) ToDeploymentGroupBlueGreenDeploymentConfigPtrOutputWithContext(ctx context.Context) DeploymentGroupBlueGreenDeploymentConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentGroupBlueGreenDeploymentConfigPtrOutput)
}

type DeploymentGroupBlueGreenDeploymentConfigOutput struct { *pulumi.OutputState }

func (DeploymentGroupBlueGreenDeploymentConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentGroupBlueGreenDeploymentConfig)(nil)).Elem()
}

func (o DeploymentGroupBlueGreenDeploymentConfigOutput) ToDeploymentGroupBlueGreenDeploymentConfigOutput() DeploymentGroupBlueGreenDeploymentConfigOutput {
	return o
}

func (o DeploymentGroupBlueGreenDeploymentConfigOutput) ToDeploymentGroupBlueGreenDeploymentConfigOutputWithContext(ctx context.Context) DeploymentGroupBlueGreenDeploymentConfigOutput {
	return o
}

func (o DeploymentGroupBlueGreenDeploymentConfigOutput) ToDeploymentGroupBlueGreenDeploymentConfigPtrOutput() DeploymentGroupBlueGreenDeploymentConfigPtrOutput {
	return o.ToDeploymentGroupBlueGreenDeploymentConfigPtrOutputWithContext(context.Background())
}

func (o DeploymentGroupBlueGreenDeploymentConfigOutput) ToDeploymentGroupBlueGreenDeploymentConfigPtrOutputWithContext(ctx context.Context) DeploymentGroupBlueGreenDeploymentConfigPtrOutput {
	return o.ApplyT(func(v DeploymentGroupBlueGreenDeploymentConfig) *DeploymentGroupBlueGreenDeploymentConfig {
		return &v
	}).(DeploymentGroupBlueGreenDeploymentConfigPtrOutput)
}
// Information about the action to take when newly provisioned instances are ready to receive traffic in a blue/green deployment (documented below).
func (o DeploymentGroupBlueGreenDeploymentConfigOutput) DeploymentReadyOption() codedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionPtrOutput {
	return o.ApplyT(func (v DeploymentGroupBlueGreenDeploymentConfig) *codedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption { return v.DeploymentReadyOption }).(codedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionPtrOutput)
}

// Information about how instances are provisioned for a replacement environment in a blue/green deployment (documented below).
func (o DeploymentGroupBlueGreenDeploymentConfigOutput) GreenFleetProvisioningOption() codedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionPtrOutput {
	return o.ApplyT(func (v DeploymentGroupBlueGreenDeploymentConfig) *codedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption { return v.GreenFleetProvisioningOption }).(codedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionPtrOutput)
}

// Information about whether to terminate instances in the original fleet during a blue/green deployment (documented below).
func (o DeploymentGroupBlueGreenDeploymentConfigOutput) TerminateBlueInstancesOnDeploymentSuccess() codedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessPtrOutput {
	return o.ApplyT(func (v DeploymentGroupBlueGreenDeploymentConfig) *codedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess { return v.TerminateBlueInstancesOnDeploymentSuccess }).(codedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessPtrOutput)
}

type DeploymentGroupBlueGreenDeploymentConfigPtrOutput struct { *pulumi.OutputState}

func (DeploymentGroupBlueGreenDeploymentConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentGroupBlueGreenDeploymentConfig)(nil)).Elem()
}

func (o DeploymentGroupBlueGreenDeploymentConfigPtrOutput) ToDeploymentGroupBlueGreenDeploymentConfigPtrOutput() DeploymentGroupBlueGreenDeploymentConfigPtrOutput {
	return o
}

func (o DeploymentGroupBlueGreenDeploymentConfigPtrOutput) ToDeploymentGroupBlueGreenDeploymentConfigPtrOutputWithContext(ctx context.Context) DeploymentGroupBlueGreenDeploymentConfigPtrOutput {
	return o
}

func (o DeploymentGroupBlueGreenDeploymentConfigPtrOutput) Elem() DeploymentGroupBlueGreenDeploymentConfigOutput {
	return o.ApplyT(func (v *DeploymentGroupBlueGreenDeploymentConfig) DeploymentGroupBlueGreenDeploymentConfig { return *v }).(DeploymentGroupBlueGreenDeploymentConfigOutput)
}

// Information about the action to take when newly provisioned instances are ready to receive traffic in a blue/green deployment (documented below).
func (o DeploymentGroupBlueGreenDeploymentConfigPtrOutput) DeploymentReadyOption() codedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionPtrOutput {
	return o.ApplyT(func (v DeploymentGroupBlueGreenDeploymentConfig) *codedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption { return v.DeploymentReadyOption }).(codedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionPtrOutput)
}

// Information about how instances are provisioned for a replacement environment in a blue/green deployment (documented below).
func (o DeploymentGroupBlueGreenDeploymentConfigPtrOutput) GreenFleetProvisioningOption() codedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionPtrOutput {
	return o.ApplyT(func (v DeploymentGroupBlueGreenDeploymentConfig) *codedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption { return v.GreenFleetProvisioningOption }).(codedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionPtrOutput)
}

// Information about whether to terminate instances in the original fleet during a blue/green deployment (documented below).
func (o DeploymentGroupBlueGreenDeploymentConfigPtrOutput) TerminateBlueInstancesOnDeploymentSuccess() codedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessPtrOutput {
	return o.ApplyT(func (v DeploymentGroupBlueGreenDeploymentConfig) *codedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess { return v.TerminateBlueInstancesOnDeploymentSuccess }).(codedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentGroupBlueGreenDeploymentConfigOutput{})
	pulumi.RegisterOutputType(DeploymentGroupBlueGreenDeploymentConfigPtrOutput{})
}
