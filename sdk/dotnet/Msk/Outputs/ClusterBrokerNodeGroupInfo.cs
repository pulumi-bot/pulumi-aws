// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Msk.Outputs
{

    [OutputType]
    public sealed class ClusterBrokerNodeGroupInfo
    {
        public readonly string? AzDistribution;
        public readonly ImmutableArray<string> ClientSubnets;
        public readonly int EbsVolumeSize;
        public readonly string InstanceType;
        public readonly ImmutableArray<string> SecurityGroups;

        [OutputConstructor]
        private ClusterBrokerNodeGroupInfo(
            string? azDistribution,

            ImmutableArray<string> clientSubnets,

            int ebsVolumeSize,

            string instanceType,

            ImmutableArray<string> securityGroups)
        {
            AzDistribution = azDistribution;
            ClientSubnets = clientSubnets;
            EbsVolumeSize = ebsVolumeSize;
            InstanceType = instanceType;
            SecurityGroups = securityGroups;
        }
    }
}
