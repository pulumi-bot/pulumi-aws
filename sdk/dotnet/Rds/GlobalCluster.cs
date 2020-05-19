// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds
{
    /// <summary>
    /// Manages a RDS Global Cluster, which is an Aurora global database spread across multiple regions. The global database contains a single primary cluster with read-write capability, and a read-only secondary cluster that receives data from the primary cluster through high-speed replication performed by the Aurora storage subsystem.
    /// 
    /// More information about Aurora global databases can be found in the [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database-creating).
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// using Pulumi = Pulumi.Pulumi;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var primary = new Aws.Provider("primary", new Aws.ProviderArgs
    ///         {
    ///             Region = "us-east-2",
    ///         });
    ///         var secondary = new Aws.Provider("secondary", new Aws.ProviderArgs
    ///         {
    ///             Region = "us-west-2",
    ///         });
    ///         var example = new Aws.Rds.GlobalCluster("example", new Aws.Rds.GlobalClusterArgs
    ///         {
    ///             GlobalClusterIdentifier = "example",
    ///         });
    ///         var primaryCluster = new Aws.Rds.Cluster("primaryCluster", new Aws.Rds.ClusterArgs
    ///         {
    ///             EngineMode = "global",
    ///             GlobalClusterIdentifier = example.Id,
    ///         });
    ///         var primaryClusterInstance = new Aws.Rds.ClusterInstance("primaryClusterInstance", new Aws.Rds.ClusterInstanceArgs
    ///         {
    ///             ClusterIdentifier = primaryCluster.Id,
    ///         });
    ///         var secondaryCluster = new Aws.Rds.Cluster("secondaryCluster", new Aws.Rds.ClusterArgs
    ///         {
    ///             EngineMode = "global",
    ///             GlobalClusterIdentifier = example.Id,
    ///         });
    ///         var secondaryClusterInstance = new Aws.Rds.ClusterInstance("secondaryClusterInstance", new Aws.Rds.ClusterInstanceArgs
    ///         {
    ///             ClusterIdentifier = secondaryCluster.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class GlobalCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// RDS Global Cluster Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Name for an automatically created database on cluster creation.
        /// </summary>
        [Output("databaseName")]
        public Output<string?> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool?> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// Name of the database engine to be used for this DB cluster. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`.
        /// </summary>
        [Output("engine")]
        public Output<string?> Engine { get; private set; } = null!;

        /// <summary>
        /// Engine version of the Aurora global database.
        /// * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// The global cluster identifier.
        /// </summary>
        [Output("globalClusterIdentifier")]
        public Output<string> GlobalClusterIdentifier { get; private set; } = null!;

        /// <summary>
        /// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
        /// </summary>
        [Output("globalClusterResourceId")]
        public Output<string> GlobalClusterResourceId { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the DB cluster is encrypted. The default is `false`.
        /// </summary>
        [Output("storageEncrypted")]
        public Output<bool?> StorageEncrypted { get; private set; } = null!;


        /// <summary>
        /// Create a GlobalCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GlobalCluster(string name, GlobalClusterArgs args, CustomResourceOptions? options = null)
            : base("aws:rds/globalCluster:GlobalCluster", name, args ?? new GlobalClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GlobalCluster(string name, Input<string> id, GlobalClusterState? state = null, CustomResourceOptions? options = null)
            : base("aws:rds/globalCluster:GlobalCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GlobalCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GlobalCluster Get(string name, Input<string> id, GlobalClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new GlobalCluster(name, id, state, options);
        }
    }

    public sealed class GlobalClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name for an automatically created database on cluster creation.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Name of the database engine to be used for this DB cluster. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Engine version of the Aurora global database.
        /// * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// The global cluster identifier.
        /// </summary>
        [Input("globalClusterIdentifier", required: true)]
        public Input<string> GlobalClusterIdentifier { get; set; } = null!;

        /// <summary>
        /// Specifies whether the DB cluster is encrypted. The default is `false`.
        /// </summary>
        [Input("storageEncrypted")]
        public Input<bool>? StorageEncrypted { get; set; }

        public GlobalClusterArgs()
        {
        }
    }

    public sealed class GlobalClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// RDS Global Cluster Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Name for an automatically created database on cluster creation.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// Name of the database engine to be used for this DB cluster. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Engine version of the Aurora global database.
        /// * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// The global cluster identifier.
        /// </summary>
        [Input("globalClusterIdentifier")]
        public Input<string>? GlobalClusterIdentifier { get; set; }

        /// <summary>
        /// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
        /// </summary>
        [Input("globalClusterResourceId")]
        public Input<string>? GlobalClusterResourceId { get; set; }

        /// <summary>
        /// Specifies whether the DB cluster is encrypted. The default is `false`.
        /// </summary>
        [Input("storageEncrypted")]
        public Input<bool>? StorageEncrypted { get; set; }

        public GlobalClusterState()
        {
        }
    }
}
