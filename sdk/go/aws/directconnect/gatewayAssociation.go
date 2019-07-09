// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Associates a Direct Connect Gateway with a VGW or transit gateway.
// 
// To create a cross-account association, create an [`aws_dx_gateway_association_proposal` resource](https://www.terraform.io/docs/providers/aws/r/dx_gateway_association_proposal.html)
// in the AWS account that owns the VGW or transit gateway and then accept the proposal in the AWS account that owns the Direct Connect Gateway
// by creating an `aws_dx_gateway_association` resource with the `proposal_id` and `associated_gateway_owner_account_id` attributes set.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_gateway_association.html.markdown.
type GatewayAssociation struct {
	s *pulumi.ResourceState
}

// NewGatewayAssociation registers a new resource with the given unique name, arguments, and options.
func NewGatewayAssociation(ctx *pulumi.Context,
	name string, args *GatewayAssociationArgs, opts ...pulumi.ResourceOpt) (*GatewayAssociation, error) {
	if args == nil || args.DxGatewayId == nil {
		return nil, errors.New("missing required argument 'DxGatewayId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allowedPrefixes"] = nil
		inputs["associatedGatewayId"] = nil
		inputs["associatedGatewayOwnerAccountId"] = nil
		inputs["dxGatewayId"] = nil
		inputs["proposalId"] = nil
		inputs["vpnGatewayId"] = nil
	} else {
		inputs["allowedPrefixes"] = args.AllowedPrefixes
		inputs["associatedGatewayId"] = args.AssociatedGatewayId
		inputs["associatedGatewayOwnerAccountId"] = args.AssociatedGatewayOwnerAccountId
		inputs["dxGatewayId"] = args.DxGatewayId
		inputs["proposalId"] = args.ProposalId
		inputs["vpnGatewayId"] = args.VpnGatewayId
	}
	inputs["associatedGatewayType"] = nil
	inputs["dxGatewayAssociationId"] = nil
	inputs["dxGatewayOwnerAccountId"] = nil
	s, err := ctx.RegisterResource("aws:directconnect/gatewayAssociation:GatewayAssociation", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &GatewayAssociation{s: s}, nil
}

// GetGatewayAssociation gets an existing GatewayAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayAssociation(ctx *pulumi.Context,
	name string, id pulumi.ID, state *GatewayAssociationState, opts ...pulumi.ResourceOpt) (*GatewayAssociation, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allowedPrefixes"] = state.AllowedPrefixes
		inputs["associatedGatewayId"] = state.AssociatedGatewayId
		inputs["associatedGatewayOwnerAccountId"] = state.AssociatedGatewayOwnerAccountId
		inputs["associatedGatewayType"] = state.AssociatedGatewayType
		inputs["dxGatewayAssociationId"] = state.DxGatewayAssociationId
		inputs["dxGatewayId"] = state.DxGatewayId
		inputs["dxGatewayOwnerAccountId"] = state.DxGatewayOwnerAccountId
		inputs["proposalId"] = state.ProposalId
		inputs["vpnGatewayId"] = state.VpnGatewayId
	}
	s, err := ctx.ReadResource("aws:directconnect/gatewayAssociation:GatewayAssociation", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &GatewayAssociation{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *GatewayAssociation) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *GatewayAssociation) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
func (r *GatewayAssociation) AllowedPrefixes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["allowedPrefixes"])
}

// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
// Used for single account Direct Connect gateway associations.
func (r *GatewayAssociation) AssociatedGatewayId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["associatedGatewayId"])
}

// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
// Used for cross-account Direct Connect gateway associations.
func (r *GatewayAssociation) AssociatedGatewayOwnerAccountId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["associatedGatewayOwnerAccountId"])
}

// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
func (r *GatewayAssociation) AssociatedGatewayType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["associatedGatewayType"])
}

// The ID of the Direct Connect gateway association.
func (r *GatewayAssociation) DxGatewayAssociationId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dxGatewayAssociationId"])
}

// The ID of the Direct Connect gateway.
func (r *GatewayAssociation) DxGatewayId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dxGatewayId"])
}

// The ID of the AWS account that owns the Direct Connect gateway.
func (r *GatewayAssociation) DxGatewayOwnerAccountId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dxGatewayOwnerAccountId"])
}

// The ID of the Direct Connect gateway association proposal.
// Used for cross-account Direct Connect gateway associations.
func (r *GatewayAssociation) ProposalId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["proposalId"])
}

// *Deprecated:* Use `associated_gateway_id` instead. The ID of the VGW with which to associate the gateway.
// Used for single account Direct Connect gateway associations.
func (r *GatewayAssociation) VpnGatewayId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpnGatewayId"])
}

// Input properties used for looking up and filtering GatewayAssociation resources.
type GatewayAssociationState struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes interface{}
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId interface{}
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId interface{}
	// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
	AssociatedGatewayType interface{}
	// The ID of the Direct Connect gateway association.
	DxGatewayAssociationId interface{}
	// The ID of the Direct Connect gateway.
	DxGatewayId interface{}
	// The ID of the AWS account that owns the Direct Connect gateway.
	DxGatewayOwnerAccountId interface{}
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId interface{}
	// *Deprecated:* Use `associated_gateway_id` instead. The ID of the VGW with which to associate the gateway.
	// Used for single account Direct Connect gateway associations.
	VpnGatewayId interface{}
}

// The set of arguments for constructing a GatewayAssociation resource.
type GatewayAssociationArgs struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes interface{}
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId interface{}
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId interface{}
	// The ID of the Direct Connect gateway.
	DxGatewayId interface{}
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId interface{}
	// *Deprecated:* Use `associated_gateway_id` instead. The ID of the VGW with which to associate the gateway.
	// Used for single account Direct Connect gateway associations.
	VpnGatewayId interface{}
}
