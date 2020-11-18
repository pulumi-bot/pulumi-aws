// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a FSx Windows File System. See the [FSx Windows Guide](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/what-is.html) for more information.
 *
 * > **NOTE:** Either the `activeDirectoryId` argument or `selfManagedActiveDirectory` configuration block must be specified.
 *
 * ## Example Usage
 * ### Using AWS Directory Service
 *
 * Additional information for using AWS Directory Service with Windows File Systems can be found in the [FSx Windows Guide](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/fsx-aws-managed-ad.html).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.fsx.WindowsFileSystem("example", {
 *     activeDirectoryId: aws_directory_service_directory.example.id,
 *     kmsKeyId: aws_kms_key.example.arn,
 *     storageCapacity: 300,
 *     subnetIds: [aws_subnet.example.id],
 *     throughputCapacity: 1024,
 * });
 * ```
 * ### Using a Self-Managed Microsoft Active Directory
 *
 * Additional information for using AWS Directory Service with Windows File Systems can be found in the [FSx Windows Guide](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/self-managed-AD.html).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.fsx.WindowsFileSystem("example", {
 *     kmsKeyId: aws_kms_key.example.arn,
 *     storageCapacity: 300,
 *     subnetIds: [aws_subnet.example.id],
 *     throughputCapacity: 1024,
 *     selfManagedActiveDirectory: {
 *         dnsIps: [
 *             "10.0.0.111",
 *             "10.0.0.222",
 *         ],
 *         domainName: "corp.example.com",
 *         password: "avoid-plaintext-passwords",
 *         username: "Admin",
 *     },
 * });
 * ```
 */
export class WindowsFileSystem extends pulumi.CustomResource {
    /**
     * Get an existing WindowsFileSystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WindowsFileSystemState, opts?: pulumi.CustomResourceOptions): WindowsFileSystem {
        return new WindowsFileSystem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:fsx/windowsFileSystem:WindowsFileSystem';

    /**
     * Returns true if the given object is an instance of WindowsFileSystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WindowsFileSystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WindowsFileSystem.__pulumiType;
    }

    /**
     * The ID for an existing Microsoft Active Directory instance that the file system should join when it's created. Cannot be specified with `selfManagedActiveDirectory`.
     */
    public readonly activeDirectoryId!: pulumi.Output<string | undefined>;
    /**
     * Amazon Resource Name of the file system.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The number of days to retain automatic backups. Minimum of `0` and maximum of `90`. Defaults to `7`. Set to `0` to disable.
     */
    public readonly automaticBackupRetentionDays!: pulumi.Output<number | undefined>;
    /**
     * A boolean flag indicating whether tags on the file system should be copied to backups. Defaults to `false`.
     */
    public readonly copyTagsToBackups!: pulumi.Output<boolean | undefined>;
    /**
     * The preferred time (in `HH:MM` format) to take daily automatic backups, in the UTC time zone.
     */
    public readonly dailyAutomaticBackupStartTime!: pulumi.Output<string>;
    /**
     * Specifies the file system deployment type, valid values are `MULTI_AZ_1`, `SINGLE_AZ_1` and `SINGLE_AZ_2`. Default value is `SINGLE_AZ_1`.
     */
    public readonly deploymentType!: pulumi.Output<string | undefined>;
    /**
     * DNS name for the file system, e.g. `fs-12345678.corp.example.com` (domain name matching the Active Directory domain name)
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * ARN for the KMS Key to encrypt the file system at rest. Defaults to an AWS managed KMS Key.
     */
    public readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * Set of Elastic Network Interface identifiers from which the file system is accessible.
     */
    public /*out*/ readonly networkInterfaceIds!: pulumi.Output<string[]>;
    /**
     * AWS account identifier that created the file system.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * The IP address of the primary, or preferred, file server.
     */
    public /*out*/ readonly preferredFileServerIp!: pulumi.Output<string>;
    /**
     * Specifies the subnet in which you want the preferred file server to be located. Required for when deployment type is `MULTI_AZ_1`.
     */
    public readonly preferredSubnetId!: pulumi.Output<string>;
    /**
     * For `MULTI_AZ_1` deployment types, use this endpoint when performing administrative tasks on the file system using Amazon FSx Remote PowerShell. For `SINGLE_AZ_1` deployment types, this is the DNS name of the file system.
     */
    public /*out*/ readonly remoteAdministrationEndpoint!: pulumi.Output<string>;
    /**
     * A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Configuration block that Amazon FSx uses to join the Windows File Server instance to your self-managed (including on-premises) Microsoft Active Directory (AD) directory. Cannot be specified with `activeDirectoryId`. Detailed below.
     */
    public readonly selfManagedActiveDirectory!: pulumi.Output<outputs.fsx.WindowsFileSystemSelfManagedActiveDirectory | undefined>;
    /**
     * When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
     */
    public readonly skipFinalBackup!: pulumi.Output<boolean | undefined>;
    /**
     * Storage capacity (GiB) of the file system. Minimum of 32 and maximum of 65536. If the storage type is set to `HDD` the minimum value is 2000.
     */
    public readonly storageCapacity!: pulumi.Output<number>;
    /**
     * Specifies the storage type, Valid values are `SSD` and `HDD`. `HDD` is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types. Default value is `SSD`.
     */
    public readonly storageType!: pulumi.Output<string | undefined>;
    /**
     * A list of IDs for the subnets that the file system will be accessible from. To specify more than a single subnet set `deploymentType` to `MULTI_AZ_1`.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * A map of tags to assign to the file system.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Throughput (megabytes per second) of the file system in power of 2 increments. Minimum of `8` and maximum of `2048`.
     */
    public readonly throughputCapacity!: pulumi.Output<number>;
    /**
     * Identifier of the Virtual Private Cloud for the file system.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
     */
    public readonly weeklyMaintenanceStartTime!: pulumi.Output<string>;

    /**
     * Create a WindowsFileSystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WindowsFileSystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WindowsFileSystemArgs | WindowsFileSystemState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as WindowsFileSystemState | undefined;
            inputs["activeDirectoryId"] = state ? state.activeDirectoryId : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["automaticBackupRetentionDays"] = state ? state.automaticBackupRetentionDays : undefined;
            inputs["copyTagsToBackups"] = state ? state.copyTagsToBackups : undefined;
            inputs["dailyAutomaticBackupStartTime"] = state ? state.dailyAutomaticBackupStartTime : undefined;
            inputs["deploymentType"] = state ? state.deploymentType : undefined;
            inputs["dnsName"] = state ? state.dnsName : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["networkInterfaceIds"] = state ? state.networkInterfaceIds : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["preferredFileServerIp"] = state ? state.preferredFileServerIp : undefined;
            inputs["preferredSubnetId"] = state ? state.preferredSubnetId : undefined;
            inputs["remoteAdministrationEndpoint"] = state ? state.remoteAdministrationEndpoint : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["selfManagedActiveDirectory"] = state ? state.selfManagedActiveDirectory : undefined;
            inputs["skipFinalBackup"] = state ? state.skipFinalBackup : undefined;
            inputs["storageCapacity"] = state ? state.storageCapacity : undefined;
            inputs["storageType"] = state ? state.storageType : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["throughputCapacity"] = state ? state.throughputCapacity : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
            inputs["weeklyMaintenanceStartTime"] = state ? state.weeklyMaintenanceStartTime : undefined;
        } else {
            const args = argsOrState as WindowsFileSystemArgs | undefined;
            if (!args || args.storageCapacity === undefined) {
                throw new Error("Missing required property 'storageCapacity'");
            }
            if (!args || args.subnetIds === undefined) {
                throw new Error("Missing required property 'subnetIds'");
            }
            if (!args || args.throughputCapacity === undefined) {
                throw new Error("Missing required property 'throughputCapacity'");
            }
            inputs["activeDirectoryId"] = args ? args.activeDirectoryId : undefined;
            inputs["automaticBackupRetentionDays"] = args ? args.automaticBackupRetentionDays : undefined;
            inputs["copyTagsToBackups"] = args ? args.copyTagsToBackups : undefined;
            inputs["dailyAutomaticBackupStartTime"] = args ? args.dailyAutomaticBackupStartTime : undefined;
            inputs["deploymentType"] = args ? args.deploymentType : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["preferredSubnetId"] = args ? args.preferredSubnetId : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["selfManagedActiveDirectory"] = args ? args.selfManagedActiveDirectory : undefined;
            inputs["skipFinalBackup"] = args ? args.skipFinalBackup : undefined;
            inputs["storageCapacity"] = args ? args.storageCapacity : undefined;
            inputs["storageType"] = args ? args.storageType : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["throughputCapacity"] = args ? args.throughputCapacity : undefined;
            inputs["weeklyMaintenanceStartTime"] = args ? args.weeklyMaintenanceStartTime : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["dnsName"] = undefined /*out*/;
            inputs["networkInterfaceIds"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
            inputs["preferredFileServerIp"] = undefined /*out*/;
            inputs["remoteAdministrationEndpoint"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(WindowsFileSystem.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WindowsFileSystem resources.
 */
export interface WindowsFileSystemState {
    /**
     * The ID for an existing Microsoft Active Directory instance that the file system should join when it's created. Cannot be specified with `selfManagedActiveDirectory`.
     */
    readonly activeDirectoryId?: pulumi.Input<string>;
    /**
     * Amazon Resource Name of the file system.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The number of days to retain automatic backups. Minimum of `0` and maximum of `90`. Defaults to `7`. Set to `0` to disable.
     */
    readonly automaticBackupRetentionDays?: pulumi.Input<number>;
    /**
     * A boolean flag indicating whether tags on the file system should be copied to backups. Defaults to `false`.
     */
    readonly copyTagsToBackups?: pulumi.Input<boolean>;
    /**
     * The preferred time (in `HH:MM` format) to take daily automatic backups, in the UTC time zone.
     */
    readonly dailyAutomaticBackupStartTime?: pulumi.Input<string>;
    /**
     * Specifies the file system deployment type, valid values are `MULTI_AZ_1`, `SINGLE_AZ_1` and `SINGLE_AZ_2`. Default value is `SINGLE_AZ_1`.
     */
    readonly deploymentType?: pulumi.Input<string>;
    /**
     * DNS name for the file system, e.g. `fs-12345678.corp.example.com` (domain name matching the Active Directory domain name)
     */
    readonly dnsName?: pulumi.Input<string>;
    /**
     * ARN for the KMS Key to encrypt the file system at rest. Defaults to an AWS managed KMS Key.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Set of Elastic Network Interface identifiers from which the file system is accessible.
     */
    readonly networkInterfaceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * AWS account identifier that created the file system.
     */
    readonly ownerId?: pulumi.Input<string>;
    /**
     * The IP address of the primary, or preferred, file server.
     */
    readonly preferredFileServerIp?: pulumi.Input<string>;
    /**
     * Specifies the subnet in which you want the preferred file server to be located. Required for when deployment type is `MULTI_AZ_1`.
     */
    readonly preferredSubnetId?: pulumi.Input<string>;
    /**
     * For `MULTI_AZ_1` deployment types, use this endpoint when performing administrative tasks on the file system using Amazon FSx Remote PowerShell. For `SINGLE_AZ_1` deployment types, this is the DNS name of the file system.
     */
    readonly remoteAdministrationEndpoint?: pulumi.Input<string>;
    /**
     * A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration block that Amazon FSx uses to join the Windows File Server instance to your self-managed (including on-premises) Microsoft Active Directory (AD) directory. Cannot be specified with `activeDirectoryId`. Detailed below.
     */
    readonly selfManagedActiveDirectory?: pulumi.Input<inputs.fsx.WindowsFileSystemSelfManagedActiveDirectory>;
    /**
     * When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
     */
    readonly skipFinalBackup?: pulumi.Input<boolean>;
    /**
     * Storage capacity (GiB) of the file system. Minimum of 32 and maximum of 65536. If the storage type is set to `HDD` the minimum value is 2000.
     */
    readonly storageCapacity?: pulumi.Input<number>;
    /**
     * Specifies the storage type, Valid values are `SSD` and `HDD`. `HDD` is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types. Default value is `SSD`.
     */
    readonly storageType?: pulumi.Input<string>;
    /**
     * A list of IDs for the subnets that the file system will be accessible from. To specify more than a single subnet set `deploymentType` to `MULTI_AZ_1`.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the file system.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Throughput (megabytes per second) of the file system in power of 2 increments. Minimum of `8` and maximum of `2048`.
     */
    readonly throughputCapacity?: pulumi.Input<number>;
    /**
     * Identifier of the Virtual Private Cloud for the file system.
     */
    readonly vpcId?: pulumi.Input<string>;
    /**
     * The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
     */
    readonly weeklyMaintenanceStartTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WindowsFileSystem resource.
 */
export interface WindowsFileSystemArgs {
    /**
     * The ID for an existing Microsoft Active Directory instance that the file system should join when it's created. Cannot be specified with `selfManagedActiveDirectory`.
     */
    readonly activeDirectoryId?: pulumi.Input<string>;
    /**
     * The number of days to retain automatic backups. Minimum of `0` and maximum of `90`. Defaults to `7`. Set to `0` to disable.
     */
    readonly automaticBackupRetentionDays?: pulumi.Input<number>;
    /**
     * A boolean flag indicating whether tags on the file system should be copied to backups. Defaults to `false`.
     */
    readonly copyTagsToBackups?: pulumi.Input<boolean>;
    /**
     * The preferred time (in `HH:MM` format) to take daily automatic backups, in the UTC time zone.
     */
    readonly dailyAutomaticBackupStartTime?: pulumi.Input<string>;
    /**
     * Specifies the file system deployment type, valid values are `MULTI_AZ_1`, `SINGLE_AZ_1` and `SINGLE_AZ_2`. Default value is `SINGLE_AZ_1`.
     */
    readonly deploymentType?: pulumi.Input<string>;
    /**
     * ARN for the KMS Key to encrypt the file system at rest. Defaults to an AWS managed KMS Key.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies the subnet in which you want the preferred file server to be located. Required for when deployment type is `MULTI_AZ_1`.
     */
    readonly preferredSubnetId?: pulumi.Input<string>;
    /**
     * A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration block that Amazon FSx uses to join the Windows File Server instance to your self-managed (including on-premises) Microsoft Active Directory (AD) directory. Cannot be specified with `activeDirectoryId`. Detailed below.
     */
    readonly selfManagedActiveDirectory?: pulumi.Input<inputs.fsx.WindowsFileSystemSelfManagedActiveDirectory>;
    /**
     * When enabled, will skip the default final backup taken when the file system is deleted. This configuration must be applied separately before attempting to delete the resource to have the desired behavior. Defaults to `false`.
     */
    readonly skipFinalBackup?: pulumi.Input<boolean>;
    /**
     * Storage capacity (GiB) of the file system. Minimum of 32 and maximum of 65536. If the storage type is set to `HDD` the minimum value is 2000.
     */
    readonly storageCapacity: pulumi.Input<number>;
    /**
     * Specifies the storage type, Valid values are `SSD` and `HDD`. `HDD` is supported on `SINGLE_AZ_2` and `MULTI_AZ_1` Windows file system deployment types. Default value is `SSD`.
     */
    readonly storageType?: pulumi.Input<string>;
    /**
     * A list of IDs for the subnets that the file system will be accessible from. To specify more than a single subnet set `deploymentType` to `MULTI_AZ_1`.
     */
    readonly subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the file system.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Throughput (megabytes per second) of the file system in power of 2 increments. Minimum of `8` and maximum of `2048`.
     */
    readonly throughputCapacity: pulumi.Input<number>;
    /**
     * The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
     */
    readonly weeklyMaintenanceStartTime?: pulumi.Input<string>;
}
