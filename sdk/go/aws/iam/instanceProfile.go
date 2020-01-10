// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an IAM instance profile.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_instance_profile.html.markdown.
type InstanceProfile struct {
	s *pulumi.ResourceState
}

// NewInstanceProfile registers a new resource with the given unique name, arguments, and options.
func NewInstanceProfile(ctx *pulumi.Context,
	name string, args *InstanceProfileArgs, opts ...pulumi.ResourceOpt) (*InstanceProfile, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["namePrefix"] = nil
		inputs["path"] = nil
		inputs["role"] = nil
		inputs["roles"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["path"] = args.Path
		inputs["role"] = args.Role
		inputs["roles"] = args.Roles
	}
	inputs["arn"] = nil
	inputs["createDate"] = nil
	inputs["uniqueId"] = nil
	s, err := ctx.RegisterResource("aws:iam/instanceProfile:InstanceProfile", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InstanceProfile{s: s}, nil
}

// GetInstanceProfile gets an existing InstanceProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceProfile(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InstanceProfileState, opts ...pulumi.ResourceOpt) (*InstanceProfile, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["createDate"] = state.CreateDate
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["path"] = state.Path
		inputs["role"] = state.Role
		inputs["roles"] = state.Roles
		inputs["uniqueId"] = state.UniqueId
	}
	s, err := ctx.ReadResource("aws:iam/instanceProfile:InstanceProfile", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InstanceProfile{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *InstanceProfile) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *InstanceProfile) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ARN assigned by AWS to the instance profile.
func (r *InstanceProfile) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The creation timestamp of the instance profile.
func (r *InstanceProfile) CreateDate() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["createDate"])
}

// The profile's name. If omitted, this provider will assign a random, unique name.
func (r *InstanceProfile) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (r *InstanceProfile) NamePrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["namePrefix"])
}

// Path in which to create the profile.
func (r *InstanceProfile) Path() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["path"])
}

// The role name to include in the profile.
func (r *InstanceProfile) Role() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["role"])
}

// 
// A list of role names to include in the profile.  The current default is 1.  If you see an error message similar to `Cannot exceed quota for InstanceSessionsPerInstanceProfile: 1`, then you must contact AWS support and ask for a limit increase.
func (r *InstanceProfile) Roles() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["roles"])
}

// The [unique ID][1] assigned by AWS.
func (r *InstanceProfile) UniqueId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["uniqueId"])
}

// Input properties used for looking up and filtering InstanceProfile resources.
type InstanceProfileState struct {
	// The ARN assigned by AWS to the instance profile.
	Arn interface{}
	// The creation timestamp of the instance profile.
	CreateDate interface{}
	// The profile's name. If omitted, this provider will assign a random, unique name.
	Name interface{}
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix interface{}
	// Path in which to create the profile.
	Path interface{}
	// The role name to include in the profile.
	Role interface{}
	// 
	// A list of role names to include in the profile.  The current default is 1.  If you see an error message similar to `Cannot exceed quota for InstanceSessionsPerInstanceProfile: 1`, then you must contact AWS support and ask for a limit increase.
	// Deprecated: Use `role` instead. Only a single role can be passed to an IAM Instance Profile
	Roles interface{}
	// The [unique ID][1] assigned by AWS.
	UniqueId interface{}
}

// The set of arguments for constructing a InstanceProfile resource.
type InstanceProfileArgs struct {
	// The profile's name. If omitted, this provider will assign a random, unique name.
	Name interface{}
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix interface{}
	// Path in which to create the profile.
	Path interface{}
	// The role name to include in the profile.
	Role interface{}
	// 
	// A list of role names to include in the profile.  The current default is 1.  If you see an error message similar to `Cannot exceed quota for InstanceSessionsPerInstanceProfile: 1`, then you must contact AWS support and ask for a limit increase.
	// Deprecated: Use `role` instead. Only a single role can be passed to an IAM Instance Profile
	Roles interface{}
}
