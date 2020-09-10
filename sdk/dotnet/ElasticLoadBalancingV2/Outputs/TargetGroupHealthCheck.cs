// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticLoadBalancingV2.Outputs
{

    [OutputType]
    public sealed class TargetGroupHealthCheck
    {
        public readonly bool? Enabled;
        public readonly int? HealthyThreshold;
        public readonly int? Interval;
        public readonly string? Matcher;
        public readonly string? Path;
        public readonly string? Port;
        public readonly string? Protocol;
        public readonly int? Timeout;
        public readonly int? UnhealthyThreshold;

        [OutputConstructor]
        private TargetGroupHealthCheck(
            bool? enabled,

            int? healthyThreshold,

            int? interval,

            string? matcher,

            string? path,

            string? port,

            string? protocol,

            int? timeout,

            int? unhealthyThreshold)
        {
            Enabled = enabled;
            HealthyThreshold = healthyThreshold;
            Interval = interval;
            Matcher = matcher;
            Path = path;
            Port = port;
            Protocol = protocol;
            Timeout = timeout;
            UnhealthyThreshold = unhealthyThreshold;
        }
    }
}
