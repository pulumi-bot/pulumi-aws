// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Vpc struct {
	pulumi.CustomResourceState

	Arn                          pulumi.StringOutput    `pulumi:"arn"`
	AssignGeneratedIpv6CidrBlock pulumi.BoolPtrOutput   `pulumi:"assignGeneratedIpv6CidrBlock"`
	CidrBlock                    pulumi.StringOutput    `pulumi:"cidrBlock"`
	DefaultNetworkAclId          pulumi.StringOutput    `pulumi:"defaultNetworkAclId"`
	DefaultRouteTableId          pulumi.StringOutput    `pulumi:"defaultRouteTableId"`
	DefaultSecurityGroupId       pulumi.StringOutput    `pulumi:"defaultSecurityGroupId"`
	DhcpOptionsId                pulumi.StringOutput    `pulumi:"dhcpOptionsId"`
	EnableClassiclink            pulumi.BoolOutput      `pulumi:"enableClassiclink"`
	EnableClassiclinkDnsSupport  pulumi.BoolOutput      `pulumi:"enableClassiclinkDnsSupport"`
	EnableDnsHostnames           pulumi.BoolOutput      `pulumi:"enableDnsHostnames"`
	EnableDnsSupport             pulumi.BoolPtrOutput   `pulumi:"enableDnsSupport"`
	InstanceTenancy              pulumi.StringPtrOutput `pulumi:"instanceTenancy"`
	Ipv6AssociationId            pulumi.StringOutput    `pulumi:"ipv6AssociationId"`
	Ipv6CidrBlock                pulumi.StringOutput    `pulumi:"ipv6CidrBlock"`
	MainRouteTableId             pulumi.StringOutput    `pulumi:"mainRouteTableId"`
	OwnerId                      pulumi.StringOutput    `pulumi:"ownerId"`
	Tags                         pulumi.StringMapOutput `pulumi:"tags"`
}

// NewVpc registers a new resource with the given unique name, arguments, and options.
func NewVpc(ctx *pulumi.Context,
	name string, args *VpcArgs, opts ...pulumi.ResourceOption) (*Vpc, error) {
	if args == nil || args.CidrBlock == nil {
		return nil, errors.New("missing required argument 'CidrBlock'")
	}
	if args == nil {
		args = &VpcArgs{}
	}
	var resource Vpc
	err := ctx.RegisterResource("aws:ec2/vpc:Vpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpc gets an existing Vpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcState, opts ...pulumi.ResourceOption) (*Vpc, error) {
	var resource Vpc
	err := ctx.ReadResource("aws:ec2/vpc:Vpc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vpc resources.
type vpcState struct {
	Arn                          *string           `pulumi:"arn"`
	AssignGeneratedIpv6CidrBlock *bool             `pulumi:"assignGeneratedIpv6CidrBlock"`
	CidrBlock                    *string           `pulumi:"cidrBlock"`
	DefaultNetworkAclId          *string           `pulumi:"defaultNetworkAclId"`
	DefaultRouteTableId          *string           `pulumi:"defaultRouteTableId"`
	DefaultSecurityGroupId       *string           `pulumi:"defaultSecurityGroupId"`
	DhcpOptionsId                *string           `pulumi:"dhcpOptionsId"`
	EnableClassiclink            *bool             `pulumi:"enableClassiclink"`
	EnableClassiclinkDnsSupport  *bool             `pulumi:"enableClassiclinkDnsSupport"`
	EnableDnsHostnames           *bool             `pulumi:"enableDnsHostnames"`
	EnableDnsSupport             *bool             `pulumi:"enableDnsSupport"`
	InstanceTenancy              *string           `pulumi:"instanceTenancy"`
	Ipv6AssociationId            *string           `pulumi:"ipv6AssociationId"`
	Ipv6CidrBlock                *string           `pulumi:"ipv6CidrBlock"`
	MainRouteTableId             *string           `pulumi:"mainRouteTableId"`
	OwnerId                      *string           `pulumi:"ownerId"`
	Tags                         map[string]string `pulumi:"tags"`
}

type VpcState struct {
	Arn                          pulumi.StringPtrInput
	AssignGeneratedIpv6CidrBlock pulumi.BoolPtrInput
	CidrBlock                    pulumi.StringPtrInput
	DefaultNetworkAclId          pulumi.StringPtrInput
	DefaultRouteTableId          pulumi.StringPtrInput
	DefaultSecurityGroupId       pulumi.StringPtrInput
	DhcpOptionsId                pulumi.StringPtrInput
	EnableClassiclink            pulumi.BoolPtrInput
	EnableClassiclinkDnsSupport  pulumi.BoolPtrInput
	EnableDnsHostnames           pulumi.BoolPtrInput
	EnableDnsSupport             pulumi.BoolPtrInput
	InstanceTenancy              pulumi.StringPtrInput
	Ipv6AssociationId            pulumi.StringPtrInput
	Ipv6CidrBlock                pulumi.StringPtrInput
	MainRouteTableId             pulumi.StringPtrInput
	OwnerId                      pulumi.StringPtrInput
	Tags                         pulumi.StringMapInput
}

func (VpcState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcState)(nil)).Elem()
}

type vpcArgs struct {
	AssignGeneratedIpv6CidrBlock *bool             `pulumi:"assignGeneratedIpv6CidrBlock"`
	CidrBlock                    string            `pulumi:"cidrBlock"`
	EnableClassiclink            *bool             `pulumi:"enableClassiclink"`
	EnableClassiclinkDnsSupport  *bool             `pulumi:"enableClassiclinkDnsSupport"`
	EnableDnsHostnames           *bool             `pulumi:"enableDnsHostnames"`
	EnableDnsSupport             *bool             `pulumi:"enableDnsSupport"`
	InstanceTenancy              *string           `pulumi:"instanceTenancy"`
	Tags                         map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Vpc resource.
type VpcArgs struct {
	AssignGeneratedIpv6CidrBlock pulumi.BoolPtrInput
	CidrBlock                    pulumi.StringInput
	EnableClassiclink            pulumi.BoolPtrInput
	EnableClassiclinkDnsSupport  pulumi.BoolPtrInput
	EnableDnsHostnames           pulumi.BoolPtrInput
	EnableDnsSupport             pulumi.BoolPtrInput
	InstanceTenancy              pulumi.StringPtrInput
	Tags                         pulumi.StringMapInput
}

func (VpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcArgs)(nil)).Elem()
}
