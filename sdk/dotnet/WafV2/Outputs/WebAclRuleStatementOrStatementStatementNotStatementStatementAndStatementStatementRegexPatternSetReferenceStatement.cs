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
    public sealed class WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatement
    {
        public readonly string Arn;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch? FieldToMatch;
        public readonly ImmutableArray<Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatement(
            string arn,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch? fieldToMatch,

            ImmutableArray<Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementTextTransformation> textTransformations)
        {
            Arn = arn;
            FieldToMatch = fieldToMatch;
            TextTransformations = textTransformations;
        }
    }
}
