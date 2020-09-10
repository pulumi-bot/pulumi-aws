// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GameLift
{
    public partial class Fleet : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("buildId")]
        public Output<string> BuildId { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("ec2InboundPermissions")]
        public Output<ImmutableArray<Outputs.FleetEc2InboundPermission>> Ec2InboundPermissions { get; private set; } = null!;

        [Output("ec2InstanceType")]
        public Output<string> Ec2InstanceType { get; private set; } = null!;

        [Output("fleetType")]
        public Output<string?> FleetType { get; private set; } = null!;

        [Output("instanceRoleArn")]
        public Output<string?> InstanceRoleArn { get; private set; } = null!;

        [Output("logPaths")]
        public Output<ImmutableArray<string>> LogPaths { get; private set; } = null!;

        [Output("metricGroups")]
        public Output<ImmutableArray<string>> MetricGroups { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("newGameSessionProtectionPolicy")]
        public Output<string?> NewGameSessionProtectionPolicy { get; private set; } = null!;

        [Output("operatingSystem")]
        public Output<string> OperatingSystem { get; private set; } = null!;

        [Output("resourceCreationLimitPolicy")]
        public Output<Outputs.FleetResourceCreationLimitPolicy?> ResourceCreationLimitPolicy { get; private set; } = null!;

        [Output("runtimeConfiguration")]
        public Output<Outputs.FleetRuntimeConfiguration?> RuntimeConfiguration { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Fleet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fleet(string name, FleetArgs args, CustomResourceOptions? options = null)
            : base("aws:gamelift/fleet:Fleet", name, args ?? new FleetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fleet(string name, Input<string> id, FleetState? state = null, CustomResourceOptions? options = null)
            : base("aws:gamelift/fleet:Fleet", name, state, MakeResourceOptions(options, id))
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
        [Input("buildId", required: true)]
        public Input<string> BuildId { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ec2InboundPermissions")]
        private InputList<Inputs.FleetEc2InboundPermissionArgs>? _ec2InboundPermissions;
        public InputList<Inputs.FleetEc2InboundPermissionArgs> Ec2InboundPermissions
        {
            get => _ec2InboundPermissions ?? (_ec2InboundPermissions = new InputList<Inputs.FleetEc2InboundPermissionArgs>());
            set => _ec2InboundPermissions = value;
        }

        [Input("ec2InstanceType", required: true)]
        public Input<string> Ec2InstanceType { get; set; } = null!;

        [Input("fleetType")]
        public Input<string>? FleetType { get; set; }

        [Input("instanceRoleArn")]
        public Input<string>? InstanceRoleArn { get; set; }

        [Input("metricGroups")]
        private InputList<string>? _metricGroups;
        public InputList<string> MetricGroups
        {
            get => _metricGroups ?? (_metricGroups = new InputList<string>());
            set => _metricGroups = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("newGameSessionProtectionPolicy")]
        public Input<string>? NewGameSessionProtectionPolicy { get; set; }

        [Input("resourceCreationLimitPolicy")]
        public Input<Inputs.FleetResourceCreationLimitPolicyArgs>? ResourceCreationLimitPolicy { get; set; }

        [Input("runtimeConfiguration")]
        public Input<Inputs.FleetRuntimeConfigurationArgs>? RuntimeConfiguration { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public FleetArgs()
        {
        }
    }

    public sealed class FleetState : Pulumi.ResourceArgs
    {
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("buildId")]
        public Input<string>? BuildId { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ec2InboundPermissions")]
        private InputList<Inputs.FleetEc2InboundPermissionGetArgs>? _ec2InboundPermissions;
        public InputList<Inputs.FleetEc2InboundPermissionGetArgs> Ec2InboundPermissions
        {
            get => _ec2InboundPermissions ?? (_ec2InboundPermissions = new InputList<Inputs.FleetEc2InboundPermissionGetArgs>());
            set => _ec2InboundPermissions = value;
        }

        [Input("ec2InstanceType")]
        public Input<string>? Ec2InstanceType { get; set; }

        [Input("fleetType")]
        public Input<string>? FleetType { get; set; }

        [Input("instanceRoleArn")]
        public Input<string>? InstanceRoleArn { get; set; }

        [Input("logPaths")]
        private InputList<string>? _logPaths;
        public InputList<string> LogPaths
        {
            get => _logPaths ?? (_logPaths = new InputList<string>());
            set => _logPaths = value;
        }

        [Input("metricGroups")]
        private InputList<string>? _metricGroups;
        public InputList<string> MetricGroups
        {
            get => _metricGroups ?? (_metricGroups = new InputList<string>());
            set => _metricGroups = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("newGameSessionProtectionPolicy")]
        public Input<string>? NewGameSessionProtectionPolicy { get; set; }

        [Input("operatingSystem")]
        public Input<string>? OperatingSystem { get; set; }

        [Input("resourceCreationLimitPolicy")]
        public Input<Inputs.FleetResourceCreationLimitPolicyGetArgs>? ResourceCreationLimitPolicy { get; set; }

        [Input("runtimeConfiguration")]
        public Input<Inputs.FleetRuntimeConfigurationGetArgs>? RuntimeConfiguration { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public FleetState()
        {
        }
    }
}
