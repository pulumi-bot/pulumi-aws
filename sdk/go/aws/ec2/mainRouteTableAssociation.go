// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource for managing the main routing table of a VPC.
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
// 		_, err := ec2.NewMainRouteTableAssociation(ctx, "mainRouteTableAssociation", &ec2.MainRouteTableAssociationArgs{
// 			VpcId:        pulumi.Any(aws_vpc.Foo.Id),
// 			RouteTableId: pulumi.Any(aws_route_table.Bar.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Notes
//
// On VPC creation, the AWS API always creates an initial Main Route Table. This
// resource records the ID of that Route Table under `originalRouteTableId`.
// The "Delete" action for a `mainRouteTableAssociation` consists of resetting
// this original table as the Main Route Table for the VPC. You'll see this
// additional Route Table in the AWS console; it must remain intact in order for
// the `mainRouteTableAssociation` delete to work properly.
type MainRouteTableAssociation struct {
	pulumi.CustomResourceState

	// Used internally, see __Notes__ below
	OriginalRouteTableId pulumi.StringOutput `pulumi:"originalRouteTableId"`
	// The ID of the Route Table to set as the new
	// main route table for the target VPC
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// The ID of the VPC whose main route table should be set
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewMainRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewMainRouteTableAssociation(ctx *pulumi.Context,
	name string, args *MainRouteTableAssociationArgs, opts ...pulumi.ResourceOption) (*MainRouteTableAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.RouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'RouteTableId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource MainRouteTableAssociation
	err := ctx.RegisterResource("aws:ec2/mainRouteTableAssociation:MainRouteTableAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMainRouteTableAssociation gets an existing MainRouteTableAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMainRouteTableAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MainRouteTableAssociationState, opts ...pulumi.ResourceOption) (*MainRouteTableAssociation, error) {
	var resource MainRouteTableAssociation
	err := ctx.ReadResource("aws:ec2/mainRouteTableAssociation:MainRouteTableAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MainRouteTableAssociation resources.
type mainRouteTableAssociationState struct {
	// Used internally, see __Notes__ below
	OriginalRouteTableId *string `pulumi:"originalRouteTableId"`
	// The ID of the Route Table to set as the new
	// main route table for the target VPC
	RouteTableId *string `pulumi:"routeTableId"`
	// The ID of the VPC whose main route table should be set
	VpcId *string `pulumi:"vpcId"`
}

type MainRouteTableAssociationState struct {
	// Used internally, see __Notes__ below
	OriginalRouteTableId pulumi.StringPtrInput
	// The ID of the Route Table to set as the new
	// main route table for the target VPC
	RouteTableId pulumi.StringPtrInput
	// The ID of the VPC whose main route table should be set
	VpcId pulumi.StringPtrInput
}

func (MainRouteTableAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*mainRouteTableAssociationState)(nil)).Elem()
}

type mainRouteTableAssociationArgs struct {
	// The ID of the Route Table to set as the new
	// main route table for the target VPC
	RouteTableId string `pulumi:"routeTableId"`
	// The ID of the VPC whose main route table should be set
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a MainRouteTableAssociation resource.
type MainRouteTableAssociationArgs struct {
	// The ID of the Route Table to set as the new
	// main route table for the target VPC
	RouteTableId pulumi.StringInput
	// The ID of the VPC whose main route table should be set
	VpcId pulumi.StringInput
}

func (MainRouteTableAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mainRouteTableAssociationArgs)(nil)).Elem()
}
