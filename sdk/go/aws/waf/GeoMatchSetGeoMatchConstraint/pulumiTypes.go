// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package GeoMatchSetGeoMatchConstraint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GeoMatchSetGeoMatchConstraint struct {
	// The type of geographical area you want AWS WAF to search for. Currently Country is the only valid value.
	Type string `pulumi:"type"`
	// The country that you want AWS WAF to search for.
	// This is the two-letter country code, e.g. `US`, `CA`, `RU`, `CN`, etc.
	// See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_GeoMatchConstraint.html) for all supported values.
	Value string `pulumi:"value"`
}

type GeoMatchSetGeoMatchConstraintInput interface {
	pulumi.Input

	ToGeoMatchSetGeoMatchConstraintOutput() GeoMatchSetGeoMatchConstraintOutput
	ToGeoMatchSetGeoMatchConstraintOutputWithContext(context.Context) GeoMatchSetGeoMatchConstraintOutput
}

type GeoMatchSetGeoMatchConstraintArgs struct {
	// The type of geographical area you want AWS WAF to search for. Currently Country is the only valid value.
	Type pulumi.StringInput `pulumi:"type"`
	// The country that you want AWS WAF to search for.
	// This is the two-letter country code, e.g. `US`, `CA`, `RU`, `CN`, etc.
	// See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_GeoMatchConstraint.html) for all supported values.
	Value pulumi.StringInput `pulumi:"value"`
}

func (GeoMatchSetGeoMatchConstraintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoMatchSetGeoMatchConstraint)(nil)).Elem()
}

func (i GeoMatchSetGeoMatchConstraintArgs) ToGeoMatchSetGeoMatchConstraintOutput() GeoMatchSetGeoMatchConstraintOutput {
	return i.ToGeoMatchSetGeoMatchConstraintOutputWithContext(context.Background())
}

func (i GeoMatchSetGeoMatchConstraintArgs) ToGeoMatchSetGeoMatchConstraintOutputWithContext(ctx context.Context) GeoMatchSetGeoMatchConstraintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoMatchSetGeoMatchConstraintOutput)
}

type GeoMatchSetGeoMatchConstraintArrayInput interface {
	pulumi.Input

	ToGeoMatchSetGeoMatchConstraintArrayOutput() GeoMatchSetGeoMatchConstraintArrayOutput
	ToGeoMatchSetGeoMatchConstraintArrayOutputWithContext(context.Context) GeoMatchSetGeoMatchConstraintArrayOutput
}

type GeoMatchSetGeoMatchConstraintArray []GeoMatchSetGeoMatchConstraintInput

func (GeoMatchSetGeoMatchConstraintArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GeoMatchSetGeoMatchConstraint)(nil)).Elem()
}

func (i GeoMatchSetGeoMatchConstraintArray) ToGeoMatchSetGeoMatchConstraintArrayOutput() GeoMatchSetGeoMatchConstraintArrayOutput {
	return i.ToGeoMatchSetGeoMatchConstraintArrayOutputWithContext(context.Background())
}

func (i GeoMatchSetGeoMatchConstraintArray) ToGeoMatchSetGeoMatchConstraintArrayOutputWithContext(ctx context.Context) GeoMatchSetGeoMatchConstraintArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeoMatchSetGeoMatchConstraintArrayOutput)
}

type GeoMatchSetGeoMatchConstraintOutput struct { *pulumi.OutputState }

func (GeoMatchSetGeoMatchConstraintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoMatchSetGeoMatchConstraint)(nil)).Elem()
}

func (o GeoMatchSetGeoMatchConstraintOutput) ToGeoMatchSetGeoMatchConstraintOutput() GeoMatchSetGeoMatchConstraintOutput {
	return o
}

func (o GeoMatchSetGeoMatchConstraintOutput) ToGeoMatchSetGeoMatchConstraintOutputWithContext(ctx context.Context) GeoMatchSetGeoMatchConstraintOutput {
	return o
}

// The type of geographical area you want AWS WAF to search for. Currently Country is the only valid value.
func (o GeoMatchSetGeoMatchConstraintOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func (v GeoMatchSetGeoMatchConstraint) string { return v.Type }).(pulumi.StringOutput)
}

// The country that you want AWS WAF to search for.
// This is the two-letter country code, e.g. `US`, `CA`, `RU`, `CN`, etc.
// See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_GeoMatchConstraint.html) for all supported values.
func (o GeoMatchSetGeoMatchConstraintOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func (v GeoMatchSetGeoMatchConstraint) string { return v.Value }).(pulumi.StringOutput)
}

type GeoMatchSetGeoMatchConstraintArrayOutput struct { *pulumi.OutputState}

func (GeoMatchSetGeoMatchConstraintArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GeoMatchSetGeoMatchConstraint)(nil)).Elem()
}

func (o GeoMatchSetGeoMatchConstraintArrayOutput) ToGeoMatchSetGeoMatchConstraintArrayOutput() GeoMatchSetGeoMatchConstraintArrayOutput {
	return o
}

func (o GeoMatchSetGeoMatchConstraintArrayOutput) ToGeoMatchSetGeoMatchConstraintArrayOutputWithContext(ctx context.Context) GeoMatchSetGeoMatchConstraintArrayOutput {
	return o
}

func (o GeoMatchSetGeoMatchConstraintArrayOutput) Index(i pulumi.IntInput) GeoMatchSetGeoMatchConstraintOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GeoMatchSetGeoMatchConstraint {
		return vs[0].([]GeoMatchSetGeoMatchConstraint)[vs[1].(int)]
	}).(GeoMatchSetGeoMatchConstraintOutput)
}

func init() {
	pulumi.RegisterOutputType(GeoMatchSetGeoMatchConstraintOutput{})
	pulumi.RegisterOutputType(GeoMatchSetGeoMatchConstraintArrayOutput{})
}
