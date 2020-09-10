// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementOrStatementStatementOrStatementStatementOrStatementStatementArgs : Pulumi.ResourceArgs
    {
        [Input("byteMatchStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementOrStatementStatementByteMatchStatementArgs>? ByteMatchStatement { get; set; }

        [Input("geoMatchStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementOrStatementStatementGeoMatchStatementArgs>? GeoMatchStatement { get; set; }

        [Input("ipSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementOrStatementStatementIpSetReferenceStatementArgs>? IpSetReferenceStatement { get; set; }

        [Input("regexPatternSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementOrStatementStatementRegexPatternSetReferenceStatementArgs>? RegexPatternSetReferenceStatement { get; set; }

        [Input("sizeConstraintStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementArgs>? SizeConstraintStatement { get; set; }

        [Input("sqliMatchStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementOrStatementStatementSqliMatchStatementArgs>? SqliMatchStatement { get; set; }

        [Input("xssMatchStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementOrStatementStatementXssMatchStatementArgs>? XssMatchStatement { get; set; }

        public WebAclRuleStatementOrStatementStatementOrStatementStatementOrStatementStatementArgs()
        {
        }
    }
}
