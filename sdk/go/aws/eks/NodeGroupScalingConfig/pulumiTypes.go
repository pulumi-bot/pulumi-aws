// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package NodeGroupScalingConfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type NodeGroupScalingConfig struct {
	// Desired number of worker nodes.
	DesiredSize int `pulumi:"desiredSize"`
	// Maximum number of worker nodes.
	MaxSize int `pulumi:"maxSize"`
	// Minimum number of worker nodes.
	MinSize int `pulumi:"minSize"`
}

type NodeGroupScalingConfigInput interface {
	pulumi.Input

	ToNodeGroupScalingConfigOutput() NodeGroupScalingConfigOutput
	ToNodeGroupScalingConfigOutputWithContext(context.Context) NodeGroupScalingConfigOutput
}

type NodeGroupScalingConfigArgs struct {
	// Desired number of worker nodes.
	DesiredSize pulumi.IntInput `pulumi:"desiredSize"`
	// Maximum number of worker nodes.
	MaxSize pulumi.IntInput `pulumi:"maxSize"`
	// Minimum number of worker nodes.
	MinSize pulumi.IntInput `pulumi:"minSize"`
}

func (NodeGroupScalingConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeGroupScalingConfig)(nil)).Elem()
}

func (i NodeGroupScalingConfigArgs) ToNodeGroupScalingConfigOutput() NodeGroupScalingConfigOutput {
	return i.ToNodeGroupScalingConfigOutputWithContext(context.Background())
}

func (i NodeGroupScalingConfigArgs) ToNodeGroupScalingConfigOutputWithContext(ctx context.Context) NodeGroupScalingConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupScalingConfigOutput)
}

func (i NodeGroupScalingConfigArgs) ToNodeGroupScalingConfigPtrOutput() NodeGroupScalingConfigPtrOutput {
	return i.ToNodeGroupScalingConfigPtrOutputWithContext(context.Background())
}

func (i NodeGroupScalingConfigArgs) ToNodeGroupScalingConfigPtrOutputWithContext(ctx context.Context) NodeGroupScalingConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupScalingConfigOutput).ToNodeGroupScalingConfigPtrOutputWithContext(ctx)
}

type NodeGroupScalingConfigPtrInput interface {
	pulumi.Input

	ToNodeGroupScalingConfigPtrOutput() NodeGroupScalingConfigPtrOutput
	ToNodeGroupScalingConfigPtrOutputWithContext(context.Context) NodeGroupScalingConfigPtrOutput
}

type nodeGroupScalingConfigPtrType NodeGroupScalingConfigArgs

func NodeGroupScalingConfigPtr(v *NodeGroupScalingConfigArgs) NodeGroupScalingConfigPtrInput {	return (*nodeGroupScalingConfigPtrType)(v)
}

func (*nodeGroupScalingConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeGroupScalingConfig)(nil)).Elem()
}

func (i *nodeGroupScalingConfigPtrType) ToNodeGroupScalingConfigPtrOutput() NodeGroupScalingConfigPtrOutput {
	return i.ToNodeGroupScalingConfigPtrOutputWithContext(context.Background())
}

func (i *nodeGroupScalingConfigPtrType) ToNodeGroupScalingConfigPtrOutputWithContext(ctx context.Context) NodeGroupScalingConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupScalingConfigPtrOutput)
}

type NodeGroupScalingConfigOutput struct { *pulumi.OutputState }

func (NodeGroupScalingConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeGroupScalingConfig)(nil)).Elem()
}

func (o NodeGroupScalingConfigOutput) ToNodeGroupScalingConfigOutput() NodeGroupScalingConfigOutput {
	return o
}

func (o NodeGroupScalingConfigOutput) ToNodeGroupScalingConfigOutputWithContext(ctx context.Context) NodeGroupScalingConfigOutput {
	return o
}

func (o NodeGroupScalingConfigOutput) ToNodeGroupScalingConfigPtrOutput() NodeGroupScalingConfigPtrOutput {
	return o.ToNodeGroupScalingConfigPtrOutputWithContext(context.Background())
}

func (o NodeGroupScalingConfigOutput) ToNodeGroupScalingConfigPtrOutputWithContext(ctx context.Context) NodeGroupScalingConfigPtrOutput {
	return o.ApplyT(func(v NodeGroupScalingConfig) *NodeGroupScalingConfig {
		return &v
	}).(NodeGroupScalingConfigPtrOutput)
}
// Desired number of worker nodes.
func (o NodeGroupScalingConfigOutput) DesiredSize() pulumi.IntOutput {
	return o.ApplyT(func (v NodeGroupScalingConfig) int { return v.DesiredSize }).(pulumi.IntOutput)
}

// Maximum number of worker nodes.
func (o NodeGroupScalingConfigOutput) MaxSize() pulumi.IntOutput {
	return o.ApplyT(func (v NodeGroupScalingConfig) int { return v.MaxSize }).(pulumi.IntOutput)
}

// Minimum number of worker nodes.
func (o NodeGroupScalingConfigOutput) MinSize() pulumi.IntOutput {
	return o.ApplyT(func (v NodeGroupScalingConfig) int { return v.MinSize }).(pulumi.IntOutput)
}

type NodeGroupScalingConfigPtrOutput struct { *pulumi.OutputState}

func (NodeGroupScalingConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeGroupScalingConfig)(nil)).Elem()
}

func (o NodeGroupScalingConfigPtrOutput) ToNodeGroupScalingConfigPtrOutput() NodeGroupScalingConfigPtrOutput {
	return o
}

func (o NodeGroupScalingConfigPtrOutput) ToNodeGroupScalingConfigPtrOutputWithContext(ctx context.Context) NodeGroupScalingConfigPtrOutput {
	return o
}

func (o NodeGroupScalingConfigPtrOutput) Elem() NodeGroupScalingConfigOutput {
	return o.ApplyT(func (v *NodeGroupScalingConfig) NodeGroupScalingConfig { return *v }).(NodeGroupScalingConfigOutput)
}

// Desired number of worker nodes.
func (o NodeGroupScalingConfigPtrOutput) DesiredSize() pulumi.IntOutput {
	return o.ApplyT(func (v NodeGroupScalingConfig) int { return v.DesiredSize }).(pulumi.IntOutput)
}

// Maximum number of worker nodes.
func (o NodeGroupScalingConfigPtrOutput) MaxSize() pulumi.IntOutput {
	return o.ApplyT(func (v NodeGroupScalingConfig) int { return v.MaxSize }).(pulumi.IntOutput)
}

// Minimum number of worker nodes.
func (o NodeGroupScalingConfigPtrOutput) MinSize() pulumi.IntOutput {
	return o.ApplyT(func (v NodeGroupScalingConfig) int { return v.MinSize }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(NodeGroupScalingConfigOutput{})
	pulumi.RegisterOutputType(NodeGroupScalingConfigPtrOutput{})
}
