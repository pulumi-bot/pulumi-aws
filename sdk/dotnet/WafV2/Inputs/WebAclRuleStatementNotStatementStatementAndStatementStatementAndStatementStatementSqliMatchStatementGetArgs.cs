// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementGetArgs : Pulumi.ResourceArgs
    {
        [Input("fieldToMatch")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchGetArgs>? FieldToMatch { get; set; }

        [Input("textTransformations", required: true)]
        private InputList<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationGetArgs>? _textTransformations;
        public InputList<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationGetArgs> TextTransformations
        {
            get => _textTransformations ?? (_textTransformations = new InputList<Inputs.WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementTextTransformationGetArgs>());
            set => _textTransformations = value;
        }

        public WebAclRuleStatementNotStatementStatementAndStatementStatementAndStatementStatementSqliMatchStatementGetArgs()
        {
        }
    }
}
