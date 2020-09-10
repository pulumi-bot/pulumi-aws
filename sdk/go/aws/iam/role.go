// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Role struct {
	pulumi.CustomResourceState

	Arn                 pulumi.StringOutput    `pulumi:"arn"`
	AssumeRolePolicy    pulumi.StringOutput    `pulumi:"assumeRolePolicy"`
	CreateDate          pulumi.StringOutput    `pulumi:"createDate"`
	Description         pulumi.StringPtrOutput `pulumi:"description"`
	ForceDetachPolicies pulumi.BoolPtrOutput   `pulumi:"forceDetachPolicies"`
	MaxSessionDuration  pulumi.IntPtrOutput    `pulumi:"maxSessionDuration"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	NamePrefix          pulumi.StringPtrOutput `pulumi:"namePrefix"`
	Path                pulumi.StringPtrOutput `pulumi:"path"`
	PermissionsBoundary pulumi.StringPtrOutput `pulumi:"permissionsBoundary"`
	Tags                pulumi.StringMapOutput `pulumi:"tags"`
	UniqueId            pulumi.StringOutput    `pulumi:"uniqueId"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil || args.AssumeRolePolicy == nil {
		return nil, errors.New("missing required argument 'AssumeRolePolicy'")
	}
	if args == nil {
		args = &RoleArgs{}
	}
	var resource Role
	err := ctx.RegisterResource("aws:iam/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("aws:iam/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	Arn                 *string           `pulumi:"arn"`
	AssumeRolePolicy    *string           `pulumi:"assumeRolePolicy"`
	CreateDate          *string           `pulumi:"createDate"`
	Description         *string           `pulumi:"description"`
	ForceDetachPolicies *bool             `pulumi:"forceDetachPolicies"`
	MaxSessionDuration  *int              `pulumi:"maxSessionDuration"`
	Name                *string           `pulumi:"name"`
	NamePrefix          *string           `pulumi:"namePrefix"`
	Path                *string           `pulumi:"path"`
	PermissionsBoundary *string           `pulumi:"permissionsBoundary"`
	Tags                map[string]string `pulumi:"tags"`
	UniqueId            *string           `pulumi:"uniqueId"`
}

type RoleState struct {
	Arn                 pulumi.StringPtrInput
	AssumeRolePolicy    pulumi.StringPtrInput
	CreateDate          pulumi.StringPtrInput
	Description         pulumi.StringPtrInput
	ForceDetachPolicies pulumi.BoolPtrInput
	MaxSessionDuration  pulumi.IntPtrInput
	Name                pulumi.StringPtrInput
	NamePrefix          pulumi.StringPtrInput
	Path                pulumi.StringPtrInput
	PermissionsBoundary pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
	UniqueId            pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	AssumeRolePolicy    interface{}       `pulumi:"assumeRolePolicy"`
	Description         *string           `pulumi:"description"`
	ForceDetachPolicies *bool             `pulumi:"forceDetachPolicies"`
	MaxSessionDuration  *int              `pulumi:"maxSessionDuration"`
	Name                *string           `pulumi:"name"`
	NamePrefix          *string           `pulumi:"namePrefix"`
	Path                *string           `pulumi:"path"`
	PermissionsBoundary *string           `pulumi:"permissionsBoundary"`
	Tags                map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	AssumeRolePolicy    pulumi.Input
	Description         pulumi.StringPtrInput
	ForceDetachPolicies pulumi.BoolPtrInput
	MaxSessionDuration  pulumi.IntPtrInput
	Name                pulumi.StringPtrInput
	NamePrefix          pulumi.StringPtrInput
	Path                pulumi.StringPtrInput
	PermissionsBoundary pulumi.StringPtrInput
	Tags                pulumi.StringMapInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}
