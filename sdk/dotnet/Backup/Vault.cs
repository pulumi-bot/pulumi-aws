// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Backup
{
    /// <summary>
    /// Provides an AWS Backup vault resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/backup_vault.html.markdown.
    /// </summary>
    public partial class Vault : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the vault.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The server-side encryption key that is used to protect your backups.
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string> KmsKeyArn { get; private set; } = null!;

        /// <summary>
        /// Name of the backup vault to create.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The number of recovery points that are stored in a backup vault.
        /// </summary>
        [Output("recoveryPoints")]
        public Output<int> RecoveryPoints { get; private set; } = null!;

        /// <summary>
        /// Metadata that you can assign to help organize the resources that you create.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Vault resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vault(string name, VaultArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:backup/vault:Vault", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Vault(string name, Input<string> id, VaultState? state = null, CustomResourceOptions? options = null)
            : base("aws:backup/vault:Vault", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vault resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vault Get(string name, Input<string> id, VaultState? state = null, CustomResourceOptions? options = null)
        {
            return new Vault(name, id, state, options);
        }
    }

    public sealed class VaultArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The server-side encryption key that is used to protect your backups.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// Name of the backup vault to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Metadata that you can assign to help organize the resources that you create.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public VaultArgs()
        {
        }
    }

    public sealed class VaultState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the vault.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The server-side encryption key that is used to protect your backups.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// Name of the backup vault to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of recovery points that are stored in a backup vault.
        /// </summary>
        [Input("recoveryPoints")]
        public Input<int>? RecoveryPoints { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Metadata that you can assign to help organize the resources that you create.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public VaultState()
        {
        }
    }
}
