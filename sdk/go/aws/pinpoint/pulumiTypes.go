// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AppCampaignHook struct {
	// Lambda function name or ARN to be called for delivery. Conflicts with `webUrl`
	LambdaFunctionName *string `pulumi:"lambdaFunctionName"`
	// What mode Lambda should be invoked in. Valid values for this parameter are `DELIVERY`, `FILTER`.
	Mode *string `pulumi:"mode"`
	// Web URL to call for hook. If the URL has authentication specified it will be added as authentication to the request. Conflicts with `lambdaFunctionName`
	WebUrl *string `pulumi:"webUrl"`
}

// AppCampaignHookInput is an input type that accepts AppCampaignHookArgs and AppCampaignHookOutput values.
// You can construct a concrete instance of `AppCampaignHookInput` via:
//
//          AppCampaignHookArgs{...}
type AppCampaignHookInput interface {
	pulumi.Input

	ToAppCampaignHookOutput() AppCampaignHookOutput
	ToAppCampaignHookOutputWithContext(context.Context) AppCampaignHookOutput
}

type AppCampaignHookArgs struct {
	// Lambda function name or ARN to be called for delivery. Conflicts with `webUrl`
	LambdaFunctionName pulumi.StringPtrInput `pulumi:"lambdaFunctionName"`
	// What mode Lambda should be invoked in. Valid values for this parameter are `DELIVERY`, `FILTER`.
	Mode pulumi.StringPtrInput `pulumi:"mode"`
	// Web URL to call for hook. If the URL has authentication specified it will be added as authentication to the request. Conflicts with `lambdaFunctionName`
	WebUrl pulumi.StringPtrInput `pulumi:"webUrl"`
}

func (AppCampaignHookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppCampaignHook)(nil)).Elem()
}

func (i AppCampaignHookArgs) ToAppCampaignHookOutput() AppCampaignHookOutput {
	return i.ToAppCampaignHookOutputWithContext(context.Background())
}

func (i AppCampaignHookArgs) ToAppCampaignHookOutputWithContext(ctx context.Context) AppCampaignHookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppCampaignHookOutput)
}

func (i AppCampaignHookArgs) ToAppCampaignHookPtrOutput() AppCampaignHookPtrOutput {
	return i.ToAppCampaignHookPtrOutputWithContext(context.Background())
}

func (i AppCampaignHookArgs) ToAppCampaignHookPtrOutputWithContext(ctx context.Context) AppCampaignHookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppCampaignHookOutput).ToAppCampaignHookPtrOutput()
}

// AppCampaignHookPtrInput is an input type that accepts AppCampaignHookArgs, AppCampaignHookPtr and AppCampaignHookPtrOutput values.
// You can construct a concrete instance of `AppCampaignHookPtrInput` via:
//
//          AppCampaignHookArgs{...}
//
//  or:
//
//          nil
type AppCampaignHookPtrInput interface {
	pulumi.Input

	ToAppCampaignHookPtrOutput() AppCampaignHookPtrOutput
	ToAppCampaignHookPtrOutputWithContext(context.Context) AppCampaignHookPtrOutput
}

type appCampaignHookPtrType AppCampaignHookArgs

func AppCampaignHookPtr(v *AppCampaignHookArgs) AppCampaignHookPtrInput {
	return (*appCampaignHookPtrType)(v)
}

func (*appCampaignHookPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppCampaignHook)(nil)).Elem()
}

func (i *appCampaignHookPtrType) ToAppCampaignHookPtrOutput() AppCampaignHookPtrOutput {
	return i.ToAppCampaignHookPtrOutputWithContext(context.Background())
}

func (i *appCampaignHookPtrType) ToAppCampaignHookPtrOutputWithContext(ctx context.Context) AppCampaignHookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppCampaignHookOutput).ToAppCampaignHookPtrOutput()
}

type AppCampaignHookOutput struct{ *pulumi.OutputState }

func (AppCampaignHookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppCampaignHook)(nil)).Elem()
}

func (o AppCampaignHookOutput) ToAppCampaignHookOutput() AppCampaignHookOutput {
	return o
}

func (o AppCampaignHookOutput) ToAppCampaignHookOutputWithContext(ctx context.Context) AppCampaignHookOutput {
	return o
}

func (o AppCampaignHookOutput) ToAppCampaignHookPtrOutput() AppCampaignHookPtrOutput {
	return o.ToAppCampaignHookPtrOutputWithContext(context.Background())
}

func (o AppCampaignHookOutput) ToAppCampaignHookPtrOutputWithContext(ctx context.Context) AppCampaignHookPtrOutput {
	return o.ApplyT(func(v AppCampaignHook) *AppCampaignHook {
		return &v
	}).(AppCampaignHookPtrOutput)
}

// Lambda function name or ARN to be called for delivery. Conflicts with `webUrl`
func (o AppCampaignHookOutput) LambdaFunctionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppCampaignHook) *string { return v.LambdaFunctionName }).(pulumi.StringPtrOutput)
}

// What mode Lambda should be invoked in. Valid values for this parameter are `DELIVERY`, `FILTER`.
func (o AppCampaignHookOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppCampaignHook) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

// Web URL to call for hook. If the URL has authentication specified it will be added as authentication to the request. Conflicts with `lambdaFunctionName`
func (o AppCampaignHookOutput) WebUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppCampaignHook) *string { return v.WebUrl }).(pulumi.StringPtrOutput)
}

type AppCampaignHookPtrOutput struct{ *pulumi.OutputState }

func (AppCampaignHookPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppCampaignHook)(nil)).Elem()
}

func (o AppCampaignHookPtrOutput) ToAppCampaignHookPtrOutput() AppCampaignHookPtrOutput {
	return o
}

func (o AppCampaignHookPtrOutput) ToAppCampaignHookPtrOutputWithContext(ctx context.Context) AppCampaignHookPtrOutput {
	return o
}

func (o AppCampaignHookPtrOutput) Elem() AppCampaignHookOutput {
	return o.ApplyT(func(v *AppCampaignHook) AppCampaignHook { return *v }).(AppCampaignHookOutput)
}

// Lambda function name or ARN to be called for delivery. Conflicts with `webUrl`
func (o AppCampaignHookPtrOutput) LambdaFunctionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppCampaignHook) *string {
		if v == nil {
			return nil
		}
		return v.LambdaFunctionName
	}).(pulumi.StringPtrOutput)
}

// What mode Lambda should be invoked in. Valid values for this parameter are `DELIVERY`, `FILTER`.
func (o AppCampaignHookPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppCampaignHook) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

// Web URL to call for hook. If the URL has authentication specified it will be added as authentication to the request. Conflicts with `lambdaFunctionName`
func (o AppCampaignHookPtrOutput) WebUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppCampaignHook) *string {
		if v == nil {
			return nil
		}
		return v.WebUrl
	}).(pulumi.StringPtrOutput)
}

type AppLimits struct {
	// The maximum number of messages that the campaign can send daily.
	Daily *int `pulumi:"daily"`
	// The length of time (in seconds) that the campaign can run before it ends and message deliveries stop. This duration begins at the scheduled start time for the campaign. The minimum value is 60.
	MaximumDuration *int `pulumi:"maximumDuration"`
	// The number of messages that the campaign can send per second. The minimum value is 50, and the maximum is 20000.
	MessagesPerSecond *int `pulumi:"messagesPerSecond"`
	// The maximum total number of messages that the campaign can send.
	Total *int `pulumi:"total"`
}

// AppLimitsInput is an input type that accepts AppLimitsArgs and AppLimitsOutput values.
// You can construct a concrete instance of `AppLimitsInput` via:
//
//          AppLimitsArgs{...}
type AppLimitsInput interface {
	pulumi.Input

	ToAppLimitsOutput() AppLimitsOutput
	ToAppLimitsOutputWithContext(context.Context) AppLimitsOutput
}

type AppLimitsArgs struct {
	// The maximum number of messages that the campaign can send daily.
	Daily pulumi.IntPtrInput `pulumi:"daily"`
	// The length of time (in seconds) that the campaign can run before it ends and message deliveries stop. This duration begins at the scheduled start time for the campaign. The minimum value is 60.
	MaximumDuration pulumi.IntPtrInput `pulumi:"maximumDuration"`
	// The number of messages that the campaign can send per second. The minimum value is 50, and the maximum is 20000.
	MessagesPerSecond pulumi.IntPtrInput `pulumi:"messagesPerSecond"`
	// The maximum total number of messages that the campaign can send.
	Total pulumi.IntPtrInput `pulumi:"total"`
}

func (AppLimitsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppLimits)(nil)).Elem()
}

func (i AppLimitsArgs) ToAppLimitsOutput() AppLimitsOutput {
	return i.ToAppLimitsOutputWithContext(context.Background())
}

func (i AppLimitsArgs) ToAppLimitsOutputWithContext(ctx context.Context) AppLimitsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppLimitsOutput)
}

func (i AppLimitsArgs) ToAppLimitsPtrOutput() AppLimitsPtrOutput {
	return i.ToAppLimitsPtrOutputWithContext(context.Background())
}

func (i AppLimitsArgs) ToAppLimitsPtrOutputWithContext(ctx context.Context) AppLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppLimitsOutput).ToAppLimitsPtrOutput()
}

// AppLimitsPtrInput is an input type that accepts AppLimitsArgs, AppLimitsPtr and AppLimitsPtrOutput values.
// You can construct a concrete instance of `AppLimitsPtrInput` via:
//
//          AppLimitsArgs{...}
//
//  or:
//
//          nil
type AppLimitsPtrInput interface {
	pulumi.Input

	ToAppLimitsPtrOutput() AppLimitsPtrOutput
	ToAppLimitsPtrOutputWithContext(context.Context) AppLimitsPtrOutput
}

type appLimitsPtrType AppLimitsArgs

func AppLimitsPtr(v *AppLimitsArgs) AppLimitsPtrInput {
	return (*appLimitsPtrType)(v)
}

func (*appLimitsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppLimits)(nil)).Elem()
}

func (i *appLimitsPtrType) ToAppLimitsPtrOutput() AppLimitsPtrOutput {
	return i.ToAppLimitsPtrOutputWithContext(context.Background())
}

func (i *appLimitsPtrType) ToAppLimitsPtrOutputWithContext(ctx context.Context) AppLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppLimitsOutput).ToAppLimitsPtrOutput()
}

type AppLimitsOutput struct{ *pulumi.OutputState }

func (AppLimitsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppLimits)(nil)).Elem()
}

func (o AppLimitsOutput) ToAppLimitsOutput() AppLimitsOutput {
	return o
}

func (o AppLimitsOutput) ToAppLimitsOutputWithContext(ctx context.Context) AppLimitsOutput {
	return o
}

func (o AppLimitsOutput) ToAppLimitsPtrOutput() AppLimitsPtrOutput {
	return o.ToAppLimitsPtrOutputWithContext(context.Background())
}

func (o AppLimitsOutput) ToAppLimitsPtrOutputWithContext(ctx context.Context) AppLimitsPtrOutput {
	return o.ApplyT(func(v AppLimits) *AppLimits {
		return &v
	}).(AppLimitsPtrOutput)
}

// The maximum number of messages that the campaign can send daily.
func (o AppLimitsOutput) Daily() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AppLimits) *int { return v.Daily }).(pulumi.IntPtrOutput)
}

// The length of time (in seconds) that the campaign can run before it ends and message deliveries stop. This duration begins at the scheduled start time for the campaign. The minimum value is 60.
func (o AppLimitsOutput) MaximumDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AppLimits) *int { return v.MaximumDuration }).(pulumi.IntPtrOutput)
}

// The number of messages that the campaign can send per second. The minimum value is 50, and the maximum is 20000.
func (o AppLimitsOutput) MessagesPerSecond() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AppLimits) *int { return v.MessagesPerSecond }).(pulumi.IntPtrOutput)
}

// The maximum total number of messages that the campaign can send.
func (o AppLimitsOutput) Total() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AppLimits) *int { return v.Total }).(pulumi.IntPtrOutput)
}

type AppLimitsPtrOutput struct{ *pulumi.OutputState }

func (AppLimitsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppLimits)(nil)).Elem()
}

func (o AppLimitsPtrOutput) ToAppLimitsPtrOutput() AppLimitsPtrOutput {
	return o
}

func (o AppLimitsPtrOutput) ToAppLimitsPtrOutputWithContext(ctx context.Context) AppLimitsPtrOutput {
	return o
}

func (o AppLimitsPtrOutput) Elem() AppLimitsOutput {
	return o.ApplyT(func(v *AppLimits) AppLimits { return *v }).(AppLimitsOutput)
}

// The maximum number of messages that the campaign can send daily.
func (o AppLimitsPtrOutput) Daily() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppLimits) *int {
		if v == nil {
			return nil
		}
		return v.Daily
	}).(pulumi.IntPtrOutput)
}

// The length of time (in seconds) that the campaign can run before it ends and message deliveries stop. This duration begins at the scheduled start time for the campaign. The minimum value is 60.
func (o AppLimitsPtrOutput) MaximumDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppLimits) *int {
		if v == nil {
			return nil
		}
		return v.MaximumDuration
	}).(pulumi.IntPtrOutput)
}

// The number of messages that the campaign can send per second. The minimum value is 50, and the maximum is 20000.
func (o AppLimitsPtrOutput) MessagesPerSecond() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppLimits) *int {
		if v == nil {
			return nil
		}
		return v.MessagesPerSecond
	}).(pulumi.IntPtrOutput)
}

// The maximum total number of messages that the campaign can send.
func (o AppLimitsPtrOutput) Total() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppLimits) *int {
		if v == nil {
			return nil
		}
		return v.Total
	}).(pulumi.IntPtrOutput)
}

type AppQuietTime struct {
	// The default end time for quiet time in ISO 8601 format. Required if `start` is set
	End *string `pulumi:"end"`
	// The default start time for quiet time in ISO 8601 format. Required if `end` is set
	Start *string `pulumi:"start"`
}

// AppQuietTimeInput is an input type that accepts AppQuietTimeArgs and AppQuietTimeOutput values.
// You can construct a concrete instance of `AppQuietTimeInput` via:
//
//          AppQuietTimeArgs{...}
type AppQuietTimeInput interface {
	pulumi.Input

	ToAppQuietTimeOutput() AppQuietTimeOutput
	ToAppQuietTimeOutputWithContext(context.Context) AppQuietTimeOutput
}

type AppQuietTimeArgs struct {
	// The default end time for quiet time in ISO 8601 format. Required if `start` is set
	End pulumi.StringPtrInput `pulumi:"end"`
	// The default start time for quiet time in ISO 8601 format. Required if `end` is set
	Start pulumi.StringPtrInput `pulumi:"start"`
}

func (AppQuietTimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppQuietTime)(nil)).Elem()
}

func (i AppQuietTimeArgs) ToAppQuietTimeOutput() AppQuietTimeOutput {
	return i.ToAppQuietTimeOutputWithContext(context.Background())
}

func (i AppQuietTimeArgs) ToAppQuietTimeOutputWithContext(ctx context.Context) AppQuietTimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppQuietTimeOutput)
}

func (i AppQuietTimeArgs) ToAppQuietTimePtrOutput() AppQuietTimePtrOutput {
	return i.ToAppQuietTimePtrOutputWithContext(context.Background())
}

func (i AppQuietTimeArgs) ToAppQuietTimePtrOutputWithContext(ctx context.Context) AppQuietTimePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppQuietTimeOutput).ToAppQuietTimePtrOutput()
}

// AppQuietTimePtrInput is an input type that accepts AppQuietTimeArgs, AppQuietTimePtr and AppQuietTimePtrOutput values.
// You can construct a concrete instance of `AppQuietTimePtrInput` via:
//
//          AppQuietTimeArgs{...}
//
//  or:
//
//          nil
type AppQuietTimePtrInput interface {
	pulumi.Input

	ToAppQuietTimePtrOutput() AppQuietTimePtrOutput
	ToAppQuietTimePtrOutputWithContext(context.Context) AppQuietTimePtrOutput
}

type appQuietTimePtrType AppQuietTimeArgs

func AppQuietTimePtr(v *AppQuietTimeArgs) AppQuietTimePtrInput {
	return (*appQuietTimePtrType)(v)
}

func (*appQuietTimePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppQuietTime)(nil)).Elem()
}

func (i *appQuietTimePtrType) ToAppQuietTimePtrOutput() AppQuietTimePtrOutput {
	return i.ToAppQuietTimePtrOutputWithContext(context.Background())
}

func (i *appQuietTimePtrType) ToAppQuietTimePtrOutputWithContext(ctx context.Context) AppQuietTimePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppQuietTimeOutput).ToAppQuietTimePtrOutput()
}

type AppQuietTimeOutput struct{ *pulumi.OutputState }

func (AppQuietTimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppQuietTime)(nil)).Elem()
}

func (o AppQuietTimeOutput) ToAppQuietTimeOutput() AppQuietTimeOutput {
	return o
}

func (o AppQuietTimeOutput) ToAppQuietTimeOutputWithContext(ctx context.Context) AppQuietTimeOutput {
	return o
}

func (o AppQuietTimeOutput) ToAppQuietTimePtrOutput() AppQuietTimePtrOutput {
	return o.ToAppQuietTimePtrOutputWithContext(context.Background())
}

func (o AppQuietTimeOutput) ToAppQuietTimePtrOutputWithContext(ctx context.Context) AppQuietTimePtrOutput {
	return o.ApplyT(func(v AppQuietTime) *AppQuietTime {
		return &v
	}).(AppQuietTimePtrOutput)
}

// The default end time for quiet time in ISO 8601 format. Required if `start` is set
func (o AppQuietTimeOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppQuietTime) *string { return v.End }).(pulumi.StringPtrOutput)
}

// The default start time for quiet time in ISO 8601 format. Required if `end` is set
func (o AppQuietTimeOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppQuietTime) *string { return v.Start }).(pulumi.StringPtrOutput)
}

type AppQuietTimePtrOutput struct{ *pulumi.OutputState }

func (AppQuietTimePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppQuietTime)(nil)).Elem()
}

func (o AppQuietTimePtrOutput) ToAppQuietTimePtrOutput() AppQuietTimePtrOutput {
	return o
}

func (o AppQuietTimePtrOutput) ToAppQuietTimePtrOutputWithContext(ctx context.Context) AppQuietTimePtrOutput {
	return o
}

func (o AppQuietTimePtrOutput) Elem() AppQuietTimeOutput {
	return o.ApplyT(func(v *AppQuietTime) AppQuietTime { return *v }).(AppQuietTimeOutput)
}

// The default end time for quiet time in ISO 8601 format. Required if `start` is set
func (o AppQuietTimePtrOutput) End() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppQuietTime) *string {
		if v == nil {
			return nil
		}
		return v.End
	}).(pulumi.StringPtrOutput)
}

// The default start time for quiet time in ISO 8601 format. Required if `end` is set
func (o AppQuietTimePtrOutput) Start() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppQuietTime) *string {
		if v == nil {
			return nil
		}
		return v.Start
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AppCampaignHookOutput{})
	pulumi.RegisterOutputType(AppCampaignHookPtrOutput{})
	pulumi.RegisterOutputType(AppLimitsOutput{})
	pulumi.RegisterOutputType(AppLimitsPtrOutput{})
	pulumi.RegisterOutputType(AppQuietTimeOutput{})
	pulumi.RegisterOutputType(AppQuietTimePtrOutput{})
}
