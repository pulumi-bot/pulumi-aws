// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Gamelift Alias resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/gamelift"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gamelift.NewAlias(ctx, "example", &gamelift.AliasArgs{
// 			Description: pulumi.String("Example Description"),
// 			RoutingStrategy: &gamelift.AliasRoutingStrategyArgs{
// 				Message: pulumi.String("Example Message"),
// 				Type:    pulumi.String("TERMINAL"),
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
// Gamelift Aliases can be imported using the ID, e.g.
//
// ```sh
//  $ pulumi import aws:gamelift/alias:Alias example <alias-id>
// ```
type Alias struct {
	pulumi.CustomResourceState

	// Alias ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the alias.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the alias.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategyOutput `pulumi:"routingStrategy"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil || args.RoutingStrategy == nil {
		return nil, errors.New("missing required argument 'RoutingStrategy'")
	}
	if args == nil {
		args = &AliasArgs{}
	}
	var resource Alias
	err := ctx.RegisterResource("aws:gamelift/alias:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("aws:gamelift/alias:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type aliasState struct {
	// Alias ARN.
	Arn *string `pulumi:"arn"`
	// Description of the alias.
	Description *string `pulumi:"description"`
	// Name of the alias.
	Name *string `pulumi:"name"`
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy *AliasRoutingStrategy `pulumi:"routingStrategy"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type AliasState struct {
	// Alias ARN.
	Arn pulumi.StringPtrInput
	// Description of the alias.
	Description pulumi.StringPtrInput
	// Name of the alias.
	Name pulumi.StringPtrInput
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategyPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	// Description of the alias.
	Description *string `pulumi:"description"`
	// Name of the alias.
	Name *string `pulumi:"name"`
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategy `pulumi:"routingStrategy"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// Description of the alias.
	Description pulumi.StringPtrInput
	// Name of the alias.
	Name pulumi.StringPtrInput
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy AliasRoutingStrategyInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}

type AliasInput interface {
	pulumi.Input

	ToAliasOutput() AliasOutput
	ToAliasOutputWithContext(ctx context.Context) AliasOutput
}

func (Alias) ElementType() reflect.Type {
	return reflect.TypeOf((*Alias)(nil)).Elem()
}

func (i Alias) ToAliasOutput() AliasOutput {
	return i.ToAliasOutputWithContext(context.Background())
}

func (i Alias) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasOutput)
}

type AliasOutput struct {
	*pulumi.OutputState
}

func (AliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasOutput)(nil)).Elem()
}

func (o AliasOutput) ToAliasOutput() AliasOutput {
	return o
}

func (o AliasOutput) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AliasOutput{})
}
