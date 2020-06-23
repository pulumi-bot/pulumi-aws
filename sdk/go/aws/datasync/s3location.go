// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an S3 Location within AWS DataSync.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/datasync"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = datasync.NewS3Location(ctx, "example", &datasync.S3LocationArgs{
// 			S3BucketArn: pulumi.String(aws_s3_bucket.Example.Arn),
// 			S3Config: &datasync.S3LocationS3ConfigArgs{
// 				BucketAccessRoleArn: pulumi.String(aws_iam_role.Example.Arn),
// 			},
// 			Subdirectory: pulumi.String("/example/prefix"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type S3Location struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Amazon Resource Name (ARN) of the S3 Bucket.
	S3BucketArn pulumi.StringOutput `pulumi:"s3BucketArn"`
	// Configuration block containing information for connecting to S3.
	S3Config S3LocationS3ConfigOutput `pulumi:"s3Config"`
	// Prefix to perform actions as source or destination.
	Subdirectory pulumi.StringOutput `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.MapOutput    `pulumi:"tags"`
	Uri  pulumi.StringOutput `pulumi:"uri"`
}

// NewS3Location registers a new resource with the given unique name, arguments, and options.
func NewS3Location(ctx *pulumi.Context,
	name string, args *S3LocationArgs, opts ...pulumi.ResourceOption) (*S3Location, error) {
	if args == nil || args.S3BucketArn == nil {
		return nil, errors.New("missing required argument 'S3BucketArn'")
	}
	if args == nil || args.S3Config == nil {
		return nil, errors.New("missing required argument 'S3Config'")
	}
	if args == nil || args.Subdirectory == nil {
		return nil, errors.New("missing required argument 'Subdirectory'")
	}
	if args == nil {
		args = &S3LocationArgs{}
	}
	var resource S3Location
	err := ctx.RegisterResource("aws:datasync/s3Location:S3Location", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetS3Location gets an existing S3Location resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetS3Location(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *S3LocationState, opts ...pulumi.ResourceOption) (*S3Location, error) {
	var resource S3Location
	err := ctx.ReadResource("aws:datasync/s3Location:S3Location", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering S3Location resources.
type s3locationState struct {
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn *string `pulumi:"arn"`
	// Amazon Resource Name (ARN) of the S3 Bucket.
	S3BucketArn *string `pulumi:"s3BucketArn"`
	// Configuration block containing information for connecting to S3.
	S3Config *S3LocationS3Config `pulumi:"s3Config"`
	// Prefix to perform actions as source or destination.
	Subdirectory *string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags map[string]interface{} `pulumi:"tags"`
	Uri  *string                `pulumi:"uri"`
}

type S3LocationState struct {
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the S3 Bucket.
	S3BucketArn pulumi.StringPtrInput
	// Configuration block containing information for connecting to S3.
	S3Config S3LocationS3ConfigPtrInput
	// Prefix to perform actions as source or destination.
	Subdirectory pulumi.StringPtrInput
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.MapInput
	Uri  pulumi.StringPtrInput
}

func (S3LocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*s3locationState)(nil)).Elem()
}

type s3locationArgs struct {
	// Amazon Resource Name (ARN) of the S3 Bucket.
	S3BucketArn string `pulumi:"s3BucketArn"`
	// Configuration block containing information for connecting to S3.
	S3Config S3LocationS3Config `pulumi:"s3Config"`
	// Prefix to perform actions as source or destination.
	Subdirectory string `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a S3Location resource.
type S3LocationArgs struct {
	// Amazon Resource Name (ARN) of the S3 Bucket.
	S3BucketArn pulumi.StringInput
	// Configuration block containing information for connecting to S3.
	S3Config S3LocationS3ConfigInput
	// Prefix to perform actions as source or destination.
	Subdirectory pulumi.StringInput
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.MapInput
}

func (S3LocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*s3locationArgs)(nil)).Elem()
}
