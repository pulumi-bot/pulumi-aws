// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecr
{
    /// <summary>
    /// Provides an Elastic Container Registry Repository Policy.
    /// 
    /// Note that currently only one policy may be applied to a repository.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Aws.Ecr.Repository("foo", new Aws.Ecr.RepositoryArgs
    ///         {
    ///         });
    ///         var foopolicy = new Aws.Ecr.RepositoryPolicy("foopolicy", new Aws.Ecr.RepositoryPolicyArgs
    ///         {
    ///             Repository = foo.Name,
    ///             Policy = @"{
    ///     ""Version"": ""2008-10-17"",
    ///     ""Statement"": [
    ///         {
    ///             ""Sid"": ""new policy"",
    ///             ""Effect"": ""Allow"",
    ///             ""Principal"": ""*"",
    ///             ""Action"": [
    ///                 ""ecr:GetDownloadUrlForLayer"",
    ///                 ""ecr:BatchGetImage"",
    ///                 ""ecr:BatchCheckLayerAvailability"",
    ///                 ""ecr:PutImage"",
    ///                 ""ecr:InitiateLayerUpload"",
    ///                 ""ecr:UploadLayerPart"",
    ///                 ""ecr:CompleteLayerUpload"",
    ///                 ""ecr:DescribeRepositories"",
    ///                 ""ecr:GetRepositoryPolicy"",
    ///                 ""ecr:ListImages"",
    ///                 ""ecr:DeleteRepository"",
    ///                 ""ecr:BatchDeleteImage"",
    ///                 ""ecr:SetRepositoryPolicy"",
    ///                 ""ecr:DeleteRepositoryPolicy""
    ///             ]
    ///         }
    ///     ]
    /// }
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECR Repository Policy can be imported using the repository name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:ecr/repositoryPolicy:RepositoryPolicy example example
    /// ```
    /// </summary>
    [AwsResourceType("aws:ecr/repositoryPolicy:RepositoryPolicy")]
    public partial class RepositoryPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// The registry ID where the repository was created.
        /// </summary>
        [Output("registryId")]
        public Output<string> RegistryId { get; private set; } = null!;

        /// <summary>
        /// Name of the repository to apply the policy.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryPolicy(string name, RepositoryPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:ecr/repositoryPolicy:RepositoryPolicy", name, args ?? new RepositoryPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryPolicy(string name, Input<string> id, RepositoryPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:ecr/repositoryPolicy:RepositoryPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RepositoryPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryPolicy Get(string name, Input<string> id, RepositoryPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryPolicy(name, id, state, options);
        }
    }

    public sealed class RepositoryPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Input("policy", required: true)]
        public string Policy { get; set; } = null!;

        /// <summary>
        /// Name of the repository to apply the policy.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public RepositoryPolicyArgs()
        {
        }
    }

    public sealed class RepositoryPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy document. This is a JSON formatted string.
        /// </summary>
        [Input("policy")]
        public string? Policy { get; set; }

        /// <summary>
        /// The registry ID where the repository was created.
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        /// <summary>
        /// Name of the repository to apply the policy.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        public RepositoryPolicyState()
        {
        }
    }
}
