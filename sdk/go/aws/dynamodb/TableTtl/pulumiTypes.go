// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package TableTtl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type TableTtl struct {
	// The name of the table attribute to store the TTL timestamp in.
	AttributeName string `pulumi:"attributeName"`
	// Whether to enable point-in-time recovery - note that it can take up to 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided then this defaults to `false`.
	Enabled *bool `pulumi:"enabled"`
}

type TableTtlInput interface {
	pulumi.Input

	ToTableTtlOutput() TableTtlOutput
	ToTableTtlOutputWithContext(context.Context) TableTtlOutput
}

type TableTtlArgs struct {
	// The name of the table attribute to store the TTL timestamp in.
	AttributeName pulumi.StringInput `pulumi:"attributeName"`
	// Whether to enable point-in-time recovery - note that it can take up to 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided then this defaults to `false`.
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (TableTtlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableTtl)(nil)).Elem()
}

func (i TableTtlArgs) ToTableTtlOutput() TableTtlOutput {
	return i.ToTableTtlOutputWithContext(context.Background())
}

func (i TableTtlArgs) ToTableTtlOutputWithContext(ctx context.Context) TableTtlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableTtlOutput)
}

func (i TableTtlArgs) ToTableTtlPtrOutput() TableTtlPtrOutput {
	return i.ToTableTtlPtrOutputWithContext(context.Background())
}

func (i TableTtlArgs) ToTableTtlPtrOutputWithContext(ctx context.Context) TableTtlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableTtlOutput).ToTableTtlPtrOutputWithContext(ctx)
}

type TableTtlPtrInput interface {
	pulumi.Input

	ToTableTtlPtrOutput() TableTtlPtrOutput
	ToTableTtlPtrOutputWithContext(context.Context) TableTtlPtrOutput
}

type tableTtlPtrType TableTtlArgs

func TableTtlPtr(v *TableTtlArgs) TableTtlPtrInput {	return (*tableTtlPtrType)(v)
}

func (*tableTtlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableTtl)(nil)).Elem()
}

func (i *tableTtlPtrType) ToTableTtlPtrOutput() TableTtlPtrOutput {
	return i.ToTableTtlPtrOutputWithContext(context.Background())
}

func (i *tableTtlPtrType) ToTableTtlPtrOutputWithContext(ctx context.Context) TableTtlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableTtlPtrOutput)
}

type TableTtlOutput struct { *pulumi.OutputState }

func (TableTtlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableTtl)(nil)).Elem()
}

func (o TableTtlOutput) ToTableTtlOutput() TableTtlOutput {
	return o
}

func (o TableTtlOutput) ToTableTtlOutputWithContext(ctx context.Context) TableTtlOutput {
	return o
}

func (o TableTtlOutput) ToTableTtlPtrOutput() TableTtlPtrOutput {
	return o.ToTableTtlPtrOutputWithContext(context.Background())
}

func (o TableTtlOutput) ToTableTtlPtrOutputWithContext(ctx context.Context) TableTtlPtrOutput {
	return o.ApplyT(func(v TableTtl) *TableTtl {
		return &v
	}).(TableTtlPtrOutput)
}
// The name of the table attribute to store the TTL timestamp in.
func (o TableTtlOutput) AttributeName() pulumi.StringOutput {
	return o.ApplyT(func (v TableTtl) string { return v.AttributeName }).(pulumi.StringOutput)
}

// Whether to enable point-in-time recovery - note that it can take up to 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided then this defaults to `false`.
func (o TableTtlOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v TableTtl) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type TableTtlPtrOutput struct { *pulumi.OutputState}

func (TableTtlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableTtl)(nil)).Elem()
}

func (o TableTtlPtrOutput) ToTableTtlPtrOutput() TableTtlPtrOutput {
	return o
}

func (o TableTtlPtrOutput) ToTableTtlPtrOutputWithContext(ctx context.Context) TableTtlPtrOutput {
	return o
}

func (o TableTtlPtrOutput) Elem() TableTtlOutput {
	return o.ApplyT(func (v *TableTtl) TableTtl { return *v }).(TableTtlOutput)
}

// The name of the table attribute to store the TTL timestamp in.
func (o TableTtlPtrOutput) AttributeName() pulumi.StringOutput {
	return o.ApplyT(func (v TableTtl) string { return v.AttributeName }).(pulumi.StringOutput)
}

// Whether to enable point-in-time recovery - note that it can take up to 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided then this defaults to `false`.
func (o TableTtlPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v TableTtl) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(TableTtlOutput{})
	pulumi.RegisterOutputType(TableTtlPtrOutput{})
}
