// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * `aws.ram.ResourceShare` Retrieve information about a RAM Resource Share.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.ram.getResourceShare({
 *     name: "example",
 *     resourceOwner: "SELF",
 * }, { async: true }));
 * ```
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
 * }, { async: true }));
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
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourceShare.
 */
export interface GetResourceShareArgs {
    /**
     * A filter used to scope the list e.g. by tags. See [related docs] (https://docs.aws.amazon.com/ram/latest/APIReference/API_TagFilter.html).
     */
    readonly filters?: inputs.ram.GetResourceShareFilter[];
    /**
     * The name of the tag key to filter on.
     */
    readonly name: string;
    /**
     * The owner of the resource share. Valid values are SELF or OTHER-ACCOUNTS
     */
    readonly resourceOwner: string;
    /**
     * The Tags attached to the RAM share
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getResourceShare.
 */
export interface GetResourceShareResult {
    /**
     * The Amazon Resource Name (ARN) of the resource share.
     */
    readonly arn: string;
    readonly filters?: outputs.ram.GetResourceShareFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The ID of the AWS account that owns the resource share.
     */
    readonly owningAccountId: string;
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
