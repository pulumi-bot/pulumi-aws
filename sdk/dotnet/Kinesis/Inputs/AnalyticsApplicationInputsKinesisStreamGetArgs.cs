// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis.Inputs
{

    public sealed class AnalyticsApplicationInputsKinesisStreamGetArgs : Pulumi.ResourceArgs
    {
        [Input("resourceArn", required: true)]
        public Input<string> ResourceArn { get; set; } = null!;

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        public AnalyticsApplicationInputsKinesisStreamGetArgs()
        {
        }
    }
}
