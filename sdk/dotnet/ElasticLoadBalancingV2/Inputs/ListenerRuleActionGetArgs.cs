// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticLoadBalancingV2.Inputs
{

    public sealed class ListenerRuleActionGetArgs : Pulumi.ResourceArgs
    {
        [Input("authenticateCognito")]
        public Input<Inputs.ListenerRuleActionAuthenticateCognitoGetArgs>? AuthenticateCognito { get; set; }

        [Input("authenticateOidc")]
        public Input<Inputs.ListenerRuleActionAuthenticateOidcGetArgs>? AuthenticateOidc { get; set; }

        [Input("fixedResponse")]
        public Input<Inputs.ListenerRuleActionFixedResponseGetArgs>? FixedResponse { get; set; }

        [Input("forward")]
        public Input<Inputs.ListenerRuleActionForwardGetArgs>? Forward { get; set; }

        [Input("order")]
        public Input<int>? Order { get; set; }

        [Input("redirect")]
        public Input<Inputs.ListenerRuleActionRedirectGetArgs>? Redirect { get; set; }

        [Input("targetGroupArn")]
        public Input<string>? TargetGroupArn { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ListenerRuleActionGetArgs()
        {
        }
    }
}
