// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws
{
    public static class GetArn
    {
        public static Task<GetArnResult> InvokeAsync(GetArnArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetArnResult>("aws:index/getArn:getArn", args ?? new GetArnArgs(), options.WithVersion());
    }


    public sealed class GetArnArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetArnArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetArnResult
    {
        public readonly string Account;
        public readonly string Arn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Partition;
        public readonly string Region;
        public readonly string Resource;
        public readonly string Service;

        [OutputConstructor]
        private GetArnResult(
            string account,

            string arn,

            string id,

            string partition,

            string region,

            string resource,

            string service)
        {
            Account = account;
            Arn = arn;
            Id = id;
            Partition = partition;
            Region = region;
            Resource = resource;
            Service = service;
        }
    }
}
