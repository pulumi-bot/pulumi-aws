// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementNotStatementStatementArgs : Pulumi.ResourceArgs
    {
        [Input("byteMatchStatement")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementNotStatementStatementByteMatchStatementArgs>? ByteMatchStatement { get; set; }

        [Input("geoMatchStatement")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementNotStatementStatementGeoMatchStatementArgs>? GeoMatchStatement { get; set; }

        [Input("ipSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementNotStatementStatementIpSetReferenceStatementArgs>? IpSetReferenceStatement { get; set; }

        [Input("regexPatternSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementNotStatementStatementRegexPatternSetReferenceStatementArgs>? RegexPatternSetReferenceStatement { get; set; }

        [Input("sizeConstraintStatement")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementNotStatementStatementSizeConstraintStatementArgs>? SizeConstraintStatement { get; set; }

        [Input("sqliMatchStatement")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementNotStatementStatementSqliMatchStatementArgs>? SqliMatchStatement { get; set; }

        [Input("xssMatchStatement")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementNotStatementStatementXssMatchStatementArgs>? XssMatchStatement { get; set; }

        public WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementNotStatementStatementArgs()
        {
        }
    }
}
