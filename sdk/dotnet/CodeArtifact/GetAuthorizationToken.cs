// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeArtifact
{
    public static class GetAuthorizationToken
    {
        /// <summary>
        /// The CodeArtifact Authorization Token data source generates a temporary authentication token for accessing repositories in a CodeArtifact domain.
        /// </summary>
        public static Task<GetAuthorizationTokenResult> InvokeAsync(GetAuthorizationTokenArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthorizationTokenResult>("aws:codeartifact/getAuthorizationToken:getAuthorizationToken", args ?? new GetAuthorizationTokenArgs(), options.WithVersion());
    }


    public sealed class GetAuthorizationTokenArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the domain that is in scope for the generated authorization token.
        /// </summary>
        [Input("domain", required: true)]
        public string Domain { get; set; } = null!;

        /// <summary>
        /// The account number of the AWS account that owns the domain.
        /// </summary>
        [Input("domainOwner")]
        public string? DomainOwner { get; set; }

        /// <summary>
        /// The time, in seconds, that the generated authorization token is valid. Valid values are `0` and between `900` and `43200`.
        /// </summary>
        [Input("durationSeconds")]
        public int? DurationSeconds { get; set; }

        public GetAuthorizationTokenArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAuthorizationTokenResult
    {
        /// <summary>
        /// Temporary authorization token.
        /// </summary>
        public readonly string AuthorizationToken;
        public readonly string Domain;
        public readonly string DomainOwner;
        public readonly int? DurationSeconds;
        /// <summary>
        /// The time in UTC RFC3339 format when the authorization token expires.
        /// </summary>
        public readonly string Expiration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAuthorizationTokenResult(
            string authorizationToken,

            string domain,

            string domainOwner,

            int? durationSeconds,

            string expiration,

            string id)
        {
            AuthorizationToken = authorizationToken;
            Domain = domain;
            DomainOwner = domainOwner;
            DurationSeconds = durationSeconds;
            Expiration = expiration;
            Id = id;
        }
    }
}
