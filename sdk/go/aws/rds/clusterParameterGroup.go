// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClusterParameterGroup struct {
	pulumi.CustomResourceState

	Arn         pulumi.StringOutput                       `pulumi:"arn"`
	Description pulumi.StringOutput                       `pulumi:"description"`
	Family      pulumi.StringOutput                       `pulumi:"family"`
	Name        pulumi.StringOutput                       `pulumi:"name"`
	NamePrefix  pulumi.StringOutput                       `pulumi:"namePrefix"`
	Parameters  ClusterParameterGroupParameterArrayOutput `pulumi:"parameters"`
	Tags        pulumi.StringMapOutput                    `pulumi:"tags"`
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
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource ClusterParameterGroup
	err := ctx.RegisterResource("aws:rds/clusterParameterGroup:ClusterParameterGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:rds/clusterParameterGroup:ClusterParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterParameterGroup resources.
type clusterParameterGroupState struct {
	Arn         *string                          `pulumi:"arn"`
	Description *string                          `pulumi:"description"`
	Family      *string                          `pulumi:"family"`
	Name        *string                          `pulumi:"name"`
	NamePrefix  *string                          `pulumi:"namePrefix"`
	Parameters  []ClusterParameterGroupParameter `pulumi:"parameters"`
	Tags        map[string]string                `pulumi:"tags"`
}

type ClusterParameterGroupState struct {
	Arn         pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	Family      pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	NamePrefix  pulumi.StringPtrInput
	Parameters  ClusterParameterGroupParameterArrayInput
	Tags        pulumi.StringMapInput
}

func (ClusterParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupState)(nil)).Elem()
}

type clusterParameterGroupArgs struct {
	Description *string                          `pulumi:"description"`
	Family      string                           `pulumi:"family"`
	Name        *string                          `pulumi:"name"`
	NamePrefix  *string                          `pulumi:"namePrefix"`
	Parameters  []ClusterParameterGroupParameter `pulumi:"parameters"`
	Tags        map[string]string                `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterParameterGroup resource.
type ClusterParameterGroupArgs struct {
	Description pulumi.StringPtrInput
	Family      pulumi.StringInput
	Name        pulumi.StringPtrInput
	NamePrefix  pulumi.StringPtrInput
	Parameters  ClusterParameterGroupParameterArrayInput
	Tags        pulumi.StringMapInput
}

func (ClusterParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupArgs)(nil)).Elem()
}
