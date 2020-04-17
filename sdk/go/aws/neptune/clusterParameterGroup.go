// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Neptune Cluster Parameter Group
type ClusterParameterGroup struct {
	pulumi.CustomResourceState

	// The ARN of the neptune cluster parameter group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the neptune cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The family of the neptune cluster parameter group.
	Family pulumi.StringOutput `pulumi:"family"`
	// The name of the neptune cluster parameter group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// A list of neptune parameters to apply.
	Parameters ClusterParameterGroupParameterArrayOutput `pulumi:"parameters"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewClusterParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewClusterParameterGroup(ctx *pulumi.Context,
	name string, args *ClusterParameterGroupArgs, opts ...pulumi.ResourceOption) (*ClusterParameterGroup, error) {
	if args == nil || args.Family == nil {
		return nil, errors.New("missing required argument 'Family'")
	}
	if args == nil {
		args = &ClusterParameterGroupArgs{}
	}
	var resource ClusterParameterGroup
	err := ctx.RegisterResource("aws:neptune/clusterParameterGroup:ClusterParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterParameterGroup gets an existing ClusterParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterParameterGroupState, opts ...pulumi.ResourceOption) (*ClusterParameterGroup, error) {
	var resource ClusterParameterGroup
	err := ctx.ReadResource("aws:neptune/clusterParameterGroup:ClusterParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterParameterGroup resources.
type clusterParameterGroupState struct {
	// The ARN of the neptune cluster parameter group.
	Arn *string `pulumi:"arn"`
	// The description of the neptune cluster parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the neptune cluster parameter group.
	Family *string `pulumi:"family"`
	// The name of the neptune cluster parameter group. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of neptune parameters to apply.
	Parameters []ClusterParameterGroupParameter `pulumi:"parameters"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ClusterParameterGroupState struct {
	// The ARN of the neptune cluster parameter group.
	Arn pulumi.StringPtrInput
	// The description of the neptune cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the neptune cluster parameter group.
	Family pulumi.StringPtrInput
	// The name of the neptune cluster parameter group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of neptune parameters to apply.
	Parameters ClusterParameterGroupParameterArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ClusterParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupState)(nil)).Elem()
}

type clusterParameterGroupArgs struct {
	// The description of the neptune cluster parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the neptune cluster parameter group.
	Family string `pulumi:"family"`
	// The name of the neptune cluster parameter group. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of neptune parameters to apply.
	Parameters []ClusterParameterGroupParameter `pulumi:"parameters"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterParameterGroup resource.
type ClusterParameterGroupArgs struct {
	// The description of the neptune cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the neptune cluster parameter group.
	Family pulumi.StringInput
	// The name of the neptune cluster parameter group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of neptune parameters to apply.
	Parameters ClusterParameterGroupParameterArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ClusterParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupArgs)(nil)).Elem()
}
