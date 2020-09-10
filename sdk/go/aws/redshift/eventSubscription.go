// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type EventSubscription struct {
	pulumi.CustomResourceState

	Arn             pulumi.StringOutput      `pulumi:"arn"`
	CustomerAwsId   pulumi.StringOutput      `pulumi:"customerAwsId"`
	Enabled         pulumi.BoolPtrOutput     `pulumi:"enabled"`
	EventCategories pulumi.StringArrayOutput `pulumi:"eventCategories"`
	Name            pulumi.StringOutput      `pulumi:"name"`
	Severity        pulumi.StringPtrOutput   `pulumi:"severity"`
	SnsTopicArn     pulumi.StringOutput      `pulumi:"snsTopicArn"`
	SourceIds       pulumi.StringArrayOutput `pulumi:"sourceIds"`
	SourceType      pulumi.StringPtrOutput   `pulumi:"sourceType"`
	Status          pulumi.StringOutput      `pulumi:"status"`
	Tags            pulumi.StringMapOutput   `pulumi:"tags"`
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
	Arn             *string           `pulumi:"arn"`
	CustomerAwsId   *string           `pulumi:"customerAwsId"`
	Enabled         *bool             `pulumi:"enabled"`
	EventCategories []string          `pulumi:"eventCategories"`
	Name            *string           `pulumi:"name"`
	Severity        *string           `pulumi:"severity"`
	SnsTopicArn     *string           `pulumi:"snsTopicArn"`
	SourceIds       []string          `pulumi:"sourceIds"`
	SourceType      *string           `pulumi:"sourceType"`
	Status          *string           `pulumi:"status"`
	Tags            map[string]string `pulumi:"tags"`
}

type EventSubscriptionState struct {
	Arn             pulumi.StringPtrInput
	CustomerAwsId   pulumi.StringPtrInput
	Enabled         pulumi.BoolPtrInput
	EventCategories pulumi.StringArrayInput
	Name            pulumi.StringPtrInput
	Severity        pulumi.StringPtrInput
	SnsTopicArn     pulumi.StringPtrInput
	SourceIds       pulumi.StringArrayInput
	SourceType      pulumi.StringPtrInput
	Status          pulumi.StringPtrInput
	Tags            pulumi.StringMapInput
}

func (EventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionState)(nil)).Elem()
}

type eventSubscriptionArgs struct {
	Enabled         *bool             `pulumi:"enabled"`
	EventCategories []string          `pulumi:"eventCategories"`
	Name            *string           `pulumi:"name"`
	Severity        *string           `pulumi:"severity"`
	SnsTopicArn     string            `pulumi:"snsTopicArn"`
	SourceIds       []string          `pulumi:"sourceIds"`
	SourceType      *string           `pulumi:"sourceType"`
	Tags            map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EventSubscription resource.
type EventSubscriptionArgs struct {
	Enabled         pulumi.BoolPtrInput
	EventCategories pulumi.StringArrayInput
	Name            pulumi.StringPtrInput
	Severity        pulumi.StringPtrInput
	SnsTopicArn     pulumi.StringInput
	SourceIds       pulumi.StringArrayInput
	SourceType      pulumi.StringPtrInput
	Tags            pulumi.StringMapInput
}

func (EventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionArgs)(nil)).Elem()
}
