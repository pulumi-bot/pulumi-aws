// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kinesis.Outputs
{

    [OutputType]
    public sealed class FirehoseDeliveryStreamS3ConfigurationCloudwatchLoggingOptions
    {
        public readonly bool? Enabled;
        public readonly string? LogGroupName;
        public readonly string? LogStreamName;

        [OutputConstructor]
        private FirehoseDeliveryStreamS3ConfigurationCloudwatchLoggingOptions(
            bool? enabled,

            string? logGroupName,

            string? logStreamName)
        {
            Enabled = enabled;
            LogGroupName = logGroupName;
            LogStreamName = logStreamName;
        }
    }
}
