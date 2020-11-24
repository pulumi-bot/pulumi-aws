// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a S3 Bucket Notification Configuration. For additional information, see the [Configuring S3 Event Notifications section in the Amazon S3 Developer Guide](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html).
//
// > **NOTE:** S3 Buckets only support a single notification configuration. Declaring multiple `s3.BucketNotification` resources to the same S3 Bucket will cause a perpetual difference in configuration. See the example "Trigger multiple Lambda functions" for an option.
//
// ## Example Usage
// ### Add notification configuration to SNS Topic
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := s3.NewBucket(ctx, "bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		topic, err := sns.NewTopic(ctx, "topic", &sns.TopicArgs{
// 			Policy: bucket.Arn.ApplyT(func(arn string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\":\"2012-10-17\",\n", "    \"Statement\":[{\n", "        \"Effect\": \"Allow\",\n", "        \"Principal\": {\"AWS\":\"*\"},\n", "        \"Action\": \"SNS:Publish\",\n", "        \"Resource\": \"arn:aws:sns:*:*:s3-event-notification-topic\",\n", "        \"Condition\":{\n", "            \"ArnLike\":{\"aws:SourceArn\":\"", arn, "\"}\n", "        }\n", "    }]\n", "}\n"), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
// 			Bucket: bucket.ID(),
// 			Topics: s3.BucketNotificationTopicArray{
// 				&s3.BucketNotificationTopicArgs{
// 					TopicArn: topic.Arn,
// 					Events: pulumi.StringArray{
// 						pulumi.String("s3:ObjectCreated:*"),
// 					},
// 					FilterSuffix: pulumi.String(".log"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Add notification configuration to SQS Queue
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sqs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := s3.NewBucket(ctx, "bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		queue, err := sqs.NewQueue(ctx, "queue", &sqs.QueueArgs{
// 			Policy: bucket.Arn.ApplyT(func(arn string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": \"*\",\n", "      \"Action\": \"sqs:SendMessage\",\n", "	  \"Resource\": \"arn:aws:sqs:*:*:s3-event-notification-queue\",\n", "      \"Condition\": {\n", "        \"ArnEquals\": { \"aws:SourceArn\": \"", arn, "\" }\n", "      }\n", "    }\n", "  ]\n", "}\n"), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
// 			Bucket: bucket.ID(),
// 			Queues: s3.BucketNotificationQueueArray{
// 				&s3.BucketNotificationQueueArgs{
// 					QueueArn: queue.Arn,
// 					Events: pulumi.StringArray{
// 						pulumi.String("s3:ObjectCreated:*"),
// 					},
// 					FilterSuffix: pulumi.String(".log"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Add multiple notification configurations to SQS Queue
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sqs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := s3.NewBucket(ctx, "bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		queue, err := sqs.NewQueue(ctx, "queue", &sqs.QueueArgs{
// 			Policy: bucket.Arn.ApplyT(func(arn string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": \"*\",\n", "      \"Action\": \"sqs:SendMessage\",\n", "	  \"Resource\": \"arn:aws:sqs:*:*:s3-event-notification-queue\",\n", "      \"Condition\": {\n", "        \"ArnEquals\": { \"aws:SourceArn\": \"", arn, "\" }\n", "      }\n", "    }\n", "  ]\n", "}\n"), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
// 			Bucket: bucket.ID(),
// 			Queues: s3.BucketNotificationQueueArray{
// 				&s3.BucketNotificationQueueArgs{
// 					Id:       pulumi.String("image-upload-event"),
// 					QueueArn: queue.Arn,
// 					Events: pulumi.StringArray{
// 						pulumi.String("s3:ObjectCreated:*"),
// 					},
// 					FilterPrefix: pulumi.String("images/"),
// 				},
// 				&s3.BucketNotificationQueueArgs{
// 					Id:       pulumi.String("video-upload-event"),
// 					QueueArn: queue.Arn,
// 					Events: pulumi.StringArray{
// 						pulumi.String("s3:ObjectCreated:*"),
// 					},
// 					FilterPrefix: pulumi.String("videos/"),
// 				},
// 			},
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
// S3 bucket notification can be imported using the `bucket`, e.g.
//
// ```sh
//  $ pulumi import aws:s3/bucketNotification:BucketNotification bucket_notification bucket-name
// ```
type BucketNotification struct {
	pulumi.CustomResourceState

	// The name of the bucket to put notification configuration.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Used to configure notifications to a Lambda Function (documented below).
	LambdaFunctions BucketNotificationLambdaFunctionArrayOutput `pulumi:"lambdaFunctions"`
	// The notification configuration to SQS Queue (documented below).
	Queues BucketNotificationQueueArrayOutput `pulumi:"queues"`
	// The notification configuration to SNS Topic (documented below).
	Topics BucketNotificationTopicArrayOutput `pulumi:"topics"`
}

// NewBucketNotification registers a new resource with the given unique name, arguments, and options.
func NewBucketNotification(ctx *pulumi.Context,
	name string, args *BucketNotificationArgs, opts ...pulumi.ResourceOption) (*BucketNotification, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	var resource BucketNotification
	err := ctx.RegisterResource("aws:s3/bucketNotification:BucketNotification", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketNotification gets an existing BucketNotification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketNotification(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketNotificationState, opts ...pulumi.ResourceOption) (*BucketNotification, error) {
	var resource BucketNotification
	err := ctx.ReadResource("aws:s3/bucketNotification:BucketNotification", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketNotification resources.
type bucketNotificationState struct {
	// The name of the bucket to put notification configuration.
	Bucket *string `pulumi:"bucket"`
	// Used to configure notifications to a Lambda Function (documented below).
	LambdaFunctions []BucketNotificationLambdaFunction `pulumi:"lambdaFunctions"`
	// The notification configuration to SQS Queue (documented below).
	Queues []BucketNotificationQueue `pulumi:"queues"`
	// The notification configuration to SNS Topic (documented below).
	Topics []BucketNotificationTopic `pulumi:"topics"`
}

type BucketNotificationState struct {
	// The name of the bucket to put notification configuration.
	Bucket pulumi.StringPtrInput
	// Used to configure notifications to a Lambda Function (documented below).
	LambdaFunctions BucketNotificationLambdaFunctionArrayInput
	// The notification configuration to SQS Queue (documented below).
	Queues BucketNotificationQueueArrayInput
	// The notification configuration to SNS Topic (documented below).
	Topics BucketNotificationTopicArrayInput
}

func (BucketNotificationState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketNotificationState)(nil)).Elem()
}

type bucketNotificationArgs struct {
	// The name of the bucket to put notification configuration.
	Bucket string `pulumi:"bucket"`
	// Used to configure notifications to a Lambda Function (documented below).
	LambdaFunctions []BucketNotificationLambdaFunction `pulumi:"lambdaFunctions"`
	// The notification configuration to SQS Queue (documented below).
	Queues []BucketNotificationQueue `pulumi:"queues"`
	// The notification configuration to SNS Topic (documented below).
	Topics []BucketNotificationTopic `pulumi:"topics"`
}

// The set of arguments for constructing a BucketNotification resource.
type BucketNotificationArgs struct {
	// The name of the bucket to put notification configuration.
	Bucket pulumi.StringInput
	// Used to configure notifications to a Lambda Function (documented below).
	LambdaFunctions BucketNotificationLambdaFunctionArrayInput
	// The notification configuration to SQS Queue (documented below).
	Queues BucketNotificationQueueArrayInput
	// The notification configuration to SNS Topic (documented below).
	Topics BucketNotificationTopicArrayInput
}

func (BucketNotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketNotificationArgs)(nil)).Elem()
}

type BucketNotificationInput interface {
	pulumi.Input

	ToBucketNotificationOutput() BucketNotificationOutput
	ToBucketNotificationOutputWithContext(ctx context.Context) BucketNotificationOutput
}

func (BucketNotification) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketNotification)(nil)).Elem()
}

func (i BucketNotification) ToBucketNotificationOutput() BucketNotificationOutput {
	return i.ToBucketNotificationOutputWithContext(context.Background())
}

func (i BucketNotification) ToBucketNotificationOutputWithContext(ctx context.Context) BucketNotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketNotificationOutput)
}

type BucketNotificationOutput struct {
	*pulumi.OutputState
}

func (BucketNotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketNotificationOutput)(nil)).Elem()
}

func (o BucketNotificationOutput) ToBucketNotificationOutput() BucketNotificationOutput {
	return o
}

func (o BucketNotificationOutput) ToBucketNotificationOutputWithContext(ctx context.Context) BucketNotificationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketNotificationOutput{})
}
