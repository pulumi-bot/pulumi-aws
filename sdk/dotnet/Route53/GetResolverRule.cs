// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53
{
    public static class GetResolverRule
    {
        public static Task<GetResolverRuleResult> InvokeAsync(GetResolverRuleArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResolverRuleResult>("aws:route53/getResolverRule:getResolverRule", args ?? new GetResolverRuleArgs(), options.WithVersion());
    }


    public sealed class GetResolverRuleArgs : Pulumi.InvokeArgs
    {
        [Input("domainName")]
        public string? DomainName { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("resolverEndpointId")]
        public string? ResolverEndpointId { get; set; }

        [Input("resolverRuleId")]
        public string? ResolverRuleId { get; set; }

        [Input("ruleType")]
        public string? RuleType { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetResolverRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResolverRuleResult
    {
        public readonly string Arn;
        public readonly string DomainName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string OwnerId;
        public readonly string ResolverEndpointId;
        public readonly string ResolverRuleId;
        public readonly string RuleType;
        public readonly string ShareStatus;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetResolverRuleResult(
            string arn,

            string domainName,

            string id,

            string name,

            string ownerId,

            string resolverEndpointId,

            string resolverRuleId,

            string ruleType,

            string shareStatus,

            ImmutableDictionary<string, string> tags)
        {
            Arn = arn;
            DomainName = domainName;
            Id = id;
            Name = name;
            OwnerId = ownerId;
            ResolverEndpointId = resolverEndpointId;
            ResolverRuleId = resolverRuleId;
            RuleType = ruleType;
            ShareStatus = shareStatus;
            Tags = tags;
        }
    }
}
