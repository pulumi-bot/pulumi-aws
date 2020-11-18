// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an EC2 Transit Gateway Route.
//
// ## Example Usage
// ### Standard usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2transitgateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2transitgateway.NewRoute(ctx, "example", &ec2transitgateway.RouteArgs{
// 			DestinationCidrBlock:       pulumi.String("0.0.0.0/0"),
// 			TransitGatewayAttachmentId: pulumi.Any(aws_ec2_transit_gateway_vpc_attachment.Example.Id),
// 			TransitGatewayRouteTableId: pulumi.Any(aws_ec2_transit_gateway.Example.Association_default_route_table_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Blackhole route
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2transitgateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2transitgateway.NewRoute(ctx, "example", &ec2transitgateway.RouteArgs{
// 			DestinationCidrBlock:       pulumi.String("0.0.0.0/0"),
// 			Blackhole:                  pulumi.Bool(true),
// 			TransitGatewayRouteTableId: pulumi.Any(aws_ec2_transit_gateway.Example.Association_default_route_table_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Route struct {
	pulumi.CustomResourceState

	// Indicates whether to drop traffic that matches this route (default to `false`).
	Blackhole pulumi.BoolPtrOutput `pulumi:"blackhole"`
	// IPv4 or IPv6 RFC1924 CIDR used for destination matches. Routing decisions are based on the most specific match.
	DestinationCidrBlock pulumi.StringOutput `pulumi:"destinationCidrBlock"`
	// Identifier of EC2 Transit Gateway Attachment (required if `blackhole` is set to false).
	TransitGatewayAttachmentId pulumi.StringPtrOutput `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringOutput `pulumi:"transitGatewayRouteTableId"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'DestinationCidrBlock'")
	}
	if args.TransitGatewayRouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayRouteTableId'")
	}
	var resource Route
	err := ctx.RegisterResource("aws:ec2transitgateway/route:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("aws:ec2transitgateway/route:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// Indicates whether to drop traffic that matches this route (default to `false`).
	Blackhole *bool `pulumi:"blackhole"`
	// IPv4 or IPv6 RFC1924 CIDR used for destination matches. Routing decisions are based on the most specific match.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// Identifier of EC2 Transit Gateway Attachment (required if `blackhole` is set to false).
	TransitGatewayAttachmentId *string `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId *string `pulumi:"transitGatewayRouteTableId"`
}

type RouteState struct {
	// Indicates whether to drop traffic that matches this route (default to `false`).
	Blackhole pulumi.BoolPtrInput
	// IPv4 or IPv6 RFC1924 CIDR used for destination matches. Routing decisions are based on the most specific match.
	DestinationCidrBlock pulumi.StringPtrInput
	// Identifier of EC2 Transit Gateway Attachment (required if `blackhole` is set to false).
	TransitGatewayAttachmentId pulumi.StringPtrInput
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// Indicates whether to drop traffic that matches this route (default to `false`).
	Blackhole *bool `pulumi:"blackhole"`
	// IPv4 or IPv6 RFC1924 CIDR used for destination matches. Routing decisions are based on the most specific match.
	DestinationCidrBlock string `pulumi:"destinationCidrBlock"`
	// Identifier of EC2 Transit Gateway Attachment (required if `blackhole` is set to false).
	TransitGatewayAttachmentId *string `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId string `pulumi:"transitGatewayRouteTableId"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// Indicates whether to drop traffic that matches this route (default to `false`).
	Blackhole pulumi.BoolPtrInput
	// IPv4 or IPv6 RFC1924 CIDR used for destination matches. Routing decisions are based on the most specific match.
	DestinationCidrBlock pulumi.StringInput
	// Identifier of EC2 Transit Gateway Attachment (required if `blackhole` is set to false).
	TransitGatewayAttachmentId pulumi.StringPtrInput
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (Route) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil)).Elem()
}

func (i Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

type RouteOutput struct {
	*pulumi.OutputState
}

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteOutput)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RouteOutput{})
}
