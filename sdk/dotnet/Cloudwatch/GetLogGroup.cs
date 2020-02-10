// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get information about an AWS Cloudwatch Log Group
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cloudwatch_log_group.html.markdown.
        /// </summary>
        public static Task<GetLogGroupResult> GetLogGroup(GetLogGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLogGroupResult>("aws:cloudwatch/getLogGroup:getLogGroup", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetLogGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Cloudwatch log group
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, object>? _tags;
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetLogGroupArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetLogGroupResult
    {
        /// <summary>
        /// The ARN of the Cloudwatch log group
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.
        /// </summary>
        public readonly int CreationTime;
        /// <summary>
        /// The ARN of the KMS Key to use when encrypting log data.
        /// </summary>
        public readonly string KmsKeyId;
        public readonly string Name;
        /// <summary>
        /// The number of days log events retained in the specified log group.
        /// </summary>
        public readonly int RetentionInDays;
        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetLogGroupResult(
            string arn,
            int creationTime,
            string kmsKeyId,
            string name,
            int retentionInDays,
            ImmutableDictionary<string, object> tags,
            string id)
        {
            Arn = arn;
            CreationTime = creationTime;
            KmsKeyId = kmsKeyId;
            Name = name;
            RetentionInDays = retentionInDays;
            Tags = tags;
            Id = id;
        }
    }
}
