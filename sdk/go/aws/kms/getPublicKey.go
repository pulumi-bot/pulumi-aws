// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the public key about the specified KMS Key with flexible key id input. This can be useful to reference key alias without having to hard code the ARN as input.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/kms"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.GetPublicKey(ctx, &kms.GetPublicKeyArgs{
// 			KeyId: "alias/my-key",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.GetPublicKey(ctx, &kms.GetPublicKeyArgs{
// 			KeyId: "1234abcd-12ab-34cd-56ef-1234567890ab",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.GetPublicKey(ctx, &kms.GetPublicKeyArgs{
// 			KeyId: "arn:aws:kms:us-east-1:111122223333:alias/my-key",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.GetPublicKey(ctx, &kms.GetPublicKeyArgs{
// 			KeyId: "arn:aws:kms:us-east-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetPublicKey(ctx *pulumi.Context, args *GetPublicKeyArgs, opts ...pulumi.InvokeOption) (*GetPublicKeyResult, error) {
	var rv GetPublicKeyResult
	err := ctx.Invoke("aws:kms/getPublicKey:getPublicKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicKey.
type GetPublicKeyArgs struct {
	// List of grant tokens
	GrantTokens []string `pulumi:"grantTokens"`
	// Key identifier which can be one of the following format:
	// * Key ID. E.g - `1234abcd-12ab-34cd-56ef-1234567890ab`
	// * Key ARN. E.g. - `arn:aws:kms:us-east-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
	// * Alias name. E.g. - `alias/my-key`
	// * Alias ARN - E.g. - `arn:aws:kms:us-east-1:111122223333:alias/my-key`
	KeyId string `pulumi:"keyId"`
}

// A collection of values returned by getPublicKey.
type GetPublicKeyResult struct {
	// Key ARN of the asymmetric CMK from which the public key was downloaded.
	Arn string `pulumi:"arn"`
	// Type of the public key that was downloaded.
	CustomerMasterKeySpec string `pulumi:"customerMasterKeySpec"`
	// Encryption algorithms that AWS KMS supports for this key. Only set when the `keyUsage` of the public key is `ENCRYPT_DECRYPT`.
	EncryptionAlgorithms []string `pulumi:"encryptionAlgorithms"`
	GrantTokens          []string `pulumi:"grantTokens"`
	// The provider-assigned unique ID for this managed resource.
	Id    string `pulumi:"id"`
	KeyId string `pulumi:"keyId"`
	// Permitted use of the public key. Valid values are `ENCRYPT_DECRYPT` or `SIGN_VERIFY`
	KeyUsage string `pulumi:"keyUsage"`
	// Exported public key. The value is a DER-encoded X.509 public key, also known as SubjectPublicKeyInfo (SPKI), as defined in [RFC 5280](https://tools.ietf.org/html/rfc5280). The value is Base64-encoded.
	PublicKey string `pulumi:"publicKey"`
	// Signing algorithms that AWS KMS supports for this key. Only set when the `keyUsage` of the public key is `SIGN_VERIFY`.
	SigningAlgorithms []string `pulumi:"signingAlgorithms"`
}

func GetPublicKeyApply(ctx *pulumi.Context, args GetPublicKeyApplyInput, opts ...pulumi.InvokeOption) GetPublicKeyResultOutput {
	return args.ToGetPublicKeyApplyOutput().ApplyT(func(v GetPublicKeyArgs) (GetPublicKeyResult, error) {
		r, err := GetPublicKey(ctx, &v, opts...)
		return *r, err

	}).(GetPublicKeyResultOutput)
}

// GetPublicKeyApplyInput is an input type that accepts GetPublicKeyApplyArgs and GetPublicKeyApplyOutput values.
// You can construct a concrete instance of `GetPublicKeyApplyInput` via:
//
//          GetPublicKeyApplyArgs{...}
type GetPublicKeyApplyInput interface {
	pulumi.Input

	ToGetPublicKeyApplyOutput() GetPublicKeyApplyOutput
	ToGetPublicKeyApplyOutputWithContext(context.Context) GetPublicKeyApplyOutput
}

// A collection of arguments for invoking getPublicKey.
type GetPublicKeyApplyArgs struct {
	// List of grant tokens
	GrantTokens pulumi.StringArrayInput `pulumi:"grantTokens"`
	// Key identifier which can be one of the following format:
	// * Key ID. E.g - `1234abcd-12ab-34cd-56ef-1234567890ab`
	// * Key ARN. E.g. - `arn:aws:kms:us-east-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
	// * Alias name. E.g. - `alias/my-key`
	// * Alias ARN - E.g. - `arn:aws:kms:us-east-1:111122223333:alias/my-key`
	KeyId pulumi.StringInput `pulumi:"keyId"`
}

func (GetPublicKeyApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicKeyArgs)(nil)).Elem()
}

func (i GetPublicKeyApplyArgs) ToGetPublicKeyApplyOutput() GetPublicKeyApplyOutput {
	return i.ToGetPublicKeyApplyOutputWithContext(context.Background())
}

func (i GetPublicKeyApplyArgs) ToGetPublicKeyApplyOutputWithContext(ctx context.Context) GetPublicKeyApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetPublicKeyApplyOutput)
}

// A collection of arguments for invoking getPublicKey.
type GetPublicKeyApplyOutput struct{ *pulumi.OutputState }

func (GetPublicKeyApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicKeyArgs)(nil)).Elem()
}

func (o GetPublicKeyApplyOutput) ToGetPublicKeyApplyOutput() GetPublicKeyApplyOutput {
	return o
}

func (o GetPublicKeyApplyOutput) ToGetPublicKeyApplyOutputWithContext(ctx context.Context) GetPublicKeyApplyOutput {
	return o
}

// List of grant tokens
func (o GetPublicKeyApplyOutput) GrantTokens() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPublicKeyArgs) []string { return v.GrantTokens }).(pulumi.StringArrayOutput)
}

// Key identifier which can be one of the following format:
// * Key ID. E.g - `1234abcd-12ab-34cd-56ef-1234567890ab`
// * Key ARN. E.g. - `arn:aws:kms:us-east-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
// * Alias name. E.g. - `alias/my-key`
// * Alias ARN - E.g. - `arn:aws:kms:us-east-1:111122223333:alias/my-key`
func (o GetPublicKeyApplyOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicKeyArgs) string { return v.KeyId }).(pulumi.StringOutput)
}

// A collection of values returned by getPublicKey.
type GetPublicKeyResultOutput struct{ *pulumi.OutputState }

func (GetPublicKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicKeyResult)(nil)).Elem()
}

func (o GetPublicKeyResultOutput) ToGetPublicKeyResultOutput() GetPublicKeyResultOutput {
	return o
}

func (o GetPublicKeyResultOutput) ToGetPublicKeyResultOutputWithContext(ctx context.Context) GetPublicKeyResultOutput {
	return o
}

// Key ARN of the asymmetric CMK from which the public key was downloaded.
func (o GetPublicKeyResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicKeyResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Type of the public key that was downloaded.
func (o GetPublicKeyResultOutput) CustomerMasterKeySpec() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicKeyResult) string { return v.CustomerMasterKeySpec }).(pulumi.StringOutput)
}

// Encryption algorithms that AWS KMS supports for this key. Only set when the `keyUsage` of the public key is `ENCRYPT_DECRYPT`.
func (o GetPublicKeyResultOutput) EncryptionAlgorithms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPublicKeyResult) []string { return v.EncryptionAlgorithms }).(pulumi.StringArrayOutput)
}

func (o GetPublicKeyResultOutput) GrantTokens() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPublicKeyResult) []string { return v.GrantTokens }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPublicKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPublicKeyResultOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicKeyResult) string { return v.KeyId }).(pulumi.StringOutput)
}

// Permitted use of the public key. Valid values are `ENCRYPT_DECRYPT` or `SIGN_VERIFY`
func (o GetPublicKeyResultOutput) KeyUsage() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicKeyResult) string { return v.KeyUsage }).(pulumi.StringOutput)
}

// Exported public key. The value is a DER-encoded X.509 public key, also known as SubjectPublicKeyInfo (SPKI), as defined in [RFC 5280](https://tools.ietf.org/html/rfc5280). The value is Base64-encoded.
func (o GetPublicKeyResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicKeyResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

// Signing algorithms that AWS KMS supports for this key. Only set when the `keyUsage` of the public key is `SIGN_VERIFY`.
func (o GetPublicKeyResultOutput) SigningAlgorithms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPublicKeyResult) []string { return v.SigningAlgorithms }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPublicKeyApplyOutput{})
	pulumi.RegisterOutputType(GetPublicKeyResultOutput{})
}
