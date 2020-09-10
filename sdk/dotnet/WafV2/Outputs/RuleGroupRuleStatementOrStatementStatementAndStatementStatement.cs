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
    public sealed class RuleGroupRuleStatementOrStatementStatementAndStatementStatement
    {
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementByteMatchStatement? ByteMatchStatement;
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementGeoMatchStatement? GeoMatchStatement;
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementIpSetReferenceStatement? IpSetReferenceStatement;
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatement? RegexPatternSetReferenceStatement;
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementSizeConstraintStatement? SizeConstraintStatement;
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementSqliMatchStatement? SqliMatchStatement;
        public readonly Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementXssMatchStatement? XssMatchStatement;

        [OutputConstructor]
        private RuleGroupRuleStatementOrStatementStatementAndStatementStatement(
            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementByteMatchStatement? byteMatchStatement,

            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementGeoMatchStatement? geoMatchStatement,

            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementIpSetReferenceStatement? ipSetReferenceStatement,

            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatement? regexPatternSetReferenceStatement,

            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementSizeConstraintStatement? sizeConstraintStatement,

            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementSqliMatchStatement? sqliMatchStatement,

            Outputs.RuleGroupRuleStatementOrStatementStatementAndStatementStatementXssMatchStatement? xssMatchStatement)
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
