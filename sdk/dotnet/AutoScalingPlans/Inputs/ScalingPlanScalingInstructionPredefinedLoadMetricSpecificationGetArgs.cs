// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScalingPlans.Inputs
{

    public sealed class ScalingPlanScalingInstructionPredefinedLoadMetricSpecificationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The metric type. Valid values: `ALBTargetGroupRequestCount`, `ASGTotalCPUUtilization`, `ASGTotalNetworkIn`, `ASGTotalNetworkOut`.
        /// </summary>
        [Input("predefinedLoadMetricType", required: true)]
        public Input<string> PredefinedLoadMetricType { get; set; } = null!;

        /// <summary>
        /// Identifies the resource associated with the metric type.
        /// </summary>
        [Input("resourceLabel")]
        public Input<string>? ResourceLabel { get; set; }

        public ScalingPlanScalingInstructionPredefinedLoadMetricSpecificationGetArgs()
        {
        }
    }
}
