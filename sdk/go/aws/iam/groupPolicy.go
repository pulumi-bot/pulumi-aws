// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an IAM policy attached to a group.
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
// 		myDevelopers, err := iam.NewGroup(ctx, "myDevelopers", &iam.GroupArgs{
// 			Path: pulumi.String("/users/"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewGroupPolicy(ctx, "myDeveloperPolicy", &iam.GroupPolicyArgs{
// 			Group:  myDevelopers.Name,
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"ec2:Describe*\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"*\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type GroupPolicy struct {
	pulumi.CustomResourceState

	// The IAM group to attach to the policy.
	Group pulumi.StringOutput `pulumi:"group"`
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The policy document. This is a JSON formatted string.
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewGroupPolicy registers a new resource with the given unique name, arguments, and options.
func NewGroupPolicy(ctx *pulumi.Context,
	name string, args *GroupPolicyArgs, opts ...pulumi.ResourceOption) (*GroupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource GroupPolicy
	err := ctx.RegisterResource("aws:iam/groupPolicy:GroupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupPolicy gets an existing GroupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupPolicyState, opts ...pulumi.ResourceOption) (*GroupPolicy, error) {
	var resource GroupPolicy
	err := ctx.ReadResource("aws:iam/groupPolicy:GroupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupPolicy resources.
type groupPolicyState struct {
	// The IAM group to attach to the policy.
	Group *string `pulumi:"group"`
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The policy document. This is a JSON formatted string.
	Policy *string `pulumi:"policy"`
}

type GroupPolicyState struct {
	// The IAM group to attach to the policy.
	Group pulumi.StringPtrInput
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The policy document. This is a JSON formatted string.
	Policy pulumi.StringPtrInput
}

func (GroupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyState)(nil)).Elem()
}

type groupPolicyArgs struct {
	// The IAM group to attach to the policy.
	Group string `pulumi:"group"`
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The policy document. This is a JSON formatted string.
	Policy interface{} `pulumi:"policy"`
}

// The set of arguments for constructing a GroupPolicy resource.
type GroupPolicyArgs struct {
	// The IAM group to attach to the policy.
	Group pulumi.StringInput
	// The name of the policy. If omitted, this provider will
	// assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The policy document. This is a JSON formatted string.
	Policy pulumi.Input
}

func (GroupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPolicyArgs)(nil)).Elem()
}
