// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides information for multiple EC2 Transit Gateway Route Tables, such as their identifiers.
 *
 * ## Example Usage
 *
 * The following shows outputing all Transit Gateway Route Table Ids.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleTransitGatewayRouteTables = aws.ec2.getTransitGatewayRouteTables({});
 * export const example = data.aws_ec2_transit_gateway_route_table.example.ids;
 * ```
 */
export function getTransitGatewayRouteTables(args?: GetTransitGatewayRouteTablesArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitGatewayRouteTablesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getTransitGatewayRouteTables:getTransitGatewayRouteTables", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getTransitGatewayRouteTables.
 */
export interface GetTransitGatewayRouteTablesArgs {
    /**
     * Custom filter block as described below.
     */
    filters?: inputs.ec2.GetTransitGatewayRouteTablesFilter[];
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired transit gateway route table.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getTransitGatewayRouteTables.
 */
export interface GetTransitGatewayRouteTablesResult {
    readonly filters?: outputs.ec2.GetTransitGatewayRouteTablesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Set of Transit Gateway Route Table identifiers.
     */
    readonly ids: string[];
    readonly tags: {[key: string]: string};
}

export function getTransitGatewayRouteTablesApply(args?: GetTransitGatewayRouteTablesApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransitGatewayRouteTablesResult> {
    return pulumi.output(args).apply(a => getTransitGatewayRouteTables(a, opts))
}

/**
 * A collection of arguments for invoking getTransitGatewayRouteTables.
 */
export interface GetTransitGatewayRouteTablesApplyArgs {
    /**
     * Custom filter block as described below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.ec2.GetTransitGatewayRouteTablesFilter>[]>;
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired transit gateway route table.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
