// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementAndStatementStatementOrStatementStatementAndStatementStatementSqliMatchStatementGetArgs : Pulumi.ResourceArgs
    {
        [Input("fieldToMatch")]
        public Input<Inputs.WebAclRuleStatementAndStatementStatementOrStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchGetArgs>? FieldToMatch { get; set; }

        [Input("textTransformations", required: true)]
        private InputList<Inputs.WebAclRuleStatementAndStatementStatementOrStatementStatementAndStatementStatementSqliMatchStatementTextTransformationGetArgs>? _textTransformations;
        public InputList<Inputs.WebAclRuleStatementAndStatementStatementOrStatementStatementAndStatementStatementSqliMatchStatementTextTransformationGetArgs> TextTransformations
        {
            get => _textTransformations ?? (_textTransformations = new InputList<Inputs.WebAclRuleStatementAndStatementStatementOrStatementStatementAndStatementStatementSqliMatchStatementTextTransformationGetArgs>());
            set => _textTransformations = value;
        }

        public WebAclRuleStatementAndStatementStatementOrStatementStatementAndStatementStatementSqliMatchStatementGetArgs()
        {
        }
    }
}
