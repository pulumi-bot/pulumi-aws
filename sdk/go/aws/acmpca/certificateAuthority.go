// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package acmpca

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type CertificateAuthority struct {
	pulumi.CustomResourceState

	Arn                               pulumi.StringOutput                                         `pulumi:"arn"`
	Certificate                       pulumi.StringOutput                                         `pulumi:"certificate"`
	CertificateAuthorityConfiguration CertificateAuthorityCertificateAuthorityConfigurationOutput `pulumi:"certificateAuthorityConfiguration"`
	CertificateChain                  pulumi.StringOutput                                         `pulumi:"certificateChain"`
	CertificateSigningRequest         pulumi.StringOutput                                         `pulumi:"certificateSigningRequest"`
	Enabled                           pulumi.BoolPtrOutput                                        `pulumi:"enabled"`
	NotAfter                          pulumi.StringOutput                                         `pulumi:"notAfter"`
	NotBefore                         pulumi.StringOutput                                         `pulumi:"notBefore"`
	PermanentDeletionTimeInDays       pulumi.IntPtrOutput                                         `pulumi:"permanentDeletionTimeInDays"`
	RevocationConfiguration           CertificateAuthorityRevocationConfigurationPtrOutput        `pulumi:"revocationConfiguration"`
	Serial                            pulumi.StringOutput                                         `pulumi:"serial"`
	Status                            pulumi.StringOutput                                         `pulumi:"status"`
	Tags                              pulumi.StringMapOutput                                      `pulumi:"tags"`
	Type                              pulumi.StringPtrOutput                                      `pulumi:"type"`
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
	Arn                               *string                                                `pulumi:"arn"`
	Certificate                       *string                                                `pulumi:"certificate"`
	CertificateAuthorityConfiguration *CertificateAuthorityCertificateAuthorityConfiguration `pulumi:"certificateAuthorityConfiguration"`
	CertificateChain                  *string                                                `pulumi:"certificateChain"`
	CertificateSigningRequest         *string                                                `pulumi:"certificateSigningRequest"`
	Enabled                           *bool                                                  `pulumi:"enabled"`
	NotAfter                          *string                                                `pulumi:"notAfter"`
	NotBefore                         *string                                                `pulumi:"notBefore"`
	PermanentDeletionTimeInDays       *int                                                   `pulumi:"permanentDeletionTimeInDays"`
	RevocationConfiguration           *CertificateAuthorityRevocationConfiguration           `pulumi:"revocationConfiguration"`
	Serial                            *string                                                `pulumi:"serial"`
	Status                            *string                                                `pulumi:"status"`
	Tags                              map[string]string                                      `pulumi:"tags"`
	Type                              *string                                                `pulumi:"type"`
}

type CertificateAuthorityState struct {
	Arn                               pulumi.StringPtrInput
	Certificate                       pulumi.StringPtrInput
	CertificateAuthorityConfiguration CertificateAuthorityCertificateAuthorityConfigurationPtrInput
	CertificateChain                  pulumi.StringPtrInput
	CertificateSigningRequest         pulumi.StringPtrInput
	Enabled                           pulumi.BoolPtrInput
	NotAfter                          pulumi.StringPtrInput
	NotBefore                         pulumi.StringPtrInput
	PermanentDeletionTimeInDays       pulumi.IntPtrInput
	RevocationConfiguration           CertificateAuthorityRevocationConfigurationPtrInput
	Serial                            pulumi.StringPtrInput
	Status                            pulumi.StringPtrInput
	Tags                              pulumi.StringMapInput
	Type                              pulumi.StringPtrInput
}

func (CertificateAuthorityState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityState)(nil)).Elem()
}

type certificateAuthorityArgs struct {
	CertificateAuthorityConfiguration CertificateAuthorityCertificateAuthorityConfiguration `pulumi:"certificateAuthorityConfiguration"`
	Enabled                           *bool                                                 `pulumi:"enabled"`
	PermanentDeletionTimeInDays       *int                                                  `pulumi:"permanentDeletionTimeInDays"`
	RevocationConfiguration           *CertificateAuthorityRevocationConfiguration          `pulumi:"revocationConfiguration"`
	Tags                              map[string]string                                     `pulumi:"tags"`
	Type                              *string                                               `pulumi:"type"`
}

// The set of arguments for constructing a CertificateAuthority resource.
type CertificateAuthorityArgs struct {
	CertificateAuthorityConfiguration CertificateAuthorityCertificateAuthorityConfigurationInput
	Enabled                           pulumi.BoolPtrInput
	PermanentDeletionTimeInDays       pulumi.IntPtrInput
	RevocationConfiguration           CertificateAuthorityRevocationConfigurationPtrInput
	Tags                              pulumi.StringMapInput
	Type                              pulumi.StringPtrInput
}

func (CertificateAuthorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateAuthorityArgs)(nil)).Elem()
}
