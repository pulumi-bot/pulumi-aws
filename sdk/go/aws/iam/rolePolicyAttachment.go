// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iam

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Attaches a Managed IAM Policy to an IAM role
// 
// > **NOTE:** The usage of this resource conflicts with the `iam.PolicyAttachment` resource and will permanently show a difference if both are defined.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_role_policy_attachment.markdown.
type RolePolicyAttachment struct {
	pulumi.CustomResourceState

	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringOutput `pulumi:"policyArn"`
	// The role the policy should be applied to
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewRolePolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewRolePolicyAttachment(ctx *pulumi.Context,
	name string, args *RolePolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*RolePolicyAttachment, error) {
	if args == nil || args.PolicyArn == nil {
		return nil, errors.New("missing required argument 'PolicyArn'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &RolePolicyAttachmentArgs{}
	}
	var resource RolePolicyAttachment
	err := ctx.RegisterResource("aws:iam/rolePolicyAttachment:RolePolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRolePolicyAttachment gets an existing RolePolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRolePolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RolePolicyAttachmentState, opts ...pulumi.ResourceOption) (*RolePolicyAttachment, error) {
	var resource RolePolicyAttachment
	err := ctx.ReadResource("aws:iam/rolePolicyAttachment:RolePolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RolePolicyAttachment resources.
type rolePolicyAttachmentState struct {
	// The ARN of the policy you want to apply
	PolicyArn *string `pulumi:"policyArn"`
	// The role the policy should be applied to
	Role *string `pulumi:"role"`
}

type RolePolicyAttachmentState struct {
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringPtrInput
	// The role the policy should be applied to
	Role pulumi.StringPtrInput
}

func (RolePolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*rolePolicyAttachmentState)(nil)).Elem()
}

type rolePolicyAttachmentArgs struct {
	// The ARN of the policy you want to apply
	PolicyArn string `pulumi:"policyArn"`
	// The role the policy should be applied to
	Role interface{} `pulumi:"role"`
}

// The set of arguments for constructing a RolePolicyAttachment resource.
type RolePolicyAttachmentArgs struct {
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringInput
	// The role the policy should be applied to
	Role pulumi.Input
}

func (RolePolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rolePolicyAttachmentArgs)(nil)).Elem()
}

