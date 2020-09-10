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
    public sealed class WebAclRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatement
    {
        public readonly string ComparisonOperator;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatch? FieldToMatch;
        public readonly int Size;
        public readonly ImmutableArray<Outputs.WebAclRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebAclRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatement(
            string comparisonOperator,

            Outputs.WebAclRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatch? fieldToMatch,

            int size,

            ImmutableArray<Outputs.WebAclRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementTextTransformation> textTransformations)
        {
            ComparisonOperator = comparisonOperator;
            FieldToMatch = fieldToMatch;
            Size = size;
            TextTransformations = textTransformations;
        }
    }
}
