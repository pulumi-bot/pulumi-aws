// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.RedShift
{
    public static class GetCluster
    {
        /// <summary>
        /// Provides details about a specific redshift cluster.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var testCluster = Output.Create(Aws.RedShift.GetCluster.InvokeAsync(new Aws.RedShift.GetClusterArgs
        ///         {
        ///             ClusterIdentifier = "test-cluster",
        ///         }));
        ///         var testStream = new Aws.Kinesis.FirehoseDeliveryStream("testStream", new Aws.Kinesis.FirehoseDeliveryStreamArgs
        ///         {
        ///             Destination = "redshift",
        ///             S3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamS3ConfigurationArgs
        ///             {
        ///                 RoleArn = aws_iam_role.Firehose_role.Arn,
        ///                 BucketArn = aws_s3_bucket.Bucket.Arn,
        ///                 BufferSize = 10,
        ///                 BufferInterval = 400,
        ///                 CompressionFormat = "GZIP",
        ///             },
        ///             RedshiftConfiguration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamRedshiftConfigurationArgs
        ///             {
        ///                 RoleArn = aws_iam_role.Firehose_role.Arn,
        ///                 ClusterJdbcurl = Output.Tuple(testCluster, testCluster).Apply(values =&gt;
        ///                 {
        ///                     var testCluster = values.Item1;
        ///                     var testCluster1 = values.Item2;
        ///                     return $"jdbc:redshift://{testCluster.Endpoint}/{testCluster1.DatabaseName}";
        ///                 }),
        ///                 Username = "testuser",
        ///                 Password = "T3stPass",
        ///                 DataTableName = "test-table",
        ///                 CopyOptions = "delimiter '|'",
        ///                 DataTableColumns = "test-col",
        ///             },
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("aws:redshift/getCluster:getCluster", args ?? new GetClusterArgs(), options.WithVersion());

        public static Output<GetClusterResult> Invoke(GetClusterOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.ClusterIdentifier.Box(),
                args.Tags.ToDict().Box()
            ).Apply(a => {
                    var args = new GetClusterArgs();
                    a[0].Set(args, nameof(args.ClusterIdentifier));
                    a[1].Set(args, nameof(args.Tags));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cluster identifier
        /// </summary>
        [Input("clusterIdentifier", required: true)]
        public string ClusterIdentifier { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// The tags associated to the cluster
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetClusterArgs()
        {
        }
    }

    public sealed class GetClusterOutputArgs
    {
        /// <summary>
        /// The cluster identifier
        /// </summary>
        [Input("clusterIdentifier", required: true)]
        public Input<string> ClusterIdentifier { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags associated to the cluster
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetClusterOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// Whether major version upgrades can be applied during maintenance period
        /// </summary>
        public readonly bool AllowVersionUpgrade;
        /// <summary>
        /// The backup retention period
        /// </summary>
        public readonly int AutomatedSnapshotRetentionPeriod;
        /// <summary>
        /// The availability zone of the cluster
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// The name of the S3 bucket where the log files are to be stored
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// The cluster identifier
        /// </summary>
        public readonly string ClusterIdentifier;
        /// <summary>
        /// The name of the parameter group to be associated with this cluster
        /// </summary>
        public readonly string ClusterParameterGroupName;
        /// <summary>
        /// The public key for the cluster
        /// </summary>
        public readonly string ClusterPublicKey;
        /// <summary>
        /// The cluster revision number
        /// </summary>
        public readonly string ClusterRevisionNumber;
        /// <summary>
        /// The security groups associated with the cluster
        /// </summary>
        public readonly ImmutableArray<string> ClusterSecurityGroups;
        /// <summary>
        /// The name of a cluster subnet group to be associated with this cluster
        /// </summary>
        public readonly string ClusterSubnetGroupName;
        /// <summary>
        /// The cluster type
        /// </summary>
        public readonly string ClusterType;
        public readonly string ClusterVersion;
        /// <summary>
        /// The name of the default database in the cluster
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// The Elastic IP of the cluster
        /// </summary>
        public readonly string ElasticIp;
        /// <summary>
        /// Whether cluster logging is enabled
        /// </summary>
        public readonly bool EnableLogging;
        /// <summary>
        /// Whether the cluster data is encrypted
        /// </summary>
        public readonly bool Encrypted;
        /// <summary>
        /// The cluster endpoint
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// Whether enhanced VPC routing is enabled
        /// </summary>
        public readonly bool EnhancedVpcRouting;
        /// <summary>
        /// The IAM roles associated to the cluster
        /// </summary>
        public readonly ImmutableArray<string> IamRoles;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The KMS encryption key associated to the cluster
        /// </summary>
        public readonly string KmsKeyId;
        /// <summary>
        /// Username for the master DB user
        /// </summary>
        public readonly string MasterUsername;
        /// <summary>
        /// The cluster node type
        /// </summary>
        public readonly string NodeType;
        /// <summary>
        /// The number of nodes in the cluster
        /// </summary>
        public readonly int NumberOfNodes;
        /// <summary>
        /// The port the cluster responds on
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The maintenance window
        /// </summary>
        public readonly string PreferredMaintenanceWindow;
        /// <summary>
        /// Whether the cluster is publicly accessible
        /// </summary>
        public readonly bool PubliclyAccessible;
        /// <summary>
        /// The folder inside the S3 bucket where the log files are stored
        /// </summary>
        public readonly string S3KeyPrefix;
        /// <summary>
        /// The tags associated to the cluster
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The VPC Id associated with the cluster
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The VPC security group Ids associated with the cluster
        /// </summary>
        public readonly ImmutableArray<string> VpcSecurityGroupIds;

        [OutputConstructor]
        private GetClusterResult(
            bool allowVersionUpgrade,

            int automatedSnapshotRetentionPeriod,

            string availabilityZone,

            string bucketName,

            string clusterIdentifier,

            string clusterParameterGroupName,

            string clusterPublicKey,

            string clusterRevisionNumber,

            ImmutableArray<string> clusterSecurityGroups,

            string clusterSubnetGroupName,

            string clusterType,

            string clusterVersion,

            string databaseName,

            string elasticIp,

            bool enableLogging,

            bool encrypted,

            string endpoint,

            bool enhancedVpcRouting,

            ImmutableArray<string> iamRoles,

            string id,

            string kmsKeyId,

            string masterUsername,

            string nodeType,

            int numberOfNodes,

            int port,

            string preferredMaintenanceWindow,

            bool publiclyAccessible,

            string s3KeyPrefix,

            ImmutableDictionary<string, string>? tags,

            string vpcId,

            ImmutableArray<string> vpcSecurityGroupIds)
        {
            AllowVersionUpgrade = allowVersionUpgrade;
            AutomatedSnapshotRetentionPeriod = automatedSnapshotRetentionPeriod;
            AvailabilityZone = availabilityZone;
            BucketName = bucketName;
            ClusterIdentifier = clusterIdentifier;
            ClusterParameterGroupName = clusterParameterGroupName;
            ClusterPublicKey = clusterPublicKey;
            ClusterRevisionNumber = clusterRevisionNumber;
            ClusterSecurityGroups = clusterSecurityGroups;
            ClusterSubnetGroupName = clusterSubnetGroupName;
            ClusterType = clusterType;
            ClusterVersion = clusterVersion;
            DatabaseName = databaseName;
            ElasticIp = elasticIp;
            EnableLogging = enableLogging;
            Encrypted = encrypted;
            Endpoint = endpoint;
            EnhancedVpcRouting = enhancedVpcRouting;
            IamRoles = iamRoles;
            Id = id;
            KmsKeyId = kmsKeyId;
            MasterUsername = masterUsername;
            NodeType = nodeType;
            NumberOfNodes = numberOfNodes;
            Port = port;
            PreferredMaintenanceWindow = preferredMaintenanceWindow;
            PubliclyAccessible = publiclyAccessible;
            S3KeyPrefix = s3KeyPrefix;
            Tags = tags;
            VpcId = vpcId;
            VpcSecurityGroupIds = vpcSecurityGroupIds;
        }
    }
}
