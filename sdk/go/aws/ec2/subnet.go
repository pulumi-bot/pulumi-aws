// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an VPC subnet resource.
//
// > **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), subnets associated with Lambda Functions can take up to 45 minutes to successfully delete.
//
// ## Example Usage
// ### Basic Usage
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
// 		_, err := ec2.NewSubnet(ctx, "main", &ec2.SubnetArgs{
// 			VpcId:     pulumi.Any(aws_vpc.Main.Id),
// 			CidrBlock: pulumi.String("10.0.1.0/24"),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("Main"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Subnets In Secondary VPC CIDR Blocks
//
// When managing subnets in one of a VPC's secondary CIDR blocks created using a `ec2.VpcIpv4CidrBlockAssociation`
// resource, it is recommended to reference that resource's `vpcId` attribute to ensure correct dependency ordering.
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
// 		secondaryCidr, err := ec2.NewVpcIpv4CidrBlockAssociation(ctx, "secondaryCidr", &ec2.VpcIpv4CidrBlockAssociationArgs{
// 			VpcId:     pulumi.Any(aws_vpc.Main.Id),
// 			CidrBlock: pulumi.String("172.2.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewSubnet(ctx, "inSecondaryCidr", &ec2.SubnetArgs{
// 			VpcId:     secondaryCidr.VpcId,
// 			CidrBlock: pulumi.String("172.2.0.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Subnet struct {
	pulumi.CustomResourceState

	// The ARN of the subnet.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specify true to indicate
	// that network interfaces created in the specified subnet should be
	// assigned an IPv6 address. Default is `false`
	AssignIpv6AddressOnCreation pulumi.BoolPtrOutput `pulumi:"assignIpv6AddressOnCreation"`
	// The AZ for the subnet.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The AZ ID of the subnet.
	AvailabilityZoneId pulumi.StringOutput `pulumi:"availabilityZoneId"`
	// The CIDR block for the subnet.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The IPv6 network range for the subnet,
	// in CIDR notation. The subnet size must use a /64 prefix length.
	Ipv6CidrBlock pulumi.StringPtrOutput `pulumi:"ipv6CidrBlock"`
	// The association ID for the IPv6 CIDR block.
	Ipv6CidrBlockAssociationId pulumi.StringOutput `pulumi:"ipv6CidrBlockAssociationId"`
	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address. Default is `false`.
	MapPublicIpOnLaunch pulumi.BoolPtrOutput `pulumi:"mapPublicIpOnLaunch"`
	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn pulumi.StringPtrOutput `pulumi:"outpostArn"`
	// The ID of the AWS account that owns the subnet.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewSubnet registers a new resource with the given unique name, arguments, and options.
func NewSubnet(ctx *pulumi.Context,
	name string, args *SubnetArgs, opts ...pulumi.ResourceOption) (*Subnet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource Subnet
	err := ctx.RegisterResource("aws:ec2/subnet:Subnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnet gets an existing Subnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetState, opts ...pulumi.ResourceOption) (*Subnet, error) {
	var resource Subnet
	err := ctx.ReadResource("aws:ec2/subnet:Subnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subnet resources.
type subnetState struct {
	// The ARN of the subnet.
	Arn *string `pulumi:"arn"`
	// Specify true to indicate
	// that network interfaces created in the specified subnet should be
	// assigned an IPv6 address. Default is `false`
	AssignIpv6AddressOnCreation *bool `pulumi:"assignIpv6AddressOnCreation"`
	// The AZ for the subnet.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The AZ ID of the subnet.
	AvailabilityZoneId *string `pulumi:"availabilityZoneId"`
	// The CIDR block for the subnet.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The IPv6 network range for the subnet,
	// in CIDR notation. The subnet size must use a /64 prefix length.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// The association ID for the IPv6 CIDR block.
	Ipv6CidrBlockAssociationId *string `pulumi:"ipv6CidrBlockAssociationId"`
	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address. Default is `false`.
	MapPublicIpOnLaunch *bool `pulumi:"mapPublicIpOnLaunch"`
	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string `pulumi:"outpostArn"`
	// The ID of the AWS account that owns the subnet.
	OwnerId *string `pulumi:"ownerId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

type SubnetState struct {
	// The ARN of the subnet.
	Arn pulumi.StringPtrInput
	// Specify true to indicate
	// that network interfaces created in the specified subnet should be
	// assigned an IPv6 address. Default is `false`
	AssignIpv6AddressOnCreation pulumi.BoolPtrInput
	// The AZ for the subnet.
	AvailabilityZone pulumi.StringPtrInput
	// The AZ ID of the subnet.
	AvailabilityZoneId pulumi.StringPtrInput
	// The CIDR block for the subnet.
	CidrBlock pulumi.StringPtrInput
	// The IPv6 network range for the subnet,
	// in CIDR notation. The subnet size must use a /64 prefix length.
	Ipv6CidrBlock pulumi.StringPtrInput
	// The association ID for the IPv6 CIDR block.
	Ipv6CidrBlockAssociationId pulumi.StringPtrInput
	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address. Default is `false`.
	MapPublicIpOnLaunch pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn pulumi.StringPtrInput
	// The ID of the AWS account that owns the subnet.
	OwnerId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VPC ID.
	VpcId pulumi.StringPtrInput
}

func (SubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetState)(nil)).Elem()
}

type subnetArgs struct {
	// Specify true to indicate
	// that network interfaces created in the specified subnet should be
	// assigned an IPv6 address. Default is `false`
	AssignIpv6AddressOnCreation *bool `pulumi:"assignIpv6AddressOnCreation"`
	// The AZ for the subnet.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The AZ ID of the subnet.
	AvailabilityZoneId *string `pulumi:"availabilityZoneId"`
	// The CIDR block for the subnet.
	CidrBlock string `pulumi:"cidrBlock"`
	// The IPv6 network range for the subnet,
	// in CIDR notation. The subnet size must use a /64 prefix length.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address. Default is `false`.
	MapPublicIpOnLaunch *bool `pulumi:"mapPublicIpOnLaunch"`
	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string `pulumi:"outpostArn"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VPC ID.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Subnet resource.
type SubnetArgs struct {
	// Specify true to indicate
	// that network interfaces created in the specified subnet should be
	// assigned an IPv6 address. Default is `false`
	AssignIpv6AddressOnCreation pulumi.BoolPtrInput
	// The AZ for the subnet.
	AvailabilityZone pulumi.StringPtrInput
	// The AZ ID of the subnet.
	AvailabilityZoneId pulumi.StringPtrInput
	// The CIDR block for the subnet.
	CidrBlock pulumi.StringInput
	// The IPv6 network range for the subnet,
	// in CIDR notation. The subnet size must use a /64 prefix length.
	Ipv6CidrBlock pulumi.StringPtrInput
	// Specify true to indicate
	// that instances launched into the subnet should be assigned
	// a public IP address. Default is `false`.
	MapPublicIpOnLaunch pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VPC ID.
	VpcId pulumi.StringInput
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetArgs)(nil)).Elem()
}

type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(ctx context.Context) SubnetOutput
}

func (Subnet) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (i Subnet) ToSubnetOutput() SubnetOutput {
	return i.ToSubnetOutputWithContext(context.Background())
}

func (i Subnet) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput)
}

type SubnetOutput struct {
	*pulumi.OutputState
}

func (SubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetOutput)(nil)).Elem()
}

func (o SubnetOutput) ToSubnetOutput() SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubnetOutput{})
}
