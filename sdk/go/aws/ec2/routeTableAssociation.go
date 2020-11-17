// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to create an association between a route table and a subnet or a route table and an
// internet gateway or virtual private gateway.
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
// 		_, err := ec2.NewRouteTableAssociation(ctx, "routeTableAssociation", &ec2.RouteTableAssociationArgs{
// 			SubnetId:     pulumi.Any(aws_subnet.Foo.Id),
// 			RouteTableId: pulumi.Any(aws_route_table.Bar.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
// 		_, err := ec2.NewRouteTableAssociation(ctx, "routeTableAssociation", &ec2.RouteTableAssociationArgs{
// 			GatewayId:    pulumi.Any(aws_internet_gateway.Foo.Id),
// 			RouteTableId: pulumi.Any(aws_route_table.Bar.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// is already associated, will result in an error (e.g., `Resource.AlreadyAssociatedthe specified association for route table rtb-4176657279 conflicts with an existing association`) unless you first import the original association. EC2 Route Table Associations can be imported using the associated resource ID and Route Table ID separated by a forward slash (`/`). For example with EC2 Subnets
//
// ```sh
//  $ pulumi import aws:ec2/routeTableAssociation:RouteTableAssociation assoc subnet-6777656e646f6c796e/rtb-656c65616e6f72
// ```
//
//  For example with EC2 Internet Gateways
//
// ```sh
//  $ pulumi import aws:ec2/routeTableAssociation:RouteTableAssociation assoc igw-01b3a60780f8d034a/rtb-656c65616e6f72
// ```
type RouteTableAssociation struct {
	pulumi.CustomResourceState

	// The gateway ID to create an association. Conflicts with `subnetId`.
	GatewayId pulumi.StringPtrOutput `pulumi:"gatewayId"`
	// The ID of the routing table to associate with.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// The subnet ID to create an association. Conflicts with `gatewayId`.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
}

// NewRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewRouteTableAssociation(ctx *pulumi.Context,
	name string, args *RouteTableAssociationArgs, opts ...pulumi.ResourceOption) (*RouteTableAssociation, error) {
	if args == nil || args.RouteTableId == nil {
		return nil, errors.New("missing required argument 'RouteTableId'")
	}
	if args == nil {
		args = &RouteTableAssociationArgs{}
	}
	var resource RouteTableAssociation
	err := ctx.RegisterResource("aws:ec2/routeTableAssociation:RouteTableAssociation", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:ec2/routeTableAssociation:RouteTableAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteTableAssociation resources.
type routeTableAssociationState struct {
	// The gateway ID to create an association. Conflicts with `subnetId`.
	GatewayId *string `pulumi:"gatewayId"`
	// The ID of the routing table to associate with.
	RouteTableId *string `pulumi:"routeTableId"`
	// The subnet ID to create an association. Conflicts with `gatewayId`.
	SubnetId *string `pulumi:"subnetId"`
}

type RouteTableAssociationState struct {
	// The gateway ID to create an association. Conflicts with `subnetId`.
	GatewayId pulumi.StringPtrInput
	// The ID of the routing table to associate with.
	RouteTableId pulumi.StringPtrInput
	// The subnet ID to create an association. Conflicts with `gatewayId`.
	SubnetId pulumi.StringPtrInput
}

func (RouteTableAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableAssociationState)(nil)).Elem()
}

type routeTableAssociationArgs struct {
	// The gateway ID to create an association. Conflicts with `subnetId`.
	GatewayId *string `pulumi:"gatewayId"`
	// The ID of the routing table to associate with.
	RouteTableId string `pulumi:"routeTableId"`
	// The subnet ID to create an association. Conflicts with `gatewayId`.
	SubnetId *string `pulumi:"subnetId"`
}

// The set of arguments for constructing a RouteTableAssociation resource.
type RouteTableAssociationArgs struct {
	// The gateway ID to create an association. Conflicts with `subnetId`.
	GatewayId pulumi.StringPtrInput
	// The ID of the routing table to associate with.
	RouteTableId pulumi.StringInput
	// The subnet ID to create an association. Conflicts with `gatewayId`.
	SubnetId pulumi.StringPtrInput
}

func (RouteTableAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableAssociationArgs)(nil)).Elem()
}

type RouteTableAssociationInput interface {
	pulumi.Input

	ToRouteTableAssociationOutput() RouteTableAssociationOutput
	ToRouteTableAssociationOutputWithContext(ctx context.Context) RouteTableAssociationOutput
}

func (RouteTableAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTableAssociation)(nil)).Elem()
}

func (i RouteTableAssociation) ToRouteTableAssociationOutput() RouteTableAssociationOutput {
	return i.ToRouteTableAssociationOutputWithContext(context.Background())
}

func (i RouteTableAssociation) ToRouteTableAssociationOutputWithContext(ctx context.Context) RouteTableAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableAssociationOutput)
}

type RouteTableAssociationOutput struct {
	*pulumi.OutputState
}

func (RouteTableAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTableAssociationOutput)(nil)).Elem()
}

func (o RouteTableAssociationOutput) ToRouteTableAssociationOutput() RouteTableAssociationOutput {
	return o
}

func (o RouteTableAssociationOutput) ToRouteTableAssociationOutputWithContext(ctx context.Context) RouteTableAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RouteTableAssociationOutput{})
}
