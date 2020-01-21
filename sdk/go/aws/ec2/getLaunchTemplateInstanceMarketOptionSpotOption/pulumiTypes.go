// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getLaunchTemplateInstanceMarketOptionSpotOption

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetLaunchTemplateInstanceMarketOptionSpotOption struct {
	BlockDurationMinutes int `pulumi:"blockDurationMinutes"`
	InstanceInterruptionBehavior string `pulumi:"instanceInterruptionBehavior"`
	MaxPrice string `pulumi:"maxPrice"`
	SpotInstanceType string `pulumi:"spotInstanceType"`
	ValidUntil string `pulumi:"validUntil"`
}

type GetLaunchTemplateInstanceMarketOptionSpotOptionInput interface {
	pulumi.Input

	ToGetLaunchTemplateInstanceMarketOptionSpotOptionOutput() GetLaunchTemplateInstanceMarketOptionSpotOptionOutput
	ToGetLaunchTemplateInstanceMarketOptionSpotOptionOutputWithContext(context.Context) GetLaunchTemplateInstanceMarketOptionSpotOptionOutput
}

type GetLaunchTemplateInstanceMarketOptionSpotOptionArgs struct {
	BlockDurationMinutes pulumi.IntInput `pulumi:"blockDurationMinutes"`
	InstanceInterruptionBehavior pulumi.StringInput `pulumi:"instanceInterruptionBehavior"`
	MaxPrice pulumi.StringInput `pulumi:"maxPrice"`
	SpotInstanceType pulumi.StringInput `pulumi:"spotInstanceType"`
	ValidUntil pulumi.StringInput `pulumi:"validUntil"`
}

func (GetLaunchTemplateInstanceMarketOptionSpotOptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLaunchTemplateInstanceMarketOptionSpotOption)(nil)).Elem()
}

func (i GetLaunchTemplateInstanceMarketOptionSpotOptionArgs) ToGetLaunchTemplateInstanceMarketOptionSpotOptionOutput() GetLaunchTemplateInstanceMarketOptionSpotOptionOutput {
	return i.ToGetLaunchTemplateInstanceMarketOptionSpotOptionOutputWithContext(context.Background())
}

func (i GetLaunchTemplateInstanceMarketOptionSpotOptionArgs) ToGetLaunchTemplateInstanceMarketOptionSpotOptionOutputWithContext(ctx context.Context) GetLaunchTemplateInstanceMarketOptionSpotOptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetLaunchTemplateInstanceMarketOptionSpotOptionOutput)
}

type GetLaunchTemplateInstanceMarketOptionSpotOptionArrayInput interface {
	pulumi.Input

	ToGetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput() GetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput
	ToGetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutputWithContext(context.Context) GetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput
}

type GetLaunchTemplateInstanceMarketOptionSpotOptionArray []GetLaunchTemplateInstanceMarketOptionSpotOptionInput

func (GetLaunchTemplateInstanceMarketOptionSpotOptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetLaunchTemplateInstanceMarketOptionSpotOption)(nil)).Elem()
}

func (i GetLaunchTemplateInstanceMarketOptionSpotOptionArray) ToGetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput() GetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput {
	return i.ToGetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutputWithContext(context.Background())
}

func (i GetLaunchTemplateInstanceMarketOptionSpotOptionArray) ToGetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutputWithContext(ctx context.Context) GetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput)
}

type GetLaunchTemplateInstanceMarketOptionSpotOptionOutput struct { *pulumi.OutputState }

func (GetLaunchTemplateInstanceMarketOptionSpotOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLaunchTemplateInstanceMarketOptionSpotOption)(nil)).Elem()
}

func (o GetLaunchTemplateInstanceMarketOptionSpotOptionOutput) ToGetLaunchTemplateInstanceMarketOptionSpotOptionOutput() GetLaunchTemplateInstanceMarketOptionSpotOptionOutput {
	return o
}

func (o GetLaunchTemplateInstanceMarketOptionSpotOptionOutput) ToGetLaunchTemplateInstanceMarketOptionSpotOptionOutputWithContext(ctx context.Context) GetLaunchTemplateInstanceMarketOptionSpotOptionOutput {
	return o
}

func (o GetLaunchTemplateInstanceMarketOptionSpotOptionOutput) BlockDurationMinutes() pulumi.IntOutput {
	return o.ApplyT(func (v GetLaunchTemplateInstanceMarketOptionSpotOption) int { return v.BlockDurationMinutes }).(pulumi.IntOutput)
}

func (o GetLaunchTemplateInstanceMarketOptionSpotOptionOutput) InstanceInterruptionBehavior() pulumi.StringOutput {
	return o.ApplyT(func (v GetLaunchTemplateInstanceMarketOptionSpotOption) string { return v.InstanceInterruptionBehavior }).(pulumi.StringOutput)
}

func (o GetLaunchTemplateInstanceMarketOptionSpotOptionOutput) MaxPrice() pulumi.StringOutput {
	return o.ApplyT(func (v GetLaunchTemplateInstanceMarketOptionSpotOption) string { return v.MaxPrice }).(pulumi.StringOutput)
}

func (o GetLaunchTemplateInstanceMarketOptionSpotOptionOutput) SpotInstanceType() pulumi.StringOutput {
	return o.ApplyT(func (v GetLaunchTemplateInstanceMarketOptionSpotOption) string { return v.SpotInstanceType }).(pulumi.StringOutput)
}

func (o GetLaunchTemplateInstanceMarketOptionSpotOptionOutput) ValidUntil() pulumi.StringOutput {
	return o.ApplyT(func (v GetLaunchTemplateInstanceMarketOptionSpotOption) string { return v.ValidUntil }).(pulumi.StringOutput)
}

type GetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput struct { *pulumi.OutputState}

func (GetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetLaunchTemplateInstanceMarketOptionSpotOption)(nil)).Elem()
}

func (o GetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput) ToGetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput() GetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput {
	return o
}

func (o GetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput) ToGetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutputWithContext(ctx context.Context) GetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput {
	return o
}

func (o GetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput) Index(i pulumi.IntInput) GetLaunchTemplateInstanceMarketOptionSpotOptionOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetLaunchTemplateInstanceMarketOptionSpotOption {
		return vs[0].([]GetLaunchTemplateInstanceMarketOptionSpotOption)[vs[1].(int)]
	}).(GetLaunchTemplateInstanceMarketOptionSpotOptionOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLaunchTemplateInstanceMarketOptionSpotOptionOutput{})
	pulumi.RegisterOutputType(GetLaunchTemplateInstanceMarketOptionSpotOptionArrayOutput{})
}
