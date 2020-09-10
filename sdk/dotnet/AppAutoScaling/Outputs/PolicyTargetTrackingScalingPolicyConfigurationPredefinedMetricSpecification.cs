// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppAutoScaling.Outputs
{

    [OutputType]
    public sealed class PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification
    {
        public readonly string PredefinedMetricType;
        public readonly string? ResourceLabel;

        [OutputConstructor]
        private PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification(
            string predefinedMetricType,

            string? resourceLabel)
        {
            PredefinedMetricType = predefinedMetricType;
            ResourceLabel = resourceLabel;
        }
    }
}
