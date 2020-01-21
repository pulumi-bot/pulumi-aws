// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package TriggerPredicate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/glue/TriggerPredicateCondition"
)

type TriggerPredicate struct {
	// A list of the conditions that determine when the trigger will fire. Defined below.
	Conditions []glueTriggerPredicateCondition.TriggerPredicateCondition `pulumi:"conditions"`
	// How to handle multiple conditions. Defaults to `AND`. Valid values are `AND` or `ANY`.
	Logical *string `pulumi:"logical"`
}

type TriggerPredicateInput interface {
	pulumi.Input

	ToTriggerPredicateOutput() TriggerPredicateOutput
	ToTriggerPredicateOutputWithContext(context.Context) TriggerPredicateOutput
}

type TriggerPredicateArgs struct {
	// A list of the conditions that determine when the trigger will fire. Defined below.
	Conditions glueTriggerPredicateCondition.TriggerPredicateConditionArrayInput `pulumi:"conditions"`
	// How to handle multiple conditions. Defaults to `AND`. Valid values are `AND` or `ANY`.
	Logical pulumi.StringPtrInput `pulumi:"logical"`
}

func (TriggerPredicateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerPredicate)(nil)).Elem()
}

func (i TriggerPredicateArgs) ToTriggerPredicateOutput() TriggerPredicateOutput {
	return i.ToTriggerPredicateOutputWithContext(context.Background())
}

func (i TriggerPredicateArgs) ToTriggerPredicateOutputWithContext(ctx context.Context) TriggerPredicateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerPredicateOutput)
}

func (i TriggerPredicateArgs) ToTriggerPredicatePtrOutput() TriggerPredicatePtrOutput {
	return i.ToTriggerPredicatePtrOutputWithContext(context.Background())
}

func (i TriggerPredicateArgs) ToTriggerPredicatePtrOutputWithContext(ctx context.Context) TriggerPredicatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerPredicateOutput).ToTriggerPredicatePtrOutputWithContext(ctx)
}

type TriggerPredicatePtrInput interface {
	pulumi.Input

	ToTriggerPredicatePtrOutput() TriggerPredicatePtrOutput
	ToTriggerPredicatePtrOutputWithContext(context.Context) TriggerPredicatePtrOutput
}

type triggerPredicatePtrType TriggerPredicateArgs

func TriggerPredicatePtr(v *TriggerPredicateArgs) TriggerPredicatePtrInput {	return (*triggerPredicatePtrType)(v)
}

func (*triggerPredicatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerPredicate)(nil)).Elem()
}

func (i *triggerPredicatePtrType) ToTriggerPredicatePtrOutput() TriggerPredicatePtrOutput {
	return i.ToTriggerPredicatePtrOutputWithContext(context.Background())
}

func (i *triggerPredicatePtrType) ToTriggerPredicatePtrOutputWithContext(ctx context.Context) TriggerPredicatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerPredicatePtrOutput)
}

type TriggerPredicateOutput struct { *pulumi.OutputState }

func (TriggerPredicateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerPredicate)(nil)).Elem()
}

func (o TriggerPredicateOutput) ToTriggerPredicateOutput() TriggerPredicateOutput {
	return o
}

func (o TriggerPredicateOutput) ToTriggerPredicateOutputWithContext(ctx context.Context) TriggerPredicateOutput {
	return o
}

func (o TriggerPredicateOutput) ToTriggerPredicatePtrOutput() TriggerPredicatePtrOutput {
	return o.ToTriggerPredicatePtrOutputWithContext(context.Background())
}

func (o TriggerPredicateOutput) ToTriggerPredicatePtrOutputWithContext(ctx context.Context) TriggerPredicatePtrOutput {
	return o.ApplyT(func(v TriggerPredicate) *TriggerPredicate {
		return &v
	}).(TriggerPredicatePtrOutput)
}
// A list of the conditions that determine when the trigger will fire. Defined below.
func (o TriggerPredicateOutput) Conditions() glueTriggerPredicateCondition.TriggerPredicateConditionArrayOutput {
	return o.ApplyT(func (v TriggerPredicate) []glueTriggerPredicateCondition.TriggerPredicateCondition { return v.Conditions }).(glueTriggerPredicateCondition.TriggerPredicateConditionArrayOutput)
}

// How to handle multiple conditions. Defaults to `AND`. Valid values are `AND` or `ANY`.
func (o TriggerPredicateOutput) Logical() pulumi.StringPtrOutput {
	return o.ApplyT(func (v TriggerPredicate) *string { return v.Logical }).(pulumi.StringPtrOutput)
}

type TriggerPredicatePtrOutput struct { *pulumi.OutputState}

func (TriggerPredicatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TriggerPredicate)(nil)).Elem()
}

func (o TriggerPredicatePtrOutput) ToTriggerPredicatePtrOutput() TriggerPredicatePtrOutput {
	return o
}

func (o TriggerPredicatePtrOutput) ToTriggerPredicatePtrOutputWithContext(ctx context.Context) TriggerPredicatePtrOutput {
	return o
}

func (o TriggerPredicatePtrOutput) Elem() TriggerPredicateOutput {
	return o.ApplyT(func (v *TriggerPredicate) TriggerPredicate { return *v }).(TriggerPredicateOutput)
}

// A list of the conditions that determine when the trigger will fire. Defined below.
func (o TriggerPredicatePtrOutput) Conditions() glueTriggerPredicateCondition.TriggerPredicateConditionArrayOutput {
	return o.ApplyT(func (v TriggerPredicate) []glueTriggerPredicateCondition.TriggerPredicateCondition { return v.Conditions }).(glueTriggerPredicateCondition.TriggerPredicateConditionArrayOutput)
}

// How to handle multiple conditions. Defaults to `AND`. Valid values are `AND` or `ANY`.
func (o TriggerPredicatePtrOutput) Logical() pulumi.StringPtrOutput {
	return o.ApplyT(func (v TriggerPredicate) *string { return v.Logical }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(TriggerPredicateOutput{})
	pulumi.RegisterOutputType(TriggerPredicatePtrOutput{})
}
