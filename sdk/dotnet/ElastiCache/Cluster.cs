// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElastiCache
{
    /// <summary>
    /// Provides an ElastiCache Cluster resource, which manages a Memcached cluster or Redis instance.
    /// For working with Redis (Cluster Mode Enabled) replication groups, see the
    /// `aws.elasticache.ReplicationGroup` resource.
    /// 
    /// &gt; **Note:** When you change an attribute, such as `node_type`, by default
    /// it is applied in the next maintenance window. Because of this, this provider may report
    /// a difference in its planning phase because the actual modification has not yet taken
    /// place. You can use the `apply_immediately` flag to instruct the service to apply the
    /// change immediately. Using `apply_immediately` can result in a brief downtime as the server reboots.
    /// See the AWS Docs on [Modifying an ElastiCache Cache Cluster](https://docs.aws.amazon.com/AmazonElastiCache/latest/UserGuide/Clusters.Modify.html) for more information.
    /// 
    /// ## Example Usage
    /// ### Memcached Cluster
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.ElastiCache.Cluster("example", new Aws.ElastiCache.ClusterArgs
    ///         {
    ///             Engine = "memcached",
    ///             NodeType = "cache.m4.large",
    ///             NumCacheNodes = 2,
    ///             ParameterGroupName = "default.memcached1.4",
    ///             Port = 11211,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Redis Instance
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.ElastiCache.Cluster("example", new Aws.ElastiCache.ClusterArgs
    ///         {
    ///             Engine = "redis",
    ///             EngineVersion = "3.2.10",
    ///             NodeType = "cache.m4.large",
    ///             NumCacheNodes = 1,
    ///             ParameterGroupName = "default.redis3.2",
    ///             Port = 6379,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Redis Cluster Mode Disabled Read Replica Instance
    /// 
    /// These inherit their settings from the replication group.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var replica = new Aws.ElastiCache.Cluster("replica", new Aws.ElastiCache.ClusterArgs
    ///         {
    ///             ReplicationGroupId = aws_elasticache_replication_group.Example.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ElastiCache Clusters can be imported using the `cluster_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:elasticache/cluster:Cluster my_cluster my_cluster
    /// ```
    /// </summary>
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether any database modifications
        /// are applied immediately, or during the next maintenance window. Default is
        /// `false`. See [Amazon ElastiCache Documentation for more information.](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheCluster.html)
        /// (Available since v0.6.0)
        /// </summary>
        [Output("applyImmediately")]
        public Output<bool> ApplyImmediately { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferred_availability_zones` instead. Default: System chosen Availability Zone.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `num_cache_nodes` must be greater than `1`
        /// </summary>
        [Output("azMode")]
        public Output<string> AzMode { get; private set; } = null!;

        /// <summary>
        /// List of node objects including `id`, `address`, `port` and `availability_zone`.
        /// Referenceable e.g. as `${aws_elasticache_cluster.bar.cache_nodes.0.address}`
        /// </summary>
        [Output("cacheNodes")]
        public Output<ImmutableArray<Outputs.ClusterCacheNode>> CacheNodes { get; private set; } = null!;

        /// <summary>
        /// (Memcached only) The DNS name of the cache cluster without the port appended.
        /// </summary>
        [Output("clusterAddress")]
        public Output<string> ClusterAddress { get; private set; } = null!;

        /// <summary>
        /// Group identifier. ElastiCache converts
        /// this name to lowercase
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// (Memcached only) The configuration endpoint to allow host discovery.
        /// </summary>
        [Output("configurationEndpoint")]
        public Output<string> ConfigurationEndpoint { get; private set; } = null!;

        /// <summary>
        /// Name of the cache engine to be used for this cache cluster.
        /// Valid values for this parameter are `memcached` or `redis`
        /// </summary>
        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        /// <summary>
        /// Version number of the cache engine to be used.
        /// See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html)
        /// in the AWS Documentation center for supported versions
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// Specifies the weekly time range for when maintenance
        /// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
        /// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
        /// </summary>
        [Output("maintenanceWindow")]
        public Output<string> MaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// The compute and memory capacity of the nodes. See
        /// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
        /// supported node types
        /// </summary>
        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        /// <summary>
        /// An Amazon Resource Name (ARN) of an
        /// SNS topic to send ElastiCache notifications to. Example:
        /// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
        /// </summary>
        [Output("notificationTopicArn")]
        public Output<string?> NotificationTopicArn { get; private set; } = null!;

        /// <summary>
        /// The initial number of cache nodes that the
        /// cache cluster will have. For Redis, this value must be 1. For Memcache, this
        /// value must be between 1 and 20. If this number is reduced on subsequent runs,
        /// the highest numbered nodes will be removed.
        /// </summary>
        [Output("numCacheNodes")]
        public Output<int> NumCacheNodes { get; private set; } = null!;

        /// <summary>
        /// Name of the parameter group to associate
        /// with this cache cluster
        /// </summary>
        [Output("parameterGroupName")]
        public Output<string> ParameterGroupName { get; private set; } = null!;

        /// <summary>
        /// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replication_group_id`.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// A list of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `num_cache_nodes`. If you want all the nodes in the same Availability Zone, use `availability_zone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
        /// </summary>
        [Output("preferredAvailabilityZones")]
        public Output<ImmutableArray<string>> PreferredAvailabilityZones { get; private set; } = null!;

        /// <summary>
        /// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
        /// </summary>
        [Output("replicationGroupId")]
        public Output<string> ReplicationGroupId { get; private set; } = null!;

        /// <summary>
        /// One or more VPC security groups associated
        /// with the cache cluster
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// List of security group
        /// names to associate with this cache cluster
        /// </summary>
        [Output("securityGroupNames")]
        public Output<ImmutableArray<string>> SecurityGroupNames { get; private set; } = null!;

        /// <summary>
        /// A single-element string list containing an
        /// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
        /// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
        /// </summary>
        [Output("snapshotArns")]
        public Output<ImmutableArray<string>> SnapshotArns { get; private set; } = null!;

        /// <summary>
        /// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshot_name` forces a new resource.
        /// </summary>
        [Output("snapshotName")]
        public Output<string?> SnapshotName { get; private set; } = null!;

        /// <summary>
        /// The number of days for which ElastiCache will
        /// retain automatic cache cluster snapshots before deleting them. For example, if you set
        /// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
        /// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
        /// Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro or cache.t2.* cache nodes
        /// </summary>
        [Output("snapshotRetentionLimit")]
        public Output<int?> SnapshotRetentionLimit { get; private set; } = null!;

        /// <summary>
        /// The daily time range (in UTC) during which ElastiCache will
        /// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
        /// </summary>
        [Output("snapshotWindow")]
        public Output<string> SnapshotWindow { get; private set; } = null!;

        /// <summary>
        /// Name of the subnet group to be used
        /// for the cache cluster.
        /// </summary>
        [Output("subnetGroupName")]
        public Output<string> SubnetGroupName { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:elasticache/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("aws:elasticache/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        /// Specifies whether any database modifications
        /// are applied immediately, or during the next maintenance window. Default is
        /// `false`. See [Amazon ElastiCache Documentation for more information.](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheCluster.html)
        /// (Available since v0.6.0)
        /// </summary>
        [Input("applyImmediately")]
        public Input<bool>? ApplyImmediately { get; set; }

        /// <summary>
        /// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferred_availability_zones` instead. Default: System chosen Availability Zone.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `num_cache_nodes` must be greater than `1`
        /// </summary>
        [Input("azMode")]
        public Input<string>? AzMode { get; set; }

        /// <summary>
        /// Group identifier. ElastiCache converts
        /// this name to lowercase
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Name of the cache engine to be used for this cache cluster.
        /// Valid values for this parameter are `memcached` or `redis`
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Version number of the cache engine to be used.
        /// See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html)
        /// in the AWS Documentation center for supported versions
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// Specifies the weekly time range for when maintenance
        /// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
        /// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<string>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The compute and memory capacity of the nodes. See
        /// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
        /// supported node types
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        /// <summary>
        /// An Amazon Resource Name (ARN) of an
        /// SNS topic to send ElastiCache notifications to. Example:
        /// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
        /// </summary>
        [Input("notificationTopicArn")]
        public Input<string>? NotificationTopicArn { get; set; }

        /// <summary>
        /// The initial number of cache nodes that the
        /// cache cluster will have. For Redis, this value must be 1. For Memcache, this
        /// value must be between 1 and 20. If this number is reduced on subsequent runs,
        /// the highest numbered nodes will be removed.
        /// </summary>
        [Input("numCacheNodes")]
        public Input<int>? NumCacheNodes { get; set; }

        /// <summary>
        /// Name of the parameter group to associate
        /// with this cache cluster
        /// </summary>
        [Input("parameterGroupName")]
        public Input<string>? ParameterGroupName { get; set; }

        /// <summary>
        /// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replication_group_id`.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("preferredAvailabilityZones")]
        private InputList<string>? _preferredAvailabilityZones;

        /// <summary>
        /// A list of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `num_cache_nodes`. If you want all the nodes in the same Availability Zone, use `availability_zone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
        /// </summary>
        public InputList<string> PreferredAvailabilityZones
        {
            get => _preferredAvailabilityZones ?? (_preferredAvailabilityZones = new InputList<string>());
            set => _preferredAvailabilityZones = value;
        }

        /// <summary>
        /// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
        /// </summary>
        [Input("replicationGroupId")]
        public Input<string>? ReplicationGroupId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// One or more VPC security groups associated
        /// with the cache cluster
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("securityGroupNames")]
        private InputList<string>? _securityGroupNames;

        /// <summary>
        /// List of security group
        /// names to associate with this cache cluster
        /// </summary>
        public InputList<string> SecurityGroupNames
        {
            get => _securityGroupNames ?? (_securityGroupNames = new InputList<string>());
            set => _securityGroupNames = value;
        }

        [Input("snapshotArns")]
        private InputList<string>? _snapshotArns;

        /// <summary>
        /// A single-element string list containing an
        /// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
        /// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
        /// </summary>
        public InputList<string> SnapshotArns
        {
            get => _snapshotArns ?? (_snapshotArns = new InputList<string>());
            set => _snapshotArns = value;
        }

        /// <summary>
        /// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshot_name` forces a new resource.
        /// </summary>
        [Input("snapshotName")]
        public Input<string>? SnapshotName { get; set; }

        /// <summary>
        /// The number of days for which ElastiCache will
        /// retain automatic cache cluster snapshots before deleting them. For example, if you set
        /// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
        /// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
        /// Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro or cache.t2.* cache nodes
        /// </summary>
        [Input("snapshotRetentionLimit")]
        public Input<int>? SnapshotRetentionLimit { get; set; }

        /// <summary>
        /// The daily time range (in UTC) during which ElastiCache will
        /// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
        /// </summary>
        [Input("snapshotWindow")]
        public Input<string>? SnapshotWindow { get; set; }

        /// <summary>
        /// Name of the subnet group to be used
        /// for the cache cluster.
        /// </summary>
        [Input("subnetGroupName")]
        public Input<string>? SubnetGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ClusterArgs()
        {
        }
    }

    public sealed class ClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether any database modifications
        /// are applied immediately, or during the next maintenance window. Default is
        /// `false`. See [Amazon ElastiCache Documentation for more information.](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheCluster.html)
        /// (Available since v0.6.0)
        /// </summary>
        [Input("applyImmediately")]
        public Input<bool>? ApplyImmediately { get; set; }

        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferred_availability_zones` instead. Default: System chosen Availability Zone.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `num_cache_nodes` must be greater than `1`
        /// </summary>
        [Input("azMode")]
        public Input<string>? AzMode { get; set; }

        [Input("cacheNodes")]
        private InputList<Inputs.ClusterCacheNodeGetArgs>? _cacheNodes;

        /// <summary>
        /// List of node objects including `id`, `address`, `port` and `availability_zone`.
        /// Referenceable e.g. as `${aws_elasticache_cluster.bar.cache_nodes.0.address}`
        /// </summary>
        public InputList<Inputs.ClusterCacheNodeGetArgs> CacheNodes
        {
            get => _cacheNodes ?? (_cacheNodes = new InputList<Inputs.ClusterCacheNodeGetArgs>());
            set => _cacheNodes = value;
        }

        /// <summary>
        /// (Memcached only) The DNS name of the cache cluster without the port appended.
        /// </summary>
        [Input("clusterAddress")]
        public Input<string>? ClusterAddress { get; set; }

        /// <summary>
        /// Group identifier. ElastiCache converts
        /// this name to lowercase
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// (Memcached only) The configuration endpoint to allow host discovery.
        /// </summary>
        [Input("configurationEndpoint")]
        public Input<string>? ConfigurationEndpoint { get; set; }

        /// <summary>
        /// Name of the cache engine to be used for this cache cluster.
        /// Valid values for this parameter are `memcached` or `redis`
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Version number of the cache engine to be used.
        /// See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html)
        /// in the AWS Documentation center for supported versions
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// Specifies the weekly time range for when maintenance
        /// on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
        /// The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<string>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The compute and memory capacity of the nodes. See
        /// [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
        /// supported node types
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        /// <summary>
        /// An Amazon Resource Name (ARN) of an
        /// SNS topic to send ElastiCache notifications to. Example:
        /// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
        /// </summary>
        [Input("notificationTopicArn")]
        public Input<string>? NotificationTopicArn { get; set; }

        /// <summary>
        /// The initial number of cache nodes that the
        /// cache cluster will have. For Redis, this value must be 1. For Memcache, this
        /// value must be between 1 and 20. If this number is reduced on subsequent runs,
        /// the highest numbered nodes will be removed.
        /// </summary>
        [Input("numCacheNodes")]
        public Input<int>? NumCacheNodes { get; set; }

        /// <summary>
        /// Name of the parameter group to associate
        /// with this cache cluster
        /// </summary>
        [Input("parameterGroupName")]
        public Input<string>? ParameterGroupName { get; set; }

        /// <summary>
        /// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replication_group_id`.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("preferredAvailabilityZones")]
        private InputList<string>? _preferredAvailabilityZones;

        /// <summary>
        /// A list of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `num_cache_nodes`. If you want all the nodes in the same Availability Zone, use `availability_zone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
        /// </summary>
        public InputList<string> PreferredAvailabilityZones
        {
            get => _preferredAvailabilityZones ?? (_preferredAvailabilityZones = new InputList<string>());
            set => _preferredAvailabilityZones = value;
        }

        /// <summary>
        /// The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
        /// </summary>
        [Input("replicationGroupId")]
        public Input<string>? ReplicationGroupId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// One or more VPC security groups associated
        /// with the cache cluster
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("securityGroupNames")]
        private InputList<string>? _securityGroupNames;

        /// <summary>
        /// List of security group
        /// names to associate with this cache cluster
        /// </summary>
        public InputList<string> SecurityGroupNames
        {
            get => _securityGroupNames ?? (_securityGroupNames = new InputList<string>());
            set => _securityGroupNames = value;
        }

        [Input("snapshotArns")]
        private InputList<string>? _snapshotArns;

        /// <summary>
        /// A single-element string list containing an
        /// Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
        /// Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
        /// </summary>
        public InputList<string> SnapshotArns
        {
            get => _snapshotArns ?? (_snapshotArns = new InputList<string>());
            set => _snapshotArns = value;
        }

        /// <summary>
        /// The name of a snapshot from which to restore data into the new node group.  Changing the `snapshot_name` forces a new resource.
        /// </summary>
        [Input("snapshotName")]
        public Input<string>? SnapshotName { get; set; }

        /// <summary>
        /// The number of days for which ElastiCache will
        /// retain automatic cache cluster snapshots before deleting them. For example, if you set
        /// SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
        /// before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
        /// Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro or cache.t2.* cache nodes
        /// </summary>
        [Input("snapshotRetentionLimit")]
        public Input<int>? SnapshotRetentionLimit { get; set; }

        /// <summary>
        /// The daily time range (in UTC) during which ElastiCache will
        /// begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
        /// </summary>
        [Input("snapshotWindow")]
        public Input<string>? SnapshotWindow { get; set; }

        /// <summary>
        /// Name of the subnet group to be used
        /// for the cache cluster.
        /// </summary>
        [Input("subnetGroupName")]
        public Input<string>? SubnetGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ClusterState()
        {
        }
    }
}
