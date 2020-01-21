// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterSetting

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClusterSetting struct {
	// The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
	Name string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type ClusterSettingInput interface {
	pulumi.Input

	ToClusterSettingOutput() ClusterSettingOutput
	ToClusterSettingOutputWithContext(context.Context) ClusterSettingOutput
}

type ClusterSettingArgs struct {
	// The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
	Name pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (ClusterSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSetting)(nil)).Elem()
}

func (i ClusterSettingArgs) ToClusterSettingOutput() ClusterSettingOutput {
	return i.ToClusterSettingOutputWithContext(context.Background())
}

func (i ClusterSettingArgs) ToClusterSettingOutputWithContext(ctx context.Context) ClusterSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSettingOutput)
}

type ClusterSettingArrayInput interface {
	pulumi.Input

	ToClusterSettingArrayOutput() ClusterSettingArrayOutput
	ToClusterSettingArrayOutputWithContext(context.Context) ClusterSettingArrayOutput
}

type ClusterSettingArray []ClusterSettingInput

func (ClusterSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterSetting)(nil)).Elem()
}

func (i ClusterSettingArray) ToClusterSettingArrayOutput() ClusterSettingArrayOutput {
	return i.ToClusterSettingArrayOutputWithContext(context.Background())
}

func (i ClusterSettingArray) ToClusterSettingArrayOutputWithContext(ctx context.Context) ClusterSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSettingArrayOutput)
}

type ClusterSettingOutput struct { *pulumi.OutputState }

func (ClusterSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSetting)(nil)).Elem()
}

func (o ClusterSettingOutput) ToClusterSettingOutput() ClusterSettingOutput {
	return o
}

func (o ClusterSettingOutput) ToClusterSettingOutputWithContext(ctx context.Context) ClusterSettingOutput {
	return o
}

// The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
func (o ClusterSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterSetting) string { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterSettingOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterSetting) string { return v.Value }).(pulumi.StringOutput)
}

type ClusterSettingArrayOutput struct { *pulumi.OutputState}

func (ClusterSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterSetting)(nil)).Elem()
}

func (o ClusterSettingArrayOutput) ToClusterSettingArrayOutput() ClusterSettingArrayOutput {
	return o
}

func (o ClusterSettingArrayOutput) ToClusterSettingArrayOutputWithContext(ctx context.Context) ClusterSettingArrayOutput {
	return o
}

func (o ClusterSettingArrayOutput) Index(i pulumi.IntInput) ClusterSettingOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ClusterSetting {
		return vs[0].([]ClusterSetting)[vs[1].(int)]
	}).(ClusterSettingOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterSettingOutput{})
	pulumi.RegisterOutputType(ClusterSettingArrayOutput{})
}
