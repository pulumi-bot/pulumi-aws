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
    public sealed class WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatch
    {
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchAllQueryArguments? AllQueryArguments;
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchBody? Body;
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchMethod? Method;
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchQueryString? QueryString;
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchSingleHeader? SingleHeader;
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchSingleQueryArgument? SingleQueryArgument;
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchUriPath? UriPath;

        [OutputConstructor]
        private WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatch(
            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchAllQueryArguments? allQueryArguments,

            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchBody? body,

            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchMethod? method,

            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchQueryString? queryString,

            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchSingleHeader? singleHeader,

            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchSingleQueryArgument? singleQueryArgument,

            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatementFieldToMatchUriPath? uriPath)
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
