// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package BucketNotificationTopic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type BucketNotificationTopic struct {
	// Specifies [event](http://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html#notification-how-to-event-types-and-destinations) for which to send notifications.
	Events []string `pulumi:"events"`
	// Specifies object key name prefix.
	FilterPrefix *string `pulumi:"filterPrefix"`
	// Specifies object key name suffix.
	FilterSuffix *string `pulumi:"filterSuffix"`
	// Specifies unique identifier for each of the notification configurations.
	Id *string `pulumi:"id"`
	// Specifies Amazon SNS topic ARN.
	TopicArn string `pulumi:"topicArn"`
}

type BucketNotificationTopicInput interface {
	pulumi.Input

	ToBucketNotificationTopicOutput() BucketNotificationTopicOutput
	ToBucketNotificationTopicOutputWithContext(context.Context) BucketNotificationTopicOutput
}

type BucketNotificationTopicArgs struct {
	// Specifies [event](http://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html#notification-how-to-event-types-and-destinations) for which to send notifications.
	Events pulumi.StringArrayInput `pulumi:"events"`
	// Specifies object key name prefix.
	FilterPrefix pulumi.StringPtrInput `pulumi:"filterPrefix"`
	// Specifies object key name suffix.
	FilterSuffix pulumi.StringPtrInput `pulumi:"filterSuffix"`
	// Specifies unique identifier for each of the notification configurations.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Specifies Amazon SNS topic ARN.
	TopicArn pulumi.StringInput `pulumi:"topicArn"`
}

func (BucketNotificationTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketNotificationTopic)(nil)).Elem()
}

func (i BucketNotificationTopicArgs) ToBucketNotificationTopicOutput() BucketNotificationTopicOutput {
	return i.ToBucketNotificationTopicOutputWithContext(context.Background())
}

func (i BucketNotificationTopicArgs) ToBucketNotificationTopicOutputWithContext(ctx context.Context) BucketNotificationTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketNotificationTopicOutput)
}

type BucketNotificationTopicArrayInput interface {
	pulumi.Input

	ToBucketNotificationTopicArrayOutput() BucketNotificationTopicArrayOutput
	ToBucketNotificationTopicArrayOutputWithContext(context.Context) BucketNotificationTopicArrayOutput
}

type BucketNotificationTopicArray []BucketNotificationTopicInput

func (BucketNotificationTopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BucketNotificationTopic)(nil)).Elem()
}

func (i BucketNotificationTopicArray) ToBucketNotificationTopicArrayOutput() BucketNotificationTopicArrayOutput {
	return i.ToBucketNotificationTopicArrayOutputWithContext(context.Background())
}

func (i BucketNotificationTopicArray) ToBucketNotificationTopicArrayOutputWithContext(ctx context.Context) BucketNotificationTopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketNotificationTopicArrayOutput)
}

type BucketNotificationTopicOutput struct { *pulumi.OutputState }

func (BucketNotificationTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketNotificationTopic)(nil)).Elem()
}

func (o BucketNotificationTopicOutput) ToBucketNotificationTopicOutput() BucketNotificationTopicOutput {
	return o
}

func (o BucketNotificationTopicOutput) ToBucketNotificationTopicOutputWithContext(ctx context.Context) BucketNotificationTopicOutput {
	return o
}

// Specifies [event](http://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html#notification-how-to-event-types-and-destinations) for which to send notifications.
func (o BucketNotificationTopicOutput) Events() pulumi.StringArrayOutput {
	return o.ApplyT(func (v BucketNotificationTopic) []string { return v.Events }).(pulumi.StringArrayOutput)
}

// Specifies object key name prefix.
func (o BucketNotificationTopicOutput) FilterPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func (v BucketNotificationTopic) *string { return v.FilterPrefix }).(pulumi.StringPtrOutput)
}

// Specifies object key name suffix.
func (o BucketNotificationTopicOutput) FilterSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func (v BucketNotificationTopic) *string { return v.FilterSuffix }).(pulumi.StringPtrOutput)
}

// Specifies unique identifier for each of the notification configurations.
func (o BucketNotificationTopicOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func (v BucketNotificationTopic) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Specifies Amazon SNS topic ARN.
func (o BucketNotificationTopicOutput) TopicArn() pulumi.StringOutput {
	return o.ApplyT(func (v BucketNotificationTopic) string { return v.TopicArn }).(pulumi.StringOutput)
}

type BucketNotificationTopicArrayOutput struct { *pulumi.OutputState}

func (BucketNotificationTopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BucketNotificationTopic)(nil)).Elem()
}

func (o BucketNotificationTopicArrayOutput) ToBucketNotificationTopicArrayOutput() BucketNotificationTopicArrayOutput {
	return o
}

func (o BucketNotificationTopicArrayOutput) ToBucketNotificationTopicArrayOutputWithContext(ctx context.Context) BucketNotificationTopicArrayOutput {
	return o
}

func (o BucketNotificationTopicArrayOutput) Index(i pulumi.IntInput) BucketNotificationTopicOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) BucketNotificationTopic {
		return vs[0].([]BucketNotificationTopic)[vs[1].(int)]
	}).(BucketNotificationTopicOutput)
}

func init() {
	pulumi.RegisterOutputType(BucketNotificationTopicOutput{})
	pulumi.RegisterOutputType(BucketNotificationTopicArrayOutput{})
}
