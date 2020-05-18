// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get information on an EC2 Transit Gateway VPN Attachment.
 *
 * ## Example Usage
 *
 * ### By Transit Gateway and VPN Connection Identifiers
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.all([aws_ec2_transit_gateway_example.id, aws_vpn_connection_example.id]).apply(([aws_ec2_transit_gateway_exampleId, aws_vpn_connection_exampleId]) => aws.ec2transitgateway.getVpnAttachment({
 *     transitGatewayId: aws_ec2_transit_gateway_exampleId,
 *     vpnConnectionId: aws_vpn_connection_exampleId,
 * }, { async: true }));
 * ```
 *
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
    readonly tags?: {[key: string]: any};
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
     * Key-value tags for the EC2 Transit Gateway VPN Attachment
     */
    readonly tags: {[key: string]: any};
    readonly transitGatewayId?: string;
    readonly vpnConnectionId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
