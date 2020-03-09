// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf
{
    public static partial class GetIpset
    {
        /// <summary>
        /// `aws.waf.IpSet` Retrieves a WAF IP Set Resource Id.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/waf_ipset.html.markdown.
        /// </summary>
        public static Task<GetIpsetResult> InvokeAsync(GetIpsetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIpsetResult>("aws:waf/getIpset:getIpset", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetIpsetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the WAF IP set.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetIpsetArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetIpsetResult
    {
        public readonly string Name;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetIpsetResult(
            string name,
            string id)
        {
            Name = name;
            Id = id;
        }
    }
}
