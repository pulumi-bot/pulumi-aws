// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about an Elasticache Replication Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const bar = pulumi.output(aws.elasticache.getReplicationGroup({
 *     replicationGroupId: "example",
 * }, { async: true }));
 * ```
 */
export function getReplicationGroup(args: GetReplicationGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetReplicationGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:elasticache/getReplicationGroup:getReplicationGroup", {
        "replicationGroupId": args.replicationGroupId,
    }, opts);
}

/**
 * A collection of arguments for invoking getReplicationGroup.
 */
export interface GetReplicationGroupArgs {
    /**
     * The identifier for the replication group.
     */
    readonly replicationGroupId: string;
}

/**
 * A collection of values returned by getReplicationGroup.
 */
export interface GetReplicationGroupResult {
    /**
     * A flag that enables using an AuthToken (password) when issuing Redis commands.
     */
    readonly authTokenEnabled: boolean;
    /**
     * A flag whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails.
     */
    readonly automaticFailoverEnabled: boolean;
    /**
     * The configuration endpoint address to allow host discovery.
     */
    readonly configurationEndpointAddress: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The identifiers of all the nodes that are part of this replication group.
     */
    readonly memberClusters: string[];
    /**
     * The cluster node type.
     */
    readonly nodeType: string;
    /**
     * The number of cache clusters that the replication group has.
     */
    readonly numberCacheClusters: number;
    /**
     * The port number on which the configuration endpoint will accept connections.
     */
    readonly port: number;
    /**
     * The endpoint of the primary node in this node group (shard).
     */
    readonly primaryEndpointAddress: string;
    /**
     * The description of the replication group.
     */
    readonly replicationGroupDescription: string;
    /**
     * The identifier for the replication group.
     */
    readonly replicationGroupId: string;
    /**
     * The number of days for which ElastiCache retains automatic cache cluster snapshots before deleting them.
     */
    readonly snapshotRetentionLimit: number;
    /**
     * The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).
     */
    readonly snapshotWindow: string;
}
