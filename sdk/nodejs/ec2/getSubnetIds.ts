// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getSubnetIds(args: GetSubnetIdsArgs, opts?: pulumi.InvokeOptions): Promise<GetSubnetIdsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getSubnetIds:getSubnetIds", {
        "filters": args.filters,
        "tags": args.tags,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubnetIds.
 */
export interface GetSubnetIdsArgs {
    readonly filters?: inputs.ec2.GetSubnetIdsFilter[];
    readonly tags?: {[key: string]: string};
    readonly vpcId: string;
}

/**
 * A collection of values returned by getSubnetIds.
 */
export interface GetSubnetIdsResult {
    readonly filters?: outputs.ec2.GetSubnetIdsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly tags: {[key: string]: string};
    readonly vpcId: string;
}
