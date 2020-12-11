// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS EBS Volume Attachment as a top level resource, to attach and
// detach volumes from AWS Instances.
//
// > **NOTE on EBS block devices:** If you use `ebsBlockDevice` on an `ec2.Instance`, this provider will assume management over the full set of non-root EBS block devices for the instance, and treats additional block devices as drift. For this reason, `ebsBlockDevice` cannot be mixed with external `ebs.Volume` + `awsEbsVolumeAttachment` resources for a given instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ebs"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		web, err := ec2.NewInstance(ctx, "web", &ec2.InstanceArgs{
// 			Ami:              pulumi.String("ami-21f78e11"),
// 			AvailabilityZone: pulumi.String("us-west-2a"),
// 			InstanceType:     pulumi.String("t2.micro"),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("HelloWorld"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example, err := ebs.NewVolume(ctx, "example", &ebs.VolumeArgs{
// 			AvailabilityZone: pulumi.String("us-west-2a"),
// 			Size:             pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewVolumeAttachment(ctx, "ebsAtt", &ec2.VolumeAttachmentArgs{
// 			DeviceName: pulumi.String("/dev/sdh"),
// 			VolumeId:   example.ID(),
// 			InstanceId: web.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// EBS Volume Attachments can be imported using `DEVICE_NAME:VOLUME_ID:INSTANCE_ID`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/volumeAttachment:VolumeAttachment example /dev/sdh:vol-049df61146c4d7901:i-12345678
// ```
//
//  [1]https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html#available-ec2-device-names [2]https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/device_naming.html#available-ec2-device-names [3]https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-detaching-volume.html
type VolumeAttachment struct {
	pulumi.CustomResourceState

	// The device name to expose to the instance (for
	// example, `/dev/sdh` or `xvdh`).  See [Device Naming on Linux Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html#available-ec2-device-names) and [Device Naming on Windows Instances](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/device_naming.html#available-ec2-device-names) for more information.
	DeviceName pulumi.StringOutput `pulumi:"deviceName"`
	// Set to `true` if you want to force the
	// volume to detach. Useful if previous attempts failed, but use this option only
	// as a last resort, as this can result in **data loss**. See
	// [Detaching an Amazon EBS Volume from an Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-detaching-volume.html) for more information.
	ForceDetach pulumi.BoolPtrOutput `pulumi:"forceDetach"`
	// ID of the Instance to attach to
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Set this to true if you do not wish
	// to detach the volume from the instance to which it is attached at destroy
	// time, and instead just remove the attachment from this provider state. This is
	// useful when destroying an instance which has volumes created by some other
	// means attached.
	SkipDestroy pulumi.BoolPtrOutput `pulumi:"skipDestroy"`
	// ID of the Volume to be attached
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
}

// NewVolumeAttachment registers a new resource with the given unique name, arguments, and options.
func NewVolumeAttachment(ctx *pulumi.Context,
	name string, args *VolumeAttachmentArgs, opts ...pulumi.ResourceOption) (*VolumeAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.VolumeId == nil {
		return nil, errors.New("invalid value for required argument 'VolumeId'")
	}
	var resource VolumeAttachment
	err := ctx.RegisterResource("aws:ec2/volumeAttachment:VolumeAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeAttachment gets an existing VolumeAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeAttachmentState, opts ...pulumi.ResourceOption) (*VolumeAttachment, error) {
	var resource VolumeAttachment
	err := ctx.ReadResource("aws:ec2/volumeAttachment:VolumeAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeAttachment resources.
type volumeAttachmentState struct {
	// The device name to expose to the instance (for
	// example, `/dev/sdh` or `xvdh`).  See [Device Naming on Linux Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html#available-ec2-device-names) and [Device Naming on Windows Instances](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/device_naming.html#available-ec2-device-names) for more information.
	DeviceName *string `pulumi:"deviceName"`
	// Set to `true` if you want to force the
	// volume to detach. Useful if previous attempts failed, but use this option only
	// as a last resort, as this can result in **data loss**. See
	// [Detaching an Amazon EBS Volume from an Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-detaching-volume.html) for more information.
	ForceDetach *bool `pulumi:"forceDetach"`
	// ID of the Instance to attach to
	InstanceId *string `pulumi:"instanceId"`
	// Set this to true if you do not wish
	// to detach the volume from the instance to which it is attached at destroy
	// time, and instead just remove the attachment from this provider state. This is
	// useful when destroying an instance which has volumes created by some other
	// means attached.
	SkipDestroy *bool `pulumi:"skipDestroy"`
	// ID of the Volume to be attached
	VolumeId *string `pulumi:"volumeId"`
}

type VolumeAttachmentState struct {
	// The device name to expose to the instance (for
	// example, `/dev/sdh` or `xvdh`).  See [Device Naming on Linux Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html#available-ec2-device-names) and [Device Naming on Windows Instances](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/device_naming.html#available-ec2-device-names) for more information.
	DeviceName pulumi.StringPtrInput
	// Set to `true` if you want to force the
	// volume to detach. Useful if previous attempts failed, but use this option only
	// as a last resort, as this can result in **data loss**. See
	// [Detaching an Amazon EBS Volume from an Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-detaching-volume.html) for more information.
	ForceDetach pulumi.BoolPtrInput
	// ID of the Instance to attach to
	InstanceId pulumi.StringPtrInput
	// Set this to true if you do not wish
	// to detach the volume from the instance to which it is attached at destroy
	// time, and instead just remove the attachment from this provider state. This is
	// useful when destroying an instance which has volumes created by some other
	// means attached.
	SkipDestroy pulumi.BoolPtrInput
	// ID of the Volume to be attached
	VolumeId pulumi.StringPtrInput
}

func (VolumeAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachmentState)(nil)).Elem()
}

type volumeAttachmentArgs struct {
	// The device name to expose to the instance (for
	// example, `/dev/sdh` or `xvdh`).  See [Device Naming on Linux Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html#available-ec2-device-names) and [Device Naming on Windows Instances](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/device_naming.html#available-ec2-device-names) for more information.
	DeviceName string `pulumi:"deviceName"`
	// Set to `true` if you want to force the
	// volume to detach. Useful if previous attempts failed, but use this option only
	// as a last resort, as this can result in **data loss**. See
	// [Detaching an Amazon EBS Volume from an Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-detaching-volume.html) for more information.
	ForceDetach *bool `pulumi:"forceDetach"`
	// ID of the Instance to attach to
	InstanceId string `pulumi:"instanceId"`
	// Set this to true if you do not wish
	// to detach the volume from the instance to which it is attached at destroy
	// time, and instead just remove the attachment from this provider state. This is
	// useful when destroying an instance which has volumes created by some other
	// means attached.
	SkipDestroy *bool `pulumi:"skipDestroy"`
	// ID of the Volume to be attached
	VolumeId string `pulumi:"volumeId"`
}

// The set of arguments for constructing a VolumeAttachment resource.
type VolumeAttachmentArgs struct {
	// The device name to expose to the instance (for
	// example, `/dev/sdh` or `xvdh`).  See [Device Naming on Linux Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html#available-ec2-device-names) and [Device Naming on Windows Instances](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/device_naming.html#available-ec2-device-names) for more information.
	DeviceName pulumi.StringInput
	// Set to `true` if you want to force the
	// volume to detach. Useful if previous attempts failed, but use this option only
	// as a last resort, as this can result in **data loss**. See
	// [Detaching an Amazon EBS Volume from an Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-detaching-volume.html) for more information.
	ForceDetach pulumi.BoolPtrInput
	// ID of the Instance to attach to
	InstanceId pulumi.StringInput
	// Set this to true if you do not wish
	// to detach the volume from the instance to which it is attached at destroy
	// time, and instead just remove the attachment from this provider state. This is
	// useful when destroying an instance which has volumes created by some other
	// means attached.
	SkipDestroy pulumi.BoolPtrInput
	// ID of the Volume to be attached
	VolumeId pulumi.StringInput
}

func (VolumeAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachmentArgs)(nil)).Elem()
}

type VolumeAttachmentInput interface {
	pulumi.Input

	ToVolumeAttachmentOutput() VolumeAttachmentOutput
	ToVolumeAttachmentOutputWithContext(ctx context.Context) VolumeAttachmentOutput
}

type VolumeAttachmentPtrInput interface {
	pulumi.Input

	ToVolumeAttachmentPtrOutput() VolumeAttachmentPtrOutput
	ToVolumeAttachmentPtrOutputWithContext(ctx context.Context) VolumeAttachmentPtrOutput
}

func (VolumeAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachment)(nil)).Elem()
}

func (i VolumeAttachment) ToVolumeAttachmentOutput() VolumeAttachmentOutput {
	return i.ToVolumeAttachmentOutputWithContext(context.Background())
}

func (i VolumeAttachment) ToVolumeAttachmentOutputWithContext(ctx context.Context) VolumeAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentOutput)
}

func (i VolumeAttachment) ToVolumeAttachmentPtrOutput() VolumeAttachmentPtrOutput {
	return i.ToVolumeAttachmentPtrOutputWithContext(context.Background())
}

func (i VolumeAttachment) ToVolumeAttachmentPtrOutputWithContext(ctx context.Context) VolumeAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentPtrOutput)
}

type VolumeAttachmentOutput struct {
	*pulumi.OutputState
}

func (VolumeAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachmentOutput)(nil)).Elem()
}

func (o VolumeAttachmentOutput) ToVolumeAttachmentOutput() VolumeAttachmentOutput {
	return o
}

func (o VolumeAttachmentOutput) ToVolumeAttachmentOutputWithContext(ctx context.Context) VolumeAttachmentOutput {
	return o
}

type VolumeAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (VolumeAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeAttachment)(nil)).Elem()
}

func (o VolumeAttachmentPtrOutput) ToVolumeAttachmentPtrOutput() VolumeAttachmentPtrOutput {
	return o
}

func (o VolumeAttachmentPtrOutput) ToVolumeAttachmentPtrOutputWithContext(ctx context.Context) VolumeAttachmentPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VolumeAttachmentOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentPtrOutput{})
}
