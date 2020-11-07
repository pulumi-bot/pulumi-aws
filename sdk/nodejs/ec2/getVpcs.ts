// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * This resource can be useful for getting back a list of VPC Ids for a region.
 *
 * The following example retrieves a list of VPC Ids with a custom tag of `service` set to a value of "production".
 */
export function getVpcs(args?: GetVpcsArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getVpcs:getVpcs", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcs.
 */
export interface GetVpcsArgs {
    /**
     * Custom filter block as described below.
     */
    readonly filters?: inputs.ec2.GetVpcsFilter[];
    /**
     * A map of tags, each pair of which must exactly match
     * a pair on the desired vpcs.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getVpcs.
 */
export interface GetVpcsResult {
    readonly filters?: outputs.ec2.GetVpcsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of all the VPC Ids found. This data source will fail if none are found.
     */
    readonly ids: string[];
    readonly tags: {[key: string]: string};
}
