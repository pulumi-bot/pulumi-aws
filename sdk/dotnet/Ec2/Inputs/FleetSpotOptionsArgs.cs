// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class FleetSpotOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The order of the launch template overrides to use in fulfilling On-Demand capacity. Valid values: `lowestPrice`, `prioritized`. Default: `lowestPrice`.
        /// </summary>
        [Input("allocationStrategy")]
        public Input<string>? AllocationStrategy { get; set; }

        /// <summary>
        /// Behavior when a Spot Instance is interrupted. Valid values: `hibernate`, `stop`, `terminate`. Default: `terminate`.
        /// </summary>
        [Input("instanceInterruptionBehavior")]
        public Input<string>? InstanceInterruptionBehavior { get; set; }

        /// <summary>
        /// Number of Spot pools across which to allocate your target Spot capacity. Valid only when Spot `allocation_strategy` is set to `lowestPrice`. Default: `1`.
        /// </summary>
        [Input("instancePoolsToUseCount")]
        public Input<int>? InstancePoolsToUseCount { get; set; }

        public FleetSpotOptionsArgs()
        {
        }
    }
}
