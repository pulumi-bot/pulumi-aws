// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Dlm.Outputs
{

    [OutputType]
    public sealed class LifecyclePolicyPolicyDetailsScheduleCreateRule
    {
        public readonly int Interval;
        public readonly string? IntervalUnit;
        public readonly string? Times;

        [OutputConstructor]
        private LifecyclePolicyPolicyDetailsScheduleCreateRule(
            int interval,

            string? intervalUnit,

            string? times)
        {
            Interval = interval;
            IntervalUnit = intervalUnit;
            Times = times;
        }
    }
}
