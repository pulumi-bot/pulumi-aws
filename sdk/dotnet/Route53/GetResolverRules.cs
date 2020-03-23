// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53
{
    public static partial class Invokes
    {
        /// <summary>
        /// `aws.route53.getResolverRules` provides details about a set of Route53 Resolver rules.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/route53_resolver_rules.html.markdown.
        /// </summary>
        [Obsolete("Use GetResolverRules.InvokeAsync() instead")]
        public static Task<GetResolverRulesResult> GetResolverRules(GetResolverRulesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResolverRulesResult>("aws:route53/getResolverRules:getResolverRules", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetResolverRules
    {
        /// <summary>
        /// `aws.route53.getResolverRules` provides details about a set of Route53 Resolver rules.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/route53_resolver_rules.html.markdown.
        /// </summary>
        public static Task<GetResolverRulesResult> InvokeAsync(GetResolverRulesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResolverRulesResult>("aws:route53/getResolverRules:getResolverRules", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetResolverRulesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// When the desired resolver rules are shared with another AWS account, the account ID of the account that the rules are shared with.
        /// </summary>
        [Input("ownerId")]
        public string? OwnerId { get; set; }

        /// <summary>
        /// The ID of the outbound resolver endpoint for the desired resolver rules.
        /// </summary>
        [Input("resolverEndpointId")]
        public string? ResolverEndpointId { get; set; }

        /// <summary>
        /// The rule type of the desired resolver rules. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
        /// </summary>
        [Input("ruleType")]
        public string? RuleType { get; set; }

        /// <summary>
        /// Whether the desired resolver rules are shared and, if so, whether the current account is sharing the rules with another account, or another account is sharing the rules with the current account.
        /// Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
        /// </summary>
        [Input("shareStatus")]
        public string? ShareStatus { get; set; }

        public GetResolverRulesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetResolverRulesResult
    {
        public readonly string? OwnerId;
        public readonly string? ResolverEndpointId;
        /// <summary>
        /// The IDs of the matched resolver rules.
        /// </summary>
        public readonly ImmutableArray<string> ResolverRuleIds;
        public readonly string? RuleType;
        public readonly string? ShareStatus;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetResolverRulesResult(
            string? ownerId,
            string? resolverEndpointId,
            ImmutableArray<string> resolverRuleIds,
            string? ruleType,
            string? shareStatus,
            string id)
        {
            OwnerId = ownerId;
            ResolverEndpointId = resolverEndpointId;
            ResolverRuleIds = resolverRuleIds;
            RuleType = ruleType;
            ShareStatus = shareStatus;
            Id = id;
        }
    }
}
