// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getAmiBlockDeviceMapping

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetAmiBlockDeviceMapping struct {
	DeviceName string `pulumi:"deviceName"`
	Ebs map[string]interface{} `pulumi:"ebs"`
	NoDevice string `pulumi:"noDevice"`
	VirtualName string `pulumi:"virtualName"`
}

type GetAmiBlockDeviceMappingInput interface {
	pulumi.Input

	ToGetAmiBlockDeviceMappingOutput() GetAmiBlockDeviceMappingOutput
	ToGetAmiBlockDeviceMappingOutputWithContext(context.Context) GetAmiBlockDeviceMappingOutput
}

type GetAmiBlockDeviceMappingArgs struct {
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	Ebs pulumi.MapInput `pulumi:"ebs"`
	NoDevice pulumi.StringInput `pulumi:"noDevice"`
	VirtualName pulumi.StringInput `pulumi:"virtualName"`
}

func (GetAmiBlockDeviceMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAmiBlockDeviceMapping)(nil)).Elem()
}

func (i GetAmiBlockDeviceMappingArgs) ToGetAmiBlockDeviceMappingOutput() GetAmiBlockDeviceMappingOutput {
	return i.ToGetAmiBlockDeviceMappingOutputWithContext(context.Background())
}

func (i GetAmiBlockDeviceMappingArgs) ToGetAmiBlockDeviceMappingOutputWithContext(ctx context.Context) GetAmiBlockDeviceMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAmiBlockDeviceMappingOutput)
}

type GetAmiBlockDeviceMappingArrayInput interface {
	pulumi.Input

	ToGetAmiBlockDeviceMappingArrayOutput() GetAmiBlockDeviceMappingArrayOutput
	ToGetAmiBlockDeviceMappingArrayOutputWithContext(context.Context) GetAmiBlockDeviceMappingArrayOutput
}

type GetAmiBlockDeviceMappingArray []GetAmiBlockDeviceMappingInput

func (GetAmiBlockDeviceMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAmiBlockDeviceMapping)(nil)).Elem()
}

func (i GetAmiBlockDeviceMappingArray) ToGetAmiBlockDeviceMappingArrayOutput() GetAmiBlockDeviceMappingArrayOutput {
	return i.ToGetAmiBlockDeviceMappingArrayOutputWithContext(context.Background())
}

func (i GetAmiBlockDeviceMappingArray) ToGetAmiBlockDeviceMappingArrayOutputWithContext(ctx context.Context) GetAmiBlockDeviceMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAmiBlockDeviceMappingArrayOutput)
}

type GetAmiBlockDeviceMappingOutput struct { *pulumi.OutputState }

func (GetAmiBlockDeviceMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAmiBlockDeviceMapping)(nil)).Elem()
}

func (o GetAmiBlockDeviceMappingOutput) ToGetAmiBlockDeviceMappingOutput() GetAmiBlockDeviceMappingOutput {
	return o
}

func (o GetAmiBlockDeviceMappingOutput) ToGetAmiBlockDeviceMappingOutputWithContext(ctx context.Context) GetAmiBlockDeviceMappingOutput {
	return o
}

func (o GetAmiBlockDeviceMappingOutput) DeviceName() pulumi.StringOutput {
	return o.ApplyT(func (v GetAmiBlockDeviceMapping) string { return v.DeviceName }).(pulumi.StringOutput)
}

func (o GetAmiBlockDeviceMappingOutput) Ebs() pulumi.MapOutput {
	return o.ApplyT(func (v GetAmiBlockDeviceMapping) map[string]interface{} { return v.Ebs }).(pulumi.MapOutput)
}

func (o GetAmiBlockDeviceMappingOutput) NoDevice() pulumi.StringOutput {
	return o.ApplyT(func (v GetAmiBlockDeviceMapping) string { return v.NoDevice }).(pulumi.StringOutput)
}

func (o GetAmiBlockDeviceMappingOutput) VirtualName() pulumi.StringOutput {
	return o.ApplyT(func (v GetAmiBlockDeviceMapping) string { return v.VirtualName }).(pulumi.StringOutput)
}

type GetAmiBlockDeviceMappingArrayOutput struct { *pulumi.OutputState}

func (GetAmiBlockDeviceMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAmiBlockDeviceMapping)(nil)).Elem()
}

func (o GetAmiBlockDeviceMappingArrayOutput) ToGetAmiBlockDeviceMappingArrayOutput() GetAmiBlockDeviceMappingArrayOutput {
	return o
}

func (o GetAmiBlockDeviceMappingArrayOutput) ToGetAmiBlockDeviceMappingArrayOutputWithContext(ctx context.Context) GetAmiBlockDeviceMappingArrayOutput {
	return o
}

func (o GetAmiBlockDeviceMappingArrayOutput) Index(i pulumi.IntInput) GetAmiBlockDeviceMappingOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetAmiBlockDeviceMapping {
		return vs[0].([]GetAmiBlockDeviceMapping)[vs[1].(int)]
	}).(GetAmiBlockDeviceMappingOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAmiBlockDeviceMappingOutput{})
	pulumi.RegisterOutputType(GetAmiBlockDeviceMappingArrayOutput{})
}
