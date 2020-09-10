// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iot.Outputs
{

    [OutputType]
    public sealed class TopicRuleErrorActionIotEvents
    {
        public readonly string InputName;
        public readonly string? MessageId;
        public readonly string RoleArn;

        [OutputConstructor]
        private TopicRuleErrorActionIotEvents(
            string inputName,

            string? messageId,

            string roleArn)
        {
            InputName = inputName;
            MessageId = messageId;
            RoleArn = roleArn;
        }
    }
}
