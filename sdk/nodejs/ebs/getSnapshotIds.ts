// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of EBS Snapshot IDs matching the specified
 * criteria.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ebsVolumes = aws.ebs.getSnapshotIds({
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
 * });
 * ```
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ebs_snapshot_ids.html.markdown.
 */
export function getSnapshotIds(args?: GetSnapshotIdsArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotIdsResult> & GetSnapshotIdsResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetSnapshotIdsResult> = pulumi.runtime.invoke("aws:ebs/getSnapshotIds:getSnapshotIds", {
        "filters": args.filters,
        "owners": args.owners,
        "restorableByUserIds": args.restorableByUserIds,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
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
    readonly filters?: inputs.ebs.GetSnapshotIdsFilter[];
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
    readonly filters?: outputs.ebs.GetSnapshotIdsFilter[];
    readonly ids: string[];
    readonly owners?: string[];
    readonly restorableByUserIds?: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
