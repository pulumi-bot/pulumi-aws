// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pinpoint.Outputs
{

    [OutputType]
    public sealed class AppLimits
    {
        public readonly int? Daily;
        public readonly int? MaximumDuration;
        public readonly int? MessagesPerSecond;
        public readonly int? Total;

        [OutputConstructor]
        private AppLimits(
            int? daily,

            int? maximumDuration,

            int? messagesPerSecond,

            int? total)
        {
            Daily = daily;
            MaximumDuration = maximumDuration;
            MessagesPerSecond = messagesPerSecond;
            Total = total;
        }
    }
}
