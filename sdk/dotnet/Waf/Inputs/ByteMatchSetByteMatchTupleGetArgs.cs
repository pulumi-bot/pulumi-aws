// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf.Inputs
{

    public sealed class ByteMatchSetByteMatchTupleGetArgs : Pulumi.ResourceArgs
    {
        [Input("fieldToMatch", required: true)]
        public Input<Inputs.ByteMatchSetByteMatchTupleFieldToMatchGetArgs> FieldToMatch { get; set; } = null!;

        [Input("positionalConstraint", required: true)]
        public Input<string> PositionalConstraint { get; set; } = null!;

        [Input("targetString")]
        public Input<string>? TargetString { get; set; }

        [Input("textTransformation", required: true)]
        public Input<string> TextTransformation { get; set; } = null!;

        public ByteMatchSetByteMatchTupleGetArgs()
        {
        }
    }
}
