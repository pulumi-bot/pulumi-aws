// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class RuleGroupRuleStatementAndStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchGetArgs : Pulumi.ResourceArgs
    {
        [Input("allQueryArguments")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchAllQueryArgumentsGetArgs>? AllQueryArguments { get; set; }

        [Input("body")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchBodyGetArgs>? Body { get; set; }

        [Input("method")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchMethodGetArgs>? Method { get; set; }

        [Input("queryString")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchQueryStringGetArgs>? QueryString { get; set; }

        [Input("singleHeader")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleHeaderGetArgs>? SingleHeader { get; set; }

        [Input("singleQueryArgument")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgumentGetArgs>? SingleQueryArgument { get; set; }

        [Input("uriPath")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchUriPathGetArgs>? UriPath { get; set; }

        public RuleGroupRuleStatementAndStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchGetArgs()
        {
        }
    }
}
