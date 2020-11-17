// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an IoT policy attachment.
type PolicyAttachment struct {
	pulumi.CustomResourceState

	// The name of the policy to attach.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The identity to which the policy is attached.
	Target pulumi.StringOutput `pulumi:"target"`
}

// NewPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewPolicyAttachment(ctx *pulumi.Context,
	name string, args *PolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*PolicyAttachment, error) {
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil || args.Target == nil {
		return nil, errors.New("missing required argument 'Target'")
	}
	if args == nil {
		args = &PolicyAttachmentArgs{}
	}
	var resource PolicyAttachment
	err := ctx.RegisterResource("aws:iot/policyAttachment:PolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyAttachment gets an existing PolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyAttachmentState, opts ...pulumi.ResourceOption) (*PolicyAttachment, error) {
	var resource PolicyAttachment
	err := ctx.ReadResource("aws:iot/policyAttachment:PolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyAttachment resources.
type policyAttachmentState struct {
	// The name of the policy to attach.
	Policy *string `pulumi:"policy"`
	// The identity to which the policy is attached.
	Target *string `pulumi:"target"`
}

type PolicyAttachmentState struct {
	// The name of the policy to attach.
	Policy pulumi.StringPtrInput
	// The identity to which the policy is attached.
	Target pulumi.StringPtrInput
}

func (PolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAttachmentState)(nil)).Elem()
}

type policyAttachmentArgs struct {
	// The name of the policy to attach.
	Policy interface{} `pulumi:"policy"`
	// The identity to which the policy is attached.
	Target string `pulumi:"target"`
}

// The set of arguments for constructing a PolicyAttachment resource.
type PolicyAttachmentArgs struct {
	// The name of the policy to attach.
	Policy pulumi.Input
	// The identity to which the policy is attached.
	Target pulumi.StringInput
}

func (PolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAttachmentArgs)(nil)).Elem()
}

type PolicyAttachmentInput interface {
	pulumi.Input

	ToPolicyAttachmentOutput() PolicyAttachmentOutput
	ToPolicyAttachmentOutputWithContext(ctx context.Context) PolicyAttachmentOutput
}

func (PolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAttachment)(nil)).Elem()
}

func (i PolicyAttachment) ToPolicyAttachmentOutput() PolicyAttachmentOutput {
	return i.ToPolicyAttachmentOutputWithContext(context.Background())
}

func (i PolicyAttachment) ToPolicyAttachmentOutputWithContext(ctx context.Context) PolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAttachmentOutput)
}

type PolicyAttachmentOutput struct {
	*pulumi.OutputState
}

func (PolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAttachmentOutput)(nil)).Elem()
}

func (o PolicyAttachmentOutput) ToPolicyAttachmentOutput() PolicyAttachmentOutput {
	return o
}

func (o PolicyAttachmentOutput) ToPolicyAttachmentOutputWithContext(ctx context.Context) PolicyAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicyAttachmentOutput{})
}
