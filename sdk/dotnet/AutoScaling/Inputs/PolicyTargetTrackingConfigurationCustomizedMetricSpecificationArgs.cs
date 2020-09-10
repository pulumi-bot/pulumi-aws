// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling.Inputs
{

    public sealed class PolicyTargetTrackingConfigurationCustomizedMetricSpecificationArgs : Pulumi.ResourceArgs
    {
        [Input("metricDimensions")]
        private InputList<Inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionArgs>? _metricDimensions;
        public InputList<Inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionArgs> MetricDimensions
        {
            get => _metricDimensions ?? (_metricDimensions = new InputList<Inputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionArgs>());
            set => _metricDimensions = value;
        }

        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        [Input("statistic", required: true)]
        public Input<string> Statistic { get; set; } = null!;

        [Input("unit")]
        public Input<string>? Unit { get; set; }

        public PolicyTargetTrackingConfigurationCustomizedMetricSpecificationArgs()
        {
        }
    }
}
