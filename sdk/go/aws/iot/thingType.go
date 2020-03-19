// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iot

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates and manages an AWS IoT Thing Type.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_thing_type.html.markdown.
type ThingType struct {
	pulumi.CustomResourceState

	// The ARN of the created AWS IoT Thing Type.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether the thing type is deprecated. If true, no new things could be associated with this type.
	Deprecated pulumi.BoolPtrOutput `pulumi:"deprecated"`
	// The name of the thing type.
	Name pulumi.StringOutput `pulumi:"name"`
	// , Configuration block that can contain the following properties of the thing type:
	Properties ThingTypePropertiesPtrOutput `pulumi:"properties"`
}

// NewThingType registers a new resource with the given unique name, arguments, and options.
func NewThingType(ctx *pulumi.Context,
	name string, args *ThingTypeArgs, opts ...pulumi.ResourceOption) (*ThingType, error) {
	if args == nil {
		args = &ThingTypeArgs{}
	}
	var resource ThingType
	err := ctx.RegisterResource("aws:iot/thingType:ThingType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetThingType gets an existing ThingType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThingType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThingTypeState, opts ...pulumi.ResourceOption) (*ThingType, error) {
	var resource ThingType
	err := ctx.ReadResource("aws:iot/thingType:ThingType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ThingType resources.
type thingTypeState struct {
	// The ARN of the created AWS IoT Thing Type.
	Arn *string `pulumi:"arn"`
	// Whether the thing type is deprecated. If true, no new things could be associated with this type.
	Deprecated *bool `pulumi:"deprecated"`
	// The name of the thing type.
	Name *string `pulumi:"name"`
	// , Configuration block that can contain the following properties of the thing type:
	Properties *ThingTypeProperties `pulumi:"properties"`
}

type ThingTypeState struct {
	// The ARN of the created AWS IoT Thing Type.
	Arn pulumi.StringPtrInput
	// Whether the thing type is deprecated. If true, no new things could be associated with this type.
	Deprecated pulumi.BoolPtrInput
	// The name of the thing type.
	Name pulumi.StringPtrInput
	// , Configuration block that can contain the following properties of the thing type:
	Properties ThingTypePropertiesPtrInput
}

func (ThingTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*thingTypeState)(nil)).Elem()
}

type thingTypeArgs struct {
	// Whether the thing type is deprecated. If true, no new things could be associated with this type.
	Deprecated *bool `pulumi:"deprecated"`
	// The name of the thing type.
	Name *string `pulumi:"name"`
	// , Configuration block that can contain the following properties of the thing type:
	Properties *ThingTypeProperties `pulumi:"properties"`
}

// The set of arguments for constructing a ThingType resource.
type ThingTypeArgs struct {
	// Whether the thing type is deprecated. If true, no new things could be associated with this type.
	Deprecated pulumi.BoolPtrInput
	// The name of the thing type.
	Name pulumi.StringPtrInput
	// , Configuration block that can contain the following properties of the thing type:
	Properties ThingTypePropertiesPtrInput
}

func (ThingTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*thingTypeArgs)(nil)).Elem()
}

