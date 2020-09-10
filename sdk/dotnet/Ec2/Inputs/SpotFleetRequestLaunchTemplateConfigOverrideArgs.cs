// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class SpotFleetRequestLaunchTemplateConfigOverrideArgs : Pulumi.ResourceArgs
    {
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("priority")]
        public Input<double>? Priority { get; set; }

        [Input("spotPrice")]
        public Input<string>? SpotPrice { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("weightedCapacity")]
        public Input<double>? WeightedCapacity { get; set; }

        public SpotFleetRequestLaunchTemplateConfigOverrideArgs()
        {
        }
    }
}
