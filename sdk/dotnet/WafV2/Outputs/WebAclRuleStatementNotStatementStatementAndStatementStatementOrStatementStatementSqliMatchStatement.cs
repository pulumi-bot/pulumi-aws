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
    public sealed class WebAclRuleStatementNotStatementStatementAndStatementStatementOrStatementStatementSqliMatchStatement
    {
        public readonly Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementOrStatementStatementSqliMatchStatementFieldToMatch? FieldToMatch;
        public readonly ImmutableArray<Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementOrStatementStatementSqliMatchStatementTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebAclRuleStatementNotStatementStatementAndStatementStatementOrStatementStatementSqliMatchStatement(
            Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementOrStatementStatementSqliMatchStatementFieldToMatch? fieldToMatch,

            ImmutableArray<Outputs.WebAclRuleStatementNotStatementStatementAndStatementStatementOrStatementStatementSqliMatchStatementTextTransformation> textTransformations)
        {
            FieldToMatch = fieldToMatch;
            TextTransformations = textTransformations;
        }
    }
}
