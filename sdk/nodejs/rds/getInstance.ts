// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about an RDS instance
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const database = pulumi.output(aws.rds.getInstance({
 *     dbInstanceIdentifier: "my-test-database",
 * }, { async: true }));
 * ```
 */
export function getInstance(args: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:rds/getInstance:getInstance", {
        "dbInstanceIdentifier": args.dbInstanceIdentifier,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    /**
     * The name of the RDS instance
     */
    readonly dbInstanceIdentifier: string;
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    /**
     * The hostname of the RDS instance. See also `endpoint` and `port`.
     */
    readonly address: string;
    /**
     * Specifies the allocated storage size specified in gigabytes.
     */
    readonly allocatedStorage: number;
    /**
     * Indicates that minor version patches are applied automatically.
     */
    readonly autoMinorVersionUpgrade: boolean;
    /**
     * Specifies the name of the Availability Zone the DB instance is located in.
     */
    readonly availabilityZone: string;
    /**
     * Specifies the number of days for which automatic DB snapshots are retained.
     */
    readonly backupRetentionPeriod: number;
    /**
     * Specifies the identifier of the CA certificate for the DB instance.
     */
    readonly caCertIdentifier: string;
    /**
     * If the DB instance is a member of a DB cluster, contains the name of the DB cluster that the DB instance is a member of.
     */
    readonly dbClusterIdentifier: string;
    /**
     * The Amazon Resource Name (ARN) for the DB instance.
     */
    readonly dbInstanceArn: string;
    /**
     * Contains the name of the compute and memory capacity class of the DB instance.
     */
    readonly dbInstanceClass: string;
    readonly dbInstanceIdentifier: string;
    /**
     * Specifies the port that the DB instance listens on.
     */
    readonly dbInstancePort: number;
    /**
     * Contains the name of the initial database of this instance that was provided at create time, if one was specified when the DB instance was created. This same name is returned for the life of the DB instance.
     */
    readonly dbName: string;
    /**
     * Provides the list of DB parameter groups applied to this DB instance.
     */
    readonly dbParameterGroups: string[];
    /**
     * Provides List of DB security groups associated to this DB instance.
     */
    readonly dbSecurityGroups: string[];
    /**
     * Specifies the name of the subnet group associated with the DB instance.
     */
    readonly dbSubnetGroup: string;
    /**
     * List of log types to export to cloudwatch.
     */
    readonly enabledCloudwatchLogsExports: string[];
    /**
     * The connection endpoint in `address:port` format.
     */
    readonly endpoint: string;
    /**
     * Provides the name of the database engine to be used for this DB instance.
     */
    readonly engine: string;
    /**
     * Indicates the database engine version.
     */
    readonly engineVersion: string;
    /**
     * The canonical hosted zone ID of the DB instance (to be used in a Route 53 Alias record).
     */
    readonly hostedZoneId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Specifies the Provisioned IOPS (I/O operations per second) value.
     */
    readonly iops: number;
    /**
     * If StorageEncrypted is true, the KMS key identifier for the encrypted DB instance.
     */
    readonly kmsKeyId: string;
    /**
     * License model information for this DB instance.
     */
    readonly licenseModel: string;
    /**
     * Contains the master username for the DB instance.
     */
    readonly masterUsername: string;
    /**
     * The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.
     */
    readonly monitoringInterval: number;
    /**
     * The ARN for the IAM role that permits RDS to send Enhanced Monitoring metrics to CloudWatch Logs.
     */
    readonly monitoringRoleArn: string;
    /**
     * Specifies if the DB instance is a Multi-AZ deployment.
     */
    readonly multiAz: boolean;
    /**
     * Provides the list of option group memberships for this DB instance.
     */
    readonly optionGroupMemberships: string[];
    /**
     * The database port.
     */
    readonly port: number;
    /**
     * Specifies the daily time range during which automated backups are created.
     */
    readonly preferredBackupWindow: string;
    /**
     * Specifies the weekly time range during which system maintenance can occur in UTC.
     */
    readonly preferredMaintenanceWindow: string;
    /**
     * Specifies the accessibility options for the DB instance.
     */
    readonly publiclyAccessible: boolean;
    /**
     * The identifier of the source DB that this is a replica of.
     */
    readonly replicateSourceDb: string;
    /**
     * The RDS Resource ID of this instance.
     */
    readonly resourceId: string;
    /**
     * Specifies whether the DB instance is encrypted.
     */
    readonly storageEncrypted: boolean;
    /**
     * Specifies the storage type associated with DB instance.
     */
    readonly storageType: string;
    readonly tags: {[key: string]: string};
    /**
     * The time zone of the DB instance.
     */
    readonly timezone: string;
    /**
     * Provides a list of VPC security group elements that the DB instance belongs to.
     */
    readonly vpcSecurityGroups: string[];
}
