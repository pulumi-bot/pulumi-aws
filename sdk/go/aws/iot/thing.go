// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates and manages an AWS IoT Thing.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iot"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iot.NewThing(ctx, "example", &iot.ThingArgs{
// 			Attributes: pulumi.StringMap{
// 				"First": pulumi.String("examplevalue"),
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
// IOT Things can be imported using the name, e.g.
//
// ```sh
//  $ pulumi import aws:iot/thing:Thing example example
// ```
type Thing struct {
	pulumi.CustomResourceState

	// The ARN of the thing.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Map of attributes of the thing.
	Attributes pulumi.StringMapOutput `pulumi:"attributes"`
	// The default client ID.
	DefaultClientId pulumi.StringOutput `pulumi:"defaultClientId"`
	// The name of the thing.
	Name pulumi.StringOutput `pulumi:"name"`
	// The thing type name.
	ThingTypeName pulumi.StringPtrOutput `pulumi:"thingTypeName"`
	// The current version of the thing record in the registry.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewThing registers a new resource with the given unique name, arguments, and options.
func NewThing(ctx *pulumi.Context,
	name string, args *ThingArgs, opts ...pulumi.ResourceOption) (*Thing, error) {
	if args == nil {
		args = &ThingArgs{}
	}

	var resource Thing
	err := ctx.RegisterResource("aws:iot/thing:Thing", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetThing gets an existing Thing resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThing(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThingState, opts ...pulumi.ResourceOption) (*Thing, error) {
	var resource Thing
	err := ctx.ReadResource("aws:iot/thing:Thing", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Thing resources.
type thingState struct {
	// The ARN of the thing.
	Arn *string `pulumi:"arn"`
	// Map of attributes of the thing.
	Attributes map[string]string `pulumi:"attributes"`
	// The default client ID.
	DefaultClientId *string `pulumi:"defaultClientId"`
	// The name of the thing.
	Name *string `pulumi:"name"`
	// The thing type name.
	ThingTypeName *string `pulumi:"thingTypeName"`
	// The current version of the thing record in the registry.
	Version *int `pulumi:"version"`
}

type ThingState struct {
	// The ARN of the thing.
	Arn pulumi.StringPtrInput
	// Map of attributes of the thing.
	Attributes pulumi.StringMapInput
	// The default client ID.
	DefaultClientId pulumi.StringPtrInput
	// The name of the thing.
	Name pulumi.StringPtrInput
	// The thing type name.
	ThingTypeName pulumi.StringPtrInput
	// The current version of the thing record in the registry.
	Version pulumi.IntPtrInput
}

func (ThingState) ElementType() reflect.Type {
	return reflect.TypeOf((*thingState)(nil)).Elem()
}

type thingArgs struct {
	// Map of attributes of the thing.
	Attributes map[string]string `pulumi:"attributes"`
	// The name of the thing.
	Name *string `pulumi:"name"`
	// The thing type name.
	ThingTypeName *string `pulumi:"thingTypeName"`
}

// The set of arguments for constructing a Thing resource.
type ThingArgs struct {
	// Map of attributes of the thing.
	Attributes pulumi.StringMapInput
	// The name of the thing.
	Name pulumi.StringPtrInput
	// The thing type name.
	ThingTypeName pulumi.StringPtrInput
}

func (ThingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*thingArgs)(nil)).Elem()
}

type ThingInput interface {
	pulumi.Input

	ToThingOutput() ThingOutput
	ToThingOutputWithContext(ctx context.Context) ThingOutput
}

type ThingPtrInput interface {
	pulumi.Input

	ToThingPtrOutput() ThingPtrOutput
	ToThingPtrOutputWithContext(ctx context.Context) ThingPtrOutput
}

func (Thing) ElementType() reflect.Type {
	return reflect.TypeOf((*Thing)(nil)).Elem()
}

func (i Thing) ToThingOutput() ThingOutput {
	return i.ToThingOutputWithContext(context.Background())
}

func (i Thing) ToThingOutputWithContext(ctx context.Context) ThingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThingOutput)
}

func (i Thing) ToThingPtrOutput() ThingPtrOutput {
	return i.ToThingPtrOutputWithContext(context.Background())
}

func (i Thing) ToThingPtrOutputWithContext(ctx context.Context) ThingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThingPtrOutput)
}

type ThingOutput struct {
	*pulumi.OutputState
}

func (ThingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThingOutput)(nil)).Elem()
}

func (o ThingOutput) ToThingOutput() ThingOutput {
	return o
}

func (o ThingOutput) ToThingOutputWithContext(ctx context.Context) ThingOutput {
	return o
}

type ThingPtrOutput struct {
	*pulumi.OutputState
}

func (ThingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Thing)(nil)).Elem()
}

func (o ThingPtrOutput) ToThingPtrOutput() ThingPtrOutput {
	return o
}

func (o ThingPtrOutput) ToThingPtrOutputWithContext(ctx context.Context) ThingPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ThingOutput{})
	pulumi.RegisterOutputType(ThingPtrOutput{})
}
