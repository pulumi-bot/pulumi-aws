// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterLogging

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClusterLogging struct {
	// The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
	// For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
	BucketName *string `pulumi:"bucketName"`
	// Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
	Enable bool `pulumi:"enable"`
	// The prefix applied to the log file names.
	S3KeyPrefix *string `pulumi:"s3KeyPrefix"`
}

type ClusterLoggingInput interface {
	pulumi.Input

	ToClusterLoggingOutput() ClusterLoggingOutput
	ToClusterLoggingOutputWithContext(context.Context) ClusterLoggingOutput
}

type ClusterLoggingArgs struct {
	// The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
	// For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
	BucketName pulumi.StringPtrInput `pulumi:"bucketName"`
	// Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
	Enable pulumi.BoolInput `pulumi:"enable"`
	// The prefix applied to the log file names.
	S3KeyPrefix pulumi.StringPtrInput `pulumi:"s3KeyPrefix"`
}

func (ClusterLoggingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterLogging)(nil)).Elem()
}

func (i ClusterLoggingArgs) ToClusterLoggingOutput() ClusterLoggingOutput {
	return i.ToClusterLoggingOutputWithContext(context.Background())
}

func (i ClusterLoggingArgs) ToClusterLoggingOutputWithContext(ctx context.Context) ClusterLoggingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterLoggingOutput)
}

func (i ClusterLoggingArgs) ToClusterLoggingPtrOutput() ClusterLoggingPtrOutput {
	return i.ToClusterLoggingPtrOutputWithContext(context.Background())
}

func (i ClusterLoggingArgs) ToClusterLoggingPtrOutputWithContext(ctx context.Context) ClusterLoggingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterLoggingOutput).ToClusterLoggingPtrOutputWithContext(ctx)
}

type ClusterLoggingPtrInput interface {
	pulumi.Input

	ToClusterLoggingPtrOutput() ClusterLoggingPtrOutput
	ToClusterLoggingPtrOutputWithContext(context.Context) ClusterLoggingPtrOutput
}

type clusterLoggingPtrType ClusterLoggingArgs

func ClusterLoggingPtr(v *ClusterLoggingArgs) ClusterLoggingPtrInput {	return (*clusterLoggingPtrType)(v)
}

func (*clusterLoggingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterLogging)(nil)).Elem()
}

func (i *clusterLoggingPtrType) ToClusterLoggingPtrOutput() ClusterLoggingPtrOutput {
	return i.ToClusterLoggingPtrOutputWithContext(context.Background())
}

func (i *clusterLoggingPtrType) ToClusterLoggingPtrOutputWithContext(ctx context.Context) ClusterLoggingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterLoggingPtrOutput)
}

type ClusterLoggingOutput struct { *pulumi.OutputState }

func (ClusterLoggingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterLogging)(nil)).Elem()
}

func (o ClusterLoggingOutput) ToClusterLoggingOutput() ClusterLoggingOutput {
	return o
}

func (o ClusterLoggingOutput) ToClusterLoggingOutputWithContext(ctx context.Context) ClusterLoggingOutput {
	return o
}

func (o ClusterLoggingOutput) ToClusterLoggingPtrOutput() ClusterLoggingPtrOutput {
	return o.ToClusterLoggingPtrOutputWithContext(context.Background())
}

func (o ClusterLoggingOutput) ToClusterLoggingPtrOutputWithContext(ctx context.Context) ClusterLoggingPtrOutput {
	return o.ApplyT(func(v ClusterLogging) *ClusterLogging {
		return &v
	}).(ClusterLoggingPtrOutput)
}
// The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
// For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
func (o ClusterLoggingOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterLogging) *string { return v.BucketName }).(pulumi.StringPtrOutput)
}

// Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
func (o ClusterLoggingOutput) Enable() pulumi.BoolOutput {
	return o.ApplyT(func (v ClusterLogging) bool { return v.Enable }).(pulumi.BoolOutput)
}

// The prefix applied to the log file names.
func (o ClusterLoggingOutput) S3KeyPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterLogging) *string { return v.S3KeyPrefix }).(pulumi.StringPtrOutput)
}

type ClusterLoggingPtrOutput struct { *pulumi.OutputState}

func (ClusterLoggingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterLogging)(nil)).Elem()
}

func (o ClusterLoggingPtrOutput) ToClusterLoggingPtrOutput() ClusterLoggingPtrOutput {
	return o
}

func (o ClusterLoggingPtrOutput) ToClusterLoggingPtrOutputWithContext(ctx context.Context) ClusterLoggingPtrOutput {
	return o
}

func (o ClusterLoggingPtrOutput) Elem() ClusterLoggingOutput {
	return o.ApplyT(func (v *ClusterLogging) ClusterLogging { return *v }).(ClusterLoggingOutput)
}

// The name of an existing S3 bucket where the log files are to be stored. Must be in the same region as the cluster and the cluster must have read bucket and put object permissions.
// For more information on the permissions required for the bucket, please read the AWS [documentation](http://docs.aws.amazon.com/redshift/latest/mgmt/db-auditing.html#db-auditing-enable-logging)
func (o ClusterLoggingPtrOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterLogging) *string { return v.BucketName }).(pulumi.StringPtrOutput)
}

// Enables logging information such as queries and connection attempts, for the specified Amazon Redshift cluster.
func (o ClusterLoggingPtrOutput) Enable() pulumi.BoolOutput {
	return o.ApplyT(func (v ClusterLogging) bool { return v.Enable }).(pulumi.BoolOutput)
}

// The prefix applied to the log file names.
func (o ClusterLoggingPtrOutput) S3KeyPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterLogging) *string { return v.S3KeyPrefix }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterLoggingOutput{})
	pulumi.RegisterOutputType(ClusterLoggingPtrOutput{})
}
