// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getAmiProductCode

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetAmiProductCode struct {
	ProductCodeId string `pulumi:"productCodeId"`
	ProductCodeType string `pulumi:"productCodeType"`
}

type GetAmiProductCodeInput interface {
	pulumi.Input

	ToGetAmiProductCodeOutput() GetAmiProductCodeOutput
	ToGetAmiProductCodeOutputWithContext(context.Context) GetAmiProductCodeOutput
}

type GetAmiProductCodeArgs struct {
	ProductCodeId pulumi.StringInput `pulumi:"productCodeId"`
	ProductCodeType pulumi.StringInput `pulumi:"productCodeType"`
}

func (GetAmiProductCodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAmiProductCode)(nil)).Elem()
}

func (i GetAmiProductCodeArgs) ToGetAmiProductCodeOutput() GetAmiProductCodeOutput {
	return i.ToGetAmiProductCodeOutputWithContext(context.Background())
}

func (i GetAmiProductCodeArgs) ToGetAmiProductCodeOutputWithContext(ctx context.Context) GetAmiProductCodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAmiProductCodeOutput)
}

type GetAmiProductCodeArrayInput interface {
	pulumi.Input

	ToGetAmiProductCodeArrayOutput() GetAmiProductCodeArrayOutput
	ToGetAmiProductCodeArrayOutputWithContext(context.Context) GetAmiProductCodeArrayOutput
}

type GetAmiProductCodeArray []GetAmiProductCodeInput

func (GetAmiProductCodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAmiProductCode)(nil)).Elem()
}

func (i GetAmiProductCodeArray) ToGetAmiProductCodeArrayOutput() GetAmiProductCodeArrayOutput {
	return i.ToGetAmiProductCodeArrayOutputWithContext(context.Background())
}

func (i GetAmiProductCodeArray) ToGetAmiProductCodeArrayOutputWithContext(ctx context.Context) GetAmiProductCodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAmiProductCodeArrayOutput)
}

type GetAmiProductCodeOutput struct { *pulumi.OutputState }

func (GetAmiProductCodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAmiProductCode)(nil)).Elem()
}

func (o GetAmiProductCodeOutput) ToGetAmiProductCodeOutput() GetAmiProductCodeOutput {
	return o
}

func (o GetAmiProductCodeOutput) ToGetAmiProductCodeOutputWithContext(ctx context.Context) GetAmiProductCodeOutput {
	return o
}

func (o GetAmiProductCodeOutput) ProductCodeId() pulumi.StringOutput {
	return o.ApplyT(func (v GetAmiProductCode) string { return v.ProductCodeId }).(pulumi.StringOutput)
}

func (o GetAmiProductCodeOutput) ProductCodeType() pulumi.StringOutput {
	return o.ApplyT(func (v GetAmiProductCode) string { return v.ProductCodeType }).(pulumi.StringOutput)
}

type GetAmiProductCodeArrayOutput struct { *pulumi.OutputState}

func (GetAmiProductCodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAmiProductCode)(nil)).Elem()
}

func (o GetAmiProductCodeArrayOutput) ToGetAmiProductCodeArrayOutput() GetAmiProductCodeArrayOutput {
	return o
}

func (o GetAmiProductCodeArrayOutput) ToGetAmiProductCodeArrayOutputWithContext(ctx context.Context) GetAmiProductCodeArrayOutput {
	return o
}

func (o GetAmiProductCodeArrayOutput) Index(i pulumi.IntInput) GetAmiProductCodeOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetAmiProductCode {
		return vs[0].([]GetAmiProductCode)[vs[1].(int)]
	}).(GetAmiProductCodeOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAmiProductCodeOutput{})
	pulumi.RegisterOutputType(GetAmiProductCodeArrayOutput{})
}
