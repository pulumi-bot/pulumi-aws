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
        /// <summary>
        /// List of CIDR blocks.
        /// </summary>
        public readonly ImmutableArray<string> CidrBlocks;
        /// <summary>
        /// The security group description. Defaults to
        /// "Managed by Pulumi". Cannot be "". __NOTE__: This field maps to the AWS
        /// `GroupDescription` attribute, for which there is no Update API. If you'd like
        /// to classify your security groups in a way that can be updated, use `tags`.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The start port (or ICMP type number if protocol is "icmp" or "icmpv6")
        /// </summary>
        public readonly int FromPort;
        /// <summary>
        /// List of IPv6 CIDR blocks.
        /// </summary>
        public readonly ImmutableArray<string> Ipv6CidrBlocks;
        /// <summary>
        /// List of prefix list IDs.
        /// </summary>
        public readonly ImmutableArray<string> PrefixListIds;
        /// <summary>
        /// The protocol. If you select a protocol of
        /// "-1" (semantically equivalent to `"all"`, which is not a valid value here), you must specify a "from_port" and "to_port" equal to 0. If not icmp, icmpv6, tcp, udp, or "-1" use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// List of security group Group Names if using
        /// EC2-Classic, or Group IDs if using a VPC.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        /// <summary>
        /// If true, the security group itself will be added as
        /// a source to this ingress rule.
        /// </summary>
        public readonly bool? Self;
        /// <summary>
        /// The end range port (or ICMP code if protocol is "icmp").
        /// </summary>
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
