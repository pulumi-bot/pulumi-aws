// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package placementGroup

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an EC2 placement group. Read more about placement groups
// in [AWS Docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html).
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/placement_group.html.markdown.
type PlacementGroup struct {
	pulumi.CustomResourceState

	// The name of the placement group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The placement strategy.
	Strategy pulumi.StringOutput `pulumi:"strategy"`
}

// NewPlacementGroup registers a new resource with the given unique name, arguments, and options.
func NewPlacementGroup(ctx *pulumi.Context,
	name string, args *PlacementGroupArgs, opts ...pulumi.ResourceOption) (*PlacementGroup, error) {
	if args == nil || args.Strategy == nil {
		return nil, errors.New("missing required argument 'Strategy'")
	}
	if args == nil {
		args = &PlacementGroupArgs{}
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
	// The name of the placement group.
	Name *string `pulumi:"name"`
	// The placement strategy.
	Strategy *string `pulumi:"strategy"`
}

type PlacementGroupState struct {
	// The name of the placement group.
	Name pulumi.StringPtrInput
	// The placement strategy.
	Strategy pulumi.StringPtrInput
}

func (PlacementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*placementGroupState)(nil)).Elem()
}

type placementGroupArgs struct {
	// The name of the placement group.
	Name *string `pulumi:"name"`
	// The placement strategy.
	Strategy string `pulumi:"strategy"`
}

// The set of arguments for constructing a PlacementGroup resource.
type PlacementGroupArgs struct {
	// The name of the placement group.
	Name pulumi.StringPtrInput
	// The placement strategy.
	Strategy pulumi.StringInput
}

func (PlacementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*placementGroupArgs)(nil)).Elem()
}

