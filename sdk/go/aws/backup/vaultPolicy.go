// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS Backup vault policy resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/backup"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleVault, err := backup.NewVault(ctx, "exampleVault", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = backup.NewVaultPolicy(ctx, "exampleVaultPolicy", &backup.VaultPolicyArgs{
// 			BackupVaultName: exampleVault.Name,
// 			Policy: exampleVault.Arn.ApplyT(func(arn string) (string, error) {
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Id\": \"default\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"default\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"AWS\": \"*\"\n", "      },\n", "      \"Action\": [\n", "		\"backup:DescribeBackupVault\",\n", "		\"backup:DeleteBackupVault\",\n", "		\"backup:PutBackupVaultAccessPolicy\",\n", "		\"backup:DeleteBackupVaultAccessPolicy\",\n", "		\"backup:GetBackupVaultAccessPolicy\",\n", "		\"backup:StartBackupJob\",\n", "		\"backup:GetBackupVaultNotifications\",\n", "		\"backup:PutBackupVaultNotifications\"\n", "      ],\n", "      \"Resource\": \"", arn, "\"\n", "    }\n", "  ]\n", "}\n"), nil
// 			}).(pulumi.StringOutput),
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
// Backup vault policy can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:backup/vaultPolicy:VaultPolicy test TestVault
// ```
type VaultPolicy struct {
	pulumi.CustomResourceState

	// The ARN of the vault.
	BackupVaultArn pulumi.StringOutput `pulumi:"backupVaultArn"`
	// Name of the backup vault to add policy for.
	BackupVaultName pulumi.StringOutput `pulumi:"backupVaultName"`
	// The backup vault access policy document in JSON format.
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewVaultPolicy registers a new resource with the given unique name, arguments, and options.
func NewVaultPolicy(ctx *pulumi.Context,
	name string, args *VaultPolicyArgs, opts ...pulumi.ResourceOption) (*VaultPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupVaultName == nil {
		return nil, errors.New("invalid value for required argument 'BackupVaultName'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource VaultPolicy
	err := ctx.RegisterResource("aws:backup/vaultPolicy:VaultPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVaultPolicy gets an existing VaultPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVaultPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VaultPolicyState, opts ...pulumi.ResourceOption) (*VaultPolicy, error) {
	var resource VaultPolicy
	err := ctx.ReadResource("aws:backup/vaultPolicy:VaultPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VaultPolicy resources.
type vaultPolicyState struct {
	// The ARN of the vault.
	BackupVaultArn *string `pulumi:"backupVaultArn"`
	// Name of the backup vault to add policy for.
	BackupVaultName *string `pulumi:"backupVaultName"`
	// The backup vault access policy document in JSON format.
	Policy *string `pulumi:"policy"`
}

type VaultPolicyState struct {
	// The ARN of the vault.
	BackupVaultArn pulumi.StringPtrInput
	// Name of the backup vault to add policy for.
	BackupVaultName pulumi.StringPtrInput
	// The backup vault access policy document in JSON format.
	Policy pulumi.StringPtrInput
}

func (VaultPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultPolicyState)(nil)).Elem()
}

type vaultPolicyArgs struct {
	// Name of the backup vault to add policy for.
	BackupVaultName string `pulumi:"backupVaultName"`
	// The backup vault access policy document in JSON format.
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a VaultPolicy resource.
type VaultPolicyArgs struct {
	// Name of the backup vault to add policy for.
	BackupVaultName pulumi.StringInput
	// The backup vault access policy document in JSON format.
	Policy pulumi.StringInput
}

func (VaultPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultPolicyArgs)(nil)).Elem()
}

type VaultPolicyInput interface {
	pulumi.Input

	ToVaultPolicyOutput() VaultPolicyOutput
	ToVaultPolicyOutputWithContext(ctx context.Context) VaultPolicyOutput
}

func (*VaultPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPolicy)(nil))
}

func (i *VaultPolicy) ToVaultPolicyOutput() VaultPolicyOutput {
	return i.ToVaultPolicyOutputWithContext(context.Background())
}

func (i *VaultPolicy) ToVaultPolicyOutputWithContext(ctx context.Context) VaultPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPolicyOutput)
}

type VaultPolicyOutput struct {
	*pulumi.OutputState
}

func (VaultPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPolicy)(nil))
}

func (o VaultPolicyOutput) ToVaultPolicyOutput() VaultPolicyOutput {
	return o
}

func (o VaultPolicyOutput) ToVaultPolicyOutputWithContext(ctx context.Context) VaultPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VaultPolicyOutput{})
}
