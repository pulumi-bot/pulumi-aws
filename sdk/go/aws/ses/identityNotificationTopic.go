// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Resource for managing SES Identity Notification Topics
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewIdentityNotificationTopic(ctx, "test", &ses.IdentityNotificationTopicArgs{
// 			TopicArn:               pulumi.Any(aws_sns_topic.Example.Arn),
// 			NotificationType:       pulumi.String("Bounce"),
// 			Identity:               pulumi.Any(aws_ses_domain_identity.Example.Domain),
// 			IncludeOriginalHeaders: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IdentityNotificationTopic struct {
	pulumi.CustomResourceState

	// The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
	Identity pulumi.StringOutput `pulumi:"identity"`
	// Whether SES should include original email headers in SNS notifications of this type. *false* by default.
	IncludeOriginalHeaders pulumi.BoolPtrOutput `pulumi:"includeOriginalHeaders"`
	// The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: *Bounce*, *Complaint* or *Delivery*.
	NotificationType pulumi.StringOutput `pulumi:"notificationType"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
	TopicArn pulumi.StringPtrOutput `pulumi:"topicArn"`
}

// NewIdentityNotificationTopic registers a new resource with the given unique name, arguments, and options.
func NewIdentityNotificationTopic(ctx *pulumi.Context,
	name string, args *IdentityNotificationTopicArgs, opts ...pulumi.ResourceOption) (*IdentityNotificationTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.NotificationType == nil {
		return nil, errors.New("invalid value for required argument 'NotificationType'")
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
	// The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
	Identity *string `pulumi:"identity"`
	// Whether SES should include original email headers in SNS notifications of this type. *false* by default.
	IncludeOriginalHeaders *bool `pulumi:"includeOriginalHeaders"`
	// The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: *Bounce*, *Complaint* or *Delivery*.
	NotificationType *string `pulumi:"notificationType"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
	TopicArn *string `pulumi:"topicArn"`
}

type IdentityNotificationTopicState struct {
	// The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
	Identity pulumi.StringPtrInput
	// Whether SES should include original email headers in SNS notifications of this type. *false* by default.
	IncludeOriginalHeaders pulumi.BoolPtrInput
	// The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: *Bounce*, *Complaint* or *Delivery*.
	NotificationType pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
	TopicArn pulumi.StringPtrInput
}

func (IdentityNotificationTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityNotificationTopicState)(nil)).Elem()
}

type identityNotificationTopicArgs struct {
	// The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
	Identity string `pulumi:"identity"`
	// Whether SES should include original email headers in SNS notifications of this type. *false* by default.
	IncludeOriginalHeaders *bool `pulumi:"includeOriginalHeaders"`
	// The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: *Bounce*, *Complaint* or *Delivery*.
	NotificationType string `pulumi:"notificationType"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
	TopicArn *string `pulumi:"topicArn"`
}

// The set of arguments for constructing a IdentityNotificationTopic resource.
type IdentityNotificationTopicArgs struct {
	// The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
	Identity pulumi.StringInput
	// Whether SES should include original email headers in SNS notifications of this type. *false* by default.
	IncludeOriginalHeaders pulumi.BoolPtrInput
	// The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: *Bounce*, *Complaint* or *Delivery*.
	NotificationType pulumi.StringInput
	// The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to "" (an empty string) to disable publishing.
	TopicArn pulumi.StringPtrInput
}

func (IdentityNotificationTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityNotificationTopicArgs)(nil)).Elem()
}
