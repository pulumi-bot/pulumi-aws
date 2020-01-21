// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package TopicRuleCloudwatchMetric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type TopicRuleCloudwatchMetric struct {
	// The CloudWatch metric name.
	MetricName string `pulumi:"metricName"`
	// The CloudWatch metric namespace name.
	MetricNamespace string `pulumi:"metricNamespace"`
	// An optional Unix timestamp (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp).
	MetricTimestamp *string `pulumi:"metricTimestamp"`
	// The metric unit (supported units can be found here: http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit)
	MetricUnit string `pulumi:"metricUnit"`
	// The CloudWatch metric value.
	MetricValue string `pulumi:"metricValue"`
	// The ARN of the IAM role that grants access.
	RoleArn string `pulumi:"roleArn"`
}

type TopicRuleCloudwatchMetricInput interface {
	pulumi.Input

	ToTopicRuleCloudwatchMetricOutput() TopicRuleCloudwatchMetricOutput
	ToTopicRuleCloudwatchMetricOutputWithContext(context.Context) TopicRuleCloudwatchMetricOutput
}

type TopicRuleCloudwatchMetricArgs struct {
	// The CloudWatch metric name.
	MetricName pulumi.StringInput `pulumi:"metricName"`
	// The CloudWatch metric namespace name.
	MetricNamespace pulumi.StringInput `pulumi:"metricNamespace"`
	// An optional Unix timestamp (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp).
	MetricTimestamp pulumi.StringPtrInput `pulumi:"metricTimestamp"`
	// The metric unit (supported units can be found here: http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit)
	MetricUnit pulumi.StringInput `pulumi:"metricUnit"`
	// The CloudWatch metric value.
	MetricValue pulumi.StringInput `pulumi:"metricValue"`
	// The ARN of the IAM role that grants access.
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
}

func (TopicRuleCloudwatchMetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicRuleCloudwatchMetric)(nil)).Elem()
}

func (i TopicRuleCloudwatchMetricArgs) ToTopicRuleCloudwatchMetricOutput() TopicRuleCloudwatchMetricOutput {
	return i.ToTopicRuleCloudwatchMetricOutputWithContext(context.Background())
}

func (i TopicRuleCloudwatchMetricArgs) ToTopicRuleCloudwatchMetricOutputWithContext(ctx context.Context) TopicRuleCloudwatchMetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicRuleCloudwatchMetricOutput)
}

func (i TopicRuleCloudwatchMetricArgs) ToTopicRuleCloudwatchMetricPtrOutput() TopicRuleCloudwatchMetricPtrOutput {
	return i.ToTopicRuleCloudwatchMetricPtrOutputWithContext(context.Background())
}

func (i TopicRuleCloudwatchMetricArgs) ToTopicRuleCloudwatchMetricPtrOutputWithContext(ctx context.Context) TopicRuleCloudwatchMetricPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicRuleCloudwatchMetricOutput).ToTopicRuleCloudwatchMetricPtrOutputWithContext(ctx)
}

type TopicRuleCloudwatchMetricPtrInput interface {
	pulumi.Input

	ToTopicRuleCloudwatchMetricPtrOutput() TopicRuleCloudwatchMetricPtrOutput
	ToTopicRuleCloudwatchMetricPtrOutputWithContext(context.Context) TopicRuleCloudwatchMetricPtrOutput
}

type topicRuleCloudwatchMetricPtrType TopicRuleCloudwatchMetricArgs

func TopicRuleCloudwatchMetricPtr(v *TopicRuleCloudwatchMetricArgs) TopicRuleCloudwatchMetricPtrInput {	return (*topicRuleCloudwatchMetricPtrType)(v)
}

func (*topicRuleCloudwatchMetricPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicRuleCloudwatchMetric)(nil)).Elem()
}

func (i *topicRuleCloudwatchMetricPtrType) ToTopicRuleCloudwatchMetricPtrOutput() TopicRuleCloudwatchMetricPtrOutput {
	return i.ToTopicRuleCloudwatchMetricPtrOutputWithContext(context.Background())
}

func (i *topicRuleCloudwatchMetricPtrType) ToTopicRuleCloudwatchMetricPtrOutputWithContext(ctx context.Context) TopicRuleCloudwatchMetricPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicRuleCloudwatchMetricPtrOutput)
}

type TopicRuleCloudwatchMetricOutput struct { *pulumi.OutputState }

func (TopicRuleCloudwatchMetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicRuleCloudwatchMetric)(nil)).Elem()
}

func (o TopicRuleCloudwatchMetricOutput) ToTopicRuleCloudwatchMetricOutput() TopicRuleCloudwatchMetricOutput {
	return o
}

func (o TopicRuleCloudwatchMetricOutput) ToTopicRuleCloudwatchMetricOutputWithContext(ctx context.Context) TopicRuleCloudwatchMetricOutput {
	return o
}

func (o TopicRuleCloudwatchMetricOutput) ToTopicRuleCloudwatchMetricPtrOutput() TopicRuleCloudwatchMetricPtrOutput {
	return o.ToTopicRuleCloudwatchMetricPtrOutputWithContext(context.Background())
}

func (o TopicRuleCloudwatchMetricOutput) ToTopicRuleCloudwatchMetricPtrOutputWithContext(ctx context.Context) TopicRuleCloudwatchMetricPtrOutput {
	return o.ApplyT(func(v TopicRuleCloudwatchMetric) *TopicRuleCloudwatchMetric {
		return &v
	}).(TopicRuleCloudwatchMetricPtrOutput)
}
// The CloudWatch metric name.
func (o TopicRuleCloudwatchMetricOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchMetric) string { return v.MetricName }).(pulumi.StringOutput)
}

// The CloudWatch metric namespace name.
func (o TopicRuleCloudwatchMetricOutput) MetricNamespace() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchMetric) string { return v.MetricNamespace }).(pulumi.StringOutput)
}

// An optional Unix timestamp (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp).
func (o TopicRuleCloudwatchMetricOutput) MetricTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchMetric) *string { return v.MetricTimestamp }).(pulumi.StringPtrOutput)
}

// The metric unit (supported units can be found here: http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit)
func (o TopicRuleCloudwatchMetricOutput) MetricUnit() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchMetric) string { return v.MetricUnit }).(pulumi.StringOutput)
}

// The CloudWatch metric value.
func (o TopicRuleCloudwatchMetricOutput) MetricValue() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchMetric) string { return v.MetricValue }).(pulumi.StringOutput)
}

// The ARN of the IAM role that grants access.
func (o TopicRuleCloudwatchMetricOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchMetric) string { return v.RoleArn }).(pulumi.StringOutput)
}

type TopicRuleCloudwatchMetricPtrOutput struct { *pulumi.OutputState}

func (TopicRuleCloudwatchMetricPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicRuleCloudwatchMetric)(nil)).Elem()
}

func (o TopicRuleCloudwatchMetricPtrOutput) ToTopicRuleCloudwatchMetricPtrOutput() TopicRuleCloudwatchMetricPtrOutput {
	return o
}

func (o TopicRuleCloudwatchMetricPtrOutput) ToTopicRuleCloudwatchMetricPtrOutputWithContext(ctx context.Context) TopicRuleCloudwatchMetricPtrOutput {
	return o
}

func (o TopicRuleCloudwatchMetricPtrOutput) Elem() TopicRuleCloudwatchMetricOutput {
	return o.ApplyT(func (v *TopicRuleCloudwatchMetric) TopicRuleCloudwatchMetric { return *v }).(TopicRuleCloudwatchMetricOutput)
}

// The CloudWatch metric name.
func (o TopicRuleCloudwatchMetricPtrOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchMetric) string { return v.MetricName }).(pulumi.StringOutput)
}

// The CloudWatch metric namespace name.
func (o TopicRuleCloudwatchMetricPtrOutput) MetricNamespace() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchMetric) string { return v.MetricNamespace }).(pulumi.StringOutput)
}

// An optional Unix timestamp (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp).
func (o TopicRuleCloudwatchMetricPtrOutput) MetricTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchMetric) *string { return v.MetricTimestamp }).(pulumi.StringPtrOutput)
}

// The metric unit (supported units can be found here: http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit)
func (o TopicRuleCloudwatchMetricPtrOutput) MetricUnit() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchMetric) string { return v.MetricUnit }).(pulumi.StringOutput)
}

// The CloudWatch metric value.
func (o TopicRuleCloudwatchMetricPtrOutput) MetricValue() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchMetric) string { return v.MetricValue }).(pulumi.StringOutput)
}

// The ARN of the IAM role that grants access.
func (o TopicRuleCloudwatchMetricPtrOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func (v TopicRuleCloudwatchMetric) string { return v.RoleArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicRuleCloudwatchMetricOutput{})
	pulumi.RegisterOutputType(TopicRuleCloudwatchMetricPtrOutput{})
}
