// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sfn

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Step Function Activity resource
type Activity struct {
	s *pulumi.ResourceState
}

// NewActivity registers a new resource with the given unique name, arguments, and options.
func NewActivity(ctx *pulumi.Context,
	name string, args *ActivityArgs, opts ...pulumi.ResourceOpt) (*Activity, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
	} else {
		inputs["name"] = args.Name
	}
	inputs["creationDate"] = nil
	s, err := ctx.RegisterResource("aws:sfn/activity:Activity", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Activity{s: s}, nil
}

// GetActivity gets an existing Activity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActivity(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ActivityState, opts ...pulumi.ResourceOpt) (*Activity, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationDate"] = state.CreationDate
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("aws:sfn/activity:Activity", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Activity{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Activity) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Activity) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The date the activity was created.
func (r *Activity) CreationDate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationDate"])
}

// The name of the activity to create.
func (r *Activity) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering Activity resources.
type ActivityState struct {
	// The date the activity was created.
	CreationDate interface{}
	// The name of the activity to create.
	Name interface{}
}

// The set of arguments for constructing a Activity resource.
type ActivityArgs struct {
	// The name of the activity to create.
	Name interface{}
}
