// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Outputs
{

    [OutputType]
    public sealed class FleetLaunchTemplateConfigOverride
    {
        public readonly string? AvailabilityZone;
        public readonly string? InstanceType;
        public readonly string? MaxPrice;
        public readonly double? Priority;
        public readonly string? SubnetId;
        public readonly double? WeightedCapacity;

        [OutputConstructor]
        private FleetLaunchTemplateConfigOverride(
            string? availabilityZone,

            string? instanceType,

            string? maxPrice,

            double? priority,

            string? subnetId,

            double? weightedCapacity)
        {
            AvailabilityZone = availabilityZone;
            InstanceType = instanceType;
            MaxPrice = maxPrice;
            Priority = priority;
            SubnetId = subnetId;
            WeightedCapacity = weightedCapacity;
        }
    }
}
