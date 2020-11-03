// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3control

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage an S3 Control Bucket Policy.
//
// > This functionality is for managing [S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html). To manage S3 Bucket Policies in an AWS Partition, see the [`s3.BucketPolicy` resource](https://www.terraform.io/docs/providers/aws/r/s3_bucket_policy.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3control"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"Id": "testBucketPolicy",
// 			"Statement": []map[string]interface{}{
// 				map[string]interface{}{
// 					"Action": "s3-outposts:PutBucketLifecycleConfiguration",
// 					"Effect": "Deny",
// 					"Principal": map[string]interface{}{
// 						"AWS": "*",
// 					},
// 					"Resource": aws_s3control_bucket.Example.Arn,
// 					"Sid":      "statement1",
// 				},
// 			},
// 			"Version": "2012-10-17",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		_, err := s3control.NewBucketPolicy(ctx, "example", &s3control.BucketPolicyArgs{
// 			Bucket: pulumi.Any(aws_s3control_bucket.Example.Arn),
// 			Policy: pulumi.String(json0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type BucketPolicy struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewBucketPolicy registers a new resource with the given unique name, arguments, and options.
func NewBucketPolicy(ctx *pulumi.Context,
	name string, args *BucketPolicyArgs, opts ...pulumi.ResourceOption) (*BucketPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource BucketPolicy
	err := ctx.RegisterResource("aws:s3control/bucketPolicy:BucketPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketPolicy gets an existing BucketPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketPolicyState, opts ...pulumi.ResourceOption) (*BucketPolicy, error) {
	var resource BucketPolicy
	err := ctx.ReadResource("aws:s3control/bucketPolicy:BucketPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketPolicy resources.
type bucketPolicyState struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket *string `pulumi:"bucket"`
	Policy *string `pulumi:"policy"`
}

type BucketPolicyState struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket pulumi.StringPtrInput
	Policy pulumi.StringPtrInput
}

func (BucketPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPolicyState)(nil)).Elem()
}

type bucketPolicyArgs struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket string `pulumi:"bucket"`
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a BucketPolicy resource.
type BucketPolicyArgs struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket pulumi.StringInput
	Policy pulumi.StringInput
}

func (BucketPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPolicyArgs)(nil)).Elem()
}
