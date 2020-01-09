// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElastiCache
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get information about an Elasticache Replication Group.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elasticache_replication_group.html.markdown.
        /// </summary>
        public static Task<GetReplicationGroupResult> GetReplicationGroup(GetReplicationGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReplicationGroupResult>("aws:elasticache/getReplicationGroup:getReplicationGroup", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetReplicationGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the replication group.
        /// </summary>
        [Input("replicationGroupId", required: true)]
        public string ReplicationGroupId { get; set; } = null!;

        public GetReplicationGroupArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetReplicationGroupResult
    {
        /// <summary>
        /// A flag that enables using an AuthToken (password) when issuing Redis commands.
        /// </summary>
        public readonly bool AuthTokenEnabled;
        /// <summary>
        /// A flag whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails.
        /// </summary>
        public readonly bool AutomaticFailoverEnabled;
        /// <summary>
        /// The configuration endpoint address to allow host discovery.
        /// </summary>
        public readonly string ConfigurationEndpointAddress;
        /// <summary>
        /// The identifiers of all the nodes that are part of this replication group.
        /// </summary>
        public readonly ImmutableArray<string> MemberClusters;
        /// <summary>
        /// The cluster node type.
        /// </summary>
        public readonly string NodeType;
        /// <summary>
        /// The number of cache clusters that the replication group has.
        /// </summary>
        public readonly int NumberCacheClusters;
        /// <summary>
        /// The port number on which the configuration endpoint will accept connections.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The endpoint of the primary node in this node group (shard).
        /// </summary>
        public readonly string PrimaryEndpointAddress;
        /// <summary>
        /// The description of the replication group.
        /// </summary>
        public readonly string ReplicationGroupDescription;
        /// <summary>
        /// The identifier for the replication group.
        /// </summary>
        public readonly string ReplicationGroupId;
        /// <summary>
        /// The number of days for which ElastiCache retains automatic cache cluster snapshots before deleting them.
        /// </summary>
        public readonly int SnapshotRetentionLimit;
        /// <summary>
        /// The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).
        /// </summary>
        public readonly string SnapshotWindow;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetReplicationGroupResult(
            bool authTokenEnabled,
            bool automaticFailoverEnabled,
            string configurationEndpointAddress,
            ImmutableArray<string> memberClusters,
            string nodeType,
            int numberCacheClusters,
            int port,
            string primaryEndpointAddress,
            string replicationGroupDescription,
            string replicationGroupId,
            int snapshotRetentionLimit,
            string snapshotWindow,
            string id)
        {
            AuthTokenEnabled = authTokenEnabled;
            AutomaticFailoverEnabled = automaticFailoverEnabled;
            ConfigurationEndpointAddress = configurationEndpointAddress;
            MemberClusters = memberClusters;
            NodeType = nodeType;
            NumberCacheClusters = numberCacheClusters;
            Port = port;
            PrimaryEndpointAddress = primaryEndpointAddress;
            ReplicationGroupDescription = replicationGroupDescription;
            ReplicationGroupId = replicationGroupId;
            SnapshotRetentionLimit = snapshotRetentionLimit;
            SnapshotWindow = snapshotWindow;
            Id = id;
        }
    }
}
