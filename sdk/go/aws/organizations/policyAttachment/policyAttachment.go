// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package policyAttachment

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to attach an AWS Organizations policy to an organization account, root, or unit.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/organizations_policy_attachment.html.markdown.
type PolicyAttachment struct {
	pulumi.CustomResourceState

	// The unique identifier (ID) of the policy that you want to attach to the target.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
	TargetId pulumi.StringOutput `pulumi:"targetId"`
}

// NewPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewPolicyAttachment(ctx *pulumi.Context,
	name string, args *PolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*PolicyAttachment, error) {
	if args == nil || args.PolicyId == nil {
		return nil, errors.New("missing required argument 'PolicyId'")
	}
	if args == nil || args.TargetId == nil {
		return nil, errors.New("missing required argument 'TargetId'")
	}
	if args == nil {
		args = &PolicyAttachmentArgs{}
	}
	var resource PolicyAttachment
	err := ctx.RegisterResource("aws:organizations/policyAttachment:PolicyAttachment", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:organizations/policyAttachment:PolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyAttachment resources.
type policyAttachmentState struct {
	// The unique identifier (ID) of the policy that you want to attach to the target.
	PolicyId *string `pulumi:"policyId"`
	// The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
	TargetId *string `pulumi:"targetId"`
}

type PolicyAttachmentState struct {
	// The unique identifier (ID) of the policy that you want to attach to the target.
	PolicyId pulumi.StringPtrInput
	// The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
	TargetId pulumi.StringPtrInput
}

func (PolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAttachmentState)(nil)).Elem()
}

type policyAttachmentArgs struct {
	// The unique identifier (ID) of the policy that you want to attach to the target.
	PolicyId string `pulumi:"policyId"`
	// The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
	TargetId string `pulumi:"targetId"`
}

// The set of arguments for constructing a PolicyAttachment resource.
type PolicyAttachmentArgs struct {
	// The unique identifier (ID) of the policy that you want to attach to the target.
	PolicyId pulumi.StringInput
	// The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
	TargetId pulumi.StringInput
}

func (PolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAttachmentArgs)(nil)).Elem()
}

