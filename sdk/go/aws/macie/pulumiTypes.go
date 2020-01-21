// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package macie

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type S3BucketAssociationClassificationType struct {
	// A string value indicating that Macie perform a one-time classification of all of the existing objects in the bucket.
	// The only valid value is the default value, `FULL`.
	Continuous *string `pulumi:"continuous"`
	// A string value indicating whether or not Macie performs a one-time classification of all of the existing objects in the bucket.
	// Valid values are `NONE` and `FULL`. Defaults to `NONE` indicating that Macie only classifies objects that are added after the association was created.
	OneTime *string `pulumi:"oneTime"`
}

type S3BucketAssociationClassificationTypeInput interface {
	pulumi.Input

	ToS3BucketAssociationClassificationTypeOutput() S3BucketAssociationClassificationTypeOutput
	ToS3BucketAssociationClassificationTypeOutputWithContext(context.Context) S3BucketAssociationClassificationTypeOutput
}

type S3BucketAssociationClassificationTypeArgs struct {
	// A string value indicating that Macie perform a one-time classification of all of the existing objects in the bucket.
	// The only valid value is the default value, `FULL`.
	Continuous pulumi.StringPtrInput `pulumi:"continuous"`
	// A string value indicating whether or not Macie performs a one-time classification of all of the existing objects in the bucket.
	// Valid values are `NONE` and `FULL`. Defaults to `NONE` indicating that Macie only classifies objects that are added after the association was created.
	OneTime pulumi.StringPtrInput `pulumi:"oneTime"`
}

func (S3BucketAssociationClassificationTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*S3BucketAssociationClassificationType)(nil)).Elem()
}

func (i S3BucketAssociationClassificationTypeArgs) ToS3BucketAssociationClassificationTypeOutput() S3BucketAssociationClassificationTypeOutput {
	return i.ToS3BucketAssociationClassificationTypeOutputWithContext(context.Background())
}

func (i S3BucketAssociationClassificationTypeArgs) ToS3BucketAssociationClassificationTypeOutputWithContext(ctx context.Context) S3BucketAssociationClassificationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketAssociationClassificationTypeOutput)
}

func (i S3BucketAssociationClassificationTypeArgs) ToS3BucketAssociationClassificationTypePtrOutput() S3BucketAssociationClassificationTypePtrOutput {
	return i.ToS3BucketAssociationClassificationTypePtrOutputWithContext(context.Background())
}

func (i S3BucketAssociationClassificationTypeArgs) ToS3BucketAssociationClassificationTypePtrOutputWithContext(ctx context.Context) S3BucketAssociationClassificationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketAssociationClassificationTypeOutput).ToS3BucketAssociationClassificationTypePtrOutputWithContext(ctx)
}

type S3BucketAssociationClassificationTypePtrInput interface {
	pulumi.Input

	ToS3BucketAssociationClassificationTypePtrOutput() S3BucketAssociationClassificationTypePtrOutput
	ToS3BucketAssociationClassificationTypePtrOutputWithContext(context.Context) S3BucketAssociationClassificationTypePtrOutput
}

type s3bucketAssociationClassificationTypePtrType S3BucketAssociationClassificationTypeArgs

func S3BucketAssociationClassificationTypePtr(v *S3BucketAssociationClassificationTypeArgs) S3BucketAssociationClassificationTypePtrInput {	return (*s3bucketAssociationClassificationTypePtrType)(v)
}

func (*s3bucketAssociationClassificationTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**S3BucketAssociationClassificationType)(nil)).Elem()
}

func (i *s3bucketAssociationClassificationTypePtrType) ToS3BucketAssociationClassificationTypePtrOutput() S3BucketAssociationClassificationTypePtrOutput {
	return i.ToS3BucketAssociationClassificationTypePtrOutputWithContext(context.Background())
}

func (i *s3bucketAssociationClassificationTypePtrType) ToS3BucketAssociationClassificationTypePtrOutputWithContext(ctx context.Context) S3BucketAssociationClassificationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketAssociationClassificationTypePtrOutput)
}

type S3BucketAssociationClassificationTypeOutput struct { *pulumi.OutputState }

func (S3BucketAssociationClassificationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*S3BucketAssociationClassificationType)(nil)).Elem()
}

func (o S3BucketAssociationClassificationTypeOutput) ToS3BucketAssociationClassificationTypeOutput() S3BucketAssociationClassificationTypeOutput {
	return o
}

func (o S3BucketAssociationClassificationTypeOutput) ToS3BucketAssociationClassificationTypeOutputWithContext(ctx context.Context) S3BucketAssociationClassificationTypeOutput {
	return o
}

func (o S3BucketAssociationClassificationTypeOutput) ToS3BucketAssociationClassificationTypePtrOutput() S3BucketAssociationClassificationTypePtrOutput {
	return o.ToS3BucketAssociationClassificationTypePtrOutputWithContext(context.Background())
}

func (o S3BucketAssociationClassificationTypeOutput) ToS3BucketAssociationClassificationTypePtrOutputWithContext(ctx context.Context) S3BucketAssociationClassificationTypePtrOutput {
	return o.ApplyT(func(v S3BucketAssociationClassificationType) *S3BucketAssociationClassificationType {
		return &v
	}).(S3BucketAssociationClassificationTypePtrOutput)
}
// A string value indicating that Macie perform a one-time classification of all of the existing objects in the bucket.
// The only valid value is the default value, `FULL`.
func (o S3BucketAssociationClassificationTypeOutput) Continuous() pulumi.StringPtrOutput {
	return o.ApplyT(func (v S3BucketAssociationClassificationType) *string { return v.Continuous }).(pulumi.StringPtrOutput)
}

// A string value indicating whether or not Macie performs a one-time classification of all of the existing objects in the bucket.
// Valid values are `NONE` and `FULL`. Defaults to `NONE` indicating that Macie only classifies objects that are added after the association was created.
func (o S3BucketAssociationClassificationTypeOutput) OneTime() pulumi.StringPtrOutput {
	return o.ApplyT(func (v S3BucketAssociationClassificationType) *string { return v.OneTime }).(pulumi.StringPtrOutput)
}

type S3BucketAssociationClassificationTypePtrOutput struct { *pulumi.OutputState}

func (S3BucketAssociationClassificationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**S3BucketAssociationClassificationType)(nil)).Elem()
}

func (o S3BucketAssociationClassificationTypePtrOutput) ToS3BucketAssociationClassificationTypePtrOutput() S3BucketAssociationClassificationTypePtrOutput {
	return o
}

func (o S3BucketAssociationClassificationTypePtrOutput) ToS3BucketAssociationClassificationTypePtrOutputWithContext(ctx context.Context) S3BucketAssociationClassificationTypePtrOutput {
	return o
}

func (o S3BucketAssociationClassificationTypePtrOutput) Elem() S3BucketAssociationClassificationTypeOutput {
	return o.ApplyT(func (v *S3BucketAssociationClassificationType) S3BucketAssociationClassificationType { return *v }).(S3BucketAssociationClassificationTypeOutput)
}

// A string value indicating that Macie perform a one-time classification of all of the existing objects in the bucket.
// The only valid value is the default value, `FULL`.
func (o S3BucketAssociationClassificationTypePtrOutput) Continuous() pulumi.StringPtrOutput {
	return o.ApplyT(func (v S3BucketAssociationClassificationType) *string { return v.Continuous }).(pulumi.StringPtrOutput)
}

// A string value indicating whether or not Macie performs a one-time classification of all of the existing objects in the bucket.
// Valid values are `NONE` and `FULL`. Defaults to `NONE` indicating that Macie only classifies objects that are added after the association was created.
func (o S3BucketAssociationClassificationTypePtrOutput) OneTime() pulumi.StringPtrOutput {
	return o.ApplyT(func (v S3BucketAssociationClassificationType) *string { return v.OneTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(S3BucketAssociationClassificationTypeOutput{})
	pulumi.RegisterOutputType(S3BucketAssociationClassificationTypePtrOutput{})
}
