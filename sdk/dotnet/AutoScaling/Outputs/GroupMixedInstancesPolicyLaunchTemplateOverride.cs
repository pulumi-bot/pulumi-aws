// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling.Outputs
{

    [OutputType]
    public sealed class GroupMixedInstancesPolicyLaunchTemplateOverride
    {
        public readonly string? InstanceType;
        public readonly string? WeightedCapacity;

        [OutputConstructor]
        private GroupMixedInstancesPolicyLaunchTemplateOverride(
            string? instanceType,

            string? weightedCapacity)
        {
            InstanceType = instanceType;
            WeightedCapacity = weightedCapacity;
        }
    }
}
