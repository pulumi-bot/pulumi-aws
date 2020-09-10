// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53.Inputs
{

    public sealed class RecordGeolocationRoutingPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("continent")]
        public Input<string>? Continent { get; set; }

        [Input("country")]
        public Input<string>? Country { get; set; }

        [Input("subdivision")]
        public Input<string>? Subdivision { get; set; }

        public RecordGeolocationRoutingPolicyArgs()
        {
        }
    }
}
