// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type TargetGroupAttachment struct {
	pulumi.CustomResourceState

	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	Port             pulumi.IntPtrOutput    `pulumi:"port"`
	TargetGroupArn   pulumi.StringOutput    `pulumi:"targetGroupArn"`
	TargetId         pulumi.StringOutput    `pulumi:"targetId"`
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("aws:applicationloadbalancing/targetGroupAttachment:TargetGroupAttachment"),
		},
	})
	opts = append(opts, aliases)
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
	AvailabilityZone *string `pulumi:"availabilityZone"`
	Port             *int    `pulumi:"port"`
	TargetGroupArn   *string `pulumi:"targetGroupArn"`
	TargetId         *string `pulumi:"targetId"`
}

type TargetGroupAttachmentState struct {
	AvailabilityZone pulumi.StringPtrInput
	Port             pulumi.IntPtrInput
	TargetGroupArn   pulumi.StringPtrInput
	TargetId         pulumi.StringPtrInput
}

func (TargetGroupAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGroupAttachmentState)(nil)).Elem()
}

type targetGroupAttachmentArgs struct {
	AvailabilityZone *string `pulumi:"availabilityZone"`
	Port             *int    `pulumi:"port"`
	TargetGroupArn   string  `pulumi:"targetGroupArn"`
	TargetId         string  `pulumi:"targetId"`
}

// The set of arguments for constructing a TargetGroupAttachment resource.
type TargetGroupAttachmentArgs struct {
	AvailabilityZone pulumi.StringPtrInput
	Port             pulumi.IntPtrInput
	TargetGroupArn   pulumi.StringInput
	TargetId         pulumi.StringInput
}

func (TargetGroupAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGroupAttachmentArgs)(nil)).Elem()
}
