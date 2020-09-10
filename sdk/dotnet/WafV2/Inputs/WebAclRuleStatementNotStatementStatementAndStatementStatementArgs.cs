// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementNotStatementStatementAndStatementStatementArgs : Pulumi.ResourceArgs
    {
        [Input("andStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementArgs>? AndStatement { get; set; }

        [Input("byteMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementArgs>? ByteMatchStatement { get; set; }

        [Input("geoMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementGeoMatchStatementArgs>? GeoMatchStatement { get; set; }

        [Input("ipSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementIpSetReferenceStatementArgs>? IpSetReferenceStatement { get; set; }

        [Input("notStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementNotStatementArgs>? NotStatement { get; set; }

        [Input("orStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementOrStatementArgs>? OrStatement { get; set; }

        [Input("regexPatternSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementArgs>? RegexPatternSetReferenceStatement { get; set; }

        [Input("sizeConstraintStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementSizeConstraintStatementArgs>? SizeConstraintStatement { get; set; }

        [Input("sqliMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementSqliMatchStatementArgs>? SqliMatchStatement { get; set; }

        [Input("xssMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementXssMatchStatementArgs>? XssMatchStatement { get; set; }

        public WebAclRuleStatementNotStatementStatementAndStatementStatementArgs()
        {
        }
    }
}
