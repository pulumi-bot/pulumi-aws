// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/acmpca_certificate_authority.html.markdown.
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
            : base("aws:acmpca/certificateAuthority:CertificateAuthority", name, args, MakeResourceOptions(options, ""))
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

    namespace Inputs
    {

    public sealed class CertificateAuthorityCertificateAuthorityConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of the public key algorithm and size, in bits, of the key pair that your key pair creates when it issues a certificate. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
        /// </summary>
        [Input("keyAlgorithm", required: true)]
        public Input<string> KeyAlgorithm { get; set; } = null!;

        /// <summary>
        /// Name of the algorithm your private CA uses to sign certificate requests. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
        /// </summary>
        [Input("signingAlgorithm", required: true)]
        public Input<string> SigningAlgorithm { get; set; } = null!;

        /// <summary>
        /// Nested argument that contains X.500 distinguished name information. At least one nested attribute must be specified.
        /// </summary>
        [Input("subject", required: true)]
        public Input<CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs> Subject { get; set; } = null!;

        public CertificateAuthorityCertificateAuthorityConfigurationArgs()
        {
        }
    }

    public sealed class CertificateAuthorityCertificateAuthorityConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of the public key algorithm and size, in bits, of the key pair that your key pair creates when it issues a certificate. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
        /// </summary>
        [Input("keyAlgorithm", required: true)]
        public Input<string> KeyAlgorithm { get; set; } = null!;

        /// <summary>
        /// Name of the algorithm your private CA uses to sign certificate requests. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
        /// </summary>
        [Input("signingAlgorithm", required: true)]
        public Input<string> SigningAlgorithm { get; set; } = null!;

        /// <summary>
        /// Nested argument that contains X.500 distinguished name information. At least one nested attribute must be specified.
        /// </summary>
        [Input("subject", required: true)]
        public Input<CertificateAuthorityCertificateAuthorityConfigurationSubjectGetArgs> Subject { get; set; } = null!;

        public CertificateAuthorityCertificateAuthorityConfigurationGetArgs()
        {
        }
    }

    public sealed class CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fully qualified domain name (FQDN) associated with the certificate subject.
        /// </summary>
        [Input("commonName")]
        public Input<string>? CommonName { get; set; }

        /// <summary>
        /// Two digit code that specifies the country in which the certificate subject located.
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Disambiguating information for the certificate subject.
        /// </summary>
        [Input("distinguishedNameQualifier")]
        public Input<string>? DistinguishedNameQualifier { get; set; }

        /// <summary>
        /// Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third.
        /// </summary>
        [Input("generationQualifier")]
        public Input<string>? GenerationQualifier { get; set; }

        /// <summary>
        /// First name.
        /// </summary>
        [Input("givenName")]
        public Input<string>? GivenName { get; set; }

        /// <summary>
        /// Concatenation that typically contains the first letter of the `given_name`, the first letter of the middle name if one exists, and the first letter of the `surname`.
        /// </summary>
        [Input("initials")]
        public Input<string>? Initials { get; set; }

        /// <summary>
        /// The locality (such as a city or town) in which the certificate subject is located.
        /// </summary>
        [Input("locality")]
        public Input<string>? Locality { get; set; }

        /// <summary>
        /// Legal name of the organization with which the certificate subject is affiliated.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
        /// </summary>
        [Input("organizationalUnit")]
        public Input<string>? OrganizationalUnit { get; set; }

        /// <summary>
        /// Typically a shortened version of a longer `given_name`. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
        /// </summary>
        [Input("pseudonym")]
        public Input<string>? Pseudonym { get; set; }

        /// <summary>
        /// State in which the subject of the certificate is located.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Family name. In the US and the UK for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.
        /// </summary>
        [Input("surname")]
        public Input<string>? Surname { get; set; }

        /// <summary>
        /// A title such as Mr. or Ms. which is pre-pended to the name to refer formally to the certificate subject.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public CertificateAuthorityCertificateAuthorityConfigurationSubjectArgs()
        {
        }
    }

    public sealed class CertificateAuthorityCertificateAuthorityConfigurationSubjectGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fully qualified domain name (FQDN) associated with the certificate subject.
        /// </summary>
        [Input("commonName")]
        public Input<string>? CommonName { get; set; }

        /// <summary>
        /// Two digit code that specifies the country in which the certificate subject located.
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Disambiguating information for the certificate subject.
        /// </summary>
        [Input("distinguishedNameQualifier")]
        public Input<string>? DistinguishedNameQualifier { get; set; }

        /// <summary>
        /// Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third.
        /// </summary>
        [Input("generationQualifier")]
        public Input<string>? GenerationQualifier { get; set; }

        /// <summary>
        /// First name.
        /// </summary>
        [Input("givenName")]
        public Input<string>? GivenName { get; set; }

        /// <summary>
        /// Concatenation that typically contains the first letter of the `given_name`, the first letter of the middle name if one exists, and the first letter of the `surname`.
        /// </summary>
        [Input("initials")]
        public Input<string>? Initials { get; set; }

        /// <summary>
        /// The locality (such as a city or town) in which the certificate subject is located.
        /// </summary>
        [Input("locality")]
        public Input<string>? Locality { get; set; }

        /// <summary>
        /// Legal name of the organization with which the certificate subject is affiliated.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
        /// </summary>
        [Input("organizationalUnit")]
        public Input<string>? OrganizationalUnit { get; set; }

        /// <summary>
        /// Typically a shortened version of a longer `given_name`. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
        /// </summary>
        [Input("pseudonym")]
        public Input<string>? Pseudonym { get; set; }

        /// <summary>
        /// State in which the subject of the certificate is located.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Family name. In the US and the UK for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.
        /// </summary>
        [Input("surname")]
        public Input<string>? Surname { get; set; }

        /// <summary>
        /// A title such as Mr. or Ms. which is pre-pended to the name to refer formally to the certificate subject.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public CertificateAuthorityCertificateAuthorityConfigurationSubjectGetArgs()
        {
        }
    }

    public sealed class CertificateAuthorityRevocationConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Nested argument containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority. Defined below.
        /// </summary>
        [Input("crlConfiguration")]
        public Input<CertificateAuthorityRevocationConfigurationCrlConfigurationArgs>? CrlConfiguration { get; set; }

        public CertificateAuthorityRevocationConfigurationArgs()
        {
        }
    }

    public sealed class CertificateAuthorityRevocationConfigurationCrlConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point. Use this value if you don't want the name of your S3 bucket to be public.
        /// </summary>
        [Input("customCname")]
        public Input<string>? CustomCname { get; set; }

        /// <summary>
        /// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Number of days until a certificate expires. Must be between 1 and 5000.
        /// </summary>
        [Input("expirationInDays", required: true)]
        public Input<int> ExpirationInDays { get; set; } = null!;

        /// <summary>
        /// Name of the S3 bucket that contains the CRL. If you do not provide a value for the `custom_cname` argument, the name of your S3 bucket is placed into the CRL Distribution Points extension of the issued certificate. You must specify a bucket policy that allows ACM PCA to write the CRL to your bucket.
        /// </summary>
        [Input("s3BucketName")]
        public Input<string>? S3BucketName { get; set; }

        public CertificateAuthorityRevocationConfigurationCrlConfigurationArgs()
        {
        }
    }

    public sealed class CertificateAuthorityRevocationConfigurationCrlConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point. Use this value if you don't want the name of your S3 bucket to be public.
        /// </summary>
        [Input("customCname")]
        public Input<string>? CustomCname { get; set; }

        /// <summary>
        /// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Number of days until a certificate expires. Must be between 1 and 5000.
        /// </summary>
        [Input("expirationInDays", required: true)]
        public Input<int> ExpirationInDays { get; set; } = null!;

        /// <summary>
        /// Name of the S3 bucket that contains the CRL. If you do not provide a value for the `custom_cname` argument, the name of your S3 bucket is placed into the CRL Distribution Points extension of the issued certificate. You must specify a bucket policy that allows ACM PCA to write the CRL to your bucket.
        /// </summary>
        [Input("s3BucketName")]
        public Input<string>? S3BucketName { get; set; }

        public CertificateAuthorityRevocationConfigurationCrlConfigurationGetArgs()
        {
        }
    }

    public sealed class CertificateAuthorityRevocationConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Nested argument containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority. Defined below.
        /// </summary>
        [Input("crlConfiguration")]
        public Input<CertificateAuthorityRevocationConfigurationCrlConfigurationGetArgs>? CrlConfiguration { get; set; }

        public CertificateAuthorityRevocationConfigurationGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class CertificateAuthorityCertificateAuthorityConfiguration
    {
        /// <summary>
        /// Type of the public key algorithm and size, in bits, of the key pair that your key pair creates when it issues a certificate. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
        /// </summary>
        public readonly string KeyAlgorithm;
        /// <summary>
        /// Name of the algorithm your private CA uses to sign certificate requests. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
        /// </summary>
        public readonly string SigningAlgorithm;
        /// <summary>
        /// Nested argument that contains X.500 distinguished name information. At least one nested attribute must be specified.
        /// </summary>
        public readonly CertificateAuthorityCertificateAuthorityConfigurationSubject Subject;

        [OutputConstructor]
        private CertificateAuthorityCertificateAuthorityConfiguration(
            string keyAlgorithm,
            string signingAlgorithm,
            CertificateAuthorityCertificateAuthorityConfigurationSubject subject)
        {
            KeyAlgorithm = keyAlgorithm;
            SigningAlgorithm = signingAlgorithm;
            Subject = subject;
        }
    }

    [OutputType]
    public sealed class CertificateAuthorityCertificateAuthorityConfigurationSubject
    {
        /// <summary>
        /// Fully qualified domain name (FQDN) associated with the certificate subject.
        /// </summary>
        public readonly string? CommonName;
        /// <summary>
        /// Two digit code that specifies the country in which the certificate subject located.
        /// </summary>
        public readonly string? Country;
        /// <summary>
        /// Disambiguating information for the certificate subject.
        /// </summary>
        public readonly string? DistinguishedNameQualifier;
        /// <summary>
        /// Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third.
        /// </summary>
        public readonly string? GenerationQualifier;
        /// <summary>
        /// First name.
        /// </summary>
        public readonly string? GivenName;
        /// <summary>
        /// Concatenation that typically contains the first letter of the `given_name`, the first letter of the middle name if one exists, and the first letter of the `surname`.
        /// </summary>
        public readonly string? Initials;
        /// <summary>
        /// The locality (such as a city or town) in which the certificate subject is located.
        /// </summary>
        public readonly string? Locality;
        /// <summary>
        /// Legal name of the organization with which the certificate subject is affiliated.
        /// </summary>
        public readonly string? Organization;
        /// <summary>
        /// A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
        /// </summary>
        public readonly string? OrganizationalUnit;
        /// <summary>
        /// Typically a shortened version of a longer `given_name`. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
        /// </summary>
        public readonly string? Pseudonym;
        /// <summary>
        /// State in which the subject of the certificate is located.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// Family name. In the US and the UK for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.
        /// </summary>
        public readonly string? Surname;
        /// <summary>
        /// A title such as Mr. or Ms. which is pre-pended to the name to refer formally to the certificate subject.
        /// </summary>
        public readonly string? Title;

        [OutputConstructor]
        private CertificateAuthorityCertificateAuthorityConfigurationSubject(
            string? commonName,
            string? country,
            string? distinguishedNameQualifier,
            string? generationQualifier,
            string? givenName,
            string? initials,
            string? locality,
            string? organization,
            string? organizationalUnit,
            string? pseudonym,
            string? state,
            string? surname,
            string? title)
        {
            CommonName = commonName;
            Country = country;
            DistinguishedNameQualifier = distinguishedNameQualifier;
            GenerationQualifier = generationQualifier;
            GivenName = givenName;
            Initials = initials;
            Locality = locality;
            Organization = organization;
            OrganizationalUnit = organizationalUnit;
            Pseudonym = pseudonym;
            State = state;
            Surname = surname;
            Title = title;
        }
    }

    [OutputType]
    public sealed class CertificateAuthorityRevocationConfiguration
    {
        /// <summary>
        /// Nested argument containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority. Defined below.
        /// </summary>
        public readonly CertificateAuthorityRevocationConfigurationCrlConfiguration? CrlConfiguration;

        [OutputConstructor]
        private CertificateAuthorityRevocationConfiguration(CertificateAuthorityRevocationConfigurationCrlConfiguration? crlConfiguration)
        {
            CrlConfiguration = crlConfiguration;
        }
    }

    [OutputType]
    public sealed class CertificateAuthorityRevocationConfigurationCrlConfiguration
    {
        /// <summary>
        /// Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point. Use this value if you don't want the name of your S3 bucket to be public.
        /// </summary>
        public readonly string? CustomCname;
        /// <summary>
        /// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Number of days until a certificate expires. Must be between 1 and 5000.
        /// </summary>
        public readonly int ExpirationInDays;
        /// <summary>
        /// Name of the S3 bucket that contains the CRL. If you do not provide a value for the `custom_cname` argument, the name of your S3 bucket is placed into the CRL Distribution Points extension of the issued certificate. You must specify a bucket policy that allows ACM PCA to write the CRL to your bucket.
        /// </summary>
        public readonly string? S3BucketName;

        [OutputConstructor]
        private CertificateAuthorityRevocationConfigurationCrlConfiguration(
            string? customCname,
            bool? enabled,
            int expirationInDays,
            string? s3BucketName)
        {
            CustomCname = customCname;
            Enabled = enabled;
            ExpirationInDays = expirationInDays;
            S3BucketName = s3BucketName;
        }
    }
    }
}
