// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DMS (Data Migration Service) event subscription resource.
//
// ## Import
//
// Event subscriptions can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:dms/eventSubscription:EventSubscription test my-awesome-event-subscription
// ```
type EventSubscription struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether the event subscription should be enabled.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of event categories to listen for, see `DescribeEventCategories` for a canonical list.
	EventCategories pulumi.StringArrayOutput `pulumi:"eventCategories"`
	// Name of event subscription.
	Name pulumi.StringOutput `pulumi:"name"`
	// SNS topic arn to send events on.
	SnsTopicArn pulumi.StringOutput `pulumi:"snsTopicArn"`
	// Ids of sources to listen to.
	SourceIds pulumi.StringArrayOutput `pulumi:"sourceIds"`
	// Type of source for events. Valid values: `replication-instance` or `replication-task`
	SourceType pulumi.StringPtrOutput `pulumi:"sourceType"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
}

// NewEventSubscription registers a new resource with the given unique name, arguments, and options.
func NewEventSubscription(ctx *pulumi.Context,
	name string, args *EventSubscriptionArgs, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	if args == nil || args.EventCategories == nil {
		return nil, errors.New("missing required argument 'EventCategories'")
	}
	if args == nil || args.SnsTopicArn == nil {
		return nil, errors.New("missing required argument 'SnsTopicArn'")
	}
	if args == nil {
		args = &EventSubscriptionArgs{}
	}
	var resource EventSubscription
	err := ctx.RegisterResource("aws:dms/eventSubscription:EventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventSubscription gets an existing EventSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventSubscriptionState, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	var resource EventSubscription
	err := ctx.ReadResource("aws:dms/eventSubscription:EventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventSubscription resources.
type eventSubscriptionState struct {
	Arn *string `pulumi:"arn"`
	// Whether the event subscription should be enabled.
	Enabled *bool `pulumi:"enabled"`
	// List of event categories to listen for, see `DescribeEventCategories` for a canonical list.
	EventCategories []string `pulumi:"eventCategories"`
	// Name of event subscription.
	Name *string `pulumi:"name"`
	// SNS topic arn to send events on.
	SnsTopicArn *string `pulumi:"snsTopicArn"`
	// Ids of sources to listen to.
	SourceIds []string `pulumi:"sourceIds"`
	// Type of source for events. Valid values: `replication-instance` or `replication-task`
	SourceType *string           `pulumi:"sourceType"`
	Tags       map[string]string `pulumi:"tags"`
}

type EventSubscriptionState struct {
	Arn pulumi.StringPtrInput
	// Whether the event subscription should be enabled.
	Enabled pulumi.BoolPtrInput
	// List of event categories to listen for, see `DescribeEventCategories` for a canonical list.
	EventCategories pulumi.StringArrayInput
	// Name of event subscription.
	Name pulumi.StringPtrInput
	// SNS topic arn to send events on.
	SnsTopicArn pulumi.StringPtrInput
	// Ids of sources to listen to.
	SourceIds pulumi.StringArrayInput
	// Type of source for events. Valid values: `replication-instance` or `replication-task`
	SourceType pulumi.StringPtrInput
	Tags       pulumi.StringMapInput
}

func (EventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionState)(nil)).Elem()
}

type eventSubscriptionArgs struct {
	// Whether the event subscription should be enabled.
	Enabled *bool `pulumi:"enabled"`
	// List of event categories to listen for, see `DescribeEventCategories` for a canonical list.
	EventCategories []string `pulumi:"eventCategories"`
	// Name of event subscription.
	Name *string `pulumi:"name"`
	// SNS topic arn to send events on.
	SnsTopicArn string `pulumi:"snsTopicArn"`
	// Ids of sources to listen to.
	SourceIds []string `pulumi:"sourceIds"`
	// Type of source for events. Valid values: `replication-instance` or `replication-task`
	SourceType *string           `pulumi:"sourceType"`
	Tags       map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EventSubscription resource.
type EventSubscriptionArgs struct {
	// Whether the event subscription should be enabled.
	Enabled pulumi.BoolPtrInput
	// List of event categories to listen for, see `DescribeEventCategories` for a canonical list.
	EventCategories pulumi.StringArrayInput
	// Name of event subscription.
	Name pulumi.StringPtrInput
	// SNS topic arn to send events on.
	SnsTopicArn pulumi.StringInput
	// Ids of sources to listen to.
	SourceIds pulumi.StringArrayInput
	// Type of source for events. Valid values: `replication-instance` or `replication-task`
	SourceType pulumi.StringPtrInput
	Tags       pulumi.StringMapInput
}

func (EventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionArgs)(nil)).Elem()
}

type EventSubscriptionInput interface {
	pulumi.Input

	ToEventSubscriptionOutput() EventSubscriptionOutput
	ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput
}

func (EventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscription)(nil)).Elem()
}

func (i EventSubscription) ToEventSubscriptionOutput() EventSubscriptionOutput {
	return i.ToEventSubscriptionOutputWithContext(context.Background())
}

func (i EventSubscription) ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionOutput)
}

type EventSubscriptionOutput struct {
	*pulumi.OutputState
}

func (EventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionOutput)(nil)).Elem()
}

func (o EventSubscriptionOutput) ToEventSubscriptionOutput() EventSubscriptionOutput {
	return o
}

func (o EventSubscriptionOutput) ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventSubscriptionOutput{})
}
