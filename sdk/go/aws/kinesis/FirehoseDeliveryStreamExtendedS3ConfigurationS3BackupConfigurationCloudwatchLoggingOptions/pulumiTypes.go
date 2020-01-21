// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions struct {
	// Defaults to `true`. Set it to `false` if you want to disable format conversion while preserving the configuration details.
	Enabled *bool `pulumi:"enabled"`
	// The CloudWatch group name for logging. This value is required if `enabled` is true.
	LogGroupName *string `pulumi:"logGroupName"`
	// The CloudWatch log stream name for logging. This value is required if `enabled` is true.
	LogStreamName *string `pulumi:"logStreamName"`
}

type FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsInput interface {
	pulumi.Input

	ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput() FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput
	ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutputWithContext(context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput
}

type FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsArgs struct {
	// Defaults to `true`. Set it to `false` if you want to disable format conversion while preserving the configuration details.
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
	// The CloudWatch group name for logging. This value is required if `enabled` is true.
	LogGroupName pulumi.StringPtrInput `pulumi:"logGroupName"`
	// The CloudWatch log stream name for logging. This value is required if `enabled` is true.
	LogStreamName pulumi.StringPtrInput `pulumi:"logStreamName"`
}

func (FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions)(nil)).Elem()
}

func (i FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsArgs) ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput() FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput {
	return i.ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutputWithContext(context.Background())
}

func (i FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsArgs) ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutputWithContext(ctx context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput)
}

func (i FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsArgs) ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput() FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput {
	return i.ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutputWithContext(context.Background())
}

func (i FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsArgs) ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutputWithContext(ctx context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput).ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutputWithContext(ctx)
}

type FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrInput interface {
	pulumi.Input

	ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput() FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput
	ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutputWithContext(context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput
}

type firehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrType FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsArgs

func FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtr(v *FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsArgs) FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrInput {	return (*firehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrType)(v)
}

func (*firehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions)(nil)).Elem()
}

func (i *firehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrType) ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput() FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput {
	return i.ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutputWithContext(context.Background())
}

func (i *firehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrType) ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutputWithContext(ctx context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput)
}

type FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput struct { *pulumi.OutputState }

func (FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions)(nil)).Elem()
}

func (o FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput) ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput() FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput {
	return o
}

func (o FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput) ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutputWithContext(ctx context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput {
	return o
}

func (o FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput) ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput() FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput {
	return o.ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutputWithContext(context.Background())
}

func (o FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput) ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutputWithContext(ctx context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput {
	return o.ApplyT(func(v FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions) *FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions {
		return &v
	}).(FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput)
}
// Defaults to `true`. Set it to `false` if you want to disable format conversion while preserving the configuration details.
func (o FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The CloudWatch group name for logging. This value is required if `enabled` is true.
func (o FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput) LogGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions) *string { return v.LogGroupName }).(pulumi.StringPtrOutput)
}

// The CloudWatch log stream name for logging. This value is required if `enabled` is true.
func (o FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput) LogStreamName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions) *string { return v.LogStreamName }).(pulumi.StringPtrOutput)
}

type FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput struct { *pulumi.OutputState}

func (FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions)(nil)).Elem()
}

func (o FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput) ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput() FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput {
	return o
}

func (o FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput) ToFirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutputWithContext(ctx context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput {
	return o
}

func (o FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput) Elem() FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput {
	return o.ApplyT(func (v *FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions) FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions { return *v }).(FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput)
}

// Defaults to `true`. Set it to `false` if you want to disable format conversion while preserving the configuration details.
func (o FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The CloudWatch group name for logging. This value is required if `enabled` is true.
func (o FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput) LogGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions) *string { return v.LogGroupName }).(pulumi.StringPtrOutput)
}

// The CloudWatch log stream name for logging. This value is required if `enabled` is true.
func (o FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput) LogStreamName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptions) *string { return v.LogStreamName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsOutput{})
	pulumi.RegisterOutputType(FirehoseDeliveryStreamExtendedS3ConfigurationS3BackupConfigurationCloudwatchLoggingOptionsPtrOutput{})
}
