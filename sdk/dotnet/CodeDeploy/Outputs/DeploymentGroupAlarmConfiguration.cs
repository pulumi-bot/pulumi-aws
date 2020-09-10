// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeDeploy.Outputs
{

    [OutputType]
    public sealed class DeploymentGroupAlarmConfiguration
    {
        public readonly ImmutableArray<string> Alarms;
        public readonly bool? Enabled;
        public readonly bool? IgnorePollAlarmFailure;

        [OutputConstructor]
        private DeploymentGroupAlarmConfiguration(
            ImmutableArray<string> alarms,

            bool? enabled,

            bool? ignorePollAlarmFailure)
        {
            Alarms = alarms;
            Enabled = enabled;
            IgnorePollAlarmFailure = ignorePollAlarmFailure;
        }
    }
}
