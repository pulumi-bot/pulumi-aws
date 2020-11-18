// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Resource Access Manager (RAM) Resource Association.
//
// > *NOTE:* Certain AWS resources (e.g. EC2 Subnets) can only be shared in an AWS account that is a member of an AWS Organizations organization with organization-wide Resource Access Manager functionality enabled. See the [Resource Access Manager User Guide](https://docs.aws.amazon.com/ram/latest/userguide/what-is.html) and AWS service specific documentation for additional information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ram"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ram.NewResourceAssociation(ctx, "example", &ram.ResourceAssociationArgs{
// 			ResourceArn:      pulumi.Any(aws_subnet.Example.Arn),
// 			ResourceShareArn: pulumi.Any(aws_ram_resource_share.Example.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ResourceAssociation struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
	// Amazon Resource Name (ARN) of the RAM Resource Share.
	ResourceShareArn pulumi.StringOutput `pulumi:"resourceShareArn"`
}

// NewResourceAssociation registers a new resource with the given unique name, arguments, and options.
func NewResourceAssociation(ctx *pulumi.Context,
	name string, args *ResourceAssociationArgs, opts ...pulumi.ResourceOption) (*ResourceAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	if args.ResourceShareArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceShareArn'")
	}
	var resource ResourceAssociation
	err := ctx.RegisterResource("aws:ram/resourceAssociation:ResourceAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceAssociation gets an existing ResourceAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceAssociationState, opts ...pulumi.ResourceOption) (*ResourceAssociation, error) {
	var resource ResourceAssociation
	err := ctx.ReadResource("aws:ram/resourceAssociation:ResourceAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceAssociation resources.
type resourceAssociationState struct {
	// Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
	ResourceArn *string `pulumi:"resourceArn"`
	// Amazon Resource Name (ARN) of the RAM Resource Share.
	ResourceShareArn *string `pulumi:"resourceShareArn"`
}

type ResourceAssociationState struct {
	// Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
	ResourceArn pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the RAM Resource Share.
	ResourceShareArn pulumi.StringPtrInput
}

func (ResourceAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceAssociationState)(nil)).Elem()
}

type resourceAssociationArgs struct {
	// Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
	ResourceArn string `pulumi:"resourceArn"`
	// Amazon Resource Name (ARN) of the RAM Resource Share.
	ResourceShareArn string `pulumi:"resourceShareArn"`
}

// The set of arguments for constructing a ResourceAssociation resource.
type ResourceAssociationArgs struct {
	// Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
	ResourceArn pulumi.StringInput
	// Amazon Resource Name (ARN) of the RAM Resource Share.
	ResourceShareArn pulumi.StringInput
}

func (ResourceAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceAssociationArgs)(nil)).Elem()
}

type ResourceAssociationInput interface {
	pulumi.Input

	ToResourceAssociationOutput() ResourceAssociationOutput
	ToResourceAssociationOutputWithContext(ctx context.Context) ResourceAssociationOutput
}

func (ResourceAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceAssociation)(nil)).Elem()
}

func (i ResourceAssociation) ToResourceAssociationOutput() ResourceAssociationOutput {
	return i.ToResourceAssociationOutputWithContext(context.Background())
}

func (i ResourceAssociation) ToResourceAssociationOutputWithContext(ctx context.Context) ResourceAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceAssociationOutput)
}

type ResourceAssociationOutput struct {
	*pulumi.OutputState
}

func (ResourceAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceAssociationOutput)(nil)).Elem()
}

func (o ResourceAssociationOutput) ToResourceAssociationOutput() ResourceAssociationOutput {
	return o
}

func (o ResourceAssociationOutput) ToResourceAssociationOutputWithContext(ctx context.Context) ResourceAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourceAssociationOutput{})
}
