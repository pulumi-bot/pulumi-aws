// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package CatalogTableStorageDescriptorSkewedInfo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type CatalogTableStorageDescriptorSkewedInfo struct {
	// A list of names of columns that contain skewed values.
	SkewedColumnNames []string `pulumi:"skewedColumnNames"`
	// A list of values that appear so frequently as to be considered skewed.
	SkewedColumnValueLocationMaps map[string]string `pulumi:"skewedColumnValueLocationMaps"`
	// A mapping of skewed values to the columns that contain them.
	SkewedColumnValues []string `pulumi:"skewedColumnValues"`
}

type CatalogTableStorageDescriptorSkewedInfoInput interface {
	pulumi.Input

	ToCatalogTableStorageDescriptorSkewedInfoOutput() CatalogTableStorageDescriptorSkewedInfoOutput
	ToCatalogTableStorageDescriptorSkewedInfoOutputWithContext(context.Context) CatalogTableStorageDescriptorSkewedInfoOutput
}

type CatalogTableStorageDescriptorSkewedInfoArgs struct {
	// A list of names of columns that contain skewed values.
	SkewedColumnNames pulumi.StringArrayInput `pulumi:"skewedColumnNames"`
	// A list of values that appear so frequently as to be considered skewed.
	SkewedColumnValueLocationMaps pulumi.StringMapInput `pulumi:"skewedColumnValueLocationMaps"`
	// A mapping of skewed values to the columns that contain them.
	SkewedColumnValues pulumi.StringArrayInput `pulumi:"skewedColumnValues"`
}

func (CatalogTableStorageDescriptorSkewedInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CatalogTableStorageDescriptorSkewedInfo)(nil)).Elem()
}

func (i CatalogTableStorageDescriptorSkewedInfoArgs) ToCatalogTableStorageDescriptorSkewedInfoOutput() CatalogTableStorageDescriptorSkewedInfoOutput {
	return i.ToCatalogTableStorageDescriptorSkewedInfoOutputWithContext(context.Background())
}

func (i CatalogTableStorageDescriptorSkewedInfoArgs) ToCatalogTableStorageDescriptorSkewedInfoOutputWithContext(ctx context.Context) CatalogTableStorageDescriptorSkewedInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogTableStorageDescriptorSkewedInfoOutput)
}

func (i CatalogTableStorageDescriptorSkewedInfoArgs) ToCatalogTableStorageDescriptorSkewedInfoPtrOutput() CatalogTableStorageDescriptorSkewedInfoPtrOutput {
	return i.ToCatalogTableStorageDescriptorSkewedInfoPtrOutputWithContext(context.Background())
}

func (i CatalogTableStorageDescriptorSkewedInfoArgs) ToCatalogTableStorageDescriptorSkewedInfoPtrOutputWithContext(ctx context.Context) CatalogTableStorageDescriptorSkewedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogTableStorageDescriptorSkewedInfoOutput).ToCatalogTableStorageDescriptorSkewedInfoPtrOutputWithContext(ctx)
}

type CatalogTableStorageDescriptorSkewedInfoPtrInput interface {
	pulumi.Input

	ToCatalogTableStorageDescriptorSkewedInfoPtrOutput() CatalogTableStorageDescriptorSkewedInfoPtrOutput
	ToCatalogTableStorageDescriptorSkewedInfoPtrOutputWithContext(context.Context) CatalogTableStorageDescriptorSkewedInfoPtrOutput
}

type catalogTableStorageDescriptorSkewedInfoPtrType CatalogTableStorageDescriptorSkewedInfoArgs

func CatalogTableStorageDescriptorSkewedInfoPtr(v *CatalogTableStorageDescriptorSkewedInfoArgs) CatalogTableStorageDescriptorSkewedInfoPtrInput {	return (*catalogTableStorageDescriptorSkewedInfoPtrType)(v)
}

func (*catalogTableStorageDescriptorSkewedInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CatalogTableStorageDescriptorSkewedInfo)(nil)).Elem()
}

func (i *catalogTableStorageDescriptorSkewedInfoPtrType) ToCatalogTableStorageDescriptorSkewedInfoPtrOutput() CatalogTableStorageDescriptorSkewedInfoPtrOutput {
	return i.ToCatalogTableStorageDescriptorSkewedInfoPtrOutputWithContext(context.Background())
}

func (i *catalogTableStorageDescriptorSkewedInfoPtrType) ToCatalogTableStorageDescriptorSkewedInfoPtrOutputWithContext(ctx context.Context) CatalogTableStorageDescriptorSkewedInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogTableStorageDescriptorSkewedInfoPtrOutput)
}

type CatalogTableStorageDescriptorSkewedInfoOutput struct { *pulumi.OutputState }

func (CatalogTableStorageDescriptorSkewedInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CatalogTableStorageDescriptorSkewedInfo)(nil)).Elem()
}

func (o CatalogTableStorageDescriptorSkewedInfoOutput) ToCatalogTableStorageDescriptorSkewedInfoOutput() CatalogTableStorageDescriptorSkewedInfoOutput {
	return o
}

func (o CatalogTableStorageDescriptorSkewedInfoOutput) ToCatalogTableStorageDescriptorSkewedInfoOutputWithContext(ctx context.Context) CatalogTableStorageDescriptorSkewedInfoOutput {
	return o
}

func (o CatalogTableStorageDescriptorSkewedInfoOutput) ToCatalogTableStorageDescriptorSkewedInfoPtrOutput() CatalogTableStorageDescriptorSkewedInfoPtrOutput {
	return o.ToCatalogTableStorageDescriptorSkewedInfoPtrOutputWithContext(context.Background())
}

func (o CatalogTableStorageDescriptorSkewedInfoOutput) ToCatalogTableStorageDescriptorSkewedInfoPtrOutputWithContext(ctx context.Context) CatalogTableStorageDescriptorSkewedInfoPtrOutput {
	return o.ApplyT(func(v CatalogTableStorageDescriptorSkewedInfo) *CatalogTableStorageDescriptorSkewedInfo {
		return &v
	}).(CatalogTableStorageDescriptorSkewedInfoPtrOutput)
}
// A list of names of columns that contain skewed values.
func (o CatalogTableStorageDescriptorSkewedInfoOutput) SkewedColumnNames() pulumi.StringArrayOutput {
	return o.ApplyT(func (v CatalogTableStorageDescriptorSkewedInfo) []string { return v.SkewedColumnNames }).(pulumi.StringArrayOutput)
}

// A list of values that appear so frequently as to be considered skewed.
func (o CatalogTableStorageDescriptorSkewedInfoOutput) SkewedColumnValueLocationMaps() pulumi.StringMapOutput {
	return o.ApplyT(func (v CatalogTableStorageDescriptorSkewedInfo) map[string]string { return v.SkewedColumnValueLocationMaps }).(pulumi.StringMapOutput)
}

// A mapping of skewed values to the columns that contain them.
func (o CatalogTableStorageDescriptorSkewedInfoOutput) SkewedColumnValues() pulumi.StringArrayOutput {
	return o.ApplyT(func (v CatalogTableStorageDescriptorSkewedInfo) []string { return v.SkewedColumnValues }).(pulumi.StringArrayOutput)
}

type CatalogTableStorageDescriptorSkewedInfoPtrOutput struct { *pulumi.OutputState}

func (CatalogTableStorageDescriptorSkewedInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CatalogTableStorageDescriptorSkewedInfo)(nil)).Elem()
}

func (o CatalogTableStorageDescriptorSkewedInfoPtrOutput) ToCatalogTableStorageDescriptorSkewedInfoPtrOutput() CatalogTableStorageDescriptorSkewedInfoPtrOutput {
	return o
}

func (o CatalogTableStorageDescriptorSkewedInfoPtrOutput) ToCatalogTableStorageDescriptorSkewedInfoPtrOutputWithContext(ctx context.Context) CatalogTableStorageDescriptorSkewedInfoPtrOutput {
	return o
}

func (o CatalogTableStorageDescriptorSkewedInfoPtrOutput) Elem() CatalogTableStorageDescriptorSkewedInfoOutput {
	return o.ApplyT(func (v *CatalogTableStorageDescriptorSkewedInfo) CatalogTableStorageDescriptorSkewedInfo { return *v }).(CatalogTableStorageDescriptorSkewedInfoOutput)
}

// A list of names of columns that contain skewed values.
func (o CatalogTableStorageDescriptorSkewedInfoPtrOutput) SkewedColumnNames() pulumi.StringArrayOutput {
	return o.ApplyT(func (v CatalogTableStorageDescriptorSkewedInfo) []string { return v.SkewedColumnNames }).(pulumi.StringArrayOutput)
}

// A list of values that appear so frequently as to be considered skewed.
func (o CatalogTableStorageDescriptorSkewedInfoPtrOutput) SkewedColumnValueLocationMaps() pulumi.StringMapOutput {
	return o.ApplyT(func (v CatalogTableStorageDescriptorSkewedInfo) map[string]string { return v.SkewedColumnValueLocationMaps }).(pulumi.StringMapOutput)
}

// A mapping of skewed values to the columns that contain them.
func (o CatalogTableStorageDescriptorSkewedInfoPtrOutput) SkewedColumnValues() pulumi.StringArrayOutput {
	return o.ApplyT(func (v CatalogTableStorageDescriptorSkewedInfo) []string { return v.SkewedColumnValues }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(CatalogTableStorageDescriptorSkewedInfoOutput{})
	pulumi.RegisterOutputType(CatalogTableStorageDescriptorSkewedInfoPtrOutput{})
}
