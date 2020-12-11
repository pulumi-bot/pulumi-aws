// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3Control
{
    /// <summary>
    /// Provides a resource to manage an S3 Control Bucket Policy.
    /// 
    /// &gt; This functionality is for managing [S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html). To manage S3 Bucket Policies in an AWS Partition, see the [`aws.s3.BucketPolicy` resource](https://www.terraform.io/docs/providers/aws/r/s3_bucket_policy.html).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.S3Control.BucketPolicy("example", new Aws.S3Control.BucketPolicyArgs
    ///         {
    ///             Bucket = aws_s3control_bucket.Example.Arn,
    ///             Policy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 { "Id", "testBucketPolicy" },
    ///                 { "Statement", new[]
    ///                     {
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             { "Action", "s3-outposts:PutBucketLifecycleConfiguration" },
    ///                             { "Effect", "Deny" },
    ///                             { "Principal", new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 { "AWS", "*" },
    ///                             } },
    ///                             { "Resource", aws_s3control_bucket.Example.Arn },
    ///                             { "Sid", "statement1" },
    ///                         },
    ///                     }
    ///                  },
    ///                 { "Version", "2012-10-17" },
    ///             }),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// S3 Control Bucket Policies can be imported using the Amazon Resource Name (ARN), e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:s3control/bucketPolicy:BucketPolicy example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/bucket/example
    /// ```
    /// </summary>
    [AwsResourceType("aws:s3control/bucketPolicy:BucketPolicy")]
    public partial class BucketPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the bucket.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;


        /// <summary>
        /// Create a BucketPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketPolicy(string name, BucketPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:s3control/bucketPolicy:BucketPolicy", name, args ?? new BucketPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketPolicy(string name, Input<string> id, BucketPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:s3control/bucketPolicy:BucketPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketPolicy Get(string name, Input<string> id, BucketPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketPolicy(name, id, state, options);
        }
    }

    public sealed class BucketPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        public BucketPolicyArgs()
        {
        }
    }

    public sealed class BucketPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the bucket.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public BucketPolicyState()
        {
        }
    }
}
