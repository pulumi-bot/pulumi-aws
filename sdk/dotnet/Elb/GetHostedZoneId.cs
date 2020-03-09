// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Elb
{
    public static partial class GetHostedZoneId
    {
        /// <summary>
        /// Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
        /// in a given region for the purpose of using in an AWS Route53 Alias.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elb_hosted_zone_id.html.markdown.
        /// </summary>
        public static Task<GetHostedZoneIdResult> InvokeAsync(GetHostedZoneIdArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHostedZoneIdResult>("aws:elb/getHostedZoneId:getHostedZoneId", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetHostedZoneIdArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the region whose AWS ELB HostedZoneId is desired.
        /// Defaults to the region from the AWS provider configuration.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetHostedZoneIdArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetHostedZoneIdResult
    {
        public readonly string? Region;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetHostedZoneIdResult(
            string? region,
            string id)
        {
            Region = region;
            Id = id;
        }
    }
}
