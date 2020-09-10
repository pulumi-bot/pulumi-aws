// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getIpSet(args: GetIpSetArgs, opts?: pulumi.InvokeOptions): Promise<GetIpSetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:wafv2/getIpSet:getIpSet", {
        "name": args.name,
        "scope": args.scope,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpSet.
 */
export interface GetIpSetArgs {
    readonly name: string;
    readonly scope: string;
}

/**
 * A collection of values returned by getIpSet.
 */
export interface GetIpSetResult {
    readonly addresses: string[];
    readonly arn: string;
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipAddressVersion: string;
    readonly name: string;
    readonly scope: string;
}
