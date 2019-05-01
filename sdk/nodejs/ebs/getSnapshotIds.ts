// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of EBS Snapshot IDs matching the specified
 * criteria.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ebsVolumes = pulumi.output(aws.ebs.getSnapshotIds({
 *     filters: [
 *         {
 *             name: "volume-size",
 *             values: ["40"],
 *         },
 *         {
 *             name: "tag:Name",
 *             values: ["Example"],
 *         },
 *     ],
 *     owners: ["self"],
 * }));
 * ```
 */
export function getSnapshotIds(args?: GetSnapshotIdsArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotIdsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ebs/getSnapshotIds:getSnapshotIds", {
        "filters": args.filters,
        "owners": args.owners,
        "restorableByUserIds": args.restorableByUserIds,
    }, opts);
}

/**
 * A collection of arguments for invoking getSnapshotIds.
 */
export interface GetSnapshotIdsArgs {
    /**
     * One or more name/value pairs to filter off of. There are
     * several valid keys, for a full reference, check out
     * [describe-volumes in the AWS CLI reference][1].
     */
    readonly filters?: { name: string, values: string[] }[];
    /**
     * Returns the snapshots owned by the specified owner id. Multiple owners can be specified.
     */
    readonly owners?: string[];
    /**
     * One or more AWS accounts IDs that can create volumes from the snapshot.
     */
    readonly restorableByUserIds?: string[];
}

/**
 * A collection of values returned by getSnapshotIds.
 */
export interface GetSnapshotIdsResult {
    readonly filters?: { name: string, values: string[] }[];
    readonly ids: string[];
    readonly owners?: string[];
    readonly restorableByUserIds?: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
