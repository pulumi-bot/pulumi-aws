// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElastiCache
{
    public static class GetReplicationGroup
    {
        public static Task<GetReplicationGroupResult> InvokeAsync(GetReplicationGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReplicationGroupResult>("aws:elasticache/getReplicationGroup:getReplicationGroup", args ?? new GetReplicationGroupArgs(), options.WithVersion());
    }


    public sealed class GetReplicationGroupArgs : Pulumi.InvokeArgs
    {
        [Input("replicationGroupId", required: true)]
        public string ReplicationGroupId { get; set; } = null!;

        public GetReplicationGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetReplicationGroupResult
    {
        public readonly bool AuthTokenEnabled;
        public readonly bool AutomaticFailoverEnabled;
        public readonly string ConfigurationEndpointAddress;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> MemberClusters;
        public readonly string NodeType;
        public readonly int NumberCacheClusters;
        public readonly int Port;
        public readonly string PrimaryEndpointAddress;
        public readonly string ReplicationGroupDescription;
        public readonly string ReplicationGroupId;
        public readonly int SnapshotRetentionLimit;
        public readonly string SnapshotWindow;

        [OutputConstructor]
        private GetReplicationGroupResult(
            bool authTokenEnabled,

            bool automaticFailoverEnabled,

            string configurationEndpointAddress,

            string id,

            ImmutableArray<string> memberClusters,

            string nodeType,

            int numberCacheClusters,

            int port,

            string primaryEndpointAddress,

            string replicationGroupDescription,

            string replicationGroupId,

            int snapshotRetentionLimit,

            string snapshotWindow)
        {
            AuthTokenEnabled = authTokenEnabled;
            AutomaticFailoverEnabled = automaticFailoverEnabled;
            ConfigurationEndpointAddress = configurationEndpointAddress;
            Id = id;
            MemberClusters = memberClusters;
            NodeType = nodeType;
            NumberCacheClusters = numberCacheClusters;
            Port = port;
            PrimaryEndpointAddress = primaryEndpointAddress;
            ReplicationGroupDescription = replicationGroupDescription;
            ReplicationGroupId = replicationGroupId;
            SnapshotRetentionLimit = snapshotRetentionLimit;
            SnapshotWindow = snapshotWindow;
        }
    }
}
