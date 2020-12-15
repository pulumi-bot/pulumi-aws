// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package acm

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the ARN of a certificate in AWS Certificate
// Manager (ACM), you can reference
// it by domain without having to hard code the ARNs as input.
//
// ## Example Usage
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
// 		_, err := acm.LookupCertificate(ctx, &acm.LookupCertificateArgs{
// 			Domain: "tf.example.com",
// 			Statuses: []string{
// 				"ISSUED",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := true
// 		_, err = acm.LookupCertificate(ctx, &acm.LookupCertificateArgs{
// 			Domain:     "tf.example.com",
// 			MostRecent: _opt0,
// 			Types: []string{
// 				"AMAZON_ISSUED",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = acm.LookupCertificate(ctx, &acm.LookupCertificateArgs{
// 			Domain: "tf.example.com",
// 			KeyTypes: []string{
// 				"RSA_4096",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("aws:acm/getCertificate:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateArgs struct {
	// The domain of the certificate to look up. If no certificate is found with this name, an error will be returned.
	Domain string `pulumi:"domain"`
	// A list of key algorithms to filter certificates. By default, ACM does not return all certificate types when searching. Valid values are `RSA_1024`, `RSA_2048`, `RSA_4096`, `EC_prime256v1`, `EC_secp384r1`, and `EC_secp521r1`.
	KeyTypes []string `pulumi:"keyTypes"`
	// If set to true, it sorts the certificates matched by previous criteria by the NotBefore field, returning only the most recent one. If set to false, it returns an error if more than one certificate is found. Defaults to false.
	MostRecent *bool `pulumi:"mostRecent"`
	// A list of statuses on which to filter the returned list. Valid values are `PENDING_VALIDATION`, `ISSUED`,
	// `INACTIVE`, `EXPIRED`, `VALIDATION_TIMED_OUT`, `REVOKED` and `FAILED`. If no value is specified, only certificates in the `ISSUED` state
	// are returned.
	Statuses []string `pulumi:"statuses"`
	// A mapping of tags for the resource.
	Tags map[string]string `pulumi:"tags"`
	// A list of types on which to filter the returned list. Valid values are `AMAZON_ISSUED` and `IMPORTED`.
	Types []string `pulumi:"types"`
}

// A collection of values returned by getCertificate.
type LookupCertificateResult struct {
	// Amazon Resource Name (ARN) of the found certificate, suitable for referencing in other resources that support ACM certificates.
	Arn    string `pulumi:"arn"`
	Domain string `pulumi:"domain"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	KeyTypes   []string `pulumi:"keyTypes"`
	MostRecent *bool    `pulumi:"mostRecent"`
	Statuses   []string `pulumi:"statuses"`
	// A mapping of tags for the resource.
	Tags  map[string]string `pulumi:"tags"`
	Types []string          `pulumi:"types"`
}
