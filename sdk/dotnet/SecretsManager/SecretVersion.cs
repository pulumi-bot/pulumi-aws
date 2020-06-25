// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecretsManager
{
    /// <summary>
    /// Provides a resource to manage AWS Secrets Manager secret version including its secret value. To manage secret metadata, see the `aws.secretsmanager.Secret` resource.
    /// 
    /// &gt; **NOTE:** If the `AWSCURRENT` staging label is present on this version during resource deletion, that label cannot be removed and will be skipped to prevent errors when fully deleting the secret. That label will leave this secret version active even after the resource is deleted from this provider unless the secret itself is deleted. Move the `AWSCURRENT` staging label before or after deleting this resource from this provider to fully trigger version deprecation if necessary.
    /// 
    /// ## Example Usage
    /// ### Simple String Value
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.SecretsManager.SecretVersion("example", new Aws.SecretsManager.SecretVersionArgs
    ///         {
    ///             SecretId = aws_secretsmanager_secret.Example.Id,
    ///             SecretString = "example-string-to-protect",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class SecretVersion : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the secret.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specifies binary data that you want to encrypt and store in this version of the secret. This is required if secret_string is not set. Needs to be encoded to base64.
        /// </summary>
        [Output("secretBinary")]
        public Output<string?> SecretBinary { get; private set; } = null!;

        /// <summary>
        /// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
        /// </summary>
        [Output("secretId")]
        public Output<string> SecretId { get; private set; } = null!;

        /// <summary>
        /// Specifies text data that you want to encrypt and store in this version of the secret. This is required if secret_binary is not set.
        /// </summary>
        [Output("secretString")]
        public Output<string?> SecretString { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the version of the secret.
        /// </summary>
        [Output("versionId")]
        public Output<string> VersionId { get; private set; } = null!;

        /// <summary>
        /// Specifies a list of staging labels that are attached to this version of the secret. A staging label must be unique to a single version of the secret. If you specify a staging label that's already associated with a different version of the same secret then that staging label is automatically removed from the other version and attached to this version. If you do not specify a value, then AWS Secrets Manager automatically moves the staging label `AWSCURRENT` to this new version on creation.
        /// </summary>
        [Output("versionStages")]
        public Output<ImmutableArray<string>> VersionStages { get; private set; } = null!;


        /// <summary>
        /// Create a SecretVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretVersion(string name, SecretVersionArgs args, CustomResourceOptions? options = null)
            : base("aws:secretsmanager/secretVersion:SecretVersion", name, args ?? new SecretVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretVersion(string name, Input<string> id, SecretVersionState? state = null, CustomResourceOptions? options = null)
            : base("aws:secretsmanager/secretVersion:SecretVersion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretVersion Get(string name, Input<string> id, SecretVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretVersion(name, id, state, options);
        }
    }

    public sealed class SecretVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies binary data that you want to encrypt and store in this version of the secret. This is required if secret_string is not set. Needs to be encoded to base64.
        /// </summary>
        [Input("secretBinary")]
        public Input<string>? SecretBinary { get; set; }

        /// <summary>
        /// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
        /// </summary>
        [Input("secretId", required: true)]
        public Input<string> SecretId { get; set; } = null!;

        /// <summary>
        /// Specifies text data that you want to encrypt and store in this version of the secret. This is required if secret_binary is not set.
        /// </summary>
        [Input("secretString")]
        public Input<string>? SecretString { get; set; }

        [Input("versionStages")]
        private InputList<string>? _versionStages;

        /// <summary>
        /// Specifies a list of staging labels that are attached to this version of the secret. A staging label must be unique to a single version of the secret. If you specify a staging label that's already associated with a different version of the same secret then that staging label is automatically removed from the other version and attached to this version. If you do not specify a value, then AWS Secrets Manager automatically moves the staging label `AWSCURRENT` to this new version on creation.
        /// </summary>
        public InputList<string> VersionStages
        {
            get => _versionStages ?? (_versionStages = new InputList<string>());
            set => _versionStages = value;
        }

        public SecretVersionArgs()
        {
        }
    }

    public sealed class SecretVersionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the secret.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Specifies binary data that you want to encrypt and store in this version of the secret. This is required if secret_string is not set. Needs to be encoded to base64.
        /// </summary>
        [Input("secretBinary")]
        public Input<string>? SecretBinary { get; set; }

        /// <summary>
        /// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        /// <summary>
        /// Specifies text data that you want to encrypt and store in this version of the secret. This is required if secret_binary is not set.
        /// </summary>
        [Input("secretString")]
        public Input<string>? SecretString { get; set; }

        /// <summary>
        /// The unique identifier of the version of the secret.
        /// </summary>
        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        [Input("versionStages")]
        private InputList<string>? _versionStages;

        /// <summary>
        /// Specifies a list of staging labels that are attached to this version of the secret. A staging label must be unique to a single version of the secret. If you specify a staging label that's already associated with a different version of the same secret then that staging label is automatically removed from the other version and attached to this version. If you do not specify a value, then AWS Secrets Manager automatically moves the staging label `AWSCURRENT` to this new version on creation.
        /// </summary>
        public InputList<string> VersionStages
        {
            get => _versionStages ?? (_versionStages = new InputList<string>());
            set => _versionStages = value;
        }

        public SecretVersionState()
        {
        }
    }
}
