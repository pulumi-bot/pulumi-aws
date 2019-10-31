// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Dax
{
    /// <summary>
    /// Provides a DAX Cluster resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dax_cluster.html.markdown.
    /// </summary>
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the DAX cluster
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// List of Availability Zones in which the
        /// nodes will be created
        /// </summary>
        [Output("availabilityZones")]
        public Output<ImmutableArray<string>> AvailabilityZones { get; private set; } = null!;

        /// <summary>
        /// The DNS name of the DAX cluster without the port appended
        /// </summary>
        [Output("clusterAddress")]
        public Output<string> ClusterAddress { get; private set; } = null!;

        /// <summary>
        /// Group identifier. DAX converts this name to
        /// lowercase
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// The configuration endpoint for this DAX cluster,
        /// consisting of a DNS name and a port number
        /// </summary>
        [Output("configurationEndpoint")]
        public Output<string> ConfigurationEndpoint { get; private set; } = null!;

        /// <summary>
        /// Description for the cluster
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A valid Amazon Resource Name (ARN) that identifies
        /// an IAM role. At runtime, DAX will assume this role and use the role's
        /// permissions to access DynamoDB on your behalf
        /// </summary>
        [Output("iamRoleArn")]
        public Output<string> IamRoleArn { get; private set; } = null!;

        /// <summary>
        /// Specifies the weekly time range for when
        /// maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
        /// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
        /// `sun:05:00-sun:09:00`
        /// </summary>
        [Output("maintenanceWindow")]
        public Output<string> MaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// The compute and memory capacity of the nodes. See
        /// [Nodes][1] for supported node types
        /// </summary>
        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        /// <summary>
        /// List of node objects including `id`, `address`, `port` and
        /// `availability_zone`. Referenceable e.g. as
        /// `${aws_dax_cluster.test.nodes.0.address}`
        /// </summary>
        [Output("nodes")]
        public Output<ImmutableArray<Outputs.ClusterNodes>> Nodes { get; private set; } = null!;

        /// <summary>
        /// An Amazon Resource Name (ARN) of an
        /// SNS topic to send DAX notifications to. Example:
        /// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
        /// </summary>
        [Output("notificationTopicArn")]
        public Output<string?> NotificationTopicArn { get; private set; } = null!;

        /// <summary>
        /// Name of the parameter group to associate
        /// with this DAX cluster
        /// </summary>
        [Output("parameterGroupName")]
        public Output<string> ParameterGroupName { get; private set; } = null!;

        /// <summary>
        /// The port used by the configuration endpoint
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The number of nodes in the DAX cluster. A
        /// replication factor of 1 will create a single-node cluster, without any read
        /// replicas
        /// </summary>
        [Output("replicationFactor")]
        public Output<int> ReplicationFactor { get; private set; } = null!;

        /// <summary>
        /// One or more VPC security groups associated
        /// with the cluster
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Encrypt at rest options
        /// </summary>
        [Output("serverSideEncryption")]
        public Output<Outputs.ClusterServerSideEncryption?> ServerSideEncryption { get; private set; } = null!;

        /// <summary>
        /// Name of the subnet group to be used for the
        /// cluster
        /// </summary>
        [Output("subnetGroupName")]
        public Output<string> SubnetGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource
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
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("aws:dax/cluster:Cluster", name, args, MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("aws:dax/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;

        /// <summary>
        /// List of Availability Zones in which the
        /// nodes will be created
        /// </summary>
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        /// <summary>
        /// Group identifier. DAX converts this name to
        /// lowercase
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Description for the cluster
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A valid Amazon Resource Name (ARN) that identifies
        /// an IAM role. At runtime, DAX will assume this role and use the role's
        /// permissions to access DynamoDB on your behalf
        /// </summary>
        [Input("iamRoleArn", required: true)]
        public Input<string> IamRoleArn { get; set; } = null!;

        /// <summary>
        /// Specifies the weekly time range for when
        /// maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
        /// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
        /// `sun:05:00-sun:09:00`
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<string>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The compute and memory capacity of the nodes. See
        /// [Nodes][1] for supported node types
        /// </summary>
        [Input("nodeType", required: true)]
        public Input<string> NodeType { get; set; } = null!;

        /// <summary>
        /// An Amazon Resource Name (ARN) of an
        /// SNS topic to send DAX notifications to. Example:
        /// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
        /// </summary>
        [Input("notificationTopicArn")]
        public Input<string>? NotificationTopicArn { get; set; }

        /// <summary>
        /// Name of the parameter group to associate
        /// with this DAX cluster
        /// </summary>
        [Input("parameterGroupName")]
        public Input<string>? ParameterGroupName { get; set; }

        /// <summary>
        /// The number of nodes in the DAX cluster. A
        /// replication factor of 1 will create a single-node cluster, without any read
        /// replicas
        /// </summary>
        [Input("replicationFactor", required: true)]
        public Input<int> ReplicationFactor { get; set; } = null!;

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// One or more VPC security groups associated
        /// with the cluster
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Encrypt at rest options
        /// </summary>
        [Input("serverSideEncryption")]
        public Input<Inputs.ClusterServerSideEncryptionArgs>? ServerSideEncryption { get; set; }

        /// <summary>
        /// Name of the subnet group to be used for the
        /// cluster
        /// </summary>
        [Input("subnetGroupName")]
        public Input<string>? SubnetGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource
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
        /// The ARN of the DAX cluster
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;

        /// <summary>
        /// List of Availability Zones in which the
        /// nodes will be created
        /// </summary>
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        /// <summary>
        /// The DNS name of the DAX cluster without the port appended
        /// </summary>
        [Input("clusterAddress")]
        public Input<string>? ClusterAddress { get; set; }

        /// <summary>
        /// Group identifier. DAX converts this name to
        /// lowercase
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// The configuration endpoint for this DAX cluster,
        /// consisting of a DNS name and a port number
        /// </summary>
        [Input("configurationEndpoint")]
        public Input<string>? ConfigurationEndpoint { get; set; }

        /// <summary>
        /// Description for the cluster
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A valid Amazon Resource Name (ARN) that identifies
        /// an IAM role. At runtime, DAX will assume this role and use the role's
        /// permissions to access DynamoDB on your behalf
        /// </summary>
        [Input("iamRoleArn")]
        public Input<string>? IamRoleArn { get; set; }

        /// <summary>
        /// Specifies the weekly time range for when
        /// maintenance on the cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi`
        /// (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example:
        /// `sun:05:00-sun:09:00`
        /// </summary>
        [Input("maintenanceWindow")]
        public Input<string>? MaintenanceWindow { get; set; }

        /// <summary>
        /// The compute and memory capacity of the nodes. See
        /// [Nodes][1] for supported node types
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        [Input("nodes")]
        private InputList<Inputs.ClusterNodesGetArgs>? _nodes;

        /// <summary>
        /// List of node objects including `id`, `address`, `port` and
        /// `availability_zone`. Referenceable e.g. as
        /// `${aws_dax_cluster.test.nodes.0.address}`
        /// </summary>
        public InputList<Inputs.ClusterNodesGetArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<Inputs.ClusterNodesGetArgs>());
            set => _nodes = value;
        }

        /// <summary>
        /// An Amazon Resource Name (ARN) of an
        /// SNS topic to send DAX notifications to. Example:
        /// `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
        /// </summary>
        [Input("notificationTopicArn")]
        public Input<string>? NotificationTopicArn { get; set; }

        /// <summary>
        /// Name of the parameter group to associate
        /// with this DAX cluster
        /// </summary>
        [Input("parameterGroupName")]
        public Input<string>? ParameterGroupName { get; set; }

        /// <summary>
        /// The port used by the configuration endpoint
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The number of nodes in the DAX cluster. A
        /// replication factor of 1 will create a single-node cluster, without any read
        /// replicas
        /// </summary>
        [Input("replicationFactor")]
        public Input<int>? ReplicationFactor { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// One or more VPC security groups associated
        /// with the cluster
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Encrypt at rest options
        /// </summary>
        [Input("serverSideEncryption")]
        public Input<Inputs.ClusterServerSideEncryptionGetArgs>? ServerSideEncryption { get; set; }

        /// <summary>
        /// Name of the subnet group to be used for the
        /// cluster
        /// </summary>
        [Input("subnetGroupName")]
        public Input<string>? SubnetGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource
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

    namespace Inputs
    {

    public sealed class ClusterNodesGetArgs : Pulumi.ResourceArgs
    {
        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The port used by the configuration endpoint
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        public ClusterNodesGetArgs()
        {
        }
    }

    public sealed class ClusterServerSideEncryptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable encryption at rest. Defaults to `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ClusterServerSideEncryptionArgs()
        {
        }
    }

    public sealed class ClusterServerSideEncryptionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable encryption at rest. Defaults to `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ClusterServerSideEncryptionGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ClusterNodes
    {
        public readonly string Address;
        public readonly string AvailabilityZone;
        public readonly string Id;
        /// <summary>
        /// The port used by the configuration endpoint
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private ClusterNodes(
            string address,
            string availabilityZone,
            string id,
            int port)
        {
            Address = address;
            AvailabilityZone = availabilityZone;
            Id = id;
            Port = port;
        }
    }

    [OutputType]
    public sealed class ClusterServerSideEncryption
    {
        /// <summary>
        /// Whether to enable encryption at rest. Defaults to `false`.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private ClusterServerSideEncryption(bool? enabled)
        {
            Enabled = enabled;
        }
    }
    }
}
