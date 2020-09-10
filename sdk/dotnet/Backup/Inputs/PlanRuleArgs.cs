// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Backup.Inputs
{

    public sealed class PlanRuleArgs : Pulumi.ResourceArgs
    {
        [Input("completionWindow")]
        public Input<int>? CompletionWindow { get; set; }

        [Input("copyActions")]
        private InputList<Inputs.PlanRuleCopyActionArgs>? _copyActions;
        public InputList<Inputs.PlanRuleCopyActionArgs> CopyActions
        {
            get => _copyActions ?? (_copyActions = new InputList<Inputs.PlanRuleCopyActionArgs>());
            set => _copyActions = value;
        }

        [Input("lifecycle")]
        public Input<Inputs.PlanRuleLifecycleArgs>? Lifecycle { get; set; }

        [Input("recoveryPointTags")]
        private InputMap<string>? _recoveryPointTags;
        public InputMap<string> RecoveryPointTags
        {
            get => _recoveryPointTags ?? (_recoveryPointTags = new InputMap<string>());
            set => _recoveryPointTags = value;
        }

        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        [Input("startWindow")]
        public Input<int>? StartWindow { get; set; }

        [Input("targetVaultName", required: true)]
        public Input<string> TargetVaultName { get; set; } = null!;

        public PlanRuleArgs()
        {
        }
    }
}
