// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getLocalGateway(args?: GetLocalGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalGatewayResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getLocalGateway:getLocalGateway", {
        "filters": args.filters,
        "id": args.id,
        "state": args.state,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocalGateway.
 */
export interface GetLocalGatewayArgs {
    readonly filters?: inputs.ec2.GetLocalGatewayFilter[];
    readonly id?: string;
    readonly state?: string;
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getLocalGateway.
 */
export interface GetLocalGatewayResult {
    readonly filters?: outputs.ec2.GetLocalGatewayFilter[];
    readonly id: string;
    readonly outpostArn: string;
    readonly ownerId: string;
    readonly state: string;
    readonly tags: {[key: string]: string};
}
