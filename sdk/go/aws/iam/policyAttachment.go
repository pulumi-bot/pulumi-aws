// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iam

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Attaches a Managed IAM Policy to user(s), role(s), and/or group(s)
//
// !> **WARNING:** The iam.PolicyAttachment resource creates **exclusive** attachments of IAM policies. Across the entire AWS account, all of the users/roles/groups to which a single policy is attached must be declared by a single iam.PolicyAttachment resource. This means that even any users/roles/groups that have the attached policy via any other mechanism (including other resources managed by this provider) will have that attached policy revoked by this resource. Consider `iam.RolePolicyAttachment`, `iam.UserPolicyAttachment`, or `iam.GroupPolicyAttachment` instead. These resources do not enforce exclusive attachment of an IAM policy.
//
// > **NOTE:** The usage of this resource conflicts with the `iam.GroupPolicyAttachment`, `iam.RolePolicyAttachment`, and `iam.UserPolicyAttachment` resources and will permanently show a difference if both are defined.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_policy_attachment.html.markdown.
type PolicyAttachment struct {
	pulumi.CustomResourceState

	// The group(s) the policy should be applied to
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// The name of the attachment. This cannot be an empty string.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringOutput `pulumi:"policyArn"`
	// The role(s) the policy should be applied to
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The user(s) the policy should be applied to
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewPolicyAttachment(ctx *pulumi.Context,
	name string, args *PolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*PolicyAttachment, error) {
	if args == nil || args.PolicyArn == nil {
		return nil, errors.New("missing required argument 'PolicyArn'")
	}
	if args == nil {
		args = &PolicyAttachmentArgs{}
	}
	var resource PolicyAttachment
	err := ctx.RegisterResource("aws:iam/policyAttachment:PolicyAttachment", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:iam/policyAttachment:PolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyAttachment resources.
type policyAttachmentState struct {
	// The group(s) the policy should be applied to
	Groups []string `pulumi:"groups"`
	// The name of the attachment. This cannot be an empty string.
	Name *string `pulumi:"name"`
	// The ARN of the policy you want to apply
	PolicyArn *string `pulumi:"policyArn"`
	// The role(s) the policy should be applied to
	Roles []string `pulumi:"roles"`
	// The user(s) the policy should be applied to
	Users []string `pulumi:"users"`
}

type PolicyAttachmentState struct {
	// The group(s) the policy should be applied to
	Groups pulumi.StringArrayInput
	// The name of the attachment. This cannot be an empty string.
	Name pulumi.StringPtrInput
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringPtrInput
	// The role(s) the policy should be applied to
	Roles pulumi.StringArrayInput
	// The user(s) the policy should be applied to
	Users pulumi.StringArrayInput
}

func (PolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAttachmentState)(nil)).Elem()
}

type policyAttachmentArgs struct {
	// The group(s) the policy should be applied to
	Groups []interface{} `pulumi:"groups"`
	// The name of the attachment. This cannot be an empty string.
	Name *string `pulumi:"name"`
	// The ARN of the policy you want to apply
	PolicyArn string `pulumi:"policyArn"`
	// The role(s) the policy should be applied to
	Roles []interface{} `pulumi:"roles"`
	// The user(s) the policy should be applied to
	Users []interface{} `pulumi:"users"`
}

// The set of arguments for constructing a PolicyAttachment resource.
type PolicyAttachmentArgs struct {
	// The group(s) the policy should be applied to
	Groups pulumi.ArrayInput
	// The name of the attachment. This cannot be an empty string.
	Name pulumi.StringPtrInput
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringInput
	// The role(s) the policy should be applied to
	Roles pulumi.ArrayInput
	// The user(s) the policy should be applied to
	Users pulumi.ArrayInput
}

func (PolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAttachmentArgs)(nil)).Elem()
}

