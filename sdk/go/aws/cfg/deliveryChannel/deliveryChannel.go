// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package deliveryChannel

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/cfg/DeliveryChannelSnapshotDeliveryProperties"
)

// Provides an AWS Config Delivery Channel.
// 
// > **Note:** Delivery Channel requires a [Configuration Recorder](https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html) to be present. Use of `dependsOn` (as shown below) is recommended to avoid race conditions.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/config_delivery_channel.html.markdown.
type DeliveryChannel struct {
	pulumi.CustomResourceState

	// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the S3 bucket used to store the configuration history.
	S3BucketName pulumi.StringOutput `pulumi:"s3BucketName"`
	// The prefix for the specified S3 bucket.
	S3KeyPrefix pulumi.StringPtrOutput `pulumi:"s3KeyPrefix"`
	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties cfgDeliveryChannelSnapshotDeliveryProperties.DeliveryChannelSnapshotDeliveryPropertiesPtrOutput `pulumi:"snapshotDeliveryProperties"`
	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn pulumi.StringPtrOutput `pulumi:"snsTopicArn"`
}

// NewDeliveryChannel registers a new resource with the given unique name, arguments, and options.
func NewDeliveryChannel(ctx *pulumi.Context,
	name string, args *DeliveryChannelArgs, opts ...pulumi.ResourceOption) (*DeliveryChannel, error) {
	if args == nil || args.S3BucketName == nil {
		return nil, errors.New("missing required argument 'S3BucketName'")
	}
	if args == nil {
		args = &DeliveryChannelArgs{}
	}
	var resource DeliveryChannel
	err := ctx.RegisterResource("aws:cfg/deliveryChannel:DeliveryChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeliveryChannel gets an existing DeliveryChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeliveryChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeliveryChannelState, opts ...pulumi.ResourceOption) (*DeliveryChannel, error) {
	var resource DeliveryChannel
	err := ctx.ReadResource("aws:cfg/deliveryChannel:DeliveryChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeliveryChannel resources.
type deliveryChannelState struct {
	// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
	Name *string `pulumi:"name"`
	// The name of the S3 bucket used to store the configuration history.
	S3BucketName *string `pulumi:"s3BucketName"`
	// The prefix for the specified S3 bucket.
	S3KeyPrefix *string `pulumi:"s3KeyPrefix"`
	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties *cfgDeliveryChannelSnapshotDeliveryProperties.DeliveryChannelSnapshotDeliveryProperties `pulumi:"snapshotDeliveryProperties"`
	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn *string `pulumi:"snsTopicArn"`
}

type DeliveryChannelState struct {
	// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringPtrInput
	// The name of the S3 bucket used to store the configuration history.
	S3BucketName pulumi.StringPtrInput
	// The prefix for the specified S3 bucket.
	S3KeyPrefix pulumi.StringPtrInput
	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties cfgDeliveryChannelSnapshotDeliveryProperties.DeliveryChannelSnapshotDeliveryPropertiesPtrInput
	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn pulumi.StringPtrInput
}

func (DeliveryChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*deliveryChannelState)(nil)).Elem()
}

type deliveryChannelArgs struct {
	// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
	Name *string `pulumi:"name"`
	// The name of the S3 bucket used to store the configuration history.
	S3BucketName string `pulumi:"s3BucketName"`
	// The prefix for the specified S3 bucket.
	S3KeyPrefix *string `pulumi:"s3KeyPrefix"`
	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties *cfgDeliveryChannelSnapshotDeliveryProperties.DeliveryChannelSnapshotDeliveryProperties `pulumi:"snapshotDeliveryProperties"`
	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn *string `pulumi:"snsTopicArn"`
}

// The set of arguments for constructing a DeliveryChannel resource.
type DeliveryChannelArgs struct {
	// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringPtrInput
	// The name of the S3 bucket used to store the configuration history.
	S3BucketName pulumi.StringInput
	// The prefix for the specified S3 bucket.
	S3KeyPrefix pulumi.StringPtrInput
	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties cfgDeliveryChannelSnapshotDeliveryProperties.DeliveryChannelSnapshotDeliveryPropertiesPtrInput
	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn pulumi.StringPtrInput
}

func (DeliveryChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deliveryChannelArgs)(nil)).Elem()
}

