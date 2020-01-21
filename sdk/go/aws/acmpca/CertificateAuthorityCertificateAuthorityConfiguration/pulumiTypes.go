// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package CertificateAuthorityCertificateAuthorityConfiguration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/acmpca/CertificateAuthorityCertificateAuthorityConfigurationSubject"
)

type CertificateAuthorityCertificateAuthorityConfiguration struct {
	// Type of the public key algorithm and size, in bits, of the key pair that your key pair creates when it issues a certificate. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
	KeyAlgorithm string `pulumi:"keyAlgorithm"`
	// Name of the algorithm your private CA uses to sign certificate requests. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
	SigningAlgorithm string `pulumi:"signingAlgorithm"`
	// Nested argument that contains X.500 distinguished name information. At least one nested attribute must be specified.
	Subject acmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject.CertificateAuthorityCertificateAuthorityConfigurationSubject `pulumi:"subject"`
}

type CertificateAuthorityCertificateAuthorityConfigurationInput interface {
	pulumi.Input

	ToCertificateAuthorityCertificateAuthorityConfigurationOutput() CertificateAuthorityCertificateAuthorityConfigurationOutput
	ToCertificateAuthorityCertificateAuthorityConfigurationOutputWithContext(context.Context) CertificateAuthorityCertificateAuthorityConfigurationOutput
}

type CertificateAuthorityCertificateAuthorityConfigurationArgs struct {
	// Type of the public key algorithm and size, in bits, of the key pair that your key pair creates when it issues a certificate. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
	KeyAlgorithm pulumi.StringInput `pulumi:"keyAlgorithm"`
	// Name of the algorithm your private CA uses to sign certificate requests. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
	SigningAlgorithm pulumi.StringInput `pulumi:"signingAlgorithm"`
	// Nested argument that contains X.500 distinguished name information. At least one nested attribute must be specified.
	Subject acmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject.CertificateAuthorityCertificateAuthorityConfigurationSubjectInput `pulumi:"subject"`
}

func (CertificateAuthorityCertificateAuthorityConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateAuthorityCertificateAuthorityConfiguration)(nil)).Elem()
}

func (i CertificateAuthorityCertificateAuthorityConfigurationArgs) ToCertificateAuthorityCertificateAuthorityConfigurationOutput() CertificateAuthorityCertificateAuthorityConfigurationOutput {
	return i.ToCertificateAuthorityCertificateAuthorityConfigurationOutputWithContext(context.Background())
}

func (i CertificateAuthorityCertificateAuthorityConfigurationArgs) ToCertificateAuthorityCertificateAuthorityConfigurationOutputWithContext(ctx context.Context) CertificateAuthorityCertificateAuthorityConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityCertificateAuthorityConfigurationOutput)
}

func (i CertificateAuthorityCertificateAuthorityConfigurationArgs) ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutput() CertificateAuthorityCertificateAuthorityConfigurationPtrOutput {
	return i.ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutputWithContext(context.Background())
}

func (i CertificateAuthorityCertificateAuthorityConfigurationArgs) ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutputWithContext(ctx context.Context) CertificateAuthorityCertificateAuthorityConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityCertificateAuthorityConfigurationOutput).ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutputWithContext(ctx)
}

type CertificateAuthorityCertificateAuthorityConfigurationPtrInput interface {
	pulumi.Input

	ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutput() CertificateAuthorityCertificateAuthorityConfigurationPtrOutput
	ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutputWithContext(context.Context) CertificateAuthorityCertificateAuthorityConfigurationPtrOutput
}

type certificateAuthorityCertificateAuthorityConfigurationPtrType CertificateAuthorityCertificateAuthorityConfigurationArgs

func CertificateAuthorityCertificateAuthorityConfigurationPtr(v *CertificateAuthorityCertificateAuthorityConfigurationArgs) CertificateAuthorityCertificateAuthorityConfigurationPtrInput {	return (*certificateAuthorityCertificateAuthorityConfigurationPtrType)(v)
}

func (*certificateAuthorityCertificateAuthorityConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateAuthorityCertificateAuthorityConfiguration)(nil)).Elem()
}

func (i *certificateAuthorityCertificateAuthorityConfigurationPtrType) ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutput() CertificateAuthorityCertificateAuthorityConfigurationPtrOutput {
	return i.ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutputWithContext(context.Background())
}

func (i *certificateAuthorityCertificateAuthorityConfigurationPtrType) ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutputWithContext(ctx context.Context) CertificateAuthorityCertificateAuthorityConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateAuthorityCertificateAuthorityConfigurationPtrOutput)
}

type CertificateAuthorityCertificateAuthorityConfigurationOutput struct { *pulumi.OutputState }

func (CertificateAuthorityCertificateAuthorityConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateAuthorityCertificateAuthorityConfiguration)(nil)).Elem()
}

func (o CertificateAuthorityCertificateAuthorityConfigurationOutput) ToCertificateAuthorityCertificateAuthorityConfigurationOutput() CertificateAuthorityCertificateAuthorityConfigurationOutput {
	return o
}

func (o CertificateAuthorityCertificateAuthorityConfigurationOutput) ToCertificateAuthorityCertificateAuthorityConfigurationOutputWithContext(ctx context.Context) CertificateAuthorityCertificateAuthorityConfigurationOutput {
	return o
}

func (o CertificateAuthorityCertificateAuthorityConfigurationOutput) ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutput() CertificateAuthorityCertificateAuthorityConfigurationPtrOutput {
	return o.ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutputWithContext(context.Background())
}

func (o CertificateAuthorityCertificateAuthorityConfigurationOutput) ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutputWithContext(ctx context.Context) CertificateAuthorityCertificateAuthorityConfigurationPtrOutput {
	return o.ApplyT(func(v CertificateAuthorityCertificateAuthorityConfiguration) *CertificateAuthorityCertificateAuthorityConfiguration {
		return &v
	}).(CertificateAuthorityCertificateAuthorityConfigurationPtrOutput)
}
// Type of the public key algorithm and size, in bits, of the key pair that your key pair creates when it issues a certificate. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
func (o CertificateAuthorityCertificateAuthorityConfigurationOutput) KeyAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func (v CertificateAuthorityCertificateAuthorityConfiguration) string { return v.KeyAlgorithm }).(pulumi.StringOutput)
}

// Name of the algorithm your private CA uses to sign certificate requests. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
func (o CertificateAuthorityCertificateAuthorityConfigurationOutput) SigningAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func (v CertificateAuthorityCertificateAuthorityConfiguration) string { return v.SigningAlgorithm }).(pulumi.StringOutput)
}

// Nested argument that contains X.500 distinguished name information. At least one nested attribute must be specified.
func (o CertificateAuthorityCertificateAuthorityConfigurationOutput) Subject() acmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject.CertificateAuthorityCertificateAuthorityConfigurationSubjectOutput {
	return o.ApplyT(func (v CertificateAuthorityCertificateAuthorityConfiguration) acmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject.CertificateAuthorityCertificateAuthorityConfigurationSubject { return v.Subject }).(acmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject.CertificateAuthorityCertificateAuthorityConfigurationSubjectOutput)
}

type CertificateAuthorityCertificateAuthorityConfigurationPtrOutput struct { *pulumi.OutputState}

func (CertificateAuthorityCertificateAuthorityConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateAuthorityCertificateAuthorityConfiguration)(nil)).Elem()
}

func (o CertificateAuthorityCertificateAuthorityConfigurationPtrOutput) ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutput() CertificateAuthorityCertificateAuthorityConfigurationPtrOutput {
	return o
}

func (o CertificateAuthorityCertificateAuthorityConfigurationPtrOutput) ToCertificateAuthorityCertificateAuthorityConfigurationPtrOutputWithContext(ctx context.Context) CertificateAuthorityCertificateAuthorityConfigurationPtrOutput {
	return o
}

func (o CertificateAuthorityCertificateAuthorityConfigurationPtrOutput) Elem() CertificateAuthorityCertificateAuthorityConfigurationOutput {
	return o.ApplyT(func (v *CertificateAuthorityCertificateAuthorityConfiguration) CertificateAuthorityCertificateAuthorityConfiguration { return *v }).(CertificateAuthorityCertificateAuthorityConfigurationOutput)
}

// Type of the public key algorithm and size, in bits, of the key pair that your key pair creates when it issues a certificate. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
func (o CertificateAuthorityCertificateAuthorityConfigurationPtrOutput) KeyAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func (v CertificateAuthorityCertificateAuthorityConfiguration) string { return v.KeyAlgorithm }).(pulumi.StringOutput)
}

// Name of the algorithm your private CA uses to sign certificate requests. Valid values can be found in the [ACM PCA Documentation](https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CertificateAuthorityConfiguration.html).
func (o CertificateAuthorityCertificateAuthorityConfigurationPtrOutput) SigningAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func (v CertificateAuthorityCertificateAuthorityConfiguration) string { return v.SigningAlgorithm }).(pulumi.StringOutput)
}

// Nested argument that contains X.500 distinguished name information. At least one nested attribute must be specified.
func (o CertificateAuthorityCertificateAuthorityConfigurationPtrOutput) Subject() acmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject.CertificateAuthorityCertificateAuthorityConfigurationSubjectOutput {
	return o.ApplyT(func (v CertificateAuthorityCertificateAuthorityConfiguration) acmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject.CertificateAuthorityCertificateAuthorityConfigurationSubject { return v.Subject }).(acmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject.CertificateAuthorityCertificateAuthorityConfigurationSubjectOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateAuthorityCertificateAuthorityConfigurationOutput{})
	pulumi.RegisterOutputType(CertificateAuthorityCertificateAuthorityConfigurationPtrOutput{})
}
