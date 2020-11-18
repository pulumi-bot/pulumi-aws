// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an EC2 placement group. Read more about placement groups
// in [AWS Docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html).
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
// 		_, err := ec2.NewPlacementGroup(ctx, "web", &ec2.PlacementGroupArgs{
// 			Strategy: pulumi.String("cluster"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type PlacementGroup struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the placement group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the placement group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the placement group.
	PlacementGroupId pulumi.StringOutput `pulumi:"placementGroupId"`
	// The placement strategy. Can be `"cluster"`, `"partition"` or `"spread"`.
	Strategy pulumi.StringOutput `pulumi:"strategy"`
	// Key-value map of resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewPlacementGroup registers a new resource with the given unique name, arguments, and options.
func NewPlacementGroup(ctx *pulumi.Context,
	name string, args *PlacementGroupArgs, opts ...pulumi.ResourceOption) (*PlacementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Strategy == nil {
		return nil, errors.New("invalid value for required argument 'Strategy'")
	}
	var resource PlacementGroup
	err := ctx.RegisterResource("aws:ec2/placementGroup:PlacementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlacementGroup gets an existing PlacementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlacementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlacementGroupState, opts ...pulumi.ResourceOption) (*PlacementGroup, error) {
	var resource PlacementGroup
	err := ctx.ReadResource("aws:ec2/placementGroup:PlacementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PlacementGroup resources.
type placementGroupState struct {
	// Amazon Resource Name (ARN) of the placement group.
	Arn *string `pulumi:"arn"`
	// The name of the placement group.
	Name *string `pulumi:"name"`
	// The ID of the placement group.
	PlacementGroupId *string `pulumi:"placementGroupId"`
	// The placement strategy. Can be `"cluster"`, `"partition"` or `"spread"`.
	Strategy *string `pulumi:"strategy"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

type PlacementGroupState struct {
	// Amazon Resource Name (ARN) of the placement group.
	Arn pulumi.StringPtrInput
	// The name of the placement group.
	Name pulumi.StringPtrInput
	// The ID of the placement group.
	PlacementGroupId pulumi.StringPtrInput
	// The placement strategy. Can be `"cluster"`, `"partition"` or `"spread"`.
	Strategy pulumi.StringPtrInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
}

func (PlacementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*placementGroupState)(nil)).Elem()
}

type placementGroupArgs struct {
	// The name of the placement group.
	Name *string `pulumi:"name"`
	// The placement strategy. Can be `"cluster"`, `"partition"` or `"spread"`.
	Strategy string `pulumi:"strategy"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PlacementGroup resource.
type PlacementGroupArgs struct {
	// The name of the placement group.
	Name pulumi.StringPtrInput
	// The placement strategy. Can be `"cluster"`, `"partition"` or `"spread"`.
	Strategy pulumi.StringInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
}

func (PlacementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*placementGroupArgs)(nil)).Elem()
}

type PlacementGroupInput interface {
	pulumi.Input

	ToPlacementGroupOutput() PlacementGroupOutput
	ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput
}

func (PlacementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementGroup)(nil)).Elem()
}

func (i PlacementGroup) ToPlacementGroupOutput() PlacementGroupOutput {
	return i.ToPlacementGroupOutputWithContext(context.Background())
}

func (i PlacementGroup) ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementGroupOutput)
}

type PlacementGroupOutput struct {
	*pulumi.OutputState
}

func (PlacementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementGroupOutput)(nil)).Elem()
}

func (o PlacementGroupOutput) ToPlacementGroupOutput() PlacementGroupOutput {
	return o
}

func (o PlacementGroupOutput) ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PlacementGroupOutput{})
}
