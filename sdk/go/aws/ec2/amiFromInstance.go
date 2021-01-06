// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewAmiFromInstance(ctx, "example", &ec2.AmiFromInstanceArgs{
// 			SourceInstanceId: pulumi.String("i-xxxxxxxx"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AmiFromInstance struct {
	pulumi.CustomResourceState

	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringOutput `pulumi:"architecture"`
	// The ARN of the AMI.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A longer, human-readable description for the AMI.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiFromInstanceEbsBlockDeviceArrayOutput `pulumi:"ebsBlockDevices"`
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolOutput `pulumi:"enaSupport"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiFromInstanceEphemeralBlockDeviceArrayOutput `pulumi:"ephemeralBlockDevices"`
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation pulumi.StringOutput `pulumi:"imageLocation"`
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId           pulumi.StringOutput `pulumi:"kernelId"`
	ManageEbsSnapshots pulumi.BoolOutput   `pulumi:"manageEbsSnapshots"`
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
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType pulumi.StringOutput `pulumi:"virtualizationType"`
}

// NewAmiFromInstance registers a new resource with the given unique name, arguments, and options.
func NewAmiFromInstance(ctx *pulumi.Context,
	name string, args *AmiFromInstanceArgs, opts ...pulumi.ResourceOption) (*AmiFromInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SourceInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceInstanceId'")
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
	// The ARN of the AMI.
	Arn *string `pulumi:"arn"`
	// A longer, human-readable description for the AMI.
	Description *string `pulumi:"description"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices []AmiFromInstanceEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport *bool `pulumi:"enaSupport"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices []AmiFromInstanceEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
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
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType *string `pulumi:"virtualizationType"`
}

type AmiFromInstanceState struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringPtrInput
	// The ARN of the AMI.
	Arn pulumi.StringPtrInput
	// A longer, human-readable description for the AMI.
	Description pulumi.StringPtrInput
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiFromInstanceEbsBlockDeviceArrayInput
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolPtrInput
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiFromInstanceEphemeralBlockDeviceArrayInput
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
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
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
	EbsBlockDevices []AmiFromInstanceEbsBlockDevice `pulumi:"ebsBlockDevices"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices []AmiFromInstanceEphemeralBlockDevice `pulumi:"ephemeralBlockDevices"`
	// A region-unique name for the AMI.
	Name *string `pulumi:"name"`
	// Boolean that overrides the behavior of stopping
	// the instance before snapshotting. This is risky since it may cause a snapshot of an
	// inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
	// guarantees that no filesystem writes will be underway at the time of snapshot.
	SnapshotWithoutReboot *bool `pulumi:"snapshotWithoutReboot"`
	// The id of the instance to use as the basis of the AMI.
	SourceInstanceId string `pulumi:"sourceInstanceId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AmiFromInstance resource.
type AmiFromInstanceArgs struct {
	// A longer, human-readable description for the AMI.
	Description pulumi.StringPtrInput
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiFromInstanceEbsBlockDeviceArrayInput
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiFromInstanceEphemeralBlockDeviceArrayInput
	// A region-unique name for the AMI.
	Name pulumi.StringPtrInput
	// Boolean that overrides the behavior of stopping
	// the instance before snapshotting. This is risky since it may cause a snapshot of an
	// inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
	// guarantees that no filesystem writes will be underway at the time of snapshot.
	SnapshotWithoutReboot pulumi.BoolPtrInput
	// The id of the instance to use as the basis of the AMI.
	SourceInstanceId pulumi.StringInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AmiFromInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*amiFromInstanceArgs)(nil)).Elem()
}

type AmiFromInstanceInput interface {
	pulumi.Input

	ToAmiFromInstanceOutput() AmiFromInstanceOutput
	ToAmiFromInstanceOutputWithContext(ctx context.Context) AmiFromInstanceOutput
}

func (*AmiFromInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*AmiFromInstance)(nil))
}

func (i *AmiFromInstance) ToAmiFromInstanceOutput() AmiFromInstanceOutput {
	return i.ToAmiFromInstanceOutputWithContext(context.Background())
}

func (i *AmiFromInstance) ToAmiFromInstanceOutputWithContext(ctx context.Context) AmiFromInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiFromInstanceOutput)
}

func (i *AmiFromInstance) ToAmiFromInstancePtrOutput() AmiFromInstancePtrOutput {
	return i.ToAmiFromInstancePtrOutputWithContext(context.Background())
}

func (i *AmiFromInstance) ToAmiFromInstancePtrOutputWithContext(ctx context.Context) AmiFromInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiFromInstancePtrOutput)
}

type AmiFromInstancePtrInput interface {
	pulumi.Input

	ToAmiFromInstancePtrOutput() AmiFromInstancePtrOutput
	ToAmiFromInstancePtrOutputWithContext(ctx context.Context) AmiFromInstancePtrOutput
}

type amiFromInstancePtrType AmiFromInstanceArgs

func (*amiFromInstancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AmiFromInstance)(nil))
}

func (i *amiFromInstancePtrType) ToAmiFromInstancePtrOutput() AmiFromInstancePtrOutput {
	return i.ToAmiFromInstancePtrOutputWithContext(context.Background())
}

func (i *amiFromInstancePtrType) ToAmiFromInstancePtrOutputWithContext(ctx context.Context) AmiFromInstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiFromInstanceOutput).ToAmiFromInstancePtrOutput()
}

// AmiFromInstanceArrayInput is an input type that accepts AmiFromInstanceArray and AmiFromInstanceArrayOutput values.
// You can construct a concrete instance of `AmiFromInstanceArrayInput` via:
//
//          AmiFromInstanceArray{ AmiFromInstanceArgs{...} }
type AmiFromInstanceArrayInput interface {
	pulumi.Input

	ToAmiFromInstanceArrayOutput() AmiFromInstanceArrayOutput
	ToAmiFromInstanceArrayOutputWithContext(context.Context) AmiFromInstanceArrayOutput
}

type AmiFromInstanceArray []AmiFromInstanceInput

func (AmiFromInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AmiFromInstance)(nil))
}

func (i AmiFromInstanceArray) ToAmiFromInstanceArrayOutput() AmiFromInstanceArrayOutput {
	return i.ToAmiFromInstanceArrayOutputWithContext(context.Background())
}

func (i AmiFromInstanceArray) ToAmiFromInstanceArrayOutputWithContext(ctx context.Context) AmiFromInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiFromInstanceArrayOutput)
}

// AmiFromInstanceMapInput is an input type that accepts AmiFromInstanceMap and AmiFromInstanceMapOutput values.
// You can construct a concrete instance of `AmiFromInstanceMapInput` via:
//
//          AmiFromInstanceMap{ "key": AmiFromInstanceArgs{...} }
type AmiFromInstanceMapInput interface {
	pulumi.Input

	ToAmiFromInstanceMapOutput() AmiFromInstanceMapOutput
	ToAmiFromInstanceMapOutputWithContext(context.Context) AmiFromInstanceMapOutput
}

type AmiFromInstanceMap map[string]AmiFromInstanceInput

func (AmiFromInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AmiFromInstance)(nil))
}

func (i AmiFromInstanceMap) ToAmiFromInstanceMapOutput() AmiFromInstanceMapOutput {
	return i.ToAmiFromInstanceMapOutputWithContext(context.Background())
}

func (i AmiFromInstanceMap) ToAmiFromInstanceMapOutputWithContext(ctx context.Context) AmiFromInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiFromInstanceMapOutput)
}

type AmiFromInstanceOutput struct {
	*pulumi.OutputState
}

func (AmiFromInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmiFromInstance)(nil))
}

func (o AmiFromInstanceOutput) ToAmiFromInstanceOutput() AmiFromInstanceOutput {
	return o
}

func (o AmiFromInstanceOutput) ToAmiFromInstanceOutputWithContext(ctx context.Context) AmiFromInstanceOutput {
	return o
}

func (o AmiFromInstanceOutput) ToAmiFromInstancePtrOutput() AmiFromInstancePtrOutput {
	return o.ToAmiFromInstancePtrOutputWithContext(context.Background())
}

func (o AmiFromInstanceOutput) ToAmiFromInstancePtrOutputWithContext(ctx context.Context) AmiFromInstancePtrOutput {
	return o.ApplyT(func(v AmiFromInstance) *AmiFromInstance {
		return &v
	}).(AmiFromInstancePtrOutput)
}

type AmiFromInstancePtrOutput struct {
	*pulumi.OutputState
}

func (AmiFromInstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AmiFromInstance)(nil))
}

func (o AmiFromInstancePtrOutput) ToAmiFromInstancePtrOutput() AmiFromInstancePtrOutput {
	return o
}

func (o AmiFromInstancePtrOutput) ToAmiFromInstancePtrOutputWithContext(ctx context.Context) AmiFromInstancePtrOutput {
	return o
}

type AmiFromInstanceArrayOutput struct{ *pulumi.OutputState }

func (AmiFromInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AmiFromInstance)(nil))
}

func (o AmiFromInstanceArrayOutput) ToAmiFromInstanceArrayOutput() AmiFromInstanceArrayOutput {
	return o
}

func (o AmiFromInstanceArrayOutput) ToAmiFromInstanceArrayOutputWithContext(ctx context.Context) AmiFromInstanceArrayOutput {
	return o
}

func (o AmiFromInstanceArrayOutput) Index(i pulumi.IntInput) AmiFromInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AmiFromInstance {
		return vs[0].([]AmiFromInstance)[vs[1].(int)]
	}).(AmiFromInstanceOutput)
}

type AmiFromInstanceMapOutput struct{ *pulumi.OutputState }

func (AmiFromInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AmiFromInstance)(nil))
}

func (o AmiFromInstanceMapOutput) ToAmiFromInstanceMapOutput() AmiFromInstanceMapOutput {
	return o
}

func (o AmiFromInstanceMapOutput) ToAmiFromInstanceMapOutputWithContext(ctx context.Context) AmiFromInstanceMapOutput {
	return o
}

func (o AmiFromInstanceMapOutput) MapIndex(k pulumi.StringInput) AmiFromInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AmiFromInstance {
		return vs[0].(map[string]AmiFromInstance)[vs[1].(string)]
	}).(AmiFromInstanceOutput)
}

func init() {
	pulumi.RegisterOutputType(AmiFromInstanceOutput{})
	pulumi.RegisterOutputType(AmiFromInstancePtrOutput{})
	pulumi.RegisterOutputType(AmiFromInstanceArrayOutput{})
	pulumi.RegisterOutputType(AmiFromInstanceMapOutput{})
}
