// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchGetArgs : Pulumi.ResourceArgs
    {
        [Input("allQueryArguments")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchAllQueryArgumentsGetArgs>? AllQueryArguments { get; set; }

        [Input("body")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchBodyGetArgs>? Body { get; set; }

        [Input("method")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchMethodGetArgs>? Method { get; set; }

        [Input("queryString")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchQueryStringGetArgs>? QueryString { get; set; }

        [Input("singleHeader")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchSingleHeaderGetArgs>? SingleHeader { get; set; }

        [Input("singleQueryArgument")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgumentGetArgs>? SingleQueryArgument { get; set; }

        [Input("uriPath")]
        public Input<Inputs.RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchUriPathGetArgs>? UriPath { get; set; }

        public RuleGroupRuleStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchGetArgs()
        {
        }
    }
}
