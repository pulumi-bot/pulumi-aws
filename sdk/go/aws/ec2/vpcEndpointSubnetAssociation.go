// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create an association between a VPC endpoint and a subnet.
// 
// ~> **NOTE on VPC Endpoints and VPC Endpoint Subnet Associations:** Terraform provides
// both a standalone VPC Endpoint Subnet Association (an association between a VPC endpoint
// and a single `subnet_id`) and a VPC Endpoint resource with a `subnet_ids`
// attribute. Do not use the same subnet ID in both a VPC Endpoint resource and a VPC Endpoint Subnet
// Association resource. Doing so will cause a conflict of associations and will overwrite the association.
type VpcEndpointSubnetAssociation struct {
	s *pulumi.ResourceState
}

// NewVpcEndpointSubnetAssociation registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointSubnetAssociation(ctx *pulumi.Context,
	name string, args *VpcEndpointSubnetAssociationArgs, opts ...pulumi.ResourceOpt) (*VpcEndpointSubnetAssociation, error) {
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	if args == nil || args.VpcEndpointId == nil {
		return nil, errors.New("missing required argument 'VpcEndpointId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["subnetId"] = nil
		inputs["vpcEndpointId"] = nil
	} else {
		inputs["subnetId"] = args.SubnetId
		inputs["vpcEndpointId"] = args.VpcEndpointId
	}
	s, err := ctx.RegisterResource("aws:ec2/vpcEndpointSubnetAssociation:VpcEndpointSubnetAssociation", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VpcEndpointSubnetAssociation{s: s}, nil
}

// GetVpcEndpointSubnetAssociation gets an existing VpcEndpointSubnetAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointSubnetAssociation(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VpcEndpointSubnetAssociationState, opts ...pulumi.ResourceOpt) (*VpcEndpointSubnetAssociation, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["subnetId"] = state.SubnetId
		inputs["vpcEndpointId"] = state.VpcEndpointId
	}
	s, err := ctx.ReadResource("aws:ec2/vpcEndpointSubnetAssociation:VpcEndpointSubnetAssociation", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VpcEndpointSubnetAssociation{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VpcEndpointSubnetAssociation) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VpcEndpointSubnetAssociation) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ID of the subnet to be associated with the VPC endpoint.
func (r *VpcEndpointSubnetAssociation) SubnetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetId"])
}

// The ID of the VPC endpoint with which the subnet will be associated.
func (r *VpcEndpointSubnetAssociation) VpcEndpointId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcEndpointId"])
}

// Input properties used for looking up and filtering VpcEndpointSubnetAssociation resources.
type VpcEndpointSubnetAssociationState struct {
	// The ID of the subnet to be associated with the VPC endpoint.
	SubnetId interface{}
	// The ID of the VPC endpoint with which the subnet will be associated.
	VpcEndpointId interface{}
}

// The set of arguments for constructing a VpcEndpointSubnetAssociation resource.
type VpcEndpointSubnetAssociationArgs struct {
	// The ID of the subnet to be associated with the VPC endpoint.
	SubnetId interface{}
	// The ID of the VPC endpoint with which the subnet will be associated.
	VpcEndpointId interface{}
}
