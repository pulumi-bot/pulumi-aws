// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatementFieldToMatchGetArgs : Pulumi.ResourceArgs
    {
        [Input("allQueryArguments")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatementFieldToMatchAllQueryArgumentsGetArgs>? AllQueryArguments { get; set; }

        [Input("body")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatementFieldToMatchBodyGetArgs>? Body { get; set; }

        [Input("method")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatementFieldToMatchMethodGetArgs>? Method { get; set; }

        [Input("queryString")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatementFieldToMatchQueryStringGetArgs>? QueryString { get; set; }

        [Input("singleHeader")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatementFieldToMatchSingleHeaderGetArgs>? SingleHeader { get; set; }

        [Input("singleQueryArgument")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatementFieldToMatchSingleQueryArgumentGetArgs>? SingleQueryArgument { get; set; }

        [Input("uriPath")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatementFieldToMatchUriPathGetArgs>? UriPath { get; set; }

        public WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatementFieldToMatchGetArgs()
        {
        }
    }
}
