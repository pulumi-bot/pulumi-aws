// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Kms
{
    /// <summary>
    /// Provides a KMS customer master key.
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
    ///         var key = new Aws.Kms.Key("key", new Aws.Kms.KeyArgs
    ///         {
    ///             DeletionWindowInDays = 10,
    ///             Description = "KMS key 1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// KMS Keys can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:kms/key:Key a 1234abcd-12ab-34cd-56ef-1234567890ab
    /// ```
    /// </summary>
    public partial class Key : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the key.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
        /// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
        /// </summary>
        [Output("customerMasterKeySpec")]
        public Output<string?> CustomerMasterKeySpec { get; private set; } = null!;

        /// <summary>
        /// Duration in days after which the key is deleted
        /// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
        /// </summary>
        [Output("deletionWindowInDays")]
        public Output<int?> DeletionWindowInDays { get; private set; } = null!;

        /// <summary>
        /// The description of the key as viewed in AWS console.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
        /// is enabled. Defaults to false.
        /// </summary>
        [Output("enableKeyRotation")]
        public Output<bool?> EnableKeyRotation { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the key is enabled. Defaults to true.
        /// </summary>
        [Output("isEnabled")]
        public Output<bool?> IsEnabled { get; private set; } = null!;

        /// <summary>
        /// The globally unique identifier for the key.
        /// </summary>
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;

        /// <summary>
        /// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT` or `SIGN_VERIFY`.
        /// Defaults to `ENCRYPT_DECRYPT`.
        /// </summary>
        [Output("keyUsage")]
        public Output<string?> KeyUsage { get; private set; } = null!;

        /// <summary>
        /// A valid policy JSON document.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the object.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Key resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Key(string name, KeyArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:kms/key:Key", name, args ?? new KeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Key(string name, Input<string> id, KeyState? state = null, CustomResourceOptions? options = null)
            : base("aws:kms/key:Key", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Key resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Key Get(string name, Input<string> id, KeyState? state = null, CustomResourceOptions? options = null)
        {
            return new Key(name, id, state, options);
        }
    }

    public sealed class KeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
        /// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
        /// </summary>
        [Input("customerMasterKeySpec")]
        public Input<string>? CustomerMasterKeySpec { get; set; }

        /// <summary>
        /// Duration in days after which the key is deleted
        /// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
        /// </summary>
        [Input("deletionWindowInDays")]
        public Input<int>? DeletionWindowInDays { get; set; }

        /// <summary>
        /// The description of the key as viewed in AWS console.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
        /// is enabled. Defaults to false.
        /// </summary>
        [Input("enableKeyRotation")]
        public Input<bool>? EnableKeyRotation { get; set; }

        /// <summary>
        /// Specifies whether the key is enabled. Defaults to true.
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        /// <summary>
        /// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT` or `SIGN_VERIFY`.
        /// Defaults to `ENCRYPT_DECRYPT`.
        /// </summary>
        [Input("keyUsage")]
        public Input<string>? KeyUsage { get; set; }

        /// <summary>
        /// A valid policy JSON document.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the object.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public KeyArgs()
        {
        }
    }

    public sealed class KeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the key.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
        /// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
        /// </summary>
        [Input("customerMasterKeySpec")]
        public Input<string>? CustomerMasterKeySpec { get; set; }

        /// <summary>
        /// Duration in days after which the key is deleted
        /// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
        /// </summary>
        [Input("deletionWindowInDays")]
        public Input<int>? DeletionWindowInDays { get; set; }

        /// <summary>
        /// The description of the key as viewed in AWS console.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
        /// is enabled. Defaults to false.
        /// </summary>
        [Input("enableKeyRotation")]
        public Input<bool>? EnableKeyRotation { get; set; }

        /// <summary>
        /// Specifies whether the key is enabled. Defaults to true.
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        /// <summary>
        /// The globally unique identifier for the key.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT` or `SIGN_VERIFY`.
        /// Defaults to `ENCRYPT_DECRYPT`.
        /// </summary>
        [Input("keyUsage")]
        public Input<string>? KeyUsage { get; set; }

        /// <summary>
        /// A valid policy JSON document.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the object.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public KeyState()
        {
        }
    }
}
