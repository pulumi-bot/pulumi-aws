// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ParameterGroupParameter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ParameterGroupParameter struct {
	// The name of the ElastiCache parameter.
	Name string `pulumi:"name"`
	// The value of the ElastiCache parameter.
	Value string `pulumi:"value"`
}

type ParameterGroupParameterInput interface {
	pulumi.Input

	ToParameterGroupParameterOutput() ParameterGroupParameterOutput
	ToParameterGroupParameterOutputWithContext(context.Context) ParameterGroupParameterOutput
}

type ParameterGroupParameterArgs struct {
	// The name of the ElastiCache parameter.
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the ElastiCache parameter.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ParameterGroupParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameter)(nil)).Elem()
}

func (i ParameterGroupParameterArgs) ToParameterGroupParameterOutput() ParameterGroupParameterOutput {
	return i.ToParameterGroupParameterOutputWithContext(context.Background())
}

func (i ParameterGroupParameterArgs) ToParameterGroupParameterOutputWithContext(ctx context.Context) ParameterGroupParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterOutput)
}

type ParameterGroupParameterArrayInput interface {
	pulumi.Input

	ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput
	ToParameterGroupParameterArrayOutputWithContext(context.Context) ParameterGroupParameterArrayOutput
}

type ParameterGroupParameterArray []ParameterGroupParameterInput

func (ParameterGroupParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameter)(nil)).Elem()
}

func (i ParameterGroupParameterArray) ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput {
	return i.ToParameterGroupParameterArrayOutputWithContext(context.Background())
}

func (i ParameterGroupParameterArray) ToParameterGroupParameterArrayOutputWithContext(ctx context.Context) ParameterGroupParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterArrayOutput)
}

type ParameterGroupParameterOutput struct { *pulumi.OutputState }

func (ParameterGroupParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameter)(nil)).Elem()
}

func (o ParameterGroupParameterOutput) ToParameterGroupParameterOutput() ParameterGroupParameterOutput {
	return o
}

func (o ParameterGroupParameterOutput) ToParameterGroupParameterOutputWithContext(ctx context.Context) ParameterGroupParameterOutput {
	return o
}

// The name of the ElastiCache parameter.
func (o ParameterGroupParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v ParameterGroupParameter) string { return v.Name }).(pulumi.StringOutput)
}

// The value of the ElastiCache parameter.
func (o ParameterGroupParameterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func (v ParameterGroupParameter) string { return v.Value }).(pulumi.StringOutput)
}

type ParameterGroupParameterArrayOutput struct { *pulumi.OutputState}

func (ParameterGroupParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameter)(nil)).Elem()
}

func (o ParameterGroupParameterArrayOutput) ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput {
	return o
}

func (o ParameterGroupParameterArrayOutput) ToParameterGroupParameterArrayOutputWithContext(ctx context.Context) ParameterGroupParameterArrayOutput {
	return o
}

func (o ParameterGroupParameterArrayOutput) Index(i pulumi.IntInput) ParameterGroupParameterOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ParameterGroupParameter {
		return vs[0].([]ParameterGroupParameter)[vs[1].(int)]
	}).(ParameterGroupParameterOutput)
}

func init() {
	pulumi.RegisterOutputType(ParameterGroupParameterOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterArrayOutput{})
}
