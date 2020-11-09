// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

export function getVpcPeeringConnections(args?: GetVpcPeeringConnectionsArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcPeeringConnectionsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getVpcPeeringConnections:getVpcPeeringConnections", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcPeeringConnections.
 */
export interface GetVpcPeeringConnectionsArgs {
    /**
     * Custom filter block as described below.
     */
    readonly filters?: inputs.ec2.GetVpcPeeringConnectionsFilter[];
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired VPC Peering Connection.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getVpcPeeringConnections.
 */
export interface GetVpcPeeringConnectionsResult {
    readonly filters?: outputs.ec2.GetVpcPeeringConnectionsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The IDs of the VPC Peering Connections.
     */
    readonly ids: string[];
    readonly tags: {[key: string]: string};
}
