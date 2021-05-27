// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeArtifact
{
    public static class GetRepositoryEndpoint
    {
        /// <summary>
        /// The CodeArtifact Repository Endpoint data source returns the endpoint of a repository for a specific package format.
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
        ///         var test = Output.Create(Aws.CodeArtifact.GetRepositoryEndpoint.InvokeAsync(new Aws.CodeArtifact.GetRepositoryEndpointArgs
        ///         {
        ///             Domain = aws_codeartifact_domain.Test.Domain,
        ///             Repository = aws_codeartifact_repository.Test.Repository,
        ///             Format = "npm",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRepositoryEndpointResult> InvokeAsync(GetRepositoryEndpointArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryEndpointResult>("aws:codeartifact/getRepositoryEndpoint:getRepositoryEndpoint", args ?? new GetRepositoryEndpointArgs(), options.WithVersion());

        public static Output<GetRepositoryEndpointResult> Invoke(GetRepositoryEndpointOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Domain.Box(),
                args.DomainOwner.Box(),
                args.Format.Box(),
                args.Repository.Box()
            ).Apply(a => {
                    var args = new GetRepositoryEndpointArgs();
                    a[0].Set(args, nameof(args.Domain));
                    a[1].Set(args, nameof(args.DomainOwner));
                    a[2].Set(args, nameof(args.Format));
                    a[3].Set(args, nameof(args.Repository));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetRepositoryEndpointArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the domain that contains the repository.
        /// </summary>
        [Input("domain", required: true)]
        public string Domain { get; set; } = null!;

        /// <summary>
        /// The account number of the AWS account that owns the domain.
        /// </summary>
        [Input("domainOwner")]
        public string? DomainOwner { get; set; }

        /// <summary>
        /// Which endpoint of a repository to return. A repository has one endpoint for each package format: `npm`, `pypi`, `maven`, and `nuget`.
        /// </summary>
        [Input("format", required: true)]
        public string Format { get; set; } = null!;

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetRepositoryEndpointArgs()
        {
        }
    }

    public sealed class GetRepositoryEndpointOutputArgs
    {
        /// <summary>
        /// The name of the domain that contains the repository.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The account number of the AWS account that owns the domain.
        /// </summary>
        [Input("domainOwner")]
        public Input<string>? DomainOwner { get; set; }

        /// <summary>
        /// Which endpoint of a repository to return. A repository has one endpoint for each package format: `npm`, `pypi`, `maven`, and `nuget`.
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetRepositoryEndpointOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRepositoryEndpointResult
    {
        public readonly string Domain;
        public readonly string DomainOwner;
        public readonly string Format;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Repository;
        /// <summary>
        /// The URL of the returned endpoint.
        /// </summary>
        public readonly string RepositoryEndpoint;

        [OutputConstructor]
        private GetRepositoryEndpointResult(
            string domain,

            string domainOwner,

            string format,

            string id,

            string repository,

            string repositoryEndpoint)
        {
            Domain = domain;
            DomainOwner = domainOwner;
            Format = format;
            Id = id;
            Repository = repository;
            RepositoryEndpoint = repositoryEndpoint;
        }
    }
}
