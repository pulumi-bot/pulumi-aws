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
    public sealed class WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatement
    {
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementByteMatchStatement? ByteMatchStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementGeoMatchStatement? GeoMatchStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementIpSetReferenceStatement? IpSetReferenceStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatement? RegexPatternSetReferenceStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementSizeConstraintStatement? SizeConstraintStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementSqliMatchStatement? SqliMatchStatement;
        public readonly Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatement? XssMatchStatement;

        [OutputConstructor]
        private WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatement(
            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementByteMatchStatement? byteMatchStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementGeoMatchStatement? geoMatchStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementIpSetReferenceStatement? ipSetReferenceStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatement? regexPatternSetReferenceStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementSizeConstraintStatement? sizeConstraintStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementSqliMatchStatement? sqliMatchStatement,

            Outputs.WebAclRuleStatementOrStatementStatementNotStatementStatementNotStatementStatementXssMatchStatement? xssMatchStatement)
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
