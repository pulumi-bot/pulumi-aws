// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a VPC resource.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Basic usage with tags:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
// 			CidrBlock:       pulumi.String("10.0.0.0/16"),
// 			InstanceTenancy: pulumi.String("default"),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("main"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Vpc struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of VPC
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Requests an Amazon-provided IPv6 CIDR
	// block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
	// the size of the CIDR block. Default is `false`.
	AssignGeneratedIpv6CidrBlock pulumi.BoolPtrOutput `pulumi:"assignGeneratedIpv6CidrBlock"`
	// The CIDR block for the VPC.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkAclId pulumi.StringOutput `pulumi:"defaultNetworkAclId"`
	// The ID of the route table created by default on VPC creation
	DefaultRouteTableId pulumi.StringOutput `pulumi:"defaultRouteTableId"`
	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupId pulumi.StringOutput `pulumi:"defaultSecurityGroupId"`
	DhcpOptionsId          pulumi.StringOutput `pulumi:"dhcpOptionsId"`
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	EnableClassiclink pulumi.BoolOutput `pulumi:"enableClassiclink"`
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	EnableClassiclinkDnsSupport pulumi.BoolOutput `pulumi:"enableClassiclinkDnsSupport"`
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames pulumi.BoolOutput `pulumi:"enableDnsHostnames"`
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport pulumi.BoolPtrOutput `pulumi:"enableDnsSupport"`
	// A tenancy option for instances launched into the VPC. Default is `default`, which
	// makes your instances shared on the host. Using either of the other options (`dedicated` or `host`) costs at least $2/hr.
	InstanceTenancy pulumi.StringPtrOutput `pulumi:"instanceTenancy"`
	// The association ID for the IPv6 CIDR block.
	Ipv6AssociationId pulumi.StringOutput `pulumi:"ipv6AssociationId"`
	// The IPv6 CIDR block.
	Ipv6CidrBlock pulumi.StringOutput `pulumi:"ipv6CidrBlock"`
	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// `ec2.MainRouteTableAssociation`.
	MainRouteTableId pulumi.StringOutput `pulumi:"mainRouteTableId"`
	// The ID of the AWS account that owns the VPC.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewVpc registers a new resource with the given unique name, arguments, and options.
func NewVpc(ctx *pulumi.Context,
	name string, args *VpcArgs, opts ...pulumi.ResourceOption) (*Vpc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
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
	// Amazon Resource Name (ARN) of VPC
	Arn *string `pulumi:"arn"`
	// Requests an Amazon-provided IPv6 CIDR
	// block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
	// the size of the CIDR block. Default is `false`.
	AssignGeneratedIpv6CidrBlock *bool `pulumi:"assignGeneratedIpv6CidrBlock"`
	// The CIDR block for the VPC.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkAclId *string `pulumi:"defaultNetworkAclId"`
	// The ID of the route table created by default on VPC creation
	DefaultRouteTableId *string `pulumi:"defaultRouteTableId"`
	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupId *string `pulumi:"defaultSecurityGroupId"`
	DhcpOptionsId          *string `pulumi:"dhcpOptionsId"`
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	EnableClassiclink *bool `pulumi:"enableClassiclink"`
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	EnableClassiclinkDnsSupport *bool `pulumi:"enableClassiclinkDnsSupport"`
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames *bool `pulumi:"enableDnsHostnames"`
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport *bool `pulumi:"enableDnsSupport"`
	// A tenancy option for instances launched into the VPC. Default is `default`, which
	// makes your instances shared on the host. Using either of the other options (`dedicated` or `host`) costs at least $2/hr.
	InstanceTenancy *string `pulumi:"instanceTenancy"`
	// The association ID for the IPv6 CIDR block.
	Ipv6AssociationId *string `pulumi:"ipv6AssociationId"`
	// The IPv6 CIDR block.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// `ec2.MainRouteTableAssociation`.
	MainRouteTableId *string `pulumi:"mainRouteTableId"`
	// The ID of the AWS account that owns the VPC.
	OwnerId *string `pulumi:"ownerId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type VpcState struct {
	// Amazon Resource Name (ARN) of VPC
	Arn pulumi.StringPtrInput
	// Requests an Amazon-provided IPv6 CIDR
	// block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
	// the size of the CIDR block. Default is `false`.
	AssignGeneratedIpv6CidrBlock pulumi.BoolPtrInput
	// The CIDR block for the VPC.
	CidrBlock pulumi.StringPtrInput
	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkAclId pulumi.StringPtrInput
	// The ID of the route table created by default on VPC creation
	DefaultRouteTableId pulumi.StringPtrInput
	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupId pulumi.StringPtrInput
	DhcpOptionsId          pulumi.StringPtrInput
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	EnableClassiclink pulumi.BoolPtrInput
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	EnableClassiclinkDnsSupport pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport pulumi.BoolPtrInput
	// A tenancy option for instances launched into the VPC. Default is `default`, which
	// makes your instances shared on the host. Using either of the other options (`dedicated` or `host`) costs at least $2/hr.
	InstanceTenancy pulumi.StringPtrInput
	// The association ID for the IPv6 CIDR block.
	Ipv6AssociationId pulumi.StringPtrInput
	// The IPv6 CIDR block.
	Ipv6CidrBlock pulumi.StringPtrInput
	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// `ec2.MainRouteTableAssociation`.
	MainRouteTableId pulumi.StringPtrInput
	// The ID of the AWS account that owns the VPC.
	OwnerId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VpcState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcState)(nil)).Elem()
}

type vpcArgs struct {
	// Requests an Amazon-provided IPv6 CIDR
	// block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
	// the size of the CIDR block. Default is `false`.
	AssignGeneratedIpv6CidrBlock *bool `pulumi:"assignGeneratedIpv6CidrBlock"`
	// The CIDR block for the VPC.
	CidrBlock string `pulumi:"cidrBlock"`
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	EnableClassiclink *bool `pulumi:"enableClassiclink"`
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	EnableClassiclinkDnsSupport *bool `pulumi:"enableClassiclinkDnsSupport"`
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames *bool `pulumi:"enableDnsHostnames"`
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport *bool `pulumi:"enableDnsSupport"`
	// A tenancy option for instances launched into the VPC. Default is `default`, which
	// makes your instances shared on the host. Using either of the other options (`dedicated` or `host`) costs at least $2/hr.
	InstanceTenancy *string `pulumi:"instanceTenancy"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Vpc resource.
type VpcArgs struct {
	// Requests an Amazon-provided IPv6 CIDR
	// block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
	// the size of the CIDR block. Default is `false`.
	AssignGeneratedIpv6CidrBlock pulumi.BoolPtrInput
	// The CIDR block for the VPC.
	CidrBlock pulumi.StringInput
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	EnableClassiclink pulumi.BoolPtrInput
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	EnableClassiclinkDnsSupport pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport pulumi.BoolPtrInput
	// A tenancy option for instances launched into the VPC. Default is `default`, which
	// makes your instances shared on the host. Using either of the other options (`dedicated` or `host`) costs at least $2/hr.
	InstanceTenancy pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcArgs)(nil)).Elem()
}
