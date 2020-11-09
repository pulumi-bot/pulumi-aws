// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeArtifact
{
    /// <summary>
    /// Provides a CodeArtifact Domains Permissions Policy Resource.
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
    ///         var exampleKey = new Aws.Kms.Key("exampleKey", new Aws.Kms.KeyArgs
    ///         {
    ///             Description = "domain key",
    ///         });
    ///         var exampleDomain = new Aws.CodeArtifact.Domain("exampleDomain", new Aws.CodeArtifact.DomainArgs
    ///         {
    ///             Domain = "example.com",
    ///             EncryptionKey = exampleKey.Arn,
    ///         });
    ///         var test = new Aws.CodeArtifact.DomainPermissions("test", new Aws.CodeArtifact.DomainPermissionsArgs
    ///         {
    ///             Domain = exampleDomain.DomainName,
    ///             PolicyDocument = exampleDomain.Arn.Apply(arn =&gt; @$"{{
    ///     ""Version"": ""2012-10-17"",
    ///     ""Statement"": [
    ///         {{
    ///             ""Action"": ""codeartifact:CreateRepository"",
    ///             ""Effect"": ""Allow"",
    ///             ""Principal"": ""*"",
    ///             ""Resource"": ""{arn}""
    ///         }}
    ///     ]
    /// }}
    /// "),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CodeArtifact Domain Permissions Policies can be imported using the CodeArtifact Domain ARN, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:codeartifact/domainPermissions:DomainPermissions example arn:aws:codeartifact:us-west-2:012345678912:domain/tf-acc-test-1928056699409417367
    /// ```
    /// </summary>
    public partial class DomainPermissions : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the domain on which to set the resource policy.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The account number of the AWS account that owns the domain.
        /// </summary>
        [Output("domainOwner")]
        public Output<string> DomainOwner { get; private set; } = null!;

        /// <summary>
        /// A JSON policy string to be set as the access control resource policy on the provided domain.
        /// </summary>
        [Output("policyDocument")]
        public Output<string> PolicyDocument { get; private set; } = null!;

        /// <summary>
        /// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
        /// </summary>
        [Output("policyRevision")]
        public Output<string> PolicyRevision { get; private set; } = null!;

        /// <summary>
        /// The ARN of the resource associated with the resource policy.
        /// </summary>
        [Output("resourceArn")]
        public Output<string> ResourceArn { get; private set; } = null!;


        /// <summary>
        /// Create a DomainPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainPermissions(string name, DomainPermissionsArgs args, CustomResourceOptions? options = null)
            : base("aws:codeartifact/domainPermissions:DomainPermissions", name, args ?? new DomainPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainPermissions(string name, Input<string> id, DomainPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("aws:codeartifact/domainPermissions:DomainPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainPermissions Get(string name, Input<string> id, DomainPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainPermissions(name, id, state, options);
        }
    }

    public sealed class DomainPermissionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the domain on which to set the resource policy.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The account number of the AWS account that owns the domain.
        /// </summary>
        [Input("domainOwner")]
        public Input<string>? DomainOwner { get; set; }

        /// <summary>
        /// A JSON policy string to be set as the access control resource policy on the provided domain.
        /// </summary>
        [Input("policyDocument", required: true)]
        public Input<string> PolicyDocument { get; set; } = null!;

        /// <summary>
        /// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
        /// </summary>
        [Input("policyRevision")]
        public Input<string>? PolicyRevision { get; set; }

        public DomainPermissionsArgs()
        {
        }
    }

    public sealed class DomainPermissionsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the domain on which to set the resource policy.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The account number of the AWS account that owns the domain.
        /// </summary>
        [Input("domainOwner")]
        public Input<string>? DomainOwner { get; set; }

        /// <summary>
        /// A JSON policy string to be set as the access control resource policy on the provided domain.
        /// </summary>
        [Input("policyDocument")]
        public Input<string>? PolicyDocument { get; set; }

        /// <summary>
        /// The current revision of the resource policy to be set. This revision is used for optimistic locking, which prevents others from overwriting your changes to the domain's resource policy.
        /// </summary>
        [Input("policyRevision")]
        public Input<string>? PolicyRevision { get; set; }

        /// <summary>
        /// The ARN of the resource associated with the resource policy.
        /// </summary>
        [Input("resourceArn")]
        public Input<string>? ResourceArn { get; set; }

        public DomainPermissionsState()
        {
        }
    }
}
