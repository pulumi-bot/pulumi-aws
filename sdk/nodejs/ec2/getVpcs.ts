// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource can be useful for getting back a list of VPC Ids for a region.
 * 
 * The following example retrieves a list of VPC Ids with a custom tag of `service` set to a value of "production".
 * 
 * ## Example Usage
 * 
 * The following shows outputing all VPC Ids.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const fooVpcs = aws.ec2.getVpcs({
 *     tags: {
 *         service: "production",
 *     },
 * });
 * 
 * export const foo = fooVpcs.ids;
 * ```
 * 
 * An example use case would be interpolate the `aws.ec2.getVpcs` output into `count` of an aws.ec2.FlowLog resource.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const fooVpcs = aws.ec2.getVpcs({});
 * const testFlowLog: aws.ec2.FlowLog[] = [];
 * for (let i = 0; i < fooVpcs.ids.length; i++) {
 *     testFlowLog.push(new aws.ec2.FlowLog(`test_flow_log-${i}`, {
 *         // ...
 *         vpcId: fooVpcs.ids[i],
 *     }));
 * }
 * 
 * export const foo = fooVpcs.ids;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/vpcs.html.markdown.
 */
export function getVpcs(args?: GetVpcsArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcsResult> & GetVpcsResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetVpcsResult> = pulumi.runtime.invoke("aws:ec2/getVpcs:getVpcs", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getVpcs.
 */
export interface GetVpcsArgs {
    /**
     * Custom filter block as described below.
     */
    readonly filters?: { name: string, values: string[] }[];
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired vpcs.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getVpcs.
 */
export interface GetVpcsResult {
    readonly filters?: { name: string, values: string[] }[];
    /**
     * A list of all the VPC Ids found. This data source will fail if none are found.
     */
    readonly ids: string[];
    readonly tags: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
