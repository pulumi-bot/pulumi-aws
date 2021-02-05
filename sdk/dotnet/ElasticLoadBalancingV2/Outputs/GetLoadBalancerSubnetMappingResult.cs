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
    public sealed class GetLoadBalancerSubnetMappingResult
    {
        public readonly string AllocationId;
        public readonly string Ipv6Address;
        public readonly string OutpostId;
        public readonly string PrivateIpv4Address;
        public readonly string SubnetId;

        [OutputConstructor]
        private GetLoadBalancerSubnetMappingResult(
            string allocationId,

            string ipv6Address,

            string outpostId,

            string privateIpv4Address,

            string subnetId)
        {
            AllocationId = allocationId;
            Ipv6Address = ipv6Address;
            OutpostId = outpostId;
            PrivateIpv4Address = privateIpv4Address;
            SubnetId = subnetId;
        }
    }
}
