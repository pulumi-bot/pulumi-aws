// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getVpcAttachment(args?: GetVpcAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcAttachmentResult> {
    args = args || {};
    return pulumi.runtime.invoke("aws:ec2transitgateway/getVpcAttachment:getVpcAttachment", {
        "filters": args.filters,
        "id": args.id,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcAttachment.
 */
export interface GetVpcAttachmentArgs {
    readonly filters?: { name: string, values: string[] }[];
    readonly id?: string;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getVpcAttachment.
 */
export interface GetVpcAttachmentResult {
    readonly dnsSupport: string;
    readonly ipv6Support: string;
    readonly subnetIds: string[];
    readonly tags: {[key: string]: any};
    readonly transitGatewayId: string;
    readonly vpcId: string;
    readonly vpcOwnerId: string;
}
