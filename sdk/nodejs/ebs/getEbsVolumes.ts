// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * `aws.ebs.getEbsVolumes` provides identifying information for EBS volumes matching given criteria.
 *
 * This data source can be useful for getting a list of volume IDs with (for example) matching tags.
 */
export function getEbsVolumes(args?: GetEbsVolumesArgs, opts?: pulumi.InvokeOptions): Promise<GetEbsVolumesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ebs/getEbsVolumes:getEbsVolumes", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getEbsVolumes.
 */
export interface GetEbsVolumesArgs {
    /**
     * Custom filter block as described below.
     */
    readonly filters?: inputs.ebs.GetEbsVolumesFilter[];
    /**
     * A map of tags, each pair of which must exactly match
     * a pair on the desired volumes.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getEbsVolumes.
 */
export interface GetEbsVolumesResult {
    readonly filters?: outputs.ebs.GetEbsVolumesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A set of all the EBS Volume IDs found. This data source will fail if
     * no volumes match the provided criteria.
     */
    readonly ids: string[];
    readonly tags?: {[key: string]: string};
}
