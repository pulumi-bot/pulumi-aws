// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class RuleGroupRuleStatementNotStatementStatementAndStatementStatementSizeConstraintStatementTextTransformationGetArgs : Pulumi.ResourceArgs
    {
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public RuleGroupRuleStatementNotStatementStatementAndStatementStatementSizeConstraintStatementTextTransformationGetArgs()
        {
        }
    }
}
