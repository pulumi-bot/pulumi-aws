// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about an Elasticache Cluster
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const myCluster = aws.elasticache.getCluster({
 *     clusterId: "my-cluster-id",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elasticache_cluster.html.markdown.
 */
export function getCluster(args: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> & GetClusterResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetClusterResult> = pulumi.runtime.invoke("aws:elasticache/getCluster:getCluster", {
        "clusterId": args.clusterId,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    /**
     * Group identifier.
     */
    readonly clusterId: string;
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    readonly arn: string;
    /**
     * The Availability Zone for the cache cluster.
     */
    readonly availabilityZone: string;
    /**
     * List of node objects including `id`, `address`, `port` and `availabilityZone`.
     * Referenceable e.g. as `${data.aws_elasticache_cluster.bar.cache_nodes.0.address}`
     */
    readonly cacheNodes: outputs.elasticache.GetClusterCacheNode[];
    /**
     * (Memcached only) The DNS name of the cache cluster without the port appended.
     */
    readonly clusterAddress: string;
    readonly clusterId: string;
    /**
     * (Memcached only) The configuration endpoint to allow host discovery.
     */
    readonly configurationEndpoint: string;
    /**
     * Name of the cache engine.
     */
    readonly engine: string;
    /**
     * Version number of the cache engine.
     */
    readonly engineVersion: string;
    /**
     * Specifies the weekly time range for when maintenance
     * on the cache cluster is performed.
     */
    readonly maintenanceWindow: string;
    /**
     * The cluster node type.
     */
    readonly nodeType: string;
    /**
     * An Amazon Resource Name (ARN) of an
     * SNS topic that ElastiCache notifications get sent to.
     */
    readonly notificationTopicArn: string;
    /**
     * The number of cache nodes that the cache cluster has.
     */
    readonly numCacheNodes: number;
    /**
     * Name of the parameter group associated with this cache cluster.
     */
    readonly parameterGroupName: string;
    /**
     * The port number on which each of the cache nodes will
     * accept connections.
     */
    readonly port: number;
    /**
     * The replication group to which this cache cluster belongs.
     */
    readonly replicationGroupId: string;
    /**
     * List VPC security groups associated with the cache cluster.
     */
    readonly securityGroupIds: string[];
    /**
     * List of security group names associated with this cache cluster.
     */
    readonly securityGroupNames: string[];
    /**
     * The number of days for which ElastiCache will
     * retain automatic cache cluster snapshots before deleting them.
     */
    readonly snapshotRetentionLimit: number;
    /**
     * The daily time range (in UTC) during which ElastiCache will
     * begin taking a daily snapshot of the cache cluster.
     */
    readonly snapshotWindow: string;
    /**
     * Name of the subnet group associated to the cache cluster.
     */
    readonly subnetGroupName: string;
    /**
     * The tags assigned to the resource
     */
    readonly tags: {[key: string]: string};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
