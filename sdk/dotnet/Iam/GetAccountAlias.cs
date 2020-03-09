// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam
{
    public static partial class GetAccountAlias
    {
        /// <summary>
        /// The IAM Account Alias data source allows access to the account alias
        /// for the effective account in which this provider is working.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/iam_account_alias.html.markdown.
        /// </summary>
        public static Task<GetAccountAliasResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountAliasResult>("aws:iam/getAccountAlias:getAccountAlias", InvokeArgs.Empty, options.WithVersion());
    }

    [OutputType]
    public sealed class GetAccountAliasResult
    {
        /// <summary>
        /// The alias associated with the AWS account.
        /// </summary>
        public readonly string AccountAlias;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAccountAliasResult(
            string accountAlias,
            string id)
        {
            AccountAlias = accountAlias;
            Id = id;
        }
    }
}
