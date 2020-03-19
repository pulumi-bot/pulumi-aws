// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis
{
    /// <summary>
    /// Provides a Kinesis Stream resource. Amazon Kinesis is a managed service that
    /// scales elastically for real-time processing of streaming big data.
    /// 
    /// For more details, see the [Amazon Kinesis Documentation][1].
    /// 
    /// {{% examples %}}
    /// {{% /examples %}}
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/kinesis_stream.html.markdown.
    /// </summary>
    public partial class Stream : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
        /// </summary>
        [Output("encryptionType")]
        public Output<string?> EncryptionType { get; private set; } = null!;

        /// <summary>
        /// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
        /// </summary>
        [Output("enforceConsumerDeletion")]
        public Output<bool?> EnforceConsumerDeletion { get; private set; } = null!;

        /// <summary>
        /// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24.
        /// </summary>
        [Output("retentionPeriod")]
        public Output<int?> RetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// The number of shards that the stream will use.
        /// Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams][2] for more.
        /// </summary>
        [Output("shardCount")]
        public Output<int> ShardCount { get; private set; } = null!;

        /// <summary>
        /// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch][3] for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
        /// </summary>
        [Output("shardLevelMetrics")]
        public Output<ImmutableArray<string>> ShardLevelMetrics { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Stream resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Stream(string name, StreamArgs args, CustomResourceOptions? options = null)
            : base("aws:kinesis/stream:Stream", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Stream(string name, Input<string> id, StreamState? state = null, CustomResourceOptions? options = null)
            : base("aws:kinesis/stream:Stream", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Stream resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Stream Get(string name, Input<string> id, StreamState? state = null, CustomResourceOptions? options = null)
        {
            return new Stream(name, id, state, options);
        }
    }

    public sealed class StreamArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
        /// </summary>
        [Input("encryptionType")]
        public Input<string>? EncryptionType { get; set; }

        /// <summary>
        /// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
        /// </summary>
        [Input("enforceConsumerDeletion")]
        public Input<bool>? EnforceConsumerDeletion { get; set; }

        /// <summary>
        /// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24.
        /// </summary>
        [Input("retentionPeriod")]
        public Input<int>? RetentionPeriod { get; set; }

        /// <summary>
        /// The number of shards that the stream will use.
        /// Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams][2] for more.
        /// </summary>
        [Input("shardCount", required: true)]
        public Input<int> ShardCount { get; set; } = null!;

        [Input("shardLevelMetrics")]
        private InputList<string>? _shardLevelMetrics;

        /// <summary>
        /// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch][3] for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
        /// </summary>
        public InputList<string> ShardLevelMetrics
        {
            get => _shardLevelMetrics ?? (_shardLevelMetrics = new InputList<string>());
            set => _shardLevelMetrics = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public StreamArgs()
        {
        }
    }

    public sealed class StreamState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
        /// </summary>
        [Input("encryptionType")]
        public Input<string>? EncryptionType { get; set; }

        /// <summary>
        /// A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
        /// </summary>
        [Input("enforceConsumerDeletion")]
        public Input<bool>? EnforceConsumerDeletion { get; set; }

        /// <summary>
        /// The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24.
        /// </summary>
        [Input("retentionPeriod")]
        public Input<int>? RetentionPeriod { get; set; }

        /// <summary>
        /// The number of shards that the stream will use.
        /// Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams][2] for more.
        /// </summary>
        [Input("shardCount")]
        public Input<int>? ShardCount { get; set; }

        [Input("shardLevelMetrics")]
        private InputList<string>? _shardLevelMetrics;

        /// <summary>
        /// A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch][3] for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
        /// </summary>
        public InputList<string> ShardLevelMetrics
        {
            get => _shardLevelMetrics ?? (_shardLevelMetrics = new InputList<string>());
            set => _shardLevelMetrics = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public StreamState()
        {
        }
    }
}
