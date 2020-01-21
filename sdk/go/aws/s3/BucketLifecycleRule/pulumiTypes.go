// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package BucketLifecycleRule

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/s3/BucketLifecycleRuleExpiration"
	"https:/github.com/pulumi/pulumi-aws/s3/BucketLifecycleRuleNoncurrentVersionExpiration"
	"https:/github.com/pulumi/pulumi-aws/s3/BucketLifecycleRuleNoncurrentVersionTransition"
	"https:/github.com/pulumi/pulumi-aws/s3/BucketLifecycleRuleTransition"
)

type BucketLifecycleRule struct {
	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	AbortIncompleteMultipartUploadDays *int `pulumi:"abortIncompleteMultipartUploadDays"`
	// Boolean which indicates if this criteria is enabled.
	Enabled bool `pulumi:"enabled"`
	// Specifies a period in the object's expire (documented below).
	Expiration *s3BucketLifecycleRuleExpiration.BucketLifecycleRuleExpiration `pulumi:"expiration"`
	// Unique identifier for the rule.
	Id *string `pulumi:"id"`
	// Specifies when noncurrent object versions expire (documented below).
	NoncurrentVersionExpiration *s3BucketLifecycleRuleNoncurrentVersionExpiration.BucketLifecycleRuleNoncurrentVersionExpiration `pulumi:"noncurrentVersionExpiration"`
	// Specifies when noncurrent object versions transitions (documented below).
	NoncurrentVersionTransitions []s3BucketLifecycleRuleNoncurrentVersionTransition.BucketLifecycleRuleNoncurrentVersionTransition `pulumi:"noncurrentVersionTransitions"`
	// Object keyname prefix that identifies subset of objects to which the rule applies.
	Prefix *string `pulumi:"prefix"`
	// A mapping of tags that identifies subset of objects to which the rule applies.
	// The rule applies only to objects having all the tags in its tagset.
	Tags map[string]interface{} `pulumi:"tags"`
	// Specifies a period in the object's transitions (documented below).
	Transitions []s3BucketLifecycleRuleTransition.BucketLifecycleRuleTransition `pulumi:"transitions"`
}

type BucketLifecycleRuleInput interface {
	pulumi.Input

	ToBucketLifecycleRuleOutput() BucketLifecycleRuleOutput
	ToBucketLifecycleRuleOutputWithContext(context.Context) BucketLifecycleRuleOutput
}

type BucketLifecycleRuleArgs struct {
	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	AbortIncompleteMultipartUploadDays pulumi.IntPtrInput `pulumi:"abortIncompleteMultipartUploadDays"`
	// Boolean which indicates if this criteria is enabled.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	// Specifies a period in the object's expire (documented below).
	Expiration s3BucketLifecycleRuleExpiration.BucketLifecycleRuleExpirationPtrInput `pulumi:"expiration"`
	// Unique identifier for the rule.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Specifies when noncurrent object versions expire (documented below).
	NoncurrentVersionExpiration s3BucketLifecycleRuleNoncurrentVersionExpiration.BucketLifecycleRuleNoncurrentVersionExpirationPtrInput `pulumi:"noncurrentVersionExpiration"`
	// Specifies when noncurrent object versions transitions (documented below).
	NoncurrentVersionTransitions s3BucketLifecycleRuleNoncurrentVersionTransition.BucketLifecycleRuleNoncurrentVersionTransitionArrayInput `pulumi:"noncurrentVersionTransitions"`
	// Object keyname prefix that identifies subset of objects to which the rule applies.
	Prefix pulumi.StringPtrInput `pulumi:"prefix"`
	// A mapping of tags that identifies subset of objects to which the rule applies.
	// The rule applies only to objects having all the tags in its tagset.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Specifies a period in the object's transitions (documented below).
	Transitions s3BucketLifecycleRuleTransition.BucketLifecycleRuleTransitionArrayInput `pulumi:"transitions"`
}

func (BucketLifecycleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketLifecycleRule)(nil)).Elem()
}

func (i BucketLifecycleRuleArgs) ToBucketLifecycleRuleOutput() BucketLifecycleRuleOutput {
	return i.ToBucketLifecycleRuleOutputWithContext(context.Background())
}

func (i BucketLifecycleRuleArgs) ToBucketLifecycleRuleOutputWithContext(ctx context.Context) BucketLifecycleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleRuleOutput)
}

type BucketLifecycleRuleArrayInput interface {
	pulumi.Input

	ToBucketLifecycleRuleArrayOutput() BucketLifecycleRuleArrayOutput
	ToBucketLifecycleRuleArrayOutputWithContext(context.Context) BucketLifecycleRuleArrayOutput
}

type BucketLifecycleRuleArray []BucketLifecycleRuleInput

func (BucketLifecycleRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BucketLifecycleRule)(nil)).Elem()
}

func (i BucketLifecycleRuleArray) ToBucketLifecycleRuleArrayOutput() BucketLifecycleRuleArrayOutput {
	return i.ToBucketLifecycleRuleArrayOutputWithContext(context.Background())
}

func (i BucketLifecycleRuleArray) ToBucketLifecycleRuleArrayOutputWithContext(ctx context.Context) BucketLifecycleRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleRuleArrayOutput)
}

type BucketLifecycleRuleOutput struct { *pulumi.OutputState }

func (BucketLifecycleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketLifecycleRule)(nil)).Elem()
}

func (o BucketLifecycleRuleOutput) ToBucketLifecycleRuleOutput() BucketLifecycleRuleOutput {
	return o
}

func (o BucketLifecycleRuleOutput) ToBucketLifecycleRuleOutputWithContext(ctx context.Context) BucketLifecycleRuleOutput {
	return o
}

// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
func (o BucketLifecycleRuleOutput) AbortIncompleteMultipartUploadDays() pulumi.IntPtrOutput {
	return o.ApplyT(func (v BucketLifecycleRule) *int { return v.AbortIncompleteMultipartUploadDays }).(pulumi.IntPtrOutput)
}

// Boolean which indicates if this criteria is enabled.
func (o BucketLifecycleRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func (v BucketLifecycleRule) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// Specifies a period in the object's expire (documented below).
func (o BucketLifecycleRuleOutput) Expiration() s3BucketLifecycleRuleExpiration.BucketLifecycleRuleExpirationPtrOutput {
	return o.ApplyT(func (v BucketLifecycleRule) *s3BucketLifecycleRuleExpiration.BucketLifecycleRuleExpiration { return v.Expiration }).(s3BucketLifecycleRuleExpiration.BucketLifecycleRuleExpirationPtrOutput)
}

// Unique identifier for the rule.
func (o BucketLifecycleRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func (v BucketLifecycleRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Specifies when noncurrent object versions expire (documented below).
func (o BucketLifecycleRuleOutput) NoncurrentVersionExpiration() s3BucketLifecycleRuleNoncurrentVersionExpiration.BucketLifecycleRuleNoncurrentVersionExpirationPtrOutput {
	return o.ApplyT(func (v BucketLifecycleRule) *s3BucketLifecycleRuleNoncurrentVersionExpiration.BucketLifecycleRuleNoncurrentVersionExpiration { return v.NoncurrentVersionExpiration }).(s3BucketLifecycleRuleNoncurrentVersionExpiration.BucketLifecycleRuleNoncurrentVersionExpirationPtrOutput)
}

// Specifies when noncurrent object versions transitions (documented below).
func (o BucketLifecycleRuleOutput) NoncurrentVersionTransitions() s3BucketLifecycleRuleNoncurrentVersionTransition.BucketLifecycleRuleNoncurrentVersionTransitionArrayOutput {
	return o.ApplyT(func (v BucketLifecycleRule) []s3BucketLifecycleRuleNoncurrentVersionTransition.BucketLifecycleRuleNoncurrentVersionTransition { return v.NoncurrentVersionTransitions }).(s3BucketLifecycleRuleNoncurrentVersionTransition.BucketLifecycleRuleNoncurrentVersionTransitionArrayOutput)
}

// Object keyname prefix that identifies subset of objects to which the rule applies.
func (o BucketLifecycleRuleOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func (v BucketLifecycleRule) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

// A mapping of tags that identifies subset of objects to which the rule applies.
// The rule applies only to objects having all the tags in its tagset.
func (o BucketLifecycleRuleOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func (v BucketLifecycleRule) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

// Specifies a period in the object's transitions (documented below).
func (o BucketLifecycleRuleOutput) Transitions() s3BucketLifecycleRuleTransition.BucketLifecycleRuleTransitionArrayOutput {
	return o.ApplyT(func (v BucketLifecycleRule) []s3BucketLifecycleRuleTransition.BucketLifecycleRuleTransition { return v.Transitions }).(s3BucketLifecycleRuleTransition.BucketLifecycleRuleTransitionArrayOutput)
}

type BucketLifecycleRuleArrayOutput struct { *pulumi.OutputState}

func (BucketLifecycleRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BucketLifecycleRule)(nil)).Elem()
}

func (o BucketLifecycleRuleArrayOutput) ToBucketLifecycleRuleArrayOutput() BucketLifecycleRuleArrayOutput {
	return o
}

func (o BucketLifecycleRuleArrayOutput) ToBucketLifecycleRuleArrayOutputWithContext(ctx context.Context) BucketLifecycleRuleArrayOutput {
	return o
}

func (o BucketLifecycleRuleArrayOutput) Index(i pulumi.IntInput) BucketLifecycleRuleOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) BucketLifecycleRule {
		return vs[0].([]BucketLifecycleRule)[vs[1].(int)]
	}).(BucketLifecycleRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(BucketLifecycleRuleOutput{})
	pulumi.RegisterOutputType(BucketLifecycleRuleArrayOutput{})
}
