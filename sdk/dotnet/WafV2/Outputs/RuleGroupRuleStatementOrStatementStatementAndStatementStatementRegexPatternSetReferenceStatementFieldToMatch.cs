// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Outputs
{

    [OutputType]
    public sealed class RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch
    {
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments? AllQueryArguments;
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody? Body;
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod? Method;
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString? QueryString;
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader? SingleHeader;
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument? SingleQueryArgument;
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath? UriPath;

        [OutputConstructor]
        private RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch(
            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments? allQueryArguments,

            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody? body,

            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod? method,

            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString? queryString,

            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader? singleHeader,

            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument? singleQueryArgument,

            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath? uriPath)
        {
            AllQueryArguments = allQueryArguments;
            Body = body;
            Method = method;
            QueryString = queryString;
            SingleHeader = singleHeader;
            SingleQueryArgument = singleQueryArgument;
            UriPath = uriPath;
        }
    }
}
