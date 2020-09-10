// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type IdentityNotificationTopic struct {
	pulumi.CustomResourceState

	Identity               pulumi.StringOutput    `pulumi:"identity"`
	IncludeOriginalHeaders pulumi.BoolPtrOutput   `pulumi:"includeOriginalHeaders"`
	NotificationType       pulumi.StringOutput    `pulumi:"notificationType"`
	TopicArn               pulumi.StringPtrOutput `pulumi:"topicArn"`
}

// NewIdentityNotificationTopic registers a new resource with the given unique name, arguments, and options.
func NewIdentityNotificationTopic(ctx *pulumi.Context,
	name string, args *IdentityNotificationTopicArgs, opts ...pulumi.ResourceOption) (*IdentityNotificationTopic, error) {
	if args == nil || args.Identity == nil {
		return nil, errors.New("missing required argument 'Identity'")
	}
	if args == nil || args.NotificationType == nil {
		return nil, errors.New("missing required argument 'NotificationType'")
	}
	if args == nil {
		args = &IdentityNotificationTopicArgs{}
	}
	var resource IdentityNotificationTopic
	err := ctx.RegisterResource("aws:ses/identityNotificationTopic:IdentityNotificationTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityNotificationTopic gets an existing IdentityNotificationTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityNotificationTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityNotificationTopicState, opts ...pulumi.ResourceOption) (*IdentityNotificationTopic, error) {
	var resource IdentityNotificationTopic
	err := ctx.ReadResource("aws:ses/identityNotificationTopic:IdentityNotificationTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityNotificationTopic resources.
type identityNotificationTopicState struct {
	Identity               *string `pulumi:"identity"`
	IncludeOriginalHeaders *bool   `pulumi:"includeOriginalHeaders"`
	NotificationType       *string `pulumi:"notificationType"`
	TopicArn               *string `pulumi:"topicArn"`
}

type IdentityNotificationTopicState struct {
	Identity               pulumi.StringPtrInput
	IncludeOriginalHeaders pulumi.BoolPtrInput
	NotificationType       pulumi.StringPtrInput
	TopicArn               pulumi.StringPtrInput
}

func (IdentityNotificationTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityNotificationTopicState)(nil)).Elem()
}

type identityNotificationTopicArgs struct {
	Identity               string  `pulumi:"identity"`
	IncludeOriginalHeaders *bool   `pulumi:"includeOriginalHeaders"`
	NotificationType       string  `pulumi:"notificationType"`
	TopicArn               *string `pulumi:"topicArn"`
}

// The set of arguments for constructing a IdentityNotificationTopic resource.
type IdentityNotificationTopicArgs struct {
	Identity               pulumi.StringInput
	IncludeOriginalHeaders pulumi.BoolPtrInput
	NotificationType       pulumi.StringInput
	TopicArn               pulumi.StringPtrInput
}

func (IdentityNotificationTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityNotificationTopicArgs)(nil)).Elem()
}
