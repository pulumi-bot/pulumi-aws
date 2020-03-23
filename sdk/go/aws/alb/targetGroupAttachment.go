// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package alb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides the ability to register instances and containers with an Application Load Balancer (ALB) or Network Load Balancer (NLB) target group. For attaching resources with Elastic Load Balancer (ELB), see the [`elb.Attachment` resource](https://www.terraform.io/docs/providers/aws/r/elb_attachment.html).
//
// > **Note:** `alb.TargetGroupAttachment` is known as `lb.TargetGroupAttachment`. The functionality is identical.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_target_group_attachment.html.markdown.
type TargetGroupAttachment struct {
	pulumi.CustomResourceState

	// The Availability Zone where the IP address of the target is to be registered.
	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	// The port on which targets receive traffic.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The ARN of the target group with which to register targets
	TargetGroupArn pulumi.StringOutput `pulumi:"targetGroupArn"`
	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
	TargetId pulumi.StringOutput `pulumi:"targetId"`
}

// NewTargetGroupAttachment registers a new resource with the given unique name, arguments, and options.
func NewTargetGroupAttachment(ctx *pulumi.Context,
	name string, args *TargetGroupAttachmentArgs, opts ...pulumi.ResourceOption) (*TargetGroupAttachment, error) {
	if args == nil || args.TargetGroupArn == nil {
		return nil, errors.New("missing required argument 'TargetGroupArn'")
	}
	if args == nil || args.TargetId == nil {
		return nil, errors.New("missing required argument 'TargetId'")
	}
	if args == nil {
		args = &TargetGroupAttachmentArgs{}
	}
	var resource TargetGroupAttachment
	err := ctx.RegisterResource("aws:alb/targetGroupAttachment:TargetGroupAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetGroupAttachment gets an existing TargetGroupAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetGroupAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetGroupAttachmentState, opts ...pulumi.ResourceOption) (*TargetGroupAttachment, error) {
	var resource TargetGroupAttachment
	err := ctx.ReadResource("aws:alb/targetGroupAttachment:TargetGroupAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetGroupAttachment resources.
type targetGroupAttachmentState struct {
	// The Availability Zone where the IP address of the target is to be registered.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The port on which targets receive traffic.
	Port *int `pulumi:"port"`
	// The ARN of the target group with which to register targets
	TargetGroupArn *string `pulumi:"targetGroupArn"`
	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
	TargetId *string `pulumi:"targetId"`
}

type TargetGroupAttachmentState struct {
	// The Availability Zone where the IP address of the target is to be registered.
	AvailabilityZone pulumi.StringPtrInput
	// The port on which targets receive traffic.
	Port pulumi.IntPtrInput
	// The ARN of the target group with which to register targets
	TargetGroupArn pulumi.StringPtrInput
	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
	TargetId pulumi.StringPtrInput
}

func (TargetGroupAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGroupAttachmentState)(nil)).Elem()
}

type targetGroupAttachmentArgs struct {
	// The Availability Zone where the IP address of the target is to be registered.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The port on which targets receive traffic.
	Port *int `pulumi:"port"`
	// The ARN of the target group with which to register targets
	TargetGroupArn string `pulumi:"targetGroupArn"`
	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
	TargetId string `pulumi:"targetId"`
}

// The set of arguments for constructing a TargetGroupAttachment resource.
type TargetGroupAttachmentArgs struct {
	// The Availability Zone where the IP address of the target is to be registered.
	AvailabilityZone pulumi.StringPtrInput
	// The port on which targets receive traffic.
	Port pulumi.IntPtrInput
	// The ARN of the target group with which to register targets
	TargetGroupArn pulumi.StringInput
	// The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
	TargetId pulumi.StringInput
}

func (TargetGroupAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGroupAttachmentArgs)(nil)).Elem()
}

