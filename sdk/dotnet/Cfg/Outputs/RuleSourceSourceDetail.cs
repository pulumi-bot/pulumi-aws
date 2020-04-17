// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cfg.Outputs
{

    [OutputType]
    public sealed class RuleSourceSourceDetail
    {
        /// <summary>
        /// The source of the event, such as an AWS service, that triggers AWS Config
        /// to evaluate your AWS resources. This defaults to `aws.config` and is the only valid value.
        /// is triggered periodically. If specified, requires `message_type` to be `ScheduledNotification`.
        /// </summary>
        public readonly string? EventSource;
        /// <summary>
        /// The maximum frequency with which AWS Config runs evaluations for a rule.
        /// </summary>
        public readonly string? MaximumExecutionFrequency;
        /// <summary>
        /// The type of notification that triggers AWS Config to run an evaluation for a rule. You can specify the following notification types:
        /// </summary>
        public readonly string? MessageType;

        [OutputConstructor]
        private RuleSourceSourceDetail(
            string? eventSource,

            string? maximumExecutionFrequency,

            string? messageType)
        {
            EventSource = eventSource;
            MaximumExecutionFrequency = maximumExecutionFrequency;
            MessageType = messageType;
        }
    }
}
