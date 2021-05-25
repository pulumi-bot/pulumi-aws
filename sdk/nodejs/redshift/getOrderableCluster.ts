// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Information about Redshift Orderable Clusters and valid parameter combinations.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.redshift.getOrderableCluster({
 *     clusterType: "multi-node",
 *     preferredNodeTypes: [
 *         "dc2.large",
 *         "ds2.xlarge",
 *     ],
 * }, { async: true }));
 * ```
 */
export function getOrderableCluster(args?: GetOrderableClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetOrderableClusterResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:redshift/getOrderableCluster:getOrderableCluster", {
        "clusterType": args.clusterType,
        "clusterVersion": args.clusterVersion,
        "nodeType": args.nodeType,
        "preferredNodeTypes": args.preferredNodeTypes,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrderableCluster.
 */
export interface GetOrderableClusterArgs {
    /**
     * Reshift Cluster type. e.g. `multi-node` or `single-node`
     */
    clusterType?: string;
    /**
     * Redshift Cluster version. e.g. `1.0`
     */
    clusterVersion?: string;
    /**
     * Redshift Cluster node type. e.g. `dc2.8xlarge`
     */
    nodeType?: string;
    /**
     * Ordered list of preferred Redshift Cluster node types. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
     */
    preferredNodeTypes?: string[];
}

/**
 * A collection of values returned by getOrderableCluster.
 */
export interface GetOrderableClusterResult {
    /**
     * List of Availability Zone names where the Redshit Cluster is available.
     */
    readonly availabilityZones: string[];
    readonly clusterType: string;
    readonly clusterVersion: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nodeType: string;
    readonly preferredNodeTypes?: string[];
}
