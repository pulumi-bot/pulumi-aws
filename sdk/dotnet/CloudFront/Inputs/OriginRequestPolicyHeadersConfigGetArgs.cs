// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Inputs
{

    public sealed class OriginRequestPolicyHeadersConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("headerBehavior")]
        public Input<string>? HeaderBehavior { get; set; }

        [Input("headers")]
        public Input<Inputs.OriginRequestPolicyHeadersConfigHeadersGetArgs>? Headers { get; set; }

        public OriginRequestPolicyHeadersConfigGetArgs()
        {
        }
    }
}
