// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementRegexPatternSetReferenceStatementGetArgs : Pulumi.ResourceArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("fieldToMatch")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementRegexPatternSetReferenceStatementFieldToMatchGetArgs>? FieldToMatch { get; set; }

        [Input("textTransformations", required: true)]
        private InputList<Inputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementRegexPatternSetReferenceStatementTextTransformationGetArgs>? _textTransformations;
        public InputList<Inputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementRegexPatternSetReferenceStatementTextTransformationGetArgs> TextTransformations
        {
            get => _textTransformations ?? (_textTransformations = new InputList<Inputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementRegexPatternSetReferenceStatementTextTransformationGetArgs>());
            set => _textTransformations = value;
        }

        public WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatementRegexPatternSetReferenceStatementGetArgs()
        {
        }
    }
}
