// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an existing backup vault.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/backup"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := backup.LookupVault(ctx, &backup.LookupVaultArgs{
// 			Name: "example_backup_vault",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVault(ctx *pulumi.Context, args *LookupVaultArgs, opts ...pulumi.InvokeOption) (*LookupVaultResult, error) {
	var rv LookupVaultResult
	err := ctx.Invoke("aws:backup/getVault:getVault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVault.
type LookupVaultArgs struct {
	// The name of the backup vault.
	Name string `pulumi:"name"`
	// Metadata that you can assign to help organize the resources that you create.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getVault.
type LookupVaultResult struct {
	// The ARN of the vault.
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The server-side encryption key that is used to protect your backups.
	KmsKeyArn string `pulumi:"kmsKeyArn"`
	Name      string `pulumi:"name"`
	// The number of recovery points that are stored in a backup vault.
	RecoveryPoints int `pulumi:"recoveryPoints"`
	// Metadata that you can assign to help organize the resources that you create.
	Tags map[string]string `pulumi:"tags"`
}

func LookupVaultApply(ctx *pulumi.Context, args LookupVaultApplyInput, opts ...pulumi.InvokeOption) LookupVaultResultOutput {
	return args.ToLookupVaultApplyOutput().ApplyT(func(v LookupVaultArgs) (LookupVaultResult, error) {
		r, err := LookupVault(ctx, &v, opts...)
		return *r, err

	}).(LookupVaultResultOutput)
}

// LookupVaultApplyInput is an input type that accepts LookupVaultApplyArgs and LookupVaultApplyOutput values.
// You can construct a concrete instance of `LookupVaultApplyInput` via:
//
//          LookupVaultApplyArgs{...}
type LookupVaultApplyInput interface {
	pulumi.Input

	ToLookupVaultApplyOutput() LookupVaultApplyOutput
	ToLookupVaultApplyOutputWithContext(context.Context) LookupVaultApplyOutput
}

// A collection of arguments for invoking getVault.
type LookupVaultApplyArgs struct {
	// The name of the backup vault.
	Name pulumi.StringInput `pulumi:"name"`
	// Metadata that you can assign to help organize the resources that you create.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupVaultApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVaultArgs)(nil)).Elem()
}

func (i LookupVaultApplyArgs) ToLookupVaultApplyOutput() LookupVaultApplyOutput {
	return i.ToLookupVaultApplyOutputWithContext(context.Background())
}

func (i LookupVaultApplyArgs) ToLookupVaultApplyOutputWithContext(ctx context.Context) LookupVaultApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupVaultApplyOutput)
}

// A collection of arguments for invoking getVault.
type LookupVaultApplyOutput struct{ *pulumi.OutputState }

func (LookupVaultApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVaultArgs)(nil)).Elem()
}

func (o LookupVaultApplyOutput) ToLookupVaultApplyOutput() LookupVaultApplyOutput {
	return o
}

func (o LookupVaultApplyOutput) ToLookupVaultApplyOutputWithContext(ctx context.Context) LookupVaultApplyOutput {
	return o
}

// The name of the backup vault.
func (o LookupVaultApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVaultArgs) string { return v.Name }).(pulumi.StringOutput)
}

// Metadata that you can assign to help organize the resources that you create.
func (o LookupVaultApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVaultArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getVault.
type LookupVaultResultOutput struct{ *pulumi.OutputState }

func (LookupVaultResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVaultResult)(nil)).Elem()
}

func (o LookupVaultResultOutput) ToLookupVaultResultOutput() LookupVaultResultOutput {
	return o
}

func (o LookupVaultResultOutput) ToLookupVaultResultOutputWithContext(ctx context.Context) LookupVaultResultOutput {
	return o
}

// The ARN of the vault.
func (o LookupVaultResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVaultResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVaultResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVaultResult) string { return v.Id }).(pulumi.StringOutput)
}

// The server-side encryption key that is used to protect your backups.
func (o LookupVaultResultOutput) KmsKeyArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVaultResult) string { return v.KmsKeyArn }).(pulumi.StringOutput)
}

func (o LookupVaultResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVaultResult) string { return v.Name }).(pulumi.StringOutput)
}

// The number of recovery points that are stored in a backup vault.
func (o LookupVaultResultOutput) RecoveryPoints() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVaultResult) int { return v.RecoveryPoints }).(pulumi.IntOutput)
}

// Metadata that you can assign to help organize the resources that you create.
func (o LookupVaultResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVaultResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVaultApplyOutput{})
	pulumi.RegisterOutputType(LookupVaultResultOutput{})
}
