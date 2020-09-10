// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class InstanceNetworkInterfaceGetArgs : Pulumi.ResourceArgs
    {
        [Input("deleteOnTermination")]
        public Input<bool>? DeleteOnTermination { get; set; }

        [Input("deviceIndex", required: true)]
        public Input<int> DeviceIndex { get; set; } = null!;

        [Input("networkInterfaceId", required: true)]
        public Input<string> NetworkInterfaceId { get; set; } = null!;

        public InstanceNetworkInterfaceGetArgs()
        {
        }
    }
}
