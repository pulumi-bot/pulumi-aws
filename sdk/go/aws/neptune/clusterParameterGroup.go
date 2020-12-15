// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Neptune Cluster Parameter Group
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/neptune"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := neptune.NewClusterParameterGroup(ctx, "example", &neptune.ClusterParameterGroupArgs{
// 			Description: pulumi.String("neptune cluster parameter group"),
// 			Family:      pulumi.String("neptune1"),
// 			Parameters: neptune.ClusterParameterGroupParameterArray{
// 				&neptune.ClusterParameterGroupParameterArgs{
// 					Name:  pulumi.String("neptune_enable_audit_log"),
// 					Value: pulumi.String("1"),
// 				},
// 			},
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
// Neptune Cluster Parameter Groups can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:neptune/clusterParameterGroup:ClusterParameterGroup cluster_pg production-pg-1
// ```
type ClusterParameterGroup struct {
	pulumi.CustomResourceState

	// The ARN of the neptune cluster parameter group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the neptune cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The family of the neptune cluster parameter group.
	Family pulumi.StringOutput `pulumi:"family"`
	// The name of the neptune parameter.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// A list of neptune parameters to apply.
	Parameters ClusterParameterGroupParameterArrayOutput `pulumi:"parameters"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewClusterParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewClusterParameterGroup(ctx *pulumi.Context,
	name string, args *ClusterParameterGroupArgs, opts ...pulumi.ResourceOption) (*ClusterParameterGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Family == nil {
		return nil, errors.New("invalid value for required argument 'Family'")
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
	// The name of the neptune parameter.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of neptune parameters to apply.
	Parameters []ClusterParameterGroupParameter `pulumi:"parameters"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ClusterParameterGroupState struct {
	// The ARN of the neptune cluster parameter group.
	Arn pulumi.StringPtrInput
	// The description of the neptune cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the neptune cluster parameter group.
	Family pulumi.StringPtrInput
	// The name of the neptune parameter.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of neptune parameters to apply.
	Parameters ClusterParameterGroupParameterArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ClusterParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupState)(nil)).Elem()
}

type clusterParameterGroupArgs struct {
	// The description of the neptune cluster parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the neptune cluster parameter group.
	Family string `pulumi:"family"`
	// The name of the neptune parameter.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of neptune parameters to apply.
	Parameters []ClusterParameterGroupParameter `pulumi:"parameters"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterParameterGroup resource.
type ClusterParameterGroupArgs struct {
	// The description of the neptune cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the neptune cluster parameter group.
	Family pulumi.StringInput
	// The name of the neptune parameter.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of neptune parameters to apply.
	Parameters ClusterParameterGroupParameterArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ClusterParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupArgs)(nil)).Elem()
}

type ClusterParameterGroupInput interface {
	pulumi.Input

	ToClusterParameterGroupOutput() ClusterParameterGroupOutput
	ToClusterParameterGroupOutputWithContext(ctx context.Context) ClusterParameterGroupOutput
}

func (*ClusterParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParameterGroup)(nil))
}

func (i *ClusterParameterGroup) ToClusterParameterGroupOutput() ClusterParameterGroupOutput {
	return i.ToClusterParameterGroupOutputWithContext(context.Background())
}

func (i *ClusterParameterGroup) ToClusterParameterGroupOutputWithContext(ctx context.Context) ClusterParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupOutput)
}

type ClusterParameterGroupOutput struct {
	*pulumi.OutputState
}

func (ClusterParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParameterGroup)(nil))
}

func (o ClusterParameterGroupOutput) ToClusterParameterGroupOutput() ClusterParameterGroupOutput {
	return o
}

func (o ClusterParameterGroupOutput) ToClusterParameterGroupOutputWithContext(ctx context.Context) ClusterParameterGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ClusterParameterGroupOutput{})
}
