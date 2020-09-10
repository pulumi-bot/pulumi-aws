// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementNotStatementStatementOrStatementStatementGetArgs : Pulumi.ResourceArgs
    {
        [Input("andStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementAndStatementGetArgs>? AndStatement { get; set; }

        [Input("byteMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementByteMatchStatementGetArgs>? ByteMatchStatement { get; set; }

        [Input("geoMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementGeoMatchStatementGetArgs>? GeoMatchStatement { get; set; }

        [Input("ipSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementIpSetReferenceStatementGetArgs>? IpSetReferenceStatement { get; set; }

        [Input("notStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementNotStatementGetArgs>? NotStatement { get; set; }

        [Input("orStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementGetArgs>? OrStatement { get; set; }

        [Input("regexPatternSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementRegexPatternSetReferenceStatementGetArgs>? RegexPatternSetReferenceStatement { get; set; }

        [Input("sizeConstraintStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementGetArgs>? SizeConstraintStatement { get; set; }

        [Input("sqliMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementSqliMatchStatementGetArgs>? SqliMatchStatement { get; set; }

        [Input("xssMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementXssMatchStatementGetArgs>? XssMatchStatement { get; set; }

        public WebAclRuleStatementNotStatementStatementOrStatementStatementGetArgs()
        {
        }
    }
}
