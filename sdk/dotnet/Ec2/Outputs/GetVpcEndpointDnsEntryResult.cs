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
    public sealed class GetVpcEndpointDnsEntryResult
    {
        public readonly string DnsName;
        public readonly string HostedZoneId;

        [OutputConstructor]
        private GetVpcEndpointDnsEntryResult(
            string dnsName,

            string hostedZoneId)
        {
            DnsName = dnsName;
            HostedZoneId = hostedZoneId;
        }
    }
}
