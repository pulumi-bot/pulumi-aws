// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Acmpca
{
    /// <summary>
    /// Provides a resource to manage AWS Certificate Manager Private Certificate Authorities (ACM PCA Certificate Authorities).
    /// 
    /// &gt; **NOTE:** Creating this resource will leave the certificate authority in a `PENDING_CERTIFICATE` status, which means it cannot yet issue certificates. To complete this setup, you must fully sign the certificate authority CSR available in the `certificate_signing_request` attribute and import the signed certificate using the AWS SDK, CLI or Console. This provider can support another resource to manage that workflow automatically in the future.
    /// 
    /// ## Example Usage
    /// </summary>
    public partial class CertificateAuthority : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the certificate authority.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing algorithms and certificate subject information. Defined below.
        /// </summary>
        [Output("certificateAuthorityConfiguration")]
        public Output<Outputs.CertificateAuthorityCertificateAuthorityConfiguration> CertificateAuthorityConfiguration { get; private set; } = null!;

        /// <summary>
        /// Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
        /// </summary>
        [Output("certificateChain")]
        public Output<string> CertificateChain { get; private set; } = null!;

        /// <summary>
        /// The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
        /// </summary>
        [Output("certificateSigningRequest")]
        public Output<string> CertificateSigningRequest { get; private set; } = null!;

        /// <summary>
        /// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        /// </summary>
        [Output("notAfter")]
        public Output<string> NotAfter { get; private set; } = null!;

        /// <summary>
        /// Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        /// </summary>
        [Output("notBefore")]
        public Output<string> NotBefore { get; private set; } = null!;

        /// <summary>
        /// The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
        /// </summary>
        [Output("permanentDeletionTimeInDays")]
        public Output<int?> PermanentDeletionTimeInDays { get; private set; } = null!;

        /// <summary>
        /// Nested argument containing revocation configuration. Defined below.
        /// </summary>
        [Output("revocationConfiguration")]
        public Output<Outputs.CertificateAuthorityRevocationConfiguration?> RevocationConfiguration { get; private set; } = null!;

        /// <summary>
        /// Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
        /// </summary>
        [Output("serial")]
        public Output<string> Serial { get; private set; } = null!;

        /// <summary>
        /// Status of the certificate authority.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies a key-value map of user-defined tags that are attached to the certificate authority.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateAuthority resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateAuthority(string name, CertificateAuthorityArgs args, CustomResourceOptions? options = null)
            : base("aws:acmpca/certificateAuthority:CertificateAuthority", name, args ?? new CertificateAuthorityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateAuthority(string name, Input<string> id, CertificateAuthorityState? state = null, CustomResourceOptions? options = null)
            : base("aws:acmpca/certificateAuthority:CertificateAuthority", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CertificateAuthority resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateAuthority Get(string name, Input<string> id, CertificateAuthorityState? state = null, CustomResourceOptions? options = null)
        {
            return new CertificateAuthority(name, id, state, options);
        }
    }

    public sealed class CertificateAuthorityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Nested argument containing algorithms and certificate subject information. Defined below.
        /// </summary>
        [Input("certificateAuthorityConfiguration", required: true)]
        public Input<Inputs.CertificateAuthorityCertificateAuthorityConfigurationArgs> CertificateAuthorityConfiguration { get; set; } = null!;

        /// <summary>
        /// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
        /// </summary>
        [Input("permanentDeletionTimeInDays")]
        public Input<int>? PermanentDeletionTimeInDays { get; set; }

        /// <summary>
        /// Nested argument containing revocation configuration. Defined below.
        /// </summary>
        [Input("revocationConfiguration")]
        public Input<Inputs.CertificateAuthorityRevocationConfigurationArgs>? RevocationConfiguration { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Specifies a key-value map of user-defined tags that are attached to the certificate authority.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public CertificateAuthorityArgs()
        {
        }
    }

    public sealed class CertificateAuthorityState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the certificate authority.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// Nested argument containing algorithms and certificate subject information. Defined below.
        /// </summary>
        [Input("certificateAuthorityConfiguration")]
        public Input<Inputs.CertificateAuthorityCertificateAuthorityConfigurationGetArgs>? CertificateAuthorityConfiguration { get; set; }

        /// <summary>
        /// Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
        /// </summary>
        [Input("certificateSigningRequest")]
        public Input<string>? CertificateSigningRequest { get; set; }

        /// <summary>
        /// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        /// </summary>
        [Input("notAfter")]
        public Input<string>? NotAfter { get; set; }

        /// <summary>
        /// Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        /// </summary>
        [Input("notBefore")]
        public Input<string>? NotBefore { get; set; }

        /// <summary>
        /// The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
        /// </summary>
        [Input("permanentDeletionTimeInDays")]
        public Input<int>? PermanentDeletionTimeInDays { get; set; }

        /// <summary>
        /// Nested argument containing revocation configuration. Defined below.
        /// </summary>
        [Input("revocationConfiguration")]
        public Input<Inputs.CertificateAuthorityRevocationConfigurationGetArgs>? RevocationConfiguration { get; set; }

        /// <summary>
        /// Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
        /// </summary>
        [Input("serial")]
        public Input<string>? Serial { get; set; }

        /// <summary>
        /// Status of the certificate authority.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Specifies a key-value map of user-defined tags that are attached to the certificate authority.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public CertificateAuthorityState()
        {
        }
    }
}
