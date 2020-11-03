// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// > **Note:** There is only a single account alias per AWS account.
//
// Manages the account alias for the AWS Account.
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
// 		_, err := iam.NewAccountAlias(ctx, "alias", &iam.AccountAliasArgs{
// 			AccountAlias: pulumi.String("my-account-alias"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AccountAlias struct {
	pulumi.CustomResourceState

	// The account alias
	AccountAlias pulumi.StringOutput `pulumi:"accountAlias"`
}

// NewAccountAlias registers a new resource with the given unique name, arguments, and options.
func NewAccountAlias(ctx *pulumi.Context,
	name string, args *AccountAliasArgs, opts ...pulumi.ResourceOption) (*AccountAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AccountAlias == nil {
		return nil, errors.New("invalid value for required argument 'AccountAlias'")
	}
	var resource AccountAlias
	err := ctx.RegisterResource("aws:iam/accountAlias:AccountAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountAlias gets an existing AccountAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountAliasState, opts ...pulumi.ResourceOption) (*AccountAlias, error) {
	var resource AccountAlias
	err := ctx.ReadResource("aws:iam/accountAlias:AccountAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountAlias resources.
type accountAliasState struct {
	// The account alias
	AccountAlias *string `pulumi:"accountAlias"`
}

type AccountAliasState struct {
	// The account alias
	AccountAlias pulumi.StringPtrInput
}

func (AccountAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountAliasState)(nil)).Elem()
}

type accountAliasArgs struct {
	// The account alias
	AccountAlias string `pulumi:"accountAlias"`
}

// The set of arguments for constructing a AccountAlias resource.
type AccountAliasArgs struct {
	// The account alias
	AccountAlias pulumi.StringInput
}

func (AccountAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountAliasArgs)(nil)).Elem()
}
