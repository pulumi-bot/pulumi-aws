// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package TopicRuleCloudwatchAlarm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type TopicRuleCloudwatchAlarm struct {
	// The CloudWatch alarm name.
	AlarmName string `pulumi:"alarmName"`
	// The ARN of the IAM role that grants access.
	RoleArn string `pulumi:"roleArn"`
	// The reason for the alarm change.
	StateReason string `pulumi:"stateReason"`
	// The value of the alarm state. Acceptable values are: OK, ALARM, INSUFFICIENT_DATA.
	StateValue string `pulumi:"stateValue"`
}

type TopicRuleCloudwatchAlarmInput interface {
	pulumi.Input

	ToTopicRuleCloudwatchAlarmOutput() TopicRuleCloudwatchAlarmOutput
	ToTopicRuleCloudwatchAlarmOutputWithContext(context.Context) TopicRuleCloudwatchAlarmOutput
}

type TopicRuleCloudwatchAlarmArgs struct {
	// The CloudWatch alarm name.
	AlarmName pulumi.StringInput `pulumi:"alarmName"`
	// The ARN of the IAM role that grants access.
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
	// The reason for the alarm change.
	StateReason pulumi.StringInput `pulumi:"stateReason"`
	// The value of the alarm state. Acceptable values are: OK, ALARM, INSUFFICIENT_DATA.
	StateValue pulumi.StringInput `pulumi:"stateValue"`
}

func (TopicRuleCloudwatchAlarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicRuleCloudwatchAlarm)(nil)).Elem()
}

func (i TopicRuleCloudwatchAlarmArgs) ToTopicRuleCloudwatchAlarmOutput() TopicRuleCloudwatchAlarmOutput {
	return i.ToTopicRuleCloudwatchAlarmOutputWithContext(context.Background())
}

func (i TopicRuleCloudwatchAlarmArgs) ToTopicRuleCloudwatchAlarmOutputWithContext(ctx context.Context) TopicRuleCloudwatchAlarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicRuleCloudwatchAlarmOutput)
}

func (i TopicRuleCloudwatchAlarmArgs) ToTopicRuleCloudwatchAlarmPtrOutput() TopicRuleCloudwatchAlarmPtrOutput {
	return i.ToTopicRuleCloudwatchAlarmPtrOutputWithContext(context.Background())
}

func (i TopicRuleCloudwatchAlarmArgs) ToTopicRuleCloudwatchAlarmPtrOutputWithContext(ctx context.Context) TopicRuleCloudwatchAlarmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicRuleCloudwatchAlarmOutput).ToTopicRuleCloudwatchAlarmPtrOutputWithContext(ctx)
}

type TopicRuleCloudwatchAlarmPtrInput interface {
	pulumi.Input

	ToTopicRuleCloudwatchAlarmPtrOutput() TopicRuleCloudwatchAlarmPtrOutput
	ToTopicRuleCloudwatchAlarmPtrOutputWithContext(context.Context) TopicRuleCloudwatchAlarmPtrOutput
}

type topicRuleCloudwatchAlarmPtrType TopicRuleCloudwatchAlarmArgs

func TopicRuleCloudwatchAlarmPtr(v *TopicRuleCloudwatchAlarmArgs) TopicRuleCloudwatchAlarmPtrInput {	return (*topicRuleCloudwatchAlarmPtrType)(v)
}

func (*topicRuleCloudwatchAlarmPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicRuleCloudwatchAlarm)(nil)).Elem()
}

func (i *topicRuleCloudwatchAlarmPtrType) ToTopicRuleCloudwatchAlarmPtrOutput() TopicRuleCloudwatchAlarmPtrOutput {
	return i.ToTopicRuleCloudwatchAlarmPtrOutputWithContext(context.Background())
}

func (i *topicRuleCloudwatchAlarmPtrType) ToTopicRuleCloudwatchAlarmPtrOutputWithContext(ctx context.Context) TopicRuleCloudwatchAlarmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicRuleCloudwatchAlarmPtrOutput)
}

type TopicRuleCloudwatchAlarmOutput struct { *pulumi.OutputState }

func (TopicRuleCloudwatchAlarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicRuleCloudwatchAlarm)(nil)).Elem()
}

func (o TopicRuleCloudwatchAlarmOutput) ToTopicRuleCloudwatchAlarmOutput() TopicRuleCloudwatchAlarmOutput {
	return o
}

func (o TopicRuleCloudwatchAlarmOutput) ToTopicRuleCloudwatchAlarmOutputWithContext(ctx context.Context) TopicRuleCloudwatchAlarmOutput {
	return o
}

func (o TopicRuleCloudwatchAlarmOutput) ToTopicRuleCloudwatchAlarmPtrOutput() TopicRuleCloudwatchAlarmPtrOutput {
	return o.ToTopicRuleCloudwatchAlarmPtrOutputWithContext(context.Background())
}

func (o TopicRuleCloudwatchAlarmOutput) ToTopicRuleCloudwatchAlarmPtrOutputWithContext(ctx context.Context) TopicRuleCloudwatchAlarmPtrOutput {
	return o.ApplyT(func(v TopicRuleCloudwatchAlarm) *TopicRuleCloudwatchAlarm {
		return &v
	}).(TopicRuleCloudwatchAlarmPtrOutput)
}
// The CloudWatch alarm name.
func (o TopicRuleCloudwatchAlarmOutput) AlarmName() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchAlarm) string { return v.AlarmName }).(pulumi.StringOutput)
}

// The ARN of the IAM role that grants access.
func (o TopicRuleCloudwatchAlarmOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchAlarm) string { return v.RoleArn }).(pulumi.StringOutput)
}

// The reason for the alarm change.
func (o TopicRuleCloudwatchAlarmOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchAlarm) string { return v.StateReason }).(pulumi.StringOutput)
}

// The value of the alarm state. Acceptable values are: OK, ALARM, INSUFFICIENT_DATA.
func (o TopicRuleCloudwatchAlarmOutput) StateValue() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchAlarm) string { return v.StateValue }).(pulumi.StringOutput)
}

type TopicRuleCloudwatchAlarmPtrOutput struct { *pulumi.OutputState}

func (TopicRuleCloudwatchAlarmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicRuleCloudwatchAlarm)(nil)).Elem()
}

func (o TopicRuleCloudwatchAlarmPtrOutput) ToTopicRuleCloudwatchAlarmPtrOutput() TopicRuleCloudwatchAlarmPtrOutput {
	return o
}

func (o TopicRuleCloudwatchAlarmPtrOutput) ToTopicRuleCloudwatchAlarmPtrOutputWithContext(ctx context.Context) TopicRuleCloudwatchAlarmPtrOutput {
	return o
}

func (o TopicRuleCloudwatchAlarmPtrOutput) Elem() TopicRuleCloudwatchAlarmOutput {
	return o.ApplyT(func (v *TopicRuleCloudwatchAlarm) TopicRuleCloudwatchAlarm { return *v }).(TopicRuleCloudwatchAlarmOutput)
}

// The CloudWatch alarm name.
func (o TopicRuleCloudwatchAlarmPtrOutput) AlarmName() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchAlarm) string { return v.AlarmName }).(pulumi.StringOutput)
}

// The ARN of the IAM role that grants access.
func (o TopicRuleCloudwatchAlarmPtrOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchAlarm) string { return v.RoleArn }).(pulumi.StringOutput)
}

// The reason for the alarm change.
func (o TopicRuleCloudwatchAlarmPtrOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchAlarm) string { return v.StateReason }).(pulumi.StringOutput)
}

// The value of the alarm state. Acceptable values are: OK, ALARM, INSUFFICIENT_DATA.
func (o TopicRuleCloudwatchAlarmPtrOutput) StateValue() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchAlarm) string { return v.StateValue }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicRuleCloudwatchAlarmOutput{})
	pulumi.RegisterOutputType(TopicRuleCloudwatchAlarmPtrOutput{})
}
