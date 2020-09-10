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
    public sealed class SecurityGroupEgress
    {
        public readonly ImmutableArray<string> CidrBlocks;
        public readonly string? Description;
        public readonly int FromPort;
        public readonly ImmutableArray<string> Ipv6CidrBlocks;
        public readonly ImmutableArray<string> PrefixListIds;
        public readonly string Protocol;
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly bool? Self;
        public readonly int ToPort;

        [OutputConstructor]
        private SecurityGroupEgress(
            ImmutableArray<string> cidrBlocks,

            string? description,

            int fromPort,

            ImmutableArray<string> ipv6CidrBlocks,

            ImmutableArray<string> prefixListIds,

            string protocol,

            ImmutableArray<string> securityGroups,

            bool? self,

            int toPort)
        {
            CidrBlocks = cidrBlocks;
            Description = description;
            FromPort = fromPort;
            Ipv6CidrBlocks = ipv6CidrBlocks;
            PrefixListIds = prefixListIds;
            Protocol = protocol;
            SecurityGroups = securityGroups;
            Self = self;
            ToPort = toPort;
        }
    }
}
