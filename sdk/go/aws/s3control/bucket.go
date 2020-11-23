// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3control

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage an S3 Control Bucket.
//
// > This functionality is for managing [S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html). To manage S3 Buckets in an AWS Partition, see the [`s3.Bucket` resource](https://www.terraform.io/docs/providers/aws/r/s3_bucket.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3control"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := s3control.NewBucket(ctx, "example", &s3control.BucketArgs{
// 			Bucket:    pulumi.String("example"),
// 			OutpostId: pulumi.Any(data.Aws_outposts_outpost.Example.Id),
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
// S3 Control Buckets can be imported using Amazon Resource Name (ARN), e.g.
//
// ```sh
//  $ pulumi import aws:s3control/bucket:Bucket example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/bucket/example
// ```
type Bucket struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the bucket.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// UTC creation date in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Identifier of the Outpost to contain this bucket.
	OutpostId pulumi.StringOutput `pulumi:"outpostId"`
	// Boolean whether Public Access Block is enabled.
	PublicAccessBlockEnabled pulumi.BoolOutput `pulumi:"publicAccessBlockEnabled"`
	// Key-value map of resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.OutpostId == nil {
		return nil, errors.New("invalid value for required argument 'OutpostId'")
	}
	var resource Bucket
	err := ctx.RegisterResource("aws:s3control/bucket:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("aws:s3control/bucket:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
	// Amazon Resource Name (ARN) of the bucket.
	Arn *string `pulumi:"arn"`
	// Name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// UTC creation date in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	CreationDate *string `pulumi:"creationDate"`
	// Identifier of the Outpost to contain this bucket.
	OutpostId *string `pulumi:"outpostId"`
	// Boolean whether Public Access Block is enabled.
	PublicAccessBlockEnabled *bool `pulumi:"publicAccessBlockEnabled"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

type BucketState struct {
	// Amazon Resource Name (ARN) of the bucket.
	Arn pulumi.StringPtrInput
	// Name of the bucket.
	Bucket pulumi.StringPtrInput
	// UTC creation date in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
	CreationDate pulumi.StringPtrInput
	// Identifier of the Outpost to contain this bucket.
	OutpostId pulumi.StringPtrInput
	// Boolean whether Public Access Block is enabled.
	PublicAccessBlockEnabled pulumi.BoolPtrInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// Name of the bucket.
	Bucket string `pulumi:"bucket"`
	// Identifier of the Outpost to contain this bucket.
	OutpostId string `pulumi:"outpostId"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// Name of the bucket.
	Bucket pulumi.StringInput
	// Identifier of the Outpost to contain this bucket.
	OutpostId pulumi.StringInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((*Bucket)(nil)).Elem()
}

func (i Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

type BucketOutput struct {
	*pulumi.OutputState
}

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketOutput)(nil)).Elem()
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketOutput{})
}
