// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sqs
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the ARN and URL of queue in AWS Simple Queue Service (SQS).
        /// By using this data source, you can reference SQS queues without having to hardcode
        /// the ARNs as input.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/sqs_queue.html.markdown.
        /// </summary>
        public static Task<GetQueueResult> GetQueue(GetQueueArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetQueueResult>("aws:sqs/getQueue:getQueue", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetQueueArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the queue to match.
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

        public GetQueueArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetQueueResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the queue.
        /// </summary>
        public readonly string Arn;
        public readonly string Name;
        /// <summary>
        /// A mapping of tags for the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The URL of the queue.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetQueueResult(
            string arn,
            string name,
            ImmutableDictionary<string, object> tags,
            string url,
            string id)
        {
            Arn = arn;
            Name = name;
            Tags = tags;
            Url = url;
            Id = id;
        }
    }
}
