// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/neptune"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultCluster, err := neptune.NewCluster(ctx, "defaultCluster", &neptune.ClusterArgs{
// 			ApplyImmediately:                 pulumi.Bool(true),
// 			BackupRetentionPeriod:            pulumi.Int(5),
// 			ClusterIdentifier:                pulumi.String("neptune-cluster-demo"),
// 			Engine:                           pulumi.String("neptune"),
// 			IamDatabaseAuthenticationEnabled: pulumi.Bool(true),
// 			PreferredBackupWindow:            pulumi.String("07:00-09:00"),
// 			SkipFinalSnapshot:                pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example, err := neptune.NewClusterInstance(ctx, "example", &neptune.ClusterInstanceArgs{
// 			ApplyImmediately:  pulumi.Bool(true),
// 			ClusterIdentifier: defaultCluster.ID(),
// 			Engine:            pulumi.String("neptune"),
// 			InstanceClass:     pulumi.String("db.r4.large"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultTopic, err := sns.NewTopic(ctx, "defaultTopic", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = neptune.NewEventSubscription(ctx, "defaultEventSubscription", &neptune.EventSubscriptionArgs{
// 			EventCategories: pulumi.StringArray{
// 				pulumi.String("maintenance"),
// 				pulumi.String("availability"),
// 				pulumi.String("creation"),
// 				pulumi.String("backup"),
// 				pulumi.String("restoration"),
// 				pulumi.String("recovery"),
// 				pulumi.String("deletion"),
// 				pulumi.String("failover"),
// 				pulumi.String("failure"),
// 				pulumi.String("notification"),
// 				pulumi.String("configuration change"),
// 				pulumi.String("read replica"),
// 			},
// 			SnsTopicArn: pulumi.String(defaultTopic.Arn),
// 			SourceIds: pulumi.StringArray{
// 				example.ID(),
// 			},
// 			SourceType: pulumi.String("db-instance"),
// 			Tags: map[string]interface{}{
// 				"env": "test",
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Attributes
//
// The following additional atttributes are provided:
//
// * `id` - The name of the Neptune event notification subscription.
// * `arn` - The Amazon Resource Name of the Neptune event notification subscription.
// * `customerAwsId` - The AWS customer account associated with the Neptune event notification subscription.
type EventSubscription struct {
	pulumi.CustomResourceState

	Arn           pulumi.StringOutput `pulumi:"arn"`
	CustomerAwsId pulumi.StringOutput `pulumi:"customerAwsId"`
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// A list of event categories for a `sourceType` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
	EventCategories pulumi.StringArrayOutput `pulumi:"eventCategories"`
	// The name of the Neptune event subscription. By default generated by this provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Neptune event subscription. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// The ARN of the SNS topic to send events to.
	SnsTopicArn pulumi.StringOutput `pulumi:"snsTopicArn"`
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `sourceType` must also be specified.
	SourceIds pulumi.StringArrayOutput `pulumi:"sourceIds"`
	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
	SourceType pulumi.StringPtrOutput `pulumi:"sourceType"`
	// A map of tags to assign to the resource.
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
	err := ctx.RegisterResource("aws:neptune/eventSubscription:EventSubscription", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:neptune/eventSubscription:EventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventSubscription resources.
type eventSubscriptionState struct {
	Arn           *string `pulumi:"arn"`
	CustomerAwsId *string `pulumi:"customerAwsId"`
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled *bool `pulumi:"enabled"`
	// A list of event categories for a `sourceType` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
	EventCategories []string `pulumi:"eventCategories"`
	// The name of the Neptune event subscription. By default generated by this provider.
	Name *string `pulumi:"name"`
	// The name of the Neptune event subscription. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The ARN of the SNS topic to send events to.
	SnsTopicArn *string `pulumi:"snsTopicArn"`
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `sourceType` must also be specified.
	SourceIds []string `pulumi:"sourceIds"`
	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
	SourceType *string `pulumi:"sourceType"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type EventSubscriptionState struct {
	Arn           pulumi.StringPtrInput
	CustomerAwsId pulumi.StringPtrInput
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled pulumi.BoolPtrInput
	// A list of event categories for a `sourceType` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
	EventCategories pulumi.StringArrayInput
	// The name of the Neptune event subscription. By default generated by this provider.
	Name pulumi.StringPtrInput
	// The name of the Neptune event subscription. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The ARN of the SNS topic to send events to.
	SnsTopicArn pulumi.StringPtrInput
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `sourceType` must also be specified.
	SourceIds pulumi.StringArrayInput
	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
	SourceType pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (EventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionState)(nil)).Elem()
}

type eventSubscriptionArgs struct {
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled *bool `pulumi:"enabled"`
	// A list of event categories for a `sourceType` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
	EventCategories []string `pulumi:"eventCategories"`
	// The name of the Neptune event subscription. By default generated by this provider.
	Name *string `pulumi:"name"`
	// The name of the Neptune event subscription. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The ARN of the SNS topic to send events to.
	SnsTopicArn string `pulumi:"snsTopicArn"`
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `sourceType` must also be specified.
	SourceIds []string `pulumi:"sourceIds"`
	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
	SourceType *string `pulumi:"sourceType"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a EventSubscription resource.
type EventSubscriptionArgs struct {
	// A boolean flag to enable/disable the subscription. Defaults to true.
	Enabled pulumi.BoolPtrInput
	// A list of event categories for a `sourceType` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
	EventCategories pulumi.StringArrayInput
	// The name of the Neptune event subscription. By default generated by this provider.
	Name pulumi.StringPtrInput
	// The name of the Neptune event subscription. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The ARN of the SNS topic to send events to.
	SnsTopicArn pulumi.StringInput
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `sourceType` must also be specified.
	SourceIds pulumi.StringArrayInput
	// The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
	SourceType pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (EventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionArgs)(nil)).Elem()
}
