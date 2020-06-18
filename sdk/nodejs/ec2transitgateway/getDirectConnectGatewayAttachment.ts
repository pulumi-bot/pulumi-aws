// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get information on an EC2 Transit Gateway's attachment to a Direct Connect Gateway.
 *
 * ## Example Usage
 * ### By Transit Gateway and Direct Connect Gateway Identifiers
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.all([aws_dx_gateway_example.id, aws_ec2_transit_gateway_example.id]).apply(([aws_dx_gateway_exampleId, aws_ec2_transit_gateway_exampleId]) => aws.ec2transitgateway.getDirectConnectGatewayAttachment({
 *     dxGatewayId: aws_dx_gateway_exampleId,
 *     transitGatewayId: aws_ec2_transit_gateway_exampleId,
 * }, { async: true }));
 * ```
 */
export function getDirectConnectGatewayAttachment(args?: GetDirectConnectGatewayAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetDirectConnectGatewayAttachmentResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2transitgateway/getDirectConnectGatewayAttachment:getDirectConnectGatewayAttachment", {
        "dxGatewayId": args.dxGatewayId,
        "filters": args.filters,
        "tags": args.tags,
        "transitGatewayId": args.transitGatewayId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDirectConnectGatewayAttachment.
 */
export interface GetDirectConnectGatewayAttachmentArgs {
    /**
     * Identifier of the Direct Connect Gateway.
     */
    readonly dxGatewayId?: string;
    /**
     * Configuration block(s) for filtering. Detailed below.
     */
    readonly filters?: inputs.ec2transitgateway.GetDirectConnectGatewayAttachmentFilter[];
    /**
     * A map of tags, each pair of which must exactly match a pair on the desired Transit Gateway Direct Connect Gateway Attachment.
     */
    readonly tags?: {[key: string]: any};
    /**
     * Identifier of the EC2 Transit Gateway.
     */
    readonly transitGatewayId?: string;
}

/**
 * A collection of values returned by getDirectConnectGatewayAttachment.
 */
export interface GetDirectConnectGatewayAttachmentResult {
    readonly dxGatewayId?: string;
    readonly filters?: outputs.ec2transitgateway.GetDirectConnectGatewayAttachmentFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Key-value tags for the EC2 Transit Gateway Attachment
     */
    readonly tags: {[key: string]: any};
    readonly transitGatewayId?: string;
}
