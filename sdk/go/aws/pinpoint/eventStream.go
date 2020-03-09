// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package pinpoint

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Pinpoint Event Stream resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/pinpoint_event_stream.markdown.
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

