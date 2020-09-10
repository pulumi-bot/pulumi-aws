// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementAndStatementStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementArgs : Pulumi.ResourceArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("fieldToMatch")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchArgs>? FieldToMatch { get; set; }

        [Input("textTransformations", required: true)]
        private InputList<Inputs.WebAclRuleStatementAndStatementStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementTextTransformationArgs>? _textTransformations;
        public InputList<Inputs.WebAclRuleStatementAndStatementStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementTextTransformationArgs> TextTransformations
        {
            get => _textTransformations ?? (_textTransformations = new InputList<Inputs.WebAclRuleStatementAndStatementStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementTextTransformationArgs>());
            set => _textTransformations = value;
        }

        public WebAclRuleStatementAndStatementStatementNotStatementStatementAndStatementStatementRegexPatternSetReferenceStatementArgs()
        {
        }
    }
}
