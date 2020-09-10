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
    public sealed class GetInstanceMetadataOptionResult
    {
        public readonly string HttpEndpoint;
        public readonly int HttpPutResponseHopLimit;
        public readonly string HttpTokens;

        [OutputConstructor]
        private GetInstanceMetadataOptionResult(
            string httpEndpoint,

            int httpPutResponseHopLimit,

            string httpTokens)
        {
            HttpEndpoint = httpEndpoint;
            HttpPutResponseHopLimit = httpPutResponseHopLimit;
            HttpTokens = httpTokens;
        }
    }
}
