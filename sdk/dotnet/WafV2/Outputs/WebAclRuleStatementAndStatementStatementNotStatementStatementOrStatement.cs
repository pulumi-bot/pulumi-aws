// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Outputs
{

    [OutputType]
    public sealed class WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatement
    {
        public readonly ImmutableArray<Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatement> Statements;

        [OutputConstructor]
        private WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatement(ImmutableArray<Outputs.WebAclRuleStatementAndStatementStatementNotStatementStatementOrStatementStatement> statements)
        {
            Statements = statements;
        }
    }
}
