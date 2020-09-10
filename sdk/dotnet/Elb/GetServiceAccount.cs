// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Elb
{
    public static class GetServiceAccount
    {
        public static Task<GetServiceAccountResult> InvokeAsync(GetServiceAccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceAccountResult>("aws:elb/getServiceAccount:getServiceAccount", args ?? new GetServiceAccountArgs(), options.WithVersion());
    }


    public sealed class GetServiceAccountArgs : Pulumi.InvokeArgs
    {
        [Input("region")]
        public string? Region { get; set; }

        public GetServiceAccountArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceAccountResult
    {
        public readonly string Arn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Region;

        [OutputConstructor]
        private GetServiceAccountResult(
            string arn,

            string id,

            string? region)
        {
            Arn = arn;
            Id = id;
            Region = region;
        }
    }
}
