// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package certificateAuthority

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/acmpca/CertificateAuthorityCertificateAuthorityConfiguration"
	"https:/github.com/pulumi/pulumi-aws/acmpca/CertificateAuthorityCertificateAuthorityConfigurationSubject"
	"https:/github.com/pulumi/pulumi-aws/acmpca/CertificateAuthorityRevocationConfiguration"
	"https:/github.com/pulumi/pulumi-aws/acmpca/CertificateAuthorityRevocationConfigurationCrlConfiguration"
)

// Provides a resource to manage AWS Certificate Manager Private Certificate Authorities (ACM PCA Certificate Authorities).
// 
// > **NOTE:** Creating this resource will leave the certificate authority in a `PENDING_CERTIFICATE` status, which means it cannot yet issue certificates. To complete this setup, you must fully sign the certificate authority CSR available in the `certificateSigningRequest` attribute and import the signed certificate using the AWS SDK, CLI or Console. This provider can support another resource to manage that workflow automatically in the future.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/acmpca_certificate_authority.html.markdown.
type CertificateAuthority struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the certificate authority.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Nested argument containing algorithms and certificate subject information. Defined below.
	CertificateAuthorityConfiguration acmpcaCertificateAuthorityCertificateAuthorityConfiguration.CertificateAuthorityCertificateAuthorityConfigurationOutput `pulumi:"certificateAuthorityConfiguration"`
	// Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
	CertificateChain pulumi.StringOutput `pulumi:"certificateChain"`
	// The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
	CertificateSigningRequest pulumi.StringOutput `pulumi:"certificateSigningRequest"`
	// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotAfter pulumi.StringOutput `pulumi:"notAfter"`
	// Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotBefore pulumi.StringOutput `pulumi:"notBefore"`
	// The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
	PermanentDeletionTimeInDays pulumi.IntPtrOutput `pulumi:"permanentDeletionTimeInDays"`
	// Nested argument containing revocation configuration. Defined below.
	RevocationConfiguration acmpcaCertificateAuthorityRevocationConfiguration.CertificateAuthorityRevocationConfigurationPtrOutput `pulumi:"revocationConfiguration"`
	// Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
	Serial pulumi.StringOutput `pulumi:"serial"`
	// Status of the certificate authority.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies a key-value map of user-defined tags that are attached to the certificate authority.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewCertificateAuthority registers a new resource with the given unique name, arguments, and options.
func NewCertificateAuthority(ctx *pulumi.Context,
	name string, args *CertificateAuthorityArgs, opts ...pulumi.ResourceOption) (*CertificateAuthority, error) {
	if args == nil || args.CertificateAuthorityConfiguration == nil {
		return nil, errors.New("missing required argument 'CertificateAuthorityConfiguration'")
	}
	if args == nil {
		args = &CertificateAuthorityArgs{}
	}
	var resource CertificateAuthority
	err := ctx.RegisterResource("aws:acmpca/certificateAuthority:CertificateAuthority", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateAuthority gets an existing CertificateAuthority resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateAuthority(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateAuthorityState, opts ...pulumi.ResourceOption) (*CertificateAuthority, error) {
	var resource CertificateAuthority
	err := ctx.ReadResource("aws:acmpca/certificateAuthority:CertificateAuthority", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateAuthority resources.
type certificateAuthorityState struct {
	// Amazon Resource Name (ARN) of the certificate authority.
	Arn *string `pulumi:"arn"`
	// Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
	Certificate *string `pulumi:"certificate"`
	// Nested argument containing algorithms and certificate subject information. Defined below.
	CertificateAuthorityConfiguration *acmpcaCertificateAuthorityCertificateAuthorityConfiguration.CertificateAuthorityCertificateAuthorityConfiguration `pulumi:"certificateAuthorityConfiguration"`
	// Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
	CertificateChain *string `pulumi:"certificateChain"`
	// The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
	CertificateSigningRequest *string `pulumi:"certificateSigningRequest"`
	// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
	Enabled *bool `pulumi:"enabled"`
	// Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotAfter *string `pulumi:"notAfter"`
	// Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotBefore *string `pulumi:"notBefore"`
	// The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
	PermanentDeletionTimeInDays *int `pulumi:"permanentDeletionTimeInDays"`
	// Nested argument containing revocation configuration. Defined below.
	RevocationConfiguration *acmpcaCertificateAuthorityRevocationConfiguration.CertificateAuthorityRevocationConfiguration `pulumi:"revocationConfiguration"`
	// Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
	Serial *string `pulumi:"serial"`
	// Status of the certificate authority.
	Status *string `pulumi:"status"`
	// Specifies a key-value map of user-defined tags that are attached to the certificate authority.
	Tags map[string]interface{} `pulumi:"tags"`
	// The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
	Type *string `pulumi:"type"`
}

type CertificateAuthorityState struct {
	// Amazon Resource Name (ARN) of the certificate authority.
	Arn pulumi.StringPtrInput
	// Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
	Certificate pulumi.StringPtrInput
	// Nested argument containing algorithms and certificate subject information. Defined below.
	CertificateAuthorityConfiguration acmpcaCertificateAuthorityCertificateAuthorityConfiguration.CertificateAuthorityCertificateAuthorityConfigurationPtrInput
	// Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
	CertificateChain pulumi.StringPtrInput
	// The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
	CertificateSigningRequest pulumi.StringPtrInput
	// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
	Enabled pulumi.BoolPtrInput
	// Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotAfter pulumi.StringPtrInput
	// Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
	NotBefore pulumi.StringPtrInput
	// The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
	PermanentDeletionTimeInDays pulumi.IntPtrInput
	// Nested argument containing revocation configuration. Defined below.
	RevocationConfiguration acmpcaCertificateAuthorityRevocationConfiguration.CertificateAuthorityRevocationConfigurationPtrInput
	// Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
	Serial pulumi.StringPtrInput
	// Status of the certificate authority.
	Status pulumi.StringPtrInput
	// Specifies a key-value map of user-defined tags that are attached to the certificate authority.
	Tags pulumi.MapInput
	// The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
	Type pulumi.StringPtrInput
}

func (CertificateAuthorityState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityState)(nil)).Elem()
}

type certificateAuthorityArgs struct {
	// Nested argument containing algorithms and certificate subject information. Defined below.
	CertificateAuthorityConfiguration acmpcaCertificateAuthorityCertificateAuthorityConfiguration.CertificateAuthorityCertificateAuthorityConfiguration `pulumi:"certificateAuthorityConfiguration"`
	// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
	Enabled *bool `pulumi:"enabled"`
	// The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
	PermanentDeletionTimeInDays *int `pulumi:"permanentDeletionTimeInDays"`
	// Nested argument containing revocation configuration. Defined below.
	RevocationConfiguration *acmpcaCertificateAuthorityRevocationConfiguration.CertificateAuthorityRevocationConfiguration `pulumi:"revocationConfiguration"`
	// Specifies a key-value map of user-defined tags that are attached to the certificate authority.
	Tags map[string]interface{} `pulumi:"tags"`
	// The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a CertificateAuthority resource.
type CertificateAuthorityArgs struct {
	// Nested argument containing algorithms and certificate subject information. Defined below.
	CertificateAuthorityConfiguration acmpcaCertificateAuthorityCertificateAuthorityConfiguration.CertificateAuthorityCertificateAuthorityConfigurationInput
	// Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
	Enabled pulumi.BoolPtrInput
	// The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
	PermanentDeletionTimeInDays pulumi.IntPtrInput
	// Nested argument containing revocation configuration. Defined below.
	RevocationConfiguration acmpcaCertificateAuthorityRevocationConfiguration.CertificateAuthorityRevocationConfigurationPtrInput
	// Specifies a key-value map of user-defined tags that are attached to the certificate authority.
	Tags pulumi.MapInput
	// The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
	Type pulumi.StringPtrInput
}

func (CertificateAuthorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityArgs)(nil)).Elem()
}

