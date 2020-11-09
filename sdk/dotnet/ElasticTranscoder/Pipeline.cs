// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticTranscoder
{
    /// <summary>
    /// Provides an Elastic Transcoder pipeline resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bar = new Aws.ElasticTranscoder.Pipeline("bar", new Aws.ElasticTranscoder.PipelineArgs
    ///         {
    ///             InputBucket = aws_s3_bucket.Input_bucket.Bucket,
    ///             Role = aws_iam_role.Test_role.Arn,
    ///             ContentConfig = new Aws.ElasticTranscoder.Inputs.PipelineContentConfigArgs
    ///             {
    ///                 Bucket = aws_s3_bucket.Content_bucket.Bucket,
    ///                 StorageClass = "Standard",
    ///             },
    ///             ThumbnailConfig = new Aws.ElasticTranscoder.Inputs.PipelineThumbnailConfigArgs
    ///             {
    ///                 Bucket = aws_s3_bucket.Thumb_bucket.Bucket,
    ///                 StorageClass = "Standard",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Elastic Transcoder pipelines can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:elastictranscoder/pipeline:Pipeline basic_pipeline 1407981661351-cttk8b
    /// ```
    /// </summary>
    public partial class Pipeline : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        /// </summary>
        [Output("awsKmsKeyArn")]
        public Output<string?> AwsKmsKeyArn { get; private set; } = null!;

        /// <summary>
        /// The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        /// </summary>
        [Output("contentConfig")]
        public Output<Outputs.PipelineContentConfig> ContentConfig { get; private set; } = null!;

        /// <summary>
        /// The permissions for the `content_config` object. (documented below)
        /// </summary>
        [Output("contentConfigPermissions")]
        public Output<ImmutableArray<Outputs.PipelineContentConfigPermission>> ContentConfigPermissions { get; private set; } = null!;

        /// <summary>
        /// The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        /// </summary>
        [Output("inputBucket")]
        public Output<string> InputBucket { get; private set; } = null!;

        /// <summary>
        /// The name of the pipeline. Maximum 40 characters
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        /// </summary>
        [Output("notifications")]
        public Output<Outputs.PipelineNotifications?> Notifications { get; private set; } = null!;

        /// <summary>
        /// The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        /// </summary>
        [Output("outputBucket")]
        public Output<string> OutputBucket { get; private set; } = null!;

        /// <summary>
        /// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        /// </summary>
        [Output("thumbnailConfig")]
        public Output<Outputs.PipelineThumbnailConfig> ThumbnailConfig { get; private set; } = null!;

        /// <summary>
        /// The permissions for the `thumbnail_config` object. (documented below)
        /// </summary>
        [Output("thumbnailConfigPermissions")]
        public Output<ImmutableArray<Outputs.PipelineThumbnailConfigPermission>> ThumbnailConfigPermissions { get; private set; } = null!;


        /// <summary>
        /// Create a Pipeline resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Pipeline(string name, PipelineArgs args, CustomResourceOptions? options = null)
            : base("aws:elastictranscoder/pipeline:Pipeline", name, args ?? new PipelineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Pipeline(string name, Input<string> id, PipelineState? state = null, CustomResourceOptions? options = null)
            : base("aws:elastictranscoder/pipeline:Pipeline", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Pipeline resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Pipeline Get(string name, Input<string> id, PipelineState? state = null, CustomResourceOptions? options = null)
        {
            return new Pipeline(name, id, state, options);
        }
    }

    public sealed class PipelineArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        /// </summary>
        [Input("awsKmsKeyArn")]
        public Input<string>? AwsKmsKeyArn { get; set; }

        /// <summary>
        /// The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        /// </summary>
        [Input("contentConfig")]
        public Input<Inputs.PipelineContentConfigArgs>? ContentConfig { get; set; }

        [Input("contentConfigPermissions")]
        private InputList<Inputs.PipelineContentConfigPermissionArgs>? _contentConfigPermissions;

        /// <summary>
        /// The permissions for the `content_config` object. (documented below)
        /// </summary>
        public InputList<Inputs.PipelineContentConfigPermissionArgs> ContentConfigPermissions
        {
            get => _contentConfigPermissions ?? (_contentConfigPermissions = new InputList<Inputs.PipelineContentConfigPermissionArgs>());
            set => _contentConfigPermissions = value;
        }

        /// <summary>
        /// The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        /// </summary>
        [Input("inputBucket", required: true)]
        public Input<string> InputBucket { get; set; } = null!;

        /// <summary>
        /// The name of the pipeline. Maximum 40 characters
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        /// </summary>
        [Input("notifications")]
        public Input<Inputs.PipelineNotificationsArgs>? Notifications { get; set; }

        /// <summary>
        /// The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        /// </summary>
        [Input("outputBucket")]
        public Input<string>? OutputBucket { get; set; }

        /// <summary>
        /// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        /// </summary>
        [Input("thumbnailConfig")]
        public Input<Inputs.PipelineThumbnailConfigArgs>? ThumbnailConfig { get; set; }

        [Input("thumbnailConfigPermissions")]
        private InputList<Inputs.PipelineThumbnailConfigPermissionArgs>? _thumbnailConfigPermissions;

        /// <summary>
        /// The permissions for the `thumbnail_config` object. (documented below)
        /// </summary>
        public InputList<Inputs.PipelineThumbnailConfigPermissionArgs> ThumbnailConfigPermissions
        {
            get => _thumbnailConfigPermissions ?? (_thumbnailConfigPermissions = new InputList<Inputs.PipelineThumbnailConfigPermissionArgs>());
            set => _thumbnailConfigPermissions = value;
        }

        public PipelineArgs()
        {
        }
    }

    public sealed class PipelineState : Pulumi.ResourceArgs
    {
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        /// </summary>
        [Input("awsKmsKeyArn")]
        public Input<string>? AwsKmsKeyArn { get; set; }

        /// <summary>
        /// The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        /// </summary>
        [Input("contentConfig")]
        public Input<Inputs.PipelineContentConfigGetArgs>? ContentConfig { get; set; }

        [Input("contentConfigPermissions")]
        private InputList<Inputs.PipelineContentConfigPermissionGetArgs>? _contentConfigPermissions;

        /// <summary>
        /// The permissions for the `content_config` object. (documented below)
        /// </summary>
        public InputList<Inputs.PipelineContentConfigPermissionGetArgs> ContentConfigPermissions
        {
            get => _contentConfigPermissions ?? (_contentConfigPermissions = new InputList<Inputs.PipelineContentConfigPermissionGetArgs>());
            set => _contentConfigPermissions = value;
        }

        /// <summary>
        /// The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        /// </summary>
        [Input("inputBucket")]
        public Input<string>? InputBucket { get; set; }

        /// <summary>
        /// The name of the pipeline. Maximum 40 characters
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        /// </summary>
        [Input("notifications")]
        public Input<Inputs.PipelineNotificationsGetArgs>? Notifications { get; set; }

        /// <summary>
        /// The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        /// </summary>
        [Input("outputBucket")]
        public Input<string>? OutputBucket { get; set; }

        /// <summary>
        /// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        /// </summary>
        [Input("thumbnailConfig")]
        public Input<Inputs.PipelineThumbnailConfigGetArgs>? ThumbnailConfig { get; set; }

        [Input("thumbnailConfigPermissions")]
        private InputList<Inputs.PipelineThumbnailConfigPermissionGetArgs>? _thumbnailConfigPermissions;

        /// <summary>
        /// The permissions for the `thumbnail_config` object. (documented below)
        /// </summary>
        public InputList<Inputs.PipelineThumbnailConfigPermissionGetArgs> ThumbnailConfigPermissions
        {
            get => _thumbnailConfigPermissions ?? (_thumbnailConfigPermissions = new InputList<Inputs.PipelineThumbnailConfigPermissionGetArgs>());
            set => _thumbnailConfigPermissions = value;
        }

        public PipelineState()
        {
        }
    }
}
