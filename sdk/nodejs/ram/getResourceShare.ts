// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `aws_ram_resource_share` Retrieve information about a RAM Resource Share.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = pulumi.output(aws.ram.getResourceShare({
 *     name: "example",
 * }));
 * ```
 * 
 * ## Search by filters
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const tagFilter = pulumi.output(aws.ram.getResourceShare({
 *     filters: [{
 *         name: "NameOfTag",
 *         values: ["exampleNameTagValue"],
 *     }],
 *     name: "MyResourceName",
 *     resourceOwner: "SELF",
 * }));
 * ```
 */
export function getResourceShare(args: GetResourceShareArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceShareResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ram/getResourceShare:getResourceShare", {
        "filters": args.filters,
        "name": args.name,
        "resourceOwner": args.resourceOwner,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourceShare.
 */
export interface GetResourceShareArgs {
    /**
     * A filter used to scope the list e.g. by tags. See [related docs] (https://docs.aws.amazon.com/ram/latest/APIReference/API_TagFilter.html).
     */
    readonly filters?: { name: string, values: string[] }[];
    /**
     * The name of the tag key to filter on.
     */
    readonly name: string;
    /**
     * The owner of the resource share. Valid values are SELF or OTHER-ACCOUNTS
     */
    readonly resourceOwner: string;
}

/**
 * A collection of values returned by getResourceShare.
 */
export interface GetResourceShareResult {
    /**
     * The Amazon Resource Name (ARN) of the resource share.
     */
    readonly arn: string;
    readonly filters?: { name: string, values: string[] }[];
    /**
     * The Amazon Resource Name (ARN) of the resource share.
     */
    readonly id: string;
    readonly name: string;
    readonly resourceOwner: string;
    /**
     * The Status of the RAM share.
     */
    readonly status: string;
    /**
     * The Tags attached to the RAM share
     */
    readonly tags: {[key: string]: any};
}
