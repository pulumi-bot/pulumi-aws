// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CloudWatch Event Target resource.
type EventTarget struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) associated of the target.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget EventTargetBatchTargetPtrOutput `pulumi:"batchTarget"`
	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget EventTargetEcsTargetPtrOutput `pulumi:"ecsTarget"`
	// Valid JSON text passed to the target.
	Input pulumi.StringPtrOutput `pulumi:"input"`
	// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
	// that is used for extracting part of the matched event when passing it to the target.
	InputPath pulumi.StringPtrOutput `pulumi:"inputPath"`
	// Parameters used when you are providing a custom input to a target based on certain event data.
	InputTransformer EventTargetInputTransformerPtrOutput `pulumi:"inputTransformer"`
	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget EventTargetKinesisTargetPtrOutput `pulumi:"kinesisTarget"`
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
	// The name of the rule you want to add targets to.
	Rule pulumi.StringOutput `pulumi:"rule"`
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets EventTargetRunCommandTargetArrayOutput `pulumi:"runCommandTargets"`
	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget EventTargetSqsTargetPtrOutput `pulumi:"sqsTarget"`
	// The unique target assignment ID.  If missing, will generate a random, unique id.
	TargetId pulumi.StringOutput `pulumi:"targetId"`
}

// NewEventTarget registers a new resource with the given unique name, arguments, and options.
func NewEventTarget(ctx *pulumi.Context,
	name string, args *EventTargetArgs, opts ...pulumi.ResourceOption) (*EventTarget, error) {
	if args == nil || args.Arn == nil {
		return nil, errors.New("missing required argument 'Arn'")
	}
	if args == nil || args.Rule == nil {
		return nil, errors.New("missing required argument 'Rule'")
	}
	if args == nil {
		args = &EventTargetArgs{}
	}
	var resource EventTarget
	err := ctx.RegisterResource("aws:cloudwatch/eventTarget:EventTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventTarget gets an existing EventTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventTargetState, opts ...pulumi.ResourceOption) (*EventTarget, error) {
	var resource EventTarget
	err := ctx.ReadResource("aws:cloudwatch/eventTarget:EventTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventTarget resources.
type eventTargetState struct {
	// The Amazon Resource Name (ARN) associated of the target.
	Arn *string `pulumi:"arn"`
	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget *EventTargetBatchTarget `pulumi:"batchTarget"`
	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget *EventTargetEcsTarget `pulumi:"ecsTarget"`
	// Valid JSON text passed to the target.
	Input *string `pulumi:"input"`
	// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
	// that is used for extracting part of the matched event when passing it to the target.
	InputPath *string `pulumi:"inputPath"`
	// Parameters used when you are providing a custom input to a target based on certain event data.
	InputTransformer *EventTargetInputTransformer `pulumi:"inputTransformer"`
	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget *EventTargetKinesisTarget `pulumi:"kinesisTarget"`
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
	RoleArn *string `pulumi:"roleArn"`
	// The name of the rule you want to add targets to.
	Rule *string `pulumi:"rule"`
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets []EventTargetRunCommandTarget `pulumi:"runCommandTargets"`
	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget *EventTargetSqsTarget `pulumi:"sqsTarget"`
	// The unique target assignment ID.  If missing, will generate a random, unique id.
	TargetId *string `pulumi:"targetId"`
}

type EventTargetState struct {
	// The Amazon Resource Name (ARN) associated of the target.
	Arn pulumi.StringPtrInput
	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget EventTargetBatchTargetPtrInput
	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget EventTargetEcsTargetPtrInput
	// Valid JSON text passed to the target.
	Input pulumi.StringPtrInput
	// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
	// that is used for extracting part of the matched event when passing it to the target.
	InputPath pulumi.StringPtrInput
	// Parameters used when you are providing a custom input to a target based on certain event data.
	InputTransformer EventTargetInputTransformerPtrInput
	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget EventTargetKinesisTargetPtrInput
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
	RoleArn pulumi.StringPtrInput
	// The name of the rule you want to add targets to.
	Rule pulumi.StringPtrInput
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets EventTargetRunCommandTargetArrayInput
	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget EventTargetSqsTargetPtrInput
	// The unique target assignment ID.  If missing, will generate a random, unique id.
	TargetId pulumi.StringPtrInput
}

func (EventTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventTargetState)(nil)).Elem()
}

type eventTargetArgs struct {
	// The Amazon Resource Name (ARN) associated of the target.
	Arn string `pulumi:"arn"`
	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget *EventTargetBatchTargetArgs `pulumi:"batchTarget"`
	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget *EventTargetEcsTargetArgs `pulumi:"ecsTarget"`
	// Valid JSON text passed to the target.
	Input *string `pulumi:"input"`
	// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
	// that is used for extracting part of the matched event when passing it to the target.
	InputPath *string `pulumi:"inputPath"`
	// Parameters used when you are providing a custom input to a target based on certain event data.
	InputTransformer *EventTargetInputTransformerArgs `pulumi:"inputTransformer"`
	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget *EventTargetKinesisTargetArgs `pulumi:"kinesisTarget"`
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
	RoleArn *string `pulumi:"roleArn"`
	// The name of the rule you want to add targets to.
	Rule string `pulumi:"rule"`
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets []EventTargetRunCommandTargetArgs `pulumi:"runCommandTargets"`
	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget *EventTargetSqsTargetArgs `pulumi:"sqsTarget"`
	// The unique target assignment ID.  If missing, will generate a random, unique id.
	TargetId *string `pulumi:"targetId"`
}

// The set of arguments for constructing a EventTarget resource.
type EventTargetArgs struct {
	// The Amazon Resource Name (ARN) associated of the target.
	Arn pulumi.StringInput
	// Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
	BatchTarget EventTargetBatchTargetArgsPtrInput
	// Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
	EcsTarget EventTargetEcsTargetArgsPtrInput
	// Valid JSON text passed to the target.
	Input pulumi.StringPtrInput
	// The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
	// that is used for extracting part of the matched event when passing it to the target.
	InputPath pulumi.StringPtrInput
	// Parameters used when you are providing a custom input to a target based on certain event data.
	InputTransformer EventTargetInputTransformerArgsPtrInput
	// Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
	KinesisTarget EventTargetKinesisTargetArgsPtrInput
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
	RoleArn pulumi.StringPtrInput
	// The name of the rule you want to add targets to.
	Rule pulumi.StringInput
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
	RunCommandTargets EventTargetRunCommandTargetArgsArrayInput
	// Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
	SqsTarget EventTargetSqsTargetArgsPtrInput
	// The unique target assignment ID.  If missing, will generate a random, unique id.
	TargetId pulumi.StringPtrInput
}

func (EventTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventTargetArgs)(nil)).Elem()
}
