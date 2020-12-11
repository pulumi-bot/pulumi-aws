// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Attaches a Managed IAM Policy to an IAM role
//
// > **NOTE:** The usage of this resource conflicts with the `iam.PolicyAttachment` resource and will permanently show a difference if both are defined.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "    {\n", "      \"Version\": \"2012-10-17\",\n", "      \"Statement\": [\n", "        {\n", "          \"Action\": \"sts:AssumeRole\",\n", "          \"Principal\": {\n", "            \"Service\": \"ec2.amazonaws.com\"\n", "          },\n", "          \"Effect\": \"Allow\",\n", "          \"Sid\": \"\"\n", "        }\n", "      ]\n", "    }\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		policy, err := iam.NewPolicy(ctx, "policy", &iam.PolicyArgs{
// 			Description: pulumi.String("A test policy"),
// 			Policy:      pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"ec2:Describe*\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"*\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "test_attach", &iam.RolePolicyAttachmentArgs{
// 			Role:      role.Name,
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
// IAM role policy attachments can be imported using the role name and policy arn separated by `/`.
//
// ```sh
//  $ pulumi import aws:iam/rolePolicyAttachment:RolePolicyAttachment test-attach test-role/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
// ```
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyArn == nil {
		return nil, errors.New("invalid value for required argument 'PolicyArn'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
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

type RolePolicyAttachmentInput interface {
	pulumi.Input

	ToRolePolicyAttachmentOutput() RolePolicyAttachmentOutput
	ToRolePolicyAttachmentOutputWithContext(ctx context.Context) RolePolicyAttachmentOutput
}

type RolePolicyAttachmentPtrInput interface {
	pulumi.Input

	ToRolePolicyAttachmentPtrOutput() RolePolicyAttachmentPtrOutput
	ToRolePolicyAttachmentPtrOutputWithContext(ctx context.Context) RolePolicyAttachmentPtrOutput
}

func (RolePolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*RolePolicyAttachment)(nil)).Elem()
}

func (i RolePolicyAttachment) ToRolePolicyAttachmentOutput() RolePolicyAttachmentOutput {
	return i.ToRolePolicyAttachmentOutputWithContext(context.Background())
}

func (i RolePolicyAttachment) ToRolePolicyAttachmentOutputWithContext(ctx context.Context) RolePolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePolicyAttachmentOutput)
}

func (i RolePolicyAttachment) ToRolePolicyAttachmentPtrOutput() RolePolicyAttachmentPtrOutput {
	return i.ToRolePolicyAttachmentPtrOutputWithContext(context.Background())
}

func (i RolePolicyAttachment) ToRolePolicyAttachmentPtrOutputWithContext(ctx context.Context) RolePolicyAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePolicyAttachmentPtrOutput)
}

type RolePolicyAttachmentOutput struct {
	*pulumi.OutputState
}

func (RolePolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RolePolicyAttachmentOutput)(nil)).Elem()
}

func (o RolePolicyAttachmentOutput) ToRolePolicyAttachmentOutput() RolePolicyAttachmentOutput {
	return o
}

func (o RolePolicyAttachmentOutput) ToRolePolicyAttachmentOutputWithContext(ctx context.Context) RolePolicyAttachmentOutput {
	return o
}

type RolePolicyAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (RolePolicyAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RolePolicyAttachment)(nil)).Elem()
}

func (o RolePolicyAttachmentPtrOutput) ToRolePolicyAttachmentPtrOutput() RolePolicyAttachmentPtrOutput {
	return o
}

func (o RolePolicyAttachmentPtrOutput) ToRolePolicyAttachmentPtrOutputWithContext(ctx context.Context) RolePolicyAttachmentPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RolePolicyAttachmentOutput{})
	pulumi.RegisterOutputType(RolePolicyAttachmentPtrOutput{})
}
