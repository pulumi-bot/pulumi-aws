// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class RuleGroupRuleStatementNotStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchGetArgs : Pulumi.ResourceArgs
    {
        [Input("allQueryArguments")]
        public Input<Inputs.RuleGroupRuleStatementNotStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchAllQueryArgumentsGetArgs>? AllQueryArguments { get; set; }

        [Input("body")]
        public Input<Inputs.RuleGroupRuleStatementNotStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchBodyGetArgs>? Body { get; set; }

        [Input("method")]
        public Input<Inputs.RuleGroupRuleStatementNotStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchMethodGetArgs>? Method { get; set; }

        [Input("queryString")]
        public Input<Inputs.RuleGroupRuleStatementNotStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchQueryStringGetArgs>? QueryString { get; set; }

        [Input("singleHeader")]
        public Input<Inputs.RuleGroupRuleStatementNotStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleHeaderGetArgs>? SingleHeader { get; set; }

        [Input("singleQueryArgument")]
        public Input<Inputs.RuleGroupRuleStatementNotStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgumentGetArgs>? SingleQueryArgument { get; set; }

        [Input("uriPath")]
        public Input<Inputs.RuleGroupRuleStatementNotStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchUriPathGetArgs>? UriPath { get; set; }

        public RuleGroupRuleStatementNotStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchGetArgs()
        {
        }
    }
}
