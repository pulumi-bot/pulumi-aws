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
    public sealed class WebAclRuleStatementOrStatementStatementAndStatementStatementXssMatchStatement
    {
        public readonly Outputs.WebAclRuleStatementOrStatementStatementAndStatementStatementXssMatchStatementFieldToMatch? FieldToMatch;
        public readonly ImmutableArray<Outputs.WebAclRuleStatementOrStatementStatementAndStatementStatementXssMatchStatementTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebAclRuleStatementOrStatementStatementAndStatementStatementXssMatchStatement(
            Outputs.WebAclRuleStatementOrStatementStatementAndStatementStatementXssMatchStatementFieldToMatch? fieldToMatch,

            ImmutableArray<Outputs.WebAclRuleStatementOrStatementStatementAndStatementStatementXssMatchStatementTextTransformation> textTransformations)
        {
            FieldToMatch = fieldToMatch;
            TextTransformations = textTransformations;
        }
    }
}
