// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Backup.Outputs
{

    [OutputType]
    public sealed class PlanRule
    {
        public readonly int? CompletionWindow;
        public readonly ImmutableArray<Outputs.PlanRuleCopyAction> CopyActions;
        public readonly Outputs.PlanRuleLifecycle? Lifecycle;
        public readonly ImmutableDictionary<string, string>? RecoveryPointTags;
        public readonly string RuleName;
        public readonly string? Schedule;
        public readonly int? StartWindow;
        public readonly string TargetVaultName;

        [OutputConstructor]
        private PlanRule(
            int? completionWindow,

            ImmutableArray<Outputs.PlanRuleCopyAction> copyActions,

            Outputs.PlanRuleLifecycle? lifecycle,

            ImmutableDictionary<string, string>? recoveryPointTags,

            string ruleName,

            string? schedule,

            int? startWindow,

            string targetVaultName)
        {
            CompletionWindow = completionWindow;
            CopyActions = copyActions;
            Lifecycle = lifecycle;
            RecoveryPointTags = recoveryPointTags;
            RuleName = ruleName;
            Schedule = schedule;
            StartWindow = startWindow;
            TargetVaultName = targetVaultName;
        }
    }
}
