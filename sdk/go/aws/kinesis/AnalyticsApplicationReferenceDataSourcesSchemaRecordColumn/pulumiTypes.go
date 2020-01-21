// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package AnalyticsApplicationReferenceDataSourcesSchemaRecordColumn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type AnalyticsApplicationReferenceDataSourcesSchemaRecordColumn struct {
	Mapping *string `pulumi:"mapping"`
	// Name of the Kinesis Analytics Application.
	Name string `pulumi:"name"`
	SqlType string `pulumi:"sqlType"`
}

type AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnInput interface {
	pulumi.Input

	ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput() AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput
	ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutputWithContext(context.Context) AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput
}

type AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArgs struct {
	Mapping pulumi.StringPtrInput `pulumi:"mapping"`
	// Name of the Kinesis Analytics Application.
	Name pulumi.StringInput `pulumi:"name"`
	SqlType pulumi.StringInput `pulumi:"sqlType"`
}

func (AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsApplicationReferenceDataSourcesSchemaRecordColumn)(nil)).Elem()
}

func (i AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArgs) ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput() AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput {
	return i.ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutputWithContext(context.Background())
}

func (i AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArgs) ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutputWithContext(ctx context.Context) AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput)
}

type AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayInput interface {
	pulumi.Input

	ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput() AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput
	ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutputWithContext(context.Context) AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput
}

type AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArray []AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnInput

func (AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AnalyticsApplicationReferenceDataSourcesSchemaRecordColumn)(nil)).Elem()
}

func (i AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArray) ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput() AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput {
	return i.ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutputWithContext(context.Background())
}

func (i AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArray) ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutputWithContext(ctx context.Context) AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput)
}

type AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput struct { *pulumi.OutputState }

func (AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsApplicationReferenceDataSourcesSchemaRecordColumn)(nil)).Elem()
}

func (o AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput) ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput() AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput {
	return o
}

func (o AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput) ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutputWithContext(ctx context.Context) AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput {
	return o
}

func (o AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput) Mapping() pulumi.StringPtrOutput {
	return o.ApplyT(func (v AnalyticsApplicationReferenceDataSourcesSchemaRecordColumn) *string { return v.Mapping }).(pulumi.StringPtrOutput)
}

// Name of the Kinesis Analytics Application.
func (o AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v AnalyticsApplicationReferenceDataSourcesSchemaRecordColumn) string { return v.Name }).(pulumi.StringOutput)
}

func (o AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput) SqlType() pulumi.StringOutput {
	return o.ApplyT(func (v AnalyticsApplicationReferenceDataSourcesSchemaRecordColumn) string { return v.SqlType }).(pulumi.StringOutput)
}

type AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput struct { *pulumi.OutputState}

func (AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AnalyticsApplicationReferenceDataSourcesSchemaRecordColumn)(nil)).Elem()
}

func (o AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput) ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput() AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput {
	return o
}

func (o AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput) ToAnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutputWithContext(ctx context.Context) AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput {
	return o
}

func (o AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput) Index(i pulumi.IntInput) AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) AnalyticsApplicationReferenceDataSourcesSchemaRecordColumn {
		return vs[0].([]AnalyticsApplicationReferenceDataSourcesSchemaRecordColumn)[vs[1].(int)]
	}).(AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput)
}

func init() {
	pulumi.RegisterOutputType(AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnOutput{})
	pulumi.RegisterOutputType(AnalyticsApplicationReferenceDataSourcesSchemaRecordColumnArrayOutput{})
}
