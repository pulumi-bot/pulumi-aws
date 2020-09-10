// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticLoadBalancingV2.Inputs
{

    public sealed class ListenerDefaultActionAuthenticateCognitoArgs : Pulumi.ResourceArgs
    {
        [Input("authenticationRequestExtraParams")]
        private InputMap<string>? _authenticationRequestExtraParams;
        public InputMap<string> AuthenticationRequestExtraParams
        {
            get => _authenticationRequestExtraParams ?? (_authenticationRequestExtraParams = new InputMap<string>());
            set => _authenticationRequestExtraParams = value;
        }

        [Input("onUnauthenticatedRequest")]
        public Input<string>? OnUnauthenticatedRequest { get; set; }

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("sessionCookieName")]
        public Input<string>? SessionCookieName { get; set; }

        [Input("sessionTimeout")]
        public Input<int>? SessionTimeout { get; set; }

        [Input("userPoolArn", required: true)]
        public Input<string> UserPoolArn { get; set; } = null!;

        [Input("userPoolClientId", required: true)]
        public Input<string> UserPoolClientId { get; set; } = null!;

        [Input("userPoolDomain", required: true)]
        public Input<string> UserPoolDomain { get; set; } = null!;

        public ListenerDefaultActionAuthenticateCognitoArgs()
        {
        }
    }
}
