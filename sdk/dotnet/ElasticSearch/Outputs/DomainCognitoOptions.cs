// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticSearch.Outputs
{

    [OutputType]
    public sealed class DomainCognitoOptions
    {
        /// <summary>
        /// Whether to enable encryption at rest. If the `encrypt_at_rest` block is not provided then this defaults to `false`.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// ID of the Cognito Identity Pool to use
        /// </summary>
        public readonly string IdentityPoolId;
        /// <summary>
        /// ARN of the IAM role that has the AmazonESCognitoAccess policy attached
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// ID of the Cognito User Pool to use
        /// </summary>
        public readonly string UserPoolId;

        [OutputConstructor]
        private DomainCognitoOptions(
            bool? enabled,

            string identityPoolId,

            string roleArn,

            string userPoolId)
        {
            Enabled = enabled;
            IdentityPoolId = identityPoolId;
            RoleArn = roleArn;
            UserPoolId = userPoolId;
        }
    }
}
