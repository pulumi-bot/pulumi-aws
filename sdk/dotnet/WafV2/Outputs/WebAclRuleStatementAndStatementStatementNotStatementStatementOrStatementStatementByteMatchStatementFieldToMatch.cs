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
    public sealed class WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatch
    {
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchAllQueryArguments? AllQueryArguments;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchBody? Body;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchMethod? Method;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchQueryString? QueryString;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchSingleHeader? SingleHeader;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchSingleQueryArgument? SingleQueryArgument;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchUriPath? UriPath;

        [OutputConstructor]
        private WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatch(
            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchAllQueryArguments? allQueryArguments,

            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchBody? body,

            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchMethod? method,

            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchQueryString? queryString,

            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchSingleHeader? singleHeader,

            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchSingleQueryArgument? singleQueryArgument,

            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatementFieldToMatchUriPath? uriPath)
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
