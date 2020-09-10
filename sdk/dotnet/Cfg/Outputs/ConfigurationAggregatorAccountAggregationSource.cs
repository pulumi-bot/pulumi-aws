// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cfg.Outputs
{

    [OutputType]
    public sealed class ConfigurationAggregatorAccountAggregationSource
    {
        public readonly ImmutableArray<string> AccountIds;
        public readonly bool? AllRegions;
        public readonly ImmutableArray<string> Regions;

        [OutputConstructor]
        private ConfigurationAggregatorAccountAggregationSource(
            ImmutableArray<string> accountIds,

            bool? allRegions,

            ImmutableArray<string> regions)
        {
            AccountIds = accountIds;
            AllRegions = allRegions;
            Regions = regions;
        }
    }
}
