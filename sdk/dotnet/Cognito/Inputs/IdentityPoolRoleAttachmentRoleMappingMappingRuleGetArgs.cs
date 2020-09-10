// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito.Inputs
{

    public sealed class IdentityPoolRoleAttachmentRoleMappingMappingRuleGetArgs : Pulumi.ResourceArgs
    {
        [Input("claim", required: true)]
        public Input<string> Claim { get; set; } = null!;

        [Input("matchType", required: true)]
        public Input<string> MatchType { get; set; } = null!;

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public IdentityPoolRoleAttachmentRoleMappingMappingRuleGetArgs()
        {
        }
    }
}
