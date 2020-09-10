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
    public sealed class RuleGroupRuleStatementAndStatementStatementOrStatementStatement
    {
        public readonly Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementByteMatchStatement? ByteMatchStatement;
        public readonly Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementGeoMatchStatement? GeoMatchStatement;
        public readonly Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementIpSetReferenceStatement? IpSetReferenceStatement;
        public readonly Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementRegexPatternSetReferenceStatement? RegexPatternSetReferenceStatement;
        public readonly Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatement? SizeConstraintStatement;
        public readonly Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementSqliMatchStatement? SqliMatchStatement;
        public readonly Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementXssMatchStatement? XssMatchStatement;

        [OutputConstructor]
        private RuleGroupRuleStatementAndStatementStatementOrStatementStatement(
            Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementByteMatchStatement? byteMatchStatement,

            Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementGeoMatchStatement? geoMatchStatement,

            Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementIpSetReferenceStatement? ipSetReferenceStatement,

            Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementRegexPatternSetReferenceStatement? regexPatternSetReferenceStatement,

            Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatement? sizeConstraintStatement,

            Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementSqliMatchStatement? sqliMatchStatement,

            Outputs.RuleGroupRuleStatementAndStatementStatementOrStatementStatementXssMatchStatement? xssMatchStatement)
        {
            ByteMatchStatement = byteMatchStatement;
            GeoMatchStatement = geoMatchStatement;
            IpSetReferenceStatement = ipSetReferenceStatement;
            RegexPatternSetReferenceStatement = regexPatternSetReferenceStatement;
            SizeConstraintStatement = sizeConstraintStatement;
            SqliMatchStatement = sqliMatchStatement;
            XssMatchStatement = xssMatchStatement;
        }
    }
}
