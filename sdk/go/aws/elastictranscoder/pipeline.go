// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elastictranscoder

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Elastic Transcoder pipeline resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/elastictranscoder"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := elastictranscoder.NewPipeline(ctx, "bar", &elastictranscoder.PipelineArgs{
// 			InputBucket: pulumi.Any(aws_s3_bucket.Input_bucket.Bucket),
// 			Role:        pulumi.Any(aws_iam_role.Test_role.Arn),
// 			ContentConfig: &elastictranscoder.PipelineContentConfigArgs{
// 				Bucket:       pulumi.Any(aws_s3_bucket.Content_bucket.Bucket),
// 				StorageClass: pulumi.String("Standard"),
// 			},
// 			ThumbnailConfig: &elastictranscoder.PipelineThumbnailConfigArgs{
// 				Bucket:       pulumi.Any(aws_s3_bucket.Thumb_bucket.Bucket),
// 				StorageClass: pulumi.String("Standard"),
// 			},
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
// Elastic Transcoder pipelines can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:elastictranscoder/pipeline:Pipeline basic_pipeline 1407981661351-cttk8b
// ```
type Pipeline struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
	AwsKmsKeyArn pulumi.StringPtrOutput `pulumi:"awsKmsKeyArn"`
	// The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
	ContentConfig PipelineContentConfigOutput `pulumi:"contentConfig"`
	// The permissions for the `contentConfig` object. (documented below)
	ContentConfigPermissions PipelineContentConfigPermissionArrayOutput `pulumi:"contentConfigPermissions"`
	// The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
	InputBucket pulumi.StringOutput `pulumi:"inputBucket"`
	// The name of the pipeline. Maximum 40 characters
	Name pulumi.StringOutput `pulumi:"name"`
	// The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
	Notifications PipelineNotificationsPtrOutput `pulumi:"notifications"`
	// The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
	OutputBucket pulumi.StringOutput `pulumi:"outputBucket"`
	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
	Role pulumi.StringOutput `pulumi:"role"`
	// The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
	ThumbnailConfig PipelineThumbnailConfigOutput `pulumi:"thumbnailConfig"`
	// The permissions for the `thumbnailConfig` object. (documented below)
	ThumbnailConfigPermissions PipelineThumbnailConfigPermissionArrayOutput `pulumi:"thumbnailConfigPermissions"`
}

// NewPipeline registers a new resource with the given unique name, arguments, and options.
func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil || args.InputBucket == nil {
		return nil, errors.New("missing required argument 'InputBucket'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &PipelineArgs{}
	}
	var resource Pipeline
	err := ctx.RegisterResource("aws:elastictranscoder/pipeline:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipeline gets an existing Pipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("aws:elastictranscoder/pipeline:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pipeline resources.
type pipelineState struct {
	Arn *string `pulumi:"arn"`
	// The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
	AwsKmsKeyArn *string `pulumi:"awsKmsKeyArn"`
	// The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
	ContentConfig *PipelineContentConfig `pulumi:"contentConfig"`
	// The permissions for the `contentConfig` object. (documented below)
	ContentConfigPermissions []PipelineContentConfigPermission `pulumi:"contentConfigPermissions"`
	// The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
	InputBucket *string `pulumi:"inputBucket"`
	// The name of the pipeline. Maximum 40 characters
	Name *string `pulumi:"name"`
	// The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
	Notifications *PipelineNotifications `pulumi:"notifications"`
	// The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
	OutputBucket *string `pulumi:"outputBucket"`
	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
	Role *string `pulumi:"role"`
	// The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
	ThumbnailConfig *PipelineThumbnailConfig `pulumi:"thumbnailConfig"`
	// The permissions for the `thumbnailConfig` object. (documented below)
	ThumbnailConfigPermissions []PipelineThumbnailConfigPermission `pulumi:"thumbnailConfigPermissions"`
}

type PipelineState struct {
	Arn pulumi.StringPtrInput
	// The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
	AwsKmsKeyArn pulumi.StringPtrInput
	// The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
	ContentConfig PipelineContentConfigPtrInput
	// The permissions for the `contentConfig` object. (documented below)
	ContentConfigPermissions PipelineContentConfigPermissionArrayInput
	// The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
	InputBucket pulumi.StringPtrInput
	// The name of the pipeline. Maximum 40 characters
	Name pulumi.StringPtrInput
	// The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
	Notifications PipelineNotificationsPtrInput
	// The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
	OutputBucket pulumi.StringPtrInput
	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
	Role pulumi.StringPtrInput
	// The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
	ThumbnailConfig PipelineThumbnailConfigPtrInput
	// The permissions for the `thumbnailConfig` object. (documented below)
	ThumbnailConfigPermissions PipelineThumbnailConfigPermissionArrayInput
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	// The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
	AwsKmsKeyArn *string `pulumi:"awsKmsKeyArn"`
	// The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
	ContentConfig *PipelineContentConfig `pulumi:"contentConfig"`
	// The permissions for the `contentConfig` object. (documented below)
	ContentConfigPermissions []PipelineContentConfigPermission `pulumi:"contentConfigPermissions"`
	// The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
	InputBucket string `pulumi:"inputBucket"`
	// The name of the pipeline. Maximum 40 characters
	Name *string `pulumi:"name"`
	// The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
	Notifications *PipelineNotifications `pulumi:"notifications"`
	// The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
	OutputBucket *string `pulumi:"outputBucket"`
	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
	Role string `pulumi:"role"`
	// The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
	ThumbnailConfig *PipelineThumbnailConfig `pulumi:"thumbnailConfig"`
	// The permissions for the `thumbnailConfig` object. (documented below)
	ThumbnailConfigPermissions []PipelineThumbnailConfigPermission `pulumi:"thumbnailConfigPermissions"`
}

// The set of arguments for constructing a Pipeline resource.
type PipelineArgs struct {
	// The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
	AwsKmsKeyArn pulumi.StringPtrInput
	// The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
	ContentConfig PipelineContentConfigPtrInput
	// The permissions for the `contentConfig` object. (documented below)
	ContentConfigPermissions PipelineContentConfigPermissionArrayInput
	// The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
	InputBucket pulumi.StringInput
	// The name of the pipeline. Maximum 40 characters
	Name pulumi.StringPtrInput
	// The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
	Notifications PipelineNotificationsPtrInput
	// The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
	OutputBucket pulumi.StringPtrInput
	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
	Role pulumi.StringInput
	// The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
	ThumbnailConfig PipelineThumbnailConfigPtrInput
	// The permissions for the `thumbnailConfig` object. (documented below)
	ThumbnailConfigPermissions PipelineThumbnailConfigPermissionArrayInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}
