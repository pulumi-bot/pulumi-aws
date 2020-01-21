// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package BucketLifecycleRuleExpiration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type BucketLifecycleRuleExpiration struct {
	// Specifies the date after which you want the corresponding action to take effect.
	Date *string `pulumi:"date"`
	// The number of days that you want to specify for the default retention period.
	Days *int `pulumi:"days"`
	// On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Amazon S3 to delete expired object delete markers.
	ExpiredObjectDeleteMarker *bool `pulumi:"expiredObjectDeleteMarker"`
}

type BucketLifecycleRuleExpirationInput interface {
	pulumi.Input

	ToBucketLifecycleRuleExpirationOutput() BucketLifecycleRuleExpirationOutput
	ToBucketLifecycleRuleExpirationOutputWithContext(context.Context) BucketLifecycleRuleExpirationOutput
}

type BucketLifecycleRuleExpirationArgs struct {
	// Specifies the date after which you want the corresponding action to take effect.
	Date pulumi.StringPtrInput `pulumi:"date"`
	// The number of days that you want to specify for the default retention period.
	Days pulumi.IntPtrInput `pulumi:"days"`
	// On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Amazon S3 to delete expired object delete markers.
	ExpiredObjectDeleteMarker pulumi.BoolPtrInput `pulumi:"expiredObjectDeleteMarker"`
}

func (BucketLifecycleRuleExpirationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketLifecycleRuleExpiration)(nil)).Elem()
}

func (i BucketLifecycleRuleExpirationArgs) ToBucketLifecycleRuleExpirationOutput() BucketLifecycleRuleExpirationOutput {
	return i.ToBucketLifecycleRuleExpirationOutputWithContext(context.Background())
}

func (i BucketLifecycleRuleExpirationArgs) ToBucketLifecycleRuleExpirationOutputWithContext(ctx context.Context) BucketLifecycleRuleExpirationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleRuleExpirationOutput)
}

func (i BucketLifecycleRuleExpirationArgs) ToBucketLifecycleRuleExpirationPtrOutput() BucketLifecycleRuleExpirationPtrOutput {
	return i.ToBucketLifecycleRuleExpirationPtrOutputWithContext(context.Background())
}

func (i BucketLifecycleRuleExpirationArgs) ToBucketLifecycleRuleExpirationPtrOutputWithContext(ctx context.Context) BucketLifecycleRuleExpirationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleRuleExpirationOutput).ToBucketLifecycleRuleExpirationPtrOutputWithContext(ctx)
}

type BucketLifecycleRuleExpirationPtrInput interface {
	pulumi.Input

	ToBucketLifecycleRuleExpirationPtrOutput() BucketLifecycleRuleExpirationPtrOutput
	ToBucketLifecycleRuleExpirationPtrOutputWithContext(context.Context) BucketLifecycleRuleExpirationPtrOutput
}

type bucketLifecycleRuleExpirationPtrType BucketLifecycleRuleExpirationArgs

func BucketLifecycleRuleExpirationPtr(v *BucketLifecycleRuleExpirationArgs) BucketLifecycleRuleExpirationPtrInput {	return (*bucketLifecycleRuleExpirationPtrType)(v)
}

func (*bucketLifecycleRuleExpirationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketLifecycleRuleExpiration)(nil)).Elem()
}

func (i *bucketLifecycleRuleExpirationPtrType) ToBucketLifecycleRuleExpirationPtrOutput() BucketLifecycleRuleExpirationPtrOutput {
	return i.ToBucketLifecycleRuleExpirationPtrOutputWithContext(context.Background())
}

func (i *bucketLifecycleRuleExpirationPtrType) ToBucketLifecycleRuleExpirationPtrOutputWithContext(ctx context.Context) BucketLifecycleRuleExpirationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleRuleExpirationPtrOutput)
}

type BucketLifecycleRuleExpirationOutput struct { *pulumi.OutputState }

func (BucketLifecycleRuleExpirationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketLifecycleRuleExpiration)(nil)).Elem()
}

func (o BucketLifecycleRuleExpirationOutput) ToBucketLifecycleRuleExpirationOutput() BucketLifecycleRuleExpirationOutput {
	return o
}

func (o BucketLifecycleRuleExpirationOutput) ToBucketLifecycleRuleExpirationOutputWithContext(ctx context.Context) BucketLifecycleRuleExpirationOutput {
	return o
}

func (o BucketLifecycleRuleExpirationOutput) ToBucketLifecycleRuleExpirationPtrOutput() BucketLifecycleRuleExpirationPtrOutput {
	return o.ToBucketLifecycleRuleExpirationPtrOutputWithContext(context.Background())
}

func (o BucketLifecycleRuleExpirationOutput) ToBucketLifecycleRuleExpirationPtrOutputWithContext(ctx context.Context) BucketLifecycleRuleExpirationPtrOutput {
	return o.ApplyT(func(v BucketLifecycleRuleExpiration) *BucketLifecycleRuleExpiration {
		return &v
	}).(BucketLifecycleRuleExpirationPtrOutput)
}
// Specifies the date after which you want the corresponding action to take effect.
func (o BucketLifecycleRuleExpirationOutput) Date() pulumi.StringPtrOutput {
	return o.ApplyT(func (v BucketLifecycleRuleExpiration) *string { return v.Date }).(pulumi.StringPtrOutput)
}

// The number of days that you want to specify for the default retention period.
func (o BucketLifecycleRuleExpirationOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func (v BucketLifecycleRuleExpiration) *int { return v.Days }).(pulumi.IntPtrOutput)
}

// On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Amazon S3 to delete expired object delete markers.
func (o BucketLifecycleRuleExpirationOutput) ExpiredObjectDeleteMarker() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v BucketLifecycleRuleExpiration) *bool { return v.ExpiredObjectDeleteMarker }).(pulumi.BoolPtrOutput)
}

type BucketLifecycleRuleExpirationPtrOutput struct { *pulumi.OutputState}

func (BucketLifecycleRuleExpirationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketLifecycleRuleExpiration)(nil)).Elem()
}

func (o BucketLifecycleRuleExpirationPtrOutput) ToBucketLifecycleRuleExpirationPtrOutput() BucketLifecycleRuleExpirationPtrOutput {
	return o
}

func (o BucketLifecycleRuleExpirationPtrOutput) ToBucketLifecycleRuleExpirationPtrOutputWithContext(ctx context.Context) BucketLifecycleRuleExpirationPtrOutput {
	return o
}

func (o BucketLifecycleRuleExpirationPtrOutput) Elem() BucketLifecycleRuleExpirationOutput {
	return o.ApplyT(func (v *BucketLifecycleRuleExpiration) BucketLifecycleRuleExpiration { return *v }).(BucketLifecycleRuleExpirationOutput)
}

// Specifies the date after which you want the corresponding action to take effect.
func (o BucketLifecycleRuleExpirationPtrOutput) Date() pulumi.StringPtrOutput {
	return o.ApplyT(func (v BucketLifecycleRuleExpiration) *string { return v.Date }).(pulumi.StringPtrOutput)
}

// The number of days that you want to specify for the default retention period.
func (o BucketLifecycleRuleExpirationPtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func (v BucketLifecycleRuleExpiration) *int { return v.Days }).(pulumi.IntPtrOutput)
}

// On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Amazon S3 to delete expired object delete markers.
func (o BucketLifecycleRuleExpirationPtrOutput) ExpiredObjectDeleteMarker() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v BucketLifecycleRuleExpiration) *bool { return v.ExpiredObjectDeleteMarker }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BucketLifecycleRuleExpirationOutput{})
	pulumi.RegisterOutputType(BucketLifecycleRuleExpirationPtrOutput{})
}
