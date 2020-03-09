// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.OpsWorks
{
    /// <summary>
    /// Provides an OpsWorks Java application layer resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/opsworks_java_app_layer.html.markdown.
    /// </summary>
    public partial class JavaAppLayer : Pulumi.CustomResource
    {
        /// <summary>
        /// Keyword for the application container to use. Defaults to "tomcat".
        /// </summary>
        [Output("appServer")]
        public Output<string?> AppServer { get; private set; } = null!;

        /// <summary>
        /// Version of the selected application container to use. Defaults to "7".
        /// </summary>
        [Output("appServerVersion")]
        public Output<string?> AppServerVersion { get; private set; } = null!;

        /// <summary>
        /// Whether to automatically assign an elastic IP address to the layer's instances.
        /// </summary>
        [Output("autoAssignElasticIps")]
        public Output<bool?> AutoAssignElasticIps { get; private set; } = null!;

        /// <summary>
        /// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
        /// </summary>
        [Output("autoAssignPublicIps")]
        public Output<bool?> AutoAssignPublicIps { get; private set; } = null!;

        /// <summary>
        /// Whether to enable auto-healing for the layer.
        /// </summary>
        [Output("autoHealing")]
        public Output<bool?> AutoHealing { get; private set; } = null!;

        [Output("customConfigureRecipes")]
        public Output<ImmutableArray<string>> CustomConfigureRecipes { get; private set; } = null!;

        [Output("customDeployRecipes")]
        public Output<ImmutableArray<string>> CustomDeployRecipes { get; private set; } = null!;

        /// <summary>
        /// The ARN of an IAM profile that will be used for the layer's instances.
        /// </summary>
        [Output("customInstanceProfileArn")]
        public Output<string?> CustomInstanceProfileArn { get; private set; } = null!;

        /// <summary>
        /// Custom JSON attributes to apply to the layer.
        /// </summary>
        [Output("customJson")]
        public Output<string?> CustomJson { get; private set; } = null!;

        /// <summary>
        /// Ids for a set of security groups to apply to the layer's instances.
        /// </summary>
        [Output("customSecurityGroupIds")]
        public Output<ImmutableArray<string>> CustomSecurityGroupIds { get; private set; } = null!;

        [Output("customSetupRecipes")]
        public Output<ImmutableArray<string>> CustomSetupRecipes { get; private set; } = null!;

        [Output("customShutdownRecipes")]
        public Output<ImmutableArray<string>> CustomShutdownRecipes { get; private set; } = null!;

        [Output("customUndeployRecipes")]
        public Output<ImmutableArray<string>> CustomUndeployRecipes { get; private set; } = null!;

        /// <summary>
        /// Whether to enable Elastic Load Balancing connection draining.
        /// </summary>
        [Output("drainElbOnShutdown")]
        public Output<bool?> DrainElbOnShutdown { get; private set; } = null!;

        /// <summary>
        /// `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
        /// </summary>
        [Output("ebsVolumes")]
        public Output<ImmutableArray<Outputs.JavaAppLayerEbsVolumes>> EbsVolumes { get; private set; } = null!;

        /// <summary>
        /// Name of an Elastic Load Balancer to attach to this layer
        /// </summary>
        [Output("elasticLoadBalancer")]
        public Output<string?> ElasticLoadBalancer { get; private set; } = null!;

        /// <summary>
        /// Whether to install OS and package updates on each instance when it boots.
        /// </summary>
        [Output("installUpdatesOnBoot")]
        public Output<bool?> InstallUpdatesOnBoot { get; private set; } = null!;

        /// <summary>
        /// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
        /// </summary>
        [Output("instanceShutdownTimeout")]
        public Output<int?> InstanceShutdownTimeout { get; private set; } = null!;

        /// <summary>
        /// Options to set for the JVM.
        /// </summary>
        [Output("jvmOptions")]
        public Output<string?> JvmOptions { get; private set; } = null!;

        /// <summary>
        /// Keyword for the type of JVM to use. Defaults to `openjdk`.
        /// </summary>
        [Output("jvmType")]
        public Output<string?> JvmType { get; private set; } = null!;

        /// <summary>
        /// Version of JVM to use. Defaults to "7".
        /// </summary>
        [Output("jvmVersion")]
        public Output<string?> JvmVersion { get; private set; } = null!;

        /// <summary>
        /// A human-readable name for the layer.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The id of the stack the layer will belong to.
        /// </summary>
        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;

        /// <summary>
        /// Names of a set of system packages to install on the layer's instances.
        /// </summary>
        [Output("systemPackages")]
        public Output<ImmutableArray<string>> SystemPackages { get; private set; } = null!;

        /// <summary>
        /// Whether to use EBS-optimized instances.
        /// </summary>
        [Output("useEbsOptimizedInstances")]
        public Output<bool?> UseEbsOptimizedInstances { get; private set; } = null!;


        /// <summary>
        /// Create a JavaAppLayer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public JavaAppLayer(string name, JavaAppLayerArgs args, CustomResourceOptions? options = null)
            : base("aws:opsworks/javaAppLayer:JavaAppLayer", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private JavaAppLayer(string name, Input<string> id, JavaAppLayerState? state = null, CustomResourceOptions? options = null)
            : base("aws:opsworks/javaAppLayer:JavaAppLayer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing JavaAppLayer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static JavaAppLayer Get(string name, Input<string> id, JavaAppLayerState? state = null, CustomResourceOptions? options = null)
        {
            return new JavaAppLayer(name, id, state, options);
        }
    }

    public sealed class JavaAppLayerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Keyword for the application container to use. Defaults to "tomcat".
        /// </summary>
        [Input("appServer")]
        public Input<string>? AppServer { get; set; }

        /// <summary>
        /// Version of the selected application container to use. Defaults to "7".
        /// </summary>
        [Input("appServerVersion")]
        public Input<string>? AppServerVersion { get; set; }

        /// <summary>
        /// Whether to automatically assign an elastic IP address to the layer's instances.
        /// </summary>
        [Input("autoAssignElasticIps")]
        public Input<bool>? AutoAssignElasticIps { get; set; }

        /// <summary>
        /// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
        /// </summary>
        [Input("autoAssignPublicIps")]
        public Input<bool>? AutoAssignPublicIps { get; set; }

        /// <summary>
        /// Whether to enable auto-healing for the layer.
        /// </summary>
        [Input("autoHealing")]
        public Input<bool>? AutoHealing { get; set; }

        [Input("customConfigureRecipes")]
        private InputList<string>? _customConfigureRecipes;
        public InputList<string> CustomConfigureRecipes
        {
            get => _customConfigureRecipes ?? (_customConfigureRecipes = new InputList<string>());
            set => _customConfigureRecipes = value;
        }

        [Input("customDeployRecipes")]
        private InputList<string>? _customDeployRecipes;
        public InputList<string> CustomDeployRecipes
        {
            get => _customDeployRecipes ?? (_customDeployRecipes = new InputList<string>());
            set => _customDeployRecipes = value;
        }

        /// <summary>
        /// The ARN of an IAM profile that will be used for the layer's instances.
        /// </summary>
        [Input("customInstanceProfileArn")]
        public Input<string>? CustomInstanceProfileArn { get; set; }

        /// <summary>
        /// Custom JSON attributes to apply to the layer.
        /// </summary>
        [Input("customJson")]
        public Input<string>? CustomJson { get; set; }

        [Input("customSecurityGroupIds")]
        private InputList<string>? _customSecurityGroupIds;

        /// <summary>
        /// Ids for a set of security groups to apply to the layer's instances.
        /// </summary>
        public InputList<string> CustomSecurityGroupIds
        {
            get => _customSecurityGroupIds ?? (_customSecurityGroupIds = new InputList<string>());
            set => _customSecurityGroupIds = value;
        }

        [Input("customSetupRecipes")]
        private InputList<string>? _customSetupRecipes;
        public InputList<string> CustomSetupRecipes
        {
            get => _customSetupRecipes ?? (_customSetupRecipes = new InputList<string>());
            set => _customSetupRecipes = value;
        }

        [Input("customShutdownRecipes")]
        private InputList<string>? _customShutdownRecipes;
        public InputList<string> CustomShutdownRecipes
        {
            get => _customShutdownRecipes ?? (_customShutdownRecipes = new InputList<string>());
            set => _customShutdownRecipes = value;
        }

        [Input("customUndeployRecipes")]
        private InputList<string>? _customUndeployRecipes;
        public InputList<string> CustomUndeployRecipes
        {
            get => _customUndeployRecipes ?? (_customUndeployRecipes = new InputList<string>());
            set => _customUndeployRecipes = value;
        }

        /// <summary>
        /// Whether to enable Elastic Load Balancing connection draining.
        /// </summary>
        [Input("drainElbOnShutdown")]
        public Input<bool>? DrainElbOnShutdown { get; set; }

        [Input("ebsVolumes")]
        private InputList<Inputs.JavaAppLayerEbsVolumesArgs>? _ebsVolumes;

        /// <summary>
        /// `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
        /// </summary>
        public InputList<Inputs.JavaAppLayerEbsVolumesArgs> EbsVolumes
        {
            get => _ebsVolumes ?? (_ebsVolumes = new InputList<Inputs.JavaAppLayerEbsVolumesArgs>());
            set => _ebsVolumes = value;
        }

        /// <summary>
        /// Name of an Elastic Load Balancer to attach to this layer
        /// </summary>
        [Input("elasticLoadBalancer")]
        public Input<string>? ElasticLoadBalancer { get; set; }

        /// <summary>
        /// Whether to install OS and package updates on each instance when it boots.
        /// </summary>
        [Input("installUpdatesOnBoot")]
        public Input<bool>? InstallUpdatesOnBoot { get; set; }

        /// <summary>
        /// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
        /// </summary>
        [Input("instanceShutdownTimeout")]
        public Input<int>? InstanceShutdownTimeout { get; set; }

        /// <summary>
        /// Options to set for the JVM.
        /// </summary>
        [Input("jvmOptions")]
        public Input<string>? JvmOptions { get; set; }

        /// <summary>
        /// Keyword for the type of JVM to use. Defaults to `openjdk`.
        /// </summary>
        [Input("jvmType")]
        public Input<string>? JvmType { get; set; }

        /// <summary>
        /// Version of JVM to use. Defaults to "7".
        /// </summary>
        [Input("jvmVersion")]
        public Input<string>? JvmVersion { get; set; }

        /// <summary>
        /// A human-readable name for the layer.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the stack the layer will belong to.
        /// </summary>
        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        [Input("systemPackages")]
        private InputList<string>? _systemPackages;

        /// <summary>
        /// Names of a set of system packages to install on the layer's instances.
        /// </summary>
        public InputList<string> SystemPackages
        {
            get => _systemPackages ?? (_systemPackages = new InputList<string>());
            set => _systemPackages = value;
        }

        /// <summary>
        /// Whether to use EBS-optimized instances.
        /// </summary>
        [Input("useEbsOptimizedInstances")]
        public Input<bool>? UseEbsOptimizedInstances { get; set; }

        public JavaAppLayerArgs()
        {
        }
    }

    public sealed class JavaAppLayerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Keyword for the application container to use. Defaults to "tomcat".
        /// </summary>
        [Input("appServer")]
        public Input<string>? AppServer { get; set; }

        /// <summary>
        /// Version of the selected application container to use. Defaults to "7".
        /// </summary>
        [Input("appServerVersion")]
        public Input<string>? AppServerVersion { get; set; }

        /// <summary>
        /// Whether to automatically assign an elastic IP address to the layer's instances.
        /// </summary>
        [Input("autoAssignElasticIps")]
        public Input<bool>? AutoAssignElasticIps { get; set; }

        /// <summary>
        /// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
        /// </summary>
        [Input("autoAssignPublicIps")]
        public Input<bool>? AutoAssignPublicIps { get; set; }

        /// <summary>
        /// Whether to enable auto-healing for the layer.
        /// </summary>
        [Input("autoHealing")]
        public Input<bool>? AutoHealing { get; set; }

        [Input("customConfigureRecipes")]
        private InputList<string>? _customConfigureRecipes;
        public InputList<string> CustomConfigureRecipes
        {
            get => _customConfigureRecipes ?? (_customConfigureRecipes = new InputList<string>());
            set => _customConfigureRecipes = value;
        }

        [Input("customDeployRecipes")]
        private InputList<string>? _customDeployRecipes;
        public InputList<string> CustomDeployRecipes
        {
            get => _customDeployRecipes ?? (_customDeployRecipes = new InputList<string>());
            set => _customDeployRecipes = value;
        }

        /// <summary>
        /// The ARN of an IAM profile that will be used for the layer's instances.
        /// </summary>
        [Input("customInstanceProfileArn")]
        public Input<string>? CustomInstanceProfileArn { get; set; }

        /// <summary>
        /// Custom JSON attributes to apply to the layer.
        /// </summary>
        [Input("customJson")]
        public Input<string>? CustomJson { get; set; }

        [Input("customSecurityGroupIds")]
        private InputList<string>? _customSecurityGroupIds;

        /// <summary>
        /// Ids for a set of security groups to apply to the layer's instances.
        /// </summary>
        public InputList<string> CustomSecurityGroupIds
        {
            get => _customSecurityGroupIds ?? (_customSecurityGroupIds = new InputList<string>());
            set => _customSecurityGroupIds = value;
        }

        [Input("customSetupRecipes")]
        private InputList<string>? _customSetupRecipes;
        public InputList<string> CustomSetupRecipes
        {
            get => _customSetupRecipes ?? (_customSetupRecipes = new InputList<string>());
            set => _customSetupRecipes = value;
        }

        [Input("customShutdownRecipes")]
        private InputList<string>? _customShutdownRecipes;
        public InputList<string> CustomShutdownRecipes
        {
            get => _customShutdownRecipes ?? (_customShutdownRecipes = new InputList<string>());
            set => _customShutdownRecipes = value;
        }

        [Input("customUndeployRecipes")]
        private InputList<string>? _customUndeployRecipes;
        public InputList<string> CustomUndeployRecipes
        {
            get => _customUndeployRecipes ?? (_customUndeployRecipes = new InputList<string>());
            set => _customUndeployRecipes = value;
        }

        /// <summary>
        /// Whether to enable Elastic Load Balancing connection draining.
        /// </summary>
        [Input("drainElbOnShutdown")]
        public Input<bool>? DrainElbOnShutdown { get; set; }

        [Input("ebsVolumes")]
        private InputList<Inputs.JavaAppLayerEbsVolumesGetArgs>? _ebsVolumes;

        /// <summary>
        /// `ebs_volume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
        /// </summary>
        public InputList<Inputs.JavaAppLayerEbsVolumesGetArgs> EbsVolumes
        {
            get => _ebsVolumes ?? (_ebsVolumes = new InputList<Inputs.JavaAppLayerEbsVolumesGetArgs>());
            set => _ebsVolumes = value;
        }

        /// <summary>
        /// Name of an Elastic Load Balancer to attach to this layer
        /// </summary>
        [Input("elasticLoadBalancer")]
        public Input<string>? ElasticLoadBalancer { get; set; }

        /// <summary>
        /// Whether to install OS and package updates on each instance when it boots.
        /// </summary>
        [Input("installUpdatesOnBoot")]
        public Input<bool>? InstallUpdatesOnBoot { get; set; }

        /// <summary>
        /// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
        /// </summary>
        [Input("instanceShutdownTimeout")]
        public Input<int>? InstanceShutdownTimeout { get; set; }

        /// <summary>
        /// Options to set for the JVM.
        /// </summary>
        [Input("jvmOptions")]
        public Input<string>? JvmOptions { get; set; }

        /// <summary>
        /// Keyword for the type of JVM to use. Defaults to `openjdk`.
        /// </summary>
        [Input("jvmType")]
        public Input<string>? JvmType { get; set; }

        /// <summary>
        /// Version of JVM to use. Defaults to "7".
        /// </summary>
        [Input("jvmVersion")]
        public Input<string>? JvmVersion { get; set; }

        /// <summary>
        /// A human-readable name for the layer.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the stack the layer will belong to.
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        [Input("systemPackages")]
        private InputList<string>? _systemPackages;

        /// <summary>
        /// Names of a set of system packages to install on the layer's instances.
        /// </summary>
        public InputList<string> SystemPackages
        {
            get => _systemPackages ?? (_systemPackages = new InputList<string>());
            set => _systemPackages = value;
        }

        /// <summary>
        /// Whether to use EBS-optimized instances.
        /// </summary>
        [Input("useEbsOptimizedInstances")]
        public Input<bool>? UseEbsOptimizedInstances { get; set; }

        public JavaAppLayerState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class JavaAppLayerEbsVolumesArgs : Pulumi.ResourceArgs
    {
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        [Input("iops")]
        public Input<int>? Iops { get; set; }

        [Input("mountPoint", required: true)]
        public Input<string> MountPoint { get; set; } = null!;

        [Input("numberOfDisks", required: true)]
        public Input<int> NumberOfDisks { get; set; } = null!;

        [Input("raidLevel")]
        public Input<string>? RaidLevel { get; set; }

        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        [Input("type")]
        public Input<string>? Type { get; set; }

        public JavaAppLayerEbsVolumesArgs()
        {
        }
    }

    public sealed class JavaAppLayerEbsVolumesGetArgs : Pulumi.ResourceArgs
    {
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        [Input("iops")]
        public Input<int>? Iops { get; set; }

        [Input("mountPoint", required: true)]
        public Input<string> MountPoint { get; set; } = null!;

        [Input("numberOfDisks", required: true)]
        public Input<int> NumberOfDisks { get; set; } = null!;

        [Input("raidLevel")]
        public Input<string>? RaidLevel { get; set; }

        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        [Input("type")]
        public Input<string>? Type { get; set; }

        public JavaAppLayerEbsVolumesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class JavaAppLayerEbsVolumes
    {
        public readonly bool? Encrypted;
        public readonly int? Iops;
        public readonly string MountPoint;
        public readonly int NumberOfDisks;
        public readonly string? RaidLevel;
        public readonly int Size;
        public readonly string? Type;

        [OutputConstructor]
        private JavaAppLayerEbsVolumes(
            bool? encrypted,
            int? iops,
            string mountPoint,
            int numberOfDisks,
            string? raidLevel,
            int size,
            string? type)
        {
            Encrypted = encrypted;
            Iops = iops;
            MountPoint = mountPoint;
            NumberOfDisks = numberOfDisks;
            RaidLevel = raidLevel;
            Size = size;
            Type = type;
        }
    }
    }
}
