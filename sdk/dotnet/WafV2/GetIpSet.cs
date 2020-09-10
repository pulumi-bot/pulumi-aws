// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2
{
    public static class GetIpSet
    {
        public static Task<GetIpSetResult> InvokeAsync(GetIpSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIpSetResult>("aws:wafv2/getIpSet:getIpSet", args ?? new GetIpSetArgs(), options.WithVersion());
    }


    public sealed class GetIpSetArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("scope", required: true)]
        public string Scope { get; set; } = null!;

        public GetIpSetArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIpSetResult
    {
        public readonly ImmutableArray<string> Addresses;
        public readonly string Arn;
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IpAddressVersion;
        public readonly string Name;
        public readonly string Scope;

        [OutputConstructor]
        private GetIpSetResult(
            ImmutableArray<string> addresses,

            string arn,

            string description,

            string id,

            string ipAddressVersion,

            string name,

            string scope)
        {
            Addresses = addresses;
            Arn = arn;
            Description = description;
            Id = id;
            IpAddressVersion = ipAddressVersion;
            Name = name;
            Scope = scope;
        }
    }
}
