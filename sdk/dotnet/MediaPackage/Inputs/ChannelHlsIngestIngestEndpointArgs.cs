// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaPackage.Inputs
{

    public sealed class ChannelHlsIngestIngestEndpointArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public ChannelHlsIngestIngestEndpointArgs()
        {
        }
    }
}
