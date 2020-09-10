// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Outputs
{

    [OutputType]
    public sealed class VirtualNodeSpecListenerHealthCheck
    {
        public readonly int HealthyThreshold;
        public readonly int IntervalMillis;
        public readonly string? Path;
        public readonly int? Port;
        public readonly string Protocol;
        public readonly int TimeoutMillis;
        public readonly int UnhealthyThreshold;

        [OutputConstructor]
        private VirtualNodeSpecListenerHealthCheck(
            int healthyThreshold,

            int intervalMillis,

            string? path,

            int? port,

            string protocol,

            int timeoutMillis,

            int unhealthyThreshold)
        {
            HealthyThreshold = healthyThreshold;
            IntervalMillis = intervalMillis;
            Path = path;
            Port = port;
            Protocol = protocol;
            TimeoutMillis = timeoutMillis;
            UnhealthyThreshold = unhealthyThreshold;
        }
    }
}
