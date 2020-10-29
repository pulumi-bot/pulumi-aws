// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an DocumentDB subnet group resource.
type SubnetGroup struct {
	pulumi.CustomResourceState

	// The ARN of the docDB subnet group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the docDB subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the docDB subnet group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// A list of VPC subnet IDs.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewSubnetGroup(ctx *pulumi.Context,
	name string, args *SubnetGroupArgs, opts ...pulumi.ResourceOption) (*SubnetGroup, error) {
	if args == nil || args.SubnetIds == nil {
		return nil, errors.New("missing required argument 'SubnetIds'")
	}
	if args == nil {
		args = &SubnetGroupArgs{}
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource SubnetGroup
	err := ctx.RegisterResource("aws:docdb/subnetGroup:SubnetGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:docdb/subnetGroup:SubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetGroup resources.
type subnetGroupState struct {
	// The ARN of the docDB subnet group.
	Arn *string `pulumi:"arn"`
	// The description of the docDB subnet group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The name of the docDB subnet group. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of VPC subnet IDs.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type SubnetGroupState struct {
	// The ARN of the docDB subnet group.
	Arn pulumi.StringPtrInput
	// The description of the docDB subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The name of the docDB subnet group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of VPC subnet IDs.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetGroupState)(nil)).Elem()
}

type subnetGroupArgs struct {
	// The description of the docDB subnet group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The name of the docDB subnet group. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of VPC subnet IDs.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SubnetGroup resource.
type SubnetGroupArgs struct {
	// The description of the docDB subnet group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The name of the docDB subnet group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of VPC subnet IDs.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetGroupArgs)(nil)).Elem()
}
