// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dax

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DAX Subnet Group resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/dax"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dax.NewSubnetGroup(ctx, "example", &dax.SubnetGroupArgs{
// 			SubnetIds: pulumi.StringArray{
// 				pulumi.Any(aws_subnet.Example1.Id),
// 				pulumi.Any(aws_subnet.Example2.Id),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SubnetGroup struct {
	pulumi.CustomResourceState

	// A description of the subnet group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the subnet group.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of VPC subnet IDs for the subnet group.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// VPC ID of the subnet group.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewSubnetGroup(ctx *pulumi.Context,
	name string, args *SubnetGroupArgs, opts ...pulumi.ResourceOption) (*SubnetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	var resource SubnetGroup
	err := ctx.RegisterResource("aws:dax/subnetGroup:SubnetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetGroup gets an existing SubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetGroupState, opts ...pulumi.ResourceOption) (*SubnetGroup, error) {
	var resource SubnetGroup
	err := ctx.ReadResource("aws:dax/subnetGroup:SubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetGroup resources.
type subnetGroupState struct {
	// A description of the subnet group.
	Description *string `pulumi:"description"`
	// The name of the subnet group.
	Name *string `pulumi:"name"`
	// A list of VPC subnet IDs for the subnet group.
	SubnetIds []string `pulumi:"subnetIds"`
	// VPC ID of the subnet group.
	VpcId *string `pulumi:"vpcId"`
}

type SubnetGroupState struct {
	// A description of the subnet group.
	Description pulumi.StringPtrInput
	// The name of the subnet group.
	Name pulumi.StringPtrInput
	// A list of VPC subnet IDs for the subnet group.
	SubnetIds pulumi.StringArrayInput
	// VPC ID of the subnet group.
	VpcId pulumi.StringPtrInput
}

func (SubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetGroupState)(nil)).Elem()
}

type subnetGroupArgs struct {
	// A description of the subnet group.
	Description *string `pulumi:"description"`
	// The name of the subnet group.
	Name *string `pulumi:"name"`
	// A list of VPC subnet IDs for the subnet group.
	SubnetIds []string `pulumi:"subnetIds"`
}

// The set of arguments for constructing a SubnetGroup resource.
type SubnetGroupArgs struct {
	// A description of the subnet group.
	Description pulumi.StringPtrInput
	// The name of the subnet group.
	Name pulumi.StringPtrInput
	// A list of VPC subnet IDs for the subnet group.
	SubnetIds pulumi.StringArrayInput
}

func (SubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetGroupArgs)(nil)).Elem()
}
