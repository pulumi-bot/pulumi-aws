// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Attaches a Managed IAM Policy to an IAM user
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
// 		user, err := iam.NewUser(ctx, "user", nil)
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
// 		_, err = iam.NewUserPolicyAttachment(ctx, "test_attach", &iam.UserPolicyAttachmentArgs{
// 			User:      user.Name,
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
// IAM user policy attachments can be imported using the user name and policy arn separated by `/`.
//
// ```sh
//  $ pulumi import aws:iam/userPolicyAttachment:UserPolicyAttachment test-attach test-user/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
// ```
type UserPolicyAttachment struct {
	pulumi.CustomResourceState

	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringOutput `pulumi:"policyArn"`
	// The user the policy should be applied to
	User pulumi.StringOutput `pulumi:"user"`
}

// NewUserPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewUserPolicyAttachment(ctx *pulumi.Context,
	name string, args *UserPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*UserPolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyArn == nil {
		return nil, errors.New("invalid value for required argument 'PolicyArn'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	var resource UserPolicyAttachment
	err := ctx.RegisterResource("aws:iam/userPolicyAttachment:UserPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPolicyAttachment gets an existing UserPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPolicyAttachmentState, opts ...pulumi.ResourceOption) (*UserPolicyAttachment, error) {
	var resource UserPolicyAttachment
	err := ctx.ReadResource("aws:iam/userPolicyAttachment:UserPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPolicyAttachment resources.
type userPolicyAttachmentState struct {
	// The ARN of the policy you want to apply
	PolicyArn *string `pulumi:"policyArn"`
	// The user the policy should be applied to
	User *string `pulumi:"user"`
}

type UserPolicyAttachmentState struct {
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringPtrInput
	// The user the policy should be applied to
	User pulumi.StringPtrInput
}

func (UserPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPolicyAttachmentState)(nil)).Elem()
}

type userPolicyAttachmentArgs struct {
	// The ARN of the policy you want to apply
	PolicyArn string `pulumi:"policyArn"`
	// The user the policy should be applied to
	User interface{} `pulumi:"user"`
}

// The set of arguments for constructing a UserPolicyAttachment resource.
type UserPolicyAttachmentArgs struct {
	// The ARN of the policy you want to apply
	PolicyArn pulumi.StringInput
	// The user the policy should be applied to
	User pulumi.Input
}

func (UserPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPolicyAttachmentArgs)(nil)).Elem()
}

type UserPolicyAttachmentInput interface {
	pulumi.Input

	ToUserPolicyAttachmentOutput() UserPolicyAttachmentOutput
	ToUserPolicyAttachmentOutputWithContext(ctx context.Context) UserPolicyAttachmentOutput
}

type UserPolicyAttachmentPtrInput interface {
	pulumi.Input

	ToUserPolicyAttachmentPtrOutput() UserPolicyAttachmentPtrOutput
	ToUserPolicyAttachmentPtrOutputWithContext(ctx context.Context) UserPolicyAttachmentPtrOutput
}

func (UserPolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPolicyAttachment)(nil)).Elem()
}

func (i UserPolicyAttachment) ToUserPolicyAttachmentOutput() UserPolicyAttachmentOutput {
	return i.ToUserPolicyAttachmentOutputWithContext(context.Background())
}

func (i UserPolicyAttachment) ToUserPolicyAttachmentOutputWithContext(ctx context.Context) UserPolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyAttachmentOutput)
}

func (i UserPolicyAttachment) ToUserPolicyAttachmentPtrOutput() UserPolicyAttachmentPtrOutput {
	return i.ToUserPolicyAttachmentPtrOutputWithContext(context.Background())
}

func (i UserPolicyAttachment) ToUserPolicyAttachmentPtrOutputWithContext(ctx context.Context) UserPolicyAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyAttachmentPtrOutput)
}

type UserPolicyAttachmentOutput struct {
	*pulumi.OutputState
}

func (UserPolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPolicyAttachmentOutput)(nil)).Elem()
}

func (o UserPolicyAttachmentOutput) ToUserPolicyAttachmentOutput() UserPolicyAttachmentOutput {
	return o
}

func (o UserPolicyAttachmentOutput) ToUserPolicyAttachmentOutputWithContext(ctx context.Context) UserPolicyAttachmentOutput {
	return o
}

type UserPolicyAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (UserPolicyAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPolicyAttachment)(nil)).Elem()
}

func (o UserPolicyAttachmentPtrOutput) ToUserPolicyAttachmentPtrOutput() UserPolicyAttachmentPtrOutput {
	return o
}

func (o UserPolicyAttachmentPtrOutput) ToUserPolicyAttachmentPtrOutputWithContext(ctx context.Context) UserPolicyAttachmentPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UserPolicyAttachmentOutput{})
	pulumi.RegisterOutputType(UserPolicyAttachmentPtrOutput{})
}
