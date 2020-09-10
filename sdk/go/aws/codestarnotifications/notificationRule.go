// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codestarnotifications

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type NotificationRule struct {
	pulumi.CustomResourceState

	Arn          pulumi.StringOutput               `pulumi:"arn"`
	DetailType   pulumi.StringOutput               `pulumi:"detailType"`
	EventTypeIds pulumi.StringArrayOutput          `pulumi:"eventTypeIds"`
	Name         pulumi.StringOutput               `pulumi:"name"`
	Resource     pulumi.StringOutput               `pulumi:"resource"`
	Status       pulumi.StringPtrOutput            `pulumi:"status"`
	Tags         pulumi.StringMapOutput            `pulumi:"tags"`
	Targets      NotificationRuleTargetArrayOutput `pulumi:"targets"`
}

// NewNotificationRule registers a new resource with the given unique name, arguments, and options.
func NewNotificationRule(ctx *pulumi.Context,
	name string, args *NotificationRuleArgs, opts ...pulumi.ResourceOption) (*NotificationRule, error) {
	if args == nil || args.DetailType == nil {
		return nil, errors.New("missing required argument 'DetailType'")
	}
	if args == nil || args.EventTypeIds == nil {
		return nil, errors.New("missing required argument 'EventTypeIds'")
	}
	if args == nil || args.Resource == nil {
		return nil, errors.New("missing required argument 'Resource'")
	}
	if args == nil {
		args = &NotificationRuleArgs{}
	}
	var resource NotificationRule
	err := ctx.RegisterResource("aws:codestarnotifications/notificationRule:NotificationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationRule gets an existing NotificationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationRuleState, opts ...pulumi.ResourceOption) (*NotificationRule, error) {
	var resource NotificationRule
	err := ctx.ReadResource("aws:codestarnotifications/notificationRule:NotificationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationRule resources.
type notificationRuleState struct {
	Arn          *string                  `pulumi:"arn"`
	DetailType   *string                  `pulumi:"detailType"`
	EventTypeIds []string                 `pulumi:"eventTypeIds"`
	Name         *string                  `pulumi:"name"`
	Resource     *string                  `pulumi:"resource"`
	Status       *string                  `pulumi:"status"`
	Tags         map[string]string        `pulumi:"tags"`
	Targets      []NotificationRuleTarget `pulumi:"targets"`
}

type NotificationRuleState struct {
	Arn          pulumi.StringPtrInput
	DetailType   pulumi.StringPtrInput
	EventTypeIds pulumi.StringArrayInput
	Name         pulumi.StringPtrInput
	Resource     pulumi.StringPtrInput
	Status       pulumi.StringPtrInput
	Tags         pulumi.StringMapInput
	Targets      NotificationRuleTargetArrayInput
}

func (NotificationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRuleState)(nil)).Elem()
}

type notificationRuleArgs struct {
	DetailType   string                   `pulumi:"detailType"`
	EventTypeIds []string                 `pulumi:"eventTypeIds"`
	Name         *string                  `pulumi:"name"`
	Resource     string                   `pulumi:"resource"`
	Status       *string                  `pulumi:"status"`
	Tags         map[string]string        `pulumi:"tags"`
	Targets      []NotificationRuleTarget `pulumi:"targets"`
}

// The set of arguments for constructing a NotificationRule resource.
type NotificationRuleArgs struct {
	DetailType   pulumi.StringInput
	EventTypeIds pulumi.StringArrayInput
	Name         pulumi.StringPtrInput
	Resource     pulumi.StringInput
	Status       pulumi.StringPtrInput
	Tags         pulumi.StringMapInput
	Targets      NotificationRuleTargetArrayInput
}

func (NotificationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRuleArgs)(nil)).Elem()
}
