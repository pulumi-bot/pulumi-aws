// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage the default customer master key (CMK) that your AWS account uses to encrypt EBS volumes.
//
// Your AWS account has an AWS-managed default CMK that is used for encrypting an EBS volume when no CMK is specified in the API call that creates the volume.
// By using the `ebs.DefaultKmsKey` resource, you can specify a customer-managed CMK to use in place of the AWS-managed default CMK.
//
// > **NOTE:** Creating an `ebs.DefaultKmsKey` resource does not enable default EBS encryption. Use the `ebs.EncryptionByDefault` to enable default EBS encryption.
//
// > **NOTE:** Destroying this resource will reset the default CMK to the account's AWS-managed default CMK for EBS.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ebs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ebs.NewDefaultKmsKey(ctx, "example", &ebs.DefaultKmsKeyArgs{
// 			KeyArn: pulumi.Any(aws_kms_key.Example.Arn),
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
// The EBS default KMS CMK can be imported with the KMS key ARN, e.g. console
//
// ```sh
//  $ pulumi import aws:ebs/defaultKmsKey:DefaultKmsKey example arn:aws:kms:us-east-1:123456789012:key/abcd-1234
// ```
type DefaultKmsKey struct {
	pulumi.CustomResourceState

	// The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
	KeyArn pulumi.StringOutput `pulumi:"keyArn"`
}

// NewDefaultKmsKey registers a new resource with the given unique name, arguments, and options.
func NewDefaultKmsKey(ctx *pulumi.Context,
	name string, args *DefaultKmsKeyArgs, opts ...pulumi.ResourceOption) (*DefaultKmsKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyArn == nil {
		return nil, errors.New("invalid value for required argument 'KeyArn'")
	}
	var resource DefaultKmsKey
	err := ctx.RegisterResource("aws:ebs/defaultKmsKey:DefaultKmsKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultKmsKey gets an existing DefaultKmsKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultKmsKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultKmsKeyState, opts ...pulumi.ResourceOption) (*DefaultKmsKey, error) {
	var resource DefaultKmsKey
	err := ctx.ReadResource("aws:ebs/defaultKmsKey:DefaultKmsKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultKmsKey resources.
type defaultKmsKeyState struct {
	// The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
	KeyArn *string `pulumi:"keyArn"`
}

type DefaultKmsKeyState struct {
	// The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
	KeyArn pulumi.StringPtrInput
}

func (DefaultKmsKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultKmsKeyState)(nil)).Elem()
}

type defaultKmsKeyArgs struct {
	// The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
	KeyArn string `pulumi:"keyArn"`
}

// The set of arguments for constructing a DefaultKmsKey resource.
type DefaultKmsKeyArgs struct {
	// The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
	KeyArn pulumi.StringInput
}

func (DefaultKmsKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultKmsKeyArgs)(nil)).Elem()
}

type DefaultKmsKeyInput interface {
	pulumi.Input

	ToDefaultKmsKeyOutput() DefaultKmsKeyOutput
	ToDefaultKmsKeyOutputWithContext(ctx context.Context) DefaultKmsKeyOutput
}

func (*DefaultKmsKey) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultKmsKey)(nil))
}

func (i *DefaultKmsKey) ToDefaultKmsKeyOutput() DefaultKmsKeyOutput {
	return i.ToDefaultKmsKeyOutputWithContext(context.Background())
}

func (i *DefaultKmsKey) ToDefaultKmsKeyOutputWithContext(ctx context.Context) DefaultKmsKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultKmsKeyOutput)
}

type DefaultKmsKeyOutput struct {
	*pulumi.OutputState
}

func (DefaultKmsKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultKmsKey)(nil))
}

func (o DefaultKmsKeyOutput) ToDefaultKmsKeyOutput() DefaultKmsKeyOutput {
	return o
}

func (o DefaultKmsKeyOutput) ToDefaultKmsKeyOutputWithContext(ctx context.Context) DefaultKmsKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DefaultKmsKeyOutput{})
}
