// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Eks
{
    /// <summary>
    /// Manages an EKS Node Group, which can provision and optionally update an Auto Scaling Group of Kubernetes worker nodes compatible with EKS. Additional documentation about this functionality can be found in the [EKS User Guide](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html).
    /// 
    /// ## Example Usage
    /// 
    /// ### Example IAM Role for EKS Node Group
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Iam.Role("example", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 { "Statement", new[]
    ///                     {
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             { "Action", "sts:AssumeRole" },
    ///                             { "Effect", "Allow" },
    ///                             { "Principal", new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 { "Service", "ec2.amazonaws.com" },
    ///                             } },
    ///                         },
    ///                     }
    ///                  },
    ///                 { "Version", "2012-10-17" },
    ///             }),
    ///         });
    ///         var example-AmazonEKSWorkerNodePolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSWorkerNodePolicy", new Aws.Iam.RolePolicyAttachmentArgs
    ///         {
    ///             PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
    ///             Role = example.Name,
    ///         });
    ///         var example-AmazonEKSCNIPolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSCNIPolicy", new Aws.Iam.RolePolicyAttachmentArgs
    ///         {
    ///             PolicyArn = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
    ///             Role = example.Name,
    ///         });
    ///         var example-AmazonEC2ContainerRegistryReadOnly = new Aws.Iam.RolePolicyAttachment("example-AmazonEC2ContainerRegistryReadOnly", new Aws.Iam.RolePolicyAttachmentArgs
    ///         {
    ///             PolicyArn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
    ///             Role = example.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class NodeGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. Defaults to `AL2_x86_64`. Valid values: `AL2_x86_64`, `AL2_x86_64_GPU`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Output("amiType")]
        public Output<string> AmiType { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the EKS Node Group.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Name of the EKS Cluster.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// Disk size in GiB for worker nodes. Defaults to `20`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Output("diskSize")]
        public Output<int> DiskSize { get; private set; } = null!;

        /// <summary>
        /// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
        /// </summary>
        [Output("forceUpdateVersion")]
        public Output<bool?> ForceUpdateVersion { get; private set; } = null!;

        /// <summary>
        /// Set of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. This provider will only perform drift detection if a configuration value is provided. Currently, the EKS API only accepts a single value in the set.
        /// </summary>
        [Output("instanceTypes")]
        public Output<string> InstanceTypes { get; private set; } = null!;

        /// <summary>
        /// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the EKS Node Group.
        /// </summary>
        [Output("nodeGroupName")]
        public Output<string> NodeGroupName { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
        /// </summary>
        [Output("nodeRoleArn")]
        public Output<string> NodeRoleArn { get; private set; } = null!;

        /// <summary>
        /// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
        /// </summary>
        [Output("releaseVersion")]
        public Output<string> ReleaseVersion { get; private set; } = null!;

        /// <summary>
        /// Configuration block with remote access settings. Detailed below.
        /// </summary>
        [Output("remoteAccess")]
        public Output<Outputs.NodeGroupRemoteAccess?> RemoteAccess { get; private set; } = null!;

        /// <summary>
        /// List of objects containing information about underlying resources.
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<Outputs.NodeGroupResource>> Resources { get; private set; } = null!;

        /// <summary>
        /// Configuration block with scaling settings. Detailed below.
        /// </summary>
        [Output("scalingConfig")]
        public Output<Outputs.NodeGroupScalingConfig> ScalingConfig { get; private set; } = null!;

        /// <summary>
        /// Status of the EKS Node Group.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Kubernetes version. Defaults to EKS Cluster Kubernetes version. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a NodeGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NodeGroup(string name, NodeGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:eks/nodeGroup:NodeGroup", name, args ?? new NodeGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NodeGroup(string name, Input<string> id, NodeGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:eks/nodeGroup:NodeGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NodeGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NodeGroup Get(string name, Input<string> id, NodeGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new NodeGroup(name, id, state, options);
        }
    }

    public sealed class NodeGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. Defaults to `AL2_x86_64`. Valid values: `AL2_x86_64`, `AL2_x86_64_GPU`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("amiType")]
        public Input<string>? AmiType { get; set; }

        /// <summary>
        /// Name of the EKS Cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Disk size in GiB for worker nodes. Defaults to `20`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
        /// </summary>
        [Input("forceUpdateVersion")]
        public Input<bool>? ForceUpdateVersion { get; set; }

        /// <summary>
        /// Set of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. This provider will only perform drift detection if a configuration value is provided. Currently, the EKS API only accepts a single value in the set.
        /// </summary>
        [Input("instanceTypes")]
        public Input<string>? InstanceTypes { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the EKS Node Group.
        /// </summary>
        [Input("nodeGroupName")]
        public Input<string>? NodeGroupName { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
        /// </summary>
        [Input("nodeRoleArn", required: true)]
        public Input<string> NodeRoleArn { get; set; } = null!;

        /// <summary>
        /// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
        /// </summary>
        [Input("releaseVersion")]
        public Input<string>? ReleaseVersion { get; set; }

        /// <summary>
        /// Configuration block with remote access settings. Detailed below.
        /// </summary>
        [Input("remoteAccess")]
        public Input<Inputs.NodeGroupRemoteAccessArgs>? RemoteAccess { get; set; }

        /// <summary>
        /// Configuration block with scaling settings. Detailed below.
        /// </summary>
        [Input("scalingConfig", required: true)]
        public Input<Inputs.NodeGroupScalingConfigArgs> ScalingConfig { get; set; } = null!;

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Kubernetes version. Defaults to EKS Cluster Kubernetes version. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public NodeGroupArgs()
        {
        }
    }

    public sealed class NodeGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of Amazon Machine Image (AMI) associated with the EKS Node Group. Defaults to `AL2_x86_64`. Valid values: `AL2_x86_64`, `AL2_x86_64_GPU`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("amiType")]
        public Input<string>? AmiType { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the EKS Node Group.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Name of the EKS Cluster.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// Disk size in GiB for worker nodes. Defaults to `20`. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        /// <summary>
        /// Force version update if existing pods are unable to be drained due to a pod disruption budget issue.
        /// </summary>
        [Input("forceUpdateVersion")]
        public Input<bool>? ForceUpdateVersion { get; set; }

        /// <summary>
        /// Set of instance types associated with the EKS Node Group. Defaults to `["t3.medium"]`. This provider will only perform drift detection if a configuration value is provided. Currently, the EKS API only accepts a single value in the set.
        /// </summary>
        [Input("instanceTypes")]
        public Input<string>? InstanceTypes { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Key-value map of Kubernetes labels. Only labels that are applied with the EKS API are managed by this argument. Other Kubernetes labels applied to the EKS Node Group will not be managed.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the EKS Node Group.
        /// </summary>
        [Input("nodeGroupName")]
        public Input<string>? NodeGroupName { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the IAM Role that provides permissions for the EKS Node Group.
        /// </summary>
        [Input("nodeRoleArn")]
        public Input<string>? NodeRoleArn { get; set; }

        /// <summary>
        /// AMI version of the EKS Node Group. Defaults to latest version for Kubernetes version.
        /// </summary>
        [Input("releaseVersion")]
        public Input<string>? ReleaseVersion { get; set; }

        /// <summary>
        /// Configuration block with remote access settings. Detailed below.
        /// </summary>
        [Input("remoteAccess")]
        public Input<Inputs.NodeGroupRemoteAccessGetArgs>? RemoteAccess { get; set; }

        [Input("resources")]
        private InputList<Inputs.NodeGroupResourceGetArgs>? _resources;

        /// <summary>
        /// List of objects containing information about underlying resources.
        /// </summary>
        public InputList<Inputs.NodeGroupResourceGetArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.NodeGroupResourceGetArgs>());
            set => _resources = value;
        }

        /// <summary>
        /// Configuration block with scaling settings. Detailed below.
        /// </summary>
        [Input("scalingConfig")]
        public Input<Inputs.NodeGroupScalingConfigGetArgs>? ScalingConfig { get; set; }

        /// <summary>
        /// Status of the EKS Node Group.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// Identifiers of EC2 Subnets to associate with the EKS Node Group. These subnets must have the following resource tag: `kubernetes.io/cluster/CLUSTER_NAME` (where `CLUSTER_NAME` is replaced with the name of the EKS Cluster).
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Kubernetes version. Defaults to EKS Cluster Kubernetes version. This provider will only perform drift detection if a configuration value is provided.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public NodeGroupState()
        {
        }
    }
}
