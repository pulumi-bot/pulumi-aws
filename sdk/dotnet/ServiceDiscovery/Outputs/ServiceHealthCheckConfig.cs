// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ServiceDiscovery.Outputs
{

    [OutputType]
    public sealed class ServiceHealthCheckConfig
    {
        /// <summary>
        /// The number of consecutive health checks. Maximum value of 10.
        /// </summary>
        public readonly int? FailureThreshold;
        /// <summary>
        /// The path that you want Route 53 to request when performing health checks. Route 53 automatically adds the DNS name for the service. If you don't specify a value, the default value is /.
        /// </summary>
        public readonly string? ResourcePath;
        /// <summary>
        /// The type of the resource, which indicates the value that Amazon Route 53 returns in response to DNS queries. Valid Values: A, AAAA, SRV, CNAME
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private ServiceHealthCheckConfig(
            int? failureThreshold,

            string? resourcePath,

            string? type)
        {
            FailureThreshold = failureThreshold;
            ResourcePath = resourcePath;
            Type = type;
        }
    }
}
