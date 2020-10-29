// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an IAM role.
//
// > *NOTE:* If policies are attached to the role via the `iam.PolicyAttachment` resource and you are modifying the role `name` or `path`, the `forceDetachPolicies` argument must be set to `true` and applied before attempting the operation otherwise you will encounter a `DeleteConflict` error. The `iam.RolePolicyAttachment` resource (recommended) does not have this requirement.
type Role struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) specifying the role.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The policy that grants an entity permission to assume the role.
	AssumeRolePolicy pulumi.StringOutput `pulumi:"assumeRolePolicy"`
	// The creation date of the IAM role.
	CreateDate pulumi.StringOutput `pulumi:"createDate"`
	// The description of the role.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies to force detaching any policies the role has before destroying it. Defaults to `false`.
	ForceDetachPolicies pulumi.BoolPtrOutput `pulumi:"forceDetachPolicies"`
	// The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration pulumi.IntPtrOutput `pulumi:"maxSessionDuration"`
	// The name of the role. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The path to the role.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary pulumi.StringPtrOutput `pulumi:"permissionsBoundary"`
	// Key-value map of tags for the IAM role
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The stable and unique string identifying the role.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
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
	// The Amazon Resource Name (ARN) specifying the role.
	Arn *string `pulumi:"arn"`
	// The policy that grants an entity permission to assume the role.
	AssumeRolePolicy *string `pulumi:"assumeRolePolicy"`
	// The creation date of the IAM role.
	CreateDate *string `pulumi:"createDate"`
	// The description of the role.
	Description *string `pulumi:"description"`
	// Specifies to force detaching any policies the role has before destroying it. Defaults to `false`.
	ForceDetachPolicies *bool `pulumi:"forceDetachPolicies"`
	// The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration *int `pulumi:"maxSessionDuration"`
	// The name of the role. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The path to the role.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path *string `pulumi:"path"`
	// The ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary *string `pulumi:"permissionsBoundary"`
	// Key-value map of tags for the IAM role
	Tags map[string]string `pulumi:"tags"`
	// The stable and unique string identifying the role.
	UniqueId *string `pulumi:"uniqueId"`
}

type RoleState struct {
	// The Amazon Resource Name (ARN) specifying the role.
	Arn pulumi.StringPtrInput
	// The policy that grants an entity permission to assume the role.
	AssumeRolePolicy pulumi.StringPtrInput
	// The creation date of the IAM role.
	CreateDate pulumi.StringPtrInput
	// The description of the role.
	Description pulumi.StringPtrInput
	// Specifies to force detaching any policies the role has before destroying it. Defaults to `false`.
	ForceDetachPolicies pulumi.BoolPtrInput
	// The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration pulumi.IntPtrInput
	// The name of the role. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The path to the role.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path pulumi.StringPtrInput
	// The ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary pulumi.StringPtrInput
	// Key-value map of tags for the IAM role
	Tags pulumi.StringMapInput
	// The stable and unique string identifying the role.
	UniqueId pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// The policy that grants an entity permission to assume the role.
	AssumeRolePolicy interface{} `pulumi:"assumeRolePolicy"`
	// The description of the role.
	Description *string `pulumi:"description"`
	// Specifies to force detaching any policies the role has before destroying it. Defaults to `false`.
	ForceDetachPolicies *bool `pulumi:"forceDetachPolicies"`
	// The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration *int `pulumi:"maxSessionDuration"`
	// The name of the role. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The path to the role.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path *string `pulumi:"path"`
	// The ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary *string `pulumi:"permissionsBoundary"`
	// Key-value map of tags for the IAM role
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// The policy that grants an entity permission to assume the role.
	AssumeRolePolicy pulumi.Input
	// The description of the role.
	Description pulumi.StringPtrInput
	// Specifies to force detaching any policies the role has before destroying it. Defaults to `false`.
	ForceDetachPolicies pulumi.BoolPtrInput
	// The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration pulumi.IntPtrInput
	// The name of the role. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The path to the role.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path pulumi.StringPtrInput
	// The ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary pulumi.StringPtrInput
	// Key-value map of tags for the IAM role
	Tags pulumi.StringMapInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}
