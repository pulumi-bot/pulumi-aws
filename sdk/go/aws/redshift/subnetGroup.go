// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a new Amazon Redshift subnet group. You must provide a list of one or more subnets in your existing Amazon Virtual Private Cloud (Amazon VPC) when creating Amazon Redshift subnet group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/redshift"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooVpc, err := ec2.NewVpc(ctx, "fooVpc", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.1.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fooSubnet, err := ec2.NewSubnet(ctx, "fooSubnet", &ec2.SubnetArgs{
// 			CidrBlock:        pulumi.String("10.1.1.0/24"),
// 			AvailabilityZone: pulumi.String("us-west-2a"),
// 			VpcId:            fooVpc.ID(),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("tf-dbsubnet-test-1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		bar, err := ec2.NewSubnet(ctx, "bar", &ec2.SubnetArgs{
// 			CidrBlock:        pulumi.String("10.1.2.0/24"),
// 			AvailabilityZone: pulumi.String("us-west-2b"),
// 			VpcId:            fooVpc.ID(),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("tf-dbsubnet-test-2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = redshift.NewSubnetGroup(ctx, "fooSubnetGroup", &redshift.SubnetGroupArgs{
// 			SubnetIds: pulumi.StringArray{
// 				fooSubnet.ID(),
// 				bar.ID(),
// 			},
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("Production"),
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

	// Amazon Resource Name (ARN) of the Redshift Subnet group name
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the Redshift Subnet group.
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of VPC subnet IDs.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
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
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource SubnetGroup
	err := ctx.RegisterResource("aws:redshift/subnetGroup:SubnetGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:redshift/subnetGroup:SubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetGroup resources.
type subnetGroupState struct {
	// Amazon Resource Name (ARN) of the Redshift Subnet group name
	Arn *string `pulumi:"arn"`
	// The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The name of the Redshift Subnet group.
	Name *string `pulumi:"name"`
	// An array of VPC subnet IDs.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type SubnetGroupState struct {
	// Amazon Resource Name (ARN) of the Redshift Subnet group name
	Arn pulumi.StringPtrInput
	// The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The name of the Redshift Subnet group.
	Name pulumi.StringPtrInput
	// An array of VPC subnet IDs.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetGroupState)(nil)).Elem()
}

type subnetGroupArgs struct {
	// The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The name of the Redshift Subnet group.
	Name *string `pulumi:"name"`
	// An array of VPC subnet IDs.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SubnetGroup resource.
type SubnetGroupArgs struct {
	// The description of the Redshift Subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The name of the Redshift Subnet group.
	Name pulumi.StringPtrInput
	// An array of VPC subnet IDs.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetGroupArgs)(nil)).Elem()
}

type SubnetGroupInput interface {
	pulumi.Input

	ToSubnetGroupOutput() SubnetGroupOutput
	ToSubnetGroupOutputWithContext(ctx context.Context) SubnetGroupOutput
}

func (SubnetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetGroup)(nil)).Elem()
}

func (i SubnetGroup) ToSubnetGroupOutput() SubnetGroupOutput {
	return i.ToSubnetGroupOutputWithContext(context.Background())
}

func (i SubnetGroup) ToSubnetGroupOutputWithContext(ctx context.Context) SubnetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetGroupOutput)
}

type SubnetGroupOutput struct {
	*pulumi.OutputState
}

func (SubnetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetGroupOutput)(nil)).Elem()
}

func (o SubnetGroupOutput) ToSubnetGroupOutput() SubnetGroupOutput {
	return o
}

func (o SubnetGroupOutput) ToSubnetGroupOutputWithContext(ctx context.Context) SubnetGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubnetGroupOutput{})
}
