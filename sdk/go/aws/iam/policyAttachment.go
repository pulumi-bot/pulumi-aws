// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Attaches a Managed IAM Policy to user(s), role(s), and/or group(s)
// 
// !> **WARNING:** The iam.PolicyAttachment resource creates **exclusive** attachments of IAM policies. Across the entire AWS account, all of the users/roles/groups to which a single policy is attached must be declared by a single iam.PolicyAttachment resource. This means that even any users/roles/groups that have the attached policy via any other mechanism (including other resources managed by this provider) will have that attached policy revoked by this resource. Consider `iam.RolePolicyAttachment`, `iam.UserPolicyAttachment`, or `iam.GroupPolicyAttachment` instead. These resources do not enforce exclusive attachment of an IAM policy.
// 
// > **NOTE:** The usage of this resource conflicts with the `iam.GroupPolicyAttachment`, `iam.RolePolicyAttachment`, and `iam.UserPolicyAttachment` resources and will permanently show a difference if both are defined.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_policy_attachment.html.markdown.
type PolicyAttachment struct {
	s *pulumi.ResourceState
}

// NewPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewPolicyAttachment(ctx *pulumi.Context,
	name string, args *PolicyAttachmentArgs, opts ...pulumi.ResourceOpt) (*PolicyAttachment, error) {
	if args == nil || args.PolicyArn == nil {
		return nil, errors.New("missing required argument 'PolicyArn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["groups"] = nil
		inputs["name"] = nil
		inputs["policyArn"] = nil
		inputs["roles"] = nil
		inputs["users"] = nil
	} else {
		inputs["groups"] = args.Groups
		inputs["name"] = args.Name
		inputs["policyArn"] = args.PolicyArn
		inputs["roles"] = args.Roles
		inputs["users"] = args.Users
	}
	s, err := ctx.RegisterResource("aws:iam/policyAttachment:PolicyAttachment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PolicyAttachment{s: s}, nil
}

// GetPolicyAttachment gets an existing PolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PolicyAttachmentState, opts ...pulumi.ResourceOpt) (*PolicyAttachment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["groups"] = state.Groups
		inputs["name"] = state.Name
		inputs["policyArn"] = state.PolicyArn
		inputs["roles"] = state.Roles
		inputs["users"] = state.Users
	}
	s, err := ctx.ReadResource("aws:iam/policyAttachment:PolicyAttachment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PolicyAttachment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *PolicyAttachment) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *PolicyAttachment) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The group(s) the policy should be applied to
func (r *PolicyAttachment) Groups() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["groups"])
}

// The name of the attachment. This cannot be an empty string.
func (r *PolicyAttachment) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ARN of the policy you want to apply
func (r *PolicyAttachment) PolicyArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyArn"])
}

// The role(s) the policy should be applied to
func (r *PolicyAttachment) Roles() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["roles"])
}

// The user(s) the policy should be applied to
func (r *PolicyAttachment) Users() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["users"])
}

// Input properties used for looking up and filtering PolicyAttachment resources.
type PolicyAttachmentState struct {
	// The group(s) the policy should be applied to
	Groups interface{}
	// The name of the attachment. This cannot be an empty string.
	Name interface{}
	// The ARN of the policy you want to apply
	PolicyArn interface{}
	// The role(s) the policy should be applied to
	Roles interface{}
	// The user(s) the policy should be applied to
	Users interface{}
}

// The set of arguments for constructing a PolicyAttachment resource.
type PolicyAttachmentArgs struct {
	// The group(s) the policy should be applied to
	Groups interface{}
	// The name of the attachment. This cannot be an empty string.
	Name interface{}
	// The ARN of the policy you want to apply
	PolicyArn interface{}
	// The role(s) the policy should be applied to
	Roles interface{}
	// The user(s) the policy should be applied to
	Users interface{}
}
