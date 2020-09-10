// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ParameterGroup struct {
	pulumi.CustomResourceState

	Arn         pulumi.StringOutput                `pulumi:"arn"`
	Description pulumi.StringOutput                `pulumi:"description"`
	Family      pulumi.StringOutput                `pulumi:"family"`
	Name        pulumi.StringOutput                `pulumi:"name"`
	NamePrefix  pulumi.StringOutput                `pulumi:"namePrefix"`
	Parameters  ParameterGroupParameterArrayOutput `pulumi:"parameters"`
	Tags        pulumi.StringMapOutput             `pulumi:"tags"`
}

// NewParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewParameterGroup(ctx *pulumi.Context,
	name string, args *ParameterGroupArgs, opts ...pulumi.ResourceOption) (*ParameterGroup, error) {
	if args == nil || args.Family == nil {
		return nil, errors.New("missing required argument 'Family'")
	}
	if args == nil {
		args = &ParameterGroupArgs{}
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource ParameterGroup
	err := ctx.RegisterResource("aws:rds/parameterGroup:ParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameterGroup gets an existing ParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParameterGroupState, opts ...pulumi.ResourceOption) (*ParameterGroup, error) {
	var resource ParameterGroup
	err := ctx.ReadResource("aws:rds/parameterGroup:ParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ParameterGroup resources.
type parameterGroupState struct {
	Arn         *string                   `pulumi:"arn"`
	Description *string                   `pulumi:"description"`
	Family      *string                   `pulumi:"family"`
	Name        *string                   `pulumi:"name"`
	NamePrefix  *string                   `pulumi:"namePrefix"`
	Parameters  []ParameterGroupParameter `pulumi:"parameters"`
	Tags        map[string]string         `pulumi:"tags"`
}

type ParameterGroupState struct {
	Arn         pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	Family      pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	NamePrefix  pulumi.StringPtrInput
	Parameters  ParameterGroupParameterArrayInput
	Tags        pulumi.StringMapInput
}

func (ParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupState)(nil)).Elem()
}

type parameterGroupArgs struct {
	Description *string                   `pulumi:"description"`
	Family      string                    `pulumi:"family"`
	Name        *string                   `pulumi:"name"`
	NamePrefix  *string                   `pulumi:"namePrefix"`
	Parameters  []ParameterGroupParameter `pulumi:"parameters"`
	Tags        map[string]string         `pulumi:"tags"`
}

// The set of arguments for constructing a ParameterGroup resource.
type ParameterGroupArgs struct {
	Description pulumi.StringPtrInput
	Family      pulumi.StringInput
	Name        pulumi.StringPtrInput
	NamePrefix  pulumi.StringPtrInput
	Parameters  ParameterGroupParameterArrayInput
	Tags        pulumi.StringMapInput
}

func (ParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupArgs)(nil)).Elem()
}
