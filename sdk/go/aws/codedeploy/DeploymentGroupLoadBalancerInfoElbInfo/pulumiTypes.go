// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package DeploymentGroupLoadBalancerInfoElbInfo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type DeploymentGroupLoadBalancerInfoElbInfo struct {
	// Name of the target group.
	Name *string `pulumi:"name"`
}

type DeploymentGroupLoadBalancerInfoElbInfoInput interface {
	pulumi.Input

	ToDeploymentGroupLoadBalancerInfoElbInfoOutput() DeploymentGroupLoadBalancerInfoElbInfoOutput
	ToDeploymentGroupLoadBalancerInfoElbInfoOutputWithContext(context.Context) DeploymentGroupLoadBalancerInfoElbInfoOutput
}

type DeploymentGroupLoadBalancerInfoElbInfoArgs struct {
	// Name of the target group.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DeploymentGroupLoadBalancerInfoElbInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentGroupLoadBalancerInfoElbInfo)(nil)).Elem()
}

func (i DeploymentGroupLoadBalancerInfoElbInfoArgs) ToDeploymentGroupLoadBalancerInfoElbInfoOutput() DeploymentGroupLoadBalancerInfoElbInfoOutput {
	return i.ToDeploymentGroupLoadBalancerInfoElbInfoOutputWithContext(context.Background())
}

func (i DeploymentGroupLoadBalancerInfoElbInfoArgs) ToDeploymentGroupLoadBalancerInfoElbInfoOutputWithContext(ctx context.Context) DeploymentGroupLoadBalancerInfoElbInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentGroupLoadBalancerInfoElbInfoOutput)
}

type DeploymentGroupLoadBalancerInfoElbInfoArrayInput interface {
	pulumi.Input

	ToDeploymentGroupLoadBalancerInfoElbInfoArrayOutput() DeploymentGroupLoadBalancerInfoElbInfoArrayOutput
	ToDeploymentGroupLoadBalancerInfoElbInfoArrayOutputWithContext(context.Context) DeploymentGroupLoadBalancerInfoElbInfoArrayOutput
}

type DeploymentGroupLoadBalancerInfoElbInfoArray []DeploymentGroupLoadBalancerInfoElbInfoInput

func (DeploymentGroupLoadBalancerInfoElbInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentGroupLoadBalancerInfoElbInfo)(nil)).Elem()
}

func (i DeploymentGroupLoadBalancerInfoElbInfoArray) ToDeploymentGroupLoadBalancerInfoElbInfoArrayOutput() DeploymentGroupLoadBalancerInfoElbInfoArrayOutput {
	return i.ToDeploymentGroupLoadBalancerInfoElbInfoArrayOutputWithContext(context.Background())
}

func (i DeploymentGroupLoadBalancerInfoElbInfoArray) ToDeploymentGroupLoadBalancerInfoElbInfoArrayOutputWithContext(ctx context.Context) DeploymentGroupLoadBalancerInfoElbInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentGroupLoadBalancerInfoElbInfoArrayOutput)
}

type DeploymentGroupLoadBalancerInfoElbInfoOutput struct { *pulumi.OutputState }

func (DeploymentGroupLoadBalancerInfoElbInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentGroupLoadBalancerInfoElbInfo)(nil)).Elem()
}

func (o DeploymentGroupLoadBalancerInfoElbInfoOutput) ToDeploymentGroupLoadBalancerInfoElbInfoOutput() DeploymentGroupLoadBalancerInfoElbInfoOutput {
	return o
}

func (o DeploymentGroupLoadBalancerInfoElbInfoOutput) ToDeploymentGroupLoadBalancerInfoElbInfoOutputWithContext(ctx context.Context) DeploymentGroupLoadBalancerInfoElbInfoOutput {
	return o
}

// Name of the target group.
func (o DeploymentGroupLoadBalancerInfoElbInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v DeploymentGroupLoadBalancerInfoElbInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DeploymentGroupLoadBalancerInfoElbInfoArrayOutput struct { *pulumi.OutputState}

func (DeploymentGroupLoadBalancerInfoElbInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentGroupLoadBalancerInfoElbInfo)(nil)).Elem()
}

func (o DeploymentGroupLoadBalancerInfoElbInfoArrayOutput) ToDeploymentGroupLoadBalancerInfoElbInfoArrayOutput() DeploymentGroupLoadBalancerInfoElbInfoArrayOutput {
	return o
}

func (o DeploymentGroupLoadBalancerInfoElbInfoArrayOutput) ToDeploymentGroupLoadBalancerInfoElbInfoArrayOutputWithContext(ctx context.Context) DeploymentGroupLoadBalancerInfoElbInfoArrayOutput {
	return o
}

func (o DeploymentGroupLoadBalancerInfoElbInfoArrayOutput) Index(i pulumi.IntInput) DeploymentGroupLoadBalancerInfoElbInfoOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) DeploymentGroupLoadBalancerInfoElbInfo {
		return vs[0].([]DeploymentGroupLoadBalancerInfoElbInfo)[vs[1].(int)]
	}).(DeploymentGroupLoadBalancerInfoElbInfoOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentGroupLoadBalancerInfoElbInfoOutput{})
	pulumi.RegisterOutputType(DeploymentGroupLoadBalancerInfoElbInfoArrayOutput{})
}
