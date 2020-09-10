// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LB.Inputs
{

    public sealed class ListenerRuleConditionGetArgs : Pulumi.ResourceArgs
    {
        [Input("hostHeader")]
        public Input<Inputs.ListenerRuleConditionHostHeaderGetArgs>? HostHeader { get; set; }

        [Input("httpHeader")]
        public Input<Inputs.ListenerRuleConditionHttpHeaderGetArgs>? HttpHeader { get; set; }

        [Input("httpRequestMethod")]
        public Input<Inputs.ListenerRuleConditionHttpRequestMethodGetArgs>? HttpRequestMethod { get; set; }

        [Input("pathPattern")]
        public Input<Inputs.ListenerRuleConditionPathPatternGetArgs>? PathPattern { get; set; }

        [Input("queryStrings")]
        private InputList<Inputs.ListenerRuleConditionQueryStringGetArgs>? _queryStrings;
        public InputList<Inputs.ListenerRuleConditionQueryStringGetArgs> QueryStrings
        {
            get => _queryStrings ?? (_queryStrings = new InputList<Inputs.ListenerRuleConditionQueryStringGetArgs>());
            set => _queryStrings = value;
        }

        [Input("sourceIp")]
        public Input<Inputs.ListenerRuleConditionSourceIpGetArgs>? SourceIp { get; set; }

        public ListenerRuleConditionGetArgs()
        {
        }
    }
}
