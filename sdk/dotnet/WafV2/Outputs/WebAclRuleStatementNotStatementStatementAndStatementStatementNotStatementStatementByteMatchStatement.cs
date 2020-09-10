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
    public sealed class WebAclRuleStatementNotStatementStatementAndStatementStatementNotStatementStatementByteMatchStatement
    {
        public readonly Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementNotStatementStatementByteMatchStatementFieldToMatch? FieldToMatch;
        public readonly string PositionalConstraint;
        public readonly string SearchString;
        public readonly ImmutableArray<Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementNotStatementStatementByteMatchStatementTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebAclRuleStatementNotStatementStatementAndStatementStatementNotStatementStatementByteMatchStatement(
            Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementNotStatementStatementByteMatchStatementFieldToMatch? fieldToMatch,

            string positionalConstraint,

            string searchString,

            ImmutableArray<Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementNotStatementStatementByteMatchStatementTextTransformation> textTransformations)
        {
            FieldToMatch = fieldToMatch;
            PositionalConstraint = positionalConstraint;
            SearchString = searchString;
            TextTransformations = textTransformations;
        }
    }
}
