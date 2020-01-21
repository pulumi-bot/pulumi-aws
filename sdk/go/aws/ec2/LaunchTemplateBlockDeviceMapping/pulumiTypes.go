// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package LaunchTemplateBlockDeviceMapping

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/ec2/LaunchTemplateBlockDeviceMappingEbs"
)

type LaunchTemplateBlockDeviceMapping struct {
	// The name of the device to mount.
	DeviceName *string `pulumi:"deviceName"`
	// Configure EBS volume properties.
	Ebs *ec2LaunchTemplateBlockDeviceMappingEbs.LaunchTemplateBlockDeviceMappingEbs `pulumi:"ebs"`
	// Suppresses the specified device included in the AMI's block device mapping.
	NoDevice *string `pulumi:"noDevice"`
	// The [Instance Store Device
	// Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames)
	// (e.g. `"ephemeral0"`).
	VirtualName *string `pulumi:"virtualName"`
}

type LaunchTemplateBlockDeviceMappingInput interface {
	pulumi.Input

	ToLaunchTemplateBlockDeviceMappingOutput() LaunchTemplateBlockDeviceMappingOutput
	ToLaunchTemplateBlockDeviceMappingOutputWithContext(context.Context) LaunchTemplateBlockDeviceMappingOutput
}

type LaunchTemplateBlockDeviceMappingArgs struct {
	// The name of the device to mount.
	DeviceName pulumi.StringPtrInput `pulumi:"deviceName"`
	// Configure EBS volume properties.
	Ebs ec2LaunchTemplateBlockDeviceMappingEbs.LaunchTemplateBlockDeviceMappingEbsPtrInput `pulumi:"ebs"`
	// Suppresses the specified device included in the AMI's block device mapping.
	NoDevice pulumi.StringPtrInput `pulumi:"noDevice"`
	// The [Instance Store Device
	// Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames)
	// (e.g. `"ephemeral0"`).
	VirtualName pulumi.StringPtrInput `pulumi:"virtualName"`
}

func (LaunchTemplateBlockDeviceMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchTemplateBlockDeviceMapping)(nil)).Elem()
}

func (i LaunchTemplateBlockDeviceMappingArgs) ToLaunchTemplateBlockDeviceMappingOutput() LaunchTemplateBlockDeviceMappingOutput {
	return i.ToLaunchTemplateBlockDeviceMappingOutputWithContext(context.Background())
}

func (i LaunchTemplateBlockDeviceMappingArgs) ToLaunchTemplateBlockDeviceMappingOutputWithContext(ctx context.Context) LaunchTemplateBlockDeviceMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateBlockDeviceMappingOutput)
}

type LaunchTemplateBlockDeviceMappingArrayInput interface {
	pulumi.Input

	ToLaunchTemplateBlockDeviceMappingArrayOutput() LaunchTemplateBlockDeviceMappingArrayOutput
	ToLaunchTemplateBlockDeviceMappingArrayOutputWithContext(context.Context) LaunchTemplateBlockDeviceMappingArrayOutput
}

type LaunchTemplateBlockDeviceMappingArray []LaunchTemplateBlockDeviceMappingInput

func (LaunchTemplateBlockDeviceMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LaunchTemplateBlockDeviceMapping)(nil)).Elem()
}

func (i LaunchTemplateBlockDeviceMappingArray) ToLaunchTemplateBlockDeviceMappingArrayOutput() LaunchTemplateBlockDeviceMappingArrayOutput {
	return i.ToLaunchTemplateBlockDeviceMappingArrayOutputWithContext(context.Background())
}

func (i LaunchTemplateBlockDeviceMappingArray) ToLaunchTemplateBlockDeviceMappingArrayOutputWithContext(ctx context.Context) LaunchTemplateBlockDeviceMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateBlockDeviceMappingArrayOutput)
}

type LaunchTemplateBlockDeviceMappingOutput struct { *pulumi.OutputState }

func (LaunchTemplateBlockDeviceMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LaunchTemplateBlockDeviceMapping)(nil)).Elem()
}

func (o LaunchTemplateBlockDeviceMappingOutput) ToLaunchTemplateBlockDeviceMappingOutput() LaunchTemplateBlockDeviceMappingOutput {
	return o
}

func (o LaunchTemplateBlockDeviceMappingOutput) ToLaunchTemplateBlockDeviceMappingOutputWithContext(ctx context.Context) LaunchTemplateBlockDeviceMappingOutput {
	return o
}

// The name of the device to mount.
func (o LaunchTemplateBlockDeviceMappingOutput) DeviceName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v LaunchTemplateBlockDeviceMapping) *string { return v.DeviceName }).(pulumi.StringPtrOutput)
}

// Configure EBS volume properties.
func (o LaunchTemplateBlockDeviceMappingOutput) Ebs() ec2LaunchTemplateBlockDeviceMappingEbs.LaunchTemplateBlockDeviceMappingEbsPtrOutput {
	return o.ApplyT(func (v LaunchTemplateBlockDeviceMapping) *ec2LaunchTemplateBlockDeviceMappingEbs.LaunchTemplateBlockDeviceMappingEbs { return v.Ebs }).(ec2LaunchTemplateBlockDeviceMappingEbs.LaunchTemplateBlockDeviceMappingEbsPtrOutput)
}

// Suppresses the specified device included in the AMI's block device mapping.
func (o LaunchTemplateBlockDeviceMappingOutput) NoDevice() pulumi.StringPtrOutput {
	return o.ApplyT(func (v LaunchTemplateBlockDeviceMapping) *string { return v.NoDevice }).(pulumi.StringPtrOutput)
}

// The [Instance Store Device
// Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames)
// (e.g. `"ephemeral0"`).
func (o LaunchTemplateBlockDeviceMappingOutput) VirtualName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v LaunchTemplateBlockDeviceMapping) *string { return v.VirtualName }).(pulumi.StringPtrOutput)
}

type LaunchTemplateBlockDeviceMappingArrayOutput struct { *pulumi.OutputState}

func (LaunchTemplateBlockDeviceMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LaunchTemplateBlockDeviceMapping)(nil)).Elem()
}

func (o LaunchTemplateBlockDeviceMappingArrayOutput) ToLaunchTemplateBlockDeviceMappingArrayOutput() LaunchTemplateBlockDeviceMappingArrayOutput {
	return o
}

func (o LaunchTemplateBlockDeviceMappingArrayOutput) ToLaunchTemplateBlockDeviceMappingArrayOutputWithContext(ctx context.Context) LaunchTemplateBlockDeviceMappingArrayOutput {
	return o
}

func (o LaunchTemplateBlockDeviceMappingArrayOutput) Index(i pulumi.IntInput) LaunchTemplateBlockDeviceMappingOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) LaunchTemplateBlockDeviceMapping {
		return vs[0].([]LaunchTemplateBlockDeviceMapping)[vs[1].(int)]
	}).(LaunchTemplateBlockDeviceMappingOutput)
}

func init() {
	pulumi.RegisterOutputType(LaunchTemplateBlockDeviceMappingOutput{})
	pulumi.RegisterOutputType(LaunchTemplateBlockDeviceMappingArrayOutput{})
}
