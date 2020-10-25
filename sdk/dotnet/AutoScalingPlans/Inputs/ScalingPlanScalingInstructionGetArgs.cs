// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScalingPlans.Inputs
{

    public sealed class ScalingPlanScalingInstructionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The customized load metric to use for predictive scaling. You must specify either `customized_load_metric_specification` or `predefined_load_metric_specification` when configuring predictive scaling.
        /// More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_CustomizedLoadMetricSpecification.html).
        /// </summary>
        [Input("customizedLoadMetricSpecification")]
        public Input<Inputs.ScalingPlanScalingInstructionCustomizedLoadMetricSpecificationGetArgs>? CustomizedLoadMetricSpecification { get; set; }

        /// <summary>
        /// Boolean controlling whether dynamic scaling by AWS Auto Scaling is disabled. Defaults to `false`.
        /// </summary>
        [Input("disableDynamicScaling")]
        public Input<bool>? DisableDynamicScaling { get; set; }

        /// <summary>
        /// The maximum capacity of the resource. The exception to this upper limit is if you specify a non-default setting for `predictive_scaling_max_capacity_behavior`.
        /// </summary>
        [Input("maxCapacity", required: true)]
        public Input<int> MaxCapacity { get; set; } = null!;

        /// <summary>
        /// The minimum capacity of the resource.
        /// </summary>
        [Input("minCapacity", required: true)]
        public Input<int> MinCapacity { get; set; } = null!;

        /// <summary>
        /// The predefined load metric to use for predictive scaling. You must specify either `predefined_load_metric_specification` or `customized_load_metric_specification` when configuring predictive scaling.
        /// More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_PredefinedLoadMetricSpecification.html).
        /// </summary>
        [Input("predefinedLoadMetricSpecification")]
        public Input<Inputs.ScalingPlanScalingInstructionPredefinedLoadMetricSpecificationGetArgs>? PredefinedLoadMetricSpecification { get; set; }

        /// <summary>
        /// Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity specified for the resource.
        /// Valid values: `SetForecastCapacityToMaxCapacity`, `SetMaxCapacityAboveForecastCapacity`, `SetMaxCapacityToForecastCapacity`.
        /// </summary>
        [Input("predictiveScalingMaxCapacityBehavior")]
        public Input<string>? PredictiveScalingMaxCapacityBehavior { get; set; }

        /// <summary>
        /// The size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity.
        /// </summary>
        [Input("predictiveScalingMaxCapacityBuffer")]
        public Input<int>? PredictiveScalingMaxCapacityBuffer { get; set; }

        /// <summary>
        /// The predictive scaling mode. Valid values: `ForecastAndScale`, `ForecastOnly`.
        /// </summary>
        [Input("predictiveScalingMode")]
        public Input<string>? PredictiveScalingMode { get; set; }

        /// <summary>
        /// The ID of the resource. This string consists of the resource type and unique identifier.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The scalable dimension associated with the resource. Valid values: `autoscaling:autoScalingGroup:DesiredCapacity`, `dynamodb:index:ReadCapacityUnits`, `dynamodb:index:WriteCapacityUnits`, `dynamodb:table:ReadCapacityUnits`, `dynamodb:table:WriteCapacityUnits`, `ecs:service:DesiredCount`, `ec2:spot-fleet-request:TargetCapacity`, `rds:cluster:ReadReplicaCount`.
        /// </summary>
        [Input("scalableDimension", required: true)]
        public Input<string> ScalableDimension { get; set; } = null!;

        /// <summary>
        /// Controls whether a resource's externally created scaling policies are kept or replaced. Valid values: `KeepExternalPolicies`, `ReplaceExternalPolicies`. Defaults to `KeepExternalPolicies`.
        /// </summary>
        [Input("scalingPolicyUpdateBehavior")]
        public Input<string>? ScalingPolicyUpdateBehavior { get; set; }

        /// <summary>
        /// The amount of time, in seconds, to buffer the run time of scheduled scaling actions when scaling out.
        /// </summary>
        [Input("scheduledActionBufferTime")]
        public Input<int>? ScheduledActionBufferTime { get; set; }

        /// <summary>
        /// The namespace of the AWS service. Valid values: `autoscaling`, `dynamodb`, `ecs`, `ec2`, `rds`.
        /// </summary>
        [Input("serviceNamespace", required: true)]
        public Input<string> ServiceNamespace { get; set; } = null!;

        [Input("targetTrackingConfigurations", required: true)]
        private InputList<Inputs.ScalingPlanScalingInstructionTargetTrackingConfigurationGetArgs>? _targetTrackingConfigurations;

        /// <summary>
        /// The structure that defines new target tracking configurations. Each of these structures includes a specific scaling metric and a target value for the metric, along with various parameters to use with dynamic scaling.
        /// More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_TargetTrackingConfiguration.html).
        /// </summary>
        public InputList<Inputs.ScalingPlanScalingInstructionTargetTrackingConfigurationGetArgs> TargetTrackingConfigurations
        {
            get => _targetTrackingConfigurations ?? (_targetTrackingConfigurations = new InputList<Inputs.ScalingPlanScalingInstructionTargetTrackingConfigurationGetArgs>());
            set => _targetTrackingConfigurations = value;
        }

        public ScalingPlanScalingInstructionGetArgs()
        {
        }
    }
}
