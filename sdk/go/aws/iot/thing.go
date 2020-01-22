// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iot

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates and manages an AWS IoT Thing.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_thing.html.markdown.
type Thing struct {
	pulumi.CustomResourceState

	// The ARN of the thing.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Map of attributes of the thing.
	Attributes pulumi.MapOutput `pulumi:"attributes"`
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
	Attributes map[string]interface{} `pulumi:"attributes"`
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
	Attributes pulumi.MapInput
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
	Attributes map[string]interface{} `pulumi:"attributes"`
	// The name of the thing.
	Name *string `pulumi:"name"`
	// The thing type name.
	ThingTypeName *string `pulumi:"thingTypeName"`
}

// The set of arguments for constructing a Thing resource.
type ThingArgs struct {
	// Map of attributes of the thing.
	Attributes pulumi.MapInput
	// The name of the thing.
	Name pulumi.StringPtrInput
	// The thing type name.
	ThingTypeName pulumi.StringPtrInput
}

func (ThingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*thingArgs)(nil)).Elem()
}

