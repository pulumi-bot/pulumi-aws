// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class RuleGroupRuleStatementOrStatementStatementAndStatementArgs : Pulumi.ResourceArgs
    {
        [Input("statements", required: true)]
        private InputList<Inputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementArgs>? _statements;
        public InputList<Inputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementArgs> Statements
        {
            get => _statements ?? (_statements = new InputList<Inputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementArgs>());
            set => _statements = value;
        }

        public RuleGroupRuleStatementOrStatementStatementAndStatementArgs()
        {
        }
    }
}
