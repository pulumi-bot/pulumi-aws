// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages S3 bucket-level Public Access Block configuration. For more information about these settings, see the [AWS S3 Block Public Access documentation](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleBucket, err := s3.NewBucket(ctx, "exampleBucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucketPublicAccessBlock(ctx, "exampleBucketPublicAccessBlock", &s3.BucketPublicAccessBlockArgs{
// 			Bucket:            exampleBucket.ID(),
// 			BlockPublicAcls:   pulumi.Bool(true),
// 			BlockPublicPolicy: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// `aws_s3_bucket_public_access_block` can be imported by using the bucket name, e.g.
//
// ```sh
//  $ pulumi import aws:s3/bucketPublicAccessBlock:BucketPublicAccessBlock example my-bucket
// ```
type BucketPublicAccessBlock struct {
	pulumi.CustomResourceState

	// Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls pulumi.BoolPtrOutput `pulumi:"blockPublicAcls"`
	// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy pulumi.BoolPtrOutput `pulumi:"blockPublicPolicy"`
	// S3 Bucket to which this Public Access Block configuration should be applied.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore public ACLs on this bucket and any objects that it contains.
	IgnorePublicAcls pulumi.BoolPtrOutput `pulumi:"ignorePublicAcls"`
	// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
	RestrictPublicBuckets pulumi.BoolPtrOutput `pulumi:"restrictPublicBuckets"`
}

// NewBucketPublicAccessBlock registers a new resource with the given unique name, arguments, and options.
func NewBucketPublicAccessBlock(ctx *pulumi.Context,
	name string, args *BucketPublicAccessBlockArgs, opts ...pulumi.ResourceOption) (*BucketPublicAccessBlock, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil {
		args = &BucketPublicAccessBlockArgs{}
	}
	var resource BucketPublicAccessBlock
	err := ctx.RegisterResource("aws:s3/bucketPublicAccessBlock:BucketPublicAccessBlock", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketPublicAccessBlock gets an existing BucketPublicAccessBlock resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketPublicAccessBlock(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketPublicAccessBlockState, opts ...pulumi.ResourceOption) (*BucketPublicAccessBlock, error) {
	var resource BucketPublicAccessBlock
	err := ctx.ReadResource("aws:s3/bucketPublicAccessBlock:BucketPublicAccessBlock", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketPublicAccessBlock resources.
type bucketPublicAccessBlockState struct {
	// Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls *bool `pulumi:"blockPublicAcls"`
	// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy *bool `pulumi:"blockPublicPolicy"`
	// S3 Bucket to which this Public Access Block configuration should be applied.
	Bucket *string `pulumi:"bucket"`
	// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore public ACLs on this bucket and any objects that it contains.
	IgnorePublicAcls *bool `pulumi:"ignorePublicAcls"`
	// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
	RestrictPublicBuckets *bool `pulumi:"restrictPublicBuckets"`
}

type BucketPublicAccessBlockState struct {
	// Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls pulumi.BoolPtrInput
	// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy pulumi.BoolPtrInput
	// S3 Bucket to which this Public Access Block configuration should be applied.
	Bucket pulumi.StringPtrInput
	// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore public ACLs on this bucket and any objects that it contains.
	IgnorePublicAcls pulumi.BoolPtrInput
	// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
	RestrictPublicBuckets pulumi.BoolPtrInput
}

func (BucketPublicAccessBlockState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPublicAccessBlockState)(nil)).Elem()
}

type bucketPublicAccessBlockArgs struct {
	// Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls *bool `pulumi:"blockPublicAcls"`
	// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy *bool `pulumi:"blockPublicPolicy"`
	// S3 Bucket to which this Public Access Block configuration should be applied.
	Bucket string `pulumi:"bucket"`
	// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore public ACLs on this bucket and any objects that it contains.
	IgnorePublicAcls *bool `pulumi:"ignorePublicAcls"`
	// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
	RestrictPublicBuckets *bool `pulumi:"restrictPublicBuckets"`
}

// The set of arguments for constructing a BucketPublicAccessBlock resource.
type BucketPublicAccessBlockArgs struct {
	// Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls pulumi.BoolPtrInput
	// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy pulumi.BoolPtrInput
	// S3 Bucket to which this Public Access Block configuration should be applied.
	Bucket pulumi.StringInput
	// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore public ACLs on this bucket and any objects that it contains.
	IgnorePublicAcls pulumi.BoolPtrInput
	// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
	RestrictPublicBuckets pulumi.BoolPtrInput
}

func (BucketPublicAccessBlockArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPublicAccessBlockArgs)(nil)).Elem()
}

type BucketPublicAccessBlockInput interface {
	pulumi.Input

	ToBucketPublicAccessBlockOutput() BucketPublicAccessBlockOutput
	ToBucketPublicAccessBlockOutputWithContext(ctx context.Context) BucketPublicAccessBlockOutput
}

func (BucketPublicAccessBlock) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketPublicAccessBlock)(nil)).Elem()
}

func (i BucketPublicAccessBlock) ToBucketPublicAccessBlockOutput() BucketPublicAccessBlockOutput {
	return i.ToBucketPublicAccessBlockOutputWithContext(context.Background())
}

func (i BucketPublicAccessBlock) ToBucketPublicAccessBlockOutputWithContext(ctx context.Context) BucketPublicAccessBlockOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPublicAccessBlockOutput)
}

type BucketPublicAccessBlockOutput struct {
	*pulumi.OutputState
}

func (BucketPublicAccessBlockOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketPublicAccessBlockOutput)(nil)).Elem()
}

func (o BucketPublicAccessBlockOutput) ToBucketPublicAccessBlockOutput() BucketPublicAccessBlockOutput {
	return o
}

func (o BucketPublicAccessBlockOutput) ToBucketPublicAccessBlockOutputWithContext(ctx context.Context) BucketPublicAccessBlockOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketPublicAccessBlockOutput{})
}
