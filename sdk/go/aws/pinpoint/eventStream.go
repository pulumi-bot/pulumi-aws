// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Pinpoint Event Stream resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kinesis"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/pinpoint"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		app, err := pinpoint.NewApp(ctx, "app", nil)
// 		if err != nil {
// 			return err
// 		}
// 		testStream, err := kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
// 			ShardCount: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"pinpoint.us-east-1.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pinpoint.NewEventStream(ctx, "stream", &pinpoint.EventStreamArgs{
// 			ApplicationId:        app.ApplicationId,
// 			DestinationStreamArn: testStream.Arn,
// 			RoleArn:              testRole.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "testRolePolicy", &iam.RolePolicyArgs{
// 			Role:   testRole.ID(),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": {\n", "    \"Action\": [\n", "      \"kinesis:PutRecords\",\n", "      \"kinesis:DescribeStream\"\n", "    ],\n", "    \"Effect\": \"Allow\",\n", "    \"Resource\": [\n", "      \"arn:aws:kinesis:us-east-1:*:*/*\"\n", "    ]\n", "  }\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type EventStream struct {
	pulumi.CustomResourceState

	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream or Firehose delivery stream to which you want to publish events.
	DestinationStreamArn pulumi.StringOutput `pulumi:"destinationStreamArn"`
	// The IAM role that authorizes Amazon Pinpoint to publish events to the stream in your account.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
}

// NewEventStream registers a new resource with the given unique name, arguments, and options.
func NewEventStream(ctx *pulumi.Context,
	name string, args *EventStreamArgs, opts ...pulumi.ResourceOption) (*EventStream, error) {
	if args == nil || args.ApplicationId == nil {
		return nil, errors.New("missing required argument 'ApplicationId'")
	}
	if args == nil || args.DestinationStreamArn == nil {
		return nil, errors.New("missing required argument 'DestinationStreamArn'")
	}
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	if args == nil {
		args = &EventStreamArgs{}
	}
	var resource EventStream
	err := ctx.RegisterResource("aws:pinpoint/eventStream:EventStream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventStream gets an existing EventStream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventStreamState, opts ...pulumi.ResourceOption) (*EventStream, error) {
	var resource EventStream
	err := ctx.ReadResource("aws:pinpoint/eventStream:EventStream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventStream resources.
type eventStreamState struct {
	// The application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream or Firehose delivery stream to which you want to publish events.
	DestinationStreamArn *string `pulumi:"destinationStreamArn"`
	// The IAM role that authorizes Amazon Pinpoint to publish events to the stream in your account.
	RoleArn *string `pulumi:"roleArn"`
}

type EventStreamState struct {
	// The application ID.
	ApplicationId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream or Firehose delivery stream to which you want to publish events.
	DestinationStreamArn pulumi.StringPtrInput
	// The IAM role that authorizes Amazon Pinpoint to publish events to the stream in your account.
	RoleArn pulumi.StringPtrInput
}

func (EventStreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventStreamState)(nil)).Elem()
}

type eventStreamArgs struct {
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream or Firehose delivery stream to which you want to publish events.
	DestinationStreamArn string `pulumi:"destinationStreamArn"`
	// The IAM role that authorizes Amazon Pinpoint to publish events to the stream in your account.
	RoleArn string `pulumi:"roleArn"`
}

// The set of arguments for constructing a EventStream resource.
type EventStreamArgs struct {
	// The application ID.
	ApplicationId pulumi.StringInput
	// The Amazon Resource Name (ARN) of the Amazon Kinesis stream or Firehose delivery stream to which you want to publish events.
	DestinationStreamArn pulumi.StringInput
	// The IAM role that authorizes Amazon Pinpoint to publish events to the stream in your account.
	RoleArn pulumi.StringInput
}

func (EventStreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventStreamArgs)(nil)).Elem()
}

type EventStreamInput interface {
	pulumi.Input

	ToEventStreamOutput() EventStreamOutput
	ToEventStreamOutputWithContext(ctx context.Context) EventStreamOutput
}

func (EventStream) ElementType() reflect.Type {
	return reflect.TypeOf((*EventStream)(nil)).Elem()
}

func (i EventStream) ToEventStreamOutput() EventStreamOutput {
	return i.ToEventStreamOutputWithContext(context.Background())
}

func (i EventStream) ToEventStreamOutputWithContext(ctx context.Context) EventStreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventStreamOutput)
}

type EventStreamOutput struct {
	*pulumi.OutputState
}

func (EventStreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventStreamOutput)(nil)).Elem()
}

func (o EventStreamOutput) ToEventStreamOutput() EventStreamOutput {
	return o
}

func (o EventStreamOutput) ToEventStreamOutputWithContext(ctx context.Context) EventStreamOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventStreamOutput{})
}
