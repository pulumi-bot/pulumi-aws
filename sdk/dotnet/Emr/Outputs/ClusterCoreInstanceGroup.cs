// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Emr.Outputs
{

    [OutputType]
    public sealed class ClusterCoreInstanceGroup
    {
        public readonly string? AutoscalingPolicy;
        public readonly string? BidPrice;
        public readonly ImmutableArray<Outputs.ClusterCoreInstanceGroupEbsConfig> EbsConfigs;
        public readonly string? Id;
        public readonly int? InstanceCount;
        public readonly string InstanceType;
        public readonly string? Name;

        [OutputConstructor]
        private ClusterCoreInstanceGroup(
            string? autoscalingPolicy,

            string? bidPrice,

            ImmutableArray<Outputs.ClusterCoreInstanceGroupEbsConfig> ebsConfigs,

            string? id,

            int? instanceCount,

            string instanceType,

            string? name)
        {
            AutoscalingPolicy = autoscalingPolicy;
            BidPrice = bidPrice;
            EbsConfigs = ebsConfigs;
            Id = id;
            InstanceCount = instanceCount;
            InstanceType = instanceType;
            Name = name;
        }
    }
}
