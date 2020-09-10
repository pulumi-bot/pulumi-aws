// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Msk.Inputs
{

    public sealed class ClusterLoggingInfoBrokerLogsS3Args : Pulumi.ResourceArgs
    {
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public ClusterLoggingInfoBrokerLogsS3Args()
        {
        }
    }
}
