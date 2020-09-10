// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.OpsWorks
{
    public partial class RailsAppLayer : Pulumi.CustomResource
    {
        [Output("appServer")]
        public Output<string?> AppServer { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("autoAssignElasticIps")]
        public Output<bool?> AutoAssignElasticIps { get; private set; } = null!;

        [Output("autoAssignPublicIps")]
        public Output<bool?> AutoAssignPublicIps { get; private set; } = null!;

        [Output("autoHealing")]
        public Output<bool?> AutoHealing { get; private set; } = null!;

        [Output("bundlerVersion")]
        public Output<string?> BundlerVersion { get; private set; } = null!;

        [Output("customConfigureRecipes")]
        public Output<ImmutableArray<string>> CustomConfigureRecipes { get; private set; } = null!;

        [Output("customDeployRecipes")]
        public Output<ImmutableArray<string>> CustomDeployRecipes { get; private set; } = null!;

        [Output("customInstanceProfileArn")]
        public Output<string?> CustomInstanceProfileArn { get; private set; } = null!;

        [Output("customJson")]
        public Output<string?> CustomJson { get; private set; } = null!;

        [Output("customSecurityGroupIds")]
        public Output<ImmutableArray<string>> CustomSecurityGroupIds { get; private set; } = null!;

        [Output("customSetupRecipes")]
        public Output<ImmutableArray<string>> CustomSetupRecipes { get; private set; } = null!;

        [Output("customShutdownRecipes")]
        public Output<ImmutableArray<string>> CustomShutdownRecipes { get; private set; } = null!;

        [Output("customUndeployRecipes")]
        public Output<ImmutableArray<string>> CustomUndeployRecipes { get; private set; } = null!;

        [Output("drainElbOnShutdown")]
        public Output<bool?> DrainElbOnShutdown { get; private set; } = null!;

        [Output("ebsVolumes")]
        public Output<ImmutableArray<Outputs.RailsAppLayerEbsVolume>> EbsVolumes { get; private set; } = null!;

        [Output("elasticLoadBalancer")]
        public Output<string?> ElasticLoadBalancer { get; private set; } = null!;

        [Output("installUpdatesOnBoot")]
        public Output<bool?> InstallUpdatesOnBoot { get; private set; } = null!;

        [Output("instanceShutdownTimeout")]
        public Output<int?> InstanceShutdownTimeout { get; private set; } = null!;

        [Output("manageBundler")]
        public Output<bool?> ManageBundler { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("passengerVersion")]
        public Output<string?> PassengerVersion { get; private set; } = null!;

        [Output("rubyVersion")]
        public Output<string?> RubyVersion { get; private set; } = null!;

        [Output("rubygemsVersion")]
        public Output<string?> RubygemsVersion { get; private set; } = null!;

        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;

        [Output("systemPackages")]
        public Output<ImmutableArray<string>> SystemPackages { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("useEbsOptimizedInstances")]
        public Output<bool?> UseEbsOptimizedInstances { get; private set; } = null!;


        /// <summary>
        /// Create a RailsAppLayer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RailsAppLayer(string name, RailsAppLayerArgs args, CustomResourceOptions? options = null)
            : base("aws:opsworks/railsAppLayer:RailsAppLayer", name, args ?? new RailsAppLayerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RailsAppLayer(string name, Input<string> id, RailsAppLayerState? state = null, CustomResourceOptions? options = null)
            : base("aws:opsworks/railsAppLayer:RailsAppLayer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RailsAppLayer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RailsAppLayer Get(string name, Input<string> id, RailsAppLayerState? state = null, CustomResourceOptions? options = null)
        {
            return new RailsAppLayer(name, id, state, options);
        }
    }

    public sealed class RailsAppLayerArgs : Pulumi.ResourceArgs
    {
        [Input("appServer")]
        public Input<string>? AppServer { get; set; }

        [Input("autoAssignElasticIps")]
        public Input<bool>? AutoAssignElasticIps { get; set; }

        [Input("autoAssignPublicIps")]
        public Input<bool>? AutoAssignPublicIps { get; set; }

        [Input("autoHealing")]
        public Input<bool>? AutoHealing { get; set; }

        [Input("bundlerVersion")]
        public Input<string>? BundlerVersion { get; set; }

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

        [Input("customInstanceProfileArn")]
        public Input<string>? CustomInstanceProfileArn { get; set; }

        [Input("customJson")]
        public Input<string>? CustomJson { get; set; }

        [Input("customSecurityGroupIds")]
        private InputList<string>? _customSecurityGroupIds;
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

        [Input("drainElbOnShutdown")]
        public Input<bool>? DrainElbOnShutdown { get; set; }

        [Input("ebsVolumes")]
        private InputList<Inputs.RailsAppLayerEbsVolumeArgs>? _ebsVolumes;
        public InputList<Inputs.RailsAppLayerEbsVolumeArgs> EbsVolumes
        {
            get => _ebsVolumes ?? (_ebsVolumes = new InputList<Inputs.RailsAppLayerEbsVolumeArgs>());
            set => _ebsVolumes = value;
        }

        [Input("elasticLoadBalancer")]
        public Input<string>? ElasticLoadBalancer { get; set; }

        [Input("installUpdatesOnBoot")]
        public Input<bool>? InstallUpdatesOnBoot { get; set; }

        [Input("instanceShutdownTimeout")]
        public Input<int>? InstanceShutdownTimeout { get; set; }

        [Input("manageBundler")]
        public Input<bool>? ManageBundler { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("passengerVersion")]
        public Input<string>? PassengerVersion { get; set; }

        [Input("rubyVersion")]
        public Input<string>? RubyVersion { get; set; }

        [Input("rubygemsVersion")]
        public Input<string>? RubygemsVersion { get; set; }

        [Input("stackId", required: true)]
        public Input<string> StackId { get; set; } = null!;

        [Input("systemPackages")]
        private InputList<string>? _systemPackages;
        public InputList<string> SystemPackages
        {
            get => _systemPackages ?? (_systemPackages = new InputList<string>());
            set => _systemPackages = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("useEbsOptimizedInstances")]
        public Input<bool>? UseEbsOptimizedInstances { get; set; }

        public RailsAppLayerArgs()
        {
        }
    }

    public sealed class RailsAppLayerState : Pulumi.ResourceArgs
    {
        [Input("appServer")]
        public Input<string>? AppServer { get; set; }

        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("autoAssignElasticIps")]
        public Input<bool>? AutoAssignElasticIps { get; set; }

        [Input("autoAssignPublicIps")]
        public Input<bool>? AutoAssignPublicIps { get; set; }

        [Input("autoHealing")]
        public Input<bool>? AutoHealing { get; set; }

        [Input("bundlerVersion")]
        public Input<string>? BundlerVersion { get; set; }

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

        [Input("customInstanceProfileArn")]
        public Input<string>? CustomInstanceProfileArn { get; set; }

        [Input("customJson")]
        public Input<string>? CustomJson { get; set; }

        [Input("customSecurityGroupIds")]
        private InputList<string>? _customSecurityGroupIds;
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

        [Input("drainElbOnShutdown")]
        public Input<bool>? DrainElbOnShutdown { get; set; }

        [Input("ebsVolumes")]
        private InputList<Inputs.RailsAppLayerEbsVolumeGetArgs>? _ebsVolumes;
        public InputList<Inputs.RailsAppLayerEbsVolumeGetArgs> EbsVolumes
        {
            get => _ebsVolumes ?? (_ebsVolumes = new InputList<Inputs.RailsAppLayerEbsVolumeGetArgs>());
            set => _ebsVolumes = value;
        }

        [Input("elasticLoadBalancer")]
        public Input<string>? ElasticLoadBalancer { get; set; }

        [Input("installUpdatesOnBoot")]
        public Input<bool>? InstallUpdatesOnBoot { get; set; }

        [Input("instanceShutdownTimeout")]
        public Input<int>? InstanceShutdownTimeout { get; set; }

        [Input("manageBundler")]
        public Input<bool>? ManageBundler { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("passengerVersion")]
        public Input<string>? PassengerVersion { get; set; }

        [Input("rubyVersion")]
        public Input<string>? RubyVersion { get; set; }

        [Input("rubygemsVersion")]
        public Input<string>? RubygemsVersion { get; set; }

        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        [Input("systemPackages")]
        private InputList<string>? _systemPackages;
        public InputList<string> SystemPackages
        {
            get => _systemPackages ?? (_systemPackages = new InputList<string>());
            set => _systemPackages = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("useEbsOptimizedInstances")]
        public Input<bool>? UseEbsOptimizedInstances { get; set; }

        public RailsAppLayerState()
        {
        }
    }
}
