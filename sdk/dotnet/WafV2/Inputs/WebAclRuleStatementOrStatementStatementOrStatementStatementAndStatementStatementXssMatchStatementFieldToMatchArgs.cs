// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementOrStatementStatementOrStatementStatementAndStatementStatementXssMatchStatementFieldToMatchArgs : Pulumi.ResourceArgs
    {
        [Input("allQueryArguments")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementAndStatementStatementXssMatchStatementFieldToMatchAllQueryArgumentsArgs>? AllQueryArguments { get; set; }

        [Input("body")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementAndStatementStatementXssMatchStatementFieldToMatchBodyArgs>? Body { get; set; }

        [Input("method")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementAndStatementStatementXssMatchStatementFieldToMatchMethodArgs>? Method { get; set; }

        [Input("queryString")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementAndStatementStatementXssMatchStatementFieldToMatchQueryStringArgs>? QueryString { get; set; }

        [Input("singleHeader")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementAndStatementStatementXssMatchStatementFieldToMatchSingleHeaderArgs>? SingleHeader { get; set; }

        [Input("singleQueryArgument")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementAndStatementStatementXssMatchStatementFieldToMatchSingleQueryArgumentArgs>? SingleQueryArgument { get; set; }

        [Input("uriPath")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementAndStatementStatementXssMatchStatementFieldToMatchUriPathArgs>? UriPath { get; set; }

        public WebAclRuleStatementOrStatementStatementOrStatementStatementAndStatementStatementXssMatchStatementFieldToMatchArgs()
        {
        }
    }
}
