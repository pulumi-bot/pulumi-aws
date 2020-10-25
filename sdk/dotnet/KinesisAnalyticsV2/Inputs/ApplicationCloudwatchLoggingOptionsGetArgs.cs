// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.KinesisAnalyticsV2.Inputs
{

    public sealed class ApplicationCloudwatchLoggingOptionsGetArgs : Pulumi.ResourceArgs
    {
        [Input("cloudwatchLoggingOptionId")]
        public Input<string>? CloudwatchLoggingOptionId { get; set; }

        /// <summary>
        /// The ARN of the CloudWatch log stream to receive application messages.
        /// </summary>
        [Input("logStreamArn", required: true)]
        public Input<string> LogStreamArn { get; set; } = null!;

        public ApplicationCloudwatchLoggingOptionsGetArgs()
        {
        }
    }
}
