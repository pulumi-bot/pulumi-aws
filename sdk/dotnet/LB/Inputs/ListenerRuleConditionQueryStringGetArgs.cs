// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LB.Inputs
{

    public sealed class ListenerRuleConditionQueryStringGetArgs : Pulumi.ResourceArgs
    {
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ListenerRuleConditionQueryStringGetArgs()
        {
        }
    }
}
