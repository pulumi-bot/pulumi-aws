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
    public sealed class WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatement
    {
        public readonly Outputs.WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatementFieldToMatch? FieldToMatch;
        public readonly ImmutableArray<Outputs.WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatementTextTransformation> TextTransformations;

        [OutputConstructor]
        private WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatement(
            Outputs.WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatementFieldToMatch? fieldToMatch,

            ImmutableArray<Outputs.WebAclRuleStatementOrStatementStatementAndStatementStatementNotStatementStatementXssMatchStatementTextTransformation> textTransformations)
        {
            FieldToMatch = fieldToMatch;
            TextTransformations = textTransformations;
        }
    }
}
