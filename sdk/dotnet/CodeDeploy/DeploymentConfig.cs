// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeDeploy
{
    public partial class DeploymentConfig : Pulumi.CustomResource
    {
        [Output("computePlatform")]
        public Output<string?> ComputePlatform { get; private set; } = null!;

        [Output("deploymentConfigId")]
        public Output<string> DeploymentConfigId { get; private set; } = null!;

        [Output("deploymentConfigName")]
        public Output<string> DeploymentConfigName { get; private set; } = null!;

        [Output("minimumHealthyHosts")]
        public Output<Outputs.DeploymentConfigMinimumHealthyHosts?> MinimumHealthyHosts { get; private set; } = null!;

        [Output("trafficRoutingConfig")]
        public Output<Outputs.DeploymentConfigTrafficRoutingConfig?> TrafficRoutingConfig { get; private set; } = null!;


        /// <summary>
        /// Create a DeploymentConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeploymentConfig(string name, DeploymentConfigArgs args, CustomResourceOptions? options = null)
            : base("aws:codedeploy/deploymentConfig:DeploymentConfig", name, args ?? new DeploymentConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeploymentConfig(string name, Input<string> id, DeploymentConfigState? state = null, CustomResourceOptions? options = null)
            : base("aws:codedeploy/deploymentConfig:DeploymentConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeploymentConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeploymentConfig Get(string name, Input<string> id, DeploymentConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new DeploymentConfig(name, id, state, options);
        }
    }

    public sealed class DeploymentConfigArgs : Pulumi.ResourceArgs
    {
        [Input("computePlatform")]
        public Input<string>? ComputePlatform { get; set; }

        [Input("deploymentConfigName", required: true)]
        public Input<string> DeploymentConfigName { get; set; } = null!;

        [Input("minimumHealthyHosts")]
        public Input<Inputs.DeploymentConfigMinimumHealthyHostsArgs>? MinimumHealthyHosts { get; set; }

        [Input("trafficRoutingConfig")]
        public Input<Inputs.DeploymentConfigTrafficRoutingConfigArgs>? TrafficRoutingConfig { get; set; }

        public DeploymentConfigArgs()
        {
        }
    }

    public sealed class DeploymentConfigState : Pulumi.ResourceArgs
    {
        [Input("computePlatform")]
        public Input<string>? ComputePlatform { get; set; }

        [Input("deploymentConfigId")]
        public Input<string>? DeploymentConfigId { get; set; }

        [Input("deploymentConfigName")]
        public Input<string>? DeploymentConfigName { get; set; }

        [Input("minimumHealthyHosts")]
        public Input<Inputs.DeploymentConfigMinimumHealthyHostsGetArgs>? MinimumHealthyHosts { get; set; }

        [Input("trafficRoutingConfig")]
        public Input<Inputs.DeploymentConfigTrafficRoutingConfigGetArgs>? TrafficRoutingConfig { get; set; }

        public DeploymentConfigState()
        {
        }
    }
}
