// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mediaconvert

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type QueueReservationPlanSettings struct {
	// The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
	Commitment string `pulumi:"commitment"`
	// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
	RenewalType string `pulumi:"renewalType"`
	// Specifies the number of reserved transcode slots (RTS) for queue.
	ReservedSlots int `pulumi:"reservedSlots"`
}

type QueueReservationPlanSettingsInput interface {
	pulumi.Input

	ToQueueReservationPlanSettingsOutput() QueueReservationPlanSettingsOutput
	ToQueueReservationPlanSettingsOutputWithContext(context.Context) QueueReservationPlanSettingsOutput
}

type QueueReservationPlanSettingsArgs struct {
	// The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
	Commitment pulumi.StringInput `pulumi:"commitment"`
	// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
	RenewalType pulumi.StringInput `pulumi:"renewalType"`
	// Specifies the number of reserved transcode slots (RTS) for queue.
	ReservedSlots pulumi.IntInput `pulumi:"reservedSlots"`
}

func (QueueReservationPlanSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueReservationPlanSettings)(nil)).Elem()
}

func (i QueueReservationPlanSettingsArgs) ToQueueReservationPlanSettingsOutput() QueueReservationPlanSettingsOutput {
	return i.ToQueueReservationPlanSettingsOutputWithContext(context.Background())
}

func (i QueueReservationPlanSettingsArgs) ToQueueReservationPlanSettingsOutputWithContext(ctx context.Context) QueueReservationPlanSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueReservationPlanSettingsOutput)
}

func (i QueueReservationPlanSettingsArgs) ToQueueReservationPlanSettingsPtrOutput() QueueReservationPlanSettingsPtrOutput {
	return i.ToQueueReservationPlanSettingsPtrOutputWithContext(context.Background())
}

func (i QueueReservationPlanSettingsArgs) ToQueueReservationPlanSettingsPtrOutputWithContext(ctx context.Context) QueueReservationPlanSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueReservationPlanSettingsOutput).ToQueueReservationPlanSettingsPtrOutputWithContext(ctx)
}

type QueueReservationPlanSettingsPtrInput interface {
	pulumi.Input

	ToQueueReservationPlanSettingsPtrOutput() QueueReservationPlanSettingsPtrOutput
	ToQueueReservationPlanSettingsPtrOutputWithContext(context.Context) QueueReservationPlanSettingsPtrOutput
}

type queueReservationPlanSettingsPtrType QueueReservationPlanSettingsArgs

func QueueReservationPlanSettingsPtr(v *QueueReservationPlanSettingsArgs) QueueReservationPlanSettingsPtrInput {
	return (*queueReservationPlanSettingsPtrType)(v)
}

func (*queueReservationPlanSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueReservationPlanSettings)(nil)).Elem()
}

func (i *queueReservationPlanSettingsPtrType) ToQueueReservationPlanSettingsPtrOutput() QueueReservationPlanSettingsPtrOutput {
	return i.ToQueueReservationPlanSettingsPtrOutputWithContext(context.Background())
}

func (i *queueReservationPlanSettingsPtrType) ToQueueReservationPlanSettingsPtrOutputWithContext(ctx context.Context) QueueReservationPlanSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueReservationPlanSettingsPtrOutput)
}

type QueueReservationPlanSettingsOutput struct{ *pulumi.OutputState }

func (QueueReservationPlanSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueReservationPlanSettings)(nil)).Elem()
}

func (o QueueReservationPlanSettingsOutput) ToQueueReservationPlanSettingsOutput() QueueReservationPlanSettingsOutput {
	return o
}

func (o QueueReservationPlanSettingsOutput) ToQueueReservationPlanSettingsOutputWithContext(ctx context.Context) QueueReservationPlanSettingsOutput {
	return o
}

func (o QueueReservationPlanSettingsOutput) ToQueueReservationPlanSettingsPtrOutput() QueueReservationPlanSettingsPtrOutput {
	return o.ToQueueReservationPlanSettingsPtrOutputWithContext(context.Background())
}

func (o QueueReservationPlanSettingsOutput) ToQueueReservationPlanSettingsPtrOutputWithContext(ctx context.Context) QueueReservationPlanSettingsPtrOutput {
	return o.ApplyT(func(v QueueReservationPlanSettings) *QueueReservationPlanSettings {
		return &v
	}).(QueueReservationPlanSettingsPtrOutput)
}

// The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
func (o QueueReservationPlanSettingsOutput) Commitment() pulumi.StringOutput {
	return o.ApplyT(func(v QueueReservationPlanSettings) string { return v.Commitment }).(pulumi.StringOutput)
}

// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
func (o QueueReservationPlanSettingsOutput) RenewalType() pulumi.StringOutput {
	return o.ApplyT(func(v QueueReservationPlanSettings) string { return v.RenewalType }).(pulumi.StringOutput)
}

// Specifies the number of reserved transcode slots (RTS) for queue.
func (o QueueReservationPlanSettingsOutput) ReservedSlots() pulumi.IntOutput {
	return o.ApplyT(func(v QueueReservationPlanSettings) int { return v.ReservedSlots }).(pulumi.IntOutput)
}

type QueueReservationPlanSettingsPtrOutput struct{ *pulumi.OutputState }

func (QueueReservationPlanSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueReservationPlanSettings)(nil)).Elem()
}

func (o QueueReservationPlanSettingsPtrOutput) ToQueueReservationPlanSettingsPtrOutput() QueueReservationPlanSettingsPtrOutput {
	return o
}

func (o QueueReservationPlanSettingsPtrOutput) ToQueueReservationPlanSettingsPtrOutputWithContext(ctx context.Context) QueueReservationPlanSettingsPtrOutput {
	return o
}

func (o QueueReservationPlanSettingsPtrOutput) Elem() QueueReservationPlanSettingsOutput {
	return o.ApplyT(func(v *QueueReservationPlanSettings) QueueReservationPlanSettings { return *v }).(QueueReservationPlanSettingsOutput)
}

// The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
func (o QueueReservationPlanSettingsPtrOutput) Commitment() pulumi.StringOutput {
	return o.ApplyT(func(v QueueReservationPlanSettings) string { return v.Commitment }).(pulumi.StringOutput)
}

// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
func (o QueueReservationPlanSettingsPtrOutput) RenewalType() pulumi.StringOutput {
	return o.ApplyT(func(v QueueReservationPlanSettings) string { return v.RenewalType }).(pulumi.StringOutput)
}

// Specifies the number of reserved transcode slots (RTS) for queue.
func (o QueueReservationPlanSettingsPtrOutput) ReservedSlots() pulumi.IntOutput {
	return o.ApplyT(func(v QueueReservationPlanSettings) int { return v.ReservedSlots }).(pulumi.IntOutput)
}

type QueueReservationPlanSettingsArgs struct {
	// The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
	Commitment string `pulumi:"commitment"`
	// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
	RenewalType string `pulumi:"renewalType"`
	// Specifies the number of reserved transcode slots (RTS) for queue.
	ReservedSlots int `pulumi:"reservedSlots"`
}

type QueueReservationPlanSettingsArgsInput interface {
	pulumi.Input

	ToQueueReservationPlanSettingsArgsOutput() QueueReservationPlanSettingsArgsOutput
	ToQueueReservationPlanSettingsArgsOutputWithContext(context.Context) QueueReservationPlanSettingsArgsOutput
}

type QueueReservationPlanSettingsArgsArgs struct {
	// The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
	Commitment pulumi.StringInput `pulumi:"commitment"`
	// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
	RenewalType pulumi.StringInput `pulumi:"renewalType"`
	// Specifies the number of reserved transcode slots (RTS) for queue.
	ReservedSlots pulumi.IntInput `pulumi:"reservedSlots"`
}

func (QueueReservationPlanSettingsArgsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueReservationPlanSettingsArgs)(nil)).Elem()
}

func (i QueueReservationPlanSettingsArgsArgs) ToQueueReservationPlanSettingsArgsOutput() QueueReservationPlanSettingsArgsOutput {
	return i.ToQueueReservationPlanSettingsArgsOutputWithContext(context.Background())
}

func (i QueueReservationPlanSettingsArgsArgs) ToQueueReservationPlanSettingsArgsOutputWithContext(ctx context.Context) QueueReservationPlanSettingsArgsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueReservationPlanSettingsArgsOutput)
}

func (i QueueReservationPlanSettingsArgsArgs) ToQueueReservationPlanSettingsArgsPtrOutput() QueueReservationPlanSettingsArgsPtrOutput {
	return i.ToQueueReservationPlanSettingsArgsPtrOutputWithContext(context.Background())
}

func (i QueueReservationPlanSettingsArgsArgs) ToQueueReservationPlanSettingsArgsPtrOutputWithContext(ctx context.Context) QueueReservationPlanSettingsArgsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueReservationPlanSettingsArgsOutput).ToQueueReservationPlanSettingsArgsPtrOutputWithContext(ctx)
}

type QueueReservationPlanSettingsArgsPtrInput interface {
	pulumi.Input

	ToQueueReservationPlanSettingsArgsPtrOutput() QueueReservationPlanSettingsArgsPtrOutput
	ToQueueReservationPlanSettingsArgsPtrOutputWithContext(context.Context) QueueReservationPlanSettingsArgsPtrOutput
}

type queueReservationPlanSettingsArgsPtrType QueueReservationPlanSettingsArgsArgs

func QueueReservationPlanSettingsArgsPtr(v *QueueReservationPlanSettingsArgsArgs) QueueReservationPlanSettingsArgsPtrInput {
	return (*queueReservationPlanSettingsArgsPtrType)(v)
}

func (*queueReservationPlanSettingsArgsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueReservationPlanSettingsArgs)(nil)).Elem()
}

func (i *queueReservationPlanSettingsArgsPtrType) ToQueueReservationPlanSettingsArgsPtrOutput() QueueReservationPlanSettingsArgsPtrOutput {
	return i.ToQueueReservationPlanSettingsArgsPtrOutputWithContext(context.Background())
}

func (i *queueReservationPlanSettingsArgsPtrType) ToQueueReservationPlanSettingsArgsPtrOutputWithContext(ctx context.Context) QueueReservationPlanSettingsArgsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueReservationPlanSettingsArgsPtrOutput)
}

type QueueReservationPlanSettingsArgsOutput struct{ *pulumi.OutputState }

func (QueueReservationPlanSettingsArgsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueReservationPlanSettingsArgs)(nil)).Elem()
}

func (o QueueReservationPlanSettingsArgsOutput) ToQueueReservationPlanSettingsArgsOutput() QueueReservationPlanSettingsArgsOutput {
	return o
}

func (o QueueReservationPlanSettingsArgsOutput) ToQueueReservationPlanSettingsArgsOutputWithContext(ctx context.Context) QueueReservationPlanSettingsArgsOutput {
	return o
}

func (o QueueReservationPlanSettingsArgsOutput) ToQueueReservationPlanSettingsArgsPtrOutput() QueueReservationPlanSettingsArgsPtrOutput {
	return o.ToQueueReservationPlanSettingsArgsPtrOutputWithContext(context.Background())
}

func (o QueueReservationPlanSettingsArgsOutput) ToQueueReservationPlanSettingsArgsPtrOutputWithContext(ctx context.Context) QueueReservationPlanSettingsArgsPtrOutput {
	return o.ApplyT(func(v QueueReservationPlanSettingsArgs) *QueueReservationPlanSettingsArgs {
		return &v
	}).(QueueReservationPlanSettingsArgsPtrOutput)
}

// The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
func (o QueueReservationPlanSettingsArgsOutput) Commitment() pulumi.StringOutput {
	return o.ApplyT(func(v QueueReservationPlanSettingsArgs) string { return v.Commitment }).(pulumi.StringOutput)
}

// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
func (o QueueReservationPlanSettingsArgsOutput) RenewalType() pulumi.StringOutput {
	return o.ApplyT(func(v QueueReservationPlanSettingsArgs) string { return v.RenewalType }).(pulumi.StringOutput)
}

// Specifies the number of reserved transcode slots (RTS) for queue.
func (o QueueReservationPlanSettingsArgsOutput) ReservedSlots() pulumi.IntOutput {
	return o.ApplyT(func(v QueueReservationPlanSettingsArgs) int { return v.ReservedSlots }).(pulumi.IntOutput)
}

type QueueReservationPlanSettingsArgsPtrOutput struct{ *pulumi.OutputState }

func (QueueReservationPlanSettingsArgsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueReservationPlanSettingsArgs)(nil)).Elem()
}

func (o QueueReservationPlanSettingsArgsPtrOutput) ToQueueReservationPlanSettingsArgsPtrOutput() QueueReservationPlanSettingsArgsPtrOutput {
	return o
}

func (o QueueReservationPlanSettingsArgsPtrOutput) ToQueueReservationPlanSettingsArgsPtrOutputWithContext(ctx context.Context) QueueReservationPlanSettingsArgsPtrOutput {
	return o
}

func (o QueueReservationPlanSettingsArgsPtrOutput) Elem() QueueReservationPlanSettingsArgsOutput {
	return o.ApplyT(func(v *QueueReservationPlanSettingsArgs) QueueReservationPlanSettingsArgs { return *v }).(QueueReservationPlanSettingsArgsOutput)
}

// The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
func (o QueueReservationPlanSettingsArgsPtrOutput) Commitment() pulumi.StringOutput {
	return o.ApplyT(func(v QueueReservationPlanSettingsArgs) string { return v.Commitment }).(pulumi.StringOutput)
}

// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
func (o QueueReservationPlanSettingsArgsPtrOutput) RenewalType() pulumi.StringOutput {
	return o.ApplyT(func(v QueueReservationPlanSettingsArgs) string { return v.RenewalType }).(pulumi.StringOutput)
}

// Specifies the number of reserved transcode slots (RTS) for queue.
func (o QueueReservationPlanSettingsArgsPtrOutput) ReservedSlots() pulumi.IntOutput {
	return o.ApplyT(func(v QueueReservationPlanSettingsArgs) int { return v.ReservedSlots }).(pulumi.IntOutput)
}

type QueueReservationPlanSettingsState struct {
	// The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
	Commitment string `pulumi:"commitment"`
	// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
	RenewalType string `pulumi:"renewalType"`
	// Specifies the number of reserved transcode slots (RTS) for queue.
	ReservedSlots int `pulumi:"reservedSlots"`
}

type QueueReservationPlanSettingsStateInput interface {
	pulumi.Input

	ToQueueReservationPlanSettingsStateOutput() QueueReservationPlanSettingsStateOutput
	ToQueueReservationPlanSettingsStateOutputWithContext(context.Context) QueueReservationPlanSettingsStateOutput
}

type QueueReservationPlanSettingsStateArgs struct {
	// The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
	Commitment pulumi.StringInput `pulumi:"commitment"`
	// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
	RenewalType pulumi.StringInput `pulumi:"renewalType"`
	// Specifies the number of reserved transcode slots (RTS) for queue.
	ReservedSlots pulumi.IntInput `pulumi:"reservedSlots"`
}

func (QueueReservationPlanSettingsStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueReservationPlanSettingsState)(nil)).Elem()
}

func (i QueueReservationPlanSettingsStateArgs) ToQueueReservationPlanSettingsStateOutput() QueueReservationPlanSettingsStateOutput {
	return i.ToQueueReservationPlanSettingsStateOutputWithContext(context.Background())
}

func (i QueueReservationPlanSettingsStateArgs) ToQueueReservationPlanSettingsStateOutputWithContext(ctx context.Context) QueueReservationPlanSettingsStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueReservationPlanSettingsStateOutput)
}

type QueueReservationPlanSettingsStateOutput struct{ *pulumi.OutputState }

func (QueueReservationPlanSettingsStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueReservationPlanSettingsState)(nil)).Elem()
}

func (o QueueReservationPlanSettingsStateOutput) ToQueueReservationPlanSettingsStateOutput() QueueReservationPlanSettingsStateOutput {
	return o
}

func (o QueueReservationPlanSettingsStateOutput) ToQueueReservationPlanSettingsStateOutputWithContext(ctx context.Context) QueueReservationPlanSettingsStateOutput {
	return o
}

// The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
func (o QueueReservationPlanSettingsStateOutput) Commitment() pulumi.StringOutput {
	return o.ApplyT(func(v QueueReservationPlanSettingsState) string { return v.Commitment }).(pulumi.StringOutput)
}

// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
func (o QueueReservationPlanSettingsStateOutput) RenewalType() pulumi.StringOutput {
	return o.ApplyT(func(v QueueReservationPlanSettingsState) string { return v.RenewalType }).(pulumi.StringOutput)
}

// Specifies the number of reserved transcode slots (RTS) for queue.
func (o QueueReservationPlanSettingsStateOutput) ReservedSlots() pulumi.IntOutput {
	return o.ApplyT(func(v QueueReservationPlanSettingsState) int { return v.ReservedSlots }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueReservationPlanSettingsOutput{})
	pulumi.RegisterOutputType(QueueReservationPlanSettingsPtrOutput{})
	pulumi.RegisterOutputType(QueueReservationPlanSettingsArgsOutput{})
	pulumi.RegisterOutputType(QueueReservationPlanSettingsArgsPtrOutput{})
	pulumi.RegisterOutputType(QueueReservationPlanSettingsStateOutput{})
}
