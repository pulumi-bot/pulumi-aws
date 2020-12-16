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
    /// Manages an EKS Cluster.
    /// 
    /// ## Example Usage
    /// ### Example IAM Role for EKS Cluster
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Iam.Role("example", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Effect"": ""Allow"",
    ///       ""Principal"": {
    ///         ""Service"": ""eks.amazonaws.com""
    ///       },
    ///       ""Action"": ""sts:AssumeRole""
    ///     }
    ///   ]
    /// }
    /// ",
    ///         });
    ///         var example_AmazonEKSClusterPolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSClusterPolicy", new Aws.Iam.RolePolicyAttachmentArgs
    ///         {
    ///             PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
    ///             Role = example.Name,
    ///         });
    ///         // Optionally, enable Security Groups for Pods
    ///         // Reference: https://docs.aws.amazon.com/eks/latest/userguide/security-groups-for-pods.html
    ///         var example_AmazonEKSVPCResourceController = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSVPCResourceController", new Aws.Iam.RolePolicyAttachmentArgs
    ///         {
    ///             PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSVPCResourceController",
    ///             Role = example.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Enabling Control Plane Logging
    /// 
    /// [EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) can be enabled via the `enabled_cluster_log_types` argument. To manage the CloudWatch Log Group retention period, the `aws.cloudwatch.LogGroup` resource can be used.
    /// 
    /// &gt; The below configuration uses [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) to prevent ordering issues with EKS automatically creating the log group first and a variable for naming consistency. Other ordering and naming methodologies may be more appropriate for your environment.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var clusterName = config.Get("clusterName") ?? "example";
    ///         var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup", new Aws.CloudWatch.LogGroupArgs
    ///         {
    ///             RetentionInDays = 7,
    ///         });
    ///         // ... potentially other configuration ...
    ///         var exampleCluster = new Aws.Eks.Cluster("exampleCluster", new Aws.Eks.ClusterArgs
    ///         {
    ///             EnabledClusterLogTypes = 
    ///             {
    ///                 "api",
    ///                 "audit",
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleLogGroup,
    ///             },
    ///         });
    ///         // ... other configuration ...
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// EKS Clusters can be imported using the `name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:eks/cluster:Cluster my_cluster my_cluster
    /// ```
    /// </summary>
    [AwsResourceType("aws:eks/cluster:Cluster")]
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the cluster.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Nested attribute containing `certificate-authority-data` for your cluster.
        /// </summary>
        [Output("certificateAuthority")]
        public Output<Outputs.ClusterCertificateAuthority> CertificateAuthority { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
        /// </summary>
        [Output("enabledClusterLogTypes")]
        public Output<ImmutableArray<string>> EnabledClusterLogTypes { get; private set; } = null!;

        /// <summary>
        /// Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
        /// </summary>
        [Output("encryptionConfig")]
        public Output<Outputs.ClusterEncryptionConfig?> EncryptionConfig { get; private set; } = null!;

        /// <summary>
        /// The endpoint for your Kubernetes API server.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Nested attribute containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019.
        /// </summary>
        [Output("identities")]
        public Output<ImmutableArray<Outputs.ClusterIdentity>> Identities { get; private set; } = null!;

        [Output("kubernetesNetworkConfig")]
        public Output<Outputs.ClusterKubernetesNetworkConfig> KubernetesNetworkConfig { get; private set; } = null!;

        /// <summary>
        /// Name of the cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The platform version for the cluster.
        /// </summary>
        [Output("platformVersion")]
        public Output<string> PlatformVersion { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the `aws.iam.RolePolicy` resource) or `aws.iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// The status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
        /// </summary>
        [Output("vpcConfig")]
        public Output<Outputs.ClusterVpcConfig> VpcConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("aws:eks/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("aws:eks/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        [Input("enabledClusterLogTypes")]
        private InputList<string>? _enabledClusterLogTypes;

        /// <summary>
        /// A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
        /// </summary>
        public InputList<string> EnabledClusterLogTypes
        {
            get => _enabledClusterLogTypes ?? (_enabledClusterLogTypes = new InputList<string>());
            set => _enabledClusterLogTypes = value;
        }

        /// <summary>
        /// Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.ClusterEncryptionConfigArgs>? EncryptionConfig { get; set; }

        [Input("kubernetesNetworkConfig")]
        public Input<Inputs.ClusterKubernetesNetworkConfigArgs>? KubernetesNetworkConfig { get; set; }

        /// <summary>
        /// Name of the cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the `aws.iam.RolePolicy` resource) or `aws.iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
        /// </summary>
        [Input("vpcConfig", required: true)]
        public Input<Inputs.ClusterVpcConfigArgs> VpcConfig { get; set; } = null!;

        public ClusterArgs()
        {
        }
    }

    public sealed class ClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the cluster.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Nested attribute containing `certificate-authority-data` for your cluster.
        /// </summary>
        [Input("certificateAuthority")]
        public Input<Inputs.ClusterCertificateAuthorityGetArgs>? CertificateAuthority { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("enabledClusterLogTypes")]
        private InputList<string>? _enabledClusterLogTypes;

        /// <summary>
        /// A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
        /// </summary>
        public InputList<string> EnabledClusterLogTypes
        {
            get => _enabledClusterLogTypes ?? (_enabledClusterLogTypes = new InputList<string>());
            set => _enabledClusterLogTypes = value;
        }

        /// <summary>
        /// Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
        /// </summary>
        [Input("encryptionConfig")]
        public Input<Inputs.ClusterEncryptionConfigGetArgs>? EncryptionConfig { get; set; }

        /// <summary>
        /// The endpoint for your Kubernetes API server.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        [Input("identities")]
        private InputList<Inputs.ClusterIdentityGetArgs>? _identities;

        /// <summary>
        /// Nested attribute containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019.
        /// </summary>
        public InputList<Inputs.ClusterIdentityGetArgs> Identities
        {
            get => _identities ?? (_identities = new InputList<Inputs.ClusterIdentityGetArgs>());
            set => _identities = value;
        }

        [Input("kubernetesNetworkConfig")]
        public Input<Inputs.ClusterKubernetesNetworkConfigGetArgs>? KubernetesNetworkConfig { get; set; }

        /// <summary>
        /// Name of the cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The platform version for the cluster.
        /// </summary>
        [Input("platformVersion")]
        public Input<string>? PlatformVersion { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) if using the `aws.iam.RolePolicy` resource) or `aws.iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// The status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
        /// </summary>
        [Input("vpcConfig")]
        public Input<Inputs.ClusterVpcConfigGetArgs>? VpcConfig { get; set; }

        public ClusterState()
        {
        }
    }
}
