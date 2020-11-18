// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Lightsail Key Pair, for use with Lightsail Instances. These key pairs
// are separate from EC2 Key Pairs, and must be created or imported for use with
// Lightsail.
//
// > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
//
// ## Example Usage
// ### Creating A New Key Pair
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lightsail"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := lightsail.NewKeyPair(ctx, "lgKeyPair", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Create new Key Pair, encrypting the private key with a PGP Key
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lightsail"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := lightsail.NewKeyPair(ctx, "lgKeyPair", &lightsail.KeyPairArgs{
// 			PgpKey: pulumi.String("keybase:keybaseusername"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type KeyPair struct {
	pulumi.CustomResourceState

	// The ARN of the Lightsail key pair
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The MD5 public key fingerprint for the encrypted
	// private key
	EncryptedFingerprint pulumi.StringOutput `pulumi:"encryptedFingerprint"`
	// the private key material, base 64 encoded and
	// encrypted with the given `pgpKey`. This is only populated when creating a new
	// key and `pgpKey` is supplied
	EncryptedPrivateKey pulumi.StringOutput `pulumi:"encryptedPrivateKey"`
	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The name of the Lightsail Key Pair. If omitted, a unique
	// name will be generated by this provider
	Name       pulumi.StringOutput    `pulumi:"name"`
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// An optional PGP key to encrypt the resulting private
	// key material. Only used when creating a new key pair
	PgpKey pulumi.StringPtrOutput `pulumi:"pgpKey"`
	// the private key, base64 encoded. This is only populated
	// when creating a new key, and when no `pgpKey` is provided
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The public key material. This public key will be
	// imported into Lightsail
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
}

// NewKeyPair registers a new resource with the given unique name, arguments, and options.
func NewKeyPair(ctx *pulumi.Context,
	name string, args *KeyPairArgs, opts ...pulumi.ResourceOption) (*KeyPair, error) {
	if args == nil {
		args = &KeyPairArgs{}
	}

	var resource KeyPair
	err := ctx.RegisterResource("aws:lightsail/keyPair:KeyPair", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyPair gets an existing KeyPair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyPair(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyPairState, opts ...pulumi.ResourceOption) (*KeyPair, error) {
	var resource KeyPair
	err := ctx.ReadResource("aws:lightsail/keyPair:KeyPair", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyPair resources.
type keyPairState struct {
	// The ARN of the Lightsail key pair
	Arn *string `pulumi:"arn"`
	// The MD5 public key fingerprint for the encrypted
	// private key
	EncryptedFingerprint *string `pulumi:"encryptedFingerprint"`
	// the private key material, base 64 encoded and
	// encrypted with the given `pgpKey`. This is only populated when creating a new
	// key and `pgpKey` is supplied
	EncryptedPrivateKey *string `pulumi:"encryptedPrivateKey"`
	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	Fingerprint *string `pulumi:"fingerprint"`
	// The name of the Lightsail Key Pair. If omitted, a unique
	// name will be generated by this provider
	Name       *string `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	// An optional PGP key to encrypt the resulting private
	// key material. Only used when creating a new key pair
	PgpKey *string `pulumi:"pgpKey"`
	// the private key, base64 encoded. This is only populated
	// when creating a new key, and when no `pgpKey` is provided
	PrivateKey *string `pulumi:"privateKey"`
	// The public key material. This public key will be
	// imported into Lightsail
	PublicKey *string `pulumi:"publicKey"`
}

type KeyPairState struct {
	// The ARN of the Lightsail key pair
	Arn pulumi.StringPtrInput
	// The MD5 public key fingerprint for the encrypted
	// private key
	EncryptedFingerprint pulumi.StringPtrInput
	// the private key material, base 64 encoded and
	// encrypted with the given `pgpKey`. This is only populated when creating a new
	// key and `pgpKey` is supplied
	EncryptedPrivateKey pulumi.StringPtrInput
	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	Fingerprint pulumi.StringPtrInput
	// The name of the Lightsail Key Pair. If omitted, a unique
	// name will be generated by this provider
	Name       pulumi.StringPtrInput
	NamePrefix pulumi.StringPtrInput
	// An optional PGP key to encrypt the resulting private
	// key material. Only used when creating a new key pair
	PgpKey pulumi.StringPtrInput
	// the private key, base64 encoded. This is only populated
	// when creating a new key, and when no `pgpKey` is provided
	PrivateKey pulumi.StringPtrInput
	// The public key material. This public key will be
	// imported into Lightsail
	PublicKey pulumi.StringPtrInput
}

func (KeyPairState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairState)(nil)).Elem()
}

type keyPairArgs struct {
	// The name of the Lightsail Key Pair. If omitted, a unique
	// name will be generated by this provider
	Name       *string `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	// An optional PGP key to encrypt the resulting private
	// key material. Only used when creating a new key pair
	PgpKey *string `pulumi:"pgpKey"`
	// The public key material. This public key will be
	// imported into Lightsail
	PublicKey *string `pulumi:"publicKey"`
}

// The set of arguments for constructing a KeyPair resource.
type KeyPairArgs struct {
	// The name of the Lightsail Key Pair. If omitted, a unique
	// name will be generated by this provider
	Name       pulumi.StringPtrInput
	NamePrefix pulumi.StringPtrInput
	// An optional PGP key to encrypt the resulting private
	// key material. Only used when creating a new key pair
	PgpKey pulumi.StringPtrInput
	// The public key material. This public key will be
	// imported into Lightsail
	PublicKey pulumi.StringPtrInput
}

func (KeyPairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairArgs)(nil)).Elem()
}

type KeyPairInput interface {
	pulumi.Input

	ToKeyPairOutput() KeyPairOutput
	ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput
}

func (KeyPair) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPair)(nil)).Elem()
}

func (i KeyPair) ToKeyPairOutput() KeyPairOutput {
	return i.ToKeyPairOutputWithContext(context.Background())
}

func (i KeyPair) ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairOutput)
}

type KeyPairOutput struct {
	*pulumi.OutputState
}

func (KeyPairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPairOutput)(nil)).Elem()
}

func (o KeyPairOutput) ToKeyPairOutput() KeyPairOutput {
	return o
}

func (o KeyPairOutput) ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KeyPairOutput{})
}
