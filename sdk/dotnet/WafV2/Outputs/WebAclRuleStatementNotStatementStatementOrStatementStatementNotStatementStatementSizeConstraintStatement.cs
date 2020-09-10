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
    public sealed class WebAclRuleStatementNotStatementStatementOrStatementStatementNotStatementStatementSizeConstraintStatement
    {
        public readonly string ComparisonOperator;
        public readonly Outputs.WebAclRuleStatementNotStatementStatementOrStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatch? FieldToMatch;
        public readonly int Size;
        public readonly ImmutableArray<Outputs.WebAclRuleStatementNotStatementStatementOrStatementStatementNotStatementStatementSizeConstraintStatementTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebAclRuleStatementNotStatementStatementOrStatementStatementNotStatementStatementSizeConstraintStatement(
            string comparisonOperator,

            Outputs.WebAclRuleStatementNotStatementStatementOrStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatch? fieldToMatch,

            int size,

            ImmutableArray<Outputs.WebAclRuleStatementNotStatementStatementOrStatementStatementNotStatementStatementSizeConstraintStatementTextTransformation> textTransformations)
        {
            ComparisonOperator = comparisonOperator;
            FieldToMatch = fieldToMatch;
            Size = size;
            TextTransformations = textTransformations;
        }
    }
}
