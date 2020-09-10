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
    public sealed class RuleGroupRuleStatementAndStatementStatementSizeConstraintStatement
    {
        public readonly string ComparisonOperator;
        public readonly Outputs.RuleGroupRuleStatementAndStatementStatementSizeConstraintStatementFieldToMatch? FieldToMatch;
        public readonly int Size;
        public readonly ImmutableArray<Outputs.RuleGroupRuleStatementAndStatementStatementSizeConstraintStatementTextTransformation> TextTransformations;

        [OutputConstructor]
        private RuleGroupRuleStatementAndStatementStatementSizeConstraintStatement(
            string comparisonOperator,

            Outputs.RuleGroupRuleStatementAndStatementStatementSizeConstraintStatementFieldToMatch? fieldToMatch,

            int size,

            ImmutableArray<Outputs.RuleGroupRuleStatementAndStatementStatementSizeConstraintStatementTextTransformation> textTransformations)
        {
            ComparisonOperator = comparisonOperator;
            FieldToMatch = fieldToMatch;
            Size = size;
            TextTransformations = textTransformations;
        }
    }
}
