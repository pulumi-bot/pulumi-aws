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
    public sealed class WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatement
    {
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatement? ByteMatchStatement;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementGeoMatchStatement? GeoMatchStatement;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementIpSetReferenceStatement? IpSetReferenceStatement;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementRegexPatternSetReferenceStatement? RegexPatternSetReferenceStatement;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementSizeConstraintStatement? SizeConstraintStatement;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementSqliMatchStatement? SqliMatchStatement;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementXssMatchStatement? XssMatchStatement;

        [OutputConstructor]
        private WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatement(
            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementByteMatchStatement? byteMatchStatement,

            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementGeoMatchStatement? geoMatchStatement,

            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementIpSetReferenceStatement? ipSetReferenceStatement,

            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementRegexPatternSetReferenceStatement? regexPatternSetReferenceStatement,

            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementSizeConstraintStatement? sizeConstraintStatement,

            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementSqliMatchStatement? sqliMatchStatement,

            Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementXssMatchStatement? xssMatchStatement)
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
