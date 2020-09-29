// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Provides information for multiple EC2 Local Gateway Route Tables, such as their identifiers.
 */
export function getLocalGatewayRouteTables(args?: GetLocalGatewayRouteTablesArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalGatewayRouteTablesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getLocalGatewayRouteTables:getLocalGatewayRouteTables", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocalGatewayRouteTables.
 */
export interface GetLocalGatewayRouteTablesArgs {
    /**
     * Custom filter block as described below.
     */
    readonly filters?: inputs.ec2.GetLocalGatewayRouteTablesFilter[];
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired local gateway route table.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getLocalGatewayRouteTables.
 */
export interface GetLocalGatewayRouteTablesResult {
    readonly filters?: outputs.ec2.GetLocalGatewayRouteTablesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Set of Local Gateway Route Table identifiers
     */
    readonly ids: string[];
    readonly tags: {[key: string]: string};
}
