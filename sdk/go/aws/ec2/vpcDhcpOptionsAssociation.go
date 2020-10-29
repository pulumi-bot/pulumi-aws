// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a VPC DHCP Options Association resource.
//
// ## Remarks
//
// * You can only associate one DHCP Options Set to a given VPC ID.
// * Removing the DHCP Options Association automatically sets AWS's `default` DHCP Options Set to the VPC.
type VpcDhcpOptionsAssociation struct {
	pulumi.CustomResourceState

	// The ID of the DHCP Options Set to associate to the VPC.
	DhcpOptionsId pulumi.StringOutput `pulumi:"dhcpOptionsId"`
	// The ID of the VPC to which we would like to associate a DHCP Options Set.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpcDhcpOptionsAssociation registers a new resource with the given unique name, arguments, and options.
func NewVpcDhcpOptionsAssociation(ctx *pulumi.Context,
	name string, args *VpcDhcpOptionsAssociationArgs, opts ...pulumi.ResourceOption) (*VpcDhcpOptionsAssociation, error) {
	if args == nil || args.DhcpOptionsId == nil {
		return nil, errors.New("missing required argument 'DhcpOptionsId'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	if args == nil {
		args = &VpcDhcpOptionsAssociationArgs{}
	}
	var resource VpcDhcpOptionsAssociation
	err := ctx.RegisterResource("aws:ec2/vpcDhcpOptionsAssociation:VpcDhcpOptionsAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcDhcpOptionsAssociation gets an existing VpcDhcpOptionsAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcDhcpOptionsAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcDhcpOptionsAssociationState, opts ...pulumi.ResourceOption) (*VpcDhcpOptionsAssociation, error) {
	var resource VpcDhcpOptionsAssociation
	err := ctx.ReadResource("aws:ec2/vpcDhcpOptionsAssociation:VpcDhcpOptionsAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcDhcpOptionsAssociation resources.
type vpcDhcpOptionsAssociationState struct {
	// The ID of the DHCP Options Set to associate to the VPC.
	DhcpOptionsId *string `pulumi:"dhcpOptionsId"`
	// The ID of the VPC to which we would like to associate a DHCP Options Set.
	VpcId *string `pulumi:"vpcId"`
}

type VpcDhcpOptionsAssociationState struct {
	// The ID of the DHCP Options Set to associate to the VPC.
	DhcpOptionsId pulumi.StringPtrInput
	// The ID of the VPC to which we would like to associate a DHCP Options Set.
	VpcId pulumi.StringPtrInput
}

func (VpcDhcpOptionsAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcDhcpOptionsAssociationState)(nil)).Elem()
}

type vpcDhcpOptionsAssociationArgs struct {
	// The ID of the DHCP Options Set to associate to the VPC.
	DhcpOptionsId string `pulumi:"dhcpOptionsId"`
	// The ID of the VPC to which we would like to associate a DHCP Options Set.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcDhcpOptionsAssociation resource.
type VpcDhcpOptionsAssociationArgs struct {
	// The ID of the DHCP Options Set to associate to the VPC.
	DhcpOptionsId pulumi.StringInput
	// The ID of the VPC to which we would like to associate a DHCP Options Set.
	VpcId pulumi.StringInput
}

func (VpcDhcpOptionsAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcDhcpOptionsAssociationArgs)(nil)).Elem()
}
