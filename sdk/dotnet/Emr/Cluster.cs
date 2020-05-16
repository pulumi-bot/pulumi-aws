// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Emr
{
    /// <summary>
    /// Provides an Elastic MapReduce Cluster, a web service that makes it easy to
    /// process large amounts of data efficiently. See [Amazon Elastic MapReduce Documentation](https://aws.amazon.com/documentation/elastic-mapreduce/)
    /// for more information.
    /// 
    /// To configure [Instance Groups](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for [task nodes](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-task), see the [`aws.emr.InstanceGroup` resource](https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html).
    /// 
    /// &gt; Support for [Instance Fleets](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-fleets) will be made available in an upcoming release.
    /// </summary>
    public partial class Cluster : Pulumi.CustomResource
    {
        /// <summary>
        /// A JSON string for selecting additional features such as adding proxy information. Note: Currently there is no API to retrieve the value of this argument after EMR cluster creation from provider, therefore this provider cannot detect drift from the actual EMR cluster if its value is changed outside this provider.
        /// </summary>
        [Output("additionalInfo")]
        public Output<string?> AdditionalInfo { get; private set; } = null!;

        /// <summary>
        /// A list of applications for the cluster. Valid values are: `Flink`, `Hadoop`, `Hive`, `Mahout`, `Pig`, `Spark`, and `JupyterHub` (as of EMR 5.14.0). Case insensitive
        /// </summary>
        [Output("applications")]
        public Output<ImmutableArray<string>> Applications { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// An IAM role for automatic scaling policies. The IAM role provides permissions that the automatic scaling feature requires to launch and terminate EC2 instances in an instance group.
        /// </summary>
        [Output("autoscalingRole")]
        public Output<string?> AutoscalingRole { get; private set; } = null!;

        /// <summary>
        /// Ordered list of bootstrap actions that will be run before Hadoop is started on the cluster nodes. Defined below.
        /// </summary>
        [Output("bootstrapActions")]
        public Output<ImmutableArray<Outputs.ClusterBootstrapAction>> BootstrapActions { get; private set; } = null!;

        [Output("clusterState")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// List of configurations supplied for the EMR cluster you are creating
        /// </summary>
        [Output("configurations")]
        public Output<string?> Configurations { get; private set; } = null!;

        /// <summary>
        /// A JSON string for supplying list of configurations for the EMR cluster.
        /// </summary>
        [Output("configurationsJson")]
        public Output<string?> ConfigurationsJson { get; private set; } = null!;

        /// <summary>
        /// Use the `core_instance_group` configuration block `instance_count` argument instead. Number of Amazon EC2 instances used to execute the job flow. EMR will use one node as the cluster's master node and use the remainder of the nodes (`core_instance_count`-1) as core nodes. Cannot be specified if `core_instance_group` or `instance_group` configuration blocks are set. Default `1`
        /// </summary>
        [Output("coreInstanceCount")]
        public Output<int> CoreInstanceCount { get; private set; } = null!;

        /// <summary>
        /// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [core node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-core). Cannot be specified if `core_instance_count` argument, `core_instance_type` argument, or `instance_group` configuration blocks are set. Detailed below.
        /// </summary>
        [Output("coreInstanceGroup")]
        public Output<Outputs.ClusterCoreInstanceGroup> CoreInstanceGroup { get; private set; } = null!;

        /// <summary>
        /// Use the `core_instance_group` configuration block `instance_type` argument instead. The EC2 instance type of the slave nodes. Cannot be specified if `core_instance_group` or `instance_group` configuration blocks are set.
        /// </summary>
        [Output("coreInstanceType")]
        public Output<string> CoreInstanceType { get; private set; } = null!;

        /// <summary>
        /// A custom Amazon Linux AMI for the cluster (instead of an EMR-owned AMI). Available in Amazon EMR version 5.7.0 and later.
        /// </summary>
        [Output("customAmiId")]
        public Output<string?> CustomAmiId { get; private set; } = null!;

        /// <summary>
        /// Size in GiB of the EBS root device volume of the Linux AMI that is used for each EC2 instance. Available in Amazon EMR version 4.x and later.
        /// </summary>
        [Output("ebsRootVolumeSize")]
        public Output<int?> EbsRootVolumeSize { get; private set; } = null!;

        /// <summary>
        /// Attributes for the EC2 instances running the job flow. Defined below
        /// </summary>
        [Output("ec2Attributes")]
        public Output<Outputs.ClusterEc2Attributes?> Ec2Attributes { get; private set; } = null!;

        /// <summary>
        /// Use the `master_instance_group` configuration block, `core_instance_group` configuration block and [`aws.emr.InstanceGroup` resource(s)](https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html) instead. A list of `instance_group` objects for each instance group in the cluster. Exactly one of `master_instance_type` and `instance_group` must be specified. If `instance_group` is set, then it must contain a configuration block for at least the `MASTER` instance group type (as well as any additional instance groups). Cannot be specified if `master_instance_group` or `core_instance_group` configuration blocks are set. Defined below
        /// </summary>
        [Output("instanceGroups")]
        public Output<ImmutableArray<Outputs.ClusterInstanceGroup>> InstanceGroups { get; private set; } = null!;

        /// <summary>
        /// Switch on/off run cluster with no steps or when all steps are complete (default is on)
        /// </summary>
        [Output("keepJobFlowAliveWhenNoSteps")]
        public Output<bool> KeepJobFlowAliveWhenNoSteps { get; private set; } = null!;

        /// <summary>
        /// Kerberos configuration for the cluster. Defined below
        /// </summary>
        [Output("kerberosAttributes")]
        public Output<Outputs.ClusterKerberosAttributes?> KerberosAttributes { get; private set; } = null!;

        /// <summary>
        /// S3 bucket to write the log files of the job flow. If a value is not provided, logs are not created
        /// </summary>
        [Output("logUri")]
        public Output<string?> LogUri { get; private set; } = null!;

        /// <summary>
        /// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [master node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-master). Cannot be specified if `master_instance_type` argument or `instance_group` configuration blocks are set. Detailed below.
        /// </summary>
        [Output("masterInstanceGroup")]
        public Output<Outputs.ClusterMasterInstanceGroup> MasterInstanceGroup { get; private set; } = null!;

        /// <summary>
        /// Use the `master_instance_group` configuration block `instance_type` argument instead. The EC2 instance type of the master node. Cannot be specified if `master_instance_group` or `instance_group` configuration blocks are set.
        /// </summary>
        [Output("masterInstanceType")]
        public Output<string> MasterInstanceType { get; private set; } = null!;

        /// <summary>
        /// The public DNS name of the master EC2 instance.
        /// * `core_instance_group.0.id` - Core node type Instance Group ID, if using Instance Group for this node type.
        /// </summary>
        [Output("masterPublicDns")]
        public Output<string> MasterPublicDns { get; private set; } = null!;

        /// <summary>
        /// The name of the step.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The release label for the Amazon EMR release
        /// </summary>
        [Output("releaseLabel")]
        public Output<string> ReleaseLabel { get; private set; } = null!;

        /// <summary>
        /// The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an `instance group` is resized.
        /// </summary>
        [Output("scaleDownBehavior")]
        public Output<string> ScaleDownBehavior { get; private set; } = null!;

        /// <summary>
        /// The security configuration name to attach to the EMR cluster. Only valid for EMR clusters with `release_label` 4.8.0 or greater
        /// </summary>
        [Output("securityConfiguration")]
        public Output<string?> SecurityConfiguration { get; private set; } = null!;

        /// <summary>
        /// IAM role that will be assumed by the Amazon EMR service to access AWS resources
        /// </summary>
        [Output("serviceRole")]
        public Output<string> ServiceRole { get; private set; } = null!;

        /// <summary>
        /// The number of steps that can be executed concurrently. You can specify a maximum of 256 steps. Only valid for EMR clusters with `release_label` 5.28.0 or greater. (default is 1)
        /// </summary>
        [Output("stepConcurrencyLevel")]
        public Output<int?> StepConcurrencyLevel { get; private set; } = null!;

        /// <summary>
        /// List of steps to run when creating the cluster. Defined below. It is highly recommended to utilize [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) if other steps are being managed outside of this provider.
        /// </summary>
        [Output("steps")]
        public Output<ImmutableArray<Outputs.ClusterStep>> Steps { get; private set; } = null!;

        /// <summary>
        /// list of tags to apply to the EMR Cluster
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Switch on/off termination protection (default is `false`, except when using multiple master nodes). Before attempting to destroy the resource when termination protection is enabled, this configuration must be applied with its value set to `false`.
        /// </summary>
        [Output("terminationProtection")]
        public Output<bool> TerminationProtection { get; private set; } = null!;

        /// <summary>
        /// Whether the job flow is visible to all IAM users of the AWS account associated with the job flow. Default `true`
        /// </summary>
        [Output("visibleToAllUsers")]
        public Output<bool?> VisibleToAllUsers { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("aws:emr/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("aws:emr/cluster:Cluster", name, state, MakeResourceOptions(options, id))
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
        /// A JSON string for selecting additional features such as adding proxy information. Note: Currently there is no API to retrieve the value of this argument after EMR cluster creation from provider, therefore this provider cannot detect drift from the actual EMR cluster if its value is changed outside this provider.
        /// </summary>
        [Input("additionalInfo")]
        public Input<string>? AdditionalInfo { get; set; }

        [Input("applications")]
        private InputList<string>? _applications;

        /// <summary>
        /// A list of applications for the cluster. Valid values are: `Flink`, `Hadoop`, `Hive`, `Mahout`, `Pig`, `Spark`, and `JupyterHub` (as of EMR 5.14.0). Case insensitive
        /// </summary>
        public InputList<string> Applications
        {
            get => _applications ?? (_applications = new InputList<string>());
            set => _applications = value;
        }

        /// <summary>
        /// An IAM role for automatic scaling policies. The IAM role provides permissions that the automatic scaling feature requires to launch and terminate EC2 instances in an instance group.
        /// </summary>
        [Input("autoscalingRole")]
        public Input<string>? AutoscalingRole { get; set; }

        [Input("bootstrapActions")]
        private InputList<Inputs.ClusterBootstrapActionArgs>? _bootstrapActions;

        /// <summary>
        /// Ordered list of bootstrap actions that will be run before Hadoop is started on the cluster nodes. Defined below.
        /// </summary>
        public InputList<Inputs.ClusterBootstrapActionArgs> BootstrapActions
        {
            get => _bootstrapActions ?? (_bootstrapActions = new InputList<Inputs.ClusterBootstrapActionArgs>());
            set => _bootstrapActions = value;
        }

        /// <summary>
        /// List of configurations supplied for the EMR cluster you are creating
        /// </summary>
        [Input("configurations")]
        public Input<string>? Configurations { get; set; }

        /// <summary>
        /// A JSON string for supplying list of configurations for the EMR cluster.
        /// </summary>
        [Input("configurationsJson")]
        public Input<string>? ConfigurationsJson { get; set; }

        /// <summary>
        /// Use the `core_instance_group` configuration block `instance_count` argument instead. Number of Amazon EC2 instances used to execute the job flow. EMR will use one node as the cluster's master node and use the remainder of the nodes (`core_instance_count`-1) as core nodes. Cannot be specified if `core_instance_group` or `instance_group` configuration blocks are set. Default `1`
        /// </summary>
        [Input("coreInstanceCount")]
        public Input<int>? CoreInstanceCount { get; set; }

        /// <summary>
        /// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [core node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-core). Cannot be specified if `core_instance_count` argument, `core_instance_type` argument, or `instance_group` configuration blocks are set. Detailed below.
        /// </summary>
        [Input("coreInstanceGroup")]
        public Input<Inputs.ClusterCoreInstanceGroupArgs>? CoreInstanceGroup { get; set; }

        /// <summary>
        /// Use the `core_instance_group` configuration block `instance_type` argument instead. The EC2 instance type of the slave nodes. Cannot be specified if `core_instance_group` or `instance_group` configuration blocks are set.
        /// </summary>
        [Input("coreInstanceType")]
        public Input<string>? CoreInstanceType { get; set; }

        /// <summary>
        /// A custom Amazon Linux AMI for the cluster (instead of an EMR-owned AMI). Available in Amazon EMR version 5.7.0 and later.
        /// </summary>
        [Input("customAmiId")]
        public Input<string>? CustomAmiId { get; set; }

        /// <summary>
        /// Size in GiB of the EBS root device volume of the Linux AMI that is used for each EC2 instance. Available in Amazon EMR version 4.x and later.
        /// </summary>
        [Input("ebsRootVolumeSize")]
        public Input<int>? EbsRootVolumeSize { get; set; }

        /// <summary>
        /// Attributes for the EC2 instances running the job flow. Defined below
        /// </summary>
        [Input("ec2Attributes")]
        public Input<Inputs.ClusterEc2AttributesArgs>? Ec2Attributes { get; set; }

        [Input("instanceGroups")]
        private InputList<Inputs.ClusterInstanceGroupArgs>? _instanceGroups;

        /// <summary>
        /// Use the `master_instance_group` configuration block, `core_instance_group` configuration block and [`aws.emr.InstanceGroup` resource(s)](https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html) instead. A list of `instance_group` objects for each instance group in the cluster. Exactly one of `master_instance_type` and `instance_group` must be specified. If `instance_group` is set, then it must contain a configuration block for at least the `MASTER` instance group type (as well as any additional instance groups). Cannot be specified if `master_instance_group` or `core_instance_group` configuration blocks are set. Defined below
        /// </summary>
        [Obsolete(@"use `[masterInstanceGroup](#/resources/aws:emr%2Fcluster:Cluster/inputProperties/masterInstanceGroup)` configuration block, `[coreInstanceGroup](#/resources/aws:emr%2Fcluster:Cluster/inputProperties/coreInstanceGroup)` configuration block, and `[aws:emr/instanceGroup:InstanceGroup](#/resources//aws:emr%252FinstanceGroup:InstanceGroup)` resource(s) instead")]
        public InputList<Inputs.ClusterInstanceGroupArgs> InstanceGroups
        {
            get => _instanceGroups ?? (_instanceGroups = new InputList<Inputs.ClusterInstanceGroupArgs>());
            set => _instanceGroups = value;
        }

        /// <summary>
        /// Switch on/off run cluster with no steps or when all steps are complete (default is on)
        /// </summary>
        [Input("keepJobFlowAliveWhenNoSteps")]
        public Input<bool>? KeepJobFlowAliveWhenNoSteps { get; set; }

        /// <summary>
        /// Kerberos configuration for the cluster. Defined below
        /// </summary>
        [Input("kerberosAttributes")]
        public Input<Inputs.ClusterKerberosAttributesArgs>? KerberosAttributes { get; set; }

        /// <summary>
        /// S3 bucket to write the log files of the job flow. If a value is not provided, logs are not created
        /// </summary>
        [Input("logUri")]
        public Input<string>? LogUri { get; set; }

        /// <summary>
        /// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [master node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-master). Cannot be specified if `master_instance_type` argument or `instance_group` configuration blocks are set. Detailed below.
        /// </summary>
        [Input("masterInstanceGroup")]
        public Input<Inputs.ClusterMasterInstanceGroupArgs>? MasterInstanceGroup { get; set; }

        /// <summary>
        /// Use the `master_instance_group` configuration block `instance_type` argument instead. The EC2 instance type of the master node. Cannot be specified if `master_instance_group` or `instance_group` configuration blocks are set.
        /// </summary>
        [Input("masterInstanceType")]
        public Input<string>? MasterInstanceType { get; set; }

        /// <summary>
        /// The name of the step.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The release label for the Amazon EMR release
        /// </summary>
        [Input("releaseLabel", required: true)]
        public Input<string> ReleaseLabel { get; set; } = null!;

        /// <summary>
        /// The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an `instance group` is resized.
        /// </summary>
        [Input("scaleDownBehavior")]
        public Input<string>? ScaleDownBehavior { get; set; }

        /// <summary>
        /// The security configuration name to attach to the EMR cluster. Only valid for EMR clusters with `release_label` 4.8.0 or greater
        /// </summary>
        [Input("securityConfiguration")]
        public Input<string>? SecurityConfiguration { get; set; }

        /// <summary>
        /// IAM role that will be assumed by the Amazon EMR service to access AWS resources
        /// </summary>
        [Input("serviceRole", required: true)]
        public Input<string> ServiceRole { get; set; } = null!;

        /// <summary>
        /// The number of steps that can be executed concurrently. You can specify a maximum of 256 steps. Only valid for EMR clusters with `release_label` 5.28.0 or greater. (default is 1)
        /// </summary>
        [Input("stepConcurrencyLevel")]
        public Input<int>? StepConcurrencyLevel { get; set; }

        [Input("steps")]
        private InputList<Inputs.ClusterStepArgs>? _steps;

        /// <summary>
        /// List of steps to run when creating the cluster. Defined below. It is highly recommended to utilize [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) if other steps are being managed outside of this provider.
        /// </summary>
        public InputList<Inputs.ClusterStepArgs> Steps
        {
            get => _steps ?? (_steps = new InputList<Inputs.ClusterStepArgs>());
            set => _steps = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// list of tags to apply to the EMR Cluster
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Switch on/off termination protection (default is `false`, except when using multiple master nodes). Before attempting to destroy the resource when termination protection is enabled, this configuration must be applied with its value set to `false`.
        /// </summary>
        [Input("terminationProtection")]
        public Input<bool>? TerminationProtection { get; set; }

        /// <summary>
        /// Whether the job flow is visible to all IAM users of the AWS account associated with the job flow. Default `true`
        /// </summary>
        [Input("visibleToAllUsers")]
        public Input<bool>? VisibleToAllUsers { get; set; }

        public ClusterArgs()
        {
        }
    }

    public sealed class ClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A JSON string for selecting additional features such as adding proxy information. Note: Currently there is no API to retrieve the value of this argument after EMR cluster creation from provider, therefore this provider cannot detect drift from the actual EMR cluster if its value is changed outside this provider.
        /// </summary>
        [Input("additionalInfo")]
        public Input<string>? AdditionalInfo { get; set; }

        [Input("applications")]
        private InputList<string>? _applications;

        /// <summary>
        /// A list of applications for the cluster. Valid values are: `Flink`, `Hadoop`, `Hive`, `Mahout`, `Pig`, `Spark`, and `JupyterHub` (as of EMR 5.14.0). Case insensitive
        /// </summary>
        public InputList<string> Applications
        {
            get => _applications ?? (_applications = new InputList<string>());
            set => _applications = value;
        }

        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// An IAM role for automatic scaling policies. The IAM role provides permissions that the automatic scaling feature requires to launch and terminate EC2 instances in an instance group.
        /// </summary>
        [Input("autoscalingRole")]
        public Input<string>? AutoscalingRole { get; set; }

        [Input("bootstrapActions")]
        private InputList<Inputs.ClusterBootstrapActionGetArgs>? _bootstrapActions;

        /// <summary>
        /// Ordered list of bootstrap actions that will be run before Hadoop is started on the cluster nodes. Defined below.
        /// </summary>
        public InputList<Inputs.ClusterBootstrapActionGetArgs> BootstrapActions
        {
            get => _bootstrapActions ?? (_bootstrapActions = new InputList<Inputs.ClusterBootstrapActionGetArgs>());
            set => _bootstrapActions = value;
        }

        [Input("clusterState")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// List of configurations supplied for the EMR cluster you are creating
        /// </summary>
        [Input("configurations")]
        public Input<string>? Configurations { get; set; }

        /// <summary>
        /// A JSON string for supplying list of configurations for the EMR cluster.
        /// </summary>
        [Input("configurationsJson")]
        public Input<string>? ConfigurationsJson { get; set; }

        /// <summary>
        /// Use the `core_instance_group` configuration block `instance_count` argument instead. Number of Amazon EC2 instances used to execute the job flow. EMR will use one node as the cluster's master node and use the remainder of the nodes (`core_instance_count`-1) as core nodes. Cannot be specified if `core_instance_group` or `instance_group` configuration blocks are set. Default `1`
        /// </summary>
        [Input("coreInstanceCount")]
        public Input<int>? CoreInstanceCount { get; set; }

        /// <summary>
        /// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [core node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-core). Cannot be specified if `core_instance_count` argument, `core_instance_type` argument, or `instance_group` configuration blocks are set. Detailed below.
        /// </summary>
        [Input("coreInstanceGroup")]
        public Input<Inputs.ClusterCoreInstanceGroupGetArgs>? CoreInstanceGroup { get; set; }

        /// <summary>
        /// Use the `core_instance_group` configuration block `instance_type` argument instead. The EC2 instance type of the slave nodes. Cannot be specified if `core_instance_group` or `instance_group` configuration blocks are set.
        /// </summary>
        [Input("coreInstanceType")]
        public Input<string>? CoreInstanceType { get; set; }

        /// <summary>
        /// A custom Amazon Linux AMI for the cluster (instead of an EMR-owned AMI). Available in Amazon EMR version 5.7.0 and later.
        /// </summary>
        [Input("customAmiId")]
        public Input<string>? CustomAmiId { get; set; }

        /// <summary>
        /// Size in GiB of the EBS root device volume of the Linux AMI that is used for each EC2 instance. Available in Amazon EMR version 4.x and later.
        /// </summary>
        [Input("ebsRootVolumeSize")]
        public Input<int>? EbsRootVolumeSize { get; set; }

        /// <summary>
        /// Attributes for the EC2 instances running the job flow. Defined below
        /// </summary>
        [Input("ec2Attributes")]
        public Input<Inputs.ClusterEc2AttributesGetArgs>? Ec2Attributes { get; set; }

        [Input("instanceGroups")]
        private InputList<Inputs.ClusterInstanceGroupGetArgs>? _instanceGroups;

        /// <summary>
        /// Use the `master_instance_group` configuration block, `core_instance_group` configuration block and [`aws.emr.InstanceGroup` resource(s)](https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html) instead. A list of `instance_group` objects for each instance group in the cluster. Exactly one of `master_instance_type` and `instance_group` must be specified. If `instance_group` is set, then it must contain a configuration block for at least the `MASTER` instance group type (as well as any additional instance groups). Cannot be specified if `master_instance_group` or `core_instance_group` configuration blocks are set. Defined below
        /// </summary>
        [Obsolete(@"use `[masterInstanceGroup](#/resources/aws:emr%2Fcluster:Cluster/properties/masterInstanceGroup)` configuration block, `[coreInstanceGroup](#/resources/aws:emr%2Fcluster:Cluster/properties/coreInstanceGroup)` configuration block, and `[aws:emr/instanceGroup:InstanceGroup](#/resources//aws:emr%252FinstanceGroup:InstanceGroup)` resource(s) instead")]
        public InputList<Inputs.ClusterInstanceGroupGetArgs> InstanceGroups
        {
            get => _instanceGroups ?? (_instanceGroups = new InputList<Inputs.ClusterInstanceGroupGetArgs>());
            set => _instanceGroups = value;
        }

        /// <summary>
        /// Switch on/off run cluster with no steps or when all steps are complete (default is on)
        /// </summary>
        [Input("keepJobFlowAliveWhenNoSteps")]
        public Input<bool>? KeepJobFlowAliveWhenNoSteps { get; set; }

        /// <summary>
        /// Kerberos configuration for the cluster. Defined below
        /// </summary>
        [Input("kerberosAttributes")]
        public Input<Inputs.ClusterKerberosAttributesGetArgs>? KerberosAttributes { get; set; }

        /// <summary>
        /// S3 bucket to write the log files of the job flow. If a value is not provided, logs are not created
        /// </summary>
        [Input("logUri")]
        public Input<string>? LogUri { get; set; }

        /// <summary>
        /// Configuration block to use an [Instance Group](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-instance-group-configuration.html#emr-plan-instance-groups) for the [master node type](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html#emr-plan-master). Cannot be specified if `master_instance_type` argument or `instance_group` configuration blocks are set. Detailed below.
        /// </summary>
        [Input("masterInstanceGroup")]
        public Input<Inputs.ClusterMasterInstanceGroupGetArgs>? MasterInstanceGroup { get; set; }

        /// <summary>
        /// Use the `master_instance_group` configuration block `instance_type` argument instead. The EC2 instance type of the master node. Cannot be specified if `master_instance_group` or `instance_group` configuration blocks are set.
        /// </summary>
        [Input("masterInstanceType")]
        public Input<string>? MasterInstanceType { get; set; }

        /// <summary>
        /// The public DNS name of the master EC2 instance.
        /// * `core_instance_group.0.id` - Core node type Instance Group ID, if using Instance Group for this node type.
        /// </summary>
        [Input("masterPublicDns")]
        public Input<string>? MasterPublicDns { get; set; }

        /// <summary>
        /// The name of the step.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The release label for the Amazon EMR release
        /// </summary>
        [Input("releaseLabel")]
        public Input<string>? ReleaseLabel { get; set; }

        /// <summary>
        /// The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an `instance group` is resized.
        /// </summary>
        [Input("scaleDownBehavior")]
        public Input<string>? ScaleDownBehavior { get; set; }

        /// <summary>
        /// The security configuration name to attach to the EMR cluster. Only valid for EMR clusters with `release_label` 4.8.0 or greater
        /// </summary>
        [Input("securityConfiguration")]
        public Input<string>? SecurityConfiguration { get; set; }

        /// <summary>
        /// IAM role that will be assumed by the Amazon EMR service to access AWS resources
        /// </summary>
        [Input("serviceRole")]
        public Input<string>? ServiceRole { get; set; }

        /// <summary>
        /// The number of steps that can be executed concurrently. You can specify a maximum of 256 steps. Only valid for EMR clusters with `release_label` 5.28.0 or greater. (default is 1)
        /// </summary>
        [Input("stepConcurrencyLevel")]
        public Input<int>? StepConcurrencyLevel { get; set; }

        [Input("steps")]
        private InputList<Inputs.ClusterStepGetArgs>? _steps;

        /// <summary>
        /// List of steps to run when creating the cluster. Defined below. It is highly recommended to utilize [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) if other steps are being managed outside of this provider.
        /// </summary>
        public InputList<Inputs.ClusterStepGetArgs> Steps
        {
            get => _steps ?? (_steps = new InputList<Inputs.ClusterStepGetArgs>());
            set => _steps = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// list of tags to apply to the EMR Cluster
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Switch on/off termination protection (default is `false`, except when using multiple master nodes). Before attempting to destroy the resource when termination protection is enabled, this configuration must be applied with its value set to `false`.
        /// </summary>
        [Input("terminationProtection")]
        public Input<bool>? TerminationProtection { get; set; }

        /// <summary>
        /// Whether the job flow is visible to all IAM users of the AWS account associated with the job flow. Default `true`
        /// </summary>
        [Input("visibleToAllUsers")]
        public Input<bool>? VisibleToAllUsers { get; set; }

        public ClusterState()
        {
        }
    }
}
