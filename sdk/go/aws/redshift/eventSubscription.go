// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package redshift

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Redshift event subscription resource.
//
// ## Attributes
//
// {{% examples %}}
//
// The following additional atttributes are provided:
//
// * `arn` - Amazon Resource Name (ARN) of the Redshift event notification subscription
// * `id` - The name of the Redshift event notification subscription
// * `customerAwsId` - The AWS customer account associated with the Redshift event notification subscription
//
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/redshift_event_subscription.html.markdown.
type EventSubscription struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	CustomerAwsId pulumi.StringOutput `pulumi:"customerAwsId"`
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html or run `aws redshift describe-event-categories`.
	EventCategories pulumi.StringArrayOutput `pulumi:"eventCategories"`
	// The name of the Redshift event subscription.
	Name pulumi.StringOutput `pulumi:"name"`
	// The event severity to be published by the notification subscription. Valid options are `INFO` or `ERROR`.
	Severity pulumi.StringPtrOutput `pulumi:"severity"`
	// The ARN of the SNS topic to send events to.
	SnsTopicArn pulumi.StringOutput `pulumi:"snsTopicArn"`
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
	SourceIds pulumi.StringArrayOutput `pulumi:"sourceIds"`
	// The type of source that will be generating the events. Valid options are `cluster`, `cluster-parameter-group`, `cluster-security-group`, or `cluster-snapshot`. If not set, all sources will be subscribed to.
	SourceType pulumi.StringPtrOutput `pulumi:"sourceType"`
	Status pulumi.StringOutput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewEventSubscription registers a new resource with the given unique name, arguments, and options.
func NewEventSubscription(ctx *pulumi.Context,
	name string, args *EventSubscriptionArgs, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	if args == nil || args.SnsTopicArn == nil {
		return nil, errors.New("missing required argument 'SnsTopicArn'")
	}
	if args == nil {
		args = &EventSubscriptionArgs{}
	}
	var resource EventSubscription
	err := ctx.RegisterResource("aws:redshift/eventSubscription:EventSubscription", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:redshift/eventSubscription:EventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventSubscription resources.
type eventSubscriptionState struct {
	Arn *string `pulumi:"arn"`
	CustomerAwsId *string `pulumi:"customerAwsId"`
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled *bool `pulumi:"enabled"`
	// A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html or run `aws redshift describe-event-categories`.
	EventCategories []string `pulumi:"eventCategories"`
	// The name of the Redshift event subscription.
	Name *string `pulumi:"name"`
	// The event severity to be published by the notification subscription. Valid options are `INFO` or `ERROR`.
	Severity *string `pulumi:"severity"`
	// The ARN of the SNS topic to send events to.
	SnsTopicArn *string `pulumi:"snsTopicArn"`
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
	SourceIds []string `pulumi:"sourceIds"`
	// The type of source that will be generating the events. Valid options are `cluster`, `cluster-parameter-group`, `cluster-security-group`, or `cluster-snapshot`. If not set, all sources will be subscribed to.
	SourceType *string `pulumi:"sourceType"`
	Status *string `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type EventSubscriptionState struct {
	Arn pulumi.StringPtrInput
	CustomerAwsId pulumi.StringPtrInput
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled pulumi.BoolPtrInput
	// A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html or run `aws redshift describe-event-categories`.
	EventCategories pulumi.StringArrayInput
	// The name of the Redshift event subscription.
	Name pulumi.StringPtrInput
	// The event severity to be published by the notification subscription. Valid options are `INFO` or `ERROR`.
	Severity pulumi.StringPtrInput
	// The ARN of the SNS topic to send events to.
	SnsTopicArn pulumi.StringPtrInput
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
	SourceIds pulumi.StringArrayInput
	// The type of source that will be generating the events. Valid options are `cluster`, `cluster-parameter-group`, `cluster-security-group`, or `cluster-snapshot`. If not set, all sources will be subscribed to.
	SourceType pulumi.StringPtrInput
	Status pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (EventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionState)(nil)).Elem()
}

type eventSubscriptionArgs struct {
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled *bool `pulumi:"enabled"`
	// A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html or run `aws redshift describe-event-categories`.
	EventCategories []string `pulumi:"eventCategories"`
	// The name of the Redshift event subscription.
	Name *string `pulumi:"name"`
	// The event severity to be published by the notification subscription. Valid options are `INFO` or `ERROR`.
	Severity *string `pulumi:"severity"`
	// The ARN of the SNS topic to send events to.
	SnsTopicArn string `pulumi:"snsTopicArn"`
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
	SourceIds []string `pulumi:"sourceIds"`
	// The type of source that will be generating the events. Valid options are `cluster`, `cluster-parameter-group`, `cluster-security-group`, or `cluster-snapshot`. If not set, all sources will be subscribed to.
	SourceType *string `pulumi:"sourceType"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a EventSubscription resource.
type EventSubscriptionArgs struct {
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled pulumi.BoolPtrInput
	// A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html or run `aws redshift describe-event-categories`.
	EventCategories pulumi.StringArrayInput
	// The name of the Redshift event subscription.
	Name pulumi.StringPtrInput
	// The event severity to be published by the notification subscription. Valid options are `INFO` or `ERROR`.
	Severity pulumi.StringPtrInput
	// The ARN of the SNS topic to send events to.
	SnsTopicArn pulumi.StringInput
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
	SourceIds pulumi.StringArrayInput
	// The type of source that will be generating the events. Valid options are `cluster`, `cluster-parameter-group`, `cluster-security-group`, or `cluster-snapshot`. If not set, all sources will be subscribed to.
	SourceType pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (EventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionArgs)(nil)).Elem()
}

