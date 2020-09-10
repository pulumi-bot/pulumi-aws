// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclVisibilityConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("cloudwatchMetricsEnabled", required: true)]
        public Input<bool> CloudwatchMetricsEnabled { get; set; } = null!;

        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        [Input("sampledRequestsEnabled", required: true)]
        public Input<bool> SampledRequestsEnabled { get; set; } = null!;

        public WebAclVisibilityConfigGetArgs()
        {
        }
    }
}
