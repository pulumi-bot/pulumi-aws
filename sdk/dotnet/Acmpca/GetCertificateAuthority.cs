// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Acmpca
{
    public static partial class GetCertificateAuthority
    {
        /// <summary>
        /// Get information on a AWS Certificate Manager Private Certificate Authority (ACM PCA Certificate Authority).
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/acmpca_certificate_authority.html.markdown.
        /// </summary>
        public static Task<GetCertificateAuthorityResult> InvokeAsync(GetCertificateAuthorityArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateAuthorityResult>("aws:acmpca/getCertificateAuthority:getCertificateAuthority", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetCertificateAuthorityArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the certificate authority.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        [Input("revocationConfigurations")]
        private List<Inputs.GetCertificateAuthorityRevocationConfigurationsArgs>? _revocationConfigurations;
        public List<Inputs.GetCertificateAuthorityRevocationConfigurationsArgs> RevocationConfigurations
        {
            get => _revocationConfigurations ?? (_revocationConfigurations = new List<Inputs.GetCertificateAuthorityRevocationConfigurationsArgs>());
            set => _revocationConfigurations = value;
        }

        [Input("tags")]
        private Dictionary<string, object>? _tags;
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetCertificateAuthorityArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetCertificateAuthorityResult
    {
        public readonly string Arn;
        /// <summary>
        /// Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
        /// </summary>
        public readonly string Certificate;
        /// <summary>
        /// Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
        /// </summary>
        public readonly string CertificateChain;
        /// <summary>
        /// The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
        /// </summary>
        public readonly string CertificateSigningRequest;
        /// <summary>
        /// Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        /// </summary>
        public readonly string NotAfter;
        /// <summary>
        /// Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
        /// </summary>
        public readonly string NotBefore;
        /// <summary>
        /// Nested attribute containing revocation configuration.
        /// * `revocation_configuration.0.crl_configuration` - Nested attribute containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority.
        /// * `revocation_configuration.0.crl_configuration.0.custom_cname` - Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point.
        /// * `revocation_configuration.0.crl_configuration.0.enabled` - Boolean value that specifies whether certificate revocation lists (CRLs) are enabled.
        /// * `revocation_configuration.0.crl_configuration.0.expiration_in_days` - Number of days until a certificate expires.
        /// * `revocation_configuration.0.crl_configuration.0.s3_bucket_name` - Name of the S3 bucket that contains the CRL.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCertificateAuthorityRevocationConfigurationsResult> RevocationConfigurations;
        /// <summary>
        /// Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
        /// </summary>
        public readonly string Serial;
        /// <summary>
        /// Status of the certificate authority.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Specifies a key-value map of user-defined tags that are attached to the certificate authority.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The type of the certificate authority.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetCertificateAuthorityResult(
            string arn,
            string certificate,
            string certificateChain,
            string certificateSigningRequest,
            string notAfter,
            string notBefore,
            ImmutableArray<Outputs.GetCertificateAuthorityRevocationConfigurationsResult> revocationConfigurations,
            string serial,
            string status,
            ImmutableDictionary<string, object> tags,
            string type,
            string id)
        {
            Arn = arn;
            Certificate = certificate;
            CertificateChain = certificateChain;
            CertificateSigningRequest = certificateSigningRequest;
            NotAfter = notAfter;
            NotBefore = notBefore;
            RevocationConfigurations = revocationConfigurations;
            Serial = serial;
            Status = status;
            Tags = tags;
            Type = type;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetCertificateAuthorityRevocationConfigurationsArgs : Pulumi.InvokeArgs
    {
        [Input("crlConfigurations")]
        private List<GetCertificateAuthorityRevocationConfigurationsCrlConfigurationsArgs>? _crlConfigurations;
        public List<GetCertificateAuthorityRevocationConfigurationsCrlConfigurationsArgs> CrlConfigurations
        {
            get => _crlConfigurations ?? (_crlConfigurations = new List<GetCertificateAuthorityRevocationConfigurationsCrlConfigurationsArgs>());
            set => _crlConfigurations = value;
        }

        public GetCertificateAuthorityRevocationConfigurationsArgs()
        {
        }
    }

    public sealed class GetCertificateAuthorityRevocationConfigurationsCrlConfigurationsArgs : Pulumi.InvokeArgs
    {
        [Input("customCname")]
        public string? CustomCname { get; set; }

        [Input("enabled")]
        public bool? Enabled { get; set; }

        [Input("expirationInDays")]
        public int? ExpirationInDays { get; set; }

        [Input("s3BucketName")]
        public string? S3BucketName { get; set; }

        public GetCertificateAuthorityRevocationConfigurationsCrlConfigurationsArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetCertificateAuthorityRevocationConfigurationsCrlConfigurationsResult
    {
        public readonly string CustomCname;
        public readonly bool Enabled;
        public readonly int ExpirationInDays;
        public readonly string S3BucketName;

        [OutputConstructor]
        private GetCertificateAuthorityRevocationConfigurationsCrlConfigurationsResult(
            string customCname,
            bool enabled,
            int expirationInDays,
            string s3BucketName)
        {
            CustomCname = customCname;
            Enabled = enabled;
            ExpirationInDays = expirationInDays;
            S3BucketName = s3BucketName;
        }
    }

    [OutputType]
    public sealed class GetCertificateAuthorityRevocationConfigurationsResult
    {
        public readonly ImmutableArray<GetCertificateAuthorityRevocationConfigurationsCrlConfigurationsResult> CrlConfigurations;

        [OutputConstructor]
        private GetCertificateAuthorityRevocationConfigurationsResult(ImmutableArray<GetCertificateAuthorityRevocationConfigurationsCrlConfigurationsResult> crlConfigurations)
        {
            CrlConfigurations = crlConfigurations;
        }
    }
    }
}
