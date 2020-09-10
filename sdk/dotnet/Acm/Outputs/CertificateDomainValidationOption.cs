// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Acm.Outputs
{

    [OutputType]
    public sealed class CertificateDomainValidationOption
    {
        public readonly string? DomainName;
        public readonly string? ResourceRecordName;
        public readonly string? ResourceRecordType;
        public readonly string? ResourceRecordValue;

        [OutputConstructor]
        private CertificateDomainValidationOption(
            string? domainName,

            string? resourceRecordName,

            string? resourceRecordType,

            string? resourceRecordValue)
        {
            DomainName = domainName;
            ResourceRecordName = resourceRecordName;
            ResourceRecordType = resourceRecordType;
            ResourceRecordValue = resourceRecordValue;
        }
    }
}
