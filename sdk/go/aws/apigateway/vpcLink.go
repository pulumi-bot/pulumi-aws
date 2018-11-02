// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an API Gateway VPC Link.
type VpcLink struct {
	s *pulumi.ResourceState
}

// NewVpcLink registers a new resource with the given unique name, arguments, and options.
func NewVpcLink(ctx *pulumi.Context,
	name string, args *VpcLinkArgs, opts ...pulumi.ResourceOpt) (*VpcLink, error) {
	if args == nil || args.TargetArn == nil {
		return nil, errors.New("missing required argument 'TargetArn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["targetArn"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["targetArn"] = args.TargetArn
	}
	s, err := ctx.RegisterResource("aws:apigateway/vpcLink:VpcLink", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VpcLink{s: s}, nil
}

// GetVpcLink gets an existing VpcLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcLink(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VpcLinkState, opts ...pulumi.ResourceOpt) (*VpcLink, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["targetArn"] = state.TargetArn
	}
	s, err := ctx.ReadResource("aws:apigateway/vpcLink:VpcLink", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VpcLink{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VpcLink) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VpcLink) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The description of the VPC link.
func (r *VpcLink) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The name used to label and identify the VPC link.
func (r *VpcLink) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
func (r *VpcLink) TargetArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["targetArn"])
}

// Input properties used for looking up and filtering VpcLink resources.
type VpcLinkState struct {
	// The description of the VPC link.
	Description interface{}
	// The name used to label and identify the VPC link.
	Name interface{}
	// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn interface{}
}

// The set of arguments for constructing a VpcLink resource.
type VpcLinkArgs struct {
	// The description of the VPC link.
	Description interface{}
	// The name used to label and identify the VPC link.
	Name interface{}
	// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn interface{}
}
