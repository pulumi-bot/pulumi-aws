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
    public sealed class DomainDomainEndpointOptions
    {
        public readonly bool EnforceHttps;
        public readonly string? TlsSecurityPolicy;

        [OutputConstructor]
        private DomainDomainEndpointOptions(
            bool enforceHttps,

            string? tlsSecurityPolicy)
        {
            EnforceHttps = enforceHttps;
            TlsSecurityPolicy = tlsSecurityPolicy;
        }
    }
}
