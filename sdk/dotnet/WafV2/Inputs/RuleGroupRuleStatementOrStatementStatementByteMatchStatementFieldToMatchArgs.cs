// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class RuleGroupRuleStatementOrStatementStatementByteMatchStatementFieldToMatchArgs : Pulumi.ResourceArgs
    {
        [Input("allQueryArguments")]
        public Input<Inputs.RuleGroupRuleStatementOrStatementStatementByteMatchStatementFieldToMatchAllQueryArgumentsArgs>? AllQueryArguments { get; set; }

        [Input("body")]
        public Input<Inputs.RuleGroupRuleStatementOrStatementStatementByteMatchStatementFieldToMatchBodyArgs>? Body { get; set; }

        [Input("method")]
        public Input<Inputs.RuleGroupRuleStatementOrStatementStatementByteMatchStatementFieldToMatchMethodArgs>? Method { get; set; }

        [Input("queryString")]
        public Input<Inputs.RuleGroupRuleStatementOrStatementStatementByteMatchStatementFieldToMatchQueryStringArgs>? QueryString { get; set; }

        [Input("singleHeader")]
        public Input<Inputs.RuleGroupRuleStatementOrStatementStatementByteMatchStatementFieldToMatchSingleHeaderArgs>? SingleHeader { get; set; }

        [Input("singleQueryArgument")]
        public Input<Inputs.RuleGroupRuleStatementOrStatementStatementByteMatchStatementFieldToMatchSingleQueryArgumentArgs>? SingleQueryArgument { get; set; }

        [Input("uriPath")]
        public Input<Inputs.RuleGroupRuleStatementOrStatementStatementByteMatchStatementFieldToMatchUriPathArgs>? UriPath { get; set; }

        public RuleGroupRuleStatementOrStatementStatementByteMatchStatementFieldToMatchArgs()
        {
        }
    }
}
