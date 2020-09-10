// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf.Inputs
{

    public sealed class RegexMatchSetRegexMatchTupleArgs : Pulumi.ResourceArgs
    {
        [Input("fieldToMatch", required: true)]
        public Input<Inputs.RegexMatchSetRegexMatchTupleFieldToMatchArgs> FieldToMatch { get; set; } = null!;

        [Input("regexPatternSetId", required: true)]
        public Input<string> RegexPatternSetId { get; set; } = null!;

        [Input("textTransformation", required: true)]
        public Input<string> TextTransformation { get; set; } = null!;

        public RegexMatchSetRegexMatchTupleArgs()
        {
        }
    }
}
