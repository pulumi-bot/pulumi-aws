// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Provides information about AWS Regions. Can be used to filter regions i.e. by Opt-In status or only regions enabled for current account. To get details like endpoint and description of each region the data source can be combined with the `aws.getRegion` data source.
 *
 * ## Example Usage
 *
 * Enabled AWS Regions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const current = pulumi.output(aws.getRegions({ async: true }));
 * ```
 *
 * All the regions regardless of the availability
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const current = pulumi.output(aws.getRegions({
 *     allRegions: true,
 * }, { async: true }));
 * ```
 *
 * To see regions that are filtered by `"not-opted-in"`, the `allRegions` argument needs to be set to `true` or no results will be returned.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const current = pulumi.output(aws.getRegions({
 *     allRegions: true,
 *     filters: [{
 *         name: "opt-in-status",
 *         values: ["not-opted-in"],
 *     }],
 * }, { async: true }));
 * ```
 */
export function getRegions(args?: GetRegionsArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:index/getRegions:getRegions", {
        "allRegions": args.allRegions,
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegions.
 */
export interface GetRegionsArgs {
    /**
     * If true the source will query all regions regardless of availability.
     */
    readonly allRegions?: boolean;
    /**
     * Configuration block(s) to use as filters. Detailed below.
     */
    readonly filters?: inputs.GetRegionsFilter[];
}

/**
 * A collection of values returned by getRegions.
 */
export interface GetRegionsResult {
    readonly allRegions?: boolean;
    readonly filters?: outputs.GetRegionsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Names of regions that meets the criteria.
     */
    readonly names: string[];
}
