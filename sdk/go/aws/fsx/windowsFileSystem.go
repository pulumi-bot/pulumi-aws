// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fsx

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a FSx Windows File System. See the [FSx Windows Guide](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/what-is.html) for more information.
//
// > **NOTE:** Either the `activeDirectoryId` argument or `selfManagedActiveDirectory` configuration block must be specified.
//
// ## Example Usage
// ### Using AWS Directory Service
//
// Additional information for using AWS Directory Service with Windows File Systems can be found in the [FSx Windows Guide](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/fsx-aws-managed-ad.html).
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/fsx"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fsx.NewWindowsFileSystem(ctx, "example", &fsx.WindowsFileSystemArgs{
// 			ActiveDirectoryId: pulumi.Any(aws_directory_service_directory.Example.Id),
// 			KmsKeyId:          pulumi.Any(aws_kms_key.Example.Arn),
// 			StorageCapacity:   pulumi.Int(300),
// 			SubnetIds: pulumi.StringArray{
// 				pulumi.Any(aws_subnet.Example.Id),
// 			},
// 			ThroughputCapacity: pulumi.Int(1024),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using a Self-Managed Microsoft Active Directory
//
// Additional information for using AWS Directory Service with Windows File Systems can be found in the [FSx Windows Guide](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/self-managed-AD.html).
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/fsx"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fsx.NewWindowsFileSystem(ctx, "example", &fsx.WindowsFileSystemArgs{
// 			KmsKeyId:        pulumi.Any(aws_kms_key.Example.Arn),
// 			StorageCapacity: pulumi.Int(300),
// 			SubnetIds: pulumi.StringArray{
// 				pulumi.Any(aws_subnet.Example.Id),
// 			},
// 			ThroughputCapacity: pulumi.Int(1024),
// 			SelfManagedActiveDirectory: &fsx.WindowsFileSystemSelfManagedActiveDirectoryArgs{
// 				DnsIps: pulumi.StringArray{
// 					pulumi.String("10.0.0.111"),
// 					pulumi.String("10.0.0.222"),
// 				},
// 				DomainName: pulumi.String("corp.example.com"),
// 				Password:   pulumi.String("avoid-plaintext-passwords"),
// 				Username:   pulumi.String("Admin"),
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
// FSx File Systems can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:fsx/windowsFileSystem:WindowsFileSystem example fs-543ab12b1ca672f33
// ```
//
//  Certain resource arguments, like `security_group_ids` and the `self_managed_active_directory` configuation block `password`, do not have a FSx API method for reading the information after creation. If these arguments are set in the provider configuration on an imported resource, the povider will always show a difference. To workaround this behavior, either omit the argument from the configuration or use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to hide the difference, e.g. hcl resource "aws_fsx_windows_file_system" "example" {
//
// # ... other configuration ...
//
//  security_group_ids = [aws_security_group.example.id]
//
// # There is no FSx API for reading security_group_ids
//
//  lifecycle {
//
//  ignore_changes = [security_group_ids]
//
//  } }
type WindowsFileSystem struct {
	pulumi.CustomResourceState

	// The ID for an existing Microsoft Active Directory instance that the file system should join when it's created. Cannot be specified with `selfManagedActiveDirectory`.
	ActiveDirectoryId pulumi.StringPtrOutput `pulumi:"activeDirectoryId"`
	// Amazon Resource Name of the file system.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The number of days to retain automatic backups. Minimum of `0` and maximum of `90`. Defaults to `7`. Set to `0` to disable.
	AutomaticBackupRetentionDays pulumi.IntPtrOutput `pulumi:"automaticBackupRetentionDays"`
	// A boolean flag indicating whether tags on the file system should be copied to backups. Defaults to `false`.
	CopyTagsToBackups pulumi.BoolPtrOutput `pulumi:"copyTagsToBackups"`
	// The preferred time (in `HH:MM` format) to take daily automatic backups, in the UTC time zone.
	DailyAutomaticBackupStartTime pulumi.StringOutput `pulumi:"dailyAutomaticBackupStartTime"`
	// Specifies the file system deployment type, valid values are `MULTI_AZ_1`, `SINGLE_AZ_1` and `SINGLE_AZ_2`. Default value is `SINGLE_AZ_1`.
	DeploymentType pulumi.StringPtrOutput `pulumi:"deploymentType"`
	// DNS name for the file system, e.g. `fs-12345678.corp.example.com` (domain name matching the Active Directory domain name)
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// ARN for the KMS Key to encrypt the file system at rest. Defaults to an AWS managed KMS Key.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// Set of Elastic Network Interface identifiers from which the file system is accessible.
	NetworkInterfaceIds pulumi.StringArrayOutput `pulumi:"networkInterfaceIds"`
	// AWS account identifier that created the file system.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The IP address of the primary, or preferred, file server.
	PreferredFileServerIp pulumi.StringOutput `pulumi:"preferredFileServerIp"`
	// Specifies the subnet in which you want the preferred file server to be located. Required for when deployment type is `MULTI_AZ_1`.
	PreferredSubnetId pulumi.StringOutput `pulumi:"preferredSubnetId"`
	// For `MULTI_AZ_1` deployment types, use this endpoint when performing administrative tasks on the file system using Amazon FSx Remote PowerShell. For `SINGLE_AZ_1` deployment types, this is the DNS name of the file system.
	RemoteAdministrationEndpoint pulumi.StringOutput `pulumi:"remoteAdministrationEndpoint"`
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Configuration block that Amazon FSx uses to join the Windows File Server instance to your self-managed (including on-premises) Microsoft Active Directory (AD) directory. Cannot be specified with `activeDirectoryId`. Detailed below.
	SelfManagedActiveDirectory WindowsFileSystemSelfManagedActiveDirectoryPtrOutput `pulumi:"selfManagedActiveDirectory"`
	// When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
	SkipFinalBackup pulumi.BoolPtrOutput `pulumi:"skipFinalBackup"`
	// Storage capacity (GiB) of the file system. Minimum of 32 and maximum of 65536. If the storage type is set to `HDD` the minimum value is 2000.
	StorageCapacity pulumi.IntOutput `pulumi:"storageCapacity"`
	// Specifies the storage type, Valid values are `SSD` and `HDD`. `HDD` is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types. Default value is `SSD`.
	StorageType pulumi.StringPtrOutput `pulumi:"storageType"`
	// A list of IDs for the subnets that the file system will be accessible from. To specify more than a single subnet set `deploymentType` to `MULTI_AZ_1`.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A map of tags to assign to the file system.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Throughput (megabytes per second) of the file system in power of 2 increments. Minimum of `8` and maximum of `2048`.
	ThroughputCapacity pulumi.IntOutput `pulumi:"throughputCapacity"`
	// Identifier of the Virtual Private Cloud for the file system.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime pulumi.StringOutput `pulumi:"weeklyMaintenanceStartTime"`
}

// NewWindowsFileSystem registers a new resource with the given unique name, arguments, and options.
func NewWindowsFileSystem(ctx *pulumi.Context,
	name string, args *WindowsFileSystemArgs, opts ...pulumi.ResourceOption) (*WindowsFileSystem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StorageCapacity == nil {
		return nil, errors.New("invalid value for required argument 'StorageCapacity'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	if args.ThroughputCapacity == nil {
		return nil, errors.New("invalid value for required argument 'ThroughputCapacity'")
	}
	var resource WindowsFileSystem
	err := ctx.RegisterResource("aws:fsx/windowsFileSystem:WindowsFileSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWindowsFileSystem gets an existing WindowsFileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWindowsFileSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WindowsFileSystemState, opts ...pulumi.ResourceOption) (*WindowsFileSystem, error) {
	var resource WindowsFileSystem
	err := ctx.ReadResource("aws:fsx/windowsFileSystem:WindowsFileSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WindowsFileSystem resources.
type windowsFileSystemState struct {
	// The ID for an existing Microsoft Active Directory instance that the file system should join when it's created. Cannot be specified with `selfManagedActiveDirectory`.
	ActiveDirectoryId *string `pulumi:"activeDirectoryId"`
	// Amazon Resource Name of the file system.
	Arn *string `pulumi:"arn"`
	// The number of days to retain automatic backups. Minimum of `0` and maximum of `90`. Defaults to `7`. Set to `0` to disable.
	AutomaticBackupRetentionDays *int `pulumi:"automaticBackupRetentionDays"`
	// A boolean flag indicating whether tags on the file system should be copied to backups. Defaults to `false`.
	CopyTagsToBackups *bool `pulumi:"copyTagsToBackups"`
	// The preferred time (in `HH:MM` format) to take daily automatic backups, in the UTC time zone.
	DailyAutomaticBackupStartTime *string `pulumi:"dailyAutomaticBackupStartTime"`
	// Specifies the file system deployment type, valid values are `MULTI_AZ_1`, `SINGLE_AZ_1` and `SINGLE_AZ_2`. Default value is `SINGLE_AZ_1`.
	DeploymentType *string `pulumi:"deploymentType"`
	// DNS name for the file system, e.g. `fs-12345678.corp.example.com` (domain name matching the Active Directory domain name)
	DnsName *string `pulumi:"dnsName"`
	// ARN for the KMS Key to encrypt the file system at rest. Defaults to an AWS managed KMS Key.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Set of Elastic Network Interface identifiers from which the file system is accessible.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// AWS account identifier that created the file system.
	OwnerId *string `pulumi:"ownerId"`
	// The IP address of the primary, or preferred, file server.
	PreferredFileServerIp *string `pulumi:"preferredFileServerIp"`
	// Specifies the subnet in which you want the preferred file server to be located. Required for when deployment type is `MULTI_AZ_1`.
	PreferredSubnetId *string `pulumi:"preferredSubnetId"`
	// For `MULTI_AZ_1` deployment types, use this endpoint when performing administrative tasks on the file system using Amazon FSx Remote PowerShell. For `SINGLE_AZ_1` deployment types, this is the DNS name of the file system.
	RemoteAdministrationEndpoint *string `pulumi:"remoteAdministrationEndpoint"`
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Configuration block that Amazon FSx uses to join the Windows File Server instance to your self-managed (including on-premises) Microsoft Active Directory (AD) directory. Cannot be specified with `activeDirectoryId`. Detailed below.
	SelfManagedActiveDirectory *WindowsFileSystemSelfManagedActiveDirectory `pulumi:"selfManagedActiveDirectory"`
	// When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
	SkipFinalBackup *bool `pulumi:"skipFinalBackup"`
	// Storage capacity (GiB) of the file system. Minimum of 32 and maximum of 65536. If the storage type is set to `HDD` the minimum value is 2000.
	StorageCapacity *int `pulumi:"storageCapacity"`
	// Specifies the storage type, Valid values are `SSD` and `HDD`. `HDD` is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types. Default value is `SSD`.
	StorageType *string `pulumi:"storageType"`
	// A list of IDs for the subnets that the file system will be accessible from. To specify more than a single subnet set `deploymentType` to `MULTI_AZ_1`.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the file system.
	Tags map[string]string `pulumi:"tags"`
	// Throughput (megabytes per second) of the file system in power of 2 increments. Minimum of `8` and maximum of `2048`.
	ThroughputCapacity *int `pulumi:"throughputCapacity"`
	// Identifier of the Virtual Private Cloud for the file system.
	VpcId *string `pulumi:"vpcId"`
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime *string `pulumi:"weeklyMaintenanceStartTime"`
}

type WindowsFileSystemState struct {
	// The ID for an existing Microsoft Active Directory instance that the file system should join when it's created. Cannot be specified with `selfManagedActiveDirectory`.
	ActiveDirectoryId pulumi.StringPtrInput
	// Amazon Resource Name of the file system.
	Arn pulumi.StringPtrInput
	// The number of days to retain automatic backups. Minimum of `0` and maximum of `90`. Defaults to `7`. Set to `0` to disable.
	AutomaticBackupRetentionDays pulumi.IntPtrInput
	// A boolean flag indicating whether tags on the file system should be copied to backups. Defaults to `false`.
	CopyTagsToBackups pulumi.BoolPtrInput
	// The preferred time (in `HH:MM` format) to take daily automatic backups, in the UTC time zone.
	DailyAutomaticBackupStartTime pulumi.StringPtrInput
	// Specifies the file system deployment type, valid values are `MULTI_AZ_1`, `SINGLE_AZ_1` and `SINGLE_AZ_2`. Default value is `SINGLE_AZ_1`.
	DeploymentType pulumi.StringPtrInput
	// DNS name for the file system, e.g. `fs-12345678.corp.example.com` (domain name matching the Active Directory domain name)
	DnsName pulumi.StringPtrInput
	// ARN for the KMS Key to encrypt the file system at rest. Defaults to an AWS managed KMS Key.
	KmsKeyId pulumi.StringPtrInput
	// Set of Elastic Network Interface identifiers from which the file system is accessible.
	NetworkInterfaceIds pulumi.StringArrayInput
	// AWS account identifier that created the file system.
	OwnerId pulumi.StringPtrInput
	// The IP address of the primary, or preferred, file server.
	PreferredFileServerIp pulumi.StringPtrInput
	// Specifies the subnet in which you want the preferred file server to be located. Required for when deployment type is `MULTI_AZ_1`.
	PreferredSubnetId pulumi.StringPtrInput
	// For `MULTI_AZ_1` deployment types, use this endpoint when performing administrative tasks on the file system using Amazon FSx Remote PowerShell. For `SINGLE_AZ_1` deployment types, this is the DNS name of the file system.
	RemoteAdministrationEndpoint pulumi.StringPtrInput
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds pulumi.StringArrayInput
	// Configuration block that Amazon FSx uses to join the Windows File Server instance to your self-managed (including on-premises) Microsoft Active Directory (AD) directory. Cannot be specified with `activeDirectoryId`. Detailed below.
	SelfManagedActiveDirectory WindowsFileSystemSelfManagedActiveDirectoryPtrInput
	// When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
	SkipFinalBackup pulumi.BoolPtrInput
	// Storage capacity (GiB) of the file system. Minimum of 32 and maximum of 65536. If the storage type is set to `HDD` the minimum value is 2000.
	StorageCapacity pulumi.IntPtrInput
	// Specifies the storage type, Valid values are `SSD` and `HDD`. `HDD` is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types. Default value is `SSD`.
	StorageType pulumi.StringPtrInput
	// A list of IDs for the subnets that the file system will be accessible from. To specify more than a single subnet set `deploymentType` to `MULTI_AZ_1`.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the file system.
	Tags pulumi.StringMapInput
	// Throughput (megabytes per second) of the file system in power of 2 increments. Minimum of `8` and maximum of `2048`.
	ThroughputCapacity pulumi.IntPtrInput
	// Identifier of the Virtual Private Cloud for the file system.
	VpcId pulumi.StringPtrInput
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime pulumi.StringPtrInput
}

func (WindowsFileSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*windowsFileSystemState)(nil)).Elem()
}

type windowsFileSystemArgs struct {
	// The ID for an existing Microsoft Active Directory instance that the file system should join when it's created. Cannot be specified with `selfManagedActiveDirectory`.
	ActiveDirectoryId *string `pulumi:"activeDirectoryId"`
	// The number of days to retain automatic backups. Minimum of `0` and maximum of `90`. Defaults to `7`. Set to `0` to disable.
	AutomaticBackupRetentionDays *int `pulumi:"automaticBackupRetentionDays"`
	// A boolean flag indicating whether tags on the file system should be copied to backups. Defaults to `false`.
	CopyTagsToBackups *bool `pulumi:"copyTagsToBackups"`
	// The preferred time (in `HH:MM` format) to take daily automatic backups, in the UTC time zone.
	DailyAutomaticBackupStartTime *string `pulumi:"dailyAutomaticBackupStartTime"`
	// Specifies the file system deployment type, valid values are `MULTI_AZ_1`, `SINGLE_AZ_1` and `SINGLE_AZ_2`. Default value is `SINGLE_AZ_1`.
	DeploymentType *string `pulumi:"deploymentType"`
	// ARN for the KMS Key to encrypt the file system at rest. Defaults to an AWS managed KMS Key.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Specifies the subnet in which you want the preferred file server to be located. Required for when deployment type is `MULTI_AZ_1`.
	PreferredSubnetId *string `pulumi:"preferredSubnetId"`
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Configuration block that Amazon FSx uses to join the Windows File Server instance to your self-managed (including on-premises) Microsoft Active Directory (AD) directory. Cannot be specified with `activeDirectoryId`. Detailed below.
	SelfManagedActiveDirectory *WindowsFileSystemSelfManagedActiveDirectory `pulumi:"selfManagedActiveDirectory"`
	// When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
	SkipFinalBackup *bool `pulumi:"skipFinalBackup"`
	// Storage capacity (GiB) of the file system. Minimum of 32 and maximum of 65536. If the storage type is set to `HDD` the minimum value is 2000.
	StorageCapacity int `pulumi:"storageCapacity"`
	// Specifies the storage type, Valid values are `SSD` and `HDD`. `HDD` is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types. Default value is `SSD`.
	StorageType *string `pulumi:"storageType"`
	// A list of IDs for the subnets that the file system will be accessible from. To specify more than a single subnet set `deploymentType` to `MULTI_AZ_1`.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the file system.
	Tags map[string]string `pulumi:"tags"`
	// Throughput (megabytes per second) of the file system in power of 2 increments. Minimum of `8` and maximum of `2048`.
	ThroughputCapacity int `pulumi:"throughputCapacity"`
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime *string `pulumi:"weeklyMaintenanceStartTime"`
}

// The set of arguments for constructing a WindowsFileSystem resource.
type WindowsFileSystemArgs struct {
	// The ID for an existing Microsoft Active Directory instance that the file system should join when it's created. Cannot be specified with `selfManagedActiveDirectory`.
	ActiveDirectoryId pulumi.StringPtrInput
	// The number of days to retain automatic backups. Minimum of `0` and maximum of `90`. Defaults to `7`. Set to `0` to disable.
	AutomaticBackupRetentionDays pulumi.IntPtrInput
	// A boolean flag indicating whether tags on the file system should be copied to backups. Defaults to `false`.
	CopyTagsToBackups pulumi.BoolPtrInput
	// The preferred time (in `HH:MM` format) to take daily automatic backups, in the UTC time zone.
	DailyAutomaticBackupStartTime pulumi.StringPtrInput
	// Specifies the file system deployment type, valid values are `MULTI_AZ_1`, `SINGLE_AZ_1` and `SINGLE_AZ_2`. Default value is `SINGLE_AZ_1`.
	DeploymentType pulumi.StringPtrInput
	// ARN for the KMS Key to encrypt the file system at rest. Defaults to an AWS managed KMS Key.
	KmsKeyId pulumi.StringPtrInput
	// Specifies the subnet in which you want the preferred file server to be located. Required for when deployment type is `MULTI_AZ_1`.
	PreferredSubnetId pulumi.StringPtrInput
	// A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
	SecurityGroupIds pulumi.StringArrayInput
	// Configuration block that Amazon FSx uses to join the Windows File Server instance to your self-managed (including on-premises) Microsoft Active Directory (AD) directory. Cannot be specified with `activeDirectoryId`. Detailed below.
	SelfManagedActiveDirectory WindowsFileSystemSelfManagedActiveDirectoryPtrInput
	// When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
	SkipFinalBackup pulumi.BoolPtrInput
	// Storage capacity (GiB) of the file system. Minimum of 32 and maximum of 65536. If the storage type is set to `HDD` the minimum value is 2000.
	StorageCapacity pulumi.IntInput
	// Specifies the storage type, Valid values are `SSD` and `HDD`. `HDD` is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types. Default value is `SSD`.
	StorageType pulumi.StringPtrInput
	// A list of IDs for the subnets that the file system will be accessible from. To specify more than a single subnet set `deploymentType` to `MULTI_AZ_1`.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the file system.
	Tags pulumi.StringMapInput
	// Throughput (megabytes per second) of the file system in power of 2 increments. Minimum of `8` and maximum of `2048`.
	ThroughputCapacity pulumi.IntInput
	// The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
	WeeklyMaintenanceStartTime pulumi.StringPtrInput
}

func (WindowsFileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*windowsFileSystemArgs)(nil)).Elem()
}

type WindowsFileSystemInput interface {
	pulumi.Input

	ToWindowsFileSystemOutput() WindowsFileSystemOutput
	ToWindowsFileSystemOutputWithContext(ctx context.Context) WindowsFileSystemOutput
}

func (*WindowsFileSystem) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsFileSystem)(nil))
}

func (i *WindowsFileSystem) ToWindowsFileSystemOutput() WindowsFileSystemOutput {
	return i.ToWindowsFileSystemOutputWithContext(context.Background())
}

func (i *WindowsFileSystem) ToWindowsFileSystemOutputWithContext(ctx context.Context) WindowsFileSystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsFileSystemOutput)
}

type WindowsFileSystemOutput struct {
	*pulumi.OutputState
}

func (WindowsFileSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsFileSystem)(nil))
}

func (o WindowsFileSystemOutput) ToWindowsFileSystemOutput() WindowsFileSystemOutput {
	return o
}

func (o WindowsFileSystemOutput) ToWindowsFileSystemOutputWithContext(ctx context.Context) WindowsFileSystemOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WindowsFileSystemOutput{})
}
