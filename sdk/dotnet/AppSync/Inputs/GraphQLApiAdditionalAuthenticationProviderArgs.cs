// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppSync.Inputs
{

    public sealed class GraphQLApiAdditionalAuthenticationProviderArgs : Pulumi.ResourceArgs
    {
        [Input("authenticationType", required: true)]
        public Input<string> AuthenticationType { get; set; } = null!;

        [Input("openidConnectConfig")]
        public Input<Inputs.GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigArgs>? OpenidConnectConfig { get; set; }

        [Input("userPoolConfig")]
        public Input<Inputs.GraphQLApiAdditionalAuthenticationProviderUserPoolConfigArgs>? UserPoolConfig { get; set; }

        public GraphQLApiAdditionalAuthenticationProviderArgs()
        {
        }
    }
}
