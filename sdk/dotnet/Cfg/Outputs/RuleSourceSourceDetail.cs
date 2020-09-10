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
        public readonly string? EventSource;
        public readonly string? MaximumExecutionFrequency;
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
