// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchGetArgs : Pulumi.ResourceArgs
    {
        [Input("allQueryArguments")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsGetArgs>? AllQueryArguments { get; set; }

        [Input("body")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBodyGetArgs>? Body { get; set; }

        [Input("method")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethodGetArgs>? Method { get; set; }

        [Input("queryString")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryStringGetArgs>? QueryString { get; set; }

        [Input("singleHeader")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderGetArgs>? SingleHeader { get; set; }

        [Input("singleQueryArgument")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentGetArgs>? SingleQueryArgument { get; set; }

        [Input("uriPath")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPathGetArgs>? UriPath { get; set; }

        public WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchGetArgs()
        {
        }
    }
}
