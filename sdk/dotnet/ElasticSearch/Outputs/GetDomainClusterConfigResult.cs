// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticSearch.Outputs
{

    [OutputType]
    public sealed class GetDomainClusterConfigResult
    {
        public readonly int DedicatedMasterCount;
        public readonly bool DedicatedMasterEnabled;
        public readonly string DedicatedMasterType;
        public readonly int InstanceCount;
        public readonly string InstanceType;
        public readonly int WarmCount;
        public readonly bool? WarmEnabled;
        public readonly string WarmType;
        public readonly ImmutableArray<Outputs.GetDomainClusterConfigZoneAwarenessConfigResult> ZoneAwarenessConfigs;
        public readonly bool ZoneAwarenessEnabled;

        [OutputConstructor]
        private GetDomainClusterConfigResult(
            int dedicatedMasterCount,

            bool dedicatedMasterEnabled,

            string dedicatedMasterType,

            int instanceCount,

            string instanceType,

            int warmCount,

            bool? warmEnabled,

            string warmType,

            ImmutableArray<Outputs.GetDomainClusterConfigZoneAwarenessConfigResult> zoneAwarenessConfigs,

            bool zoneAwarenessEnabled)
        {
            DedicatedMasterCount = dedicatedMasterCount;
            DedicatedMasterEnabled = dedicatedMasterEnabled;
            DedicatedMasterType = dedicatedMasterType;
            InstanceCount = instanceCount;
            InstanceType = instanceType;
            WarmCount = warmCount;
            WarmEnabled = warmEnabled;
            WarmType = warmType;
            ZoneAwarenessConfigs = zoneAwarenessConfigs;
            ZoneAwarenessEnabled = zoneAwarenessEnabled;
        }
    }
}
