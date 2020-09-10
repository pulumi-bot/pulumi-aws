// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GameLift.Inputs
{

    public sealed class GameSessionQueuePlayerLatencyPolicyGetArgs : Pulumi.ResourceArgs
    {
        [Input("maximumIndividualPlayerLatencyMilliseconds", required: true)]
        public Input<int> MaximumIndividualPlayerLatencyMilliseconds { get; set; } = null!;

        [Input("policyDurationSeconds")]
        public Input<int>? PolicyDurationSeconds { get; set; }

        public GameSessionQueuePlayerLatencyPolicyGetArgs()
        {
        }
    }
}
