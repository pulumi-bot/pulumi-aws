// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterS3Import

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClusterS3Import struct {
	// The bucket name where your backup is stored
	BucketName string `pulumi:"bucketName"`
	// Can be blank, but is the path to your backup
	BucketPrefix *string `pulumi:"bucketPrefix"`
	// Role applied to load the data.
	IngestionRole string `pulumi:"ingestionRole"`
	// Source engine for the backup
	SourceEngine string `pulumi:"sourceEngine"`
	// Version of the source engine used to make the backup
	SourceEngineVersion string `pulumi:"sourceEngineVersion"`
}

type ClusterS3ImportInput interface {
	pulumi.Input

	ToClusterS3ImportOutput() ClusterS3ImportOutput
	ToClusterS3ImportOutputWithContext(context.Context) ClusterS3ImportOutput
}

type ClusterS3ImportArgs struct {
	// The bucket name where your backup is stored
	BucketName pulumi.StringInput `pulumi:"bucketName"`
	// Can be blank, but is the path to your backup
	BucketPrefix pulumi.StringPtrInput `pulumi:"bucketPrefix"`
	// Role applied to load the data.
	IngestionRole pulumi.StringInput `pulumi:"ingestionRole"`
	// Source engine for the backup
	SourceEngine pulumi.StringInput `pulumi:"sourceEngine"`
	// Version of the source engine used to make the backup
	SourceEngineVersion pulumi.StringInput `pulumi:"sourceEngineVersion"`
}

func (ClusterS3ImportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterS3Import)(nil)).Elem()
}

func (i ClusterS3ImportArgs) ToClusterS3ImportOutput() ClusterS3ImportOutput {
	return i.ToClusterS3ImportOutputWithContext(context.Background())
}

func (i ClusterS3ImportArgs) ToClusterS3ImportOutputWithContext(ctx context.Context) ClusterS3ImportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterS3ImportOutput)
}

func (i ClusterS3ImportArgs) ToClusterS3ImportPtrOutput() ClusterS3ImportPtrOutput {
	return i.ToClusterS3ImportPtrOutputWithContext(context.Background())
}

func (i ClusterS3ImportArgs) ToClusterS3ImportPtrOutputWithContext(ctx context.Context) ClusterS3ImportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterS3ImportOutput).ToClusterS3ImportPtrOutputWithContext(ctx)
}

type ClusterS3ImportPtrInput interface {
	pulumi.Input

	ToClusterS3ImportPtrOutput() ClusterS3ImportPtrOutput
	ToClusterS3ImportPtrOutputWithContext(context.Context) ClusterS3ImportPtrOutput
}

type clusterS3ImportPtrType ClusterS3ImportArgs

func ClusterS3ImportPtr(v *ClusterS3ImportArgs) ClusterS3ImportPtrInput {	return (*clusterS3ImportPtrType)(v)
}

func (*clusterS3ImportPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterS3Import)(nil)).Elem()
}

func (i *clusterS3ImportPtrType) ToClusterS3ImportPtrOutput() ClusterS3ImportPtrOutput {
	return i.ToClusterS3ImportPtrOutputWithContext(context.Background())
}

func (i *clusterS3ImportPtrType) ToClusterS3ImportPtrOutputWithContext(ctx context.Context) ClusterS3ImportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterS3ImportPtrOutput)
}

type ClusterS3ImportOutput struct { *pulumi.OutputState }

func (ClusterS3ImportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterS3Import)(nil)).Elem()
}

func (o ClusterS3ImportOutput) ToClusterS3ImportOutput() ClusterS3ImportOutput {
	return o
}

func (o ClusterS3ImportOutput) ToClusterS3ImportOutputWithContext(ctx context.Context) ClusterS3ImportOutput {
	return o
}

func (o ClusterS3ImportOutput) ToClusterS3ImportPtrOutput() ClusterS3ImportPtrOutput {
	return o.ToClusterS3ImportPtrOutputWithContext(context.Background())
}

func (o ClusterS3ImportOutput) ToClusterS3ImportPtrOutputWithContext(ctx context.Context) ClusterS3ImportPtrOutput {
	return o.ApplyT(func(v ClusterS3Import) *ClusterS3Import {
		return &v
	}).(ClusterS3ImportPtrOutput)
}
// The bucket name where your backup is stored
func (o ClusterS3ImportOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterS3Import) string { return v.BucketName }).(pulumi.StringOutput)
}

// Can be blank, but is the path to your backup
func (o ClusterS3ImportOutput) BucketPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterS3Import) *string { return v.BucketPrefix }).(pulumi.StringPtrOutput)
}

// Role applied to load the data.
func (o ClusterS3ImportOutput) IngestionRole() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterS3Import) string { return v.IngestionRole }).(pulumi.StringOutput)
}

// Source engine for the backup
func (o ClusterS3ImportOutput) SourceEngine() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterS3Import) string { return v.SourceEngine }).(pulumi.StringOutput)
}

// Version of the source engine used to make the backup
func (o ClusterS3ImportOutput) SourceEngineVersion() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterS3Import) string { return v.SourceEngineVersion }).(pulumi.StringOutput)
}

type ClusterS3ImportPtrOutput struct { *pulumi.OutputState}

func (ClusterS3ImportPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterS3Import)(nil)).Elem()
}

func (o ClusterS3ImportPtrOutput) ToClusterS3ImportPtrOutput() ClusterS3ImportPtrOutput {
	return o
}

func (o ClusterS3ImportPtrOutput) ToClusterS3ImportPtrOutputWithContext(ctx context.Context) ClusterS3ImportPtrOutput {
	return o
}

func (o ClusterS3ImportPtrOutput) Elem() ClusterS3ImportOutput {
	return o.ApplyT(func (v *ClusterS3Import) ClusterS3Import { return *v }).(ClusterS3ImportOutput)
}

// The bucket name where your backup is stored
func (o ClusterS3ImportPtrOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterS3Import) string { return v.BucketName }).(pulumi.StringOutput)
}

// Can be blank, but is the path to your backup
func (o ClusterS3ImportPtrOutput) BucketPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterS3Import) *string { return v.BucketPrefix }).(pulumi.StringPtrOutput)
}

// Role applied to load the data.
func (o ClusterS3ImportPtrOutput) IngestionRole() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterS3Import) string { return v.IngestionRole }).(pulumi.StringOutput)
}

// Source engine for the backup
func (o ClusterS3ImportPtrOutput) SourceEngine() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterS3Import) string { return v.SourceEngine }).(pulumi.StringOutput)
}

// Version of the source engine used to make the backup
func (o ClusterS3ImportPtrOutput) SourceEngineVersion() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterS3Import) string { return v.SourceEngineVersion }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterS3ImportOutput{})
	pulumi.RegisterOutputType(ClusterS3ImportPtrOutput{})
}
