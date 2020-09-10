// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecs
{
    public partial class Service : Pulumi.CustomResource
    {
        [Output("capacityProviderStrategies")]
        public Output<ImmutableArray<Outputs.ServiceCapacityProviderStrategy>> CapacityProviderStrategies { get; private set; } = null!;

        [Output("cluster")]
        public Output<string> Cluster { get; private set; } = null!;

        [Output("deploymentController")]
        public Output<Outputs.ServiceDeploymentController?> DeploymentController { get; private set; } = null!;

        [Output("deploymentMaximumPercent")]
        public Output<int?> DeploymentMaximumPercent { get; private set; } = null!;

        [Output("deploymentMinimumHealthyPercent")]
        public Output<int?> DeploymentMinimumHealthyPercent { get; private set; } = null!;

        [Output("desiredCount")]
        public Output<int?> DesiredCount { get; private set; } = null!;

        [Output("enableEcsManagedTags")]
        public Output<bool?> EnableEcsManagedTags { get; private set; } = null!;

        [Output("forceNewDeployment")]
        public Output<bool?> ForceNewDeployment { get; private set; } = null!;

        [Output("healthCheckGracePeriodSeconds")]
        public Output<int?> HealthCheckGracePeriodSeconds { get; private set; } = null!;

        [Output("iamRole")]
        public Output<string> IamRole { get; private set; } = null!;

        [Output("launchType")]
        public Output<string> LaunchType { get; private set; } = null!;

        [Output("loadBalancers")]
        public Output<ImmutableArray<Outputs.ServiceLoadBalancer>> LoadBalancers { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("networkConfiguration")]
        public Output<Outputs.ServiceNetworkConfiguration?> NetworkConfiguration { get; private set; } = null!;

        [Output("orderedPlacementStrategies")]
        public Output<ImmutableArray<Outputs.ServiceOrderedPlacementStrategy>> OrderedPlacementStrategies { get; private set; } = null!;

        [Output("placementConstraints")]
        public Output<ImmutableArray<Outputs.ServicePlacementConstraint>> PlacementConstraints { get; private set; } = null!;

        [Output("platformVersion")]
        public Output<string> PlatformVersion { get; private set; } = null!;

        [Output("propagateTags")]
        public Output<string?> PropagateTags { get; private set; } = null!;

        [Output("schedulingStrategy")]
        public Output<string?> SchedulingStrategy { get; private set; } = null!;

        [Output("serviceRegistries")]
        public Output<Outputs.ServiceServiceRegistries?> ServiceRegistries { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("taskDefinition")]
        public Output<string?> TaskDefinition { get; private set; } = null!;

        [Output("waitForSteadyState")]
        public Output<bool?> WaitForSteadyState { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:ecs/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("aws:ecs/service:Service", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new Service(name, id, state, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        [Input("capacityProviderStrategies")]
        private InputList<Inputs.ServiceCapacityProviderStrategyArgs>? _capacityProviderStrategies;
        public InputList<Inputs.ServiceCapacityProviderStrategyArgs> CapacityProviderStrategies
        {
            get => _capacityProviderStrategies ?? (_capacityProviderStrategies = new InputList<Inputs.ServiceCapacityProviderStrategyArgs>());
            set => _capacityProviderStrategies = value;
        }

        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        [Input("deploymentController")]
        public Input<Inputs.ServiceDeploymentControllerArgs>? DeploymentController { get; set; }

        [Input("deploymentMaximumPercent")]
        public Input<int>? DeploymentMaximumPercent { get; set; }

        [Input("deploymentMinimumHealthyPercent")]
        public Input<int>? DeploymentMinimumHealthyPercent { get; set; }

        [Input("desiredCount")]
        public Input<int>? DesiredCount { get; set; }

        [Input("enableEcsManagedTags")]
        public Input<bool>? EnableEcsManagedTags { get; set; }

        [Input("forceNewDeployment")]
        public Input<bool>? ForceNewDeployment { get; set; }

        [Input("healthCheckGracePeriodSeconds")]
        public Input<int>? HealthCheckGracePeriodSeconds { get; set; }

        [Input("iamRole")]
        public Input<string>? IamRole { get; set; }

        [Input("launchType")]
        public Input<string>? LaunchType { get; set; }

        [Input("loadBalancers")]
        private InputList<Inputs.ServiceLoadBalancerArgs>? _loadBalancers;
        public InputList<Inputs.ServiceLoadBalancerArgs> LoadBalancers
        {
            get => _loadBalancers ?? (_loadBalancers = new InputList<Inputs.ServiceLoadBalancerArgs>());
            set => _loadBalancers = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkConfiguration")]
        public Input<Inputs.ServiceNetworkConfigurationArgs>? NetworkConfiguration { get; set; }

        [Input("orderedPlacementStrategies")]
        private InputList<Inputs.ServiceOrderedPlacementStrategyArgs>? _orderedPlacementStrategies;
        public InputList<Inputs.ServiceOrderedPlacementStrategyArgs> OrderedPlacementStrategies
        {
            get => _orderedPlacementStrategies ?? (_orderedPlacementStrategies = new InputList<Inputs.ServiceOrderedPlacementStrategyArgs>());
            set => _orderedPlacementStrategies = value;
        }

        [Input("placementConstraints")]
        private InputList<Inputs.ServicePlacementConstraintArgs>? _placementConstraints;
        public InputList<Inputs.ServicePlacementConstraintArgs> PlacementConstraints
        {
            get => _placementConstraints ?? (_placementConstraints = new InputList<Inputs.ServicePlacementConstraintArgs>());
            set => _placementConstraints = value;
        }

        [Input("platformVersion")]
        public Input<string>? PlatformVersion { get; set; }

        [Input("propagateTags")]
        public Input<string>? PropagateTags { get; set; }

        [Input("schedulingStrategy")]
        public Input<string>? SchedulingStrategy { get; set; }

        [Input("serviceRegistries")]
        public Input<Inputs.ServiceServiceRegistriesArgs>? ServiceRegistries { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("taskDefinition")]
        public Input<string>? TaskDefinition { get; set; }

        [Input("waitForSteadyState")]
        public Input<bool>? WaitForSteadyState { get; set; }

        public ServiceArgs()
        {
        }
    }

    public sealed class ServiceState : Pulumi.ResourceArgs
    {
        [Input("capacityProviderStrategies")]
        private InputList<Inputs.ServiceCapacityProviderStrategyGetArgs>? _capacityProviderStrategies;
        public InputList<Inputs.ServiceCapacityProviderStrategyGetArgs> CapacityProviderStrategies
        {
            get => _capacityProviderStrategies ?? (_capacityProviderStrategies = new InputList<Inputs.ServiceCapacityProviderStrategyGetArgs>());
            set => _capacityProviderStrategies = value;
        }

        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        [Input("deploymentController")]
        public Input<Inputs.ServiceDeploymentControllerGetArgs>? DeploymentController { get; set; }

        [Input("deploymentMaximumPercent")]
        public Input<int>? DeploymentMaximumPercent { get; set; }

        [Input("deploymentMinimumHealthyPercent")]
        public Input<int>? DeploymentMinimumHealthyPercent { get; set; }

        [Input("desiredCount")]
        public Input<int>? DesiredCount { get; set; }

        [Input("enableEcsManagedTags")]
        public Input<bool>? EnableEcsManagedTags { get; set; }

        [Input("forceNewDeployment")]
        public Input<bool>? ForceNewDeployment { get; set; }

        [Input("healthCheckGracePeriodSeconds")]
        public Input<int>? HealthCheckGracePeriodSeconds { get; set; }

        [Input("iamRole")]
        public Input<string>? IamRole { get; set; }

        [Input("launchType")]
        public Input<string>? LaunchType { get; set; }

        [Input("loadBalancers")]
        private InputList<Inputs.ServiceLoadBalancerGetArgs>? _loadBalancers;
        public InputList<Inputs.ServiceLoadBalancerGetArgs> LoadBalancers
        {
            get => _loadBalancers ?? (_loadBalancers = new InputList<Inputs.ServiceLoadBalancerGetArgs>());
            set => _loadBalancers = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkConfiguration")]
        public Input<Inputs.ServiceNetworkConfigurationGetArgs>? NetworkConfiguration { get; set; }

        [Input("orderedPlacementStrategies")]
        private InputList<Inputs.ServiceOrderedPlacementStrategyGetArgs>? _orderedPlacementStrategies;
        public InputList<Inputs.ServiceOrderedPlacementStrategyGetArgs> OrderedPlacementStrategies
        {
            get => _orderedPlacementStrategies ?? (_orderedPlacementStrategies = new InputList<Inputs.ServiceOrderedPlacementStrategyGetArgs>());
            set => _orderedPlacementStrategies = value;
        }

        [Input("placementConstraints")]
        private InputList<Inputs.ServicePlacementConstraintGetArgs>? _placementConstraints;
        public InputList<Inputs.ServicePlacementConstraintGetArgs> PlacementConstraints
        {
            get => _placementConstraints ?? (_placementConstraints = new InputList<Inputs.ServicePlacementConstraintGetArgs>());
            set => _placementConstraints = value;
        }

        [Input("platformVersion")]
        public Input<string>? PlatformVersion { get; set; }

        [Input("propagateTags")]
        public Input<string>? PropagateTags { get; set; }

        [Input("schedulingStrategy")]
        public Input<string>? SchedulingStrategy { get; set; }

        [Input("serviceRegistries")]
        public Input<Inputs.ServiceServiceRegistriesGetArgs>? ServiceRegistries { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("taskDefinition")]
        public Input<string>? TaskDefinition { get; set; }

        [Input("waitForSteadyState")]
        public Input<bool>? WaitForSteadyState { get; set; }

        public ServiceState()
        {
        }
    }
}
