// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafRegional
{
    public static partial class Invokes
    {
        /// <summary>
        /// `aws.wafregional.RateBasedRule` Retrieves a WAF Regional Rate Based Rule Resource Id.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/wafregional_rate_based_rule.html.markdown.
        /// </summary>
        public static Task<GetRateBasedModResult> GetRateBasedMod(GetRateBasedModArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRateBasedModResult>("aws:wafregional/getRateBasedMod:getRateBasedMod", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetRateBasedModArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the WAF Regional rate based rule.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetRateBasedModArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetRateBasedModResult
    {
        public readonly string Name;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRateBasedModResult(
            string name,
            string id)
        {
            Name = name;
            Id = id;
        }
    }
}
