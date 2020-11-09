// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Neptune
{
    /// <summary>
    /// Provides an Neptune Cluster Resource. A Cluster Resource defines attributes that are
    /// applied to the entire cluster of Neptune Cluster Instances.
    /// 
    /// Changes to a Neptune Cluster can occur when you manually change a
    /// parameter, such as `backup_retention_period`, and are reflected in the next maintenance
    /// window. Because of this, this provider may report a difference in its planning
    /// phase because a modification has not yet taken place. You can use the
    /// `apply_immediately` flag to instruct the service to apply the change immediately
    /// (see documentation below).
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
    ///         var @default = new Aws.Neptune.Cluster("default", new Aws.Neptune.ClusterArgs
    ///         {
    ///             ApplyImmediately = true,
    ///             BackupRetentionPeriod = 5,
    ///             ClusterIdentifier = "neptune-cluster-demo",
    ///             Engine = "neptune",
    ///             IamDatabaseAuthenticationEnabled = true,
    ///             PreferredBackupWindow = "07:00-09:00",
    ///             SkipFinalSnapshot = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// &gt; **Note:** AWS Neptune does not support user name/password–based access control.
    /// See the AWS [Docs](https://docs.aws.amazon.com/neptune/latest/userguide/limits.html) for more information.
    /// 
    /// ## Import
    /// 
    /// `aws_neptune_cluster` can be imported by using the cluster identifier, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:neptune/cluster:Cluster example my-cluster
    /// ```
    /// </summary>
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is `false`.
        /// </summary>
        [Output("applyImmediately")]
        public Output<bool> ApplyImmediately { get; private set; } = null!;

        /// <summary>
        /// The Neptune Cluster Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A list of EC2 Availability Zones that instances in the Neptune cluster can be created in.
        /// </summary>
        [Output("availabilityZones")]
        public Output<ImmutableArray<string>> AvailabilityZones { get; private set; } = null!;

        /// <summary>
        /// The days to retain backups for. Default `1`
        /// </summary>
        [Output("backupRetentionPeriod")]
        public Output<int?> BackupRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// The cluster identifier. If omitted, this provider will assign a random, unique identifier.
        /// </summary>
        [Output("clusterIdentifier")]
        public Output<string> ClusterIdentifier { get; private set; } = null!;

        /// <summary>
        /// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `cluster_identifier`.
        /// </summary>
        [Output("clusterIdentifierPrefix")]
        public Output<string> ClusterIdentifierPrefix { get; private set; } = null!;

        /// <summary>
        /// List of Neptune Instances that are a part of this cluster
        /// </summary>
        [Output("clusterMembers")]
        public Output<ImmutableArray<string>> ClusterMembers { get; private set; } = null!;

        /// <summary>
        /// The Neptune Cluster Resource ID
        /// </summary>
        [Output("clusterResourceId")]
        public Output<string> ClusterResourceId { get; private set; } = null!;

        /// <summary>
        /// A value that indicates whether the DB cluster has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool?> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// A list of the log types this DB cluster is configured to export to Cloudwatch Logs. Currently only supports `audit`.
        /// </summary>
        [Output("enableCloudwatchLogsExports")]
        public Output<ImmutableArray<string>> EnableCloudwatchLogsExports { get; private set; } = null!;

        /// <summary>
        /// The DNS address of the Neptune instance
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// The name of the database engine to be used for this Neptune cluster. Defaults to `neptune`.
        /// </summary>
        [Output("engine")]
        public Output<string?> Engine { get; private set; } = null!;

        /// <summary>
        /// The database engine version.
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// The name of your final Neptune snapshot when this Neptune cluster is deleted. If omitted, no final snapshot will be made.
        /// </summary>
        [Output("finalSnapshotIdentifier")]
        public Output<string?> FinalSnapshotIdentifier { get; private set; } = null!;

        /// <summary>
        /// The Route53 Hosted Zone ID of the endpoint
        /// </summary>
        [Output("hostedZoneId")]
        public Output<string> HostedZoneId { get; private set; } = null!;

        /// <summary>
        /// Specifies whether or mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.
        /// </summary>
        [Output("iamDatabaseAuthenticationEnabled")]
        public Output<bool?> IamDatabaseAuthenticationEnabled { get; private set; } = null!;

        /// <summary>
        /// A List of ARNs for the IAM roles to associate to the Neptune Cluster.
        /// </summary>
        [Output("iamRoles")]
        public Output<ImmutableArray<string>> IamRoles { get; private set; } = null!;

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying `kms_key_arn`, `storage_encrypted` needs to be set to true.
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string> KmsKeyArn { get; private set; } = null!;

        /// <summary>
        /// A cluster parameter group to associate with the cluster.
        /// </summary>
        [Output("neptuneClusterParameterGroupName")]
        public Output<string?> NeptuneClusterParameterGroupName { get; private set; } = null!;

        /// <summary>
        /// A Neptune subnet group to associate with this Neptune instance.
        /// </summary>
        [Output("neptuneSubnetGroupName")]
        public Output<string> NeptuneSubnetGroupName { get; private set; } = null!;

        /// <summary>
        /// The port on which the Neptune accepts connections. Default is `8182`.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. Time in UTC. Default: A 30-minute window selected at random from an 8-hour block of time per region. e.g. 04:00-09:00
        /// </summary>
        [Output("preferredBackupWindow")]
        public Output<string> PreferredBackupWindow { get; private set; } = null!;

        /// <summary>
        /// The weekly time range during which system maintenance can occur, in (UTC) e.g. wed:04:00-wed:04:30
        /// </summary>
        [Output("preferredMaintenanceWindow")]
        public Output<string> PreferredMaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// A read-only endpoint for the Neptune cluster, automatically load-balanced across replicas
        /// </summary>
        [Output("readerEndpoint")]
        public Output<string> ReaderEndpoint { get; private set; } = null!;

        /// <summary>
        /// ARN of a source Neptune cluster or Neptune instance if this Neptune cluster is to be created as a Read Replica.
        /// </summary>
        [Output("replicationSourceIdentifier")]
        public Output<string?> ReplicationSourceIdentifier { get; private set; } = null!;

        /// <summary>
        /// Determines whether a final Neptune snapshot is created before the Neptune cluster is deleted. If true is specified, no Neptune snapshot is created. If false is specified, a Neptune snapshot is created before the Neptune cluster is deleted, using the value from `final_snapshot_identifier`. Default is `false`.
        /// </summary>
        [Output("skipFinalSnapshot")]
        public Output<bool?> SkipFinalSnapshot { get; private set; } = null!;

        /// <summary>
        /// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a Neptune cluster snapshot, or the ARN when specifying a Neptune snapshot.
        /// </summary>
        [Output("snapshotIdentifier")]
        public Output<string?> SnapshotIdentifier { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the Neptune cluster is encrypted. The default is `false` if not specified.
        /// </summary>
        [Output("storageEncrypted")]
        public Output<bool?> StorageEncrypted { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the Neptune cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// List of VPC security groups to associate with the Cluster
        /// </summary>
        [Output("vpcSecurityGroupIds")]
        public Output<ImmutableArray<string>> VpcSecurityGroupIds { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:neptune/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("aws:neptune/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is `false`.
        /// </summary>
        [Input("applyImmediately")]
        public Input<bool>? ApplyImmediately { get; set; }

        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;

        /// <summary>
        /// A list of EC2 Availability Zones that instances in the Neptune cluster can be created in.
        /// </summary>
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        /// <summary>
        /// The days to retain backups for. Default `1`
        /// </summary>
        [Input("backupRetentionPeriod")]
        public Input<int>? BackupRetentionPeriod { get; set; }

        /// <summary>
        /// The cluster identifier. If omitted, this provider will assign a random, unique identifier.
        /// </summary>
        [Input("clusterIdentifier")]
        public Input<string>? ClusterIdentifier { get; set; }

        /// <summary>
        /// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `cluster_identifier`.
        /// </summary>
        [Input("clusterIdentifierPrefix")]
        public Input<string>? ClusterIdentifierPrefix { get; set; }

        /// <summary>
        /// A value that indicates whether the DB cluster has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        [Input("enableCloudwatchLogsExports")]
        private InputList<string>? _enableCloudwatchLogsExports;

        /// <summary>
        /// A list of the log types this DB cluster is configured to export to Cloudwatch Logs. Currently only supports `audit`.
        /// </summary>
        public InputList<string> EnableCloudwatchLogsExports
        {
            get => _enableCloudwatchLogsExports ?? (_enableCloudwatchLogsExports = new InputList<string>());
            set => _enableCloudwatchLogsExports = value;
        }

        /// <summary>
        /// The name of the database engine to be used for this Neptune cluster. Defaults to `neptune`.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// The database engine version.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// The name of your final Neptune snapshot when this Neptune cluster is deleted. If omitted, no final snapshot will be made.
        /// </summary>
        [Input("finalSnapshotIdentifier")]
        public Input<string>? FinalSnapshotIdentifier { get; set; }

        /// <summary>
        /// Specifies whether or mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.
        /// </summary>
        [Input("iamDatabaseAuthenticationEnabled")]
        public Input<bool>? IamDatabaseAuthenticationEnabled { get; set; }

        [Input("iamRoles")]
        private InputList<string>? _iamRoles;

        /// <summary>
        /// A List of ARNs for the IAM roles to associate to the Neptune Cluster.
        /// </summary>
        public InputList<string> IamRoles
        {
            get => _iamRoles ?? (_iamRoles = new InputList<string>());
            set => _iamRoles = value;
        }

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying `kms_key_arn`, `storage_encrypted` needs to be set to true.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// A cluster parameter group to associate with the cluster.
        /// </summary>
        [Input("neptuneClusterParameterGroupName")]
        public Input<string>? NeptuneClusterParameterGroupName { get; set; }

        /// <summary>
        /// A Neptune subnet group to associate with this Neptune instance.
        /// </summary>
        [Input("neptuneSubnetGroupName")]
        public Input<string>? NeptuneSubnetGroupName { get; set; }

        /// <summary>
        /// The port on which the Neptune accepts connections. Default is `8182`.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. Time in UTC. Default: A 30-minute window selected at random from an 8-hour block of time per region. e.g. 04:00-09:00
        /// </summary>
        [Input("preferredBackupWindow")]
        public Input<string>? PreferredBackupWindow { get; set; }

        /// <summary>
        /// The weekly time range during which system maintenance can occur, in (UTC) e.g. wed:04:00-wed:04:30
        /// </summary>
        [Input("preferredMaintenanceWindow")]
        public Input<string>? PreferredMaintenanceWindow { get; set; }

        /// <summary>
        /// ARN of a source Neptune cluster or Neptune instance if this Neptune cluster is to be created as a Read Replica.
        /// </summary>
        [Input("replicationSourceIdentifier")]
        public Input<string>? ReplicationSourceIdentifier { get; set; }

        /// <summary>
        /// Determines whether a final Neptune snapshot is created before the Neptune cluster is deleted. If true is specified, no Neptune snapshot is created. If false is specified, a Neptune snapshot is created before the Neptune cluster is deleted, using the value from `final_snapshot_identifier`. Default is `false`.
        /// </summary>
        [Input("skipFinalSnapshot")]
        public Input<bool>? SkipFinalSnapshot { get; set; }

        /// <summary>
        /// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a Neptune cluster snapshot, or the ARN when specifying a Neptune snapshot.
        /// </summary>
        [Input("snapshotIdentifier")]
        public Input<string>? SnapshotIdentifier { get; set; }

        /// <summary>
        /// Specifies whether the Neptune cluster is encrypted. The default is `false` if not specified.
        /// </summary>
        [Input("storageEncrypted")]
        public Input<bool>? StorageEncrypted { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the Neptune cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("vpcSecurityGroupIds")]
        private InputList<string>? _vpcSecurityGroupIds;

        /// <summary>
        /// List of VPC security groups to associate with the Cluster
        /// </summary>
        public InputList<string> VpcSecurityGroupIds
        {
            get => _vpcSecurityGroupIds ?? (_vpcSecurityGroupIds = new InputList<string>());
            set => _vpcSecurityGroupIds = value;
        }

        public ClusterArgs()
        {
        }
    }

    public sealed class ClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether any cluster modifications are applied immediately, or during the next maintenance window. Default is `false`.
        /// </summary>
        [Input("applyImmediately")]
        public Input<bool>? ApplyImmediately { get; set; }

        /// <summary>
        /// The Neptune Cluster Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;

        /// <summary>
        /// A list of EC2 Availability Zones that instances in the Neptune cluster can be created in.
        /// </summary>
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        /// <summary>
        /// The days to retain backups for. Default `1`
        /// </summary>
        [Input("backupRetentionPeriod")]
        public Input<int>? BackupRetentionPeriod { get; set; }

        /// <summary>
        /// The cluster identifier. If omitted, this provider will assign a random, unique identifier.
        /// </summary>
        [Input("clusterIdentifier")]
        public Input<string>? ClusterIdentifier { get; set; }

        /// <summary>
        /// Creates a unique cluster identifier beginning with the specified prefix. Conflicts with `cluster_identifier`.
        /// </summary>
        [Input("clusterIdentifierPrefix")]
        public Input<string>? ClusterIdentifierPrefix { get; set; }

        [Input("clusterMembers")]
        private InputList<string>? _clusterMembers;

        /// <summary>
        /// List of Neptune Instances that are a part of this cluster
        /// </summary>
        public InputList<string> ClusterMembers
        {
            get => _clusterMembers ?? (_clusterMembers = new InputList<string>());
            set => _clusterMembers = value;
        }

        /// <summary>
        /// The Neptune Cluster Resource ID
        /// </summary>
        [Input("clusterResourceId")]
        public Input<string>? ClusterResourceId { get; set; }

        /// <summary>
        /// A value that indicates whether the DB cluster has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        [Input("enableCloudwatchLogsExports")]
        private InputList<string>? _enableCloudwatchLogsExports;

        /// <summary>
        /// A list of the log types this DB cluster is configured to export to Cloudwatch Logs. Currently only supports `audit`.
        /// </summary>
        public InputList<string> EnableCloudwatchLogsExports
        {
            get => _enableCloudwatchLogsExports ?? (_enableCloudwatchLogsExports = new InputList<string>());
            set => _enableCloudwatchLogsExports = value;
        }

        /// <summary>
        /// The DNS address of the Neptune instance
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The name of the database engine to be used for this Neptune cluster. Defaults to `neptune`.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// The database engine version.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// The name of your final Neptune snapshot when this Neptune cluster is deleted. If omitted, no final snapshot will be made.
        /// </summary>
        [Input("finalSnapshotIdentifier")]
        public Input<string>? FinalSnapshotIdentifier { get; set; }

        /// <summary>
        /// The Route53 Hosted Zone ID of the endpoint
        /// </summary>
        [Input("hostedZoneId")]
        public Input<string>? HostedZoneId { get; set; }

        /// <summary>
        /// Specifies whether or mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled.
        /// </summary>
        [Input("iamDatabaseAuthenticationEnabled")]
        public Input<bool>? IamDatabaseAuthenticationEnabled { get; set; }

        [Input("iamRoles")]
        private InputList<string>? _iamRoles;

        /// <summary>
        /// A List of ARNs for the IAM roles to associate to the Neptune Cluster.
        /// </summary>
        public InputList<string> IamRoles
        {
            get => _iamRoles ?? (_iamRoles = new InputList<string>());
            set => _iamRoles = value;
        }

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying `kms_key_arn`, `storage_encrypted` needs to be set to true.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// A cluster parameter group to associate with the cluster.
        /// </summary>
        [Input("neptuneClusterParameterGroupName")]
        public Input<string>? NeptuneClusterParameterGroupName { get; set; }

        /// <summary>
        /// A Neptune subnet group to associate with this Neptune instance.
        /// </summary>
        [Input("neptuneSubnetGroupName")]
        public Input<string>? NeptuneSubnetGroupName { get; set; }

        /// <summary>
        /// The port on which the Neptune accepts connections. Default is `8182`.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. Time in UTC. Default: A 30-minute window selected at random from an 8-hour block of time per region. e.g. 04:00-09:00
        /// </summary>
        [Input("preferredBackupWindow")]
        public Input<string>? PreferredBackupWindow { get; set; }

        /// <summary>
        /// The weekly time range during which system maintenance can occur, in (UTC) e.g. wed:04:00-wed:04:30
        /// </summary>
        [Input("preferredMaintenanceWindow")]
        public Input<string>? PreferredMaintenanceWindow { get; set; }

        /// <summary>
        /// A read-only endpoint for the Neptune cluster, automatically load-balanced across replicas
        /// </summary>
        [Input("readerEndpoint")]
        public Input<string>? ReaderEndpoint { get; set; }

        /// <summary>
        /// ARN of a source Neptune cluster or Neptune instance if this Neptune cluster is to be created as a Read Replica.
        /// </summary>
        [Input("replicationSourceIdentifier")]
        public Input<string>? ReplicationSourceIdentifier { get; set; }

        /// <summary>
        /// Determines whether a final Neptune snapshot is created before the Neptune cluster is deleted. If true is specified, no Neptune snapshot is created. If false is specified, a Neptune snapshot is created before the Neptune cluster is deleted, using the value from `final_snapshot_identifier`. Default is `false`.
        /// </summary>
        [Input("skipFinalSnapshot")]
        public Input<bool>? SkipFinalSnapshot { get; set; }

        /// <summary>
        /// Specifies whether or not to create this cluster from a snapshot. You can use either the name or ARN when specifying a Neptune cluster snapshot, or the ARN when specifying a Neptune snapshot.
        /// </summary>
        [Input("snapshotIdentifier")]
        public Input<string>? SnapshotIdentifier { get; set; }

        /// <summary>
        /// Specifies whether the Neptune cluster is encrypted. The default is `false` if not specified.
        /// </summary>
        [Input("storageEncrypted")]
        public Input<bool>? StorageEncrypted { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the Neptune cluster.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("vpcSecurityGroupIds")]
        private InputList<string>? _vpcSecurityGroupIds;

        /// <summary>
        /// List of VPC security groups to associate with the Cluster
        /// </summary>
        public InputList<string> VpcSecurityGroupIds
        {
            get => _vpcSecurityGroupIds ?? (_vpcSecurityGroupIds = new InputList<string>());
            set => _vpcSecurityGroupIds = value;
        }

        public ClusterState()
        {
        }
    }
}
