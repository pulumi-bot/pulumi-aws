// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementAndStatementStatementAndStatementStatementAndStatementStatementByteMatchStatementArgs : Pulumi.ResourceArgs
    {
        [Input("fieldToMatch")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementAndStatementStatementByteMatchStatementFieldToMatchArgs>? FieldToMatch { get; set; }

        [Input("positionalConstraint", required: true)]
        public Input<string> PositionalConstraint { get; set; } = null!;

        [Input("searchString", required: true)]
        public Input<string> SearchString { get; set; } = null!;

        [Input("textTransformations", required: true)]
        private InputList<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementAndStatementStatementByteMatchStatementTextTransformationArgs>? _textTransformations;
        public InputList<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementAndStatementStatementByteMatchStatementTextTransformationArgs> TextTransformations
        {
            get => _textTransformations ?? (_textTransformations = new InputList<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementAndStatementStatementByteMatchStatementTextTransformationArgs>());
            set => _textTransformations = value;
        }

        public WebAclRuleStatementAndStatementStatementAndStatementStatementAndStatementStatementByteMatchStatementArgs()
        {
        }
    }
}
