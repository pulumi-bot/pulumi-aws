// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dax

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClusterNode struct {
	Address          *string `pulumi:"address"`
	AvailabilityZone *string `pulumi:"availabilityZone"`
	Id               *string `pulumi:"id"`
	// The port used by the configuration endpoint
	Port *int `pulumi:"port"`
}

// ClusterNodeInput is an input type that accepts ClusterNodeArgs and ClusterNodeOutput values.
// You can construct a concrete instance of `ClusterNodeInput` via:
//
// 		 ClusterNodeArgs{...}
//
type ClusterNodeInput interface {
	pulumi.Input

	ToClusterNodeOutput() ClusterNodeOutput
	ToClusterNodeOutputWithContext(context.Context) ClusterNodeOutput
}

type ClusterNodeArgs struct {
	Address          pulumi.StringPtrInput `pulumi:"address"`
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	Id               pulumi.StringPtrInput `pulumi:"id"`
	// The port used by the configuration endpoint
	Port pulumi.IntPtrInput `pulumi:"port"`
}

func (ClusterNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterNode)(nil)).Elem()
}

func (i ClusterNodeArgs) ToClusterNodeOutput() ClusterNodeOutput {
	return i.ToClusterNodeOutputWithContext(context.Background())
}

func (i ClusterNodeArgs) ToClusterNodeOutputWithContext(ctx context.Context) ClusterNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterNodeOutput)
}

// ClusterNodeArrayInput is an input type that accepts ClusterNodeArray and ClusterNodeArrayOutput values.
// You can construct a concrete instance of `ClusterNodeArrayInput` via:
//
// 		 ClusterNodeArray{ ClusterNodeArgs{...} }
//
type ClusterNodeArrayInput interface {
	pulumi.Input

	ToClusterNodeArrayOutput() ClusterNodeArrayOutput
	ToClusterNodeArrayOutputWithContext(context.Context) ClusterNodeArrayOutput
}

type ClusterNodeArray []ClusterNodeInput

func (ClusterNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterNode)(nil)).Elem()
}

func (i ClusterNodeArray) ToClusterNodeArrayOutput() ClusterNodeArrayOutput {
	return i.ToClusterNodeArrayOutputWithContext(context.Background())
}

func (i ClusterNodeArray) ToClusterNodeArrayOutputWithContext(ctx context.Context) ClusterNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterNodeArrayOutput)
}

type ClusterNodeOutput struct{ *pulumi.OutputState }

func (ClusterNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterNode)(nil)).Elem()
}

func (o ClusterNodeOutput) ToClusterNodeOutput() ClusterNodeOutput {
	return o
}

func (o ClusterNodeOutput) ToClusterNodeOutputWithContext(ctx context.Context) ClusterNodeOutput {
	return o
}

func (o ClusterNodeOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterNode) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o ClusterNodeOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterNode) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o ClusterNodeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterNode) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The port used by the configuration endpoint
func (o ClusterNodeOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterNode) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type ClusterNodeArrayOutput struct{ *pulumi.OutputState }

func (ClusterNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterNode)(nil)).Elem()
}

func (o ClusterNodeArrayOutput) ToClusterNodeArrayOutput() ClusterNodeArrayOutput {
	return o
}

func (o ClusterNodeArrayOutput) ToClusterNodeArrayOutputWithContext(ctx context.Context) ClusterNodeArrayOutput {
	return o
}

func (o ClusterNodeArrayOutput) Index(i pulumi.IntInput) ClusterNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterNode {
		return vs[0].([]ClusterNode)[vs[1].(int)]
	}).(ClusterNodeOutput)
}

type ClusterServerSideEncryption struct {
	// Whether to enable encryption at rest. Defaults to `false`.
	Enabled *bool `pulumi:"enabled"`
}

// ClusterServerSideEncryptionInput is an input type that accepts ClusterServerSideEncryptionArgs and ClusterServerSideEncryptionOutput values.
// You can construct a concrete instance of `ClusterServerSideEncryptionInput` via:
//
// 		 ClusterServerSideEncryptionArgs{...}
//
type ClusterServerSideEncryptionInput interface {
	pulumi.Input

	ToClusterServerSideEncryptionOutput() ClusterServerSideEncryptionOutput
	ToClusterServerSideEncryptionOutputWithContext(context.Context) ClusterServerSideEncryptionOutput
}

type ClusterServerSideEncryptionArgs struct {
	// Whether to enable encryption at rest. Defaults to `false`.
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (ClusterServerSideEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterServerSideEncryption)(nil)).Elem()
}

func (i ClusterServerSideEncryptionArgs) ToClusterServerSideEncryptionOutput() ClusterServerSideEncryptionOutput {
	return i.ToClusterServerSideEncryptionOutputWithContext(context.Background())
}

func (i ClusterServerSideEncryptionArgs) ToClusterServerSideEncryptionOutputWithContext(ctx context.Context) ClusterServerSideEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterServerSideEncryptionOutput)
}

func (i ClusterServerSideEncryptionArgs) ToClusterServerSideEncryptionPtrOutput() ClusterServerSideEncryptionPtrOutput {
	return i.ToClusterServerSideEncryptionPtrOutputWithContext(context.Background())
}

func (i ClusterServerSideEncryptionArgs) ToClusterServerSideEncryptionPtrOutputWithContext(ctx context.Context) ClusterServerSideEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterServerSideEncryptionOutput).ToClusterServerSideEncryptionPtrOutputWithContext(ctx)
}

// ClusterServerSideEncryptionPtrInput is an input type that accepts ClusterServerSideEncryptionArgs, ClusterServerSideEncryptionPtr and ClusterServerSideEncryptionPtrOutput values.
// You can construct a concrete instance of `ClusterServerSideEncryptionPtrInput` via:
//
// 		 ClusterServerSideEncryptionArgs{...}
//
//  or:
//
// 		 nil
//
type ClusterServerSideEncryptionPtrInput interface {
	pulumi.Input

	ToClusterServerSideEncryptionPtrOutput() ClusterServerSideEncryptionPtrOutput
	ToClusterServerSideEncryptionPtrOutputWithContext(context.Context) ClusterServerSideEncryptionPtrOutput
}

type clusterServerSideEncryptionPtrType ClusterServerSideEncryptionArgs

func ClusterServerSideEncryptionPtr(v *ClusterServerSideEncryptionArgs) ClusterServerSideEncryptionPtrInput {
	return (*clusterServerSideEncryptionPtrType)(v)
}

func (*clusterServerSideEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterServerSideEncryption)(nil)).Elem()
}

func (i *clusterServerSideEncryptionPtrType) ToClusterServerSideEncryptionPtrOutput() ClusterServerSideEncryptionPtrOutput {
	return i.ToClusterServerSideEncryptionPtrOutputWithContext(context.Background())
}

func (i *clusterServerSideEncryptionPtrType) ToClusterServerSideEncryptionPtrOutputWithContext(ctx context.Context) ClusterServerSideEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterServerSideEncryptionPtrOutput)
}

type ClusterServerSideEncryptionOutput struct{ *pulumi.OutputState }

func (ClusterServerSideEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterServerSideEncryption)(nil)).Elem()
}

func (o ClusterServerSideEncryptionOutput) ToClusterServerSideEncryptionOutput() ClusterServerSideEncryptionOutput {
	return o
}

func (o ClusterServerSideEncryptionOutput) ToClusterServerSideEncryptionOutputWithContext(ctx context.Context) ClusterServerSideEncryptionOutput {
	return o
}

func (o ClusterServerSideEncryptionOutput) ToClusterServerSideEncryptionPtrOutput() ClusterServerSideEncryptionPtrOutput {
	return o.ToClusterServerSideEncryptionPtrOutputWithContext(context.Background())
}

func (o ClusterServerSideEncryptionOutput) ToClusterServerSideEncryptionPtrOutputWithContext(ctx context.Context) ClusterServerSideEncryptionPtrOutput {
	return o.ApplyT(func(v ClusterServerSideEncryption) *ClusterServerSideEncryption {
		return &v
	}).(ClusterServerSideEncryptionPtrOutput)
}

// Whether to enable encryption at rest. Defaults to `false`.
func (o ClusterServerSideEncryptionOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterServerSideEncryption) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type ClusterServerSideEncryptionPtrOutput struct{ *pulumi.OutputState }

func (ClusterServerSideEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterServerSideEncryption)(nil)).Elem()
}

func (o ClusterServerSideEncryptionPtrOutput) ToClusterServerSideEncryptionPtrOutput() ClusterServerSideEncryptionPtrOutput {
	return o
}

func (o ClusterServerSideEncryptionPtrOutput) ToClusterServerSideEncryptionPtrOutputWithContext(ctx context.Context) ClusterServerSideEncryptionPtrOutput {
	return o
}

func (o ClusterServerSideEncryptionPtrOutput) Elem() ClusterServerSideEncryptionOutput {
	return o.ApplyT(func(v *ClusterServerSideEncryption) ClusterServerSideEncryption { return *v }).(ClusterServerSideEncryptionOutput)
}

// Whether to enable encryption at rest. Defaults to `false`.
func (o ClusterServerSideEncryptionPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterServerSideEncryption) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type ParameterGroupParameter struct {
	// The name of the parameter group.
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

// ParameterGroupParameterInput is an input type that accepts ParameterGroupParameterArgs and ParameterGroupParameterOutput values.
// You can construct a concrete instance of `ParameterGroupParameterInput` via:
//
// 		 ParameterGroupParameterArgs{...}
//
type ParameterGroupParameterInput interface {
	pulumi.Input

	ToParameterGroupParameterOutput() ParameterGroupParameterOutput
	ToParameterGroupParameterOutputWithContext(context.Context) ParameterGroupParameterOutput
}

type ParameterGroupParameterArgs struct {
	// The name of the parameter group.
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (ParameterGroupParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameter)(nil)).Elem()
}

func (i ParameterGroupParameterArgs) ToParameterGroupParameterOutput() ParameterGroupParameterOutput {
	return i.ToParameterGroupParameterOutputWithContext(context.Background())
}

func (i ParameterGroupParameterArgs) ToParameterGroupParameterOutputWithContext(ctx context.Context) ParameterGroupParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterOutput)
}

// ParameterGroupParameterArrayInput is an input type that accepts ParameterGroupParameterArray and ParameterGroupParameterArrayOutput values.
// You can construct a concrete instance of `ParameterGroupParameterArrayInput` via:
//
// 		 ParameterGroupParameterArray{ ParameterGroupParameterArgs{...} }
//
type ParameterGroupParameterArrayInput interface {
	pulumi.Input

	ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput
	ToParameterGroupParameterArrayOutputWithContext(context.Context) ParameterGroupParameterArrayOutput
}

type ParameterGroupParameterArray []ParameterGroupParameterInput

func (ParameterGroupParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameter)(nil)).Elem()
}

func (i ParameterGroupParameterArray) ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput {
	return i.ToParameterGroupParameterArrayOutputWithContext(context.Background())
}

func (i ParameterGroupParameterArray) ToParameterGroupParameterArrayOutputWithContext(ctx context.Context) ParameterGroupParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterArrayOutput)
}

type ParameterGroupParameterOutput struct{ *pulumi.OutputState }

func (ParameterGroupParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameter)(nil)).Elem()
}

func (o ParameterGroupParameterOutput) ToParameterGroupParameterOutput() ParameterGroupParameterOutput {
	return o
}

func (o ParameterGroupParameterOutput) ToParameterGroupParameterOutputWithContext(ctx context.Context) ParameterGroupParameterOutput {
	return o
}

// The name of the parameter group.
func (o ParameterGroupParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupParameter) string { return v.Name }).(pulumi.StringOutput)
}

func (o ParameterGroupParameterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupParameter) string { return v.Value }).(pulumi.StringOutput)
}

type ParameterGroupParameterArrayOutput struct{ *pulumi.OutputState }

func (ParameterGroupParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameter)(nil)).Elem()
}

func (o ParameterGroupParameterArrayOutput) ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput {
	return o
}

func (o ParameterGroupParameterArrayOutput) ToParameterGroupParameterArrayOutputWithContext(ctx context.Context) ParameterGroupParameterArrayOutput {
	return o
}

func (o ParameterGroupParameterArrayOutput) Index(i pulumi.IntInput) ParameterGroupParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterGroupParameter {
		return vs[0].([]ParameterGroupParameter)[vs[1].(int)]
	}).(ParameterGroupParameterOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterNodeOutput{})
	pulumi.RegisterOutputType(ClusterNodeArrayOutput{})
	pulumi.RegisterOutputType(ClusterServerSideEncryptionOutput{})
	pulumi.RegisterOutputType(ClusterServerSideEncryptionPtrOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterArrayOutput{})
}
