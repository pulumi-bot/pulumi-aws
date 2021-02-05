// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticLoadBalancingV2.Inputs
{

    public sealed class LoadBalancerSubnetMappingGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The allocation ID of the Elastic IP address.
        /// </summary>
        [Input("allocationId")]
        public Input<string>? AllocationId { get; set; }

        /// <summary>
        /// An ipv6 address within the subnet to assign to the internet-facing load balancer.
        /// </summary>
        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        [Input("outpostId")]
        public Input<string>? OutpostId { get; set; }

        /// <summary>
        /// A private ipv4 address within the subnet to assign to the internal-facing load balancer.
        /// </summary>
        [Input("privateIpv4Address")]
        public Input<string>? PrivateIpv4Address { get; set; }

        /// <summary>
        /// The id of the subnet of which to attach to the load balancer. You can specify only one subnet per Availability Zone.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public LoadBalancerSubnetMappingGetArgs()
        {
        }
    }
}
