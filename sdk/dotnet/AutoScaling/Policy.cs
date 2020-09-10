// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling
{
    public partial class Policy : Pulumi.CustomResource
    {
        [Output("adjustmentType")]
        public Output<string?> AdjustmentType { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("autoscalingGroupName")]
        public Output<string> AutoscalingGroupName { get; private set; } = null!;

        [Output("cooldown")]
        public Output<int?> Cooldown { get; private set; } = null!;

        [Output("estimatedInstanceWarmup")]
        public Output<int?> EstimatedInstanceWarmup { get; private set; } = null!;

        [Output("metricAggregationType")]
        public Output<string> MetricAggregationType { get; private set; } = null!;

        [Output("minAdjustmentMagnitude")]
        public Output<int?> MinAdjustmentMagnitude { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("policyType")]
        public Output<string?> PolicyType { get; private set; } = null!;

        [Output("scalingAdjustment")]
        public Output<int?> ScalingAdjustment { get; private set; } = null!;

        [Output("stepAdjustments")]
        public Output<ImmutableArray<Outputs.PolicyStepAdjustment>> StepAdjustments { get; private set; } = null!;

        [Output("targetTrackingConfiguration")]
        public Output<Outputs.PolicyTargetTrackingConfiguration?> TargetTrackingConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:autoscaling/policy:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:autoscaling/policy:Policy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, state, options);
        }
    }

    public sealed class PolicyArgs : Pulumi.ResourceArgs
    {
        [Input("adjustmentType")]
        public Input<string>? AdjustmentType { get; set; }

        [Input("autoscalingGroupName", required: true)]
        public Input<string> AutoscalingGroupName { get; set; } = null!;

        [Input("cooldown")]
        public Input<int>? Cooldown { get; set; }

        [Input("estimatedInstanceWarmup")]
        public Input<int>? EstimatedInstanceWarmup { get; set; }

        [Input("metricAggregationType")]
        public Input<string>? MetricAggregationType { get; set; }

        [Input("minAdjustmentMagnitude")]
        public Input<int>? MinAdjustmentMagnitude { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        [Input("scalingAdjustment")]
        public Input<int>? ScalingAdjustment { get; set; }

        [Input("stepAdjustments")]
        private InputList<Inputs.PolicyStepAdjustmentArgs>? _stepAdjustments;
        public InputList<Inputs.PolicyStepAdjustmentArgs> StepAdjustments
        {
            get => _stepAdjustments ?? (_stepAdjustments = new InputList<Inputs.PolicyStepAdjustmentArgs>());
            set => _stepAdjustments = value;
        }

        [Input("targetTrackingConfiguration")]
        public Input<Inputs.PolicyTargetTrackingConfigurationArgs>? TargetTrackingConfiguration { get; set; }

        public PolicyArgs()
        {
        }
    }

    public sealed class PolicyState : Pulumi.ResourceArgs
    {
        [Input("adjustmentType")]
        public Input<string>? AdjustmentType { get; set; }

        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("autoscalingGroupName")]
        public Input<string>? AutoscalingGroupName { get; set; }

        [Input("cooldown")]
        public Input<int>? Cooldown { get; set; }

        [Input("estimatedInstanceWarmup")]
        public Input<int>? EstimatedInstanceWarmup { get; set; }

        [Input("metricAggregationType")]
        public Input<string>? MetricAggregationType { get; set; }

        [Input("minAdjustmentMagnitude")]
        public Input<int>? MinAdjustmentMagnitude { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        [Input("scalingAdjustment")]
        public Input<int>? ScalingAdjustment { get; set; }

        [Input("stepAdjustments")]
        private InputList<Inputs.PolicyStepAdjustmentGetArgs>? _stepAdjustments;
        public InputList<Inputs.PolicyStepAdjustmentGetArgs> StepAdjustments
        {
            get => _stepAdjustments ?? (_stepAdjustments = new InputList<Inputs.PolicyStepAdjustmentGetArgs>());
            set => _stepAdjustments = value;
        }

        [Input("targetTrackingConfiguration")]
        public Input<Inputs.PolicyTargetTrackingConfigurationGetArgs>? TargetTrackingConfiguration { get; set; }

        public PolicyState()
        {
        }
    }
}
