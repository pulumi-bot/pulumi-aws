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
    public sealed class WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatch
    {
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchAllQueryArguments? AllQueryArguments;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchBody? Body;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchMethod? Method;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchQueryString? QueryString;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchSingleHeader? SingleHeader;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchSingleQueryArgument? SingleQueryArgument;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchUriPath? UriPath;

        [OutputConstructor]
        private WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatch(
            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchAllQueryArguments? allQueryArguments,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchBody? body,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchMethod? method,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchQueryString? queryString,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchSingleHeader? singleHeader,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchSingleQueryArgument? singleQueryArgument,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatementFieldToMatchUriPath? uriPath)
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
