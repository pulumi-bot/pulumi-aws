// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package TrailEventSelectorDataResource

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type TrailEventSelectorDataResource struct {
	// The resource type in which you want to log data events. You can specify only the follwing value: "AWS::S3::Object", "AWS::Lambda::Function"
	Type string `pulumi:"type"`
	// A list of ARN for the specified S3 buckets and object prefixes..
	Values []string `pulumi:"values"`
}

type TrailEventSelectorDataResourceInput interface {
	pulumi.Input

	ToTrailEventSelectorDataResourceOutput() TrailEventSelectorDataResourceOutput
	ToTrailEventSelectorDataResourceOutputWithContext(context.Context) TrailEventSelectorDataResourceOutput
}

type TrailEventSelectorDataResourceArgs struct {
	// The resource type in which you want to log data events. You can specify only the follwing value: "AWS::S3::Object", "AWS::Lambda::Function"
	Type pulumi.StringInput `pulumi:"type"`
	// A list of ARN for the specified S3 buckets and object prefixes..
	Values pulumi.StringArrayInput `pulumi:"values"`
}

func (TrailEventSelectorDataResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailEventSelectorDataResource)(nil)).Elem()
}

func (i TrailEventSelectorDataResourceArgs) ToTrailEventSelectorDataResourceOutput() TrailEventSelectorDataResourceOutput {
	return i.ToTrailEventSelectorDataResourceOutputWithContext(context.Background())
}

func (i TrailEventSelectorDataResourceArgs) ToTrailEventSelectorDataResourceOutputWithContext(ctx context.Context) TrailEventSelectorDataResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailEventSelectorDataResourceOutput)
}

type TrailEventSelectorDataResourceArrayInput interface {
	pulumi.Input

	ToTrailEventSelectorDataResourceArrayOutput() TrailEventSelectorDataResourceArrayOutput
	ToTrailEventSelectorDataResourceArrayOutputWithContext(context.Context) TrailEventSelectorDataResourceArrayOutput
}

type TrailEventSelectorDataResourceArray []TrailEventSelectorDataResourceInput

func (TrailEventSelectorDataResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrailEventSelectorDataResource)(nil)).Elem()
}

func (i TrailEventSelectorDataResourceArray) ToTrailEventSelectorDataResourceArrayOutput() TrailEventSelectorDataResourceArrayOutput {
	return i.ToTrailEventSelectorDataResourceArrayOutputWithContext(context.Background())
}

func (i TrailEventSelectorDataResourceArray) ToTrailEventSelectorDataResourceArrayOutputWithContext(ctx context.Context) TrailEventSelectorDataResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailEventSelectorDataResourceArrayOutput)
}

type TrailEventSelectorDataResourceOutput struct { *pulumi.OutputState }

func (TrailEventSelectorDataResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailEventSelectorDataResource)(nil)).Elem()
}

func (o TrailEventSelectorDataResourceOutput) ToTrailEventSelectorDataResourceOutput() TrailEventSelectorDataResourceOutput {
	return o
}

func (o TrailEventSelectorDataResourceOutput) ToTrailEventSelectorDataResourceOutputWithContext(ctx context.Context) TrailEventSelectorDataResourceOutput {
	return o
}

// The resource type in which you want to log data events. You can specify only the follwing value: "AWS::S3::Object", "AWS::Lambda::Function"
func (o TrailEventSelectorDataResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func (v TrailEventSelectorDataResource) string { return v.Type }).(pulumi.StringOutput)
}

// A list of ARN for the specified S3 buckets and object prefixes..
func (o TrailEventSelectorDataResourceOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func (v TrailEventSelectorDataResource) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type TrailEventSelectorDataResourceArrayOutput struct { *pulumi.OutputState}

func (TrailEventSelectorDataResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrailEventSelectorDataResource)(nil)).Elem()
}

func (o TrailEventSelectorDataResourceArrayOutput) ToTrailEventSelectorDataResourceArrayOutput() TrailEventSelectorDataResourceArrayOutput {
	return o
}

func (o TrailEventSelectorDataResourceArrayOutput) ToTrailEventSelectorDataResourceArrayOutputWithContext(ctx context.Context) TrailEventSelectorDataResourceArrayOutput {
	return o
}

func (o TrailEventSelectorDataResourceArrayOutput) Index(i pulumi.IntInput) TrailEventSelectorDataResourceOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) TrailEventSelectorDataResource {
		return vs[0].([]TrailEventSelectorDataResource)[vs[1].(int)]
	}).(TrailEventSelectorDataResourceOutput)
}

func init() {
	pulumi.RegisterOutputType(TrailEventSelectorDataResourceOutput{})
	pulumi.RegisterOutputType(TrailEventSelectorDataResourceArrayOutput{})
}
