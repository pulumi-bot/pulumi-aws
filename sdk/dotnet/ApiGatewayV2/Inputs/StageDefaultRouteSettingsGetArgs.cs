// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGatewayV2.Inputs
{

    public sealed class StageDefaultRouteSettingsGetArgs : Pulumi.ResourceArgs
    {
        [Input("dataTraceEnabled")]
        public Input<bool>? DataTraceEnabled { get; set; }

        [Input("detailedMetricsEnabled")]
        public Input<bool>? DetailedMetricsEnabled { get; set; }

        [Input("loggingLevel")]
        public Input<string>? LoggingLevel { get; set; }

        [Input("throttlingBurstLimit")]
        public Input<int>? ThrottlingBurstLimit { get; set; }

        [Input("throttlingRateLimit")]
        public Input<double>? ThrottlingRateLimit { get; set; }

        public StageDefaultRouteSettingsGetArgs()
        {
        }
    }
}
