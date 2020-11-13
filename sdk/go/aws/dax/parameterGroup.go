// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dax

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DAX Parameter Group resource.
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
// 		_, err := dax.NewParameterGroup(ctx, "example", &dax.ParameterGroupArgs{
// 			Parameters: dax.ParameterGroupParameterArray{
// 				&dax.ParameterGroupParameterArgs{
// 					Name:  pulumi.String("query-ttl-millis"),
// 					Value: pulumi.String("100000"),
// 				},
// 				&dax.ParameterGroupParameterArgs{
// 					Name:  pulumi.String("record-ttl-millis"),
// 					Value: pulumi.String("100000"),
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
type ParameterGroup struct {
	pulumi.CustomResourceState

	// A description of the parameter group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the parameter group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameters of the parameter group.
	Parameters ParameterGroupParameterArrayOutput `pulumi:"parameters"`
}

// NewParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewParameterGroup(ctx *pulumi.Context,
	name string, args *ParameterGroupArgs, opts ...pulumi.ResourceOption) (*ParameterGroup, error) {
	if args == nil {
		args = &ParameterGroupArgs{}
	}
	var resource ParameterGroup
	err := ctx.RegisterResource("aws:dax/parameterGroup:ParameterGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:dax/parameterGroup:ParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ParameterGroup resources.
type parameterGroupState struct {
	// A description of the parameter group.
	Description *string `pulumi:"description"`
	// The name of the parameter group.
	Name *string `pulumi:"name"`
	// The parameters of the parameter group.
	Parameters []ParameterGroupParameter `pulumi:"parameters"`
}

type ParameterGroupState struct {
	// A description of the parameter group.
	Description pulumi.StringPtrInput
	// The name of the parameter group.
	Name pulumi.StringPtrInput
	// The parameters of the parameter group.
	Parameters ParameterGroupParameterArrayInput
}

func (ParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupState)(nil)).Elem()
}

type parameterGroupArgs struct {
	// A description of the parameter group.
	Description *string `pulumi:"description"`
	// The name of the parameter group.
	Name *string `pulumi:"name"`
	// The parameters of the parameter group.
	Parameters []ParameterGroupParameter `pulumi:"parameters"`
}

// The set of arguments for constructing a ParameterGroup resource.
type ParameterGroupArgs struct {
	// A description of the parameter group.
	Description pulumi.StringPtrInput
	// The name of the parameter group.
	Name pulumi.StringPtrInput
	// The parameters of the parameter group.
	Parameters ParameterGroupParameterArrayInput
}

func (ParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupArgs)(nil)).Elem()
}

type ParameterGroupInput interface {
	pulumi.Input

	ToParameterGroupOutput() ParameterGroupOutput
	ToParameterGroupOutputWithContext(ctx context.Context) ParameterGroupOutput
}

func (ParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroup)(nil)).Elem()
}

func (i ParameterGroup) ToParameterGroupOutput() ParameterGroupOutput {
	return i.ToParameterGroupOutputWithContext(context.Background())
}

func (i ParameterGroup) ToParameterGroupOutputWithContext(ctx context.Context) ParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupOutput)
}

type ParameterGroupOutput struct {
	*pulumi.OutputState
}

func (ParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupOutput)(nil)).Elem()
}

func (o ParameterGroupOutput) ToParameterGroupOutput() ParameterGroupOutput {
	return o
}

func (o ParameterGroupOutput) ToParameterGroupOutputWithContext(ctx context.Context) ParameterGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ParameterGroupOutput{})
}
