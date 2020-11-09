// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GameLift
{
    /// <summary>
    /// Provides a Gamelift Fleet resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.GameLift.Fleet("example", new Aws.GameLift.FleetArgs
    ///         {
    ///             BuildId = aws_gamelift_build.Example.Id,
    ///             Ec2InstanceType = "t2.micro",
    ///             FleetType = "ON_DEMAND",
    ///             RuntimeConfiguration = new Aws.GameLift.Inputs.FleetRuntimeConfigurationArgs
    ///             {
    ///                 ServerProcesses = 
    ///                 {
    ///                     new Aws.GameLift.Inputs.FleetRuntimeConfigurationServerProcessArgs
    ///                     {
    ///                         ConcurrentExecutions = 1,
    ///                         LaunchPath = "C:\\game\\GomokuServer.exe",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Gamelift Fleets cannot be imported at this time.
    /// </summary>
    public partial class Fleet : Pulumi.CustomResource
    {
        /// <summary>
        /// Fleet ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// ID of the Gamelift Build to be deployed on the fleet.
        /// </summary>
        [Output("buildId")]
        public Output<string> BuildId { get; private set; } = null!;

        /// <summary>
        /// Human-readable description of the fleet.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
        /// </summary>
        [Output("ec2InboundPermissions")]
        public Output<ImmutableArray<Outputs.FleetEc2InboundPermission>> Ec2InboundPermissions { get; private set; } = null!;

        /// <summary>
        /// Name of an EC2 instance type. e.g. `t2.micro`
        /// </summary>
        [Output("ec2InstanceType")]
        public Output<string> Ec2InstanceType { get; private set; } = null!;

        /// <summary>
        /// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
        /// </summary>
        [Output("fleetType")]
        public Output<string?> FleetType { get; private set; } = null!;

        /// <summary>
        /// ARN of an IAM role that instances in the fleet can assume.
        /// </summary>
        [Output("instanceRoleArn")]
        public Output<string?> InstanceRoleArn { get; private set; } = null!;

        [Output("logPaths")]
        public Output<ImmutableArray<string>> LogPaths { get; private set; } = null!;

        /// <summary>
        /// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
        /// </summary>
        [Output("metricGroups")]
        public Output<ImmutableArray<string>> MetricGroups { get; private set; } = null!;

        /// <summary>
        /// The name of the fleet.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
        /// </summary>
        [Output("newGameSessionProtectionPolicy")]
        public Output<string?> NewGameSessionProtectionPolicy { get; private set; } = null!;

        /// <summary>
        /// Operating system of the fleet's computing resources.
        /// </summary>
        [Output("operatingSystem")]
        public Output<string> OperatingSystem { get; private set; } = null!;

        /// <summary>
        /// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
        /// </summary>
        [Output("resourceCreationLimitPolicy")]
        public Output<Outputs.FleetResourceCreationLimitPolicy?> ResourceCreationLimitPolicy { get; private set; } = null!;

        /// <summary>
        /// Instructions for launching server processes on each instance in the fleet. See below.
        /// </summary>
        [Output("runtimeConfiguration")]
        public Output<Outputs.FleetRuntimeConfiguration?> RuntimeConfiguration { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
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
        /// <summary>
        /// ID of the Gamelift Build to be deployed on the fleet.
        /// </summary>
        [Input("buildId", required: true)]
        public Input<string> BuildId { get; set; } = null!;

        /// <summary>
        /// Human-readable description of the fleet.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ec2InboundPermissions")]
        private InputList<Inputs.FleetEc2InboundPermissionArgs>? _ec2InboundPermissions;

        /// <summary>
        /// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
        /// </summary>
        public InputList<Inputs.FleetEc2InboundPermissionArgs> Ec2InboundPermissions
        {
            get => _ec2InboundPermissions ?? (_ec2InboundPermissions = new InputList<Inputs.FleetEc2InboundPermissionArgs>());
            set => _ec2InboundPermissions = value;
        }

        /// <summary>
        /// Name of an EC2 instance type. e.g. `t2.micro`
        /// </summary>
        [Input("ec2InstanceType", required: true)]
        public Input<string> Ec2InstanceType { get; set; } = null!;

        /// <summary>
        /// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
        /// </summary>
        [Input("fleetType")]
        public Input<string>? FleetType { get; set; }

        /// <summary>
        /// ARN of an IAM role that instances in the fleet can assume.
        /// </summary>
        [Input("instanceRoleArn")]
        public Input<string>? InstanceRoleArn { get; set; }

        [Input("metricGroups")]
        private InputList<string>? _metricGroups;

        /// <summary>
        /// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
        /// </summary>
        public InputList<string> MetricGroups
        {
            get => _metricGroups ?? (_metricGroups = new InputList<string>());
            set => _metricGroups = value;
        }

        /// <summary>
        /// The name of the fleet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
        /// </summary>
        [Input("newGameSessionProtectionPolicy")]
        public Input<string>? NewGameSessionProtectionPolicy { get; set; }

        /// <summary>
        /// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
        /// </summary>
        [Input("resourceCreationLimitPolicy")]
        public Input<Inputs.FleetResourceCreationLimitPolicyArgs>? ResourceCreationLimitPolicy { get; set; }

        /// <summary>
        /// Instructions for launching server processes on each instance in the fleet. See below.
        /// </summary>
        [Input("runtimeConfiguration")]
        public Input<Inputs.FleetRuntimeConfigurationArgs>? RuntimeConfiguration { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
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
        /// <summary>
        /// Fleet ARN.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// ID of the Gamelift Build to be deployed on the fleet.
        /// </summary>
        [Input("buildId")]
        public Input<string>? BuildId { get; set; }

        /// <summary>
        /// Human-readable description of the fleet.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ec2InboundPermissions")]
        private InputList<Inputs.FleetEc2InboundPermissionGetArgs>? _ec2InboundPermissions;

        /// <summary>
        /// Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
        /// </summary>
        public InputList<Inputs.FleetEc2InboundPermissionGetArgs> Ec2InboundPermissions
        {
            get => _ec2InboundPermissions ?? (_ec2InboundPermissions = new InputList<Inputs.FleetEc2InboundPermissionGetArgs>());
            set => _ec2InboundPermissions = value;
        }

        /// <summary>
        /// Name of an EC2 instance type. e.g. `t2.micro`
        /// </summary>
        [Input("ec2InstanceType")]
        public Input<string>? Ec2InstanceType { get; set; }

        /// <summary>
        /// Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
        /// </summary>
        [Input("fleetType")]
        public Input<string>? FleetType { get; set; }

        /// <summary>
        /// ARN of an IAM role that instances in the fleet can assume.
        /// </summary>
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

        /// <summary>
        /// List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
        /// </summary>
        public InputList<string> MetricGroups
        {
            get => _metricGroups ?? (_metricGroups = new InputList<string>());
            set => _metricGroups = value;
        }

        /// <summary>
        /// The name of the fleet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
        /// </summary>
        [Input("newGameSessionProtectionPolicy")]
        public Input<string>? NewGameSessionProtectionPolicy { get; set; }

        /// <summary>
        /// Operating system of the fleet's computing resources.
        /// </summary>
        [Input("operatingSystem")]
        public Input<string>? OperatingSystem { get; set; }

        /// <summary>
        /// Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
        /// </summary>
        [Input("resourceCreationLimitPolicy")]
        public Input<Inputs.FleetResourceCreationLimitPolicyGetArgs>? ResourceCreationLimitPolicy { get; set; }

        /// <summary>
        /// Instructions for launching server processes on each instance in the fleet. See below.
        /// </summary>
        [Input("runtimeConfiguration")]
        public Input<Inputs.FleetRuntimeConfigurationGetArgs>? RuntimeConfiguration { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
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
