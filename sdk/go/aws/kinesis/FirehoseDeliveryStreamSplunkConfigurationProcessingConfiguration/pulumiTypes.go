// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/kinesis/FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor"
	"https:/github.com/pulumi/pulumi-aws/kinesis/FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessorParameter"
)

type FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration struct {
	// Defaults to `true`. Set it to `false` if you want to disable format conversion while preserving the configuration details.
	Enabled *bool `pulumi:"enabled"`
	// Array of data processors. More details are given below
	Processors []kinesisFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor.FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor `pulumi:"processors"`
}

type FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationInput interface {
	pulumi.Input

	ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput() FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput
	ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutputWithContext(context.Context) FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput
}

type FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationArgs struct {
	// Defaults to `true`. Set it to `false` if you want to disable format conversion while preserving the configuration details.
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
	// Array of data processors. More details are given below
	Processors kinesisFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor.FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessorArrayInput `pulumi:"processors"`
}

func (FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration)(nil)).Elem()
}

func (i FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationArgs) ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput() FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput {
	return i.ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutputWithContext(context.Background())
}

func (i FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationArgs) ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutputWithContext(ctx context.Context) FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput)
}

func (i FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationArgs) ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput() FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput {
	return i.ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutputWithContext(context.Background())
}

func (i FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationArgs) ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutputWithContext(ctx context.Context) FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput).ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutputWithContext(ctx)
}

type FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrInput interface {
	pulumi.Input

	ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput() FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput
	ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutputWithContext(context.Context) FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput
}

type firehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrType FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationArgs

func FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtr(v *FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationArgs) FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrInput {	return (*firehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrType)(v)
}

func (*firehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration)(nil)).Elem()
}

func (i *firehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrType) ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput() FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput {
	return i.ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutputWithContext(context.Background())
}

func (i *firehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrType) ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutputWithContext(ctx context.Context) FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput)
}

type FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput struct { *pulumi.OutputState }

func (FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration)(nil)).Elem()
}

func (o FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput) ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput() FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput {
	return o
}

func (o FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput) ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutputWithContext(ctx context.Context) FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput {
	return o
}

func (o FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput) ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput() FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput {
	return o.ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutputWithContext(context.Background())
}

func (o FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput) ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutputWithContext(ctx context.Context) FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput {
	return o.ApplyT(func(v FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration) *FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration {
		return &v
	}).(FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput)
}
// Defaults to `true`. Set it to `false` if you want to disable format conversion while preserving the configuration details.
func (o FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Array of data processors. More details are given below
func (o FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput) Processors() kinesisFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor.FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessorArrayOutput {
	return o.ApplyT(func (v FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration) []kinesisFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor.FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor { return v.Processors }).(kinesisFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor.FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessorArrayOutput)
}

type FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput struct { *pulumi.OutputState}

func (FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration)(nil)).Elem()
}

func (o FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput) ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput() FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput {
	return o
}

func (o FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput) ToFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutputWithContext(ctx context.Context) FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput {
	return o
}

func (o FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput) Elem() FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput {
	return o.ApplyT(func (v *FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration) FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration { return *v }).(FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput)
}

// Defaults to `true`. Set it to `false` if you want to disable format conversion while preserving the configuration details.
func (o FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Array of data processors. More details are given below
func (o FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput) Processors() kinesisFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor.FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessorArrayOutput {
	return o.ApplyT(func (v FirehoseDeliveryStreamSplunkConfigurationProcessingConfiguration) []kinesisFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor.FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor { return v.Processors }).(kinesisFirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessor.FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationProcessorArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationOutput{})
	pulumi.RegisterOutputType(FirehoseDeliveryStreamSplunkConfigurationProcessingConfigurationPtrOutput{})
}
