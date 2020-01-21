// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package JobDefinitionTimeout

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type JobDefinitionTimeout struct {
	AttemptDurationSeconds *int `pulumi:"attemptDurationSeconds"`
}

type JobDefinitionTimeoutInput interface {
	pulumi.Input

	ToJobDefinitionTimeoutOutput() JobDefinitionTimeoutOutput
	ToJobDefinitionTimeoutOutputWithContext(context.Context) JobDefinitionTimeoutOutput
}

type JobDefinitionTimeoutArgs struct {
	AttemptDurationSeconds pulumi.IntPtrInput `pulumi:"attemptDurationSeconds"`
}

func (JobDefinitionTimeoutArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDefinitionTimeout)(nil)).Elem()
}

func (i JobDefinitionTimeoutArgs) ToJobDefinitionTimeoutOutput() JobDefinitionTimeoutOutput {
	return i.ToJobDefinitionTimeoutOutputWithContext(context.Background())
}

func (i JobDefinitionTimeoutArgs) ToJobDefinitionTimeoutOutputWithContext(ctx context.Context) JobDefinitionTimeoutOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionTimeoutOutput)
}

func (i JobDefinitionTimeoutArgs) ToJobDefinitionTimeoutPtrOutput() JobDefinitionTimeoutPtrOutput {
	return i.ToJobDefinitionTimeoutPtrOutputWithContext(context.Background())
}

func (i JobDefinitionTimeoutArgs) ToJobDefinitionTimeoutPtrOutputWithContext(ctx context.Context) JobDefinitionTimeoutPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionTimeoutOutput).ToJobDefinitionTimeoutPtrOutputWithContext(ctx)
}

type JobDefinitionTimeoutPtrInput interface {
	pulumi.Input

	ToJobDefinitionTimeoutPtrOutput() JobDefinitionTimeoutPtrOutput
	ToJobDefinitionTimeoutPtrOutputWithContext(context.Context) JobDefinitionTimeoutPtrOutput
}

type jobDefinitionTimeoutPtrType JobDefinitionTimeoutArgs

func JobDefinitionTimeoutPtr(v *JobDefinitionTimeoutArgs) JobDefinitionTimeoutPtrInput {	return (*jobDefinitionTimeoutPtrType)(v)
}

func (*jobDefinitionTimeoutPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinitionTimeout)(nil)).Elem()
}

func (i *jobDefinitionTimeoutPtrType) ToJobDefinitionTimeoutPtrOutput() JobDefinitionTimeoutPtrOutput {
	return i.ToJobDefinitionTimeoutPtrOutputWithContext(context.Background())
}

func (i *jobDefinitionTimeoutPtrType) ToJobDefinitionTimeoutPtrOutputWithContext(ctx context.Context) JobDefinitionTimeoutPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionTimeoutPtrOutput)
}

type JobDefinitionTimeoutOutput struct { *pulumi.OutputState }

func (JobDefinitionTimeoutOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDefinitionTimeout)(nil)).Elem()
}

func (o JobDefinitionTimeoutOutput) ToJobDefinitionTimeoutOutput() JobDefinitionTimeoutOutput {
	return o
}

func (o JobDefinitionTimeoutOutput) ToJobDefinitionTimeoutOutputWithContext(ctx context.Context) JobDefinitionTimeoutOutput {
	return o
}

func (o JobDefinitionTimeoutOutput) ToJobDefinitionTimeoutPtrOutput() JobDefinitionTimeoutPtrOutput {
	return o.ToJobDefinitionTimeoutPtrOutputWithContext(context.Background())
}

func (o JobDefinitionTimeoutOutput) ToJobDefinitionTimeoutPtrOutputWithContext(ctx context.Context) JobDefinitionTimeoutPtrOutput {
	return o.ApplyT(func(v JobDefinitionTimeout) *JobDefinitionTimeout {
		return &v
	}).(JobDefinitionTimeoutPtrOutput)
}
func (o JobDefinitionTimeoutOutput) AttemptDurationSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobDefinitionTimeout) *int { return v.AttemptDurationSeconds }).(pulumi.IntPtrOutput)
}

type JobDefinitionTimeoutPtrOutput struct { *pulumi.OutputState}

func (JobDefinitionTimeoutPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinitionTimeout)(nil)).Elem()
}

func (o JobDefinitionTimeoutPtrOutput) ToJobDefinitionTimeoutPtrOutput() JobDefinitionTimeoutPtrOutput {
	return o
}

func (o JobDefinitionTimeoutPtrOutput) ToJobDefinitionTimeoutPtrOutputWithContext(ctx context.Context) JobDefinitionTimeoutPtrOutput {
	return o
}

func (o JobDefinitionTimeoutPtrOutput) Elem() JobDefinitionTimeoutOutput {
	return o.ApplyT(func (v *JobDefinitionTimeout) JobDefinitionTimeout { return *v }).(JobDefinitionTimeoutOutput)
}

func (o JobDefinitionTimeoutPtrOutput) AttemptDurationSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func (v JobDefinitionTimeout) *int { return v.AttemptDurationSeconds }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(JobDefinitionTimeoutOutput{})
	pulumi.RegisterOutputType(JobDefinitionTimeoutPtrOutput{})
}
