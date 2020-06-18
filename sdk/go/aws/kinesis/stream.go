// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Kinesis Stream resource. Amazon Kinesis is a managed service that
// scales elastically for real-time processing of streaming big data.
//
// For more details, see the [Amazon Kinesis Documentation](https://aws.amazon.com/documentation/kinesis/).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/kinesis"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
// 			RetentionPeriod: pulumi.Int(48),
// 			ShardCount:      pulumi.Int(1),
// 			ShardLevelMetrics: pulumi.StringArray{
// 				pulumi.String("IncomingBytes"),
// 				pulumi.String("OutgoingBytes"),
// 			},
// 			Tags: map[string]interface{}{
// 				"Environment": "test",
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Stream struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
	EncryptionType pulumi.StringPtrOutput `pulumi:"encryptionType"`
	// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
	EnforceConsumerDeletion pulumi.BoolPtrOutput `pulumi:"enforceConsumerDeletion"`
	// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
	Name pulumi.StringOutput `pulumi:"name"`
	// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24.
	RetentionPeriod pulumi.IntPtrOutput `pulumi:"retentionPeriod"`
	// The number of shards that the stream will use.
	// Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
	ShardCount pulumi.IntOutput `pulumi:"shardCount"`
	// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
	ShardLevelMetrics pulumi.StringArrayOutput `pulumi:"shardLevelMetrics"`
	// A map of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewStream registers a new resource with the given unique name, arguments, and options.
func NewStream(ctx *pulumi.Context,
	name string, args *StreamArgs, opts ...pulumi.ResourceOption) (*Stream, error) {
	if args == nil || args.ShardCount == nil {
		return nil, errors.New("missing required argument 'ShardCount'")
	}
	if args == nil {
		args = &StreamArgs{}
	}
	var resource Stream
	err := ctx.RegisterResource("aws:kinesis/stream:Stream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStream gets an existing Stream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamState, opts ...pulumi.ResourceOption) (*Stream, error) {
	var resource Stream
	err := ctx.ReadResource("aws:kinesis/stream:Stream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stream resources.
type streamState struct {
	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn *string `pulumi:"arn"`
	// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
	EncryptionType *string `pulumi:"encryptionType"`
	// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
	EnforceConsumerDeletion *bool `pulumi:"enforceConsumerDeletion"`
	// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
	Name *string `pulumi:"name"`
	// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// The number of shards that the stream will use.
	// Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
	ShardCount *int `pulumi:"shardCount"`
	// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
	ShardLevelMetrics []string `pulumi:"shardLevelMetrics"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type StreamState struct {
	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn pulumi.StringPtrInput
	// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
	EncryptionType pulumi.StringPtrInput
	// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
	EnforceConsumerDeletion pulumi.BoolPtrInput
	// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
	KmsKeyId pulumi.StringPtrInput
	// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
	Name pulumi.StringPtrInput
	// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24.
	RetentionPeriod pulumi.IntPtrInput
	// The number of shards that the stream will use.
	// Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
	ShardCount pulumi.IntPtrInput
	// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
	ShardLevelMetrics pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (StreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamState)(nil)).Elem()
}

type streamArgs struct {
	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn *string `pulumi:"arn"`
	// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
	EncryptionType *string `pulumi:"encryptionType"`
	// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
	EnforceConsumerDeletion *bool `pulumi:"enforceConsumerDeletion"`
	// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
	Name *string `pulumi:"name"`
	// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// The number of shards that the stream will use.
	// Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
	ShardCount int `pulumi:"shardCount"`
	// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
	ShardLevelMetrics []string `pulumi:"shardLevelMetrics"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Stream resource.
type StreamArgs struct {
	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn pulumi.StringPtrInput
	// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
	EncryptionType pulumi.StringPtrInput
	// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
	EnforceConsumerDeletion pulumi.BoolPtrInput
	// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
	KmsKeyId pulumi.StringPtrInput
	// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
	Name pulumi.StringPtrInput
	// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24.
	RetentionPeriod pulumi.IntPtrInput
	// The number of shards that the stream will use.
	// Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
	ShardCount pulumi.IntInput
	// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
	ShardLevelMetrics pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (StreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamArgs)(nil)).Elem()
}
