// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticSearch.Inputs
{

    public sealed class DomainCognitoOptionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable encryption at rest. If the `encrypt_at_rest` block is not provided then this defaults to `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// ID of the Cognito Identity Pool to use
        /// </summary>
        [Input("identityPoolId", required: true)]
        public Input<string> IdentityPoolId { get; set; } = null!;

        /// <summary>
        /// ARN of the IAM role that has the AmazonESCognitoAccess policy attached
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// ID of the Cognito User Pool to use
        /// </summary>
        [Input("userPoolId", required: true)]
        public Input<string> UserPoolId { get; set; } = null!;

        public DomainCognitoOptionsGetArgs()
        {
        }
    }
}
