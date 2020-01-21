// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package logDestination

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CloudWatch Logs destination resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_destination.html.markdown.
type LogDestination struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) specifying the log destination.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A name for the log destination
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The ARN of the target Amazon Kinesis stream resource for the destination
	TargetArn pulumi.StringOutput `pulumi:"targetArn"`
}

// NewLogDestination registers a new resource with the given unique name, arguments, and options.
func NewLogDestination(ctx *pulumi.Context,
	name string, args *LogDestinationArgs, opts ...pulumi.ResourceOption) (*LogDestination, error) {
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	if args == nil || args.TargetArn == nil {
		return nil, errors.New("missing required argument 'TargetArn'")
	}
	if args == nil {
		args = &LogDestinationArgs{}
	}
	var resource LogDestination
	err := ctx.RegisterResource("aws:cloudwatch/logDestination:LogDestination", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogDestination gets an existing LogDestination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogDestinationState, opts ...pulumi.ResourceOption) (*LogDestination, error) {
	var resource LogDestination
	err := ctx.ReadResource("aws:cloudwatch/logDestination:LogDestination", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogDestination resources.
type logDestinationState struct {
	// The Amazon Resource Name (ARN) specifying the log destination.
	Arn *string `pulumi:"arn"`
	// A name for the log destination
	Name *string `pulumi:"name"`
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target
	RoleArn *string `pulumi:"roleArn"`
	// The ARN of the target Amazon Kinesis stream resource for the destination
	TargetArn *string `pulumi:"targetArn"`
}

type LogDestinationState struct {
	// The Amazon Resource Name (ARN) specifying the log destination.
	Arn pulumi.StringPtrInput
	// A name for the log destination
	Name pulumi.StringPtrInput
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target
	RoleArn pulumi.StringPtrInput
	// The ARN of the target Amazon Kinesis stream resource for the destination
	TargetArn pulumi.StringPtrInput
}

func (LogDestinationState) ElementType() reflect.Type {
	return reflect.TypeOf((*logDestinationState)(nil)).Elem()
}

type logDestinationArgs struct {
	// A name for the log destination
	Name *string `pulumi:"name"`
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target
	RoleArn string `pulumi:"roleArn"`
	// The ARN of the target Amazon Kinesis stream resource for the destination
	TargetArn string `pulumi:"targetArn"`
}

// The set of arguments for constructing a LogDestination resource.
type LogDestinationArgs struct {
	// A name for the log destination
	Name pulumi.StringPtrInput
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target
	RoleArn pulumi.StringInput
	// The ARN of the target Amazon Kinesis stream resource for the destination
	TargetArn pulumi.StringInput
}

func (LogDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logDestinationArgs)(nil)).Elem()
}

