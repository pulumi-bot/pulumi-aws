// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getInternetGateway(args?: GetInternetGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetInternetGatewayResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getInternetGateway:getInternetGateway", {
        "filters": args.filters,
        "internetGatewayId": args.internetGatewayId,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getInternetGateway.
 */
export interface GetInternetGatewayArgs {
    readonly filters?: inputs.ec2.GetInternetGatewayFilter[];
    readonly internetGatewayId?: string;
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getInternetGateway.
 */
export interface GetInternetGatewayResult {
    readonly arn: string;
    readonly attachments: outputs.ec2.GetInternetGatewayAttachment[];
    readonly filters?: outputs.ec2.GetInternetGatewayFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly internetGatewayId: string;
    readonly ownerId: string;
    readonly tags: {[key: string]: string};
}
