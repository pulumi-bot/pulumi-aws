// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iot.Inputs
{

    public sealed class TopicRuleSnsArgs : Pulumi.ResourceArgs
    {
        [Input("messageFormat")]
        public Input<string>? MessageFormat { get; set; }

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("targetArn", required: true)]
        public Input<string> TargetArn { get; set; } = null!;

        public TopicRuleSnsArgs()
        {
        }
    }
}
