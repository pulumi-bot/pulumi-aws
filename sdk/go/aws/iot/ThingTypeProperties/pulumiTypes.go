// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ThingTypeProperties

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ThingTypeProperties struct {
	// The description of the thing type.
	Description *string `pulumi:"description"`
	// A list of searchable thing attribute names.
	SearchableAttributes []string `pulumi:"searchableAttributes"`
}

type ThingTypePropertiesInput interface {
	pulumi.Input

	ToThingTypePropertiesOutput() ThingTypePropertiesOutput
	ToThingTypePropertiesOutputWithContext(context.Context) ThingTypePropertiesOutput
}

type ThingTypePropertiesArgs struct {
	// The description of the thing type.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// A list of searchable thing attribute names.
	SearchableAttributes pulumi.StringArrayInput `pulumi:"searchableAttributes"`
}

func (ThingTypePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThingTypeProperties)(nil)).Elem()
}

func (i ThingTypePropertiesArgs) ToThingTypePropertiesOutput() ThingTypePropertiesOutput {
	return i.ToThingTypePropertiesOutputWithContext(context.Background())
}

func (i ThingTypePropertiesArgs) ToThingTypePropertiesOutputWithContext(ctx context.Context) ThingTypePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThingTypePropertiesOutput)
}

func (i ThingTypePropertiesArgs) ToThingTypePropertiesPtrOutput() ThingTypePropertiesPtrOutput {
	return i.ToThingTypePropertiesPtrOutputWithContext(context.Background())
}

func (i ThingTypePropertiesArgs) ToThingTypePropertiesPtrOutputWithContext(ctx context.Context) ThingTypePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThingTypePropertiesOutput).ToThingTypePropertiesPtrOutputWithContext(ctx)
}

type ThingTypePropertiesPtrInput interface {
	pulumi.Input

	ToThingTypePropertiesPtrOutput() ThingTypePropertiesPtrOutput
	ToThingTypePropertiesPtrOutputWithContext(context.Context) ThingTypePropertiesPtrOutput
}

type thingTypePropertiesPtrType ThingTypePropertiesArgs

func ThingTypePropertiesPtr(v *ThingTypePropertiesArgs) ThingTypePropertiesPtrInput {	return (*thingTypePropertiesPtrType)(v)
}

func (*thingTypePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ThingTypeProperties)(nil)).Elem()
}

func (i *thingTypePropertiesPtrType) ToThingTypePropertiesPtrOutput() ThingTypePropertiesPtrOutput {
	return i.ToThingTypePropertiesPtrOutputWithContext(context.Background())
}

func (i *thingTypePropertiesPtrType) ToThingTypePropertiesPtrOutputWithContext(ctx context.Context) ThingTypePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThingTypePropertiesPtrOutput)
}

type ThingTypePropertiesOutput struct { *pulumi.OutputState }

func (ThingTypePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThingTypeProperties)(nil)).Elem()
}

func (o ThingTypePropertiesOutput) ToThingTypePropertiesOutput() ThingTypePropertiesOutput {
	return o
}

func (o ThingTypePropertiesOutput) ToThingTypePropertiesOutputWithContext(ctx context.Context) ThingTypePropertiesOutput {
	return o
}

func (o ThingTypePropertiesOutput) ToThingTypePropertiesPtrOutput() ThingTypePropertiesPtrOutput {
	return o.ToThingTypePropertiesPtrOutputWithContext(context.Background())
}

func (o ThingTypePropertiesOutput) ToThingTypePropertiesPtrOutputWithContext(ctx context.Context) ThingTypePropertiesPtrOutput {
	return o.ApplyT(func(v ThingTypeProperties) *ThingTypeProperties {
		return &v
	}).(ThingTypePropertiesPtrOutput)
}
// The description of the thing type.
func (o ThingTypePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ThingTypeProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A list of searchable thing attribute names.
func (o ThingTypePropertiesOutput) SearchableAttributes() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ThingTypeProperties) []string { return v.SearchableAttributes }).(pulumi.StringArrayOutput)
}

type ThingTypePropertiesPtrOutput struct { *pulumi.OutputState}

func (ThingTypePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThingTypeProperties)(nil)).Elem()
}

func (o ThingTypePropertiesPtrOutput) ToThingTypePropertiesPtrOutput() ThingTypePropertiesPtrOutput {
	return o
}

func (o ThingTypePropertiesPtrOutput) ToThingTypePropertiesPtrOutputWithContext(ctx context.Context) ThingTypePropertiesPtrOutput {
	return o
}

func (o ThingTypePropertiesPtrOutput) Elem() ThingTypePropertiesOutput {
	return o.ApplyT(func (v *ThingTypeProperties) ThingTypeProperties { return *v }).(ThingTypePropertiesOutput)
}

// The description of the thing type.
func (o ThingTypePropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ThingTypeProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A list of searchable thing attribute names.
func (o ThingTypePropertiesPtrOutput) SearchableAttributes() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ThingTypeProperties) []string { return v.SearchableAttributes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ThingTypePropertiesOutput{})
	pulumi.RegisterOutputType(ThingTypePropertiesPtrOutput{})
}
