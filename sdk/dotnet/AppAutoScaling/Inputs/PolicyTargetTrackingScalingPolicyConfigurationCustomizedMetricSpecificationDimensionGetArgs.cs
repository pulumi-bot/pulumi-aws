// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppAutoScaling.Inputs
{

    public sealed class PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionGetArgs : Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public PolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensionGetArgs()
        {
        }
    }
}
