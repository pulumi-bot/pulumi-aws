// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Get information on an EC2 Transit Gateway VPN Attachment.
 *
 * ## Example Usage
 * ### By Transit Gateway and VPN Connection Identifiers
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.ec2transitgateway.getVpnAttachment({
 *     transitGatewayId: aws_ec2_transit_gateway.example.id,
 *     vpnConnectionId: aws_vpn_connection.example.id,
 * });
 * ```
 * ### Filter
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.ec2transitgateway.getVpnAttachment({
 *     filters: [{
 *         name: "resource-id",
 *         values: ["some-resource"],
 *     }],
 * }, { async: true }));
 * ```
 */
export function getVpnAttachment(args?: GetVpnAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetVpnAttachmentResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2transitgateway/getVpnAttachment:getVpnAttachment", {
        "filters": args.filters,
        "tags": args.tags,
        "transitGatewayId": args.transitGatewayId,
        "vpnConnectionId": args.vpnConnectionId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpnAttachment.
 */
export interface GetVpnAttachmentArgs {
    /**
     * Configuration block(s) for filtering. Detailed below.
     */
    readonly filters?: inputs.ec2transitgateway.GetVpnAttachmentFilter[];
    /**
     * A map of tags, each pair of which must exactly match a pair on the desired Transit Gateway VPN Attachment.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Identifier of the EC2 Transit Gateway.
     */
    readonly transitGatewayId?: string;
    /**
     * Identifier of the EC2 VPN Connection.
     */
    readonly vpnConnectionId?: string;
}

/**
 * A collection of values returned by getVpnAttachment.
 */
export interface GetVpnAttachmentResult {
    readonly filters?: outputs.ec2transitgateway.GetVpnAttachmentFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Key-value tags for the EC2 Transit Gateway VPN Attachment
     */
    readonly tags: {[key: string]: string};
    readonly transitGatewayId?: string;
    readonly vpnConnectionId?: string;
}
