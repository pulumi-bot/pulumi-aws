// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides details about an EC2 Local Gateway.
 */
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
    /**
     * Custom filter block as described below.
     */
    readonly filters?: inputs.ec2.GetLocalGatewayFilter[];
    /**
     * The id of the specific Local Gateway to retrieve.
     */
    readonly id?: string;
    /**
     * The current state of the desired Local Gateway.
     * Can be either `"pending"` or `"available"`.
     */
    readonly state?: string;
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired Local Gateway.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getLocalGateway.
 */
export interface GetLocalGatewayResult {
    readonly filters?: outputs.ec2.GetLocalGatewayFilter[];
    readonly id: string;
    /**
     * Amazon Resource Name (ARN) of Outpost
     */
    readonly outpostArn: string;
    /**
     * AWS account identifier that owns the Local Gateway.
     */
    readonly ownerId: string;
    /**
     * State of the local gateway.
     */
    readonly state: string;
    readonly tags: {[key: string]: string};
}
