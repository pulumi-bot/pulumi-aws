// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package autoscaling

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AutoScaling Attachment resource.
//
// > **NOTE on AutoScaling Groups and ASG Attachments:** This provider currently provides
// both a standalone ASG Attachment resource (describing an ASG attached to
// an ELB), and an AutoScaling Group resource with
// `loadBalancers` defined in-line. At this time you cannot use an ASG with in-line
// load balancers in conjunction with an ASG Attachment resource. Doing so will cause a
// conflict and will overwrite attachments.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/autoscaling_attachment.html.markdown.
type Attachment struct {
	pulumi.CustomResourceState

	// The ARN of an ALB Target Group.
	AlbTargetGroupArn pulumi.StringPtrOutput `pulumi:"albTargetGroupArn"`
	// Name of ASG to associate with the ELB.
	AutoscalingGroupName pulumi.StringOutput `pulumi:"autoscalingGroupName"`
	// The name of the ELB.
	Elb pulumi.StringPtrOutput `pulumi:"elb"`
}

// NewAttachment registers a new resource with the given unique name, arguments, and options.
func NewAttachment(ctx *pulumi.Context,
	name string, args *AttachmentArgs, opts ...pulumi.ResourceOption) (*Attachment, error) {
	if args == nil || args.AutoscalingGroupName == nil {
		return nil, errors.New("missing required argument 'AutoscalingGroupName'")
	}
	if args == nil {
		args = &AttachmentArgs{}
	}
	var resource Attachment
	err := ctx.RegisterResource("aws:autoscaling/attachment:Attachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachment gets an existing Attachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachmentState, opts ...pulumi.ResourceOption) (*Attachment, error) {
	var resource Attachment
	err := ctx.ReadResource("aws:autoscaling/attachment:Attachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Attachment resources.
type attachmentState struct {
	// The ARN of an ALB Target Group.
	AlbTargetGroupArn *string `pulumi:"albTargetGroupArn"`
	// Name of ASG to associate with the ELB.
	AutoscalingGroupName *string `pulumi:"autoscalingGroupName"`
	// The name of the ELB.
	Elb *string `pulumi:"elb"`
}

type AttachmentState struct {
	// The ARN of an ALB Target Group.
	AlbTargetGroupArn pulumi.StringPtrInput
	// Name of ASG to associate with the ELB.
	AutoscalingGroupName pulumi.StringPtrInput
	// The name of the ELB.
	Elb pulumi.StringPtrInput
}

func (AttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentState)(nil)).Elem()
}

type attachmentArgs struct {
	// The ARN of an ALB Target Group.
	AlbTargetGroupArn *string `pulumi:"albTargetGroupArn"`
	// Name of ASG to associate with the ELB.
	AutoscalingGroupName string `pulumi:"autoscalingGroupName"`
	// The name of the ELB.
	Elb *string `pulumi:"elb"`
}

// The set of arguments for constructing a Attachment resource.
type AttachmentArgs struct {
	// The ARN of an ALB Target Group.
	AlbTargetGroupArn pulumi.StringPtrInput
	// Name of ASG to associate with the ELB.
	AutoscalingGroupName pulumi.StringInput
	// The name of the ELB.
	Elb pulumi.StringPtrInput
}

func (AttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentArgs)(nil)).Elem()
}

