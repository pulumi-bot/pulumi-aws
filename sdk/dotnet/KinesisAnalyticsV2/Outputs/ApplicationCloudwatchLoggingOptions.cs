// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.KinesisAnalyticsV2.Outputs
{

    [OutputType]
    public sealed class ApplicationCloudwatchLoggingOptions
    {
        public readonly string? CloudwatchLoggingOptionId;
        /// <summary>
        /// The ARN of the CloudWatch log stream to receive application messages.
        /// </summary>
        public readonly string LogStreamArn;

        [OutputConstructor]
        private ApplicationCloudwatchLoggingOptions(
            string? cloudwatchLoggingOptionId,

            string logStreamArn)
        {
            CloudwatchLoggingOptionId = cloudwatchLoggingOptionId;
            LogStreamArn = logStreamArn;
        }
    }
}
