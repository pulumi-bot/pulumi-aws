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
    public sealed class RuleGroupRuleStatementNotStatementStatementAndStatementStatementSqliMatchStatement
    {
        public readonly Outputs.RuleGroupRuleStatementNotStatementStatementAndStatementStatementSqliMatchStatementFieldToMatch? FieldToMatch;
        public readonly ImmutableArray<Outputs.RuleGroupRuleStatementNotStatementStatementAndStatementStatementSqliMatchStatementTextTransformation> TextTransformations;

        [OutputConstructor]
        private RuleGroupRuleStatementNotStatementStatementAndStatementStatementSqliMatchStatement(
            Outputs.RuleGroupRuleStatementNotStatementStatementAndStatementStatementSqliMatchStatementFieldToMatch? fieldToMatch,

            ImmutableArray<Outputs.RuleGroupRuleStatementNotStatementStatementAndStatementStatementSqliMatchStatementTextTransformation> textTransformations)
        {
            FieldToMatch = fieldToMatch;
            TextTransformations = textTransformations;
        }
    }
}
