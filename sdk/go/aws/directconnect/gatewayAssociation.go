// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Associates a Direct Connect Gateway with a VGW or transit gateway.
//
// To create a cross-account association, create an [`directconnect.GatewayAssociationProposal` resource](https://www.terraform.io/docs/providers/aws/r/dx_gateway_association_proposal.html)
// in the AWS account that owns the VGW or transit gateway and then accept the proposal in the AWS account that owns the Direct Connect Gateway
// by creating an `directconnect.GatewayAssociation` resource with the `proposalId` and `associatedGatewayOwnerAccountId` attributes set.
type GatewayAssociation struct {
	pulumi.CustomResourceState

	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes pulumi.StringArrayOutput `pulumi:"allowedPrefixes"`
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId pulumi.StringOutput `pulumi:"associatedGatewayId"`
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId pulumi.StringOutput `pulumi:"associatedGatewayOwnerAccountId"`
	// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
	AssociatedGatewayType pulumi.StringOutput `pulumi:"associatedGatewayType"`
	// The ID of the Direct Connect gateway association.
	DxGatewayAssociationId pulumi.StringOutput `pulumi:"dxGatewayAssociationId"`
	// The ID of the Direct Connect gateway.
	DxGatewayId pulumi.StringOutput `pulumi:"dxGatewayId"`
	// The ID of the AWS account that owns the Direct Connect gateway.
	DxGatewayOwnerAccountId pulumi.StringOutput `pulumi:"dxGatewayOwnerAccountId"`
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId pulumi.StringPtrOutput `pulumi:"proposalId"`
	// *Deprecated:* Use `associatedGatewayId` instead. The ID of the VGW with which to associate the gateway.
	// Used for single account Direct Connect gateway associations.
	//
	// use 'associated_gateway_id' argument instead
	VpnGatewayId pulumi.StringPtrOutput `pulumi:"vpnGatewayId"`
}

// NewGatewayAssociation registers a new resource with the given unique name, arguments, and options.
func NewGatewayAssociation(ctx *pulumi.Context,
	name string, args *GatewayAssociationArgs, opts ...pulumi.ResourceOption) (*GatewayAssociation, error) {
	if args == nil || args.DxGatewayId == nil {
		return nil, errors.New("missing required argument 'DxGatewayId'")
	}
	if args == nil {
		args = &GatewayAssociationArgs{}
	}
	var resource GatewayAssociation
	err := ctx.RegisterResource("aws:directconnect/gatewayAssociation:GatewayAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayAssociation gets an existing GatewayAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayAssociationState, opts ...pulumi.ResourceOption) (*GatewayAssociation, error) {
	var resource GatewayAssociation
	err := ctx.ReadResource("aws:directconnect/gatewayAssociation:GatewayAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayAssociation resources.
type gatewayAssociationState struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes []string `pulumi:"allowedPrefixes"`
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId *string `pulumi:"associatedGatewayId"`
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId *string `pulumi:"associatedGatewayOwnerAccountId"`
	// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
	AssociatedGatewayType *string `pulumi:"associatedGatewayType"`
	// The ID of the Direct Connect gateway association.
	DxGatewayAssociationId *string `pulumi:"dxGatewayAssociationId"`
	// The ID of the Direct Connect gateway.
	DxGatewayId *string `pulumi:"dxGatewayId"`
	// The ID of the AWS account that owns the Direct Connect gateway.
	DxGatewayOwnerAccountId *string `pulumi:"dxGatewayOwnerAccountId"`
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId *string `pulumi:"proposalId"`
	// *Deprecated:* Use `associatedGatewayId` instead. The ID of the VGW with which to associate the gateway.
	// Used for single account Direct Connect gateway associations.
	//
	// use 'associated_gateway_id' argument instead
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

type GatewayAssociationState struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes pulumi.StringArrayInput
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId pulumi.StringPtrInput
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId pulumi.StringPtrInput
	// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
	AssociatedGatewayType pulumi.StringPtrInput
	// The ID of the Direct Connect gateway association.
	DxGatewayAssociationId pulumi.StringPtrInput
	// The ID of the Direct Connect gateway.
	DxGatewayId pulumi.StringPtrInput
	// The ID of the AWS account that owns the Direct Connect gateway.
	DxGatewayOwnerAccountId pulumi.StringPtrInput
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId pulumi.StringPtrInput
	// *Deprecated:* Use `associatedGatewayId` instead. The ID of the VGW with which to associate the gateway.
	// Used for single account Direct Connect gateway associations.
	//
	// use 'associated_gateway_id' argument instead
	VpnGatewayId pulumi.StringPtrInput
}

func (GatewayAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayAssociationState)(nil)).Elem()
}

type gatewayAssociationArgs struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes []string `pulumi:"allowedPrefixes"`
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId *string `pulumi:"associatedGatewayId"`
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId *string `pulumi:"associatedGatewayOwnerAccountId"`
	// The ID of the Direct Connect gateway.
	DxGatewayId string `pulumi:"dxGatewayId"`
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId *string `pulumi:"proposalId"`
	// *Deprecated:* Use `associatedGatewayId` instead. The ID of the VGW with which to associate the gateway.
	// Used for single account Direct Connect gateway associations.
	//
	// use 'associated_gateway_id' argument instead
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a GatewayAssociation resource.
type GatewayAssociationArgs struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes pulumi.StringArrayInput
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId pulumi.StringPtrInput
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId pulumi.StringPtrInput
	// The ID of the Direct Connect gateway.
	DxGatewayId pulumi.StringInput
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId pulumi.StringPtrInput
	// *Deprecated:* Use `associatedGatewayId` instead. The ID of the VGW with which to associate the gateway.
	// Used for single account Direct Connect gateway associations.
	//
	// use 'associated_gateway_id' argument instead
	VpnGatewayId pulumi.StringPtrInput
}

func (GatewayAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayAssociationArgs)(nil)).Elem()
}
