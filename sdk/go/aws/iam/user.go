// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an IAM user.
type User struct {
	s *pulumi.ResourceState
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOpt) (*User, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["forceDestroy"] = nil
		inputs["name"] = nil
		inputs["path"] = nil
		inputs["permissionsBoundary"] = nil
	} else {
		inputs["forceDestroy"] = args.ForceDestroy
		inputs["name"] = args.Name
		inputs["path"] = args.Path
		inputs["permissionsBoundary"] = args.PermissionsBoundary
	}
	inputs["arn"] = nil
	inputs["uniqueId"] = nil
	s, err := ctx.RegisterResource("aws:iam/user:User", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &User{s: s}, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.ID, state *UserState, opts ...pulumi.ResourceOpt) (*User, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["forceDestroy"] = state.ForceDestroy
		inputs["name"] = state.Name
		inputs["path"] = state.Path
		inputs["permissionsBoundary"] = state.PermissionsBoundary
		inputs["uniqueId"] = state.UniqueId
	}
	s, err := ctx.ReadResource("aws:iam/user:User", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &User{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *User) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *User) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN assigned by AWS for this user.
func (r *User) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// When destroying this user, destroy even if it
// has non-Terraform-managed IAM access keys, login profile or MFA devices. Without `force_destroy`
// a user with non-Terraform-managed access keys and login profile will fail to be destroyed.
func (r *User) ForceDestroy() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["forceDestroy"])
}

// The user's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. User names are not distinguished by case. For example, you cannot create users named both "TESTUSER" and "testuser".
func (r *User) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Path in which to create the user.
func (r *User) Path() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["path"])
}

// The ARN of the policy that is used to set the permissions boundary for the user.
func (r *User) PermissionsBoundary() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["permissionsBoundary"])
}

// The [unique ID][1] assigned by AWS.
func (r *User) UniqueId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["uniqueId"])
}

// Input properties used for looking up and filtering User resources.
type UserState struct {
	// The ARN assigned by AWS for this user.
	Arn interface{}
	// When destroying this user, destroy even if it
	// has non-Terraform-managed IAM access keys, login profile or MFA devices. Without `force_destroy`
	// a user with non-Terraform-managed access keys and login profile will fail to be destroyed.
	ForceDestroy interface{}
	// The user's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. User names are not distinguished by case. For example, you cannot create users named both "TESTUSER" and "testuser".
	Name interface{}
	// Path in which to create the user.
	Path interface{}
	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary interface{}
	// The [unique ID][1] assigned by AWS.
	UniqueId interface{}
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// When destroying this user, destroy even if it
	// has non-Terraform-managed IAM access keys, login profile or MFA devices. Without `force_destroy`
	// a user with non-Terraform-managed access keys and login profile will fail to be destroyed.
	ForceDestroy interface{}
	// The user's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. User names are not distinguished by case. For example, you cannot create users named both "TESTUSER" and "testuser".
	Name interface{}
	// Path in which to create the user.
	Path interface{}
	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary interface{}
}
