// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleStatementNotStatementArgs : Pulumi.ResourceArgs
    {
        [Input("statements", required: true)]
        private InputList<Inputs.WebAclRuleStatementNotStatementStatementArgs>? _statements;
        public InputList<Inputs.WebAclRuleStatementNotStatementStatementArgs> Statements
        {
            get => _statements ?? (_statements = new InputList<Inputs.WebAclRuleStatementNotStatementStatementArgs>());
            set => _statements = value;
        }

        public WebAclRuleStatementNotStatementArgs()
        {
        }
    }
}
