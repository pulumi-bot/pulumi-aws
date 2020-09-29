// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Provides information for multiple EC2 Local Gateways, such as their identifiers.
 *
 * ## Example Usage
 *
 * The following example retrieves Local Gateways with a resource tag of `service` set to `production`.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const fooLocalGateways = aws.ec2.getLocalGateways({
 *     tags: {
 *         service: "production",
 *     },
 * });
 * export const foo = fooLocalGateways.then(fooLocalGateways => fooLocalGateways.ids);
 * ```
 */
export function getLocalGateways(args?: GetLocalGatewaysArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalGatewaysResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getLocalGateways:getLocalGateways", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocalGateways.
 */
export interface GetLocalGatewaysArgs {
    /**
     * Custom filter block as described below.
     */
    readonly filters?: inputs.ec2.GetLocalGatewaysFilter[];
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired local_gateways.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getLocalGateways.
 */
export interface GetLocalGatewaysResult {
    readonly filters?: outputs.ec2.GetLocalGatewaysFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Set of all the Local Gateway identifiers
     */
    readonly ids: string[];
    readonly tags: {[key: string]: string};
}
