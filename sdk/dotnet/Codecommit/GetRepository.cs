// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeCommit
{
    public static partial class Invokes
    {
        /// <summary>
        /// The CodeCommit Repository data source allows the ARN, Repository ID, Repository URL for HTTP and Repository URL for SSH to be retrieved for an CodeCommit repository.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/codecommit_repository.html.markdown.
        /// </summary>
        [Obsolete("Use GetRepository.InvokeAsync() instead")]
        public static Task<GetRepositoryResult> GetRepository(GetRepositoryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryResult>("aws:codecommit/getRepository:getRepository", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetRepository
    {
        /// <summary>
        /// The CodeCommit Repository data source allows the ARN, Repository ID, Repository URL for HTTP and Repository URL for SSH to be retrieved for an CodeCommit repository.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/codecommit_repository.html.markdown.
        /// </summary>
        public static Task<GetRepositoryResult> InvokeAsync(GetRepositoryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryResult>("aws:codecommit/getRepository:getRepository", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetRepositoryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name for the repository. This needs to be less than 100 characters.
        /// </summary>
        [Input("repositoryName", required: true)]
        public string RepositoryName { get; set; } = null!;

        public GetRepositoryArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetRepositoryResult
    {
        /// <summary>
        /// The ARN of the repository
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The URL to use for cloning the repository over HTTPS.
        /// </summary>
        public readonly string CloneUrlHttp;
        /// <summary>
        /// The URL to use for cloning the repository over SSH.
        /// </summary>
        public readonly string CloneUrlSsh;
        /// <summary>
        /// The ID of the repository
        /// </summary>
        public readonly string RepositoryId;
        public readonly string RepositoryName;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRepositoryResult(
            string arn,
            string cloneUrlHttp,
            string cloneUrlSsh,
            string repositoryId,
            string repositoryName,
            string id)
        {
            Arn = arn;
            CloneUrlHttp = cloneUrlHttp;
            CloneUrlSsh = cloneUrlSsh;
            RepositoryId = repositoryId;
            RepositoryName = repositoryName;
            Id = id;
        }
    }
}
