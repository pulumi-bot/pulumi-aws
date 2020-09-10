// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementGetArgs : Pulumi.ResourceArgs
    {
        [Input("byteMatchStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementByteMatchStatementGetArgs>? ByteMatchStatement { get; set; }

        [Input("geoMatchStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementGeoMatchStatementGetArgs>? GeoMatchStatement { get; set; }

        [Input("ipSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementIpSetReferenceStatementGetArgs>? IpSetReferenceStatement { get; set; }

        [Input("regexPatternSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementRegexPatternSetReferenceStatementGetArgs>? RegexPatternSetReferenceStatement { get; set; }

        [Input("sizeConstraintStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementSizeConstraintStatementGetArgs>? SizeConstraintStatement { get; set; }

        [Input("sqliMatchStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementSqliMatchStatementGetArgs>? SqliMatchStatement { get; set; }

        [Input("xssMatchStatement")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementXssMatchStatementGetArgs>? XssMatchStatement { get; set; }

        public WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementGetArgs()
        {
        }
    }
}
