// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementGetArgs : Pulumi.ResourceArgs
    {
        [Input("comparisonOperator", required: true)]
        public Input<string> ComparisonOperator { get; set; } = null!;

        [Input("fieldToMatch")]
        public Input<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchGetArgs>? FieldToMatch { get; set; }

        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        [Input("textTransformations", required: true)]
        private InputList<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementTextTransformationGetArgs>? _textTransformations;
        public InputList<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementTextTransformationGetArgs> TextTransformations
        {
            get => _textTransformations ?? (_textTransformations = new InputList<Inputs.WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementTextTransformationGetArgs>());
            set => _textTransformations = value;
        }

        public WebAclRuleStatementNotStatementStatementOrStatementStatementOrStatementStatementSizeConstraintStatementGetArgs()
        {
        }
    }
}
