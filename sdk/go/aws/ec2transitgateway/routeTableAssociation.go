// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an EC2 Transit Gateway Route Table association.
//
// ## Example Usage
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
// 		_, err := ec2transitgateway.NewRouteTableAssociation(ctx, "example", &ec2transitgateway.RouteTableAssociationArgs{
// 			TransitGatewayAttachmentId: pulumi.Any(aws_ec2_transit_gateway_vpc_attachment.Example.Id),
// 			TransitGatewayRouteTableId: pulumi.Any(aws_ec2_transit_gateway_route_table.Example.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RouteTableAssociation struct {
	pulumi.CustomResourceState

	// Identifier of the resource
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Type of the resource
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId pulumi.StringOutput `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringOutput `pulumi:"transitGatewayRouteTableId"`
}

// NewRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewRouteTableAssociation(ctx *pulumi.Context,
	name string, args *RouteTableAssociationArgs, opts ...pulumi.ResourceOption) (*RouteTableAssociation, error) {
	if args == nil || args.TransitGatewayAttachmentId == nil {
		return nil, errors.New("missing required argument 'TransitGatewayAttachmentId'")
	}
	if args == nil || args.TransitGatewayRouteTableId == nil {
		return nil, errors.New("missing required argument 'TransitGatewayRouteTableId'")
	}
	if args == nil {
		args = &RouteTableAssociationArgs{}
	}
	var resource RouteTableAssociation
	err := ctx.RegisterResource("aws:ec2transitgateway/routeTableAssociation:RouteTableAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTableAssociation gets an existing RouteTableAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTableAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteTableAssociationState, opts ...pulumi.ResourceOption) (*RouteTableAssociation, error) {
	var resource RouteTableAssociation
	err := ctx.ReadResource("aws:ec2transitgateway/routeTableAssociation:RouteTableAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteTableAssociation resources.
type routeTableAssociationState struct {
	// Identifier of the resource
	ResourceId *string `pulumi:"resourceId"`
	// Type of the resource
	ResourceType *string `pulumi:"resourceType"`
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId *string `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId *string `pulumi:"transitGatewayRouteTableId"`
}

type RouteTableAssociationState struct {
	// Identifier of the resource
	ResourceId pulumi.StringPtrInput
	// Type of the resource
	ResourceType pulumi.StringPtrInput
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId pulumi.StringPtrInput
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringPtrInput
}

func (RouteTableAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableAssociationState)(nil)).Elem()
}

type routeTableAssociationArgs struct {
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId string `pulumi:"transitGatewayAttachmentId"`
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId string `pulumi:"transitGatewayRouteTableId"`
}

// The set of arguments for constructing a RouteTableAssociation resource.
type RouteTableAssociationArgs struct {
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId pulumi.StringInput
	// Identifier of EC2 Transit Gateway Route Table.
	TransitGatewayRouteTableId pulumi.StringInput
}

func (RouteTableAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableAssociationArgs)(nil)).Elem()
}
