// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getTableTtl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetTableTtl struct {
	AttributeName string `pulumi:"attributeName"`
	Enabled bool `pulumi:"enabled"`
}

type GetTableTtlInput interface {
	pulumi.Input

	ToGetTableTtlOutput() GetTableTtlOutput
	ToGetTableTtlOutputWithContext(context.Context) GetTableTtlOutput
}

type GetTableTtlArgs struct {
	AttributeName pulumi.StringInput `pulumi:"attributeName"`
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (GetTableTtlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTableTtl)(nil)).Elem()
}

func (i GetTableTtlArgs) ToGetTableTtlOutput() GetTableTtlOutput {
	return i.ToGetTableTtlOutputWithContext(context.Background())
}

func (i GetTableTtlArgs) ToGetTableTtlOutputWithContext(ctx context.Context) GetTableTtlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetTableTtlOutput)
}

type GetTableTtlOutput struct { *pulumi.OutputState }

func (GetTableTtlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTableTtl)(nil)).Elem()
}

func (o GetTableTtlOutput) ToGetTableTtlOutput() GetTableTtlOutput {
	return o
}

func (o GetTableTtlOutput) ToGetTableTtlOutputWithContext(ctx context.Context) GetTableTtlOutput {
	return o
}

func (o GetTableTtlOutput) AttributeName() pulumi.StringOutput {
	return o.ApplyT(func (v GetTableTtl) string { return v.AttributeName }).(pulumi.StringOutput)
}

func (o GetTableTtlOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func (v GetTableTtl) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTableTtlOutput{})
}
