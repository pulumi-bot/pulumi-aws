// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates and manages an AWS IoT Thing Type.
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
// 		_, err := iot.NewThingType(ctx, "foo", nil)
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
// IOT Thing Types can be imported using the name, e.g.
//
// ```sh
//  $ pulumi import aws:iot/thingType:ThingType example example
// ```
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

type ThingTypeInput interface {
	pulumi.Input

	ToThingTypeOutput() ThingTypeOutput
	ToThingTypeOutputWithContext(ctx context.Context) ThingTypeOutput
}

func (*ThingType) ElementType() reflect.Type {
	return reflect.TypeOf((*ThingType)(nil))
}

func (i *ThingType) ToThingTypeOutput() ThingTypeOutput {
	return i.ToThingTypeOutputWithContext(context.Background())
}

func (i *ThingType) ToThingTypeOutputWithContext(ctx context.Context) ThingTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThingTypeOutput)
}

func (i *ThingType) ToThingTypePtrOutput() ThingTypePtrOutput {
	return i.ToThingTypePtrOutputWithContext(context.Background())
}

func (i *ThingType) ToThingTypePtrOutputWithContext(ctx context.Context) ThingTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThingTypePtrOutput)
}

type ThingTypePtrInput interface {
	pulumi.Input

	ToThingTypePtrOutput() ThingTypePtrOutput
	ToThingTypePtrOutputWithContext(ctx context.Context) ThingTypePtrOutput
}

type thingTypePtrType ThingTypeArgs

func (*thingTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ThingType)(nil))
}

func (i *thingTypePtrType) ToThingTypePtrOutput() ThingTypePtrOutput {
	return i.ToThingTypePtrOutputWithContext(context.Background())
}

func (i *thingTypePtrType) ToThingTypePtrOutputWithContext(ctx context.Context) ThingTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThingTypePtrOutput)
}

// ThingTypeArrayInput is an input type that accepts ThingTypeArray and ThingTypeArrayOutput values.
// You can construct a concrete instance of `ThingTypeArrayInput` via:
//
//          ThingTypeArray{ ThingTypeArgs{...} }
type ThingTypeArrayInput interface {
	pulumi.Input

	ToThingTypeArrayOutput() ThingTypeArrayOutput
	ToThingTypeArrayOutputWithContext(context.Context) ThingTypeArrayOutput
}

type ThingTypeArray []ThingTypeInput

func (ThingTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ThingType)(nil))
}

func (i ThingTypeArray) ToThingTypeArrayOutput() ThingTypeArrayOutput {
	return i.ToThingTypeArrayOutputWithContext(context.Background())
}

func (i ThingTypeArray) ToThingTypeArrayOutputWithContext(ctx context.Context) ThingTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThingTypeArrayOutput)
}

// ThingTypeMapInput is an input type that accepts ThingTypeMap and ThingTypeMapOutput values.
// You can construct a concrete instance of `ThingTypeMapInput` via:
//
//          ThingTypeMap{ "key": ThingTypeArgs{...} }
type ThingTypeMapInput interface {
	pulumi.Input

	ToThingTypeMapOutput() ThingTypeMapOutput
	ToThingTypeMapOutputWithContext(context.Context) ThingTypeMapOutput
}

type ThingTypeMap map[string]ThingTypeInput

func (ThingTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ThingType)(nil))
}

func (i ThingTypeMap) ToThingTypeMapOutput() ThingTypeMapOutput {
	return i.ToThingTypeMapOutputWithContext(context.Background())
}

func (i ThingTypeMap) ToThingTypeMapOutputWithContext(ctx context.Context) ThingTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThingTypeMapOutput)
}

type ThingTypeOutput struct {
	*pulumi.OutputState
}

func (ThingTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThingType)(nil))
}

func (o ThingTypeOutput) ToThingTypeOutput() ThingTypeOutput {
	return o
}

func (o ThingTypeOutput) ToThingTypeOutputWithContext(ctx context.Context) ThingTypeOutput {
	return o
}

func (o ThingTypeOutput) ToThingTypePtrOutput() ThingTypePtrOutput {
	return o.ToThingTypePtrOutputWithContext(context.Background())
}

func (o ThingTypeOutput) ToThingTypePtrOutputWithContext(ctx context.Context) ThingTypePtrOutput {
	return o.ApplyT(func(v ThingType) *ThingType {
		return &v
	}).(ThingTypePtrOutput)
}

type ThingTypePtrOutput struct {
	*pulumi.OutputState
}

func (ThingTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThingType)(nil))
}

func (o ThingTypePtrOutput) ToThingTypePtrOutput() ThingTypePtrOutput {
	return o
}

func (o ThingTypePtrOutput) ToThingTypePtrOutputWithContext(ctx context.Context) ThingTypePtrOutput {
	return o
}

type ThingTypeArrayOutput struct{ *pulumi.OutputState }

func (ThingTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThingType)(nil))
}

func (o ThingTypeArrayOutput) ToThingTypeArrayOutput() ThingTypeArrayOutput {
	return o
}

func (o ThingTypeArrayOutput) ToThingTypeArrayOutputWithContext(ctx context.Context) ThingTypeArrayOutput {
	return o
}

func (o ThingTypeArrayOutput) Index(i pulumi.IntInput) ThingTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThingType {
		return vs[0].([]ThingType)[vs[1].(int)]
	}).(ThingTypeOutput)
}

type ThingTypeMapOutput struct{ *pulumi.OutputState }

func (ThingTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ThingType)(nil))
}

func (o ThingTypeMapOutput) ToThingTypeMapOutput() ThingTypeMapOutput {
	return o
}

func (o ThingTypeMapOutput) ToThingTypeMapOutputWithContext(ctx context.Context) ThingTypeMapOutput {
	return o
}

func (o ThingTypeMapOutput) MapIndex(k pulumi.StringInput) ThingTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ThingType {
		return vs[0].(map[string]ThingType)[vs[1].(string)]
	}).(ThingTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(ThingTypeOutput{})
	pulumi.RegisterOutputType(ThingTypePtrOutput{})
	pulumi.RegisterOutputType(ThingTypeArrayOutput{})
	pulumi.RegisterOutputType(ThingTypeMapOutput{})
}
