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
    public sealed class WebAclRuleStatementAndStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatement
    {
        public readonly Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatch? FieldToMatch;
        public readonly ImmutableArray<Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebAclRuleStatementAndStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatement(
            Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatch? fieldToMatch,

            ImmutableArray<Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformation> textTransformations)
        {
            FieldToMatch = fieldToMatch;
            TextTransformations = textTransformations;
        }
    }
}
