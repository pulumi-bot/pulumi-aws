// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects struct {
	// Boolean which indicates if this criteria is enabled.
	Enabled bool `pulumi:"enabled"`
}

type BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsInput interface {
	pulumi.Input

	ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput() BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput
	ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutputWithContext(context.Context) BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput
}

type BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsArgs struct {
	// Boolean which indicates if this criteria is enabled.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects)(nil)).Elem()
}

func (i BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsArgs) ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput() BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput {
	return i.ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutputWithContext(context.Background())
}

func (i BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsArgs) ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutputWithContext(ctx context.Context) BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput)
}

func (i BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsArgs) ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput() BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput {
	return i.ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutputWithContext(context.Background())
}

func (i BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsArgs) ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutputWithContext(ctx context.Context) BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput).ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutputWithContext(ctx)
}

type BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrInput interface {
	pulumi.Input

	ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput() BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput
	ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutputWithContext(context.Context) BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput
}

type bucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrType BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsArgs

func BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtr(v *BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsArgs) BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrInput {	return (*bucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrType)(v)
}

func (*bucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects)(nil)).Elem()
}

func (i *bucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrType) ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput() BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput {
	return i.ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutputWithContext(context.Background())
}

func (i *bucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrType) ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutputWithContext(ctx context.Context) BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput)
}

type BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput struct { *pulumi.OutputState }

func (BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects)(nil)).Elem()
}

func (o BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput) ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput() BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput {
	return o
}

func (o BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput) ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutputWithContext(ctx context.Context) BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput {
	return o
}

func (o BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput) ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput() BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput {
	return o.ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutputWithContext(context.Background())
}

func (o BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput) ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutputWithContext(ctx context.Context) BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput {
	return o.ApplyT(func(v BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects) *BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects {
		return &v
	}).(BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput)
}
// Boolean which indicates if this criteria is enabled.
func (o BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func (v BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput struct { *pulumi.OutputState}

func (BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects)(nil)).Elem()
}

func (o BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput) ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput() BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput {
	return o
}

func (o BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput) ToBucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutputWithContext(ctx context.Context) BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput {
	return o
}

func (o BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput) Elem() BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput {
	return o.ApplyT(func (v *BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects) BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects { return *v }).(BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput)
}

// Boolean which indicates if this criteria is enabled.
func (o BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func (v BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsOutput{})
	pulumi.RegisterOutputType(BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjectsPtrOutput{})
}
