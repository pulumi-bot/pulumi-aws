// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Neptune Parameter Group
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/neptune_parameter_group.html.markdown.
type ParameterGroup struct {
	s *pulumi.ResourceState
}

// NewParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewParameterGroup(ctx *pulumi.Context,
	name string, args *ParameterGroupArgs, opts ...pulumi.ResourceOpt) (*ParameterGroup, error) {
	if args == nil || args.Family == nil {
		return nil, errors.New("missing required argument 'Family'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["family"] = nil
		inputs["name"] = nil
		inputs["parameters"] = nil
		inputs["tags"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["family"] = args.Family
		inputs["name"] = args.Name
		inputs["parameters"] = args.Parameters
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:neptune/parameterGroup:ParameterGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ParameterGroup{s: s}, nil
}

// GetParameterGroup gets an existing ParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ParameterGroupState, opts ...pulumi.ResourceOpt) (*ParameterGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["description"] = state.Description
		inputs["family"] = state.Family
		inputs["name"] = state.Name
		inputs["parameters"] = state.Parameters
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("aws:neptune/parameterGroup:ParameterGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ParameterGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ParameterGroup) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ParameterGroup) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The Neptune parameter group Amazon Resource Name (ARN).
func (r *ParameterGroup) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

func (r *ParameterGroup) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The family of the Neptune parameter group.
func (r *ParameterGroup) Family() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["family"])
}

// The name of the Neptune parameter.
func (r *ParameterGroup) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// A list of Neptune parameters to apply.
func (r *ParameterGroup) Parameters() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["parameters"])
}

// A mapping of tags to assign to the resource.
func (r *ParameterGroup) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering ParameterGroup resources.
type ParameterGroupState struct {
	// The Neptune parameter group Amazon Resource Name (ARN).
	Arn interface{}
	Description interface{}
	// The family of the Neptune parameter group.
	Family interface{}
	// The name of the Neptune parameter.
	Name interface{}
	// A list of Neptune parameters to apply.
	Parameters interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a ParameterGroup resource.
type ParameterGroupArgs struct {
	Description interface{}
	// The family of the Neptune parameter group.
	Family interface{}
	// The name of the Neptune parameter.
	Name interface{}
	// A list of Neptune parameters to apply.
	Parameters interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
