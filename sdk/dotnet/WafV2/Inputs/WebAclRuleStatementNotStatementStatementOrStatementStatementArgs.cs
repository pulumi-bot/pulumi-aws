// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementNotStatementStatementOrStatementStatementArgs : Pulumi.ResourceArgs
    {
        [Input("andStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementAndStatementArgs>? AndStatement { get; set; }

        [Input("byteMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementByteMatchStatementArgs>? ByteMatchStatement { get; set; }

        [Input("geoMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementGeoMatchStatementArgs>? GeoMatchStatement { get; set; }

        [Input("ipSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementIpSetReferenceStatementArgs>? IpSetReferenceStatement { get; set; }

        [Input("notStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementNotStatementArgs>? NotStatement { get; set; }

        [Input("orStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementArgs>? OrStatement { get; set; }

        [Input("regexPatternSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementRegexPatternSetReferenceStatementArgs>? RegexPatternSetReferenceStatement { get; set; }

        [Input("sizeConstraintStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementArgs>? SizeConstraintStatement { get; set; }

        [Input("sqliMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementSqliMatchStatementArgs>? SqliMatchStatement { get; set; }

        [Input("xssMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementXssMatchStatementArgs>? XssMatchStatement { get; set; }

        public WebAclRuleStatementNotStatementStatementOrStatementStatementArgs()
        {
        }
    }
}
