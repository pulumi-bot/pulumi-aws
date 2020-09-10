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
    public sealed class GetDomainVpcOptionResult
    {
        public readonly ImmutableArray<string> AvailabilityZones;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly ImmutableArray<string> SubnetIds;
        public readonly string VpcId;

        [OutputConstructor]
        private GetDomainVpcOptionResult(
            ImmutableArray<string> availabilityZones,

            ImmutableArray<string> securityGroupIds,

            ImmutableArray<string> subnetIds,

            string vpcId)
        {
            AvailabilityZones = availabilityZones;
            SecurityGroupIds = securityGroupIds;
            SubnetIds = subnetIds;
            VpcId = vpcId;
        }
    }
}
