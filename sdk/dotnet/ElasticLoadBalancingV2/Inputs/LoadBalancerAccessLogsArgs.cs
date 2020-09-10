// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticLoadBalancingV2.Inputs
{

    public sealed class LoadBalancerAccessLogsArgs : Pulumi.ResourceArgs
    {
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public LoadBalancerAccessLogsArgs()
        {
        }
    }
}
