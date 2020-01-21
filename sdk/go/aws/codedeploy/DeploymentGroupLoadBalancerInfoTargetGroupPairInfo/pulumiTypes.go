// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package DeploymentGroupLoadBalancerInfoTargetGroupPairInfo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/codedeploy/DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute"
	"https:/github.com/pulumi/pulumi-aws/codedeploy/DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup"
	"https:/github.com/pulumi/pulumi-aws/codedeploy/DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute"
)

type DeploymentGroupLoadBalancerInfoTargetGroupPairInfo struct {
	// Configuration block for the production traffic route (documented below).
	ProdTrafficRoute codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute `pulumi:"prodTrafficRoute"`
	// Configuration blocks for a target group within a target group pair (documented below).
	TargetGroups []codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup `pulumi:"targetGroups"`
	// Configuration block for the test traffic route (documented below).
	TestTrafficRoute *codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute `pulumi:"testTrafficRoute"`
}

type DeploymentGroupLoadBalancerInfoTargetGroupPairInfoInput interface {
	pulumi.Input

	ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput() DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput
	ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputWithContext(context.Context) DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput
}

type DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs struct {
	// Configuration block for the production traffic route (documented below).
	ProdTrafficRoute codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteInput `pulumi:"prodTrafficRoute"`
	// Configuration blocks for a target group within a target group pair (documented below).
	TargetGroups codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArrayInput `pulumi:"targetGroups"`
	// Configuration block for the test traffic route (documented below).
	TestTrafficRoute codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoutePtrInput `pulumi:"testTrafficRoute"`
}

func (DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentGroupLoadBalancerInfoTargetGroupPairInfo)(nil)).Elem()
}

func (i DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs) ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput() DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput {
	return i.ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputWithContext(context.Background())
}

func (i DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs) ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputWithContext(ctx context.Context) DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput)
}

func (i DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs) ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput() DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput {
	return i.ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutputWithContext(context.Background())
}

func (i DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs) ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutputWithContext(ctx context.Context) DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput).ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutputWithContext(ctx)
}

type DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrInput interface {
	pulumi.Input

	ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput() DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput
	ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutputWithContext(context.Context) DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput
}

type deploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrType DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs

func DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtr(v *DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs) DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrInput {	return (*deploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrType)(v)
}

func (*deploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentGroupLoadBalancerInfoTargetGroupPairInfo)(nil)).Elem()
}

func (i *deploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrType) ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput() DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput {
	return i.ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutputWithContext(context.Background())
}

func (i *deploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrType) ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutputWithContext(ctx context.Context) DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput)
}

type DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput struct { *pulumi.OutputState }

func (DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentGroupLoadBalancerInfoTargetGroupPairInfo)(nil)).Elem()
}

func (o DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput) ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput() DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput {
	return o
}

func (o DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput) ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputWithContext(ctx context.Context) DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput {
	return o
}

func (o DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput) ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput() DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput {
	return o.ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutputWithContext(context.Background())
}

func (o DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput) ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutputWithContext(ctx context.Context) DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput {
	return o.ApplyT(func(v DeploymentGroupLoadBalancerInfoTargetGroupPairInfo) *DeploymentGroupLoadBalancerInfoTargetGroupPairInfo {
		return &v
	}).(DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput)
}
// Configuration block for the production traffic route (documented below).
func (o DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput) ProdTrafficRoute() codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutput {
	return o.ApplyT(func (v DeploymentGroupLoadBalancerInfoTargetGroupPairInfo) codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute { return v.ProdTrafficRoute }).(codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutput)
}

// Configuration blocks for a target group within a target group pair (documented below).
func (o DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput) TargetGroups() codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArrayOutput {
	return o.ApplyT(func (v DeploymentGroupLoadBalancerInfoTargetGroupPairInfo) []codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup { return v.TargetGroups }).(codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArrayOutput)
}

// Configuration block for the test traffic route (documented below).
func (o DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput) TestTrafficRoute() codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoutePtrOutput {
	return o.ApplyT(func (v DeploymentGroupLoadBalancerInfoTargetGroupPairInfo) *codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute { return v.TestTrafficRoute }).(codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoutePtrOutput)
}

type DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput struct { *pulumi.OutputState}

func (DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentGroupLoadBalancerInfoTargetGroupPairInfo)(nil)).Elem()
}

func (o DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput) ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput() DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput {
	return o
}

func (o DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput) ToDeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutputWithContext(ctx context.Context) DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput {
	return o
}

func (o DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput) Elem() DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput {
	return o.ApplyT(func (v *DeploymentGroupLoadBalancerInfoTargetGroupPairInfo) DeploymentGroupLoadBalancerInfoTargetGroupPairInfo { return *v }).(DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput)
}

// Configuration block for the production traffic route (documented below).
func (o DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput) ProdTrafficRoute() codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutput {
	return o.ApplyT(func (v DeploymentGroupLoadBalancerInfoTargetGroupPairInfo) codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute { return v.ProdTrafficRoute }).(codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutput)
}

// Configuration blocks for a target group within a target group pair (documented below).
func (o DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput) TargetGroups() codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArrayOutput {
	return o.ApplyT(func (v DeploymentGroupLoadBalancerInfoTargetGroupPairInfo) []codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup { return v.TargetGroups }).(codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArrayOutput)
}

// Configuration block for the test traffic route (documented below).
func (o DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput) TestTrafficRoute() codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoutePtrOutput {
	return o.ApplyT(func (v DeploymentGroupLoadBalancerInfoTargetGroupPairInfo) *codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute { return v.TestTrafficRoute }).(codedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoutePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutput{})
	pulumi.RegisterOutputType(DeploymentGroupLoadBalancerInfoTargetGroupPairInfoPtrOutput{})
}
