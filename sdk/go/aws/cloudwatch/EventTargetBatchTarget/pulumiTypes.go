// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package EventTargetBatchTarget

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type EventTargetBatchTarget struct {
	// The size of the array, if this is an array batch job. Valid values are integers between 2 and 10,000.
	ArraySize *int `pulumi:"arraySize"`
	// The number of times to attempt to retry, if the job fails. Valid values are 1 to 10.
	JobAttempts *int `pulumi:"jobAttempts"`
	// The ARN or name of the job definition to use if the event target is an AWS Batch job. This job definition must already exist.
	JobDefinition string `pulumi:"jobDefinition"`
	// The name to use for this execution of the job, if the target is an AWS Batch job.
	JobName string `pulumi:"jobName"`
}

type EventTargetBatchTargetInput interface {
	pulumi.Input

	ToEventTargetBatchTargetOutput() EventTargetBatchTargetOutput
	ToEventTargetBatchTargetOutputWithContext(context.Context) EventTargetBatchTargetOutput
}

type EventTargetBatchTargetArgs struct {
	// The size of the array, if this is an array batch job. Valid values are integers between 2 and 10,000.
	ArraySize pulumi.IntPtrInput `pulumi:"arraySize"`
	// The number of times to attempt to retry, if the job fails. Valid values are 1 to 10.
	JobAttempts pulumi.IntPtrInput `pulumi:"jobAttempts"`
	// The ARN or name of the job definition to use if the event target is an AWS Batch job. This job definition must already exist.
	JobDefinition pulumi.StringInput `pulumi:"jobDefinition"`
	// The name to use for this execution of the job, if the target is an AWS Batch job.
	JobName pulumi.StringInput `pulumi:"jobName"`
}

func (EventTargetBatchTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventTargetBatchTarget)(nil)).Elem()
}

func (i EventTargetBatchTargetArgs) ToEventTargetBatchTargetOutput() EventTargetBatchTargetOutput {
	return i.ToEventTargetBatchTargetOutputWithContext(context.Background())
}

func (i EventTargetBatchTargetArgs) ToEventTargetBatchTargetOutputWithContext(ctx context.Context) EventTargetBatchTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTargetBatchTargetOutput)
}

func (i EventTargetBatchTargetArgs) ToEventTargetBatchTargetPtrOutput() EventTargetBatchTargetPtrOutput {
	return i.ToEventTargetBatchTargetPtrOutputWithContext(context.Background())
}

func (i EventTargetBatchTargetArgs) ToEventTargetBatchTargetPtrOutputWithContext(ctx context.Context) EventTargetBatchTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTargetBatchTargetOutput).ToEventTargetBatchTargetPtrOutputWithContext(ctx)
}

type EventTargetBatchTargetPtrInput interface {
	pulumi.Input

	ToEventTargetBatchTargetPtrOutput() EventTargetBatchTargetPtrOutput
	ToEventTargetBatchTargetPtrOutputWithContext(context.Context) EventTargetBatchTargetPtrOutput
}

type eventTargetBatchTargetPtrType EventTargetBatchTargetArgs

func EventTargetBatchTargetPtr(v *EventTargetBatchTargetArgs) EventTargetBatchTargetPtrInput {	return (*eventTargetBatchTargetPtrType)(v)
}

func (*eventTargetBatchTargetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventTargetBatchTarget)(nil)).Elem()
}

func (i *eventTargetBatchTargetPtrType) ToEventTargetBatchTargetPtrOutput() EventTargetBatchTargetPtrOutput {
	return i.ToEventTargetBatchTargetPtrOutputWithContext(context.Background())
}

func (i *eventTargetBatchTargetPtrType) ToEventTargetBatchTargetPtrOutputWithContext(ctx context.Context) EventTargetBatchTargetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTargetBatchTargetPtrOutput)
}

type EventTargetBatchTargetOutput struct { *pulumi.OutputState }

func (EventTargetBatchTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventTargetBatchTarget)(nil)).Elem()
}

func (o EventTargetBatchTargetOutput) ToEventTargetBatchTargetOutput() EventTargetBatchTargetOutput {
	return o
}

func (o EventTargetBatchTargetOutput) ToEventTargetBatchTargetOutputWithContext(ctx context.Context) EventTargetBatchTargetOutput {
	return o
}

func (o EventTargetBatchTargetOutput) ToEventTargetBatchTargetPtrOutput() EventTargetBatchTargetPtrOutput {
	return o.ToEventTargetBatchTargetPtrOutputWithContext(context.Background())
}

func (o EventTargetBatchTargetOutput) ToEventTargetBatchTargetPtrOutputWithContext(ctx context.Context) EventTargetBatchTargetPtrOutput {
	return o.ApplyT(func(v EventTargetBatchTarget) *EventTargetBatchTarget {
		return &v
	}).(EventTargetBatchTargetPtrOutput)
}
// The size of the array, if this is an array batch job. Valid values are integers between 2 and 10,000.
func (o EventTargetBatchTargetOutput) ArraySize() pulumi.IntPtrOutput {
	return o.ApplyT(func (v EventTargetBatchTarget) *int { return v.ArraySize }).(pulumi.IntPtrOutput)
}

// The number of times to attempt to retry, if the job fails. Valid values are 1 to 10.
func (o EventTargetBatchTargetOutput) JobAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func (v EventTargetBatchTarget) *int { return v.JobAttempts }).(pulumi.IntPtrOutput)
}

// The ARN or name of the job definition to use if the event target is an AWS Batch job. This job definition must already exist.
func (o EventTargetBatchTargetOutput) JobDefinition() pulumi.StringOutput {
	return o.ApplyT(func (v EventTargetBatchTarget) string { return v.JobDefinition }).(pulumi.StringOutput)
}

// The name to use for this execution of the job, if the target is an AWS Batch job.
func (o EventTargetBatchTargetOutput) JobName() pulumi.StringOutput {
	return o.ApplyT(func (v EventTargetBatchTarget) string { return v.JobName }).(pulumi.StringOutput)
}

type EventTargetBatchTargetPtrOutput struct { *pulumi.OutputState}

func (EventTargetBatchTargetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventTargetBatchTarget)(nil)).Elem()
}

func (o EventTargetBatchTargetPtrOutput) ToEventTargetBatchTargetPtrOutput() EventTargetBatchTargetPtrOutput {
	return o
}

func (o EventTargetBatchTargetPtrOutput) ToEventTargetBatchTargetPtrOutputWithContext(ctx context.Context) EventTargetBatchTargetPtrOutput {
	return o
}

func (o EventTargetBatchTargetPtrOutput) Elem() EventTargetBatchTargetOutput {
	return o.ApplyT(func (v *EventTargetBatchTarget) EventTargetBatchTarget { return *v }).(EventTargetBatchTargetOutput)
}

// The size of the array, if this is an array batch job. Valid values are integers between 2 and 10,000.
func (o EventTargetBatchTargetPtrOutput) ArraySize() pulumi.IntPtrOutput {
	return o.ApplyT(func (v EventTargetBatchTarget) *int { return v.ArraySize }).(pulumi.IntPtrOutput)
}

// The number of times to attempt to retry, if the job fails. Valid values are 1 to 10.
func (o EventTargetBatchTargetPtrOutput) JobAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func (v EventTargetBatchTarget) *int { return v.JobAttempts }).(pulumi.IntPtrOutput)
}

// The ARN or name of the job definition to use if the event target is an AWS Batch job. This job definition must already exist.
func (o EventTargetBatchTargetPtrOutput) JobDefinition() pulumi.StringOutput {
	return o.ApplyT(func (v EventTargetBatchTarget) string { return v.JobDefinition }).(pulumi.StringOutput)
}

// The name to use for this execution of the job, if the target is an AWS Batch job.
func (o EventTargetBatchTargetPtrOutput) JobName() pulumi.StringOutput {
	return o.ApplyT(func (v EventTargetBatchTarget) string { return v.JobName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventTargetBatchTargetOutput{})
	pulumi.RegisterOutputType(EventTargetBatchTargetPtrOutput{})
}
