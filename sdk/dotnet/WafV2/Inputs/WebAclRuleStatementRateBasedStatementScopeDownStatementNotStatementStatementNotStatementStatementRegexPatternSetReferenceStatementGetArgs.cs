// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementGetArgs : Pulumi.ResourceArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("fieldToMatch")]
        public Input<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchGetArgs>? FieldToMatch { get; set; }

        [Input("textTransformations", required: true)]
        private InputList<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementTextTransformationGetArgs>? _textTransformations;
        public InputList<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementTextTransformationGetArgs> TextTransformations
        {
            get => _textTransformations ?? (_textTransformations = new InputList<Inputs.WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementTextTransformationGetArgs>());
            set => _textTransformations = value;
        }

        public WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementGetArgs()
        {
        }
    }
}
