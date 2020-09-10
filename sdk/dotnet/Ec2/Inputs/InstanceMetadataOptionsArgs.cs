// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Inputs
{

    public sealed class InstanceMetadataOptionsArgs : Pulumi.ResourceArgs
    {
        [Input("httpEndpoint")]
        public Input<string>? HttpEndpoint { get; set; }

        [Input("httpPutResponseHopLimit")]
        public Input<int>? HttpPutResponseHopLimit { get; set; }

        [Input("httpTokens")]
        public Input<string>? HttpTokens { get; set; }

        public InstanceMetadataOptionsArgs()
        {
        }
    }
}
