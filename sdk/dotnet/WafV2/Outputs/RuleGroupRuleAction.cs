// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Outputs
{

    [OutputType]
    public sealed class RuleGroupRuleAction
    {
        public readonly Outputs.RuleGroupRuleActionAllow? Allow;
        public readonly Outputs.RuleGroupRuleActionBlock? Block;
        public readonly Outputs.RuleGroupRuleActionCount? Count;

        [OutputConstructor]
        private RuleGroupRuleAction(
            Outputs.RuleGroupRuleActionAllow? allow,

            Outputs.RuleGroupRuleActionBlock? block,

            Outputs.RuleGroupRuleActionCount? count)
        {
            Allow = allow;
            Block = block;
            Count = count;
        }
    }
}
