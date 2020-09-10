// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getClusterSnapshot(args?: GetClusterSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterSnapshotResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:rds/getClusterSnapshot:getClusterSnapshot", {
        "dbClusterIdentifier": args.dbClusterIdentifier,
        "dbClusterSnapshotIdentifier": args.dbClusterSnapshotIdentifier,
        "includePublic": args.includePublic,
        "includeShared": args.includeShared,
        "mostRecent": args.mostRecent,
        "snapshotType": args.snapshotType,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusterSnapshot.
 */
export interface GetClusterSnapshotArgs {
    readonly dbClusterIdentifier?: string;
    readonly dbClusterSnapshotIdentifier?: string;
    readonly includePublic?: boolean;
    readonly includeShared?: boolean;
    readonly mostRecent?: boolean;
    readonly snapshotType?: string;
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getClusterSnapshot.
 */
export interface GetClusterSnapshotResult {
    readonly allocatedStorage: number;
    readonly availabilityZones: string[];
    readonly dbClusterIdentifier?: string;
    readonly dbClusterSnapshotArn: string;
    readonly dbClusterSnapshotIdentifier?: string;
    readonly engine: string;
    readonly engineVersion: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includePublic?: boolean;
    readonly includeShared?: boolean;
    readonly kmsKeyId: string;
    readonly licenseModel: string;
    readonly mostRecent?: boolean;
    readonly port: number;
    readonly snapshotCreateTime: string;
    readonly snapshotType?: string;
    readonly sourceDbClusterSnapshotArn: string;
    readonly status: string;
    readonly storageEncrypted: boolean;
    readonly tags: {[key: string]: string};
    readonly vpcId: string;
}
