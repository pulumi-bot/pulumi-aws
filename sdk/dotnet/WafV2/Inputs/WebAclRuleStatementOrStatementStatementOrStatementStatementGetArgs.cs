// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementOrStatementStatementOrStatementStatementGetArgs : Pulumi.ResourceArgs
    {
        [Input("andStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementAndStatementGetArgs>? AndStatement { get; set; }

        [Input("byteMatchStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementByteMatchStatementGetArgs>? ByteMatchStatement { get; set; }

        [Input("geoMatchStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementGeoMatchStatementGetArgs>? GeoMatchStatement { get; set; }

        [Input("ipSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementIpSetReferenceStatementGetArgs>? IpSetReferenceStatement { get; set; }

        [Input("notStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementNotStatementGetArgs>? NotStatement { get; set; }

        [Input("orStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementOrStatementGetArgs>? OrStatement { get; set; }

        [Input("regexPatternSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementRegexPatternSetReferenceStatementGetArgs>? RegexPatternSetReferenceStatement { get; set; }

        [Input("sizeConstraintStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementSizeConstraintStatementGetArgs>? SizeConstraintStatement { get; set; }

        [Input("sqliMatchStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementGetArgs>? SqliMatchStatement { get; set; }

        [Input("xssMatchStatement")]
        public Input<Inputs.WebAclRuleStatementOrStatementStatementOrStatementStatementXssMatchStatementGetArgs>? XssMatchStatement { get; set; }

        public WebAclRuleStatementOrStatementStatementOrStatementStatementGetArgs()
        {
        }
    }
}
