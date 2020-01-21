// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package amiFromInstance

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/ec2/AmiFromInstanceEbsBlockDevice"
	"https:/github.com/pulumi/pulumi-aws/ec2/AmiFromInstanceEphemeralBlockDevice"
)

// The "AMI from instance" resource allows the creation of an Amazon Machine
// Image (AMI) modelled after an existing EBS-backed EC2 instance.
// 
// The created AMI will refer to implicitly-created snapshots of the instance's
// EBS volumes and mimick its assigned block device configuration at the time
// the resource is created.
// 
// This resource is best applied to an instance that is stopped when this instance
// is created, so that the contents of the created image are predictable. When
// applied to an instance that is running, *the instance will be stopped before taking
// the snapshots and then started back up again*, resulting in a period of
// downtime.
// 
// Note that the source instance is inspected only at the initial creation of this
// resource. Ongoing updates to the referenced instance will not be propagated into
// the generated AMI. Users may taint or otherwise recreate the resource in order
// to produce a fresh snapshot.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ami_from_instance.html.markdown.
type AmiFromInstance struct {
	pulumi.CustomResourceState

	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringOutput `pulumi:"architecture"`
	// A longer, human-readable description for the AMI.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices ec2AmiFromInstanceEbsBlockDevice.AmiFromInstanceEbsBlockDeviceArrayOutput `pulumi:"ebsBlockDevices"`
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolOutput `pulumi:"enaSupport"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices ec2AmiFromInstanceEphemeralBlockDevice.AmiFromInstanceEphemeralBlockDeviceArrayOutput `pulumi:"ephemeralBlockDevices"`
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation pulumi.StringOutput `pulumi:"imageLocation"`
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId pulumi.StringOutput `pulumi:"kernelId"`
	ManageEbsSnapshots pulumi.BoolOutput `pulumi:"manageEbsSnapshots"`
	// A region-unique name for the AMI.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId pulumi.StringOutput `pulumi:"ramdiskId"`
	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName pulumi.StringOutput `pulumi:"rootDeviceName"`
	RootSnapshotId pulumi.StringOutput `pulumi:"rootSnapshotId"`
	// Boolean that overrides the behavior of stopping
	// the instance before snapshotting. This is risky since it may cause a snapshot of an
	// inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
	// guarantees that no filesystem writes will be underway at the time of snapshot.
	SnapshotWithoutReboot pulumi.BoolPtrOutput `pulumi:"snapshotWithoutReboot"`
	// The id of the instance to use as the basis of the AMI.
	SourceInstanceId pulumi.StringOutput `pulumi:"sourceInstanceId"`
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport pulumi.StringOutput `pulumi:"sriovNetSupport"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType pulumi.StringOutput `pulumi:"virtualizationType"`
}

// NewAmiFromInstance registers a new resource with the given unique name, arguments, and options.
func NewAmiFromInstance(ctx *pulumi.Context,
	name string, args *AmiFromInstanceArgs, opts ...pulumi.ResourceOption) (*AmiFromInstance, error) {
	if args == nil || args.SourceInstanceId == nil {
		return nil, errors.New("missing required argument 'SourceInstanceId'")
	}
	if args == nil {
		args = &AmiFromInstanceArgs{}
	}
	var resource AmiFromInstance
	err := ctx.RegisterResource("aws:ec2/amiFromInstance:AmiFromInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAmiFromInstance gets an existing AmiFromInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAmiFromInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AmiFromInstanceState, opts ...pulumi.ResourceOption) (*AmiFromInstance, error) {
	var resource AmiFromInstance
	err := ctx.ReadResource("aws:ec2/amiFromInstance:AmiFromInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AmiFromInstance resources.
type amiFromInstanceState struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture *string `pulumi:"architecture"`
	// A longer, human-readable description for the AMI.
	Description *string `pulumi:"description"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices []ec2AmiFromInstanceEbsBlockDevice.AmiFromInstanceEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport *bool `pulumi:"enaSupport"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices []ec2AmiFromInstanceEphemeralBlockDevice.AmiFromInstanceEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation *string `pulumi:"imageLocation"`
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId *string `pulumi:"kernelId"`
	ManageEbsSnapshots *bool `pulumi:"manageEbsSnapshots"`
	// A region-unique name for the AMI.
	Name *string `pulumi:"name"`
	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId *string `pulumi:"ramdiskId"`
	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName *string `pulumi:"rootDeviceName"`
	RootSnapshotId *string `pulumi:"rootSnapshotId"`
	// Boolean that overrides the behavior of stopping
	// the instance before snapshotting. This is risky since it may cause a snapshot of an
	// inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
	// guarantees that no filesystem writes will be underway at the time of snapshot.
	SnapshotWithoutReboot *bool `pulumi:"snapshotWithoutReboot"`
	// The id of the instance to use as the basis of the AMI.
	SourceInstanceId *string `pulumi:"sourceInstanceId"`
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport *string `pulumi:"sriovNetSupport"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType *string `pulumi:"virtualizationType"`
}

type AmiFromInstanceState struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringPtrInput
	// A longer, human-readable description for the AMI.
	Description pulumi.StringPtrInput
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices ec2AmiFromInstanceEbsBlockDevice.AmiFromInstanceEbsBlockDeviceArrayInput
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolPtrInput
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices ec2AmiFromInstanceEphemeralBlockDevice.AmiFromInstanceEphemeralBlockDeviceArrayInput
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation pulumi.StringPtrInput
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId pulumi.StringPtrInput
	ManageEbsSnapshots pulumi.BoolPtrInput
	// A region-unique name for the AMI.
	Name pulumi.StringPtrInput
	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId pulumi.StringPtrInput
	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName pulumi.StringPtrInput
	RootSnapshotId pulumi.StringPtrInput
	// Boolean that overrides the behavior of stopping
	// the instance before snapshotting. This is risky since it may cause a snapshot of an
	// inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
	// guarantees that no filesystem writes will be underway at the time of snapshot.
	SnapshotWithoutReboot pulumi.BoolPtrInput
	// The id of the instance to use as the basis of the AMI.
	SourceInstanceId pulumi.StringPtrInput
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType pulumi.StringPtrInput
}

func (AmiFromInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*amiFromInstanceState)(nil)).Elem()
}

type amiFromInstanceArgs struct {
	// A longer, human-readable description for the AMI.
	Description *string `pulumi:"description"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices []ec2AmiFromInstanceEbsBlockDevice.AmiFromInstanceEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices []ec2AmiFromInstanceEphemeralBlockDevice.AmiFromInstanceEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	// A region-unique name for the AMI.
	Name *string `pulumi:"name"`
	// Boolean that overrides the behavior of stopping
	// the instance before snapshotting. This is risky since it may cause a snapshot of an
	// inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
	// guarantees that no filesystem writes will be underway at the time of snapshot.
	SnapshotWithoutReboot *bool `pulumi:"snapshotWithoutReboot"`
	// The id of the instance to use as the basis of the AMI.
	SourceInstanceId string `pulumi:"sourceInstanceId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a AmiFromInstance resource.
type AmiFromInstanceArgs struct {
	// A longer, human-readable description for the AMI.
	Description pulumi.StringPtrInput
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices ec2AmiFromInstanceEbsBlockDevice.AmiFromInstanceEbsBlockDeviceArrayInput
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices ec2AmiFromInstanceEphemeralBlockDevice.AmiFromInstanceEphemeralBlockDeviceArrayInput
	// A region-unique name for the AMI.
	Name pulumi.StringPtrInput
	// Boolean that overrides the behavior of stopping
	// the instance before snapshotting. This is risky since it may cause a snapshot of an
	// inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
	// guarantees that no filesystem writes will be underway at the time of snapshot.
	SnapshotWithoutReboot pulumi.BoolPtrInput
	// The id of the instance to use as the basis of the AMI.
	SourceInstanceId pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (AmiFromInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*amiFromInstanceArgs)(nil)).Elem()
}

