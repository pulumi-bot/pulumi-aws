// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mediaconvert

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS Elemental MediaConvert Queue.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/mediaconvert"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mediaconvert.NewQueue(ctx, "test", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Queue struct {
	pulumi.CustomResourceState

	// The Arn of the queue
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description of the queue
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique identifier describing the queue
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies whether the pricing plan for the queue is on-demand or reserved. Valid values are `ON_DEMAND` or `RESERVED`. Default to `ON_DEMAND`.
	PricingPlan pulumi.StringPtrOutput `pulumi:"pricingPlan"`
	// A detail pricing plan of the  reserved queue. See below.
	ReservationPlanSettings QueueReservationPlanSettingsOutput `pulumi:"reservationPlanSettings"`
	// A status of the queue. Valid values are `ACTIVE` or `RESERVED`. Default to `PAUSED`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil {
		args = &QueueArgs{}
	}

	var resource Queue
	err := ctx.RegisterResource("aws:mediaconvert/queue:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("aws:mediaconvert/queue:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
	// The Arn of the queue
	Arn *string `pulumi:"arn"`
	// A description of the queue
	Description *string `pulumi:"description"`
	// A unique identifier describing the queue
	Name *string `pulumi:"name"`
	// Specifies whether the pricing plan for the queue is on-demand or reserved. Valid values are `ON_DEMAND` or `RESERVED`. Default to `ON_DEMAND`.
	PricingPlan *string `pulumi:"pricingPlan"`
	// A detail pricing plan of the  reserved queue. See below.
	ReservationPlanSettings *QueueReservationPlanSettings `pulumi:"reservationPlanSettings"`
	// A status of the queue. Valid values are `ACTIVE` or `RESERVED`. Default to `PAUSED`.
	Status *string `pulumi:"status"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type QueueState struct {
	// The Arn of the queue
	Arn pulumi.StringPtrInput
	// A description of the queue
	Description pulumi.StringPtrInput
	// A unique identifier describing the queue
	Name pulumi.StringPtrInput
	// Specifies whether the pricing plan for the queue is on-demand or reserved. Valid values are `ON_DEMAND` or `RESERVED`. Default to `ON_DEMAND`.
	PricingPlan pulumi.StringPtrInput
	// A detail pricing plan of the  reserved queue. See below.
	ReservationPlanSettings QueueReservationPlanSettingsPtrInput
	// A status of the queue. Valid values are `ACTIVE` or `RESERVED`. Default to `PAUSED`.
	Status pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// A description of the queue
	Description *string `pulumi:"description"`
	// A unique identifier describing the queue
	Name *string `pulumi:"name"`
	// Specifies whether the pricing plan for the queue is on-demand or reserved. Valid values are `ON_DEMAND` or `RESERVED`. Default to `ON_DEMAND`.
	PricingPlan *string `pulumi:"pricingPlan"`
	// A detail pricing plan of the  reserved queue. See below.
	ReservationPlanSettings *QueueReservationPlanSettings `pulumi:"reservationPlanSettings"`
	// A status of the queue. Valid values are `ACTIVE` or `RESERVED`. Default to `PAUSED`.
	Status *string `pulumi:"status"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// A description of the queue
	Description pulumi.StringPtrInput
	// A unique identifier describing the queue
	Name pulumi.StringPtrInput
	// Specifies whether the pricing plan for the queue is on-demand or reserved. Valid values are `ON_DEMAND` or `RESERVED`. Default to `ON_DEMAND`.
	PricingPlan pulumi.StringPtrInput
	// A detail pricing plan of the  reserved queue. See below.
	ReservationPlanSettings QueueReservationPlanSettingsPtrInput
	// A status of the queue. Valid values are `ACTIVE` or `RESERVED`. Default to `PAUSED`.
	Status pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}

type QueueInput interface {
	pulumi.Input

	ToQueueOutput() QueueOutput
	ToQueueOutputWithContext(ctx context.Context) QueueOutput
}

func (Queue) ElementType() reflect.Type {
	return reflect.TypeOf((*Queue)(nil)).Elem()
}

func (i Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

type QueueOutput struct {
	*pulumi.OutputState
}

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueOutput)(nil)).Elem()
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(QueueOutput{})
}
