// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchArgs : Pulumi.ResourceArgs
    {
        [Input("allQueryArguments")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsArgs>? AllQueryArguments { get; set; }

        [Input("body")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBodyArgs>? Body { get; set; }

        [Input("method")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethodArgs>? Method { get; set; }

        [Input("queryString")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringArgs>? QueryString { get; set; }

        [Input("singleHeader")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderArgs>? SingleHeader { get; set; }

        [Input("singleQueryArgument")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentArgs>? SingleQueryArgument { get; set; }

        [Input("uriPath")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPathArgs>? UriPath { get; set; }

        public WebAclRuleStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchArgs()
        {
        }
    }
}
