// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package BucketReplicationConfigurationRuleDestinationAccessControlTranslation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type BucketReplicationConfigurationRuleDestinationAccessControlTranslation struct {
	// The override value for the owner on replicated objects. Currently only `Destination` is supported.
	Owner string `pulumi:"owner"`
}

type BucketReplicationConfigurationRuleDestinationAccessControlTranslationInput interface {
	pulumi.Input

	ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput() BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput
	ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationOutputWithContext(context.Context) BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput
}

type BucketReplicationConfigurationRuleDestinationAccessControlTranslationArgs struct {
	// The override value for the owner on replicated objects. Currently only `Destination` is supported.
	Owner pulumi.StringInput `pulumi:"owner"`
}

func (BucketReplicationConfigurationRuleDestinationAccessControlTranslationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketReplicationConfigurationRuleDestinationAccessControlTranslation)(nil)).Elem()
}

func (i BucketReplicationConfigurationRuleDestinationAccessControlTranslationArgs) ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput() BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput {
	return i.ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationOutputWithContext(context.Background())
}

func (i BucketReplicationConfigurationRuleDestinationAccessControlTranslationArgs) ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationOutputWithContext(ctx context.Context) BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput)
}

func (i BucketReplicationConfigurationRuleDestinationAccessControlTranslationArgs) ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput() BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput {
	return i.ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutputWithContext(context.Background())
}

func (i BucketReplicationConfigurationRuleDestinationAccessControlTranslationArgs) ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutputWithContext(ctx context.Context) BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput).ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutputWithContext(ctx)
}

type BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrInput interface {
	pulumi.Input

	ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput() BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput
	ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutputWithContext(context.Context) BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput
}

type bucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrType BucketReplicationConfigurationRuleDestinationAccessControlTranslationArgs

func BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtr(v *BucketReplicationConfigurationRuleDestinationAccessControlTranslationArgs) BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrInput {	return (*bucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrType)(v)
}

func (*bucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketReplicationConfigurationRuleDestinationAccessControlTranslation)(nil)).Elem()
}

func (i *bucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrType) ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput() BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput {
	return i.ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutputWithContext(context.Background())
}

func (i *bucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrType) ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutputWithContext(ctx context.Context) BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput)
}

type BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput struct { *pulumi.OutputState }

func (BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketReplicationConfigurationRuleDestinationAccessControlTranslation)(nil)).Elem()
}

func (o BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput) ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput() BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput {
	return o
}

func (o BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput) ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationOutputWithContext(ctx context.Context) BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput {
	return o
}

func (o BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput) ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput() BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput {
	return o.ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutputWithContext(context.Background())
}

func (o BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput) ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutputWithContext(ctx context.Context) BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput {
	return o.ApplyT(func(v BucketReplicationConfigurationRuleDestinationAccessControlTranslation) *BucketReplicationConfigurationRuleDestinationAccessControlTranslation {
		return &v
	}).(BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput)
}
// The override value for the owner on replicated objects. Currently only `Destination` is supported.
func (o BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func (v BucketReplicationConfigurationRuleDestinationAccessControlTranslation) string { return v.Owner }).(pulumi.StringOutput)
}

type BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput struct { *pulumi.OutputState}

func (BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketReplicationConfigurationRuleDestinationAccessControlTranslation)(nil)).Elem()
}

func (o BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput) ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput() BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput {
	return o
}

func (o BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput) ToBucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutputWithContext(ctx context.Context) BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput {
	return o
}

func (o BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput) Elem() BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput {
	return o.ApplyT(func (v *BucketReplicationConfigurationRuleDestinationAccessControlTranslation) BucketReplicationConfigurationRuleDestinationAccessControlTranslation { return *v }).(BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput)
}

// The override value for the owner on replicated objects. Currently only `Destination` is supported.
func (o BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func (v BucketReplicationConfigurationRuleDestinationAccessControlTranslation) string { return v.Owner }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BucketReplicationConfigurationRuleDestinationAccessControlTranslationOutput{})
	pulumi.RegisterOutputType(BucketReplicationConfigurationRuleDestinationAccessControlTranslationPtrOutput{})
}
