// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticLoadBalancingV2.Inputs
{

    public sealed class LoadBalancerSubnetMappingArgs : Pulumi.ResourceArgs
    {
        [Input("allocationId")]
        public Input<string>? AllocationId { get; set; }

        [Input("privateIpv4Address")]
        public Input<string>? PrivateIpv4Address { get; set; }

        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public LoadBalancerSubnetMappingArgs()
        {
        }
    }
}
