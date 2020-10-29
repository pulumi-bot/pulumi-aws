// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an AWS Storage Gateway stored iSCSI volume.
//
// > **NOTE:** The gateway must have a working storage added (e.g. via the [`storagegateway.WorkingStorage`](https://www.terraform.io/docs/providers/aws/r/storagegateway_working_storage.html) resource) before the volume is operational to clients, however the Storage Gateway API will allow volume creation without error in that case and return volume status as `WORKING STORAGE NOT CONFIGURED`.
//
// ## Example Usage
type StoredIscsiVolume struct {
	pulumi.CustomResourceState

	// Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether mutual CHAP is enabled for the iSCSI target.
	ChapEnabled pulumi.BoolOutput `pulumi:"chapEnabled"`
	// The unique identifier for the gateway local disk that is configured as a stored volume.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringOutput `pulumi:"gatewayArn"`
	// `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Optional.
	KmsEncrypted pulumi.BoolPtrOutput `pulumi:"kmsEncrypted"`
	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is `true`.
	KmsKey pulumi.StringPtrOutput `pulumi:"kmsKey"`
	// Logical disk number.
	LunNumber pulumi.IntOutput `pulumi:"lunNumber"`
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The port used to communicate with iSCSI targets.
	NetworkInterfacePort pulumi.IntOutput `pulumi:"networkInterfacePort"`
	// Specify this field as `true` if you want to preserve the data on the local disk. Otherwise, specifying this field as false creates an empty volume.
	PreserveExistingData pulumi.BoolOutput `pulumi:"preserveExistingData"`
	// The snapshot ID of the snapshot to restore as the new stored volume. e.g. `snap-1122aabb`.
	SnapshotId pulumi.StringPtrOutput `pulumi:"snapshotId"`
	// Key-value mapping of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Target Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
	TargetArn pulumi.StringOutput `pulumi:"targetArn"`
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName pulumi.StringOutput `pulumi:"targetName"`
	// A value that indicates whether a storage volume is attached to, detached from, or is in the process of detaching from a gateway.
	VolumeAttachmentStatus pulumi.StringOutput `pulumi:"volumeAttachmentStatus"`
	// Volume ID, e.g. `vol-12345678`.
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
	// The size of the data stored on the volume in bytes.
	VolumeSizeInBytes pulumi.IntOutput `pulumi:"volumeSizeInBytes"`
	// indicates the state of the storage volume.
	VolumeStatus pulumi.StringOutput `pulumi:"volumeStatus"`
	// indicates the type of the volume.
	VolumeType pulumi.StringOutput `pulumi:"volumeType"`
}

// NewStoredIscsiVolume registers a new resource with the given unique name, arguments, and options.
func NewStoredIscsiVolume(ctx *pulumi.Context,
	name string, args *StoredIscsiVolumeArgs, opts ...pulumi.ResourceOption) (*StoredIscsiVolume, error) {
	if args == nil || args.DiskId == nil {
		return nil, errors.New("missing required argument 'DiskId'")
	}
	if args == nil || args.GatewayArn == nil {
		return nil, errors.New("missing required argument 'GatewayArn'")
	}
	if args == nil || args.NetworkInterfaceId == nil {
		return nil, errors.New("missing required argument 'NetworkInterfaceId'")
	}
	if args == nil || args.PreserveExistingData == nil {
		return nil, errors.New("missing required argument 'PreserveExistingData'")
	}
	if args == nil || args.TargetName == nil {
		return nil, errors.New("missing required argument 'TargetName'")
	}
	if args == nil {
		args = &StoredIscsiVolumeArgs{}
	}
	var resource StoredIscsiVolume
	err := ctx.RegisterResource("aws:storagegateway/storedIscsiVolume:StoredIscsiVolume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStoredIscsiVolume gets an existing StoredIscsiVolume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStoredIscsiVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StoredIscsiVolumeState, opts ...pulumi.ResourceOption) (*StoredIscsiVolume, error) {
	var resource StoredIscsiVolume
	err := ctx.ReadResource("aws:storagegateway/storedIscsiVolume:StoredIscsiVolume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StoredIscsiVolume resources.
type storedIscsiVolumeState struct {
	// Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	Arn *string `pulumi:"arn"`
	// Whether mutual CHAP is enabled for the iSCSI target.
	ChapEnabled *bool `pulumi:"chapEnabled"`
	// The unique identifier for the gateway local disk that is configured as a stored volume.
	DiskId *string `pulumi:"diskId"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn *string `pulumi:"gatewayArn"`
	// `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Optional.
	KmsEncrypted *bool `pulumi:"kmsEncrypted"`
	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is `true`.
	KmsKey *string `pulumi:"kmsKey"`
	// Logical disk number.
	LunNumber *int `pulumi:"lunNumber"`
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The port used to communicate with iSCSI targets.
	NetworkInterfacePort *int `pulumi:"networkInterfacePort"`
	// Specify this field as `true` if you want to preserve the data on the local disk. Otherwise, specifying this field as false creates an empty volume.
	PreserveExistingData *bool `pulumi:"preserveExistingData"`
	// The snapshot ID of the snapshot to restore as the new stored volume. e.g. `snap-1122aabb`.
	SnapshotId *string `pulumi:"snapshotId"`
	// Key-value mapping of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Target Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
	TargetArn *string `pulumi:"targetArn"`
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName *string `pulumi:"targetName"`
	// A value that indicates whether a storage volume is attached to, detached from, or is in the process of detaching from a gateway.
	VolumeAttachmentStatus *string `pulumi:"volumeAttachmentStatus"`
	// Volume ID, e.g. `vol-12345678`.
	VolumeId *string `pulumi:"volumeId"`
	// The size of the data stored on the volume in bytes.
	VolumeSizeInBytes *int `pulumi:"volumeSizeInBytes"`
	// indicates the state of the storage volume.
	VolumeStatus *string `pulumi:"volumeStatus"`
	// indicates the type of the volume.
	VolumeType *string `pulumi:"volumeType"`
}

type StoredIscsiVolumeState struct {
	// Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	Arn pulumi.StringPtrInput
	// Whether mutual CHAP is enabled for the iSCSI target.
	ChapEnabled pulumi.BoolPtrInput
	// The unique identifier for the gateway local disk that is configured as a stored volume.
	DiskId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringPtrInput
	// `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Optional.
	KmsEncrypted pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is `true`.
	KmsKey pulumi.StringPtrInput
	// Logical disk number.
	LunNumber pulumi.IntPtrInput
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId pulumi.StringPtrInput
	// The port used to communicate with iSCSI targets.
	NetworkInterfacePort pulumi.IntPtrInput
	// Specify this field as `true` if you want to preserve the data on the local disk. Otherwise, specifying this field as false creates an empty volume.
	PreserveExistingData pulumi.BoolPtrInput
	// The snapshot ID of the snapshot to restore as the new stored volume. e.g. `snap-1122aabb`.
	SnapshotId pulumi.StringPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.StringMapInput
	// Target Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
	TargetArn pulumi.StringPtrInput
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName pulumi.StringPtrInput
	// A value that indicates whether a storage volume is attached to, detached from, or is in the process of detaching from a gateway.
	VolumeAttachmentStatus pulumi.StringPtrInput
	// Volume ID, e.g. `vol-12345678`.
	VolumeId pulumi.StringPtrInput
	// The size of the data stored on the volume in bytes.
	VolumeSizeInBytes pulumi.IntPtrInput
	// indicates the state of the storage volume.
	VolumeStatus pulumi.StringPtrInput
	// indicates the type of the volume.
	VolumeType pulumi.StringPtrInput
}

func (StoredIscsiVolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*storedIscsiVolumeState)(nil)).Elem()
}

type storedIscsiVolumeArgs struct {
	// The unique identifier for the gateway local disk that is configured as a stored volume.
	DiskId string `pulumi:"diskId"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn string `pulumi:"gatewayArn"`
	// `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Optional.
	KmsEncrypted *bool `pulumi:"kmsEncrypted"`
	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is `true`.
	KmsKey *string `pulumi:"kmsKey"`
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// Specify this field as `true` if you want to preserve the data on the local disk. Otherwise, specifying this field as false creates an empty volume.
	PreserveExistingData bool `pulumi:"preserveExistingData"`
	// The snapshot ID of the snapshot to restore as the new stored volume. e.g. `snap-1122aabb`.
	SnapshotId *string `pulumi:"snapshotId"`
	// Key-value mapping of resource tags
	Tags map[string]string `pulumi:"tags"`
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName string `pulumi:"targetName"`
}

// The set of arguments for constructing a StoredIscsiVolume resource.
type StoredIscsiVolumeArgs struct {
	// The unique identifier for the gateway local disk that is configured as a stored volume.
	DiskId pulumi.StringInput
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringInput
	// `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Optional.
	KmsEncrypted pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. This value can only be set when `kmsEncrypted` is `true`.
	KmsKey pulumi.StringPtrInput
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId pulumi.StringInput
	// Specify this field as `true` if you want to preserve the data on the local disk. Otherwise, specifying this field as false creates an empty volume.
	PreserveExistingData pulumi.BoolInput
	// The snapshot ID of the snapshot to restore as the new stored volume. e.g. `snap-1122aabb`.
	SnapshotId pulumi.StringPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.StringMapInput
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName pulumi.StringInput
}

func (StoredIscsiVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storedIscsiVolumeArgs)(nil)).Elem()
}
