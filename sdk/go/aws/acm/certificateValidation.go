// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package acm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource represents a successful validation of an ACM certificate in concert
// with other resources.
//
// Most commonly, this resource is used together with `route53.Record` and
// `acm.Certificate` to request a DNS validated certificate,
// deploy the required validation records and wait for validation to complete.
//
// > **WARNING:** This resource implements a part of the validation workflow. It does not represent a real-world entity in AWS, therefore changing or deleting this resource on its own has no immediate effect.
//
// ## Example Usage
// ### Email Validation
//
// In this situation, the resource is simply a waiter for manual email approval of ACM certificates.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/acm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleCertificate, err := acm.NewCertificate(ctx, "exampleCertificate", &acm.CertificateArgs{
// 			DomainName:       pulumi.String("example.com"),
// 			ValidationMethod: pulumi.String("EMAIL"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = acm.NewCertificateValidation(ctx, "exampleCertificateValidation", &acm.CertificateValidationArgs{
// 			CertificateArn: exampleCertificate.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type CertificateValidation struct {
	pulumi.CustomResourceState

	// The ARN of the certificate that is being validated.
	CertificateArn pulumi.StringOutput `pulumi:"certificateArn"`
	// List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
	ValidationRecordFqdns pulumi.StringArrayOutput `pulumi:"validationRecordFqdns"`
}

// NewCertificateValidation registers a new resource with the given unique name, arguments, and options.
func NewCertificateValidation(ctx *pulumi.Context,
	name string, args *CertificateValidationArgs, opts ...pulumi.ResourceOption) (*CertificateValidation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateArn == nil {
		return nil, errors.New("invalid value for required argument 'CertificateArn'")
	}
	var resource CertificateValidation
	err := ctx.RegisterResource("aws:acm/certificateValidation:CertificateValidation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateValidation gets an existing CertificateValidation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateValidation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateValidationState, opts ...pulumi.ResourceOption) (*CertificateValidation, error) {
	var resource CertificateValidation
	err := ctx.ReadResource("aws:acm/certificateValidation:CertificateValidation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateValidation resources.
type certificateValidationState struct {
	// The ARN of the certificate that is being validated.
	CertificateArn *string `pulumi:"certificateArn"`
	// List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
	ValidationRecordFqdns []string `pulumi:"validationRecordFqdns"`
}

type CertificateValidationState struct {
	// The ARN of the certificate that is being validated.
	CertificateArn pulumi.StringPtrInput
	// List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
	ValidationRecordFqdns pulumi.StringArrayInput
}

func (CertificateValidationState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateValidationState)(nil)).Elem()
}

type certificateValidationArgs struct {
	// The ARN of the certificate that is being validated.
	CertificateArn string `pulumi:"certificateArn"`
	// List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
	ValidationRecordFqdns []string `pulumi:"validationRecordFqdns"`
}

// The set of arguments for constructing a CertificateValidation resource.
type CertificateValidationArgs struct {
	// The ARN of the certificate that is being validated.
	CertificateArn pulumi.StringInput
	// List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
	ValidationRecordFqdns pulumi.StringArrayInput
}

func (CertificateValidationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateValidationArgs)(nil)).Elem()
}

type CertificateValidationInput interface {
	pulumi.Input

	ToCertificateValidationOutput() CertificateValidationOutput
	ToCertificateValidationOutputWithContext(ctx context.Context) CertificateValidationOutput
}

type CertificateValidationPtrInput interface {
	pulumi.Input

	ToCertificateValidationPtrOutput() CertificateValidationPtrOutput
	ToCertificateValidationPtrOutputWithContext(ctx context.Context) CertificateValidationPtrOutput
}

func (CertificateValidation) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateValidation)(nil)).Elem()
}

func (i CertificateValidation) ToCertificateValidationOutput() CertificateValidationOutput {
	return i.ToCertificateValidationOutputWithContext(context.Background())
}

func (i CertificateValidation) ToCertificateValidationOutputWithContext(ctx context.Context) CertificateValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateValidationOutput)
}

func (i CertificateValidation) ToCertificateValidationPtrOutput() CertificateValidationPtrOutput {
	return i.ToCertificateValidationPtrOutputWithContext(context.Background())
}

func (i CertificateValidation) ToCertificateValidationPtrOutputWithContext(ctx context.Context) CertificateValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateValidationPtrOutput)
}

type CertificateValidationOutput struct {
	*pulumi.OutputState
}

func (CertificateValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateValidationOutput)(nil)).Elem()
}

func (o CertificateValidationOutput) ToCertificateValidationOutput() CertificateValidationOutput {
	return o
}

func (o CertificateValidationOutput) ToCertificateValidationOutputWithContext(ctx context.Context) CertificateValidationOutput {
	return o
}

type CertificateValidationPtrOutput struct {
	*pulumi.OutputState
}

func (CertificateValidationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateValidation)(nil)).Elem()
}

func (o CertificateValidationPtrOutput) ToCertificateValidationPtrOutput() CertificateValidationPtrOutput {
	return o
}

func (o CertificateValidationPtrOutput) ToCertificateValidationPtrOutputWithContext(ctx context.Context) CertificateValidationPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CertificateValidationOutput{})
	pulumi.RegisterOutputType(CertificateValidationPtrOutput{})
}
