// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a S3 bucket [metrics configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html) resource.
//
// ## Example Usage
// ### Add metrics configuration for entire S3 bucket
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
// 		example, err := s3.NewBucket(ctx, "example", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucketMetric(ctx, "example_entire_bucket", &s3.BucketMetricArgs{
// 			Bucket: example.Bucket,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Add metrics configuration with S3 bucket object filter
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
// 		example, err := s3.NewBucket(ctx, "example", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucketMetric(ctx, "example_filtered", &s3.BucketMetricArgs{
// 			Bucket: example.Bucket,
// 			Filter: &s3.BucketMetricFilterArgs{
// 				Prefix: pulumi.String("documents/"),
// 				Tags: pulumi.StringMap{
// 					"priority": pulumi.String("high"),
// 					"class":    pulumi.String("blue"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type BucketMetric struct {
	pulumi.CustomResourceState

	// The name of the bucket to put metric configuration.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter BucketMetricFilterPtrOutput `pulumi:"filter"`
	// Unique identifier of the metrics configuration for the bucket.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewBucketMetric registers a new resource with the given unique name, arguments, and options.
func NewBucketMetric(ctx *pulumi.Context,
	name string, args *BucketMetricArgs, opts ...pulumi.ResourceOption) (*BucketMetric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	var resource BucketMetric
	err := ctx.RegisterResource("aws:s3/bucketMetric:BucketMetric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketMetric gets an existing BucketMetric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketMetric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketMetricState, opts ...pulumi.ResourceOption) (*BucketMetric, error) {
	var resource BucketMetric
	err := ctx.ReadResource("aws:s3/bucketMetric:BucketMetric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketMetric resources.
type bucketMetricState struct {
	// The name of the bucket to put metric configuration.
	Bucket *string `pulumi:"bucket"`
	// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter *BucketMetricFilter `pulumi:"filter"`
	// Unique identifier of the metrics configuration for the bucket.
	Name *string `pulumi:"name"`
}

type BucketMetricState struct {
	// The name of the bucket to put metric configuration.
	Bucket pulumi.StringPtrInput
	// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter BucketMetricFilterPtrInput
	// Unique identifier of the metrics configuration for the bucket.
	Name pulumi.StringPtrInput
}

func (BucketMetricState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketMetricState)(nil)).Elem()
}

type bucketMetricArgs struct {
	// The name of the bucket to put metric configuration.
	Bucket string `pulumi:"bucket"`
	// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter *BucketMetricFilter `pulumi:"filter"`
	// Unique identifier of the metrics configuration for the bucket.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a BucketMetric resource.
type BucketMetricArgs struct {
	// The name of the bucket to put metric configuration.
	Bucket pulumi.StringInput
	// [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter BucketMetricFilterPtrInput
	// Unique identifier of the metrics configuration for the bucket.
	Name pulumi.StringPtrInput
}

func (BucketMetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketMetricArgs)(nil)).Elem()
}

type BucketMetricInput interface {
	pulumi.Input

	ToBucketMetricOutput() BucketMetricOutput
	ToBucketMetricOutputWithContext(ctx context.Context) BucketMetricOutput
}

func (BucketMetric) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketMetric)(nil)).Elem()
}

func (i BucketMetric) ToBucketMetricOutput() BucketMetricOutput {
	return i.ToBucketMetricOutputWithContext(context.Background())
}

func (i BucketMetric) ToBucketMetricOutputWithContext(ctx context.Context) BucketMetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketMetricOutput)
}

type BucketMetricOutput struct {
	*pulumi.OutputState
}

func (BucketMetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketMetricOutput)(nil)).Elem()
}

func (o BucketMetricOutput) ToBucketMetricOutput() BucketMetricOutput {
	return o
}

func (o BucketMetricOutput) ToBucketMetricOutputWithContext(ctx context.Context) BucketMetricOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketMetricOutput{})
}
