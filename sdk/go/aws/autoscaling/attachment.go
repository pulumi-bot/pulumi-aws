// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AutoScaling Attachment resource.
// 
// ~> **NOTE on AutoScaling Groups and ASG Attachments:** Terraform currently provides
// both a standalone ASG Attachment resource (describing an ASG attached to
// an ELB), and an AutoScaling Group resource with
// `load_balancers` defined in-line. At this time you cannot use an ASG with in-line
// load balancers in conjunction with an ASG Attachment resource. Doing so will cause a
// conflict and will overwrite attachments.
type Attachment struct {
	s *pulumi.ResourceState
}

// NewAttachment registers a new resource with the given unique name, arguments, and options.
func NewAttachment(ctx *pulumi.Context,
	name string, args *AttachmentArgs, opts ...pulumi.ResourceOpt) (*Attachment, error) {
	if args == nil || args.AutoscalingGroupName == nil {
		return nil, errors.New("missing required argument 'AutoscalingGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["albTargetGroupArn"] = nil
		inputs["autoscalingGroupName"] = nil
		inputs["elb"] = nil
	} else {
		inputs["albTargetGroupArn"] = args.AlbTargetGroupArn
		inputs["autoscalingGroupName"] = args.AutoscalingGroupName
		inputs["elb"] = args.Elb
	}
	s, err := ctx.RegisterResource("aws:autoscaling/attachment:Attachment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Attachment{s: s}, nil
}

// GetAttachment gets an existing Attachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AttachmentState, opts ...pulumi.ResourceOpt) (*Attachment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["albTargetGroupArn"] = state.AlbTargetGroupArn
		inputs["autoscalingGroupName"] = state.AutoscalingGroupName
		inputs["elb"] = state.Elb
	}
	s, err := ctx.ReadResource("aws:autoscaling/attachment:Attachment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Attachment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Attachment) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Attachment) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of an ALB Target Group.
func (r *Attachment) AlbTargetGroupArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["albTargetGroupArn"])
}

// Name of ASG to associate with the ELB.
func (r *Attachment) AutoscalingGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["autoscalingGroupName"])
}

// The name of the ELB.
func (r *Attachment) Elb() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["elb"])
}

// Input properties used for looking up and filtering Attachment resources.
type AttachmentState struct {
	// The ARN of an ALB Target Group.
	AlbTargetGroupArn interface{}
	// Name of ASG to associate with the ELB.
	AutoscalingGroupName interface{}
	// The name of the ELB.
	Elb interface{}
}

// The set of arguments for constructing a Attachment resource.
type AttachmentArgs struct {
	// The ARN of an ALB Target Group.
	AlbTargetGroupArn interface{}
	// Name of ASG to associate with the ELB.
	AutoscalingGroupName interface{}
	// The name of the ELB.
	Elb interface{}
}
