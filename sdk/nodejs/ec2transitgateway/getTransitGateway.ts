// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getTransitGateway(args?: GetTransitGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitGatewayResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2transitgateway/getTransitGateway:getTransitGateway", {
        "filters": args.filters,
        "id": args.id,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getTransitGateway.
 */
export interface GetTransitGatewayArgs {
    readonly filters?: inputs.ec2transitgateway.GetTransitGatewayFilter[];
    readonly id?: string;
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getTransitGateway.
 */
export interface GetTransitGatewayResult {
    readonly amazonSideAsn: number;
    readonly arn: string;
    readonly associationDefaultRouteTableId: string;
    readonly autoAcceptSharedAttachments: string;
    readonly defaultRouteTableAssociation: string;
    readonly defaultRouteTablePropagation: string;
    readonly description: string;
    readonly dnsSupport: string;
    readonly filters?: outputs.ec2transitgateway.GetTransitGatewayFilter[];
    readonly id?: string;
    readonly ownerId: string;
    readonly propagationDefaultRouteTableId: string;
    readonly tags: {[key: string]: string};
    readonly vpnEcmpSupport: string;
}
