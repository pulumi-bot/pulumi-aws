// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iot.Inputs
{

    public sealed class TopicRuleFirehoseGetArgs : Pulumi.ResourceArgs
    {
        [Input("deliveryStreamName", required: true)]
        public Input<string> DeliveryStreamName { get; set; } = null!;

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("separator")]
        public Input<string>? Separator { get; set; }

        public TopicRuleFirehoseGetArgs()
        {
        }
    }
}
