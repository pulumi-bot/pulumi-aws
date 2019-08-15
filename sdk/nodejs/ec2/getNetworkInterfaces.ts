// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * 
 * The following shows outputing all network interface ids in a region.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleNetworkInterfaces = pulumi.output(aws.ec2.getNetworkInterfaces({}));
 * 
 * export const example = exampleNetworkInterfaces.ids;
 * ```
 * 
 * The following example retrieves a list of all network interface ids with a custom tag of `Name` set to a value of `test`.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = pulumi.output(aws.ec2.getNetworkInterfaces({
 *     tags: {
 *         Name: "test",
 *     },
 * }));
 * 
 * export const example1 = example.ids;
 * ```
 * 
 * The following example retrieves a network interface ids which associated
 * with specific subnet.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleNetworkInterfaces = aws_subnet_test.id.apply(id => aws.ec2.getNetworkInterfaces({
 *     filters: [{
 *         name: "subnet-id",
 *         values: [id],
 *     }],
 * }));
 * 
 * export const example = exampleNetworkInterfaces.ids;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/network_interfaces.html.markdown.
 */
export function getNetworkInterfaces(args?: GetNetworkInterfacesArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkInterfacesResult> & GetNetworkInterfacesResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetNetworkInterfacesResult> = pulumi.runtime.invoke("aws:ec2/getNetworkInterfaces:getNetworkInterfaces", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getNetworkInterfaces.
 */
export interface GetNetworkInterfacesArgs {
    /**
     * Custom filter block as described below.
     */
    readonly filters?: inputApi.ec2.GetNetworkInterfacesFilter[];
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired network interfaces.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getNetworkInterfaces.
 */
export interface GetNetworkInterfacesResult {
    readonly filters?: outputApi.ec2.GetNetworkInterfacesFilter[];
    /**
     * A list of all the network interface ids found. This data source will fail if none are found.
     */
    readonly ids: string[];
    readonly tags: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
