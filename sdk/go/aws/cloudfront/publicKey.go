// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type PublicKey struct {
	pulumi.CustomResourceState

	// Internal value used by CloudFront to allow future updates to the public key configuration.
	CallerReference pulumi.StringOutput `pulumi:"callerReference"`
	// An optional comment about the public key.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	EncodedKey pulumi.StringOutput `pulumi:"encodedKey"`
	// The current version of the public key. For example: `E2QWRUHAPOMQZL`.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name for the public key. By default generated by this provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name for the public key. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
}

// NewPublicKey registers a new resource with the given unique name, arguments, and options.
func NewPublicKey(ctx *pulumi.Context,
	name string, args *PublicKeyArgs, opts ...pulumi.ResourceOption) (*PublicKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EncodedKey == nil {
		return nil, errors.New("invalid value for required argument 'EncodedKey'")
	}
	var resource PublicKey
	err := ctx.RegisterResource("aws:cloudfront/publicKey:PublicKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicKey gets an existing PublicKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicKeyState, opts ...pulumi.ResourceOption) (*PublicKey, error) {
	var resource PublicKey
	err := ctx.ReadResource("aws:cloudfront/publicKey:PublicKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicKey resources.
type publicKeyState struct {
	// Internal value used by CloudFront to allow future updates to the public key configuration.
	CallerReference *string `pulumi:"callerReference"`
	// An optional comment about the public key.
	Comment *string `pulumi:"comment"`
	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	EncodedKey *string `pulumi:"encodedKey"`
	// The current version of the public key. For example: `E2QWRUHAPOMQZL`.
	Etag *string `pulumi:"etag"`
	// The name for the public key. By default generated by this provider.
	Name *string `pulumi:"name"`
	// The name for the public key. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
}

type PublicKeyState struct {
	// Internal value used by CloudFront to allow future updates to the public key configuration.
	CallerReference pulumi.StringPtrInput
	// An optional comment about the public key.
	Comment pulumi.StringPtrInput
	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	EncodedKey pulumi.StringPtrInput
	// The current version of the public key. For example: `E2QWRUHAPOMQZL`.
	Etag pulumi.StringPtrInput
	// The name for the public key. By default generated by this provider.
	Name pulumi.StringPtrInput
	// The name for the public key. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
}

func (PublicKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicKeyState)(nil)).Elem()
}

type publicKeyArgs struct {
	// An optional comment about the public key.
	Comment *string `pulumi:"comment"`
	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	EncodedKey string `pulumi:"encodedKey"`
	// The name for the public key. By default generated by this provider.
	Name *string `pulumi:"name"`
	// The name for the public key. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
}

// The set of arguments for constructing a PublicKey resource.
type PublicKeyArgs struct {
	// An optional comment about the public key.
	Comment pulumi.StringPtrInput
	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	EncodedKey pulumi.StringInput
	// The name for the public key. By default generated by this provider.
	Name pulumi.StringPtrInput
	// The name for the public key. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
}

func (PublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicKeyArgs)(nil)).Elem()
}

type PublicKeyInput interface {
	pulumi.Input

	ToPublicKeyOutput() PublicKeyOutput
	ToPublicKeyOutputWithContext(ctx context.Context) PublicKeyOutput
}

func (*PublicKey) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicKey)(nil))
}

func (i *PublicKey) ToPublicKeyOutput() PublicKeyOutput {
	return i.ToPublicKeyOutputWithContext(context.Background())
}

func (i *PublicKey) ToPublicKeyOutputWithContext(ctx context.Context) PublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicKeyOutput)
}

type PublicKeyOutput struct {
	*pulumi.OutputState
}

func (PublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicKey)(nil))
}

func (o PublicKeyOutput) ToPublicKeyOutput() PublicKeyOutput {
	return o
}

func (o PublicKeyOutput) ToPublicKeyOutputWithContext(ctx context.Context) PublicKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PublicKeyOutput{})
}
