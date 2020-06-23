// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type EfsLocationEc2Config struct {
	// List of Amazon Resource Names (ARNs) of the EC2 Security Groups that are associated with the EFS Mount Target.
	SecurityGroupArns []string `pulumi:"securityGroupArns"`
	// Amazon Resource Name (ARN) of the EC2 Subnet that is associated with the EFS Mount Target.
	SubnetArn string `pulumi:"subnetArn"`
}

// EfsLocationEc2ConfigInput is an input type that accepts EfsLocationEc2ConfigArgs and EfsLocationEc2ConfigOutput values.
// You can construct a concrete instance of `EfsLocationEc2ConfigInput` via:
//
//          EfsLocationEc2ConfigArgs{...}
type EfsLocationEc2ConfigInput interface {
	pulumi.Input

	ToEfsLocationEc2ConfigOutput() EfsLocationEc2ConfigOutput
	ToEfsLocationEc2ConfigOutputWithContext(context.Context) EfsLocationEc2ConfigOutput
}

type EfsLocationEc2ConfigArgs struct {
	// List of Amazon Resource Names (ARNs) of the EC2 Security Groups that are associated with the EFS Mount Target.
	SecurityGroupArns pulumi.StringArrayInput `pulumi:"securityGroupArns"`
	// Amazon Resource Name (ARN) of the EC2 Subnet that is associated with the EFS Mount Target.
	SubnetArn pulumi.StringInput `pulumi:"subnetArn"`
}

func (EfsLocationEc2ConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EfsLocationEc2Config)(nil)).Elem()
}

func (i EfsLocationEc2ConfigArgs) ToEfsLocationEc2ConfigOutput() EfsLocationEc2ConfigOutput {
	return i.ToEfsLocationEc2ConfigOutputWithContext(context.Background())
}

func (i EfsLocationEc2ConfigArgs) ToEfsLocationEc2ConfigOutputWithContext(ctx context.Context) EfsLocationEc2ConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EfsLocationEc2ConfigOutput)
}

func (i EfsLocationEc2ConfigArgs) ToEfsLocationEc2ConfigPtrOutput() EfsLocationEc2ConfigPtrOutput {
	return i.ToEfsLocationEc2ConfigPtrOutputWithContext(context.Background())
}

func (i EfsLocationEc2ConfigArgs) ToEfsLocationEc2ConfigPtrOutputWithContext(ctx context.Context) EfsLocationEc2ConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EfsLocationEc2ConfigOutput).ToEfsLocationEc2ConfigPtrOutputWithContext(ctx)
}

// EfsLocationEc2ConfigPtrInput is an input type that accepts EfsLocationEc2ConfigArgs, EfsLocationEc2ConfigPtr and EfsLocationEc2ConfigPtrOutput values.
// You can construct a concrete instance of `EfsLocationEc2ConfigPtrInput` via:
//
//          EfsLocationEc2ConfigArgs{...}
//
//  or:
//
//          nil
type EfsLocationEc2ConfigPtrInput interface {
	pulumi.Input

	ToEfsLocationEc2ConfigPtrOutput() EfsLocationEc2ConfigPtrOutput
	ToEfsLocationEc2ConfigPtrOutputWithContext(context.Context) EfsLocationEc2ConfigPtrOutput
}

type efsLocationEc2ConfigPtrType EfsLocationEc2ConfigArgs

func EfsLocationEc2ConfigPtr(v *EfsLocationEc2ConfigArgs) EfsLocationEc2ConfigPtrInput {
	return (*efsLocationEc2ConfigPtrType)(v)
}

func (*efsLocationEc2ConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EfsLocationEc2Config)(nil)).Elem()
}

func (i *efsLocationEc2ConfigPtrType) ToEfsLocationEc2ConfigPtrOutput() EfsLocationEc2ConfigPtrOutput {
	return i.ToEfsLocationEc2ConfigPtrOutputWithContext(context.Background())
}

func (i *efsLocationEc2ConfigPtrType) ToEfsLocationEc2ConfigPtrOutputWithContext(ctx context.Context) EfsLocationEc2ConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EfsLocationEc2ConfigPtrOutput)
}

type EfsLocationEc2ConfigOutput struct{ *pulumi.OutputState }

func (EfsLocationEc2ConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EfsLocationEc2Config)(nil)).Elem()
}

func (o EfsLocationEc2ConfigOutput) ToEfsLocationEc2ConfigOutput() EfsLocationEc2ConfigOutput {
	return o
}

func (o EfsLocationEc2ConfigOutput) ToEfsLocationEc2ConfigOutputWithContext(ctx context.Context) EfsLocationEc2ConfigOutput {
	return o
}

func (o EfsLocationEc2ConfigOutput) ToEfsLocationEc2ConfigPtrOutput() EfsLocationEc2ConfigPtrOutput {
	return o.ToEfsLocationEc2ConfigPtrOutputWithContext(context.Background())
}

func (o EfsLocationEc2ConfigOutput) ToEfsLocationEc2ConfigPtrOutputWithContext(ctx context.Context) EfsLocationEc2ConfigPtrOutput {
	return o.ApplyT(func(v EfsLocationEc2Config) *EfsLocationEc2Config {
		return &v
	}).(EfsLocationEc2ConfigPtrOutput)
}

// List of Amazon Resource Names (ARNs) of the EC2 Security Groups that are associated with the EFS Mount Target.
func (o EfsLocationEc2ConfigOutput) SecurityGroupArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EfsLocationEc2Config) []string { return v.SecurityGroupArns }).(pulumi.StringArrayOutput)
}

// Amazon Resource Name (ARN) of the EC2 Subnet that is associated with the EFS Mount Target.
func (o EfsLocationEc2ConfigOutput) SubnetArn() pulumi.StringOutput {
	return o.ApplyT(func(v EfsLocationEc2Config) string { return v.SubnetArn }).(pulumi.StringOutput)
}

type EfsLocationEc2ConfigPtrOutput struct{ *pulumi.OutputState }

func (EfsLocationEc2ConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EfsLocationEc2Config)(nil)).Elem()
}

func (o EfsLocationEc2ConfigPtrOutput) ToEfsLocationEc2ConfigPtrOutput() EfsLocationEc2ConfigPtrOutput {
	return o
}

func (o EfsLocationEc2ConfigPtrOutput) ToEfsLocationEc2ConfigPtrOutputWithContext(ctx context.Context) EfsLocationEc2ConfigPtrOutput {
	return o
}

func (o EfsLocationEc2ConfigPtrOutput) Elem() EfsLocationEc2ConfigOutput {
	return o.ApplyT(func(v *EfsLocationEc2Config) EfsLocationEc2Config { return *v }).(EfsLocationEc2ConfigOutput)
}

// List of Amazon Resource Names (ARNs) of the EC2 Security Groups that are associated with the EFS Mount Target.
func (o EfsLocationEc2ConfigPtrOutput) SecurityGroupArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EfsLocationEc2Config) []string {
		if v == nil {
			return nil
		}
		return v.SecurityGroupArns
	}).(pulumi.StringArrayOutput)
}

// Amazon Resource Name (ARN) of the EC2 Subnet that is associated with the EFS Mount Target.
func (o EfsLocationEc2ConfigPtrOutput) SubnetArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EfsLocationEc2Config) *string {
		if v == nil {
			return nil
		}
		return &v.SubnetArn
	}).(pulumi.StringPtrOutput)
}

type LocationSmbMountOptions struct {
	// The specific SMB version that you want DataSync to use for mounting your SMB share. Valid values: `AUTOMATIC`, `SMB2`, and `SMB3`. Default: `AUTOMATIC`
	Version *string `pulumi:"version"`
}

// LocationSmbMountOptionsInput is an input type that accepts LocationSmbMountOptionsArgs and LocationSmbMountOptionsOutput values.
// You can construct a concrete instance of `LocationSmbMountOptionsInput` via:
//
//          LocationSmbMountOptionsArgs{...}
type LocationSmbMountOptionsInput interface {
	pulumi.Input

	ToLocationSmbMountOptionsOutput() LocationSmbMountOptionsOutput
	ToLocationSmbMountOptionsOutputWithContext(context.Context) LocationSmbMountOptionsOutput
}

type LocationSmbMountOptionsArgs struct {
	// The specific SMB version that you want DataSync to use for mounting your SMB share. Valid values: `AUTOMATIC`, `SMB2`, and `SMB3`. Default: `AUTOMATIC`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (LocationSmbMountOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationSmbMountOptions)(nil)).Elem()
}

func (i LocationSmbMountOptionsArgs) ToLocationSmbMountOptionsOutput() LocationSmbMountOptionsOutput {
	return i.ToLocationSmbMountOptionsOutputWithContext(context.Background())
}

func (i LocationSmbMountOptionsArgs) ToLocationSmbMountOptionsOutputWithContext(ctx context.Context) LocationSmbMountOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationSmbMountOptionsOutput)
}

func (i LocationSmbMountOptionsArgs) ToLocationSmbMountOptionsPtrOutput() LocationSmbMountOptionsPtrOutput {
	return i.ToLocationSmbMountOptionsPtrOutputWithContext(context.Background())
}

func (i LocationSmbMountOptionsArgs) ToLocationSmbMountOptionsPtrOutputWithContext(ctx context.Context) LocationSmbMountOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationSmbMountOptionsOutput).ToLocationSmbMountOptionsPtrOutputWithContext(ctx)
}

// LocationSmbMountOptionsPtrInput is an input type that accepts LocationSmbMountOptionsArgs, LocationSmbMountOptionsPtr and LocationSmbMountOptionsPtrOutput values.
// You can construct a concrete instance of `LocationSmbMountOptionsPtrInput` via:
//
//          LocationSmbMountOptionsArgs{...}
//
//  or:
//
//          nil
type LocationSmbMountOptionsPtrInput interface {
	pulumi.Input

	ToLocationSmbMountOptionsPtrOutput() LocationSmbMountOptionsPtrOutput
	ToLocationSmbMountOptionsPtrOutputWithContext(context.Context) LocationSmbMountOptionsPtrOutput
}

type locationSmbMountOptionsPtrType LocationSmbMountOptionsArgs

func LocationSmbMountOptionsPtr(v *LocationSmbMountOptionsArgs) LocationSmbMountOptionsPtrInput {
	return (*locationSmbMountOptionsPtrType)(v)
}

func (*locationSmbMountOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationSmbMountOptions)(nil)).Elem()
}

func (i *locationSmbMountOptionsPtrType) ToLocationSmbMountOptionsPtrOutput() LocationSmbMountOptionsPtrOutput {
	return i.ToLocationSmbMountOptionsPtrOutputWithContext(context.Background())
}

func (i *locationSmbMountOptionsPtrType) ToLocationSmbMountOptionsPtrOutputWithContext(ctx context.Context) LocationSmbMountOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationSmbMountOptionsPtrOutput)
}

type LocationSmbMountOptionsOutput struct{ *pulumi.OutputState }

func (LocationSmbMountOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationSmbMountOptions)(nil)).Elem()
}

func (o LocationSmbMountOptionsOutput) ToLocationSmbMountOptionsOutput() LocationSmbMountOptionsOutput {
	return o
}

func (o LocationSmbMountOptionsOutput) ToLocationSmbMountOptionsOutputWithContext(ctx context.Context) LocationSmbMountOptionsOutput {
	return o
}

func (o LocationSmbMountOptionsOutput) ToLocationSmbMountOptionsPtrOutput() LocationSmbMountOptionsPtrOutput {
	return o.ToLocationSmbMountOptionsPtrOutputWithContext(context.Background())
}

func (o LocationSmbMountOptionsOutput) ToLocationSmbMountOptionsPtrOutputWithContext(ctx context.Context) LocationSmbMountOptionsPtrOutput {
	return o.ApplyT(func(v LocationSmbMountOptions) *LocationSmbMountOptions {
		return &v
	}).(LocationSmbMountOptionsPtrOutput)
}

// The specific SMB version that you want DataSync to use for mounting your SMB share. Valid values: `AUTOMATIC`, `SMB2`, and `SMB3`. Default: `AUTOMATIC`
func (o LocationSmbMountOptionsOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationSmbMountOptions) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type LocationSmbMountOptionsPtrOutput struct{ *pulumi.OutputState }

func (LocationSmbMountOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationSmbMountOptions)(nil)).Elem()
}

func (o LocationSmbMountOptionsPtrOutput) ToLocationSmbMountOptionsPtrOutput() LocationSmbMountOptionsPtrOutput {
	return o
}

func (o LocationSmbMountOptionsPtrOutput) ToLocationSmbMountOptionsPtrOutputWithContext(ctx context.Context) LocationSmbMountOptionsPtrOutput {
	return o
}

func (o LocationSmbMountOptionsPtrOutput) Elem() LocationSmbMountOptionsOutput {
	return o.ApplyT(func(v *LocationSmbMountOptions) LocationSmbMountOptions { return *v }).(LocationSmbMountOptionsOutput)
}

// The specific SMB version that you want DataSync to use for mounting your SMB share. Valid values: `AUTOMATIC`, `SMB2`, and `SMB3`. Default: `AUTOMATIC`
func (o LocationSmbMountOptionsPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationSmbMountOptions) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type NfsLocationOnPremConfig struct {
	// List of Amazon Resource Names (ARNs) of the DataSync Agents used to connect to the NFS server.
	AgentArns []string `pulumi:"agentArns"`
}

// NfsLocationOnPremConfigInput is an input type that accepts NfsLocationOnPremConfigArgs and NfsLocationOnPremConfigOutput values.
// You can construct a concrete instance of `NfsLocationOnPremConfigInput` via:
//
//          NfsLocationOnPremConfigArgs{...}
type NfsLocationOnPremConfigInput interface {
	pulumi.Input

	ToNfsLocationOnPremConfigOutput() NfsLocationOnPremConfigOutput
	ToNfsLocationOnPremConfigOutputWithContext(context.Context) NfsLocationOnPremConfigOutput
}

type NfsLocationOnPremConfigArgs struct {
	// List of Amazon Resource Names (ARNs) of the DataSync Agents used to connect to the NFS server.
	AgentArns pulumi.StringArrayInput `pulumi:"agentArns"`
}

func (NfsLocationOnPremConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsLocationOnPremConfig)(nil)).Elem()
}

func (i NfsLocationOnPremConfigArgs) ToNfsLocationOnPremConfigOutput() NfsLocationOnPremConfigOutput {
	return i.ToNfsLocationOnPremConfigOutputWithContext(context.Background())
}

func (i NfsLocationOnPremConfigArgs) ToNfsLocationOnPremConfigOutputWithContext(ctx context.Context) NfsLocationOnPremConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsLocationOnPremConfigOutput)
}

func (i NfsLocationOnPremConfigArgs) ToNfsLocationOnPremConfigPtrOutput() NfsLocationOnPremConfigPtrOutput {
	return i.ToNfsLocationOnPremConfigPtrOutputWithContext(context.Background())
}

func (i NfsLocationOnPremConfigArgs) ToNfsLocationOnPremConfigPtrOutputWithContext(ctx context.Context) NfsLocationOnPremConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsLocationOnPremConfigOutput).ToNfsLocationOnPremConfigPtrOutputWithContext(ctx)
}

// NfsLocationOnPremConfigPtrInput is an input type that accepts NfsLocationOnPremConfigArgs, NfsLocationOnPremConfigPtr and NfsLocationOnPremConfigPtrOutput values.
// You can construct a concrete instance of `NfsLocationOnPremConfigPtrInput` via:
//
//          NfsLocationOnPremConfigArgs{...}
//
//  or:
//
//          nil
type NfsLocationOnPremConfigPtrInput interface {
	pulumi.Input

	ToNfsLocationOnPremConfigPtrOutput() NfsLocationOnPremConfigPtrOutput
	ToNfsLocationOnPremConfigPtrOutputWithContext(context.Context) NfsLocationOnPremConfigPtrOutput
}

type nfsLocationOnPremConfigPtrType NfsLocationOnPremConfigArgs

func NfsLocationOnPremConfigPtr(v *NfsLocationOnPremConfigArgs) NfsLocationOnPremConfigPtrInput {
	return (*nfsLocationOnPremConfigPtrType)(v)
}

func (*nfsLocationOnPremConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NfsLocationOnPremConfig)(nil)).Elem()
}

func (i *nfsLocationOnPremConfigPtrType) ToNfsLocationOnPremConfigPtrOutput() NfsLocationOnPremConfigPtrOutput {
	return i.ToNfsLocationOnPremConfigPtrOutputWithContext(context.Background())
}

func (i *nfsLocationOnPremConfigPtrType) ToNfsLocationOnPremConfigPtrOutputWithContext(ctx context.Context) NfsLocationOnPremConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsLocationOnPremConfigPtrOutput)
}

type NfsLocationOnPremConfigOutput struct{ *pulumi.OutputState }

func (NfsLocationOnPremConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsLocationOnPremConfig)(nil)).Elem()
}

func (o NfsLocationOnPremConfigOutput) ToNfsLocationOnPremConfigOutput() NfsLocationOnPremConfigOutput {
	return o
}

func (o NfsLocationOnPremConfigOutput) ToNfsLocationOnPremConfigOutputWithContext(ctx context.Context) NfsLocationOnPremConfigOutput {
	return o
}

func (o NfsLocationOnPremConfigOutput) ToNfsLocationOnPremConfigPtrOutput() NfsLocationOnPremConfigPtrOutput {
	return o.ToNfsLocationOnPremConfigPtrOutputWithContext(context.Background())
}

func (o NfsLocationOnPremConfigOutput) ToNfsLocationOnPremConfigPtrOutputWithContext(ctx context.Context) NfsLocationOnPremConfigPtrOutput {
	return o.ApplyT(func(v NfsLocationOnPremConfig) *NfsLocationOnPremConfig {
		return &v
	}).(NfsLocationOnPremConfigPtrOutput)
}

// List of Amazon Resource Names (ARNs) of the DataSync Agents used to connect to the NFS server.
func (o NfsLocationOnPremConfigOutput) AgentArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NfsLocationOnPremConfig) []string { return v.AgentArns }).(pulumi.StringArrayOutput)
}

type NfsLocationOnPremConfigPtrOutput struct{ *pulumi.OutputState }

func (NfsLocationOnPremConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NfsLocationOnPremConfig)(nil)).Elem()
}

func (o NfsLocationOnPremConfigPtrOutput) ToNfsLocationOnPremConfigPtrOutput() NfsLocationOnPremConfigPtrOutput {
	return o
}

func (o NfsLocationOnPremConfigPtrOutput) ToNfsLocationOnPremConfigPtrOutputWithContext(ctx context.Context) NfsLocationOnPremConfigPtrOutput {
	return o
}

func (o NfsLocationOnPremConfigPtrOutput) Elem() NfsLocationOnPremConfigOutput {
	return o.ApplyT(func(v *NfsLocationOnPremConfig) NfsLocationOnPremConfig { return *v }).(NfsLocationOnPremConfigOutput)
}

// List of Amazon Resource Names (ARNs) of the DataSync Agents used to connect to the NFS server.
func (o NfsLocationOnPremConfigPtrOutput) AgentArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NfsLocationOnPremConfig) []string {
		if v == nil {
			return nil
		}
		return v.AgentArns
	}).(pulumi.StringArrayOutput)
}

type S3LocationS3Config struct {
	// Amazon Resource Names (ARN) of the IAM Role used to connect to the S3 Bucket.
	BucketAccessRoleArn string `pulumi:"bucketAccessRoleArn"`
}

// S3LocationS3ConfigInput is an input type that accepts S3LocationS3ConfigArgs and S3LocationS3ConfigOutput values.
// You can construct a concrete instance of `S3LocationS3ConfigInput` via:
//
//          S3LocationS3ConfigArgs{...}
type S3LocationS3ConfigInput interface {
	pulumi.Input

	ToS3LocationS3ConfigOutput() S3LocationS3ConfigOutput
	ToS3LocationS3ConfigOutputWithContext(context.Context) S3LocationS3ConfigOutput
}

type S3LocationS3ConfigArgs struct {
	// Amazon Resource Names (ARN) of the IAM Role used to connect to the S3 Bucket.
	BucketAccessRoleArn pulumi.StringInput `pulumi:"bucketAccessRoleArn"`
}

func (S3LocationS3ConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*S3LocationS3Config)(nil)).Elem()
}

func (i S3LocationS3ConfigArgs) ToS3LocationS3ConfigOutput() S3LocationS3ConfigOutput {
	return i.ToS3LocationS3ConfigOutputWithContext(context.Background())
}

func (i S3LocationS3ConfigArgs) ToS3LocationS3ConfigOutputWithContext(ctx context.Context) S3LocationS3ConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3LocationS3ConfigOutput)
}

func (i S3LocationS3ConfigArgs) ToS3LocationS3ConfigPtrOutput() S3LocationS3ConfigPtrOutput {
	return i.ToS3LocationS3ConfigPtrOutputWithContext(context.Background())
}

func (i S3LocationS3ConfigArgs) ToS3LocationS3ConfigPtrOutputWithContext(ctx context.Context) S3LocationS3ConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3LocationS3ConfigOutput).ToS3LocationS3ConfigPtrOutputWithContext(ctx)
}

// S3LocationS3ConfigPtrInput is an input type that accepts S3LocationS3ConfigArgs, S3LocationS3ConfigPtr and S3LocationS3ConfigPtrOutput values.
// You can construct a concrete instance of `S3LocationS3ConfigPtrInput` via:
//
//          S3LocationS3ConfigArgs{...}
//
//  or:
//
//          nil
type S3LocationS3ConfigPtrInput interface {
	pulumi.Input

	ToS3LocationS3ConfigPtrOutput() S3LocationS3ConfigPtrOutput
	ToS3LocationS3ConfigPtrOutputWithContext(context.Context) S3LocationS3ConfigPtrOutput
}

type s3locationS3ConfigPtrType S3LocationS3ConfigArgs

func S3LocationS3ConfigPtr(v *S3LocationS3ConfigArgs) S3LocationS3ConfigPtrInput {
	return (*s3locationS3ConfigPtrType)(v)
}

func (*s3locationS3ConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**S3LocationS3Config)(nil)).Elem()
}

func (i *s3locationS3ConfigPtrType) ToS3LocationS3ConfigPtrOutput() S3LocationS3ConfigPtrOutput {
	return i.ToS3LocationS3ConfigPtrOutputWithContext(context.Background())
}

func (i *s3locationS3ConfigPtrType) ToS3LocationS3ConfigPtrOutputWithContext(ctx context.Context) S3LocationS3ConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3LocationS3ConfigPtrOutput)
}

type S3LocationS3ConfigOutput struct{ *pulumi.OutputState }

func (S3LocationS3ConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*S3LocationS3Config)(nil)).Elem()
}

func (o S3LocationS3ConfigOutput) ToS3LocationS3ConfigOutput() S3LocationS3ConfigOutput {
	return o
}

func (o S3LocationS3ConfigOutput) ToS3LocationS3ConfigOutputWithContext(ctx context.Context) S3LocationS3ConfigOutput {
	return o
}

func (o S3LocationS3ConfigOutput) ToS3LocationS3ConfigPtrOutput() S3LocationS3ConfigPtrOutput {
	return o.ToS3LocationS3ConfigPtrOutputWithContext(context.Background())
}

func (o S3LocationS3ConfigOutput) ToS3LocationS3ConfigPtrOutputWithContext(ctx context.Context) S3LocationS3ConfigPtrOutput {
	return o.ApplyT(func(v S3LocationS3Config) *S3LocationS3Config {
		return &v
	}).(S3LocationS3ConfigPtrOutput)
}

// Amazon Resource Names (ARN) of the IAM Role used to connect to the S3 Bucket.
func (o S3LocationS3ConfigOutput) BucketAccessRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v S3LocationS3Config) string { return v.BucketAccessRoleArn }).(pulumi.StringOutput)
}

type S3LocationS3ConfigPtrOutput struct{ *pulumi.OutputState }

func (S3LocationS3ConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**S3LocationS3Config)(nil)).Elem()
}

func (o S3LocationS3ConfigPtrOutput) ToS3LocationS3ConfigPtrOutput() S3LocationS3ConfigPtrOutput {
	return o
}

func (o S3LocationS3ConfigPtrOutput) ToS3LocationS3ConfigPtrOutputWithContext(ctx context.Context) S3LocationS3ConfigPtrOutput {
	return o
}

func (o S3LocationS3ConfigPtrOutput) Elem() S3LocationS3ConfigOutput {
	return o.ApplyT(func(v *S3LocationS3Config) S3LocationS3Config { return *v }).(S3LocationS3ConfigOutput)
}

// Amazon Resource Names (ARN) of the IAM Role used to connect to the S3 Bucket.
func (o S3LocationS3ConfigPtrOutput) BucketAccessRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *S3LocationS3Config) *string {
		if v == nil {
			return nil
		}
		return &v.BucketAccessRoleArn
	}).(pulumi.StringPtrOutput)
}

type TaskOptions struct {
	// A file metadata that shows the last time a file was accessed (that is when the file was read or written to). If set to `BEST_EFFORT`, the DataSync Task attempts to preserve the original (that is, the version before sync `PREPARING` phase) `atime` attribute on all source files. Valid values: `BEST_EFFORT`, `NONE`. Default: `BEST_EFFORT`.
	Atime *string `pulumi:"atime"`
	// Limits the bandwidth utilized. For example, to set a maximum of 1 MB, set this value to `1048576`. Value values: `-1` or greater. Default: `-1` (unlimited).
	BytesPerSecond *int `pulumi:"bytesPerSecond"`
	// Group identifier of the file's owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
	Gid *string `pulumi:"gid"`
	// A file metadata that indicates the last time a file was modified (written to) before the sync `PREPARING` phase. Value values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
	Mtime *string `pulumi:"mtime"`
	// Determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file. Valid values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
	PosixPermissions *string `pulumi:"posixPermissions"`
	// Whether files deleted in the source should be removed or preserved in the destination file system. Valid values: `PRESERVE`, `REMOVE`. Default: `PRESERVE`.
	PreserveDeletedFiles *string `pulumi:"preserveDeletedFiles"`
	// Whether the DataSync Task should preserve the metadata of block and character devices in the source files system, and recreate the files with that device name and metadata on the destination. The DataSync Task can’t sync the actual contents of such devices, because many of the devices are non-terminal and don’t return an end of file (EOF) marker. Valid values: `NONE`, `PRESERVE`. Default: `NONE` (ignore special devices).
	PreserveDevices *string `pulumi:"preserveDevices"`
	// User identifier of the file's owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
	Uid *string `pulumi:"uid"`
	// Whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred. Valid values: `NONE`, `POINT_IN_TIME_CONSISTENT`, `ONLY_FILES_TRANSFERRED`. Default: `POINT_IN_TIME_CONSISTENT`.
	VerifyMode *string `pulumi:"verifyMode"`
}

// TaskOptionsInput is an input type that accepts TaskOptionsArgs and TaskOptionsOutput values.
// You can construct a concrete instance of `TaskOptionsInput` via:
//
//          TaskOptionsArgs{...}
type TaskOptionsInput interface {
	pulumi.Input

	ToTaskOptionsOutput() TaskOptionsOutput
	ToTaskOptionsOutputWithContext(context.Context) TaskOptionsOutput
}

type TaskOptionsArgs struct {
	// A file metadata that shows the last time a file was accessed (that is when the file was read or written to). If set to `BEST_EFFORT`, the DataSync Task attempts to preserve the original (that is, the version before sync `PREPARING` phase) `atime` attribute on all source files. Valid values: `BEST_EFFORT`, `NONE`. Default: `BEST_EFFORT`.
	Atime pulumi.StringPtrInput `pulumi:"atime"`
	// Limits the bandwidth utilized. For example, to set a maximum of 1 MB, set this value to `1048576`. Value values: `-1` or greater. Default: `-1` (unlimited).
	BytesPerSecond pulumi.IntPtrInput `pulumi:"bytesPerSecond"`
	// Group identifier of the file's owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
	Gid pulumi.StringPtrInput `pulumi:"gid"`
	// A file metadata that indicates the last time a file was modified (written to) before the sync `PREPARING` phase. Value values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
	Mtime pulumi.StringPtrInput `pulumi:"mtime"`
	// Determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file. Valid values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
	PosixPermissions pulumi.StringPtrInput `pulumi:"posixPermissions"`
	// Whether files deleted in the source should be removed or preserved in the destination file system. Valid values: `PRESERVE`, `REMOVE`. Default: `PRESERVE`.
	PreserveDeletedFiles pulumi.StringPtrInput `pulumi:"preserveDeletedFiles"`
	// Whether the DataSync Task should preserve the metadata of block and character devices in the source files system, and recreate the files with that device name and metadata on the destination. The DataSync Task can’t sync the actual contents of such devices, because many of the devices are non-terminal and don’t return an end of file (EOF) marker. Valid values: `NONE`, `PRESERVE`. Default: `NONE` (ignore special devices).
	PreserveDevices pulumi.StringPtrInput `pulumi:"preserveDevices"`
	// User identifier of the file's owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
	Uid pulumi.StringPtrInput `pulumi:"uid"`
	// Whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred. Valid values: `NONE`, `POINT_IN_TIME_CONSISTENT`, `ONLY_FILES_TRANSFERRED`. Default: `POINT_IN_TIME_CONSISTENT`.
	VerifyMode pulumi.StringPtrInput `pulumi:"verifyMode"`
}

func (TaskOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskOptions)(nil)).Elem()
}

func (i TaskOptionsArgs) ToTaskOptionsOutput() TaskOptionsOutput {
	return i.ToTaskOptionsOutputWithContext(context.Background())
}

func (i TaskOptionsArgs) ToTaskOptionsOutputWithContext(ctx context.Context) TaskOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskOptionsOutput)
}

func (i TaskOptionsArgs) ToTaskOptionsPtrOutput() TaskOptionsPtrOutput {
	return i.ToTaskOptionsPtrOutputWithContext(context.Background())
}

func (i TaskOptionsArgs) ToTaskOptionsPtrOutputWithContext(ctx context.Context) TaskOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskOptionsOutput).ToTaskOptionsPtrOutputWithContext(ctx)
}

// TaskOptionsPtrInput is an input type that accepts TaskOptionsArgs, TaskOptionsPtr and TaskOptionsPtrOutput values.
// You can construct a concrete instance of `TaskOptionsPtrInput` via:
//
//          TaskOptionsArgs{...}
//
//  or:
//
//          nil
type TaskOptionsPtrInput interface {
	pulumi.Input

	ToTaskOptionsPtrOutput() TaskOptionsPtrOutput
	ToTaskOptionsPtrOutputWithContext(context.Context) TaskOptionsPtrOutput
}

type taskOptionsPtrType TaskOptionsArgs

func TaskOptionsPtr(v *TaskOptionsArgs) TaskOptionsPtrInput {
	return (*taskOptionsPtrType)(v)
}

func (*taskOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskOptions)(nil)).Elem()
}

func (i *taskOptionsPtrType) ToTaskOptionsPtrOutput() TaskOptionsPtrOutput {
	return i.ToTaskOptionsPtrOutputWithContext(context.Background())
}

func (i *taskOptionsPtrType) ToTaskOptionsPtrOutputWithContext(ctx context.Context) TaskOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskOptionsPtrOutput)
}

type TaskOptionsOutput struct{ *pulumi.OutputState }

func (TaskOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskOptions)(nil)).Elem()
}

func (o TaskOptionsOutput) ToTaskOptionsOutput() TaskOptionsOutput {
	return o
}

func (o TaskOptionsOutput) ToTaskOptionsOutputWithContext(ctx context.Context) TaskOptionsOutput {
	return o
}

func (o TaskOptionsOutput) ToTaskOptionsPtrOutput() TaskOptionsPtrOutput {
	return o.ToTaskOptionsPtrOutputWithContext(context.Background())
}

func (o TaskOptionsOutput) ToTaskOptionsPtrOutputWithContext(ctx context.Context) TaskOptionsPtrOutput {
	return o.ApplyT(func(v TaskOptions) *TaskOptions {
		return &v
	}).(TaskOptionsPtrOutput)
}

// A file metadata that shows the last time a file was accessed (that is when the file was read or written to). If set to `BEST_EFFORT`, the DataSync Task attempts to preserve the original (that is, the version before sync `PREPARING` phase) `atime` attribute on all source files. Valid values: `BEST_EFFORT`, `NONE`. Default: `BEST_EFFORT`.
func (o TaskOptionsOutput) Atime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskOptions) *string { return v.Atime }).(pulumi.StringPtrOutput)
}

// Limits the bandwidth utilized. For example, to set a maximum of 1 MB, set this value to `1048576`. Value values: `-1` or greater. Default: `-1` (unlimited).
func (o TaskOptionsOutput) BytesPerSecond() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TaskOptions) *int { return v.BytesPerSecond }).(pulumi.IntPtrOutput)
}

// Group identifier of the file's owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
func (o TaskOptionsOutput) Gid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskOptions) *string { return v.Gid }).(pulumi.StringPtrOutput)
}

// A file metadata that indicates the last time a file was modified (written to) before the sync `PREPARING` phase. Value values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
func (o TaskOptionsOutput) Mtime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskOptions) *string { return v.Mtime }).(pulumi.StringPtrOutput)
}

// Determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file. Valid values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
func (o TaskOptionsOutput) PosixPermissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskOptions) *string { return v.PosixPermissions }).(pulumi.StringPtrOutput)
}

// Whether files deleted in the source should be removed or preserved in the destination file system. Valid values: `PRESERVE`, `REMOVE`. Default: `PRESERVE`.
func (o TaskOptionsOutput) PreserveDeletedFiles() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskOptions) *string { return v.PreserveDeletedFiles }).(pulumi.StringPtrOutput)
}

// Whether the DataSync Task should preserve the metadata of block and character devices in the source files system, and recreate the files with that device name and metadata on the destination. The DataSync Task can’t sync the actual contents of such devices, because many of the devices are non-terminal and don’t return an end of file (EOF) marker. Valid values: `NONE`, `PRESERVE`. Default: `NONE` (ignore special devices).
func (o TaskOptionsOutput) PreserveDevices() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskOptions) *string { return v.PreserveDevices }).(pulumi.StringPtrOutput)
}

// User identifier of the file's owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
func (o TaskOptionsOutput) Uid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskOptions) *string { return v.Uid }).(pulumi.StringPtrOutput)
}

// Whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred. Valid values: `NONE`, `POINT_IN_TIME_CONSISTENT`, `ONLY_FILES_TRANSFERRED`. Default: `POINT_IN_TIME_CONSISTENT`.
func (o TaskOptionsOutput) VerifyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskOptions) *string { return v.VerifyMode }).(pulumi.StringPtrOutput)
}

type TaskOptionsPtrOutput struct{ *pulumi.OutputState }

func (TaskOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskOptions)(nil)).Elem()
}

func (o TaskOptionsPtrOutput) ToTaskOptionsPtrOutput() TaskOptionsPtrOutput {
	return o
}

func (o TaskOptionsPtrOutput) ToTaskOptionsPtrOutputWithContext(ctx context.Context) TaskOptionsPtrOutput {
	return o
}

func (o TaskOptionsPtrOutput) Elem() TaskOptionsOutput {
	return o.ApplyT(func(v *TaskOptions) TaskOptions { return *v }).(TaskOptionsOutput)
}

// A file metadata that shows the last time a file was accessed (that is when the file was read or written to). If set to `BEST_EFFORT`, the DataSync Task attempts to preserve the original (that is, the version before sync `PREPARING` phase) `atime` attribute on all source files. Valid values: `BEST_EFFORT`, `NONE`. Default: `BEST_EFFORT`.
func (o TaskOptionsPtrOutput) Atime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskOptions) *string {
		if v == nil {
			return nil
		}
		return v.Atime
	}).(pulumi.StringPtrOutput)
}

// Limits the bandwidth utilized. For example, to set a maximum of 1 MB, set this value to `1048576`. Value values: `-1` or greater. Default: `-1` (unlimited).
func (o TaskOptionsPtrOutput) BytesPerSecond() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TaskOptions) *int {
		if v == nil {
			return nil
		}
		return v.BytesPerSecond
	}).(pulumi.IntPtrOutput)
}

// Group identifier of the file's owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
func (o TaskOptionsPtrOutput) Gid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskOptions) *string {
		if v == nil {
			return nil
		}
		return v.Gid
	}).(pulumi.StringPtrOutput)
}

// A file metadata that indicates the last time a file was modified (written to) before the sync `PREPARING` phase. Value values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
func (o TaskOptionsPtrOutput) Mtime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskOptions) *string {
		if v == nil {
			return nil
		}
		return v.Mtime
	}).(pulumi.StringPtrOutput)
}

// Determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file. Valid values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
func (o TaskOptionsPtrOutput) PosixPermissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskOptions) *string {
		if v == nil {
			return nil
		}
		return v.PosixPermissions
	}).(pulumi.StringPtrOutput)
}

// Whether files deleted in the source should be removed or preserved in the destination file system. Valid values: `PRESERVE`, `REMOVE`. Default: `PRESERVE`.
func (o TaskOptionsPtrOutput) PreserveDeletedFiles() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskOptions) *string {
		if v == nil {
			return nil
		}
		return v.PreserveDeletedFiles
	}).(pulumi.StringPtrOutput)
}

// Whether the DataSync Task should preserve the metadata of block and character devices in the source files system, and recreate the files with that device name and metadata on the destination. The DataSync Task can’t sync the actual contents of such devices, because many of the devices are non-terminal and don’t return an end of file (EOF) marker. Valid values: `NONE`, `PRESERVE`. Default: `NONE` (ignore special devices).
func (o TaskOptionsPtrOutput) PreserveDevices() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskOptions) *string {
		if v == nil {
			return nil
		}
		return v.PreserveDevices
	}).(pulumi.StringPtrOutput)
}

// User identifier of the file's owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
func (o TaskOptionsPtrOutput) Uid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskOptions) *string {
		if v == nil {
			return nil
		}
		return v.Uid
	}).(pulumi.StringPtrOutput)
}

// Whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred. Valid values: `NONE`, `POINT_IN_TIME_CONSISTENT`, `ONLY_FILES_TRANSFERRED`. Default: `POINT_IN_TIME_CONSISTENT`.
func (o TaskOptionsPtrOutput) VerifyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskOptions) *string {
		if v == nil {
			return nil
		}
		return v.VerifyMode
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EfsLocationEc2ConfigOutput{})
	pulumi.RegisterOutputType(EfsLocationEc2ConfigPtrOutput{})
	pulumi.RegisterOutputType(LocationSmbMountOptionsOutput{})
	pulumi.RegisterOutputType(LocationSmbMountOptionsPtrOutput{})
	pulumi.RegisterOutputType(NfsLocationOnPremConfigOutput{})
	pulumi.RegisterOutputType(NfsLocationOnPremConfigPtrOutput{})
	pulumi.RegisterOutputType(S3LocationS3ConfigOutput{})
	pulumi.RegisterOutputType(S3LocationS3ConfigPtrOutput{})
	pulumi.RegisterOutputType(TaskOptionsOutput{})
	pulumi.RegisterOutputType(TaskOptionsPtrOutput{})
}
