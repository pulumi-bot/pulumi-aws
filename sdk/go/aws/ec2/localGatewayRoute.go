// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an EC2 Local Gateway Route. More information can be found in the [Outposts User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#routing).
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
// 		_, err := ec2.NewLocalGatewayRoute(ctx, "example", &ec2.LocalGatewayRouteArgs{
// 			DestinationCidrBlock:                pulumi.String("172.16.0.0/16"),
// 			LocalGatewayRouteTableId:            pulumi.Any(data.Aws_ec2_local_gateway_route_table.Example.Id),
// 			LocalGatewayVirtualInterfaceGroupId: pulumi.Any(data.Aws_ec2_local_gateway_virtual_interface_group.Example.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocalGatewayRoute struct {
	pulumi.CustomResourceState

	// IPv4 CIDR range used for destination matches. Routing decisions are based on the most specific match.
	DestinationCidrBlock pulumi.StringOutput `pulumi:"destinationCidrBlock"`
	// Identifier of EC2 Local Gateway Route Table.
	LocalGatewayRouteTableId pulumi.StringOutput `pulumi:"localGatewayRouteTableId"`
	// Identifier of EC2 Local Gateway Virtual Interface Group.
	LocalGatewayVirtualInterfaceGroupId pulumi.StringOutput `pulumi:"localGatewayVirtualInterfaceGroupId"`
}

// NewLocalGatewayRoute registers a new resource with the given unique name, arguments, and options.
func NewLocalGatewayRoute(ctx *pulumi.Context,
	name string, args *LocalGatewayRouteArgs, opts ...pulumi.ResourceOption) (*LocalGatewayRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'DestinationCidrBlock'")
	}
	if args.LocalGatewayRouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'LocalGatewayRouteTableId'")
	}
	if args.LocalGatewayVirtualInterfaceGroupId == nil {
		return nil, errors.New("invalid value for required argument 'LocalGatewayVirtualInterfaceGroupId'")
	}
	var resource LocalGatewayRoute
	err := ctx.RegisterResource("aws:ec2/localGatewayRoute:LocalGatewayRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalGatewayRoute gets an existing LocalGatewayRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalGatewayRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalGatewayRouteState, opts ...pulumi.ResourceOption) (*LocalGatewayRoute, error) {
	var resource LocalGatewayRoute
	err := ctx.ReadResource("aws:ec2/localGatewayRoute:LocalGatewayRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalGatewayRoute resources.
type localGatewayRouteState struct {
	// IPv4 CIDR range used for destination matches. Routing decisions are based on the most specific match.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// Identifier of EC2 Local Gateway Route Table.
	LocalGatewayRouteTableId *string `pulumi:"localGatewayRouteTableId"`
	// Identifier of EC2 Local Gateway Virtual Interface Group.
	LocalGatewayVirtualInterfaceGroupId *string `pulumi:"localGatewayVirtualInterfaceGroupId"`
}

type LocalGatewayRouteState struct {
	// IPv4 CIDR range used for destination matches. Routing decisions are based on the most specific match.
	DestinationCidrBlock pulumi.StringPtrInput
	// Identifier of EC2 Local Gateway Route Table.
	LocalGatewayRouteTableId pulumi.StringPtrInput
	// Identifier of EC2 Local Gateway Virtual Interface Group.
	LocalGatewayVirtualInterfaceGroupId pulumi.StringPtrInput
}

func (LocalGatewayRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*localGatewayRouteState)(nil)).Elem()
}

type localGatewayRouteArgs struct {
	// IPv4 CIDR range used for destination matches. Routing decisions are based on the most specific match.
	DestinationCidrBlock string `pulumi:"destinationCidrBlock"`
	// Identifier of EC2 Local Gateway Route Table.
	LocalGatewayRouteTableId string `pulumi:"localGatewayRouteTableId"`
	// Identifier of EC2 Local Gateway Virtual Interface Group.
	LocalGatewayVirtualInterfaceGroupId string `pulumi:"localGatewayVirtualInterfaceGroupId"`
}

// The set of arguments for constructing a LocalGatewayRoute resource.
type LocalGatewayRouteArgs struct {
	// IPv4 CIDR range used for destination matches. Routing decisions are based on the most specific match.
	DestinationCidrBlock pulumi.StringInput
	// Identifier of EC2 Local Gateway Route Table.
	LocalGatewayRouteTableId pulumi.StringInput
	// Identifier of EC2 Local Gateway Virtual Interface Group.
	LocalGatewayVirtualInterfaceGroupId pulumi.StringInput
}

func (LocalGatewayRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localGatewayRouteArgs)(nil)).Elem()
}

type LocalGatewayRouteInput interface {
	pulumi.Input

	ToLocalGatewayRouteOutput() LocalGatewayRouteOutput
	ToLocalGatewayRouteOutputWithContext(ctx context.Context) LocalGatewayRouteOutput
}

func (LocalGatewayRoute) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalGatewayRoute)(nil)).Elem()
}

func (i LocalGatewayRoute) ToLocalGatewayRouteOutput() LocalGatewayRouteOutput {
	return i.ToLocalGatewayRouteOutputWithContext(context.Background())
}

func (i LocalGatewayRoute) ToLocalGatewayRouteOutputWithContext(ctx context.Context) LocalGatewayRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGatewayRouteOutput)
}

type LocalGatewayRouteOutput struct {
	*pulumi.OutputState
}

func (LocalGatewayRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalGatewayRouteOutput)(nil)).Elem()
}

func (o LocalGatewayRouteOutput) ToLocalGatewayRouteOutput() LocalGatewayRouteOutput {
	return o
}

func (o LocalGatewayRouteOutput) ToLocalGatewayRouteOutputWithContext(ctx context.Context) LocalGatewayRouteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LocalGatewayRouteOutput{})
}
