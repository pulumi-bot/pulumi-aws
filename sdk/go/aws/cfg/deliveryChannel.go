// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS Config Delivery Channel.
//
// > **Note:** Delivery Channel requires a `Configuration Recorder` to be present. Use of `dependsOn` (as shown below) is recommended to avoid race conditions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cfg"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
// 			ForceDestroy: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"config.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fooRecorder, err := cfg.NewRecorder(ctx, "fooRecorder", &cfg.RecorderArgs{
// 			RoleArn: role.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cfg.NewDeliveryChannel(ctx, "fooDeliveryChannel", &cfg.DeliveryChannelArgs{
// 			S3BucketName: bucket.Bucket,
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			fooRecorder,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "rolePolicy", &iam.RolePolicyArgs{
// 			Role: role.ID(),
// 			Policy: pulumi.All(bucket.Arn, bucket.Arn).ApplyT(func(_args []interface{}) (string, error) {
// 				bucketArn := _args[0].(string)
// 				bucketArn1 := _args[1].(string)
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"s3:*\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": [\n", "        \"", bucketArn, "\",\n", "        \"", bucketArn1, "/*\"\n", "      ]\n", "    }\n", "  ]\n", "}\n"), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Delivery Channel can be imported using the name, e.g.
//
// ```sh
//  $ pulumi import aws:cfg/deliveryChannel:DeliveryChannel foo example
// ```
type DeliveryChannel struct {
	pulumi.CustomResourceState

	// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the S3 bucket used to store the configuration history.
	S3BucketName pulumi.StringOutput `pulumi:"s3BucketName"`
	// The prefix for the specified S3 bucket.
	S3KeyPrefix pulumi.StringPtrOutput `pulumi:"s3KeyPrefix"`
	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties DeliveryChannelSnapshotDeliveryPropertiesPtrOutput `pulumi:"snapshotDeliveryProperties"`
	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn pulumi.StringPtrOutput `pulumi:"snsTopicArn"`
}

// NewDeliveryChannel registers a new resource with the given unique name, arguments, and options.
func NewDeliveryChannel(ctx *pulumi.Context,
	name string, args *DeliveryChannelArgs, opts ...pulumi.ResourceOption) (*DeliveryChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.S3BucketName == nil {
		return nil, errors.New("invalid value for required argument 'S3BucketName'")
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
	SnapshotDeliveryProperties *DeliveryChannelSnapshotDeliveryProperties `pulumi:"snapshotDeliveryProperties"`
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
	SnapshotDeliveryProperties DeliveryChannelSnapshotDeliveryPropertiesPtrInput
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
	SnapshotDeliveryProperties *DeliveryChannelSnapshotDeliveryProperties `pulumi:"snapshotDeliveryProperties"`
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
	SnapshotDeliveryProperties DeliveryChannelSnapshotDeliveryPropertiesPtrInput
	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn pulumi.StringPtrInput
}

func (DeliveryChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deliveryChannelArgs)(nil)).Elem()
}

type DeliveryChannelInput interface {
	pulumi.Input

	ToDeliveryChannelOutput() DeliveryChannelOutput
	ToDeliveryChannelOutputWithContext(ctx context.Context) DeliveryChannelOutput
}

func (*DeliveryChannel) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryChannel)(nil))
}

func (i *DeliveryChannel) ToDeliveryChannelOutput() DeliveryChannelOutput {
	return i.ToDeliveryChannelOutputWithContext(context.Background())
}

func (i *DeliveryChannel) ToDeliveryChannelOutputWithContext(ctx context.Context) DeliveryChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryChannelOutput)
}

func (i *DeliveryChannel) ToDeliveryChannelPtrOutput() DeliveryChannelPtrOutput {
	return i.ToDeliveryChannelPtrOutputWithContext(context.Background())
}

func (i *DeliveryChannel) ToDeliveryChannelPtrOutputWithContext(ctx context.Context) DeliveryChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryChannelPtrOutput)
}

type DeliveryChannelPtrInput interface {
	pulumi.Input

	ToDeliveryChannelPtrOutput() DeliveryChannelPtrOutput
	ToDeliveryChannelPtrOutputWithContext(ctx context.Context) DeliveryChannelPtrOutput
}

type deliveryChannelPtrType DeliveryChannelArgs

func (*deliveryChannelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryChannel)(nil))
}

func (i *deliveryChannelPtrType) ToDeliveryChannelPtrOutput() DeliveryChannelPtrOutput {
	return i.ToDeliveryChannelPtrOutputWithContext(context.Background())
}

func (i *deliveryChannelPtrType) ToDeliveryChannelPtrOutputWithContext(ctx context.Context) DeliveryChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryChannelOutput).ToDeliveryChannelPtrOutput()
}

// DeliveryChannelArrayInput is an input type that accepts DeliveryChannelArray and DeliveryChannelArrayOutput values.
// You can construct a concrete instance of `DeliveryChannelArrayInput` via:
//
//          DeliveryChannelArray{ DeliveryChannelArgs{...} }
type DeliveryChannelArrayInput interface {
	pulumi.Input

	ToDeliveryChannelArrayOutput() DeliveryChannelArrayOutput
	ToDeliveryChannelArrayOutputWithContext(context.Context) DeliveryChannelArrayOutput
}

type DeliveryChannelArray []DeliveryChannelInput

func (DeliveryChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeliveryChannel)(nil))
}

func (i DeliveryChannelArray) ToDeliveryChannelArrayOutput() DeliveryChannelArrayOutput {
	return i.ToDeliveryChannelArrayOutputWithContext(context.Background())
}

func (i DeliveryChannelArray) ToDeliveryChannelArrayOutputWithContext(ctx context.Context) DeliveryChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryChannelArrayOutput)
}

// DeliveryChannelMapInput is an input type that accepts DeliveryChannelMap and DeliveryChannelMapOutput values.
// You can construct a concrete instance of `DeliveryChannelMapInput` via:
//
//          DeliveryChannelMap{ "key": DeliveryChannelArgs{...} }
type DeliveryChannelMapInput interface {
	pulumi.Input

	ToDeliveryChannelMapOutput() DeliveryChannelMapOutput
	ToDeliveryChannelMapOutputWithContext(context.Context) DeliveryChannelMapOutput
}

type DeliveryChannelMap map[string]DeliveryChannelInput

func (DeliveryChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DeliveryChannel)(nil))
}

func (i DeliveryChannelMap) ToDeliveryChannelMapOutput() DeliveryChannelMapOutput {
	return i.ToDeliveryChannelMapOutputWithContext(context.Background())
}

func (i DeliveryChannelMap) ToDeliveryChannelMapOutputWithContext(ctx context.Context) DeliveryChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryChannelMapOutput)
}

type DeliveryChannelOutput struct {
	*pulumi.OutputState
}

func (DeliveryChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryChannel)(nil))
}

func (o DeliveryChannelOutput) ToDeliveryChannelOutput() DeliveryChannelOutput {
	return o
}

func (o DeliveryChannelOutput) ToDeliveryChannelOutputWithContext(ctx context.Context) DeliveryChannelOutput {
	return o
}

func (o DeliveryChannelOutput) ToDeliveryChannelPtrOutput() DeliveryChannelPtrOutput {
	return o.ToDeliveryChannelPtrOutputWithContext(context.Background())
}

func (o DeliveryChannelOutput) ToDeliveryChannelPtrOutputWithContext(ctx context.Context) DeliveryChannelPtrOutput {
	return o.ApplyT(func(v DeliveryChannel) *DeliveryChannel {
		return &v
	}).(DeliveryChannelPtrOutput)
}

type DeliveryChannelPtrOutput struct {
	*pulumi.OutputState
}

func (DeliveryChannelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryChannel)(nil))
}

func (o DeliveryChannelPtrOutput) ToDeliveryChannelPtrOutput() DeliveryChannelPtrOutput {
	return o
}

func (o DeliveryChannelPtrOutput) ToDeliveryChannelPtrOutputWithContext(ctx context.Context) DeliveryChannelPtrOutput {
	return o
}

type DeliveryChannelArrayOutput struct{ *pulumi.OutputState }

func (DeliveryChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeliveryChannel)(nil))
}

func (o DeliveryChannelArrayOutput) ToDeliveryChannelArrayOutput() DeliveryChannelArrayOutput {
	return o
}

func (o DeliveryChannelArrayOutput) ToDeliveryChannelArrayOutputWithContext(ctx context.Context) DeliveryChannelArrayOutput {
	return o
}

func (o DeliveryChannelArrayOutput) Index(i pulumi.IntInput) DeliveryChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeliveryChannel {
		return vs[0].([]DeliveryChannel)[vs[1].(int)]
	}).(DeliveryChannelOutput)
}

type DeliveryChannelMapOutput struct{ *pulumi.OutputState }

func (DeliveryChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DeliveryChannel)(nil))
}

func (o DeliveryChannelMapOutput) ToDeliveryChannelMapOutput() DeliveryChannelMapOutput {
	return o
}

func (o DeliveryChannelMapOutput) ToDeliveryChannelMapOutputWithContext(ctx context.Context) DeliveryChannelMapOutput {
	return o
}

func (o DeliveryChannelMapOutput) MapIndex(k pulumi.StringInput) DeliveryChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DeliveryChannel {
		return vs[0].(map[string]DeliveryChannel)[vs[1].(string)]
	}).(DeliveryChannelOutput)
}

func init() {
	pulumi.RegisterOutputType(DeliveryChannelOutput{})
	pulumi.RegisterOutputType(DeliveryChannelPtrOutput{})
	pulumi.RegisterOutputType(DeliveryChannelArrayOutput{})
	pulumi.RegisterOutputType(DeliveryChannelMapOutput{})
}
