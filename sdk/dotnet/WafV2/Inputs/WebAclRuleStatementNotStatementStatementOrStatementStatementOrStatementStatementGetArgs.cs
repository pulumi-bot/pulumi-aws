// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementGetArgs : Pulumi.ResourceArgs
    {
        [Input("byteMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementByteMatchStatementGetArgs>? ByteMatchStatement { get; set; }

        [Input("geoMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementGeoMatchStatementGetArgs>? GeoMatchStatement { get; set; }

        [Input("ipSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementIpSetReferenceStatementGetArgs>? IpSetReferenceStatement { get; set; }

        [Input("regexPatternSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementRegexPatternSetReferenceStatementGetArgs>? RegexPatternSetReferenceStatement { get; set; }

        [Input("sizeConstraintStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementGetArgs>? SizeConstraintStatement { get; set; }

        [Input("sqliMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSqliMatchStatementGetArgs>? SqliMatchStatement { get; set; }

        [Input("xssMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementXssMatchStatementGetArgs>? XssMatchStatement { get; set; }

        public WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementGetArgs()
        {
        }
    }
}
