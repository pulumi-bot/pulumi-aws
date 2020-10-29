// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws
{
    public static class GetPartition
    {
        /// <summary>
        /// Use this data source to lookup current AWS partition in which this provider is working
        /// </summary>
        public static Task<GetPartitionResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPartitionResult>("aws:index/getPartition:getPartition", InvokeArgs.Empty, options.WithVersion());
    }


    [OutputType]
    public sealed class GetPartitionResult
    {
        /// <summary>
        /// Base DNS domain name for the current partition (e.g. `amazonaws.com` in AWS Commercial, `amazonaws.com.cn` in AWS China).
        /// </summary>
        public readonly string DnsSuffix;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Identifier of the current partition (e.g. `aws` in AWS Commercial, `aws-cn` in AWS China).
        /// </summary>
        public readonly string Partition;

        [OutputConstructor]
        private GetPartitionResult(
            string dnsSuffix,

            string id,

            string partition)
        {
            DnsSuffix = dnsSuffix;
            Id = id;
            Partition = partition;
        }
    }
}
