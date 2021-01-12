// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glacier

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Glacier Vault Resource. You can refer to the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-vaults.html) for a full explanation of the Glacier Vault functionality
//
// > **NOTE:** When removing a Glacier Vault, the Vault must be empty.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/glacier"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/sns"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		awsSnsTopic, err := sns.NewTopic(ctx, "awsSnsTopic", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = glacier.NewVault(ctx, "myArchive", &glacier.VaultArgs{
// 			Notifications: glacier.VaultNotificationArray{
// 				&glacier.VaultNotificationArgs{
// 					SnsTopic: awsSnsTopic.Arn,
// 					Events: pulumi.StringArray{
// 						pulumi.String("ArchiveRetrievalCompleted"),
// 						pulumi.String("InventoryRetrievalCompleted"),
// 					},
// 				},
// 			},
// 			AccessPolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\":\"2012-10-17\",\n", "    \"Statement\":[\n", "       {\n", "          \"Sid\": \"add-read-only-perm\",\n", "          \"Principal\": \"*\",\n", "          \"Effect\": \"Allow\",\n", "          \"Action\": [\n", "             \"glacier:InitiateJob\",\n", "             \"glacier:GetJobOutput\"\n", "          ],\n", "          \"Resource\": \"arn:aws:glacier:eu-west-1:432981146916:vaults/MyArchive\"\n", "       }\n", "    ]\n", "}\n")),
// 			Tags: pulumi.StringMap{
// 				"Test": pulumi.String("MyArchive"),
// 			},
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
// Glacier Vaults can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:glacier/vault:Vault archive my_archive
// ```
type Vault struct {
	pulumi.CustomResourceState

	// The policy document. This is a JSON formatted string.
	// The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
	AccessPolicy pulumi.StringPtrOutput `pulumi:"accessPolicy"`
	// The ARN of the vault.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The URI of the vault that was created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
	Name pulumi.StringOutput `pulumi:"name"`
	// The notifications for the Vault. Fields documented below.
	Notifications VaultNotificationArrayOutput `pulumi:"notifications"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewVault registers a new resource with the given unique name, arguments, and options.
func NewVault(ctx *pulumi.Context,
	name string, args *VaultArgs, opts ...pulumi.ResourceOption) (*Vault, error) {
	if args == nil {
		args = &VaultArgs{}
	}

	var resource Vault
	err := ctx.RegisterResource("aws:glacier/vault:Vault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVault gets an existing Vault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VaultState, opts ...pulumi.ResourceOption) (*Vault, error) {
	var resource Vault
	err := ctx.ReadResource("aws:glacier/vault:Vault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vault resources.
type vaultState struct {
	// The policy document. This is a JSON formatted string.
	// The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
	AccessPolicy *string `pulumi:"accessPolicy"`
	// The ARN of the vault.
	Arn *string `pulumi:"arn"`
	// The URI of the vault that was created.
	Location *string `pulumi:"location"`
	// The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
	Name *string `pulumi:"name"`
	// The notifications for the Vault. Fields documented below.
	Notifications []VaultNotification `pulumi:"notifications"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type VaultState struct {
	// The policy document. This is a JSON formatted string.
	// The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
	AccessPolicy pulumi.StringPtrInput
	// The ARN of the vault.
	Arn pulumi.StringPtrInput
	// The URI of the vault that was created.
	Location pulumi.StringPtrInput
	// The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
	Name pulumi.StringPtrInput
	// The notifications for the Vault. Fields documented below.
	Notifications VaultNotificationArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultState)(nil)).Elem()
}

type vaultArgs struct {
	// The policy document. This is a JSON formatted string.
	// The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
	AccessPolicy *string `pulumi:"accessPolicy"`
	// The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
	Name *string `pulumi:"name"`
	// The notifications for the Vault. Fields documented below.
	Notifications []VaultNotification `pulumi:"notifications"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Vault resource.
type VaultArgs struct {
	// The policy document. This is a JSON formatted string.
	// The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
	AccessPolicy pulumi.StringPtrInput
	// The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
	Name pulumi.StringPtrInput
	// The notifications for the Vault. Fields documented below.
	Notifications VaultNotificationArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultArgs)(nil)).Elem()
}

type VaultInput interface {
	pulumi.Input

	ToVaultOutput() VaultOutput
	ToVaultOutputWithContext(ctx context.Context) VaultOutput
}

func (*Vault) ElementType() reflect.Type {
	return reflect.TypeOf((*Vault)(nil))
}

func (i *Vault) ToVaultOutput() VaultOutput {
	return i.ToVaultOutputWithContext(context.Background())
}

func (i *Vault) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultOutput)
}

func (i *Vault) ToVaultPtrOutput() VaultPtrOutput {
	return i.ToVaultPtrOutputWithContext(context.Background())
}

func (i *Vault) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPtrOutput)
}

type VaultPtrInput interface {
	pulumi.Input

	ToVaultPtrOutput() VaultPtrOutput
	ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput
}

type vaultPtrType VaultArgs

func (*vaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Vault)(nil))
}

func (i *vaultPtrType) ToVaultPtrOutput() VaultPtrOutput {
	return i.ToVaultPtrOutputWithContext(context.Background())
}

func (i *vaultPtrType) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultOutput).ToVaultPtrOutput()
}

// VaultArrayInput is an input type that accepts VaultArray and VaultArrayOutput values.
// You can construct a concrete instance of `VaultArrayInput` via:
//
//          VaultArray{ VaultArgs{...} }
type VaultArrayInput interface {
	pulumi.Input

	ToVaultArrayOutput() VaultArrayOutput
	ToVaultArrayOutputWithContext(context.Context) VaultArrayOutput
}

type VaultArray []VaultInput

func (VaultArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Vault)(nil))
}

func (i VaultArray) ToVaultArrayOutput() VaultArrayOutput {
	return i.ToVaultArrayOutputWithContext(context.Background())
}

func (i VaultArray) ToVaultArrayOutputWithContext(ctx context.Context) VaultArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultArrayOutput)
}

// VaultMapInput is an input type that accepts VaultMap and VaultMapOutput values.
// You can construct a concrete instance of `VaultMapInput` via:
//
//          VaultMap{ "key": VaultArgs{...} }
type VaultMapInput interface {
	pulumi.Input

	ToVaultMapOutput() VaultMapOutput
	ToVaultMapOutputWithContext(context.Context) VaultMapOutput
}

type VaultMap map[string]VaultInput

func (VaultMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Vault)(nil))
}

func (i VaultMap) ToVaultMapOutput() VaultMapOutput {
	return i.ToVaultMapOutputWithContext(context.Background())
}

func (i VaultMap) ToVaultMapOutputWithContext(ctx context.Context) VaultMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultMapOutput)
}

type VaultOutput struct {
	*pulumi.OutputState
}

func (VaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Vault)(nil))
}

func (o VaultOutput) ToVaultOutput() VaultOutput {
	return o
}

func (o VaultOutput) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return o
}

func (o VaultOutput) ToVaultPtrOutput() VaultPtrOutput {
	return o.ToVaultPtrOutputWithContext(context.Background())
}

func (o VaultOutput) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return o.ApplyT(func(v Vault) *Vault {
		return &v
	}).(VaultPtrOutput)
}

type VaultPtrOutput struct {
	*pulumi.OutputState
}

func (VaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vault)(nil))
}

func (o VaultPtrOutput) ToVaultPtrOutput() VaultPtrOutput {
	return o
}

func (o VaultPtrOutput) ToVaultPtrOutputWithContext(ctx context.Context) VaultPtrOutput {
	return o
}

type VaultArrayOutput struct{ *pulumi.OutputState }

func (VaultArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Vault)(nil))
}

func (o VaultArrayOutput) ToVaultArrayOutput() VaultArrayOutput {
	return o
}

func (o VaultArrayOutput) ToVaultArrayOutputWithContext(ctx context.Context) VaultArrayOutput {
	return o
}

func (o VaultArrayOutput) Index(i pulumi.IntInput) VaultOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Vault {
		return vs[0].([]Vault)[vs[1].(int)]
	}).(VaultOutput)
}

type VaultMapOutput struct{ *pulumi.OutputState }

func (VaultMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Vault)(nil))
}

func (o VaultMapOutput) ToVaultMapOutput() VaultMapOutput {
	return o
}

func (o VaultMapOutput) ToVaultMapOutputWithContext(ctx context.Context) VaultMapOutput {
	return o
}

func (o VaultMapOutput) MapIndex(k pulumi.StringInput) VaultOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Vault {
		return vs[0].(map[string]Vault)[vs[1].(string)]
	}).(VaultOutput)
}

func init() {
	pulumi.RegisterOutputType(VaultOutput{})
	pulumi.RegisterOutputType(VaultPtrOutput{})
	pulumi.RegisterOutputType(VaultArrayOutput{})
	pulumi.RegisterOutputType(VaultMapOutput{})
}
