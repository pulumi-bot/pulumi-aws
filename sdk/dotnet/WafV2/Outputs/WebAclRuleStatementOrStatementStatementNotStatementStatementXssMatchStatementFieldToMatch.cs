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
    public sealed class WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatch
    {
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchAllQueryArguments? AllQueryArguments;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchBody? Body;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchMethod? Method;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchQueryString? QueryString;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchSingleHeader? SingleHeader;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchSingleQueryArgument? SingleQueryArgument;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchUriPath? UriPath;

        [OutputConstructor]
        private WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatch(
            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchAllQueryArguments? allQueryArguments,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchBody? body,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchMethod? method,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchQueryString? queryString,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchSingleHeader? singleHeader,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchSingleQueryArgument? singleQueryArgument,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatementFieldToMatchUriPath? uriPath)
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
