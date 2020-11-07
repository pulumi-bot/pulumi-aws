// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The VPN Gateway data source provides details about
 * a specific VPN gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const selected = aws.ec2.getVpnGateway({
 *     filters: [{
 *         name: "tag:Name",
 *         values: ["vpn-gw"],
 *     }],
 * });
 * export const vpnGatewayId = selected.then(selected => selected.id);
 * ```
 */
export function getVpnGateway(args?: GetVpnGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetVpnGatewayResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getVpnGateway:getVpnGateway", {
        "amazonSideAsn": args.amazonSideAsn,
        "attachedVpcId": args.attachedVpcId,
        "availabilityZone": args.availabilityZone,
        "filters": args.filters,
        "id": args.id,
        "state": args.state,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpnGateway.
 */
export interface GetVpnGatewayArgs {
    /**
     * The Autonomous System Number (ASN) for the Amazon side of the specific VPN Gateway to retrieve.
     */
    readonly amazonSideAsn?: string;
    /**
     * The ID of a VPC attached to the specific VPN Gateway to retrieve.
     */
    readonly attachedVpcId?: string;
    /**
     * The Availability Zone of the specific VPN Gateway to retrieve.
     */
    readonly availabilityZone?: string;
    /**
     * Custom filter block as described below.
     */
    readonly filters?: inputs.ec2.GetVpnGatewayFilter[];
    /**
     * The ID of the specific VPN Gateway to retrieve.
     */
    readonly id?: string;
    /**
     * The state of the specific VPN Gateway to retrieve.
     */
    readonly state?: string;
    /**
     * A map of tags, each pair of which must exactly match
     * a pair on the desired VPN Gateway.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getVpnGateway.
 */
export interface GetVpnGatewayResult {
    readonly amazonSideAsn: string;
    readonly arn: string;
    readonly attachedVpcId: string;
    readonly availabilityZone: string;
    readonly filters?: outputs.ec2.GetVpnGatewayFilter[];
    readonly id: string;
    readonly state: string;
    readonly tags: {[key: string]: string};
}
