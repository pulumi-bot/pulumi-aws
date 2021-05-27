// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
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
 * }));
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
    replicationGroupId: string;
}

/**
 * A collection of values returned by getReplicationGroup.
 */
export interface GetReplicationGroupResult {
    /**
     * The Amazon Resource Name (ARN) of the created ElastiCache Replication Group.
     */
    readonly arn: string;
    /**
     * Specifies whether an AuthToken (password) is enabled.
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
     * Specifies whether Multi-AZ Support is enabled for the replication group.
     */
    readonly multiAzEnabled: boolean;
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
     * The endpoint of the reader node in this node group (shard).
     */
    readonly readerEndpointAddress: string;
    /**
     * The description of the replication group.
     */
    readonly replicationGroupDescription: string;
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

export function getReplicationGroupApply(args: GetReplicationGroupApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReplicationGroupResult> {
    return pulumi.output(args).apply(a => getReplicationGroup(a, opts))
}

/**
 * A collection of arguments for invoking getReplicationGroup.
 */
export interface GetReplicationGroupApplyArgs {
    /**
     * The identifier for the replication group.
     */
    replicationGroupId: pulumi.Input<string>;
}
