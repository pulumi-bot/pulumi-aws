// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class RuleGroupRuleStatementNotStatementStatementNotStatementArgs : Pulumi.ResourceArgs
    {
        [Input("statements", required: true)]
        private InputList<Inputs.RuleGroupRuleStatementNotStatementStatementNotStatementStatementArgs>? _statements;
        public InputList<Inputs.RuleGroupRuleStatementNotStatementStatementNotStatementStatementArgs> Statements
        {
            get => _statements ?? (_statements = new InputList<Inputs.RuleGroupRuleStatementNotStatementStatementNotStatementStatementArgs>());
            set => _statements = value;
        }

        public RuleGroupRuleStatementNotStatementStatementNotStatementArgs()
        {
        }
    }
}
