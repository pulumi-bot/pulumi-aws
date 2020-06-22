// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Neptune Parameter Group
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/neptune"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = neptune.NewParameterGroup(ctx, "example", &neptune.ParameterGroupArgs{
// 			Family: pulumi.String("neptune1"),
// 			Parameters: neptune.ParameterGroupParameterArray{
// 				&neptune.ParameterGroupParameterArgs{
// 					Name:  pulumi.String("neptune_query_timeout"),
// 					Value: pulumi.String("25"),
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

	// The Neptune parameter group Amazon Resource Name (ARN).
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the Neptune parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The family of the Neptune parameter group.
	Family pulumi.StringOutput `pulumi:"family"`
	// The name of the Neptune parameter.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of Neptune parameters to apply.
	Parameters ParameterGroupParameterArrayOutput `pulumi:"parameters"`
	// A map of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
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
	var resource ParameterGroup
	err := ctx.RegisterResource("aws:neptune/parameterGroup:ParameterGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:neptune/parameterGroup:ParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ParameterGroup resources.
type parameterGroupState struct {
	// The Neptune parameter group Amazon Resource Name (ARN).
	Arn *string `pulumi:"arn"`
	// The description of the Neptune parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the Neptune parameter group.
	Family *string `pulumi:"family"`
	// The name of the Neptune parameter.
	Name *string `pulumi:"name"`
	// A list of Neptune parameters to apply.
	Parameters []ParameterGroupParameter `pulumi:"parameters"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type ParameterGroupState struct {
	// The Neptune parameter group Amazon Resource Name (ARN).
	Arn pulumi.StringPtrInput
	// The description of the Neptune parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the Neptune parameter group.
	Family pulumi.StringPtrInput
	// The name of the Neptune parameter.
	Name pulumi.StringPtrInput
	// A list of Neptune parameters to apply.
	Parameters ParameterGroupParameterArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupState)(nil)).Elem()
}

type parameterGroupArgs struct {
	// The description of the Neptune parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the Neptune parameter group.
	Family string `pulumi:"family"`
	// The name of the Neptune parameter.
	Name *string `pulumi:"name"`
	// A list of Neptune parameters to apply.
	Parameters []ParameterGroupParameter `pulumi:"parameters"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ParameterGroup resource.
type ParameterGroupArgs struct {
	// The description of the Neptune parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the Neptune parameter group.
	Family pulumi.StringInput
	// The name of the Neptune parameter.
	Name pulumi.StringPtrInput
	// A list of Neptune parameters to apply.
	Parameters ParameterGroupParameterArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterGroupArgs)(nil)).Elem()
}
