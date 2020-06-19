// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glacier

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Glacier Vault Lock. You can refer to the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html) for a full explanation of the Glacier Vault Lock functionality.
//
// > **NOTE:** This resource allows you to test Glacier Vault Lock policies by setting the `completeLock` argument to `false`. When testing policies in this manner, the Glacier Vault Lock automatically expires after 24 hours and this provider will show this resource as needing recreation after that time. To permanently apply the policy, set the `completeLock` argument to `true`. When changing `completeLock` to `true`, it is expected the resource will show as recreating.
//
// !> **WARNING:** Once a Glacier Vault Lock is completed, it is immutable. The deletion of the Glacier Vault Lock is not be possible and attempting to remove it from this provider will return an error. Set the `ignoreDeletionError` argument to `true` and apply this configuration before attempting to delete this resource via this provider or remove this resource from this provider's management.
//
// ## Example Usage
// ### Testing Glacier Vault Lock Policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glacier"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleVault, err := glacier.NewVault(ctx, "exampleVault", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = glacier.NewVaultLock(ctx, "exampleVaultLock", &glacier.VaultLockArgs{
// 			CompleteLock: pulumi.Bool(false),
// 			Policy: examplePolicyDocument.ApplyT(func(examplePolicyDocument iam.GetPolicyDocumentResult) (string, error) {
// 				return examplePolicyDocument.Json, nil
// 			}).(pulumi.StringOutput),
// 			VaultName: exampleVault.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Permanently Applying Glacier Vault Lock Policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glacier"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = glacier.NewVaultLock(ctx, "example", &glacier.VaultLockArgs{
// 			CompleteLock: pulumi.Bool(true),
// 			Policy:       pulumi.String(data.Aws_iam_policy_document.Example.Json),
// 			VaultName:    pulumi.String(aws_glacier_vault.Example.Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VaultLock struct {
	pulumi.CustomResourceState

	// Boolean whether to permanently apply this Glacier Lock Policy. Once completed, this cannot be undone. If set to `false`, the Glacier Lock Policy remains in a testing mode for 24 hours. After that time, the Glacier Lock Policy is automatically removed by Glacier and the this provider resource will show as needing recreation. Changing this from `false` to `true` will show as resource recreation, which is expected. Changing this from `true` to `false` is not possible unless the Glacier Vault is recreated at the same time.
	CompleteLock pulumi.BoolOutput `pulumi:"completeLock"`
	// Allow this provider to ignore the error returned when attempting to delete the Glacier Lock Policy. This can be used to delete or recreate the Glacier Vault via this provider, for example, if the Glacier Vault Lock policy permits that action. This should only be used in conjunction with `completeLock` being set to `true`.
	IgnoreDeletionError pulumi.BoolPtrOutput `pulumi:"ignoreDeletionError"`
	// JSON string containing the IAM policy to apply as the Glacier Vault Lock policy.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The name of the Glacier Vault.
	VaultName pulumi.StringOutput `pulumi:"vaultName"`
}

// NewVaultLock registers a new resource with the given unique name, arguments, and options.
func NewVaultLock(ctx *pulumi.Context,
	name string, args *VaultLockArgs, opts ...pulumi.ResourceOption) (*VaultLock, error) {
	if args == nil || args.CompleteLock == nil {
		return nil, errors.New("missing required argument 'CompleteLock'")
	}
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil || args.VaultName == nil {
		return nil, errors.New("missing required argument 'VaultName'")
	}
	if args == nil {
		args = &VaultLockArgs{}
	}
	var resource VaultLock
	err := ctx.RegisterResource("aws:glacier/vaultLock:VaultLock", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVaultLock gets an existing VaultLock resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVaultLock(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VaultLockState, opts ...pulumi.ResourceOption) (*VaultLock, error) {
	var resource VaultLock
	err := ctx.ReadResource("aws:glacier/vaultLock:VaultLock", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VaultLock resources.
type vaultLockState struct {
	// Boolean whether to permanently apply this Glacier Lock Policy. Once completed, this cannot be undone. If set to `false`, the Glacier Lock Policy remains in a testing mode for 24 hours. After that time, the Glacier Lock Policy is automatically removed by Glacier and the this provider resource will show as needing recreation. Changing this from `false` to `true` will show as resource recreation, which is expected. Changing this from `true` to `false` is not possible unless the Glacier Vault is recreated at the same time.
	CompleteLock *bool `pulumi:"completeLock"`
	// Allow this provider to ignore the error returned when attempting to delete the Glacier Lock Policy. This can be used to delete or recreate the Glacier Vault via this provider, for example, if the Glacier Vault Lock policy permits that action. This should only be used in conjunction with `completeLock` being set to `true`.
	IgnoreDeletionError *bool `pulumi:"ignoreDeletionError"`
	// JSON string containing the IAM policy to apply as the Glacier Vault Lock policy.
	Policy *string `pulumi:"policy"`
	// The name of the Glacier Vault.
	VaultName *string `pulumi:"vaultName"`
}

type VaultLockState struct {
	// Boolean whether to permanently apply this Glacier Lock Policy. Once completed, this cannot be undone. If set to `false`, the Glacier Lock Policy remains in a testing mode for 24 hours. After that time, the Glacier Lock Policy is automatically removed by Glacier and the this provider resource will show as needing recreation. Changing this from `false` to `true` will show as resource recreation, which is expected. Changing this from `true` to `false` is not possible unless the Glacier Vault is recreated at the same time.
	CompleteLock pulumi.BoolPtrInput
	// Allow this provider to ignore the error returned when attempting to delete the Glacier Lock Policy. This can be used to delete or recreate the Glacier Vault via this provider, for example, if the Glacier Vault Lock policy permits that action. This should only be used in conjunction with `completeLock` being set to `true`.
	IgnoreDeletionError pulumi.BoolPtrInput
	// JSON string containing the IAM policy to apply as the Glacier Vault Lock policy.
	Policy pulumi.StringPtrInput
	// The name of the Glacier Vault.
	VaultName pulumi.StringPtrInput
}

func (VaultLockState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultLockState)(nil)).Elem()
}

type vaultLockArgs struct {
	// Boolean whether to permanently apply this Glacier Lock Policy. Once completed, this cannot be undone. If set to `false`, the Glacier Lock Policy remains in a testing mode for 24 hours. After that time, the Glacier Lock Policy is automatically removed by Glacier and the this provider resource will show as needing recreation. Changing this from `false` to `true` will show as resource recreation, which is expected. Changing this from `true` to `false` is not possible unless the Glacier Vault is recreated at the same time.
	CompleteLock bool `pulumi:"completeLock"`
	// Allow this provider to ignore the error returned when attempting to delete the Glacier Lock Policy. This can be used to delete or recreate the Glacier Vault via this provider, for example, if the Glacier Vault Lock policy permits that action. This should only be used in conjunction with `completeLock` being set to `true`.
	IgnoreDeletionError *bool `pulumi:"ignoreDeletionError"`
	// JSON string containing the IAM policy to apply as the Glacier Vault Lock policy.
	Policy string `pulumi:"policy"`
	// The name of the Glacier Vault.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a VaultLock resource.
type VaultLockArgs struct {
	// Boolean whether to permanently apply this Glacier Lock Policy. Once completed, this cannot be undone. If set to `false`, the Glacier Lock Policy remains in a testing mode for 24 hours. After that time, the Glacier Lock Policy is automatically removed by Glacier and the this provider resource will show as needing recreation. Changing this from `false` to `true` will show as resource recreation, which is expected. Changing this from `true` to `false` is not possible unless the Glacier Vault is recreated at the same time.
	CompleteLock pulumi.BoolInput
	// Allow this provider to ignore the error returned when attempting to delete the Glacier Lock Policy. This can be used to delete or recreate the Glacier Vault via this provider, for example, if the Glacier Vault Lock policy permits that action. This should only be used in conjunction with `completeLock` being set to `true`.
	IgnoreDeletionError pulumi.BoolPtrInput
	// JSON string containing the IAM policy to apply as the Glacier Vault Lock policy.
	Policy pulumi.StringInput
	// The name of the Glacier Vault.
	VaultName pulumi.StringInput
}

func (VaultLockArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultLockArgs)(nil)).Elem()
}
