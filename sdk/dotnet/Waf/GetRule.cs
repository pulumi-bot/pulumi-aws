// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf
{
    public static class GetRule
    {
        /// <summary>
        /// `aws.waf.Rule` Retrieves a WAF Rule Resource Id.
        /// </summary>
        public static Task<GetRuleResult> InvokeAsync(GetRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRuleResult>("aws:waf/getRule:getRule", args ?? new GetRuleArgs(), options.WithVersion());
    }


    public sealed class GetRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the WAF rule.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRuleResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetRuleResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
