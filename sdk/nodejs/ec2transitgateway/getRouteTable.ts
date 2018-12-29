// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get information on an EC2 Transit Gateway Route Table.
 * 
 * ## Example Usage
 * 
 * ### By Filter
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_ec2_transit_gateway_route_table_example = pulumi.output(aws.ec2transitgateway.getRouteTable({
 *     filters: [
 *         {
 *             name: "default-association-route-table",
 *             values: ["true"],
 *         },
 *         {
 *             name: "transit-gateway-id",
 *             values: ["tgw-12345678"],
 *         },
 *     ],
 * }));
 * ```
 * ### By Identifier
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_ec2_transit_gateway_route_table_example = pulumi.output(aws.ec2transitgateway.getRouteTable({
 *     id: "tgw-rtb-12345678",
 * }));
 * ```
 */
export function getRouteTable(args?: GetRouteTableArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteTableResult> {
    args = args || {};
    return pulumi.runtime.invoke("aws:ec2transitgateway/getRouteTable:getRouteTable", {
        "filters": args.filters,
        "id": args.id,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteTable.
 */
export interface GetRouteTableArgs {
    /**
     * One or more configuration blocks containing name-values filters. Detailed below.
     */
    readonly filters?: { name: string, values: string[] }[];
    /**
     * Identifier of the EC2 Transit Gateway Route Table.
     */
    readonly id?: string;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getRouteTable.
 */
export interface GetRouteTableResult {
    /**
     * Boolean whether this is the default association route table for the EC2 Transit Gateway
     */
    readonly defaultAssociationRouteTable: boolean;
    /**
     * Boolean whether this is the default propagation route table for the EC2 Transit Gateway
     */
    readonly defaultPropagationRouteTable: boolean;
    /**
     * Key-value tags for the EC2 Transit Gateway Route Table
     */
    readonly tags: {[key: string]: any};
    /**
     * EC2 Transit Gateway identifier
     */
    readonly transitGatewayId: string;
}
