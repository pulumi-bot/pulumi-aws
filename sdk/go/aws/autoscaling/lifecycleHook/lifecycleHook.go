// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package lifecycleHook

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AutoScaling Lifecycle Hook resource.
// 
// > **NOTE:** This provider has two types of ways you can add lifecycle hooks - via
// the `initialLifecycleHook` attribute from the
// [`autoscaling.Group`](https://www.terraform.io/docs/providers/aws/r/autoscaling_group.html)
// resource, or via this one. Hooks added via this resource will not be added
// until the autoscaling group has been created, and depending on your
// [capacity](https://www.terraform.io/docs/providers/aws/r/autoscaling_group.html#waiting-for-capacity)
// settings, after the initial instances have been launched, creating unintended
// behavior. If you need hooks to run on all instances, add them with
// `initialLifecycleHook` in
// [`autoscaling.Group`](https://www.terraform.io/docs/providers/aws/r/autoscaling_group.html),
// but take care to not duplicate those hooks with this resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/autoscaling_lifecycle_hook.html.markdown.
type LifecycleHook struct {
	pulumi.CustomResourceState

	// The name of the Auto Scaling group to which you want to assign the lifecycle hook
	AutoscalingGroupName pulumi.StringOutput `pulumi:"autoscalingGroupName"`
	// Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
	DefaultResult pulumi.StringOutput `pulumi:"defaultResult"`
	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
	HeartbeatTimeout pulumi.IntPtrOutput `pulumi:"heartbeatTimeout"`
	// The instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
	LifecycleTransition pulumi.StringOutput `pulumi:"lifecycleTransition"`
	// The name of the lifecycle hook.
	Name pulumi.StringOutput `pulumi:"name"`
	// Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
	NotificationMetadata pulumi.StringPtrOutput `pulumi:"notificationMetadata"`
	// The ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
	NotificationTargetArn pulumi.StringPtrOutput `pulumi:"notificationTargetArn"`
	// The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
}

// NewLifecycleHook registers a new resource with the given unique name, arguments, and options.
func NewLifecycleHook(ctx *pulumi.Context,
	name string, args *LifecycleHookArgs, opts ...pulumi.ResourceOption) (*LifecycleHook, error) {
	if args == nil || args.AutoscalingGroupName == nil {
		return nil, errors.New("missing required argument 'AutoscalingGroupName'")
	}
	if args == nil || args.LifecycleTransition == nil {
		return nil, errors.New("missing required argument 'LifecycleTransition'")
	}
	if args == nil {
		args = &LifecycleHookArgs{}
	}
	var resource LifecycleHook
	err := ctx.RegisterResource("aws:autoscaling/lifecycleHook:LifecycleHook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLifecycleHook gets an existing LifecycleHook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLifecycleHook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LifecycleHookState, opts ...pulumi.ResourceOption) (*LifecycleHook, error) {
	var resource LifecycleHook
	err := ctx.ReadResource("aws:autoscaling/lifecycleHook:LifecycleHook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LifecycleHook resources.
type lifecycleHookState struct {
	// The name of the Auto Scaling group to which you want to assign the lifecycle hook
	AutoscalingGroupName *string `pulumi:"autoscalingGroupName"`
	// Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
	DefaultResult *string `pulumi:"defaultResult"`
	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
	HeartbeatTimeout *int `pulumi:"heartbeatTimeout"`
	// The instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
	LifecycleTransition *string `pulumi:"lifecycleTransition"`
	// The name of the lifecycle hook.
	Name *string `pulumi:"name"`
	// Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
	NotificationMetadata *string `pulumi:"notificationMetadata"`
	// The ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
	NotificationTargetArn *string `pulumi:"notificationTargetArn"`
	// The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
	RoleArn *string `pulumi:"roleArn"`
}

type LifecycleHookState struct {
	// The name of the Auto Scaling group to which you want to assign the lifecycle hook
	AutoscalingGroupName pulumi.StringPtrInput
	// Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
	DefaultResult pulumi.StringPtrInput
	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
	HeartbeatTimeout pulumi.IntPtrInput
	// The instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
	LifecycleTransition pulumi.StringPtrInput
	// The name of the lifecycle hook.
	Name pulumi.StringPtrInput
	// Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
	NotificationMetadata pulumi.StringPtrInput
	// The ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
	NotificationTargetArn pulumi.StringPtrInput
	// The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
	RoleArn pulumi.StringPtrInput
}

func (LifecycleHookState) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecycleHookState)(nil)).Elem()
}

type lifecycleHookArgs struct {
	// The name of the Auto Scaling group to which you want to assign the lifecycle hook
	AutoscalingGroupName string `pulumi:"autoscalingGroupName"`
	// Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
	DefaultResult *string `pulumi:"defaultResult"`
	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
	HeartbeatTimeout *int `pulumi:"heartbeatTimeout"`
	// The instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
	LifecycleTransition string `pulumi:"lifecycleTransition"`
	// The name of the lifecycle hook.
	Name *string `pulumi:"name"`
	// Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
	NotificationMetadata *string `pulumi:"notificationMetadata"`
	// The ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
	NotificationTargetArn *string `pulumi:"notificationTargetArn"`
	// The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
	RoleArn *string `pulumi:"roleArn"`
}

// The set of arguments for constructing a LifecycleHook resource.
type LifecycleHookArgs struct {
	// The name of the Auto Scaling group to which you want to assign the lifecycle hook
	AutoscalingGroupName pulumi.StringInput
	// Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
	DefaultResult pulumi.StringPtrInput
	// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
	HeartbeatTimeout pulumi.IntPtrInput
	// The instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
	LifecycleTransition pulumi.StringInput
	// The name of the lifecycle hook.
	Name pulumi.StringPtrInput
	// Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
	NotificationMetadata pulumi.StringPtrInput
	// The ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
	NotificationTargetArn pulumi.StringPtrInput
	// The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
	RoleArn pulumi.StringPtrInput
}

func (LifecycleHookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecycleHookArgs)(nil)).Elem()
}

