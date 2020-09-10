// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getBundle(args?: GetBundleArgs, opts?: pulumi.InvokeOptions): Promise<GetBundleResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:workspaces/getBundle:getBundle", {
        "bundleId": args.bundleId,
        "name": args.name,
        "owner": args.owner,
    }, opts);
}

/**
 * A collection of arguments for invoking getBundle.
 */
export interface GetBundleArgs {
    readonly bundleId?: string;
    readonly name?: string;
    readonly owner?: string;
}

/**
 * A collection of values returned by getBundle.
 */
export interface GetBundleResult {
    readonly bundleId?: string;
    readonly computeTypes: outputs.workspaces.GetBundleComputeType[];
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    readonly owner?: string;
    readonly rootStorages: outputs.workspaces.GetBundleRootStorage[];
    readonly userStorages: outputs.workspaces.GetBundleUserStorage[];
}
