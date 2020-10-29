// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an [EC2 key pair](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) resource. A key pair is used to control login access to EC2 instances.
//
// Currently this resource requires an existing user-supplied key pair. This key pair's public key will be registered with AWS to allow logging-in to EC2 instances.
//
// When importing an existing key pair the public key material may be in any format supported by AWS. Supported formats (per the [AWS documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html#how-to-generate-your-own-key-and-import-it-to-aws)) are:
//
// * OpenSSH public key format (the format in ~/.ssh/authorized_keys)
// * Base64 encoded DER format
// * SSH public key file format as specified in RFC4716
type KeyPair struct {
	pulumi.CustomResourceState

	// The key pair ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The name for the key pair.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `keyName`.
	KeyNamePrefix pulumi.StringPtrOutput `pulumi:"keyNamePrefix"`
	// The key pair ID.
	KeyPairId pulumi.StringOutput `pulumi:"keyPairId"`
	// The public key material.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewKeyPair registers a new resource with the given unique name, arguments, and options.
func NewKeyPair(ctx *pulumi.Context,
	name string, args *KeyPairArgs, opts ...pulumi.ResourceOption) (*KeyPair, error) {
	if args == nil || args.PublicKey == nil {
		return nil, errors.New("missing required argument 'PublicKey'")
	}
	if args == nil {
		args = &KeyPairArgs{}
	}
	var resource KeyPair
	err := ctx.RegisterResource("aws:ec2/keyPair:KeyPair", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:ec2/keyPair:KeyPair", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyPair resources.
type keyPairState struct {
	// The key pair ARN.
	Arn *string `pulumi:"arn"`
	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	Fingerprint *string `pulumi:"fingerprint"`
	// The name for the key pair.
	KeyName *string `pulumi:"keyName"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `keyName`.
	KeyNamePrefix *string `pulumi:"keyNamePrefix"`
	// The key pair ID.
	KeyPairId *string `pulumi:"keyPairId"`
	// The public key material.
	PublicKey *string `pulumi:"publicKey"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type KeyPairState struct {
	// The key pair ARN.
	Arn pulumi.StringPtrInput
	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	Fingerprint pulumi.StringPtrInput
	// The name for the key pair.
	KeyName pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `keyName`.
	KeyNamePrefix pulumi.StringPtrInput
	// The key pair ID.
	KeyPairId pulumi.StringPtrInput
	// The public key material.
	PublicKey pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (KeyPairState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairState)(nil)).Elem()
}

type keyPairArgs struct {
	// The name for the key pair.
	KeyName *string `pulumi:"keyName"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `keyName`.
	KeyNamePrefix *string `pulumi:"keyNamePrefix"`
	// The public key material.
	PublicKey string `pulumi:"publicKey"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a KeyPair resource.
type KeyPairArgs struct {
	// The name for the key pair.
	KeyName pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `keyName`.
	KeyNamePrefix pulumi.StringPtrInput
	// The public key material.
	PublicKey pulumi.StringInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (KeyPairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairArgs)(nil)).Elem()
}
