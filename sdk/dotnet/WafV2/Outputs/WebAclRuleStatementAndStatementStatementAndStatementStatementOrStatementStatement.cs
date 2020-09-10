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
    public sealed class WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatement
    {
        public readonly Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementByteMatchStatement? ByteMatchStatement;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementGeoMatchStatement? GeoMatchStatement;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementIpSetReferenceStatement? IpSetReferenceStatement;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementRegexPatternSetReferenceStatement? RegexPatternSetReferenceStatement;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementSizeConstraintStatement? SizeConstraintStatement;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementSqliMatchStatement? SqliMatchStatement;
        public readonly Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementXssMatchStatement? XssMatchStatement;

        [OutputConstructor]
        private WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatement(
            Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementByteMatchStatement? byteMatchStatement,

            Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementGeoMatchStatement? geoMatchStatement,

            Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementIpSetReferenceStatement? ipSetReferenceStatement,

            Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementRegexPatternSetReferenceStatement? regexPatternSetReferenceStatement,

            Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementSizeConstraintStatement? sizeConstraintStatement,

            Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementSqliMatchStatement? sqliMatchStatement,

            Outputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementXssMatchStatement? xssMatchStatement)
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
