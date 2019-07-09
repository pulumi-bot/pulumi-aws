// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dax

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a DAX Subnet Group resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dax_subnet_group.html.markdown.
type SubnetGroup struct {
	s *pulumi.ResourceState
}

// NewSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewSubnetGroup(ctx *pulumi.Context,
	name string, args *SubnetGroupArgs, opts ...pulumi.ResourceOpt) (*SubnetGroup, error) {
	if args == nil || args.SubnetIds == nil {
		return nil, errors.New("missing required argument 'SubnetIds'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["subnetIds"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["subnetIds"] = args.SubnetIds
	}
	inputs["vpcId"] = nil
	s, err := ctx.RegisterResource("aws:dax/subnetGroup:SubnetGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SubnetGroup{s: s}, nil
}

// GetSubnetGroup gets an existing SubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SubnetGroupState, opts ...pulumi.ResourceOpt) (*SubnetGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["subnetIds"] = state.SubnetIds
		inputs["vpcId"] = state.VpcId
	}
	s, err := ctx.ReadResource("aws:dax/subnetGroup:SubnetGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SubnetGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SubnetGroup) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SubnetGroup) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A description of the subnet group.
func (r *SubnetGroup) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The name of the subnet group.
func (r *SubnetGroup) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// A list of VPC subnet IDs for the subnet group.
func (r *SubnetGroup) SubnetIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["subnetIds"])
}

// VPC ID of the subnet group.
func (r *SubnetGroup) VpcId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcId"])
}

// Input properties used for looking up and filtering SubnetGroup resources.
type SubnetGroupState struct {
	// A description of the subnet group.
	Description interface{}
	// The name of the subnet group.
	Name interface{}
	// A list of VPC subnet IDs for the subnet group.
	SubnetIds interface{}
	// VPC ID of the subnet group.
	VpcId interface{}
}

// The set of arguments for constructing a SubnetGroup resource.
type SubnetGroupArgs struct {
	// A description of the subnet group.
	Description interface{}
	// The name of the subnet group.
	Name interface{}
	// A list of VPC subnet IDs for the subnet group.
	SubnetIds interface{}
}
