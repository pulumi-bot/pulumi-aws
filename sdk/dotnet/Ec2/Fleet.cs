// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Provides a resource to manage EC2 Fleets.
    /// </summary>
    public partial class Fleet : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`.
        /// </summary>
        [Output("excessCapacityTerminationPolicy")]
        public Output<string?> ExcessCapacityTerminationPolicy { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing EC2 Launch Template configurations. Defined below.
        /// </summary>
        [Output("launchTemplateConfig")]
        public Output<Outputs.FleetLaunchTemplateConfig> LaunchTemplateConfig { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing On-Demand configurations. Defined below.
        /// </summary>
        [Output("onDemandOptions")]
        public Output<Outputs.FleetOnDemandOptions?> OnDemandOptions { get; private set; } = null!;

        /// <summary>
        /// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`.
        /// </summary>
        [Output("replaceUnhealthyInstances")]
        public Output<bool?> ReplaceUnhealthyInstances { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing Spot configurations. Defined below.
        /// </summary>
        [Output("spotOptions")]
        public Output<Outputs.FleetSpotOptions?> SpotOptions { get; private set; } = null!;

        /// <summary>
        /// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing target capacity configurations. Defined below.
        /// </summary>
        [Output("targetCapacitySpecification")]
        public Output<Outputs.FleetTargetCapacitySpecification> TargetCapacitySpecification { get; private set; } = null!;

        /// <summary>
        /// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
        /// </summary>
        [Output("terminateInstances")]
        public Output<bool?> TerminateInstances { get; private set; } = null!;

        /// <summary>
        /// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
        /// </summary>
        [Output("terminateInstancesWithExpiration")]
        public Output<bool?> TerminateInstancesWithExpiration { get; private set; } = null!;

        /// <summary>
        /// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`. Defaults to `maintain`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Fleet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fleet(string name, FleetArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/fleet:Fleet", name, args ?? new FleetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fleet(string name, Input<string> id, FleetState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/fleet:Fleet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Fleet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fleet Get(string name, Input<string> id, FleetState? state = null, CustomResourceOptions? options = null)
        {
            return new Fleet(name, id, state, options);
        }
    }

    public sealed class FleetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`.
        /// </summary>
        [Input("excessCapacityTerminationPolicy")]
        public Input<string>? ExcessCapacityTerminationPolicy { get; set; }

        /// <summary>
        /// Nested argument containing EC2 Launch Template configurations. Defined below.
        /// </summary>
        [Input("launchTemplateConfig", required: true)]
        public Input<Inputs.FleetLaunchTemplateConfigArgs> LaunchTemplateConfig { get; set; } = null!;

        /// <summary>
        /// Nested argument containing On-Demand configurations. Defined below.
        /// </summary>
        [Input("onDemandOptions")]
        public Input<Inputs.FleetOnDemandOptionsArgs>? OnDemandOptions { get; set; }

        /// <summary>
        /// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`.
        /// </summary>
        [Input("replaceUnhealthyInstances")]
        public Input<bool>? ReplaceUnhealthyInstances { get; set; }

        /// <summary>
        /// Nested argument containing Spot configurations. Defined below.
        /// </summary>
        [Input("spotOptions")]
        public Input<Inputs.FleetSpotOptionsArgs>? SpotOptions { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Nested argument containing target capacity configurations. Defined below.
        /// </summary>
        [Input("targetCapacitySpecification", required: true)]
        public Input<Inputs.FleetTargetCapacitySpecificationArgs> TargetCapacitySpecification { get; set; } = null!;

        /// <summary>
        /// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
        /// </summary>
        [Input("terminateInstances")]
        public Input<bool>? TerminateInstances { get; set; }

        /// <summary>
        /// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
        /// </summary>
        [Input("terminateInstancesWithExpiration")]
        public Input<bool>? TerminateInstancesWithExpiration { get; set; }

        /// <summary>
        /// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`. Defaults to `maintain`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FleetArgs()
        {
        }
    }

    public sealed class FleetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2. Valid values: `no-termination`, `termination`. Defaults to `termination`.
        /// </summary>
        [Input("excessCapacityTerminationPolicy")]
        public Input<string>? ExcessCapacityTerminationPolicy { get; set; }

        /// <summary>
        /// Nested argument containing EC2 Launch Template configurations. Defined below.
        /// </summary>
        [Input("launchTemplateConfig")]
        public Input<Inputs.FleetLaunchTemplateConfigGetArgs>? LaunchTemplateConfig { get; set; }

        /// <summary>
        /// Nested argument containing On-Demand configurations. Defined below.
        /// </summary>
        [Input("onDemandOptions")]
        public Input<Inputs.FleetOnDemandOptionsGetArgs>? OnDemandOptions { get; set; }

        /// <summary>
        /// Whether EC2 Fleet should replace unhealthy instances. Defaults to `false`.
        /// </summary>
        [Input("replaceUnhealthyInstances")]
        public Input<bool>? ReplaceUnhealthyInstances { get; set; }

        /// <summary>
        /// Nested argument containing Spot configurations. Defined below.
        /// </summary>
        [Input("spotOptions")]
        public Input<Inputs.FleetSpotOptionsGetArgs>? SpotOptions { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of Fleet tags. To tag instances at launch, specify the tags in the Launch Template.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Nested argument containing target capacity configurations. Defined below.
        /// </summary>
        [Input("targetCapacitySpecification")]
        public Input<Inputs.FleetTargetCapacitySpecificationGetArgs>? TargetCapacitySpecification { get; set; }

        /// <summary>
        /// Whether to terminate instances for an EC2 Fleet if it is deleted successfully. Defaults to `false`.
        /// </summary>
        [Input("terminateInstances")]
        public Input<bool>? TerminateInstances { get; set; }

        /// <summary>
        /// Whether running instances should be terminated when the EC2 Fleet expires. Defaults to `false`.
        /// </summary>
        [Input("terminateInstancesWithExpiration")]
        public Input<bool>? TerminateInstancesWithExpiration { get; set; }

        /// <summary>
        /// The type of request. Indicates whether the EC2 Fleet only requests the target capacity, or also attempts to maintain it. Valid values: `maintain`, `request`. Defaults to `maintain`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FleetState()
        {
        }
    }
}
