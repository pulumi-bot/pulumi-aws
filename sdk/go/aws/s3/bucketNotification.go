// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type BucketNotification struct {
	pulumi.CustomResourceState

	Bucket          pulumi.StringOutput                         `pulumi:"bucket"`
	LambdaFunctions BucketNotificationLambdaFunctionArrayOutput `pulumi:"lambdaFunctions"`
	Queues          BucketNotificationQueueArrayOutput          `pulumi:"queues"`
	Topics          BucketNotificationTopicArrayOutput          `pulumi:"topics"`
}

// NewBucketNotification registers a new resource with the given unique name, arguments, and options.
func NewBucketNotification(ctx *pulumi.Context,
	name string, args *BucketNotificationArgs, opts ...pulumi.ResourceOption) (*BucketNotification, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil {
		args = &BucketNotificationArgs{}
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
	Bucket          *string                            `pulumi:"bucket"`
	LambdaFunctions []BucketNotificationLambdaFunction `pulumi:"lambdaFunctions"`
	Queues          []BucketNotificationQueue          `pulumi:"queues"`
	Topics          []BucketNotificationTopic          `pulumi:"topics"`
}

type BucketNotificationState struct {
	Bucket          pulumi.StringPtrInput
	LambdaFunctions BucketNotificationLambdaFunctionArrayInput
	Queues          BucketNotificationQueueArrayInput
	Topics          BucketNotificationTopicArrayInput
}

func (BucketNotificationState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketNotificationState)(nil)).Elem()
}

type bucketNotificationArgs struct {
	Bucket          string                             `pulumi:"bucket"`
	LambdaFunctions []BucketNotificationLambdaFunction `pulumi:"lambdaFunctions"`
	Queues          []BucketNotificationQueue          `pulumi:"queues"`
	Topics          []BucketNotificationTopic          `pulumi:"topics"`
}

// The set of arguments for constructing a BucketNotification resource.
type BucketNotificationArgs struct {
	Bucket          pulumi.StringInput
	LambdaFunctions BucketNotificationLambdaFunctionArrayInput
	Queues          BucketNotificationQueueArrayInput
	Topics          BucketNotificationTopicArrayInput
}

func (BucketNotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketNotificationArgs)(nil)).Elem()
}
