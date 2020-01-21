// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ReceiptRuleAddHeaderAction

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ReceiptRuleAddHeaderAction struct {
	// The name of the header to add
	HeaderName string `pulumi:"headerName"`
	// The value of the header to add
	HeaderValue string `pulumi:"headerValue"`
	// The position of the action in the receipt rule
	Position int `pulumi:"position"`
}

type ReceiptRuleAddHeaderActionInput interface {
	pulumi.Input

	ToReceiptRuleAddHeaderActionOutput() ReceiptRuleAddHeaderActionOutput
	ToReceiptRuleAddHeaderActionOutputWithContext(context.Context) ReceiptRuleAddHeaderActionOutput
}

type ReceiptRuleAddHeaderActionArgs struct {
	// The name of the header to add
	HeaderName pulumi.StringInput `pulumi:"headerName"`
	// The value of the header to add
	HeaderValue pulumi.StringInput `pulumi:"headerValue"`
	// The position of the action in the receipt rule
	Position pulumi.IntInput `pulumi:"position"`
}

func (ReceiptRuleAddHeaderActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReceiptRuleAddHeaderAction)(nil)).Elem()
}

func (i ReceiptRuleAddHeaderActionArgs) ToReceiptRuleAddHeaderActionOutput() ReceiptRuleAddHeaderActionOutput {
	return i.ToReceiptRuleAddHeaderActionOutputWithContext(context.Background())
}

func (i ReceiptRuleAddHeaderActionArgs) ToReceiptRuleAddHeaderActionOutputWithContext(ctx context.Context) ReceiptRuleAddHeaderActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptRuleAddHeaderActionOutput)
}

type ReceiptRuleAddHeaderActionArrayInput interface {
	pulumi.Input

	ToReceiptRuleAddHeaderActionArrayOutput() ReceiptRuleAddHeaderActionArrayOutput
	ToReceiptRuleAddHeaderActionArrayOutputWithContext(context.Context) ReceiptRuleAddHeaderActionArrayOutput
}

type ReceiptRuleAddHeaderActionArray []ReceiptRuleAddHeaderActionInput

func (ReceiptRuleAddHeaderActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReceiptRuleAddHeaderAction)(nil)).Elem()
}

func (i ReceiptRuleAddHeaderActionArray) ToReceiptRuleAddHeaderActionArrayOutput() ReceiptRuleAddHeaderActionArrayOutput {
	return i.ToReceiptRuleAddHeaderActionArrayOutputWithContext(context.Background())
}

func (i ReceiptRuleAddHeaderActionArray) ToReceiptRuleAddHeaderActionArrayOutputWithContext(ctx context.Context) ReceiptRuleAddHeaderActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptRuleAddHeaderActionArrayOutput)
}

type ReceiptRuleAddHeaderActionOutput struct { *pulumi.OutputState }

func (ReceiptRuleAddHeaderActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReceiptRuleAddHeaderAction)(nil)).Elem()
}

func (o ReceiptRuleAddHeaderActionOutput) ToReceiptRuleAddHeaderActionOutput() ReceiptRuleAddHeaderActionOutput {
	return o
}

func (o ReceiptRuleAddHeaderActionOutput) ToReceiptRuleAddHeaderActionOutputWithContext(ctx context.Context) ReceiptRuleAddHeaderActionOutput {
	return o
}

// The name of the header to add
func (o ReceiptRuleAddHeaderActionOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func (v ReceiptRuleAddHeaderAction) string { return v.HeaderName }).(pulumi.StringOutput)
}

// The value of the header to add
func (o ReceiptRuleAddHeaderActionOutput) HeaderValue() pulumi.StringOutput {
	return o.ApplyT(func (v ReceiptRuleAddHeaderAction) string { return v.HeaderValue }).(pulumi.StringOutput)
}

// The position of the action in the receipt rule
func (o ReceiptRuleAddHeaderActionOutput) Position() pulumi.IntOutput {
	return o.ApplyT(func (v ReceiptRuleAddHeaderAction) int { return v.Position }).(pulumi.IntOutput)
}

type ReceiptRuleAddHeaderActionArrayOutput struct { *pulumi.OutputState}

func (ReceiptRuleAddHeaderActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReceiptRuleAddHeaderAction)(nil)).Elem()
}

func (o ReceiptRuleAddHeaderActionArrayOutput) ToReceiptRuleAddHeaderActionArrayOutput() ReceiptRuleAddHeaderActionArrayOutput {
	return o
}

func (o ReceiptRuleAddHeaderActionArrayOutput) ToReceiptRuleAddHeaderActionArrayOutputWithContext(ctx context.Context) ReceiptRuleAddHeaderActionArrayOutput {
	return o
}

func (o ReceiptRuleAddHeaderActionArrayOutput) Index(i pulumi.IntInput) ReceiptRuleAddHeaderActionOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ReceiptRuleAddHeaderAction {
		return vs[0].([]ReceiptRuleAddHeaderAction)[vs[1].(int)]
	}).(ReceiptRuleAddHeaderActionOutput)
}

func init() {
	pulumi.RegisterOutputType(ReceiptRuleAddHeaderActionOutput{})
	pulumi.RegisterOutputType(ReceiptRuleAddHeaderActionArrayOutput{})
}
