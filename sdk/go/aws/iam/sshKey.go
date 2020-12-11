// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Uploads an SSH public key and associates it with the specified IAM user.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		userUser, err := iam.NewUser(ctx, "userUser", &iam.UserArgs{
// 			Path: pulumi.String("/"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewSshKey(ctx, "userSshKey", &iam.SshKeyArgs{
// 			Username:  userUser.Name,
// 			Encoding:  pulumi.String("SSH"),
// 			PublicKey: pulumi.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD3F6tyPEFEzV0LX3X8BsXdMsQz1x2cEikKDEY0aIj41qgxMCP/iteneqXSIFZBp5vizPvaoIR3Um9xK7PGoW8giupGn+EPuxIA4cDM4vzOqOkiMPhz5XK0whEjkVzTo4+S0puvDZuwIsdiW9mxhJc7tgBNL0cYlWSYVkz4G/fslNfRPW5mYAM49f4fhtxPb5ok4Q2Lg9dPKVHO/Bgeu5woMc7RY0p1ej6D4CKFE6lymSDJpW0YHX/wqE9+cfEauh7xZcG0q9t2ta6F6fmX0agvpFyZo8aFbXeUBr7osSCJNgvavWbM/06niWrOvYX2xwWdhXmXSrbX8ZbabVohBK41 mytest@mydomain.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// SSH public keys can be imported using the `username`, `ssh_public_key_id`, and `encoding` e.g.
//
// ```sh
//  $ pulumi import aws:iam/sshKey:SshKey user user:APKAJNCNNJICVN7CFKCA:SSH
// ```
type SshKey struct {
	pulumi.CustomResourceState

	// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use `SSH`. To retrieve the public key in PEM format, use `PEM`.
	Encoding pulumi.StringOutput `pulumi:"encoding"`
	// The MD5 message digest of the SSH public key.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// The unique identifier for the SSH public key.
	SshPublicKeyId pulumi.StringOutput `pulumi:"sshPublicKeyId"`
	// The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is `active`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The name of the IAM user to associate the SSH public key with.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewSshKey registers a new resource with the given unique name, arguments, and options.
func NewSshKey(ctx *pulumi.Context,
	name string, args *SshKeyArgs, opts ...pulumi.ResourceOption) (*SshKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Encoding == nil {
		return nil, errors.New("invalid value for required argument 'Encoding'")
	}
	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource SshKey
	err := ctx.RegisterResource("aws:iam/sshKey:SshKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSshKey gets an existing SshKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSshKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SshKeyState, opts ...pulumi.ResourceOption) (*SshKey, error) {
	var resource SshKey
	err := ctx.ReadResource("aws:iam/sshKey:SshKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SshKey resources.
type sshKeyState struct {
	// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use `SSH`. To retrieve the public key in PEM format, use `PEM`.
	Encoding *string `pulumi:"encoding"`
	// The MD5 message digest of the SSH public key.
	Fingerprint *string `pulumi:"fingerprint"`
	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
	PublicKey *string `pulumi:"publicKey"`
	// The unique identifier for the SSH public key.
	SshPublicKeyId *string `pulumi:"sshPublicKeyId"`
	// The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is `active`.
	Status *string `pulumi:"status"`
	// The name of the IAM user to associate the SSH public key with.
	Username *string `pulumi:"username"`
}

type SshKeyState struct {
	// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use `SSH`. To retrieve the public key in PEM format, use `PEM`.
	Encoding pulumi.StringPtrInput
	// The MD5 message digest of the SSH public key.
	Fingerprint pulumi.StringPtrInput
	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
	PublicKey pulumi.StringPtrInput
	// The unique identifier for the SSH public key.
	SshPublicKeyId pulumi.StringPtrInput
	// The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is `active`.
	Status pulumi.StringPtrInput
	// The name of the IAM user to associate the SSH public key with.
	Username pulumi.StringPtrInput
}

func (SshKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshKeyState)(nil)).Elem()
}

type sshKeyArgs struct {
	// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use `SSH`. To retrieve the public key in PEM format, use `PEM`.
	Encoding string `pulumi:"encoding"`
	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
	PublicKey string `pulumi:"publicKey"`
	// The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is `active`.
	Status *string `pulumi:"status"`
	// The name of the IAM user to associate the SSH public key with.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a SshKey resource.
type SshKeyArgs struct {
	// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use `SSH`. To retrieve the public key in PEM format, use `PEM`.
	Encoding pulumi.StringInput
	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
	PublicKey pulumi.StringInput
	// The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is `active`.
	Status pulumi.StringPtrInput
	// The name of the IAM user to associate the SSH public key with.
	Username pulumi.StringInput
}

func (SshKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshKeyArgs)(nil)).Elem()
}

type SshKeyInput interface {
	pulumi.Input

	ToSshKeyOutput() SshKeyOutput
	ToSshKeyOutputWithContext(ctx context.Context) SshKeyOutput
}

type SshKeyPtrInput interface {
	pulumi.Input

	ToSshKeyPtrOutput() SshKeyPtrOutput
	ToSshKeyPtrOutputWithContext(ctx context.Context) SshKeyPtrOutput
}

func (SshKey) ElementType() reflect.Type {
	return reflect.TypeOf((*SshKey)(nil)).Elem()
}

func (i SshKey) ToSshKeyOutput() SshKeyOutput {
	return i.ToSshKeyOutputWithContext(context.Background())
}

func (i SshKey) ToSshKeyOutputWithContext(ctx context.Context) SshKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshKeyOutput)
}

func (i SshKey) ToSshKeyPtrOutput() SshKeyPtrOutput {
	return i.ToSshKeyPtrOutputWithContext(context.Background())
}

func (i SshKey) ToSshKeyPtrOutputWithContext(ctx context.Context) SshKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshKeyPtrOutput)
}

type SshKeyOutput struct {
	*pulumi.OutputState
}

func (SshKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshKeyOutput)(nil)).Elem()
}

func (o SshKeyOutput) ToSshKeyOutput() SshKeyOutput {
	return o
}

func (o SshKeyOutput) ToSshKeyOutputWithContext(ctx context.Context) SshKeyOutput {
	return o
}

type SshKeyPtrOutput struct {
	*pulumi.OutputState
}

func (SshKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshKey)(nil)).Elem()
}

func (o SshKeyPtrOutput) ToSshKeyPtrOutput() SshKeyPtrOutput {
	return o
}

func (o SshKeyPtrOutput) ToSshKeyPtrOutputWithContext(ctx context.Context) SshKeyPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SshKeyOutput{})
	pulumi.RegisterOutputType(SshKeyPtrOutput{})
}
