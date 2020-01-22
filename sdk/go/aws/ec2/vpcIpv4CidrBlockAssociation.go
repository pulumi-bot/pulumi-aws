// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to associate additional IPv4 CIDR blocks with a VPC.
// 
// When a VPC is created, a primary IPv4 CIDR block for the VPC must be specified.
// The `ec2.VpcIpv4CidrBlockAssociation` resource allows further IPv4 CIDR blocks to be added to the VPC.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_ipv4_cidr_block_association.html.markdown.
type VpcIpv4CidrBlockAssociation struct {
	pulumi.CustomResourceState

	// The additional IPv4 CIDR block to associate with the VPC.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The ID of the VPC to make the association with.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpcIpv4CidrBlockAssociation registers a new resource with the given unique name, arguments, and options.
func NewVpcIpv4CidrBlockAssociation(ctx *pulumi.Context,
	name string, args *VpcIpv4CidrBlockAssociationArgs, opts ...pulumi.ResourceOption) (*VpcIpv4CidrBlockAssociation, error) {
	if args == nil || args.CidrBlock == nil {
		return nil, errors.New("missing required argument 'CidrBlock'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	if args == nil {
		args = &VpcIpv4CidrBlockAssociationArgs{}
	}
	var resource VpcIpv4CidrBlockAssociation
	err := ctx.RegisterResource("aws:ec2/vpcIpv4CidrBlockAssociation:VpcIpv4CidrBlockAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcIpv4CidrBlockAssociation gets an existing VpcIpv4CidrBlockAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcIpv4CidrBlockAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcIpv4CidrBlockAssociationState, opts ...pulumi.ResourceOption) (*VpcIpv4CidrBlockAssociation, error) {
	var resource VpcIpv4CidrBlockAssociation
	err := ctx.ReadResource("aws:ec2/vpcIpv4CidrBlockAssociation:VpcIpv4CidrBlockAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcIpv4CidrBlockAssociation resources.
type vpcIpv4CidrBlockAssociationState struct {
	// The additional IPv4 CIDR block to associate with the VPC.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The ID of the VPC to make the association with.
	VpcId *string `pulumi:"vpcId"`
}

type VpcIpv4CidrBlockAssociationState struct {
	// The additional IPv4 CIDR block to associate with the VPC.
	CidrBlock pulumi.StringPtrInput
	// The ID of the VPC to make the association with.
	VpcId pulumi.StringPtrInput
}

func (VpcIpv4CidrBlockAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpv4CidrBlockAssociationState)(nil)).Elem()
}

type vpcIpv4CidrBlockAssociationArgs struct {
	// The additional IPv4 CIDR block to associate with the VPC.
	CidrBlock string `pulumi:"cidrBlock"`
	// The ID of the VPC to make the association with.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcIpv4CidrBlockAssociation resource.
type VpcIpv4CidrBlockAssociationArgs struct {
	// The additional IPv4 CIDR block to associate with the VPC.
	CidrBlock pulumi.StringInput
	// The ID of the VPC to make the association with.
	VpcId pulumi.StringInput
}

func (VpcIpv4CidrBlockAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpv4CidrBlockAssociationArgs)(nil)).Elem()
}

