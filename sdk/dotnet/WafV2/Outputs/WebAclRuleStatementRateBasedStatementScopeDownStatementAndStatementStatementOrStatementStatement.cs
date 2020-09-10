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
    public sealed class WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatement
    {
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementByteMatchStatement? ByteMatchStatement;
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementGeoMatchStatement? GeoMatchStatement;
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementIpSetReferenceStatement? IpSetReferenceStatement;
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementRegexPatternSetReferenceStatement? RegexPatternSetReferenceStatement;
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementSizeConstraintStatement? SizeConstraintStatement;
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementSqliMatchStatement? SqliMatchStatement;
        public readonly Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementXssMatchStatement? XssMatchStatement;

        [OutputConstructor]
        private WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatement(
            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementByteMatchStatement? byteMatchStatement,

            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementGeoMatchStatement? geoMatchStatement,

            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementIpSetReferenceStatement? ipSetReferenceStatement,

            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementRegexPatternSetReferenceStatement? regexPatternSetReferenceStatement,

            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementSizeConstraintStatement? sizeConstraintStatement,

            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementSqliMatchStatement? sqliMatchStatement,

            Outputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementOrStatementStatementXssMatchStatement? xssMatchStatement)
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
