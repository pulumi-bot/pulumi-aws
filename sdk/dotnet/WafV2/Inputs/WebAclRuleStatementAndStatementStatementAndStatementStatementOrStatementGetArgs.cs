// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementGetArgs : Pulumi.ResourceArgs
    {
        [Input("statements", required: true)]
        private InputList<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementGetArgs>? _statements;
        public InputList<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementGetArgs> Statements
        {
            get => _statements ?? (_statements = new InputList<Inputs.WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementStatementGetArgs>());
            set => _statements = value;
        }

        public WebAclRuleStatementAndStatementStatementAndStatementStatementOrStatementGetArgs()
        {
        }
    }
}
