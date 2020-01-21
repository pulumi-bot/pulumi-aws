// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/kinesis/FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameter"
)

type FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessor struct {
	// Array of processor parameters. More details are given below
	Parameters []kinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameter.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameter `pulumi:"parameters"`
	// The type of processor. Valid Values: `Lambda`
	Type string `pulumi:"type"`
}

type FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorInput interface {
	pulumi.Input

	ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput() FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput
	ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutputWithContext(context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput
}

type FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs struct {
	// Array of processor parameters. More details are given below
	Parameters kinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameter.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArrayInput `pulumi:"parameters"`
	// The type of processor. Valid Values: `Lambda`
	Type pulumi.StringInput `pulumi:"type"`
}

func (FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessor)(nil)).Elem()
}

func (i FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs) ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput() FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput {
	return i.ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutputWithContext(context.Background())
}

func (i FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArgs) ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutputWithContext(ctx context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput)
}

type FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayInput interface {
	pulumi.Input

	ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput() FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput
	ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutputWithContext(context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput
}

type FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArray []FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorInput

func (FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessor)(nil)).Elem()
}

func (i FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArray) ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput() FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput {
	return i.ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutputWithContext(context.Background())
}

func (i FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArray) ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutputWithContext(ctx context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput)
}

type FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput struct { *pulumi.OutputState }

func (FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessor)(nil)).Elem()
}

func (o FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput) ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput() FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput {
	return o
}

func (o FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput) ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutputWithContext(ctx context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput {
	return o
}

// Array of processor parameters. More details are given below
func (o FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput) Parameters() kinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameter.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArrayOutput {
	return o.ApplyT(func (v FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessor) []kinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameter.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameter { return v.Parameters }).(kinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameter.FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorParameterArrayOutput)
}

// The type of processor. Valid Values: `Lambda`
func (o FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func (v FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessor) string { return v.Type }).(pulumi.StringOutput)
}

type FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput struct { *pulumi.OutputState}

func (FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessor)(nil)).Elem()
}

func (o FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput) ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput() FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput {
	return o
}

func (o FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput) ToFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutputWithContext(ctx context.Context) FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput {
	return o
}

func (o FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput) Index(i pulumi.IntInput) FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessor {
		return vs[0].([]FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessor)[vs[1].(int)]
	}).(FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput)
}

func init() {
	pulumi.RegisterOutputType(FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorOutput{})
	pulumi.RegisterOutputType(FirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorArrayOutput{})
}
