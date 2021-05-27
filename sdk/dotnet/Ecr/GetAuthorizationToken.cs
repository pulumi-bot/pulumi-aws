// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecr
{
    public static class GetAuthorizationToken
    {
        /// <summary>
        /// The ECR Authorization Token data source allows the authorization token, proxy endpoint, token expiration date, user name and password to be retrieved for an ECR repository.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var token = Output.Create(Aws.Ecr.GetAuthorizationToken.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAuthorizationTokenResult> InvokeAsync(GetAuthorizationTokenArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthorizationTokenResult>("aws:ecr/getAuthorizationToken:getAuthorizationToken", args ?? new GetAuthorizationTokenArgs(), options.WithVersion());

        public static Output<GetAuthorizationTokenResult> Invoke(GetAuthorizationTokenOutputArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetAuthorizationTokenOutputArgs();
            return Pulumi.Output.All(
                args.RegistryId.Box()
            ).Apply(a => {
                    var args = new GetAuthorizationTokenArgs();
                    a[0].Set(args, nameof(args.RegistryId));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetAuthorizationTokenArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// AWS account ID of the ECR Repository. If not specified the default account is assumed.
        /// </summary>
        [Input("registryId")]
        public string? RegistryId { get; set; }

        public GetAuthorizationTokenArgs()
        {
        }
    }

    public sealed class GetAuthorizationTokenOutputArgs
    {
        /// <summary>
        /// AWS account ID of the ECR Repository. If not specified the default account is assumed.
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        public GetAuthorizationTokenOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAuthorizationTokenResult
    {
        /// <summary>
        /// Temporary IAM authentication credentials to access the ECR repository encoded in base64 in the form of `user_name:password`.
        /// </summary>
        public readonly string AuthorizationToken;
        /// <summary>
        /// The time in UTC RFC3339 format when the authorization token expires.
        /// </summary>
        public readonly string ExpiresAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Password decoded from the authorization token.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The registry URL to use in the docker login command.
        /// </summary>
        public readonly string ProxyEndpoint;
        public readonly string? RegistryId;
        /// <summary>
        /// User name decoded from the authorization token.
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private GetAuthorizationTokenResult(
            string authorizationToken,

            string expiresAt,

            string id,

            string password,

            string proxyEndpoint,

            string? registryId,

            string userName)
        {
            AuthorizationToken = authorizationToken;
            ExpiresAt = expiresAt;
            Id = id;
            Password = password;
            ProxyEndpoint = proxyEndpoint;
            RegistryId = registryId;
            UserName = userName;
        }
    }
}
