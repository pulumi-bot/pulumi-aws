// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an alias for a KMS customer master key. AWS Console enforces 1-to-1 mapping between aliases & keys,
// but API (hence this provider too) allows you to create as many aliases as
// the [account limits](http://docs.aws.amazon.com/kms/latest/developerguide/limits.html) allow you.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		key, err := kms.NewKey(ctx, "key", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.NewAlias(ctx, "alias", &kms.AliasArgs{
// 			TargetKeyId: key.KeyId,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Alias struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the key alias.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates an unique alias beginning with the specified prefix.
	// The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The Amazon Resource Name (ARN) of the target key identifier.
	TargetKeyArn pulumi.StringOutput `pulumi:"targetKeyArn"`
	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	TargetKeyId pulumi.StringOutput `pulumi:"targetKeyId"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TargetKeyId == nil {
		return nil, errors.New("invalid value for required argument 'TargetKeyId'")
	}
	var resource Alias
	err := ctx.RegisterResource("aws:kms/alias:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("aws:kms/alias:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type aliasState struct {
	// The Amazon Resource Name (ARN) of the key alias.
	Arn *string `pulumi:"arn"`
	// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
	Name *string `pulumi:"name"`
	// Creates an unique alias beginning with the specified prefix.
	// The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The Amazon Resource Name (ARN) of the target key identifier.
	TargetKeyArn *string `pulumi:"targetKeyArn"`
	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	TargetKeyId *string `pulumi:"targetKeyId"`
}

type AliasState struct {
	// The Amazon Resource Name (ARN) of the key alias.
	Arn pulumi.StringPtrInput
	// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
	Name pulumi.StringPtrInput
	// Creates an unique alias beginning with the specified prefix.
	// The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the target key identifier.
	TargetKeyArn pulumi.StringPtrInput
	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	TargetKeyId pulumi.StringPtrInput
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
	Name *string `pulumi:"name"`
	// Creates an unique alias beginning with the specified prefix.
	// The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	TargetKeyId string `pulumi:"targetKeyId"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
	Name pulumi.StringPtrInput
	// Creates an unique alias beginning with the specified prefix.
	// The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	TargetKeyId pulumi.StringInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}

type AliasInput interface {
	pulumi.Input

	ToAliasOutput() AliasOutput
	ToAliasOutputWithContext(ctx context.Context) AliasOutput
}

func (Alias) ElementType() reflect.Type {
	return reflect.TypeOf((*Alias)(nil)).Elem()
}

func (i Alias) ToAliasOutput() AliasOutput {
	return i.ToAliasOutputWithContext(context.Background())
}

func (i Alias) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasOutput)
}

type AliasOutput struct {
	*pulumi.OutputState
}

func (AliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AliasOutput)(nil)).Elem()
}

func (o AliasOutput) ToAliasOutput() AliasOutput {
	return o
}

func (o AliasOutput) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AliasOutput{})
}
