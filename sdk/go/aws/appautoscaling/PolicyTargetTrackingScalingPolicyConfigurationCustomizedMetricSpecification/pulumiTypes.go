// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/appautoscaling/PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension"
)

type PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification struct {
	Dimensions []appautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension `pulumi:"dimensions"`
	MetricName string `pulumi:"metricName"`
	Namespace string `pulumi:"namespace"`
	Statistic string `pulumi:"statistic"`
	Unit *string `pulumi:"unit"`
}

type PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationInput interface {
	pulumi.Input

	ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput() PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput
	ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputWithContext(context.Context) PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput
}

type PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationArgs struct {
	Dimensions appautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionArrayInput `pulumi:"dimensions"`
	MetricName pulumi.StringInput `pulumi:"metricName"`
	Namespace pulumi.StringInput `pulumi:"namespace"`
	Statistic pulumi.StringInput `pulumi:"statistic"`
	Unit pulumi.StringPtrInput `pulumi:"unit"`
}

func (PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification)(nil)).Elem()
}

func (i PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationArgs) ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput() PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput {
	return i.ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputWithContext(context.Background())
}

func (i PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationArgs) ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputWithContext(ctx context.Context) PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput)
}

func (i PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationArgs) ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput() PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput {
	return i.ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutputWithContext(context.Background())
}

func (i PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationArgs) ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutputWithContext(ctx context.Context) PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput).ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutputWithContext(ctx)
}

type PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrInput interface {
	pulumi.Input

	ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput() PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput
	ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutputWithContext(context.Context) PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput
}

type policyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrType PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationArgs

func PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtr(v *PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationArgs) PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrInput {	return (*policyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrType)(v)
}

func (*policyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification)(nil)).Elem()
}

func (i *policyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrType) ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput() PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput {
	return i.ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutputWithContext(context.Background())
}

func (i *policyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrType) ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutputWithContext(ctx context.Context) PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput)
}

type PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput struct { *pulumi.OutputState }

func (PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification)(nil)).Elem()
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput) ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput() PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput {
	return o
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput) ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputWithContext(ctx context.Context) PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput {
	return o
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput) ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput() PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput {
	return o.ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutputWithContext(context.Background())
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput) ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutputWithContext(ctx context.Context) PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput {
	return o.ApplyT(func(v PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) *PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification {
		return &v
	}).(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput)
}
func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput) Dimensions() appautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionArrayOutput {
	return o.ApplyT(func (v PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) []appautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension { return v.Dimensions }).(appautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionArrayOutput)
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func (v PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func (v PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) string { return v.Namespace }).(pulumi.StringOutput)
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput) Statistic() pulumi.StringOutput {
	return o.ApplyT(func (v PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) string { return v.Statistic }).(pulumi.StringOutput)
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) *string { return v.Unit }).(pulumi.StringPtrOutput)
}

type PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput struct { *pulumi.OutputState}

func (PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification)(nil)).Elem()
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput) ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput() PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput {
	return o
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput) ToPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutputWithContext(ctx context.Context) PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput {
	return o
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput) Elem() PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput {
	return o.ApplyT(func (v *PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification { return *v }).(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput)
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput) Dimensions() appautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionArrayOutput {
	return o.ApplyT(func (v PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) []appautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension { return v.Dimensions }).(appautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension.PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionArrayOutput)
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func (v PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func (v PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) string { return v.Namespace }).(pulumi.StringOutput)
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput) Statistic() pulumi.StringOutput {
	return o.ApplyT(func (v PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) string { return v.Statistic }).(pulumi.StringOutput)
}

func (o PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func (v PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) *string { return v.Unit }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutput{})
	pulumi.RegisterOutputType(PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationPtrOutput{})
}
