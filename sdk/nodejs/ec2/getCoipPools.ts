// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides information for multiple EC2 Customer-Owned IP Pools, such as their identifiers.
 *
 * ## Example Usage
 *
 * The following shows outputing all COIP Pool Ids.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const fooCoipPools = pulumi.output(aws.ec2.getCoipPools({ async: true }));
 *
 * export const foo = fooCoipPools.ids;
 * ```
 */
export function getCoipPools(args?: GetCoipPoolsArgs, opts?: pulumi.InvokeOptions): Promise<GetCoipPoolsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getCoipPools:getCoipPools", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getCoipPools.
 */
export interface GetCoipPoolsArgs {
    /**
     * Custom filter block as described below.
     */
    readonly filters?: inputs.ec2.GetCoipPoolsFilter[];
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired aws_ec2_coip_pools.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getCoipPools.
 */
export interface GetCoipPoolsResult {
    readonly filters?: outputs.ec2.GetCoipPoolsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Set of COIP Pool Identifiers
     */
    readonly poolIds: string[];
    readonly tags: {[key: string]: any};
}
