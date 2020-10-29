// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a KMS customer master key.
type Key struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the key.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
	// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
	CustomerMasterKeySpec pulumi.StringPtrOutput `pulumi:"customerMasterKeySpec"`
	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	DeletionWindowInDays pulumi.IntPtrOutput `pulumi:"deletionWindowInDays"`
	// The description of the key as viewed in AWS console.
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
	// is enabled. Defaults to false.
	EnableKeyRotation pulumi.BoolPtrOutput `pulumi:"enableKeyRotation"`
	// Specifies whether the key is enabled. Defaults to true.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// The globally unique identifier for the key.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT` or `SIGN_VERIFY`.
	// Defaults to `ENCRYPT_DECRYPT`.
	KeyUsage pulumi.StringPtrOutput `pulumi:"keyUsage"`
	// A valid policy JSON document.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// A map of tags to assign to the object.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		args = &KeyArgs{}
	}
	var resource Key
	err := ctx.RegisterResource("aws:kms/key:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("aws:kms/key:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
	// The Amazon Resource Name (ARN) of the key.
	Arn *string `pulumi:"arn"`
	// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
	// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
	CustomerMasterKeySpec *string `pulumi:"customerMasterKeySpec"`
	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// The description of the key as viewed in AWS console.
	Description *string `pulumi:"description"`
	// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
	// is enabled. Defaults to false.
	EnableKeyRotation *bool `pulumi:"enableKeyRotation"`
	// Specifies whether the key is enabled. Defaults to true.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The globally unique identifier for the key.
	KeyId *string `pulumi:"keyId"`
	// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT` or `SIGN_VERIFY`.
	// Defaults to `ENCRYPT_DECRYPT`.
	KeyUsage *string `pulumi:"keyUsage"`
	// A valid policy JSON document.
	Policy *string `pulumi:"policy"`
	// A map of tags to assign to the object.
	Tags map[string]string `pulumi:"tags"`
}

type KeyState struct {
	// The Amazon Resource Name (ARN) of the key.
	Arn pulumi.StringPtrInput
	// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
	// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
	CustomerMasterKeySpec pulumi.StringPtrInput
	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	DeletionWindowInDays pulumi.IntPtrInput
	// The description of the key as viewed in AWS console.
	Description pulumi.StringPtrInput
	// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
	// is enabled. Defaults to false.
	EnableKeyRotation pulumi.BoolPtrInput
	// Specifies whether the key is enabled. Defaults to true.
	IsEnabled pulumi.BoolPtrInput
	// The globally unique identifier for the key.
	KeyId pulumi.StringPtrInput
	// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT` or `SIGN_VERIFY`.
	// Defaults to `ENCRYPT_DECRYPT`.
	KeyUsage pulumi.StringPtrInput
	// A valid policy JSON document.
	Policy pulumi.StringPtrInput
	// A map of tags to assign to the object.
	Tags pulumi.StringMapInput
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
	// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
	CustomerMasterKeySpec *string `pulumi:"customerMasterKeySpec"`
	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	DeletionWindowInDays *int `pulumi:"deletionWindowInDays"`
	// The description of the key as viewed in AWS console.
	Description *string `pulumi:"description"`
	// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
	// is enabled. Defaults to false.
	EnableKeyRotation *bool `pulumi:"enableKeyRotation"`
	// Specifies whether the key is enabled. Defaults to true.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT` or `SIGN_VERIFY`.
	// Defaults to `ENCRYPT_DECRYPT`.
	KeyUsage *string `pulumi:"keyUsage"`
	// A valid policy JSON document.
	Policy *string `pulumi:"policy"`
	// A map of tags to assign to the object.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports.
	// Valid values: `SYMMETRIC_DEFAULT`,  `RSA_2048`, `RSA_3072`, `RSA_4096`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. Defaults to `SYMMETRIC_DEFAULT`. For help with choosing a key spec, see the [AWS KMS Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html).
	CustomerMasterKeySpec pulumi.StringPtrInput
	// Duration in days after which the key is deleted
	// after destruction of the resource, must be between 7 and 30 days. Defaults to 30 days.
	DeletionWindowInDays pulumi.IntPtrInput
	// The description of the key as viewed in AWS console.
	Description pulumi.StringPtrInput
	// Specifies whether [key rotation](http://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html)
	// is enabled. Defaults to false.
	EnableKeyRotation pulumi.BoolPtrInput
	// Specifies whether the key is enabled. Defaults to true.
	IsEnabled pulumi.BoolPtrInput
	// Specifies the intended use of the key. Valid values: `ENCRYPT_DECRYPT` or `SIGN_VERIFY`.
	// Defaults to `ENCRYPT_DECRYPT`.
	KeyUsage pulumi.StringPtrInput
	// A valid policy JSON document.
	Policy pulumi.StringPtrInput
	// A map of tags to assign to the object.
	Tags pulumi.StringMapInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}
