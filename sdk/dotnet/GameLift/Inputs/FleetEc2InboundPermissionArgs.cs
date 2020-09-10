// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GameLift.Inputs
{

    public sealed class FleetEc2InboundPermissionArgs : Pulumi.ResourceArgs
    {
        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        [Input("ipRange", required: true)]
        public Input<string> IpRange { get; set; } = null!;

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public FleetEc2InboundPermissionArgs()
        {
        }
    }
}
