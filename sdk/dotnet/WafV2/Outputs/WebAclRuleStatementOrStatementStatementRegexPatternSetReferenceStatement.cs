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
    public sealed class WebAclRuleStatementOrStatementStatementRegexPatternSetReferenceStatement
    {
        public readonly string Arn;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementRegexPatternSetReferenceStatementFieldToMatch? FieldToMatch;
        public readonly ImmutableArray<Outputs.WebAclRuleStatementOrStatementStatementRegexPatternSetReferenceStatementTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebAclRuleStatementOrStatementStatementRegexPatternSetReferenceStatement(
            string arn,

            Outputs.WebAclRuleStatementOrStatementStatementRegexPatternSetReferenceStatementFieldToMatch? fieldToMatch,

            ImmutableArray<Outputs.WebAclRuleStatementOrStatementStatementRegexPatternSetReferenceStatementTextTransformation> textTransformations)
        {
            Arn = arn;
            FieldToMatch = fieldToMatch;
            TextTransformations = textTransformations;
        }
    }
}
