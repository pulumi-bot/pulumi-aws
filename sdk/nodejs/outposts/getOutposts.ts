// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides details about multiple Outposts.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.outposts.getOutposts({
 *     siteId: data.aws_outposts_site.id,
 * });
 * ```
 */
export function getOutposts(args?: GetOutpostsArgs, opts?: pulumi.InvokeOptions): Promise<GetOutpostsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:outposts/getOutposts:getOutposts", {
        "availabilityZone": args.availabilityZone,
        "availabilityZoneId": args.availabilityZoneId,
        "siteId": args.siteId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOutposts.
 */
export interface GetOutpostsArgs {
    /**
     * Availability Zone name.
     */
    readonly availabilityZone?: string;
    /**
     * Availability Zone identifier.
     */
    readonly availabilityZoneId?: string;
    /**
     * Site identifier.
     */
    readonly siteId?: string;
}

/**
 * A collection of values returned by getOutposts.
 */
export interface GetOutpostsResult {
    /**
     * Set of Amazon Resource Names (ARNs).
     */
    readonly arns: string[];
    readonly availabilityZone: string;
    readonly availabilityZoneId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Set of identifiers.
     */
    readonly ids: string[];
    readonly siteId: string;
}
