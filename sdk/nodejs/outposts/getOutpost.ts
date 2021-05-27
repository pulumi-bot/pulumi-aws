// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides details about an Outposts Outpost.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.outposts.getOutpost({
 *     name: "example",
 * }));
 * ```
 */
export function getOutpost(args?: GetOutpostArgs, opts?: pulumi.InvokeOptions): Promise<GetOutpostResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:outposts/getOutpost:getOutpost", {
        "arn": args.arn,
        "id": args.id,
        "name": args.name,
        "ownerId": args.ownerId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOutpost.
 */
export interface GetOutpostArgs {
    /**
     * Amazon Resource Name (ARN).
     */
    arn?: string;
    /**
     * Identifier of the Outpost.
     */
    id?: string;
    /**
     * Name of the Outpost.
     */
    name?: string;
    /**
     * AWS Account identifier of the Outpost owner.
     */
    ownerId?: string;
}

/**
 * A collection of values returned by getOutpost.
 */
export interface GetOutpostResult {
    readonly arn: string;
    /**
     * Availability Zone name.
     */
    readonly availabilityZone: string;
    /**
     * Availability Zone identifier.
     */
    readonly availabilityZoneId: string;
    /**
     * Description.
     */
    readonly description: string;
    readonly id: string;
    readonly name: string;
    readonly ownerId: string;
    /**
     * Site identifier.
     */
    readonly siteId: string;
}

export function getOutpostApply(args?: GetOutpostApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOutpostResult> {
    return pulumi.output(args).apply(a => getOutpost(a, opts))
}

/**
 * A collection of arguments for invoking getOutpost.
 */
export interface GetOutpostApplyArgs {
    /**
     * Amazon Resource Name (ARN).
     */
    arn?: pulumi.Input<string>;
    /**
     * Identifier of the Outpost.
     */
    id?: pulumi.Input<string>;
    /**
     * Name of the Outpost.
     */
    name?: pulumi.Input<string>;
    /**
     * AWS Account identifier of the Outpost owner.
     */
    ownerId?: pulumi.Input<string>;
}
