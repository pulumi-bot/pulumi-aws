// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package AnalyticsApplicationInputsKinesisFirehose

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type AnalyticsApplicationInputsKinesisFirehose struct {
	ResourceArn string `pulumi:"resourceArn"`
	RoleArn string `pulumi:"roleArn"`
}

type AnalyticsApplicationInputsKinesisFirehoseInput interface {
	pulumi.Input

	ToAnalyticsApplicationInputsKinesisFirehoseOutput() AnalyticsApplicationInputsKinesisFirehoseOutput
	ToAnalyticsApplicationInputsKinesisFirehoseOutputWithContext(context.Context) AnalyticsApplicationInputsKinesisFirehoseOutput
}

type AnalyticsApplicationInputsKinesisFirehoseArgs struct {
	ResourceArn pulumi.StringInput `pulumi:"resourceArn"`
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
}

func (AnalyticsApplicationInputsKinesisFirehoseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsApplicationInputsKinesisFirehose)(nil)).Elem()
}

func (i AnalyticsApplicationInputsKinesisFirehoseArgs) ToAnalyticsApplicationInputsKinesisFirehoseOutput() AnalyticsApplicationInputsKinesisFirehoseOutput {
	return i.ToAnalyticsApplicationInputsKinesisFirehoseOutputWithContext(context.Background())
}

func (i AnalyticsApplicationInputsKinesisFirehoseArgs) ToAnalyticsApplicationInputsKinesisFirehoseOutputWithContext(ctx context.Context) AnalyticsApplicationInputsKinesisFirehoseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsApplicationInputsKinesisFirehoseOutput)
}

func (i AnalyticsApplicationInputsKinesisFirehoseArgs) ToAnalyticsApplicationInputsKinesisFirehosePtrOutput() AnalyticsApplicationInputsKinesisFirehosePtrOutput {
	return i.ToAnalyticsApplicationInputsKinesisFirehosePtrOutputWithContext(context.Background())
}

func (i AnalyticsApplicationInputsKinesisFirehoseArgs) ToAnalyticsApplicationInputsKinesisFirehosePtrOutputWithContext(ctx context.Context) AnalyticsApplicationInputsKinesisFirehosePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsApplicationInputsKinesisFirehoseOutput).ToAnalyticsApplicationInputsKinesisFirehosePtrOutputWithContext(ctx)
}

type AnalyticsApplicationInputsKinesisFirehosePtrInput interface {
	pulumi.Input

	ToAnalyticsApplicationInputsKinesisFirehosePtrOutput() AnalyticsApplicationInputsKinesisFirehosePtrOutput
	ToAnalyticsApplicationInputsKinesisFirehosePtrOutputWithContext(context.Context) AnalyticsApplicationInputsKinesisFirehosePtrOutput
}

type analyticsApplicationInputsKinesisFirehosePtrType AnalyticsApplicationInputsKinesisFirehoseArgs

func AnalyticsApplicationInputsKinesisFirehosePtr(v *AnalyticsApplicationInputsKinesisFirehoseArgs) AnalyticsApplicationInputsKinesisFirehosePtrInput {	return (*analyticsApplicationInputsKinesisFirehosePtrType)(v)
}

func (*analyticsApplicationInputsKinesisFirehosePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsApplicationInputsKinesisFirehose)(nil)).Elem()
}

func (i *analyticsApplicationInputsKinesisFirehosePtrType) ToAnalyticsApplicationInputsKinesisFirehosePtrOutput() AnalyticsApplicationInputsKinesisFirehosePtrOutput {
	return i.ToAnalyticsApplicationInputsKinesisFirehosePtrOutputWithContext(context.Background())
}

func (i *analyticsApplicationInputsKinesisFirehosePtrType) ToAnalyticsApplicationInputsKinesisFirehosePtrOutputWithContext(ctx context.Context) AnalyticsApplicationInputsKinesisFirehosePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsApplicationInputsKinesisFirehosePtrOutput)
}

type AnalyticsApplicationInputsKinesisFirehoseOutput struct { *pulumi.OutputState }

func (AnalyticsApplicationInputsKinesisFirehoseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsApplicationInputsKinesisFirehose)(nil)).Elem()
}

func (o AnalyticsApplicationInputsKinesisFirehoseOutput) ToAnalyticsApplicationInputsKinesisFirehoseOutput() AnalyticsApplicationInputsKinesisFirehoseOutput {
	return o
}

func (o AnalyticsApplicationInputsKinesisFirehoseOutput) ToAnalyticsApplicationInputsKinesisFirehoseOutputWithContext(ctx context.Context) AnalyticsApplicationInputsKinesisFirehoseOutput {
	return o
}

func (o AnalyticsApplicationInputsKinesisFirehoseOutput) ToAnalyticsApplicationInputsKinesisFirehosePtrOutput() AnalyticsApplicationInputsKinesisFirehosePtrOutput {
	return o.ToAnalyticsApplicationInputsKinesisFirehosePtrOutputWithContext(context.Background())
}

func (o AnalyticsApplicationInputsKinesisFirehoseOutput) ToAnalyticsApplicationInputsKinesisFirehosePtrOutputWithContext(ctx context.Context) AnalyticsApplicationInputsKinesisFirehosePtrOutput {
	return o.ApplyT(func(v AnalyticsApplicationInputsKinesisFirehose) *AnalyticsApplicationInputsKinesisFirehose {
		return &v
	}).(AnalyticsApplicationInputsKinesisFirehosePtrOutput)
}
func (o AnalyticsApplicationInputsKinesisFirehoseOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func (v AnalyticsApplicationInputsKinesisFirehose) string { return v.ResourceArn }).(pulumi.StringOutput)
}

func (o AnalyticsApplicationInputsKinesisFirehoseOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func (v AnalyticsApplicationInputsKinesisFirehose) string { return v.RoleArn }).(pulumi.StringOutput)
}

type AnalyticsApplicationInputsKinesisFirehosePtrOutput struct { *pulumi.OutputState}

func (AnalyticsApplicationInputsKinesisFirehosePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsApplicationInputsKinesisFirehose)(nil)).Elem()
}

func (o AnalyticsApplicationInputsKinesisFirehosePtrOutput) ToAnalyticsApplicationInputsKinesisFirehosePtrOutput() AnalyticsApplicationInputsKinesisFirehosePtrOutput {
	return o
}

func (o AnalyticsApplicationInputsKinesisFirehosePtrOutput) ToAnalyticsApplicationInputsKinesisFirehosePtrOutputWithContext(ctx context.Context) AnalyticsApplicationInputsKinesisFirehosePtrOutput {
	return o
}

func (o AnalyticsApplicationInputsKinesisFirehosePtrOutput) Elem() AnalyticsApplicationInputsKinesisFirehoseOutput {
	return o.ApplyT(func (v *AnalyticsApplicationInputsKinesisFirehose) AnalyticsApplicationInputsKinesisFirehose { return *v }).(AnalyticsApplicationInputsKinesisFirehoseOutput)
}

func (o AnalyticsApplicationInputsKinesisFirehosePtrOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func (v AnalyticsApplicationInputsKinesisFirehose) string { return v.ResourceArn }).(pulumi.StringOutput)
}

func (o AnalyticsApplicationInputsKinesisFirehosePtrOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func (v AnalyticsApplicationInputsKinesisFirehose) string { return v.RoleArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AnalyticsApplicationInputsKinesisFirehoseOutput{})
	pulumi.RegisterOutputType(AnalyticsApplicationInputsKinesisFirehosePtrOutput{})
}
