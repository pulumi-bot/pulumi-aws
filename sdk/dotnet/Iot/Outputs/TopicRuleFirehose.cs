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
    public sealed class TopicRuleFirehose
    {
        public readonly string DeliveryStreamName;
        public readonly string RoleArn;
        public readonly string? Separator;

        [OutputConstructor]
        private TopicRuleFirehose(
            string deliveryStreamName,

            string roleArn,

            string? separator)
        {
            DeliveryStreamName = deliveryStreamName;
            RoleArn = roleArn;
            Separator = separator;
        }
    }
}
