// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeDeploy
{
    public partial class DeploymentGroup : Pulumi.CustomResource
    {
        [Output("alarmConfiguration")]
        public Output<Outputs.DeploymentGroupAlarmConfiguration?> AlarmConfiguration { get; private set; } = null!;

        [Output("appName")]
        public Output<string> AppName { get; private set; } = null!;

        [Output("autoRollbackConfiguration")]
        public Output<Outputs.DeploymentGroupAutoRollbackConfiguration?> AutoRollbackConfiguration { get; private set; } = null!;

        [Output("autoscalingGroups")]
        public Output<ImmutableArray<string>> AutoscalingGroups { get; private set; } = null!;

        [Output("blueGreenDeploymentConfig")]
        public Output<Outputs.DeploymentGroupBlueGreenDeploymentConfig> BlueGreenDeploymentConfig { get; private set; } = null!;

        [Output("deploymentConfigName")]
        public Output<string?> DeploymentConfigName { get; private set; } = null!;

        [Output("deploymentGroupName")]
        public Output<string> DeploymentGroupName { get; private set; } = null!;

        [Output("deploymentStyle")]
        public Output<Outputs.DeploymentGroupDeploymentStyle?> DeploymentStyle { get; private set; } = null!;

        [Output("ec2TagFilters")]
        public Output<ImmutableArray<Outputs.DeploymentGroupEc2TagFilter>> Ec2TagFilters { get; private set; } = null!;

        [Output("ec2TagSets")]
        public Output<ImmutableArray<Outputs.DeploymentGroupEc2TagSet>> Ec2TagSets { get; private set; } = null!;

        [Output("ecsService")]
        public Output<Outputs.DeploymentGroupEcsService?> EcsService { get; private set; } = null!;

        [Output("loadBalancerInfo")]
        public Output<Outputs.DeploymentGroupLoadBalancerInfo?> LoadBalancerInfo { get; private set; } = null!;

        [Output("onPremisesInstanceTagFilters")]
        public Output<ImmutableArray<Outputs.DeploymentGroupOnPremisesInstanceTagFilter>> OnPremisesInstanceTagFilters { get; private set; } = null!;

        [Output("serviceRoleArn")]
        public Output<string> ServiceRoleArn { get; private set; } = null!;

        [Output("triggerConfigurations")]
        public Output<ImmutableArray<Outputs.DeploymentGroupTriggerConfiguration>> TriggerConfigurations { get; private set; } = null!;


        /// <summary>
        /// Create a DeploymentGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeploymentGroup(string name, DeploymentGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:codedeploy/deploymentGroup:DeploymentGroup", name, args ?? new DeploymentGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeploymentGroup(string name, Input<string> id, DeploymentGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:codedeploy/deploymentGroup:DeploymentGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeploymentGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeploymentGroup Get(string name, Input<string> id, DeploymentGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new DeploymentGroup(name, id, state, options);
        }
    }

    public sealed class DeploymentGroupArgs : Pulumi.ResourceArgs
    {
        [Input("alarmConfiguration")]
        public Input<Inputs.DeploymentGroupAlarmConfigurationArgs>? AlarmConfiguration { get; set; }

        [Input("appName", required: true)]
        public Input<string> AppName { get; set; } = null!;

        [Input("autoRollbackConfiguration")]
        public Input<Inputs.DeploymentGroupAutoRollbackConfigurationArgs>? AutoRollbackConfiguration { get; set; }

        [Input("autoscalingGroups")]
        private InputList<string>? _autoscalingGroups;
        public InputList<string> AutoscalingGroups
        {
            get => _autoscalingGroups ?? (_autoscalingGroups = new InputList<string>());
            set => _autoscalingGroups = value;
        }

        [Input("blueGreenDeploymentConfig")]
        public Input<Inputs.DeploymentGroupBlueGreenDeploymentConfigArgs>? BlueGreenDeploymentConfig { get; set; }

        [Input("deploymentConfigName")]
        public Input<string>? DeploymentConfigName { get; set; }

        [Input("deploymentGroupName", required: true)]
        public Input<string> DeploymentGroupName { get; set; } = null!;

        [Input("deploymentStyle")]
        public Input<Inputs.DeploymentGroupDeploymentStyleArgs>? DeploymentStyle { get; set; }

        [Input("ec2TagFilters")]
        private InputList<Inputs.DeploymentGroupEc2TagFilterArgs>? _ec2TagFilters;
        public InputList<Inputs.DeploymentGroupEc2TagFilterArgs> Ec2TagFilters
        {
            get => _ec2TagFilters ?? (_ec2TagFilters = new InputList<Inputs.DeploymentGroupEc2TagFilterArgs>());
            set => _ec2TagFilters = value;
        }

        [Input("ec2TagSets")]
        private InputList<Inputs.DeploymentGroupEc2TagSetArgs>? _ec2TagSets;
        public InputList<Inputs.DeploymentGroupEc2TagSetArgs> Ec2TagSets
        {
            get => _ec2TagSets ?? (_ec2TagSets = new InputList<Inputs.DeploymentGroupEc2TagSetArgs>());
            set => _ec2TagSets = value;
        }

        [Input("ecsService")]
        public Input<Inputs.DeploymentGroupEcsServiceArgs>? EcsService { get; set; }

        [Input("loadBalancerInfo")]
        public Input<Inputs.DeploymentGroupLoadBalancerInfoArgs>? LoadBalancerInfo { get; set; }

        [Input("onPremisesInstanceTagFilters")]
        private InputList<Inputs.DeploymentGroupOnPremisesInstanceTagFilterArgs>? _onPremisesInstanceTagFilters;
        public InputList<Inputs.DeploymentGroupOnPremisesInstanceTagFilterArgs> OnPremisesInstanceTagFilters
        {
            get => _onPremisesInstanceTagFilters ?? (_onPremisesInstanceTagFilters = new InputList<Inputs.DeploymentGroupOnPremisesInstanceTagFilterArgs>());
            set => _onPremisesInstanceTagFilters = value;
        }

        [Input("serviceRoleArn", required: true)]
        public Input<string> ServiceRoleArn { get; set; } = null!;

        [Input("triggerConfigurations")]
        private InputList<Inputs.DeploymentGroupTriggerConfigurationArgs>? _triggerConfigurations;
        public InputList<Inputs.DeploymentGroupTriggerConfigurationArgs> TriggerConfigurations
        {
            get => _triggerConfigurations ?? (_triggerConfigurations = new InputList<Inputs.DeploymentGroupTriggerConfigurationArgs>());
            set => _triggerConfigurations = value;
        }

        public DeploymentGroupArgs()
        {
        }
    }

    public sealed class DeploymentGroupState : Pulumi.ResourceArgs
    {
        [Input("alarmConfiguration")]
        public Input<Inputs.DeploymentGroupAlarmConfigurationGetArgs>? AlarmConfiguration { get; set; }

        [Input("appName")]
        public Input<string>? AppName { get; set; }

        [Input("autoRollbackConfiguration")]
        public Input<Inputs.DeploymentGroupAutoRollbackConfigurationGetArgs>? AutoRollbackConfiguration { get; set; }

        [Input("autoscalingGroups")]
        private InputList<string>? _autoscalingGroups;
        public InputList<string> AutoscalingGroups
        {
            get => _autoscalingGroups ?? (_autoscalingGroups = new InputList<string>());
            set => _autoscalingGroups = value;
        }

        [Input("blueGreenDeploymentConfig")]
        public Input<Inputs.DeploymentGroupBlueGreenDeploymentConfigGetArgs>? BlueGreenDeploymentConfig { get; set; }

        [Input("deploymentConfigName")]
        public Input<string>? DeploymentConfigName { get; set; }

        [Input("deploymentGroupName")]
        public Input<string>? DeploymentGroupName { get; set; }

        [Input("deploymentStyle")]
        public Input<Inputs.DeploymentGroupDeploymentStyleGetArgs>? DeploymentStyle { get; set; }

        [Input("ec2TagFilters")]
        private InputList<Inputs.DeploymentGroupEc2TagFilterGetArgs>? _ec2TagFilters;
        public InputList<Inputs.DeploymentGroupEc2TagFilterGetArgs> Ec2TagFilters
        {
            get => _ec2TagFilters ?? (_ec2TagFilters = new InputList<Inputs.DeploymentGroupEc2TagFilterGetArgs>());
            set => _ec2TagFilters = value;
        }

        [Input("ec2TagSets")]
        private InputList<Inputs.DeploymentGroupEc2TagSetGetArgs>? _ec2TagSets;
        public InputList<Inputs.DeploymentGroupEc2TagSetGetArgs> Ec2TagSets
        {
            get => _ec2TagSets ?? (_ec2TagSets = new InputList<Inputs.DeploymentGroupEc2TagSetGetArgs>());
            set => _ec2TagSets = value;
        }

        [Input("ecsService")]
        public Input<Inputs.DeploymentGroupEcsServiceGetArgs>? EcsService { get; set; }

        [Input("loadBalancerInfo")]
        public Input<Inputs.DeploymentGroupLoadBalancerInfoGetArgs>? LoadBalancerInfo { get; set; }

        [Input("onPremisesInstanceTagFilters")]
        private InputList<Inputs.DeploymentGroupOnPremisesInstanceTagFilterGetArgs>? _onPremisesInstanceTagFilters;
        public InputList<Inputs.DeploymentGroupOnPremisesInstanceTagFilterGetArgs> OnPremisesInstanceTagFilters
        {
            get => _onPremisesInstanceTagFilters ?? (_onPremisesInstanceTagFilters = new InputList<Inputs.DeploymentGroupOnPremisesInstanceTagFilterGetArgs>());
            set => _onPremisesInstanceTagFilters = value;
        }

        [Input("serviceRoleArn")]
        public Input<string>? ServiceRoleArn { get; set; }

        [Input("triggerConfigurations")]
        private InputList<Inputs.DeploymentGroupTriggerConfigurationGetArgs>? _triggerConfigurations;
        public InputList<Inputs.DeploymentGroupTriggerConfigurationGetArgs> TriggerConfigurations
        {
            get => _triggerConfigurations ?? (_triggerConfigurations = new InputList<Inputs.DeploymentGroupTriggerConfigurationGetArgs>());
            set => _triggerConfigurations = value;
        }

        public DeploymentGroupState()
        {
        }
    }
}
