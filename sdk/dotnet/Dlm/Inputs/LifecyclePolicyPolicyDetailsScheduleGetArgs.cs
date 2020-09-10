// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Dlm.Inputs
{

    public sealed class LifecyclePolicyPolicyDetailsScheduleGetArgs : Pulumi.ResourceArgs
    {
        [Input("copyTags")]
        public Input<bool>? CopyTags { get; set; }

        [Input("createRule", required: true)]
        public Input<Inputs.LifecyclePolicyPolicyDetailsScheduleCreateRuleGetArgs> CreateRule { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("retainRule", required: true)]
        public Input<Inputs.LifecyclePolicyPolicyDetailsScheduleRetainRuleGetArgs> RetainRule { get; set; } = null!;

        [Input("tagsToAdd")]
        private InputMap<string>? _tagsToAdd;
        public InputMap<string> TagsToAdd
        {
            get => _tagsToAdd ?? (_tagsToAdd = new InputMap<string>());
            set => _tagsToAdd = value;
        }

        public LifecyclePolicyPolicyDetailsScheduleGetArgs()
        {
        }
    }
}
