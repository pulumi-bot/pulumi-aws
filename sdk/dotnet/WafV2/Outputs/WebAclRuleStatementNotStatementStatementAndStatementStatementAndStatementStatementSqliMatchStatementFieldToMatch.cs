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
    public sealed class WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatch
    {
        public readonly Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchAllQueryArguments? AllQueryArguments;
        public readonly Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchBody? Body;
        public readonly Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchMethod? Method;
        public readonly Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchQueryString? QueryString;
        public readonly Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchSingleHeader? SingleHeader;
        public readonly Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgument? SingleQueryArgument;
        public readonly Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchUriPath? UriPath;

        [OutputConstructor]
        private WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatch(
            Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchAllQueryArguments? allQueryArguments,

            Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchBody? body,

            Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchMethod? method,

            Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchQueryString? queryString,

            Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchSingleHeader? singleHeader,

            Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgument? singleQueryArgument,

            Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchUriPath? uriPath)
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
