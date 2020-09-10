// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getOutpost(args?: GetOutpostArgs, opts?: pulumi.InvokeOptions): Promise<GetOutpostResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:outposts/getOutpost:getOutpost", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getOutpost.
 */
export interface GetOutpostArgs {
    readonly id?: string;
    readonly name?: string;
}

/**
 * A collection of values returned by getOutpost.
 */
export interface GetOutpostResult {
    readonly arn: string;
    readonly availabilityZone: string;
    readonly availabilityZoneId: string;
    readonly description: string;
    readonly id: string;
    readonly name: string;
    readonly ownerId: string;
    readonly siteId: string;
}
