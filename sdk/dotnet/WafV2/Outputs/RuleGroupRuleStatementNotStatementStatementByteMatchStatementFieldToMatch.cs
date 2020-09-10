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
    public sealed class RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatch
    {
        public readonly Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchAllQueryArguments? AllQueryArguments;
        public readonly Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchBody? Body;
        public readonly Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchMethod? Method;
        public readonly Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchQueryString? QueryString;
        public readonly Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchSingleHeader? SingleHeader;
        public readonly Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchSingleQueryArgument? SingleQueryArgument;
        public readonly Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchUriPath? UriPath;

        [OutputConstructor]
        private RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatch(
            Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchAllQueryArguments? allQueryArguments,

            Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchBody? body,

            Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchMethod? method,

            Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchQueryString? queryString,

            Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchSingleHeader? singleHeader,

            Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchSingleQueryArgument? singleQueryArgument,

            Outputs.RuleGroupRuleStatementNotStatementStatementByteMatchStatementFieldToMatchUriPath? uriPath)
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
