// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementAndStatementStatementOrStatementStatementOrStatementStatementGetArgs : Pulumi.ResourceArgs
    {
        [Input("byteMatchStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementOrStatementStatementOrStatementStatementByteMatchStatementGetArgs>? ByteMatchStatement { get; set; }

        [Input("geoMatchStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementOrStatementStatementOrStatementStatementGeoMatchStatementGetArgs>? GeoMatchStatement { get; set; }

        [Input("ipSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementOrStatementStatementOrStatementStatementIpSetReferenceStatementGetArgs>? IpSetReferenceStatement { get; set; }

        [Input("regexPatternSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementOrStatementStatementOrStatementStatementRegexPatternSetReferenceStatementGetArgs>? RegexPatternSetReferenceStatement { get; set; }

        [Input("sizeConstraintStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementGetArgs>? SizeConstraintStatement { get; set; }

        [Input("sqliMatchStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementOrStatementStatementOrStatementStatementSqliMatchStatementGetArgs>? SqliMatchStatement { get; set; }

        [Input("xssMatchStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementOrStatementStatementOrStatementStatementXssMatchStatementGetArgs>? XssMatchStatement { get; set; }

        public WebAclRuleStatementAndStatementStatementOrStatementStatementOrStatementStatementGetArgs()
        {
        }
    }
}
