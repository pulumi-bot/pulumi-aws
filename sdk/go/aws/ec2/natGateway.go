// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to create a VPC NAT Gateway.
//
// ## Example Usage
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
// 		_, err := ec2.NewNatGateway(ctx, "gw", &ec2.NatGatewayArgs{
// 			AllocationId: pulumi.Any(aws_eip.Nat.Id),
// 			SubnetId:     pulumi.Any(aws_subnet.Example.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Usage with tags:
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
// 		_, err := ec2.NewNatGateway(ctx, "gw", &ec2.NatGatewayArgs{
// 			AllocationId: pulumi.Any(aws_eip.Nat.Id),
// 			SubnetId:     pulumi.Any(aws_subnet.Example.Id),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("gw NAT"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NatGateway struct {
	pulumi.CustomResourceState

	// The Allocation ID of the Elastic IP address for the gateway.
	AllocationId pulumi.StringOutput `pulumi:"allocationId"`
	// The ENI ID of the network interface created by the NAT gateway.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The private IP address of the NAT Gateway.
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// The public IP address of the NAT Gateway.
	PublicIp pulumi.StringOutput `pulumi:"publicIp"`
	// The Subnet ID of the subnet in which to place the gateway.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewNatGateway registers a new resource with the given unique name, arguments, and options.
func NewNatGateway(ctx *pulumi.Context,
	name string, args *NatGatewayArgs, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	if args == nil || args.AllocationId == nil {
		return nil, errors.New("missing required argument 'AllocationId'")
	}
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	if args == nil {
		args = &NatGatewayArgs{}
	}
	var resource NatGateway
	err := ctx.RegisterResource("aws:ec2/natGateway:NatGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatGateway gets an existing NatGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatGatewayState, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	var resource NatGateway
	err := ctx.ReadResource("aws:ec2/natGateway:NatGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatGateway resources.
type natGatewayState struct {
	// The Allocation ID of the Elastic IP address for the gateway.
	AllocationId *string `pulumi:"allocationId"`
	// The ENI ID of the network interface created by the NAT gateway.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The private IP address of the NAT Gateway.
	PrivateIp *string `pulumi:"privateIp"`
	// The public IP address of the NAT Gateway.
	PublicIp *string `pulumi:"publicIp"`
	// The Subnet ID of the subnet in which to place the gateway.
	SubnetId *string `pulumi:"subnetId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type NatGatewayState struct {
	// The Allocation ID of the Elastic IP address for the gateway.
	AllocationId pulumi.StringPtrInput
	// The ENI ID of the network interface created by the NAT gateway.
	NetworkInterfaceId pulumi.StringPtrInput
	// The private IP address of the NAT Gateway.
	PrivateIp pulumi.StringPtrInput
	// The public IP address of the NAT Gateway.
	PublicIp pulumi.StringPtrInput
	// The Subnet ID of the subnet in which to place the gateway.
	SubnetId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (NatGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayState)(nil)).Elem()
}

type natGatewayArgs struct {
	// The Allocation ID of the Elastic IP address for the gateway.
	AllocationId string `pulumi:"allocationId"`
	// The Subnet ID of the subnet in which to place the gateway.
	SubnetId string `pulumi:"subnetId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NatGateway resource.
type NatGatewayArgs struct {
	// The Allocation ID of the Elastic IP address for the gateway.
	AllocationId pulumi.StringInput
	// The Subnet ID of the subnet in which to place the gateway.
	SubnetId pulumi.StringInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (NatGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayArgs)(nil)).Elem()
}

type NatGatewayInput interface {
	pulumi.Input

	ToNatGatewayOutput() NatGatewayOutput
	ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput
}

func (NatGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*NatGateway)(nil)).Elem()
}

func (i NatGateway) ToNatGatewayOutput() NatGatewayOutput {
	return i.ToNatGatewayOutputWithContext(context.Background())
}

func (i NatGateway) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayOutput)
}

type NatGatewayOutput struct {
	*pulumi.OutputState
}

func (NatGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NatGatewayOutput)(nil)).Elem()
}

func (o NatGatewayOutput) ToNatGatewayOutput() NatGatewayOutput {
	return o
}

func (o NatGatewayOutput) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NatGatewayOutput{})
}
