// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Provides information about an RDS cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const clusterName = pulumi.output(aws.rds.getCluster({
 *     clusterIdentifier: "clusterName",
 * }, { async: true }));
 * ```
 */
export function getCluster(args: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:rds/getCluster:getCluster", {
        "clusterIdentifier": args.clusterIdentifier,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    /**
     * The cluster identifier of the RDS cluster.
     */
    readonly clusterIdentifier: string;
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    readonly arn: string;
    readonly availabilityZones: string[];
    readonly backtrackWindow: number;
    readonly backupRetentionPeriod: number;
    readonly clusterIdentifier: string;
    readonly clusterMembers: string[];
    readonly clusterResourceId: string;
    readonly databaseName: string;
    readonly dbClusterParameterGroupName: string;
    readonly dbSubnetGroupName: string;
    readonly enabledCloudwatchLogsExports: string[];
    readonly endpoint: string;
    readonly engine: string;
    readonly engineVersion: string;
    readonly finalSnapshotIdentifier: string;
    readonly hostedZoneId: string;
    readonly iamDatabaseAuthenticationEnabled: boolean;
    readonly iamRoles: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly kmsKeyId: string;
    readonly masterUsername: string;
    readonly port: number;
    readonly preferredBackupWindow: string;
    readonly preferredMaintenanceWindow: string;
    readonly readerEndpoint: string;
    readonly replicationSourceIdentifier: string;
    readonly storageEncrypted: boolean;
    readonly tags: {[key: string]: string};
    readonly vpcSecurityGroupIds: string[];
}
