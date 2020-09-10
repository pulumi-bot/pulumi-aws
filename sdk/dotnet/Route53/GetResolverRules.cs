// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53
{
    public static class GetResolverRules
    {
        public static Task<GetResolverRulesResult> InvokeAsync(GetResolverRulesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResolverRulesResult>("aws:route53/getResolverRules:getResolverRules", args ?? new GetResolverRulesArgs(), options.WithVersion());
    }


    public sealed class GetResolverRulesArgs : Pulumi.InvokeArgs
    {
        [Input("ownerId")]
        public string? OwnerId { get; set; }

        [Input("resolverEndpointId")]
        public string? ResolverEndpointId { get; set; }

        [Input("ruleType")]
        public string? RuleType { get; set; }

        [Input("shareStatus")]
        public string? ShareStatus { get; set; }

        public GetResolverRulesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResolverRulesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OwnerId;
        public readonly string? ResolverEndpointId;
        public readonly ImmutableArray<string> ResolverRuleIds;
        public readonly string? RuleType;
        public readonly string? ShareStatus;

        [OutputConstructor]
        private GetResolverRulesResult(
            string id,

            string? ownerId,

            string? resolverEndpointId,

            ImmutableArray<string> resolverRuleIds,

            string? ruleType,

            string? shareStatus)
        {
            Id = id;
            OwnerId = ownerId;
            ResolverEndpointId = resolverEndpointId;
            ResolverRuleIds = resolverRuleIds;
            RuleType = ruleType;
            ShareStatus = shareStatus;
        }
    }
}
