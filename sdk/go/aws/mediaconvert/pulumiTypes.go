// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mediaconvert

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type QueueReservationPlanSettings struct {
	// The length of the term of your reserved queue pricing plan commitment. Valid value is `ONE_YEAR`.
	Commitment string `pulumi:"commitment"`
	// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
	RenewalType string `pulumi:"renewalType"`
	// Specifies the number of reserved transcode slots (RTS) for queue.
	ReservedSlots int `pulumi:"reservedSlots"`
}

// QueueReservationPlanSettingsInput is an input type that accepts QueueReservationPlanSettingsArgs and QueueReservationPlanSettingsOutput values.
// You can construct a concrete instance of `QueueReservationPlanSettingsInput` via:
//
//          QueueReservationPlanSettingsArgs{...}
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
	return pulumi.ToOutputWithContext(ctx, i).(QueueReservationPlanSettingsOutput).ToQueueReservationPlanSettingsPtrOutput()
}

// QueueReservationPlanSettingsPtrInput is an input type that accepts QueueReservationPlanSettingsArgs, QueueReservationPlanSettingsPtr and QueueReservationPlanSettingsPtrOutput values.
// You can construct a concrete instance of `QueueReservationPlanSettingsPtrInput` via:
//
//          QueueReservationPlanSettingsArgs{...}
//
//  or:
//
//          nil
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
	return pulumi.ToOutputWithContext(ctx, i).(QueueReservationPlanSettingsOutput).ToQueueReservationPlanSettingsPtrOutput()
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
func (o QueueReservationPlanSettingsPtrOutput) Commitment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueReservationPlanSettings) *string {
		if v == nil {
			return nil
		}
		return &v.Commitment
	}).(pulumi.StringPtrOutput)
}

// Specifies whether the term of your reserved queue pricing plan. Valid values are `AUTO_RENEW` or `EXPIRE`.
func (o QueueReservationPlanSettingsPtrOutput) RenewalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueReservationPlanSettings) *string {
		if v == nil {
			return nil
		}
		return &v.RenewalType
	}).(pulumi.StringPtrOutput)
}

// Specifies the number of reserved transcode slots (RTS) for queue.
func (o QueueReservationPlanSettingsPtrOutput) ReservedSlots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QueueReservationPlanSettings) *int {
		if v == nil {
			return nil
		}
		return &v.ReservedSlots
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(QueueReservationPlanSettingsOutput{})
	pulumi.RegisterOutputType(QueueReservationPlanSettingsPtrOutput{})
}
