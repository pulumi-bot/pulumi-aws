// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Attaches a Managed IAM Policy to an IAM group
//
// > **NOTE:** The usage of this resource conflicts with the `iam.PolicyAttachment` resource and will permanently show a difference if both are defined.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		group, err := iam.NewGroup(ctx, "group", nil)
// 		if err != nil {
// 			return err
// 		}
// 		policy, err := iam.NewPolicy(ctx, "policy", &iam.PolicyArgs{
// 			Description: pulumi.String("A test policy"),
// 			Policy:      pulumi.String("{ ... policy JSON ... }"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewGroupPolicyAttachment(ctx, "test_attach", &iam.GroupPolicyAttachmentArgs{
// 			Group:     group.Name,
// 			PolicyArn: policy.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// IAM group policy attachments can be imported using the group name and policy arn separated by `/`.
//
// ```sh
//  $ pulumi import aws:iam/groupPolicyAttachment:GroupPolicyAttachment test-attach test-group/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
// ```
type GroupPolicyAttachment struct {
	pulumi.CustomResourceState

	// The group the policy should be applied to
	Group pulumi.StringOutput `pulumi:"group"`
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringOutput `pulumi:"policyArn"`
}

// NewGroupPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewGroupPolicyAttachment(ctx *pulumi.Context,
	name string, args *GroupPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*GroupPolicyAttachment, error) {
	if args == nil || args.Group == nil {
		return nil, errors.New("missing required argument 'Group'")
	}
	if args == nil || args.PolicyArn == nil {
		return nil, errors.New("missing required argument 'PolicyArn'")
	}
	if args == nil {
		args = &GroupPolicyAttachmentArgs{}
	}
	var resource GroupPolicyAttachment
	err := ctx.RegisterResource("aws:iam/groupPolicyAttachment:GroupPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupPolicyAttachment gets an existing GroupPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupPolicyAttachmentState, opts ...pulumi.ResourceOption) (*GroupPolicyAttachment, error) {
	var resource GroupPolicyAttachment
	err := ctx.ReadResource("aws:iam/groupPolicyAttachment:GroupPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupPolicyAttachment resources.
type groupPolicyAttachmentState struct {
	// The group the policy should be applied to
	Group *string `pulumi:"group"`
	// The ARN of the policy you want to apply
	PolicyArn *string `pulumi:"policyArn"`
}

type GroupPolicyAttachmentState struct {
	// The group the policy should be applied to
	Group pulumi.StringPtrInput
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringPtrInput
}

func (GroupPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyAttachmentState)(nil)).Elem()
}

type groupPolicyAttachmentArgs struct {
	// The group the policy should be applied to
	Group interface{} `pulumi:"group"`
	// The ARN of the policy you want to apply
	PolicyArn string `pulumi:"policyArn"`
}

// The set of arguments for constructing a GroupPolicyAttachment resource.
type GroupPolicyAttachmentArgs struct {
	// The group the policy should be applied to
	Group pulumi.Input
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringInput
}

func (GroupPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyAttachmentArgs)(nil)).Elem()
}
