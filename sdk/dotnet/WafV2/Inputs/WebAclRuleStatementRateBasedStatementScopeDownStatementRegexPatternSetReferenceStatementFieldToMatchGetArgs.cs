// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchGetArgs : Pulumi.ResourceArgs
    {
        [Input("allQueryArguments")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsGetArgs>? AllQueryArguments { get; set; }

        [Input("body")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchBodyGetArgs>? Body { get; set; }

        [Input("method")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodGetArgs>? Method { get; set; }

        [Input("queryString")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringGetArgs>? QueryString { get; set; }

        [Input("singleHeader")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderGetArgs>? SingleHeader { get; set; }

        [Input("singleQueryArgument")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentGetArgs>? SingleQueryArgument { get; set; }

        [Input("uriPath")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriPathGetArgs>? UriPath { get; set; }

        public WebAclRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchGetArgs()
        {
        }
    }
}
