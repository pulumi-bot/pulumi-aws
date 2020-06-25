// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The AMI resource allows the creation and management of a completely-custom
// *Amazon Machine Image* (AMI).
//
// If you just want to duplicate an existing AMI, possibly copying it to another
// region, it's better to use `ec2.AmiCopy` instead.
//
// If you just want to share an existing AMI with another AWS account,
// it's better to use `ec2.AmiLaunchPermission` instead.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = ec2.NewAmi(ctx, "example", &ec2.AmiArgs{
// 			EbsBlockDevices: ec2.AmiEbsBlockDeviceArray{
// 				&ec2.AmiEbsBlockDeviceArgs{
// 					DeviceName: pulumi.String("/dev/xvda"),
// 					SnapshotId: pulumi.String("snap-xxxxxxxx"),
// 					VolumeSize: pulumi.Int(8),
// 				},
// 			},
// 			RootDeviceName:     pulumi.String("/dev/xvda"),
// 			VirtualizationType: pulumi.String("hvm"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Ami struct {
	pulumi.CustomResourceState

	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringPtrOutput `pulumi:"architecture"`
	// A longer, human-readable description for the AMI.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiEbsBlockDeviceArrayOutput `pulumi:"ebsBlockDevices"`
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolPtrOutput `pulumi:"enaSupport"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiEphemeralBlockDeviceArrayOutput `pulumi:"ephemeralBlockDevices"`
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation pulumi.StringOutput `pulumi:"imageLocation"`
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId           pulumi.StringPtrOutput `pulumi:"kernelId"`
	ManageEbsSnapshots pulumi.BoolOutput      `pulumi:"manageEbsSnapshots"`
	// A region-unique name for the AMI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId pulumi.StringPtrOutput `pulumi:"ramdiskId"`
	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName pulumi.StringPtrOutput `pulumi:"rootDeviceName"`
	// The Snapshot ID for the root volume (for EBS-backed AMIs)
	RootSnapshotId pulumi.StringOutput `pulumi:"rootSnapshotId"`
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport pulumi.StringPtrOutput `pulumi:"sriovNetSupport"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType pulumi.StringPtrOutput `pulumi:"virtualizationType"`
}

// NewAmi registers a new resource with the given unique name, arguments, and options.
func NewAmi(ctx *pulumi.Context,
	name string, args *AmiArgs, opts ...pulumi.ResourceOption) (*Ami, error) {
	if args == nil {
		args = &AmiArgs{}
	}
	var resource Ami
	err := ctx.RegisterResource("aws:ec2/ami:Ami", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAmi gets an existing Ami resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAmi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AmiState, opts ...pulumi.ResourceOption) (*Ami, error) {
	var resource Ami
	err := ctx.ReadResource("aws:ec2/ami:Ami", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ami resources.
type amiState struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture *string `pulumi:"architecture"`
	// A longer, human-readable description for the AMI.
	Description *string `pulumi:"description"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices []AmiEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport *bool `pulumi:"enaSupport"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices []AmiEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation *string `pulumi:"imageLocation"`
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId           *string `pulumi:"kernelId"`
	ManageEbsSnapshots *bool   `pulumi:"manageEbsSnapshots"`
	// A region-unique name for the AMI.
	Name *string `pulumi:"name"`
	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId *string `pulumi:"ramdiskId"`
	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName *string `pulumi:"rootDeviceName"`
	// The Snapshot ID for the root volume (for EBS-backed AMIs)
	RootSnapshotId *string `pulumi:"rootSnapshotId"`
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport *string `pulumi:"sriovNetSupport"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType *string `pulumi:"virtualizationType"`
}

type AmiState struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringPtrInput
	// A longer, human-readable description for the AMI.
	Description pulumi.StringPtrInput
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiEbsBlockDeviceArrayInput
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolPtrInput
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiEphemeralBlockDeviceArrayInput
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation pulumi.StringPtrInput
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId           pulumi.StringPtrInput
	ManageEbsSnapshots pulumi.BoolPtrInput
	// A region-unique name for the AMI.
	Name pulumi.StringPtrInput
	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId pulumi.StringPtrInput
	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName pulumi.StringPtrInput
	// The Snapshot ID for the root volume (for EBS-backed AMIs)
	RootSnapshotId pulumi.StringPtrInput
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType pulumi.StringPtrInput
}

func (AmiState) ElementType() reflect.Type {
	return reflect.TypeOf((*amiState)(nil)).Elem()
}

type amiArgs struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture *string `pulumi:"architecture"`
	// A longer, human-readable description for the AMI.
	Description *string `pulumi:"description"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices []AmiEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport *bool `pulumi:"enaSupport"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices []AmiEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation *string `pulumi:"imageLocation"`
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId *string `pulumi:"kernelId"`
	// A region-unique name for the AMI.
	Name *string `pulumi:"name"`
	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId *string `pulumi:"ramdiskId"`
	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName *string `pulumi:"rootDeviceName"`
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport *string `pulumi:"sriovNetSupport"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType *string `pulumi:"virtualizationType"`
}

// The set of arguments for constructing a Ami resource.
type AmiArgs struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringPtrInput
	// A longer, human-readable description for the AMI.
	Description pulumi.StringPtrInput
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiEbsBlockDeviceArrayInput
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolPtrInput
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiEphemeralBlockDeviceArrayInput
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation pulumi.StringPtrInput
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId pulumi.StringPtrInput
	// A region-unique name for the AMI.
	Name pulumi.StringPtrInput
	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId pulumi.StringPtrInput
	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName pulumi.StringPtrInput
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType pulumi.StringPtrInput
}

func (AmiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*amiArgs)(nil)).Elem()
}
