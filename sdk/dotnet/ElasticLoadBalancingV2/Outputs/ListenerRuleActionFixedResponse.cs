// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticLoadBalancingV2.Outputs
{

    [OutputType]
    public sealed class ListenerRuleActionFixedResponse
    {
        public readonly string ContentType;
        public readonly string? MessageBody;
        public readonly string? StatusCode;

        [OutputConstructor]
        private ListenerRuleActionFixedResponse(
            string contentType,

            string? messageBody,

            string? statusCode)
        {
            ContentType = contentType;
            MessageBody = messageBody;
            StatusCode = statusCode;
        }
    }
}
