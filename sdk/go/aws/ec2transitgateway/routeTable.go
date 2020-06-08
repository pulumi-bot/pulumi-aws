// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an EC2 Transit Gateway Route Table.
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := ec2transitgateway.NewRouteTable(ctx, "example", &ec2transitgateway.RouteTableArgs{
// 			TransitGatewayId: pulumi.String(aws_ec2_transit_gateway.Example.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RouteTable struct {
	pulumi.CustomResourceState

	// Boolean whether this is the default association route table for the EC2 Transit Gateway.
	DefaultAssociationRouteTable pulumi.BoolOutput `pulumi:"defaultAssociationRouteTable"`
	// Boolean whether this is the default propagation route table for the EC2 Transit Gateway.
	DefaultPropagationRouteTable pulumi.BoolOutput `pulumi:"defaultPropagationRouteTable"`
	// Key-value tags for the EC2 Transit Gateway Route Table.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId pulumi.StringOutput `pulumi:"transitGatewayId"`
}

// NewRouteTable registers a new resource with the given unique name, arguments, and options.
func NewRouteTable(ctx *pulumi.Context,
	name string, args *RouteTableArgs, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	if args == nil || args.TransitGatewayId == nil {
		return nil, errors.New("missing required argument 'TransitGatewayId'")
	}
	if args == nil {
		args = &RouteTableArgs{}
	}
	var resource RouteTable
	err := ctx.RegisterResource("aws:ec2transitgateway/routeTable:RouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTable gets an existing RouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteTableState, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	var resource RouteTable
	err := ctx.ReadResource("aws:ec2transitgateway/routeTable:RouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteTable resources.
type routeTableState struct {
	// Boolean whether this is the default association route table for the EC2 Transit Gateway.
	DefaultAssociationRouteTable *bool `pulumi:"defaultAssociationRouteTable"`
	// Boolean whether this is the default propagation route table for the EC2 Transit Gateway.
	DefaultPropagationRouteTable *bool `pulumi:"defaultPropagationRouteTable"`
	// Key-value tags for the EC2 Transit Gateway Route Table.
	Tags map[string]interface{} `pulumi:"tags"`
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId *string `pulumi:"transitGatewayId"`
}

type RouteTableState struct {
	// Boolean whether this is the default association route table for the EC2 Transit Gateway.
	DefaultAssociationRouteTable pulumi.BoolPtrInput
	// Boolean whether this is the default propagation route table for the EC2 Transit Gateway.
	DefaultPropagationRouteTable pulumi.BoolPtrInput
	// Key-value tags for the EC2 Transit Gateway Route Table.
	Tags pulumi.MapInput
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId pulumi.StringPtrInput
}

func (RouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableState)(nil)).Elem()
}

type routeTableArgs struct {
	// Key-value tags for the EC2 Transit Gateway Route Table.
	Tags map[string]interface{} `pulumi:"tags"`
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId string `pulumi:"transitGatewayId"`
}

// The set of arguments for constructing a RouteTable resource.
type RouteTableArgs struct {
	// Key-value tags for the EC2 Transit Gateway Route Table.
	Tags pulumi.MapInput
	// Identifier of EC2 Transit Gateway.
	TransitGatewayId pulumi.StringInput
}

func (RouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableArgs)(nil)).Elem()
}
