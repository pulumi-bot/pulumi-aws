// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codecommit

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type TriggerTrigger struct {
	// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
	Branches []string `pulumi:"branches"`
	// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
	CustomData *string `pulumi:"customData"`
	// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
	DestinationArn string `pulumi:"destinationArn"`
	// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
	Events []string `pulumi:"events"`
	// The name of the trigger.
	Name string `pulumi:"name"`
}

type TriggerTriggerInput interface {
	pulumi.Input

	ToTriggerTriggerOutput() TriggerTriggerOutput
	ToTriggerTriggerOutputWithContext(context.Context) TriggerTriggerOutput
}

type TriggerTriggerArgs struct {
	// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
	Branches pulumi.StringArrayInput `pulumi:"branches"`
	// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
	CustomData pulumi.StringPtrInput `pulumi:"customData"`
	// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
	DestinationArn pulumi.StringInput `pulumi:"destinationArn"`
	// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
	Events pulumi.StringArrayInput `pulumi:"events"`
	// The name of the trigger.
	Name pulumi.StringInput `pulumi:"name"`
}

func (TriggerTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerTrigger)(nil)).Elem()
}

func (i TriggerTriggerArgs) ToTriggerTriggerOutput() TriggerTriggerOutput {
	return i.ToTriggerTriggerOutputWithContext(context.Background())
}

func (i TriggerTriggerArgs) ToTriggerTriggerOutputWithContext(ctx context.Context) TriggerTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerTriggerOutput)
}

type TriggerTriggerArrayInput interface {
	pulumi.Input

	ToTriggerTriggerArrayOutput() TriggerTriggerArrayOutput
	ToTriggerTriggerArrayOutputWithContext(context.Context) TriggerTriggerArrayOutput
}

type TriggerTriggerArray []TriggerTriggerInput

func (TriggerTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerTrigger)(nil)).Elem()
}

func (i TriggerTriggerArray) ToTriggerTriggerArrayOutput() TriggerTriggerArrayOutput {
	return i.ToTriggerTriggerArrayOutputWithContext(context.Background())
}

func (i TriggerTriggerArray) ToTriggerTriggerArrayOutputWithContext(ctx context.Context) TriggerTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerTriggerArrayOutput)
}

type TriggerTriggerOutput struct{ *pulumi.OutputState }

func (TriggerTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerTrigger)(nil)).Elem()
}

func (o TriggerTriggerOutput) ToTriggerTriggerOutput() TriggerTriggerOutput {
	return o
}

func (o TriggerTriggerOutput) ToTriggerTriggerOutputWithContext(ctx context.Context) TriggerTriggerOutput {
	return o
}

// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
func (o TriggerTriggerOutput) Branches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TriggerTrigger) []string { return v.Branches }).(pulumi.StringArrayOutput)
}

// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
func (o TriggerTriggerOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerTrigger) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
func (o TriggerTriggerOutput) DestinationArn() pulumi.StringOutput {
	return o.ApplyT(func(v TriggerTrigger) string { return v.DestinationArn }).(pulumi.StringOutput)
}

// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
func (o TriggerTriggerOutput) Events() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TriggerTrigger) []string { return v.Events }).(pulumi.StringArrayOutput)
}

// The name of the trigger.
func (o TriggerTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TriggerTrigger) string { return v.Name }).(pulumi.StringOutput)
}

type TriggerTriggerArrayOutput struct{ *pulumi.OutputState }

func (TriggerTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerTrigger)(nil)).Elem()
}

func (o TriggerTriggerArrayOutput) ToTriggerTriggerArrayOutput() TriggerTriggerArrayOutput {
	return o
}

func (o TriggerTriggerArrayOutput) ToTriggerTriggerArrayOutputWithContext(ctx context.Context) TriggerTriggerArrayOutput {
	return o
}

func (o TriggerTriggerArrayOutput) Index(i pulumi.IntInput) TriggerTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TriggerTrigger {
		return vs[0].([]TriggerTrigger)[vs[1].(int)]
	}).(TriggerTriggerOutput)
}

type TriggerTriggerArgs struct {
	// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
	Branches []string `pulumi:"branches"`
	// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
	CustomData *string `pulumi:"customData"`
	// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
	DestinationArn string `pulumi:"destinationArn"`
	// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
	Events []string `pulumi:"events"`
	// The name of the trigger.
	Name string `pulumi:"name"`
}

type TriggerTriggerArgsInput interface {
	pulumi.Input

	ToTriggerTriggerArgsOutput() TriggerTriggerArgsOutput
	ToTriggerTriggerArgsOutputWithContext(context.Context) TriggerTriggerArgsOutput
}

type TriggerTriggerArgsArgs struct {
	// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
	Branches pulumi.StringArrayInput `pulumi:"branches"`
	// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
	CustomData pulumi.StringPtrInput `pulumi:"customData"`
	// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
	DestinationArn pulumi.StringInput `pulumi:"destinationArn"`
	// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
	Events pulumi.StringArrayInput `pulumi:"events"`
	// The name of the trigger.
	Name pulumi.StringInput `pulumi:"name"`
}

func (TriggerTriggerArgsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerTriggerArgs)(nil)).Elem()
}

func (i TriggerTriggerArgsArgs) ToTriggerTriggerArgsOutput() TriggerTriggerArgsOutput {
	return i.ToTriggerTriggerArgsOutputWithContext(context.Background())
}

func (i TriggerTriggerArgsArgs) ToTriggerTriggerArgsOutputWithContext(ctx context.Context) TriggerTriggerArgsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerTriggerArgsOutput)
}

type TriggerTriggerArgsArrayInput interface {
	pulumi.Input

	ToTriggerTriggerArgsArrayOutput() TriggerTriggerArgsArrayOutput
	ToTriggerTriggerArgsArrayOutputWithContext(context.Context) TriggerTriggerArgsArrayOutput
}

type TriggerTriggerArgsArray []TriggerTriggerArgsInput

func (TriggerTriggerArgsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerTriggerArgs)(nil)).Elem()
}

func (i TriggerTriggerArgsArray) ToTriggerTriggerArgsArrayOutput() TriggerTriggerArgsArrayOutput {
	return i.ToTriggerTriggerArgsArrayOutputWithContext(context.Background())
}

func (i TriggerTriggerArgsArray) ToTriggerTriggerArgsArrayOutputWithContext(ctx context.Context) TriggerTriggerArgsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerTriggerArgsArrayOutput)
}

type TriggerTriggerArgsOutput struct{ *pulumi.OutputState }

func (TriggerTriggerArgsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerTriggerArgs)(nil)).Elem()
}

func (o TriggerTriggerArgsOutput) ToTriggerTriggerArgsOutput() TriggerTriggerArgsOutput {
	return o
}

func (o TriggerTriggerArgsOutput) ToTriggerTriggerArgsOutputWithContext(ctx context.Context) TriggerTriggerArgsOutput {
	return o
}

// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
func (o TriggerTriggerArgsOutput) Branches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TriggerTriggerArgs) []string { return v.Branches }).(pulumi.StringArrayOutput)
}

// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
func (o TriggerTriggerArgsOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerTriggerArgs) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
func (o TriggerTriggerArgsOutput) DestinationArn() pulumi.StringOutput {
	return o.ApplyT(func(v TriggerTriggerArgs) string { return v.DestinationArn }).(pulumi.StringOutput)
}

// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
func (o TriggerTriggerArgsOutput) Events() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TriggerTriggerArgs) []string { return v.Events }).(pulumi.StringArrayOutput)
}

// The name of the trigger.
func (o TriggerTriggerArgsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TriggerTriggerArgs) string { return v.Name }).(pulumi.StringOutput)
}

type TriggerTriggerArgsArrayOutput struct{ *pulumi.OutputState }

func (TriggerTriggerArgsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerTriggerArgs)(nil)).Elem()
}

func (o TriggerTriggerArgsArrayOutput) ToTriggerTriggerArgsArrayOutput() TriggerTriggerArgsArrayOutput {
	return o
}

func (o TriggerTriggerArgsArrayOutput) ToTriggerTriggerArgsArrayOutputWithContext(ctx context.Context) TriggerTriggerArgsArrayOutput {
	return o
}

func (o TriggerTriggerArgsArrayOutput) Index(i pulumi.IntInput) TriggerTriggerArgsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TriggerTriggerArgs {
		return vs[0].([]TriggerTriggerArgs)[vs[1].(int)]
	}).(TriggerTriggerArgsOutput)
}

type TriggerTriggerState struct {
	// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
	Branches []string `pulumi:"branches"`
	// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
	CustomData *string `pulumi:"customData"`
	// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
	DestinationArn string `pulumi:"destinationArn"`
	// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
	Events []string `pulumi:"events"`
	// The name of the trigger.
	Name string `pulumi:"name"`
}

type TriggerTriggerStateInput interface {
	pulumi.Input

	ToTriggerTriggerStateOutput() TriggerTriggerStateOutput
	ToTriggerTriggerStateOutputWithContext(context.Context) TriggerTriggerStateOutput
}

type TriggerTriggerStateArgs struct {
	// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
	Branches pulumi.StringArrayInput `pulumi:"branches"`
	// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
	CustomData pulumi.StringPtrInput `pulumi:"customData"`
	// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
	DestinationArn pulumi.StringInput `pulumi:"destinationArn"`
	// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
	Events pulumi.StringArrayInput `pulumi:"events"`
	// The name of the trigger.
	Name pulumi.StringInput `pulumi:"name"`
}

func (TriggerTriggerStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerTriggerState)(nil)).Elem()
}

func (i TriggerTriggerStateArgs) ToTriggerTriggerStateOutput() TriggerTriggerStateOutput {
	return i.ToTriggerTriggerStateOutputWithContext(context.Background())
}

func (i TriggerTriggerStateArgs) ToTriggerTriggerStateOutputWithContext(ctx context.Context) TriggerTriggerStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerTriggerStateOutput)
}

type TriggerTriggerStateArrayInput interface {
	pulumi.Input

	ToTriggerTriggerStateArrayOutput() TriggerTriggerStateArrayOutput
	ToTriggerTriggerStateArrayOutputWithContext(context.Context) TriggerTriggerStateArrayOutput
}

type TriggerTriggerStateArray []TriggerTriggerStateInput

func (TriggerTriggerStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerTriggerState)(nil)).Elem()
}

func (i TriggerTriggerStateArray) ToTriggerTriggerStateArrayOutput() TriggerTriggerStateArrayOutput {
	return i.ToTriggerTriggerStateArrayOutputWithContext(context.Background())
}

func (i TriggerTriggerStateArray) ToTriggerTriggerStateArrayOutputWithContext(ctx context.Context) TriggerTriggerStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerTriggerStateArrayOutput)
}

type TriggerTriggerStateOutput struct{ *pulumi.OutputState }

func (TriggerTriggerStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerTriggerState)(nil)).Elem()
}

func (o TriggerTriggerStateOutput) ToTriggerTriggerStateOutput() TriggerTriggerStateOutput {
	return o
}

func (o TriggerTriggerStateOutput) ToTriggerTriggerStateOutputWithContext(ctx context.Context) TriggerTriggerStateOutput {
	return o
}

// The branches that will be included in the trigger configuration. If no branches are specified, the trigger will apply to all branches.
func (o TriggerTriggerStateOutput) Branches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TriggerTriggerState) []string { return v.Branches }).(pulumi.StringArrayOutput)
}

// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
func (o TriggerTriggerStateOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TriggerTriggerState) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
func (o TriggerTriggerStateOutput) DestinationArn() pulumi.StringOutput {
	return o.ApplyT(func(v TriggerTriggerState) string { return v.DestinationArn }).(pulumi.StringOutput)
}

// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
func (o TriggerTriggerStateOutput) Events() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TriggerTriggerState) []string { return v.Events }).(pulumi.StringArrayOutput)
}

// The name of the trigger.
func (o TriggerTriggerStateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TriggerTriggerState) string { return v.Name }).(pulumi.StringOutput)
}

type TriggerTriggerStateArrayOutput struct{ *pulumi.OutputState }

func (TriggerTriggerStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerTriggerState)(nil)).Elem()
}

func (o TriggerTriggerStateArrayOutput) ToTriggerTriggerStateArrayOutput() TriggerTriggerStateArrayOutput {
	return o
}

func (o TriggerTriggerStateArrayOutput) ToTriggerTriggerStateArrayOutputWithContext(ctx context.Context) TriggerTriggerStateArrayOutput {
	return o
}

func (o TriggerTriggerStateArrayOutput) Index(i pulumi.IntInput) TriggerTriggerStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TriggerTriggerState {
		return vs[0].([]TriggerTriggerState)[vs[1].(int)]
	}).(TriggerTriggerStateOutput)
}

func init() {
	pulumi.RegisterOutputType(TriggerTriggerOutput{})
	pulumi.RegisterOutputType(TriggerTriggerArrayOutput{})
	pulumi.RegisterOutputType(TriggerTriggerArgsOutput{})
	pulumi.RegisterOutputType(TriggerTriggerArgsArrayOutput{})
	pulumi.RegisterOutputType(TriggerTriggerStateOutput{})
	pulumi.RegisterOutputType(TriggerTriggerStateArrayOutput{})
}
