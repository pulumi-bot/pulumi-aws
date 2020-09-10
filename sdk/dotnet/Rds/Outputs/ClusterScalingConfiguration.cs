// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds.Outputs
{

    [OutputType]
    public sealed class ClusterScalingConfiguration
    {
        public readonly bool? AutoPause;
        public readonly int? MaxCapacity;
        public readonly int? MinCapacity;
        public readonly int? SecondsUntilAutoPause;
        public readonly string? TimeoutAction;

        [OutputConstructor]
        private ClusterScalingConfiguration(
            bool? autoPause,

            int? maxCapacity,

            int? minCapacity,

            int? secondsUntilAutoPause,

            string? timeoutAction)
        {
            AutoPause = autoPause;
            MaxCapacity = maxCapacity;
            MinCapacity = minCapacity;
            SecondsUntilAutoPause = secondsUntilAutoPause;
            TimeoutAction = timeoutAction;
        }
    }
}
