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
    public sealed class WebAclRuleStatementOrStatementStatementNotStatementStatement
    {
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatement? AndStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementByteMatchStatement? ByteMatchStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementGeoMatchStatement? GeoMatchStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementIpSetReferenceStatement? IpSetReferenceStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatement? NotStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementOrStatement? OrStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementRegexPatternSetReferenceStatement? RegexPatternSetReferenceStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementSizeConstraintStatement? SizeConstraintStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatement? SqliMatchStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatement? XssMatchStatement;

        [OutputConstructor]
        private WebAclRuleStatementOrStatementStatementNotStatementStatement(
            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementAndStatement? andStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementByteMatchStatement? byteMatchStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementGeoMatchStatement? geoMatchStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementIpSetReferenceStatement? ipSetReferenceStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatement? notStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementOrStatement? orStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementRegexPatternSetReferenceStatement? regexPatternSetReferenceStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementSizeConstraintStatement? sizeConstraintStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatement? sqliMatchStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementXssMatchStatement? xssMatchStatement)
        {
            AndStatement = andStatement;
            ByteMatchStatement = byteMatchStatement;
            GeoMatchStatement = geoMatchStatement;
            IpSetReferenceStatement = ipSetReferenceStatement;
            NotStatement = notStatement;
            OrStatement = orStatement;
            RegexPatternSetReferenceStatement = regexPatternSetReferenceStatement;
            SizeConstraintStatement = sizeConstraintStatement;
            SqliMatchStatement = sqliMatchStatement;
            XssMatchStatement = xssMatchStatement;
        }
    }
}
