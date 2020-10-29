// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds
{
    public static class GetCluster
    {
        /// <summary>
        /// Provides information about an RDS cluster.
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("aws:rds/getCluster:getCluster", args ?? new GetClusterArgs(), options.WithVersion());
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cluster identifier of the RDS cluster.
        /// </summary>
        [Input("clusterIdentifier", required: true)]
        public string ClusterIdentifier { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetClusterArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        public readonly string Arn;
        public readonly ImmutableArray<string> AvailabilityZones;
        public readonly int BacktrackWindow;
        public readonly int BackupRetentionPeriod;
        public readonly string ClusterIdentifier;
        public readonly ImmutableArray<string> ClusterMembers;
        public readonly string ClusterResourceId;
        public readonly string DatabaseName;
        public readonly string DbClusterParameterGroupName;
        public readonly string DbSubnetGroupName;
        public readonly ImmutableArray<string> EnabledCloudwatchLogsExports;
        public readonly string Endpoint;
        public readonly string Engine;
        public readonly string EngineVersion;
        public readonly string FinalSnapshotIdentifier;
        public readonly string HostedZoneId;
        public readonly bool IamDatabaseAuthenticationEnabled;
        public readonly ImmutableArray<string> IamRoles;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string KmsKeyId;
        public readonly string MasterUsername;
        public readonly int Port;
        public readonly string PreferredBackupWindow;
        public readonly string PreferredMaintenanceWindow;
        public readonly string ReaderEndpoint;
        public readonly string ReplicationSourceIdentifier;
        public readonly bool StorageEncrypted;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly ImmutableArray<string> VpcSecurityGroupIds;

        [OutputConstructor]
        private GetClusterResult(
            string arn,

            ImmutableArray<string> availabilityZones,

            int backtrackWindow,

            int backupRetentionPeriod,

            string clusterIdentifier,

            ImmutableArray<string> clusterMembers,

            string clusterResourceId,

            string databaseName,

            string dbClusterParameterGroupName,

            string dbSubnetGroupName,

            ImmutableArray<string> enabledCloudwatchLogsExports,

            string endpoint,

            string engine,

            string engineVersion,

            string finalSnapshotIdentifier,

            string hostedZoneId,

            bool iamDatabaseAuthenticationEnabled,

            ImmutableArray<string> iamRoles,

            string id,

            string kmsKeyId,

            string masterUsername,

            int port,

            string preferredBackupWindow,

            string preferredMaintenanceWindow,

            string readerEndpoint,

            string replicationSourceIdentifier,

            bool storageEncrypted,

            ImmutableDictionary<string, string> tags,

            ImmutableArray<string> vpcSecurityGroupIds)
        {
            Arn = arn;
            AvailabilityZones = availabilityZones;
            BacktrackWindow = backtrackWindow;
            BackupRetentionPeriod = backupRetentionPeriod;
            ClusterIdentifier = clusterIdentifier;
            ClusterMembers = clusterMembers;
            ClusterResourceId = clusterResourceId;
            DatabaseName = databaseName;
            DbClusterParameterGroupName = dbClusterParameterGroupName;
            DbSubnetGroupName = dbSubnetGroupName;
            EnabledCloudwatchLogsExports = enabledCloudwatchLogsExports;
            Endpoint = endpoint;
            Engine = engine;
            EngineVersion = engineVersion;
            FinalSnapshotIdentifier = finalSnapshotIdentifier;
            HostedZoneId = hostedZoneId;
            IamDatabaseAuthenticationEnabled = iamDatabaseAuthenticationEnabled;
            IamRoles = iamRoles;
            Id = id;
            KmsKeyId = kmsKeyId;
            MasterUsername = masterUsername;
            Port = port;
            PreferredBackupWindow = preferredBackupWindow;
            PreferredMaintenanceWindow = preferredMaintenanceWindow;
            ReaderEndpoint = readerEndpoint;
            ReplicationSourceIdentifier = replicationSourceIdentifier;
            StorageEncrypted = storageEncrypted;
            Tags = tags;
            VpcSecurityGroupIds = vpcSecurityGroupIds;
        }
    }
}
