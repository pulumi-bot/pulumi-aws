// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Attaches a policy to an S3 bucket resource.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucket, err := s3.NewBucket(ctx, "bucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucketPolicy(ctx, "bucketPolicy", &s3.BucketPolicyArgs{
// 			Bucket: bucket.ID(),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Id\": \"MYBUCKETPOLICY\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"IPAllow\",\n", "      \"Effect\": \"Deny\",\n", "      \"Principal\": \"*\",\n", "      \"Action\": \"s3:*\",\n", "      \"Resource\": \"arn:aws:s3:::my_tf_test_bucket/*\",\n", "      \"Condition\": {\n", "         \"IpAddress\": {\"aws:SourceIp\": \"8.8.8.8/32\"}\n", "      }\n", "    }\n", "  ]\n", "}\n")),
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

	// The name of the bucket to which to apply the policy.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The text of the policy.
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewBucketPolicy registers a new resource with the given unique name, arguments, and options.
func NewBucketPolicy(ctx *pulumi.Context,
	name string, args *BucketPolicyArgs, opts ...pulumi.ResourceOption) (*BucketPolicy, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil {
		args = &BucketPolicyArgs{}
	}
	var resource BucketPolicy
	err := ctx.RegisterResource("aws:s3/bucketPolicy:BucketPolicy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:s3/bucketPolicy:BucketPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketPolicy resources.
type bucketPolicyState struct {
	// The name of the bucket to which to apply the policy.
	Bucket *string `pulumi:"bucket"`
	// The text of the policy.
	Policy *string `pulumi:"policy"`
}

type BucketPolicyState struct {
	// The name of the bucket to which to apply the policy.
	Bucket pulumi.StringPtrInput
	// The text of the policy.
	Policy pulumi.StringPtrInput
}

func (BucketPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPolicyState)(nil)).Elem()
}

type bucketPolicyArgs struct {
	// The name of the bucket to which to apply the policy.
	Bucket string `pulumi:"bucket"`
	// The text of the policy.
	Policy interface{} `pulumi:"policy"`
}

// The set of arguments for constructing a BucketPolicy resource.
type BucketPolicyArgs struct {
	// The name of the bucket to which to apply the policy.
	Bucket pulumi.StringInput
	// The text of the policy.
	Policy pulumi.Input
}

func (BucketPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPolicyArgs)(nil)).Elem()
}

type BucketPolicyInput interface {
	pulumi.Input

	ToBucketPolicyOutput() BucketPolicyOutput
	ToBucketPolicyOutputWithContext(ctx context.Context) BucketPolicyOutput
}

func (BucketPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketPolicy)(nil)).Elem()
}

func (i BucketPolicy) ToBucketPolicyOutput() BucketPolicyOutput {
	return i.ToBucketPolicyOutputWithContext(context.Background())
}

func (i BucketPolicy) ToBucketPolicyOutputWithContext(ctx context.Context) BucketPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPolicyOutput)
}

type BucketPolicyOutput struct {
	*pulumi.OutputState
}

func (BucketPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketPolicyOutput)(nil)).Elem()
}

func (o BucketPolicyOutput) ToBucketPolicyOutput() BucketPolicyOutput {
	return o
}

func (o BucketPolicyOutput) ToBucketPolicyOutputWithContext(ctx context.Context) BucketPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketPolicyOutput{})
}
