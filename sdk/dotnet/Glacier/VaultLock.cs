// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glacier
{
    /// <summary>
    /// Manages a Glacier Vault Lock. You can refer to the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html) for a full explanation of the Glacier Vault Lock functionality.
    /// 
    /// &gt; **NOTE:** This resource allows you to test Glacier Vault Lock policies by setting the `complete_lock` argument to `false`. When testing policies in this manner, the Glacier Vault Lock automatically expires after 24 hours and this provider will show this resource as needing recreation after that time. To permanently apply the policy, set the `complete_lock` argument to `true`. When changing `complete_lock` to `true`, it is expected the resource will show as recreating.
    /// 
    /// !&gt; **WARNING:** Once a Glacier Vault Lock is completed, it is immutable. The deletion of the Glacier Vault Lock is not be possible and attempting to remove it from this provider will return an error. Set the `ignore_deletion_error` argument to `true` and apply this configuration before attempting to delete this resource via this provider or remove this resource from this provider's management.
    /// 
    /// {{% examples %}}
    /// {{% /examples %}}
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/glacier_vault_lock.html.markdown.
    /// </summary>
    public partial class VaultLock : Pulumi.CustomResource
    {
        /// <summary>
        /// Boolean whether to permanently apply this Glacier Lock Policy. Once completed, this cannot be undone. If set to `false`, the Glacier Lock Policy remains in a testing mode for 24 hours. After that time, the Glacier Lock Policy is automatically removed by Glacier and the this provider resource will show as needing recreation. Changing this from `false` to `true` will show as resource recreation, which is expected. Changing this from `true` to `false` is not possible unless the Glacier Vault is recreated at the same time.
        /// </summary>
        [Output("completeLock")]
        public Output<bool> CompleteLock { get; private set; } = null!;

        /// <summary>
        /// Allow this provider to ignore the error returned when attempting to delete the Glacier Lock Policy. This can be used to delete or recreate the Glacier Vault via this provider, for example, if the Glacier Vault Lock policy permits that action. This should only be used in conjunction with `complete_lock` being set to `true`.
        /// </summary>
        [Output("ignoreDeletionError")]
        public Output<bool?> IgnoreDeletionError { get; private set; } = null!;

        /// <summary>
        /// JSON string containing the IAM policy to apply as the Glacier Vault Lock policy.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// The name of the Glacier Vault.
        /// </summary>
        [Output("vaultName")]
        public Output<string> VaultName { get; private set; } = null!;


        /// <summary>
        /// Create a VaultLock resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VaultLock(string name, VaultLockArgs args, CustomResourceOptions? options = null)
            : base("aws:glacier/vaultLock:VaultLock", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private VaultLock(string name, Input<string> id, VaultLockState? state = null, CustomResourceOptions? options = null)
            : base("aws:glacier/vaultLock:VaultLock", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VaultLock resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VaultLock Get(string name, Input<string> id, VaultLockState? state = null, CustomResourceOptions? options = null)
        {
            return new VaultLock(name, id, state, options);
        }
    }

    public sealed class VaultLockArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean whether to permanently apply this Glacier Lock Policy. Once completed, this cannot be undone. If set to `false`, the Glacier Lock Policy remains in a testing mode for 24 hours. After that time, the Glacier Lock Policy is automatically removed by Glacier and the this provider resource will show as needing recreation. Changing this from `false` to `true` will show as resource recreation, which is expected. Changing this from `true` to `false` is not possible unless the Glacier Vault is recreated at the same time.
        /// </summary>
        [Input("completeLock", required: true)]
        public Input<bool> CompleteLock { get; set; } = null!;

        /// <summary>
        /// Allow this provider to ignore the error returned when attempting to delete the Glacier Lock Policy. This can be used to delete or recreate the Glacier Vault via this provider, for example, if the Glacier Vault Lock policy permits that action. This should only be used in conjunction with `complete_lock` being set to `true`.
        /// </summary>
        [Input("ignoreDeletionError")]
        public Input<bool>? IgnoreDeletionError { get; set; }

        /// <summary>
        /// JSON string containing the IAM policy to apply as the Glacier Vault Lock policy.
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        /// <summary>
        /// The name of the Glacier Vault.
        /// </summary>
        [Input("vaultName", required: true)]
        public Input<string> VaultName { get; set; } = null!;

        public VaultLockArgs()
        {
        }
    }

    public sealed class VaultLockState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean whether to permanently apply this Glacier Lock Policy. Once completed, this cannot be undone. If set to `false`, the Glacier Lock Policy remains in a testing mode for 24 hours. After that time, the Glacier Lock Policy is automatically removed by Glacier and the this provider resource will show as needing recreation. Changing this from `false` to `true` will show as resource recreation, which is expected. Changing this from `true` to `false` is not possible unless the Glacier Vault is recreated at the same time.
        /// </summary>
        [Input("completeLock")]
        public Input<bool>? CompleteLock { get; set; }

        /// <summary>
        /// Allow this provider to ignore the error returned when attempting to delete the Glacier Lock Policy. This can be used to delete or recreate the Glacier Vault via this provider, for example, if the Glacier Vault Lock policy permits that action. This should only be used in conjunction with `complete_lock` being set to `true`.
        /// </summary>
        [Input("ignoreDeletionError")]
        public Input<bool>? IgnoreDeletionError { get; set; }

        /// <summary>
        /// JSON string containing the IAM policy to apply as the Glacier Vault Lock policy.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The name of the Glacier Vault.
        /// </summary>
        [Input("vaultName")]
        public Input<string>? VaultName { get; set; }

        public VaultLockState()
        {
        }
    }
}
