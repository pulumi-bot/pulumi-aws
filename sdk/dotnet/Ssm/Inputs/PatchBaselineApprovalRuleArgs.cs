// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ssm.Inputs
{

    public sealed class PatchBaselineApprovalRuleArgs : Pulumi.ResourceArgs
    {
        [Input("approveAfterDays", required: true)]
        public Input<int> ApproveAfterDays { get; set; } = null!;

        [Input("complianceLevel")]
        public Input<string>? ComplianceLevel { get; set; }

        [Input("enableNonSecurity")]
        public Input<bool>? EnableNonSecurity { get; set; }

        [Input("patchFilters", required: true)]
        private InputList<Inputs.PatchBaselineApprovalRulePatchFilterArgs>? _patchFilters;
        public InputList<Inputs.PatchBaselineApprovalRulePatchFilterArgs> PatchFilters
        {
            get => _patchFilters ?? (_patchFilters = new InputList<Inputs.PatchBaselineApprovalRulePatchFilterArgs>());
            set => _patchFilters = value;
        }

        public PatchBaselineApprovalRuleArgs()
        {
        }
    }
}
