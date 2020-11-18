// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an EventBridge event bus resource.
//
// > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudwatch.NewEventBus(ctx, "messenger", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type EventBus struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the event bus.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the new event bus.
	// The names of custom event buses can't contain the / character.
	// Please note that a partner event bus is not supported at the moment.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewEventBus registers a new resource with the given unique name, arguments, and options.
func NewEventBus(ctx *pulumi.Context,
	name string, args *EventBusArgs, opts ...pulumi.ResourceOption) (*EventBus, error) {
	if args == nil {
		args = &EventBusArgs{}
	}

	var resource EventBus
	err := ctx.RegisterResource("aws:cloudwatch/eventBus:EventBus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventBus gets an existing EventBus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventBus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventBusState, opts ...pulumi.ResourceOption) (*EventBus, error) {
	var resource EventBus
	err := ctx.ReadResource("aws:cloudwatch/eventBus:EventBus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventBus resources.
type eventBusState struct {
	// The Amazon Resource Name (ARN) of the event bus.
	Arn *string `pulumi:"arn"`
	// The name of the new event bus.
	// The names of custom event buses can't contain the / character.
	// Please note that a partner event bus is not supported at the moment.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type EventBusState struct {
	// The Amazon Resource Name (ARN) of the event bus.
	Arn pulumi.StringPtrInput
	// The name of the new event bus.
	// The names of custom event buses can't contain the / character.
	// Please note that a partner event bus is not supported at the moment.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (EventBusState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventBusState)(nil)).Elem()
}

type eventBusArgs struct {
	// The name of the new event bus.
	// The names of custom event buses can't contain the / character.
	// Please note that a partner event bus is not supported at the moment.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EventBus resource.
type EventBusArgs struct {
	// The name of the new event bus.
	// The names of custom event buses can't contain the / character.
	// Please note that a partner event bus is not supported at the moment.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (EventBusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventBusArgs)(nil)).Elem()
}

type EventBusInput interface {
	pulumi.Input

	ToEventBusOutput() EventBusOutput
	ToEventBusOutputWithContext(ctx context.Context) EventBusOutput
}

func (EventBus) ElementType() reflect.Type {
	return reflect.TypeOf((*EventBus)(nil)).Elem()
}

func (i EventBus) ToEventBusOutput() EventBusOutput {
	return i.ToEventBusOutputWithContext(context.Background())
}

func (i EventBus) ToEventBusOutputWithContext(ctx context.Context) EventBusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventBusOutput)
}

type EventBusOutput struct {
	*pulumi.OutputState
}

func (EventBusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventBusOutput)(nil)).Elem()
}

func (o EventBusOutput) ToEventBusOutput() EventBusOutput {
	return o
}

func (o EventBusOutput) ToEventBusOutputWithContext(ctx context.Context) EventBusOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventBusOutput{})
}
