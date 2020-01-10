// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the IP ranges of various AWS products and services. For more information about the contents of this data source and required JSON syntax if referencing a custom URL, see the [AWS IP Address Ranges documention][1].
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ip_ranges.html.markdown.
        /// </summary>
        public static Task<GetIpRangesResult> GetIpRanges(GetIpRangesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIpRangesResult>("aws:index/getIpRanges:getIpRanges", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetIpRangesArgs : Pulumi.InvokeArgs
    {
        [Input("regions")]
        private List<string>? _regions;

        /// <summary>
        /// Filter IP ranges by regions (or include all regions, if
        /// omitted). Valid items are `global` (for `cloudfront`) as well as all AWS regions
        /// (e.g. `eu-central-1`)
        /// </summary>
        public List<string> Regions
        {
            get => _regions ?? (_regions = new List<string>());
            set => _regions = value;
        }

        [Input("services", required: true)]
        private List<string>? _services;

        /// <summary>
        /// Filter IP ranges by services. Valid items are `amazon`
        /// (for amazon.com), `cloudfront`, `codebuild`, `ec2`, `route53`, `route53_healthchecks` and `S3`.
        /// </summary>
        public List<string> Services
        {
            get => _services ?? (_services = new List<string>());
            set => _services = value;
        }

        /// <summary>
        /// Custom URL for source JSON file. Syntax must match [AWS IP Address Ranges documention][1]. Defaults to `https://ip-ranges.amazonaws.com/ip-ranges.json`.
        /// </summary>
        [Input("url")]
        public string? Url { get; set; }

        public GetIpRangesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetIpRangesResult
    {
        /// <summary>
        /// The lexically ordered list of CIDR blocks.
        /// </summary>
        public readonly ImmutableArray<string> CidrBlocks;
        /// <summary>
        /// The publication time of the IP ranges (e.g. `2016-08-03-23-46-05`).
        /// </summary>
        public readonly string CreateDate;
        /// <summary>
        /// The lexically ordered list of IPv6 CIDR blocks.
        /// </summary>
        public readonly ImmutableArray<string> Ipv6CidrBlocks;
        public readonly ImmutableArray<string> Regions;
        public readonly ImmutableArray<string> Services;
        /// <summary>
        /// The publication time of the IP ranges, in Unix epoch time format
        /// (e.g. `1470267965`).
        /// </summary>
        public readonly int SyncToken;
        public readonly string? Url;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetIpRangesResult(
            ImmutableArray<string> cidrBlocks,
            string createDate,
            ImmutableArray<string> ipv6CidrBlocks,
            ImmutableArray<string> regions,
            ImmutableArray<string> services,
            int syncToken,
            string? url,
            string id)
        {
            CidrBlocks = cidrBlocks;
            CreateDate = createDate;
            Ipv6CidrBlocks = ipv6CidrBlocks;
            Regions = regions;
            Services = services;
            SyncToken = syncToken;
            Url = url;
            Id = id;
        }
    }
}
