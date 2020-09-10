// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementNotStatementStatementNotStatementStatementAndStatementStatementGetArgs : Pulumi.ResourceArgs
    {
        [Input("byteMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementNotStatementStatementAndStatementStatementByteMatchStatementGetArgs>? ByteMatchStatement { get; set; }

        [Input("geoMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementNotStatementStatementAndStatementStatementGeoMatchStatementGetArgs>? GeoMatchStatement { get; set; }

        [Input("ipSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementNotStatementStatementAndStatementStatementIpSetReferenceStatementGetArgs>? IpSetReferenceStatement { get; set; }

        [Input("regexPatternSetReferenceStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementGetArgs>? RegexPatternSetReferenceStatement { get; set; }

        [Input("sizeConstraintStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementNotStatementStatementAndStatementStatementSizeConstraintStatementGetArgs>? SizeConstraintStatement { get; set; }

        [Input("sqliMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementNotStatementStatementAndStatementStatementSqliMatchStatementGetArgs>? SqliMatchStatement { get; set; }

        [Input("xssMatchStatement")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementNotStatementStatementAndStatementStatementXssMatchStatementGetArgs>? XssMatchStatement { get; set; }

        public WebAclRuleStatementNotStatementStatementNotStatementStatementAndStatementStatementGetArgs()
        {
        }
    }
}
