// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Information about RDS orderable DB instances and valid parameter combinations.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.rds.getOrderableDbInstance({
 *     engine: "mysql",
 *     engineVersion: "5.7.22",
 *     licenseModel: "general-public-license",
 *     preferredInstanceClasses: [
 *         "db.r6.xlarge",
 *         "db.m4.large",
 *         "db.t3.small",
 *     ],
 *     storageType: "standard",
 * }));
 * ```
 *
 * Valid parameter combinations can also be found with `preferredEngineVersions` and/or `preferredInstanceClasses`.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.rds.getOrderableDbInstance({
 *     engine: "mysql",
 *     licenseModel: "general-public-license",
 *     preferredEngineVersions: [
 *         "5.6.35",
 *         "5.6.41",
 *         "5.6.44",
 *     ],
 *     preferredInstanceClasses: [
 *         "db.t2.small",
 *         "db.t3.medium",
 *         "db.t3.large",
 *     ],
 * }));
 * ```
 */
export function getOrderableDbInstance(args: GetOrderableDbInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetOrderableDbInstanceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:rds/getOrderableDbInstance:getOrderableDbInstance", {
        "availabilityZoneGroup": args.availabilityZoneGroup,
        "engine": args.engine,
        "engineVersion": args.engineVersion,
        "instanceClass": args.instanceClass,
        "licenseModel": args.licenseModel,
        "preferredEngineVersions": args.preferredEngineVersions,
        "preferredInstanceClasses": args.preferredInstanceClasses,
        "storageType": args.storageType,
        "supportsEnhancedMonitoring": args.supportsEnhancedMonitoring,
        "supportsGlobalDatabases": args.supportsGlobalDatabases,
        "supportsIamDatabaseAuthentication": args.supportsIamDatabaseAuthentication,
        "supportsIops": args.supportsIops,
        "supportsKerberosAuthentication": args.supportsKerberosAuthentication,
        "supportsPerformanceInsights": args.supportsPerformanceInsights,
        "supportsStorageAutoscaling": args.supportsStorageAutoscaling,
        "supportsStorageEncryption": args.supportsStorageEncryption,
        "vpc": args.vpc,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrderableDbInstance.
 */
export interface GetOrderableDbInstanceArgs {
    /**
     * Availability zone group.
     */
    availabilityZoneGroup?: string;
    /**
     * DB engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
     */
    engine: string;
    /**
     * Version of the DB engine. If none is provided, the AWS-defined default version will be used.
     */
    engineVersion?: string;
    /**
     * DB instance class. Examples of classes are `db.m3.2xlarge`, `db.t2.small`, and `db.m3.medium`.
     */
    instanceClass?: string;
    /**
     * License model. Examples of license models are `general-public-license`, `bring-your-own-license`, and `amazon-license`.
     */
    licenseModel?: string;
    /**
     * Ordered list of preferred RDS DB instance engine versions. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
     */
    preferredEngineVersions?: string[];
    /**
     * Ordered list of preferred RDS DB instance classes. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
     */
    preferredInstanceClasses?: string[];
    /**
     * Storage types. Examples of storage types are `standard`, `io1`, `gp2`, and `aurora`.
     */
    storageType?: string;
    /**
     * Enable this to ensure a DB instance supports Enhanced Monitoring at intervals from 1 to 60 seconds.
     */
    supportsEnhancedMonitoring?: boolean;
    /**
     * Enable this to ensure a DB instance supports Aurora global databases with a specific combination of other DB engine attributes.
     */
    supportsGlobalDatabases?: boolean;
    /**
     * Enable this to ensure a DB instance supports IAM database authentication.
     */
    supportsIamDatabaseAuthentication?: boolean;
    /**
     * Enable this to ensure a DB instance supports provisioned IOPS.
     */
    supportsIops?: boolean;
    /**
     * Enable this to ensure a DB instance supports Kerberos Authentication.
     */
    supportsKerberosAuthentication?: boolean;
    /**
     * Enable this to ensure a DB instance supports Performance Insights.
     */
    supportsPerformanceInsights?: boolean;
    /**
     * Enable this to ensure Amazon RDS can automatically scale storage for DB instances that use the specified DB instance class.
     */
    supportsStorageAutoscaling?: boolean;
    /**
     * Enable this to ensure a DB instance supports encrypted storage.
     */
    supportsStorageEncryption?: boolean;
    /**
     * Boolean that indicates whether to show only VPC or non-VPC offerings.
     */
    vpc?: boolean;
}

/**
 * A collection of values returned by getOrderableDbInstance.
 */
export interface GetOrderableDbInstanceResult {
    readonly availabilityZoneGroup: string;
    /**
     * Availability zones where the instance is available.
     */
    readonly availabilityZones: string[];
    readonly engine: string;
    readonly engineVersion: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceClass: string;
    readonly licenseModel: string;
    /**
     * Maximum total provisioned IOPS for a DB instance.
     */
    readonly maxIopsPerDbInstance: number;
    /**
     * Maximum provisioned IOPS per GiB for a DB instance.
     */
    readonly maxIopsPerGib: number;
    /**
     * Maximum storage size for a DB instance.
     */
    readonly maxStorageSize: number;
    /**
     * Minimum total provisioned IOPS for a DB instance.
     */
    readonly minIopsPerDbInstance: number;
    /**
     * Minimum provisioned IOPS per GiB for a DB instance.
     */
    readonly minIopsPerGib: number;
    /**
     * Minimum storage size for a DB instance.
     */
    readonly minStorageSize: number;
    /**
     * Whether a DB instance is Multi-AZ capable.
     */
    readonly multiAzCapable: boolean;
    /**
     * Whether a DB instance supports RDS on Outposts.
     */
    readonly outpostCapable: boolean;
    readonly preferredEngineVersions?: string[];
    readonly preferredInstanceClasses?: string[];
    /**
     * Whether a DB instance can have a read replica.
     */
    readonly readReplicaCapable: boolean;
    readonly storageType: string;
    /**
     * A list of the supported DB engine modes.
     */
    readonly supportedEngineModes: string[];
    readonly supportsEnhancedMonitoring: boolean;
    readonly supportsGlobalDatabases: boolean;
    readonly supportsIamDatabaseAuthentication: boolean;
    readonly supportsIops: boolean;
    readonly supportsKerberosAuthentication: boolean;
    readonly supportsPerformanceInsights: boolean;
    readonly supportsStorageAutoscaling: boolean;
    readonly supportsStorageEncryption: boolean;
    readonly vpc: boolean;
}

export function getOrderableDbInstanceApply(args: GetOrderableDbInstanceApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrderableDbInstanceResult> {
    return pulumi.output(args).apply(a => getOrderableDbInstance(a, opts))
}

/**
 * A collection of arguments for invoking getOrderableDbInstance.
 */
export interface GetOrderableDbInstanceApplyArgs {
    /**
     * Availability zone group.
     */
    availabilityZoneGroup?: pulumi.Input<string>;
    /**
     * DB engine. Engine values include `aurora`, `aurora-mysql`, `aurora-postgresql`, `docdb`, `mariadb`, `mysql`, `neptune`, `oracle-ee`, `oracle-se`, `oracle-se1`, `oracle-se2`, `postgres`, `sqlserver-ee`, `sqlserver-ex`, `sqlserver-se`, and `sqlserver-web`.
     */
    engine: pulumi.Input<string>;
    /**
     * Version of the DB engine. If none is provided, the AWS-defined default version will be used.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * DB instance class. Examples of classes are `db.m3.2xlarge`, `db.t2.small`, and `db.m3.medium`.
     */
    instanceClass?: pulumi.Input<string>;
    /**
     * License model. Examples of license models are `general-public-license`, `bring-your-own-license`, and `amazon-license`.
     */
    licenseModel?: pulumi.Input<string>;
    /**
     * Ordered list of preferred RDS DB instance engine versions. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
     */
    preferredEngineVersions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Ordered list of preferred RDS DB instance classes. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
     */
    preferredInstanceClasses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Storage types. Examples of storage types are `standard`, `io1`, `gp2`, and `aurora`.
     */
    storageType?: pulumi.Input<string>;
    /**
     * Enable this to ensure a DB instance supports Enhanced Monitoring at intervals from 1 to 60 seconds.
     */
    supportsEnhancedMonitoring?: pulumi.Input<boolean>;
    /**
     * Enable this to ensure a DB instance supports Aurora global databases with a specific combination of other DB engine attributes.
     */
    supportsGlobalDatabases?: pulumi.Input<boolean>;
    /**
     * Enable this to ensure a DB instance supports IAM database authentication.
     */
    supportsIamDatabaseAuthentication?: pulumi.Input<boolean>;
    /**
     * Enable this to ensure a DB instance supports provisioned IOPS.
     */
    supportsIops?: pulumi.Input<boolean>;
    /**
     * Enable this to ensure a DB instance supports Kerberos Authentication.
     */
    supportsKerberosAuthentication?: pulumi.Input<boolean>;
    /**
     * Enable this to ensure a DB instance supports Performance Insights.
     */
    supportsPerformanceInsights?: pulumi.Input<boolean>;
    /**
     * Enable this to ensure Amazon RDS can automatically scale storage for DB instances that use the specified DB instance class.
     */
    supportsStorageAutoscaling?: pulumi.Input<boolean>;
    /**
     * Enable this to ensure a DB instance supports encrypted storage.
     */
    supportsStorageEncryption?: pulumi.Input<boolean>;
    /**
     * Boolean that indicates whether to show only VPC or non-VPC offerings.
     */
    vpc?: pulumi.Input<boolean>;
}
